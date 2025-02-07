// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package sdk_test

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/memory"
	"github.com/akave-ai/akavesdk/private/testrand"
	"github.com/akave-ai/akavesdk/sdk"
)

func TestBuildChunkDag(t *testing.T) {
	file := generate10MiBFile(t, 2024)
	actual, err := sdk.BuildDAG(context.Background(), file, 1*memory.MiB.ToInt64(), nil)
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
		chunkDAG, err := sdk.BuildDAG(context.Background(), f, memory.MiB.ToInt64(), nil)
		require.NoError(t, err)
		require.Len(t, chunkDAG.Blocks, 1)

		require.NoError(t, builder.AddLink(chunkDAG.CID, chunkDAG.RawDataSize, chunkDAG.ProtoNodeSize))

		rootCID, err := builder.Build()
		require.NoError(t, err)
		require.Equal(t, chunkDAG.CID.String(), rootCID.String())
	})

	t.Run("add chunk with multiple blocks", func(t *testing.T) {
		builder, err := sdk.NewDAGRoot()
		require.NoError(t, err)

		f := bytes.NewBuffer(testrand.BytesD(t, 2024, 10*memory.MiB.ToInt64()))
		chunkDAG, err := sdk.BuildDAG(context.Background(), f, memory.MiB.ToInt64(), nil)
		require.NoError(t, err)
		require.Len(t, chunkDAG.Blocks, 10)

		require.NoError(t, builder.AddLink(chunkDAG.CID, chunkDAG.RawDataSize, chunkDAG.ProtoNodeSize))

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

		chunk1DAG, err := sdk.BuildDAG(context.Background(), chunk1, memory.MiB.ToInt64(), nil)
		require.NoError(t, err)
		require.NoError(t, builder.AddLink(chunk1DAG.CID, chunk1DAG.RawDataSize, chunk1DAG.ProtoNodeSize))

		chunk2DAG, err := sdk.BuildDAG(context.Background(), chunk2, memory.MiB.ToInt64(), nil)
		require.NoError(t, err)
		require.NoError(t, builder.AddLink(chunk2DAG.CID, chunk2DAG.RawDataSize, chunk2DAG.ProtoNodeSize))

		rootCid, err := builder.Build()
		require.NoError(t, err)

		require.Equal(t, "bafybeigik3o6tjpdncjam4ymxqav3772lhbo6abevzezjxxm4ghzafst4i", rootCid.String())
	})
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
