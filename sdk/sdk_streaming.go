// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package sdk

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"sync/atomic"

	"github.com/ipfs/go-cid"
	"golang.org/x/exp/maps"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/akave-ai/akavesdk/private/cryptoutils"
	"github.com/akave-ai/akavesdk/private/encryption"
	"github.com/akave-ai/akavesdk/private/erasurecode"
	"github.com/akave-ai/akavesdk/private/pb"
	"github.com/akave-ai/akavesdk/private/spclient"
)

// StreamingAPI exposes SDK file streaming API.
type StreamingAPI struct {
	client            pb.StreamAPIClient
	conn              *grpc.ClientConn
	spClient          *spclient.SPClient
	uploadEC          *erasurecode.ErasureCode
	maxConcurrency    int
	blockPartSize     int64
	useConnectionPool bool
	encryptionKey     []byte // empty means no encryption
	maxBlocksInChunk  int
	chunkBuffer       int // number of chunks to buffer during streaming operations
}

// FileInfo returns meta information for single file by bucket and file name.
func (sdk *StreamingAPI) FileInfo(ctx context.Context, bucketName string, fileName string) (_ FileMeta, err error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&err)
	if bucketName == "" {
		return FileMeta{}, errSDK.Errorf("empty bucket name")
	}
	if fileName == "" {
		return FileMeta{}, errSDK.Errorf("empty file name")
	}

	res, err := sdk.client.FileView(ctx, &pb.StreamFileViewRequest{
		BucketName: bucketName,
		FileName:   fileName,
	})
	if err != nil {
		return FileMeta{}, errSDK.Wrap(err)
	}

	return FileMeta{
		StreamID:    res.StreamId,
		RootCID:     res.RootCid,
		BucketName:  res.BucketName,
		Name:        res.FileName,
		Size:        res.Size,
		EncodedSize: res.EncodedSize,
		DataBlocks:  res.DataBlocks,
		TotalBlocks: res.TotalBlocks,
		CreatedAt:   res.CreatedAt.AsTime(),
		CommitedAt:  res.CommittedAt.AsTime(),
	}, nil
}

// ListFiles returns list of files in a particular bucket.
func (sdk *StreamingAPI) ListFiles(ctx context.Context, bucketName string) (_ []FileMeta, err error) {
	defer mon.Task()(&ctx, bucketName)(&err)

	if bucketName == "" {
		return nil, errSDK.Errorf("empty bucket name")
	}

	resp, err := sdk.client.FileList(ctx, &pb.StreamFileListRequest{
		BucketName: bucketName,
	})
	if err != nil {
		return nil, errSDK.Wrap(err)
	}

	files := make([]FileMeta, 0, len(resp.Files))
	for _, fileMeta := range resp.Files {
		files = append(files, toFileMeta(fileMeta, bucketName))
	}

	return files, nil
}

// FileVersions returns list of files in a particular bucket.
func (sdk *StreamingAPI) FileVersions(ctx context.Context, bucketName, fileName string) (_ []FileMeta, err error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&err)

	if bucketName == "" {
		return nil, errSDK.Errorf("empty bucket name")
	}

	resp, err := sdk.client.FileVersions(ctx, &pb.StreamFileListVersionsRequest{
		BucketName: bucketName,
		FileName:   fileName,
	})
	if err != nil {
		return nil, errSDK.Wrap(err)
	}

	files := make([]FileMeta, 0, len(resp.Versions))
	for _, fileMeta := range resp.Versions {
		files = append(files, toFileMeta(fileMeta, bucketName))
	}

	return files, nil
}

