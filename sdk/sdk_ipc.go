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
	"sort"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ipfs/go-cid"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	"akave.ai/akavesdk/private/encryption"
	"akave.ai/akavesdk/private/ipc"
	"akave.ai/akavesdk/private/pb"
)

// IPC exposes SDK ipc API.
type IPC struct {
	client pb.IPCNodeAPIClient
	conn   *grpc.ClientConn
	ipc    *ipc.Client

	maxConcurrency    int
	blockPartSize     int64
	useConnectionPool bool
	encryptionKey     []byte // empty means no encryption
}

// CreateBucket creates a new bucket.
func (sdk *IPC) CreateBucket(ctx context.Context, name string) (_ *BucketCreateResult, err error) {
	defer mon.Task()(&ctx, name)(&err)

	if len(name) < minBucketNameLength {
		return nil, errSDK.Errorf("invalid bucket name")
	}

	tx, err := sdk.ipc.Storage.CreateBucket(sdk.ipc.Auth, name)
	if err != nil {
		return &BucketCreateResult{}, errSDK.Wrap(err)
	}

	if err := sdk.ipc.WaitForTx(ctx, tx.Hash()); err != nil {
		return &BucketCreateResult{}, errSDK.Wrap(err)
	}

	bucket, err := sdk.ipc.Storage.GetBucketByName(&bind.CallOpts{From: sdk.ipc.Auth.From}, name)
	if err != nil {
		return &BucketCreateResult{}, errSDK.Wrap(err)
	}

	return &BucketCreateResult{
		ID:        hex.EncodeToString(bucket.Id[:]),
		CreatedAt: time.Unix(bucket.CreatedAt.Int64(), 0),
	}, nil
}

// ViewBucket returns bucket's metadata.
func (sdk *IPC) ViewBucket(ctx context.Context, bucketName string) (_ Bucket, err error) {
	defer mon.Task()(&ctx, bucketName)(&err)

	if bucketName == "" {
		return Bucket{}, errSDK.Errorf("empty bucket name")
	}

	res, err := sdk.client.BucketView(ctx, &pb.IPCBucketViewRequest{
		BucketName: bucketName,
		Address:    sdk.ipc.Auth.From.String(),
	})
	if err != nil {
		return Bucket{}, errSDK.Wrap(err)
	}

	return Bucket{
		ID:        res.GetId(),
		Name:      res.GetName(),
		CreatedAt: res.GetCreatedAt().AsTime(),
	}, nil
}

// ListBuckets returns list of buckets.
func (sdk *IPC) ListBuckets(ctx context.Context) (_ []Bucket, err error) {
	defer mon.Task()(&ctx)(&err)

	res, err := sdk.client.BucketList(ctx, &pb.IPCBucketListRequest{
		Address: sdk.ipc.Auth.From.String(),
	})
	if err != nil {
		return nil, errSDK.Wrap(err)
	}

	buckets := make([]Bucket, 0, len(res.Buckets))
	for _, bucket := range res.Buckets {
		buckets = append(buckets, Bucket{
			ID:        bucket.GetId(),
			Name:      bucket.GetName(),
			CreatedAt: bucket.GetCreatedAt().AsTime(),
		})
	}

	return buckets, nil
}

// DeleteBucket deletes bucket by name.
func (sdk *IPC) DeleteBucket(ctx context.Context, name string) (err error) {
	defer mon.Task()(&ctx)(&err)

	bucket, err := sdk.client.BucketView(ctx, &pb.IPCBucketViewRequest{
		BucketName: name,
		Address:    sdk.ipc.Auth.From.String(),
	})
	if err != nil {
		return errSDK.Wrap(err)
	}

	id, err := hex.DecodeString(bucket.Id)
	if err != nil {
		return errSDK.Wrap(err)
	}

	var bucketID [32]byte
	copy(bucketID[:], id)

	tx, err := sdk.ipc.Storage.DeleteBucket(sdk.ipc.Auth, bucketID, name)
	if err != nil {
		return errSDK.Wrap(err)
	}

	return errSDK.Wrap(sdk.ipc.WaitForTx(ctx, tx.Hash()))
}

