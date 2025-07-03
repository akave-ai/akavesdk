// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package ipc_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/ipc"
	"github.com/akave-ai/akavesdk/private/ipctest"
	"github.com/akave-ai/akavesdk/private/testrand"
)

var (
	// DialURI is flag to set ipc dial uri.
	DialURI = flag.String("ipc-rpc-uri", os.Getenv("DIAL_URI"), "flag to set ipc dial uri for client tests")
	// PrivateKey is flag to set deployers hex private key.
	PrivateKey = flag.String("private-key", os.Getenv("PRIVATE_KEY"), "flag to set deployers hex private key for client tests")
)

// PickDialURI picks IPC provider URI.
func PickDialURI(t testing.TB) string {
	if *DialURI == "" || strings.EqualFold(*DialURI, "omit") {
		t.Skip("dial uri flag missing, example: -DIAL_URI=<dial uri>")
	}
	return *DialURI
}

// PickPrivateKey picks hex private key of deployer.
func PickPrivateKey(t testing.TB) string {
	if *PrivateKey == "" || strings.EqualFold(*PrivateKey, "omit") {
		t.Skip("private key flag missing, example: -PRIVATE_KEY=<deployers hex private key>")
	}
	return *PrivateKey
}

func TestContracts(t *testing.T) {
	var (
		ctx            = context.Background()
		testBucketName = "test-bucket-1"
		testFileName   = "test-file-1"
		dialUri        = PickDialURI(t)
		privateKey     = PickPrivateKey(t)
		address        = generateRandomAddress(t)
	)

	pk := ipctest.NewFundedAccount(t, privateKey, dialUri, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	client, storageAddress, _, err := ipc.DeployStorage(ctx, ipc.Config{
		DialURI:    dialUri,
		PrivateKey: newPk,
	})
	require.NoError(t, err)

	listIDs, err := client.Storage.GetOwnerBuckets(&bind.CallOpts{}, client.Auth.From)
	require.NoError(t, err)
	require.Len(t, listIDs, 0)

	tx, err := client.Storage.CreateBucket(client.Auth, testBucketName)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	bucket, err := client.Storage.GetBucketByName(&bind.CallOpts{From: client.Auth.From}, testBucketName)
	require.NoError(t, err)
	require.Equal(t, testBucketName, bucket.Name)

	listIDs, err = client.Storage.GetOwnerBuckets(&bind.CallOpts{}, client.Auth.From)
	require.NoError(t, err)
	require.Len(t, listIDs, 1)

	buckets, err := client.Storage.GetBucketsByIds(&bind.CallOpts{}, listIDs)
	require.NoError(t, err)
	require.Equal(t, testBucketName, buckets[0].Name)

	tx, err = client.Storage.CreateFile(client.Auth, buckets[0].Id, testFileName)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	chunkCid, err := hex.DecodeString("aaaa")
	require.NoError(t, err)
	c, err := hex.DecodeString("2e508ef32df4ed7026f552020fde3e522d032fa3fde0e33d06bb5485c9c82cd3")
	require.NoError(t, err)

	cids := make([][32]byte, 0)
	nonces := make([]*big.Int, 0)
	sizes := make([]*big.Int, 0)
	var (
		cid [32]byte
	)
	copy(cid[:], c)

	for i := range 32 {
		nonce := testrand.GenerateRandomNonce(t)
		cid[31] = byte(i)
		cids = append(cids, cid)
		nonces = append(nonces, nonce)
		sizes = append(sizes, big.NewInt(int64(1)))
	}

	tx, err = client.Storage.AddFileChunk(client.Auth, chunkCid, buckets[0].Id, testFileName, big.NewInt(32),
		cids, sizes, big.NewInt(0))
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	nodeId, err := hex.DecodeString("c39cd1e86738c302a2fc3eb6f6cc5d2f8a964ad29490c4335b2a2e089e0dcaf5")
	require.NoError(t, err)

	var nodeId32 [32]byte
	copy(nodeId32[:], nodeId)

	deadline := time.Now().Add(time.Hour * 24).Unix()

	for j := range 32 {
		index := uint8(j)

		data := ipc.StorageData{
			ChunkCID:   chunkCid,
			BlockCID:   cids[j],
			ChunkIndex: big.NewInt(0),
			BlockIndex: index,
			NodeID:     nodeId32,
			Nonce:      nonces[j],
			Deadline:   big.NewInt(deadline),
			BucketID:   bucket.Id,
		}

		sign, err := ipc.SignBlock(pk, storageAddress, big.NewInt(31337), data)
		require.NoError(t, err)

		tx, err = client.Storage.FillChunkBlock(client.Auth, cids[j], nodeId32, bucket.Id, big.NewInt(0), nonces[j], index, testFileName, sign, big.NewInt(deadline))
		require.NoError(t, err)
		require.NoError(t, client.WaitForTx(ctx, tx.Hash()))
	}

	tx, err = client.Storage.CommitFile(client.Auth, buckets[0].Id, testFileName, big.NewInt(32), big.NewInt(32), chunkCid)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	file, err := client.Storage.GetFileByName(&bind.CallOpts{}, buckets[0].Id, testFileName)
	require.NoError(t, err)
	require.Equal(t, testFileName, file.Name)
	require.Equal(t, chunkCid, file.FileCID)
	require.Equal(t, int64(32), file.EncodedSize.Int64())

	file, err = client.Storage.GetFileById(&bind.CallOpts{}, file.Id)
	require.NoError(t, err)
	require.Equal(t, testFileName, file.Name)
	require.Equal(t, chunkCid, file.FileCID)
	require.Equal(t, int64(32), file.EncodedSize.Int64())

	policyFactory, err := client.DeployListPolicy(ctx, client.Auth.From)
	require.NoError(t, err)
	require.NotNil(t, policyFactory)

	tx, err = policyFactory.AssignRole(client.Auth, address)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	isValid, err := policyFactory.ValidateAccess(&bind.CallOpts{}, address, nil)
	require.NoError(t, err)
	require.True(t, isValid)

	tx, err = client.AccessManager.ChangePublicAccess(client.Auth, file.Id, true)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	access, isPublic, err := client.AccessManager.GetFileAccessInfo(&bind.CallOpts{}, file.Id)
	require.NoError(t, err)
	require.NotNil(t, access)
	require.True(t, isPublic)

	fileIdx, err := client.Storage.GetFileIndexById(&bind.CallOpts{}, file.Name, file.Id)
	require.NoError(t, err)

	tx, err = client.Storage.DeleteFile(client.Auth, file.Id, file.BucketId, file.Name, fileIdx)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	bucketIdx, err := client.Storage.GetBucketIndexByName(&bind.CallOpts{}, bucket.Name, client.Auth.From)
	require.NoError(t, err)

	tx, err = client.Storage.DeleteBucket(client.Auth, buckets[0].Id, buckets[0].Name, bucketIdx)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	listIDs, err = client.Storage.GetOwnerBuckets(&bind.CallOpts{}, client.Auth.From)
	require.NoError(t, err)
	require.Len(t, listIDs, 0)

	var peerBlockCid, peerId1byte32, peerId2byte32 [32]byte
	copy(peerBlockCid[:], "new test CID")

	peerId1, err := testrand.GenPeerID(t, "peer1").MarshalBinary()
	require.NoError(t, err)
	copy(peerId1byte32[:], peerId1[6:])

	peerId2, err := testrand.GenPeerID(t, "peer2").MarshalBinary()
	require.NoError(t, err)
	copy(peerId2byte32[:], peerId2[6:])

	tx, err = client.Storage.AddPeerBlock(client.Auth, peerId1byte32, peerBlockCid, testFileName, false)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	tx, err = client.Storage.AddPeerBlock(client.Auth, peerId2byte32, peerBlockCid, testFileName, true)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	listPeerIDs, err := client.Storage.GetPeersByPeerBlockCid(&bind.CallOpts{}, peerBlockCid, testFileName)
	require.NoError(t, err)
	require.Equal(t, listPeerIDs[0], peerId1byte32)
	require.Equal(t, listPeerIDs[1], peerId2byte32)

	peersMap, err := client.Storage.GetPeersArrayByPeerBlockCid(&bind.CallOpts{}, [][32]byte{peerBlockCid}, testFileName)
	require.NoError(t, err)
	require.Len(t, peersMap, 1)

	var b []byte
	b = append(b, peerId1byte32[:]...)
	b = append(b, peerBlockCid[:]...)
	b = append(b, testFileName...)

	id := crypto.Keccak256Hash(b)

	idx, err := client.Storage.GetPeerBlockIndexById(&bind.CallOpts{}, peerId1byte32, peerBlockCid, testFileName)
	require.NoError(t, err)

	tx, err = client.Storage.DeletePeerBlock(client.Auth, id, peerId1byte32, peerBlockCid, testFileName, idx)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))
}

