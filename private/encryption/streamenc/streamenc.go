// Copyright (C) 2026 Akave
// See LICENSE for copying information.

// Package streamenc implements a custom encryption format that splits plaintext into blocks
package streamenc

import (
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/akave-ai/akavesdk/private/encryption"
	"github.com/akave-ai/akavesdk/private/memory"
)

const (
	nonceSize          = 12
	tagSize            = 16
	ecDataBlocks       = 16 // TODO: factor out erasure coding specific constants from this package
	versionSize        = 1
	plaintextSizeBytes = 4

	// MaxBlockSize is the total ciphertext block size (including overhead).
	MaxBlockSize = int(32 * memory.KiB)

	// HeaderSize is the size of the header (version + nonce + plaintext size), stored in block 0.
	HeaderSize = versionSize + nonceSize + plaintextSizeBytes // 1 byte version + 12 bytes nonce + 4 bytes plaintext size

	// Version is the current format version written into every ciphertext.
	Version = 1

	// Block0DataSize is the max plaintext capacity of block 0 (header occupies part of it).
	Block0DataSize = MaxBlockSize - HeaderSize - tagSize

	// BlockNDataSize is the max plaintext capacity of blocks 1..N-1.
	BlockNDataSize = MaxBlockSize - tagSize

	// MinCipherTextSize is the smallest valid ciphertext size: one aligned block that fits at least 1 byte of plaintext.
	MinCipherTextSize = ((HeaderSize + tagSize + 1 + ecDataBlocks - 1) / ecDataBlocks) * ecDataBlocks
)

var (
	// ErrTargetSizeNotAligned is returned when targetSize is not a multiple of ecDataBlocks.
	ErrTargetSizeNotAligned = errors.New("targetSize must be a multiple of 16 bytes")

	// ErrCiphertextSizeTooSmall is returned when targetSize is smaller than the minimum possible ciphertext.
	ErrCiphertextSizeTooSmall = fmt.Errorf("target size must be at least %d bytes", MinCipherTextSize)

	// ErrVersionMismatch is returned when the ciphertext version does not match the expected version.
	ErrVersionMismatch = errors.New("ciphertext version mismatch")
)

// ParseHeader parses the header from block 0 and returns the format version,
// initial nonce, and plaintext size. The data slice must be at least HeaderSize bytes.
func ParseHeader(data []byte) (uint8, [nonceSize]byte, uint32, error) {
	if len(data) < HeaderSize {
		return 0, [nonceSize]byte{}, 0, errors.New("header too short")
	}

	ver := data[0]
	var initialNonce [nonceSize]byte
	copy(initialNonce[:], data[versionSize:versionSize+nonceSize])
	plaintextSize := binary.BigEndian.Uint32(data[versionSize+nonceSize : HeaderSize])

	return ver, initialNonce, plaintextSize, nil
}

// NumBlocks returns the number of ciphertext blocks for a given plaintext size.
func NumBlocks(plaintextSize int) int {
	if plaintextSize == 0 {
		return 0
	}
	if plaintextSize <= Block0DataSize {
		return 1
	}
	return 1 + encryption.CeilDiv(plaintextSize-Block0DataSize, BlockNDataSize)
}

// BlockDataSize returns the actual (un-padded) count of bytes in the block at the given index.
func BlockDataSize(plaintextSize, blockIndex int) int {
	numBlocks := NumBlocks(plaintextSize)

	if blockIndex == 0 {
		if numBlocks == 1 {
			return plaintextSize
		}
		return Block0DataSize
	}

	if blockIndex < numBlocks-1 {
		return BlockNDataSize
	}

	// last block
	filled := Block0DataSize + (blockIndex-1)*BlockNDataSize
	return plaintextSize - filled
}

// EncryptedBlock returns all ciphertext bytes for the given block index,
// including the header for block 0. The ciphertext parameter is the full
// encrypted buffer returned by Encrypt.
func EncryptedBlock(ciphertext []byte, blockIndex int) []byte {
	start := blockIndex * MaxBlockSize
	if blockIndex < 0 || start >= len(ciphertext) {
		panic("blockIndex out of bounds")
	}

	end := min(start+MaxBlockSize, len(ciphertext))

	return ciphertext[start:end]
}