// CreateFileUpload creates a new file upload request.
func (sdk *StreamingAPI) CreateFileUpload(ctx context.Context, bucketName string, fileName string) (_ FileUpload, err error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&err)

	if bucketName == "" {
		return FileUpload{}, errSDK.Errorf("empty bucket name")
	}

	var dataBlocks, totalBlocks int64
	if sdk.uploadEC != nil {
		dataBlocks = int64(sdk.uploadEC.DataBlocks)
		totalBlocks = int64(sdk.uploadEC.DataBlocks + sdk.uploadEC.ParityBlocks)
	}

	req := &pb.StreamFileUploadCreateRequest{
		BucketName:  bucketName,
		FileName:    fileName,
		DataBlocks:  dataBlocks,
		TotalBlocks: totalBlocks,
	}

	res, err := sdk.client.FileUploadCreate(ctx, req)
	if err != nil {
		return FileUpload{}, errSDK.Wrap(err)
	}

	return FileUpload{
		BucketName:   res.BucketName,
		Name:         res.FileName,
		StreamID:     res.StreamId,
		DataBlocks:   res.DataBlocks,
		TotalBlocks:  res.TotalBlocks,
		CreatedAt:    res.CreatedAt.AsTime(),
		dataCounters: newDataCounters(),
	}, nil
}

// Upload uploads a file using streaming api.
func (sdk *StreamingAPI) Upload(ctx context.Context, upload FileUpload, reader io.Reader) (_ FileMeta, err error) {
	defer mon.Task()(&ctx, upload)(&err)

	chunkEncOverhead := 0
	fileEncKey, err := encryptionKey(sdk.encryptionKey, upload.BucketName, upload.Name)
	if err != nil {
		return FileMeta{}, errSDK.Wrap(err)
	}
	if len(fileEncKey) > 0 {
		chunkEncOverhead = EncryptionOverhead
	}

	bufferSize := sdk.maxBlocksInChunk * int(BlockSize)
	if sdk.uploadEC != nil { // erasure coding enabled
		bufferSize = sdk.uploadEC.DataBlocks * int(BlockSize)
	}
	bufferSize -= chunkEncOverhead

	dagRoot, err := NewDAGRoot()
	if err != nil {
		return FileMeta{}, errSDK.Wrap(err)
	}

	g, chunkCtx := errgroup.WithContext(ctx)
	fileUploadChunksCh := make(chan FileChunkUpload, sdk.chunkBuffer)
	var totalSize atomic.Int64

	// Start goroutine for reading data and creating chunks
	g.Go(func() error {
		defer close(fileUploadChunksCh)

		buf := make([]byte, bufferSize)
		var index int64

		for {
			n, readErr := io.ReadFull(reader, buf)
			if readErr != nil {
				if errors.Is(readErr, io.EOF) {
					if index == 0 {
						return fmt.Errorf("empty file")
					}
					return nil
				}
				if !errors.Is(readErr, io.ErrUnexpectedEOF) {
					return readErr
				}
			}

			if n > 0 {
				chunkUpload, err := sdk.createChunkUpload(chunkCtx, upload, index, fileEncKey, buf[:n])
				if err != nil {
					return err
				}

				select {
				case <-chunkCtx.Done():
					return chunkCtx.Err()
				case fileUploadChunksCh <- chunkUpload:
					index++
					totalSize.Add(int64(n))
				}
			}

			// If this was the last chunk (ErrUnexpectedEOF), we're done
			if errors.Is(readErr, io.ErrUnexpectedEOF) {
				return nil
			}
		}
	})

	var chunkCount int64

	// Start goroutine for processing uploaded chunks
	g.Go(func() error {
		for {
			select {
			case <-chunkCtx.Done():
				return chunkCtx.Err()
			case chunkUpload, ok := <-fileUploadChunksCh:
				if !ok {
					return nil
				}

				if err := dagRoot.AddLink(chunkUpload.ChunkCID, chunkUpload.RawDataSize, chunkUpload.ProtoNodeSize); err != nil {
					return errSDK.Wrap(err)
				}

				if err := sdk.uploadChunk(chunkCtx, chunkUpload, upload.blocksCounter, upload.bytesCounter); err != nil {
					return errSDK.Wrap(err)
				}
				upload.chunksCounter.Add(1)
				chunkCount++
			}
		}
	})

	if err := g.Wait(); err != nil {
		return FileMeta{}, errSDK.Wrap(err)
	}

	rootCID, err := dagRoot.Build()
	if err != nil {
		return FileMeta{}, errSDK.Wrap(err)
	}

	// chunkCtx is already cancelled at this point, so parent ctx is used
	fileMeta, err := sdk.commitStream(ctx, upload, rootCID.String(), chunkCount, totalSize.Load())
	if err != nil {
		return FileMeta{}, errSDK.Wrap(err)
	}

	return fileMeta, nil
}

