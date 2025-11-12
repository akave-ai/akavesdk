// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package eip712_test

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/eip712"
	"github.com/akave-ai/akavesdk/private/ipc"
)

func TestSignatureAgainstContract(t *testing.T) {
	testCases := []struct {
		chunkCID       string
		blockCID       string
		nodeID         string
		nonce          int64
		deadline       int64
		bucketID       string
		storageAddress string
		signature      string
	}{{
		chunkCID:       "86b258127d599eb74c729f97",
		blockCID:       "c00612ae8af29b5437ba40df50c46c0175c69b6dc3b3014ed19bda51e318f0f3",
		nodeID:         "5a604f924e185f6ec5754156e331e9d52df8a669de7e1a060b90e636e0e9e818",
		nonce:          3456789012,
		deadline:       1759859212,
		bucketID:       "930c2de1e6a9a0726f2d7bde19428453d9fdc11fa5c98205ce9b9e794bbd93a2",
		storageAddress: "0x4e7B1E9c3214C973Ff2fc680A9789E8579a5eD9d",
		signature:      "726683359604ffe042e73afd7adef9b7f6e13ffd0078999d31bd1cc8c119e1e8324d44cffdc2f771912e500c522082ee94e5f30ac5844c06497e3c49dab8b6de1b",
	}, {
		chunkCID:       "edf5fb5fdd325e462cd806f2",
		blockCID:       "fbeeb197dd90574c97d5993fab0610403197db0f18133033755ec39cab7596c9",
		nodeID:         "3a59ed631290287c86c90777b2d45926c1a860b1e90828963358d72fa8834389",
		nonce:          2345678901,
		deadline:       1759862780,
		bucketID:       "95f7f023dbf92b2ab036280c44037485c0deec1d854046443bae8ae16c37bc86",
		storageAddress: "0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f",
		signature:      "47569b36d69bde9e8953cc8c6a01599f0a307850d25e9101c4b1338fbf562d58017bd4ecae535eb330ea7c7ca710fb0055d9d3697e2ebc18902aa32d252eb7361c",
	}, {
		chunkCID:       "2e3adffef0437b35f247022b",
		blockCID:       "fc785a432d1c6d45671f60ed36f44378f63ae4fbbf4ef2a9f0d4951e77e81272",
		nodeID:         "050f9e0347ebfbdcf50fddf89713b7f37e667d19279d9f550fa7b93237ce29fa",
		nonce:          1234567890,
		deadline:       1759866325,
		bucketID:       "a928e74732b6ca5fd1bf7f3eedfdca3c578a05297157e239e7f7861de2b40f42",
		storageAddress: "0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc",
		signature:      "8ccd5143f4b87e898021c4b3a4bf73e3e8d6e8b97e39106374fac72be610629463a0ba6fc4c975c41fbb1ad3940f76a30e6cb916a8e01d09afbe24538ce151ca1b",
	}}

	dataTypes := apitypes.Types{
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

	for _, tc := range testCases {
		domain := apitypes.TypedDataDomain{
			Name:              "Storage",
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(31337),
			VerifyingContract: tc.storageAddress,
		}

		var (
			blockCid32, nodeID32, bucketID32     [32]byte
			chunkCid, blockCid, nodeID, bucketID []byte
		)

		chunkCid, err := hex.DecodeString(tc.chunkCID)
		require.NoError(t, err)

		blockCid, err = hex.DecodeString(tc.blockCID)
		require.NoError(t, err)
		copy(blockCid32[:], blockCid)

		nodeID, err = hex.DecodeString(tc.nodeID)
		require.NoError(t, err)
		copy(nodeID32[:], nodeID)

		bucketID, err = hex.DecodeString(tc.bucketID)
		require.NoError(t, err)
		copy(bucketID32[:], bucketID)

		data := ipc.StorageData{
			ChunkCID:   chunkCid,
			BlockCID:   blockCid32,
			ChunkIndex: big.NewInt(0),
			BlockIndex: big.NewInt(0),
			NodeID:     nodeID32,
			Nonce:      big.NewInt(tc.nonce),
			Deadline:   big.NewInt(tc.deadline),
			BucketID:   bucketID32,
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

		pk, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
		require.NoError(t, err)

		sign, err := eip712.Sign(pk, domain, "StorageData", dataTypes, dataMessage)
		require.NoError(t, err)
		require.Equal(t, tc.signature, hex.EncodeToString(sign))
	}
}

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

	dataTypes := apitypes.Types{
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
		BlockIndex: big.NewInt(0),
		NodeID:     nodeID32,
		Nonce:      big.NewInt(1234567890),
		Deadline:   big.NewInt(12345),
		BucketID:   bucketId32,
	}

	domain := apitypes.TypedDataDomain{
		Name:              "Storage",
		Version:           "1",
		ChainId:           math.NewHexOrDecimal256(31337),
		VerifyingContract: "0x1234567890123456789012345678901234567890",
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

	sign, err := eip712.Sign(privateKey, domain, "StorageData", dataTypes, dataMessage)
	require.NoError(t, err)

	recoveredAddr, err := eip712.RecoverSignerAddress(sign, domain, "StorageData", dataTypes, dataMessage)
	if err != nil {
		fmt.Printf("Error recovering address: %v\n", err)
		return
	}
	require.Equal(t, address.Hex(), recoveredAddr.Hex())
}
