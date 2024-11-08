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

	"akave.ai/akavesdk/private/memory"
	"akave.ai/akavesdk/private/testrand"
	"akave.ai/akavesdk/sdk"
)

const (
	maxConcurrency = 5
	blockPartSize  = 128 * memory.KiB
	secretKey      = "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"
)

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
	require.NotEmpty(t, bucket.ID)
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
		ID:        bucket.ID,
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
			ID:        bucket.ID,
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

func TestDeleteFile(t *testing.T) {
	ctx := context.Background()

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	bucketName := randomBucketName(t, 10)
	_, err = akave.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	rootCID := uploadSimpleFile(ctx, t, akave, bucketName, "example.txt")

	// check that file exists
	fMeta, err := akave.FileInfo(ctx, bucketName, "example.txt")
	require.NoError(t, err)
	require.Equal(t, rootCID, fMeta.RootCID)
	require.Equal(t, "example.txt", fMeta.Name)

	// delete file
	err = akave.FileDelete(ctx, bucketName, "example.txt")
	require.NoError(t, err)

	// check that file does not exist
	_, err = akave.FileInfo(ctx, bucketName, "example.txt")
	require.Error(t, err)
	st, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.NotFound, st.Code())

	// check that list of files is empty
	files, err := akave.ListFiles(ctx, bucketName)
	require.NoError(t, err)
	require.Empty(t, files)
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
func TestUploadDownloadSmallFiles(t *testing.T) {
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
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.BytesD(t, 2024, tc.fileSize)
				testUploadDownload(t, akave, data)
			})
		}
	})

	t.Run("with encyption", func(t *testing.T) {
		akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithEncryptionKey([]byte(secretKey)))
		require.NoError(t, err)
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.BytesD(t, 2024, tc.fileSize)
				testUploadDownload(t, akave, data)
			})
		}
	})
}

// Less than or close to block size.
func TestUploadDownloadSmallFilesV2(t *testing.T) {
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
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.BytesD(t, 2024, tc.fileSize)
				testUploadDownloadV2(t, akave, data)
			})
		}
	})

	t.Run("with encyption", func(t *testing.T) {
		akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithEncryptionKey([]byte(secretKey)))
		require.NoError(t, err)
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.BytesD(t, 2024, tc.fileSize)
				testUploadDownloadV2(t, akave, data)
			})
		}
	})
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
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.BytesD(t, 2024, tc.fileSize)
				testUploadDownloadStreamingAPI(t, akave, data)
			})
		}
	})

	t.Run("with encyption", func(t *testing.T) {
		akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithEncryptionKey([]byte(secretKey)))
		require.NoError(t, err)
		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.BytesD(t, 2024, tc.fileSize)
				testUploadDownloadStreamingAPI(t, akave, data)
			})
		}
	})
}

// Greater than block size.
func TestUploadDownload(t *testing.T) {
	tests := []struct {
		name     string
		fileSize int64 // Size in MB
	}{
		{"100 MB", 100},
		{"256 MB", 256},
		{"512 MB", 512},
		{"1 GB", 1000},
	}

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := testrand.BytesD(t, 2024, tc.fileSize*memory.MB.ToInt64())
			testUploadDownload(t, akave, data)
		})
	}
}

// Test case when a file has the same blocks.
func TestUploadDownloadSameBlocks(t *testing.T) {
	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)

	data := make([]byte, 10*memory.MB.ToInt64())
	testUploadDownload(t, akave, data)
}

func TestUploadDownloadV2(t *testing.T) {
	tests := []struct {
		name     string
		fileSize int64 // Size in MB
	}{
		{"100 MB", 100},
		{"256 MB", 256},
		{"512 MB", 512},
		{"1 GB", 1000},
	}

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := testrand.BytesD(t, 2024, tc.fileSize*memory.MB.ToInt64())
			testUploadDownloadV2(t, akave, data)
		})
	}
}

// Test case when a file has the same blocks.
func TestUploadDownloadV2SameBlocks(t *testing.T) {
	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)

	data := make([]byte, 10*memory.MB.ToInt64())
	testUploadDownloadV2(t, akave, data)
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

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := testrand.BytesD(t, 2024, tc.fileSize*memory.MB.ToInt64())
			testUploadDownloadStreamingAPI(t, akave, data)
		})
	}
}