func TestGetTransactionReceiptsBatch(t *testing.T) {
	var (
		ctx        = context.Background()
		dialUri    = PickDialURI(t)
		privateKey = PickPrivateKey(t)
	)

	pk := ipctest.NewFundedAccount(t, privateKey, dialUri, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	client, _, _, err := ipc.DeployStorage(ctx, ipc.Config{
		DialURI:    dialUri,
		PrivateKey: newPk,
	})
	require.NoError(t, err)

	const numTxs = 55
	var requests []ipc.BatchReceiptRequest

	fromAddress := crypto.PubkeyToAddress(pk.PublicKey)
	nonce, err := client.Eth.PendingNonceAt(ctx, fromAddress)
	require.NoError(t, err)

	gasPrice, err := client.Eth.SuggestGasPrice(ctx)
	require.NoError(t, err)

	chainID, err := client.Eth.NetworkID(ctx)
	require.NoError(t, err)

	for i := 0; i < numTxs; i++ {
		toAddress := common.HexToAddress("0x000000000000000000000000000000000000dead")

		tx := types.NewTx(&types.LegacyTx{
			Nonce:    nonce + uint64(i),
			To:       &toAddress,
			Value:    big.NewInt(0),
			Gas:      21000,
			GasPrice: gasPrice,
		})

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), pk)
		require.NoError(t, err)
		require.NoError(t, client.Eth.SendTransaction(ctx, signedTx))

		hash := signedTx.Hash()
		requests = append(requests, ipc.BatchReceiptRequest{
			Hash: hash,
			Key:  fmt.Sprintf("tx-%d", i),
		})
	}

	result, err := client.GetTransactionReceiptsBatch(ctx, requests, 30*time.Second)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, len(requests), len(result.Responses))

	receiptsFound := 0
	receiptsNotFound := 0
	individualErrors := 0

	for i, response := range result.Responses {
		expectedKey := fmt.Sprintf("tx-%d", i)
		require.Equal(t, expectedKey, response.Key)

		if response.Error != nil {
			individualErrors++
			t.Logf("Transaction %s has error: %v", requests[i].Hash.Hex(), response.Error)
		} else if response.Receipt == nil {
			receiptsNotFound++
			t.Logf("Transaction %s not yet mined (receipt is nil)", requests[i].Hash.Hex())
		} else {
			receiptsFound++
			require.Equal(t, requests[i].Hash, response.Receipt.TxHash)
			require.True(t, response.Receipt.Status == 0 || response.Receipt.Status == 1)
			require.True(t, response.Receipt.BlockNumber.Uint64() > 0)

			t.Logf("Transaction %s mined in block %d with status %d",
				response.Receipt.TxHash.Hex(),
				response.Receipt.BlockNumber.Uint64(),
				response.Receipt.Status)
		}
	}

	require.Equal(t, len(requests), receiptsFound+receiptsNotFound+individualErrors)
}

func generateRandomAddress(t *testing.T) common.Address {
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	require.True(t, ok)

	return crypto.PubkeyToAddress(*publicKeyECDSA)
}
