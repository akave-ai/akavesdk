// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package ipc_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/ipc"
)

func TestParseBlockFromJSONValidBlock(t *testing.T) {
	// Generate a private key for signing
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	tx := types.NewTransaction(
		0, // nonce
		common.HexToAddress("0x1234567890123456789012345678901234567891"), // to
		big.NewInt(1000),       // value
		21000,                  // gas limit
		big.NewInt(1000000000), // gas price
		nil,                    // data
	)

	chainID := big.NewInt(1)
	signer := types.LatestSignerForChainID(chainID)
	signedTx, err := types.SignTx(tx, signer, privateKey)
	require.NoError(t, err)

	uncleHash := common.HexToHash("0xdef7890123456789012345678901234567890123456789012345678901234567")

	header := &types.Header{
		ParentHash:  common.Hash{},
		UncleHash:   types.EmptyUncleHash,
		Root:        common.HexToHash("0xd7f8974fb5ac78d9ac099b9ad5018bedc2ce0a72dad1827a1709da30580f0544"),
		TxHash:      types.EmptyRootHash,
		ReceiptHash: types.EmptyRootHash,
		MixDigest:   common.Hash{},
		Nonce:       types.BlockNonce{},
	}

	blockData := map[string]any{
		"number":           "0x1",
		"parentHash":       header.ParentHash.Hex(),
		"nonce":            "0x0000000000000000",
		"sha3Uncles":       header.UncleHash.Hex(),
		"logsBloom":        "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		"transactionsRoot": header.TxHash.Hex(),
		"stateRoot":        header.Root.Hex(),
		"receiptsRoot":     header.ReceiptHash.Hex(),
		"miner":            "0x0000000000000000000000000000000000000000",
		"difficulty":       "0x0",
		"totalDifficulty":  "0x0",
		"extraData":        "0x",
		"size":             "0x219",
		"gasLimit":         "0x1c9c380",
		"gasUsed":          "0x5208",
		"timestamp":        "0x61bc6b5d",
		"transactions":     []any{},
		"uncles":           []string{uncleHash.Hex()},
	}

	// Add the signed transaction to the block data
	txData := make(map[string]any)
	txJSON, err := signedTx.MarshalJSON()
	require.NoError(t, err)
	err = json.Unmarshal(txJSON, &txData)
	require.NoError(t, err)
	blockData["transactions"] = []any{txData}

	// Marshal the complete block data to JSON
	blockJSON, err := json.Marshal(blockData)
	require.NoError(t, err)

	expectedTxHash := signedTx.Hash()
	expectedUncleHash := uncleHash

	parsedBlock, err := ipc.BlockFromJSON(json.RawMessage(blockJSON))
	require.NoError(t, err)
	require.NotNil(t, parsedBlock)
	require.Len(t, parsedBlock.Transactions(), 1)
	require.Equal(t, expectedTxHash, parsedBlock.Transactions()[0].Hash())
	require.Len(t, parsedBlock.Uncles(), 1)
	require.Equal(t, expectedUncleHash, parsedBlock.Uncles()[0].Root)
}

func TestParseBlockFromJSONNullBlock(t *testing.T) {
	block, err := ipc.BlockFromJSON(json.RawMessage("null"))
	require.ErrorIs(t, err, ethereum.NotFound)
	require.Nil(t, block)
}

func TestParseBlockFromJSONInvalidJSON(t *testing.T) {
	block, err := ipc.BlockFromJSON(json.RawMessage("{invalid json"))
	require.Error(t, err)
	require.Nil(t, block)
}

func TestParseBlockFromJSONEmptyTransactionsAndUncles(t *testing.T) {
	blockData := map[string]any{
		"number":           "0x2",
		"parentHash":       "0x0000000000000000000000000000000000000000000000000000000000000000",
		"nonce":            "0x0000000000000000",
		"sha3Uncles":       "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
		"logsBloom":        "0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		"transactionsRoot": "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
		"stateRoot":        "0xd7f8974fb5ac78d9ac099b9ad5018bedc2ce0a72dad1827a1709da30580f0544",
		"receiptsRoot":     "0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421",
		"miner":            "0x0000000000000000000000000000000000000000",
		"difficulty":       "0x0",
		"totalDifficulty":  "0x0",
		"extraData":        "0x",
		"size":             "0x219",
		"gasLimit":         "0x1c9c380",
		"gasUsed":          "0x0",
		"timestamp":        "0x61bc6b5d",
		"transactions":     []any{},
		"uncles":           []any{},
	}

	blockJSON, err := json.Marshal(blockData)
	require.NoError(t, err)

	block, err := ipc.BlockFromJSON(json.RawMessage(blockJSON))
	require.NoError(t, err)
	require.NotNil(t, block)
	// The block hash will be calculated based on the header fields, not the JSON "hash" field
	require.NotEqual(t, common.Hash{}, block.Hash()) // Just verify it's not zero
	require.Len(t, block.Transactions(), 0)
	require.Len(t, block.Uncles(), 0)
}
