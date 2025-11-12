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
	ChunkCIDs       [][]byte
	ChunkSize       []*big.Int
	FulfilledBlocks []uint32
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

// IStorageFillChunkBlockArgs is an auto generated low-level Go binding around an user-defined struct.
type IStorageFillChunkBlockArgs struct {
	BlockCID   [32]byte
	NodeId     [32]byte
	BucketId   [32]byte
	ChunkIndex *big.Int
	Nonce      *big.Int
	BlockIndex uint8
	FileName   string
	Signature  []byte
	Deadline   *big.Int
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockAlreadyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockNonexists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketInvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketNonempty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketNonexists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"}],\"name\":\"ChunkCIDMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileChunkDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileFullyUploaded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNameDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNonempty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNotFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cidsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sizesLength\",\"type\":\"uint256\"}],\"name\":\"InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlocksAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEncodedSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFileBlocksCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFileCID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLastBlockSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPeerIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LastChunkDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPeersForCID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OffsetOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyBlockCIDs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"AddFileBlocks\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AddFileChunk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"}],\"name\":\"AddPeerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CommitFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CreateBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CreateFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DeleteBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DeleteFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"}],\"name\":\"DeletePeerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockCID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"}],\"name\":\"FillChunkBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_BLOCKS_PER_FILE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BLOCK_SIZE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessManager\",\"outputs\":[{\"internalType\":\"contractIAccessManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"chunkCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedChunkSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"cids\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkBlocksSizes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"addFileChunk\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"cids\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"encodedChunkSizes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"chunkBlocksCIDs\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"chunkBlockSizes\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256\",\"name\":\"startingChunkIndex\",\"type\":\"uint256\"}],\"name\":\"addFileChunks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedFileSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"}],\"name\":\"commitFile\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"createBucket\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"createFile\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deleteBucket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deleteFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fileFillCounter\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fileRewardClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockCID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIStorage.FillChunkBlockArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"fillChunkBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockCID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIStorage.FillChunkBlockArgs[]\",\"name\":\"args\",\"type\":\"tuple[]\"}],\"name\":\"fillChunkBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"blockCIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"fileID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"getBlockPeersOfChunk\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"peers\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fileOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileLimit\",\"type\":\"uint256\"}],\"name\":\"getBucketByName\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket\",\"name\":\"bucket\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getBucketIndexByName\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"}],\"name\":\"getBucketsByIds\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"bucketOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bucketLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileLimit\",\"type\":\"uint256\"}],\"name\":\"getBucketsByIdsWithFiles\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket[]\",\"name\":\"buckets\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getChunkByIndex\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getFileById\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"chunkCIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkSize\",\"type\":\"uint256[]\"},{\"internalType\":\"uint32[]\",\"name\":\"fulfilledBlocks\",\"type\":\"uint32[]\"}],\"internalType\":\"structIStorage.Chunk\",\"name\":\"chunks\",\"type\":\"tuple\"}],\"internalType\":\"structIStorage.File\",\"name\":\"file\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"getFileByName\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"chunkCIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkSize\",\"type\":\"uint256[]\"},{\"internalType\":\"uint32[]\",\"name\":\"fulfilledBlocks\",\"type\":\"uint32[]\"}],\"internalType\":\"structIStorage.Chunk\",\"name\":\"chunks\",\"type\":\"tuple\"}],\"internalType\":\"structIStorage.File\",\"name\":\"file\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getFileIndexById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getFileOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getFullFileInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"chunkCIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkSize\",\"type\":\"uint256[]\"},{\"internalType\":\"uint32[]\",\"name\":\"fulfilledBlocks\",\"type\":\"uint32[]\"}],\"internalType\":\"structIStorage.Chunk\",\"name\":\"chunks\",\"type\":\"tuple\"}],\"internalType\":\"structIStorage.File\",\"name\":\"file\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileLimit\",\"type\":\"uint256\"}],\"name\":\"getOwnerBuckets\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket[]\",\"name\":\"buckets\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"peerBlockID\",\"type\":\"bytes32\"}],\"name\":\"getPeerByPeerBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"isBlockFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"isChunkFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"isFileFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"isFileFilledV2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accessManagerAddress\",\"type\":\"address\"}],\"name\":\"setAccessManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIAkaveToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516156746100f95f395f81816133680152818161339101526134ca01526156745ff3fe60806040526004361061023e575f3560e01c80636af0f80111610134578063d6d3110b116100b3578063f8a3e41a11610078578063f8a3e41a146107a8578063faec0542146107ea578063fc0c546a14610809578063fd1d3c0c14610827578063fd21c28414610846578063fdcb606814610865575f5ffd5b8063d6d3110b146106f3578063e3f787e814610712578063e4ba8a5814610731578063f3b348fc14610750578063f83533d21461077c575f5ffd5b8063ad3cb1cc116100f9578063ad3cb1cc14610628578063b80777ea14610658578063c4d66de81461066a578063c958080414610689578063d4967ce6146106c5575f5ffd5b80636af0f801146105805780636ce023631461059f57806384b0196e146105cd5780639a2e82b3146105f45780639ccd464614610613575f5ffd5b806349cdf6ac116101c057806354fd4d501161018557806354fd4d50146104ac578063564b81ef146104e25780635ecdfb53146104f457806362838c231461052057806368e6408f1461053f575f5ffd5b806349cdf6ac1461041a5780634d7dc614146104395780634f1ef2861461046657806351dc0c231461047957806352d1902d14610498575f5ffd5b8063308c3cf211610206578063308c3cf21461036557806330b91d071461038457806335bdb711146103b15780633d840ee9146103d05780633f383980146103fb575f5ffd5b8063018c1e9c146102425780631b475ef4146102855780631ed572b7146102e4578063262a61ce14610305578063287e677f14610331575b5f5ffd5b34801561024d575f5ffd5b5061027061025c36600461458c565b60076020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b348015610290575f5ffd5b506102cc61029f36600461458c565b5f90815260046020908152604080832060029081015484529091529020600301546001600160a01b031690565b6040516001600160a01b03909116815260200161027c565b3480156102ef575f5ffd5b506103036102fe3660046145a3565b610884565b005b348015610310575f5ffd5b5061032461031f3660046145f5565b61093c565b60405161027c91906146e3565b34801561033c575f5ffd5b5061035061034b3660046147fb565b6109bc565b6040805192835290151560208301520161027c565b348015610370575f5ffd5b5061030361037f366004614885565b610a4e565b34801561038f575f5ffd5b506103a361039e3660046148c3565b610c60565b60405190815260200161027c565b3480156103bc575f5ffd5b506103246103cb366004614885565b610fa8565b3480156103db575f5ffd5b506103a36103ea36600461458c565b5f9081526008602052604090205490565b348015610406575f5ffd5b5061027061041536600461499c565b611217565b348015610425575f5ffd5b506103506104343660046149bc565b611337565b348015610444575f5ffd5b5061045861045336600461499c565b6113c2565b60405161027c929190614a0e565b610303610474366004614a2f565b6114ba565b348015610484575f5ffd5b50610303610493366004614b79565b6114d0565b3480156104a3575f5ffd5b506103a3611657565b3480156104b7575f5ffd5b506040805180820190915260058152640312e302e360dc1b60208201525b60405161027c9190614c22565b3480156104ed575f5ffd5b50466103a3565b3480156104ff575f5ffd5b5061051361050e366004614c34565b611672565b60405161027c9190614ddc565b34801561052b575f5ffd5b5061032461053a366004614dee565b6119d2565b34801561054a575f5ffd5b5061027061055936600461458c565b5f9081526005602090815260408083205460049092529091206007015461ffff9091161490565b34801561058b575f5ffd5b506103a361059a366004614c34565b611a17565b3480156105aa575f5ffd5b506105b5620f424081565b6040516001600160401b03909116815260200161027c565b3480156105d8575f5ffd5b506105e1611a9e565b60405161027c9796959493929190614e4a565b3480156105ff575f5ffd5b506103a361060e366004614eb9565b611b4c565b34801561061e575f5ffd5b506105b561040081565b348015610633575f5ffd5b506104d5604051806040016040528060058152602001640352e302e360dc1b81525081565b348015610663575f5ffd5b50426103a3565b348015610675575f5ffd5b50610303610684366004614f37565b611ea7565b348015610694575f5ffd5b506103036106a3366004614f37565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b3480156106d0575f5ffd5b506106e46106df366004614f50565b611ff9565b60405161027c93929190614fcc565b3480156106fe575f5ffd5b5061027061070d366004614ff5565b612057565b34801561071d575f5ffd5b506103a361072c366004615048565b6122fe565b34801561073c575f5ffd5b5061027061074b366004615089565b612464565b34801561075b575f5ffd5b5061076f61076a3660046150bc565b6124e8565b60405161027c9190615108565b348015610787575f5ffd5b5061079b61079636600461514a565b61263a565b60405161027c91906151a2565b3480156107b3575f5ffd5b506107d76107c236600461458c565b60056020525f908152604090205461ffff1681565b60405161ffff909116815260200161027c565b3480156107f5575f5ffd5b5061051361080436600461458c565b6126e5565b348015610814575f5ffd5b505f546102cc906001600160a01b031681565b348015610832575f5ffd5b506102706108413660046151b4565b612a1f565b348015610851575f5ffd5b5061027061086036600461458c565b612c5c565b348015610870575f5ffd5b506001546102cc906001600160a01b031681565b610939813560208301356040840135606085013560808601356108ad60c0880160a08901615200565b6108ba60c0890189615219565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506108fb9250505060e08a018a615219565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152505050506101008a0135612ca2565b50565b6001600160a01b0385165f90815260036020908152604091829020805483518184028101840190945280845260609391926109b1929184918301828280156109a157602002820191905f5260205f20905b81548152602001906001019080831161098d575b5050505050878787876001612fa7565b979650505050505050565b5f5f5f84846040516020016109d2929190615272565b60408051601f1981840301815291815281516020928301206001600160a01b0387165f9081526003909352908220909250905b8154811015610a435782828281548110610a2157610a2161529e565b905f5260205f20015403610a3b5780945060019350610a43565b600101610a05565b5050505b9250929050565b5f5b81811015610c5b57610c53838383818110610a6d57610a6d61529e565b9050602002810190610a7f91906152b2565b35848484818110610a9257610a9261529e565b9050602002810190610aa491906152b2565b60200135858585818110610aba57610aba61529e565b9050602002810190610acc91906152b2565b60400135868686818110610ae257610ae261529e565b9050602002810190610af491906152b2565b60600135878787818110610b0a57610b0a61529e565b9050602002810190610b1c91906152b2565b60800135888888818110610b3257610b3261529e565b9050602002810190610b4491906152b2565b610b559060c081019060a001615200565b898989818110610b6757610b6761529e565b9050602002810190610b7991906152b2565b610b879060c0810190615219565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508d92508c91508b9050818110610bcf57610bcf61529e565b9050602002810190610be191906152b2565b610bef9060e0810190615219565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508e92508d91508c9050818110610c3757610c3761529e565b9050602002810190610c4991906152b2565b6101000135612ca2565b600101610a50565b505050565b5f8881526002602052604081205489908203610c8f5760405163938a92b760e01b815260040160405180910390fd5b5f818152600260205260409020600301546001600160a01b03163314610cd557610cb8816132c5565b610cd55760405163dc64d0ad60e01b815260040160405180910390fd5b89895f8282604051602001610ceb9291906152d1565b60405160208183030381529060405280519060200120905060045f8281526020019081526020015f20600401545f14610d375760405163d96b03b160e01b815260040160405180910390fd5b5f8d8d604051602001610d4b9291906152d1565b60408051601f1981840301815291815281516020928301205f8181526004909352908220549092509003610d9257604051632abde33960e01b815260040160405180910390fd5b5f818152600460205260409020600701548714610dc2576040516301c0b3dd60e61b815260040160405180910390fd5b8988141580610dd1575060208a115b15610def576040516373d8ccd360e11b815260040160405180910390fd5b60208a1015610e45575f8181526006602052604090205461ffff1615610e28576040516355cbc83160e01b815260040160405180910390fd5b5f818152600660205260409020805461ffff191661ffff8c161790555b8b5f03610e6557604051631b6fdfeb60e01b815260040160405180910390fd5b8c604051610e7391906152e2565b604051908190038120338252908f9083907f0d877d3656ad561e6c8224e31cc98f1b3e62b1d828fd396c47ef0daccb6838e59060200160405180910390a460045f8281526020019081526020015f206007015f018f908060018154018082558091505060019003905f5260205f20015f909190919091509081610ef69190615370565b5060045f8281526020019081526020015f206007016001018c908060018154018082558091505060019003905f5260205f20015f909190919091505560045f8281526020019081526020015f206007016002015f908060018154018082558091505060019003905f5260205f2090600891828204019190066004029091909190916101000a81548163ffffffff021916908363ffffffff16021790555080955050505050509998505050505050505050565b60605f826001600160401b03811115610fc357610fc3614746565b604051908082528060200260200182016040528015610ffc57816020015b610fe96142fb565b815260200190600190039081610fe15790505b5090505f5b8381101561120d575f60025f87878581811061101f5761101f61529e565b9050602002013581526020019081526020015f206040518060a00160405290815f8201548152602001600182018054611057906152ed565b80601f0160208091040260200160405190810160405280929190818152602001828054611083906152ed565b80156110ce5780601f106110a5576101008083540402835291602001916110ce565b820191905f5260205f20905b8154815290600101906020018083116110b157829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020016004820180548060200260200160405190810160405280929190818152602001828054801561115c57602002820191905f5260205f20905b815481526020019060010190808311611148575b50505050508152505090506040518060a00160405280825f01518152602001826020015181526020018260400151815260200182606001516001600160a01b031681526020015f6001600160401b038111156111ba576111ba614746565b6040519080825280602002602001820160405280156111e3578160200160208202803683370190505b508152508383815181106111f9576111f961529e565b602090810291909101015250600101611001565b5090505b92915050565b5f8281526004602052604081206007015481906112369060019061543e565b9050808314801561125757505f8481526006602052604090205461ffff1615155b156112e2575f8481526006602052604081205461127d9060019061ffff1681901b61543e565b5f868152600460205260409020600901805491925063ffffffff831691839190879081106112ad576112ad61529e565b905f5260205f2090600891828204019190066004029054906101000a900463ffffffff161663ffffffff161492505050611211565b5f848152600460205260409020600901805463ffffffff91908590811061130b5761130b61529e565b5f918252602090912060088204015460079091166004026101000a900463ffffffff1614949350505050565b5f5f5f858460405160200161134d929190615272565b60408051601f1981840301815291815281516020928301205f8181526002909352908220909250600401905b81548110156113b757868282815481106113955761139561529e565b905f5260205f200154036113af57809450600193506113b7565b600101611379565b505050935093915050565b5f8281526004602052604081206007018054606092918291859081106113ea576113ea61529e565b905f5260205f200180546113fd906152ed565b80601f0160208091040260200160405190810160405280929190818152602001828054611429906152ed565b80156114745780601f1061144b57610100808354040283529160200191611474565b820191905f5260205f20905b81548152906001019060200180831161145757829003601f168201915b5050505f8881526004602052604081206008018054949550909390925087915081106114a2576114a261529e565b5f918252602090912001549196919550909350505050565b6114c261335d565b6114cc8282613403565b5050565b5f888152600260205260408120548991036114fe5760405163938a92b760e01b815260040160405180910390fd5b5f818152600260205260409020600301546001600160a01b0316331461154457611527816132c5565b6115445760405163dc64d0ad60e01b815260040160405180910390fd5b88885f828260405160200161155a9291906152d1565b60405160208183030381529060405280519060200120905060045f8281526020019081526020015f20600401545f146115a65760405163d96b03b160e01b815260040160405180910390fd5b5f5b8d518110156116475761163e8e82815181106115c6576115c661529e565b60200260200101518e8e8e85815181106115e2576115e261529e565b60200260200101518e8e878181106115fc576115fc61529e565b905060200281019061160e9190615451565b8e8e898181106116205761162061529e565b90506020028101906116329190615451565b898f61039e9190615496565b506001016115a8565b5050505050505050505050505050565b5f6116606134bf565b505f51602061561f5f395f51905f5290565b61167a614330565b5f838360405160200161168e9291906152d1565b60408051601f1981840301815282825280516020918201205f8181526004835283902061010085019093528254845260018301805491955091840191906116d4906152ed565b80601f0160208091040260200160405190810160405280929190818152602001828054611700906152ed565b801561174b5780601f106117225761010080835404028352916020019161174b565b820191905f5260205f20905b81548152906001019060200180831161172e57829003601f168201915b505050505081526020016002820154815260200160038201805461176e906152ed565b80601f016020809104026020016040519081016040528092919081815260200182805461179a906152ed565b80156117e55780601f106117bc576101008083540402835291602001916117e5565b820191905f5260205f20905b8154815290600101906020018083116117c857829003601f168201915b50505050508152602001600482015481526020016005820154815260200160068201548152602001600782016040518060600160405290815f8201805480602002602001604051908101604052809291908181526020015f905b828210156118e7578382905f5260205f2001805461185c906152ed565b80601f0160208091040260200160405190810160405280929190818152602001828054611888906152ed565b80156118d35780601f106118aa576101008083540402835291602001916118d3565b820191905f5260205f20905b8154815290600101906020018083116118b657829003601f168201915b50505050508152602001906001019061183f565b5050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561193c57602002820191905f5260205f20905b815481526020019060010190808311611928575b50505050508152602001600282018054806020026020016040519081016040528092919081815260200182805480156119bd57602002820191905f5260205f20905f905b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116119805790505b50505091909252505050905250949350505050565b60606109b18787808060200260200160405190810160405280939291908181526020018383602002808284375f92019190915250899250889150879050866001612fa7565b5f8281526002602052604081205483908203611a465760405163938a92b760e01b815260040160405180910390fd5b5f818152600260205260409020600301546001600160a01b03163314611a8c57611a6f816132c5565b611a8c5760405163dc64d0ad60e01b815260040160405180910390fd5b611a968484613508565b949350505050565b5f60608082808083815f5160206155705f395f51905f528054909150158015611ac957506001810154155b611b125760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064015b60405180910390fd5b611b1a613772565b611b22613832565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b5f8581526002602052604081205486908203611b7b5760405163938a92b760e01b815260040160405180910390fd5b5f818152600260205260409020600301546001600160a01b03163314611bc157611ba4816132c5565b611bc15760405163dc64d0ad60e01b815260040160405180910390fd5b86865f8282604051602001611bd79291906152d1565b60405160208183030381529060405280519060200120905060045f8281526020019081526020015f20600401545f14611c235760405163d96b03b160e01b815260040160405180910390fd5b5f8a8a604051602001611c379291906152d1565b60408051601f1981840301815291815281516020928301205f81815260058452828120546004909452919091206007015490925061ffff90911614611c8f57604051632e1b8f4960e11b815260040160405180910390fd5b885f03611caf57604051631b6fdfeb60e01b815260040160405180910390fd5b86515f03611cd057604051637f19edc960e11b815260040160405180910390fd5b5f805b5f83815260046020526040902060080154811015611d2b575f838152600460205260409020600801805482908110611d0d57611d0d61529e565b905f5260205f20015482611d219190615496565b9150600101611cd3565b50898114611d4c57604051631b6fdfeb60e01b815260040160405180910390fd5b5f828152600460205260409020600101611d668982615370565b505f8281526004602081815260408084209283018e905560069092018c90556007905290205460ff16611e4c575f828152600760208181526040808420805460ff1916600117905560049091528220015490611dc382600a6154a9565b611dd590670de0b6b3a76400006154a9565b90505f82118015611de557505f81115b15611e49575f546040516340c10f1960e01b8152336004820152602481018390526001600160a01b03909116906340c10f19906044015f604051808303815f87803b158015611e32575f5ffd5b505af1158015611e44573d5f5f3e3d5ffd5b505050505b50505b8a604051611e5a91906152e2565b604051908190038120338252908d9084907fc208b556b9717832827b84e296c7f4bb33292ba520fed2eae2f7c0d61f1c0a619060200160405180910390a4509a9950505050505050505050565b5f611eb0613870565b805490915060ff600160401b82041615906001600160401b03165f81158015611ed65750825b90505f826001600160401b03166001148015611ef15750303b155b905081158015611eff575080155b15611f1d5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315611f4757845460ff60401b1916600160401b1785555b611f896040518060400160405280600781526020016653746f7261676560c81b815250604051806040016040528060018152602001603160f81b815250613898565b611f916138aa565b5f80546001600160a01b0319166001600160a01b0388161790558315611ff157845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b612001614330565b5f5f5f85876040516020016120179291906152d1565b6040516020818303038152906040528051906020012090506120398688611672565b9350612046888287611337565b9093509150505b9450945094915050565b5f83815260026020526040812054849082036120865760405163938a92b760e01b815260040160405180910390fd5b5f818152600260205260409020600301546001600160a01b031633146120cc576120af816132c5565b6120cc5760405163dc64d0ad60e01b815260040160405180910390fd5b5f8681526004602052604081205490036120f957604051632abde33960e01b815260040160405180910390fd5b848460405160200161210c9291906152d1565b60405160208183030381529060405280519060200120861461214157604051630ef4797b60e31b815260040160405180910390fd5b5f85815260026020526040902060040180548410158061217b57508681858154811061216f5761216f61529e565b905f5260205f20015414155b15612199576040516337c7f25560e01b815260040160405180910390fd5b5f878152600660209081526040808320805461ffff1916905560049091528120818155906121ca6001830182614391565b600282015f9055600382015f6121e09190614391565b5f600483018190556005830181905560068301819055600783019061220582826143c8565b612212600183015f6143e3565b61221f600283015f6143fe565b5050505f888152600560205260409020805461ffff1916905550805481906122499060019061543e565b815481106122595761225961529e565b905f5260205f2001548185815481106122745761227461529e565b905f5260205f20018190555080805480612290576122906154c0565b600190038181905f5260205f20015f90559055846040516122b191906152e2565b60405190819003812033825290879089907f0e1bf50f5cca6659c62146db5b60160121a3752011b621d8c8953a1e0e23d0489060200160405180910390a460019250505b50949350505050565b5f8133604051602001612312929190615272565b60408051601f1981840301815291815281516020928301205f818152600290935291205490915015612357576040516324bf796160e11b815260040160405180910390fd5b6040805160a0810182528281526020808201858152428385015233606084015283515f80825281840186526080850191909152858152600290925292902081518155915190919060018201906123ad9082615370565b506040820151600282015560608201516003820180546001600160a01b0319166001600160a01b03909216919091179055608082015180516123f9916004840191602090910190614420565b5050335f818152600360209081526040808320805460018101825590845291909220018490555190915061242e9084906152e2565b6040519081900381209083907fb8be57bce74a717a1bbd4acf428df655720fce75c1854b02a88484388df241a8905f90a4919050565b5f60208360ff16111561248a576040516359b452ef60e01b815260040160405180910390fd5b5f8481526004602052604090206009018054600160ff86161b9190849081106124b5576124b561529e565b5f918252602090912060088204015460046007909216919091026101000a90041663ffffffff16151590505b9392505050565b6060602084111561250c57604051637ee7f45f60e11b815260040160405180910390fd5b836001600160401b0381111561252457612524614746565b60405190808252806020026020018201604052801561254d578160200160208202803683370190505b5090505f5b60ff81168511156122f5575f84848389898660ff168181106125765761257661529e565b905060200201356040516020016125b39493929190938452602084019290925260f81b6001600160f81b0319166040830152604182015260610190565b60408051601f1981840301815291815281516020928301205f81815260089093529120549091506125f757604051633e80ff7760e01b815260040160405180910390fd5b5f818152600860205260409020548351849060ff851690811061261c5761261c61529e565b60209081029190910101525080612632816154d4565b915050612552565b6126426142fb565b5f8585604051602001612656929190615272565b60408051601f198184030181528282528051602091820120600180855284840190935293505f929190808301908036833701905050905081815f815181106126a0576126a061529e565b6020026020010181815250505f6126bc825f5f89896001612fa7565b9050805f815181106126d0576126d061529e565b60200260200101519350505050949350505050565b6126ed614330565b60045f8381526020019081526020015f20604051806101000160405290815f8201548152602001600182018054612723906152ed565b80601f016020809104026020016040519081016040528092919081815260200182805461274f906152ed565b801561279a5780601f106127715761010080835404028352916020019161279a565b820191905f5260205f20905b81548152906001019060200180831161277d57829003601f168201915b50505050508152602001600282015481526020016003820180546127bd906152ed565b80601f01602080910402602001604051908101604052809291908181526020018280546127e9906152ed565b80156128345780601f1061280b57610100808354040283529160200191612834565b820191905f5260205f20905b81548152906001019060200180831161281757829003601f168201915b50505050508152602001600482015481526020016005820154815260200160068201548152602001600782016040518060600160405290815f8201805480602002602001604051908101604052809291908181526020015f905b82821015612936578382905f5260205f200180546128ab906152ed565b80601f01602080910402602001604051908101604052809291908181526020018280546128d7906152ed565b80156129225780601f106128f957610100808354040283529160200191612922565b820191905f5260205f20905b81548152906001019060200180831161290557829003601f168201915b50505050508152602001906001019061288e565b5050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561298b57602002820191905f5260205f20905b815481526020019060010190808311612977575b5050505050815260200160028201805480602002602001604051908101604052809291908181526020018280548015612a0c57602002820191905f5260205f20905f905b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116129cf5790505b5050509190925250505090525092915050565b5f8381526002602052604081205484908203612a4e5760405163938a92b760e01b815260040160405180910390fd5b5f818152600260205260409020600301546001600160a01b03163314612a9457612a77816132c5565b612a945760405163dc64d0ad60e01b815260040160405180910390fd5b8333604051602001612aa7929190615272565b604051602081830303815290604052805190602001208514612adc576040516327a5901560e11b815260040160405180910390fd5b5f8581526002602052604090206004015415612b0a5760405162227f7760ea1b815260040160405180910390fd5b5f85815260026020526040812081815590612b286001830182614391565b5f600283018190556003830180546001600160a01b0319169055612b509060048401906143e3565b5050335f90815260036020526040902080548690829086908110612b7657612b7661529e565b905f5260205f20015414612b9d576040516337c7f25560e01b815260040160405180910390fd5b80548190612bad9060019061543e565b81548110612bbd57612bbd61529e565b905f5260205f200154818581548110612bd857612bd861529e565b905f5260205f20018190555080805480612bf457612bf46154c0565b600190038181905f5260205f20015f90559055336001600160a01b031685604051612c1f91906152e2565b6040519081900381209088907feda2cc8e002ead8000b1e0c1debfc9a88bd7ee6e94b8dc0763db17849fcf0389905f90a450600195945050505050565b5f805b5f83815260046020526040902060070154811015612c9957612c818382611217565b15155f03612c9157505f92915050565b600101612c5f565b50600192915050565b5f8784604051602001612cb69291906152d1565b60408051601f1981840301815291815281516020928301205f818152600490935290822060070180549193509089908110612cf357612cf361529e565b905f5260205f20018054612d06906152ed565b80601f0160208091040260200160405190810160405280929190818152602001828054612d32906152ed565b8015612d7d5780601f10612d5457610100808354040283529160200191612d7d565b820191905f5260205f20905b815481529060010190602001808311612d6057829003601f168201915b5050505f85815260046020526040812054939450929092039150612db6905057604051632abde33960e01b815260040160405180910390fd5b5f82815260046020526040812060080180548a908110612dd857612dd861529e565b905f5260205f200154905080876001612df191906154f2565b60ff161115612e13576040516359b452ef60e01b815260040160405180910390fd5b5f6040518061010001604052808481526020018e81526020018b81526020018960ff1681526020018d81526020018a81526020018681526020018c8152509050612e5d81876138b2565b612e6884898c612464565b15612e8657604051636d680ca160e11b815260040160405180910390fd5b5050612e958a838d8b8a613b24565b50612ea182878a613baf565b612eab8289611217565b15612eed575f828152600560205260408120805460019290612ed290849061ffff1661550b565b92506101000a81548161ffff021916908361ffff1602179055505b5f546040516340c10f1960e01b8152336004820152670de0b6b3a764000060248201526001600160a01b03909116906340c10f19906044015f604051808303815f87803b158015612f3c575f5ffd5b505af1158015612f4e573d5f5f3e3d5ffd5b505050508560ff1688837fdd94ce85349373ab535d0282b48fc5974464abc207369050d660ab5296e96db18e8e604051612f92929190918252602082015260400190565b60405180910390a45050505050505050505050565b60605f86158015612fb6575085155b15612fc357506001612fd2565b612fcf87895188613c49565b90505b806001600160401b03811115612fea57612fea614746565b60405190808252806020026020018201604052801561302357816020015b6130106142fb565b8152602001906001900390816130085790505b5091505f5b818110156132b9575f8961303c838b615496565b8151811061304c5761304c61529e565b602002602001015190505f60025f8381526020019081526020015f206040518060a00160405290815f820154815260200160018201805461308c906152ed565b80601f01602080910402602001604051908101604052809291908181526020018280546130b8906152ed565b80156131035780601f106130da57610100808354040283529160200191613103565b820191905f5260205f20905b8154815290600101906020018083116130e657829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020016004820180548060200260200160405190810160405280929190818152602001828054801561319157602002820191905f5260205f20905b81548152602001906001019080831161317d575b50505050508152505090508515613211576040518060a00160405280825f01518152602001826020015181526020018260400151815260200182606001516001600160a01b031681526020016131ec83608001518b8b613c9b565b8152508584815181106132015761320161529e565b60200260200101819052506132af565b6040518060a00160405280825f01518152602001826020015181526020018260400151815260200182606001516001600160a01b031681526020015f6001600160401b0381111561326457613264614746565b60405190808252806020026020018201604052801561328d578160200160208202803683370190505b508152508584815181106132a3576132a361529e565b60200260200101819052505b5050600101613028565b50509695505050505050565b6001545f906001600160a01b0316156133565760015460405163e124bdd960e01b8152600481018490523360248201526001600160a01b039091169063e124bdd990604401602060405180830381865afa158015613325573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906133499190615525565b1561335657506001919050565b505f919050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806133e357507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166133d75f51602061561f5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b156134015760405163703e46dd60e11b815260040160405180910390fd5b565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561345d575060408051601f3d908101601f1916820190925261345a91810190615544565b60015b61348557604051634c9c8ce360e01b81526001600160a01b0383166004820152602401611b09565b5f51602061561f5f395f51905f5281146134b557604051632a87526960e21b815260048101829052602401611b09565b610c5b8383613d4d565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146134015760405163703e46dd60e11b815260040160405180910390fd5b5f5f838360405160200161351d9291906152d1565b60408051601f1981840301815291815281516020928301205f818152600490935291205490915015613562576040516303448eef60e51b815260040160405180910390fd5b5f84815260026020908152604080832060040180546001810182559084529190922001829055516135949084906152e2565b60405190819003812033825290859083907fb018e47bdb983351e1bee22415a8f41eef5c2bf1c43c6d3d0992e678ae762df59060200160405180910390a4604080515f606082018181526080830190935291819083613603565b60608152602001906001900390816135ee5790505b5081526020015f604051908082528060200260200182016040528015613633578160200160208202803683370190505b5081526020015f604051908082528060200260200182016040528015613663578160200160208202803683370190505b50905260408051610100810182528481528151602081810184525f8083528184019283528385018b9052606084018a9052608084018190524260a085015260c0840181905260e084018690528781526004909152929092208151815591519293509160018201906136d49082615370565b5060408201516002820155606082015160038201906136f39082615370565b506080820151600482015560a0820151600582015560c0820151600682015560e082015180518051600784019161372f91839160200190614469565b5060208281015180516137489260018501920190614420565b50604082015180516137649160028401916020909101906144b9565b509498975050505050505050565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f5160206155705f395f51905f52916137b0906152ed565b80601f01602080910402602001604051908101604052809291908181526020018280546137dc906152ed565b80156138275780601f106137fe57610100808354040283529160200191613827565b820191905f5260205f20905b81548152906001019060200180831161380a57829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f5160206155705f395f51905f52916137b0906152ed565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00611211565b6138a0613da2565b6114cc8282613dc7565b613401613da2565b8160c001514211156138fa5760405162461bcd60e51b815260206004820152601160248201527014da59db985d1d5c9948195e1c1a5c9959607a1b6044820152606401611b09565b60e08201515f90815260026020908152604080832060030154815160c08101909252608f8083526001600160a01b0390911693926155909083013980519060200120845f01518051906020012085602001518660400151876060015188608001518960a001518a60c001518b60e001516040516020016139bf9998979695949392919098895260208901979097526040880195909552606087019390935260ff91909116608086015260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012090505f6139e182613e26565b90505f6139ee8286613e52565b9050836001600160a01b0316816001600160a01b031614613a675760405162461bcd60e51b815260206004820152602d60248201527f496e76616c6964207369676e61747572653a204e6f74207369676e656420627960448201526c10313ab1b5b2ba1037bbb732b960991b6064820152608401611b09565b6001600160a01b0381165f90815260096020908152604080832060e08a01518452825280832060a08a0151845290915290205460ff1615613adf5760405162461bcd60e51b8152602060048201526012602482015271139bdb98d948185b1c9958591e481d5cd95960721b6044820152606401611b09565b6001600160a01b03165f90815260096020908152604080832060e08901518452825280832060a090980151835296905294909420805460ff1916600117905550505050565b60408051602081018690529081018390526001600160f81b031960f883901b166060820152606181018490525f9060810160408051601f1981840301815282825280516020918201205f8181526008909252918120899055909250879183917f9f91e060f6a60bc2a5e695f07d854d6a659eab9f5e1441ac24f6e62614bb8b5b91a395945050505050565b60208260ff1610613bd3576040516359b452ef60e01b815260040160405180910390fd5b5f8381526004602052604090206009018054600160ff85161b919083908110613bfe57613bfe61529e565b905f5260205f2090600891828204019190066004028282829054906101000a900463ffffffff161792506101000a81548163ffffffff021916908363ffffffff160217905550505050565b5f82841115613c6b576040516309605a0160e41b815260040160405180910390fd5b5f613c768386615496565b905083811115613c9257613c8a858561543e565b9150506124e1565b829150506124e1565b60605f613caa84865185613c49565b90505f816001600160401b03811115613cc557613cc5614746565b604051908082528060200260200182016040528015613cee578160200160208202803683370190505b5090505f5b82811015613d435786613d068288615496565b81518110613d1657613d1661529e565b6020026020010151828281518110613d3057613d3061529e565b6020908102919091010152600101613cf3565b5095945050505050565b613d5682613e7a565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115613d9a57610c5b8282613edd565b6114cc613f4f565b613daa613f6e565b61340157604051631afcd79f60e31b815260040160405180910390fd5b613dcf613da2565b5f5160206155705f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102613e088482615370565b5060038101613e178382615370565b505f8082556001909101555050565b5f611211613e32613f87565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f613e608686613f95565b925092509250613e708282613fde565b5090949350505050565b806001600160a01b03163b5f03613eaf57604051634c9c8ce360e01b81526001600160a01b0382166004820152602401611b09565b5f51602061561f5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051613ef991906152e2565b5f60405180830381855af49150503d805f8114613f31576040519150601f19603f3d011682016040523d82523d5f602084013e613f36565b606091505b5091509150613f46858383614096565b95945050505050565b34156134015760405163b398979f60e01b815260040160405180910390fd5b5f613f77613870565b54600160401b900460ff16919050565b5f613f906140f2565b905090565b5f5f5f8351604103613fcc576020840151604085015160608601515f1a613fbe88828585614165565b955095509550505050613fd7565b505081515f91506002905b9250925092565b5f826003811115613ff157613ff161555b565b03613ffa575050565b600182600381111561400e5761400e61555b565b0361402c5760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156140405761404061555b565b036140615760405163fce698f760e01b815260048101829052602401611b09565b60038260038111156140755761407561555b565b036114cc576040516335e2f38360e21b815260048101829052602401611b09565b6060826140ab576140a682614229565b6124e1565b81511580156140c257506001600160a01b0384163b155b156140eb57604051639996b31560e01b81526001600160a01b0385166004820152602401611b09565b50806124e1565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f61411c614251565b6141246142b9565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561419e57505f9150600390508261204d565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156141ef573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661421a57505f92506001915082905061204d565b975f9750879650945050505050565b80511561423857805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f5f5160206155705f395f51905f5281614269613772565b80519091501561428157805160209091012092915050565b81548015614290579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f5160206155705f395f51905f52816142d1613832565b8051909150156142e957805160209091012092915050565b60018201548015614290579392505050565b6040518060a001604052805f8152602001606081526020015f81526020015f6001600160a01b03168152602001606081525090565b6040518061010001604052805f8152602001606081526020015f8152602001606081526020015f81526020015f81526020015f815260200161438c60405180606001604052806060815260200160608152602001606081525090565b905290565b50805461439d906152ed565b5f825580601f106143ac575050565b601f0160209004905f5260205f2090810190610939919061455c565b5080545f8255905f5260205f20908101906109399190614570565b5080545f8255905f5260205f2090810190610939919061455c565b5080545f825560070160089004905f5260205f2090810190610939919061455c565b828054828255905f5260205f20908101928215614459579160200282015b8281111561445957825182559160200191906001019061443e565b5061446592915061455c565b5090565b828054828255905f5260205f209081019282156144ad579160200282015b828111156144ad578251829061449d9082615370565b5091602001919060010190614487565b50614465929150614570565b828054828255905f5260205f2090600701600890048101928215614459579160200282015f5b8382111561452357835183826101000a81548163ffffffff021916908363ffffffff16021790555092602001926004016020816003010492830192600103026144df565b80156145535782816101000a81549063ffffffff0219169055600401602081600301049283019260010302614523565b50506144659291505b5b80821115614465575f815560010161455d565b80821115614465575f6145838282614391565b50600101614570565b5f6020828403121561459c575f5ffd5b5035919050565b5f602082840312156145b3575f5ffd5b81356001600160401b038111156145c8575f5ffd5b820161012081850312156124e1575f5ffd5b80356001600160a01b03811681146145f0575f5ffd5b919050565b5f5f5f5f5f60a08688031215614609575f5ffd5b614612866145da565b97602087013597506040870135966060810135965060800135945092505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b805182525f602082015160a0602085015261467f60a0850182614633565b604084810151908601526060808501516001600160a01b031690860152608080850151868303918701919091528051808352602091820193505f9291909101905b80831015613d4357835182526020820191506020840193506001830192506146c0565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561473a57603f19878603018452614725858351614661565b94506020938401939190910190600101614709565b50929695505050505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b038111828210171561478257614782614746565b604052919050565b5f82601f830112614799575f5ffd5b8135602083015f5f6001600160401b038411156147b8576147b8614746565b50601f8301601f19166020016147cd8161475a565b9150508281528583830111156147e1575f5ffd5b828260208301375f92810160200192909252509392505050565b5f5f6040838503121561480c575f5ffd5b82356001600160401b03811115614821575f5ffd5b61482d8582860161478a565b92505061483c602084016145da565b90509250929050565b5f5f83601f840112614855575f5ffd5b5081356001600160401b0381111561486b575f5ffd5b6020830191508360208260051b8501011115610a47575f5ffd5b5f5f60208385031215614896575f5ffd5b82356001600160401b038111156148ab575f5ffd5b6148b785828601614845565b90969095509350505050565b5f5f5f5f5f5f5f5f5f60e08a8c0312156148db575f5ffd5b89356001600160401b038111156148f0575f5ffd5b6148fc8c828d0161478a565b99505060208a0135975060408a01356001600160401b0381111561491e575f5ffd5b61492a8c828d0161478a565b97505060608a0135955060808a01356001600160401b0381111561494c575f5ffd5b6149588c828d01614845565b90965094505060a08a01356001600160401b03811115614976575f5ffd5b6149828c828d01614845565b9a9d999c50979a9699959894979660c00135949350505050565b5f5f604083850312156149ad575f5ffd5b50508035926020909101359150565b5f5f5f606084860312156149ce575f5ffd5b83356001600160401b038111156149e3575f5ffd5b6149ef8682870161478a565b93505060208401359150614a05604085016145da565b90509250925092565b604081525f614a206040830185614633565b90508260208301529392505050565b5f5f60408385031215614a40575f5ffd5b614a49836145da565b915060208301356001600160401b03811115614a63575f5ffd5b614a6f8582860161478a565b9150509250929050565b5f6001600160401b03821115614a9157614a91614746565b5060051b60200190565b5f82601f830112614aaa575f5ffd5b8135614abd614ab882614a79565b61475a565b8082825260208201915060208360051b860101925085831115614ade575f5ffd5b602085015b83811015613d435780356001600160401b03811115614b00575f5ffd5b614b0f886020838a010161478a565b84525060209283019201614ae3565b5f82601f830112614b2d575f5ffd5b8135614b3b614ab882614a79565b8082825260208201915060208360051b860101925085831115614b5c575f5ffd5b602085015b83811015613d43578035835260209283019201614b61565b5f5f5f5f5f5f5f5f5f60e08a8c031215614b91575f5ffd5b89356001600160401b03811115614ba6575f5ffd5b614bb28c828d01614a9b565b99505060208a0135975060408a01356001600160401b03811115614bd4575f5ffd5b614be08c828d0161478a565b97505060608a01356001600160401b03811115614bfb575f5ffd5b614c078c828d01614b1e565b96505060808a01356001600160401b0381111561494c575f5ffd5b602081525f6124e16020830184614633565b5f5f60408385031215614c45575f5ffd5b8235915060208301356001600160401b03811115614a63575f5ffd5b5f8151808452602084019350602083015f5b82811015614c91578151865260209586019590910190600101614c73565b5093949350505050565b5f8151808452602084019350602083015f5b82811015614c9157815163ffffffff16865260209586019590910190600101614cad565b5f6060830182516060855281815180845260808701915060808160051b88010193506020830192505f5b81811015614d2c57607f19888603018352614d17858551614633565b94506020938401939290920191600101614cfb565b5050505060208301518482036020860152614d478282614c61565b91505060408301518482036040860152613f468282614c9b565b805182525f60208201516101006020850152614d81610100850182614633565b90506040830151604085015260608301518482036060860152614da48282614633565b9150506080830151608085015260a083015160a085015260c083015160c085015260e083015184820360e0860152613f468282614cd1565b602081525f6124e16020830184614d61565b5f5f5f5f5f5f60a08789031215614e03575f5ffd5b86356001600160401b03811115614e18575f5ffd5b614e2489828a01614845565b909a90995060208901359860408101359850606081013597506080013595509350505050565b60ff60f81b8816815260e060208201525f614e6860e0830189614633565b8281036040840152614e7a8189614633565b606084018890526001600160a01b038716608085015260a0840186905283810360c08501529050614eab8185614c61565b9a9950505050505050505050565b5f5f5f5f5f60a08688031215614ecd575f5ffd5b8535945060208601356001600160401b03811115614ee9575f5ffd5b614ef58882890161478a565b945050604086013592506060860135915060808601356001600160401b03811115614f1e575f5ffd5b614f2a8882890161478a565b9150509295509295909350565b5f60208284031215614f47575f5ffd5b6124e1826145da565b5f5f5f5f60808587031215614f63575f5ffd5b84356001600160401b03811115614f78575f5ffd5b614f848782880161478a565b94505060208501356001600160401b03811115614f9f575f5ffd5b614fab8782880161478a565b93505060408501359150614fc1606086016145da565b905092959194509250565b606081525f614fde6060830186614d61565b602083019490945250901515604090910152919050565b5f5f5f5f60808587031215615008575f5ffd5b843593506020850135925060408501356001600160401b0381111561502b575f5ffd5b6150378782880161478a565b949793965093946060013593505050565b5f60208284031215615058575f5ffd5b81356001600160401b0381111561506d575f5ffd5b611a968482850161478a565b803560ff811681146145f0575f5ffd5b5f5f5f6060848603121561509b575f5ffd5b833592506150ab60208501615079565b929592945050506040919091013590565b5f5f5f5f606085870312156150cf575f5ffd5b84356001600160401b038111156150e4575f5ffd5b6150f087828801614845565b90989097506020870135966040013595509350505050565b602080825282518282018190525f918401906040840190835b8181101561513f578351835260209384019390920191600101615121565b509095945050505050565b5f5f5f5f6080858703121561515d575f5ffd5b84356001600160401b03811115615172575f5ffd5b61517e8782880161478a565b94505061518d602086016145da565b93969395505050506040820135916060013590565b602081525f6124e16020830184614661565b5f5f5f606084860312156151c6575f5ffd5b8335925060208401356001600160401b038111156151e2575f5ffd5b6151ee8682870161478a565b93969395505050506040919091013590565b5f60208284031215615210575f5ffd5b6124e182615079565b5f5f8335601e1984360301811261522e575f5ffd5b8301803591506001600160401b03821115615247575f5ffd5b602001915036819003821315610a47575f5ffd5b5f81518060208401855e5f93019283525090919050565b5f61527d828561525b565b60609390931b6bffffffffffffffffffffffff191683525050601401919050565b634e487b7160e01b5f52603260045260245ffd5b5f823561011e198336030181126152c7575f5ffd5b9190910192915050565b8281525f611a96602083018461525b565b5f6124e1828461525b565b600181811c9082168061530157607f821691505b60208210810361531f57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115610c5b57805f5260205f20601f840160051c8101602085101561534a5750805b601f840160051c820191505b81811015615369575f8155600101615356565b5050505050565b81516001600160401b0381111561538957615389614746565b61539d8161539784546152ed565b84615325565b6020601f8211600181146153cf575f83156153b85750848201515b5f19600385901b1c1916600184901b178455615369565b5f84815260208120601f198516915b828110156153fe57878501518255602094850194600190920191016153de565b508482101561541b57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b818103818111156112115761121161542a565b5f5f8335601e19843603018112615466575f5ffd5b8301803591506001600160401b0382111561547f575f5ffd5b6020019150600581901b3603821315610a47575f5ffd5b808201808211156112115761121161542a565b80820281158282048414176112115761121161542a565b634e487b7160e01b5f52603160045260245ffd5b5f60ff821660ff81036154e9576154e961542a565b60010192915050565b60ff81811683821601908111156112115761121161542a565b61ffff81811683821601908111156112115761121161542a565b5f60208284031215615535575f5ffd5b815180151581146124e1575f5ffd5b5f60208284031215615554575f5ffd5b5051919050565b634e487b7160e01b5f52602160045260245ffdfea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10053746f7261676544617461286279746573206368756e6b4349442c6279746573333220626c6f636b4349442c75696e74323536206368756e6b496e6465782c75696e743820626c6f636b496e6465782c62797465733332206e6f646549642c75696e74323536206e6f6e63652c75696e7432353620646561646c696e652c62797465733332206275636b6574496429360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca2646970667358221220fd0b021a43c36bfc0fa782e95d0b01fd9ebb4355633e115c86965be272fc399564736f6c634300081c0033",
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

