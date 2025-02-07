// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package sdk_test

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/akave-ai/akavesdk/private/memory"
	"github.com/akave-ai/akavesdk/private/testrand"
	"github.com/akave-ai/akavesdk/sdk"
)

const (
	maxConcurrency = 5
	blockPartSize  = 128 * memory.KiB
	secretKey      = "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"
)

func TestCreateSDKClient(t *testing.T) {
	t.Run("invalid max blocks in chunk", func(t *testing.T) {
		_, err := sdk.New("", maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithStreamingMaxBlocksInChunk(1))
		require.Error(t, err)
		require.Equal(t, "sdk: streaming max blocks in chunk 1 should be >= 2", err.Error())
	})

	t.Run("invalid erasure coding config", func(t *testing.T) {
		_, err := sdk.New("", maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithErasureCoding(17))
		require.Error(t, err)
		require.Equal(t, "sdk: parity blocks count 17 should be <= 16", err.Error())

		_, err = sdk.New("", maxConcurrency, blockPartSize.ToInt64(), true,
			sdk.WithErasureCoding(40),
			sdk.WithStreamingMaxBlocksInChunk(64),
		)
		require.Error(t, err)
		require.Equal(t, "sdk: parity blocks count 40 should be <= 32", err.Error())
	})

	t.Run("invalid encryption key size", func(t *testing.T) {
		_, err := sdk.New("", maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithEncryptionKey([]byte("short")))
		require.Error(t, err)
		require.Equal(t, "sdk: encyption key length should be 32 bytes long", err.Error())
	})
}

func TestCreateBucket(t *testing.T) {
	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	bucketName := randomBucketName(t, 12)
	bucket, err := akave.CreateBucket(context.Background(), bucketName)
	require.NoError(t, err)
	require.NotEmpty(t, bucket.CreatedAt)
	require.NotEmpty(t, bucket.Name)
}

func TestViewBucket(t *testing.T) {
	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	bucketName := randomBucketName(t, 10)
	// create bucket
	bucket, err := akave.CreateBucket(context.Background(), bucketName)
	require.NoError(t, err)
	expected := sdk.Bucket{
		Name:      bucketName,
		CreatedAt: bucket.CreatedAt,
	}
	actual, err := akave.ViewBucket(context.Background(), bucketName)
	require.NoError(t, err)
	require.Equal(t, expected, actual)
}

func TestListBuckets(t *testing.T) {
	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	expected := make([]sdk.Bucket, 0)
	for i := 0; i < 3; i++ {
		bucketName := randomBucketName(t, 10)
		// create bucket
		bucket, err := akave.CreateBucket(context.Background(), bucketName)
		require.NoError(t, err)
		expected = append(expected, sdk.Bucket{
			Name:      bucketName,
			CreatedAt: bucket.CreatedAt,
		})
	}

	buckets, err := akave.ListBuckets(context.Background())
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(buckets), 3)
	require.True(t, assert.Subset(t, buckets, expected))
}

func TestDeleteBucket(t *testing.T) {
	ctx := context.Background()

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	bucketName := randomBucketName(t, 10)
	_, err = akave.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	// check that bucket exists
	bucket, err := akave.ViewBucket(ctx, bucketName)
	require.NoError(t, err)
	require.Equal(t, bucketName, bucket.Name)

	// delete bucket
	err = akave.DeleteBucket(ctx, bucketName)
	require.NoError(t, err)

	// check that bucket does not exist
	_, err = akave.ViewBucket(ctx, bucketName)
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.NotFound, st.Code())
}

func TestStreamDeleteFile(t *testing.T) {
	ctx := context.Background()

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	streamingAPI := akave.StreamingAPI()

	bucketName := randomBucketName(t, 10)
	_, err = akave.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	fMeta := uploadSimpleFileStream(ctx, t, akave, bucketName, "example.txt")

	// check that file exists
	fMeta, err = streamingAPI.FileInfo(ctx, bucketName, fMeta.Name)
	require.NoError(t, err)
	require.Equal(t, "example.txt", fMeta.Name)

	// delete file
	err = streamingAPI.FileDelete(ctx, bucketName, fMeta.Name)
	require.NoError(t, err)

	// check that file does not exist
	_, err = streamingAPI.FileInfo(ctx, bucketName, fMeta.Name)
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.NotFound, st.Code())

	// check that list of files is empty
	files, err := streamingAPI.ListFiles(ctx, bucketName)
	require.NoError(t, err)
	require.Empty(t, files)
}

