// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package erasurecode_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/klauspost/reedsolomon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/encryption"
	"github.com/akave-ai/akavesdk/private/erasurecode"
	"github.com/akave-ai/akavesdk/private/testrand"
)

func TestErasureCodeInvalidParams(t *testing.T) {
	_, err := erasurecode.New(0, 0)
	require.Error(t, err)

	_, err = erasurecode.New(16, 0)
	require.Error(t, err)
}

func TestErasureCode(t *testing.T) {
	data := []byte("Quick brown fox jumps over the lazy dog")
	dataShards := 5
	parityShards := 3

	encoder, err := erasurecode.New(dataShards, parityShards)
	require.NoError(t, err)
	assert.Equal(t, dataShards, encoder.DataBlocks)
	assert.Equal(t, parityShards, encoder.ParityBlocks)

	encoded, err := encoder.Encode(data)
	require.NoError(t, err)

	shardSize := len(encoded) / (dataShards + parityShards)

	t.Run("no missing shards", func(t *testing.T) {
		blocks := splitIntoBlocks(encoded, shardSize)

		extracted, err := encoder.ExtractData(blocks)
		require.NoError(t, err)
		assert.Equal(t, data, extracted)
	})

	t.Run(fmt.Sprintf("missing no more than %d shards shards", parityShards), func(t *testing.T) {
		// Generate all possible missing shard combinations of length <= parityShards
		var allCombos [][]int
		for k := 1; k <= 3; k++ {
			combos := missingShardsIdx(dataShards+parityShards, k)
			allCombos = append(allCombos, combos...)
		}

		encoded, err := encoder.Encode(data)
		require.NoError(t, err)

		// Test each combination of missing shards
		for _, missingIdxs := range allCombos {
			blocks := splitIntoBlocks(encoded, shardSize)

			// Set missing blocks to nil
			for _, idx := range missingIdxs {
				blocks[idx] = nil
			}

			extracted, err := encoder.ExtractData(blocks)
			require.NoError(t, err)
			assert.Equal(t, data, extracted)
		}
	})

	t.Run(fmt.Sprintf("missing more than %d shards", parityShards), func(t *testing.T) {
		blocks := splitIntoBlocks(encoded, shardSize)
		for i := range parityShards + 1 {
			blocks[i] = nil
		}
		_, err := encoder.ExtractData(blocks)
		require.Error(t, err)
	})

	t.Run("non-erasure-coded data", func(t *testing.T) {
		// Create blocks of the correct size, but with random (non-encoded) data
		blocks := make([][]byte, dataShards+parityShards)
		for i := range blocks {
			block := make([]byte, shardSize)
			for j := range block {
				block[j] = byte(i + j)
			}
			blocks[i] = block
		}
		_, err := encoder.ExtractData(blocks)
		require.Error(t, err)
	})

	t.Run("non nil shard should fail on reconstruct", func(t *testing.T) {
		blocks := splitIntoBlocks(encoded, shardSize)
		// Corrupt one block's size
		blocks[0] = blocks[0][:len(blocks[0])-1]

		_, err := encoder.ExtractData(blocks)
		require.Error(t, err)
	})

	t.Run("too few shards should fail without reconstruction", func(t *testing.T) {
		blocks := splitIntoBlocks(encoded, shardSize)
		// Remove all but one shard (triggers ErrTooFewShards, not ErrShardSize)
		blocks = blocks[:1]

		_, err := encoder.ExtractData(blocks)
		require.Error(t, err)
		require.False(t, errors.Is(err, reedsolomon.ErrShardSize))
	})

	t.Run("all nil shards should fail without reconstruction", func(t *testing.T) {
		// Create blocks with all nil shards (triggers ErrShardNoData, not ErrShardSize)
		blocks := make([][]byte, dataShards+parityShards)

		_, err := encoder.ExtractData(blocks)
		require.Error(t, err)
		require.False(t, errors.Is(err, reedsolomon.ErrShardSize))
	})

	t.Run("too many nil shards should fail without reconstruction", func(t *testing.T) {
		blocks := splitIntoBlocks(encoded, shardSize)
		// Set more than parity shards to nil (triggers ErrShardSize)
		for i := 0; i < parityShards+1; i++ {
			blocks[i] = nil
		}

		// This will fail because even though ErrShardSize is triggered,
		// Reconstruct will fail when trying to reconstruct >parityShards missing shards
		_, err := encoder.ExtractData(blocks)
		require.Error(t, err)
	})
}