// CreateFileDownload creates a new download request.
// rootCID is optional and can be empty. Required when you want to dwonload a specific version of the file.
func (sdk *StreamingAPI) CreateFileDownload(ctx context.Context, bucketName, fileName, rootCID string) (_ FileDownload, err error) {
	defer mon.Task()(&ctx, bucketName, fileName, rootCID)(&err)

	res, err := sdk.client.FileDownloadCreate(ctx, &pb.StreamFileDownloadCreateRequest{
		BucketName: bucketName,
		FileName:   fileName,
		RootCid:    rootCID,
	})
	if err != nil {
		return FileDownload{}, errSDK.Wrap(err)
	}

	chunks := make([]Chunk, len(res.Chunks))
	for i, chunk := range res.Chunks {
		chunks[i] = Chunk{
			CID:         chunk.Cid,
			EncodedSize: chunk.EncodedSize,
			Size:        chunk.Size,
			Index:       int64(i),
		}
	}

	return FileDownload{
		StreamID:    res.StreamId,
		BucketName:  res.BucketName,
		Name:        fileName,
		DataBlocks:  res.DataBlocks,
		TotalBlocks: res.TotalBlocks,
		Chunks:      chunks,
	}, nil
}

// CreateRangeFileDownload creates a new download request with block ranges.
func (sdk *StreamingAPI) CreateRangeFileDownload(ctx context.Context, bucketName, fileName string, start, end int64) (_ FileDownload, err error) {
	defer mon.Task()(&ctx, bucketName, fileName, start, end)(&err)

	res, err := sdk.client.FileDownloadRangeCreate(ctx, &pb.StreamFileDownloadRangeCreateRequest{
		BucketName: bucketName,
		FileName:   fileName,
		StartIndex: start,
		EndIndex:   end,
	})
	if err != nil {
		return FileDownload{}, errSDK.Wrap(err)
	}

	chunks := make([]Chunk, len(res.Chunks))
	for i, chunk := range res.Chunks {
		chunks[i] = Chunk{
			CID:         chunk.Cid,
			EncodedSize: chunk.EncodedSize,
			Size:        chunk.Size,
			Index:       int64(i) + start,
		}
	}

	return FileDownload{
		StreamID:    res.StreamId,
		BucketName:  res.BucketName,
		Name:        fileName,
		DataBlocks:  res.DataBlocks,
		TotalBlocks: res.TotalBlocks,
		Chunks:      chunks,
	}, nil
}

// Download downloads a file using streaming api.
func (sdk *StreamingAPI) Download(ctx context.Context, fileDownload FileDownload, writer io.Writer) (err error) {
	defer mon.Task()(&ctx, fileDownload)(&err)

	fileEncKey, err := encryptionKey(sdk.encryptionKey, fileDownload.BucketName, fileDownload.Name)
	if err != nil {
		return errSDK.Wrap(err)
	}

	var downloadEC *erasurecode.ErasureCode
	if fileDownload.DataBlocks < fileDownload.TotalBlocks { // erasure coding was applied to a file if data blocks < total blocks
		downloadEC, err = erasurecode.New(int(fileDownload.DataBlocks), int(fileDownload.TotalBlocks-fileDownload.DataBlocks))
		if err != nil {
			return errSDK.Wrap(err)
		}
	}

	for _, chunk := range fileDownload.Chunks {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		chunkDownload, err := sdk.createChunkDownload(ctx, fileDownload.StreamID, chunk)
		if err != nil {
			return err
		}

		if err := sdk.downloadChunkBlocks(ctx, fileDownload.StreamID, chunkDownload, fileEncKey, downloadEC, writer); err != nil {
			return err
		}
	}

	return nil
}

