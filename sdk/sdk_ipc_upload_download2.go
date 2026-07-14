// Copyright (C) 2026 Akave
// See LICENSE for copying information.

package sdk

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/sync/errgroup"

	"github.com/akave-ai/akavesdk/private/encryption"
	"github.com/akave-ai/akavesdk/private/encryption/streamenc"
	"github.com/akave-ai/akavesdk/private/erasurecode"
	"github.com/akave-ai/akavesdk/private/ipc"
	"github.com/akave-ai/akavesdk/private/ipc/contracts"
	"github.com/akave-ai/akavesdk/private/memory"
	"github.com/akave-ai/akavesdk/private/pb"
)

const (
	// targetCiphertextSize is the target size of encrypted blocks produced by Upload2.
	targetCiphertextSize = int(16 * memory.MB)

	maxStripeSize = int(2 * memory.MiB)
)

// Upload2 uploads a file using streamenc encryption and per-stripe erasure coding.
// It requires both erasure coding and encryption to be configured on the SDK.
func (sdk *IPC) Upload2(ctx context.Context, fileUpload *IPCFileUpload, reader io.Reader) (_ IPCFileMetaV2, err error) {
	if sdk.ec == nil {
		return IPCFileMetaV2{}, errSDK.Errorf("erasure coding is required for Upload2")
	}
	if len(sdk.encryptionKey) == 0 {
		return IPCFileMetaV2{}, errSDK.Errorf("encryption is required for Upload2")
	}

	if fileUpload.state.isCommitted {
		return IPCFileMetaV2{}, errSDK.Errorf("file is already committed")
	}

	var isContinuation bool
	if fileUpload.state.chunkCount > 0 {
		isContinuation = true
	}

	var bucket contracts.IStorageBucket
	err = sdk.withRetry.Do(ctx, func() (bool, error) {
		bucket, err = sdk.ipc.Storage.GetBucketByName(
			&bind.CallOpts{Context: ctx, From: sdk.ipc.Auth.From},
			fileUpload.BucketName,
			sdk.ipc.Auth.From,
			big.NewInt(0), big.NewInt(0),
		)
		return true, err
	})
	if err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	fileEncKey, err := encryptionKey(sdk.encryptionKey, fileUpload.BucketName, fileUpload.Name)
	if err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(err)
	}

	bufSize, err := streamenc.MaxPlaintextSizeForTarget(targetCiphertextSize)
	if err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(err)
	}
	buf := make([]byte, bufSize, targetCiphertextSize)

	g, chunkCtx := errgroup.WithContext(ctx)
	fileUploadChunksCh := make(chan IPCFileChunkUploadV2)
	waitTransactionsCh := make(chan BatchTransaction, sdk.chunkBuffer)

	g.Go(func() error {
		defer close(waitTransactionsCh)

		chunkIndex := int64(0)
		if isContinuation {
			if err := skipToPosition(reader, fileUpload.state.actualFileSize); err != nil {
				return err
			}
			chunkIndex = fileUpload.state.chunkCount
		}

		stopReading := false
		for !stopReading {
			var batch []ChunkData
			for range sdk.chunkBatchSize {
				buf = buf[:bufSize]
				n, err := readUpTo(reader, buf)
				if err != nil && !errors.Is(err, io.EOF) {
					return err
				}
				if errors.Is(err, io.EOF) && chunkIndex == 0 {
					return fmt.Errorf("empty file")
				}

				if n == 0 {
					stopReading = true
					break
				}

				result, err := sdk.createChunkUpload2(chunkCtx, chunkIndex, fileEncKey, buf[:n], bucket.Id, fileUpload.Name)
				if err != nil {
					return err
				}

				batch = append(batch, result)
				chunkIndex++
			}

			if len(batch) > 0 {
				tx, err := sdk.createBatchedChunkTransaction(chunkCtx, batch, bucket.Id, fileUpload.Name)
				if err != nil {
					return err
				}

				chunks := make([]IPCFileChunkUploadV2, len(batch))
				for i, result := range batch {
					if err := fileUpload.state.preCreateChunk(result.ChunkUpload, tx); err != nil {
						return err
					}
					chunks[i] = result.ChunkUpload
				}

				select {
				case <-chunkCtx.Done():
					return chunkCtx.Err()
				case waitTransactionsCh <- BatchTransaction{Chunks: chunks, Tx: tx}:
				}
			}
		}
		return nil
	})

	g.Go(func() error {
		defer close(fileUploadChunksCh)

		if isContinuation {
			for _, chunkWithTx := range fileUpload.state.listPreCreatedChunks() {
				if err := sdk.ipc.WaitForTx(chunkCtx, chunkWithTx.tx.Hash()); err != nil {
					return err
				}

				select {
				case <-chunkCtx.Done():
					return chunkCtx.Err()
				case fileUploadChunksCh <- chunkWithTx.chunk:
				}
			}
		}

		for {
			select {
			case <-chunkCtx.Done():
				return chunkCtx.Err()
			case batchResult, ok := <-waitTransactionsCh:
				if !ok {
					return nil
				}

				if err := sdk.ipc.WaitForTx(chunkCtx, batchResult.Tx.Hash()); err != nil {
					return err
				}

				for _, chunk := range batchResult.Chunks {
					select {
					case <-chunkCtx.Done():
						return chunkCtx.Err()
					case fileUploadChunksCh <- chunk:
					}
				}
			}
		}
	})

	g.Go(func() error {
		for {
			select {
			case <-chunkCtx.Done():
				return chunkCtx.Err()
			case chunkUpload, ok := <-fileUploadChunksCh:
				if !ok {
					return nil
				}

				if err := sdk.uploadChunk(chunkCtx, chunkUpload, fileUpload.blocksCounter, fileUpload.bytesCounter, isContinuation); err != nil {
					return err
				}

				fileUpload.state.chunkUploaded(chunkUpload)
				fileUpload.chunksCounter.Add(1)
			}
		}
	})

	if err := g.Wait(); err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(err)
	}

	rootCID, err := fileUpload.state.dagRoot.Build()
	if err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(err)
	}

	var fileMeta contracts.IStorageFile
	err = sdk.withRetry.Do(ctx, func() (bool, error) {
		fileMeta, err = sdk.ipc.Storage.GetFileByName(
			&bind.CallOpts{Context: ctx, From: sdk.ipc.Auth.From},
			bucket.Id,
			fileUpload.Name,
		)
		return true, err
	})
	if err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	fileID := ipc.CalculateFileID(bucket.Id[:], fileUpload.Name)
	var isFilled bool
	for !isFilled {
		err = sdk.withRetry.Do(ctx, func() (bool, error) {
			isFilled, err = sdk.ipc.Storage.IsFileFilled(&bind.CallOpts{Context: ctx}, fileID)
			return true, err
		})
		if err != nil {
			return IPCFileMetaV2{}, errSDK.Wrap(ipc.ErrorHashToError(err))
		}

		time.Sleep(time.Second)
	}

	var tx *types.Transaction
	err = sdk.withRetry.Do(ctx, func() (bool, error) {
		tx, err = sdk.ipc.Storage.CommitFile(
			sdk.ipc.Auth,
			bucket.Id,
			fileUpload.Name,
			big.NewInt(fileUpload.state.encodedFileSize),
			big.NewInt(fileUpload.state.actualFileSize),
			rootCID.Bytes(),
		)
		return isRetryableTxError(err), err
	})
	if err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	fileUpload.state.isCommitted = true

	return IPCFileMetaV2{
		RootCID:     rootCID.String(),
		BucketName:  fileUpload.BucketName,
		Name:        fileUpload.Name,
		Size:        fileUpload.state.actualFileSize,
		EncodedSize: fileUpload.state.encodedFileSize,
		CreatedAt:   time.Unix(fileMeta.CreatedAt.Int64(), 0),
		CommittedAt: time.Now().UTC(),
	}, errSDK.Wrap(sdk.ipc.WaitForTx(ctx, tx.Hash()))
}

