// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package ipc

import (
	"crypto/rand"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
