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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockAlreadyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockNonexists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketInvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketNonempty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketNonexists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"}],\"name\":\"ChunkCIDMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileChunkDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileFullyUploaded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNameDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNonempty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNotFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cidsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sizesLength\",\"type\":\"uint256\"}],\"name\":\"InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlocksAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEncodedSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFileBlocksCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFileCID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLastBlockSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPeerIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LastChunkDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPeersForCID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEligibleToUpgrade\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OffsetOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyBlockCIDs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongAuthority\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"AddFileBlocks\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AddFileChunk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"}],\"name\":\"AddPeerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CommitFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CreateBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CreateFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DeleteBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DeleteFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"}],\"name\":\"DeletePeerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blockCID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"}],\"name\":\"FillChunkBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_BLOCKS_PER_FILE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BLOCK_SIZE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessManager\",\"outputs\":[{\"internalType\":\"contractIAccessManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"chunkCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedChunkSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"cids\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkBlocksSizes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"addFileChunk\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"cids\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"encodedChunkSizes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"chunkBlocksCIDs\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"chunkBlockSizes\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256\",\"name\":\"startingChunkIndex\",\"type\":\"uint256\"}],\"name\":\"addFileChunks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedFileSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"}],\"name\":\"commitFile\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"createBucket\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"createFile\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deleteBucket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deleteFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fileFillCounter\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fileRewardClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockCID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIStorage.FillChunkBlockArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"fillChunkBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockCID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIStorage.FillChunkBlockArgs[]\",\"name\":\"args\",\"type\":\"tuple[]\"}],\"name\":\"fillChunkBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"blockCIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"fileID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"getBlockPeersOfChunk\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"peers\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fileOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileLimit\",\"type\":\"uint256\"}],\"name\":\"getBucketByName\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket\",\"name\":\"bucket\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getBucketIndexByName\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"}],\"name\":\"getBucketsByIds\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"bucketOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bucketLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileLimit\",\"type\":\"uint256\"}],\"name\":\"getBucketsByIdsWithFiles\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket[]\",\"name\":\"buckets\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getChunkByIndex\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getFileById\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"chunkCIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkSize\",\"type\":\"uint256[]\"},{\"internalType\":\"uint32[]\",\"name\":\"fulfilledBlocks\",\"type\":\"uint32[]\"}],\"internalType\":\"structIStorage.Chunk\",\"name\":\"chunks\",\"type\":\"tuple\"}],\"internalType\":\"structIStorage.File\",\"name\":\"file\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"getFileByName\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"chunkCIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkSize\",\"type\":\"uint256[]\"},{\"internalType\":\"uint32[]\",\"name\":\"fulfilledBlocks\",\"type\":\"uint32[]\"}],\"internalType\":\"structIStorage.Chunk\",\"name\":\"chunks\",\"type\":\"tuple\"}],\"internalType\":\"structIStorage.File\",\"name\":\"file\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getFileIndexById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getFileOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getFullFileInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"chunkCIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkSize\",\"type\":\"uint256[]\"},{\"internalType\":\"uint32[]\",\"name\":\"fulfilledBlocks\",\"type\":\"uint32[]\"}],\"internalType\":\"structIStorage.Chunk\",\"name\":\"chunks\",\"type\":\"tuple\"}],\"internalType\":\"structIStorage.File\",\"name\":\"file\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileLimit\",\"type\":\"uint256\"}],\"name\":\"getOwnerBuckets\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket[]\",\"name\":\"buckets\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"peerBlockID\",\"type\":\"bytes32\"}],\"name\":\"getPeerByPeerBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"peer\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"isBlockFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"isChunkFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"isFileFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"isFileFilledV2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accessManagerAddress\",\"type\":\"address\"}],\"name\":\"setAccessManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_upgradeAuthority\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIAkaveToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516155c06100f95f395f818161305c01528181613085015261321301526155c05ff3fe608060405260043610610249575f3560e01c80636ce0236311610134578063d6d3110b116100b3578063f8a3e41a11610078578063f8a3e41a146107d2578063faec054214610814578063fc0c546a14610833578063fd1d3c0c14610851578063fd21c28414610870578063fdcb60681461088f575f5ffd5b8063d6d3110b1461071d578063e3f787e81461073c578063e4ba8a581461075b578063f3b348fc1461077a578063f83533d2146107a6575f5ffd5b8063ad3cb1cc116100f9578063ad3cb1cc14610652578063b80777ea14610682578063c4d66de814610694578063c9580804146106b3578063d4967ce6146106ef575f5ffd5b80636ce02363146105aa5780637a9e5e4b146105d857806384b0196e146105f75780639a2e82b31461061e5780639ccd46461461063d575f5ffd5b806349cdf6ac116101cb57806354fd4d501161019057806354fd4d50146104b7578063564b81ef146104ed5780635ecdfb53146104ff57806362838c231461052b57806368e6408f1461054a5780636af0f8011461058b575f5ffd5b806349cdf6ac146104255780634d7dc614146104445780634f1ef2861461047157806351dc0c231461048457806352d1902d146104a3575f5ffd5b8063308c3cf211610211578063308c3cf21461037057806330b91d071461038f57806335bdb711146103bc5780633d840ee9146103db5780633f38398014610406575f5ffd5b8063018c1e9c1461024d5780631b475ef4146102905780631ed572b7146102ef578063262a61ce14610310578063287e677f1461033c575b5f5ffd5b348015610258575f5ffd5b5061027b6102673660046144ce565b60086020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b34801561029b575f5ffd5b506102d76102aa3660046144ce565b5f90815260056020908152604080832060020154835260039182905290912001546001600160a01b031690565b6040516001600160a01b039091168152602001610287565b3480156102fa575f5ffd5b5061030e6103093660046144e5565b6108ae565b005b34801561031b575f5ffd5b5061032f61032a366004614537565b610966565b604051610287919061462f565b348015610347575f5ffd5b5061035b610356366004614747565b610998565b60408051928352901515602083015201610287565b34801561037b575f5ffd5b5061030e61038a3660046147d1565b610a2a565b34801561039a575f5ffd5b506103ae6103a936600461480f565b610c3c565b604051908152602001610287565b3480156103c7575f5ffd5b5061032f6103d63660046147d1565b610f81565b3480156103e6575f5ffd5b506103ae6103f53660046144ce565b5f9081526009602052604090205490565b348015610411575f5ffd5b5061027b6104203660046148e8565b61112e565b348015610430575f5ffd5b5061035b61043f366004614908565b61124e565b34801561044f575f5ffd5b5061046361045e3660046148e8565b6112d9565b60405161028792919061495a565b61030e61047f36600461497b565b6113d1565b34801561048f575f5ffd5b5061030e61049e366004614ac5565b6113f0565b3480156104ae575f5ffd5b506103ae611574565b3480156104c2575f5ffd5b506040805180820190915260058152640312e302e360dc1b60208201525b6040516102879190614b6e565b3480156104f8575f5ffd5b50466103ae565b34801561050a575f5ffd5b5061051e610519366004614b80565b61158f565b6040516102879190614d28565b348015610536575f5ffd5b5061032f610545366004614d3a565b6118ef565b348015610555575f5ffd5b5061027b6105643660046144ce565b5f9081526006602090815260408083205460059092529091206007015461ffff9091161490565b348015610596575f5ffd5b506103ae6105a5366004614b80565b611932565b3480156105b5575f5ffd5b506105c0620f424081565b6040516001600160401b039091168152602001610287565b3480156105e3575f5ffd5b5061030e6105f2366004614d96565b6119ba565b348015610602575f5ffd5b5061060b611a1f565b6040516102879796959493929190614daf565b348015610629575f5ffd5b506103ae610638366004614e1e565b611acd565b348015610648575f5ffd5b506105c061040081565b34801561065d575f5ffd5b506104e0604051806040016040528060058152602001640352e302e360dc1b81525081565b34801561068d575f5ffd5b50426103ae565b34801561069f575f5ffd5b5061030e6106ae366004614d96565b611e27565b3480156106be575f5ffd5b5061030e6106cd366004614d96565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b3480156106fa575f5ffd5b5061070e610709366004614e9c565b611f82565b60405161028793929190614f18565b348015610728575f5ffd5b5061027b610737366004614f41565b611fe0565b348015610747575f5ffd5b506103ae610756366004614f94565b612288565b348015610766575f5ffd5b5061027b610775366004614fd5565b6123ee565b348015610785575f5ffd5b50610799610794366004615008565b612472565b6040516102879190615054565b3480156107b1575f5ffd5b506107c56107c0366004615096565b6125c4565b60405161028791906150ee565b3480156107dd575f5ffd5b506108016107ec3660046144ce565b60066020525f908152604090205461ffff1681565b60405161ffff9091168152602001610287565b34801561081f575f5ffd5b5061051e61082e3660046144ce565b61260d565b34801561083e575f5ffd5b505f546102d7906001600160a01b031681565b34801561085c575f5ffd5b5061027b61086b366004615100565b612947565b34801561087b575f5ffd5b5061027b61088a3660046144ce565b612b85565b34801561089a575f5ffd5b506001546102d7906001600160a01b031681565b610963813560208301356040840135606085013560808601356108d760c0880160a0890161514c565b6108e460c0890189615165565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152506109259250505060e08a018a615165565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152505050506101008a0135612bcb565b50565b6001600160a01b0385165f90815260046020526040902060609061098d8187878787612ed0565b979650505050505050565b5f5f5f84846040516020016109ae9291906151be565b60408051601f1981840301815291815281516020928301206001600160a01b0387165f9081526004909352908220909250905b8154811015610a1f57828282815481106109fd576109fd6151ea565b905f5260205f20015403610a175780945060019350610a1f565b6001016109e1565b5050505b9250929050565b5f5b81811015610c3757610c2f838383818110610a4957610a496151ea565b9050602002810190610a5b91906151fe565b35848484818110610a6e57610a6e6151ea565b9050602002810190610a8091906151fe565b60200135858585818110610a9657610a966151ea565b9050602002810190610aa891906151fe565b60400135868686818110610abe57610abe6151ea565b9050602002810190610ad091906151fe565b60600135878787818110610ae657610ae66151ea565b9050602002810190610af891906151fe565b60800135888888818110610b0e57610b0e6151ea565b9050602002810190610b2091906151fe565b610b319060c081019060a00161514c565b898989818110610b4357610b436151ea565b9050602002810190610b5591906151fe565b610b639060c0810190615165565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508d92508c91508b9050818110610bab57610bab6151ea565b9050602002810190610bbd91906151fe565b610bcb9060e0810190615165565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508e92508d91508c9050818110610c1357610c136151ea565b9050602002810190610c2591906151fe565b6101000135612bcb565b600101610a2c565b505050565b5f8881526003602052604081205489908203610c6b5760405163938a92b760e01b815260040160405180910390fd5b5f81815260036020819052604090912001546001600160a01b03163314610cb257610c9581612fb9565b610cb25760405163dc64d0ad60e01b815260040160405180910390fd5b89895f8282604051602001610cc892919061521d565b60408051601f1981840301815291815281516020928301205f818152600590935291206004015490915015610d105760405163d96b03b160e01b815260040160405180910390fd5b5f8d8d604051602001610d2492919061521d565b60408051601f1981840301815291815281516020928301205f8181526005909352908220549092509003610d6b57604051632abde33960e01b815260040160405180910390fd5b5f818152600560205260409020600701548714610d9b576040516301c0b3dd60e61b815260040160405180910390fd5b8988141580610daa575060208a115b15610dc8576040516373d8ccd360e11b815260040160405180910390fd5b60208a1015610e1e575f8181526007602052604090205461ffff1615610e01576040516355cbc83160e01b815260040160405180910390fd5b5f818152600760205260409020805461ffff191661ffff8c161790555b8b5f03610e3e57604051631b6fdfeb60e01b815260040160405180910390fd5b8c604051610e4c919061522e565b604051908190038120338252908f9083907f0d877d3656ad561e6c8224e31cc98f1b3e62b1d828fd396c47ef0daccb6838e59060200160405180910390a460055f8281526020019081526020015f206007015f018f908060018154018082558091505060019003905f5260205f20015f909190919091509081610ecf91906152bc565b5060055f8281526020019081526020015f206007016001018c908060018154018082558091505060019003905f5260205f20015f909190919091505560055f8281526020019081526020015f206007016002015f908060018154018082558091505060019003905f5260205f2090600891828204019190066004029091909190916101000a81548163ffffffff021916908363ffffffff16021790555080955050505050509998505050505050505050565b60605f826001600160401b03811115610f9c57610f9c614692565b604051908082528060200260200182016040528015610fd557816020015b610fc261423d565b815260200190600190039081610fba5790505b5090505f5b83811015611124575f60035f878785818110610ff857610ff86151ea565b9050602002013581526020019081526020015f2090506040518060a00160405280825f0154815260200182600101805461103190615239565b80601f016020809104026020016040519081016040528092919081815260200182805461105d90615239565b80156110a85780601f1061107f576101008083540402835291602001916110a8565b820191905f5260205f20905b81548152906001019060200180831161108b57829003601f168201915b50505091835250506002830154602082015260038301546001600160a01b031660408201526060015f6040519080825280602002602001820160405280156110fa578160200160208202803683370190505b50815250838381518110611110576111106151ea565b602090810291909101015250600101610fda565b5090505b92915050565b5f82815260056020526040812060070154819061114d9060019061538a565b9050808314801561116e57505f8481526007602052604090205461ffff1615155b156111f9575f848152600760205260408120546111949060019061ffff1681901b61538a565b5f868152600560205260409020600901805491925063ffffffff831691839190879081106111c4576111c46151ea565b905f5260205f2090600891828204019190066004029054906101000a900463ffffffff161663ffffffff161492505050611128565b5f848152600560205260409020600901805463ffffffff919085908110611222576112226151ea565b5f918252602090912060088204015460079091166004026101000a900463ffffffff1614949350505050565b5f5f5f85846040516020016112649291906151be565b60408051601f1981840301815291815281516020928301205f8181526003909352908220909250600401905b81548110156112ce57868282815481106112ac576112ac6151ea565b905f5260205f200154036112c657809450600193506112ce565b600101611290565b505050935093915050565b5f828152600560205260408120600701805460609291829185908110611301576113016151ea565b905f5260205f2001805461131490615239565b80601f016020809104026020016040519081016040528092919081815260200182805461134090615239565b801561138b5780601f106113625761010080835404028352916020019161138b565b820191905f5260205f20905b81548152906001019060200180831161136e57829003601f168201915b5050505f8881526005602052604081206008018054949550909390925087915081106113b9576113b96151ea565b5f918252602090912001549196919550909350505050565b6113d9613051565b6113e2826130f7565b6113ec828261314c565b5050565b5f8881526003602052604081205489910361141e5760405163938a92b760e01b815260040160405180910390fd5b5f81815260036020819052604090912001546001600160a01b031633146114655761144881612fb9565b6114655760405163dc64d0ad60e01b815260040160405180910390fd5b88885f828260405160200161147b92919061521d565b60408051601f1981840301815291815281516020928301205f8181526005909352912060040154909150156114c35760405163d96b03b160e01b815260040160405180910390fd5b5f5b8d518110156115645761155b8e82815181106114e3576114e36151ea565b60200260200101518e8e8e85815181106114ff576114ff6151ea565b60200260200101518e8e87818110611519576115196151ea565b905060200281019061152b919061539d565b8e8e8981811061153d5761153d6151ea565b905060200281019061154f919061539d565b898f6103a991906153e2565b506001016114c5565b5050505050505050505050505050565b5f61157d613208565b505f51602061556b5f395f51905f5290565b611597614272565b5f83836040516020016115ab92919061521d565b60408051601f1981840301815282825280516020918201205f8181526005835283902061010085019093528254845260018301805491955091840191906115f190615239565b80601f016020809104026020016040519081016040528092919081815260200182805461161d90615239565b80156116685780601f1061163f57610100808354040283529160200191611668565b820191905f5260205f20905b81548152906001019060200180831161164b57829003601f168201915b505050505081526020016002820154815260200160038201805461168b90615239565b80601f01602080910402602001604051908101604052809291908181526020018280546116b790615239565b80156117025780601f106116d957610100808354040283529160200191611702565b820191905f5260205f20905b8154815290600101906020018083116116e557829003601f168201915b50505050508152602001600482015481526020016005820154815260200160068201548152602001600782016040518060600160405290815f8201805480602002602001604051908101604052809291908181526020015f905b82821015611804578382905f5260205f2001805461177990615239565b80601f01602080910402602001604051908101604052809291908181526020018280546117a590615239565b80156117f05780601f106117c7576101008083540402835291602001916117f0565b820191905f5260205f20905b8154815290600101906020018083116117d357829003601f168201915b50505050508152602001906001019061175c565b5050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561185957602002820191905f5260205f20905b815481526020019060010190808311611845575b50505050508152602001600282018054806020026020016040519081016040528092919081815260200182805480156118da57602002820191905f5260205f20905f905b82829054906101000a900463ffffffff1663ffffffff168152602001906004019060208260030104928301926001038202915080841161189d5790505b50505091909252505050905250949350505050565b606061098d8787808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525089925088915087905086613251565b5f82815260036020526040812054839082036119615760405163938a92b760e01b815260040160405180910390fd5b5f81815260036020819052604090912001546001600160a01b031633146119a85761198b81612fb9565b6119a85760405163dc64d0ad60e01b815260040160405180910390fd5b6119b2848461332c565b949350505050565b6002546001600160a01b031633148015906119df57506002546001600160a01b031615155b156119fd57604051631fc0e70760e21b815260040160405180910390fd5b600280546001600160a01b0319166001600160a01b0392909216919091179055565b5f60608082808083815f5160206154bc5f395f51905f528054909150158015611a4a57506001810154155b611a935760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064015b60405180910390fd5b611a9b613596565b611aa3613656565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b5f8581526003602052604081205486908203611afc5760405163938a92b760e01b815260040160405180910390fd5b5f81815260036020819052604090912001546001600160a01b03163314611b4357611b2681612fb9565b611b435760405163dc64d0ad60e01b815260040160405180910390fd5b86865f8282604051602001611b5992919061521d565b60408051601f1981840301815291815281516020928301205f818152600590935291206004015490915015611ba15760405163d96b03b160e01b815260040160405180910390fd5b5f8a8a604051602001611bb592919061521d565b60408051601f1981840301815291815281516020928301205f81815260068452828120546005909452919091206007015490925061ffff90911614611c0d57604051632e1b8f4960e11b815260040160405180910390fd5b885f03611c2d57604051631b6fdfeb60e01b815260040160405180910390fd5b86515f03611c4e57604051637f19edc960e11b815260040160405180910390fd5b5f805b5f83815260056020526040902060080154811015611ca9575f838152600560205260409020600801805482908110611c8b57611c8b6151ea565b905f5260205f20015482611c9f91906153e2565b9150600101611c51565b50898114611cca57604051631b6fdfeb60e01b815260040160405180910390fd5b5f828152600560205260409020600101611ce489826152bc565b505f828152600560209081526040808320600481018e90556006018c9055600890915290205460ff16611dcc575f828152600860209081526040808320805460ff19166001179055600590915281206007015490611d4382600a6153f5565b611d5590670de0b6b3a76400006153f5565b90505f82118015611d6557505f81115b15611dc9575f546040516340c10f1960e01b8152336004820152602481018390526001600160a01b03909116906340c10f19906044015f604051808303815f87803b158015611db2575f5ffd5b505af1158015611dc4573d5f5f3e3d5ffd5b505050505b50505b8a604051611dda919061522e565b604051908190038120338252908d9084907fc208b556b9717832827b84e296c7f4bb33292ba520fed2eae2f7c0d61f1c0a619060200160405180910390a4509a9950505050505050505050565b5f611e30613694565b805490915060ff600160401b82041615906001600160401b03165f81158015611e565750825b90505f826001600160401b03166001148015611e715750303b155b905081158015611e7f575080155b15611e9d5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315611ec757845460ff60401b1916600160401b1785555b611f096040518060400160405280600781526020016653746f7261676560c81b815250604051806040016040528060018152602001603160f81b8152506136bc565b611f116136ce565b611f1a336119ba565b5f80546001600160a01b0319166001600160a01b0388161790558315611f7a57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b611f8a614272565b5f5f5f8587604051602001611fa092919061521d565b604051602081830303815290604052805190602001209050611fc2868861158f565b9350611fcf88828761124e565b9093509150505b9450945094915050565b5f838152600360205260408120548490820361200f5760405163938a92b760e01b815260040160405180910390fd5b5f81815260036020819052604090912001546001600160a01b031633146120565761203981612fb9565b6120565760405163dc64d0ad60e01b815260040160405180910390fd5b5f86815260056020526040812054900361208357604051632abde33960e01b815260040160405180910390fd5b848460405160200161209692919061521d565b6040516020818303038152906040528051906020012086146120cb57604051630ef4797b60e31b815260040160405180910390fd5b5f8581526003602052604090206004018054841015806121055750868185815481106120f9576120f96151ea565b905f5260205f20015414155b15612123576040516337c7f25560e01b815260040160405180910390fd5b5f878152600760209081526040808320805461ffff19169055600590915281208181559061215460018301826142d3565b600282015f9055600382015f61216a91906142d3565b5f600483018190556005830181905560068301819055600783019061218f828261430a565b61219c600183015f614325565b6121a9600283015f614340565b5050505f888152600660205260409020805461ffff1916905550805481906121d39060019061538a565b815481106121e3576121e36151ea565b905f5260205f2001548185815481106121fe576121fe6151ea565b905f5260205f2001819055508080548061221a5761221a61540c565b600190038181905f5260205f20015f905590558460405161223b919061522e565b60405190819003812033825290879089907f0e1bf50f5cca6659c62146db5b60160121a3752011b621d8c8953a1e0e23d0489060200160405180910390a460019250505b50949350505050565b5f813360405160200161229c9291906151be565b60408051601f1981840301815291815281516020928301205f8181526003909352912054909150156122e1576040516324bf796160e11b815260040160405180910390fd5b6040805160a0810182528281526020808201858152428385015233606084015283515f808252818401865260808501919091528581526003909252929020815181559151909190600182019061233790826152bc565b506040820151600282015560608201516003820180546001600160a01b0319166001600160a01b0390921691909117905560808201518051612383916004840191602090910190614362565b5050335f81815260046020908152604080832080546001810182559084529190922001849055519091506123b890849061522e565b6040519081900381209083907fb8be57bce74a717a1bbd4acf428df655720fce75c1854b02a88484388df241a8905f90a4919050565b5f60208360ff161115612414576040516359b452ef60e01b815260040160405180910390fd5b5f8481526005602052604090206009018054600160ff86161b91908490811061243f5761243f6151ea565b5f918252602090912060088204015460046007909216919091026101000a90041663ffffffff16151590505b9392505050565b6060602084111561249657604051637ee7f45f60e11b815260040160405180910390fd5b836001600160401b038111156124ae576124ae614692565b6040519080825280602002602001820160405280156124d7578160200160208202803683370190505b5090505f5b60ff811685111561227f575f84848389898660ff16818110612500576125006151ea565b9050602002013560405160200161253d9493929190938452602084019290925260f81b6001600160f81b0319166040830152604182015260610190565b60408051601f1981840301815291815281516020928301205f818152600990935291205490915061258157604051633e80ff7760e01b815260040160405180910390fd5b5f818152600960205260409020548351849060ff85169081106125a6576125a66151ea565b602090810291909101015250806125bc81615420565b9150506124dc565b6125cc61423d565b5f85856040516020016125e09291906151be565b6040516020818303038152906040528051906020012090506126038185856136d6565b9695505050505050565b612615614272565b60055f8381526020019081526020015f20604051806101000160405290815f820154815260200160018201805461264b90615239565b80601f016020809104026020016040519081016040528092919081815260200182805461267790615239565b80156126c25780601f10612699576101008083540402835291602001916126c2565b820191905f5260205f20905b8154815290600101906020018083116126a557829003601f168201915b50505050508152602001600282015481526020016003820180546126e590615239565b80601f016020809104026020016040519081016040528092919081815260200182805461271190615239565b801561275c5780601f106127335761010080835404028352916020019161275c565b820191905f5260205f20905b81548152906001019060200180831161273f57829003601f168201915b50505050508152602001600482015481526020016005820154815260200160068201548152602001600782016040518060600160405290815f8201805480602002602001604051908101604052809291908181526020015f905b8282101561285e578382905f5260205f200180546127d390615239565b80601f01602080910402602001604051908101604052809291908181526020018280546127ff90615239565b801561284a5780601f106128215761010080835404028352916020019161284a565b820191905f5260205f20905b81548152906001019060200180831161282d57829003601f168201915b5050505050815260200190600101906127b6565b505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156128b357602002820191905f5260205f20905b81548152602001906001019080831161289f575b505050505081526020016002820180548060200260200160405190810160405280929190818152602001828054801561293457602002820191905f5260205f20905f905b82829054906101000a900463ffffffff1663ffffffff16815260200190600401906020826003010492830192600103820291508084116128f75790505b5050509190925250505090525092915050565b5f83815260036020526040812054849082036129765760405163938a92b760e01b815260040160405180910390fd5b5f81815260036020819052604090912001546001600160a01b031633146129bd576129a081612fb9565b6129bd5760405163dc64d0ad60e01b815260040160405180910390fd5b83336040516020016129d09291906151be565b604051602081830303815290604052805190602001208514612a05576040516327a5901560e11b815260040160405180910390fd5b5f8581526003602052604090206004015415612a335760405162227f7760ea1b815260040160405180910390fd5b5f85815260036020526040812081815590612a5160018301826142d3565b5f600283018190556003830180546001600160a01b0319169055612a79906004840190614325565b5050335f90815260046020526040902080548690829086908110612a9f57612a9f6151ea565b905f5260205f20015414612ac6576040516337c7f25560e01b815260040160405180910390fd5b80548190612ad69060019061538a565b81548110612ae657612ae66151ea565b905f5260205f200154818581548110612b0157612b016151ea565b905f5260205f20018190555080805480612b1d57612b1d61540c565b600190038181905f5260205f20015f90559055336001600160a01b031685604051612b48919061522e565b6040519081900381209088907feda2cc8e002ead8000b1e0c1debfc9a88bd7ee6e94b8dc0763db17849fcf0389905f90a450600195945050505050565b5f805b5f83815260056020526040902060070154811015612bc257612baa838261112e565b15155f03612bba57505f92915050565b600101612b88565b50600192915050565b5f8784604051602001612bdf92919061521d565b60408051601f1981840301815291815281516020928301205f818152600590935290822060070180549193509089908110612c1c57612c1c6151ea565b905f5260205f20018054612c2f90615239565b80601f0160208091040260200160405190810160405280929190818152602001828054612c5b90615239565b8015612ca65780601f10612c7d57610100808354040283529160200191612ca6565b820191905f5260205f20905b815481529060010190602001808311612c8957829003601f168201915b5050505f85815260056020526040812054939450929092039150612cdf905057604051632abde33960e01b815260040160405180910390fd5b5f82815260056020526040812060080180548a908110612d0157612d016151ea565b905f5260205f200154905080876001612d1a919061543e565b60ff161115612d3c576040516359b452ef60e01b815260040160405180910390fd5b5f6040518061010001604052808481526020018e81526020018b81526020018960ff1681526020018d81526020018a81526020018681526020018c8152509050612d8681876137de565b612d9184898c6123ee565b15612daf57604051636d680ca160e11b815260040160405180910390fd5b5050612dbe8a838d8b8a613a51565b50612dca82878a613adc565b612dd4828961112e565b15612e16575f828152600660205260408120805460019290612dfb90849061ffff16615457565b92506101000a81548161ffff021916908361ffff1602179055505b5f546040516340c10f1960e01b8152336004820152670de0b6b3a764000060248201526001600160a01b03909116906340c10f19906044015f604051808303815f87803b158015612e65575f5ffd5b505af1158015612e77573d5f5f3e3d5ffd5b505050508560ff1688837fdd94ce85349373ab535d0282b48fc5974464abc207369050d660ab5296e96db18e8e604051612ebb929190918252602082015260400190565b60405180910390a45050505050505050505050565b60605f85158015612edf575084155b15612eec57506001612efd565b8654612efa90879087613b76565b90505b806001600160401b03811115612f1557612f15614692565b604051908082528060200260200182016040528015612f4e57816020015b612f3b61423d565b815260200190600190039081612f335790505b5091505f5b81811015612fae57612f8988612f69838a6153e2565b81548110612f7957612f796151ea565b905f5260205f20015486866136d6565b838281518110612f9b57612f9b6151ea565b6020908102919091010152600101612f53565b505095945050505050565b6001545f906001600160a01b03161561304a5760015460405163e124bdd960e01b8152600481018490523360248201526001600160a01b039091169063e124bdd990604401602060405180830381865afa158015613019573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061303d9190615471565b1561304a57506001919050565b505f919050565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806130d757507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166130cb5f51602061556b5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b156130f55760405163703e46dd60e11b815260040160405180910390fd5b565b6002546001600160a01b0316331461312257604051636e3a1f4160e01b815260040160405180910390fd5b806001600160a01b03163b5f036109635760405163340aafcd60e11b815260040160405180910390fd5b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156131a6575060408051601f3d908101601f191682019092526131a391810190615490565b60015b6131ce57604051634c9c8ce360e01b81526001600160a01b0383166004820152602401611a8a565b5f51602061556b5f395f51905f5281146131fe57604051632a87526960e21b815260048101829052602401611a8a565b610c378383613bc8565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146130f55760405163703e46dd60e11b815260040160405180910390fd5b60605f85158015613260575084155b1561326d5750600161327c565b61327986885187613b76565b90505b806001600160401b0381111561329457613294614692565b6040519080825280602002602001820160405280156132cd57816020015b6132ba61423d565b8152602001906001900390816132b25790505b5091505f5b81811015612fae57613307886132e8838a6153e2565b815181106132f8576132f86151ea565b602002602001015186866136d6565b838281518110613319576133196151ea565b60209081029190910101526001016132d2565b5f5f838360405160200161334192919061521d565b60408051601f1981840301815291815281516020928301205f818152600590935291205490915015613386576040516303448eef60e51b815260040160405180910390fd5b5f84815260036020908152604080832060040180546001810182559084529190922001829055516133b890849061522e565b60405190819003812033825290859083907fb018e47bdb983351e1bee22415a8f41eef5c2bf1c43c6d3d0992e678ae762df59060200160405180910390a4604080515f606082018181526080830190935291819083613427565b60608152602001906001900390816134125790505b5081526020015f604051908082528060200260200182016040528015613457578160200160208202803683370190505b5081526020015f604051908082528060200260200182016040528015613487578160200160208202803683370190505b50905260408051610100810182528481528151602081810184525f8083528184019283528385018b9052606084018a9052608084018190524260a085015260c0840181905260e084018690528781526005909152929092208151815591519293509160018201906134f890826152bc565b50604082015160028201556060820151600382019061351790826152bc565b506080820151600482015560a0820151600582015560c0820151600682015560e0820151805180516007840191613553918391602001906143ab565b50602082810151805161356c9260018501920190614362565b50604082015180516135889160028401916020909101906143fb565b509498975050505050505050565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f5160206154bc5f395f51905f52916135d490615239565b80601f016020809104026020016040519081016040528092919081815260200182805461360090615239565b801561364b5780601f106136225761010080835404028352916020019161364b565b820191905f5260205f20905b81548152906001019060200180831161362e57829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f5160206154bc5f395f51905f52916135d490615239565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00611128565b6136c4613c1d565b6113ec8282613c42565b6130f5613c1d565b6136de61423d565b5f84815260036020908152604091829020825160a081019093528054835260018101805491939283019161371190615239565b80601f016020809104026020016040519081016040528092919081815260200182805461373d90615239565b80156137885780601f1061375f57610100808354040283529160200191613788565b820191905f5260205f20905b81548152906001019060200180831161376b57829003601f168201915b50505091835250506002830154602082015260038301546001600160a01b03166040820152606001846137c857604080515f8152602081019091526137d3565b6137d3878787613ca1565b905295945050505050565b8160c001514211156138265760405162461bcd60e51b815260206004820152601160248201527014da59db985d1d5c9948195e1c1a5c9959607a1b6044820152606401611a8a565b60e08201515f90815260036020818152604080842090920154825160c08101909352608f8084526001600160a01b039091169392916154dc9083013980519060200120845f01518051906020012085602001518660400151876060015188608001518960a001518a60c001518b60e001516040516020016138ec9998979695949392919098895260208901979097526040880195909552606087019390935260ff91909116608086015260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012090505f61390e82613d68565b90505f61391b8286613d94565b9050836001600160a01b0316816001600160a01b0316146139945760405162461bcd60e51b815260206004820152602d60248201527f496e76616c6964207369676e61747572653a204e6f74207369676e656420627960448201526c10313ab1b5b2ba1037bbb732b960991b6064820152608401611a8a565b6001600160a01b0381165f908152600a6020908152604080832060e08a01518452825280832060a08a0151845290915290205460ff1615613a0c5760405162461bcd60e51b8152602060048201526012602482015271139bdb98d948185b1c9958591e481d5cd95960721b6044820152606401611a8a565b6001600160a01b03165f908152600a6020908152604080832060e08901518452825280832060a090980151835296905294909420805460ff1916600117905550505050565b60408051602081018690529081018390526001600160f81b031960f883901b166060820152606181018490525f9060810160408051601f1981840301815282825280516020918201205f8181526009909252918120899055909250879183917f9f91e060f6a60bc2a5e695f07d854d6a659eab9f5e1441ac24f6e62614bb8b5b91a395945050505050565b60208260ff1610613b00576040516359b452ef60e01b815260040160405180910390fd5b5f8381526005602052604090206009018054600160ff85161b919083908110613b2b57613b2b6151ea565b905f5260205f2090600891828204019190066004028282829054906101000a900463ffffffff161792506101000a81548163ffffffff021916908363ffffffff160217905550505050565b5f82841115613b98576040516309605a0160e41b815260040160405180910390fd5b5f613ba383866153e2565b905083811115613bbf57613bb7858561538a565b91505061246b565b8291505061246b565b613bd182613dbc565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115613c1557610c378282613e1f565b6113ec613e91565b613c25613eb0565b6130f557604051631afcd79f60e31b815260040160405180910390fd5b613c4a613c1d565b5f5160206154bc5f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102613c8384826152bc565b5060038101613c9283826152bc565b505f8082556001909101555050565b5f838152600360205260408120600401805460609290613cc390869086613b76565b90505f816001600160401b03811115613cde57613cde614692565b604051908082528060200260200182016040528015613d07578160200160208202803683370190505b5090505f5b82811015613d5d5783613d1f82896153e2565b81548110613d2f57613d2f6151ea565b905f5260205f200154828281518110613d4a57613d4a6151ea565b6020908102919091010152600101613d0c565b509695505050505050565b5f611128613d74613ec9565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f613da28686613ed7565b925092509250613db28282613f20565b5090949350505050565b806001600160a01b03163b5f03613df157604051634c9c8ce360e01b81526001600160a01b0382166004820152602401611a8a565b5f51602061556b5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b031684604051613e3b919061522e565b5f60405180830381855af49150503d805f8114613e73576040519150601f19603f3d011682016040523d82523d5f602084013e613e78565b606091505b5091509150613e88858383613fd8565b95945050505050565b34156130f55760405163b398979f60e01b815260040160405180910390fd5b5f613eb9613694565b54600160401b900460ff16919050565b5f613ed2614034565b905090565b5f5f5f8351604103613f0e576020840151604085015160608601515f1a613f00888285856140a7565b955095509550505050613f19565b505081515f91506002905b9250925092565b5f826003811115613f3357613f336154a7565b03613f3c575050565b6001826003811115613f5057613f506154a7565b03613f6e5760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115613f8257613f826154a7565b03613fa35760405163fce698f760e01b815260048101829052602401611a8a565b6003826003811115613fb757613fb76154a7565b036113ec576040516335e2f38360e21b815260048101829052602401611a8a565b606082613fed57613fe88261416b565b61246b565b815115801561400457506001600160a01b0384163b155b1561402d57604051639996b31560e01b81526001600160a01b0385166004820152602401611a8a565b508061246b565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f61405e614193565b6140666141fb565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156140e057505f91506003905082611fd6565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015614131573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661415c57505f925060019150829050611fd6565b975f9750879650945050505050565b80511561417a57805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f5f5160206154bc5f395f51905f52816141ab613596565b8051909150156141c357805160209091012092915050565b815480156141d2579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f5160206154bc5f395f51905f5281614213613656565b80519091501561422b57805160209091012092915050565b600182015480156141d2579392505050565b6040518060a001604052805f8152602001606081526020015f81526020015f6001600160a01b03168152602001606081525090565b6040518061010001604052805f8152602001606081526020015f8152602001606081526020015f81526020015f81526020015f81526020016142ce60405180606001604052806060815260200160608152602001606081525090565b905290565b5080546142df90615239565b5f825580601f106142ee575050565b601f0160209004905f5260205f2090810190610963919061449e565b5080545f8255905f5260205f209081019061096391906144b2565b5080545f8255905f5260205f2090810190610963919061449e565b5080545f825560070160089004905f5260205f2090810190610963919061449e565b828054828255905f5260205f2090810192821561439b579160200282015b8281111561439b578251825591602001919060010190614380565b506143a792915061449e565b5090565b828054828255905f5260205f209081019282156143ef579160200282015b828111156143ef57825182906143df90826152bc565b50916020019190600101906143c9565b506143a79291506144b2565b828054828255905f5260205f209060070160089004810192821561439b579160200282015f5b8382111561446557835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302614421565b80156144955782816101000a81549063ffffffff0219169055600401602081600301049283019260010302614465565b50506143a79291505b5b808211156143a7575f815560010161449f565b808211156143a7575f6144c582826142d3565b506001016144b2565b5f602082840312156144de575f5ffd5b5035919050565b5f602082840312156144f5575f5ffd5b81356001600160401b0381111561450a575f5ffd5b8201610120818503121561246b575f5ffd5b80356001600160a01b0381168114614532575f5ffd5b919050565b5f5f5f5f5f60a0868803121561454b575f5ffd5b6145548661451c565b97602087013597506040870135966060810135965060800135945092505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b805182525f602082015160a060208501526145c160a0850182614575565b604084810151908601526060808501516001600160a01b031690860152608080850151868303918701919091528051808352602091820193505f9291909101905b808310156146255783518252602082019150602084019350600183019250614602565b5095945050505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561468657603f198786030184526146718583516145a3565b94506020938401939190910190600101614655565b50929695505050505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156146ce576146ce614692565b604052919050565b5f82601f8301126146e5575f5ffd5b8135602083015f5f6001600160401b0384111561470457614704614692565b50601f8301601f1916602001614719816146a6565b91505082815285838301111561472d575f5ffd5b828260208301375f92810160200192909252509392505050565b5f5f60408385031215614758575f5ffd5b82356001600160401b0381111561476d575f5ffd5b614779858286016146d6565b9250506147886020840161451c565b90509250929050565b5f5f83601f8401126147a1575f5ffd5b5081356001600160401b038111156147b7575f5ffd5b6020830191508360208260051b8501011115610a23575f5ffd5b5f5f602083850312156147e2575f5ffd5b82356001600160401b038111156147f7575f5ffd5b61480385828601614791565b90969095509350505050565b5f5f5f5f5f5f5f5f5f60e08a8c031215614827575f5ffd5b89356001600160401b0381111561483c575f5ffd5b6148488c828d016146d6565b99505060208a0135975060408a01356001600160401b0381111561486a575f5ffd5b6148768c828d016146d6565b97505060608a0135955060808a01356001600160401b03811115614898575f5ffd5b6148a48c828d01614791565b90965094505060a08a01356001600160401b038111156148c2575f5ffd5b6148ce8c828d01614791565b9a9d999c50979a9699959894979660c00135949350505050565b5f5f604083850312156148f9575f5ffd5b50508035926020909101359150565b5f5f5f6060848603121561491a575f5ffd5b83356001600160401b0381111561492f575f5ffd5b61493b868287016146d6565b935050602084013591506149516040850161451c565b90509250925092565b604081525f61496c6040830185614575565b90508260208301529392505050565b5f5f6040838503121561498c575f5ffd5b6149958361451c565b915060208301356001600160401b038111156149af575f5ffd5b6149bb858286016146d6565b9150509250929050565b5f6001600160401b038211156149dd576149dd614692565b5060051b60200190565b5f82601f8301126149f6575f5ffd5b8135614a09614a04826149c5565b6146a6565b8082825260208201915060208360051b860101925085831115614a2a575f5ffd5b602085015b838110156146255780356001600160401b03811115614a4c575f5ffd5b614a5b886020838a01016146d6565b84525060209283019201614a2f565b5f82601f830112614a79575f5ffd5b8135614a87614a04826149c5565b8082825260208201915060208360051b860101925085831115614aa8575f5ffd5b602085015b83811015614625578035835260209283019201614aad565b5f5f5f5f5f5f5f5f5f60e08a8c031215614add575f5ffd5b89356001600160401b03811115614af2575f5ffd5b614afe8c828d016149e7565b99505060208a0135975060408a01356001600160401b03811115614b20575f5ffd5b614b2c8c828d016146d6565b97505060608a01356001600160401b03811115614b47575f5ffd5b614b538c828d01614a6a565b96505060808a01356001600160401b03811115614898575f5ffd5b602081525f61246b6020830184614575565b5f5f60408385031215614b91575f5ffd5b8235915060208301356001600160401b038111156149af575f5ffd5b5f8151808452602084019350602083015f5b82811015614bdd578151865260209586019590910190600101614bbf565b5093949350505050565b5f8151808452602084019350602083015f5b82811015614bdd57815163ffffffff16865260209586019590910190600101614bf9565b5f6060830182516060855281815180845260808701915060808160051b88010193506020830192505f5b81811015614c7857607f19888603018352614c63858551614575565b94506020938401939290920191600101614c47565b5050505060208301518482036020860152614c938282614bad565b91505060408301518482036040860152613e888282614be7565b805182525f60208201516101006020850152614ccd610100850182614575565b90506040830151604085015260608301518482036060860152614cf08282614575565b9150506080830151608085015260a083015160a085015260c083015160c085015260e083015184820360e0860152613e888282614c1d565b602081525f61246b6020830184614cad565b5f5f5f5f5f5f60a08789031215614d4f575f5ffd5b86356001600160401b03811115614d64575f5ffd5b614d7089828a01614791565b909a90995060208901359860408101359850606081013597506080013595509350505050565b5f60208284031215614da6575f5ffd5b61246b8261451c565b60ff60f81b8816815260e060208201525f614dcd60e0830189614575565b8281036040840152614ddf8189614575565b606084018890526001600160a01b038716608085015260a0840186905283810360c08501529050614e108185614bad565b9a9950505050505050505050565b5f5f5f5f5f60a08688031215614e32575f5ffd5b8535945060208601356001600160401b03811115614e4e575f5ffd5b614e5a888289016146d6565b945050604086013592506060860135915060808601356001600160401b03811115614e83575f5ffd5b614e8f888289016146d6565b9150509295509295909350565b5f5f5f5f60808587031215614eaf575f5ffd5b84356001600160401b03811115614ec4575f5ffd5b614ed0878288016146d6565b94505060208501356001600160401b03811115614eeb575f5ffd5b614ef7878288016146d6565b93505060408501359150614f0d6060860161451c565b905092959194509250565b606081525f614f2a6060830186614cad565b602083019490945250901515604090910152919050565b5f5f5f5f60808587031215614f54575f5ffd5b843593506020850135925060408501356001600160401b03811115614f77575f5ffd5b614f83878288016146d6565b949793965093946060013593505050565b5f60208284031215614fa4575f5ffd5b81356001600160401b03811115614fb9575f5ffd5b6119b2848285016146d6565b803560ff81168114614532575f5ffd5b5f5f5f60608486031215614fe7575f5ffd5b83359250614ff760208501614fc5565b929592945050506040919091013590565b5f5f5f5f6060858703121561501b575f5ffd5b84356001600160401b03811115615030575f5ffd5b61503c87828801614791565b90989097506020870135966040013595509350505050565b602080825282518282018190525f918401906040840190835b8181101561508b57835183526020938401939092019160010161506d565b509095945050505050565b5f5f5f5f608085870312156150a9575f5ffd5b84356001600160401b038111156150be575f5ffd5b6150ca878288016146d6565b9450506150d96020860161451c565b93969395505050506040820135916060013590565b602081525f61246b60208301846145a3565b5f5f5f60608486031215615112575f5ffd5b8335925060208401356001600160401b0381111561512e575f5ffd5b61513a868287016146d6565b93969395505050506040919091013590565b5f6020828403121561515c575f5ffd5b61246b82614fc5565b5f5f8335601e1984360301811261517a575f5ffd5b8301803591506001600160401b03821115615193575f5ffd5b602001915036819003821315610a23575f5ffd5b5f81518060208401855e5f93019283525090919050565b5f6151c982856151a7565b60609390931b6bffffffffffffffffffffffff191683525050601401919050565b634e487b7160e01b5f52603260045260245ffd5b5f823561011e19833603018112615213575f5ffd5b9190910192915050565b8281525f6119b260208301846151a7565b5f61246b82846151a7565b600181811c9082168061524d57607f821691505b60208210810361526b57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115610c3757805f5260205f20601f840160051c810160208510156152965750805b601f840160051c820191505b818110156152b5575f81556001016152a2565b5050505050565b81516001600160401b038111156152d5576152d5614692565b6152e9816152e38454615239565b84615271565b6020601f82116001811461531b575f83156153045750848201515b5f19600385901b1c1916600184901b1784556152b5565b5f84815260208120601f198516915b8281101561534a578785015182556020948501946001909201910161532a565b508482101561536757868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561112857611128615376565b5f5f8335601e198436030181126153b2575f5ffd5b8301803591506001600160401b038211156153cb575f5ffd5b6020019150600581901b3603821315610a23575f5ffd5b8082018082111561112857611128615376565b808202811582820484141761112857611128615376565b634e487b7160e01b5f52603160045260245ffd5b5f60ff821660ff810361543557615435615376565b60010192915050565b60ff818116838216019081111561112857611128615376565b61ffff818116838216019081111561112857611128615376565b5f60208284031215615481575f5ffd5b8151801515811461246b575f5ffd5b5f602082840312156154a0575f5ffd5b5051919050565b634e487b7160e01b5f52602160045260245ffdfea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10053746f7261676544617461286279746573206368756e6b4349442c6279746573333220626c6f636b4349442c75696e74323536206368756e6b496e6465782c75696e743820626c6f636b496e6465782c62797465733332206e6f646549642c75696e74323536206e6f6e63652c75696e7432353620646561646c696e652c62797465733332206275636b6574496429360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca2646970667358221220cae7b8dfb48cad5c9a83ea145a84cf43fb72d57143f575dfdc24b23083d3849164736f6c634300081c0033",
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

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address _upgradeAuthority) returns()
func (_Storage *StorageTransactor) SetAuthority(opts *bind.TransactOpts, _upgradeAuthority common.Address) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "setAuthority", _upgradeAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address _upgradeAuthority) returns()
func (_Storage *StorageSession) SetAuthority(_upgradeAuthority common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetAuthority(&_Storage.TransactOpts, _upgradeAuthority)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(address _upgradeAuthority) returns()
func (_Storage *StorageTransactorSession) SetAuthority(_upgradeAuthority common.Address) (*types.Transaction, error) {
	return _Storage.Contract.SetAuthority(&_Storage.TransactOpts, _upgradeAuthority)
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