// Download2 downloads a file that was uploaded with Upload2.
func (sdk *IPC) Download2(ctx context.Context, fileDownload IPCFileDownload, writer io.Writer) (err error) {
	defer mon.Task()(&ctx, fileDownload)(&err)

	if sdk.ec == nil {
		return errSDK.Errorf("erasure coding is required for Download2")
	}
	if len(sdk.encryptionKey) == 0 {
		return errSDK.Errorf("encryption is required for Download2")
	}

	fileEncKey, err := encryptionKey(sdk.encryptionKey, fileDownload.BucketName, fileDownload.Name)
	if err != nil {
		return errSDK.Wrap(err)
	}

	g, ctx := errgroup.WithContext(ctx)
	chunkDownloadCh := make(chan FileChunkDownload, sdk.chunkBuffer)

	g.Go(func() error {
		defer close(chunkDownloadCh)

		for _, chunk := range fileDownload.Chunks {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			chunkDownload, err := sdk.createChunkDownload(ctx, fileDownload.BucketName, fileDownload.Name, chunk)
			if err != nil {
				return err
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			case chunkDownloadCh <- chunkDownload:
			}
		}
		return nil
	})

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()

			case chunkDownload, ok := <-chunkDownloadCh:
				if !ok {
					return nil
				}

				err = sdk.downloadChunkBlocksV2(
					ctx,
					fileDownload.BucketName,
					fileDownload.Name,
					sdk.ipc.Auth.From.String(),
					chunkDownload,
					fileEncKey,
					writer,
					fileDownload.blocksCounter,
					fileDownload.bytesCounter,
				)
				if err != nil {
					return err
				}

				fileDownload.chunksCounter.Add(1)
			}
		}
	})

	return g.Wait()
}