// FileInfo returns meta information for single file by bucket and file name.
func (sdk *IPC) FileInfo(ctx context.Context, bucketName string, fileName string) (_ FileMeta, err error) {
	res, err := sdk.client.FileView(ctx, &pb.IPCFileViewRequest{
		BucketName: bucketName,
		FileName:   fileName,
		Address:    sdk.ipc.Auth.From.String(),
	})
	if err != nil {
		return FileMeta{}, errSDK.Wrap(err)
	}

	return FileMeta{
		RootCID:   res.GetRootCid(),
		Name:      res.GetFileName(),
		Size:      res.GetSize(),
		CreatedAt: res.CreatedAt.AsTime(),
	}, nil
}

// ListFiles returns list of files in a particular bucket.
func (sdk *IPC) ListFiles(ctx context.Context, bucketName string) (_ []FileListItem, err error) {
	defer mon.Task()(&ctx, bucketName)(&err)

	if bucketName == "" {
		return nil, errSDK.Errorf("empty bucket name")
	}

	resp, err := sdk.client.FileList(ctx, &pb.IPCFileListRequest{
		BucketName: bucketName,
		Address:    sdk.ipc.Auth.From.String(),
	})
	if err != nil {
		return nil, errSDK.Wrap(err)
	}

	files := make([]FileListItem, 0, len(resp.List))
	for _, fileMeta := range resp.List {
		files = append(files, FileListItem{
			RootCID:   fileMeta.GetRootCid(),
			Name:      fileMeta.GetName(),
			Size:      fileMeta.GetSize(),
			CreatedAt: fileMeta.GetCreatedAt().AsTime(),
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

	bucket, err := sdk.ipc.Storage.GetBucketByName(&bind.CallOpts{From: sdk.ipc.Auth.From}, bucketName)
	if err != nil {
		return errSDK.Wrap(err)
	}

	file, err := sdk.ipc.Storage.GetFileByName(&bind.CallOpts{}, bucket.Id, fileName)
	if err != nil {
		return errSDK.Wrap(err)
	}

	tx, err := sdk.ipc.Storage.DeleteFile(sdk.ipc.Auth, file.Id, bucket.Id, fileName, true)
	if err != nil {
		return errSDK.Wrap(err)
	}

	return errSDK.Wrap(sdk.ipc.WaitForTx(ctx, tx.Hash()))
}

// CreateFileUpload creates a new file upload request.
func (sdk *IPC) CreateFileUpload(ctx context.Context, bucketName string, fileName string, fileSize int64, reader io.Reader) (_ FileUpload, err error) {
	defer mon.Task()(&ctx, bucketName, fileName, fileSize)(&err)
	if bucketName == "" {
		return FileUpload{}, errSDK.Errorf("empty bucket name")
	}
	if fileSize < minFileSize {
		return FileUpload{}, errSDK.Errorf("file size is too small")
	}

	bucket, err := sdk.ipc.Storage.GetBucketByName(&bind.CallOpts{From: sdk.ipc.Auth.From}, bucketName)
	if err != nil {
		return FileUpload{}, errSDK.Wrap(err)
	}

	encKey, err := encryptionKey(sdk.encryptionKey, bucketName, fileName)
	if err != nil {
		return FileUpload{}, errSDK.Wrap(err)
	}

	dag, err := BuildDAG(ctx, reader, BlockSize.ToInt64(), encKey)
	if err != nil {
		return FileUpload{}, errSDK.Wrap(err)
	}
	rootCID := dag.CID

	cides := make([][]byte, 0, len(dag.Blocks))
	sizes := make([]*big.Int, 0, len(dag.Blocks))

	req := &pb.IPCFileUploadCreateRequest{
		RootCid: rootCID.String(),
		Size:    fileSize,
		Blocks:  make([]*pb.IPCFileUploadCreateRequest_IPCBlock, 0),
	}

	for _, block := range dag.Blocks {
		blockCID, err := cid.Decode(block.CID)
		if err != nil {
			return FileUpload{}, errSDK.Wrap(err)
		}

		cides = append(cides, blockCID.Bytes())
		sizes = append(sizes, new(big.Int).SetInt64(int64(len(block.Data))))

		req.Blocks = append(req.Blocks, &pb.IPCFileUploadCreateRequest_IPCBlock{
			Cid:  block.CID,
			Size: int64(len(block.Data)),
		})
	}

	tx, err := sdk.ipc.Storage.AddFile(sdk.ipc.Auth, rootCID.Bytes(), bucket.Id, fileName, new(big.Int).SetInt64(fileSize), cides, sizes)
	if err != nil {
		return FileUpload{}, errSDK.Wrap(err)
	}

	if err := sdk.ipc.WaitForTx(ctx, tx.Hash()); err != nil {
		return FileUpload{}, errSDK.Wrap(err)
	}

	res, err := sdk.client.FileUploadCreate(ctx, req)
	if err != nil {
		return FileUpload{}, errSDK.Wrap(err)
	}

	blocks := make([]FileBlock, len(res.Blocks))
	for i, upload := range res.Blocks {
		ch, found := blockByCID(dag.Blocks, upload.Cid)
		if !found {
			return FileUpload{}, errSDK.Errorf("block not found")
		}
		blocks[i] = FileBlock{
			CID:         upload.Cid,
			Data:        ch.Data,
			Permit:      upload.Permit,
			NodeAddress: upload.NodeAddress,
			NodeID:      upload.NodeId,
		}
	}

	return FileUpload{
		RootCID:    rootCID.String(),
		BucketName: bucketName,
		FileName:   fileName,
		FileSize:   fileSize,
		Blocks:     blocks,
	}, nil
}

// Upload uploads a file.
func (sdk *IPC) Upload(ctx context.Context, fileUpload FileUpload) (err error) {
	defer mon.Task()(&ctx, fileUpload)(&err)
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(sdk.maxConcurrency)

	pool := newConnectionPool()
	defer func() {
		if err := pool.close(); err != nil {
			slog.Warn("failed to close connection", slog.String("error", err.Error()))
		}
	}()

	for i, block := range fileUpload.Blocks {
		i, block := i, block
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

			err = sdk.uploadIPCBlockSegments(ctx, &pb.IPCFileBlockData{
				Data: block.Data,
				Cid:  block.CID,
			}, sender.Send)
			if err != nil {
				return err
			}

			_, closeErr := sender.CloseAndRecv()
			return closeErr
		})
	}

	return errSDK.Wrap(g.Wait())
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
		end := i + sdk.blockPartSize
		if end > dataLen {
			end = dataLen
		}

		fileBlock := &pb.IPCFileBlockData{
			Data: data[i:end:end],
			Cid:  block.Cid,
		}

		if err := sender(fileBlock); err != nil {
			return err
		}
	}

	return nil
}

