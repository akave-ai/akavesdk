// Copyright (C) 2026 Akave
// See LICENSE for copying information.

package sdk

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/sync/errgroup"

	"github.com/akave-ai/akavesdk/private/ipc"
	"github.com/akave-ai/akavesdk/private/ipc/contracts"
)

// CreateFileUploadParam contains parameters for file upload creation.
type CreateFileUploadParam struct {
	BucketName string
	FileName   string
}

// UploadParam contains parameters for file upload execution.
type UploadParam struct {
	FileUpload *IPCFileUpload // tracks state of the upload, changes during upload
	Reader     io.Reader
}

// CreateFileUploadResult represents the result of a single file upload creation.
type CreateFileUploadResult struct {
	BucketName string
	FileName   string
	FileUpload *IPCFileUpload
	Error      error
}

// UploadResult represents the result of a single file upload.
type UploadResult struct {
	BucketName string
	FileName   string
	Meta       *IPCFileMetaV2
	Error      error
}

// MultiUpload provides batched/parallel file upload operations.
//
// Implementation mostly copied from SDK.Upload and SDK.CreateFileUpload, but adapted for multiple files.
// TODO: consider using multiupload in SDK itself to reduce code duplication.
type MultiUpload struct {
	mu              sync.Mutex
	ipcSDK          *IPC
	fileConcurrency int
}

