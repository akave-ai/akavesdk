// Copyright (C) 2025 Akave
// See LICENSE for copying information.

package ipc

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ipfs/go-cid"
)

const (
	contractMethodSignatureLen = 4
)

// AddChunkTransactionData represents the parsed parameters from an addFileChunk transaction.
type AddChunkTransactionData struct {
	CID         cid.Cid
	BucketID    [32]byte
	FileName    string
	EncodedSize int64
	BlockCIDs   []cid.Cid
	BlockSizes  []int64
	Index       int64
}

// ParseAddChunkTx parses the transaction data of an addFileChunk call and returns structured metadata.
func ParseAddChunkTx(storageContractABI *abi.ABI, txData []byte) (AddChunkTransactionData, error) {
	if len(txData) < contractMethodSignatureLen {
		return AddChunkTransactionData{}, fmt.Errorf("invalid transaction data length")
	}

	method, ok := storageContractABI.Methods["addFileChunk"]
	if !ok {
		return AddChunkTransactionData{}, fmt.Errorf("method addFileChunk not found in ABI")
	}

	v, err := method.Inputs.Unpack(txData[contractMethodSignatureLen:])
	if err != nil {
		return AddChunkTransactionData{}, fmt.Errorf("failed to unpack transaction data: %w", err)
	}

	// Parse all parameters in one go using method.Inputs.Copy
	var params addFileChunkParams
	if err := method.Inputs.Copy(&params, v); err != nil {
		return AddChunkTransactionData{}, fmt.Errorf("failed to parse transaction data: %w", err)
	}

	chunkCID, err := cid.Cast(params.ChunkCID)
	if err != nil {
		return AddChunkTransactionData{}, fmt.Errorf("invalid chunk CID: %w", err)
	}

	blockCIDs := make([]cid.Cid, len(params.BlockCIDs))
	for i, b := range params.BlockCIDs {
		blockCIDs[i], err = FromByteArrayCID(b)
		if err != nil {
			return AddChunkTransactionData{}, fmt.Errorf("invalid block CID at index %d: %w", i, err)
		}
	}

	blockSizes := make([]int64, len(params.ChunkBlocksSizes))
	for i, size := range params.ChunkBlocksSizes {
		blockSizes[i] = size.Int64()
	}

	if len(blockCIDs) != len(blockSizes) {
		return AddChunkTransactionData{}, fmt.Errorf("mismatched block CIDs and sizes lengths")
	}

	return AddChunkTransactionData{
		CID:         chunkCID,
		BucketID:    params.BucketID,
		FileName:    params.FileName,
		EncodedSize: params.EncodedChunkSize.Int64(),
		BlockCIDs:   blockCIDs,
		BlockSizes:  blockSizes,
		Index:       params.ChunkIndex.Int64(),
	}, nil
}

// ParseAddChunksTx parses the transaction data of an addFileChunks call and returns structured metadata for multiple chunks.
func ParseAddChunksTx(storageContractABI *abi.ABI, txData []byte) ([]AddChunkTransactionData, error) {
	if len(txData) < contractMethodSignatureLen {
		return nil, fmt.Errorf("invalid transaction data length")
	}

	method, ok := storageContractABI.Methods["addFileChunks"]
	if !ok {
		return nil, fmt.Errorf("method addFileChunks not found in ABI")
	}

	v, err := method.Inputs.Unpack(txData[contractMethodSignatureLen:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack transaction data: %w", err)
	}

	var params addFileChunksParams
	if err := method.Inputs.Copy(&params, v); err != nil {
		return nil, fmt.Errorf("failed to parse transaction data: %w", err)
	}

	numChunks := len(params.ChunkCIDs)
	chunks := make([]AddChunkTransactionData, numChunks)
	startingIndex := params.StartingChunkIndex.Int64()

	for i := 0; i < numChunks; i++ {
		chunkCID, err := cid.Cast(params.ChunkCIDs[i])
		if err != nil {
			return nil, fmt.Errorf("invalid chunk CID at index %d: %w", i, err)
		}

		blockCIDs := make([]cid.Cid, len(params.ChunkBlocksCIDs[i]))
		for j, b := range params.ChunkBlocksCIDs[i] {
			blockCIDs[j], err = FromByteArrayCID(b)
			if err != nil {
				return nil, fmt.Errorf("invalid block CID at chunk %d, block %d: %w", i, j, err)
			}
		}

		blockSizes := make([]int64, len(params.ChunkBlocksSizes[i]))
		for j, size := range params.ChunkBlocksSizes[i] {
			blockSizes[j] = size.Int64()
		}

		if len(blockCIDs) != len(blockSizes) {
			return nil, fmt.Errorf("mismatched block CIDs and sizes for chunk %d", i)
		}

		chunks[i] = AddChunkTransactionData{
			CID:         chunkCID,
			BucketID:    params.BucketID,
			FileName:    params.FileName,
			EncodedSize: params.EncodedChunkSizes[i].Int64(),
			BlockCIDs:   blockCIDs,
			BlockSizes:  blockSizes,
			Index:       startingIndex + int64(i),
		}
	}

	return chunks, nil
}

// addFileChunkParams represents the raw ABI-decoded parameters from addFileChunk transaction.
type addFileChunkParams struct {
	ChunkCID         []byte     `abi:"chunkCID"`
	BucketID         [32]byte   `abi:"bucketId"`
	FileName         string     `abi:"fileName"`
	EncodedChunkSize *big.Int   `abi:"encodedChunkSize"`
	BlockCIDs        [][32]byte `abi:"cids"`
	ChunkBlocksSizes []*big.Int `abi:"chunkBlocksSizes"`
	ChunkIndex       *big.Int   `abi:"chunkIndex"`
}

// addFileChunksParams represents the raw ABI-decoded parameters from addFileChunks transaction.
type addFileChunksParams struct {
	ChunkCIDs          [][]byte     `abi:"cids"`
	BucketID           [32]byte     `abi:"bucketId"`
	FileName           string       `abi:"fileName"`
	EncodedChunkSizes  []*big.Int   `abi:"encodedChunkSizes"`
	ChunkBlocksCIDs    [][][32]byte `abi:"chunkBlocksCIDs"`
	ChunkBlocksSizes   [][]*big.Int `abi:"chunkBlockSizes"`
	StartingChunkIndex *big.Int     `abi:"startingChunkIndex"`
}
