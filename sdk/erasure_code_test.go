// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package sdk_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/sdk"
)

func TestErasureCodeInvalidParams(t *testing.T) {
	_, err := sdk.NewErasureCode(0, 0)
	require.Error(t, err)

	_, err = sdk.NewErasureCode(16, 0)
	require.Error(t, err)
}

func TestErasureCode(t *testing.T) {
	data := []byte("Quick brown fox jumps over the lazy dog")
	dataShards := 5
	parityShards := 3

	encoder, err := sdk.NewErasureCode(dataShards, parityShards)
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
		end := offset + shardSize
		if end > len(encoded) {
			end = len(encoded)
		}
		block := make([]byte, shardSize)
		copy(block, encoded[offset:end])
		blocks = append(blocks, block)
	}
	return blocks
}