func (sdk *IPC) createChunkUpload2(
	ctx context.Context,
	index int64,
	fileEncKey, data []byte,
	bucketID [32]byte,
	fileName string,
) (_ ChunkData, err error) {

	actualSize := int64(len(data))

	n, err := streamenc.Encrypt(fileEncKey, data, strconv.FormatInt(index, 10))
	if err != nil {
		return ChunkData{}, errSDK.Wrap(err)
	}
	data = data[:n]

	stripedData := erasurecode.SplitStripes(data, int(2*memory.MiB.ToInt64()))

	stripedShards := make([][][]byte, len(stripedData))
	for i, stripe := range stripedData {
		shards, err := sdk.ec.EncodeRaw(stripe)
		if err != nil {
			return ChunkData{}, errSDK.Wrap(err)
		}
		stripedShards[i] = shards
	}

	numShards := sdk.ec.DataBlocks + sdk.ec.ParityBlocks
	var blockSize int64
	for _, shards := range stripedShards {
		blockSize += int64(len(shards[0]))
	}

	data = make([]byte, 0, int(blockSize)*numShards)
	for i := range numShards {
		for _, shards := range stripedShards {
			data = append(data, shards[i]...)
		}
	}

	chunkDAG, err := BuildDAG(ctx, bytes.NewBuffer(data), blockSize)
	if err != nil {
		return ChunkData{}, errSDK.Wrap(err)
	}

	cids, sizes, protoChunk, err := toIPCProtoChunk(chunkDAG.CID.String(), index, actualSize, chunkDAG.Blocks)
	if err != nil {
		return ChunkData{}, err
	}
	req := &pb.IPCFileUploadChunkCreateRequest{
		Chunk:    protoChunk,
		BucketId: bucketID[:],
		FileName: fileName,
	}

	res, err := sdk.client.FileUploadChunkCreate(ctx, req)
	if err != nil {
		return ChunkData{}, errSDK.Wrap(err)
	}

	if len(res.Blocks) != len(chunkDAG.Blocks) {
		return ChunkData{}, errSDK.Errorf("received unexpected amount of blocks %d, expected %d", len(res.Blocks), len(chunkDAG.Blocks))
	}
	for i, upload := range res.Blocks {
		if chunkDAG.Blocks[i].CID != upload.Cid {
			return ChunkData{}, errSDK.Errorf("block CID mismatch at position %d", i)
		}
		chunkDAG.Blocks[i].NodeAddress = upload.NodeAddress
		chunkDAG.Blocks[i].NodeID = upload.NodeId
		chunkDAG.Blocks[i].Permit = upload.Permit
	}

	return ChunkData{
		ChunkUpload: IPCFileChunkUploadV2{
			Index:       index,
			ChunkCID:    chunkDAG.CID,
			ActualSize:  actualSize,
			RawDataSize: uint64(actualSize),
			EncodedSize: chunkDAG.EncodedSize,
			Blocks:      chunkDAG.Blocks,
			BucketID:    bucketID,
			FileName:    fileName,
		},
		CIDs:  cids,
		Sizes: sizes,
	}, nil
}