// DownloadV2 downloads a file using streaming api.
// This version of the function can download blocks either from filecoin's storage provider of from akave nodes(whichever is available).
func (sdk *StreamingAPI) DownloadV2(ctx context.Context, fileDownload FileDownload, writer io.Writer) (err error) {
	defer mon.Task()(&ctx, fileDownload)(&err)

	fileEncKey, err := encryptionKey(sdk.encryptionKey, fileDownload.BucketName, fileDownload.Name)
	if err != nil {
		return errSDK.Wrap(err)
	}

	var downloadEC *erasurecode.ErasureCode
	if fileDownload.DataBlocks < fileDownload.TotalBlocks { // erasure coding was applied to a file if data blocks < total blocks
		downloadEC, err = erasurecode.New(int(fileDownload.DataBlocks), int(fileDownload.TotalBlocks-fileDownload.DataBlocks))
		if err != nil {
			return errSDK.Wrap(err)
		}
	}

	for _, chunk := range fileDownload.Chunks {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		chunkDownload, err := sdk.createChunkDownloadV2(ctx, fileDownload.StreamID, chunk)
		if err != nil {
			return err
		}

		if err := sdk.downloadChunkBlocks(ctx, fileDownload.StreamID, chunkDownload, fileEncKey, downloadEC, writer); err != nil {
			return err
		}
	}

	return nil
}

// DownloadRandom downloads a file using streaming api and fetches only randomly half of the blocks.
// TODO: do we need V2 version of this function?
func (sdk *StreamingAPI) DownloadRandom(ctx context.Context, fileDownload FileDownload, writer io.Writer) (err error) {
	defer mon.Task()(&ctx, fileDownload)(&err)

	if fileDownload.DataBlocks == fileDownload.TotalBlocks {
		return errSDK.Errorf("file is not erasure coded")
	}

	fileEncKey, err := encryptionKey(sdk.encryptionKey, fileDownload.BucketName, fileDownload.Name)
	if err != nil {
		return errSDK.Wrap(err)
	}

	downloadEC, err := erasurecode.New(int(fileDownload.DataBlocks), int(fileDownload.TotalBlocks-fileDownload.DataBlocks))
	if err != nil {
		return errSDK.Wrap(err)
	}

	for _, chunk := range fileDownload.Chunks {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		chunkDownload, err := sdk.createChunkDownload(ctx, fileDownload.StreamID, chunk)
		if err != nil {
			return err
		}

		if err := sdk.downloadRandomChunkBlocks(ctx, fileDownload.StreamID, chunkDownload, fileEncKey, downloadEC, writer); err != nil {
			return err
		}
	}

	return nil
}

// FileDelete deletes a file from a bucket.
func (sdk *StreamingAPI) FileDelete(ctx context.Context, bucketName, fileName string) (err error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&err)

	_, err = sdk.client.FileDelete(ctx, &pb.StreamFileDeleteRequest{
		BucketName: bucketName,
		FileName:   fileName,
	})
	return errSDK.Wrap(err)
}

