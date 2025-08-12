// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IStorageBucket is an auto generated low-level Go binding around an user-defined struct.
type IStorageBucket struct {
	Id        [32]byte
	Name      string
	CreatedAt *big.Int
	Owner     common.Address
	Files     [][32]byte
}

// IStorageChunk is an auto generated low-level Go binding around an user-defined struct.
type IStorageChunk struct {
	ChunkCIDs [][]byte
	ChunkSize []*big.Int
}

// IStorageFile is an auto generated low-level Go binding around an user-defined struct.
type IStorageFile struct {
	Id          [32]byte
	FileCID     []byte
	BucketId    [32]byte
	Name        string
	EncodedSize *big.Int
	CreatedAt   *big.Int
	ActualSize  *big.Int
	Chunks      IStorageChunk
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockAlreadyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockNonexists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketInvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketNonempty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketNonexists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"}],\"name\":\"ChunkCIDMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileChunkDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileFullyUploaded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNameDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNonempty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNotFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cidsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sizesLength\",\"type\":\"uint256\"}],\"name\":\"InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlocksAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEncodedSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFileBlocksCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFileCID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLastBlockSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LastChunkDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OffsetOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"AddFileBlocks\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AddFileChunk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"}],\"name\":\"AddPeerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CommitFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CreateBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CreateFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DeleteBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DeleteFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"}],\"name\":\"DeletePeerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"}],\"name\":\"FillChunkBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_BLOCKS_PER_FILE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BLOCK_SIZE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessManager\",\"outputs\":[{\"internalType\":\"contractIAccessManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"cid\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedChunkSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"cids\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkBlocksSizes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"addFileChunk\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isReplica\",\"type\":\"bool\"}],\"name\":\"addPeerBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedFileSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"}],\"name\":\"commitFile\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"createBucket\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"createFile\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deleteBucket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deleteFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deletePeerBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fileFillCounter\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fileRewardClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockCID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"fillChunkBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fulfilledBlocks\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fileOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileLimit\",\"type\":\"uint256\"}],\"name\":\"getBucketByName\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket\",\"name\":\"bucket\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getBucketIndexByName\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"}],\"name\":\"getBucketsByIds\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"bucketOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bucketLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileLimit\",\"type\":\"uint256\"}],\"name\":\"getBucketsByIdsWithFiles\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket[]\",\"name\":\"buckets\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getChunkByIndex\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getFileById\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"chunkCIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkSize\",\"type\":\"uint256[]\"}],\"internalType\":\"structIStorage.Chunk\",\"name\":\"chunks\",\"type\":\"tuple\"}],\"internalType\":\"structIStorage.File\",\"name\":\"file\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"getFileByName\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"chunkCIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkSize\",\"type\":\"uint256[]\"}],\"internalType\":\"structIStorage.Chunk\",\"name\":\"chunks\",\"type\":\"tuple\"}],\"internalType\":\"structIStorage.File\",\"name\":\"file\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getFileIndexById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getFileOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getFullFileInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"chunkCIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkSize\",\"type\":\"uint256[]\"}],\"internalType\":\"structIStorage.Chunk\",\"name\":\"chunks\",\"type\":\"tuple\"}],\"internalType\":\"structIStorage.File\",\"name\":\"file\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileLimit\",\"type\":\"uint256\"}],\"name\":\"getOwnerBuckets\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket[]\",\"name\":\"buckets\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"getPeerBlockIndexById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"cids\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"getPeersArrayByPeerBlockCid\",\"outputs\":[{\"internalType\":\"bytes32[][]\",\"name\":\"peers\",\"type\":\"bytes32[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"getPeersByPeerBlockCid\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"peers\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"isBlockFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"isChunkFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"isFileFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"isFileFilledV2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"}],\"name\":\"isPeerBlockReplica\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accessManagerAddress\",\"type\":\"address\"}],\"name\":\"setAccessManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIAkaveToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516151ab6100f95f395f8181612e5501528181612e7e01526132f601526151ab5ff3fe60806040526004361061025f575f3560e01c80637912bf681161013f578063d6d3110b116100b3578063f8a3e41a11610078578063f8a3e41a14610859578063faec05421461089b578063fc0c546a146108ba578063fd1d3c0c146108d8578063fd21c284146108f7578063fdcb606814610916575f5ffd5b8063d6d3110b146107b1578063e3f787e8146107d0578063e4ba8a58146107ef578063f83533d21461080e578063f855169a1461083a575f5ffd5b8063ad3cb1cc11610104578063ad3cb1cc146106ba578063b80777ea146106ea578063c4d66de8146106fc578063c95808041461071b578063d4967ce614610757578063d606205d14610785575f5ffd5b80637912bf681461062157806384b0196e1461064057806387c1ac07146106675780639a2e82b3146106865780639ccd4646146106a5575f5ffd5b80634f55ba82116101d65780635ecdfb531161019b5780635ecdfb531461052957806362838c231461055557806368e6408f146105745780636a6e658b146105b55780636af0f801146105d45780636ce02363146105f3575f5ffd5b80634f55ba821461045d57806352d1902d1461047c57806354fd4d5014610490578063564b81ef146104c65780635a4e9564146104d8575f5ffd5b806335bdb7111161022757806335bdb711146103925780633f383980146103b157806348b49875146103d057806349cdf6ac146103fc5780634d7dc6141461041b5780634f1ef28614610448575f5ffd5b8063018c1e9c146102635780631b475ef4146102a6578063262a61ce14610305578063287e677f1461033157806330b91d0714610365575b5f5ffd5b34801561026e575f5ffd5b5061029161027d3660046141e3565b60086020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b3480156102b1575f5ffd5b506102ed6102c03660046141e3565b5f90815260046020908152604080832060029081015484529091529020600301546001600160a01b031690565b6040516001600160a01b03909116815260200161029d565b348015610310575f5ffd5b5061032461031f366004614215565b610935565b60405161029d919061430e565b34801561033c575f5ffd5b5061035061034b366004614414565b6109b5565b6040805192835290151560208301520161029d565b348015610370575f5ffd5b5061038461037f36600461449e565b610a47565b60405190815260200161029d565b34801561039d575f5ffd5b506103246103ac366004614577565b610d14565b3480156103bc575f5ffd5b506102916103cb3660046145b5565b610f83565b3480156103db575f5ffd5b506103ef6103ea3660046145d5565b611040565b60405161029d9190614618565b348015610407575f5ffd5b5061035061041636600461465a565b6110be565b348015610426575f5ffd5b5061043a6104353660046145b5565b611149565b60405161029d9291906146ac565b61045b6104563660046146cd565b611241565b005b348015610468575f5ffd5b5061045b610477366004614711565b611257565b348015610487575f5ffd5b50610384611273565b34801561049b575f5ffd5b506040805180820190915260058152640312e302e360dc1b60208201525b60405161029d91906147be565b3480156104d1575f5ffd5b5046610384565b3480156104e3575f5ffd5b506105146104f23660046145b5565b600560209081525f928352604080842090915290825290205463ffffffff1681565b60405163ffffffff909116815260200161029d565b348015610534575f5ffd5b506105486105433660046145d5565b61128e565b60405161029d91906148c1565b348015610560575f5ffd5b5061032461056f3660046148d3565b61156d565b34801561057f575f5ffd5b5061029161058e3660046141e3565b5f9081526006602090815260408083205460049092529091206007015461ffff9091161490565b3480156105c0575f5ffd5b506102916105cf36600461492f565b6115b2565b3480156105df575f5ffd5b506103846105ee3660046145d5565b611760565b3480156105fe575f5ffd5b50610609620f424081565b6040516001600160401b03909116815260200161029d565b34801561062c575f5ffd5b5061029161063b3660046145b5565b6117e7565b34801561064b575f5ffd5b50610654611834565b60405161029d979695949392919061498b565b348015610672575f5ffd5b50610384610681366004614a07565b6118e2565b348015610691575f5ffd5b506103846106a0366004614a65565b6119e4565b3480156106b0575f5ffd5b5061060961040081565b3480156106c5575f5ffd5b506104b9604051806040016040528060058152602001640352e302e360dc1b81525081565b3480156106f5575f5ffd5b5042610384565b348015610707575f5ffd5b5061045b610716366004614ae3565b611d41565b348015610726575f5ffd5b5061045b610735366004614ae3565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b348015610762575f5ffd5b50610776610771366004614afc565b611e93565b60405161029d93929190614b78565b348015610790575f5ffd5b506107a461079f366004614ba1565b611ef1565b60405161029d9190614c08565b3480156107bc575f5ffd5b506102916107cb366004614c5f565b612004565b3480156107db575f5ffd5b506103846107ea366004614cb2565b612275565b3480156107fa575f5ffd5b50610291610809366004614ce3565b6123db565b348015610819575f5ffd5b5061082d610828366004614d16565b612433565b60405161029d9190614d6e565b348015610845575f5ffd5b50610350610854366004614d80565b6124de565b348015610864575f5ffd5b506108886108733660046141e3565b60066020525f908152604090205461ffff1681565b60405161ffff909116815260200161029d565b3480156108a6575f5ffd5b506105486108b53660046141e3565b612558565b3480156108c5575f5ffd5b505f546102ed906001600160a01b031681565b3480156108e3575f5ffd5b506102916108f2366004614db5565b612811565b348015610902575f5ffd5b506102916109113660046141e3565b612a4e565b348015610921575f5ffd5b506001546102ed906001600160a01b031681565b6001600160a01b0385165f90815260036020908152604091829020805483518184028101840190945280845260609391926109aa9291849183018282801561099a57602002820191905f5260205f20905b815481526020019060010190808311610986575b5050505050878787876001612a94565b979650505050505050565b5f5f5f84846040516020016109cb929190614e18565b60408051601f1981840301815291815281516020928301206001600160a01b0387165f9081526003909352908220909250905b8154811015610a3c5782828281548110610a1a57610a1a614e44565b905f5260205f20015403610a345780945060019350610a3c565b6001016109fe565b5050505b9250929050565b5f8881526002602052604081205489908203610a765760405163938a92b760e01b815260040160405180910390fd5b5f818152600260205260409020600301546001600160a01b03163314610abc57610a9f81612db2565b610abc5760405163dc64d0ad60e01b815260040160405180910390fd5b89895f8282604051602001610ad2929190614e58565b60405160208183030381529060405280519060200120905060045f8281526020019081526020015f20600401545f14610b1e5760405163d96b03b160e01b815260040160405180910390fd5b5f8d8d604051602001610b32929190614e58565b60408051601f1981840301815291815281516020928301205f8181526004909352908220549092509003610b7957604051632abde33960e01b815260040160405180910390fd5b5f818152600460205260409020600701548714610ba9576040516301c0b3dd60e61b815260040160405180910390fd5b8988141580610bb8575060208a115b15610bd6576040516373d8ccd360e11b815260040160405180910390fd5b60208a1015610c2c575f8181526007602052604090205461ffff1615610c0f576040516355cbc83160e01b815260040160405180910390fd5b5f818152600760205260409020805461ffff191661ffff8c161790555b8b5f03610c4c57604051631b6fdfeb60e01b815260040160405180910390fd5b8c604051610c5a9190614e69565b604051908190038120338252908f9083907f0d877d3656ad561e6c8224e31cc98f1b3e62b1d828fd396c47ef0daccb6838e59060200160405180910390a460045f8281526020019081526020015f206007015f018f908060018154018082558091505060019003905f5260205f20015f909190919091509081610cdd9190614ef7565b505f818152600460209081526040822060080180546001810182559083529120018c90559450505050509998505050505050505050565b60605f826001600160401b03811115610d2f57610d2f614371565b604051908082528060200260200182016040528015610d6857816020015b610d5561401e565b815260200190600190039081610d4d5790505b5090505f5b83811015610f79575f60025f878785818110610d8b57610d8b614e44565b9050602002013581526020019081526020015f206040518060a00160405290815f8201548152602001600182018054610dc390614e74565b80601f0160208091040260200160405190810160405280929190818152602001828054610def90614e74565b8015610e3a5780601f10610e1157610100808354040283529160200191610e3a565b820191905f5260205f20905b815481529060010190602001808311610e1d57829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b0316815260200160048201805480602002602001604051908101604052809291908181526020018280548015610ec857602002820191905f5260205f20905b815481526020019060010190808311610eb4575b50505050508152505090506040518060a00160405280825f01518152602001826020015181526020018260400151815260200182606001516001600160a01b031681526020015f6001600160401b03811115610f2657610f26614371565b604051908082528060200260200182016040528015610f4f578160200160208202803683370190505b50815250838381518110610f6557610f65614e44565b602090810291909101015250600101610d6d565b5090505b92915050565b5f828152600460205260408120600701548190610fa290600190614fc5565b90508083148015610fc357505f8481526007602052604090205461ffff1615155b15611019575f84815260076020526040812054610fe99060019061ffff1681901b614fc5565b5f868152600560209081526040808320888452909152902054811663ffffffff9081169116149250610f7d915050565b50505f91825260056020908152604080842092845291905290205463ffffffff9081161490565b6060600a5f8481526020019081526020015f20826040516110619190614e69565b90815260408051918290036020908101832080548083028501830190935282845291908301828280156110b157602002820191905f5260205f20905b81548152602001906001019080831161109d575b5050505050905092915050565b5f5f5f85846040516020016110d4929190614e18565b60408051601f1981840301815291815281516020928301205f8181526002909352908220909250600401905b815481101561113e578682828154811061111c5761111c614e44565b905f5260205f20015403611136578094506001935061113e565b600101611100565b505050935093915050565b5f82815260046020526040812060070180546060929182918590811061117157611171614e44565b905f5260205f2001805461118490614e74565b80601f01602080910402602001604051908101604052809291908181526020018280546111b090614e74565b80156111fb5780601f106111d2576101008083540402835291602001916111fb565b820191905f5260205f20905b8154815290600101906020018083116111de57829003601f168201915b5050505f88815260046020526040812060080180549495509093909250879150811061122957611229614e44565b5f918252602090912001549196919550909350505050565b611249612e4a565b6112538282612ef3565b5050565b611268898989898989898989612fb4565b505050505050505050565b5f61127c6132eb565b505f5160206151565f395f51905f5290565b611296614053565b5f83836040516020016112aa929190614e58565b60408051601f1981840301815282825280516020918201205f8181526004835283902061010085019093528254845260018301805491955091840191906112f090614e74565b80601f016020809104026020016040519081016040528092919081815260200182805461131c90614e74565b80156113675780601f1061133e57610100808354040283529160200191611367565b820191905f5260205f20905b81548152906001019060200180831161134a57829003601f168201915b505050505081526020016002820154815260200160038201805461138a90614e74565b80601f01602080910402602001604051908101604052809291908181526020018280546113b690614e74565b80156114015780601f106113d857610100808354040283529160200191611401565b820191905f5260205f20905b8154815290600101906020018083116113e457829003601f168201915b50505050508152602001600482015481526020016005820154815260200160068201548152602001600782016040518060400160405290815f8201805480602002602001604051908101604052809291908181526020015f905b82821015611503578382905f5260205f2001805461147890614e74565b80601f01602080910402602001604051908101604052809291908181526020018280546114a490614e74565b80156114ef5780601f106114c6576101008083540402835291602001916114ef565b820191905f5260205f20905b8154815290600101906020018083116114d257829003601f168201915b50505050508152602001906001019061145b565b5050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561155857602002820191905f5260205f20905b815481526020019060010190808311611544575b50505091909252505050905250949350505050565b60606109aa8787808060200260200160405190810160405280939291908181526020018383602002808284375f92019190915250899250889150879050866001612a94565b5f8581526009602052604081205460ff166115e057604051631512312160e01b815260040160405180910390fd5b8484846040516020016115f593929190614fd8565b60405160208183030381529060405280519060200120861461162a576040516332c83a2360e21b815260040160405180910390fd5b5f85858560405160200161164093929190614fd8565b60405160208183030381529060405280519060200120905086811461167857604051637ae9080f60e11b815260040160405180910390fd5b5f878152600960209081526040808320805460ff19169055878352600a90915280822090516116a8908790614e69565b9081526020016040518091039020905080600182805490506116ca9190614fc5565b815481106116da576116da614e44565b905f5260205f2001548185815481106116f5576116f5614e44565b905f5260205f2001819055508080548061171157611711614fef565b600190038181905f5260205f20015f9055905586887f37505c6d2cdf9e778d6110c5ad2e49c649d521a18d2da6f007d3364bd83a45ae60405160405180910390a3506001979650505050505050565b5f828152600260205260408120548390820361178f5760405163938a92b760e01b815260040160405180910390fd5b5f818152600260205260409020600301546001600160a01b031633146117d5576117b881612db2565b6117d55760405163dc64d0ad60e01b815260040160405180910390fd5b6117df8484613334565b949350505050565b5f5f8284604051602001611805929190918252602082015260400190565b60408051808303601f1901815291815281516020928301205f908152600b90925290205460ff16949350505050565b5f60608082808083815f5160206150a75f395f51905f52805490915015801561185f57506001810154155b6118a85760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064015b60405180910390fd5b6118b0613551565b6118b8613611565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b5f8484846040516020016118f893929190614fd8565b60408051601f19818403018152828252805160209182012090830188905290820186905291505f9060600160408051601f1981840301815291815281516020928301205f85815260099093529120805460ff1916600117905590508215611972575f818152600b60205260409020805460ff191660011790555b5f858152600a602052604090819020905161198e908690614e69565b908152604051908190036020908101822080546001810182555f918252918120909101889055879184917f9f91e060f6a60bc2a5e695f07d854d6a659eab9f5e1441ac24f6e62614bb8b5b91a350949350505050565b5f8581526002602052604081205486908203611a135760405163938a92b760e01b815260040160405180910390fd5b5f818152600260205260409020600301546001600160a01b03163314611a5957611a3c81612db2565b611a595760405163dc64d0ad60e01b815260040160405180910390fd5b86865f8282604051602001611a6f929190614e58565b60405160208183030381529060405280519060200120905060045f8281526020019081526020015f20600401545f14611abb5760405163d96b03b160e01b815260040160405180910390fd5b5f8a8a604051602001611acf929190614e58565b60408051601f1981840301815291815281516020928301205f81815260068452828120546004909452919091206007015490925061ffff90911614611b2757604051632e1b8f4960e11b815260040160405180910390fd5b885f03611b4757604051631b6fdfeb60e01b815260040160405180910390fd5b86515f03611b6857604051637f19edc960e11b815260040160405180910390fd5b5f805b5f83815260046020526040902060080154811015611bc3575f838152600460205260409020600801805482908110611ba557611ba5614e44565b905f5260205f20015482611bb99190615003565b9150600101611b6b565b50898114611be457604051631b6fdfeb60e01b815260040160405180910390fd5b5f828152600460205260409020600101611bfe8982614ef7565b505f8281526004602081815260408084209283018e905560069092018c90556008905290205460ff16611ce6575f828152600860209081526040808320805460ff19166001179055600490915281206007015490611c5d82600a615016565b611c6f90670de0b6b3a7640000615016565b90505f82118015611c7f57505f81115b15611ce3575f546040516340c10f1960e01b8152336004820152602481018390526001600160a01b03909116906340c10f19906044015f604051808303815f87803b158015611ccc575f5ffd5b505af1158015611cde573d5f5f3e3d5ffd5b505050505b50505b8a604051611cf49190614e69565b604051908190038120338252908d9084907fc208b556b9717832827b84e296c7f4bb33292ba520fed2eae2f7c0d61f1c0a619060200160405180910390a4509a9950505050505050505050565b5f611d4a61364f565b805490915060ff600160401b82041615906001600160401b03165f81158015611d705750825b90505f826001600160401b03166001148015611d8b5750303b155b905081158015611d99575080155b15611db75760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315611de157845460ff60401b1916600160401b1785555b611e236040518060400160405280600781526020016653746f7261676560c81b815250604051806040016040528060018152602001603160f81b815250613677565b611e2b613689565b5f80546001600160a01b0319166001600160a01b0388161790558315611e8b57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b611e9b614053565b5f5f5f8587604051602001611eb1929190614e58565b604051602081830303815290604052805190602001209050611ed3868861128e565b9350611ee08882876110be565b9093509150505b9450945094915050565b6060826001600160401b03811115611f0b57611f0b614371565b604051908082528060200260200182016040528015611f3e57816020015b6060815260200190600190039081611f295790505b5090505f5b83811015611ffc57600a5f868684818110611f6057611f60614e44565b9050602002013581526020019081526020015f2083604051611f829190614e69565b9081526040805191829003602090810183208054808302850183019093528284529190830182828015611fd257602002820191905f5260205f20905b815481526020019060010190808311611fbe575b5050505050828281518110611fe957611fe9614e44565b6020908102919091010152600101611f43565b509392505050565b5f83815260026020526040812054849082036120335760405163938a92b760e01b815260040160405180910390fd5b5f818152600260205260409020600301546001600160a01b031633146120795761205c81612db2565b6120795760405163dc64d0ad60e01b815260040160405180910390fd5b5f8681526004602052604081205490036120a657604051632abde33960e01b815260040160405180910390fd5b84846040516020016120b9929190614e58565b6040516020818303038152906040528051906020012086146120ee57604051630ef4797b60e31b815260040160405180910390fd5b5f8681526004602052604081208181559061210c60018301826140ad565b600282015f9055600382015f61212291906140ad565b5f600483018190556005830181905560068301819055600783019061214782826140e4565b612154600183015f6140ff565b5050505f86815260026020526040902060040180549091508410158061219457508681858154811061218857612188614e44565b905f5260205f20015414155b156121b2576040516337c7f25560e01b815260040160405180910390fd5b805481906121c290600190614fc5565b815481106121d2576121d2614e44565b905f5260205f2001548185815481106121ed576121ed614e44565b905f5260205f2001819055508080548061220957612209614fef565b600190038181905f5260205f20015f905590558460405161222a9190614e69565b60405190819003812033825290879089907f0e1bf50f5cca6659c62146db5b60160121a3752011b621d8c8953a1e0e23d0489060200160405180910390a45060019695505050505050565b5f8133604051602001612289929190614e18565b60408051601f1981840301815291815281516020928301205f8181526002909352912054909150156122ce576040516324bf796160e11b815260040160405180910390fd5b6040805160a0810182528281526020808201858152428385015233606084015283515f80825281840186526080850191909152858152600290925292902081518155915190919060018201906123249082614ef7565b506040820151600282015560608201516003820180546001600160a01b0319166001600160a01b039092169190911790556080820151805161237091600484019160209091019061411a565b5050335f81815260036020908152604080832080546001810182559084529190922001849055519091506123a5908490614e69565b6040519081900381209083907fb8be57bce74a717a1bbd4acf428df655720fce75c1854b02a88484388df241a8905f90a4919050565b5f60208360ff161115612401576040516359b452ef60e01b815260040160405180910390fd5b505f838152600560209081526040808320848452909152902054600160ff84161b1663ffffffff1615155b9392505050565b61243b61401e565b5f858560405160200161244f929190614e18565b60408051601f198184030181528282528051602091820120600180855284840190935293505f929190808301908036833701905050905081815f8151811061249957612499614e44565b6020026020010181815250505f6124b5825f5f89896001612a94565b9050805f815181106124c9576124c9614e44565b60200260200101519350505050949350505050565b5f828152600a60205260408082209051829182916124fd908690614e69565b90815260405190819003602001902090505f5b815481101561254e578682828154811061252c5761252c614e44565b905f5260205f20015403612546578093506001925061254e565b600101612510565b5050935093915050565b612560614053565b60045f8381526020019081526020015f20604051806101000160405290815f820154815260200160018201805461259690614e74565b80601f01602080910402602001604051908101604052809291908181526020018280546125c290614e74565b801561260d5780601f106125e45761010080835404028352916020019161260d565b820191905f5260205f20905b8154815290600101906020018083116125f057829003601f168201915b505050505081526020016002820154815260200160038201805461263090614e74565b80601f016020809104026020016040519081016040528092919081815260200182805461265c90614e74565b80156126a75780601f1061267e576101008083540402835291602001916126a7565b820191905f5260205f20905b81548152906001019060200180831161268a57829003601f168201915b50505050508152602001600482015481526020016005820154815260200160068201548152602001600782016040518060400160405290815f8201805480602002602001604051908101604052809291908181526020015f905b828210156127a9578382905f5260205f2001805461271e90614e74565b80601f016020809104026020016040519081016040528092919081815260200182805461274a90614e74565b80156127955780601f1061276c57610100808354040283529160200191612795565b820191905f5260205f20905b81548152906001019060200180831161277857829003601f168201915b505050505081526020019060010190612701565b505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156127fe57602002820191905f5260205f20905b8154815260200190600101908083116127ea575b5050509190925250505090525092915050565b5f83815260026020526040812054849082036128405760405163938a92b760e01b815260040160405180910390fd5b5f818152600260205260409020600301546001600160a01b031633146128865761286981612db2565b6128865760405163dc64d0ad60e01b815260040160405180910390fd5b8333604051602001612899929190614e18565b6040516020818303038152906040528051906020012085146128ce576040516327a5901560e11b815260040160405180910390fd5b5f85815260026020526040902060040154156128fc5760405162227f7760ea1b815260040160405180910390fd5b5f8581526002602052604081208181559061291a60018301826140ad565b5f600283018190556003830180546001600160a01b03191690556129429060048401906140ff565b5050335f9081526003602052604090208054869082908690811061296857612968614e44565b905f5260205f2001541461298f576040516337c7f25560e01b815260040160405180910390fd5b8054819061299f90600190614fc5565b815481106129af576129af614e44565b905f5260205f2001548185815481106129ca576129ca614e44565b905f5260205f200181905550808054806129e6576129e6614fef565b600190038181905f5260205f20015f90559055336001600160a01b031685604051612a119190614e69565b6040519081900381209088907feda2cc8e002ead8000b1e0c1debfc9a88bd7ee6e94b8dc0763db17849fcf0389905f90a450600195945050505050565b5f805b5f83815260046020526040902060070154811015612a8b57612a738382610f83565b15155f03612a8357505f92915050565b600101612a51565b50600192915050565b60605f86158015612aa3575085155b15612ab057506001612abf565b612abc87895188613691565b90505b806001600160401b03811115612ad757612ad7614371565b604051908082528060200260200182016040528015612b1057816020015b612afd61401e565b815260200190600190039081612af55790505b5091505f5b81811015612da6575f89612b29838b615003565b81518110612b3957612b39614e44565b602002602001015190505f60025f8381526020019081526020015f206040518060a00160405290815f8201548152602001600182018054612b7990614e74565b80601f0160208091040260200160405190810160405280929190818152602001828054612ba590614e74565b8015612bf05780601f10612bc757610100808354040283529160200191612bf0565b820191905f5260205f20905b815481529060010190602001808311612bd357829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b0316815260200160048201805480602002602001604051908101604052809291908181526020018280548015612c7e57602002820191905f5260205f20905b815481526020019060010190808311612c6a575b50505050508152505090508515612cfe576040518060a00160405280825f01518152602001826020015181526020018260400151815260200182606001516001600160a01b03168152602001612cd983608001518b8b6136e3565b815250858481518110612cee57612cee614e44565b6020026020010181905250612d9c565b6040518060a00160405280825f01518152602001826020015181526020018260400151815260200182606001516001600160a01b031681526020015f6001600160401b03811115612d5157612d51614371565b604051908082528060200260200182016040528015612d7a578160200160208202803683370190505b50815250858481518110612d9057612d90614e44565b60200260200101819052505b5050600101612b15565b50509695505050505050565b6001545f906001600160a01b031615612e435760015460405163e124bdd960e01b8152600481018490523360248201526001600160a01b039091169063e124bdd990604401602060405180830381865afa158015612e12573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612e36919061502d565b15612e4357506001919050565b505f919050565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480612ed057507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316612ec45f5160206151565f395f51905f52546001600160a01b031690565b6001600160a01b031614155b15612eee5760405163703e46dd60e11b815260040160405180910390fd5b565b50565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612f4d575060408051601f3d908101601f19168201909252612f4a91810190615048565b60015b612f7557604051634c9c8ce360e01b81526001600160a01b038316600482015260240161189f565b5f5160206151565f395f51905f528114612fa557604051632a87526960e21b81526004810182905260240161189f565b612faf8383613795565b505050565b5f8784604051602001612fc8929190614e58565b60408051601f1981840301815291815281516020928301205f818152600490935290822054909250900361300f57604051632abde33960e01b815260040160405180910390fd5b5f81815260046020526040812060080180548990811061303157613031614e44565b905f5260205f20015490508086600161304a919061505f565b60ff16111561306c576040516359b452ef60e01b815260040160405180910390fd5b60408051610100810182525f848152600460205291822060070180548291908c90811061309b5761309b614e44565b905f5260205f200180546130ae90614e74565b80601f01602080910402602001604051908101604052809291908181526020018280546130da90614e74565b80156131255780601f106130fc57610100808354040283529160200191613125565b820191905f5260205f20905b81548152906001019060200180831161310857829003601f168201915b505050505081526020018d81526020018a81526020018860ff1681526020018c81526020018981526020018581526020018b815250905061316681866137ea565b61317183888b6123db565b1561318f57604051636d680ca160e11b815260040160405180910390fd5b50505f898b866040516020016131a793929190614fd8565b60408051601f1981840301815291815281516020928301205f818152600990935291205490915060ff166131e3576131e18a8c875f6118e2565b505b6131ee82878a613a5c565b6131f88289610f83565b1561323a575f82815260066020526040812080546001929061321f90849061ffff16615078565b92506101000a81548161ffff021916908361ffff1602179055505b5f546040516340c10f1960e01b8152336004820152670de0b6b3a764000060248201526001600160a01b03909116906340c10f19906044015f604051808303815f87803b158015613289575f5ffd5b505af115801561329b573d5f5f3e3d5ffd5b5050604080518b815260ff8a1660208201528d93508e92507feded691eefce1ca615096347126f78807fee63b86aab874574e9c733082ce867910160405180910390a35050505050505050505050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612eee5760405163703e46dd60e11b815260040160405180910390fd5b5f5f8383604051602001613349929190614e58565b60408051601f1981840301815291815281516020928301205f81815260049093529120549091501561338e576040516303448eef60e51b815260040160405180910390fd5b5f84815260026020908152604080832060040180546001810182559084529190922001829055516133c0908490614e69565b60405190819003812033825290859083907fb018e47bdb983351e1bee22415a8f41eef5c2bf1c43c6d3d0992e678ae762df59060200160405180910390a4604080515f818301818152606083019093529181908361342e565b60608152602001906001900390816134195790505b5081526020015f60405190808252806020026020018201604052801561345e578160200160208202803683370190505b50905260408051610100810182528481528151602081810184525f8083528184019283528385018b9052606084018a9052608084018190524260a085015260c0840181905260e084018690528781526004909152929092208151815591519293509160018201906134cf9082614ef7565b5060408201516002820155606082015160038201906134ee9082614ef7565b506080820151600482015560a0820151600582015560c0820151600682015560e082015180518051600784019161352a91839160200190614163565b506020828101518051613543926001850192019061411a565b509498975050505050505050565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f5160206150a75f395f51905f529161358f90614e74565b80601f01602080910402602001604051908101604052809291908181526020018280546135bb90614e74565b80156136065780601f106135dd57610100808354040283529160200191613606565b820191905f5260205f20905b8154815290600101906020018083116135e957829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f5160206150a75f395f51905f529161358f90614e74565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610f7d565b61367f613ac5565b6112538282613aea565b612eee613ac5565b5f828411156136b3576040516309605a0160e41b815260040160405180910390fd5b5f6136be8386615003565b9050838111156136da576136d28585614fc5565b91505061242c565b8291505061242c565b60605f6136f284865185613691565b90505f816001600160401b0381111561370d5761370d614371565b604051908082528060200260200182016040528015613736578160200160208202803683370190505b5090505f5b8281101561378b578661374e8288615003565b8151811061375e5761375e614e44565b602002602001015182828151811061377857613778614e44565b602090810291909101015260010161373b565b5095945050505050565b61379e82613b49565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156137e257612faf8282613bac565b611253613c1e565b8160c001514211156138325760405162461bcd60e51b815260206004820152601160248201527014da59db985d1d5c9948195e1c1a5c9959607a1b604482015260640161189f565b60e08201515f90815260026020908152604080832060030154815160c08101909252608f8083526001600160a01b0390911693926150c79083013980519060200120845f01518051906020012085602001518660400151876060015188608001518960a001518a60c001518b60e001516040516020016138f79998979695949392919098895260208901979097526040880195909552606087019390935260ff91909116608086015260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012090505f61391982613c3d565b90505f6139268286613c69565b9050836001600160a01b0316816001600160a01b03161461399f5760405162461bcd60e51b815260206004820152602d60248201527f496e76616c6964207369676e61747572653a204e6f74207369676e656420627960448201526c10313ab1b5b2ba1037bbb732b960991b606482015260840161189f565b6001600160a01b0381165f908152600c6020908152604080832060e08a01518452825280832060a08a0151845290915290205460ff1615613a175760405162461bcd60e51b8152602060048201526012602482015271139bdb98d948185b1c9958591e481d5cd95960721b604482015260640161189f565b6001600160a01b03165f908152600c6020908152604080832060e08901518452825280832060a090980151835296905294909420805460ff1916600117905550505050565b60208260ff1610613a80576040516359b452ef60e01b815260040160405180910390fd5b5f928352600560209081526040808520928552919052909120805463ffffffff600160ff9094169390931b83169281169290921763ffffffff19909216919091179055565b613acd613c91565b612eee57604051631afcd79f60e31b815260040160405180910390fd5b613af2613ac5565b5f5160206150a75f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102613b2b8482614ef7565b5060038101613b3a8382614ef7565b505f8082556001909101555050565b806001600160a01b03163b5f03613b7e57604051634c9c8ce360e01b81526001600160a01b038216600482015260240161189f565b5f5160206151565f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051613bc89190614e69565b5f60405180830381855af49150503d805f8114613c00576040519150601f19603f3d011682016040523d82523d5f602084013e613c05565b606091505b5091509150613c15858383613caa565b95945050505050565b3415612eee5760405163b398979f60e01b815260040160405180910390fd5b5f610f7d613c49613d06565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f613c778686613d14565b925092509250613c878282613d5d565b5090949350505050565b5f613c9a61364f565b54600160401b900460ff16919050565b606082613cbf57613cba82613e15565b61242c565b8151158015613cd657506001600160a01b0384163b155b15613cff57604051639996b31560e01b81526001600160a01b038516600482015260240161189f565b508061242c565b5f613d0f613e3d565b905090565b5f5f5f8351604103613d4b576020840151604085015160608601515f1a613d3d88828585613eb0565b955095509550505050613d56565b505081515f91506002905b9250925092565b5f826003811115613d7057613d70615092565b03613d79575050565b6001826003811115613d8d57613d8d615092565b03613dab5760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115613dbf57613dbf615092565b03613de05760405163fce698f760e01b81526004810182905260240161189f565b6003826003811115613df457613df4615092565b03611253576040516335e2f38360e21b81526004810182905260240161189f565b805115613e2457805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f613e67613f74565b613e6f613fdc565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115613ee957505f91506003905082611ee7565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613f3a573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116613f6557505f925060019150829050611ee7565b975f9750879650945050505050565b5f5f5160206150a75f395f51905f5281613f8c613551565b805190915015613fa457805160209091012092915050565b81548015613fb3579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f5160206150a75f395f51905f5281613ff4613611565b80519091501561400c57805160209091012092915050565b60018201548015613fb3579392505050565b6040518060a001604052805f8152602001606081526020015f81526020015f6001600160a01b03168152602001606081525090565b6040518061010001604052805f8152602001606081526020015f8152602001606081526020015f81526020015f81526020015f81526020016140a8604051806040016040528060608152602001606081525090565b905290565b5080546140b990614e74565b5f825580601f106140c8575050565b601f0160209004905f5260205f2090810190612ef091906141b3565b5080545f8255905f5260205f2090810190612ef091906141c7565b5080545f8255905f5260205f2090810190612ef091906141b3565b828054828255905f5260205f20908101928215614153579160200282015b82811115614153578251825591602001919060010190614138565b5061415f9291506141b3565b5090565b828054828255905f5260205f209081019282156141a7579160200282015b828111156141a757825182906141979082614ef7565b5091602001919060010190614181565b5061415f9291506141c7565b5b8082111561415f575f81556001016141b4565b8082111561415f575f6141da82826140ad565b506001016141c7565b5f602082840312156141f3575f5ffd5b5035919050565b80356001600160a01b0381168114614210575f5ffd5b919050565b5f5f5f5f5f60a08688031215614229575f5ffd5b614232866141fa565b97602087013597506040870135966060810135965060800135945092505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f8151808452602084019350602083015f5b828110156142b1578151865260209586019590910190600101614293565b5093949350505050565b805182525f602082015160a060208501526142d960a0850182614253565b90506040830151604085015260018060a01b03606084015116606085015260808301518482036080860152613c158282614281565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561436557603f198786030184526143508583516142bb565b94506020938401939190910190600101614334565b50929695505050505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112614394575f5ffd5b8135602083015f5f6001600160401b038411156143b3576143b3614371565b50604051601f19601f85018116603f011681018181106001600160401b03821117156143e1576143e1614371565b6040528381529050808284018710156143f8575f5ffd5b838360208301375f602085830101528094505050505092915050565b5f5f60408385031215614425575f5ffd5b82356001600160401b0381111561443a575f5ffd5b61444685828601614385565b925050614455602084016141fa565b90509250929050565b5f5f83601f84011261446e575f5ffd5b5081356001600160401b03811115614484575f5ffd5b6020830191508360208260051b8501011115610a40575f5ffd5b5f5f5f5f5f5f5f5f5f60e08a8c0312156144b6575f5ffd5b89356001600160401b038111156144cb575f5ffd5b6144d78c828d01614385565b99505060208a0135975060408a01356001600160401b038111156144f9575f5ffd5b6145058c828d01614385565b97505060608a0135955060808a01356001600160401b03811115614527575f5ffd5b6145338c828d0161445e565b90965094505060a08a01356001600160401b03811115614551575f5ffd5b61455d8c828d0161445e565b9a9d999c50979a9699959894979660c00135949350505050565b5f5f60208385031215614588575f5ffd5b82356001600160401b0381111561459d575f5ffd5b6145a98582860161445e565b90969095509350505050565b5f5f604083850312156145c6575f5ffd5b50508035926020909101359150565b5f5f604083850312156145e6575f5ffd5b8235915060208301356001600160401b03811115614602575f5ffd5b61460e85828601614385565b9150509250929050565b602080825282518282018190525f918401906040840190835b8181101561464f578351835260209384019390920191600101614631565b509095945050505050565b5f5f5f6060848603121561466c575f5ffd5b83356001600160401b03811115614681575f5ffd5b61468d86828701614385565b935050602084013591506146a3604085016141fa565b90509250925092565b604081525f6146be6040830185614253565b90508260208301529392505050565b5f5f604083850312156146de575f5ffd5b6146e7836141fa565b915060208301356001600160401b03811115614602575f5ffd5b803560ff81168114614210575f5ffd5b5f5f5f5f5f5f5f5f5f6101208a8c03121561472a575f5ffd5b8935985060208a0135975060408a0135965060608a0135955060808a0135945061475660a08b01614701565b935060c08a01356001600160401b03811115614770575f5ffd5b61477c8c828d01614385565b93505060e08a01356001600160401b03811115614797575f5ffd5b6147a38c828d01614385565b999c989b509699959894979396509194610100013592915050565b602081525f61242c6020830184614253565b5f6040830182516040855281815180845260608701915060608160051b88010193506020830192505f5b8181101561482b57605f19888603018352614816858551614253565b945060209384019392909201916001016147fa565b5050505060208301518482036020860152613c158282614281565b805182525f60208201516101006020850152614866610100850182614253565b905060408301516040850152606083015184820360608601526148898282614253565b9150506080830151608085015260a083015160a085015260c083015160c085015260e083015184820360e0860152613c1582826147d0565b602081525f61242c6020830184614846565b5f5f5f5f5f5f60a087890312156148e8575f5ffd5b86356001600160401b038111156148fd575f5ffd5b61490989828a0161445e565b909a90995060208901359860408101359850606081013597506080013595509350505050565b5f5f5f5f5f60a08688031215614943575f5ffd5b85359450602086013593506040860135925060608601356001600160401b0381111561496d575f5ffd5b61497988828901614385565b95989497509295608001359392505050565b60ff60f81b8816815260e060208201525f6149a960e0830189614253565b82810360408401526149bb8189614253565b606084018890526001600160a01b038716608085015260a0840186905283810360c085015290506149ec8185614281565b9a9950505050505050505050565b8015158114612ef0575f5ffd5b5f5f5f5f60808587031215614a1a575f5ffd5b843593506020850135925060408501356001600160401b03811115614a3d575f5ffd5b614a4987828801614385565b9250506060850135614a5a816149fa565b939692955090935050565b5f5f5f5f5f60a08688031215614a79575f5ffd5b8535945060208601356001600160401b03811115614a95575f5ffd5b614aa188828901614385565b945050604086013592506060860135915060808601356001600160401b03811115614aca575f5ffd5b614ad688828901614385565b9150509295509295909350565b5f60208284031215614af3575f5ffd5b61242c826141fa565b5f5f5f5f60808587031215614b0f575f5ffd5b84356001600160401b03811115614b24575f5ffd5b614b3087828801614385565b94505060208501356001600160401b03811115614b4b575f5ffd5b614b5787828801614385565b93505060408501359150614b6d606086016141fa565b905092959194509250565b606081525f614b8a6060830186614846565b602083019490945250901515604090910152919050565b5f5f5f60408486031215614bb3575f5ffd5b83356001600160401b03811115614bc8575f5ffd5b614bd48682870161445e565b90945092505060208401356001600160401b03811115614bf2575f5ffd5b614bfe86828701614385565b9150509250925092565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561436557603f19878603018452614c4a858351614281565b94506020938401939190910190600101614c2e565b5f5f5f5f60808587031215614c72575f5ffd5b843593506020850135925060408501356001600160401b03811115614c95575f5ffd5b614ca187828801614385565b949793965093946060013593505050565b5f60208284031215614cc2575f5ffd5b81356001600160401b03811115614cd7575f5ffd5b6117df84828501614385565b5f5f5f60608486031215614cf5575f5ffd5b83359250614d0560208501614701565b929592945050506040919091013590565b5f5f5f5f60808587031215614d29575f5ffd5b84356001600160401b03811115614d3e575f5ffd5b614d4a87828801614385565b945050614d59602086016141fa565b93969395505050506040820135916060013590565b602081525f61242c60208301846142bb565b5f5f5f60608486031215614d92575f5ffd5b833592506020840135915060408401356001600160401b03811115614bf2575f5ffd5b5f5f5f60608486031215614dc7575f5ffd5b8335925060208401356001600160401b03811115614de3575f5ffd5b614def86828701614385565b93969395505050506040919091013590565b5f81518060208401855e5f93019283525090919050565b5f614e238285614e01565b60609390931b6bffffffffffffffffffffffff191683525050601401919050565b634e487b7160e01b5f52603260045260245ffd5b8281525f6117df6020830184614e01565b5f61242c8284614e01565b600181811c90821680614e8857607f821691505b602082108103614ea657634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115612faf57805f5260205f20601f840160051c81016020851015614ed15750805b601f840160051c820191505b81811015614ef0575f8155600101614edd565b5050505050565b81516001600160401b03811115614f1057614f10614371565b614f2481614f1e8454614e74565b84614eac565b6020601f821160018114614f56575f8315614f3f5750848201515b5f19600385901b1c1916600184901b178455614ef0565b5f84815260208120601f198516915b82811015614f855787850151825560209485019460019092019101614f65565b5084821015614fa257868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b81810381811115610f7d57610f7d614fb1565b8381528260208201525f613c156040830184614e01565b634e487b7160e01b5f52603160045260245ffd5b80820180821115610f7d57610f7d614fb1565b8082028115828204841417610f7d57610f7d614fb1565b5f6020828403121561503d575f5ffd5b815161242c816149fa565b5f60208284031215615058575f5ffd5b5051919050565b60ff8181168382160190811115610f7d57610f7d614fb1565b61ffff8181168382160190811115610f7d57610f7d614fb1565b634e487b7160e01b5f52602160045260245ffdfea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10053746f7261676544617461286279746573206368756e6b4349442c6279746573333220626c6f636b4349442c75696e74323536206368756e6b496e6465782c75696e743820626c6f636b496e6465782c62797465733332206e6f646549642c75696e74323536206e6f6e63652c75696e7432353620646561646c696e652c62797465733332206275636b6574496429360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca264697066735822122018816c5541f07001830d4fbf6d37bbd4014e06af55d1d56f4d628ccdce364f6f64736f6c634300081c0033",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// StorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageMetaData.Bin instead.
var StorageBin = StorageMetaData.Bin

