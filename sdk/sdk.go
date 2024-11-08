// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package sdk

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"sort"
	"strings"

	"github.com/ipfs/boxo/ipld/merkledag"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-unixfs"
	"github.com/ipfs/go-unixfs/importer/helpers"
	"github.com/spacemonkeygo/monkit/v3"
	"github.com/zeebo/errs/v2"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"akave.ai/akavesdk/private/encryption"
	"akave.ai/akavesdk/private/ipc"
	"akave.ai/akavesdk/private/memory"
	"akave.ai/akavesdk/private/pb"
	"akave.ai/akavesdk/private/spclient"
)

const (
	// BlockSize is the size of a block. Keep in mind that encryption adds some overhead and max supported block size(with added encryption) is 1MiB.
	BlockSize           = 1 * memory.MB
	minBucketNameLength = 3
	minFileSize         = 127 // 127 bytes
	maxFileSize         = 1 * memory.GiB
)

var errSDK = errs.Tag("sdk")
var mon = monkit.Package()

// Option is a SDK option.
type Option func(*SDK)

// SDK is the Akave SDK.
type SDK struct {
	client   pb.NodeAPIClient
	conn     *grpc.ClientConn
	spClient *spclient.SPClient

	maxConcurrency            int
	blockPartSize             int64
	useConnectionPool         bool
	privateKey                string
	encryptionKey             []byte // empty means no encryption
	streamingMaxBlocksInChunk int
}

// WithEncryptionKey sets the encryption key for the SDK.
func WithEncryptionKey(key []byte) func(*SDK) {
	return func(s *SDK) {
		s.encryptionKey = key
	}
}

// WithPrivateKey sets the private key for the SDK.
func WithPrivateKey(key string) func(*SDK) {
	return func(s *SDK) {
		s.privateKey = key
	}
}

// WithStreamingMaxBlocksInChunk sets the max blocks in chunk for streaming.
func WithStreamingMaxBlocksInChunk(maxBlocksInChunk int) func(*SDK) {
	return func(s *SDK) {
		s.streamingMaxBlocksInChunk = maxBlocksInChunk
	}
}

// New returns a new SDK.
func New(address string, maxConcurrency int, blockPartSize int64, useConnectionPool bool, options ...Option) (*SDK, error) {
	if blockPartSize <= 0 || blockPartSize > int64(helpers.BlockSizeLimit) {
		return nil, fmt.Errorf("invalid blockPartSize: %d. Valid range is 1-%d", blockPartSize, helpers.BlockSizeLimit)
	}

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	s := &SDK{
		client:                    pb.NewNodeAPIClient(conn),
		conn:                      conn,
		maxConcurrency:            maxConcurrency,
		blockPartSize:             blockPartSize,
		useConnectionPool:         useConnectionPool,
		streamingMaxBlocksInChunk: 32,
	}

	for _, opt := range options {
		opt(s)
	}

	keyLength := len(s.encryptionKey)
	if keyLength != 0 && keyLength != 32 {
		return nil, errSDK.Errorf("encyption key length should be 32 bytes long")
	}

	s.spClient = spclient.New()

	return s, nil
}

// Close closes the SDK internal connection.
func (sdk *SDK) Close() error {
	return sdk.conn.Close()
}

// StreamingAPI returns SDK streaming API.
func (sdk *SDK) StreamingAPI() *StreamingAPI {
	return &StreamingAPI{
		client:            pb.NewStreamAPIClient(sdk.conn),
		conn:              sdk.conn,
		maxConcurrency:    sdk.maxConcurrency,
		blockPartSize:     sdk.blockPartSize,
		useConnectionPool: sdk.useConnectionPool,
		encryptionKey:     sdk.encryptionKey,
		maxBlocksInChunk:  sdk.streamingMaxBlocksInChunk,
	}
}