// Less than or close to block size.
func TestUploadDownloadStreamingSmallFiles(t *testing.T) {
	tests := []struct {
		name     string
		fileSize int64 // Size in bytes
	}{
		{"127 Bytes", 127},
		{"1 KB", 1 * memory.KB.ToInt64()},
		{"1 MB", 1 * memory.MB.ToInt64()},
		{"1 MiB", 1 * memory.MiB.ToInt64()},
	}

	t.Run("no encyption", func(t *testing.T) {
		akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, akave.Close())
		})

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.BytesD(t, 2024, tc.fileSize)
				testUploadDownloadStreamingAPI(t, akave, data)
				testUploadDownloadStreamingAPIV2(t, akave, data)
			})
		}
	})

	t.Run("with encyption", func(t *testing.T) {
		akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithEncryptionKey([]byte(secretKey)))
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, akave.Close())
		})

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.BytesD(t, 2024, tc.fileSize)
				testUploadDownloadStreamingAPI(t, akave, data)
				testUploadDownloadStreamingAPIV2(t, akave, data)
			})
		}
	})

	t.Run("with erasure coding", func(t *testing.T) {
		akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithErasureCoding(16))
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, akave.Close())
		})

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.BytesD(t, 2024, tc.fileSize)
				testUploadDownloadStreamingAPI(t, akave, data)
				testUploadDownloadStreamingAPIV2(t, akave, data)
			})
		}
	})

	t.Run("with erasure coding and encryption", func(t *testing.T) {
		akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true,
			sdk.WithErasureCoding(16), sdk.WithEncryptionKey([]byte(secretKey)),
		)
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, akave.Close())
		})

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.BytesD(t, 2024, tc.fileSize)
				testUploadDownloadStreamingAPI(t, akave, data)
				testUploadDownloadStreamingAPIV2(t, akave, data)
			})
		}
	})
}

func TestUploadRandomDownloadStreamingSmallFiles(t *testing.T) {
	tests := []struct {
		name     string
		fileSize int64 // Size in bytes
	}{
		{"127 Bytes", 127},
		{"1 KB", 1 * memory.KB.ToInt64()},
		{"1 MB", 1 * memory.MB.ToInt64()},
		{"1 MiB", 1 * memory.MiB.ToInt64()},
	}

	t.Run("without encryption", func(t *testing.T) {
		akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithErasureCoding(16))
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, akave.Close())
		})

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.BytesD(t, 2024, tc.fileSize)
				testUploadRandomDownloadStreamingAPI(t, akave, data)
			})
		}
	})

	t.Run("with encryption", func(t *testing.T) {
		akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true,
			sdk.WithErasureCoding(16), sdk.WithEncryptionKey([]byte(secretKey)),
		)
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, akave.Close())
		})

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.BytesD(t, 2024, tc.fileSize)
				testUploadRandomDownloadStreamingAPI(t, akave, data)
			})
		}
	})
}

// Greater than block size.
func TestUploadDownloadStreaming(t *testing.T) {
	tests := []struct {
		name     string
		fileSize int64 // Size in MB
	}{
		{"100 MB", 100},
		{"256 MB", 256},
		{"512 MB", 512},
		{"1 GB", 1000},
	}

	t.Run("without erasure coding", func(t *testing.T) {
		akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, akave.Close())
		})

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.BytesD(t, 2024, tc.fileSize*memory.MB.ToInt64())
				testUploadDownloadStreamingAPI(t, akave, data)
			})
		}
	})

	t.Run("with erasure coding", func(t *testing.T) {
		akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithErasureCoding(16))
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, akave.Close())
		})

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				if tc.fileSize > 256 {
					t.Skip("skipping test because of large file size")
				}
				data := testrand.BytesD(t, 2024, tc.fileSize*memory.MB.ToInt64())
				testUploadDownloadStreamingAPI(t, akave, data)
			})
		}
	})

	t.Run("with erasure coding and encryption", func(t *testing.T) {
		akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true,
			sdk.WithErasureCoding(16), sdk.WithEncryptionKey([]byte(secretKey)),
		)
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, akave.Close())
		})

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				if tc.fileSize > 256 {
					t.Skip("skipping test because of large file size")
				}
				data := testrand.BytesD(t, 2024, tc.fileSize*memory.MB.ToInt64())
				testUploadDownloadStreamingAPI(t, akave, data)
				testUploadDownloadStreamingAPIV2(t, akave, data)
			})
		}
	})
}

