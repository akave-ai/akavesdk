// Copyright (C) 2026 Akave
// See LICENSE for copying information.

package sdk_test

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/ipctest"
	"github.com/akave-ai/akavesdk/private/memory"
	"github.com/akave-ai/akavesdk/private/testrand"
	"github.com/akave-ai/akavesdk/sdk"
)

func TestMultiUploadCreateFileUploadsInvalidParams(t *testing.T) {
	t.Run("function returns error when empty params", func(t *testing.T) {
		multiUpload := &sdk.MultiUpload{}
		_, err := multiUpload.CreateFileUploads(context.Background(), []sdk.CreateFileUploadParam{})
		require.Error(t, err)
		require.Contains(t, err.Error(), "no files to upload")
	})

	t.Run("function returns error when bucket name is empty", func(t *testing.T) {
		multiUpload := &sdk.MultiUpload{}
		params := []sdk.CreateFileUploadParam{
			{
				BucketName: "",
				FileName:   "file.txt",
			},
		}
		_, err := multiUpload.CreateFileUploads(context.Background(), params)
		require.Error(t, err)
		require.Contains(t, err.Error(), "empty bucket name at index 0")
	})

	t.Run("function returns error when file name is empty", func(t *testing.T) {
		multiUpload := &sdk.MultiUpload{}
		params := []sdk.CreateFileUploadParam{
			{
				BucketName: "my-bucket",
				FileName:   "",
			},
		}
		_, err := multiUpload.CreateFileUploads(context.Background(), params)
		require.Error(t, err)
		require.Contains(t, err.Error(), "empty file name at index 0")
	})

	t.Run("function returns error when multiple params have empty bucket names", func(t *testing.T) {
		multiUpload := &sdk.MultiUpload{}
		params := []sdk.CreateFileUploadParam{
			{
				BucketName: "",
				FileName:   "file1.txt",
			},
			{
				BucketName: "my-bucket",
				FileName:   "file2.txt",
			},
			{
				BucketName: "",
				FileName:   "file3.txt",
			},
		}
		_, err := multiUpload.CreateFileUploads(context.Background(), params)
		require.Error(t, err)
		require.Contains(t, err.Error(), "empty bucket name at index 0")
		require.Contains(t, err.Error(), "empty bucket name at index 2")
	})

	t.Run("function returns error when multiple params have empty file names", func(t *testing.T) {
		multiUpload := &sdk.MultiUpload{}
		params := []sdk.CreateFileUploadParam{
			{
				BucketName: "my-bucket",
				FileName:   "",
			},
			{
				BucketName: "my-bucket",
				FileName:   "file2.txt",
			},
			{
				BucketName: "my-bucket",
				FileName:   "",
			},
		}
		_, err := multiUpload.CreateFileUploads(context.Background(), params)
		require.Error(t, err)
		require.Contains(t, err.Error(), "empty file name at index 0")
		require.Contains(t, err.Error(), "empty file name at index 2")
	})

	t.Run("function returns error when mixed empty bucket and file names", func(t *testing.T) {
		multiUpload := &sdk.MultiUpload{}
		params := []sdk.CreateFileUploadParam{
			{
				BucketName: "",
				FileName:   "",
			},
			{
				BucketName: "my-bucket",
				FileName:   "file.txt",
			},
			{
				BucketName: "",
				FileName:   "file2.txt",
			},
		}
		_, err := multiUpload.CreateFileUploads(context.Background(), params)
		require.Error(t, err)
		require.Contains(t, err.Error(), "empty bucket name at index 0")
		require.Contains(t, err.Error(), "empty file name at index 0")
		require.Contains(t, err.Error(), "empty bucket name at index 2")
	})
}

