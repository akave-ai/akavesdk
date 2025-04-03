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
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p/core/peer"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"github.com/akave-ai/akavesdk/private/eip712"
	"github.com/akave-ai/akavesdk/private/encryption"
	"github.com/akave-ai/akavesdk/private/ipc"
	"github.com/akave-ai/akavesdk/private/pb"
)

// IPC exposes SDK ipc API.
type IPC struct {
	client      pb.IPCNodeAPIClient
	conn        *grpc.ClientConn
	ipc         *ipc.Client
	erasureCode *ErasureCode

	privateKey            string
	chainID               *big.Int
	storageAddress        string
	maxConcurrency        int
	blockPartSize         int64
	useConnectionPool     bool
	encryptionKey         []byte // empty means no encryption
	maxBlocksInChunk      int
	useMetadataEncryption bool
}

// CreateBucket creates a new bucket.
func (sdk *IPC) CreateBucket(ctx context.Context, name string) (_ *IPCBucketCreateResult, err error) {
	defer mon.Task()(&ctx, name)(&err)

	if len(name) < minBucketNameLength {
		return nil, errSDK.Errorf("invalid bucket name")
	}

	name, err = sdk.maybeEncryptMetadata(name, name)
	if err != nil {
		return nil, errSDK.Wrap(err)
	}

	tx, err := sdk.ipc.Storage.CreateBucket(sdk.ipc.Auth, name)
	if err != nil {
		return &IPCBucketCreateResult{}, errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	if err := sdk.ipc.WaitForTx(ctx, tx.Hash()); err != nil {
		return &IPCBucketCreateResult{}, errSDK.Wrap(err)
	}

	bucket, err := sdk.ipc.Storage.GetBucketByName(&bind.CallOpts{From: sdk.ipc.Auth.From}, name)
	if err != nil {
		return &IPCBucketCreateResult{}, errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	return &IPCBucketCreateResult{
		ID:        hex.EncodeToString(bucket.Id[:]),
		Name:      bucket.Name,
		CreatedAt: time.Unix(bucket.CreatedAt.Int64(), 0),
	}, nil
}

// ViewBucket returns bucket's metadata.
func (sdk *IPC) ViewBucket(ctx context.Context, bucketName string) (_ IPCBucket, err error) {
	defer mon.Task()(&ctx, bucketName)(&err)

	if bucketName == "" {
		return IPCBucket{}, errSDK.Errorf("empty bucket name")
	}

	bucketName, err = sdk.maybeEncryptMetadata(bucketName, bucketName)
	if err != nil {
		return IPCBucket{}, errSDK.Wrap(err)
	}

	res, err := sdk.client.BucketView(ctx, &pb.IPCBucketViewRequest{
		Name:    bucketName,
		Address: sdk.ipc.Auth.From.String(),
	})
	if err != nil {
		return IPCBucket{}, errSDK.Wrap(err)
	}

	return IPCBucket{
		ID:        res.GetId(),
		Name:      res.GetName(),
		CreatedAt: res.GetCreatedAt().AsTime(),
	}, nil
}

// ListBuckets returns list of buckets.
func (sdk *IPC) ListBuckets(ctx context.Context) (_ []IPCBucket, err error) {
	defer mon.Task()(&ctx)(&err)

	res, err := sdk.client.BucketList(ctx, &pb.IPCBucketListRequest{
		Address: sdk.ipc.Auth.From.String(),
	})
	if err != nil {
		return nil, errSDK.Wrap(err)
	}

	buckets := make([]IPCBucket, 0, len(res.Buckets))
	for _, bucket := range res.Buckets {
		buckets = append(buckets, IPCBucket{
			Name:      bucket.GetName(),
			CreatedAt: bucket.GetCreatedAt().AsTime(),
		})
	}

	return buckets, nil
}

// DeleteBucket deletes bucket by name.
func (sdk *IPC) DeleteBucket(ctx context.Context, name string) (err error) {
	defer mon.Task()(&ctx)(&err)

	name, err = sdk.maybeEncryptMetadata(name, name)
	if err != nil {
		return errSDK.Wrap(err)
	}

	bucket, err := sdk.client.BucketView(ctx, &pb.IPCBucketViewRequest{
		Name:    name,
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

	bucketIdx, err := sdk.ipc.Storage.GetBucketIndexByName(&bind.CallOpts{}, bucket.Name, sdk.ipc.Auth.From)
	if err != nil {
		return errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	tx, err := sdk.ipc.Storage.DeleteBucket(sdk.ipc.Auth, bucketID, bucket.Name, bucketIdx)
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

	fileName, err = sdk.maybeEncryptMetadata(fileName, bucketName+"/"+fileName)
	if err != nil {
		return IPCFileMeta{}, errSDK.Wrap(err)
	}
	bucketName, err = sdk.maybeEncryptMetadata(bucketName, bucketName)
	if err != nil {
		return IPCFileMeta{}, errSDK.Wrap(err)
	}

	res, err := sdk.client.FileView(ctx, &pb.IPCFileViewRequest{
		FileName:   fileName,
		BucketName: bucketName,
		Address:    sdk.ipc.Auth.From.String(),
	})
	if err != nil {
		return IPCFileMeta{}, errSDK.Wrap(err)
	}

	return IPCFileMeta{
		RootCID:     res.GetRootCid(),
		Name:        res.GetFileName(),
		BucketName:  res.GetBucketName(),
		IsPublic:    res.GetIsPublic(),
		EncodedSize: res.GetEncodedSize(),
		CreatedAt:   res.CreatedAt.AsTime(),
	}, nil
}

// ListFiles returns list of files in a particular bucket.
func (sdk *IPC) ListFiles(ctx context.Context, bucketName string) (_ []IPCFileListItem, err error) {
	defer mon.Task()(&ctx, bucketName)(&err)

	if bucketName == "" {
		return nil, errSDK.Errorf("empty bucket name")
	}

	bucketName, err = sdk.maybeEncryptMetadata(bucketName, bucketName)
	if err != nil {
		return []IPCFileListItem{}, errSDK.Wrap(err)
	}

	resp, err := sdk.client.FileList(ctx, &pb.IPCFileListRequest{
		BucketName: bucketName,
		Address:    sdk.ipc.Auth.From.String(),
	})
	if err != nil {
		return nil, errSDK.Wrap(err)
	}

	files := make([]IPCFileListItem, 0, len(resp.List))
	for _, fileMeta := range resp.List {
		files = append(files, IPCFileListItem{
			RootCID:     fileMeta.GetRootCid(),
			Name:        fileMeta.GetName(),
			EncodedSize: fileMeta.GetEncodedSize(),
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

	fileName, err = sdk.maybeEncryptMetadata(fileName, bucketName+"/"+fileName)
	if err != nil {
		return errSDK.Wrap(err)
	}
	bucketName, err = sdk.maybeEncryptMetadata(bucketName, bucketName)
	if err != nil {
		return errSDK.Wrap(err)
	}

	bucket, err := sdk.ipc.Storage.GetBucketByName(&bind.CallOpts{From: sdk.ipc.Auth.From}, bucketName)
	if err != nil {
		return errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	file, err := sdk.ipc.Storage.GetFileByName(&bind.CallOpts{}, bucket.Id, fileName)
	if err != nil {
		return errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	fileIdx, err := sdk.ipc.Storage.GetFileIndexById(&bind.CallOpts{}, fileName, bucket.Id)
	if err != nil {
		return errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	tx, err := sdk.ipc.Storage.DeleteFile(sdk.ipc.Auth, file.Id, bucket.Id, fileName, fileIdx)
	if err != nil {
		return errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	return errSDK.Wrap(sdk.ipc.WaitForTx(ctx, tx.Hash()))
}

// CreateFileUpload creates a new file upload request.
func (sdk *IPC) CreateFileUpload(ctx context.Context, bucketName, fileName string) (err error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&err)
	if bucketName == "" {
		return errSDK.Errorf("empty bucket name")
	}

	fileName, err = sdk.maybeEncryptMetadata(fileName, bucketName+"/"+fileName)
	if err != nil {
		return errSDK.Wrap(err)
	}
	bucketName, err = sdk.maybeEncryptMetadata(bucketName, bucketName)
	if err != nil {
		return errSDK.Wrap(err)
	}

	bucket, err := sdk.ipc.Storage.GetBucketByName(&bind.CallOpts{From: sdk.ipc.Auth.From}, bucketName)
	if err != nil {
		return errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	tx, err := sdk.ipc.Storage.CreateFile(sdk.ipc.Auth, bucket.Id, fileName)
	if err != nil {
		return errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	return errSDK.Wrap(sdk.ipc.WaitForTx(ctx, tx.Hash()))
}

// Upload uploads a file using ipc api.
func (sdk *IPC) Upload(ctx context.Context, bucketName, fileName string, reader io.Reader) (_ IPCFileMetaV2, err error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&err)

	if bucketName == "" {
		return IPCFileMetaV2{}, errSDK.Errorf("empty bucket name")
	}
	if fileName == "" {
		return IPCFileMetaV2{}, errSDK.Errorf("empty file name")
	}

	fileName, err = sdk.maybeEncryptMetadata(fileName, bucketName+"/"+fileName)
	if err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(err)
	}
	bucketName, err = sdk.maybeEncryptMetadata(bucketName, bucketName)
	if err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(err)
	}

	bucket, err := sdk.ipc.Storage.GetBucketByName(&bind.CallOpts{From: sdk.ipc.Auth.From}, bucketName)
	if err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	chunkEncOverhead := 0
	fileEncKey, err := encryptionKey(sdk.encryptionKey, bucketName, fileName)
	if err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(err)
	}
	if len(fileEncKey) > 0 {
		chunkEncOverhead = EncryptionOverhead
	}

	bufferSize := sdk.maxBlocksInChunk * int(BlockSize)
	if sdk.erasureCode != nil { // erasure coding enabled
		bufferSize = sdk.erasureCode.DataBlocks * int(BlockSize)
	}
	bufferSize -= chunkEncOverhead
	buf := make([]byte, bufferSize)

	dagRoot, err := NewDAGRoot()
	if err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(err)
	}

	var chunkCount, fileSize int64
	for {
		select {
		case <-ctx.Done():
			return IPCFileMetaV2{}, ctx.Err()
		default:
		}

		isLastChunk := false

		n, readErr := io.ReadFull(reader, buf)
		if readErr != nil {
			if errors.Is(readErr, io.EOF) {
				if chunkCount == 0 {
					return IPCFileMetaV2{}, errSDK.Errorf("empty file")
				}
				break
			}
			if !errors.Is(readErr, io.ErrUnexpectedEOF) {
				return IPCFileMetaV2{}, errSDK.Wrap(readErr)
			}

			isLastChunk = true
		}

		chunkUpload, err := sdk.createChunkUpload(ctx, chunkCount, fileEncKey, buf[:n], bucket.Id, fileName)
		if err != nil {
			return IPCFileMetaV2{}, err
		}
		fileSize += chunkUpload.ActualSize

		if err := dagRoot.AddLink(chunkUpload.ChunkCID, chunkUpload.RawDataSize, chunkUpload.ProtoNodeSize); err != nil {
			return IPCFileMetaV2{}, errSDK.Wrap(err)
		}

		if err := sdk.uploadChunk(ctx, chunkUpload); err != nil {
			return IPCFileMetaV2{}, err
		}

		chunkCount++

		if isLastChunk {
			break
		}
	}

	rootCID, err := dagRoot.Build()
	if err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(err)
	}

	fileMeta, err := sdk.ipc.Storage.GetFileByName(&bind.CallOpts{From: sdk.ipc.Auth.From}, bucket.Id, fileName)
	if err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	fileID := calculateFileID(bucket.Id[:], fileName)

	var isFilled bool

	for !isFilled {
		isFilled, err = sdk.ipc.Storage.IsFileFilled(&bind.CallOpts{}, fileID)
		if err != nil {
			return IPCFileMetaV2{}, errSDK.Wrap(ipc.ErrorHashToError(err))
		}

		time.Sleep(time.Second)
	}

	tx, err := sdk.ipc.Storage.CommitFile(sdk.ipc.Auth, bucket.Id, fileName, new(big.Int).SetInt64(fileSize), rootCID.Bytes())
	if err != nil {
		return IPCFileMetaV2{}, errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	return IPCFileMetaV2{
		RootCID:     rootCID.String(),
		BucketName:  bucketName,
		Name:        fileName,
		EncodedSize: fileSize,
		CreatedAt:   time.Unix(fileMeta.CreatedAt.Int64(), 0),
		CommittedAt: time.Now(),
	}, errSDK.Wrap(sdk.ipc.WaitForTx(ctx, tx.Hash()))
}

func (sdk *IPC) createChunkUpload(ctx context.Context, index int64, fileEncryptionKey, data []byte, bucketID [32]byte, fileName string) (_ IPCFileChunkUploadV2, err error) {
	defer mon.Task()(&ctx, index)(&err)

	if len(fileEncryptionKey) > 0 {
		data, err = encryption.Encrypt(fileEncryptionKey, data, []byte(fmt.Sprintf("%d", index)))
		if err != nil {
			return IPCFileChunkUploadV2{}, errSDK.Wrap(err)
		}
	}

	size := int64(len(data))
	blockSize := BlockSize.ToInt64()
	if sdk.erasureCode != nil { // erasure coding is enabled
		data, err = sdk.erasureCode.Encode(data)
		if err != nil {
			return IPCFileChunkUploadV2{}, errSDK.Wrap(err)
		}
		// equivalent to notion of shard size in erasure coding terminology
		blockSize = int64(len(data) / (sdk.erasureCode.DataBlocks + sdk.erasureCode.ParityBlocks))
	}

	chunkDAG, err := BuildDAG(ctx, bytes.NewBuffer(data), blockSize, nil)
	if err != nil {
		return IPCFileChunkUploadV2{}, err
	}

	cids, sizes, protoChunk, err := toIPCProtoChunk(chunkDAG.CID.String(), index, size, chunkDAG.Blocks)
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

	tx, err := sdk.ipc.Storage.AddFileChunk(sdk.ipc.Auth, chunkDAG.CID.Bytes(), bucketID, fileName, new(big.Int).SetInt64(size),
		cids, sizes, new(big.Int).SetInt64(index))
	if err != nil {
		return IPCFileChunkUploadV2{}, errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	return IPCFileChunkUploadV2{
		Index:         index,
		ChunkCID:      chunkDAG.CID,
		ActualSize:    size,
		RawDataSize:   chunkDAG.RawDataSize,
		ProtoNodeSize: chunkDAG.ProtoNodeSize,
		Blocks:        chunkDAG.Blocks,
		BucketID:      bucketID,
		FileName:      fileName,
	}, errSDK.Wrap(sdk.ipc.WaitForTx(ctx, tx.Hash()))
}

func (sdk *IPC) uploadChunk(ctx context.Context, fileChunkUpload IPCFileChunkUploadV2) (err error) {
	defer mon.Task()(&ctx, fileChunkUpload)(&err)

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(sdk.maxConcurrency)

	pool := newConnectionPool()
	defer func() {
		if err := pool.close(); err != nil {
			slog.Warn("failed to close connection", slog.String("error", err.Error()))
		}
	}()

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
			defer mon.Task()(&deriveCtx, block.CID)(&err)

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

			dataTypes := map[string][]eip712.TypedData{
				"StorageData": {
					{Name: "chunkCID", Type: "bytes"},
					{Name: "blockCID", Type: "bytes32"},
					{Name: "chunkIndex", Type: "uint256"},
					{Name: "blockIndex", Type: "uint8"},
					{Name: "nodeId", Type: "bytes"},
					{Name: "nonce", Type: "uint256"},
				},
			}

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

			dataMessage := map[string]interface{}{
				"chunkCID":   chunkCid.Bytes(),
				"blockCID":   bcid,
				"chunkIndex": big.NewInt(fileChunkUpload.Index),
				"blockIndex": uint8(i),
				"nodeId":     id,
				"nonce":      nonce,
			}

			signature, err := eip712.Sign(privateKey, eip712.Domain{
				Name:              "Storage",
				Version:           "1",
				ChainID:           sdk.chainID,
				VerifyingContract: common.HexToAddress(sdk.storageAddress),
			}, dataMessage, dataTypes)
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
			}, sender.Send); err != nil && !errors.Is(err, io.EOF) {
				return errSDK.Wrap(err)
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
		return IPCFileDownload{}, errSDK.Errorf("empty bucket id")
	}
	if fileName == "" {
		return IPCFileDownload{}, errSDK.Errorf("empty file name")
	}

	fileName, err := sdk.maybeEncryptMetadata(fileName, bucketName+"/"+fileName)
	if err != nil {
		return IPCFileDownload{}, errSDK.Wrap(err)
	}
	bucketName, err = sdk.maybeEncryptMetadata(bucketName, bucketName)
	if err != nil {
		return IPCFileDownload{}, errSDK.Wrap(err)
	}

	res, err := sdk.client.FileDownloadCreate(ctx, &pb.IPCFileDownloadCreateRequest{
		BucketName: bucketName,
		FileName:   fileName,
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
		BucketName: res.BucketName,
		Name:       fileName,
		Chunks:     chunks,
	}, nil
}

// FileSetPublicAccess change file status from/to public.
func (sdk *IPC) FileSetPublicAccess(ctx context.Context, bucketName, fileName string, isPublic bool) (err error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&err)

	fileName, err = sdk.maybeEncryptMetadata(fileName, bucketName+"/"+fileName)
	if err != nil {
		return errSDK.Wrap(err)
	}
	bucketName, err = sdk.maybeEncryptMetadata(bucketName, bucketName)
	if err != nil {
		return errSDK.Wrap(err)
	}

	bucket, err := sdk.ipc.Storage.GetBucketByName(&bind.CallOpts{From: sdk.ipc.Auth.From}, bucketName)
	if err != nil {
		return errSDK.Wrap(ipc.ErrorHashToError(err))
	}

	file, err := sdk.ipc.Storage.GetFileByName(&bind.CallOpts{}, bucket.Id, fileName)
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

		if err := sdk.downloadChunkBlocks(ctx, fileDownload.BucketName, fileDownload.Name, sdk.ipc.Auth.From.String(), chunkDownload, fileEncKey, writer); err != nil {
			return err
		}
	}

	return nil
}

func (sdk *IPC) createChunkDownload(ctx context.Context, bucketName, fileName string, chunk Chunk) (_ FileChunkDownload, err error) {
	defer mon.Task()(&ctx, chunk)(&err)

	res, err := sdk.client.FileDownloadChunkCreate(ctx, &pb.IPCFileDownloadChunkCreateRequest{
		BucketName: bucketName,
		FileName:   fileName,
		ChunkCid:   chunk.CID,
		Address:    sdk.ipc.Auth.From.String(),
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

func (sdk *IPC) downloadChunkBlocks(
	ctx context.Context,
	bucketName, fileName, address string,
	chunkDownload FileChunkDownload,
	fileEncryptionKey []byte,
	writer io.Writer,
) (err error) {

	defer mon.Task()(&ctx, chunkDownload)(&err)

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

			blockData, err := sdk.fetchBlockData(ctx, pool, chunkDownload.CID, bucketName, fileName, address, chunkDownload.Index, int64(i), block)
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
	if sdk.erasureCode != nil { // erasure coding is enabled
		data, err = sdk.erasureCode.ExtractData(blocks, int(chunkDownload.Size))
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

func (sdk *IPC) fetchBlockData(
	ctx context.Context,
	pool *connectionPool,
	chunkCID, bucketName, fileName, address string,
	chunkIndex, blockIdndex int64,
	block FileBlockDownload,
) ([]byte, error) {

	if block.Akave == nil && block.Filecoin == nil {
		return nil, errMissingBlockMetadata
	}

	client, closer, err := pool.createIPCClient(block.Akave.NodeAddress, sdk.useConnectionPool)
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

// Encrypts the given metadata if metadata encyption is enabled and encryption key is set.
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
		sizes = append(sizes, new(big.Int).SetInt64(int64(len(block.Data))))
	}
	return cids, sizes, &pb.IPCChunk{
		Cid:    chunkCid,
		Index:  index,
		Size:   size,
		Blocks: pbBlocks,
	}, nil
}

func calculateFileID(bucketID []byte, name string) common.Hash {
	var b []byte
	b = append(b, bucketID...)
	b = append(b, name...)

	return crypto.Keccak256Hash(b)
}
