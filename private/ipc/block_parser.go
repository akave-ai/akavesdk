// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package ipc

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// rpcBlock represents a block as returned by the JSON-RPC API.
type rpcBlock struct {
	Hash         *common.Hash        `json:"hash"`
	Transactions []rpcTransaction    `json:"transactions"`
	UncleHashes  []common.Hash       `json:"uncles"`
	Withdrawals  []*types.Withdrawal `json:"withdrawals,omitempty"`
}

// rpcTransaction represents a transaction as returned by the JSON-RPC API.
type rpcTransaction struct {
	tx *types.Transaction
	txExtraInfo
}

// txExtraInfo contains additional transaction information from JSON-RPC.
type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler for rpcTransaction.
func (tx *rpcTransaction) UnmarshalJSON(msg []byte) error {
	if err := json.Unmarshal(msg, &tx.tx); err != nil {
		return err
	}
	return json.Unmarshal(msg, &tx.txExtraInfo)
}

// BlockFromJSON parses a types.Block from raw JSON data, similar to ethclient.getBlock.
// code is mostly copied from ethclient.getBlock except for the use of caching senders and fetching uncles.
func BlockFromJSON(raw json.RawMessage) (*types.Block, error) {
	var head *types.Header
	if err := json.Unmarshal(raw, &head); err != nil {
		return nil, err
	}

	// When the block is not found, the API returns JSON null.
	if head == nil {
		return nil, ethereum.NotFound
	}

	var body rpcBlock
	if err := json.Unmarshal(raw, &body); err != nil {
		return nil, err
	}

	txs := make([]*types.Transaction, len(body.Transactions))
	for i, tx := range body.Transactions {
		txs[i] = tx.tx
	}

	uncleHeadersHashOnly := make([]*types.Header, len(body.UncleHashes))
	for i, uncleHash := range body.UncleHashes {
		uncleHeadersHashOnly[i] = &types.Header{
			Root: uncleHash, // just for the info saving uncleHash in Root field to keep track of possible uncles
		}
	}

	return types.NewBlockWithHeader(head).WithBody(
		types.Body{
			Transactions: txs,
			Uncles:       uncleHeadersHashOnly,
			Withdrawals:  body.Withdrawals,
		}), nil
}