func (sdk *StreamingAPI) createChunkUpload(ctx context.Context, fileUpload FileUpload, index int64, fileEncryptionKey, data []byte) (_ FileChunkUpload, err error) {
	defer mon.Task()(&ctx, fileUpload, index)(&err)

	if len(fileEncryptionKey) > 0 {
		data, err = encryption.Encrypt(fileEncryptionKey, data, []byte(fmt.Sprintf("%d", index)))
		if err != nil {
			return FileChunkUpload{}, errSDK.Wrap(err)
		}
	}

	blockSize := BlockSize.ToInt64()
	if sdk.uploadEC != nil { // erasure coding is enabled
		data, err = sdk.uploadEC.Encode(data)
		if err != nil {
			return FileChunkUpload{}, errSDK.Wrap(err)
		}
		// equivalent to notion of shard size in erasure coding terminology
		blockSize = int64(len(data) / (sdk.uploadEC.DataBlocks + sdk.uploadEC.ParityBlocks))
	}

	chunkDAG, err := BuildDAG(ctx, bytes.NewBuffer(data), blockSize, nil)
	if err != nil {
		return FileChunkUpload{}, errSDK.Wrap(err)
	}

	protoChunk := toProtoChunk(fileUpload.StreamID, chunkDAG.CID.String(), index, chunkDAG.Blocks)
	req := &pb.StreamFileUploadChunkCreateRequest{Chunk: protoChunk}

	res, err := sdk.client.FileUploadChunkCreate(ctx, req)
	if err != nil {
		return FileChunkUpload{}, errSDK.Wrap(err)
	}

	if len(res.Blocks) != len(chunkDAG.Blocks) {
		return FileChunkUpload{}, errSDK.Errorf("received unexpected amount of blocks %d, expected %d", len(res.Blocks), len(chunkDAG.Blocks))
	}
	for i, upload := range res.Blocks {
		if chunkDAG.Blocks[i].CID != upload.Cid {
			return FileChunkUpload{}, errSDK.Errorf("block CID mismatch at position %d", i)
		}
		chunkDAG.Blocks[i].NodeAddress = upload.NodeAddress
		chunkDAG.Blocks[i].NodeID = upload.NodeId
		chunkDAG.Blocks[i].Permit = upload.Permit
	}

	return FileChunkUpload{
		StreamID:      fileUpload.StreamID,
		Index:         index,
		ChunkCID:      chunkDAG.CID,
		RawDataSize:   chunkDAG.RawDataSize,
		ProtoNodeSize: chunkDAG.ProtoNodeSize,
		Blocks:        chunkDAG.Blocks,
	}, nil
}

func (sdk *StreamingAPI) uploadChunk(ctx context.Context, fileChunkUpload FileChunkUpload, blockCount, bytesCount *atomic.Int64) (err error) {
	defer mon.Task()(&ctx, fileChunkUpload)(&err)

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(sdk.maxConcurrency)

	pool := newConnectionPool()
	defer func() {
		if err := pool.close(); err != nil {
			slog.Warn("failed to close connection", slog.String("error", err.Error()))
		}
	}()

	protoChunk := toProtoChunk(
		fileChunkUpload.StreamID,
		fileChunkUpload.ChunkCID.String(),
		fileChunkUpload.Index,
		fileChunkUpload.Blocks,
	)
	for i, block := range fileChunkUpload.Blocks {
		deriveCtx := context.WithoutCancel(ctx)
		g.Go(func() (err error) {
			defer mon.Task()(&deriveCtx, block.CID)(&err)

			client, closer, err := pool.createStreamingClient(block.NodeAddress, sdk.useConnectionPool)
			if err != nil {
				return err
			}
			if closer != nil {
				defer func() {
					if closeErr := closer(); closeErr != nil {
						slog.Warn("failed to close connection",
							slog.Int("block_index", i),
							slog.String("block_cid", block.CID),
							slog.String("chunk_cid", protoChunk.Cid),
							slog.Int64("chunk_index", fileChunkUpload.Index),
							slog.String("node_address", block.NodeAddress),
							slog.String("error", closeErr.Error()),
						)
					}
				}()
			}

			sender, err := client.FileUploadBlock(ctx)
			if err != nil {
				return err
			}

			err = sdk.uploadBlock(ctx, &pb.StreamFileBlockData{
				Data:  block.Data,
				Cid:   block.CID,
				Index: int64(i),
				Chunk: protoChunk,
			}, sender.Send)
			if err != nil && !errors.Is(err, io.EOF) {
				return err
			}

			_, closeErr := sender.CloseAndRecv()
			blockCount.Add(1)
			bytesCount.Add(int64(len(block.Data)))
			return closeErr
		})
	}

	if err := g.Wait(); err != nil {
		return errSDK.Wrap(err)
	}

	return nil
}