// IPC returns SDK ipc API.
func (sdk *SDK) IPC() (*IPC, error) {
	client := pb.NewIPCNodeAPIClient(sdk.conn)

	res := &IPC{
		client:            client,
		conn:              sdk.conn,
		maxConcurrency:    sdk.maxConcurrency,
		blockPartSize:     sdk.blockPartSize,
		useConnectionPool: sdk.useConnectionPool,
		encryptionKey:     sdk.encryptionKey,
	}

	connParams, err := client.ConnectionParams(context.Background(), &pb.ConnectionParamsRequest{})
	if err != nil {
		return nil, err
	}

	res.ipc, err = ipc.Dial(context.Background(), ipc.Config{
		DialURI:         connParams.DialUri,
		PrivateKey:      sdk.privateKey,
		ContractAddress: connParams.ContractAddress,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// CreateBucket creates a new bucket.
func (sdk *SDK) CreateBucket(ctx context.Context, name string) (_ *BucketCreateResult, err error) {
	defer mon.Task()(&ctx, name)(&err)

	if len(name) < minBucketNameLength {
		return nil, errSDK.Errorf("invalid bucket name")
	}

	res, err := sdk.client.BucketCreate(ctx, &pb.BucketCreateRequest{Name: name})
	if err != nil {
		return nil, errSDK.Wrap(err)
	}

	return &BucketCreateResult{
		ID:        res.Id,
		CreatedAt: res.CreatedAt.AsTime(),
	}, nil
}

// ViewBucket creates a new bucket.
func (sdk *SDK) ViewBucket(ctx context.Context, bucketName string) (_ Bucket, err error) {
	defer mon.Task()(&ctx, bucketName)(&err)

	if bucketName == "" {
		return Bucket{}, errSDK.Errorf("empty bucket name")
	}

	res, err := sdk.client.BucketView(ctx, &pb.BucketViewRequest{
		BucketName: bucketName,
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
func (sdk *SDK) ListBuckets(ctx context.Context) (_ []Bucket, err error) {
	defer mon.Task()(&ctx)(&err)

	res, err := sdk.client.BucketList(ctx, &pb.BucketListRequest{})
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

// DeleteBucket deletes a bucket by name.
func (sdk *SDK) DeleteBucket(ctx context.Context, bucketName string) (err error) {
	defer mon.Task()(&ctx, bucketName)(&err)

	// TODO: add validation?

	_, err = sdk.client.BucketDelete(ctx, &pb.BucketDeleteRequest{BucketName: bucketName})
	if err != nil {
		return errSDK.Wrap(err)
	}

	return nil
}

// ListFiles returns list of files in a particular bucket.
func (sdk *SDK) ListFiles(ctx context.Context, bucketName string) (_ []FileListItem, err error) {
	defer mon.Task()(&ctx, bucketName)(&err)

	if bucketName == "" {
		return nil, errSDK.Errorf("empty bucket name")
	}

	resp, err := sdk.client.FileList(ctx, &pb.FileListRequest{
		BucketName: bucketName,
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

// FileInfo returns meta information for single file by bucket and file name.
func (sdk *SDK) FileInfo(ctx context.Context, bucketName string, fileName string) (_ FileMeta, err error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&err)
	if bucketName == "" {
		return FileMeta{}, errSDK.Errorf("empty bucket name")
	}
	if fileName == "" {
		return FileMeta{}, errSDK.Errorf("empty file name")
	}

	res, err := sdk.client.FileView(ctx, &pb.FileViewRequest{
		BucketName: bucketName,
		FileName:   fileName,
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

// CreateFileUpload creates a new file upload request.
func (sdk *SDK) CreateFileUpload(ctx context.Context, bucketName string, fileName string, fileSize int64, reader io.Reader) (_ FileUpload, err error) {
	defer mon.Task()(&ctx, bucketName, fileName, fileSize)(&err)
	if bucketName == "" {
		return FileUpload{}, errSDK.Errorf("empty bucket name")
	}
	if fileSize < minFileSize {
		return FileUpload{}, errSDK.Errorf("file size is too small")
	}
	if fileSize > int64(maxFileSize) {
		return FileUpload{}, errSDK.Errorf("file size is too large")
	}

	encKey, err := encryptionKey(sdk.encryptionKey, bucketName, fileName)
	if err != nil {
		return FileUpload{}, errSDK.Wrap(err)
	}

	dag, err := BuildDAG(ctx, reader, BlockSize.ToInt64(), encKey)
	if err != nil {
		return FileUpload{}, err
	}

	rootCID := dag.CID.String()
	req := &pb.FileUploadCreateRequest{
		BucketName: bucketName,
		RootCid:    rootCID,
		FileName:   fileName,
		Size:       fileSize,
	}
	req.Blocks = make([]*pb.FileUploadCreateRequest_Block, len(dag.Blocks))
	for i, block := range dag.Blocks {
		req.Blocks[i] = &pb.FileUploadCreateRequest_Block{
			Cid:  block.CID,
			Size: int64(len(block.Data)),
		}
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
		RootCID:    rootCID,
		BucketName: bucketName,
		FileName:   fileName,
		FileSize:   fileSize,
		Blocks:     blocks,
	}, nil
}

// Upload uploads a file.
func (sdk *SDK) Upload(ctx context.Context, fileUpload FileUpload) (err error) {
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

			client, closer, err := pool.createClient(block.NodeAddress, sdk.useConnectionPool)
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

			err = sdk.uploadBlockSegments(ctx, &pb.FileBlockData{
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

	if err := g.Wait(); err != nil {
		return errSDK.Wrap(err)
	}

	return nil
}

// CreateFileDownload creates a new download request.
func (sdk *SDK) CreateFileDownload(ctx context.Context, bucketName string, fileName string) (_ FileDownload, innerErr error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&innerErr)
	if bucketName == "" {
		return FileDownload{}, errSDK.Errorf("empty bucket id")
	}

	if fileName == "" {
		return FileDownload{}, errSDK.Errorf("empty file name")
	}

	res, err := sdk.client.FileDownloadCreate(ctx, &pb.FileDownloadCreateRequest{
		BucketName: bucketName,
		FileName:   fileName,
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

// CreateFileDownloadV2 creates a new download request.
func (sdk *SDK) CreateFileDownloadV2(ctx context.Context, bucketName string, fileName string) (_ FileDownloadSP, innerErr error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&innerErr)
	if bucketName == "" {
		return FileDownloadSP{}, errSDK.Errorf("empty bucket id")
	}

	if fileName == "" {
		return FileDownloadSP{}, errSDK.Errorf("empty file name")
	}

	res, err := sdk.client.FileDownloadCreateV2(ctx, &pb.FileDownloadCreateRequestV2{
		BucketName: bucketName,
		FileName:   fileName,
	})
	if err != nil {
		return FileDownloadSP{}, errSDK.Wrap(err)
	}

	blocks := make([]FileBlockSP, len(res.Blocks))
	for i, block := range res.Blocks {
		switch source := block.Source.(type) {
		case *pb.FileDownloadCreateResponseV2_BlockDownloadV2_NodeBlock:
			blocks[i] = FileBlockSP{
				CID:         block.GetCid(),
				Permit:      source.NodeBlock.GetPermit(),
				NodeAddress: source.NodeBlock.GetNodeAddress(),
				NodeID:      source.NodeBlock.GetNodeId(),
			}
		case *pb.FileDownloadCreateResponseV2_BlockDownloadV2_ServiceProviderBlock:
			blocks[i] = FileBlockSP{
				CID:       block.GetCid(),
				SPBaseURL: source.ServiceProviderBlock.GetSpAddress(),
			}
		default:
			return FileDownloadSP{}, errSDK.Wrap(err)
		}
	}

	return FileDownloadSP{
		BucketName: bucketName,
		FileName:   fileName,
		Blocks:     blocks,
	}, nil
}

// CreateRangeFileDownload creates a new download request with block ranges.
func (sdk *SDK) CreateRangeFileDownload(ctx context.Context, bucketName string, fileName string, start, end int64) (_ FileDownload, innerErr error) {
	defer mon.Task()(&ctx, bucketName, fileName, start, end)(&innerErr)
	if bucketName == "" {
		return FileDownload{}, errSDK.Errorf("empty bucket id")
	}

	if fileName == "" {
		return FileDownload{}, errSDK.Errorf("empty file name")
	}

	res, err := sdk.client.FileDownloadRangeCreate(ctx, &pb.FileDownloadRangeCreateRequest{
		BucketName: bucketName,
		FileName:   fileName,
		Start:      start,
		End:        end,
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
func (sdk *SDK) Download(ctx context.Context, fileDownload FileDownload, writer io.Writer) (err error) {
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

			client, closer, err := pool.createClient(block.NodeAddress, sdk.useConnectionPool)
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

			downloadClient, err := client.FileDownloadBlock(ctx, &pb.FileDownloadBlockRequest{
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

		if len(encKey) > 0 { // if ecnryption is enabled
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

// DownloadV2 function downloads a file from FileDownload.
func (sdk *SDK) DownloadV2(ctx context.Context, fileDownload FileDownloadSP, writer io.Writer) (err error) {
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

			data, err := sdk.downloadBlock(ctx, block, pool, i)
			if err != nil {
				return err
			}

			ch <- retrievedBlock{
				Pos:  i,
				CID:  block.CID,
				Data: data,
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

// downloadBlock handles downloading a block, either via spclient's FetchBlock or through a direct node connection.
func (sdk *SDK) downloadBlock(ctx context.Context, block FileBlockSP, pool *connectionPool, index int) ([]byte, error) {
	if block.SPBaseURL != "" && block.CID != "" {
		// Handle case when only NodeAddress is provided, use spclient's FetchBlock method
		cid, err := cid.Decode(block.CID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse block's CID %s: %w", block.CID, err)
		}
		fetchedBlock, err := sdk.spClient.FetchBlock(ctx, block.SPBaseURL, cid)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch block with CID %s: %w", block.CID, err)
		}
		return fetchedBlock.RawData(), nil
	}

	client, closer, err := pool.createClient(block.NodeAddress, sdk.useConnectionPool)
	if err != nil {
		return nil, fmt.Errorf("failed to create client for block %s: %w", block.CID, err)
	}
	if closer != nil {
		defer func() {
			if closeErr := closer(); closeErr != nil {
				slog.Warn("failed to close connection",
					slog.Int("block_index", index),
					slog.String("block_cid", block.CID),
					slog.String("node_address", block.NodeAddress),
					slog.String("error", closeErr.Error()),
				)
			}
		}()
	}

	downloadClient, err := client.FileDownloadBlock(ctx, &pb.FileDownloadBlockRequest{
		BlockCid: block.CID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download block %s: %w", block.CID, err)
	}

	var buf bytes.Buffer
	for {
		blockData, err := downloadClient.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("failed to receive block data for CID %s: %w", block.CID, err)
		}

		_, _ = buf.Write(blockData.Data)
	}

	return buf.Bytes(), nil
}

// FileDelete deletes a file by bucket name and file name.
func (sdk *SDK) FileDelete(ctx context.Context, bucketName, fileName string) (err error) {
	defer mon.Task()(&ctx, bucketName, fileName)(&err)

	if strings.TrimSpace(bucketName) == "" || strings.TrimSpace(fileName) == "" {
		return errSDK.Errorf("empty bucket or file name. Bucket: '%s', File: '%s'", bucketName, fileName)
	}

	_, err = sdk.client.FileDelete(ctx, &pb.FileDeleteRequest{
		BucketName: bucketName,
		FileName:   fileName,
	})
	if err != nil {
		return errSDK.Wrap(err)
	}

	return nil
}

func (sdk *SDK) uploadBlockSegments(ctx context.Context, block *pb.FileBlockData, sender func(*pb.FileBlockData) error) (err error) {
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

		fileBlock := &pb.FileBlockData{
			Data: data[i:end:end],
			Cid:  block.Cid,
		}

		if err := sender(fileBlock); err != nil {
			return err
		}
	}

	return nil
}

// ExtractBlockData unwraps the block data from the block(either protobuf or raw).
func ExtractBlockData(idStr string, data []byte) ([]byte, error) {
	id, err := cid.Decode(idStr)
	if err != nil {
		return nil, err
	}
	switch id.Type() {
	case cid.DagProtobuf:
		node, err := merkledag.DecodeProtobuf(data)
		if err != nil {
			return nil, err
		}
		fsNode, err := unixfs.FSNodeFromBytes(node.Data())
		if err != nil {
			return nil, err
		}
		return fsNode.Data(), nil
	case cid.Raw:
		return data, nil
	default:
		return nil, fmt.Errorf("unknown cid type: %v", id.Type())
	}
}

func BucketIDFromName(bucketName string) string {
	h := sha256.New()
	h.Write([]byte(bucketName))
	return string(h.Sum(nil))
}

func encryptionKey(masterKey []byte, bucketName, fileName string) ([]byte, error) {
	if len(masterKey) == 0 {
		return nil, nil
	}

	infoString := fmt.Sprintf("%s/%s", bucketName, fileName)
	key, err := encryption.DeriveKey(masterKey, []byte(infoString))
	if err != nil {
		return nil, err
	}

	return key, nil
}
