// Copyright (C) 2024 Akave
// See LICENSE for copying information.

// Package testrand provides utilities for generating random data in tests.
package testrand

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
	rand2 "math/rand"
	"testing"

	"github.com/ipfs/go-cid"
	libp2pCrypto "github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/require"
)

// Bytes returns a slice of random bytes of the given size.
func Bytes(t testing.TB, size int64) []byte {
	t.Helper()
	data := make([]byte, size)
	_, err := rand.Read(data)
	if err != nil {
		t.Fatalf("failed to read random data: %v", err)
	}
	return data
}

// BytesD returns a slice of random bytes of the given size with a deterministic seed.
func BytesD(t testing.TB, seed, size int64) []byte {
	t.Helper()
	data := make([]byte, size)
	random := rand2.New(rand2.NewSource(seed))
	_, err := random.Read(data)
	if err != nil {
		t.Fatalf("failed to read random data: %v", err)
	}
	return data
}

// String returns a random string of the given length using math/rand.
func String(length int) string {
	dictionary := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	dictionaryLen := len(dictionary)

	buffer := make([]byte, length)
	for i := range buffer {
		buffer[i] = dictionary[rand2.Intn(dictionaryLen)]
	}
	return string(buffer)
}

// GenerateRandomNonce generates a random bit.Int nonce.
func GenerateRandomNonce(t testing.TB) *big.Int {
	b := make([]byte, 32)

	_, err := rand.Read(b)
	if err != nil {
		t.Fatalf("failed to read random data: %v", err)
	}

	return big.NewInt(0).SetBytes(b)
}

// GenPeerID generates a peer.ID from a seed string deterministically.
func GenPeerID(t testing.TB, seed string) peer.ID {
	t.Helper()
	hash := sha256.Sum256([]byte(seed))
	privateKey, _, err := libp2pCrypto.GenerateEd25519Key(bytes.NewReader(hash[:]))
	require.NoError(t, err)
	peerId, err := peer.IDFromPrivateKey(privateKey)
	require.NoError(t, err)
	return peerId
}

// GenerateRandomCID returns cid, generated from random byte data.
func GenerateRandomCID(t testing.TB) cid.Cid {
	data := make([]byte, 32)
	_, err := rand.Read(data)
	require.NoError(t, err)

	mh, err := multihash.Sum(data, multihash.SHA2_256, -1)
	require.NoError(t, err)
	return cid.NewCidV1(cid.DagProtobuf, mh)
}