func (sdk *StreamingAPI) createChunkDownload(ctx context.Context, streamID string, chunk Chunk) (_ FileChunkDownload, err error) {
	defer mon.Task()(&ctx, streamID, chunk)(&err)

	res, err := sdk.client.FileDownloadChunkCreate(ctx, &pb.StreamFileDownloadChunkCreateRequest{
		StreamId: streamID,
		ChunkCid: chunk.CID,
	})
	if err != nil {
		return FileChunkDownload{}, errSDK.Wrap(err)
	}

	blocks := make([]FileBlockDownload, len(res.Blocks))
	for i, block := range res.Blocks {
		blocks[i] = FileBlockDownload{
			CID: block.Cid,
			Akave: &AkaveBlockData{
				NodeID:      block.NodeId,
				NodeAddress: block.NodeAddress,
				Permit:      block.Permit,
			},
		}
	}

	return FileChunkDownload{
		CID:         chunk.CID,
		Index:       chunk.Index,
		EncodedSize: chunk.EncodedSize,
		Size:        chunk.Size,
		Blocks:      blocks,
	}, nil
}

func (sdk *StreamingAPI) createChunkDownloadV2(ctx context.Context, streamID string, chunk Chunk) (_ FileChunkDownload, err error) {
	defer mon.Task()(&ctx, streamID, chunk)(&err)

	res, err := sdk.client.FileDownloadChunkCreateV2(ctx, &pb.StreamFileDownloadChunkCreateRequest{
		StreamId: streamID,
		ChunkCid: chunk.CID,
	})
	if err != nil {
		return FileChunkDownload{}, errSDK.Wrap(err)
	}

	blocks := make([]FileBlockDownload, len(res.Blocks))
	for i, block := range res.Blocks {
		if block.Akave == nil && block.Filecoin == nil {
			return FileChunkDownload{}, errMissingBlockMetadata
		}

		if block.Filecoin != nil {
			blocks[i] = FileBlockDownload{
				CID: block.Cid,
				Filecoin: &FilecoinBlockData{
					BaseURL: block.Filecoin.SpAddress,
				},
			}
		} else {
			blocks[i] = FileBlockDownload{
				CID: block.Cid,
				Akave: &AkaveBlockData{
					NodeID:      block.Akave.NodeId,
					NodeAddress: block.Akave.NodeAddress,
				},
			}
		}
	}

	return FileChunkDownload{
		CID:         chunk.CID,
		Index:       chunk.Index,
		EncodedSize: chunk.EncodedSize,
		Size:        chunk.Size,
		Blocks:      blocks,
	}, nil
}

