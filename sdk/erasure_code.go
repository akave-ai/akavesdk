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

type ErasureCodeError struct {
	Code    string
	Message string
	Err     error
}

func (e *ErasureCodeError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *ErasureCodeError) Unwrap() error {
	return e.Err
}

const (
	ErrInvalidParams    = "EC001"
	ErrEncodingFailed   = "EC002"
	ErrDecodingFailed   = "EC003"
	ErrShardingFailed   = "EC004"
)

// ErasureCode is a wrapper around the reedsolomon.Encoder type, providing a more user-friendly interface.
type ErasureCode struct {
	DataBlocks   int
	ParityBlocks int

	enc reedsolomon.Encoder
}

// NewErasureCode creates a new ErasureEncode instance with improved error handling
func NewErasureCode(dataBlocks, parityBlocks int) (*ErasureCode, error) {
	if dataBlocks <= 0 || parityBlocks <= 0 {
		return nil, &ErasureCodeError{
			Code:    ErrInvalidParams,
			Message: "data and parity blocks must be greater than 0",
		}
	}

	enc, err := reedsolomon.New(dataBlocks, parityBlocks)
	if err != nil {
		return nil, &ErasureCodeError{
			Code:    ErrEncodingFailed,
			Message: "failed to create erasure encoder",
			Err:     err,
		}
	}

	return &ErasureCode{
		DataBlocks:   dataBlocks,
		ParityBlocks: parityBlocks,
		enc:          enc,
	}, nil
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
