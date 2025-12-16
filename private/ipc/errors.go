// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package ipc

import (
	"errors"

	"github.com/ethereum/go-ethereum/rpc"
)

var (
	// ErrBucketAlreadyExists indicates a bucket with the same name already exists.
	ErrBucketAlreadyExists = errors.New("BucketAlreadyExists")
	// ErrBucketInvalid indicates the bucket data is invalid.
	ErrBucketInvalid = errors.New("BucketInvalid")
	// ErrBucketInvalidOwner indicates the bucket owner is invalid.
	ErrBucketInvalidOwner = errors.New("BucketInvalidOwner")
	// ErrBucketNonexists indicates the bucket does not exist.
	ErrBucketNonexists = errors.New("BucketNonexists")
	// ErrBucketNonempty indicates the bucket is not empty.
	ErrBucketNonempty = errors.New("BucketNonempty")
	// ErrFileAlreadyExists indicates a file with the same name already exists.
	ErrFileAlreadyExists = errors.New("FileAlreadyExists")
	// ErrFileInvalid indicates the file data is invalid.
	ErrFileInvalid = errors.New("FileInvalid")
	// ErrFileNonexists indicates the file does not exist.
	ErrFileNonexists = errors.New("FileNonexists")
	// ErrFileNonempty indicates the file is not empty.
	ErrFileNonempty = errors.New("FileNonempty")
	// ErrFileNameDuplicate indicates a file name is duplicated.
	ErrFileNameDuplicate = errors.New("FileNameDuplicate")
	// ErrFileFullyUploaded indicates the file is already fully uploaded.
	ErrFileFullyUploaded = errors.New("FileFullyUploaded")
	// ErrFileChunkDuplicate indicates a file chunk is duplicated.
	ErrFileChunkDuplicate = errors.New("FileChunkDuplicate")
	// ErrBlockAlreadyExists indicates a block with the same identifier already exists.
	ErrBlockAlreadyExists = errors.New("BlockAlreadyExists")
	// ErrBlockInvalid indicates the block data is invalid.
	ErrBlockInvalid = errors.New("BlockInvalid")
	// ErrBlockNonexists indicates the block does not exist.
	ErrBlockNonexists = errors.New("BlockNonexists")
	// ErrInvalidArrayLength indicates an array length is invalid.
	ErrInvalidArrayLength = errors.New("InvalidArrayLength")
	// ErrInvalidFileBlocksCount indicates the file blocks count is invalid.
	ErrInvalidFileBlocksCount = errors.New("InvalidFileBlocksCount")
	// ErrInvalidLastBlockSize indicates the last block size is invalid.
	ErrInvalidLastBlockSize = errors.New("InvalidLastBlockSize")
	// ErrInvalidEncodedSize indicates the encoded size is invalid.
	ErrInvalidEncodedSize = errors.New("InvalidEncodedSize")
	// ErrInvalidFileCID indicates the file CID is invalid.
	ErrInvalidFileCID = errors.New("InvalidFileCID")
	// ErrIndexMismatch indicates an index mismatch error.
	ErrIndexMismatch = errors.New("IndexMismatch")
	// ErrNoPolicy indicates no storage policy is set.
	ErrNoPolicy = errors.New("NoPolicy")
	// ErrFileNotFilled indicates the file is not fully filled.
	ErrFileNotFilled = errors.New("FileNotFilled")
	// ErrBlockAlreadyFilled indicates the block is already filled.
	ErrBlockAlreadyFilled = errors.New("BlockAlreadyFilled")
	// ErrChunkCIDMismatch indicates a chunk CID mismatch error.
	ErrChunkCIDMismatch = errors.New("ChunkCIDMismatch")
	// ErrNotBucketOwner indicates the caller is not the bucket owner.
	ErrNotBucketOwner = errors.New("NotBucketOwner")
	// ErrBucketNotFound indicates the bucket was not found.
	ErrBucketNotFound = errors.New("BucketNotFound")
	// ErrFileDoesNotExist indicates the file does not exist.
	ErrFileDoesNotExist = errors.New("FileDoesNotExist")
	// ErrNotThePolicyOwner indicates the caller is not the policy owner.
	ErrNotThePolicyOwner = errors.New("NotThePolicyOwner")
	// ErrCloneArgumentsTooLong indicates clone arguments are too long.
	ErrCloneArgumentsTooLong = errors.New("CloneArgumentsTooLong")
	// ErrCreate2EmptyBytecode indicates empty bytecode in create2 operation.
	ErrCreate2EmptyBytecode = errors.New("Create2EmptyBytecode")
	// ErrECDSAInvalidSignatureS indicates an invalid ECDSA signature S value.
	ErrECDSAInvalidSignatureS = errors.New("ECDSAInvalidSignatureS")
	// ErrECDSAInvalidSignatureLength indicates an invalid ECDSA signature length.
	ErrECDSAInvalidSignatureLength = errors.New("ECDSAInvalidSignatureLength")
	// ErrECDSAInvalidSignature indicates an invalid ECDSA signature.
	ErrECDSAInvalidSignature = errors.New("ECDSAInvalidSignature")
	// ErrAlreadyWhitelisted indicates the address is already whitelisted.
	ErrAlreadyWhitelisted = errors.New("AlreadyWhitelisted")
	// ErrInvalidAddress indicates an invalid address.
	ErrInvalidAddress = errors.New("InvalidAddress")
	// ErrNotWhitelisted indicates the address is not whitelisted.
	ErrNotWhitelisted = errors.New("NotWhitelisted")
	// ErrMathOverflowedMulDiv indicates a math operation overflow in multiplication or division.
	ErrMathOverflowedMulDiv = errors.New("MathOverflowedMulDiv")
	// ErrInvalidBlocksAmount indicates an invalid amount of blocks.
	ErrInvalidBlocksAmount = errors.New("InvalidBlocksAmount")
	// ErrInvalidBlockIndex indicates an invalid block index.
	ErrInvalidBlockIndex = errors.New("InvalidBlockIndex")
	// ErrLastChunkDuplicate indicates a duplicate last chunk.
	ErrLastChunkDuplicate = errors.New("LastChunkDuplicate")
	// ErrFileNotExists indicates the file does not exist.
	ErrFileNotExists = errors.New("FileNotExists")
	// ErrNotSignedByBucketOwner indicates the transaction is not signed by the bucket owner.
	ErrNotSignedByBucketOwner = errors.New("NotSignedByBucketOwner")
	// ErrNonceAlreadyUsed indicates a nonce has already been used.
	ErrNonceAlreadyUsed = errors.New("NonceAlreadyUsed")
	// ErrOffsetOutOfBounds indicates an offset is out of bounds.
	ErrOffsetOutOfBounds = errors.New("OffsetOutOfBounds")
)