func TestUploadRandomDownloadStreaming(t *testing.T) {
	tests := []struct {
		name     string
		fileSize int64
	}{
		{"100 MB", 100},
		{"256 MB", 256},
	}

	t.Run("without encryption", func(t *testing.T) {
		akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithErasureCoding(16))
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, akave.Close())
		})

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.BytesD(t, 2024, tc.fileSize*memory.MB.ToInt64())
				testUploadRandomDownloadStreamingAPI(t, akave, data)
			})
		}
	})

	t.Run("with encryption", func(t *testing.T) {
		akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true,
			sdk.WithErasureCoding(16), sdk.WithEncryptionKey([]byte(secretKey)),
		)
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, akave.Close())
		})

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.BytesD(t, 2024, tc.fileSize*memory.MB.ToInt64())
				testUploadRandomDownloadStreamingAPI(t, akave, data)
			})
		}
	})
}

// Test case when a file has the same blocks.
func TestUploadDownloadStreamingSameBlocks(t *testing.T) {
	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	data := make([]byte, 10*memory.MB.ToInt64())
	testUploadDownloadStreamingAPI(t, akave, data)
}

// Greater than block size.
func TestUploadDownloadStreamingWithEncryption(t *testing.T) {
	tests := []struct {
		name     string
		fileSize int64 // Size in MB
	}{
		{"100 MB", 100},
		{"256 MB", 256},
		{"512 MB", 512},
		{"1 GB", 1000},
	}

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithEncryptionKey([]byte(secretKey)))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := testrand.BytesD(t, 2024, tc.fileSize*memory.MB.ToInt64())
			testUploadDownloadStreamingAPI(t, akave, data)
		})
	}
}

func TestUploadDownloadIPC(t *testing.T) {
	tests := []struct {
		name     string
		fileSize int64 // Size in MB
	}{
		{"1 MB", 1},
		{"5 MB", 5},
		{"15 MB", 15},
		{"35 MB", 35},
	}

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(PickPrivateKey(t)))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// TODO: when BytesD used cause "block not found on any live peer" in CI, and block already exist locally. Find issue and fix.
			data := testrand.BytesD(t, 1, tc.fileSize*memory.MB.ToInt64())
			testUploadDownloadIPC(t, ipc, data)
		})
	}
}

func TestStreamingRangeDownload(t *testing.T) {
	ctx := context.Background()
	fileSize := 80 * memory.MB.ToInt64()
	blocksInChunk := 30

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithStreamingMaxBlocksInChunk(blocksInChunk))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	streamingAPI := akave.StreamingAPI()

	fileData := testrand.BytesD(t, 2024, fileSize)
	require.NoError(t, err)

	bucketName := randomBucketName(t, 10)
	_, err = akave.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	fileUpload, err := streamingAPI.CreateFileUpload(ctx, bucketName, "example.txt")
	require.NoError(t, err)

	_, err = streamingAPI.Upload(ctx, fileUpload, bytes.NewBuffer(fileData))
	require.NoError(t, err)

	var downloaded bytes.Buffer
	fileDownload, err := streamingAPI.CreateRangeFileDownload(ctx, bucketName, "example.txt", 1, 3)
	require.NoError(t, err)
	require.True(t, len(fileDownload.Chunks) == 2)
	assert.True(t, fileDownload.Chunks[0].Index == 1)
	assert.True(t, fileDownload.Chunks[1].Index == 2)

	// download file blocks
	err = streamingAPI.Download(ctx, fileDownload, &downloaded)
	require.NoError(t, err)

	// check downloaded partial file contents
	expected := fileData[blocksInChunk*int(sdk.BlockSize.ToInt64()):] // first chunk is skipped
	checkFileContents(t, 10, expected, downloaded.Bytes())
}

