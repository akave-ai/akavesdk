// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package sdk_test

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/memory"
	"github.com/akave-ai/akavesdk/private/testrand"
	"github.com/akave-ai/akavesdk/sdk"
)

func TestCreateBucketIPC(t *testing.T) {
	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(PickPrivateKey(t)))
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

func TestViewBucketIPC(t *testing.T) {
	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(PickPrivateKey(t)))
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

func TestListBucketsIPC(t *testing.T) {
	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(PickPrivateKey(t)))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	bucketName := randomBucketName(t, 10)
	bucketName2 := randomBucketName(t, 10)

	createBucketResult, err := ipc.CreateBucket(context.Background(), bucketName)
	require.NoError(t, err)
	require.Equal(t, bucketName, createBucketResult.Name)

	createBucketResult2, err := ipc.CreateBucket(context.Background(), bucketName2)
	require.NoError(t, err)
	require.Equal(t, bucketName2, createBucketResult2.Name)

	bucketList, err := ipc.ListBuckets(context.Background())
	require.NoError(t, err)
	require.Len(t, bucketList, 2)
}

func TestDeleteBucketIPC(t *testing.T) {
	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(PickPrivateKey(t)))
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
	bucketList, err := ipc.ListBuckets(context.Background())
	require.NoError(t, err)
	require.Len(t, bucketList, 0)
}

func TestFileInfo(t *testing.T) {
	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(PickPrivateKey(t)))
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

	require.NoError(t, ipc.CreateFileUpload(context.Background(), bucketName, fileName))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	upResult, err := ipc.Upload(ctx, bucketName, fileName, file)
	require.NoError(t, err)
	assert.Equal(t, upResult.Name, fileName)

	info, err := ipc.FileInfo(ctx, bucketName, fileName)
	require.NoError(t, err)
	assert.False(t, info.IsPublic)
	assert.Equal(t, fileName, info.Name)
	assert.Equal(t, bucketName, info.BucketName)
	assert.Equal(t, file.Len(), info.EncodedSize)
}

func TestListFiles(t *testing.T) {
	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(PickPrivateKey(t)))
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

	for i := 0; i < 5; i++ {
		file := bytes.NewBuffer(testrand.BytesD(t, 1, int64(i+1)*memory.MB.ToInt64()))
		fileName := randomBucketName(t, 10)

		require.NoError(t, ipc.CreateFileUpload(context.Background(), bucketName, fileName))

		upResult, err := ipc.Upload(ctx, bucketName, fileName, file)
		require.NoError(t, err)
		assert.Equal(t, upResult.Name, fileName)
	}

	list, err := ipc.ListFiles(context.Background(), bucketName)
	require.NoError(t, err)
	assert.Len(t, list, 5)
}

func TestFileDeleteIPC(t *testing.T) {
	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(PickPrivateKey(t)))
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

	require.NoError(t, ipc.CreateFileUpload(ctx, bucketName, fileName))

	_, err = ipc.Upload(ctx, bucketName, fileName, file)
	require.NoError(t, err)

	require.NoError(t, ipc.FileDelete(ctx, bucketName, fileName))

	list, err := ipc.ListFiles(ctx, bucketName)
	require.NoError(t, err)
	assert.Len(t, list, 0)
}

func TestFileSetPublicAccess(t *testing.T) {
	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(PickPrivateKey(t)))
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

	require.NoError(t, ipc.CreateFileUpload(ctx, bucketName, fileName))

	upResult, err := ipc.Upload(ctx, bucketName, fileName, file)
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
			data := testrand.Bytes(t, tc.fileSize*memory.MB.ToInt64())
			testUploadDownloadIPC(t, ipc, data)
		})
	}
}

func testUploadDownloadIPC(t *testing.T, ipc *sdk.IPC, data []byte) {
	file := bytes.NewBuffer(data)

	bucketName := randomBucketName(t, 10)
	fileName := randomBucketName(t, 10)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := ipc.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	now := time.Now()
	require.NoError(t, ipc.CreateFileUpload(ctx, bucketName, fileName))
	fileUploadDuration := time.Since(now)
	t.Logf("Create file upload duration: %v", fileUploadDuration)

	now = time.Now()
	_, err = ipc.Upload(ctx, bucketName, fileName, file)
	require.NoError(t, err)
	t.Logf("Upload duration: %v", time.Since(now))

	var downloaded bytes.Buffer
	fileDownload, err := ipc.CreateFileDownload(ctx, bucketName, fileName)
	require.NoError(t, err)
	assert.True(t, len(fileDownload.Chunks) > 0)

	now = time.Now()
	require.NoError(t, ipc.Download(ctx, fileDownload, &downloaded))
	t.Logf("Download duration: %v", time.Since(now))

	checkFileContents(t, 10, data, downloaded.Bytes())
}

func TestUploadDownloadIPCWithEncryption(t *testing.T) {
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

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true,
		sdk.WithPrivateKey(PickPrivateKey(t)),
		sdk.WithEncryptionKey([]byte(secretKey)),
		sdk.WithMetadataEncryption(),
	)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := testrand.Bytes(t, tc.fileSize*memory.MB.ToInt64())
			testUploadDownloadIPC(t, ipc, data)
		})
	}
}
