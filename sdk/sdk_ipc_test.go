// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package sdk_test

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/akave-ai/akavesdk/private/encryption"
	"github.com/akave-ai/akavesdk/private/ipc/contracts"
	"github.com/akave-ai/akavesdk/private/ipctest"
	"github.com/akave-ai/akavesdk/private/memory"
	"github.com/akave-ai/akavesdk/private/pb"
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

	bucketName := testrand.String(10)

	createBucketResult, err := ipc.CreateBucket(t.Context(), bucketName)
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

	bucketName := testrand.String(10)

	createBucketResult, err := ipc.CreateBucket(t.Context(), bucketName)
	require.NoError(t, err)
	require.Equal(t, bucketName, createBucketResult.Name)

	viewBucketResult, err := ipc.ViewBucket(t.Context(), bucketName)
	require.NoError(t, err)
	require.Equal(t, bucketName, viewBucketResult.Name)
	require.Equal(t, createBucketResult.ID, viewBucketResult.ID)
}

func TestIPCViewBucketWithEncryption(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]
	address := PickNodeRPCAddress(t)

	akave, err := sdk.New(address, maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk),
		sdk.WithEncryptionKey([]byte(secretKey)), sdk.WithMetadataEncryption())
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	bucketName := testrand.String(10)

	createBucketResult, err := ipc.CreateBucket(t.Context(), bucketName)
	require.NoError(t, err)
	require.Equal(t, bucketName, createBucketResult.Name)

	bucketNameEnc := encryptMetadataHex(t, bucketName, "bucket")

	storage := getStorage(t, address)

	bucket, err := storage.GetBucketByName(&bind.CallOpts{Context: t.Context()}, bucketNameEnc, crypto.PubkeyToAddress(pk.PublicKey),
		big.NewInt(0), big.NewInt(0))
	require.NoError(t, err)
	require.Equal(t, bucketNameEnc, bucket.Name)

	viewBucketResult, err := ipc.ViewBucket(t.Context(), bucketName)
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
		name := testrand.String(10)
		bucketNames[i] = name
		createBucketResult, err := ipc.CreateBucket(t.Context(), name)
		require.NoError(t, err)
		require.Equal(t, name, createBucketResult.Name)
	}

	t.Run("list all", func(t *testing.T) {
		bucketList, err := ipc.ListBuckets(t.Context(), 0, 20)
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
			bucketList, err := ipc.ListBuckets(t.Context(), int64(offset), int64(pageSize))
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

func TestIPCListBucketWithEncryption(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk),
		sdk.WithEncryptionKey([]byte(secretKey)), sdk.WithMetadataEncryption())
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	bucketNames := make([]string, 10)
	for i := 0; i < 10; i++ {
		name := testrand.String(10)
		bucketNames[i] = name
		createBucketResult, err := ipc.CreateBucket(t.Context(), name)
		require.NoError(t, err)
		require.Equal(t, name, createBucketResult.Name)
	}

	t.Run("list all", func(t *testing.T) {
		bucketList, err := ipc.ListBuckets(t.Context(), 0, 20)
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
			bucketList, err := ipc.ListBuckets(t.Context(), int64(offset), int64(pageSize))
			require.NoError(t, err)
			end := min(offset+pageSize, 10)
			require.Len(t, bucketList, end-offset)
			for i, b := range bucketList {
				require.Equal(t, bucketNames[offset+i], b.Name)
			}
			pages = append(pages, bucketNames[offset:end])
		}
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

func TestIPCListBucketsWithDifferentEncKeys(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	akaveWithEncryption, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk),
		sdk.WithEncryptionKey([]byte(secretKey)), sdk.WithMetadataEncryption())
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akaveWithEncryption.Close())
	})

	akaveWithDifferentSecretKey, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk),
		sdk.WithEncryptionKey([]byte("N1PCdw3M2B1TfJhoaY2mL736p2vCUc46")), sdk.WithMetadataEncryption())
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akaveWithDifferentSecretKey.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	ipcEnc, err := akaveWithEncryption.IPC()
	require.NoError(t, err)

	ipcEncDiff, err := akaveWithDifferentSecretKey.IPC()
	require.NoError(t, err)

	bucketNames := make([]string, 5)
	for i := 0; i < 5; i++ {
		name := testrand.String(10)
		bucketNames[i] = name
		createBucketResult, err := ipc.CreateBucket(t.Context(), name)
		require.NoError(t, err)
		require.Equal(t, name, createBucketResult.Name)
	}

	bucketNamesEncrypted := make([]string, 5)
	for i := 0; i < 5; i++ {
		name := testrand.String(10)
		bucketNamesEncrypted[i] = name
		createBucketResult, err := ipcEnc.CreateBucket(t.Context(), name)
		require.NoError(t, err)
		require.Equal(t, name, createBucketResult.Name)
	}

	bucketNamesEncryptedDiff := make([]string, 5)
	for i := 0; i < 5; i++ {
		name := testrand.String(10)
		bucketNamesEncryptedDiff[i] = name
		createBucketResult, err := ipcEncDiff.CreateBucket(t.Context(), name)
		require.NoError(t, err)
		require.Equal(t, name, createBucketResult.Name)
	}

	list, err := ipcEnc.ListBuckets(t.Context(), 0, 20)
	require.NoError(t, err)
	require.Len(t, list, 15)
	require.True(t, existsInBuckets(bucketNames, list))
	require.True(t, existsInBuckets(bucketNamesEncrypted, list))
	require.False(t, existsInBuckets(bucketNamesEncryptedDiff, list))
}