func testUploadDownloadStreamingAPI(t *testing.T, akave *sdk.SDK, data []byte) {
	file := bytes.NewBuffer(data)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bucketName := randomBucketName(t, 10)
	_, err := akave.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	streaming := akave.StreamingAPI()

	now := time.Now()
	fileUpload, err := streaming.CreateFileUpload(ctx, bucketName, "example.txt")
	require.NoError(t, err)
	assert.Equal(t, bucketName, fileUpload.BucketName)
	assert.Equal(t, "example.txt", fileUpload.Name)
	assert.GreaterOrEqual(t, fileUpload.CreatedAt.UnixNano(), now.UnixNano())
	assert.NotEmpty(t, fileUpload.StreamID)
	_, err = uuid.Parse(fileUpload.StreamID)
	require.NoError(t, err)

	now = time.Now()
	fileMeta, err := streaming.Upload(ctx, fileUpload, file)
	require.NoError(t, err)
	t.Logf("Upload duration: %v", time.Since(now))
	assert.Equal(t, fileUpload.StreamID, fileMeta.StreamID)
	assert.NotEmpty(t, fileMeta.RootCID)
	assert.Equal(t, fileUpload.BucketName, fileMeta.BucketName)
	assert.Equal(t, fileUpload.Name, fileMeta.Name)
	assert.Greater(t, fileMeta.EncodedSize, int64(len(data)))
	assert.True(t, (fileMeta.Size-int64(len(data)))%sdk.EncryptionOverhead == 0)
	assert.LessOrEqual(t, fileMeta.CreatedAt.UnixNano(), now.UnixNano())
	assert.GreaterOrEqual(t, fileMeta.CommitedAt.UnixNano(), fileMeta.CreatedAt.UnixNano())

	var downloaded bytes.Buffer
	fileDownload, err := streaming.CreateFileDownload(ctx, bucketName, fileUpload.Name, "")
	require.NoError(t, err)
	assert.Equal(t, fileUpload.StreamID, fileDownload.StreamID)
	assert.Equal(t, fileUpload.Name, fileDownload.Name)
	assert.Equal(t, fileUpload.BucketName, fileDownload.BucketName)
	require.True(t, len(fileDownload.Chunks) > 0)
	size := int64(0)
	for _, chunk := range fileDownload.Chunks {
		size += chunk.Size
	}
	assert.Equal(t, size, fileMeta.Size)

	now = time.Now()
	err = streaming.Download(ctx, fileDownload, &downloaded)
	require.NoError(t, err)
	t.Logf("Download duration: %v", time.Since(now))

	checkFileContents(t, 10, data, downloaded.Bytes())
}

// checks version V2 of Download method.
func testUploadDownloadStreamingAPIV2(t *testing.T, akave *sdk.SDK, data []byte) {
	file := bytes.NewBuffer(data)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bucketName := randomBucketName(t, 10)
	_, err := akave.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	streaming := akave.StreamingAPI()

	now := time.Now()
	fileUpload, err := streaming.CreateFileUpload(ctx, bucketName, "example.txt")
	require.NoError(t, err)
	assert.Equal(t, bucketName, fileUpload.BucketName)
	assert.Equal(t, "example.txt", fileUpload.Name)
	assert.GreaterOrEqual(t, fileUpload.CreatedAt.UnixNano(), now.UnixNano())
	assert.NotEmpty(t, fileUpload.StreamID)
	_, err = uuid.Parse(fileUpload.StreamID)
	require.NoError(t, err)

	now = time.Now()
	fileMeta, err := streaming.Upload(ctx, fileUpload, file)
	require.NoError(t, err)
	t.Logf("Upload duration: %v", time.Since(now))
	assert.Equal(t, fileUpload.StreamID, fileMeta.StreamID)
	assert.NotEmpty(t, fileMeta.RootCID)
	assert.Equal(t, fileUpload.BucketName, fileMeta.BucketName)
	assert.Equal(t, fileUpload.Name, fileMeta.Name)
	assert.Greater(t, fileMeta.EncodedSize, int64(len(data)))
	assert.True(t, (fileMeta.Size-int64(len(data)))%sdk.EncryptionOverhead == 0)
	assert.LessOrEqual(t, fileMeta.CreatedAt.UnixNano(), now.UnixNano())
	assert.GreaterOrEqual(t, fileMeta.CommitedAt.UnixNano(), fileMeta.CreatedAt.UnixNano())

	var downloaded bytes.Buffer
	fileDownload, err := streaming.CreateFileDownload(ctx, bucketName, fileUpload.Name, "")
	require.NoError(t, err)
	assert.Equal(t, fileUpload.StreamID, fileDownload.StreamID)
	assert.Equal(t, fileUpload.Name, fileDownload.Name)
	assert.Equal(t, fileUpload.BucketName, fileDownload.BucketName)
	require.True(t, len(fileDownload.Chunks) > 0)
	size := int64(0)
	for _, chunk := range fileDownload.Chunks {
		size += chunk.Size
	}
	assert.Equal(t, size, fileMeta.Size)

	now = time.Now()
	err = streaming.DownloadV2(ctx, fileDownload, &downloaded)
	require.NoError(t, err)
	t.Logf("Download duration: %v", time.Since(now))

	checkFileContents(t, 10, data, downloaded.Bytes())
}