func (sdk *StreamingAPI) downloadChunkBlocks(
	ctx context.Context,
	streamID string,
	chunkDownload FileChunkDownload,
	fileEncryptionKey []byte,
	downloadEC *erasurecode.ErasureCode,
	writer io.Writer,
) (err error) {

	defer mon.Task()(&ctx, streamID, chunkDownload)(&err)

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(sdk.maxConcurrency)

	pool := newConnectionPool()
	defer func() {
		if err := pool.close(); err != nil {
			slog.Warn("failed to close connection", slog.String("error", err.Error()))
		}
	}()

	type retrievedBlock struct {
		Pos  int
		CID  string
		Data []byte
	}
	ch := make(chan retrievedBlock, len(chunkDownload.Blocks))

	for i, block := range chunkDownload.Blocks {
		deriveCtx := context.WithoutCancel(ctx)

		g.Go(func() (err error) {
			defer mon.Task()(&deriveCtx, block.CID)(&err)

			blockData, err := sdk.fetchBlockData(ctx, pool, streamID, chunkDownload.CID, chunkDownload.Index, int64(i), block)
			if err != nil {
				return err
			}

			ch <- retrievedBlock{
				Pos:  i,
				CID:  block.CID,
				Data: blockData,
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return errSDK.Wrap(err)
	}

	close(ch)

	blocks := make([][]byte, len(chunkDownload.Blocks))
	for retrieved := range ch {
		data, err := ExtractBlockData(retrieved.CID, retrieved.Data)
		if err != nil {
			return errSDK.Wrap(err)
		}
		blocks[retrieved.Pos] = data
	}

	var data []byte
	if downloadEC != nil { // erasure coding is enabled
		data, err = downloadEC.ExtractData(blocks, int(chunkDownload.Size))
		if err != nil {
			return errSDK.Wrap(err)
		}
	} else {
		data = bytes.Join(blocks, nil)
	}

	if len(fileEncryptionKey) > 0 {
		data, err = encryption.Decrypt(fileEncryptionKey, data, []byte(fmt.Sprintf("%d", chunkDownload.Index)))
		if err != nil {
			return errSDK.Wrap(err)
		}
	}

	if _, err := writer.Write(data); err != nil {
		return errSDK.Wrap(err)
	}

	return nil
}

// !!!!use only with erasure coding!!!!
func (sdk *StreamingAPI) downloadRandomChunkBlocks(
	ctx context.Context,
	streamID string,
	chunkDownload FileChunkDownload,
	fileEncryptionKey []byte,
	downloadEC *erasurecode.ErasureCode,
	writer io.Writer,
) (err error) {

	defer mon.Task()(&ctx, streamID, chunkDownload)(&err)

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(sdk.maxConcurrency)

	pool := newConnectionPool()
	defer func() {
		if err := pool.close(); err != nil {
			slog.Warn("failed to close connection", slog.String("error", err.Error()))
		}
	}()

	type retrievedBlock struct {
		Pos  int
		CID  string
		Data []byte
	}
	ch := make(chan retrievedBlock, len(chunkDownload.Blocks))

	blocksMap := make(map[int]FileBlockDownload, len(chunkDownload.Blocks))
	for i, block := range chunkDownload.Blocks {
		blocksMap[i] = block
	}

	blockIndexes := maps.Keys(blocksMap)
	cryptoutils.Shuffle(blockIndexes)

	// take only half of the blocks
	for _, i := range blockIndexes[:downloadEC.DataBlocks] {
		delete(blocksMap, i)
	}

	for index, block := range blocksMap {
		deriveCtx := context.WithoutCancel(ctx)
		g.Go(func() (err error) {
			defer mon.Task()(&deriveCtx, block.CID)(&err)

			blockData, err := sdk.fetchBlockData(ctx, pool, streamID, chunkDownload.CID, chunkDownload.Index, int64(index), block)
			if err != nil {
				return err
			}

			ch <- retrievedBlock{
				Pos:  index,
				CID:  block.CID,
				Data: blockData,
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return errSDK.Wrap(err)
	}

	close(ch)

	blocks := make([][]byte, len(chunkDownload.Blocks))
	for retrieved := range ch {
		data, err := ExtractBlockData(retrieved.CID, retrieved.Data)
		if err != nil {
			return errSDK.Wrap(err)
		}
		blocks[retrieved.Pos] = data
	}

	var data []byte
	if downloadEC != nil { // erasure coding is enabled
		data, err = downloadEC.ExtractData(blocks, int(chunkDownload.Size))
		if err != nil {
			return errSDK.Wrap(err)
		}
	} else {
		data = bytes.Join(blocks, nil)
	}

	if len(fileEncryptionKey) > 0 {
		data, err = encryption.Decrypt(fileEncryptionKey, data, fmt.Appendf(nil, "%d", chunkDownload.Index))
		if err != nil {
			return errSDK.Wrap(err)
		}
	}

	if _, err := writer.Write(data); err != nil {
		return errSDK.Wrap(err)
	}

	return nil
}

func (sdk *StreamingAPI) commitStream(ctx context.Context, upload FileUpload, rootCID string, chunkCount, fileSize int64) (_ FileMeta, err error) {
	defer mon.Task()(&ctx, upload, rootCID, chunkCount)(&err)

	res, err := sdk.client.FileUploadCommit(ctx, &pb.StreamFileUploadCommitRequest{
		StreamId:   upload.StreamID,
		RootCid:    rootCID,
		ChunkCount: chunkCount,
		Size:       fileSize,
	})
	if err != nil {
		return FileMeta{}, errSDK.Wrap(err)
	}

	return FileMeta{
		StreamID:    res.StreamId,
		RootCID:     rootCID,
		BucketName:  res.BucketName,
		Name:        res.FileName,
		EncodedSize: res.EncodedSize,
		Size:        res.Size,
		DataBlocks:  upload.DataBlocks,
		TotalBlocks: upload.TotalBlocks,
		CreatedAt:   upload.CreatedAt,
		CommitedAt:  res.CommittedAt.AsTime(),
	}, nil
}

func (sdk *StreamingAPI) uploadBlock(ctx context.Context, block *pb.StreamFileBlockData, sender func(*pb.StreamFileBlockData) error) (err error) {
	defer mon.Task()(&ctx, block.Cid, block.Index, block.Chunk.Cid, block.Chunk.Index, block.Chunk.StreamId)(&err)

	data := block.Data
	dataLen := int64(len(data))
	if dataLen == 0 {
		return nil
	}

	i := int64(0)
	for ; i < dataLen; i += sdk.blockPartSize {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			end := i + sdk.blockPartSize
			if end > dataLen {
				end = dataLen
			}

			block.Data = data[i:end:end]

			if err := sender(block); err != nil {
				return err
			}

			// these fields are only required for the first part, skip them for the rest.
			block.Chunk = nil
			block.Cid = ""
		}
	}

	return nil
}

func (sdk *StreamingAPI) fetchBlockData(
	ctx context.Context,
	pool *connectionPool,
	streamID, chunkCID string,
	chunkIndex, blockIdndex int64,
	block FileBlockDownload,
) ([]byte, error) {

	if block.Akave == nil && block.Filecoin == nil {
		return nil, errMissingBlockMetadata
	}

	if block.Filecoin != nil {
		cid, err := cid.Decode(block.CID)
		if err != nil {
			return nil, err
		}
		data, err := sdk.spClient.FetchBlock(ctx, block.Filecoin.BaseURL, cid)
		if err != nil {
			return nil, err
		}
		return data.RawData(), nil
	}

	client, closer, err := pool.createStreamingClient(block.Akave.NodeAddress, sdk.useConnectionPool)
	if err != nil {
		return nil, err
	}
	if closer != nil {
		defer func() {
			if closeErr := closer(); closeErr != nil {
				slog.Warn("failed to close connection", slog.String("block_cid", block.CID), slog.String("error", closeErr.Error()))
			}
		}()
	}

	downloadClient, err := client.FileDownloadBlock(ctx, &pb.StreamFileDownloadBlockRequest{
		StreamId:   streamID,
		ChunkCid:   chunkCID,
		ChunkIndex: chunkIndex,
		BlockCid:   block.CID,
		BlockIndex: blockIdndex,
	})
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	for {
		blockData, err := downloadClient.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		_, _ = buf.Write(blockData.Data)
	}

	return buf.Bytes(), nil
}

func toProtoChunk(streamID, cid string, index int64, blocks []FileBlockUpload) *pb.Chunk {
	pbBlocks := make([]*pb.Chunk_Block, len(blocks))
	for i, block := range blocks {
		pbBlocks[i] = &pb.Chunk_Block{
			Cid:  block.CID,
			Size: int64(len(block.Data)),
		}
	}
	return &pb.Chunk{
		StreamId: streamID,
		Cid:      cid,
		Index:    index,
		Size:     int64(0), // will be removed soon
		Blocks:   pbBlocks,
	}
}

func toFileMeta(protoFile *pb.File, bucketName string) FileMeta {
	return FileMeta{
		StreamID:    protoFile.StreamId,
		RootCID:     protoFile.RootCid,
		BucketName:  bucketName,
		Name:        protoFile.Name,
		EncodedSize: protoFile.EncodedSize,
		DataBlocks:  protoFile.DataBlocks,
		TotalBlocks: protoFile.TotalBlocks,
		Size:        protoFile.Size,
		CreatedAt:   protoFile.CreatedAt.AsTime(),
		CommitedAt:  protoFile.CommitedAt.AsTime(),
	}
}