func existsInBuckets(bucketNames []string, buckets []sdk.IPCBucket) bool {
	for _, bucketName := range bucketNames {
		found := false
		for _, bucket := range buckets {
			if bucket.Name == bucketName {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
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

	bucketName := testrand.String(10)

	createBucketResult, err := ipc.CreateBucket(t.Context(), bucketName)
	require.NoError(t, err)
	require.Equal(t, bucketName, createBucketResult.Name)

	require.NoError(t, ipc.DeleteBucket(t.Context(), bucketName))
	bucketList, err := ipc.ListBuckets(t.Context(), 0, 10)
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

	bucketName := testrand.String(10)
	fileName := testrand.String(10)

	_, err = ipc.CreateBucket(t.Context(), bucketName)
	require.NoError(t, err)

	fileUpload, err := ipc.CreateFileUpload(t.Context(), bucketName, fileName)
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(t.Context())
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

func TestIPCFileInfoWithEncryption(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]
	address := PickNodeRPCAddress(t)

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk),
		sdk.WithMetadataEncryption(), sdk.WithEncryptionKey([]byte(secretKey)))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	file := bytes.NewBuffer(testrand.BytesD(t, 1, memory.MB.ToInt64()))

	bucketName := testrand.String(10)
	fileName := testrand.String(10)

	_, err = ipc.CreateBucket(t.Context(), bucketName)
	require.NoError(t, err)

	fileUpload, err := ipc.CreateFileUpload(t.Context(), bucketName, fileName)
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	_, err = ipc.Upload(ctx, fileUpload, file)
	require.NoError(t, err)

	info, err := ipc.FileInfo(ctx, bucketName, fileName)
	require.NoError(t, err)
	assert.False(t, info.IsPublic)
	assert.Equal(t, fileName, info.Name)
	assert.Equal(t, bucketName, info.BucketName)
	assert.GreaterOrEqual(t, info.EncodedSize, memory.MB.ToInt64())

	bucketNameEnc := encryptMetadataHex(t, bucketName, "bucket")
	fileNameEnc := encryptMetadataHex(t, fileName, bucketName)

	storage := getStorage(t, address)

	bucket, err := storage.GetBucketByName(&bind.CallOpts{Context: t.Context()}, bucketNameEnc, crypto.PubkeyToAddress(pk.PublicKey),
		big.NewInt(0), big.NewInt(0))
	require.NoError(t, err)

	fileStored, err := storage.GetFileByName(&bind.CallOpts{Context: t.Context()}, bucket.Id, fileNameEnc)
	require.NoError(t, err)
	require.Equal(t, fileNameEnc, fileStored.Name)
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

	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	bucketName := testrand.String(10)
	_, err = ipc.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	fileNames := make([]string, 10)
	for i := 0; i < 10; i++ {
		file := bytes.NewBuffer(testrand.BytesD(t, 1, int64(i+1)*memory.MB.ToInt64()))
		fileName := testrand.String(11)
		fileNames[i] = fileName

		fileUpload, err := ipc.CreateFileUpload(ctx, bucketName, fileName)
		require.NoError(t, err)

		upResult, err := ipc.Upload(ctx, fileUpload, file)
		require.NoError(t, err)
		assert.Equal(t, upResult.Name, fileName)
	}

	t.Run("list all", func(t *testing.T) {
		list, err := ipc.ListFiles(ctx, bucketName, 0, 20)
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
			list, err := ipc.ListFiles(ctx, bucketName, int64(offset), int64(pageSize))
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

func TestIPCListFilesWithEncryption(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk),
		sdk.WithMetadataEncryption(), sdk.WithEncryptionKey([]byte(secretKey)))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	bucketName := testrand.String(10)
	_, err = ipc.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	fileNames := make([]string, 10)
	for i := 0; i < 10; i++ {
		file := bytes.NewBuffer(testrand.BytesD(t, 1, int64(i+1)*memory.MB.ToInt64()))
		fileName := testrand.String(11)
		fileNames[i] = fileName

		fileUpload, err := ipc.CreateFileUpload(ctx, bucketName, fileName)
		require.NoError(t, err)

		_, err = ipc.Upload(ctx, fileUpload, file)
		require.NoError(t, err)
	}

	t.Run("list all", func(t *testing.T) {
		list, err := ipc.ListFiles(ctx, bucketName, 0, 20)
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
			list, err := ipc.ListFiles(ctx, bucketName, int64(offset), int64(pageSize))
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

	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	bucketName := testrand.String(10)
	_, err = ipc.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	file := bytes.NewBuffer(testrand.BytesD(t, 1, 5*memory.MB.ToInt64()))
	fileName := testrand.String(10)

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

	bucketName := testrand.String(10)
	fileName := testrand.String(10)

	ctx, cancel := context.WithCancel(t.Context())
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

func TestIPCMetadataEncryption(t *testing.T) {
	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]
	address := PickNodeRPCAddress(t)

	akave, err := sdk.New(
		address,
		maxConcurrency,
		blockPartSize.ToInt64(),
		true,
		sdk.WithPrivateKey(newPk),
		sdk.WithEncryptionKey([]byte(secretKey)),
		sdk.WithMetadataEncryption(),
	)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	ctx := t.Context()

	bucketName := testrand.String(10)
	fileName := testrand.String(12)
	fileData := testrand.BytesD(t, 2026, 5*memory.MB.ToInt64())

	bucketNameEnc := encryptMetadataHex(t, bucketName, "bucket")
	fileNameEnc := encryptMetadataHex(t, fileName, bucketName)

	var (
		createBucketResult *sdk.IPCBucketCreateResult
		fileUpload         *sdk.IPCFileUpload
		fileMeta           sdk.IPCFileMeta
		fileDownload       sdk.IPCFileDownload
		rangeDownload      sdk.IPCFileDownload
	)

	storage := getStorage(t, address)

	t.Run("CreateBucket", func(t *testing.T) {
		createBucketResult, err = ipc.CreateBucket(ctx, bucketName)
		require.NoError(t, err)
		require.Equal(t, bucketName, createBucketResult.Name)

		bucket, err := storage.GetBucketByName(&bind.CallOpts{Context: t.Context()}, bucketNameEnc, crypto.PubkeyToAddress(pk.PublicKey),
			big.NewInt(0), big.NewInt(0))
		require.NoError(t, err)
		require.Equal(t, bucketNameEnc, bucket.Name)
	})

	t.Run("ViewBucket", func(t *testing.T) {
		viewBucketResult, err := ipc.ViewBucket(ctx, bucketName)
		require.NoError(t, err)
		require.Equal(t, bucketName, viewBucketResult.Name)
	})

	t.Run("ListBuckets", func(t *testing.T) {
		bucketList, err := ipc.ListBuckets(ctx, 0, 25)
		require.NoError(t, err)
		foundBucket := false
		for _, bucket := range bucketList {
			if bucket.Name == bucketName {
				foundBucket = true
				break
			}
		}
		require.True(t, foundBucket)
	})

	t.Run("CreateFileUpload", func(t *testing.T) {
		fileUpload, err = ipc.CreateFileUpload(ctx, bucketName, fileName)
		require.NoError(t, err)
		require.Equal(t, bucketNameEnc, fileUpload.BucketName)
		require.Equal(t, fileNameEnc, fileUpload.Name)

		bucket, err := storage.GetBucketByName(&bind.CallOpts{Context: t.Context()}, bucketNameEnc, crypto.PubkeyToAddress(pk.PublicKey),
			big.NewInt(0), big.NewInt(0))
		require.NoError(t, err)

		file, err := storage.GetFileByName(&bind.CallOpts{Context: t.Context()}, bucket.Id, fileNameEnc)
		require.NoError(t, err)
		require.Equal(t, fileNameEnc, file.Name)
	})

	t.Run("Upload", func(t *testing.T) {
		uploadResult, err := ipc.Upload(ctx, fileUpload, bytes.NewBuffer(fileData))
		require.NoError(t, err)
		require.Equal(t, bucketNameEnc, uploadResult.BucketName)
		require.Equal(t, fileNameEnc, uploadResult.Name)
	})

	t.Run("FileInfo", func(t *testing.T) {
		fileMeta, err = ipc.FileInfo(ctx, bucketName, fileName)
		require.NoError(t, err)
		require.Equal(t, bucketName, fileMeta.BucketName)
		require.Equal(t, fileName, fileMeta.Name)
	})

	t.Run("ListFiles", func(t *testing.T) {
		fileList, err := ipc.ListFiles(ctx, bucketName, 0, 10)
		require.NoError(t, err)
		require.Len(t, fileList, 1)
		require.Equal(t, fileName, fileList[0].Name)
	})

	t.Run("FileSetPublicAccess", func(t *testing.T) {
		require.NoError(t, ipc.FileSetPublicAccess(ctx, bucketName, fileName, true))
		fileMeta, err = ipc.FileInfo(ctx, bucketName, fileName)
		require.NoError(t, err)
		require.Equal(t, bucketName, fileMeta.BucketName)
		require.Equal(t, fileName, fileMeta.Name)
		require.True(t, fileMeta.IsPublic)
	})

	t.Run("CreateFileDownload", func(t *testing.T) {
		fileDownload, err = ipc.CreateFileDownload(ctx, bucketName, fileName)
		require.NoError(t, err)
		require.Equal(t, bucketNameEnc, fileDownload.BucketName)
		require.Equal(t, fileNameEnc, fileDownload.Name)
	})

	t.Run("Download", func(t *testing.T) {
		var downloaded bytes.Buffer
		require.NoError(t, ipc.Download(ctx, fileDownload, &downloaded))
		checkFileContents(t, len(fileData), fileData, downloaded.Bytes())
	})

	t.Run("CreateRangeFileDownload", func(t *testing.T) {
		rangeDownload, err = ipc.CreateRangeFileDownload(ctx, bucketName, fileName, 0, 1)
		require.NoError(t, err)
		require.NotEmpty(t, rangeDownload.Chunks)
		require.Equal(t, bucketNameEnc, rangeDownload.BucketName)
		require.Equal(t, fileNameEnc, rangeDownload.Name)
	})

	t.Run("DownloadRange", func(t *testing.T) {
		var rangeDownloaded bytes.Buffer
		require.NoError(t, ipc.Download(ctx, rangeDownload, &rangeDownloaded))
		checkFileContents(t, len(fileData), fileData, rangeDownloaded.Bytes())
	})

	t.Run("FileDelete", func(t *testing.T) {
		require.NoError(t, ipc.FileDelete(ctx, bucketName, fileName))
		filesAfterDelete, err := ipc.ListFiles(ctx, bucketName, 0, 10)
		require.NoError(t, err)
		require.Empty(t, filesAfterDelete)
	})

	t.Run("DeleteBucket", func(t *testing.T) {
		require.NoError(t, ipc.DeleteBucket(ctx, bucketName))
		bucketsAfterDelete, err := ipc.ListBuckets(ctx, 0, 25)
		require.NoError(t, err)
		for _, bucket := range bucketsAfterDelete {
			require.NotEqual(t, bucketName, bucket.Name)
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

func TestIPCUploadWithChunksBatchSize(t *testing.T) {
	tests := []struct {
		name      string
		fileSize  int64 // Size in MB
		batchSize int
	}{
		{"128 MB", 128, 2},
		{"256 MB", 256, 3},
	}

	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]
	address := crypto.PubkeyToAddress(pk.PublicKey)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk), sdk.WithBatchSize(tc.batchSize))
			require.NoError(t, err)
			t.Cleanup(func() {
				require.NoError(t, akave.Close())
			})

			ipc, err := akave.IPC()
			require.NoError(t, err)

			client, err := ethclient.Dial(dialURI)
			require.NoError(t, err)
			t.Cleanup(client.Close)

			nonceBefore, err := client.NonceAt(t.Context(), address, nil)
			require.NoError(t, err)

			data := testrand.Bytes(t, tc.fileSize*memory.MB.ToInt64())
			testUploadDownloadIPC(t, ipc, data, false)

			nonceAfter, err := client.NonceAt(t.Context(), address, nil)
			require.NoError(t, err)

			expectedChunks := tc.fileSize / 32
			expectedBatchTransactions := int64(math.Ceil(float64(expectedChunks) / float64(tc.batchSize)))

			expectedTotalTransactions := uint64(3 + expectedBatchTransactions)
			actualTransactions := nonceAfter - nonceBefore

			require.Equal(t, expectedTotalTransactions, actualTransactions)
		})
	}
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

	ctx := t.Context()
	fileSize := 80 * memory.MB.ToInt64()

	fileData := testrand.BytesD(t, 2024, fileSize)

	bucketName := testrand.String(10)
	fileName := testrand.String(10)

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

	bucketName := testrand.String(10)
	fileName := testrand.String(10)
	ctx := t.Context()

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

// create bucket, upload 4 files, delete files, delete bucket,
// recreate bucket, upload files again, attempt to download new files.
func TestUploadSameFilesAfterRemoval(t *testing.T) {
	fileSizesInMB := []int64{1, 16, 32, rand.Int63n(31) + 1} // include a random size between 1 and 32 MB

	privateKey := PickPrivateKey(t)
	dialURI := PickDialURI(t)
	pk := ipctest.NewFundedAccount(t, privateKey, dialURI, ipctest.ToWei(10))
	newPk := ipctest.PrivateKeyToHex(pk)

	akave, err := sdk.New(PickNodeRPCAddress(t), maxConcurrency, blockPartSize.ToInt64(), true, sdk.WithPrivateKey(newPk))
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, akave.Close())
	})

	ipc, err := akave.IPC()
	require.NoError(t, err)

	ctx := t.Context()
	bucketName := testrand.String(10)

	// Create bucket
	_, err = ipc.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	// Upload files
	fileNames := make([]string, len(fileSizesInMB))
	files := make([][]byte, len(fileSizesInMB))
	for i, size := range fileSizesInMB {
		fileNames[i] = testrand.String(10)
		files[i] = testrand.BytesD(t, int64(i), size*memory.MB.ToInt64())

		fileUpload, err := ipc.CreateFileUpload(ctx, bucketName, fileNames[i])
		require.NoError(t, err)

		upResult, err := ipc.Upload(ctx, fileUpload, bytes.NewBuffer(files[i]))
		require.NoError(t, err)
		assert.Equal(t, upResult.Name, fileNames[i])
	}

	// Verify files exist
	list, err := ipc.ListFiles(ctx, bucketName, 0, 100)
	require.NoError(t, err)
	assert.Len(t, list, len(fileSizesInMB))
	for i, f := range list {
		assert.Equal(t, fileNames[i], f.Name)
	}

	// Delete all files
	for _, fileName := range fileNames {
		require.NoError(t, ipc.FileDelete(ctx, bucketName, fileName))
	}

	// Verify files are deleted
	list, err = ipc.ListFiles(ctx, bucketName, 0, 100)
	require.NoError(t, err)
	assert.Len(t, list, 0)

	// Delete bucket
	require.NoError(t, ipc.DeleteBucket(ctx, bucketName))

	// Verify bucket is deleted
	bucketList, err := ipc.ListBuckets(ctx, 0, 100)
	require.NoError(t, err)
	for _, b := range bucketList {
		assert.NotEqual(t, bucketName, b.Name)
	}

	// Create same bucket again
	_, err = ipc.CreateBucket(ctx, bucketName)
	require.NoError(t, err)

	// Upload files again
	newFileNames := make([]string, len(fileSizesInMB))
	newFiles := make([][]byte, len(fileSizesInMB))

	// First half of files reuse original names, second half get new names
	for i, size := range fileSizesInMB {
		if i < len(fileSizesInMB)/2 {
			newFileNames[i] = fileNames[i] // Reuse original name
		} else {
			newFileNames[i] = testrand.String(10) // New name
		}
		newFiles[i] = testrand.BytesD(t, int64(100+i), size*memory.MB.ToInt64())

		fileUpload, err := ipc.CreateFileUpload(ctx, bucketName, newFileNames[i])
		require.NoError(t, err)

		upResult, err := ipc.Upload(ctx, fileUpload, bytes.NewBuffer(newFiles[i]))
		require.NoError(t, err)
		assert.Equal(t, upResult.Name, newFileNames[i])
	}

	// Verify new files exist
	list, err = ipc.ListFiles(ctx, bucketName, 0, 100)
	require.NoError(t, err)
	assert.Len(t, list, len(fileSizesInMB))
	for i, f := range list {
		assert.Equal(t, newFileNames[i], f.Name)
	}

	// Verify we can download all files
	for i, fileName := range newFileNames {
		var downloaded bytes.Buffer
		fileDownload, err := ipc.CreateFileDownload(ctx, bucketName, fileName)
		require.NoError(t, err)

		err = ipc.Download(ctx, fileDownload, &downloaded)
		require.NoError(t, err)

		// Check that downloaded content matches what we uploaded
		checkFileContents(t, 10, newFiles[i], downloaded.Bytes())
	}
}

func testUploadDownloadIPC(t *testing.T, ipc *sdk.IPC, data []byte, erasureCoding bool) {
	file := bytes.NewBuffer(data)
	fileSize := len(file.Bytes())

	bucketName := testrand.String(10)
	fileName := testrand.String(10)

	ctx, cancel := context.WithCancel(t.Context())
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

// corresponds to maybeEncrypt metadata in sdk_ipc.
func encryptMetadataHex(t *testing.T, value, derivation string) string {
	t.Helper()
	encrypted, err := encryption.EncryptD([]byte(secretKey), []byte(value), []byte(derivation))
	require.NoError(t, err)
	return hex.EncodeToString(encrypted)
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

// getStorage returns storage contract, used by node, using provided address.
func getStorage(t *testing.T, address string) *contracts.Storage {
	t.Helper()
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)

	client := pb.NewIPCNodeAPIClient(conn)

	connParams, err := client.ConnectionParams(t.Context(), &pb.ConnectionParamsRequest{})
	require.NoError(t, err)
	ethClient, err := ethclient.Dial(connParams.DialUri)
	require.NoError(t, err)

	storage, err := contracts.NewStorage(common.HexToAddress(connParams.StorageAddress), ethClient)
	require.NoError(t, err)

	return storage
}