// Overhead returns the total overhead bytes for encrypting the given plaintext size.
// Total ciphertext size = plaintextSize + Overhead(plaintextSize).
func Overhead(plaintextSize int) int {
	if plaintextSize == 0 {
		return 0
	}

	numBlocks := NumBlocks(plaintextSize)
	lastData := BlockDataSize(plaintextSize, numBlocks-1)
	paddingSize := lastBlockPadding(numBlocks, lastData)

	lastBlockCipherSize := lastData + paddingSize + tagSize
	var totalSize int
	if numBlocks == 1 {
		totalSize = HeaderSize + lastBlockCipherSize
	} else {
		totalSize = (numBlocks-1)*MaxBlockSize + lastBlockCipherSize
	}

	return totalSize - plaintextSize
}

// MaxPlaintextSizeForTarget returns the maximum plaintext size whose ciphertext is exactly
// targetSize bytes. targetSize must be a positive multiple of ecDataBlocks.
func MaxPlaintextSizeForTarget(targetSize int) (int, error) {
	if targetSize < MinCipherTextSize {
		return 0, ErrCiphertextSizeTooSmall
	}

	if targetSize%ecDataBlocks != 0 {
		return 0, fmt.Errorf("%w: got %d, ecDataBlocks is %d", ErrTargetSizeNotAligned, targetSize, ecDataBlocks)
	}

	numBlocks := encryption.CeilDiv(targetSize, MaxBlockSize)
	if numBlocks == 1 {
		// Single block: totalSize = HeaderSize + lastData + lastPad + tagSize.
		// lastData + lastPad = targetSize - HeaderSize - tagSize; max lastData is when lastPad is 0
		return min(targetSize-HeaderSize-tagSize, Block0DataSize), nil
	}

	// Multi-block: last block contributes targetSize - (numBlocks-1)*MaxBlockSize.
	lastBlockCipherSize := targetSize - (numBlocks-1)*MaxBlockSize
	lastData := min(lastBlockCipherSize-tagSize, BlockNDataSize)

	return Block0DataSize + (numBlocks-2)*BlockNDataSize + lastData, nil
}

// Encrypt encrypts buf in-place by splitting it into fixed-size blocks, each encrypted
// with a nonce derived from a single random initial nonce. The last block is zero-padded
// to the nearest 16 byte boundary so the total ciphertext length is divisible
// by 16. The header (version + nonce + plaintext length) is packed into block 0. buf must have
// sufficient capacity to hold the encrypted result; returns encryption.ErrBufferTooSmall
// if capacity is insufficient. Returns the number of bytes written.
func Encrypt(key, buf []byte, info string) (int, error) {
	if len(buf) == 0 {
		return 0, errors.New("plaintext cannot be empty")
	}

	plaintextLen := len(buf)
	numBlocks := NumBlocks(plaintextLen)
	lastBlockIndex := numBlocks - 1
	lastData := BlockDataSize(plaintextLen, lastBlockIndex)
	padding := lastBlockPadding(numBlocks, lastData)
	lastBlockCipherSize := lastData + padding + tagSize

	var totalSize int
	if numBlocks == 1 {
		totalSize = HeaderSize + lastBlockCipherSize
	} else {
		totalSize = (numBlocks-1)*MaxBlockSize + lastBlockCipherSize
	}

	if cap(buf) < totalSize {
		return 0, encryption.ErrBufferTooSmall
	}

	gcm, err := encryption.GCMCipher(key, info)
	if err != nil {
		return 0, err
	}

	var initialNonce [nonceSize]byte
	if _, err := io.ReadFull(rand.Reader, initialNonce[:]); err != nil {
		return 0, err
	}

	out := buf[:totalSize]

	// Encrypt blocks right-to-left so each block's ciphertext never overwrites unprocessed plaintext.
	// Use a fixed temporary buffer large enough for the largest block capacity.
	var tmp [BlockNDataSize]byte
	// TODO: consider left-to-right encryption
	for i := numBlocks - 1; i >= 0; i-- {
		var capacity, plaintextStart int
		if i == 0 {
			capacity = Block0DataSize
			plaintextStart = 0
		} else {
			capacity = BlockNDataSize
			plaintextStart = Block0DataSize + (i-1)*BlockNDataSize
		}
		plaintextEnd := min(plaintextStart+capacity, plaintextLen)
		blockLen := plaintextEnd - plaintextStart

		// paddedLen is the number of bytes fed into GCM: actual plaintext plus zero-padding.
		// For the last block this is smaller than capacity to align the ciphertext to ecDataBlocks.
		// For all other blocks it equals capacity (no padding needed).
		paddedLen := capacity
		if i == lastBlockIndex {
			paddedLen = lastData + padding
		}

		copy(tmp[:blockLen], buf[plaintextStart:plaintextEnd])
		if blockLen < paddedLen {
			clear(tmp[blockLen:paddedLen]) // zero-pad the last block to a multiple of ecDataBlocks for consistent ciphertext size. The padding is stripped on decryption.
		}

		nonce := BlockNonce(initialNonce, i)

		var offset int
		if i == 0 {
			offset = HeaderSize
		} else {
			offset = i * MaxBlockSize
		}
		gcm.Seal(out[offset:offset:offset+paddedLen+tagSize], nonce[:], tmp[:paddedLen], nil)
	}

	// Write header into the start of block 0.
	out[0] = Version
	copy(out[versionSize:versionSize+nonceSize], initialNonce[:])
	binary.BigEndian.PutUint32(out[versionSize+nonceSize:HeaderSize], uint32(plaintextLen))

	return totalSize, nil
}

