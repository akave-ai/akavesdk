// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package ipc

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// BatchReceiptRequest represents a single receipt request in a batch.
type BatchReceiptRequest struct {
	Hash common.Hash
	Key  string
}

// BatchReceiptResponse represents the response for a single receipt request.
type BatchReceiptResponse struct {
	Receipt *types.Receipt
	Error   error
	Key     string
}

// BatchReceiptResult contains all responses from a batch request.
type BatchReceiptResult struct {
	Responses []BatchReceiptResponse
}

// BatchBlockResponse represents the response for a single block request.
type BatchBlockResponse struct {
	BlockNumber rpc.BlockNumber
	Block       *types.Block
	Error       error
}

// BatchClient provides batch operations for the IPC client, allowing batched requests for receipts and blocks.
type BatchClient struct {
	client *rpc.Client
}

// NewBatchClient creates a new BatchClient with the given regular client and timeout.
func NewBatchClient(client *rpc.Client) *BatchClient {
	return &BatchClient{client: client}
}

// GetTransactionReceiptsBatch fetches multiple transaction receipts in a single batch call.
func (bc *BatchClient) GetTransactionReceiptsBatch(ctx context.Context, requests []BatchReceiptRequest, timeout time.Duration) (*BatchReceiptResult, error) {
	batchReqs := make([]rpc.BatchElem, len(requests))
	for i, req := range requests {
		batchReqs[i] = rpc.BatchElem{
			Method: "eth_getTransactionReceipt",
			Args:   []any{req.Hash},
			Result: new(*types.Receipt),
		}
	}

	batchCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := bc.client.BatchCallContext(batchCtx, batchReqs); err != nil {
		return nil, err
	}

	responses := make([]BatchReceiptResponse, len(requests))
	for i, req := range requests {
		batchElem := batchReqs[i]
		response := BatchReceiptResponse{
			Key: req.Key,
		}

		if batchElem.Error != nil {
			response.Error = batchElem.Error
		} else {
			response.Receipt = *batchElem.Result.(**types.Receipt)

			if response.Receipt == nil {
				response.Error = ethereum.NotFound
			}
		}

		responses[i] = response
	}

	return &BatchReceiptResult{Responses: responses}, nil
}

// GetBlocksBatch fetches multiple blocks in a single batch call.
func (bc *BatchClient) GetBlocksBatch(ctx context.Context, blockNumbers []rpc.BlockNumber) ([]BatchBlockResponse, error) {
	results := make([]json.RawMessage, len(blockNumbers))
	batchReqs := make([]rpc.BatchElem, len(blockNumbers))

	for i, blockNumber := range blockNumbers {
		batchReqs[i] = rpc.BatchElem{
			Method: "eth_getBlockByNumber",
			Args:   []any{blockNumber.String(), true},
			Result: &results[i],
		}
	}

	if err := bc.client.BatchCallContext(ctx, batchReqs); err != nil {
		return nil, err
	}

	responses := make([]BatchBlockResponse, len(blockNumbers))
	for i, blockNumber := range blockNumbers {
		batchElem := batchReqs[i]
		response := BatchBlockResponse{
			BlockNumber: blockNumber,
		}

		if batchElem.Error != nil {
			response.Error = batchElem.Error
		} else {
			block, err := BlockFromJSON(results[i])
			if err != nil {
				response.Error = err
			} else {
				response.Block = block
			}
		}

		responses[i] = response
	}

	return responses, nil
}
