// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package ipc_test

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/require"

	"github.com/akave-ai/akavesdk/private/ipc"
	"github.com/akave-ai/akavesdk/private/ipctest"
)

func TestGetTransactionReceiptsBatch(t *testing.T) {
	var (
		ctx        = t.Context()
		dialUri    = ipctest.PickDialURI(t)
		privateKey = ipctest.PickPrivateKey(t)
	)

	pk := ipctest.NewFundedAccount(t, privateKey, dialUri, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	client, err := ipc.DeployContracts(ctx, ipc.Config{
		DialURI:    dialUri,
		PrivateKey: newPk,
	})
	require.NoError(t, err)
	batchClient := ipc.NewBatchClient(client.Eth.Client())

	const numTxs = 55
	var requests []ipc.BatchReceiptRequest

	fromAddress := crypto.PubkeyToAddress(pk.PublicKey)
	nonce, err := client.Eth.PendingNonceAt(ctx, fromAddress)
	require.NoError(t, err)

	gasPrice, err := client.Eth.SuggestGasPrice(ctx)
	require.NoError(t, err)

	chainID, err := client.Eth.NetworkID(ctx)
	require.NoError(t, err)

	for i := 0; i < numTxs; i++ {
		toAddress := common.HexToAddress("0x000000000000000000000000000000000000dead")

		tx := types.NewTx(&types.LegacyTx{
			Nonce:    nonce + uint64(i),
			To:       &toAddress,
			Value:    big.NewInt(0),
			Gas:      21000,
			GasPrice: gasPrice,
		})

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), pk)
		require.NoError(t, err)
		require.NoError(t, client.Eth.SendTransaction(ctx, signedTx))

		hash := signedTx.Hash()
		requests = append(requests, ipc.BatchReceiptRequest{
			Hash: hash,
			Key:  fmt.Sprintf("tx-%d", i),
		})
	}

	result, err := batchClient.GetTransactionReceiptsBatch(ctx, requests, 10*time.Second)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, len(requests), len(result.Responses))

	receiptsFound := 0
	receiptsNotFound := 0
	individualErrors := 0

	for i, response := range result.Responses {
		expectedKey := fmt.Sprintf("tx-%d", i)
		require.Equal(t, expectedKey, response.Key)

		if response.Error != nil {
			individualErrors++
			t.Logf("Transaction %s has error: %v", requests[i].Hash.Hex(), response.Error)
		} else if response.Receipt == nil {
			receiptsNotFound++
			t.Logf("Transaction %s not yet mined (receipt is nil)", requests[i].Hash.Hex())
		} else {
			receiptsFound++
			require.Equal(t, requests[i].Hash, response.Receipt.TxHash)
			require.True(t, response.Receipt.Status == 0 || response.Receipt.Status == 1)
			require.True(t, response.Receipt.BlockNumber.Uint64() > 0)

			t.Logf("Transaction %s mined in block %d with status %d",
				response.Receipt.TxHash.Hex(),
				response.Receipt.BlockNumber.Uint64(),
				response.Receipt.Status)
		}
	}

	require.Equal(t, len(requests), receiptsFound+receiptsNotFound+individualErrors)
}

func TestGetBlocksBatch(t *testing.T) {
	var (
		ctx        = t.Context()
		dialUri    = ipctest.PickDialURI(t)
		privateKey = ipctest.PickPrivateKey(t)
	)

	pk := ipctest.NewFundedAccount(t, privateKey, dialUri, ipctest.ToWei(10))
	newPk := hexutil.Encode(crypto.FromECDSA(pk))[2:]

	client, err := ipc.DeployContracts(ctx, ipc.Config{
		DialURI:    dialUri,
		PrivateKey: newPk,
	})
	require.NoError(t, err)
	t.Cleanup(client.Eth.Close)

	batchClient := ipc.NewBatchClient(client.Eth.Client())

	fromAddress := crypto.PubkeyToAddress(pk.PublicKey)
	nonce, err := client.Eth.PendingNonceAt(ctx, fromAddress)
	require.NoError(t, err)

	gasPrice, err := client.Eth.SuggestGasPrice(ctx)
	require.NoError(t, err)

	chainID, err := client.Eth.NetworkID(ctx)
	require.NoError(t, err)

	numTxs := 5
	txHashes := make([]common.Hash, 0, numTxs)
	for i := range numTxs {
		toAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")
		tx := types.NewTx(&types.LegacyTx{
			Nonce:    nonce + uint64(i),
			To:       &toAddress,
			Value:    big.NewInt(0),
			Gas:      21000,
			GasPrice: gasPrice,
		})
		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), pk)
		require.NoError(t, err)
		require.NoError(t, client.Eth.SendTransaction(ctx, signedTx))
		txHashes = append(txHashes, signedTx.Hash())

		// Wait for the transaction to be mined
		require.NoError(t, client.WaitForTx(ctx, signedTx.Hash()))
	}

	blockNumbers := make([]rpc.BlockNumber, 0, numTxs)
	for _, hash := range txHashes {
		receipt, err := client.Eth.TransactionReceipt(ctx, hash)
		require.NoError(t, err)
		blockNumbers = append(blockNumbers, rpc.BlockNumber(receipt.BlockNumber.Int64()))
	}

	// Add a non-existent block number to test NotFound error
	nonExistentBlock := 99999999
	blockNumbers = append(blockNumbers, rpc.BlockNumber(nonExistentBlock))

	responses, err := batchClient.GetBlocksBatch(ctx, blockNumbers)
	require.NoError(t, err)
	require.Equal(t, len(blockNumbers), len(responses))

	foundBlocks := 0
	notFoundBlocks := 0
	for i, resp := range responses {
		if resp.Error != nil {
			require.ErrorIs(t, resp.Error, ethereum.NotFound)
			require.Nil(t, resp.Block)
			notFoundBlocks++
		} else {
			require.NotNil(t, resp.Block)
			require.Equal(t, blockNumbers[i].Int64(), resp.Block.Number().Int64())
			foundBlocks++
		}
	}
	require.GreaterOrEqual(t, foundBlocks, 1)
	require.Equal(t, 1, notFoundBlocks)
}