// GetBlockPeersOfChunk is a free data retrieval call binding the contract method 0xf3b348fc.
//
// Solidity: function getBlockPeersOfChunk(bytes32[] blockCIDs, bytes32 fileID, uint256 chunkIndex) view returns(bytes32[] peers)
func (_Storage *StorageCaller) GetBlockPeersOfChunk(opts *bind.CallOpts, blockCIDs [][32]byte, fileID [32]byte, chunkIndex *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getBlockPeersOfChunk", blockCIDs, fileID, chunkIndex)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBlockPeersOfChunk is a free data retrieval call binding the contract method 0xf3b348fc.
//
// Solidity: function getBlockPeersOfChunk(bytes32[] blockCIDs, bytes32 fileID, uint256 chunkIndex) view returns(bytes32[] peers)
func (_Storage *StorageSession) GetBlockPeersOfChunk(blockCIDs [][32]byte, fileID [32]byte, chunkIndex *big.Int) ([][32]byte, error) {
	return _Storage.Contract.GetBlockPeersOfChunk(&_Storage.CallOpts, blockCIDs, fileID, chunkIndex)
}

// GetBlockPeersOfChunk is a free data retrieval call binding the contract method 0xf3b348fc.
//
// Solidity: function getBlockPeersOfChunk(bytes32[] blockCIDs, bytes32 fileID, uint256 chunkIndex) view returns(bytes32[] peers)
func (_Storage *StorageCallerSession) GetBlockPeersOfChunk(blockCIDs [][32]byte, fileID [32]byte, chunkIndex *big.Int) ([][32]byte, error) {
	return _Storage.Contract.GetBlockPeersOfChunk(&_Storage.CallOpts, blockCIDs, fileID, chunkIndex)
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
// Solidity: function getFileById(bytes32 id) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[],uint32[])) file)
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
// Solidity: function getFileById(bytes32 id) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[],uint32[])) file)
func (_Storage *StorageSession) GetFileById(id [32]byte) (IStorageFile, error) {
	return _Storage.Contract.GetFileById(&_Storage.CallOpts, id)
}

// GetFileById is a free data retrieval call binding the contract method 0xfaec0542.
//
// Solidity: function getFileById(bytes32 id) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[],uint32[])) file)
func (_Storage *StorageCallerSession) GetFileById(id [32]byte) (IStorageFile, error) {
	return _Storage.Contract.GetFileById(&_Storage.CallOpts, id)
}

// GetFileByName is a free data retrieval call binding the contract method 0x5ecdfb53.
//
// Solidity: function getFileByName(bytes32 bucketId, string fileName) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[],uint32[])) file)
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
// Solidity: function getFileByName(bytes32 bucketId, string fileName) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[],uint32[])) file)
func (_Storage *StorageSession) GetFileByName(bucketId [32]byte, fileName string) (IStorageFile, error) {
	return _Storage.Contract.GetFileByName(&_Storage.CallOpts, bucketId, fileName)
}

// GetFileByName is a free data retrieval call binding the contract method 0x5ecdfb53.
//
// Solidity: function getFileByName(bytes32 bucketId, string fileName) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[],uint32[])) file)
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
// Solidity: function getFullFileInfo(string bucketName, string fileName, bytes32 bucketId, address owner) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[],uint32[])) file, uint256 index, bool exists)
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
// Solidity: function getFullFileInfo(string bucketName, string fileName, bytes32 bucketId, address owner) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[],uint32[])) file, uint256 index, bool exists)
func (_Storage *StorageSession) GetFullFileInfo(bucketName string, fileName string, bucketId [32]byte, owner common.Address) (struct {
	File   IStorageFile
	Index  *big.Int
	Exists bool
}, error) {
	return _Storage.Contract.GetFullFileInfo(&_Storage.CallOpts, bucketName, fileName, bucketId, owner)
}

// GetFullFileInfo is a free data retrieval call binding the contract method 0xd4967ce6.
//
// Solidity: function getFullFileInfo(string bucketName, string fileName, bytes32 bucketId, address owner) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[],uint32[])) file, uint256 index, bool exists)
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

// GetPeerByPeerBlock is a free data retrieval call binding the contract method 0x3d840ee9.
//
// Solidity: function getPeerByPeerBlock(bytes32 peerBlockID) view returns(bytes32 peer)
func (_Storage *StorageCaller) GetPeerByPeerBlock(opts *bind.CallOpts, peerBlockID [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getPeerByPeerBlock", peerBlockID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPeerByPeerBlock is a free data retrieval call binding the contract method 0x3d840ee9.
//
// Solidity: function getPeerByPeerBlock(bytes32 peerBlockID) view returns(bytes32 peer)
func (_Storage *StorageSession) GetPeerByPeerBlock(peerBlockID [32]byte) ([32]byte, error) {
	return _Storage.Contract.GetPeerByPeerBlock(&_Storage.CallOpts, peerBlockID)
}

// GetPeerByPeerBlock is a free data retrieval call binding the contract method 0x3d840ee9.
//
// Solidity: function getPeerByPeerBlock(bytes32 peerBlockID) view returns(bytes32 peer)
func (_Storage *StorageCallerSession) GetPeerByPeerBlock(peerBlockID [32]byte) ([32]byte, error) {
	return _Storage.Contract.GetPeerByPeerBlock(&_Storage.CallOpts, peerBlockID)
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
// Solidity: function addFileChunk(bytes chunkCID, bytes32 bucketId, string fileName, uint256 encodedChunkSize, bytes32[] cids, uint256[] chunkBlocksSizes, uint256 chunkIndex) returns(bytes32)
func (_Storage *StorageTransactor) AddFileChunk(opts *bind.TransactOpts, chunkCID []byte, bucketId [32]byte, fileName string, encodedChunkSize *big.Int, cids [][32]byte, chunkBlocksSizes []*big.Int, chunkIndex *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addFileChunk", chunkCID, bucketId, fileName, encodedChunkSize, cids, chunkBlocksSizes, chunkIndex)
}

// AddFileChunk is a paid mutator transaction binding the contract method 0x30b91d07.
//
// Solidity: function addFileChunk(bytes chunkCID, bytes32 bucketId, string fileName, uint256 encodedChunkSize, bytes32[] cids, uint256[] chunkBlocksSizes, uint256 chunkIndex) returns(bytes32)
func (_Storage *StorageSession) AddFileChunk(chunkCID []byte, bucketId [32]byte, fileName string, encodedChunkSize *big.Int, cids [][32]byte, chunkBlocksSizes []*big.Int, chunkIndex *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddFileChunk(&_Storage.TransactOpts, chunkCID, bucketId, fileName, encodedChunkSize, cids, chunkBlocksSizes, chunkIndex)
}

// AddFileChunk is a paid mutator transaction binding the contract method 0x30b91d07.
//
// Solidity: function addFileChunk(bytes chunkCID, bytes32 bucketId, string fileName, uint256 encodedChunkSize, bytes32[] cids, uint256[] chunkBlocksSizes, uint256 chunkIndex) returns(bytes32)
func (_Storage *StorageTransactorSession) AddFileChunk(chunkCID []byte, bucketId [32]byte, fileName string, encodedChunkSize *big.Int, cids [][32]byte, chunkBlocksSizes []*big.Int, chunkIndex *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddFileChunk(&_Storage.TransactOpts, chunkCID, bucketId, fileName, encodedChunkSize, cids, chunkBlocksSizes, chunkIndex)
}

// AddFileChunks is a paid mutator transaction binding the contract method 0x51dc0c23.
//
// Solidity: function addFileChunks(bytes[] cids, bytes32 bucketId, string fileName, uint256[] encodedChunkSizes, bytes32[][] chunkBlocksCIDs, uint256[][] chunkBlockSizes, uint256 startingChunkIndex) returns()
func (_Storage *StorageTransactor) AddFileChunks(opts *bind.TransactOpts, cids [][]byte, bucketId [32]byte, fileName string, encodedChunkSizes []*big.Int, chunkBlocksCIDs [][][32]byte, chunkBlockSizes [][]*big.Int, startingChunkIndex *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addFileChunks", cids, bucketId, fileName, encodedChunkSizes, chunkBlocksCIDs, chunkBlockSizes, startingChunkIndex)
}

// AddFileChunks is a paid mutator transaction binding the contract method 0x51dc0c23.
//
// Solidity: function addFileChunks(bytes[] cids, bytes32 bucketId, string fileName, uint256[] encodedChunkSizes, bytes32[][] chunkBlocksCIDs, uint256[][] chunkBlockSizes, uint256 startingChunkIndex) returns()
func (_Storage *StorageSession) AddFileChunks(cids [][]byte, bucketId [32]byte, fileName string, encodedChunkSizes []*big.Int, chunkBlocksCIDs [][][32]byte, chunkBlockSizes [][]*big.Int, startingChunkIndex *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddFileChunks(&_Storage.TransactOpts, cids, bucketId, fileName, encodedChunkSizes, chunkBlocksCIDs, chunkBlockSizes, startingChunkIndex)
}

// AddFileChunks is a paid mutator transaction binding the contract method 0x51dc0c23.
//
// Solidity: function addFileChunks(bytes[] cids, bytes32 bucketId, string fileName, uint256[] encodedChunkSizes, bytes32[][] chunkBlocksCIDs, uint256[][] chunkBlockSizes, uint256 startingChunkIndex) returns()
func (_Storage *StorageTransactorSession) AddFileChunks(cids [][]byte, bucketId [32]byte, fileName string, encodedChunkSizes []*big.Int, chunkBlocksCIDs [][][32]byte, chunkBlockSizes [][]*big.Int, startingChunkIndex *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddFileChunks(&_Storage.TransactOpts, cids, bucketId, fileName, encodedChunkSizes, chunkBlocksCIDs, chunkBlockSizes, startingChunkIndex)
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

// FillChunkBlock is a paid mutator transaction binding the contract method 0x1ed572b7.
//
// Solidity: function fillChunkBlock((bytes32,bytes32,bytes32,uint256,uint256,uint8,string,bytes,uint256) args) returns()
func (_Storage *StorageTransactor) FillChunkBlock(opts *bind.TransactOpts, args IStorageFillChunkBlockArgs) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "fillChunkBlock", args)
}

// FillChunkBlock is a paid mutator transaction binding the contract method 0x1ed572b7.
//
// Solidity: function fillChunkBlock((bytes32,bytes32,bytes32,uint256,uint256,uint8,string,bytes,uint256) args) returns()
func (_Storage *StorageSession) FillChunkBlock(args IStorageFillChunkBlockArgs) (*types.Transaction, error) {
	return _Storage.Contract.FillChunkBlock(&_Storage.TransactOpts, args)
}

// FillChunkBlock is a paid mutator transaction binding the contract method 0x1ed572b7.
//
// Solidity: function fillChunkBlock((bytes32,bytes32,bytes32,uint256,uint256,uint8,string,bytes,uint256) args) returns()
func (_Storage *StorageTransactorSession) FillChunkBlock(args IStorageFillChunkBlockArgs) (*types.Transaction, error) {
	return _Storage.Contract.FillChunkBlock(&_Storage.TransactOpts, args)
}

// FillChunkBlocks is a paid mutator transaction binding the contract method 0x308c3cf2.
//
// Solidity: function fillChunkBlocks((bytes32,bytes32,bytes32,uint256,uint256,uint8,string,bytes,uint256)[] args) returns()
func (_Storage *StorageTransactor) FillChunkBlocks(opts *bind.TransactOpts, args []IStorageFillChunkBlockArgs) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "fillChunkBlocks", args)
}

// FillChunkBlocks is a paid mutator transaction binding the contract method 0x308c3cf2.
//
// Solidity: function fillChunkBlocks((bytes32,bytes32,bytes32,uint256,uint256,uint8,string,bytes,uint256)[] args) returns()
func (_Storage *StorageSession) FillChunkBlocks(args []IStorageFillChunkBlockArgs) (*types.Transaction, error) {
	return _Storage.Contract.FillChunkBlocks(&_Storage.TransactOpts, args)
}

// FillChunkBlocks is a paid mutator transaction binding the contract method 0x308c3cf2.
//
// Solidity: function fillChunkBlocks((bytes32,bytes32,bytes32,uint256,uint256,uint8,string,bytes,uint256)[] args) returns()
func (_Storage *StorageTransactorSession) FillChunkBlocks(args []IStorageFillChunkBlockArgs) (*types.Transaction, error) {
	return _Storage.Contract.FillChunkBlocks(&_Storage.TransactOpts, args)
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
	FileId     [32]byte
	ChunkIndex *big.Int
	BlockIndex *big.Int
	BlockCID   [32]byte
	NodeId     [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFillChunkBlock is a free log retrieval operation binding the contract event 0xdd94ce85349373ab535d0282b48fc5974464abc207369050d660ab5296e96db1.
//
// Solidity: event FillChunkBlock(bytes32 indexed fileId, uint256 indexed chunkIndex, uint256 indexed blockIndex, bytes32 blockCID, bytes32 nodeId)
func (_Storage *StorageFilterer) FilterFillChunkBlock(opts *bind.FilterOpts, fileId [][32]byte, chunkIndex []*big.Int, blockIndex []*big.Int) (*StorageFillChunkBlockIterator, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}
	var chunkIndexRule []interface{}
	for _, chunkIndexItem := range chunkIndex {
		chunkIndexRule = append(chunkIndexRule, chunkIndexItem)
	}
	var blockIndexRule []interface{}
	for _, blockIndexItem := range blockIndex {
		blockIndexRule = append(blockIndexRule, blockIndexItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "FillChunkBlock", fileIdRule, chunkIndexRule, blockIndexRule)
	if err != nil {
		return nil, err
	}
	return &StorageFillChunkBlockIterator{contract: _Storage.contract, event: "FillChunkBlock", logs: logs, sub: sub}, nil
}

// WatchFillChunkBlock is a free log subscription operation binding the contract event 0xdd94ce85349373ab535d0282b48fc5974464abc207369050d660ab5296e96db1.
//
// Solidity: event FillChunkBlock(bytes32 indexed fileId, uint256 indexed chunkIndex, uint256 indexed blockIndex, bytes32 blockCID, bytes32 nodeId)
func (_Storage *StorageFilterer) WatchFillChunkBlock(opts *bind.WatchOpts, sink chan<- *StorageFillChunkBlock, fileId [][32]byte, chunkIndex []*big.Int, blockIndex []*big.Int) (event.Subscription, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}
	var chunkIndexRule []interface{}
	for _, chunkIndexItem := range chunkIndex {
		chunkIndexRule = append(chunkIndexRule, chunkIndexItem)
	}
	var blockIndexRule []interface{}
	for _, blockIndexItem := range blockIndex {
		blockIndexRule = append(blockIndexRule, blockIndexItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "FillChunkBlock", fileIdRule, chunkIndexRule, blockIndexRule)
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

// ParseFillChunkBlock is a log parse operation binding the contract event 0xdd94ce85349373ab535d0282b48fc5974464abc207369050d660ab5296e96db1.
//
// Solidity: event FillChunkBlock(bytes32 indexed fileId, uint256 indexed chunkIndex, uint256 indexed blockIndex, bytes32 blockCID, bytes32 nodeId)
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
