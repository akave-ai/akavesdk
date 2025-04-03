// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package ipc

import (
	"encoding/hex"
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
)

// ErrorHashToError maps error hashes to human-readable errors.
func ErrorHashToError(err error) error {
	var x rpc.DataError

	if errors.As(err, &x) {
		data := x.ErrorData()
		if hashCode, ok := data.(string); ok {
			switch hashCode {
			case "0x497ef2c2":
				return errors.New("BucketAlreadyExists")
			case "0x4f4b202a":
				return errors.New("BucketInvalid")
			case "0xdc64d0ad":
				return errors.New("BucketInvalidOwner")
			case "0x938a92b7":
				return errors.New("BucketNonexists")
			case "0x89fddc00":
				return errors.New("BucketNonempty")
			case "0x6891dde0":
				return errors.New("FileAlreadyExists")
			case "0x77a3cbd8":
				return errors.New("FileInvalid")
			case "0x21584586":
				return errors.New("FileNonexists")
			case "0xc4a3b6f1":
				return errors.New("FileNonempty")
			case "0xd09ec7af":
				return errors.New("FileNameDuplicate")
			case "0xd96b03b1":
				return errors.New("FileFullyUploaded")
			case "0x702cf740":
				return errors.New("FileChunkDuplicate")
			case "0xc1edd16a":
				return errors.New("BlockAlreadyExists")
			case "0xcb20e88c":
				return errors.New("BlockInvalid")
			case "0x15123121":
				return errors.New("BlockNonexists")
			case "0x856b300d":
				return errors.New("InvalidArrayLength")
			case "0x17ec8370":
				return errors.New("InvalidFileBlocksCount")
			case "0x5660ebd2":
				return errors.New("InvalidLastBlockSize")
			case "0x1b6fdfeb":
				return errors.New("InvalidEncodedSize")
			case "0xfe33db92":
				return errors.New("InvalidFileCID")
			case "0x37c7f255":
				return errors.New("IndexMismatch")
			case "0xcefa6b05":
				return errors.New("NoPolicy")
			case "0x5c371e92":
				return errors.New("FileNotFilled")
			case "0xdad01942":
				return errors.New("BlockAlreadyFilled")
			case "0x4b6b8ec8":
				return errors.New("ChunkCIDMismatch")
			case "0x0d6b18f0":
				return errors.New("NotBucketOwner")
			case "0xc4c1a0c5":
				return errors.New("BucketNotFound")
			case "0x3bcbb0de":
				return errors.New("FileDoesNotExist")
			case "0xa2c09fea":
				return errors.New("NotThePolicyOwner")
			case "0x94289054":
				return errors.New("CloneArgumentsTooLong")
			case "0x4ca249dc":
				return errors.New("Create2EmptyBytecode")
			case "0xf3714a9b":
				return errors.New("ECDSAInvalidSignatureS")
			case "0x367e2e27":
				return errors.New("ECDSAInvalidSignatureLength")
			case "0xf645eedf":
				return errors.New("ECDSAInvalidSignature")
			case "0xb73e95e1":
				return errors.New("AlreadyWhitelisted")
			case "0xe6c4247b":
				return errors.New("InvalidAddress")
			case "0x584a7938":
				return errors.New("NotWhitelisted")
			case "0x227bc153":
				return errors.New("MathOverflowedMulDiv")
			case "0xe7b199a6":
				return errors.New("InvalidBlocksAmount")
			case "0x59b452ef":
				return errors.New("InvalidBlockIndex")
			case "0x55cbc831":
				return errors.New("LastChunkDuplicate")
			case "0x2abde339":
				return errors.New("FileNotExists")
			case "0x48e0ed68":
				return errors.New("NotSignedByBucketOwner")
			case "0x923b8cbb":
				return errors.New("NonceAlreadyUsed")
			default:
				return err
			}
		}
	}

	return err
}

func parseErrorsToHashes() []string {
	errorsContract := []string{"BucketAlreadyExists()", "BucketInvalid()", "BucketInvalidOwner()", "BucketNonexists()", "BucketNonempty()",
		"FileAlreadyExists()", "FileInvalid()", "FileNonexists()", "FileNonempty()", "FileNameDuplicate()", "FileFullyUploaded()", "FileChunkDuplicate()",
		"BlockAlreadyExists()", "BlockInvalid()", "BlockNonexists()", "InvalidArrayLength(uint256 cidsLength, uint256 sizesLength)", "InvalidFileBlocksCount()",
		"InvalidLastBlockSize()", "InvalidEncodedSize()", "InvalidFileCID()", "IndexMismatch()", "NoPolicy()", "FileDoesNotExist()", "BucketNotFound()", "NotBucketOwner()",
		"ChunkCIDMismatch(bytes fileCID)", "FileNotFilled()", "BlockAlreadyFilled()", "MathOverflowedMulDiv()", "NotWhitelisted()", "InvalidAddress()", "AlreadyWhitelisted()",
		"ECDSAInvalidSignature()", "ECDSAInvalidSignatureLength(uint256 length)", "error ECDSAInvalidSignatureS(bytes32 s)", "Create2EmptyBytecode()", "CloneArgumentsTooLong()",
		"NotThePolicyOwner()", "InvalidBlocksAmount()", "InvalidBlockIndex()", "LastChunkDuplicate()", "FileNotExists()", "Invalid signature: Not signed by bucket owner",
		"Nonce already used"}

	errHashes := make([]string, 0)

	for _, errC := range errorsContract {
		hash := crypto.Keccak256([]byte(errC))
		errMsg := "0x" + hex.EncodeToString(hash[:4])
		errHashes = append(errHashes, errMsg)
	}

	return errHashes
}