func TestMultiUploadUploadInvalidParams(t *testing.T) {
	t.Run("function returns error when empty params", func(t *testing.T) {
		multiUpload := &sdk.MultiUpload{}
		_, err := multiUpload.Upload(context.Background(), []sdk.UploadParam{})
		require.Error(t, err)
		require.Contains(t, err.Error(), "no files to upload")
	})

	t.Run("function returns error when FileUpload is nil", func(t *testing.T) {
		multiUpload := &sdk.MultiUpload{}
		params := []sdk.UploadParam{
			{
				FileUpload: nil,
				Reader:     bytes.NewReader([]byte("test data")),
			},
		}
		_, err := multiUpload.Upload(context.Background(), params)
		require.Error(t, err)
		require.Contains(t, err.Error(), "nil FileUpload at index 0")
	})

	t.Run("function returns error when Reader is nil", func(t *testing.T) {
		multiUpload := &sdk.MultiUpload{}
		mockFileUpload := &sdk.IPCFileUpload{}
		params := []sdk.UploadParam{
			{
				FileUpload: mockFileUpload,
				Reader:     nil,
			},
		}
		_, err := multiUpload.Upload(context.Background(), params)
		require.Error(t, err)
		require.Contains(t, err.Error(), "nil Reader at index 0")
	})

	t.Run("function returns error when multiple FileUploads are nil", func(t *testing.T) {
		multiUpload := &sdk.MultiUpload{}
		mockFileUpload := &sdk.IPCFileUpload{}
		params := []sdk.UploadParam{
			{
				FileUpload: nil,
				Reader:     bytes.NewReader([]byte("test data")),
			},
			{
				FileUpload: mockFileUpload,
				Reader:     bytes.NewReader([]byte("test data")),
			},
			{
				FileUpload: nil,
				Reader:     bytes.NewReader([]byte("test data")),
			},
		}
		_, err := multiUpload.Upload(context.Background(), params)
		require.Error(t, err)
		require.Contains(t, err.Error(), "nil FileUpload at index 0")
		require.Contains(t, err.Error(), "nil FileUpload at index 2")
	})

	t.Run("function returns error when multiple Readers are nil", func(t *testing.T) {
		multiUpload := &sdk.MultiUpload{}
		mockFileUpload := &sdk.IPCFileUpload{}
		params := []sdk.UploadParam{
			{
				FileUpload: mockFileUpload,
				Reader:     nil,
			},
			{
				FileUpload: mockFileUpload,
				Reader:     bytes.NewReader([]byte("test data")),
			},
			{
				FileUpload: mockFileUpload,
				Reader:     nil,
			},
		}
		_, err := multiUpload.Upload(context.Background(), params)
		require.Error(t, err)
		require.Contains(t, err.Error(), "nil Reader at index 0")
		require.Contains(t, err.Error(), "nil Reader at index 2")
	})

	t.Run("function returns error when mixed nil FileUpload and Reader", func(t *testing.T) {
		multiUpload := &sdk.MultiUpload{}
		mockFileUpload := &sdk.IPCFileUpload{}
		params := []sdk.UploadParam{
			{
				FileUpload: nil,
				Reader:     nil,
			},
			{
				FileUpload: mockFileUpload,
				Reader:     bytes.NewReader([]byte("test data")),
			},
			{
				FileUpload: mockFileUpload,
				Reader:     nil,
			},
		}
		_, err := multiUpload.Upload(context.Background(), params)
		require.Error(t, err)
		require.Contains(t, err.Error(), "nil FileUpload at index 0")
		require.Contains(t, err.Error(), "nil Reader at index 0")
		require.Contains(t, err.Error(), "nil Reader at index 2")
	})
}

type errorReader struct {
	err error
}

func (er *errorReader) Read(p []byte) (int, error) {
	return 0, er.err
}