// CreateFileUploads creates file upload requests on-chain for the given parameters.
func (m *MultiUpload) CreateFileUploads(ctx context.Context, params []CreateFileUploadParam) (_ []CreateFileUploadResult, err error) {
	defer mon.Task()(&ctx)(&err)

	if len(params) == 0 {
		return nil, errSDK.Errorf("no files to upload")
	}

	var errs []error
	for i, p := range params {
		if p.BucketName == "" {
			errs = append(errs, errSDK.Errorf("empty bucket name at index %d", i))
		}
		if p.FileName == "" {
			errs = append(errs, errSDK.Errorf("empty file name at index %d", i))
		}
	}
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	type uploadResult struct {
		index  int
		upload *IPCFileUpload
		err    error
	}

	resultsCh := make(chan uploadResult, len(params))
	g, gctx := errgroup.WithContext(ctx)
	g.SetLimit(m.fileConcurrency)

	for i, p := range params {
		g.Go(func() error {
			var result uploadResult
			result.index = i

			bucketNameEnc, errBucket := m.ipcSDK.maybeEncryptMetadata(p.BucketName, "bucket")
			fileNameEnc, errFile := m.ipcSDK.maybeEncryptMetadata(p.FileName, p.BucketName)
			if joinErr := errors.Join(errBucket, errFile); joinErr != nil {
				result.err = errSDK.Wrap(joinErr)
				resultsCh <- result
				return nil
			}

			fileUpload, fileUploadErr := NewIPCFileUpload(bucketNameEnc, fileNameEnc)
			if fileUploadErr != nil {
				result.err = errSDK.Wrap(fileUploadErr)
				resultsCh <- result
				return nil
			}

			var bucket contracts.IStorageBucket
			bucketGetErr := m.ipcSDK.withRetry.Do(gctx, func() (bool, error) {
				bucket, err = m.ipcSDK.ipc.Storage.GetBucketByName(
					&bind.CallOpts{Context: gctx, From: m.ipcSDK.ipc.Auth.From},
					bucketNameEnc,
					m.ipcSDK.ipc.Auth.From,
					big.NewInt(0), big.NewInt(0),
				)
				return true, err
			})
			if bucketGetErr != nil {
				result.err = errSDK.Wrap(ipc.ErrorHashToError(bucketGetErr))
				resultsCh <- result
				return nil
			}

			var tx *types.Transaction
			txErr := m.ipcSDK.withRetry.Do(gctx, func() (bool, error) {
				m.mu.Lock()
				tx, err = m.ipcSDK.ipc.Storage.CreateFile(m.ipcSDK.ipc.Auth, bucket.Id, fileNameEnc)
				m.mu.Unlock()
				return isRetryableTxError(err), err
			})
			if txErr != nil {
				result.err = errSDK.Wrap(ipc.ErrorHashToError(txErr))
				resultsCh <- result
				return nil
			}

			if waitErr := m.ipcSDK.ipc.WaitForTx(gctx, tx.Hash()); waitErr != nil {
				result.err = errSDK.Wrap(waitErr)
				resultsCh <- result
				return nil
			}

			result.upload = fileUpload
			resultsCh <- result
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		close(resultsCh)
		return nil, err
	}

	close(resultsCh)

	results := make([]CreateFileUploadResult, len(params))
	for result := range resultsCh {
		results[result.index] = CreateFileUploadResult{
			BucketName: params[result.index].BucketName,
			FileName:   params[result.index].FileName,
			FileUpload: result.upload,
			Error:      result.err,
		}
	}

	return results, nil
}

// Upload uploads all files in parallel.
func (m *MultiUpload) Upload(ctx context.Context, params []UploadParam) (_ []UploadResult, err error) {
	defer mon.Task()(&ctx)(&err)

	if len(params) == 0 {
		return nil, errSDK.Errorf("no files to upload")
	}

	var errs []error
	for i, p := range params {
		if p.FileUpload == nil {
			errs = append(errs, errSDK.Errorf("nil FileUpload at index %d", i))
		}
		if p.Reader == nil {
			errs = append(errs, errSDK.Errorf("nil Reader at index %d", i))
		}
	}
	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	type uploadResult struct {
		index int
		meta  *IPCFileMetaV2
		err   error
	}

	resultsCh := make(chan uploadResult, len(params))
	g, gctx := errgroup.WithContext(ctx)
	g.SetLimit(m.fileConcurrency)

	for i, p := range params {
		g.Go(func() error {
			var result uploadResult
			result.index = i

			meta, singleErr := m.uploadSingleFile(gctx, p.FileUpload, p.Reader)
			if singleErr != nil {
				result.err = singleErr
				resultsCh <- result
				return nil //nolint:nilerr
			}

			result.meta = &meta
			resultsCh <- result
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		close(resultsCh)
		return nil, err
	}

	close(resultsCh)

	results := make([]UploadResult, len(params))
	for result := range resultsCh {
		fileUpload := params[result.index].FileUpload // safe to access, validated earlier
		results[result.index] = UploadResult{
			BucketName: fileUpload.BucketName,
			FileName:   fileUpload.Name,
			Meta:       result.meta,
			Error:      result.err,
		}
	}

	return results, nil
}

func (m *MultiUpload) uploadSingleFile(ctx context.Context, fileUpload *IPCFileUpload, reader io.Reader) (_ IPCFileMetaV2, err error) {
	if fileUpload.state.isCommitted {
		return IPCFileMetaV2{}, errSDK.Errorf("file is already committed")
	}

	var isContinuation bool
	if fileUpload.state.chunkCount > 0 {
		isContinuation = true
	}

	var bucket contracts.IStorageBucket
	err = m.ipcSDK.withRetry.Do(ctx, func() (bool, error) {
		bucket, err = m.ipcSDK.ipc.Storage.GetBucketByName(
			&bind.CallOpts{Context: ctx, From: m.ipcSDK.ipc.Auth.From},
			fileUpload.BucketName,
			m.ipcSDK.ipc.Auth.From,
			big.NewInt(0), big.NewInt(0), // no need to fetch file ids here
		)
		return true, err
	})
	if err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	fileEncKey, err := encryptionKey(m.ipcSDK.encryptionKey, fileUpload.BucketName, fileUpload.Name)
	if err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(err)
	}

	g, chunkCtx := errgroup.WithContext(ctx)
	fileUploadChunksCh := make(chan IPCFileChunkUploadV2)
	waitTransactionsCh := make(chan BatchTransaction, m.ipcSDK.chunkBuffer)

	// Start goroutine for reading data and creating chunks
	g.Go(func() error {
		defer close(waitTransactionsCh)

		chunkIndex := int64(0)
		if isContinuation {
			if err := skipToPosition(reader, fileUpload.state.actualFileSize); err != nil {
				return err
			}
			chunkIndex = fileUpload.state.chunkCount
		}

		bufferSize, bufferCapacity := m.ipcSDK.calculateBufferSizeAndCapacity(fileEncKey)
		buf := make([]byte, bufferSize, bufferCapacity)

		stopReading := false
		for !stopReading {
			var batch []ChunkData
			for range m.ipcSDK.chunkBatchSize {
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

				chunkUpload, err := m.ipcSDK.createChunkUpload(
					chunkCtx,
					chunkIndex,
					fileEncKey,
					buf[:n],
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

			if len(batch) > 0 {
				m.mu.Lock()
				tx, err := m.ipcSDK.createBatchedChunkTransaction(chunkCtx, batch, bucket.Id, fileUpload.Name)
				m.mu.Unlock()
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
		}
		return nil
	})

	g.Go(func() error {
		defer close(fileUploadChunksCh)

		if isContinuation {
			for _, chunkWithTx := range fileUpload.state.listPreCreatedChunks() {
				if err := m.ipcSDK.ipc.WaitForTx(chunkCtx, chunkWithTx.tx.Hash()); err != nil {
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

				if err := m.ipcSDK.ipc.WaitForTx(chunkCtx, batchResult.Tx.Hash()); err != nil {
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

				if err := m.ipcSDK.uploadChunk(chunkCtx, chunkUpload, fileUpload.blocksCounter, fileUpload.bytesCounter, isContinuation); err != nil {
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

	err = m.ipcSDK.withRetry.Do(ctx, func() (bool, error) {
		fileMeta, err = m.ipcSDK.ipc.Storage.GetFileByName(
			&bind.CallOpts{Context: ctx, From: m.ipcSDK.ipc.Auth.From},
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
		err = m.ipcSDK.withRetry.Do(ctx, func() (bool, error) {
			isFilled, err = m.ipcSDK.ipc.Storage.IsFileFilled(&bind.CallOpts{Context: ctx}, fileID)
			return true, err
		})
		if err != nil {
			return IPCFileMetaV2{}, errSDK.Wrap(ipc.ErrorHashToError(err))
		}

		time.Sleep(time.Second) // TODO: make configurable
	}

	var tx *types.Transaction
	err = m.ipcSDK.withRetry.Do(ctx, func() (bool, error) {
		m.mu.Lock()
		tx, err = m.ipcSDK.ipc.Storage.CommitFile(
			m.ipcSDK.ipc.Auth,
			bucket.Id,
			fileUpload.Name,
			big.NewInt(fileUpload.state.encodedFileSize),
			big.NewInt(fileUpload.state.actualFileSize),
			rootCID.Bytes(),
		)
		m.mu.Unlock()
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
	}, errSDK.Wrap(m.ipcSDK.ipc.WaitForTx(ctx, tx.Hash()))
}
