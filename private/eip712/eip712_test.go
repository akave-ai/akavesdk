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
	var blockCid [32]byte
	copy(blockCid[:], "blockCID1")

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

	data := ipc.StorageData{
		ChunkCID:   []byte("rootCID1"),
		BlockCID:   blockCid,
		ChunkIndex: big.NewInt(0),
		BlockIndex: 0,
		NodeID:     []byte("node id"),
		Nonce:      big.NewInt(1234567890),
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