func TestIPCMultiUploadUploadResultContainsErrors(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.PrivateKeyToHex(ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10)))
	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), pk)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	t.Run("result contains error when bucket does not exist", func(t *testing.T) {
		multiUpload := ipc.MultiUpload(1)

		params := []sdk.UploadParam{
			{
				FileUpload: &sdk.IPCFileUpload{
					BucketName: "non-existent-bucket",
					Name:       "file.txt",
				},
				Reader: bytes.NewReader([]byte("test data")),
			},
		}

		results, err := multiUpload.Upload(context.Background(), params)
		require.NoError(t, err)
		require.Len(t, results, 1)
		require.NotNil(t, results[0].Error, "expected error in result when bucket does not exist")
	})

	t.Run("result contains individual errors for each file upload failure", func(t *testing.T) {
		ctx := context.Background()

		bucketName := testrand.String(10)
		_, err := ipc.CreateBucket(ctx, bucketName)
		require.NoError(t, err)

		multiUpload := ipc.MultiUpload(2)

		fileUpload1, err := ipc.CreateFileUpload(ctx, bucketName, "file1.txt")
		require.NoError(t, err)

		params := []sdk.UploadParam{
			{
				FileUpload: fileUpload1,
				Reader:     bytes.NewReader([]byte("valid data")),
			},
			{
				FileUpload: &sdk.IPCFileUpload{
					BucketName: "invalid-bucket",
					Name:       "file2.txt",
				},
				Reader: bytes.NewReader([]byte("data")),
			},
		}

		results, err := multiUpload.Upload(ctx, params)
		require.NoError(t, err)
		require.Len(t, results, 2)
		require.Nil(t, results[0].Error, "expected no error for valid upload")
		require.NotNil(t, results[1].Error, "expected error for invalid bucket upload")
	})

	t.Run("result contains error when reader fails during upload", func(t *testing.T) {
		ctx := context.Background()

		bucketName := testrand.String(10)
		_, err := ipc.CreateBucket(ctx, bucketName)
		require.NoError(t, err)

		multiUpload := ipc.MultiUpload(1)

		fileUpload, err := ipc.CreateFileUpload(ctx, bucketName, "file-with-error.txt")
		require.NoError(t, err)

		failingReader := &errorReader{err: context.Canceled}

		params := []sdk.UploadParam{
			{
				FileUpload: fileUpload,
				Reader:     failingReader,
			},
		}

		results, err := multiUpload.Upload(ctx, params)
		require.NoError(t, err)
		require.Len(t, results, 1)
		require.NotNil(t, results[0].Error, "expected error in result when reader fails")
	})
}

func TestIPCMultiUploadCreateFileUploadsResultErrors(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.PrivateKeyToHex(ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10)))
	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), pk)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	t.Run("result contains error when creating file in non-existent bucket", func(t *testing.T) {
		multiUpload := ipc.MultiUpload(1)

		params := []sdk.CreateFileUploadParam{
			{
				BucketName: "non-existent-bucket-xyz",
				FileName:   "file1.txt",
			},
		}

		results, err := multiUpload.CreateFileUploads(context.Background(), params)
		require.NoError(t, err)
		require.Len(t, results, 1)
		require.NotNil(t, results[0].Error, "result should contain an error when bucket does not exist")
		require.Equal(t, params[0].BucketName, results[0].BucketName)
		require.Equal(t, params[0].FileName, results[0].FileName)
	})

	t.Run("results contain independent errors for each file creation failure", func(t *testing.T) {
		ctx := context.Background()

		validBucketName := testrand.String(10)
		_, err := ipc.CreateBucket(ctx, validBucketName)
		require.NoError(t, err)

		multiUpload := ipc.MultiUpload(2)

		params := []sdk.CreateFileUploadParam{
			{
				BucketName: validBucketName,
				FileName:   "valid-file.txt",
			},
			{
				BucketName: "invalid-bucket-xyz",
				FileName:   "file2.txt",
			},
			{
				BucketName: validBucketName,
				FileName:   "another-file.txt",
			},
		}

		results, err := multiUpload.CreateFileUploads(ctx, params)
		require.NoError(t, err)

		require.Len(t, results, 3)
		require.Nil(t, results[0].Error, "result[0] should not contain an error for valid bucket")
		require.Equal(t, validBucketName, results[0].BucketName)
		require.Equal(t, "valid-file.txt", results[0].FileName)
		require.NotNil(t, results[1].Error, "result[1] should contain an error for non-existent bucket")
		require.Equal(t, "invalid-bucket-xyz", results[1].BucketName)
		require.Equal(t, "file2.txt", results[1].FileName)
		require.Nil(t, results[2].Error, "result[2] should not contain an error for valid bucket")
		require.Equal(t, validBucketName, results[2].BucketName)
		require.Equal(t, "another-file.txt", results[2].FileName)
	})
}

