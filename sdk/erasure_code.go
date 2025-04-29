// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package sdk

import (
	"bytes"
	"fmt"

	"github.com/klauspost/reedsolomon"
	"github.com/zeebo/errs/v2"
)

var erasureCodeErr = errs.Tag("erasure coding")

// ErasureCode is a wrapper around the reedsolomon.Encoder type, providing a more user-friendly interface.
type ErasureCode struct {
	DataBlocks   int
	ParityBlocks int

	enc reedsolomon.Encoder
}

// NewErasureCode creates a new ErasureEncode instance with the specified number of data and parity shards.
func NewErasureCode(dataBlocks, parityBlocks int) (*ErasureCode, error) {
	if dataBlocks <= 0 || parityBlocks <= 0 {
		return &ErasureCode{}, erasureCodeErr.Wrap(fmt.Errorf("data and parity shards must be > 0"))
	}

	enc, err := reedsolomon.New(dataBlocks, parityBlocks)
	if err != nil {
		return &ErasureCode{}, erasureCodeErr.Wrap(err)
	}
	return &ErasureCode{DataBlocks: dataBlocks, ParityBlocks: parityBlocks, enc: enc}, nil
}

// Encode encodes the input data using Reed-Solomon erasure coding, returning the encoded data.
func (e *ErasureCode) Encode(data []byte) ([]byte, error) {
	shards, err := e.enc.Split(data)
	if err != nil {
		return nil, erasureCodeErr.Wrap(err)
	}

	if err := e.enc.Encode(shards); err != nil {
		return nil, erasureCodeErr.Wrap(err)
	}

	// Concatenate all shards into a single byte slice
	shardSize := len(shards[0])
	result := make([]byte, 0, shardSize*len(shards))
	for _, shard := range shards {
		result = append(result, shard...)
	}

	return result, nil
}

// ExtractData extracts the original data from the encoded data using Reed-Solomon erasure coding.
func (e *ErasureCode) ExtractData(blocks [][]byte, originalDataSize int) ([]byte, error) {
	ok, _ := e.enc.Verify(blocks)
	if !ok {
		if err := e.enc.ReconstructData(blocks); err != nil {
			return nil, erasureCodeErr.Wrap(err)
		}
	}

	var buf bytes.Buffer
	if err := e.enc.Join(&buf, blocks, originalDataSize); err != nil {
		return nil, erasureCodeErr.Wrap(err)
	}
	return buf.Bytes(), nil
}
