// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package sdk

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/ipfs/boxo/ipld/merkledag"
	"github.com/ipfs/boxo/ipld/unixfs"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/erasurecode"
	"github.com/akave-ai/akavesdk/private/testrand"
)

// ecBlocks returns the DAG blocks produced by BuildDAG over erasure-coded data,
// mirroring the logic used in TestExtractBlockDataDecodeProtobufMatchesProtowire.
func ecBlocks(b *testing.B, dataBlocks, parityBlocks int) [][]byte {
	b.Helper()

	ec, err := erasurecode.New(dataBlocks, parityBlocks)
	require.NoError(b, err)

	chunkSize := int64(dataBlocks) * BlockSize.ToInt64()
	payload := testrand.Bytes(b, chunkSize)
	encoded, err := ec.Encode(payload)
	require.NoError(b, err)

	blockSize := int64(len(encoded) / (ec.DataBlocks + ec.ParityBlocks))
	dag, err := BuildDAG(b.Context(), bytes.NewBuffer(encoded), blockSize)
	require.NoError(b, err)

	raw := make([][]byte, len(dag.Blocks))
	for i, blk := range dag.Blocks {
		raw[i] = blk.Data
	}
	return raw
}

// BenchmarkExtractBlockData_Legacy measures allocations for the old path:
// merkledag.DecodeProtobuf → FSNodeFromBytes → Data().
func BenchmarkExtractBlockData_Legacy(b *testing.B) {
	for _, tc := range []struct {
		dataBlocks, parityBlocks int
	}{
		{4, 4},
		{8, 8},
		{16, 16},
	} {
		blocks := ecBlocks(b, tc.dataBlocks, tc.parityBlocks)
		b.Run(fmt.Sprintf("%d+%d", tc.dataBlocks, tc.parityBlocks), func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				for _, raw := range blocks {
					node, err := merkledag.DecodeProtobuf(raw)
					if err != nil {
						b.Fatal(err)
					}
					fsNode, err := unixfs.FSNodeFromBytes(node.Data())
					if err != nil {
						b.Fatal(err)
					}
					_ = fsNode.Data()
				}
			}
		})
	}
}

// BenchmarkExtractBlockData_Protowire measures allocations for the new path:
// consumeDAGPBDataField → consumeUnixFSDataField (zero-copy protowire scan).
func BenchmarkExtractBlockData_Protowire(b *testing.B) {
	for _, tc := range []struct {
		dataBlocks, parityBlocks int
	}{
		{4, 4},
		{8, 8},
		{16, 16},
	} {
		blocks := ecBlocks(b, tc.dataBlocks, tc.parityBlocks)
		b.Run(fmt.Sprintf("%d+%d", tc.dataBlocks, tc.parityBlocks), func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				for _, raw := range blocks {
					inner, err := consumeDAGPBDataField(raw)
					if err != nil {
						b.Fatal(err)
					}
					_, err = consumeUnixFSDataField(inner)
					if err != nil {
						b.Fatal(err)
					}
				}
			}
		})
	}
}