func TestReconstructData(t *testing.T) {
	data := []byte("Quick brown fox jumps over the lazy dog")
	dataShards := 5
	parityShards := 3

	encoder, err := erasurecode.New(dataShards, parityShards)
	require.NoError(t, err)

	originalEncoded, err := encoder.Encode(data)
	require.NoError(t, err)

	shardSize := len(originalEncoded) / (dataShards + parityShards)

	t.Run("reconstruct all blocks", func(t *testing.T) {
		originalBlocks := splitIntoBlocks(originalEncoded, shardSize)

		blocksToReconstruct := make([][]byte, len(originalBlocks))
		for i, block := range originalBlocks {
			clone := make([]byte, len(block))
			copy(clone, block)
			blocksToReconstruct[i] = clone
		}

		blocksToReconstruct[0] = nil
		blocksToReconstruct[1] = nil
		blocksToReconstruct[len(blocksToReconstruct)-1] = nil

		err = encoder.ReconstructAll(blocksToReconstruct)
		require.NoError(t, err)

		assert.Equal(t, originalBlocks, blocksToReconstruct)
	})
}

func TestEncodeWithPreallocatedSlice(t *testing.T) {
	data := []byte("Quick brown fox jumps over the lazy dog")
	dataShards := 5
	parityShards := 3

	encoder, err := erasurecode.New(dataShards, parityShards)
	require.NoError(t, err)

	// Preallocate byte slice with data, taking WrapOverhead into account
	preallocated := make([]byte, len(data), len(data)+erasurecode.WrapOverhead)
	copy(preallocated, data)

	encoded, err := encoder.Encode(preallocated)
	require.NoError(t, err)

	shardSize := len(encoded) / (dataShards + parityShards)
	blocks := splitIntoBlocks(encoded, shardSize)

	// Verify that we can extract the original data
	extracted, err := encoder.ExtractData(blocks)
	require.NoError(t, err)
	assert.Equal(t, data, extracted)
}

func missingShardsIdx(n, k int) [][]int {
	if k <= 0 || k > n {
		return nil
	}

	result := make([][]int, 0)
	current := make([]int, 0, k)

	var generateCombinations func(start, remaining int)
	generateCombinations = func(start, remaining int) {
		if remaining == 0 {
			combination := make([]int, k)
			copy(combination, current)
			result = append(result, combination)
			return
		}

		for i := start; i <= n-remaining; i++ {
			if len(current) == 0 || i > current[len(current)-1] {
				current = append(current, i)
				generateCombinations(i+1, remaining-1)
				current = current[:len(current)-1]
			}
		}
	}

	generateCombinations(0, k)
	return result
}

func splitIntoBlocks(encoded []byte, shardSize int) [][]byte {
	var blocks [][]byte
	for offset := 0; offset < len(encoded); offset += shardSize {
		end := min(offset+shardSize, len(encoded))
		block := make([]byte, shardSize)
		copy(block, encoded[offset:end])
		blocks = append(blocks, block)
	}
	return blocks
}

func TestEncryptAndErasureCode16x16_16MB(t *testing.T) {
	const (
		dataSize     = 16 * 1024 * 1024
		dataShards   = 16
		parityShards = 16
	)

	encryptionKey := testrand.Bytes(t, 32)
	plainData := testrand.Bytes(t, int64(dataSize))

	encoder, err := erasurecode.New(dataShards, parityShards)
	require.NoError(t, err)

	requiredCapacity := dataSize + encryption.Overhead + erasurecode.WrapOverhead
	encryptedData := make([]byte, dataSize, requiredCapacity)
	copy(encryptedData, plainData)
	encryptedData, err = encryption.Encrypt(encryptionKey, encryptedData, "erasure-coding")
	require.NoError(t, err)

	encoded, err := encoder.Encode(encryptedData)
	require.NoError(t, err)

	shardSize := len(encoded) / (dataShards + parityShards)

	t.Run("no missing shards", func(t *testing.T) {
		blocks := splitIntoBlocks(encoded, shardSize)

		extracted, err := encoder.ExtractData(blocks)
		require.NoError(t, err)

		decrypted, err := encryption.Decrypt(encryptionKey, extracted, "erasure-coding")
		require.NoError(t, err)
		assert.Equal(t, plainData, decrypted)
	})
}
