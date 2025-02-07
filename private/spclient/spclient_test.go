// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package spclient_test

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/ipfs/boxo/blockservice"
	blockstore "github.com/ipfs/boxo/blockstore"
	chunker "github.com/ipfs/boxo/chunker"
	"github.com/ipfs/boxo/ipld/merkledag"
	blockformat "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/sync"
	offline "github.com/ipfs/go-ipfs-exchange-offline"
	"github.com/ipfs/go-unixfs/importer/balanced"
	"github.com/ipfs/go-unixfs/importer/helpers"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/memory"
	"github.com/akave-ai/akavesdk/private/spclient"
	"github.com/akave-ai/akavesdk/private/testrand"
)

func TestFetchBlockFromFilecoin(t *testing.T) {
	ctx := context.Background()
	file := generateDeterministic100MiBFile(t, 2024)
	chunkSize := 1 * memory.MB.ToInt64()
	blocks := splitToBlocks(t, file, chunkSize)
	boosterHttpAddress := PickBoosterHTTPAddress(t)

	spClient := spclient.New()

	for _, block := range blocks {
		fetchedBlock, err := spClient.FetchBlock(ctx, boosterHttpAddress, block.Cid())
		if err != nil {
			t.Errorf("failed to fetch block with CID %s: %v", block.Cid().String(), err)
			continue
		}

		require.Equal(t, block.Cid(), fetchedBlock.Cid())
		require.Equal(t, block.RawData(), fetchedBlock.RawData())
		require.Equal(t, len(block.RawData()), len(fetchedBlock.RawData()))

		fmt.Printf("Block %s: data matches between blockstore and Filecoin node; original size=%d filecoin size=%d\n", block.Cid().String(), len(block.RawData()), len(fetchedBlock.RawData()))
	}
}

func generateDeterministic100MiBFile(t *testing.T, seed int64) *bytes.Buffer {
	size := 100 * memory.MiB.ToInt64()
	randomData := testrand.BytesD(t, seed, size)

	return bytes.NewBuffer(randomData)
}

// BoosterHTTPAddress is flag to set RPC address of akave node.
var BoosterHTTPAddress = flag.String("booster-http-address", os.Getenv("BOOSTER_HTTP_ADDRESS"), "flag to set booster HTTP address")

// PickBoosterHTTPAddress picks booster HTTP address from flag.
func PickBoosterHTTPAddress(t testing.TB) string {
	if *BoosterHTTPAddress == "" || strings.EqualFold(*BoosterHTTPAddress, "omit") {
		t.Skip("booster http address flag missing, example: -BOOSTER_HTTP_ADDRESS=<booster http address>")
	}
	return *BoosterHTTPAddress
}

func splitToBlocks(t *testing.T, reader io.Reader, chunkSize int64) []*blockformat.BasicBlock {
	cidBuilder, err := merkledag.PrefixForCidVersion(1)
	require.NoError(t, err)

	splitter := chunker.NewSizeSplitter(io.NopCloser(reader), chunkSize)

	batching := sync.MutexWrap(datastore.NewMapDatastore())
	store := blockstore.NewBlockstore(batching)
	dagServ := merkledag.NewDAGService(blockservice.New(store, offline.Exchange(store)))

	params := helpers.DagBuilderParams{
		Maxlinks:   1024,
		RawLeaves:  false, // 1048590 pb encoded ,1048576 - raw; merkledag.ProtoNode merkledag.RawNode
		CidBuilder: cidBuilder,
		Dagserv:    dagServ,
		NoCopy:     false,
	}

	db, err := params.New(splitter)
	require.NoError(t, err)

	rootNode, err := balanced.Layout(db)
	require.NoError(t, err)

	ctx := context.Background()
	var blocks = make([]*blockformat.BasicBlock, 0)
	if len(rootNode.Links()) == 0 {
		// Single chunk case, the root node itself is the chunk
		node, err := dagServ.Get(ctx, rootNode.Cid())
		require.NoError(t, err)
		block, err := blockformat.NewBlockWithCid(node.RawData(), rootNode.Cid())
		require.NoError(t, err)
		blocks = append(blocks, block)
	} else {
		for _, l := range rootNode.Links() {
			node, err := dagServ.Get(ctx, l.Cid)
			require.NoError(t, err)
			block, err := blockformat.NewBlockWithCid(node.RawData(), node.Cid())
			require.NoError(t, err)
			blocks = append(blocks, block)
		}
	}
	return blocks
}
