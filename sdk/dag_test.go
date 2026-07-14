// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package sdk_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/erasurecode"
	"github.com/akave-ai/akavesdk/private/memory"
	"github.com/akave-ai/akavesdk/private/testrand"
	"github.com/akave-ai/akavesdk/sdk"
)

func TestBuildChunkDag(t *testing.T) {
	file := generate10MiBFile(t, 2024)
	actual, err := sdk.BuildDAG(t.Context(), file, 1*memory.MiB.ToInt64())
	require.NotNil(t, actual)
	require.NoError(t, err)

	expected := expectedDAG(t)

	require.Equal(t, actual.CID.String(), expected.CID.String())
	require.Equal(t, len(actual.Blocks), len(expected.Blocks))
	require.Equal(t, actual.RawDataSize, expected.RawDataSize)

	for i := range actual.Blocks {
		require.Equal(t, expected.Blocks[i].CID, actual.Blocks[i].CID)
		require.Len(t, actual.Blocks[i].Data, 1048590)
	}
}

func TestChunkSize(t *testing.T) {
	t.Run("without erasure coding", func(t *testing.T) {
		file := generate10MiBFile(t, 2024)

		actual, err := sdk.BuildDAG(t.Context(), file, memory.MiB.ToInt64())
		require.NoError(t, err)
		require.Equal(t, uint64(10485900), actual.EncodedSize)

		var blocksTotal uint64
		for _, block := range actual.Blocks {
			blocksTotal += uint64(len(block.Data))
		}
		require.Equal(t, blocksTotal, actual.EncodedSize)
	})

	t.Run("with erasure coding", func(t *testing.T) {
		ec, err := erasurecode.New(16, 16)
		require.NoError(t, err)

		file := generate10MiBFile(t, 2024)

		data, err := ec.Encode(file.Bytes())
		require.NoError(t, err)

		blockSize := int64(len(data) / (ec.DataBlocks + ec.ParityBlocks))

		actual, err := sdk.BuildDAG(t.Context(), bytes.NewBuffer(data), blockSize)
		require.NoError(t, err)
		require.Equal(t, uint64(20972000), actual.EncodedSize)

		var blocksTotal uint64
		for _, block := range actual.Blocks {
			blocksTotal += uint64(len(block.Data))
		}
		require.Equal(t, blocksTotal, actual.EncodedSize)
	})
}

func TestRootCIDBuilder(t *testing.T) {
	t.Run("build root cid with no chunks", func(t *testing.T) {
		builder, err := sdk.NewDAGRoot()
		require.NoError(t, err)

		rootCID, err := builder.Build()
		require.Error(t, err)

		require.Equal(t, "no chunks added", err.Error())
		require.Equal(t, cid.Undef, rootCID)
	})

	t.Run("add chunk with one block", func(t *testing.T) {
		builder, err := sdk.NewDAGRoot()
		require.NoError(t, err)

		f := bytes.NewBuffer(testrand.BytesD(t, 2024, memory.MiB.ToInt64()))
		chunkDAG, err := sdk.BuildDAG(t.Context(), f, memory.MiB.ToInt64())
		require.NoError(t, err)
		require.Len(t, chunkDAG.Blocks, 1)

		require.NoError(t, builder.AddLink(chunkDAG.CID, chunkDAG.RawDataSize, chunkDAG.EncodedSize))

		rootCID, err := builder.Build()
		require.NoError(t, err)
		require.Equal(t, chunkDAG.CID.String(), rootCID.String())
	})

	t.Run("add chunk with multiple blocks", func(t *testing.T) {
		builder, err := sdk.NewDAGRoot()
		require.NoError(t, err)

		f := bytes.NewBuffer(testrand.BytesD(t, 2024, 10*memory.MiB.ToInt64()))
		chunkDAG, err := sdk.BuildDAG(t.Context(), f, memory.MiB.ToInt64())
		require.NoError(t, err)
		require.Len(t, chunkDAG.Blocks, 10)

		require.NoError(t, builder.AddLink(chunkDAG.CID, chunkDAG.RawDataSize, chunkDAG.EncodedSize))

		rootCID, err := builder.Build()
		require.NoError(t, err)
		require.Equal(t, chunkDAG.CID.String(), rootCID.String())
	})

	t.Run("add multiple chunks", func(t *testing.T) {
		builder, err := sdk.NewDAGRoot()
		require.NoError(t, err)

		f := bytes.NewBuffer(testrand.BytesD(t, 2024, 64*memory.MiB.ToInt64()))
		chunk1 := io.LimitReader(f, 32*memory.MiB.ToInt64())
		chunk2 := io.LimitReader(f, 32*memory.MiB.ToInt64())

		chunk1DAG, err := sdk.BuildDAG(t.Context(), chunk1, memory.MiB.ToInt64())
		require.NoError(t, err)
		require.NoError(t, builder.AddLink(chunk1DAG.CID, chunk1DAG.RawDataSize, chunk1DAG.EncodedSize))

		chunk2DAG, err := sdk.BuildDAG(t.Context(), chunk2, memory.MiB.ToInt64())
		require.NoError(t, err)
		require.NoError(t, builder.AddLink(chunk2DAG.CID, chunk2DAG.RawDataSize, chunk2DAG.EncodedSize))

		rootCid, err := builder.Build()
		require.NoError(t, err)

		require.Equal(t, "bafybeiamgcn2bye63nlzuvilqrc6pqt7jsn5oyp2wx7pkeu4eyxgzglvpi", rootCid.String())
	})
}