// DeployStorage deploys a new Ethereum contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// MAXBLOCKSPERFILE is a free data retrieval call binding the contract method 0x9ccd4646.
//
// Solidity: function MAX_BLOCKS_PER_FILE() view returns(uint64)
func (_Storage *StorageCaller) MAXBLOCKSPERFILE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "MAX_BLOCKS_PER_FILE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXBLOCKSPERFILE is a free data retrieval call binding the contract method 0x9ccd4646.
//
// Solidity: function MAX_BLOCKS_PER_FILE() view returns(uint64)
func (_Storage *StorageSession) MAXBLOCKSPERFILE() (uint64, error) {
	return _Storage.Contract.MAXBLOCKSPERFILE(&_Storage.CallOpts)
}

// MAXBLOCKSPERFILE is a free data retrieval call binding the contract method 0x9ccd4646.
//
// Solidity: function MAX_BLOCKS_PER_FILE() view returns(uint64)
func (_Storage *StorageCallerSession) MAXBLOCKSPERFILE() (uint64, error) {
	return _Storage.Contract.MAXBLOCKSPERFILE(&_Storage.CallOpts)
}

// MAXBLOCKSIZE is a free data retrieval call binding the contract method 0x6ce02363.
//
// Solidity: function MAX_BLOCK_SIZE() view returns(uint64)
func (_Storage *StorageCaller) MAXBLOCKSIZE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "MAX_BLOCK_SIZE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXBLOCKSIZE is a free data retrieval call binding the contract method 0x6ce02363.
//
// Solidity: function MAX_BLOCK_SIZE() view returns(uint64)
func (_Storage *StorageSession) MAXBLOCKSIZE() (uint64, error) {
	return _Storage.Contract.MAXBLOCKSIZE(&_Storage.CallOpts)
}