// ErrorHashToError maps error hashes to human-readable errors.
func ErrorHashToError(err error) error {
	var x rpc.DataError

	if errors.As(err, &x) {
		data := x.ErrorData()
		if hashCode, ok := data.(string); ok {
			switch hashCode {
			case "0x497ef2c2":
				return ErrBucketAlreadyExists
			case "0x4f4b202a":
				return ErrBucketInvalid
			case "0xdc64d0ad":
				return ErrBucketInvalidOwner
			case "0x938a92b7":
				return ErrBucketNonexists
			case "0x89fddc00":
				return ErrBucketNonempty
			case "0x6891dde0":
				return ErrFileAlreadyExists
			case "0x77a3cbd8":
				return ErrFileInvalid
			case "0x21584586":
				return ErrFileNonexists
			case "0xc4a3b6f1":
				return ErrFileNonempty
			case "0xd09ec7af":
				return ErrFileNameDuplicate
			case "0xd96b03b1":
				return ErrFileFullyUploaded
			case "0x702cf740":
				return ErrFileChunkDuplicate
			case "0xc1edd16a":
				return ErrBlockAlreadyExists
			case "0xcb20e88c":
				return ErrBlockInvalid
			case "0x15123121":
				return ErrBlockNonexists
			case "0x856b300d":
				return ErrInvalidArrayLength
			case "0x17ec8370":
				return ErrInvalidFileBlocksCount
			case "0x5660ebd2":
				return ErrInvalidLastBlockSize
			case "0x1b6fdfeb":
				return ErrInvalidEncodedSize
			case "0xfe33db92":
				return ErrInvalidFileCID
			case "0x37c7f255":
				return ErrIndexMismatch
			case "0xcefa6b05":
				return ErrNoPolicy
			case "0x5c371e92":
				return ErrFileNotFilled
			case "0xdad01942":
				return ErrBlockAlreadyFilled
			case "0x4b6b8ec8":
				return ErrChunkCIDMismatch
			case "0x0d6b18f0":
				return ErrNotBucketOwner
			case "0xc4c1a0c5":
				return ErrBucketNotFound
			case "0x3bcbb0de":
				return ErrFileDoesNotExist
			case "0xa2c09fea":
				return ErrNotThePolicyOwner
			case "0x94289054":
				return ErrCloneArgumentsTooLong
			case "0x4ca249dc":
				return ErrCreate2EmptyBytecode
			case "0xf3714a9b":
				return ErrECDSAInvalidSignatureS
			case "0x367e2e27":
				return ErrECDSAInvalidSignatureLength
			case "0xf645eedf":
				return ErrECDSAInvalidSignature
			case "0xb73e95e1":
				return ErrAlreadyWhitelisted
			case "0xe6c4247b":
				return ErrInvalidAddress
			case "0x584a7938":
				return ErrNotWhitelisted
			case "0x227bc153":
				return ErrMathOverflowedMulDiv
			case "0xe7b199a6":
				return ErrInvalidBlocksAmount
			case "0x59b452ef":
				return ErrInvalidBlockIndex
			case "0x55cbc831":
				return ErrLastChunkDuplicate
			case "0x2abde339":
				return ErrFileNotExists
			case "0x48e0ed68":
				return ErrNotSignedByBucketOwner
			case "0x923b8cbb":
				return ErrNonceAlreadyUsed
			case "0x9605a010":
				return ErrOffsetOutOfBounds
			default:
				return err
			}
		}
	}

	return err
}

// IgnoreOffsetError checks if the error is related to an offset out of bounds and returns nil if so.
func IgnoreOffsetError(err error) error {
	err = ErrorHashToError(err)
	if err != nil && errors.Is(err, ErrOffsetOutOfBounds) {
		return nil
	}

	return err
}
