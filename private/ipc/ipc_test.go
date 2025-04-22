// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package ipc_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/ipc"
)

func TestGenerateNonce(t *testing.T) {
	var nonce *big.Int
	var err error

	for i := range 10 {
		nonce, err = ipc.GenerateNonce()
		require.NoError(t, err)

		if i > 0 {
			t.Log("retrying to sample nonce", i)
		}
		if len(nonce.Bytes()) == 32 {
			break
		}
	}
	require.Len(t, nonce.Bytes(), 32)
}

func TestCalculateFileID(t *testing.T) {
	for _, testCase := range []struct {
		BucketID []byte
		Name     string
		Expected common.Hash
	}{
		{
			BucketID: common.Hex2Bytes("c10fad62c0224052065576135ed2ae4d85d34432b4fb40796eadd9a991f064b9"),
			Name:     "file1",
			Expected: common.HexToHash("eea1eddf9f4be315e978c6d0d25d1b870ec0162ebf0acf173f47b738ff0cb421"),
		},
		{
			BucketID: common.Hex2Bytes("f855c5499b442e6b57ea3ec0c260d1e23a74415451ec5a4796560cc9b7d89be0"),
			Name:     "file2",
			Expected: common.HexToHash("f8d6d1f6e7ba4f69f00df4e4b53b3e51eb8610f0774f16ea1f02162e0034487b"),
		},
		{
			BucketID: common.Hex2Bytes("f06eac67910341b595f1ef319ca12713a79f180b96a685430d806701dc42e9aa"),
			Name:     "file3",
			Expected: common.HexToHash("3eb92385cd986662e9885c47364fa5b2f154cd6fca8d99f4aed68160064991cb"),
		},
	} {
		fileID := ipc.CalculateFileID(testCase.BucketID, testCase.Name)
		require.Equal(t, testCase.Expected, fileID)
	}
}
