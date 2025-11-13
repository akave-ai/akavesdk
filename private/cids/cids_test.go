// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package cids

import (
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/testrand"
)

func TestVerifyRaw(t *testing.T) {
	testData := testrand.Bytes(t, 128)

	// Calculate the expected CID
	hash, err := multihash.Sum(testData, multihash.SHA2_256, -1)
	require.NoError(t, err)
	expectedCID := cid.NewCidV1(cid.DagProtobuf, hash)

	t.Run("valid CIDv0 matches", func(t *testing.T) {
		v0hash, err := multihash.Sum(testData, multihash.SHA2_256, -1)
		require.NoError(t, err)
		cidv0 := cid.NewCidV0(v0hash)
		err = VerifyRaw(cidv0.String(), testData)
		require.NoError(t, err)
	})

	t.Run("valid CIDv1 matches", func(t *testing.T) {
		err := VerifyRaw(expectedCID.String(), testData)
		require.NoError(t, err)
	})

	t.Run("CID mismatch", func(t *testing.T) {
		wrongData := []byte("different data")
		err := VerifyRaw(expectedCID.String(), wrongData)
		require.Error(t, err)
		require.Contains(t, err.Error(), "CID mismatch")
	})

	t.Run("invalid CID format", func(t *testing.T) {
		err := VerifyRaw("invalid-cid", testData)
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to decode provided CID")
	})

	t.Run("empty data", func(t *testing.T) {
		emptyData := []byte{}
		hash, err := multihash.Sum(emptyData, multihash.SHA2_256, -1)
		require.NoError(t, err)
		emptyCID := cid.NewCidV1(cid.DagProtobuf, hash)

		err = VerifyRaw(emptyCID.String(), emptyData)
		require.NoError(t, err)
	})
}

func TestVerify(t *testing.T) {
	testData := testrand.Bytes(t, 127)

	t.Run("valid CIDv1 matches", func(t *testing.T) {
		hash, err := multihash.Sum(testData, multihash.SHA2_256, -1)
		require.NoError(t, err)
		expectedCID := cid.NewCidV1(cid.DagProtobuf, hash)

		err = Verify(expectedCID, testData)
		require.NoError(t, err)
	})

	t.Run("valid CIDv0 matches", func(t *testing.T) {
		hash, err := multihash.Sum(testData, multihash.SHA2_256, -1)
		require.NoError(t, err)
		cidv0 := cid.NewCidV0(hash)

		err = Verify(cidv0, testData)
		require.NoError(t, err)
	})

	t.Run("CID mismatch", func(t *testing.T) {
		hash, err := multihash.Sum(testData, multihash.SHA2_256, -1)
		require.NoError(t, err)
		expectedCID := cid.NewCidV1(cid.DagProtobuf, hash)

		wrongData := []byte("different data")
		err = Verify(expectedCID, wrongData)
		require.Error(t, err)
		require.Contains(t, err.Error(), "CID mismatch")
	})
}
