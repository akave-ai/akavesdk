// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package sdk

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math/big"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/peer"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/akave-ai/akavesdk/private/cids"
	"github.com/akave-ai/akavesdk/private/encryption"
	"github.com/akave-ai/akavesdk/private/erasurecode"
	"github.com/akave-ai/akavesdk/private/httpext"
	"github.com/akave-ai/akavesdk/private/ipc"
	"github.com/akave-ai/akavesdk/private/ipc/contracts"
	"github.com/akave-ai/akavesdk/private/pb"
	"github.com/akave-ai/akavesdk/private/retry"
)

// IPC exposes SDK ipc API.
type IPC struct {
	client     pb.IPCNodeAPIClient
	conn       *grpc.ClientConn
	ipc        *ipc.Client
	httpClient *http.Client
	ec         *erasurecode.ErasureCode

	privateKey            string
	chainID               *big.Int
	storageAddress        string
	maxConcurrency        int
	blockPartSize         int64
	useConnectionPool     bool
	encryptionKey         []byte // empty means no encryption
	maxBlocksInChunk      int
	useMetadataEncryption bool
	// chunkBuffer controls the size of the buffer for chunk uploads.
	chunkBuffer int
	batchSize   int

	withRetry retry.WithRetry
}

// CreateBucket creates a new bucket.
func (sdk *IPC) CreateBucket(ctx context.Context, name string) (_ *IPCBucketCreateResult, err error) {
	defer mon.Task()(&ctx, name)(&err)

	if len(name) < minBucketNameLength {
		return nil, errSDK.Errorf("invalid bucket name")
	}

	nameEnc, err := sdk.maybeEncryptMetadata(name, "bucket")
	if err != nil {
		return nil, errSDK.Wrap(err)
	}

	var tx *types.Transaction

	err = sdk.withRetry.Do(ctx, func() (bool, error) {
		tx, err = sdk.ipc.Storage.CreateBucket(sdk.ipc.Auth, nameEnc)
		return isRetryableTxError(err), err
	})
	if err != nil {
		return &IPCBucketCreateResult{}, errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	if err := sdk.ipc.WaitForTx(ctx, tx.Hash()); err != nil {
		return &IPCBucketCreateResult{}, errSDK.Wrap(err)
	}

	bucket, err := sdk.ipc.Storage.GetBucketByName(
		&bind.CallOpts{Context: ctx, From: sdk.ipc.Auth.From},
		nameEnc,
		sdk.ipc.Auth.From,
		big.NewInt(0), big.NewInt(0), // no need to fetch file ids here
	)
	if err != nil {
		return &IPCBucketCreateResult{}, errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	bName, err := sdk.maybeDecryptMetadata(bucket.Name, "bucket")
	if err != nil {
		return &IPCBucketCreateResult{}, errSDK.Wrap(err)
	}

	return &IPCBucketCreateResult{
		ID:        hex.EncodeToString(bucket.Id[:]),
		Name:      bName,
		CreatedAt: time.Unix(bucket.CreatedAt.Int64(), 0),
	}, nil
}

// ViewBucket returns bucket's metadata.
func (sdk *IPC) ViewBucket(ctx context.Context, bucketName string) (_ IPCBucket, err error) {
	defer mon.Task()(&ctx, bucketName)(&err)

	if bucketName == "" {
		return IPCBucket{}, errSDK.Errorf("empty bucket name")
	}

	bucketNameEnc, err := sdk.maybeEncryptMetadata(bucketName, "bucket")
	if err != nil {
		return IPCBucket{}, errSDK.Wrap(err)
	}

	res, err := sdk.client.BucketView(ctx, &pb.IPCBucketViewRequest{
		Name:    bucketNameEnc,
		Address: sdk.ipc.Auth.From.String(),
	})
	if err != nil {
		return IPCBucket{}, errSDK.Wrap(err)
	}

	bName, err := sdk.maybeDecryptMetadata(res.GetName(), "bucket")
	if err != nil {
		return IPCBucket{}, errSDK.Wrap(err)
	}

	return IPCBucket{
		ID:        res.GetId(),
		Name:      bName,
		CreatedAt: res.GetCreatedAt().AsTime(),
	}, nil
}

// ListBuckets returns list of buckets.
func (sdk *IPC) ListBuckets(ctx context.Context, offset, limit int64) (_ []IPCBucket, err error) {
	defer mon.Task()(&ctx, offset, limit)(&err)

	res, err := sdk.client.BucketList(ctx, &pb.IPCBucketListRequest{
		Address: sdk.ipc.Auth.From.String(),
		Offset:  offset,
		Limit:   limit,
	})
	if err != nil {
		return nil, errSDK.Wrap(err)
	}

	buckets := make([]IPCBucket, 0, len(res.Buckets))
	for _, bucket := range res.Buckets {
		var bName string
		bName, err = sdk.maybeDecryptMetadata(bucket.Name, "bucket")
		if err != nil {
			bName = bucket.Name
		}

		buckets = append(buckets, IPCBucket{
			Name:      bName,
			CreatedAt: bucket.GetCreatedAt().AsTime(),
		})
	}

	return buckets, nil
}

// DeleteBucket deletes bucket by name.
func (sdk *IPC) DeleteBucket(ctx context.Context, name string) (err error) {
	defer mon.Task()(&ctx)(&err)

	nameEnc, err := sdk.maybeEncryptMetadata(name, "bucket")
	if err != nil {
		return errSDK.Wrap(err)
	}

	bucket, err := sdk.client.BucketView(ctx, &pb.IPCBucketViewRequest{
		Name:    nameEnc,
		Address: sdk.ipc.Auth.From.String(),
	})
	if err != nil {
		return errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	// TODO: temp solution, when contract will remove bucket id from DeleteBucket method, this code should be removed
	id, err := hex.DecodeString(bucket.Id)
	if err != nil {
		return errSDK.Wrap(err)
	}

	var bucketID [32]byte
	copy(bucketID[:], id)

	bucketIdx, err := sdk.ipc.Storage.GetBucketIndexByName(&bind.CallOpts{Context: ctx}, bucket.Name, sdk.ipc.Auth.From)
	if err != nil {
		return errSDK.Wrap(ipc.ErrorHashToError(err))
	}
	if !bucketIdx.Exists {
		return errSDK.Errorf("bucket does not exist: %s", nameEnc)
	}

	tx, err := sdk.ipc.Storage.DeleteBucket(sdk.ipc.Auth, bucketID, bucket.Name, bucketIdx.Index)
	if err != nil {
		return errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	return errSDK.Wrap(sdk.ipc.WaitForTx(ctx, tx.Hash()))
}

// FileInfo returns meta information for single file by bucket and file name.
func (sdk *IPC) FileInfo(ctx context.Context, bucketName, fileName string) (_ IPCFileMeta, err error) {
	if bucketName == "" {
		return IPCFileMeta{}, errSDK.Errorf("empty bucket name")
	}

	bucketNameEnc, err := sdk.maybeEncryptMetadata(bucketName, "bucket")
	if err != nil {
		return IPCFileMeta{}, errSDK.Wrap(err)
	}
	fileNameEnc, err := sdk.maybeEncryptMetadata(fileName, bucketName)
	if err != nil {
		return IPCFileMeta{}, errSDK.Wrap(err)
	}

	res, err := sdk.client.FileView(ctx, &pb.IPCFileViewRequest{
		FileName:   fileNameEnc,
		BucketName: bucketNameEnc,
		Address:    sdk.ipc.Auth.From.String(),
	})
	if err != nil {
		return IPCFileMeta{}, errSDK.Wrap(err)
	}

	bName, err := sdk.maybeDecryptMetadata(res.GetBucketName(), "bucket")
	if err != nil {
		return IPCFileMeta{}, errSDK.Wrap(err)
	}
	fName, err := sdk.maybeDecryptMetadata(res.GetFileName(), bName)
	if err != nil {
		return IPCFileMeta{}, errSDK.Wrap(err)
	}

	return IPCFileMeta{
		RootCID:     res.GetRootCid(),
		Name:        fName,
		BucketName:  bName,
		IsPublic:    res.GetIsPublic(),
		EncodedSize: res.GetEncodedSize(),
		ActualSize:  res.GetActualSize(),
		CreatedAt:   res.CreatedAt.AsTime(),
	}, nil
}

// ListFiles returns list of files in a particular bucket.
func (sdk *IPC) ListFiles(ctx context.Context, bucketName string, offset, limit int64) (_ []IPCFileListItem, err error) {
	defer mon.Task()(&ctx, bucketName, offset, limit)(&err)

	if bucketName == "" {
		return nil, errSDK.Errorf("empty bucket name")
	}

	bucketNameEnc, err := sdk.maybeEncryptMetadata(bucketName, "bucket")
	if err != nil {
		return []IPCFileListItem{}, errSDK.Wrap(err)
	}

	resp, err := sdk.client.FileList(ctx, &pb.IPCFileListRequest{
		BucketName: bucketNameEnc,
		Address:    sdk.ipc.Auth.From.String(),
		Offset:     offset,
		Limit:      limit,
	})
	if err != nil {
		return nil, errSDK.Wrap(err)
	}

	files := make([]IPCFileListItem, 0, len(resp.List))
	for _, fileMeta := range resp.List {
		fName, err := sdk.maybeDecryptMetadata(fileMeta.Name, bucketName)
		if err != nil {
			return nil, errSDK.Wrap(fmt.Errorf("failed to decrypt metadata: %w", err))
		}

		files = append(files, IPCFileListItem{
			RootCID:     fileMeta.GetRootCid(),
			Name:        fName,
			EncodedSize: fileMeta.GetEncodedSize(),
			ActualSize:  fileMeta.GetActualSize(),
			CreatedAt:   fileMeta.GetCreatedAt().AsTime(),
		})
	}

	return files, nil
}

// FileDelete deletes a file by bucket name and file name.
func (sdk *IPC) FileDelete(ctx context.Context, bucketName, fileName string) (err error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&err)

	if strings.TrimSpace(bucketName) == "" || strings.TrimSpace(fileName) == "" {
		return errSDK.Errorf("empty bucket or file name. Bucket: '%s', File: '%s'", bucketName, fileName)
	}

	bucketNameEnc, err := sdk.maybeEncryptMetadata(bucketName, "bucket")
	if err != nil {
		return errSDK.Wrap(err)
	}
	fileNameEnc, err := sdk.maybeEncryptMetadata(fileName, bucketName)
	if err != nil {
		return errSDK.Wrap(err)
	}

	bucket, err := sdk.ipc.Storage.GetBucketByName(
		&bind.CallOpts{Context: ctx, From: sdk.ipc.Auth.From},
		bucketNameEnc,
		sdk.ipc.Auth.From,
		big.NewInt(0), big.NewInt(0), // no need to fetch file ids here
	)
	if err != nil {
		return errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	file, err := sdk.ipc.Storage.GetFileByName(&bind.CallOpts{Context: ctx}, bucket.Id, fileNameEnc)
	if err != nil {
		return errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	fileIdx, err := sdk.ipc.Storage.GetFileIndexById(&bind.CallOpts{Context: ctx, From: sdk.ipc.Auth.From}, bucket.Name, file.Id, sdk.ipc.Auth.From)
	if err != nil {
		return errSDK.Wrap(ipc.ErrorHashToError(err))
	}
	if !fileIdx.Exists {
		return errSDK.Errorf("file index not exists: %s", fileNameEnc)
	}

	tx, err := sdk.ipc.Storage.DeleteFile(sdk.ipc.Auth, file.Id, bucket.Id, fileNameEnc, fileIdx.Index)
	if err != nil {
		return errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	return errSDK.Wrap(sdk.ipc.WaitForTx(ctx, tx.Hash()))
}

// CreateFileUpload creates a new file upload request.
func (sdk *IPC) CreateFileUpload(ctx context.Context, bucketName, fileName string) (_ *IPCFileUpload, err error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&err)

	if bucketName == "" {
		return nil, errSDK.Errorf("empty bucket name")
	}

	fileNameEnc, err := sdk.maybeEncryptMetadata(fileName, bucketName)
	if err != nil {
		return nil, errSDK.Wrap(err)
	}
	bucketNameEnc, err := sdk.maybeEncryptMetadata(bucketName, "bucket")
	if err != nil {
		return nil, errSDK.Wrap(err)
	}

	// create fileUpload template before interacting with the blockchain
	fileUpload, err := NewIPCFileUpload(bucketNameEnc, fileNameEnc)
	if err != nil {
		return nil, errSDK.Wrap(err)
	}

	var bucket contracts.IStorageBucket

	err = sdk.withRetry.Do(ctx, func() (bool, error) {
		bucket, err = sdk.ipc.Storage.GetBucketByName(
			&bind.CallOpts{Context: ctx, From: sdk.ipc.Auth.From},
			bucketNameEnc,
			sdk.ipc.Auth.From,
			big.NewInt(0), big.NewInt(0), // no need to fetch file ids here
		)
		return true, err
	})
	if err != nil {
		return nil, errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	var tx *types.Transaction

	err = sdk.withRetry.Do(ctx, func() (bool, error) {
		tx, err = sdk.ipc.Storage.CreateFile(sdk.ipc.Auth, bucket.Id, fileNameEnc)
		return isRetryableTxError(err), err
	})
	if err != nil {
		return nil, errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	return fileUpload, errSDK.Wrap(sdk.ipc.WaitForTx(ctx, tx.Hash()))
}

// ChunkData contains chunk data with CIDs and sizes for batching.
type ChunkData struct {
	ChunkUpload IPCFileChunkUploadV2
	CIDs        [][32]byte
	Sizes       []*big.Int
}

// BatchTransaction contains transaction with associated chunks.
type BatchTransaction struct {
	Chunks []IPCFileChunkUploadV2
	Tx     *types.Transaction
}

// Upload uploads a file using ipc api.
func (sdk *IPC) Upload(ctx context.Context, fileUpload *IPCFileUpload, reader io.Reader) (_ IPCFileMetaV2, err error) {
	defer mon.Task()(&ctx, fileUpload)(&err)

	if fileUpload == nil {
		return IPCFileMetaV2{}, errSDK.Errorf("empty file upload")
	}
	if fileUpload.state.isCommitted {
		return IPCFileMetaV2{}, errSDK.Errorf("file is already committed")
	}
	if fileUpload.BucketName == "" {
		return IPCFileMetaV2{}, errSDK.Errorf("empty bucket name")
	}
	if fileUpload.Name == "" {
		return IPCFileMetaV2{}, errSDK.Errorf("empty file name")
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
			big.NewInt(0), big.NewInt(0), // no need to fetch file ids here
		)
		return true, err
	})
	if err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	chunkEncOverhead := 0
	fileEncKey, err := encryptionKey(sdk.encryptionKey, fileUpload.BucketName, fileUpload.Name)
	if err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(err)
	}
	if len(fileEncKey) > 0 {
		chunkEncOverhead = EncryptionOverhead
	}

	bufferSize := sdk.maxBlocksInChunk * int(BlockSize)
	if sdk.ec != nil { // erasure coding enabled
		bufferSize = sdk.ec.DataBlocks * int(BlockSize)
	}
	bufferSize -= chunkEncOverhead

	g, chunkCtx := errgroup.WithContext(ctx)
	fileUploadChunksCh := make(chan IPCFileChunkUploadV2)
	waitTransactionsCh := make(chan BatchTransaction, sdk.chunkBuffer)

	pool := newConnectionPool()
	defer func() {
		if err := pool.close(); err != nil {
			slog.Warn("failed to close connection", slog.String("error", err.Error()))
		}
	}()

	// Start goroutine for reading data and creating chunks
	g.Go(func() error {
		defer close(waitTransactionsCh)

		buf := make([]byte, bufferSize)

		chunkIndex := int64(0)
		if isContinuation {
			if err := skipToPosition(reader, fileUpload.state.actualFileSize); err != nil {
				return err
			}
			chunkIndex = fileUpload.state.chunkCount
		}

		for {
			var batch []ChunkData
			var readErr error

			for range sdk.batchSize {
				var n int
				n, readErr = io.ReadFull(reader, buf)

				if readErr != nil {
					switch {
					case errors.Is(readErr, io.EOF):
						if chunkIndex == 0 {
							return fmt.Errorf("empty file")
						}
					case errors.Is(readErr, io.ErrUnexpectedEOF):
						// do nothing.
					default:
						return readErr
					}
				}

				if n > 0 {
					chunkData := make([]byte, n)
					copy(chunkData, buf[:n])

					chunkUpload, err := sdk.createChunkUpload(
						chunkCtx,
						chunkIndex,
						fileEncKey,
						chunkData,
						bucket.Id,
						fileUpload.Name,
					)
					if err != nil {
						return err
					}

					cids, sizes, _, err := toIPCProtoChunk(
						chunkUpload.ChunkCID.String(),
						chunkUpload.Index,
						chunkUpload.ActualSize,
						chunkUpload.Blocks,
					)
					if err != nil {
						return err
					}

					batch = append(batch, ChunkData{
						ChunkUpload: chunkUpload,
						CIDs:        cids,
						Sizes:       sizes,
					})
					chunkIndex++
				}

				if errors.Is(readErr, io.ErrUnexpectedEOF) || errors.Is(readErr, io.EOF) {
					break
				}
			}

			if len(batch) > 0 {
				tx, err := sdk.createBatchedChunkTransaction(chunkCtx, batch, bucket.Id, fileUpload.Name)
				if err != nil {
					return err
				}

				for _, chunkData := range batch {
					if err := fileUpload.state.preCreateChunk(chunkData.ChunkUpload, tx); err != nil {
						return err
					}
				}

				chunks := make([]IPCFileChunkUploadV2, len(batch))
				for i, chunkData := range batch {
					chunks[i] = chunkData.ChunkUpload
				}

				select {
				case <-chunkCtx.Done():
					return chunkCtx.Err()
				case waitTransactionsCh <- BatchTransaction{Chunks: chunks, Tx: tx}:
				}
			}

			if errors.Is(readErr, io.ErrUnexpectedEOF) || errors.Is(readErr, io.EOF) {
				return nil
			}
		}
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

		// normal processing mode
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

	// Start goroutine for uploading chunks
	g.Go(func() error {
		for {
			select {
			case <-chunkCtx.Done():
				return chunkCtx.Err()
			case chunkUpload, ok := <-fileUploadChunksCh:
				if !ok {
					return nil
				}

				if err := sdk.uploadChunk(chunkCtx, chunkUpload, pool, fileUpload.blocksCounter, fileUpload.bytesCounter, isContinuation); err != nil {
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

		time.Sleep(time.Second) // TODO: make configurable
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
		CommittedAt: time.Now(), // TODO: is it ok to rely on time zone settings of the client?
	}, errSDK.Wrap(sdk.ipc.WaitForTx(ctx, tx.Hash()))
}

// createBatchedChunkTransaction creates AddFileChunks batched transaction from chunk data and sends it.
func (sdk *IPC) createBatchedChunkTransaction(ctx context.Context, batch []ChunkData, bucketID [32]byte, fileName string) (*types.Transaction, error) {
	if len(batch) == 0 {
		return nil, fmt.Errorf("empty batch")
	}

	opts := *sdk.ipc.Auth

	perChunkGas := uint64(270000)
	opts.GasLimit = uint64(len(batch)) * perChunkGas
	opts.GasTipCap = big.NewInt(5e9)
	opts.GasFeeCap = big.NewInt(150e9)

	var (
		cids              = make([][]byte, 0)
		encodedChunkSizes = make([]*big.Int, 0)
		chunkBlocksCIDs   = make([][][32]byte, 0)
		chunkBlockSizes   = make([][]*big.Int, 0)
	)

	startingChunkIndex := batch[0].ChunkUpload.Index

	for _, batchedData := range batch {
		cids = append(cids, batchedData.ChunkUpload.ChunkCID.Bytes())
		encodedChunkSizes = append(encodedChunkSizes, big.NewInt(int64(batchedData.ChunkUpload.EncodedSize)))
		chunkBlocksCIDs = append(chunkBlocksCIDs, batchedData.CIDs)
		chunkBlockSizes = append(chunkBlockSizes, batchedData.Sizes)
	}

	var tx *types.Transaction
	var err error

	err = sdk.withRetry.Do(ctx, func() (bool, error) {
		tx, err = sdk.ipc.Storage.AddFileChunks(
			&opts,
			cids,
			bucketID,
			fileName,
			encodedChunkSizes,
			chunkBlocksCIDs,
			chunkBlockSizes,
			big.NewInt(startingChunkIndex),
		)
		return isRetryableTxError(err), err
	})
	if err != nil {
		return nil, errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	return tx, nil
}

func (sdk *IPC) createChunkUpload(ctx context.Context, index int64, fileEncryptionKey, data []byte, bucketID [32]byte, fileName string) (_ IPCFileChunkUploadV2, err error) {
	defer mon.Task()(&ctx, index)(&err)

	size := int64(len(data))
	if len(fileEncryptionKey) > 0 {
		data, err = encryption.Encrypt(fileEncryptionKey, data, []byte(fmt.Sprintf("%d", index)))
		if err != nil {
			return IPCFileChunkUploadV2{}, errSDK.Wrap(err)
		}
	}

	blockSize := BlockSize.ToInt64()
	if sdk.ec != nil { // erasure coding is enabled
		data, err = sdk.ec.Encode(data)
		if err != nil {
			return IPCFileChunkUploadV2{}, errSDK.Wrap(err)
		}
		// equivalent to notion of shard size in erasure coding terminology
		blockSize = int64(len(data) / (sdk.ec.DataBlocks + sdk.ec.ParityBlocks))
	}

	chunkDAG, err := BuildDAG(ctx, bytes.NewBuffer(data), blockSize)
	if err != nil {
		return IPCFileChunkUploadV2{}, errSDK.Wrap(err)
	}

	_, _, protoChunk, err := toIPCProtoChunk(chunkDAG.CID.String(), index, size, chunkDAG.Blocks)
	if err != nil {
		return IPCFileChunkUploadV2{}, err
	}
	req := &pb.IPCFileUploadChunkCreateRequest{
		Chunk:    protoChunk,
		BucketId: bucketID[:],
		FileName: fileName,
	}

	res, err := sdk.client.FileUploadChunkCreate(ctx, req)
	if err != nil {
		return IPCFileChunkUploadV2{}, errSDK.Wrap(err)
	}

	if len(res.Blocks) != len(chunkDAG.Blocks) {
		return IPCFileChunkUploadV2{}, errSDK.Errorf("received unexpected amount of blocks %d, expected %d", len(res.Blocks), len(chunkDAG.Blocks))
	}
	for i, upload := range res.Blocks {
		if chunkDAG.Blocks[i].CID != upload.Cid {
			return IPCFileChunkUploadV2{}, errSDK.Errorf("block CID mismatch at position %d", i)
		}
		chunkDAG.Blocks[i].NodeAddress = upload.NodeAddress
		chunkDAG.Blocks[i].NodeID = upload.NodeId
		chunkDAG.Blocks[i].Permit = upload.Permit
	}

	return IPCFileChunkUploadV2{
		Index:       index,
		ChunkCID:    chunkDAG.CID,
		ActualSize:  size,
		RawDataSize: chunkDAG.RawDataSize,
		EncodedSize: chunkDAG.EncodedSize,
		Blocks:      chunkDAG.Blocks,
		BucketID:    bucketID,
		FileName:    fileName,
	}, nil
}

func (sdk *IPC) uploadChunk(ctx context.Context, fileChunkUpload IPCFileChunkUploadV2, pool *connectionPool, blockCount, bytesCount *atomic.Int64, isResuming bool) (err error) {
	defer mon.Task()(&ctx, fileChunkUpload)(&err)

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(sdk.maxConcurrency)

	_, _, protoChunk, err := toIPCProtoChunk(
		fileChunkUpload.ChunkCID.String(),
		fileChunkUpload.Index,
		fileChunkUpload.ActualSize,
		fileChunkUpload.Blocks,
	)
	if err != nil {
		return errSDK.Wrap(err)
	}

	privateKeyBytes, err := hex.DecodeString(sdk.privateKey)
	if err != nil {
		return errSDK.Wrap(err)
	}

	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return errSDK.Wrap(err)
	}

	chunkCid, err := cid.Decode(protoChunk.Cid)
	if err != nil {
		return errSDK.Wrap(err)
	}

	for i, block := range fileChunkUpload.Blocks {
		deriveCtx := context.WithoutCancel(ctx)
		g.Go(func() (err error) {
			defer mon.TaskNamed("(*IPC).uploadIPCBlock")(&deriveCtx)(&err)

			timer := time.AfterFunc(30*time.Second, func() {
				slog.Debug("stale block",
					slog.String("node", block.NodeAddress),
					slog.String("cid", block.CID),
					slog.Int("index", i))
			})
			defer timer.Stop()

			client, closer, err := pool.createIPCClient(block.NodeAddress, sdk.useConnectionPool)
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
				return errSDK.Wrap(err)
			}

			nonce, err := ipc.GenerateNonce()
			if err != nil {
				return errSDK.Wrap(err)
			}

			deadline := time.Now().Add(time.Hour * 24).Unix()

			blockCid, err := cid.Decode(block.CID)
			if err != nil {
				return errSDK.Wrap(err)
			}

			var bcid [32]byte
			copy(bcid[:], blockCid.Bytes()[4:])

			nodeID, err := peer.Decode(block.NodeID)
			if err != nil {
				return errSDK.Wrap(err)
			}

			id, err := nodeID.MarshalBinary()
			if err != nil {
				return errSDK.Wrap(err)
			}

			var id32 [32]byte
			copy(id32[:], id[6:])

			storageData := ipc.StorageData{
				ChunkCID:   chunkCid.Bytes(),
				BlockCID:   bcid,
				ChunkIndex: big.NewInt(fileChunkUpload.Index),
				BlockIndex: big.NewInt(int64(i)),
				NodeID:     id32,
				Nonce:      nonce,
				Deadline:   big.NewInt(deadline),
				BucketID:   fileChunkUpload.BucketID,
			}

			signature, err := ipc.SignBlock(privateKey, sdk.storageAddress, sdk.chainID, storageData)
			if err != nil {
				return errSDK.Wrap(err)
			}

			if err := sdk.uploadIPCBlockSegments(ctx, &pb.IPCFileBlockData{
				Data:      block.Data,
				Cid:       block.CID,
				Index:     int64(i),
				Chunk:     protoChunk,
				BucketId:  fileChunkUpload.BucketID[:],
				FileName:  fileChunkUpload.FileName,
				Signature: hex.EncodeToString(signature),
				Nonce:     nonce.Bytes(),
				NodeId:    id,
				Deadline:  deadline,
			}, sender.Send); err != nil && !errors.Is(err, io.EOF) {
				return errSDK.Wrap(err)
			}

			_, closeErr := sender.CloseAndRecv()
			if closeErr == nil {
				blockCount.Add(1)
				bytesCount.Add(int64(len(block.Data)))
				return nil
			}

			// If block is already filled, we can skip it in resumable uploads.
			if strings.Contains(closeErr.Error(), "BlockAlreadyFilled") && isResuming {
				return nil
			}

			return closeErr
		})
	}

	if err := g.Wait(); err != nil {
		return errSDK.Wrap(err)
	}

	return nil
}

func (sdk *IPC) uploadIPCBlockSegments(ctx context.Context, block *pb.IPCFileBlockData, sender func(*pb.IPCFileBlockData) error) (err error) {
	defer mon.Task()(&ctx, block.Cid)(&err)

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

// CreateFileDownload creates a new download request.
func (sdk *IPC) CreateFileDownload(ctx context.Context, bucketName, fileName string) (_ IPCFileDownload, innerErr error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&innerErr)
	if bucketName == "" {
		return IPCFileDownload{}, errSDK.Errorf("empty bucket name")
	}
	if fileName == "" {
		return IPCFileDownload{}, errSDK.Errorf("empty file name")
	}

	bucketNameEnc, err := sdk.maybeEncryptMetadata(bucketName, "bucket")
	if err != nil {
		return IPCFileDownload{}, errSDK.Wrap(err)
	}
	fileNameEnc, err := sdk.maybeEncryptMetadata(fileName, bucketName)
	if err != nil {
		return IPCFileDownload{}, errSDK.Wrap(err)
	}

	res, err := sdk.client.FileDownloadCreate(ctx, &pb.IPCFileDownloadCreateRequest{
		BucketName: bucketNameEnc,
		FileName:   fileNameEnc,
		Address:    sdk.ipc.Auth.From.String(),
	})
	if err != nil {
		return IPCFileDownload{}, errSDK.Wrap(err)
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

	return IPCFileDownload{
		BucketName:   res.BucketName,
		Name:         fileNameEnc,
		Chunks:       chunks,
		dataCounters: newDataCounters(),
	}, nil
}

// CreateRangeFileDownload creates a new download request with block ranges.
func (sdk *IPC) CreateRangeFileDownload(ctx context.Context, bucketName, fileName string, start, end int64) (_ IPCFileDownload, err error) {
	defer mon.Task()(&ctx, bucketName, fileName, start, end)(&err)

	if bucketName == "" {
		return IPCFileDownload{}, errSDK.Errorf("empty bucket name")
	}
	if fileName == "" {
		return IPCFileDownload{}, errSDK.Errorf("empty file name")
	}

	bucketNameEnc, err := sdk.maybeEncryptMetadata(bucketName, "bucket")
	if err != nil {
		return IPCFileDownload{}, errSDK.Wrap(err)
	}
	fileNameEnc, err := sdk.maybeEncryptMetadata(fileName, bucketName)
	if err != nil {
		return IPCFileDownload{}, errSDK.Wrap(err)
	}

	res, err := sdk.client.FileDownloadRangeCreate(ctx, &pb.IPCFileDownloadRangeCreateRequest{
		BucketName: bucketNameEnc,
		FileName:   fileNameEnc,
		Address:    sdk.ipc.Auth.From.String(),
		StartIndex: start,
		EndIndex:   end,
	})
	if err != nil {
		return IPCFileDownload{}, errSDK.Wrap(err)
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

	return IPCFileDownload{
		BucketName:   res.BucketName,
		Name:         fileNameEnc,
		Chunks:       chunks,
		dataCounters: newDataCounters(),
	}, nil
}

// FileSetPublicAccess change file status from/to public.
func (sdk *IPC) FileSetPublicAccess(ctx context.Context, bucketName, fileName string, isPublic bool) (err error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&err)

	bucketNameEnc, err := sdk.maybeEncryptMetadata(bucketName, "bucket")
	if err != nil {
		return errSDK.Wrap(err)
	}
	fileNameEnc, err := sdk.maybeEncryptMetadata(fileName, bucketName)
	if err != nil {
		return errSDK.Wrap(err)
	}

	bucket, err := sdk.ipc.Storage.GetBucketByName(
		&bind.CallOpts{Context: ctx, From: sdk.ipc.Auth.From},
		bucketNameEnc,
		sdk.ipc.Auth.From,
		big.NewInt(0), big.NewInt(0), // no need to fetch file ids here
	)
	if err != nil {
		return errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	file, err := sdk.ipc.Storage.GetFileByName(&bind.CallOpts{Context: ctx}, bucket.Id, fileNameEnc)
	if err != nil {
		return errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	tx, err := sdk.ipc.AccessManager.ChangePublicAccess(sdk.ipc.Auth, file.Id, isPublic)
	if err != nil {
		return errSDK.Wrap(err)
	}

	return errSDK.Wrap(sdk.ipc.WaitForTx(ctx, tx.Hash()))
}

// Download downloads a file using ipc api.
func (sdk *IPC) Download(ctx context.Context, fileDownload IPCFileDownload, writer io.Writer) (err error) {
	defer mon.Task()(&ctx, fileDownload)(&err)

	fileEncKey, err := encryptionKey(sdk.encryptionKey, fileDownload.BucketName, fileDownload.Name)
	if err != nil {
		return errSDK.Wrap(err)
	}

	pool := newConnectionPool()
	defer func() {
		if err := pool.close(); err != nil {
			slog.Warn("failed to close connection", slog.String("error", err.Error()))
		}
	}()

	g, ctx := errgroup.WithContext(ctx)
	chunkDownloadCh := make(chan FileChunkDownload, sdk.chunkBuffer)

	// Start goroutine to create chunk downloads
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

	// Start goroutine to process chunk downloads
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case chunkDownload, ok := <-chunkDownloadCh:
				if !ok {
					return nil
				}

				err = sdk.downloadChunkBlocks(
					ctx,
					pool,
					fileDownload.BucketName,
					fileDownload.Name,
					sdk.ipc.Auth.From.String(),
					chunkDownload,
					fileEncKey,
					writer,
					fileDownload.blocksCounter,
					fileDownload.bytesCounter)
				if err != nil {
					return err
				}

				fileDownload.chunksCounter.Add(1)
			}
		}
	})

	return g.Wait()
}

// DownloadArchival downloads a file from PDP provider.
func (sdk *IPC) DownloadArchival(ctx context.Context, fileDownload IPCFileDownload, writer io.Writer) (err error) {
	defer mon.Task()(&ctx, fileDownload)(&err)

	fileEncKey, err := encryptionKey(sdk.encryptionKey, fileDownload.BucketName, fileDownload.Name)
	if err != nil {
		return errSDK.Wrap(err)
	}

	pool := newConnectionPool()
	defer func() {
		if err := pool.close(); err != nil {
			slog.Warn("failed to close connection", slog.String("error", err.Error()))
		}
	}()

	g, ctx := errgroup.WithContext(ctx)
	chunkDownloadCh := make(chan FileChunkDownload, sdk.chunkBuffer)

	// Start goroutine to create chunk downloads
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

	// Start goroutine to process chunk downloads
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case chunkDownload, ok := <-chunkDownloadCh:
				if !ok {
					return nil
				}

				err = sdk.downloadChunkBlocks2(
					ctx,
					pool,
					chunkDownload,
					fileEncKey,
					writer,
					fileDownload.blocksCounter,
					fileDownload.bytesCounter)
				if err != nil {
					return err
				}

				fileDownload.chunksCounter.Add(1)
			}
		}
	})

	return g.Wait()
}

// ArchivalMetadata returns file metadata with chunks and blocks including PDP data.
func (sdk *IPC) ArchivalMetadata(ctx context.Context, bucketName, fileName string) (_ ArchivalMetadata, err error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&err)

	fileNameEnc, err := sdk.maybeEncryptMetadata(fileName, bucketName)
	if err != nil {
		return ArchivalMetadata{}, errSDK.Wrap(err)
	}
	bucketNameEnc, err := sdk.maybeEncryptMetadata(bucketName, "bucket")
	if err != nil {
		return ArchivalMetadata{}, errSDK.Wrap(err)
	}

	res, err := sdk.client.FileDownloadCreate(ctx, &pb.IPCFileDownloadCreateRequest{
		BucketName: bucketNameEnc,
		FileName:   fileNameEnc,
		Address:    sdk.ipc.Auth.From.String(),
	})
	if err != nil {
		return ArchivalMetadata{}, errSDK.Wrap(err)
	}

	result := ArchivalMetadata{
		BucketName: bucketNameEnc,
		Name:       fileNameEnc,
		Chunks:     make([]ArchivalChunk, 0, len(res.Chunks)),
	}

	pool := newConnectionPool()
	defer func() {
		if closeErr := pool.close(); closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}()

	type resolvedBlock struct {
		Pos     int
		CID     string
		Size    int64
		PDPData *PDPBlockData
	}

	for i, chunk := range res.Chunks {
		chunkRes, err := sdk.client.FileDownloadChunkCreate(ctx, &pb.IPCFileDownloadChunkCreateRequest{
			BucketName: bucketNameEnc,
			FileName:   fileNameEnc,
			ChunkCid:   chunk.Cid,
			ChunkIndex: int64(i),
			Address:    sdk.ipc.Auth.From.String(),
		})
		if err != nil {
			return ArchivalMetadata{}, errSDK.Wrap(err)
		}

		archivalChunk := ArchivalChunk{
			Chunk: Chunk{
				CID:         chunk.Cid,
				EncodedSize: chunk.EncodedSize,
				Size:        chunk.Size,
				Index:       int64(i),
			},
			Blocks: make([]ArchivalBlock, len(chunkRes.Blocks)),
		}

		g, gctx := errgroup.WithContext(ctx)
		g.SetLimit(sdk.maxConcurrency)
		ch := make(chan resolvedBlock, len(chunkRes.Blocks))

		for idx, block := range chunkRes.Blocks {
			rb := resolvedBlock{
				Pos:  idx,
				CID:  block.Cid,
				Size: block.Size,
			}

			g.Go(func() error {
				pdpBlock, err := sdk.resolveBlock(gctx, pool, FileBlockDownload{
					CID:         block.Cid,
					NodeAddress: block.NodeAddress,
				})

				if err != nil {
					if !errors.As(err, &ErrMissingArchivalBlock{}) {
						return err
					}
				} else {
					// if resolved, set PDP data
					rb.PDPData = &pdpBlock
				}

				ch <- rb

				return nil
			})
		}

		if err := g.Wait(); err != nil {
			close(ch)
			return ArchivalMetadata{}, errSDK.Wrap(err)
		}
		close(ch)

		for resolved := range ch {
			archivalChunk.Blocks[resolved.Pos] = ArchivalBlock{
				CID:     resolved.CID,
				Size:    resolved.Size,
				PDPData: resolved.PDPData,
			}
		}

		result.Chunks = append(result.Chunks, archivalChunk)
	}

	return result, nil
}

func (sdk *IPC) createChunkDownload(ctx context.Context, bucketName, fileName string, chunk Chunk) (_ FileChunkDownload, err error) {
	defer mon.Task()(&ctx, chunk)(&err)

	res, err := sdk.client.FileDownloadChunkCreate(ctx, &pb.IPCFileDownloadChunkCreateRequest{
		BucketName: bucketName,
		FileName:   fileName,
		ChunkCid:   chunk.CID,
		ChunkIndex: chunk.Index,
		Address:    sdk.ipc.Auth.From.String(),
	})
	if err != nil {
		return FileChunkDownload{}, errSDK.Wrap(err)
	}

	blocks := make([]FileBlockDownload, len(res.Blocks))
	for i, block := range res.Blocks {
		blocks[i] = FileBlockDownload{
			CID:         block.Cid,
			NodeID:      block.NodeId,
			NodeAddress: block.NodeAddress,
			Permit:      block.Permit,
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

func (sdk *IPC) downloadChunkBlocks(
	ctx context.Context,
	pool *connectionPool,
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
			defer mon.TaskNamed("(*IPC).downloadBlock")(&deriveCtx, block.CID)(&err)

			blockData, err := sdk.fetchBlockData(ctx, pool, chunkDownload.CID, bucketName, fileName, address, chunkDownload.Index, int64(i), block)
			if err != nil {
				return err
			}

			ch <- retrievedBlock{
				Pos:  i,
				CID:  block.CID,
				Data: blockData,
			}
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

	var data []byte
	if sdk.ec != nil { // erasure coding is enabled
		data, err = sdk.ec.ExtractData(blocks, 0)
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

// downloadChunkBlocks2 resolves all chunk block references and download blocks strictly from PDP provider.
func (sdk *IPC) downloadChunkBlocks2(
	ctx context.Context,
	pool *connectionPool,
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
			defer mon.TaskNamed("(*IPC).downloadBlock2")(&deriveCtx, block.CID)(&err)

			pdpBlock, err := sdk.resolveBlock(ctx, pool, block)
			if err != nil {
				return err
			}

			blockData, err := httpext.RangeDownload(ctx,
				sdk.httpClient,
				pdpBlock.URL,
				pdpBlock.Offset,
				pdpBlock.Size,
			)
			if err != nil {
				return err
			}

			// Verify that the downloaded data matches the expected CID.
			if err := cids.VerifyRaw(block.CID, blockData); err != nil {
				return fmt.Errorf("block CID verification failed: %w", err)
			}

			ch <- retrievedBlock{
				Pos:  i,
				CID:  block.CID,
				Data: blockData,
			}
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

	var data []byte
	if sdk.ec != nil { // erasure coding is enabled
		data, err = sdk.ec.ExtractData(blocks, 0)
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

// resolveBlock resolves block metadata from PDP provider.
func (sdk *IPC) resolveBlock(ctx context.Context, pool *connectionPool, block FileBlockDownload) (_ PDPBlockData, err error) {
	defer mon.Task()(&ctx, block.CID)(&err)

	client, closer, err := pool.createArchivalClient(block.NodeAddress, sdk.useConnectionPool)
	if err != nil {
		return PDPBlockData{}, err
	}
	if closer != nil {
		defer func() {
			if closeErr := closer(); closeErr != nil {
				slog.Warn("failed to close connection",
					slog.String("block_cid", block.CID),
					slog.String("error", closeErr.Error()))
			}
		}()
	}
	resp, err := client.FileResolveBlock(ctx, &pb.IPCFileResolveBlockRequest{BlockCid: block.CID})
	if err != nil {
		if st, ok := status.FromError(err); ok && st.Code() == codes.NotFound {
			return PDPBlockData{}, ErrMissingArchivalBlock{BlockCID: block.CID}
		}
		return PDPBlockData{}, err
	}

	return PDPBlockData{
		URL:       resp.Block.GetUrl(),
		Offset:    resp.Block.GetOffset(),
		Size:      resp.Block.GetSize(),
		DataSetID: resp.Block.GetDataSetId(),
	}, nil
}

func (sdk *IPC) fetchBlockData(
	ctx context.Context,
	pool *connectionPool,
	chunkCID, bucketName, fileName, address string,
	chunkIndex, blockIdndex int64,
	block FileBlockDownload,
) ([]byte, error) {

	if block.NodeAddress == "" {
		return nil, errMissingBlockMetadata
	}

	client, closer, err := pool.createIPCClient(block.NodeAddress, sdk.useConnectionPool)
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

	downloadClient, err := client.FileDownloadBlock(ctx, &pb.IPCFileDownloadBlockRequest{
		ChunkCid:   chunkCID,
		ChunkIndex: chunkIndex,
		BlockCid:   block.CID,
		BlockIndex: blockIdndex,
		BucketName: bucketName,
		FileName:   fileName,
		Address:    address,
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

// maybeEncryptMetadata encrypts the given metadata if metadata encryption is enabled and encryption key is set.
func (sdk *IPC) maybeEncryptMetadata(value, derivationPath string) (string, error) {
	if len(sdk.encryptionKey) > 0 && sdk.useMetadataEncryption {
		encrypted, err := encryption.EncryptD(sdk.encryptionKey, []byte(value), []byte(derivationPath))
		if err != nil {
			return "", err
		}
		return hex.EncodeToString(encrypted), nil
	}

	return value, nil
}

// maybeDecryptMetadata decrypts the given metadata if metadata encryption is enabled and encryption key is set.
func (sdk *IPC) maybeDecryptMetadata(value, derivationPath string) (string, error) {
	if len(sdk.encryptionKey) > 0 && sdk.useMetadataEncryption {
		encrypted, err := hex.DecodeString(value)
		if err != nil {
			return "", errSDK.Wrap(err)
		}

		decrypted, err := encryption.Decrypt(sdk.encryptionKey, encrypted, []byte(derivationPath))
		if err != nil {
			return "", errSDK.Wrap(err)
		}

		return string(decrypted), nil
	}

	return value, nil
}

func toIPCProtoChunk(chunkCid string, index, size int64, blocks []FileBlockUpload) ([][32]byte, []*big.Int, *pb.IPCChunk, error) {
	var (
		cids  = make([][32]byte, 0)
		sizes = make([]*big.Int, 0)
	)
	pbBlocks := make([]*pb.IPCChunk_Block, len(blocks))
	for i, block := range blocks {
		pbBlocks[i] = &pb.IPCChunk_Block{
			Cid:  block.CID,
			Size: int64(len(block.Data)),
		}
		var bcid [32]byte
		c, err := cid.Decode(block.CID)
		if err != nil {
			return nil, nil, nil, errSDK.Wrap(err)
		}

		copy(bcid[:], c.Bytes()[4:])
		cids = append(cids, bcid)
		sizes = append(sizes, big.NewInt(int64(len(block.Data))))
	}
	return cids, sizes, &pb.IPCChunk{
		Cid:    chunkCid,
		Index:  index,
		Size:   size,
		Blocks: pbBlocks,
	}, nil
}