// MAXBLOCKSIZE is a free data retrieval call binding the contract method 0x6ce02363.
//
// Solidity: function MAX_BLOCK_SIZE() view returns(uint64)
func (_Storage *StorageCallerSession) MAXBLOCKSIZE() (uint64, error) {
	return _Storage.Contract.MAXBLOCKSIZE(&_Storage.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Storage *StorageCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Storage *StorageSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Storage.Contract.UPGRADEINTERFACEVERSION(&_Storage.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Storage *StorageCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Storage.Contract.UPGRADEINTERFACEVERSION(&_Storage.CallOpts)
}

// AccessManager is a free data retrieval call binding the contract method 0xfdcb6068.
//
// Solidity: function accessManager() view returns(address)
func (_Storage *StorageCaller) AccessManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "accessManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccessManager is a free data retrieval call binding the contract method 0xfdcb6068.
//
// Solidity: function accessManager() view returns(address)
func (_Storage *StorageSession) AccessManager() (common.Address, error) {
	return _Storage.Contract.AccessManager(&_Storage.CallOpts)
}

// AccessManager is a free data retrieval call binding the contract method 0xfdcb6068.
//
// Solidity: function accessManager() view returns(address)
func (_Storage *StorageCallerSession) AccessManager() (common.Address, error) {
	return _Storage.Contract.AccessManager(&_Storage.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Storage *StorageCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Storage *StorageSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Storage.Contract.Eip712Domain(&_Storage.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Storage *StorageCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Storage.Contract.Eip712Domain(&_Storage.CallOpts)
}

// FileFillCounter is a free data retrieval call binding the contract method 0xf8a3e41a.
//
// Solidity: function fileFillCounter(bytes32 ) view returns(uint16)
func (_Storage *StorageCaller) FileFillCounter(opts *bind.CallOpts, arg0 [32]byte) (uint16, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "fileFillCounter", arg0)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// FileFillCounter is a free data retrieval call binding the contract method 0xf8a3e41a.
//
// Solidity: function fileFillCounter(bytes32 ) view returns(uint16)
func (_Storage *StorageSession) FileFillCounter(arg0 [32]byte) (uint16, error) {
	return _Storage.Contract.FileFillCounter(&_Storage.CallOpts, arg0)
}

// FileFillCounter is a free data retrieval call binding the contract method 0xf8a3e41a.
//
// Solidity: function fileFillCounter(bytes32 ) view returns(uint16)
func (_Storage *StorageCallerSession) FileFillCounter(arg0 [32]byte) (uint16, error) {
	return _Storage.Contract.FileFillCounter(&_Storage.CallOpts, arg0)
}

// FileRewardClaimed is a free data retrieval call binding the contract method 0x018c1e9c.
//
// Solidity: function fileRewardClaimed(bytes32 ) view returns(bool)
func (_Storage *StorageCaller) FileRewardClaimed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "fileRewardClaimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FileRewardClaimed is a free data retrieval call binding the contract method 0x018c1e9c.
//
// Solidity: function fileRewardClaimed(bytes32 ) view returns(bool)
func (_Storage *StorageSession) FileRewardClaimed(arg0 [32]byte) (bool, error) {
	return _Storage.Contract.FileRewardClaimed(&_Storage.CallOpts, arg0)
}

// FileRewardClaimed is a free data retrieval call binding the contract method 0x018c1e9c.
//
// Solidity: function fileRewardClaimed(bytes32 ) view returns(bool)
func (_Storage *StorageCallerSession) FileRewardClaimed(arg0 [32]byte) (bool, error) {
	return _Storage.Contract.FileRewardClaimed(&_Storage.CallOpts, arg0)
}

// FulfilledBlocks is a free data retrieval call binding the contract method 0x5a4e9564.
//
// Solidity: function fulfilledBlocks(bytes32 , uint256 ) view returns(uint32)
func (_Storage *StorageCaller) FulfilledBlocks(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "fulfilledBlocks", arg0, arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FulfilledBlocks is a free data retrieval call binding the contract method 0x5a4e9564.
//
// Solidity: function fulfilledBlocks(bytes32 , uint256 ) view returns(uint32)
func (_Storage *StorageSession) FulfilledBlocks(arg0 [32]byte, arg1 *big.Int) (uint32, error) {
	return _Storage.Contract.FulfilledBlocks(&_Storage.CallOpts, arg0, arg1)
}

// FulfilledBlocks is a free data retrieval call binding the contract method 0x5a4e9564.
//
// Solidity: function fulfilledBlocks(bytes32 , uint256 ) view returns(uint32)
func (_Storage *StorageCallerSession) FulfilledBlocks(arg0 [32]byte, arg1 *big.Int) (uint32, error) {
	return _Storage.Contract.FulfilledBlocks(&_Storage.CallOpts, arg0, arg1)
}

// GetBucketByName is a free data retrieval call binding the contract method 0xf83533d2.
//
// Solidity: function getBucketByName(string bucketName, address owner, uint256 fileOffset, uint256 fileLimit) view returns((bytes32,string,uint256,address,bytes32[]) bucket)
func (_Storage *StorageCaller) GetBucketByName(opts *bind.CallOpts, bucketName string, owner common.Address, fileOffset *big.Int, fileLimit *big.Int) (IStorageBucket, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getBucketByName", bucketName, owner, fileOffset, fileLimit)

	if err != nil {
		return *new(IStorageBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(IStorageBucket)).(*IStorageBucket)

	return out0, err

}

// GetBucketByName is a free data retrieval call binding the contract method 0xf83533d2.
//
// Solidity: function getBucketByName(string bucketName, address owner, uint256 fileOffset, uint256 fileLimit) view returns((bytes32,string,uint256,address,bytes32[]) bucket)
func (_Storage *StorageSession) GetBucketByName(bucketName string, owner common.Address, fileOffset *big.Int, fileLimit *big.Int) (IStorageBucket, error) {
	return _Storage.Contract.GetBucketByName(&_Storage.CallOpts, bucketName, owner, fileOffset, fileLimit)
}

// GetBucketByName is a free data retrieval call binding the contract method 0xf83533d2.
//
// Solidity: function getBucketByName(string bucketName, address owner, uint256 fileOffset, uint256 fileLimit) view returns((bytes32,string,uint256,address,bytes32[]) bucket)
func (_Storage *StorageCallerSession) GetBucketByName(bucketName string, owner common.Address, fileOffset *big.Int, fileLimit *big.Int) (IStorageBucket, error) {
	return _Storage.Contract.GetBucketByName(&_Storage.CallOpts, bucketName, owner, fileOffset, fileLimit)
}

// GetBucketIndexByName is a free data retrieval call binding the contract method 0x287e677f.
//
// Solidity: function getBucketIndexByName(string bucketName, address owner) view returns(uint256 index, bool exists)
func (_Storage *StorageCaller) GetBucketIndexByName(opts *bind.CallOpts, bucketName string, owner common.Address) (struct {
	Index  *big.Int
	Exists bool
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getBucketIndexByName", bucketName, owner)

	outstruct := new(struct {
		Index  *big.Int
		Exists bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Exists = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetBucketIndexByName is a free data retrieval call binding the contract method 0x287e677f.
//
// Solidity: function getBucketIndexByName(string bucketName, address owner) view returns(uint256 index, bool exists)
func (_Storage *StorageSession) GetBucketIndexByName(bucketName string, owner common.Address) (struct {
	Index  *big.Int
	Exists bool
}, error) {
	return _Storage.Contract.GetBucketIndexByName(&_Storage.CallOpts, bucketName, owner)
}

// GetBucketIndexByName is a free data retrieval call binding the contract method 0x287e677f.
//
// Solidity: function getBucketIndexByName(string bucketName, address owner) view returns(uint256 index, bool exists)
func (_Storage *StorageCallerSession) GetBucketIndexByName(bucketName string, owner common.Address) (struct {
	Index  *big.Int
	Exists bool
}, error) {
	return _Storage.Contract.GetBucketIndexByName(&_Storage.CallOpts, bucketName, owner)
}

// GetBucketsByIds is a free data retrieval call binding the contract method 0x35bdb711.
//
// Solidity: function getBucketsByIds(bytes32[] ids) view returns((bytes32,string,uint256,address,bytes32[])[])
func (_Storage *StorageCaller) GetBucketsByIds(opts *bind.CallOpts, ids [][32]byte) ([]IStorageBucket, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getBucketsByIds", ids)

	if err != nil {
		return *new([]IStorageBucket), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStorageBucket)).(*[]IStorageBucket)

	return out0, err

}

// GetBucketsByIds is a free data retrieval call binding the contract method 0x35bdb711.
//
// Solidity: function getBucketsByIds(bytes32[] ids) view returns((bytes32,string,uint256,address,bytes32[])[])
func (_Storage *StorageSession) GetBucketsByIds(ids [][32]byte) ([]IStorageBucket, error) {
	return _Storage.Contract.GetBucketsByIds(&_Storage.CallOpts, ids)
}

// GetBucketsByIds is a free data retrieval call binding the contract method 0x35bdb711.
//
// Solidity: function getBucketsByIds(bytes32[] ids) view returns((bytes32,string,uint256,address,bytes32[])[])
func (_Storage *StorageCallerSession) GetBucketsByIds(ids [][32]byte) ([]IStorageBucket, error) {
	return _Storage.Contract.GetBucketsByIds(&_Storage.CallOpts, ids)
}

// GetBucketsByIdsWithFiles is a free data retrieval call binding the contract method 0x62838c23.
//
// Solidity: function getBucketsByIdsWithFiles(bytes32[] ids, uint256 bucketOffset, uint256 bucketLimit, uint256 fileOffset, uint256 fileLimit) view returns((bytes32,string,uint256,address,bytes32[])[] buckets)
func (_Storage *StorageCaller) GetBucketsByIdsWithFiles(opts *bind.CallOpts, ids [][32]byte, bucketOffset *big.Int, bucketLimit *big.Int, fileOffset *big.Int, fileLimit *big.Int) ([]IStorageBucket, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getBucketsByIdsWithFiles", ids, bucketOffset, bucketLimit, fileOffset, fileLimit)

	if err != nil {
		return *new([]IStorageBucket), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStorageBucket)).(*[]IStorageBucket)

	return out0, err

}

// GetBucketsByIdsWithFiles is a free data retrieval call binding the contract method 0x62838c23.
//
// Solidity: function getBucketsByIdsWithFiles(bytes32[] ids, uint256 bucketOffset, uint256 bucketLimit, uint256 fileOffset, uint256 fileLimit) view returns((bytes32,string,uint256,address,bytes32[])[] buckets)
func (_Storage *StorageSession) GetBucketsByIdsWithFiles(ids [][32]byte, bucketOffset *big.Int, bucketLimit *big.Int, fileOffset *big.Int, fileLimit *big.Int) ([]IStorageBucket, error) {
	return _Storage.Contract.GetBucketsByIdsWithFiles(&_Storage.CallOpts, ids, bucketOffset, bucketLimit, fileOffset, fileLimit)
}

// GetBucketsByIdsWithFiles is a free data retrieval call binding the contract method 0x62838c23.
//
// Solidity: function getBucketsByIdsWithFiles(bytes32[] ids, uint256 bucketOffset, uint256 bucketLimit, uint256 fileOffset, uint256 fileLimit) view returns((bytes32,string,uint256,address,bytes32[])[] buckets)
func (_Storage *StorageCallerSession) GetBucketsByIdsWithFiles(ids [][32]byte, bucketOffset *big.Int, bucketLimit *big.Int, fileOffset *big.Int, fileLimit *big.Int) ([]IStorageBucket, error) {
	return _Storage.Contract.GetBucketsByIdsWithFiles(&_Storage.CallOpts, ids, bucketOffset, bucketLimit, fileOffset, fileLimit)
}

// GetChainID is a free data retrieval call binding the contract method 0x564b81ef.
//
// Solidity: function getChainID() view returns(uint256)
func (_Storage *StorageCaller) GetChainID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getChainID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainID is a free data retrieval call binding the contract method 0x564b81ef.
//
// Solidity: function getChainID() view returns(uint256)
func (_Storage *StorageSession) GetChainID() (*big.Int, error) {
	return _Storage.Contract.GetChainID(&_Storage.CallOpts)
}

// GetChainID is a free data retrieval call binding the contract method 0x564b81ef.
//
// Solidity: function getChainID() view returns(uint256)
func (_Storage *StorageCallerSession) GetChainID() (*big.Int, error) {
	return _Storage.Contract.GetChainID(&_Storage.CallOpts)
}

// GetChunkByIndex is a free data retrieval call binding the contract method 0x4d7dc614.
//
// Solidity: function getChunkByIndex(bytes32 id, uint256 index) view returns(bytes, uint256)
func (_Storage *StorageCaller) GetChunkByIndex(opts *bind.CallOpts, id [32]byte, index *big.Int) ([]byte, *big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getChunkByIndex", id, index)

	if err != nil {
		return *new([]byte), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetChunkByIndex is a free data retrieval call binding the contract method 0x4d7dc614.
//
// Solidity: function getChunkByIndex(bytes32 id, uint256 index) view returns(bytes, uint256)
func (_Storage *StorageSession) GetChunkByIndex(id [32]byte, index *big.Int) ([]byte, *big.Int, error) {
	return _Storage.Contract.GetChunkByIndex(&_Storage.CallOpts, id, index)
}

// GetChunkByIndex is a free data retrieval call binding the contract method 0x4d7dc614.
//
// Solidity: function getChunkByIndex(bytes32 id, uint256 index) view returns(bytes, uint256)
func (_Storage *StorageCallerSession) GetChunkByIndex(id [32]byte, index *big.Int) ([]byte, *big.Int, error) {
	return _Storage.Contract.GetChunkByIndex(&_Storage.CallOpts, id, index)
}

// GetFileById is a free data retrieval call binding the contract method 0xfaec0542.
//
// Solidity: function getFileById(bytes32 id) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[])) file)
func (_Storage *StorageCaller) GetFileById(opts *bind.CallOpts, id [32]byte) (IStorageFile, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getFileById", id)

	if err != nil {
		return *new(IStorageFile), err
	}

	out0 := *abi.ConvertType(out[0], new(IStorageFile)).(*IStorageFile)

	return out0, err

}

// GetFileById is a free data retrieval call binding the contract method 0xfaec0542.
//
// Solidity: function getFileById(bytes32 id) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[])) file)
func (_Storage *StorageSession) GetFileById(id [32]byte) (IStorageFile, error) {
	return _Storage.Contract.GetFileById(&_Storage.CallOpts, id)
}

// GetFileById is a free data retrieval call binding the contract method 0xfaec0542.
//
// Solidity: function getFileById(bytes32 id) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[])) file)
func (_Storage *StorageCallerSession) GetFileById(id [32]byte) (IStorageFile, error) {
	return _Storage.Contract.GetFileById(&_Storage.CallOpts, id)
}

// GetFileByName is a free data retrieval call binding the contract method 0x5ecdfb53.
//
// Solidity: function getFileByName(bytes32 bucketId, string fileName) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[])) file)
func (_Storage *StorageCaller) GetFileByName(opts *bind.CallOpts, bucketId [32]byte, fileName string) (IStorageFile, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getFileByName", bucketId, fileName)

	if err != nil {
		return *new(IStorageFile), err
	}

	out0 := *abi.ConvertType(out[0], new(IStorageFile)).(*IStorageFile)

	return out0, err

}

// GetFileByName is a free data retrieval call binding the contract method 0x5ecdfb53.
//
// Solidity: function getFileByName(bytes32 bucketId, string fileName) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[])) file)
func (_Storage *StorageSession) GetFileByName(bucketId [32]byte, fileName string) (IStorageFile, error) {
	return _Storage.Contract.GetFileByName(&_Storage.CallOpts, bucketId, fileName)
}

// GetFileByName is a free data retrieval call binding the contract method 0x5ecdfb53.
//
// Solidity: function getFileByName(bytes32 bucketId, string fileName) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[])) file)
func (_Storage *StorageCallerSession) GetFileByName(bucketId [32]byte, fileName string) (IStorageFile, error) {
	return _Storage.Contract.GetFileByName(&_Storage.CallOpts, bucketId, fileName)
}

// GetFileIndexById is a free data retrieval call binding the contract method 0x49cdf6ac.
//
// Solidity: function getFileIndexById(string bucketName, bytes32 fileId, address owner) view returns(uint256 index, bool exists)
func (_Storage *StorageCaller) GetFileIndexById(opts *bind.CallOpts, bucketName string, fileId [32]byte, owner common.Address) (struct {
	Index  *big.Int
	Exists bool
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getFileIndexById", bucketName, fileId, owner)

	outstruct := new(struct {
		Index  *big.Int
		Exists bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Exists = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetFileIndexById is a free data retrieval call binding the contract method 0x49cdf6ac.
//
// Solidity: function getFileIndexById(string bucketName, bytes32 fileId, address owner) view returns(uint256 index, bool exists)
func (_Storage *StorageSession) GetFileIndexById(bucketName string, fileId [32]byte, owner common.Address) (struct {
	Index  *big.Int
	Exists bool
}, error) {
	return _Storage.Contract.GetFileIndexById(&_Storage.CallOpts, bucketName, fileId, owner)
}

// GetFileIndexById is a free data retrieval call binding the contract method 0x49cdf6ac.
//
// Solidity: function getFileIndexById(string bucketName, bytes32 fileId, address owner) view returns(uint256 index, bool exists)
func (_Storage *StorageCallerSession) GetFileIndexById(bucketName string, fileId [32]byte, owner common.Address) (struct {
	Index  *big.Int
	Exists bool
}, error) {
	return _Storage.Contract.GetFileIndexById(&_Storage.CallOpts, bucketName, fileId, owner)
}

// GetFileOwner is a free data retrieval call binding the contract method 0x1b475ef4.
//
// Solidity: function getFileOwner(bytes32 id) view returns(address)
func (_Storage *StorageCaller) GetFileOwner(opts *bind.CallOpts, id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getFileOwner", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFileOwner is a free data retrieval call binding the contract method 0x1b475ef4.
//
// Solidity: function getFileOwner(bytes32 id) view returns(address)
func (_Storage *StorageSession) GetFileOwner(id [32]byte) (common.Address, error) {
	return _Storage.Contract.GetFileOwner(&_Storage.CallOpts, id)
}

// GetFileOwner is a free data retrieval call binding the contract method 0x1b475ef4.
//
// Solidity: function getFileOwner(bytes32 id) view returns(address)
func (_Storage *StorageCallerSession) GetFileOwner(id [32]byte) (common.Address, error) {
	return _Storage.Contract.GetFileOwner(&_Storage.CallOpts, id)
}

// GetFullFileInfo is a free data retrieval call binding the contract method 0xd4967ce6.
//
// Solidity: function getFullFileInfo(string bucketName, string fileName, bytes32 bucketId, address owner) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[])) file, uint256 index, bool exists)
func (_Storage *StorageCaller) GetFullFileInfo(opts *bind.CallOpts, bucketName string, fileName string, bucketId [32]byte, owner common.Address) (struct {
	File   IStorageFile
	Index  *big.Int
	Exists bool
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getFullFileInfo", bucketName, fileName, bucketId, owner)

	outstruct := new(struct {
		File   IStorageFile
		Index  *big.Int
		Exists bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.File = *abi.ConvertType(out[0], new(IStorageFile)).(*IStorageFile)
	outstruct.Index = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Exists = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetFullFileInfo is a free data retrieval call binding the contract method 0xd4967ce6.
//
// Solidity: function getFullFileInfo(string bucketName, string fileName, bytes32 bucketId, address owner) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[])) file, uint256 index, bool exists)
func (_Storage *StorageSession) GetFullFileInfo(bucketName string, fileName string, bucketId [32]byte, owner common.Address) (struct {
	File   IStorageFile
	Index  *big.Int
	Exists bool
}, error) {
	return _Storage.Contract.GetFullFileInfo(&_Storage.CallOpts, bucketName, fileName, bucketId, owner)
}

// GetFullFileInfo is a free data retrieval call binding the contract method 0xd4967ce6.
//
// Solidity: function getFullFileInfo(string bucketName, string fileName, bytes32 bucketId, address owner) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[])) file, uint256 index, bool exists)
func (_Storage *StorageCallerSession) GetFullFileInfo(bucketName string, fileName string, bucketId [32]byte, owner common.Address) (struct {
	File   IStorageFile
	Index  *big.Int
	Exists bool
}, error) {
	return _Storage.Contract.GetFullFileInfo(&_Storage.CallOpts, bucketName, fileName, bucketId, owner)
}

// GetOwnerBuckets is a free data retrieval call binding the contract method 0x262a61ce.
//
// Solidity: function getOwnerBuckets(address owner, uint256 offset, uint256 limit, uint256 fileOffset, uint256 fileLimit) view returns((bytes32,string,uint256,address,bytes32[])[] buckets)
func (_Storage *StorageCaller) GetOwnerBuckets(opts *bind.CallOpts, owner common.Address, offset *big.Int, limit *big.Int, fileOffset *big.Int, fileLimit *big.Int) ([]IStorageBucket, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getOwnerBuckets", owner, offset, limit, fileOffset, fileLimit)

	if err != nil {
		return *new([]IStorageBucket), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStorageBucket)).(*[]IStorageBucket)

	return out0, err

}

// GetOwnerBuckets is a free data retrieval call binding the contract method 0x262a61ce.
//
// Solidity: function getOwnerBuckets(address owner, uint256 offset, uint256 limit, uint256 fileOffset, uint256 fileLimit) view returns((bytes32,string,uint256,address,bytes32[])[] buckets)
func (_Storage *StorageSession) GetOwnerBuckets(owner common.Address, offset *big.Int, limit *big.Int, fileOffset *big.Int, fileLimit *big.Int) ([]IStorageBucket, error) {
	return _Storage.Contract.GetOwnerBuckets(&_Storage.CallOpts, owner, offset, limit, fileOffset, fileLimit)
}

// GetOwnerBuckets is a free data retrieval call binding the contract method 0x262a61ce.
//
// Solidity: function getOwnerBuckets(address owner, uint256 offset, uint256 limit, uint256 fileOffset, uint256 fileLimit) view returns((bytes32,string,uint256,address,bytes32[])[] buckets)
func (_Storage *StorageCallerSession) GetOwnerBuckets(owner common.Address, offset *big.Int, limit *big.Int, fileOffset *big.Int, fileLimit *big.Int) ([]IStorageBucket, error) {
	return _Storage.Contract.GetOwnerBuckets(&_Storage.CallOpts, owner, offset, limit, fileOffset, fileLimit)
}

// GetPeerBlockIndexById is a free data retrieval call binding the contract method 0xf855169a.
//
// Solidity: function getPeerBlockIndexById(bytes32 peerId, bytes32 cid, string fileName) view returns(uint256 index, bool exists)
func (_Storage *StorageCaller) GetPeerBlockIndexById(opts *bind.CallOpts, peerId [32]byte, cid [32]byte, fileName string) (struct {
	Index  *big.Int
	Exists bool
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getPeerBlockIndexById", peerId, cid, fileName)

	outstruct := new(struct {
		Index  *big.Int
		Exists bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Exists = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetPeerBlockIndexById is a free data retrieval call binding the contract method 0xf855169a.
//
// Solidity: function getPeerBlockIndexById(bytes32 peerId, bytes32 cid, string fileName) view returns(uint256 index, bool exists)
func (_Storage *StorageSession) GetPeerBlockIndexById(peerId [32]byte, cid [32]byte, fileName string) (struct {
	Index  *big.Int
	Exists bool
}, error) {
	return _Storage.Contract.GetPeerBlockIndexById(&_Storage.CallOpts, peerId, cid, fileName)
}

// GetPeerBlockIndexById is a free data retrieval call binding the contract method 0xf855169a.
//
// Solidity: function getPeerBlockIndexById(bytes32 peerId, bytes32 cid, string fileName) view returns(uint256 index, bool exists)
func (_Storage *StorageCallerSession) GetPeerBlockIndexById(peerId [32]byte, cid [32]byte, fileName string) (struct {
	Index  *big.Int
	Exists bool
}, error) {
	return _Storage.Contract.GetPeerBlockIndexById(&_Storage.CallOpts, peerId, cid, fileName)
}

// GetPeersArrayByPeerBlockCid is a free data retrieval call binding the contract method 0xd606205d.
//
// Solidity: function getPeersArrayByPeerBlockCid(bytes32[] cids, string fileName) view returns(bytes32[][] peers)
func (_Storage *StorageCaller) GetPeersArrayByPeerBlockCid(opts *bind.CallOpts, cids [][32]byte, fileName string) ([][][32]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getPeersArrayByPeerBlockCid", cids, fileName)

	if err != nil {
		return *new([][][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][][32]byte)).(*[][][32]byte)

	return out0, err

}

// GetPeersArrayByPeerBlockCid is a free data retrieval call binding the contract method 0xd606205d.
//
// Solidity: function getPeersArrayByPeerBlockCid(bytes32[] cids, string fileName) view returns(bytes32[][] peers)
func (_Storage *StorageSession) GetPeersArrayByPeerBlockCid(cids [][32]byte, fileName string) ([][][32]byte, error) {
	return _Storage.Contract.GetPeersArrayByPeerBlockCid(&_Storage.CallOpts, cids, fileName)
}

// GetPeersArrayByPeerBlockCid is a free data retrieval call binding the contract method 0xd606205d.
//
// Solidity: function getPeersArrayByPeerBlockCid(bytes32[] cids, string fileName) view returns(bytes32[][] peers)
func (_Storage *StorageCallerSession) GetPeersArrayByPeerBlockCid(cids [][32]byte, fileName string) ([][][32]byte, error) {
	return _Storage.Contract.GetPeersArrayByPeerBlockCid(&_Storage.CallOpts, cids, fileName)
}

// GetPeersByPeerBlockCid is a free data retrieval call binding the contract method 0x48b49875.
//
// Solidity: function getPeersByPeerBlockCid(bytes32 cid, string fileName) view returns(bytes32[] peers)
func (_Storage *StorageCaller) GetPeersByPeerBlockCid(opts *bind.CallOpts, cid [32]byte, fileName string) ([][32]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getPeersByPeerBlockCid", cid, fileName)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetPeersByPeerBlockCid is a free data retrieval call binding the contract method 0x48b49875.
//
// Solidity: function getPeersByPeerBlockCid(bytes32 cid, string fileName) view returns(bytes32[] peers)
func (_Storage *StorageSession) GetPeersByPeerBlockCid(cid [32]byte, fileName string) ([][32]byte, error) {
	return _Storage.Contract.GetPeersByPeerBlockCid(&_Storage.CallOpts, cid, fileName)
}

// GetPeersByPeerBlockCid is a free data retrieval call binding the contract method 0x48b49875.
//
// Solidity: function getPeersByPeerBlockCid(bytes32 cid, string fileName) view returns(bytes32[] peers)
func (_Storage *StorageCallerSession) GetPeersByPeerBlockCid(cid [32]byte, fileName string) ([][32]byte, error) {
	return _Storage.Contract.GetPeersByPeerBlockCid(&_Storage.CallOpts, cid, fileName)
}

// IsBlockFilled is a free data retrieval call binding the contract method 0xe4ba8a58.
//
// Solidity: function isBlockFilled(bytes32 fileId, uint8 blockIndex, uint256 chunkIndex) view returns(bool)
func (_Storage *StorageCaller) IsBlockFilled(opts *bind.CallOpts, fileId [32]byte, blockIndex uint8, chunkIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "isBlockFilled", fileId, blockIndex, chunkIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlockFilled is a free data retrieval call binding the contract method 0xe4ba8a58.
//
// Solidity: function isBlockFilled(bytes32 fileId, uint8 blockIndex, uint256 chunkIndex) view returns(bool)
func (_Storage *StorageSession) IsBlockFilled(fileId [32]byte, blockIndex uint8, chunkIndex *big.Int) (bool, error) {
	return _Storage.Contract.IsBlockFilled(&_Storage.CallOpts, fileId, blockIndex, chunkIndex)
}

// IsBlockFilled is a free data retrieval call binding the contract method 0xe4ba8a58.
//
// Solidity: function isBlockFilled(bytes32 fileId, uint8 blockIndex, uint256 chunkIndex) view returns(bool)
func (_Storage *StorageCallerSession) IsBlockFilled(fileId [32]byte, blockIndex uint8, chunkIndex *big.Int) (bool, error) {
	return _Storage.Contract.IsBlockFilled(&_Storage.CallOpts, fileId, blockIndex, chunkIndex)
}

// IsChunkFilled is a free data retrieval call binding the contract method 0x3f383980.
//
// Solidity: function isChunkFilled(bytes32 fileId, uint256 chunkIndex) view returns(bool)
func (_Storage *StorageCaller) IsChunkFilled(opts *bind.CallOpts, fileId [32]byte, chunkIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "isChunkFilled", fileId, chunkIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsChunkFilled is a free data retrieval call binding the contract method 0x3f383980.
//
// Solidity: function isChunkFilled(bytes32 fileId, uint256 chunkIndex) view returns(bool)
func (_Storage *StorageSession) IsChunkFilled(fileId [32]byte, chunkIndex *big.Int) (bool, error) {
	return _Storage.Contract.IsChunkFilled(&_Storage.CallOpts, fileId, chunkIndex)
}

// IsChunkFilled is a free data retrieval call binding the contract method 0x3f383980.
//
// Solidity: function isChunkFilled(bytes32 fileId, uint256 chunkIndex) view returns(bool)
func (_Storage *StorageCallerSession) IsChunkFilled(fileId [32]byte, chunkIndex *big.Int) (bool, error) {
	return _Storage.Contract.IsChunkFilled(&_Storage.CallOpts, fileId, chunkIndex)
}

// IsFileFilled is a free data retrieval call binding the contract method 0x68e6408f.
//
// Solidity: function isFileFilled(bytes32 fileId) view returns(bool)
func (_Storage *StorageCaller) IsFileFilled(opts *bind.CallOpts, fileId [32]byte) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "isFileFilled", fileId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFileFilled is a free data retrieval call binding the contract method 0x68e6408f.
//
// Solidity: function isFileFilled(bytes32 fileId) view returns(bool)
func (_Storage *StorageSession) IsFileFilled(fileId [32]byte) (bool, error) {
	return _Storage.Contract.IsFileFilled(&_Storage.CallOpts, fileId)
}

// IsFileFilled is a free data retrieval call binding the contract method 0x68e6408f.
//
// Solidity: function isFileFilled(bytes32 fileId) view returns(bool)
func (_Storage *StorageCallerSession) IsFileFilled(fileId [32]byte) (bool, error) {
	return _Storage.Contract.IsFileFilled(&_Storage.CallOpts, fileId)
}

// IsFileFilledV2 is a free data retrieval call binding the contract method 0xfd21c284.
//
// Solidity: function isFileFilledV2(bytes32 fileId) view returns(bool)
func (_Storage *StorageCaller) IsFileFilledV2(opts *bind.CallOpts, fileId [32]byte) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "isFileFilledV2", fileId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFileFilledV2 is a free data retrieval call binding the contract method 0xfd21c284.
//
// Solidity: function isFileFilledV2(bytes32 fileId) view returns(bool)
func (_Storage *StorageSession) IsFileFilledV2(fileId [32]byte) (bool, error) {
	return _Storage.Contract.IsFileFilledV2(&_Storage.CallOpts, fileId)
}

// IsFileFilledV2 is a free data retrieval call binding the contract method 0xfd21c284.
//
// Solidity: function isFileFilledV2(bytes32 fileId) view returns(bool)
func (_Storage *StorageCallerSession) IsFileFilledV2(fileId [32]byte) (bool, error) {
	return _Storage.Contract.IsFileFilledV2(&_Storage.CallOpts, fileId)
}

// IsPeerBlockReplica is a free data retrieval call binding the contract method 0x7912bf68.
//
// Solidity: function isPeerBlockReplica(bytes32 cid, bytes32 peerId) view returns(bool)
func (_Storage *StorageCaller) IsPeerBlockReplica(opts *bind.CallOpts, cid [32]byte, peerId [32]byte) (bool, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "isPeerBlockReplica", cid, peerId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPeerBlockReplica is a free data retrieval call binding the contract method 0x7912bf68.
//
// Solidity: function isPeerBlockReplica(bytes32 cid, bytes32 peerId) view returns(bool)
func (_Storage *StorageSession) IsPeerBlockReplica(cid [32]byte, peerId [32]byte) (bool, error) {
	return _Storage.Contract.IsPeerBlockReplica(&_Storage.CallOpts, cid, peerId)
}

// IsPeerBlockReplica is a free data retrieval call binding the contract method 0x7912bf68.
//
// Solidity: function isPeerBlockReplica(bytes32 cid, bytes32 peerId) view returns(bool)
func (_Storage *StorageCallerSession) IsPeerBlockReplica(cid [32]byte, peerId [32]byte) (bool, error) {
	return _Storage.Contract.IsPeerBlockReplica(&_Storage.CallOpts, cid, peerId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Storage *StorageCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Storage *StorageSession) ProxiableUUID() ([32]byte, error) {
	return _Storage.Contract.ProxiableUUID(&_Storage.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Storage *StorageCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Storage.Contract.ProxiableUUID(&_Storage.CallOpts)
}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() view returns(uint256)
func (_Storage *StorageCaller) Timestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "timestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() view returns(uint256)
func (_Storage *StorageSession) Timestamp() (*big.Int, error) {
	return _Storage.Contract.Timestamp(&_Storage.CallOpts)
}

// Timestamp is a free data retrieval call binding the contract method 0xb80777ea.
//
// Solidity: function timestamp() view returns(uint256)
func (_Storage *StorageCallerSession) Timestamp() (*big.Int, error) {
	return _Storage.Contract.Timestamp(&_Storage.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Storage *StorageCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Storage *StorageSession) Token() (common.Address, error) {
	return _Storage.Contract.Token(&_Storage.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Storage *StorageCallerSession) Token() (common.Address, error) {
	return _Storage.Contract.Token(&_Storage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Storage *StorageCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Storage *StorageSession) Version() (string, error) {
	return _Storage.Contract.Version(&_Storage.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Storage *StorageCallerSession) Version() (string, error) {
	return _Storage.Contract.Version(&_Storage.CallOpts)
}

// AddFileChunk is a paid mutator transaction binding the contract method 0x30b91d07.
//
// Solidity: function addFileChunk(bytes cid, bytes32 bucketId, string fileName, uint256 encodedChunkSize, bytes32[] cids, uint256[] chunkBlocksSizes, uint256 chunkIndex) returns(bytes32)
func (_Storage *StorageTransactor) AddFileChunk(opts *bind.TransactOpts, cid []byte, bucketId [32]byte, fileName string, encodedChunkSize *big.Int, cids [][32]byte, chunkBlocksSizes []*big.Int, chunkIndex *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addFileChunk", cid, bucketId, fileName, encodedChunkSize, cids, chunkBlocksSizes, chunkIndex)
}

// AddFileChunk is a paid mutator transaction binding the contract method 0x30b91d07.
//
// Solidity: function addFileChunk(bytes cid, bytes32 bucketId, string fileName, uint256 encodedChunkSize, bytes32[] cids, uint256[] chunkBlocksSizes, uint256 chunkIndex) returns(bytes32)
func (_Storage *StorageSession) AddFileChunk(cid []byte, bucketId [32]byte, fileName string, encodedChunkSize *big.Int, cids [][32]byte, chunkBlocksSizes []*big.Int, chunkIndex *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddFileChunk(&_Storage.TransactOpts, cid, bucketId, fileName, encodedChunkSize, cids, chunkBlocksSizes, chunkIndex)
}

// AddFileChunk is a paid mutator transaction binding the contract method 0x30b91d07.
//
// Solidity: function addFileChunk(bytes cid, bytes32 bucketId, string fileName, uint256 encodedChunkSize, bytes32[] cids, uint256[] chunkBlocksSizes, uint256 chunkIndex) returns(bytes32)
func (_Storage *StorageTransactorSession) AddFileChunk(cid []byte, bucketId [32]byte, fileName string, encodedChunkSize *big.Int, cids [][32]byte, chunkBlocksSizes []*big.Int, chunkIndex *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddFileChunk(&_Storage.TransactOpts, cid, bucketId, fileName, encodedChunkSize, cids, chunkBlocksSizes, chunkIndex)
}

// AddPeerBlock is a paid mutator transaction binding the contract method 0x87c1ac07.
//
// Solidity: function addPeerBlock(bytes32 peerId, bytes32 cid, string fileName, bool isReplica) returns(bytes32 id)
func (_Storage *StorageTransactor) AddPeerBlock(opts *bind.TransactOpts, peerId [32]byte, cid [32]byte, fileName string, isReplica bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addPeerBlock", peerId, cid, fileName, isReplica)
}

// AddPeerBlock is a paid mutator transaction binding the contract method 0x87c1ac07.
//
// Solidity: function addPeerBlock(bytes32 peerId, bytes32 cid, string fileName, bool isReplica) returns(bytes32 id)
func (_Storage *StorageSession) AddPeerBlock(peerId [32]byte, cid [32]byte, fileName string, isReplica bool) (*types.Transaction, error) {
	return _Storage.Contract.AddPeerBlock(&_Storage.TransactOpts, peerId, cid, fileName, isReplica)
}

// AddPeerBlock is a paid mutator transaction binding the contract method 0x87c1ac07.
//
// Solidity: function addPeerBlock(bytes32 peerId, bytes32 cid, string fileName, bool isReplica) returns(bytes32 id)
func (_Storage *StorageTransactorSession) AddPeerBlock(peerId [32]byte, cid [32]byte, fileName string, isReplica bool) (*types.Transaction, error) {
	return _Storage.Contract.AddPeerBlock(&_Storage.TransactOpts, peerId, cid, fileName, isReplica)
}

// CommitFile is a paid mutator transaction binding the contract method 0x9a2e82b3.
//
// Solidity: function commitFile(bytes32 bucketId, string fileName, uint256 encodedFileSize, uint256 actualSize, bytes fileCID) returns(bytes32)
func (_Storage *StorageTransactor) CommitFile(opts *bind.TransactOpts, bucketId [32]byte, fileName string, encodedFileSize *big.Int, actualSize *big.Int, fileCID []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "commitFile", bucketId, fileName, encodedFileSize, actualSize, fileCID)
}

// CommitFile is a paid mutator transaction binding the contract method 0x9a2e82b3.
//
// Solidity: function commitFile(bytes32 bucketId, string fileName, uint256 encodedFileSize, uint256 actualSize, bytes fileCID) returns(bytes32)
func (_Storage *StorageSession) CommitFile(bucketId [32]byte, fileName string, encodedFileSize *big.Int, actualSize *big.Int, fileCID []byte) (*types.Transaction, error) {
	return _Storage.Contract.CommitFile(&_Storage.TransactOpts, bucketId, fileName, encodedFileSize, actualSize, fileCID)
}

// CommitFile is a paid mutator transaction binding the contract method 0x9a2e82b3.
//
// Solidity: function commitFile(bytes32 bucketId, string fileName, uint256 encodedFileSize, uint256 actualSize, bytes fileCID) returns(bytes32)
func (_Storage *StorageTransactorSession) CommitFile(bucketId [32]byte, fileName string, encodedFileSize *big.Int, actualSize *big.Int, fileCID []byte) (*types.Transaction, error) {
	return _Storage.Contract.CommitFile(&_Storage.TransactOpts, bucketId, fileName, encodedFileSize, actualSize, fileCID)
}

// CreateBucket is a paid mutator transaction binding the contract method 0xe3f787e8.
//
// Solidity: function createBucket(string bucketName) returns(bytes32 id)
func (_Storage *StorageTransactor) CreateBucket(opts *bind.TransactOpts, bucketName string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "createBucket", bucketName)
}

// CreateBucket is a paid mutator transaction binding the contract method 0xe3f787e8.
//
// Solidity: function createBucket(string bucketName) returns(bytes32 id)
func (_Storage *StorageSession) CreateBucket(bucketName string) (*types.Transaction, error) {
	return _Storage.Contract.CreateBucket(&_Storage.TransactOpts, bucketName)
}

// CreateBucket is a paid mutator transaction binding the contract method 0xe3f787e8.
//
// Solidity: function createBucket(string bucketName) returns(bytes32 id)
func (_Storage *StorageTransactorSession) CreateBucket(bucketName string) (*types.Transaction, error) {
	return _Storage.Contract.CreateBucket(&_Storage.TransactOpts, bucketName)
}

// CreateFile is a paid mutator transaction binding the contract method 0x6af0f801.
//
// Solidity: function createFile(bytes32 bucketId, string fileName) returns(bytes32)
func (_Storage *StorageTransactor) CreateFile(opts *bind.TransactOpts, bucketId [32]byte, fileName string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "createFile", bucketId, fileName)
}

// CreateFile is a paid mutator transaction binding the contract method 0x6af0f801.
//
// Solidity: function createFile(bytes32 bucketId, string fileName) returns(bytes32)
func (_Storage *StorageSession) CreateFile(bucketId [32]byte, fileName string) (*types.Transaction, error) {
	return _Storage.Contract.CreateFile(&_Storage.TransactOpts, bucketId, fileName)
}

// CreateFile is a paid mutator transaction binding the contract method 0x6af0f801.
//
// Solidity: function createFile(bytes32 bucketId, string fileName) returns(bytes32)
func (_Storage *StorageTransactorSession) CreateFile(bucketId [32]byte, fileName string) (*types.Transaction, error) {
	return _Storage.Contract.CreateFile(&_Storage.TransactOpts, bucketId, fileName)
}

// DeleteBucket is a paid mutator transaction binding the contract method 0xfd1d3c0c.
//
// Solidity: function deleteBucket(bytes32 id, string bucketName, uint256 index) returns(bool)
func (_Storage *StorageTransactor) DeleteBucket(opts *bind.TransactOpts, id [32]byte, bucketName string, index *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteBucket", id, bucketName, index)
}

// DeleteBucket is a paid mutator transaction binding the contract method 0xfd1d3c0c.
//
// Solidity: function deleteBucket(bytes32 id, string bucketName, uint256 index) returns(bool)
func (_Storage *StorageSession) DeleteBucket(id [32]byte, bucketName string, index *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DeleteBucket(&_Storage.TransactOpts, id, bucketName, index)
}

// DeleteBucket is a paid mutator transaction binding the contract method 0xfd1d3c0c.
//
// Solidity: function deleteBucket(bytes32 id, string bucketName, uint256 index) returns(bool)
func (_Storage *StorageTransactorSession) DeleteBucket(id [32]byte, bucketName string, index *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DeleteBucket(&_Storage.TransactOpts, id, bucketName, index)
}

// DeleteFile is a paid mutator transaction binding the contract method 0xd6d3110b.
//
// Solidity: function deleteFile(bytes32 fileID, bytes32 bucketId, string fileName, uint256 index) returns(bool)
func (_Storage *StorageTransactor) DeleteFile(opts *bind.TransactOpts, fileID [32]byte, bucketId [32]byte, fileName string, index *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteFile", fileID, bucketId, fileName, index)
}

// DeleteFile is a paid mutator transaction binding the contract method 0xd6d3110b.
//
// Solidity: function deleteFile(bytes32 fileID, bytes32 bucketId, string fileName, uint256 index) returns(bool)
func (_Storage *StorageSession) DeleteFile(fileID [32]byte, bucketId [32]byte, fileName string, index *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DeleteFile(&_Storage.TransactOpts, fileID, bucketId, fileName, index)
}

// DeleteFile is a paid mutator transaction binding the contract method 0xd6d3110b.
//
// Solidity: function deleteFile(bytes32 fileID, bytes32 bucketId, string fileName, uint256 index) returns(bool)
func (_Storage *StorageTransactorSession) DeleteFile(fileID [32]byte, bucketId [32]byte, fileName string, index *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DeleteFile(&_Storage.TransactOpts, fileID, bucketId, fileName, index)
}

// DeletePeerBlock is a paid mutator transaction binding the contract method 0x6a6e658b.
//
// Solidity: function deletePeerBlock(bytes32 id, bytes32 peerId, bytes32 cid, string fileName, uint256 index) returns(bool)
func (_Storage *StorageTransactor) DeletePeerBlock(opts *bind.TransactOpts, id [32]byte, peerId [32]byte, cid [32]byte, fileName string, index *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deletePeerBlock", id, peerId, cid, fileName, index)
}

// DeletePeerBlock is a paid mutator transaction binding the contract method 0x6a6e658b.
//
// Solidity: function deletePeerBlock(bytes32 id, bytes32 peerId, bytes32 cid, string fileName, uint256 index) returns(bool)
func (_Storage *StorageSession) DeletePeerBlock(id [32]byte, peerId [32]byte, cid [32]byte, fileName string, index *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DeletePeerBlock(&_Storage.TransactOpts, id, peerId, cid, fileName, index)
}

// DeletePeerBlock is a paid mutator transaction binding the contract method 0x6a6e658b.
//
// Solidity: function deletePeerBlock(bytes32 id, bytes32 peerId, bytes32 cid, string fileName, uint256 index) returns(bool)
func (_Storage *StorageTransactorSession) DeletePeerBlock(id [32]byte, peerId [32]byte, cid [32]byte, fileName string, index *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DeletePeerBlock(&_Storage.TransactOpts, id, peerId, cid, fileName, index)
}

// FillChunkBlock is a paid mutator transaction binding the contract method 0x4f55ba82.
//
// Solidity: function fillChunkBlock(bytes32 blockCID, bytes32 nodeId, bytes32 bucketId, uint256 chunkIndex, uint256 nonce, uint8 blockIndex, string fileName, bytes signature, uint256 deadline) returns()
func (_Storage *StorageTransactor) FillChunkBlock(opts *bind.TransactOpts, blockCID [32]byte, nodeId [32]byte, bucketId [32]byte, chunkIndex *big.Int, nonce *big.Int, blockIndex uint8, fileName string, signature []byte, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "fillChunkBlock", blockCID, nodeId, bucketId, chunkIndex, nonce, blockIndex, fileName, signature, deadline)
}

// FillChunkBlock is a paid mutator transaction binding the contract method 0x4f55ba82.
//
// Solidity: function fillChunkBlock(bytes32 blockCID, bytes32 nodeId, bytes32 bucketId, uint256 chunkIndex, uint256 nonce, uint8 blockIndex, string fileName, bytes signature, uint256 deadline) returns()
func (_Storage *StorageSession) FillChunkBlock(blockCID [32]byte, nodeId [32]byte, bucketId [32]byte, chunkIndex *big.Int, nonce *big.Int, blockIndex uint8, fileName string, signature []byte, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.FillChunkBlock(&_Storage.TransactOpts, blockCID, nodeId, bucketId, chunkIndex, nonce, blockIndex, fileName, signature, deadline)
}

// FillChunkBlock is a paid mutator transaction binding the contract method 0x4f55ba82.
//
// Solidity: function fillChunkBlock(bytes32 blockCID, bytes32 nodeId, bytes32 bucketId, uint256 chunkIndex, uint256 nonce, uint8 blockIndex, string fileName, bytes signature, uint256 deadline) returns()
func (_Storage *StorageTransactorSession) FillChunkBlock(blockCID [32]byte, nodeId [32]byte, bucketId [32]byte, chunkIndex *big.Int, nonce *big.Int, blockIndex uint8, fileName string, signature []byte, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.FillChunkBlock(&_Storage.TransactOpts, blockCID, nodeId, bucketId, chunkIndex, nonce, blockIndex, fileName, signature, deadline)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address tokenAddress) returns()
func (_Storage *StorageTransactor) Initialize(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "initialize", tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address tokenAddress) returns()
func (_Storage *StorageSession) Initialize(tokenAddress common.Address) (*types.Transaction, error) {
	return _Storage.Contract.Initialize(&_Storage.TransactOpts, tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address tokenAddress) returns()
func (_Storage *StorageTransactorSession) Initialize(tokenAddress common.Address) (*types.Transaction, error) {
	return _Storage.Contract.Initialize(&_Storage.TransactOpts, tokenAddress)
}

// SetAccessManager is a paid mutator transaction binding the contract method 0xc9580804.
//
// Solidity: function setAccessManager(address accessManagerAddress) returns()
func (_Storage *StorageTransactor) SetAccessManager(opts *bind.TransactOpts, accessManagerAddress common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setAccessManager", accessManagerAddress)
}

// SetAccessManager is a paid mutator transaction binding the contract method 0xc9580804.
//
// Solidity: function setAccessManager(address accessManagerAddress) returns()
func (_Storage *StorageSession) SetAccessManager(accessManagerAddress common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetAccessManager(&_Storage.TransactOpts, accessManagerAddress)
}

// SetAccessManager is a paid mutator transaction binding the contract method 0xc9580804.
//
// Solidity: function setAccessManager(address accessManagerAddress) returns()
func (_Storage *StorageTransactorSession) SetAccessManager(accessManagerAddress common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetAccessManager(&_Storage.TransactOpts, accessManagerAddress)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Storage *StorageTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Storage *StorageSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.UpgradeToAndCall(&_Storage.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Storage *StorageTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Storage.Contract.UpgradeToAndCall(&_Storage.TransactOpts, newImplementation, data)
}

// StorageAddFileBlocksIterator is returned from FilterAddFileBlocks and is used to iterate over the raw logs and unpacked data for AddFileBlocks events raised by the Storage contract.
type StorageAddFileBlocksIterator struct {
	Event *StorageAddFileBlocks // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageAddFileBlocksIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageAddFileBlocks)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageAddFileBlocks)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageAddFileBlocksIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageAddFileBlocksIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageAddFileBlocks represents a AddFileBlocks event raised by the Storage contract.
type StorageAddFileBlocks struct {
	Ids    [][32]byte
	FileId [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddFileBlocks is a free log retrieval operation binding the contract event 0x304b7f3b7c43506589700f0069a783fad42cfd6ef15dd044d805192bd79d3030.
//
// Solidity: event AddFileBlocks(bytes32[] indexed ids, bytes32 indexed fileId)
func (_Storage *StorageFilterer) FilterAddFileBlocks(opts *bind.FilterOpts, ids [][][32]byte, fileId [][32]byte) (*StorageAddFileBlocksIterator, error) {

	var idsRule []interface{}
	for _, idsItem := range ids {
		idsRule = append(idsRule, idsItem)
	}
	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "AddFileBlocks", idsRule, fileIdRule)
	if err != nil {
		return nil, err
	}
	return &StorageAddFileBlocksIterator{contract: _Storage.contract, event: "AddFileBlocks", logs: logs, sub: sub}, nil
}

// WatchAddFileBlocks is a free log subscription operation binding the contract event 0x304b7f3b7c43506589700f0069a783fad42cfd6ef15dd044d805192bd79d3030.
//
// Solidity: event AddFileBlocks(bytes32[] indexed ids, bytes32 indexed fileId)
func (_Storage *StorageFilterer) WatchAddFileBlocks(opts *bind.WatchOpts, sink chan<- *StorageAddFileBlocks, ids [][][32]byte, fileId [][32]byte) (event.Subscription, error) {

	var idsRule []interface{}
	for _, idsItem := range ids {
		idsRule = append(idsRule, idsItem)
	}
	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "AddFileBlocks", idsRule, fileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageAddFileBlocks)
				if err := _Storage.contract.UnpackLog(event, "AddFileBlocks", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddFileBlocks is a log parse operation binding the contract event 0x304b7f3b7c43506589700f0069a783fad42cfd6ef15dd044d805192bd79d3030.
//
// Solidity: event AddFileBlocks(bytes32[] indexed ids, bytes32 indexed fileId)
func (_Storage *StorageFilterer) ParseAddFileBlocks(log types.Log) (*StorageAddFileBlocks, error) {
	event := new(StorageAddFileBlocks)
	if err := _Storage.contract.UnpackLog(event, "AddFileBlocks", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageAddFileChunkIterator is returned from FilterAddFileChunk and is used to iterate over the raw logs and unpacked data for AddFileChunk events raised by the Storage contract.
type StorageAddFileChunkIterator struct {
	Event *StorageAddFileChunk // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageAddFileChunkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageAddFileChunk)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageAddFileChunk)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageAddFileChunkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageAddFileChunkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageAddFileChunk represents a AddFileChunk event raised by the Storage contract.
type StorageAddFileChunk struct {
	Id       [32]byte
	BucketId [32]byte
	Name     common.Hash
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddFileChunk is a free log retrieval operation binding the contract event 0x0d877d3656ad561e6c8224e31cc98f1b3e62b1d828fd396c47ef0daccb6838e5.
//
// Solidity: event AddFileChunk(bytes32 indexed id, bytes32 indexed bucketId, string indexed name, address owner)
func (_Storage *StorageFilterer) FilterAddFileChunk(opts *bind.FilterOpts, id [][32]byte, bucketId [][32]byte, name []string) (*StorageAddFileChunkIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var bucketIdRule []interface{}
	for _, bucketIdItem := range bucketId {
		bucketIdRule = append(bucketIdRule, bucketIdItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "AddFileChunk", idRule, bucketIdRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &StorageAddFileChunkIterator{contract: _Storage.contract, event: "AddFileChunk", logs: logs, sub: sub}, nil
}

// WatchAddFileChunk is a free log subscription operation binding the contract event 0x0d877d3656ad561e6c8224e31cc98f1b3e62b1d828fd396c47ef0daccb6838e5.
//
// Solidity: event AddFileChunk(bytes32 indexed id, bytes32 indexed bucketId, string indexed name, address owner)
func (_Storage *StorageFilterer) WatchAddFileChunk(opts *bind.WatchOpts, sink chan<- *StorageAddFileChunk, id [][32]byte, bucketId [][32]byte, name []string) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var bucketIdRule []interface{}
	for _, bucketIdItem := range bucketId {
		bucketIdRule = append(bucketIdRule, bucketIdItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "AddFileChunk", idRule, bucketIdRule, nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageAddFileChunk)
				if err := _Storage.contract.UnpackLog(event, "AddFileChunk", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddFileChunk is a log parse operation binding the contract event 0x0d877d3656ad561e6c8224e31cc98f1b3e62b1d828fd396c47ef0daccb6838e5.
//
// Solidity: event AddFileChunk(bytes32 indexed id, bytes32 indexed bucketId, string indexed name, address owner)
func (_Storage *StorageFilterer) ParseAddFileChunk(log types.Log) (*StorageAddFileChunk, error) {
	event := new(StorageAddFileChunk)
	if err := _Storage.contract.UnpackLog(event, "AddFileChunk", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageAddPeerBlockIterator is returned from FilterAddPeerBlock and is used to iterate over the raw logs and unpacked data for AddPeerBlock events raised by the Storage contract.
type StorageAddPeerBlockIterator struct {
	Event *StorageAddPeerBlock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageAddPeerBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageAddPeerBlock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageAddPeerBlock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageAddPeerBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageAddPeerBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageAddPeerBlock represents a AddPeerBlock event raised by the Storage contract.
type StorageAddPeerBlock struct {
	BlockId [32]byte
	PeerId  [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddPeerBlock is a free log retrieval operation binding the contract event 0x9f91e060f6a60bc2a5e695f07d854d6a659eab9f5e1441ac24f6e62614bb8b5b.
//
// Solidity: event AddPeerBlock(bytes32 indexed blockId, bytes32 indexed peerId)
func (_Storage *StorageFilterer) FilterAddPeerBlock(opts *bind.FilterOpts, blockId [][32]byte, peerId [][32]byte) (*StorageAddPeerBlockIterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var peerIdRule []interface{}
	for _, peerIdItem := range peerId {
		peerIdRule = append(peerIdRule, peerIdItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "AddPeerBlock", blockIdRule, peerIdRule)
	if err != nil {
		return nil, err
	}
	return &StorageAddPeerBlockIterator{contract: _Storage.contract, event: "AddPeerBlock", logs: logs, sub: sub}, nil
}

// WatchAddPeerBlock is a free log subscription operation binding the contract event 0x9f91e060f6a60bc2a5e695f07d854d6a659eab9f5e1441ac24f6e62614bb8b5b.
//
// Solidity: event AddPeerBlock(bytes32 indexed blockId, bytes32 indexed peerId)
func (_Storage *StorageFilterer) WatchAddPeerBlock(opts *bind.WatchOpts, sink chan<- *StorageAddPeerBlock, blockId [][32]byte, peerId [][32]byte) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var peerIdRule []interface{}
	for _, peerIdItem := range peerId {
		peerIdRule = append(peerIdRule, peerIdItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "AddPeerBlock", blockIdRule, peerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageAddPeerBlock)
				if err := _Storage.contract.UnpackLog(event, "AddPeerBlock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddPeerBlock is a log parse operation binding the contract event 0x9f91e060f6a60bc2a5e695f07d854d6a659eab9f5e1441ac24f6e62614bb8b5b.
//
// Solidity: event AddPeerBlock(bytes32 indexed blockId, bytes32 indexed peerId)
func (_Storage *StorageFilterer) ParseAddPeerBlock(log types.Log) (*StorageAddPeerBlock, error) {
	event := new(StorageAddPeerBlock)
	if err := _Storage.contract.UnpackLog(event, "AddPeerBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageCommitFileIterator is returned from FilterCommitFile and is used to iterate over the raw logs and unpacked data for CommitFile events raised by the Storage contract.
type StorageCommitFileIterator struct {
	Event *StorageCommitFile // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageCommitFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageCommitFile)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageCommitFile)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageCommitFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageCommitFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageCommitFile represents a CommitFile event raised by the Storage contract.
type StorageCommitFile struct {
	Id       [32]byte
	BucketId [32]byte
	Name     common.Hash
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommitFile is a free log retrieval operation binding the contract event 0xc208b556b9717832827b84e296c7f4bb33292ba520fed2eae2f7c0d61f1c0a61.
//
// Solidity: event CommitFile(bytes32 indexed id, bytes32 indexed bucketId, string indexed name, address owner)
func (_Storage *StorageFilterer) FilterCommitFile(opts *bind.FilterOpts, id [][32]byte, bucketId [][32]byte, name []string) (*StorageCommitFileIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var bucketIdRule []interface{}
	for _, bucketIdItem := range bucketId {
		bucketIdRule = append(bucketIdRule, bucketIdItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "CommitFile", idRule, bucketIdRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &StorageCommitFileIterator{contract: _Storage.contract, event: "CommitFile", logs: logs, sub: sub}, nil
}

// WatchCommitFile is a free log subscription operation binding the contract event 0xc208b556b9717832827b84e296c7f4bb33292ba520fed2eae2f7c0d61f1c0a61.
//
// Solidity: event CommitFile(bytes32 indexed id, bytes32 indexed bucketId, string indexed name, address owner)
func (_Storage *StorageFilterer) WatchCommitFile(opts *bind.WatchOpts, sink chan<- *StorageCommitFile, id [][32]byte, bucketId [][32]byte, name []string) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var bucketIdRule []interface{}
	for _, bucketIdItem := range bucketId {
		bucketIdRule = append(bucketIdRule, bucketIdItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "CommitFile", idRule, bucketIdRule, nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageCommitFile)
				if err := _Storage.contract.UnpackLog(event, "CommitFile", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommitFile is a log parse operation binding the contract event 0xc208b556b9717832827b84e296c7f4bb33292ba520fed2eae2f7c0d61f1c0a61.
//
// Solidity: event CommitFile(bytes32 indexed id, bytes32 indexed bucketId, string indexed name, address owner)
func (_Storage *StorageFilterer) ParseCommitFile(log types.Log) (*StorageCommitFile, error) {
	event := new(StorageCommitFile)
	if err := _Storage.contract.UnpackLog(event, "CommitFile", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageCreateBucketIterator is returned from FilterCreateBucket and is used to iterate over the raw logs and unpacked data for CreateBucket events raised by the Storage contract.
type StorageCreateBucketIterator struct {
	Event *StorageCreateBucket // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageCreateBucketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageCreateBucket)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageCreateBucket)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageCreateBucketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageCreateBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageCreateBucket represents a CreateBucket event raised by the Storage contract.
type StorageCreateBucket struct {
	Id    [32]byte
	Name  common.Hash
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCreateBucket is a free log retrieval operation binding the contract event 0xb8be57bce74a717a1bbd4acf428df655720fce75c1854b02a88484388df241a8.
//
// Solidity: event CreateBucket(bytes32 indexed id, string indexed name, address indexed owner)
func (_Storage *StorageFilterer) FilterCreateBucket(opts *bind.FilterOpts, id [][32]byte, name []string, owner []common.Address) (*StorageCreateBucketIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "CreateBucket", idRule, nameRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &StorageCreateBucketIterator{contract: _Storage.contract, event: "CreateBucket", logs: logs, sub: sub}, nil
}

// WatchCreateBucket is a free log subscription operation binding the contract event 0xb8be57bce74a717a1bbd4acf428df655720fce75c1854b02a88484388df241a8.
//
// Solidity: event CreateBucket(bytes32 indexed id, string indexed name, address indexed owner)
func (_Storage *StorageFilterer) WatchCreateBucket(opts *bind.WatchOpts, sink chan<- *StorageCreateBucket, id [][32]byte, name []string, owner []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "CreateBucket", idRule, nameRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageCreateBucket)
				if err := _Storage.contract.UnpackLog(event, "CreateBucket", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreateBucket is a log parse operation binding the contract event 0xb8be57bce74a717a1bbd4acf428df655720fce75c1854b02a88484388df241a8.
//
// Solidity: event CreateBucket(bytes32 indexed id, string indexed name, address indexed owner)
func (_Storage *StorageFilterer) ParseCreateBucket(log types.Log) (*StorageCreateBucket, error) {
	event := new(StorageCreateBucket)
	if err := _Storage.contract.UnpackLog(event, "CreateBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageCreateFileIterator is returned from FilterCreateFile and is used to iterate over the raw logs and unpacked data for CreateFile events raised by the Storage contract.
type StorageCreateFileIterator struct {
	Event *StorageCreateFile // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageCreateFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageCreateFile)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageCreateFile)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageCreateFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageCreateFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageCreateFile represents a CreateFile event raised by the Storage contract.
type StorageCreateFile struct {
	Id       [32]byte
	BucketId [32]byte
	Name     common.Hash
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCreateFile is a free log retrieval operation binding the contract event 0xb018e47bdb983351e1bee22415a8f41eef5c2bf1c43c6d3d0992e678ae762df5.
//
// Solidity: event CreateFile(bytes32 indexed id, bytes32 indexed bucketId, string indexed name, address owner)
func (_Storage *StorageFilterer) FilterCreateFile(opts *bind.FilterOpts, id [][32]byte, bucketId [][32]byte, name []string) (*StorageCreateFileIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var bucketIdRule []interface{}
	for _, bucketIdItem := range bucketId {
		bucketIdRule = append(bucketIdRule, bucketIdItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "CreateFile", idRule, bucketIdRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &StorageCreateFileIterator{contract: _Storage.contract, event: "CreateFile", logs: logs, sub: sub}, nil
}

// WatchCreateFile is a free log subscription operation binding the contract event 0xb018e47bdb983351e1bee22415a8f41eef5c2bf1c43c6d3d0992e678ae762df5.
//
// Solidity: event CreateFile(bytes32 indexed id, bytes32 indexed bucketId, string indexed name, address owner)
func (_Storage *StorageFilterer) WatchCreateFile(opts *bind.WatchOpts, sink chan<- *StorageCreateFile, id [][32]byte, bucketId [][32]byte, name []string) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var bucketIdRule []interface{}
	for _, bucketIdItem := range bucketId {
		bucketIdRule = append(bucketIdRule, bucketIdItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "CreateFile", idRule, bucketIdRule, nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageCreateFile)
				if err := _Storage.contract.UnpackLog(event, "CreateFile", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreateFile is a log parse operation binding the contract event 0xb018e47bdb983351e1bee22415a8f41eef5c2bf1c43c6d3d0992e678ae762df5.
//
// Solidity: event CreateFile(bytes32 indexed id, bytes32 indexed bucketId, string indexed name, address owner)
func (_Storage *StorageFilterer) ParseCreateFile(log types.Log) (*StorageCreateFile, error) {
	event := new(StorageCreateFile)
	if err := _Storage.contract.UnpackLog(event, "CreateFile", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageDeleteBucketIterator is returned from FilterDeleteBucket and is used to iterate over the raw logs and unpacked data for DeleteBucket events raised by the Storage contract.
type StorageDeleteBucketIterator struct {
	Event *StorageDeleteBucket // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageDeleteBucketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageDeleteBucket)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageDeleteBucket)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageDeleteBucketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageDeleteBucketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageDeleteBucket represents a DeleteBucket event raised by the Storage contract.
type StorageDeleteBucket struct {
	Id    [32]byte
	Name  common.Hash
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeleteBucket is a free log retrieval operation binding the contract event 0xeda2cc8e002ead8000b1e0c1debfc9a88bd7ee6e94b8dc0763db17849fcf0389.
//
// Solidity: event DeleteBucket(bytes32 indexed id, string indexed name, address indexed owner)
func (_Storage *StorageFilterer) FilterDeleteBucket(opts *bind.FilterOpts, id [][32]byte, name []string, owner []common.Address) (*StorageDeleteBucketIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "DeleteBucket", idRule, nameRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &StorageDeleteBucketIterator{contract: _Storage.contract, event: "DeleteBucket", logs: logs, sub: sub}, nil
}

// WatchDeleteBucket is a free log subscription operation binding the contract event 0xeda2cc8e002ead8000b1e0c1debfc9a88bd7ee6e94b8dc0763db17849fcf0389.
//
// Solidity: event DeleteBucket(bytes32 indexed id, string indexed name, address indexed owner)
func (_Storage *StorageFilterer) WatchDeleteBucket(opts *bind.WatchOpts, sink chan<- *StorageDeleteBucket, id [][32]byte, name []string, owner []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "DeleteBucket", idRule, nameRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageDeleteBucket)
				if err := _Storage.contract.UnpackLog(event, "DeleteBucket", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeleteBucket is a log parse operation binding the contract event 0xeda2cc8e002ead8000b1e0c1debfc9a88bd7ee6e94b8dc0763db17849fcf0389.
//
// Solidity: event DeleteBucket(bytes32 indexed id, string indexed name, address indexed owner)
func (_Storage *StorageFilterer) ParseDeleteBucket(log types.Log) (*StorageDeleteBucket, error) {
	event := new(StorageDeleteBucket)
	if err := _Storage.contract.UnpackLog(event, "DeleteBucket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageDeleteFileIterator is returned from FilterDeleteFile and is used to iterate over the raw logs and unpacked data for DeleteFile events raised by the Storage contract.
type StorageDeleteFileIterator struct {
	Event *StorageDeleteFile // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageDeleteFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageDeleteFile)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageDeleteFile)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageDeleteFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageDeleteFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageDeleteFile represents a DeleteFile event raised by the Storage contract.
type StorageDeleteFile struct {
	Id       [32]byte
	BucketId [32]byte
	Name     common.Hash
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeleteFile is a free log retrieval operation binding the contract event 0x0e1bf50f5cca6659c62146db5b60160121a3752011b621d8c8953a1e0e23d048.
//
// Solidity: event DeleteFile(bytes32 indexed id, bytes32 indexed bucketId, string indexed name, address owner)
func (_Storage *StorageFilterer) FilterDeleteFile(opts *bind.FilterOpts, id [][32]byte, bucketId [][32]byte, name []string) (*StorageDeleteFileIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var bucketIdRule []interface{}
	for _, bucketIdItem := range bucketId {
		bucketIdRule = append(bucketIdRule, bucketIdItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "DeleteFile", idRule, bucketIdRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &StorageDeleteFileIterator{contract: _Storage.contract, event: "DeleteFile", logs: logs, sub: sub}, nil
}

// WatchDeleteFile is a free log subscription operation binding the contract event 0x0e1bf50f5cca6659c62146db5b60160121a3752011b621d8c8953a1e0e23d048.
//
// Solidity: event DeleteFile(bytes32 indexed id, bytes32 indexed bucketId, string indexed name, address owner)
func (_Storage *StorageFilterer) WatchDeleteFile(opts *bind.WatchOpts, sink chan<- *StorageDeleteFile, id [][32]byte, bucketId [][32]byte, name []string) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var bucketIdRule []interface{}
	for _, bucketIdItem := range bucketId {
		bucketIdRule = append(bucketIdRule, bucketIdItem)
	}
	var nameRule []interface{}
	for _, nameItem := range name {
		nameRule = append(nameRule, nameItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "DeleteFile", idRule, bucketIdRule, nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageDeleteFile)
				if err := _Storage.contract.UnpackLog(event, "DeleteFile", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeleteFile is a log parse operation binding the contract event 0x0e1bf50f5cca6659c62146db5b60160121a3752011b621d8c8953a1e0e23d048.
//
// Solidity: event DeleteFile(bytes32 indexed id, bytes32 indexed bucketId, string indexed name, address owner)
func (_Storage *StorageFilterer) ParseDeleteFile(log types.Log) (*StorageDeleteFile, error) {
	event := new(StorageDeleteFile)
	if err := _Storage.contract.UnpackLog(event, "DeleteFile", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageDeletePeerBlockIterator is returned from FilterDeletePeerBlock and is used to iterate over the raw logs and unpacked data for DeletePeerBlock events raised by the Storage contract.
type StorageDeletePeerBlockIterator struct {
	Event *StorageDeletePeerBlock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageDeletePeerBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageDeletePeerBlock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageDeletePeerBlock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageDeletePeerBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageDeletePeerBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageDeletePeerBlock represents a DeletePeerBlock event raised by the Storage contract.
type StorageDeletePeerBlock struct {
	BlockId [32]byte
	PeerId  [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeletePeerBlock is a free log retrieval operation binding the contract event 0x37505c6d2cdf9e778d6110c5ad2e49c649d521a18d2da6f007d3364bd83a45ae.
//
// Solidity: event DeletePeerBlock(bytes32 indexed blockId, bytes32 indexed peerId)
func (_Storage *StorageFilterer) FilterDeletePeerBlock(opts *bind.FilterOpts, blockId [][32]byte, peerId [][32]byte) (*StorageDeletePeerBlockIterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var peerIdRule []interface{}
	for _, peerIdItem := range peerId {
		peerIdRule = append(peerIdRule, peerIdItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "DeletePeerBlock", blockIdRule, peerIdRule)
	if err != nil {
		return nil, err
	}
	return &StorageDeletePeerBlockIterator{contract: _Storage.contract, event: "DeletePeerBlock", logs: logs, sub: sub}, nil
}

// WatchDeletePeerBlock is a free log subscription operation binding the contract event 0x37505c6d2cdf9e778d6110c5ad2e49c649d521a18d2da6f007d3364bd83a45ae.
//
// Solidity: event DeletePeerBlock(bytes32 indexed blockId, bytes32 indexed peerId)
func (_Storage *StorageFilterer) WatchDeletePeerBlock(opts *bind.WatchOpts, sink chan<- *StorageDeletePeerBlock, blockId [][32]byte, peerId [][32]byte) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var peerIdRule []interface{}
	for _, peerIdItem := range peerId {
		peerIdRule = append(peerIdRule, peerIdItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "DeletePeerBlock", blockIdRule, peerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageDeletePeerBlock)
				if err := _Storage.contract.UnpackLog(event, "DeletePeerBlock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeletePeerBlock is a log parse operation binding the contract event 0x37505c6d2cdf9e778d6110c5ad2e49c649d521a18d2da6f007d3364bd83a45ae.
//
// Solidity: event DeletePeerBlock(bytes32 indexed blockId, bytes32 indexed peerId)
func (_Storage *StorageFilterer) ParseDeletePeerBlock(log types.Log) (*StorageDeletePeerBlock, error) {
	event := new(StorageDeletePeerBlock)
	if err := _Storage.contract.UnpackLog(event, "DeletePeerBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Storage contract.
type StorageEIP712DomainChangedIterator struct {
	Event *StorageEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageEIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageEIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageEIP712DomainChanged represents a EIP712DomainChanged event raised by the Storage contract.
type StorageEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Storage *StorageFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*StorageEIP712DomainChangedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &StorageEIP712DomainChangedIterator{contract: _Storage.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Storage *StorageFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *StorageEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageEIP712DomainChanged)
				if err := _Storage.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Storage *StorageFilterer) ParseEIP712DomainChanged(log types.Log) (*StorageEIP712DomainChanged, error) {
	event := new(StorageEIP712DomainChanged)
	if err := _Storage.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageFillChunkBlockIterator is returned from FilterFillChunkBlock and is used to iterate over the raw logs and unpacked data for FillChunkBlock events raised by the Storage contract.
type StorageFillChunkBlockIterator struct {
	Event *StorageFillChunkBlock // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageFillChunkBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageFillChunkBlock)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageFillChunkBlock)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageFillChunkBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageFillChunkBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageFillChunkBlock represents a FillChunkBlock event raised by the Storage contract.
type StorageFillChunkBlock struct {
	Cid        [32]byte
	ChunkIndex *big.Int
	BlockIndex uint8
	NodeId     [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFillChunkBlock is a free log retrieval operation binding the contract event 0xeded691eefce1ca615096347126f78807fee63b86aab874574e9c733082ce867.
//
// Solidity: event FillChunkBlock(bytes32 indexed cid, uint256 chunkIndex, uint8 blockIndex, bytes32 indexed nodeId)
func (_Storage *StorageFilterer) FilterFillChunkBlock(opts *bind.FilterOpts, cid [][32]byte, nodeId [][32]byte) (*StorageFillChunkBlockIterator, error) {

	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "FillChunkBlock", cidRule, nodeIdRule)
	if err != nil {
		return nil, err
	}
	return &StorageFillChunkBlockIterator{contract: _Storage.contract, event: "FillChunkBlock", logs: logs, sub: sub}, nil
}

// WatchFillChunkBlock is a free log subscription operation binding the contract event 0xeded691eefce1ca615096347126f78807fee63b86aab874574e9c733082ce867.
//
// Solidity: event FillChunkBlock(bytes32 indexed cid, uint256 chunkIndex, uint8 blockIndex, bytes32 indexed nodeId)
func (_Storage *StorageFilterer) WatchFillChunkBlock(opts *bind.WatchOpts, sink chan<- *StorageFillChunkBlock, cid [][32]byte, nodeId [][32]byte) (event.Subscription, error) {

	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "FillChunkBlock", cidRule, nodeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageFillChunkBlock)
				if err := _Storage.contract.UnpackLog(event, "FillChunkBlock", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFillChunkBlock is a log parse operation binding the contract event 0xeded691eefce1ca615096347126f78807fee63b86aab874574e9c733082ce867.
//
// Solidity: event FillChunkBlock(bytes32 indexed cid, uint256 chunkIndex, uint8 blockIndex, bytes32 indexed nodeId)
func (_Storage *StorageFilterer) ParseFillChunkBlock(log types.Log) (*StorageFillChunkBlock, error) {
	event := new(StorageFillChunkBlock)
	if err := _Storage.contract.UnpackLog(event, "FillChunkBlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Storage contract.
type StorageInitializedIterator struct {
	Event *StorageInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageInitialized represents a Initialized event raised by the Storage contract.
type StorageInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Storage *StorageFilterer) FilterInitialized(opts *bind.FilterOpts) (*StorageInitializedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StorageInitializedIterator{contract: _Storage.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Storage *StorageFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StorageInitialized) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageInitialized)
				if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Storage *StorageFilterer) ParseInitialized(log types.Log) (*StorageInitialized, error) {
	event := new(StorageInitialized)
	if err := _Storage.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Storage contract.
type StorageUpgradedIterator struct {
	Event *StorageUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageUpgraded represents a Upgraded event raised by the Storage contract.
type StorageUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Storage *StorageFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StorageUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StorageUpgradedIterator{contract: _Storage.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Storage *StorageFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StorageUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageUpgraded)
				if err := _Storage.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Storage *StorageFilterer) ParseUpgraded(log types.Log) (*StorageUpgraded, error) {
	event := new(StorageUpgraded)
	if err := _Storage.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