func TestDAGRootBuildMatchesBuildDAG(t *testing.T) {
	const maxBlocks = 32

	tests := []struct {
		name      string
		totalSize int64
	}{
		{"32MiB", 32 * memory.MiB.ToInt64()},
		{"30MiB", 30 * memory.MiB.ToInt64()},
		{"16MiB", 16 * memory.MiB.ToInt64()},
		{"5MiB", 5 * memory.MiB.ToInt64()},
		{"1MB", 1 * memory.MB.ToInt64()},
		{"142 bytes", 142},
		{"13 bytes", 13},
		{"1 byte", 1},
	}

	for _, tc := range tests {
		blockSize := (tc.totalSize + maxBlocks - 1) / maxBlocks
		expectedBlocks := int((tc.totalSize + blockSize - 1) / blockSize)
		t.Run(fmt.Sprintf("%s/blockSize=%d", tc.name, blockSize), func(t *testing.T) {
			data := testrand.BytesD(t, 2026, tc.totalSize)

			dagResult, err := sdk.BuildDAG(t.Context(), bytes.NewBuffer(data), blockSize)
			require.NoError(t, err)
			require.Len(t, dagResult.Blocks, expectedBlocks)

			root, err := sdk.NewDAGRoot()
			require.NoError(t, err)

			var totalRawDataSize, totalEncodedSize uint64
			for i, block := range dagResult.Blocks {
				start := int64(i) * blockSize
				end := min(start+blockSize, tc.totalSize)
				blockData := data[start:end]

				node, err := sdk.BuildLeafNode(blockData)
				require.NoError(t, err)

				rawDataSize := uint64(len(blockData))
				encodedSize := uint64(len(node.RawData()))

				require.Equal(t, block.CID, node.Cid().String(), "block %d CID mismatch", i)
				require.Equal(t, uint64(len(block.Data)), encodedSize, "block %d encoded size mismatch", i)

				totalRawDataSize += rawDataSize
				totalEncodedSize += encodedSize

				require.NoError(t, root.AddLink(node.Cid(), rawDataSize, encodedSize))
			}

			rootCID, err := root.Build()
			require.NoError(t, err)
			require.Equal(t, dagResult.CID.String(), rootCID.String())
			require.Equal(t, dagResult.RawDataSize, totalRawDataSize)
			require.Equal(t, dagResult.EncodedSize, totalEncodedSize)
		})
	}
}

func expectedDAG(t *testing.T) sdk.ChunkDAG {
	// retrieved using following command:
	// ipfs add --cid-version=1 --nocopy=false --chunker=size-1048576 --raw-leaves=false file.txt
	rootCid, err := cid.Parse("bafybeifir7qtrwocso27rscbwlf53p7na4ry3pyauoyilc22lotjkx4pji")
	require.NoError(t, err)

	return sdk.ChunkDAG{
		CID:         rootCid,
		RawDataSize: uint64(10 * memory.MiB.ToInt64()),
		Blocks: []sdk.FileBlockUpload{
			{CID: "bafybeid3roxuooczpetsejm7xblw26rxohzjjl3xy3cnf6ovzfxxi3sapa"},
			{CID: "bafybeigfjuysrwis5ynbcmrq2skbqx4htxx4i6dstqaqxgveje4wlw6b3m"},
			{CID: "bafybeicth7txqbqzbv522rigdlznzf2d4t4fkbeaio4bznholhcjycydpa"},
			{CID: "bafybeigf3eobgp665rmxndubsdft5pw7l6pgbgzmj4whhsplzyihdpkfzq"},
			{CID: "bafybeidgkteds7m3h7vewpk5p2lbuqkijjyyzmi43tt4dpwaywwgzsaaui"},
			{CID: "bafybeidje7v5yqm4vocwcsu44gvchdkfh6cc7ddycx43zrbxp5h7zzw5fe"},
			{CID: "bafybeibb3v6eeo7diwpjjmvt7ikca4akbrusk3itzp2xshje3g35b26gie"},
			{CID: "bafybeiflijdz4ia7yqsa736iws7nqwsvkrwot7x3aagc2mimgy4tbp4p3a"},
			{CID: "bafybeib77hlwg5gn46ycgh4ml4iavt4a3byoc24grcmmy5gxznqwqnwkfa"},
			{CID: "bafybeifqwspkmotwkeaxhus6mvvwys4rqem4p2h46bc3veeyqh5xbndgbm"},
		},
	}
}