// CreateFileDownload creates a new download request.
func (sdk *IPC) CreateFileDownload(ctx context.Context, bucketName string, fileName string) (_ FileDownload, innerErr error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&innerErr)
	if bucketName == "" {
		return FileDownload{}, errSDK.Errorf("empty bucket id")
	}

	if fileName == "" {
		return FileDownload{}, errSDK.Errorf("empty file name")
	}

	res, err := sdk.client.FileDownloadCreate(ctx, &pb.IPCFileDownloadCreateRequest{
		BucketName: bucketName,
		FileName:   fileName,
		Address:    sdk.ipc.Auth.From.String(),
	})
	if err != nil {
		return FileDownload{}, errSDK.Wrap(err)
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

	return FileDownload{
		BucketName: bucketName,
		FileName:   fileName,
		Blocks:     blocks,
	}, nil
}

// Download function downloads a file from FileDownload.
func (sdk *IPC) Download(ctx context.Context, fileDownload FileDownload, writer io.Writer) (err error) {
	defer mon.Task()(&ctx, fileDownload)(&err)
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
	ch := make(chan retrievedBlock, len(fileDownload.Blocks))

	encKey, err := encryptionKey(sdk.encryptionKey, fileDownload.BucketName, fileDownload.FileName)
	if err != nil {
		return errSDK.Wrap(err)
	}

	for i, block := range fileDownload.Blocks {
		i, block := i, block
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
							slog.String("node_address", block.NodeAddress),
							slog.String("error", closeErr.Error()),
						)
					}
				}()
			}

			downloadClient, err := client.FileDownloadBlock(ctx, &pb.IPCFileDownloadBlockRequest{
				BlockCid: block.CID,
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

	blocks := make([]retrievedBlock, 0)
	for retrieved := range ch {
		blocks = append(blocks, retrieved)
	}
	sort.SliceStable(blocks, func(i, j int) bool {
		return blocks[j].Pos > blocks[i].Pos
	})
	for i, block := range blocks {
		data, err := ExtractBlockData(block.CID, block.Data)
		if err != nil {
			return errSDK.Wrap(err)
		}

		if len(encKey) > 0 { // if encryption is enabled
			info := fmt.Sprintf("block_%d", i)
			data, err = encryption.Decrypt(encKey, data, []byte(info))
			if err != nil {
				return errSDK.Wrap(err)
			}
		}

		if _, err := writer.Write(data); err != nil {
			return err
		}
	}

	return nil
}