// Test case when a file has the same blocks.
func TestUploadDownloadStreamingSameBlocks(t *testing.T) {
	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)

	data := make([]byte, 10*memory.MB.ToInt64())
	testUploadDownloadStreamingAPI(t, akave, data)
}

// Greater than block size.
func TestUploadDownloadWithEncryption(t *testing.T) {
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

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := testrand.BytesD(t, 2024, tc.fileSize*memory.MB.ToInt64())
			testUploadDownload(t, akave, data)
		})
	}
}

func TestUploadDownloadV2WithEncryption(t *testing.T) {
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

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := testrand.BytesD(t, 2024, tc.fileSize*memory.MB.ToInt64())
			testUploadDownloadV2(t, akave, data)
		})
	}
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

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := testrand.BytesD(t, 2024, tc.fileSize*memory.MB.ToInt64())
			testUploadDownloadStreamingAPI(t, akave, data)
		})
	}
}

func TestRangeDownload(t *testing.T) {
	ctx := context.Background()
	fileSize := 5 * memory.MB.ToInt64()

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	fileData := testrand.BytesD(t, 2024, fileSize)
	require.NoError(t, err)

	bucketName := randomBucketName(t, 10)
	_, err = akave.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	fileUpload, err := akave.CreateFileUpload(ctx, bucketName, "example.txt", fileSize, bytes.NewBuffer(fileData))
	require.NoError(t, err)

	upResult := akave.Upload(ctx, fileUpload)
	require.NoError(t, upResult)

	var downloaded bytes.Buffer
	fileDownload, err := akave.CreateRangeFileDownload(ctx, fileUpload.BucketName, "example.txt", 0, 2)
	require.NoError(t, err)
	require.True(t, len(fileDownload.Blocks) == 2)

	// download file blocks
	err = akave.Download(ctx, fileDownload, &downloaded)
	require.NoError(t, err)

	// check downloaded file contents with first 2 block content
	fileDAG, err := sdk.BuildDAG(ctx, bytes.NewBuffer(fileData), sdk.BlockSize.ToInt64(), nil)
	require.NoError(t, err)
	blockData1, err := sdk.ExtractBlockData(fileDAG.Blocks[0].CID, fileDAG.Blocks[0].Data)
	require.NoError(t, err)
	blockData2, err := sdk.ExtractBlockData(fileDAG.Blocks[1].CID, fileDAG.Blocks[1].Data)
	require.NoError(t, err)
	expected := blockData1
	expected = append(expected, blockData2...)
	checkFileContents(t, 10, expected, downloaded.Bytes())
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

func testUploadDownload(t *testing.T, akave *sdk.SDK, data []byte) {
	file := bytes.NewBuffer(data)

	bucketName := randomBucketName(t, 10)
	_, err := akave.CreateBucket(context.Background(), bucketName)
	require.NoError(t, err)

	now := time.Now()
	fileUpload, err := akave.CreateFileUpload(context.Background(), bucketName, "example.txt", int64(len(data)), file)
	require.NoError(t, err)
	fileUploadDuration := time.Since(now)
	t.Logf("Create file upload duration: %v", fileUploadDuration)

	require.Equal(t, bucketName, fileUpload.BucketName)
	require.Equal(t, "example.txt", fileUpload.FileName)
	require.Equal(t, int64(len(data)), fileUpload.FileSize)
	require.Greater(t, len(fileUpload.Blocks), 0)
	require.NotEmpty(t, fileUpload.RootCID)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	now = time.Now()
	upResult := akave.Upload(ctx, fileUpload)
	require.NoError(t, upResult)
	t.Logf("Upload duration: %v", time.Since(now))

	var downloaded bytes.Buffer
	fileDownload, err := akave.CreateFileDownload(context.Background(), fileUpload.BucketName, fileUpload.FileName)
	require.NoError(t, err)
	require.True(t, len(fileDownload.Blocks) > 0)

	now = time.Now()
	err = akave.Download(ctx, fileDownload, &downloaded)
	require.NoError(t, err)
	t.Logf("Download duration: %v", time.Since(now))

	checkFileContents(t, 10, data, downloaded.Bytes())
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
	assert.Equal(t, sdk.BucketIDFromName(bucketName), fileUpload.BucketID)
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
	assert.Equal(t, fileUpload.BucketID, fileMeta.BucketID)
	assert.Equal(t, fileUpload.Name, fileMeta.Name)
	assert.GreaterOrEqual(t, fileMeta.Size, int64(len(data))) // encrypted file size is larger
	assert.LessOrEqual(t, fileMeta.CreatedAt.UnixNano(), now.UnixNano())
	assert.GreaterOrEqual(t, fileMeta.CommitedAt.UnixNano(), fileMeta.CreatedAt.UnixNano())

	var downloaded bytes.Buffer
	fileDownload, err := streaming.CreateFileDownload(ctx, bucketName, fileUpload.Name)
	require.NoError(t, err)
	assert.Equal(t, fileUpload.StreamID, fileDownload.StreamID)
	assert.Equal(t, fileUpload.Name, fileDownload.Name)
	assert.Equal(t, fileUpload.BucketID, fileDownload.BucketID)
	require.True(t, len(fileDownload.Chunks) > 0)

	now = time.Now()
	err = streaming.Download(ctx, fileDownload, &downloaded)
	require.NoError(t, err)
	t.Logf("Download duration: %v", time.Since(now))

	checkFileContents(t, 10, data, downloaded.Bytes())
}

// Checks that encrypted file is not the same as unencrypted file.
func TestUploadDownloadEncryption(t *testing.T) {
	fileSize := 10 * memory.MB.ToInt64()

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)

	akaveEnc, err := sdk.New(
		PickNodeRPCAddress(t),
		maxConcurrency,
		blockPartSize.ToInt64(),
		true,
		sdk.WithEncryptionKey([]byte(secretKey)),
	)
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, akave.Close())
		require.NoError(t, akaveEnc.Close())
	})

	file := bytes.NewBuffer(testrand.BytesD(t, 2024, fileSize))

	bucketName := randomBucketName(t, 10)
	// create bucket
	_, err = akave.CreateBucket(context.Background(), bucketName)
	require.NoError(t, err)

	fileUpload, err := akaveEnc.CreateFileUpload(context.Background(), bucketName, "example.txt", fileSize, file)
	require.NoError(t, err)

	upResult := akaveEnc.Upload(context.Background(), fileUpload)
	require.NoError(t, upResult)

	var downloaded bytes.Buffer
	fileDownload, err := akave.CreateFileDownload(context.Background(), fileUpload.BucketName, "example.txt")
	require.NoError(t, err)
	require.True(t, len(fileDownload.Blocks) > 0)

	// download file blocks
	err = akave.Download(context.Background(), fileDownload, &downloaded)
	require.NoError(t, err)

	expected := testrand.BytesD(t, 2024, fileSize)
	checkFileContentsNot(t, 10, expected, downloaded.Bytes())
}