func testUploadRandomDownloadStreamingAPI(t *testing.T, akave *sdk.SDK, data []byte) {
	file := bytes.NewBuffer(data)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bucketName := randomBucketName(t, 10)
	_, err := akave.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	streaming := akave.StreamingAPI()

	now := time.Now()
	fileUpload, err := streaming.CreateFileUpload(ctx, bucketName, "example.txt")
	require.NoError(t, err)
	assert.Equal(t, bucketName, fileUpload.BucketName)
	assert.Equal(t, "example.txt", fileUpload.Name)
	assert.GreaterOrEqual(t, fileUpload.CreatedAt.UnixNano(), now.UnixNano())
	assert.NotEmpty(t, fileUpload.StreamID)
	_, err = uuid.Parse(fileUpload.StreamID)
	require.NoError(t, err)

	now = time.Now()
	fileMeta, err := streaming.Upload(ctx, fileUpload, file)
	require.NoError(t, err)
	t.Logf("Upload duration: %v", time.Since(now))
	assert.Equal(t, fileUpload.StreamID, fileMeta.StreamID)
	assert.NotEmpty(t, fileMeta.RootCID)
	assert.Equal(t, fileUpload.BucketName, fileMeta.BucketName)
	assert.Equal(t, fileUpload.Name, fileMeta.Name)
	assert.Greater(t, fileMeta.EncodedSize, int64(len(data)))
	assert.True(t, (fileMeta.Size-int64(len(data)))%sdk.EncryptionOverhead == 0)
	assert.LessOrEqual(t, fileMeta.CreatedAt.UnixNano(), now.UnixNano())
	assert.GreaterOrEqual(t, fileMeta.CommitedAt.UnixNano(), fileMeta.CreatedAt.UnixNano())

	var downloaded bytes.Buffer
	fileDownload, err := streaming.CreateFileDownload(ctx, bucketName, fileUpload.Name, "")
	require.NoError(t, err)
	assert.Equal(t, fileUpload.StreamID, fileDownload.StreamID)
	assert.Equal(t, fileUpload.Name, fileDownload.Name)
	assert.Equal(t, fileUpload.BucketName, fileDownload.BucketName)
	require.True(t, len(fileDownload.Chunks) > 0)
	size := int64(0)
	for _, chunk := range fileDownload.Chunks {
		size += chunk.Size
	}
	assert.Equal(t, size, fileMeta.Size)

	now = time.Now()
	err = streaming.DownloadRandom(ctx, fileDownload, &downloaded)
	require.NoError(t, err)
	t.Logf("Download duration: %v", time.Since(now))

	checkFileContents(t, 10, data, downloaded.Bytes())
}

func testUploadDownloadIPC(t *testing.T, ipc *sdk.IPC, data []byte) {
	file := bytes.NewBuffer(data)

	bucketName := randomBucketName(t, 10)
	fileName := randomBucketName(t, 10)
	_, err := ipc.CreateBucket(context.Background(), bucketName)
	require.NoError(t, err)

	now := time.Now()
	require.NoError(t, ipc.CreateFileUpload(context.Background(), bucketName, fileName))
	fileUploadDuration := time.Since(now)
	t.Logf("Create file upload duration: %v", fileUploadDuration)

	time.Sleep(5 * time.Second)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	now = time.Now()
	upResult, err := ipc.Upload(ctx, bucketName, fileName, file)
	require.NoError(t, err)
	require.Equal(t, upResult.Name, fileName)
	t.Logf("Upload duration: %v", time.Since(now))

	time.Sleep(5 * time.Second)

	var downloaded bytes.Buffer
	fileDownload, err := ipc.CreateFileDownload(context.Background(), upResult.BucketName, upResult.Name)
	require.NoError(t, err)
	require.True(t, len(fileDownload.Chunks) > 0)

	time.Sleep(5 * time.Second)

	now = time.Now()
	require.NoError(t, ipc.Download(ctx, fileDownload, &downloaded))
	t.Logf("Download duration: %v", time.Since(now))

	checkFileContents(t, 10, data, downloaded.Bytes())
}

