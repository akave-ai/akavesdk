// Copyright (C) 2024 Akave
// See LICENSE for copying information.

// Package erasurecode provides an implementation of Reed-Solomon erasure coding.
package erasurecode

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/klauspost/reedsolomon"
	"github.com/zeebo/errs/v2"
)

var magicSuffix = []byte{0xDE, 0xAD, 0xBE, 0xEF}

var erasureCodeErr = errs.Tag("erasure coding")

// ErasureCode is a wrapper around the reedsolomon.Encoder type, providing a more user-friendly interface.
type ErasureCode struct {
	DataBlocks   int
	ParityBlocks int

	enc reedsolomon.Encoder
}

// New creates a new ErasureEncode instance with the specified number of data and parity shards.
func New(dataBlocks, parityBlocks int) (*ErasureCode, error) {
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
	shards, err := e.enc.Split(wrapData(data))
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

// ReconstructAll reconstructs all missing blocks using Reed-Solomon erasure coding.
func (e *ErasureCode) ReconstructAll(blocks [][]byte) error {
	return erasureCodeErr.Wrap(e.enc.Reconstruct(blocks))
}

// ExtractData extracts the original data from the encoded data using Reed-Solomon erasure coding.
// TODO: remove doUnwrap once migration is completed.
func (e *ErasureCode) ExtractData(blocks [][]byte, originalSize int) ([]byte, error) {
	ok, _ := e.enc.Verify(blocks)
	if !ok {
		if err := e.enc.ReconstructData(blocks); err != nil {
			return nil, erasureCodeErr.Wrap(err)
		}
	}

	var buf bytes.Buffer
	// at this point, blocks are all reconstructed or valid, so it's safe to take length of 1st
	outSize := e.DataBlocks * len(blocks[0])
	if originalSize > 0 {
		outSize = originalSize
	}
	if err := e.enc.Join(&buf, blocks, outSize); err != nil {
		return nil, erasureCodeErr.Wrap(err)
	}

	if originalSize == 0 {
		return unwrapData(buf.Bytes())
	}

	return buf.Bytes(), nil
}

func wrapData(data []byte) []byte {
	size := uint64(len(data))
	buf := make([]byte, 8+len(data)+len(magicSuffix))
	binary.BigEndian.PutUint64(buf[:8], size)
	copy(buf[8:], data)
	copy(buf[8+len(data):], magicSuffix)
	return buf
}

func unwrapData(buf []byte) ([]byte, error) {
	unwrapDataErr := fmt.Errorf("buffer too short")

	if len(buf) < 8+len(magicSuffix) {
		return nil, unwrapDataErr
	}
	size := binary.BigEndian.Uint64(buf[:8])
	dataStart := 8
	dataEnd := dataStart + int(size)

	n := dataEnd + len(magicSuffix)

	if n > len(buf) || n < 0 {
		return nil, unwrapDataErr
	}

	if !bytes.Equal(buf[dataEnd:n], magicSuffix) {
		return nil, fmt.Errorf("missing suffix or corrupted data")
	}
	return buf[dataStart:dataEnd], nil
}