func TestUploadFileTooLarge(t *testing.T) {
	fileSize := 2 * memory.GiB.ToInt64()

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)

	bucketName := randomBucketName(t, 10)
	_, err = akave.CreateBucket(context.Background(), bucketName)
	require.NoError(t, err)

	_, err = akave.CreateFileUpload(context.Background(), bucketName, "example.txt", fileSize, nil)
	require.Error(t, err)
	require.Equal(t, "sdk: file size is too large", err.Error())
}

func TestListFiles(t *testing.T) {
	ctx := context.Background()

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	bucketName := randomBucketName(t, 10)
	_, err = akave.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	expected := make([]sdk.FileListItem, 0)
	for _, fileName := range []string{"foo", "bar", "baz"} {
		rootCID := uploadSimpleFile(ctx, t, akave, bucketName, fileName)

		expected = append(expected, sdk.FileListItem{
			RootCID: rootCID,
			Name:    fileName,
			Size:    2 * memory.MB.ToInt64(),
		})
	}

	list, err := akave.ListFiles(ctx, bucketName)
	require.NoError(t, err)
	// TODO: look for better solution
	for i, file := range list {
		require.WithinDuration(t, time.Now(), file.CreatedAt, 5*time.Second)
		list[i].CreatedAt = time.Time{}
	}
	require.ElementsMatch(t, expected, list)
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

	expected := make([]sdk.FileMetaV2, 0)
	for _, fileName := range []string{"foo", "bar", "baz"} {
		fm := uploadSimpleFileStream(ctx, t, akave, bucketName, fileName)

		expected = append(expected, sdk.FileMetaV2{
			StreamID:   fm.StreamID,
			RootCID:    fm.RootCID,
			BucketID:   sdk.BucketIDFromName(bucketName),
			Name:       fileName,
			Size:       fm.Size,
			CreatedAt:  fm.CreatedAt,
			CommitedAt: fm.CommitedAt,
		})
	}

	list, err := akave.StreamingAPI().ListFiles(ctx, bucketName)
	require.NoError(t, err)
	require.ElementsMatch(t, expected, list)
}

