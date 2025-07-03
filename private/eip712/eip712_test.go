// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package eip712_test

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/eip712"
	"github.com/akave-ai/akavesdk/private/ipc"
)

func TestSignature(t *testing.T) {
	privateKeyHex := "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	require.NoError(t, err)
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	require.NoError(t, err)

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	var blockCid, nodeID32, bucketId32 [32]byte
	copy(blockCid[:], "blockCID1")
	nodeID := []byte("node id")
	copy(nodeID32[:], nodeID)
	bucketId := []byte("bucket id")
	copy(bucketId32[:], bucketId)

	dataTypes := map[string][]eip712.TypedData{
		"StorageData": {
			{Name: "chunkCID", Type: "bytes"},
			{Name: "blockCID", Type: "bytes32"},
			{Name: "chunkIndex", Type: "uint256"},
			{Name: "blockIndex", Type: "uint8"},
			{Name: "nodeId", Type: "bytes32"},
			{Name: "nonce", Type: "uint256"},
			{Name: "deadline", Type: "uint256"},
			{Name: "bucketId", Type: "bytes32"},
		},
	}

	data := ipc.StorageData{
		ChunkCID:   []byte("rootCID1"),
		BlockCID:   blockCid,
		ChunkIndex: big.NewInt(0),
		BlockIndex: 0,
		NodeID:     nodeID32,
		Nonce:      big.NewInt(1234567890),
		Deadline:   big.NewInt(12345),
		BucketID:   bucketId32,
	}

	domain := eip712.Domain{
		Name:              "Storage",
		Version:           "1",
		ChainID:           big.NewInt(31337),
		VerifyingContract: common.HexToAddress("0x1234567890123456789012345678901234567890"),
	}

	dataMessage := map[string]interface{}{
		"chunkCID":   data.ChunkCID,
		"blockCID":   data.BlockCID,
		"chunkIndex": data.ChunkIndex,
		"blockIndex": data.BlockIndex,
		"nodeId":     data.NodeID,
		"nonce":      data.Nonce,
		"deadline":   data.Deadline,
		"bucketId":   data.BucketID,
	}

	sign, err := eip712.Sign(privateKey, domain, dataMessage, dataTypes)
	require.NoError(t, err)

	recoveredAddr, err := eip712.RecoverSignerAddress(sign, domain, dataMessage, dataTypes)
	if err != nil {
		fmt.Printf("Error recovering address: %v\n", err)
		return
	}
	require.Equal(t, address.Hex(), recoveredAddr.Hex())
}
