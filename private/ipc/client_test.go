// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package ipc_test

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/eip712"
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

	dataTypes := map[string][]eip712.TypedData{
		"StorageData": {
			{Name: "chunkCID", Type: "bytes"},
			{Name: "blockCID", Type: "bytes32"},
			{Name: "chunkIndex", Type: "uint256"},
			{Name: "blockIndex", Type: "uint8"},
			{Name: "nodeId", Type: "bytes"},
			{Name: "nonce", Type: "uint256"},
		},
	}

	domain := eip712.Domain{
		Name:              "Storage",
		Version:           "1",
		ChainID:           big.NewInt(31337),
		VerifyingContract: common.HexToAddress(storageAddress),
	}

	for j := range 32 {
		index := uint8(j)

		data := ipc.StorageData{
			ChunkCID:   chunkCid,
			BlockCID:   cids[j],
			ChunkIndex: big.NewInt(0),
			BlockIndex: index,
			NodeID:     nodeId,
			Nonce:      nonces[j],
		}

		dataMessage := map[string]interface{}{
			"chunkCID":   data.ChunkCID,
			"blockCID":   data.BlockCID,
			"chunkIndex": data.ChunkIndex,
			"blockIndex": data.BlockIndex,
			"nodeId":     data.NodeID,
			"nonce":      data.Nonce,
		}

		sign, err := eip712.Sign(pk, domain, dataMessage, dataTypes)
		require.NoError(t, err)

		tx, err = client.Storage.FillChunkBlock(client.Auth, cids[j], nodeId, bucket.Id, big.NewInt(0), nonces[j], index, testFileName, sign)
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

	var peerBlockCid [32]byte
	copy(peerBlockCid[:], "new test CID")

	tx, err = client.Storage.AddPeerBlock(client.Auth, []byte(testrand.GenPeerID(t, "peer1")), peerBlockCid, false)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	tx, err = client.Storage.AddPeerBlock(client.Auth, []byte(testrand.GenPeerID(t, "peer2")), peerBlockCid, true)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	listPeerIDs, err := client.Storage.GetPeersByPeerBlockCid(&bind.CallOpts{}, peerBlockCid)
	require.NoError(t, err)
	require.True(t, bytes.Equal(listPeerIDs[0], []byte(testrand.GenPeerID(t, "peer1"))))
	require.True(t, bytes.Equal(listPeerIDs[1], []byte(testrand.GenPeerID(t, "peer2"))))

	var b []byte
	b = append(b, []byte(testrand.GenPeerID(t, "peer1"))...)
	b = append(b, peerBlockCid[:]...)

	id := crypto.Keccak256Hash(b)

	idx, err := client.Storage.GetPeerBlockIndexById(&bind.CallOpts{}, []byte(testrand.GenPeerID(t, "peer1")), peerBlockCid)
	require.NoError(t, err)

	tx, err = client.Storage.DeletePeerBlock(client.Auth, id, []byte(testrand.GenPeerID(t, "peer1")), peerBlockCid, idx)
	require.NoError(t, err)
	require.NoError(t, client.WaitForTx(ctx, tx.Hash()))

	pb, err := client.Storage.GetPeerBlockById(&bind.CallOpts{}, id)
	require.NoError(t, err)
	require.Equal(t, []byte{}, pb.PeerId)
}

func generateRandomAddress(t *testing.T) common.Address {
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	require.True(t, ok)

	return crypto.PubkeyToAddress(*publicKeyECDSA)
}
