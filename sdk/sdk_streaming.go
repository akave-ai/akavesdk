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
	"slices"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"akave.ai/akavesdk/private/encryption"
	"akave.ai/akavesdk/private/pb"
)

// StreamingAPI exposes SDK file streaming API.
type StreamingAPI struct {
	client            pb.StreamAPIClient
	conn              *grpc.ClientConn
	maxConcurrency    int
	blockPartSize     int64
	useConnectionPool bool
	encryptionKey     []byte // empty means no encryption
	maxBlocksInChunk  int
}

// FileInfo returns meta information for single file by bucket and file name.
func (sdk *StreamingAPI) FileInfo(ctx context.Context, bucketName string, fileName string) (_ FileMetaV2, err error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&err)
	if bucketName == "" {
		return FileMetaV2{}, errSDK.Errorf("empty bucket name")
	}
	if fileName == "" {
		return FileMetaV2{}, errSDK.Errorf("empty file name")
	}

	res, err := sdk.client.FileView(ctx, &pb.StreamFileViewRequest{
		BucketName: bucketName,
		FileName:   fileName,
	})
	if err != nil {
		return FileMetaV2{}, errSDK.Wrap(err)
	}

	return FileMetaV2{
		StreamID:   res.GetStreamId(),
		RootCID:    res.GetRootCid(),
		BucketID:   res.GetBucketId(),
		Name:       res.GetFileName(),
		Size:       res.GetSize(),
		CreatedAt:  res.CreatedAt.AsTime(),
		CommitedAt: res.CommittedAt.AsTime(),
	}, nil
}

// ListFiles returns list of files in a particular bucket.
func (sdk *StreamingAPI) ListFiles(ctx context.Context, bucketName string) (_ []FileMetaV2, err error) {
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

	files := make([]FileMetaV2, 0, len(resp.Files))
	for _, fileMeta := range resp.Files {
		files = append(files, FileMetaV2{
			StreamID:   fileMeta.GetStreamId(),
			RootCID:    fileMeta.GetRootCid(),
			BucketID:   BucketIDFromName(bucketName),
			Name:       fileMeta.GetName(),
			Size:       fileMeta.GetSize(),
			CreatedAt:  fileMeta.GetCreatedAt().AsTime(),
			CommitedAt: fileMeta.GetCommitedAt().AsTime(),
		})
	}

	return files, nil
}

// CreateFileUpload creates a new file upload request.
func (sdk *StreamingAPI) CreateFileUpload(ctx context.Context, bucketName string, fileName string) (_ FileUploadV2, err error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&err)

	if bucketName == "" {
		return FileUploadV2{}, errSDK.Errorf("empty bucket name")
	}

	req := &pb.StreamFileUploadCreateRequest{
		BucketName: bucketName,
		FileName:   fileName,
	}

	res, err := sdk.client.FileUploadCreate(ctx, req)
	if err != nil {
		return FileUploadV2{}, errSDK.Wrap(err)
	}

	return FileUploadV2{
		BucketID:  res.BucketId,
		Name:      res.FileName,
		StreamID:  res.StreamId,
		CreatedAt: res.CreatedAt.AsTime(),
	}, nil
}

// Upload uploads a file using streaming api.
func (sdk *StreamingAPI) Upload(ctx context.Context, upload FileUploadV2, reader io.Reader) (_ FileMetaV2, err error) {
	defer mon.Task()(&ctx, upload)(&err)

	isEmptyFile := true
	buf := make([]byte, sdk.maxBlocksInChunk*int(BlockSize.ToInt64()))

	dagRoot, err := NewDAGRoot()
	if err != nil {
		return FileMetaV2{}, errSDK.Wrap(err)
	}

	var i int64
	for ; ; i++ {
		select {
		case <-ctx.Done():
			return FileMetaV2{}, ctx.Err()
		default:
		}

		n, err := io.ReadAtLeast(reader, buf, 1)
		if err != nil {
			if errors.Is(err, io.EOF) {
				if isEmptyFile {
					return FileMetaV2{}, errSDK.Errorf("empty file")
				}
				break
			}
			return FileMetaV2{}, err
		}
		isEmptyFile = false

		chunkUpload, err := sdk.createChunkUpload(ctx, upload, i, bytes.NewReader(buf[:n]))
		if err != nil {
			return FileMetaV2{}, err
		}

		if err := dagRoot.AddLink(chunkUpload.ChunkCID, chunkUpload.RawDataSize, chunkUpload.ProtoNodeSize); err != nil {
			return FileMetaV2{}, errSDK.Wrap(err)
		}

		if err := sdk.uploadChunkV2(ctx, chunkUpload); err != nil {
			return FileMetaV2{}, err
		}
	}

	rootCID, err := dagRoot.Build()
	if err != nil {
		return FileMetaV2{}, errSDK.Wrap(err)
	}

	fileMeta, err := sdk.commitStream(ctx, upload, rootCID.String())
	if err != nil {
		return FileMetaV2{}, err
	}

	return fileMeta, nil
}

