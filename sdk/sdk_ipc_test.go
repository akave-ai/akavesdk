// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package sdk_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/ipctest"
	"github.com/akave-ai/akavesdk/private/memory"
	"github.com/akave-ai/akavesdk/private/testrand"
	"github.com/akave-ai/akavesdk/sdk"
)

func TestIPCCreateBucket(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	bucketName := randomBucketName(t, 10)

	createBucketResult, err := ipc.CreateBucket(context.Background(), bucketName)
	require.NoError(t, err)
	require.Equal(t, bucketName, createBucketResult.Name)
}

func TestIPCViewBucket(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	bucketName := randomBucketName(t, 10)

	createBucketResult, err := ipc.CreateBucket(context.Background(), bucketName)
	require.NoError(t, err)
	require.Equal(t, bucketName, createBucketResult.Name)

	viewBucketResult, err := ipc.ViewBucket(context.Background(), bucketName)
	require.NoError(t, err)
	require.Equal(t, bucketName, viewBucketResult.Name)
	require.Equal(t, createBucketResult.ID, viewBucketResult.ID)
}

func TestIPCListBuckets(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	// Create 10 buckets
	bucketNames := make([]string, 10)
	for i := 0; i < 10; i++ {
		name := randomBucketName(t, 10)
		bucketNames[i] = name
		createBucketResult, err := ipc.CreateBucket(context.Background(), name)
		require.NoError(t, err)
		require.Equal(t, name, createBucketResult.Name)
	}

	t.Run("list all", func(t *testing.T) {
		bucketList, err := ipc.ListBuckets(context.Background(), 0, 20)
		require.NoError(t, err)
		require.Len(t, bucketList, 10)
		for i, b := range bucketList {
			require.Equal(t, bucketNames[i], b.Name)
		}
	})

	t.Run("list paginated", func(t *testing.T) {
		pageSize := 4
		pages := [][]string{}
		for offset := 0; offset < 10; offset += pageSize {
			bucketList, err := ipc.ListBuckets(context.Background(), int64(offset), int64(pageSize))
			require.NoError(t, err)
			end := min(offset+pageSize, 10)
			require.Len(t, bucketList, end-offset)
			for i, b := range bucketList {
				require.Equal(t, bucketNames[offset+i], b.Name)
			}
			pages = append(pages, bucketNames[offset:end])
		}
		// Check that we have 3 pages: [0-3], [4-7], [8-9]
		require.Len(t, pages, 3)
		for i, page := range pages {
			if i < 2 {
				require.Len(t, page, 4)
			} else {
				require.Len(t, page, 2)
			}
		}
	})
}

func TestIPCDeleteBucket(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	bucketName := randomBucketName(t, 10)

	createBucketResult, err := ipc.CreateBucket(context.Background(), bucketName)
	require.NoError(t, err)
	require.Equal(t, bucketName, createBucketResult.Name)

	require.NoError(t, ipc.DeleteBucket(context.Background(), bucketName))
	bucketList, err := ipc.ListBuckets(context.Background(), 0, 10)
	require.NoError(t, err)
	require.Len(t, bucketList, 0)
}

func TestIPCFileInfo(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	file := bytes.NewBuffer(testrand.BytesD(t, 1, memory.MB.ToInt64()))

	bucketName := randomBucketName(t, 10)
	fileName := randomBucketName(t, 10)

	_, err = ipc.CreateBucket(context.Background(), bucketName)
	require.NoError(t, err)

	fileUpload, err := ipc.CreateFileUpload(context.Background(), bucketName, fileName)
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	upResult, err := ipc.Upload(ctx, fileUpload, file)
	require.NoError(t, err)
	assert.Equal(t, upResult.Name, fileName)

	info, err := ipc.FileInfo(ctx, bucketName, fileName)
	require.NoError(t, err)
	assert.False(t, info.IsPublic)
	assert.Equal(t, fileName, info.Name)
	assert.Equal(t, bucketName, info.BucketName)
	assert.GreaterOrEqual(t, info.EncodedSize, memory.MB.ToInt64())
}