func TestStreamListFiles(t *testing.T) {
	ctx := context.Background()

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	bucketName := randomBucketName(t, 10)
	_, err = akave.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	expected := make([]sdk.FileMeta, 0)
	for _, fileName := range []string{"foo", "bar", "baz"} {
		fm := uploadSimpleFileStream(ctx, t, akave, bucketName, fileName)

		expected = append(expected, sdk.FileMeta{
			StreamID:    fm.StreamID,
			RootCID:     fm.RootCID,
			BucketName:  bucketName,
			Name:        fileName,
			EncodedSize: fm.EncodedSize,
			Size:        fm.Size,
			CreatedAt:   fm.CreatedAt,
			CommitedAt:  fm.CommitedAt,
		})
	}

	list, err := akave.StreamingAPI().ListFiles(ctx, bucketName)
	require.NoError(t, err)
	require.ElementsMatch(t, expected, list)
}

func TestStreamFileVersions(t *testing.T) {
	ctx := context.Background()

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	bucketName := randomBucketName(t, 10)
	_, err = akave.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	fm1 := uploadSimpleFileStream(ctx, t, akave, bucketName, "foo.txt")
	fm2 := uploadSimpleFileStream(ctx, t, akave, bucketName, "foo.txt")

	list, err := akave.StreamingAPI().FileVersions(ctx, bucketName, "foo.txt")
	require.NoError(t, err)
	assert.Equal(t, []sdk.FileMeta{fm2, fm1}, list)
}

func TestStreamFileInfo(t *testing.T) {
	ctx := context.Background()
	bucketName := randomBucketName(t, 10)

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	_, err = akave.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	filemeta := uploadSimpleFileStream(ctx, t, akave, bucketName, "example.txt")

	info, err := akave.StreamingAPI().FileInfo(ctx, bucketName, "example.txt")
	require.NoError(t, err)
	require.Equal(t, filemeta, info)
}

func uploadSimpleFileStream(ctx context.Context, t *testing.T, akave *sdk.SDK, bucket string, fileName string) sdk.FileMeta {
	f := generateAny2MBFile(t)

	streaming := akave.StreamingAPI()

	upload, err := streaming.CreateFileUpload(ctx, bucket, fileName)
	require.NoError(t, err)
	fm, err := streaming.Upload(ctx, upload, f)
	require.NoError(t, err)

	return fm
}

func TestRandomFile(t *testing.T) {
	f1 := generate10MiBFile(t, 2025)
	f2 := generate10MiBFile(t, 2025)

	t.Log(f1.Bytes()[:20])
	t.Log(f2.Bytes()[:20])
	require.EqualValues(t, f1.Bytes(), f2.Bytes())
}

func generateAny2MBFile(t *testing.T) *bytes.Buffer {
	data := testrand.Bytes(t, 2*memory.MB.ToInt64())
	return bytes.NewBuffer(data)
}

func generate10MiBFile(t *testing.T, seed int64) *bytes.Buffer {
	data := testrand.BytesD(t, seed, 10*memory.MiB.ToInt64())
	return bytes.NewBuffer(data)
}

func randomBucketName(t require.TestingT, size int) string {
	b := make([]byte, size)
	_, err := rand.Read(b)
	require.NoError(t, err)
	return hex.EncodeToString(b)
}

var (
	// NodeRPCAddress is flag to set RPC address of akave node.
	NodeRPCAddress = flag.String("node-rpc-address", os.Getenv("NODE_RPC_ADDRESS"), "flag to set node rpc address")
	// PrivateKey is flag set to private key.
	PrivateKey = flag.String("sdk-private-key", os.Getenv("PRIVATE_KEY"), "flag to set private key")
)

// PickNodeRPCAddress picks node PRC address from flag.
func PickNodeRPCAddress(t testing.TB) string {
	if *NodeRPCAddress == "" || strings.EqualFold(*NodeRPCAddress, "omit") {
		t.Skip("node rpc address flag missing, example: -NODE_RPC_ADDRESS=<node rpc address>")
	}
	return *NodeRPCAddress
}

// PickPrivateKey picks hex private key of deployer.
func PickPrivateKey(t testing.TB) string {
	if *PrivateKey == "" || strings.EqualFold(*PrivateKey, "omit") {
		t.Skip("private key flag missing, example: -PRIVATE_KEY=<deployers hex private key>")
	}
	return *PrivateKey
}

// checks lengths, first and last N bytes.
func checkFileContents(t *testing.T, n int, expected, actual []byte) {
	t.Helper()
	require.Equal(t, len(expected), len(actual))
	// check first 10 bytes
	require.EqualValues(t, expected[:n], actual[:n])
	// check last 10 bytes
	require.EqualValues(t, expected[len(expected)-n:], actual[len(actual)-n:])
}
