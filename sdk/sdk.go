// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package sdk

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ipfs/boxo/ipld/merkledag"
	"github.com/ipfs/boxo/ipld/unixfs"
	"github.com/ipfs/boxo/ipld/unixfs/importer/helpers"
	"github.com/ipfs/go-cid"
	"github.com/spacemonkeygo/monkit/v3"
	"github.com/zeebo/errs/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/akave-ai/akavesdk/private/encryption"
	"github.com/akave-ai/akavesdk/private/erasurecode"
	"github.com/akave-ai/akavesdk/private/ipc"
	"github.com/akave-ai/akavesdk/private/memory"
	"github.com/akave-ai/akavesdk/private/pb"
	"github.com/akave-ai/akavesdk/private/retry"
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
	client     pb.NodeAPIClient
	conn       *grpc.ClientConn
	httpClient *http.Client
	ec         *erasurecode.ErasureCode

	maxConcurrency            int
	blockPartSize             int64
	useConnectionPool         bool
	privateKey                string
	encryptionKey             []byte // empty means no encryption
	streamingMaxBlocksInChunk int
	parityBlocksCount         int  // 0 means no erasure coding applied
	useMetadataEncryption     bool // encrypts bucket and file names if true
	chunkBuffer               int
	batchSize                 int

	withRetry retry.WithRetry
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

// WithBatchSize sets the chunk batch size for ipc operations.
func WithBatchSize(batchSize int) func(*SDK) {
	if batchSize < 1 {
		batchSize = 1
	}
	return func(s *SDK) {
		s.batchSize = batchSize
	}
}

// WithCustomHttpClient sets a custom HTTP client for the SDK.
func WithCustomHttpClient(client *http.Client) func(*SDK) {
	return func(s *SDK) {
		s.httpClient = client
	}
}

// WithoutRetry disables retries for bucket creation and file upload ops.
func WithoutRetry() func(*SDK) {
	return func(s *SDK) {
		s.withRetry = retry.WithRetry{}
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
		batchSize:                 1,
		// enable retires by default
		withRetry: retry.WithRetry{
			MaxAttempts: 5,
			BaseDelay:   100 * time.Millisecond,
		},
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
		s.ec, err = erasurecode.New(s.streamingMaxBlocksInChunk-s.parityBlocksCount, s.parityBlocksCount)
		if err != nil {
			return nil, errSDK.Wrap(err)
		}
	}

	if s.httpClient == nil {
		s.httpClient = http.DefaultClient
	}

	// sanitize possibly faulty retry params.
	if s.withRetry.MaxAttempts < 0 {
		s.withRetry.MaxAttempts = 0
	}
	if s.withRetry.BaseDelay <= 100*time.Millisecond {
		s.withRetry.BaseDelay = 100 * time.Millisecond
	}

	return s, nil
}

// Close closes the SDK internal connection.
func (sdk *SDK) Close() error {
	return sdk.conn.Close()
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
		httpClient:            sdk.httpClient,
		ec:                    sdk.ec,
		privateKey:            sdk.privateKey,
		maxConcurrency:        sdk.maxConcurrency,
		blockPartSize:         sdk.blockPartSize,
		useConnectionPool:     sdk.useConnectionPool,
		encryptionKey:         sdk.encryptionKey,
		maxBlocksInChunk:      sdk.streamingMaxBlocksInChunk,
		useMetadataEncryption: sdk.useMetadataEncryption,
		chunkBuffer:           sdk.chunkBuffer,
		batchSize:             sdk.batchSize,
		withRetry:             sdk.withRetry,
	}, nil
}

// MonkitStats are the monkit stats for the sdk package.
type MonkitStats struct {
	Name         string
	Successes    int64
	Errors       map[string]int64
	Highwater    int64
	SuccessTimes *monkit.DurationDist
	FailureTimes *monkit.DurationDist
}

// GetMonkitStats returns monkit statistics for the sdk package.
func GetMonkitStats() []MonkitStats {
	var stats []MonkitStats

	mon.Funcs(func(f *monkit.Func) {
		name := f.FullName()
		success := f.Success()
		errors := f.Errors()
		highwater := f.Highwater()

		// Only include stats for functions that have been called
		if success > 0 || len(errors) > 0 {
			stats = append(stats, MonkitStats{
				Name:         name,
				Successes:    success,
				Errors:       errors,
				Highwater:    highwater,
				SuccessTimes: f.SuccessTimes(),
				FailureTimes: f.FailureTimes(),
			})
		}
	})

	return stats
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

// isRetryableTxError checks if error on sending transaction should be retried.
func isRetryableTxError(err error) bool {
	if err == nil {
		return false
	}

	msg := err.Error()
	switch {
	case strings.Contains(msg, "nonce too low"),
		strings.Contains(msg, "replacement transaction underpriced"),
		strings.Contains(msg, "EOF"):
		return true
	default:
		return false
	}
}

// skipToPosition skips the reader to the current upload position for resumable uploads.
func skipToPosition(reader io.Reader, position int64) error {
	if position > 0 {
		// Try to seek if the reader supports it
		if seeker, ok := reader.(io.Seeker); ok {
			_, err := seeker.Seek(position, io.SeekStart)
			if err != nil {
				return fmt.Errorf("failed to seek for resumable upload: %w", err)
			}
		} else {
			// Fallback to copying to discard if seeking is not supported
			_, err := io.CopyN(io.Discard, reader, position)
			if err != nil && !errors.Is(err, io.EOF) {
				return fmt.Errorf("failed to skip bytes for resumable upload: %w", err)
			}
		}
	}

	return nil
}