func TestIPCListFiles(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bucketName := randomBucketName(t, 10)
	_, err = ipc.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	fileNames := make([]string, 10)
	for i := 0; i < 10; i++ {
		file := bytes.NewBuffer(testrand.BytesD(t, 1, int64(i+1)*memory.MB.ToInt64()))
		fileName := randomBucketName(t, 10)
		fileNames[i] = fileName

		fileUpload, err := ipc.CreateFileUpload(context.Background(), bucketName, fileName)
		require.NoError(t, err)

		upResult, err := ipc.Upload(ctx, fileUpload, file)
		require.NoError(t, err)
		assert.Equal(t, upResult.Name, fileName)
	}

	t.Run("list all", func(t *testing.T) {
		list, err := ipc.ListFiles(context.Background(), bucketName, 0, 20)
		require.NoError(t, err)
		assert.Len(t, list, 10)
		for i, f := range list {
			assert.Equal(t, fileNames[i], f.Name)
		}
	})

	t.Run("list paginated", func(t *testing.T) {
		pageSize := 4
		pages := [][]string{}
		for offset := 0; offset < 10; offset += pageSize {
			list, err := ipc.ListFiles(context.Background(), bucketName, int64(offset), int64(pageSize))
			require.NoError(t, err)
			end := min(offset+pageSize, 10)
			assert.Len(t, list, end-offset)
			for i, f := range list {
				assert.Equal(t, fileNames[offset+i], f.Name)
			}
			pages = append(pages, fileNames[offset:end])
		}
		// Check that we have 3 pages: [0-3], [4-7], [8-9]
		assert.Len(t, pages, 3)
		for i, page := range pages {
			if i < 2 {
				assert.Len(t, page, 4)
			} else {
				assert.Len(t, page, 2)
			}
		}
	})
}

func TestIPCFileDelete(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bucketName := randomBucketName(t, 10)
	_, err = ipc.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	file := bytes.NewBuffer(testrand.BytesD(t, 1, 5*memory.MB.ToInt64()))
	fileName := randomBucketName(t, 10)

	fileUpload, err := ipc.CreateFileUpload(ctx, bucketName, fileName)
	require.NoError(t, err)

	_, err = ipc.Upload(ctx, fileUpload, file)
	require.NoError(t, err)

	require.NoError(t, ipc.FileDelete(ctx, bucketName, fileName))

	list, err := ipc.ListFiles(ctx, bucketName, 0, 10)
	require.NoError(t, err)
	assert.Len(t, list, 0)
}

func TestIPCFileSetPublicAccess(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	file := bytes.NewBuffer(testrand.BytesD(t, 1, memory.MB.ToInt64()))

	bucketName := randomBucketName(t, 10)
	fileName := randomBucketName(t, 10)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err = ipc.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	fileUpload, err := ipc.CreateFileUpload(ctx, bucketName, fileName)
	require.NoError(t, err)

	upResult, err := ipc.Upload(ctx, fileUpload, file)
	require.NoError(t, err)
	assert.Equal(t, upResult.Name, fileName)

	info, err := ipc.FileInfo(ctx, bucketName, fileName)
	require.NoError(t, err)
	assert.False(t, info.IsPublic)

	require.NoError(t, ipc.FileSetPublicAccess(ctx, bucketName, fileName, true))

	info, err = ipc.FileInfo(ctx, bucketName, fileName)
	require.NoError(t, err)
	assert.True(t, info.IsPublic)
}

func TestIPCUploadDownload(t *testing.T) {
	tests := []struct {
		name     string
		fileSize int64 // Size in MB
	}{
		{"1 MB", 1},
		{"5 MB", 5},
		{"15 MB", 15},
		{"35 MB", 35},
		{"256 MB", 256},
	}

	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	t.Run("without encryption", func(t *testing.T) {
		akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk))
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, akave.Close())
		})

		ipc, err := akave.IPC()
		require.NoError(t, err)

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.Bytes(t, tc.fileSize*memory.MB.ToInt64())
				testUploadDownloadIPC(t, ipc, data, false)
			})
		}
	})

	t.Run("with encryption", func(t *testing.T) {
		akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk),
			sdk.WithEncryptionKey([]byte(secretKey)))
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, akave.Close())
		})

		ipc, err := akave.IPC()
		require.NoError(t, err)

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.BytesD(t, 2024, tc.fileSize*memory.MB.ToInt64())
				testUploadDownloadIPC(t, ipc, data, false)
			})
		}
	})
}

