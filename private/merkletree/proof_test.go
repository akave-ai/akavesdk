// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package merkletree_test

import (
	"bytes"
	"testing"

	commcid "github.com/filecoin-project/go-fil-commcid"
	commp "github.com/filecoin-project/go-fil-commp-hashhash"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/merkletree"
)

func TestMemtreeProof(t *testing.T) {
	cp := new(commp.Calc)

	data := []byte{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
		1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
		2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
		3, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
		4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
		5, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
		6, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
		7, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
	}

	_, err := cp.Write(data)
	require.NoError(t, err)

	rawCommP, paddedSize, err := cp.Digest()
	require.NoError(t, err)

	commCid, err := commcid.DataCommitmentV1ToCID(rawCommP)
	require.NoError(t, err)

	tree, err := merkletree.BuildSha254Memtree(
		merkletree.NewPaddedReader(bytes.NewReader(data), int64(paddedSize)),
		abi.PaddedPieceSize(paddedSize).Unpadded())
	require.NoError(t, err)
	require.NotNil(t, tree)

	proof, err := merkletree.MemtreeProof(tree, 3)
	require.NoError(t, err)

	commCid2, err := commcid.DataCommitmentV1ToCID(proof.Root[:])
	require.NoError(t, err)
	require.True(t, commCid.Equals(commCid2))

	// Verify the proof with known root CID(rawCommP)
	proof.Root = [32]byte(rawCommP)
	require.True(t, merkletree.VerifyMerkleProof(proof, 3))
}
