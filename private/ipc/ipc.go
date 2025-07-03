// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package ipc

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/akave-ai/akavesdk/private/eip712"
)

// GenerateNonce generates a random 256 bit nonce.
func GenerateNonce() (*big.Int, error) {
	b := make([]byte, 32)

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return big.NewInt(0).SetBytes(b), nil
}

// CalculateFileID calculates on-chain file id the way it is done on the contract.
func CalculateFileID(bucketID []byte, name string) common.Hash {
	var b []byte
	b = append(b, bucketID...)
	b = append(b, name...)

	return crypto.Keccak256Hash(b)
}

// CalculateBucketID calculates on-chain bucket id the way it is done on the contract.
func CalculateBucketID(bucketName, address string) []byte {
	b := make([]byte, 0)
	b = append(b, bucketName...)
	addr := common.HexToAddress(address)
	b = append(b, addr.Bytes()...)
	return crypto.Keccak256(b)
}

// SignBlock signs StorageData using EIP712 standard.
func SignBlock(privateKey *ecdsa.PrivateKey, storageAddress string, chainId *big.Int, data StorageData) ([]byte, error) {
	domain := eip712.Domain{
		Name:              "Storage",
		Version:           "1",
		ChainID:           chainId,
		VerifyingContract: common.HexToAddress(storageAddress),
	}
	storageDataTypes := map[string][]eip712.TypedData{
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

	dataMessage := map[string]any{
		"chunkCID":   data.ChunkCID,
		"blockCID":   data.BlockCID,
		"chunkIndex": data.ChunkIndex,
		"blockIndex": data.BlockIndex,
		"nodeId":     data.NodeID,
		"nonce":      data.Nonce,
		"deadline":   data.Deadline,
		"bucketId":   data.BucketID,
	}

	return eip712.Sign(privateKey, domain, dataMessage, storageDataTypes)
}
