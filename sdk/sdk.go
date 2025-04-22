// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package sdk

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ipfs/boxo/ipld/merkledag"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-unixfs"
	"github.com/ipfs/go-unixfs/importer/helpers"
	"github.com/spacemonkeygo/monkit/v3"
	"github.com/zeebo/errs/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/akave-ai/akavesdk/private/encryption"
	"github.com/akave-ai/akavesdk/private/ipc"
	"github.com/akave-ai/akavesdk/private/memory"
	"github.com/akave-ai/akavesdk/private/pb"
	"github.com/akave-ai/akavesdk/private/spclient"
)

const (
	// BlockSize is the size of a block. Keep in mind that encryption adds some overhead and max supported block size(with added encryption) is 1MiB.
	// TODO: after removing normal api, chnage it to 1MiB.
	BlockSize = 1 * memory.MB
	// EncryptionOverhead is the overhead of encryption.
	EncryptionOverhead = 28 // 16 bytes for AES-GCM tag, 12 bytes for nonce

	minBucketNameLength = 3
	minFileSize         = 127 // 127 bytes
)

var (
	errSDK                  = errs.Tag("sdk")
	errMissingBlockMetadata = errors.New("missing block metadata")

	mon = monkit.Package()
)

// Option is a SDK option.
type Option func(*SDK)

// SDK is the Akave SDK.
type SDK struct {
	client               pb.NodeAPIClient
	conn                 *grpc.ClientConn
	spClient             *spclient.SPClient
	streamingErasureCode *ErasureCode

	maxConcurrency            int
	blockPartSize             int64
	useConnectionPool         bool
	privateKey                string
	encryptionKey             []byte // empty means no encryption
	streamingMaxBlocksInChunk int
	parityBlocksCount         int  // 0 means no erasure coding applied
	useMetadataEncryption     bool // encrypts bucket and file names if true
	chunkBuffer               int
}

// WithMetadataEncryption sets the metadata encryption for the SDK.
func WithMetadataEncryption() func(*SDK) {
	return func(s *SDK) {
		s.useMetadataEncryption = true
	}
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

// WithErasureCoding sets the erasure coding for the SDK.
func WithErasureCoding(parityBlocks int) func(*SDK) {
	return func(s *SDK) {
		s.parityBlocksCount = parityBlocks
	}
}

// WithChunkBuffer sets the chunk buffer size for streaming operations.
func WithChunkBuffer(bufferSize int) func(*SDK) {
	return func(s *SDK) {
		s.chunkBuffer = bufferSize
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
		chunkBuffer:               0, // Default value for chunk buffer
	}

	for _, opt := range options {
		opt(s)
	}

	if s.streamingMaxBlocksInChunk < 2 {
		return nil, errSDK.Errorf("streaming max blocks in chunk %d should be >= 2", s.streamingMaxBlocksInChunk)
	}

	keyLength := len(s.encryptionKey)
	if keyLength != 0 && keyLength != 32 {
		return nil, errSDK.Errorf("encyption key length should be 32 bytes long")
	}

	if s.parityBlocksCount > s.streamingMaxBlocksInChunk/2 {
		return nil, errSDK.Errorf("parity blocks count %d should be <= %d", s.parityBlocksCount, s.streamingMaxBlocksInChunk/2)
	}

	if s.parityBlocksCount > 0 { // erasure coding enabled
		s.streamingErasureCode, err = NewErasureCode(s.streamingMaxBlocksInChunk-s.parityBlocksCount, s.parityBlocksCount)
		if err != nil {
			return nil, errSDK.Wrap(err)
		}
	}

	s.spClient = spclient.New()

	return s, nil
}

// Close closes the SDK internal connection.
func (sdk *SDK) Close() error {
	sdk.spClient.Close()
	return sdk.conn.Close()
}

// StreamingAPI returns SDK streaming API.
func (sdk *SDK) StreamingAPI() *StreamingAPI {
	return &StreamingAPI{
		client:            pb.NewStreamAPIClient(sdk.conn),
		conn:              sdk.conn,
		spClient:          sdk.spClient,
		erasureCode:       sdk.streamingErasureCode,
		maxConcurrency:    sdk.maxConcurrency,
		blockPartSize:     sdk.blockPartSize,
		useConnectionPool: sdk.useConnectionPool,
		encryptionKey:     sdk.encryptionKey,
		maxBlocksInChunk:  sdk.streamingMaxBlocksInChunk,
		chunkBuffer:       sdk.chunkBuffer,
	}
}

// IPC returns SDK ipc API.
func (sdk *SDK) IPC() (*IPC, error) {
	client := pb.NewIPCNodeAPIClient(sdk.conn)

	connParams, err := client.ConnectionParams(context.Background(), &pb.ConnectionParamsRequest{})
	if err != nil {
		return nil, err
	}

	ipcClient, err := ipc.Dial(context.Background(), ipc.Config{
		DialURI:                connParams.DialUri,
		PrivateKey:             sdk.privateKey,
		StorageContractAddress: connParams.StorageAddress,
		AccessContractAddress:  connParams.AccessAddress,
	})
	if err != nil {
		return nil, err
	}

	return &IPC{
		client:                client,
		ipc:                   ipcClient,
		chainID:               ipcClient.ChainID(),
		storageAddress:        connParams.StorageAddress,
		conn:                  sdk.conn,
		erasureCode:           sdk.streamingErasureCode,
		privateKey:            sdk.privateKey,
		maxConcurrency:        sdk.maxConcurrency,
		blockPartSize:         sdk.blockPartSize,
		useConnectionPool:     sdk.useConnectionPool,
		encryptionKey:         sdk.encryptionKey,
		maxBlocksInChunk:      sdk.streamingMaxBlocksInChunk,
		useMetadataEncryption: sdk.useMetadataEncryption,
		chunkBuffer:           sdk.chunkBuffer,
	}, nil
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
		Name:      res.Name,
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

	_, err = sdk.client.BucketDelete(ctx, &pb.BucketDeleteRequest{Name: bucketName})
	if err != nil {
		return errSDK.Wrap(err)
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

func encryptionKey(parentKey []byte, infoData ...string) ([]byte, error) {
	if len(parentKey) == 0 {
		return nil, nil
	}

	info := strings.Join(infoData, "/")
	key, err := encryption.DeriveKey(parentKey, []byte(info))
	if err != nil {
		return nil, err
	}

	return key, nil
}