func (sdk *IPC) downloadChunkBlocksV2(
	ctx context.Context,
	bucketName, fileName, address string,
	chunkDownload FileChunkDownload,
	fileEncryptionKey []byte,
	writer io.Writer,
	blockCount, bytesCount *atomic.Int64,
) (err error) {

	defer mon.Task()(&ctx, chunkDownload)(&err)

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(sdk.maxConcurrency)

	type retrievedBlock struct {
		Pos  int
		CID  string
		Data []byte
	}
	ch := make(chan retrievedBlock, len(chunkDownload.Blocks))

	for i, block := range chunkDownload.Blocks {
		deriveCtx := context.WithoutCancel(ctx)
		g.Go(func() (err error) {
			defer mon.TaskNamed("(*IPC).downloadBlockV2")(&deriveCtx, block.CID)(&err)

			blockData, err := sdk.fetchBlockData(ctx, chunkDownload.CID, bucketName, fileName, address, chunkDownload.Index, int64(i), block)
			if err != nil {
				return err
			}
			ch <- retrievedBlock{Pos: i, CID: block.CID, Data: blockData}
			blockCount.Add(1)
			bytesCount.Add(int64(len(blockData)))
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

	numShards := sdk.ec.DataBlocks + sdk.ec.ParityBlocks
	if len(blocks) != numShards {
		return errSDK.Errorf("expected %d blocks, got %d", numShards, len(blocks))
	}

	// Recover the plaintext size from the ciphertext header.
	// Each DAG block is a shard column. For a single-stripe chunk len(blocks[i])
	// equals shardSize0, so consecutive blocks form contiguous ciphertext bytes.
	// For a multi-stripe chunk shardSize0 == ceil(maxStripeSize/DataBlocks) >= HeaderSize,
	// so the header fits entirely in blocks[0].
	// In both cases, concatenating whole blocks until we have >= HeaderSize bytes is correct.
	headerBuf := make([]byte, 0, streamenc.HeaderSize)
	for _, b := range blocks[:sdk.ec.DataBlocks] {
		need := streamenc.HeaderSize - len(headerBuf)
		headerBuf = append(headerBuf, b[:min(need, len(b))]...)
		if len(headerBuf) >= streamenc.HeaderSize {
			break
		}
	}
	_, _, plaintextSize32, err := streamenc.ParseHeader(headerBuf[:streamenc.HeaderSize])
	if err != nil {
		return errSDK.Wrap(err)
	}
	plaintextSize := int(plaintextSize32)
	ciphertextSize := plaintextSize + streamenc.Overhead(plaintextSize)
	ciphertext := make([]byte, 0, ciphertextSize)
	numStripes := encryption.CeilDiv(ciphertextSize, maxStripeSize)

	// blockOffset tracks how far into each block we have consumed across stripes.
	// All blocks advance by the same shardSize per stripe because they were built
	// column-major during upload.
	blockOffset := 0
	for j := range numStripes {
		// Last stripe may be shorter than maxStripeSize.
		stripeSize := min(maxStripeSize, ciphertextSize-j*maxStripeSize)
		// Shard size matches what EncodePlain produced: ceil(stripeSize / DataBlocks).
		shardSize := encryption.CeilDiv(stripeSize, sdk.ec.DataBlocks)

		// Slice the j-th shard segment from each block (column → row conversion).
		shards := make([][]byte, numShards)
		for i, block := range blocks {
			shards[i] = block[blockOffset : blockOffset+shardSize]
		}

		// Decode the stripe data from the shards, which may involve erasure code reconstruction if some blocks are missing.
		stripeData, err := sdk.ec.ExtractDataRaw(shards, stripeSize)
		if err != nil {
			return errSDK.Wrap(err)
		}
		ciphertext = append(ciphertext, stripeData...)
		blockOffset += shardSize
	}

	n, err := streamenc.DecryptAllBlocks(fileEncryptionKey, ciphertext, strconv.FormatInt(chunkDownload.Index, 10), streamenc.Version)
	if err != nil {
		return errSDK.Wrap(err)
	}

	if _, err := writer.Write(ciphertext[:n]); err != nil {
		return errSDK.Wrap(err)
	}

	return nil
}
