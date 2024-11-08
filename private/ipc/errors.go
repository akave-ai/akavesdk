// Copyright (C) 2024 Akave
// See LICENSE for copying information.

package ipc

import (
	"errors"

	"github.com/ethereum/go-ethereum/rpc"
)

// ErrorHashToString maps error hashes to human-readable strings.
func ErrorHashToString(err error) string {
	var x rpc.DataError

	if errors.As(err, &x) {
		data := x.ErrorData()
		if hashCode, ok := data.(string); ok {
			switch hashCode {
			case "0x497ef2c2":
				return "BucketAlreadyExists"
			case "0x4f4b202a":
				return "BucketInvalid"
			case "0xdc64d0ad":
				return "BucketInvalidOwner"
			case "0x938a92b7":
				return "BucketNonexists"
			case "0x89fddc00":
				return "BucketNonempty"
			case "0x6891dde0":
				return "FileAlreadyExists"
			case "0x77a3cbd8":
				return "FileInvalid"
			case "0x21584586":
				return "FileNonexists"
			case "0xc4a3b6f1":
				return "FileNonempty"
			case "0xc1edd16a":
				return "BlockAlreadyExists"
			case "0xcb20e88c":
				return "BlockInvalid"
			case "0x15123121":
				return "BlockNonexists"
			case "0x856b300d":
				return "InvalidArrayLength"
			case "0x17ec8370":
				return "InvalidFileBlocksCount"
			default:
				return err.Error()
			}
		}
	}

	return err.Error()
}