func TestFileInfo(t *testing.T) {
	ctx := context.Background()
	bucketName := randomBucketName(t, 10)

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)
	_, err = akave.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	rootCID := uploadSimpleFile(ctx, t, akave, bucketName, "example.txt")

	info, err := akave.FileInfo(ctx, bucketName, "example.txt")
	require.NoError(t, err)
	require.Equal(t, rootCID, info.RootCID)
	require.Equal(t, info.Name, "example.txt")
	require.Equal(t, info.Size, 2*memory.MB.ToInt64())
	require.WithinDuration(t, time.Now(), info.CreatedAt, 5*time.Second)
}

func TestStreamFileInfo(t *testing.T) {
	ctx := context.Background()
	bucketName := randomBucketName(t, 10)

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true)
	require.NoError(t, err)
	_, err = akave.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	filemeta := uploadSimpleFileStream(ctx, t, akave, bucketName, "example.txt")

	info, err := akave.StreamingAPI().FileInfo(ctx, bucketName, "example.txt")
	require.NoError(t, err)
	require.Equal(t, filemeta, info)
}

func testUploadDownloadV2(t *testing.T, akave *sdk.SDK, data []byte) {
	file := bytes.NewBuffer(data)

	bucketName := randomBucketName(t, 10)
	_, err := akave.CreateBucket(context.Background(), bucketName)
	require.NoError(t, err)

	now := time.Now()
	fileUpload, err := akave.CreateFileUpload(context.Background(), bucketName, "example.txt", int64(len(data)), file)
	require.NoError(t, err)
	fileUploadDuration := time.Since(now)
	t.Logf("Create file upload duration: %v", fileUploadDuration)

	require.Equal(t, bucketName, fileUpload.BucketName)
	require.Equal(t, "example.txt", fileUpload.FileName)
	require.Equal(t, int64(len(data)), fileUpload.FileSize)
	require.Greater(t, len(fileUpload.Blocks), 0)
	require.NotEmpty(t, fileUpload.RootCID)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	now = time.Now()
	upResult := akave.Upload(ctx, fileUpload)
	require.NoError(t, upResult)
	t.Logf("Upload duration: %v", time.Since(now))

	var downloaded bytes.Buffer
	fileDownload, err := akave.CreateFileDownloadV2(context.Background(), fileUpload.BucketName, fileUpload.FileName)
	require.NoError(t, err)
	require.True(t, len(fileDownload.Blocks) > 0)

	now = time.Now()
	err = akave.DownloadV2(ctx, fileDownload, &downloaded)
	require.NoError(t, err)
	t.Logf("Download duration: %v", time.Since(now))

	checkFileContents(t, 10, data, downloaded.Bytes())
}

func uploadSimpleFile(ctx context.Context, t *testing.T, akave *sdk.SDK, bucket string, fileName string) string {
	f := generateAny2MBFile(t)

	upload, err := akave.CreateFileUpload(ctx, bucket, fileName, 2*memory.MB.ToInt64(), f)
	require.NoError(t, err)
	err = akave.Upload(ctx, upload)
	require.NoError(t, err)

	return upload.RootCID
}

func uploadSimpleFileStream(ctx context.Context, t *testing.T, akave *sdk.SDK, bucket string, fileName string) sdk.FileMetaV2 {
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

// NodeRPCAddress is flag to set RPC address of akave node.
var NodeRPCAddress = flag.String("node-rpc-address", os.Getenv("NODE_RPC_ADDRESS"), "flag to set node rpc address")

// PickNodeRPCAddress picks node PRC address from flag.
func PickNodeRPCAddress(t testing.TB) string {
	if *NodeRPCAddress == "" || strings.EqualFold(*NodeRPCAddress, "omit") {
		t.Skip("node rpc address flag missing, example: -NODE_RPC_ADDRESS=<node rpc address>")
	}
	return *NodeRPCAddress
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

// checks lengths, first and last N bytes, expects them to be different.
func checkFileContentsNot(t *testing.T, n int, expected, actual []byte) {
	t.Helper()
	require.NotEqual(t, len(expected), len(actual))
	// check first 10 bytes
	require.NotEqualValues(t, expected[:n], actual[:n])
	// check last 10 bytes
	require.NotEqualValues(t, expected[len(expected)-n:], actual[len(actual)-n:])
}
