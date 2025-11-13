// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package ipc_test

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/ipc"
	"github.com/akave-ai/akavesdk/private/ipc/contracts"
	"github.com/akave-ai/akavesdk/private/ipctest"
	"github.com/akave-ai/akavesdk/private/testrand"
)

func TestContracts(t *testing.T) {
	var (
		ctx            = t.Context()
		testBucketName = "test-bucket-1"
		testFileName   = "test-file-1"
		dialUri        = ipctest.PickDialURI(t)
		privateKey     = ipctest.PickPrivateKey(t)
		address        = generateRandomAddress(t)
	)

	pk := ipctest.NewFundedAccount(t, privateKey, dialUri, ipctest.ToWei(10))
	newPk := ipctest.PrivateKeyToHex(pk)

	client, err := ipc.DeployContracts(ctx, ipc.Config{
		DialURI:    dialUri,
		PrivateKey: newPk,
	})
	require.NoError(t, err)

	listIDs, err := client.Storage.GetOwnerBuckets(&bind.CallOpts{Context: ctx}, client.Auth.From, big.NewInt(0), big.NewInt(10), big.NewInt(0), big.NewInt(0))
	require.NoError(t, err)
	require.Len(t, listIDs, 0)

	tx, err := client.Storage.CreateBucket(client.Auth, testBucketName)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	bucket, err := client.Storage.GetBucketByName(&bind.CallOpts{Context: ctx, From: client.Auth.From}, testBucketName, client.Auth.From, big.NewInt(0), big.NewInt(0))
	require.NoError(t, err)
	require.Equal(t, testBucketName, bucket.Name)

	listIDs, err = client.Storage.GetOwnerBuckets(&bind.CallOpts{Context: ctx}, client.Auth.From, big.NewInt(0), big.NewInt(10), big.NewInt(0), big.NewInt(0))
	require.NoError(t, err)
	require.Len(t, listIDs, 1)

	bucketIDs := make([][32]byte, len(listIDs))
	for i, bucketStorage := range listIDs {
		bucketIDs[i] = bucketStorage.Id
	}

	buckets, err := client.Storage.GetBucketsByIds(&bind.CallOpts{Context: ctx}, bucketIDs)
	require.NoError(t, err)
	require.Equal(t, testBucketName, buckets[0].Name)

	tx, err = client.Storage.CreateFile(client.Auth, buckets[0].Id, testFileName)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	chunkCid := testrand.CID(t)

	cids := make([][32]byte, 0)
	nonces := make([]*big.Int, 0)
	sizes := make([]*big.Int, 0)

	for range 32 {
		blockCid := testrand.CID(t)
		var blockCidBytes [32]byte
		copy(blockCidBytes[:], blockCid.Bytes()[4:])

		nonce := testrand.Nonce(t)
		cids = append(cids, blockCidBytes)
		nonces = append(nonces, nonce)
		sizes = append(sizes, big.NewInt(int64(1)))
	}

	tx, err = client.Storage.AddFileChunk(client.Auth, chunkCid.Bytes(), buckets[0].Id, testFileName, big.NewInt(32), cids, sizes, big.NewInt(0))
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	nodeId, err := testrand.PeerID(t, "node1").MarshalBinary()
	require.NoError(t, err)

	var nodeId32 [32]byte
	copy(nodeId32[:], nodeId[6:])

	deadline := time.Now().Add(time.Hour * 24).Unix()

	for j := range 32 {
		index := uint8(j)

		data := ipc.StorageData{
			ChunkCID:   chunkCid.Bytes(),
			BlockCID:   cids[j],
			ChunkIndex: big.NewInt(0),
			BlockIndex: big.NewInt(int64(index)),
			NodeID:     nodeId32,
			Nonce:      nonces[j],
			Deadline:   big.NewInt(deadline),
			BucketID:   bucket.Id,
		}

		sign, err := ipc.SignBlock(pk, client.Addresses.Storage, big.NewInt(31337), data)
		require.NoError(t, err)

		tx, err = client.Storage.FillChunkBlock(client.Auth, contracts.IStorageFillChunkBlockArgs{
			BlockCID:   cids[j],
			NodeId:     nodeId32,
			BucketId:   bucket.Id,
			ChunkIndex: big.NewInt(0),
			Nonce:      nonces[j],
			BlockIndex: index,
			FileName:   testFileName,
			Signature:  sign,
			Deadline:   big.NewInt(deadline),
		})
		require.NoError(t, err)
		require.NoError(t, client.WaitForTx(ctx, tx.Hash()))
	}

	rootCID := testrand.CID(t)
	tx, err = client.Storage.CommitFile(client.Auth, buckets[0].Id, testFileName, big.NewInt(32), big.NewInt(32), rootCID.Bytes())
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	file, err := client.Storage.GetFileByName(&bind.CallOpts{Context: ctx}, buckets[0].Id, testFileName)
	require.NoError(t, err)
	require.Equal(t, testFileName, file.Name)
	require.Equal(t, rootCID.Bytes(), file.FileCID)
	require.Equal(t, int64(32), file.EncodedSize.Int64())

	file, err = client.Storage.GetFileById(&bind.CallOpts{Context: ctx}, file.Id)
	require.NoError(t, err)
	require.Equal(t, testFileName, file.Name)
	require.Equal(t, rootCID.Bytes(), file.FileCID)
	require.Equal(t, int64(32), file.EncodedSize.Int64())

	policyFactory, err := client.DeployListPolicy(ctx, client.Auth.From)
	require.NoError(t, err)
	require.NotNil(t, policyFactory)

	tx, err = policyFactory.AssignRole(client.Auth, address)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	isValid, err := policyFactory.ValidateAccess(&bind.CallOpts{Context: ctx}, address, nil)
	require.NoError(t, err)
	require.True(t, isValid)

	tx, err = client.AccessManager.ChangePublicAccess(client.Auth, file.Id, true)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	access, isPublic, err := client.AccessManager.GetFileAccessInfo(&bind.CallOpts{Context: ctx}, file.Id)
	require.NoError(t, err)
	require.NotNil(t, access)
	require.True(t, isPublic)

	fileIdx, err := client.Storage.GetFileIndexById(&bind.CallOpts{Context: ctx, From: client.Auth.From}, bucket.Name, file.Id, client.Auth.From)
	require.NoError(t, err)
	require.True(t, fileIdx.Exists)

	tx, err = client.Storage.DeleteFile(client.Auth, file.Id, file.BucketId, file.Name, fileIdx.Index)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	bucketIdx, err := client.Storage.GetBucketIndexByName(&bind.CallOpts{Context: ctx, From: client.Auth.From}, bucket.Name, client.Auth.From)
	require.NoError(t, err)
	require.True(t, bucketIdx.Exists)

	tx, err = client.Storage.DeleteBucket(client.Auth, buckets[0].Id, buckets[0].Name, bucketIdx.Index)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	listIDs, err = client.Storage.GetOwnerBuckets(&bind.CallOpts{Context: ctx}, client.Auth.From, big.NewInt(0), big.NewInt(10), big.NewInt(0), big.NewInt(0))
	require.NoError(t, err)
	require.Len(t, listIDs, 0)

	// Test GetPeers method instead
	peers, err := client.Storage.GetBlockPeersOfChunk(&bind.CallOpts{Context: ctx}, cids, file.Id, big.NewInt(0))
	require.NoError(t, err)
	require.Len(t, peers, 32)
	require.Contains(t, peers, nodeId32)
}