// CreateFileDownload creates a new download request.
func (sdk *StreamingAPI) CreateFileDownload(ctx context.Context, bucketName, fileName string) (_ FileDownloadV2, err error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&err)

	res, err := sdk.client.FileDownloadCreate(ctx, &pb.StreamFileDownloadCreateRequest{
		BucketName: bucketName,
		FileName:   fileName,
	})
	if err != nil {
		return FileDownloadV2{}, errSDK.Wrap(err)
	}

	chunks := make([]Chunk, len(res.Chunks))
	for i, chunk := range res.Chunks {
		chunks[i] = Chunk{
			CID:   chunk.Cid,
			Size:  chunk.Size,
			Index: int64(i),
		}
	}

	return FileDownloadV2{
		StreamID: res.StreamId,
		BucketID: res.BucketId,
		Name:     fileName,
		Chunks:   chunks,
	}, nil
}

// CreateRangeFileDownload creates a new download request with block ranges.
func (sdk *StreamingAPI) CreateRangeFileDownload(ctx context.Context, bucketName, fileName string, start, end int64) (_ FileDownloadV2, err error) {
	defer mon.Task()(&ctx, bucketName, fileName, start, end)(&err)

	res, err := sdk.client.FileDownloadRangeCreate(ctx, &pb.StreamFileDownloadRangeCreateRequest{
		BucketName: bucketName,
		FileName:   fileName,
		StartIndex: start,
		EndIndex:   end,
	})
	if err != nil {
		return FileDownloadV2{}, errSDK.Wrap(err)
	}

	chunks := make([]Chunk, len(res.Chunks))
	for i, chunk := range res.Chunks {
		chunks[i] = Chunk{
			CID:   chunk.Cid,
			Size:  chunk.Size,
			Index: int64(i) + start,
		}
	}

	return FileDownloadV2{
		StreamID: res.StreamId,
		BucketID: res.BucketId,
		Name:     fileName,
		Chunks:   chunks,
	}, nil
}

// Download downloads a file using streaming api.
func (sdk *StreamingAPI) Download(ctx context.Context, fileDownload FileDownloadV2, writer io.Writer) (err error) {
	defer mon.Task()(&ctx, fileDownload)(&err)

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

		if err := sdk.downloadChunkBlocksV2(ctx, fileDownload.BucketID, fileDownload.StreamID, chunkDownload, writer); err != nil {
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

func (sdk *StreamingAPI) createChunkUpload(ctx context.Context, fileUpload FileUploadV2, index int64, reader io.Reader) (_ FileChunkUploadV2, err error) {
	defer mon.Task()(&ctx, fileUpload, index)(&err)

	encKey, err := encryptionKey(sdk.encryptionKey, fileUpload.BucketID, fileUpload.StreamID)
	if err != nil {
		return FileChunkUploadV2{}, errSDK.Wrap(err)
	}

	chunkDAG, err := BuildDAG(ctx, reader, BlockSize.ToInt64(), encKey)
	if err != nil {
		return FileChunkUploadV2{}, err
	}

	protoChunk := toProtoChunk(fileUpload.StreamID, chunkDAG.CID.String(), index, chunkDAG.Blocks)
	req := &pb.StreamFileUploadChunkCreateRequest{Chunk: protoChunk}

	res, err := sdk.client.FileUploadChunkCreate(ctx, req)
	if err != nil {
		return FileChunkUploadV2{}, errSDK.Wrap(err)
	}

	if len(res.Blocks) != len(chunkDAG.Blocks) {
		return FileChunkUploadV2{}, errSDK.Errorf("received unexpected amount of blocks %d, expected %d", len(res.Blocks), len(chunkDAG.Blocks))
	}
	for i, upload := range res.Blocks {
		if chunkDAG.Blocks[i].CID != upload.Cid {
			return FileChunkUploadV2{}, errSDK.Errorf("block CID mismatch at position %d", i)
		}
		chunkDAG.Blocks[i].NodeAddress = upload.NodeAddress
		chunkDAG.Blocks[i].NodeID = upload.NodeId
		chunkDAG.Blocks[i].Permit = upload.Permit
	}

	return FileChunkUploadV2{
		StreamID:      fileUpload.StreamID,
		Index:         index,
		ChunkCID:      chunkDAG.CID,
		RawDataSize:   chunkDAG.RawDataSize,
		ProtoNodeSize: chunkDAG.ProtoNodeSize,
		Blocks:        chunkDAG.Blocks,
	}, nil
}

func (sdk *StreamingAPI) uploadChunkV2(ctx context.Context, fileChunkUpload FileChunkUploadV2) (err error) {
	defer mon.Task()(&ctx, fileChunkUpload)(&err)

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(sdk.maxConcurrency)

	pool := newConnectionPool()
	defer func() {
		if err := pool.close(); err != nil {
			slog.Warn("failed to close connection", slog.String("error", err.Error()))
		}
	}()

	protoChunk := toProtoChunk(fileChunkUpload.StreamID, fileChunkUpload.ChunkCID.String(), fileChunkUpload.Index, fileChunkUpload.Blocks)
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

			err = sdk.uploadBlockV2(ctx, &pb.StreamFileBlockData{
				Data:  block.Data,
				Cid:   block.CID,
				Index: int64(i),
				Chunk: protoChunk,
			}, sender.Send)
			if err != nil {
				return err
			}

			_, closeErr := sender.CloseAndRecv()
			return closeErr
		})
	}

	if err := g.Wait(); err != nil {
		return errSDK.Wrap(err)
	}

	return nil
}