func TestIPCUploadDownloadWithErasureCode(t *testing.T) {
	tests := []struct {
		name     string
		fileSize int64 // Size in MB
	}{
		{"1 MB", 1},
		{"5 MB", 5},
		{"15 MB", 15},
		{"35 MB", 35},
		{"255 MB", 255},
	}

	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	t.Run("without encryption", func(t *testing.T) {
		akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk),
			sdk.WithErasureCoding(16))
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, akave.Close())
		})

		ipc, err := akave.IPC()
		require.NoError(t, err)

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.Bytes(t, tc.fileSize*memory.MB.ToInt64())
				testUploadDownloadIPC(t, ipc, data, true)
			})
		}
	})

	t.Run("with encryption", func(t *testing.T) {
		akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk),
			sdk.WithErasureCoding(16), sdk.WithEncryptionKey([]byte(secretKey)))
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, akave.Close())
		})

		ipc, err := akave.IPC()
		require.NoError(t, err)

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				data := testrand.Bytes(t, tc.fileSize*memory.MB.ToInt64())
				testUploadDownloadIPC(t, ipc, data, true)
			})
		}
	})
}

func TestIPCRangeFileDownload(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]
	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})
	ipc, err := akave.IPC()
	require.NoError(t, err)

	ctx := context.Background()
	fileSize := 80 * memory.MB.ToInt64()

	fileData := testrand.BytesD(t, 2024, fileSize)

	bucketName := randomBucketName(t, 10)
	fileName := randomBucketName(t, 10)

	_, err = ipc.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	fileUpload, err := ipc.CreateFileUpload(ctx, bucketName, fileName)
	require.NoError(t, err)

	_, err = ipc.Upload(ctx, fileUpload, bytes.NewBuffer(fileData))
	require.NoError(t, err)

	var downloaded bytes.Buffer
	fileDownload, err := ipc.CreateRangeFileDownload(ctx, bucketName, fileName, 1, 3)
	require.NoError(t, err)
	require.True(t, len(fileDownload.Chunks) == 2)
	assert.True(t, fileDownload.Chunks[0].Index == 1)
	assert.True(t, fileDownload.Chunks[1].Index == 2)

	// download file blocks
	err = ipc.Download(ctx, fileDownload, &downloaded)
	require.NoError(t, err)

	chunks, blocks, bytesCount := fileDownload.Stats()
	t.Logf("Range download stats - Chunks: %d, Blocks: %d, Bytes: %d", chunks, blocks, bytesCount)

	assert.True(t, chunks > 0, "chunks counter should be greater than 0")
	assert.True(t, blocks > 0, "blocks counter should be greater than 0")
	assert.True(t, bytesCount > 0, "bytes counter should be greater than 0")

	assert.Equal(t, int64(len(fileDownload.Chunks)), chunks, "chunks counter should match number of chunks")
	assert.Equal(t, int64(2), chunks, "should have downloaded exactly 2 chunks for range 1-3")

	expectedMinBlocks := (chunks - 1) * 32
	assert.True(t, blocks >= expectedMinBlocks, "blocks counter (%d) should be >= (chunkCount-1)*32 (%d)", blocks, expectedMinBlocks)

	// check downloaded partial file contents
	expected := fileData[32*int(sdk.BlockSize.ToInt64()):] // first chunk is skipped
	checkFileContents(t, 10, expected, downloaded.Bytes())
}

func TestIPCResumeUpload(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	data := testrand.BytesD(t, 2024, 70*memory.MB.ToInt64())

	bucketName := randomBucketName(t, 10)
	fileName := randomBucketName(t, 10)

	ctx := context.Background()

	_, err = ipc.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	fileUpload, err := ipc.CreateFileUpload(ctx, bucketName, fileName)
	require.NoError(t, err)

	reader := &FailingReader{
		data:           data,
		failAfterBytes: 65 * memory.MB.ToInt64(), // Fail after 65MB
	}

	// Start upload - this should fail partway through
	_, err = ipc.Upload(ctx, fileUpload, reader)
	require.Error(t, err, "Upload should fail")
	require.Contains(t, err.Error(), "simulated read failure", "Error should be due to simulated failure")

	// Verify that some data was uploaded (reader position should have advanced)
	require.Greater(t, reader.pos, int64(0), "Some data should have been read before cancellation")
	require.Less(t, reader.pos, int64(len(data)), "Not all data should have been read due to cancellation")

	t.Logf("Upload was canceled after reading %d bytes out of %d total", reader.pos, len(data))

	// Now resume the upload with the same fileUpload and new reader with original data
	resumeResult, err := ipc.Upload(ctx, fileUpload, bytes.NewBuffer(data))
	require.NoError(t, err)

	assert.Equal(t, fileName, resumeResult.Name)
	assert.GreaterOrEqual(t, resumeResult.EncodedSize, int64(len(data)))

	// Verify the complete file by downloading it
	var downloaded bytes.Buffer
	fileDownload, err := ipc.CreateFileDownload(ctx, bucketName, fileName)
	require.NoError(t, err)

	err = ipc.Download(ctx, fileDownload, &downloaded)
	require.NoError(t, err)

	// Check that the downloaded data matches the original
	checkFileContents(t, len(data), data, downloaded.Bytes())

	_, err = ipc.Upload(ctx, fileUpload, bytes.NewBuffer(data))
	require.Error(t, err)
	require.Contains(t, err.Error(), "file is already committed")
}