// DecryptBlock decrypts a single ciphertext block, writing the actual plaintext bytes
// to buf[0:n] where n is the returned count. For block 0, buf must be the full MaxBlockSize
// bytes including the header prefix. plaintextSize is the total unencrypted chunk size; it is
// used to strip trailing zero-padding from the last block so only real data bytes are returned.
// version is the expected format version; for block 0 the header version is checked against it.
func DecryptBlock(gcm cipher.AEAD, buf []byte, initialNonce [nonceSize]byte, version uint8, blockIndex, plaintextSize int) (int, error) {
	ciphertext := buf
	if blockIndex == 0 {
		if len(buf) < HeaderSize {
			return 0, errors.New("block 0 too short to contain header")
		}
		if buf[0] != version {
			return 0, fmt.Errorf("%w: got %d, want %d", ErrVersionMismatch, buf[0], version)
		}
		ciphertext = buf[HeaderSize:]
	}

	nonce := BlockNonce(initialNonce, blockIndex)
	plain, err := gcm.Open(ciphertext[:0], nonce[:], ciphertext, nil)
	if err != nil {
		return 0, err
	}

	actual := BlockDataSize(plaintextSize, blockIndex)
	clear(plain[actual:]) // zero out any trailing bytes of ciphertext
	if blockIndex == 0 {
		copy(buf[:actual], plain[:actual])
	}

	return actual, nil
}

// DecryptAllBlocks decrypts all blocks in-place, compacting the plaintext to the start of buf.
// The plaintext size is read from the header. version is the expected format version; the header
// version is checked against it and ErrVersionMismatch is returned if they differ.
// After a successful call, buf[:n] holds the recovered plaintext, where n is the returned integer.
func DecryptAllBlocks(key, buf []byte, info string, version uint8) (int, error) {
	ver, initialNonce, plaintextSize32, err := ParseHeader(buf)
	if err != nil {
		return 0, err
	}
	if ver != version {
		return 0, fmt.Errorf("%w: got %d, want %d", ErrVersionMismatch, ver, version)
	}
	plaintextSize := int(plaintextSize32)

	gcm, err := encryption.GCMCipher(key, info)
	if err != nil {
		return 0, err
	}

	numBlocks := NumBlocks(plaintextSize)
	offset := 0
	for i := range numBlocks {
		start := i * MaxBlockSize
		end := min(start+MaxBlockSize, len(buf))

		n, err := DecryptBlock(gcm, buf[start:end], initialNonce, version, i, plaintextSize)
		if err != nil {
			return 0, err
		}
		copy(buf[offset:offset+n], buf[start:start+n])
		offset += n
	}

	clear(buf[offset:]) // zero out any remaining bytes in the buffer after the plaintext

	return offset, nil
}

// BlockNonce derives the nonce for a given block index by incrementing the last 4 bytes
// of the initial nonce by the block index.
func BlockNonce(initialNonce [nonceSize]byte, blockIndex int) [nonceSize]byte {
	v := binary.BigEndian.Uint32(initialNonce[8:])
	binary.BigEndian.PutUint32(initialNonce[8:], v+uint32(blockIndex))
	return initialNonce
}

// lastBlockPadding returns the number of zero-padding bytes appended to the last block
// so the total ciphertext length is a multiple of ecDataBlocks.
func lastBlockPadding(numBlocks, lastData int) int {
	if numBlocks == 1 {
		return (ecDataBlocks - (HeaderSize+lastData)%ecDataBlocks) % ecDataBlocks
	}

	return (ecDataBlocks - lastData%ecDataBlocks) % ecDataBlocks
}