func TestAddFileChunksFillChunkBlocks(t *testing.T) {
	var (
		ctx            = t.Context()
		testBucketName = "test-bucket-2"
		testFileName   = "test-file-2"
		dialUri        = ipctest.PickDialURI(t)
		privateKey     = ipctest.PickPrivateKey(t)
	)

	pk := ipctest.NewFundedAccount(t, privateKey, dialUri, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	client, err := ipc.DeployContracts(ctx, ipc.Config{
		DialURI:    dialUri,
		PrivateKey: newPk,
	})
	require.NoError(t, err)

	tx, err := client.Storage.CreateBucket(client.Auth, testBucketName)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	bucket, err := client.Storage.GetBucketByName(&bind.CallOpts{Context: ctx, From: client.Auth.From}, testBucketName, client.Auth.From, big.NewInt(0), big.NewInt(0))
	require.NoError(t, err)
	require.Equal(t, testBucketName, bucket.Name)

	tx, err = client.Storage.CreateFile(client.Auth, bucket.Id, testFileName)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	chunkCid := testrand.CID(t)
	chunkCid2 := testrand.CID(t)
	c, err := hex.DecodeString("2e508ef32df4ed7026f552020fde3e522d032fa3fde0e33d06bb5485c9c82cd3")
	require.NoError(t, err)

	var (
		cids              = make([][32]byte, 0)
		firstChunkNonces  = make([]*big.Int, 0)
		secondChunkNonces = make([]*big.Int, 0)
		sizes             = make([]*big.Int, 0)
		cidsTotal         = make([][][32]byte, 0)
		sizesTotal        = make([][]*big.Int, 0)
		cid               [32]byte
	)
	copy(cid[:], c)

	for i := range 32 {
		cid[31] = byte(i)
		cids = append(cids, cid)
		sizes = append(sizes, big.NewInt(int64(1)))
		firstChunkNonces = append(firstChunkNonces, testrand.Nonce(t))
		secondChunkNonces = append(secondChunkNonces, testrand.Nonce(t))
	}

	cidsTotal = append(cidsTotal, cids, cids)     // same cids in both chunk is acceptable.
	sizesTotal = append(sizesTotal, sizes, sizes) // same sizes of blocks in both chunk is acceptable.

	tx, err = client.Storage.AddFileChunks(client.Auth, [][]byte{chunkCid.Bytes(), chunkCid2.Bytes()}, bucket.Id, testFileName, []*big.Int{big.NewInt(32), big.NewInt(32)},
		cidsTotal, sizesTotal, big.NewInt(0))
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	nodeId, err := hex.DecodeString("c39cd1e86738c302a2fc3eb6f6cc5d2f8a964ad29490c4335b2a2e089e0dcaf5")
	require.NoError(t, err)

	var nodeId32 [32]byte
	copy(nodeId32[:], nodeId)

	deadline := time.Now().Add(time.Hour * 24).Unix()
	chunkCIDs := [][]byte{chunkCid.Bytes(), chunkCid2.Bytes()}
	chunkNonces := [][]*big.Int{firstChunkNonces, secondChunkNonces}

	for chunkIndex := range len(chunkCIDs) {
		args := make([]contracts.IStorageFillChunkBlockArgs, 0)

		for j := range 32 {
			index := uint8(j)

			data := ipc.StorageData{
				ChunkCID:   chunkCIDs[chunkIndex],
				BlockCID:   cids[j],
				ChunkIndex: big.NewInt(int64(chunkIndex)),
				BlockIndex: big.NewInt(int64(index)),
				NodeID:     nodeId32,
				Nonce:      chunkNonces[chunkIndex][j],
				Deadline:   big.NewInt(deadline),
				BucketID:   bucket.Id,
			}

			sign, err := ipc.SignBlock(pk, client.Addresses.Storage, big.NewInt(31337), data)
			require.NoError(t, err)

			args = append(args, contracts.IStorageFillChunkBlockArgs{
				BlockCID:   cids[j],
				NodeId:     nodeId32,
				BucketId:   bucket.Id,
				ChunkIndex: big.NewInt(int64(chunkIndex)),
				Nonce:      chunkNonces[chunkIndex][j],
				BlockIndex: index,
				FileName:   testFileName,
				Signature:  sign,
				Deadline:   big.NewInt(deadline),
			})
		}

		tx, err = client.Storage.FillChunkBlocks(client.Auth, args)
		require.NoError(t, err)
		require.NoError(t, client.WaitForTx(ctx, tx.Hash()))
	}

	tx, err = client.Storage.CommitFile(client.Auth, bucket.Id, testFileName, big.NewInt(64), big.NewInt(64), chunkCid.Bytes())
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))
}

func generateRandomAddress(t *testing.T) common.Address {
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	require.True(t, ok)

	return crypto.PubkeyToAddress(*publicKeyECDSA)
}