func testUploadDownloadIPC(t *testing.T, ipc *sdk.IPC, data []byte, erasureCoding bool) {
	file := bytes.NewBuffer(data)
	fileSize := len(file.Bytes())

	bucketName := randomBucketName(t, 10)
	fileName := randomBucketName(t, 10)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := ipc.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	now := time.Now()
	fileUpload, err := ipc.CreateFileUpload(ctx, bucketName, fileName)
	require.NoError(t, err)
	fileUploadDuration := time.Since(now)
	t.Logf("Create file upload duration: %v", fileUploadDuration)

	now = time.Now()
	u, err := ipc.Upload(ctx, fileUpload, file)
	require.NoError(t, err)
	require.Equal(t, u.Size, int64(fileSize))

	t.Logf("Upload duration: %v", time.Since(now))

	var downloaded bytes.Buffer
	fileDownload, err := ipc.CreateFileDownload(ctx, bucketName, fileName)
	require.NoError(t, err)
	assert.True(t, len(fileDownload.Chunks) > 0)
	if erasureCoding {
		require.Equal(t, ceilDiv(u.EncodedSize, 32*memory.MB.ToInt64()), int64(len(fileDownload.Chunks)))
	}

	now = time.Now()
	require.NoError(t, ipc.Download(ctx, fileDownload, &downloaded))
	t.Logf("Download duration: %v", time.Since(now))

	chunks, blocks, bytesCount := fileDownload.Stats()
	t.Logf("Download stats - Chunks: %d, Blocks: %d, Bytes: %d", chunks, blocks, bytesCount)

	assert.True(t, chunks > 0, "chunks counter should be greater than 0")
	assert.True(t, blocks > 0, "blocks counter should be greater than 0")
	assert.True(t, bytesCount > 0, "bytes counter should be greater than 0")

	assert.Equal(t, int64(len(fileDownload.Chunks)), chunks, "chunks counter should match number of chunks")

	if erasureCoding {
		expectedBlocks := chunks * 32
		assert.Equal(t, expectedBlocks, blocks, "with erasure coding, blocks counter should be exactly chunkCount*32")
	} else {
		expectedMinBlocks := (chunks - 1) * 32
		assert.True(t, blocks >= expectedMinBlocks, "blocks counter (%d) should be >= (chunkCount-1)*32 (%d)", blocks, expectedMinBlocks)
	}

	// Some overhead is added to real data during upload
	assert.GreaterOrEqual(t, bytesCount, int64(len(data)), "bytes counter should be slightly larger than data size")

	checkFileContents(t, 10, data, downloaded.Bytes())
}

func ceilDiv(a, b int64) int64 {
	return (a + b - 1) / b
}

// FailingReader simulates a reader that fails after reading a specific number of bytes.
type FailingReader struct {
	data           []byte
	pos            int64
	failAfterBytes int64
}

func (r *FailingReader) Read(p []byte) (n int, err error) {
	if r.pos >= int64(len(r.data)) {
		return 0, io.EOF
	}

	// Calculate how much to read
	toRead := len(p)
	if r.pos+int64(toRead) > int64(len(r.data)) {
		toRead = int(int64(len(r.data)) - r.pos)
	}

	copy(p, r.data[r.pos:r.pos+int64(toRead)])
	r.pos += int64(toRead)

	// Check if we should fail
	if r.pos >= r.failAfterBytes {
		r.pos = r.failAfterBytes // Ensure we don't read more than intended
		return 0, fmt.Errorf("simulated read failure after %d bytes", r.pos)
	}
	return toRead, nil
}