func TestIPCMultiUploadWithEncryptionAndErasureCoding(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	nodeAddress := PickNodeRPCAddress(t)
	pk := ipctest.PrivateKeyToHex(ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10)))

	akave, err := sdk.New(
		nodeAddress,
		maxConcurrency,
		blockPartSize.ToInt64(),
		pk,
		sdk.WithErasureCoding(16),
		sdk.WithEncryptionKey([]byte(secretKey)),
	)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	ctx := t.Context()

	bucketName := testrand.String(10)
	createBucketResult, err := ipc.CreateBucket(ctx, bucketName)
	require.NoError(t, err)
	require.Equal(t, bucketName, createBucketResult.Name)

	const numFiles = 5
	fileSizes := []int64{
		100 * memory.MB.ToInt64(),
		200 * memory.MB.ToInt64(),
		1 * memory.MB.ToInt64(),
		5 * memory.MB.ToInt64(),
		50 * memory.MB.ToInt64(),
	}

	fileNames := make([]string, numFiles)
	fileData := make([][]byte, numFiles)
	for i := range numFiles {
		fileNames[i] = fmt.Sprintf("file-%d-%s", i, testrand.String(8))
		fileData[i] = testrand.BytesD(t, int64(2025+i), fileSizes[i])
	}

	multiUpload := ipc.MultiUpload(3)

	createParams := make([]sdk.CreateFileUploadParam, numFiles)
	for i := range numFiles {
		createParams[i] = sdk.CreateFileUploadParam{
			BucketName: bucketName,
			FileName:   fileNames[i],
		}
	}

	var fileUploads []*sdk.IPCFileUpload
	var fileMetas []sdk.IPCFileMetaV2

	t.Run("create file uploads", func(t *testing.T) {
		createResults, err := multiUpload.CreateFileUploads(ctx, createParams)
		require.NoError(t, err)
		require.Len(t, createResults, numFiles)

		fileUploads = make([]*sdk.IPCFileUpload, numFiles)
		for i, result := range createResults {
			require.NoError(t, result.Error)
			require.NotNil(t, result.FileUpload)
			require.Equal(t, fileNames[i], result.FileName)
			require.Equal(t, bucketName, result.BucketName)
			fileUploads[i] = result.FileUpload
		}
	})

	t.Run("upload files in parallel", func(t *testing.T) {
		uploadParams := make([]sdk.UploadParam, numFiles)
		for i := range numFiles {
			uploadParams[i] = sdk.UploadParam{
				FileUpload: fileUploads[i],
				Reader:     bytes.NewReader(fileData[i]),
			}
		}

		uploadResults, err := multiUpload.Upload(ctx, uploadParams)
		require.NoError(t, err)
		require.Len(t, uploadResults, numFiles)

		fileMetas = make([]sdk.IPCFileMetaV2, numFiles)
		for i, result := range uploadResults {
			require.NoError(t, result.Error)
			require.Equal(t, fileNames[i], result.FileName)
			require.Equal(t, bucketName, result.BucketName)
			require.Equal(t, fileSizes[i], result.Meta.Size)
			require.Greater(t, result.Meta.EncodedSize, result.Meta.Size)
			require.NotEmpty(t, result.Meta.RootCID)
			fileMetas[i] = *result.Meta
		}
	})

	t.Run("download and verify files", func(t *testing.T) {
		for i := range numFiles {
			fileDownload, err := ipc.CreateFileDownload(ctx, bucketName, fileNames[i])
			require.NoError(t, err)
			require.Equal(t, fileNames[i], fileDownload.Name)

			var buf bytes.Buffer
			err = ipc.Download(ctx, fileDownload, &buf)
			require.NoError(t, err)

			checkFileContents(t, 200, fileData[i], buf.Bytes())
		}
	})
}