func (sdk *StreamingAPI) createChunkDownload(ctx context.Context, streamID string, chunk Chunk) (_ FileChunkDownloadV2, err error) {
	defer mon.Task()(&ctx, streamID, chunk)(&err)

	res, err := sdk.client.FileDownloadChunkCreate(ctx, &pb.StreamFileDownloadChunkCreateRequest{
		StreamId: streamID,
		ChunkCid: chunk.CID,
	})
	if err != nil {
		return FileChunkDownloadV2{}, errSDK.Wrap(err)
	}

	blocks := make([]FileBlock, len(res.Blocks))
	for i, block := range res.Blocks {
		blocks[i] = FileBlock{
			CID:         block.Cid,
			NodeID:      block.NodeId,
			NodeAddress: block.NodeAddress,
			Permit:      block.Permit,
		}
	}

	return FileChunkDownloadV2{
		CID:    chunk.CID,
		Index:  chunk.Index,
		Size:   chunk.Size,
		Blocks: blocks,
	}, nil
}

func (sdk *StreamingAPI) downloadChunkBlocksV2(
	ctx context.Context,
	bucketID, streamID string,
	chunkDownload FileChunkDownloadV2,
	writer io.Writer,
) (err error) {

	defer mon.Task()(&ctx, bucketID, streamID, chunkDownload)(&err)

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

	encKey, err := encryptionKey(sdk.encryptionKey, bucketID, streamID)
	if err != nil {
		return errSDK.Wrap(err)
	}

	for i, block := range chunkDownload.Blocks {
		i, block := i, block
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
							slog.String("error", closeErr.Error()),
						)
					}
				}()
			}

			downloadClient, err := client.FileDownloadBlock(ctx, &pb.StreamFileDownloadBlockRequest{
				StreamId:   streamID,
				ChunkCid:   chunkDownload.CID,
				ChunkIndex: chunkDownload.Index,
				BlockCid:   block.CID,
				BlockIndex: int64(i),
			})
			if err != nil {
				return err
			}

			var buf bytes.Buffer
			for {
				blockData, err := downloadClient.Recv()
				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					return err
				}
				_, _ = buf.Write(blockData.Data)
			}

			ch <- retrievedBlock{
				Pos:  i,
				CID:  block.CID,
				Data: buf.Bytes(),
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return errSDK.Wrap(err)
	}

	close(ch)

	recvBlocks := make([]retrievedBlock, 0, len(chunkDownload.Blocks))
	for retrieved := range ch {
		recvBlocks = append(recvBlocks, retrieved)
	}
	slices.SortStableFunc(recvBlocks, func(a, b retrievedBlock) int {
		return a.Pos - b.Pos
	})

	for i, block := range recvBlocks {
		data, err := ExtractBlockData(block.CID, block.Data)
		if err != nil {
			return errSDK.Wrap(err)
		}

		if len(encKey) > 0 { // if ecnryption is enabled
			info := fmt.Sprintf("block_%d", i)
			data, err = encryption.Decrypt(encKey, data, []byte(info))
			if err != nil {
				return errSDK.Wrap(err)
			}
		}

		if _, err := writer.Write(data); err != nil {
			return errSDK.Wrap(err)
		}
	}

	return nil
}

func (sdk *StreamingAPI) commitStream(ctx context.Context, upload FileUploadV2, rootCID string) (_ FileMetaV2, err error) {
	defer mon.Task()(&ctx, upload, rootCID)(&err)

	res, err := sdk.client.FileUploadCommit(ctx, &pb.StreamFileUploadCommitRequest{
		StreamId: upload.StreamID,
		RootCid:  rootCID,
	})
	if err != nil {
		return FileMetaV2{}, errSDK.Wrap(err)
	}

	return FileMetaV2{
		StreamID:   res.StreamId,
		RootCID:    rootCID,
		BucketID:   res.BucketId,
		Name:       res.FileName,
		Size:       res.FileSize,
		CreatedAt:  upload.CreatedAt,
		CommitedAt: res.CommittedAt.AsTime(),
	}, nil
}

func (sdk *StreamingAPI) uploadBlockV2(ctx context.Context, block *pb.StreamFileBlockData, sender func(*pb.StreamFileBlockData) error) (err error) {
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

func toProtoChunk(streamID, cid string, index int64, blocks []FileBlock) *pb.Chunk {
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
		Blocks:   pbBlocks,
	}
}
