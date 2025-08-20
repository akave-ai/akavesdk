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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockAlreadyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockNonexists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketInvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketNonempty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketNonexists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"}],\"name\":\"ChunkCIDMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileChunkDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileFullyUploaded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNameDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNonempty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNotFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cidsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sizesLength\",\"type\":\"uint256\"}],\"name\":\"InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlocksAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEncodedSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFileBlocksCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFileCID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLastBlockSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LastChunkDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OffsetOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"AddFileBlocks\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AddFileChunk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"}],\"name\":\"AddPeerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CommitFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CreateBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CreateFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DeleteBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DeleteFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"}],\"name\":\"DeletePeerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"}],\"name\":\"FillChunkBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_BLOCKS_PER_FILE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BLOCK_SIZE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessManager\",\"outputs\":[{\"internalType\":\"contractIAccessManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"cid\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedChunkSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"cids\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkBlocksSizes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"addFileChunk\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"cids\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"encodedChunkSizes\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[][]\",\"name\":\"chunkBlocksCIDs\",\"type\":\"bytes32[][]\"},{\"internalType\":\"uint256[][]\",\"name\":\"chunkBlockSizes\",\"type\":\"uint256[][]\"},{\"internalType\":\"uint256\",\"name\":\"startingChunkIndex\",\"type\":\"uint256\"}],\"name\":\"addFileChunks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isReplica\",\"type\":\"bool\"}],\"name\":\"addPeerBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedFileSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"}],\"name\":\"commitFile\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"}],\"name\":\"createBucket\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"createFile\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deleteBucket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deleteFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deletePeerBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fileFillCounter\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fileRewardClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockCID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIStorage.FillChunkBlockArgs\",\"name\":\"args\",\"type\":\"tuple\"}],\"name\":\"fillChunkBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"blockCID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIStorage.FillChunkBlockArgs[]\",\"name\":\"args\",\"type\":\"tuple[]\"}],\"name\":\"fillChunkBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fulfilledBlocks\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fileOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileLimit\",\"type\":\"uint256\"}],\"name\":\"getBucketByName\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket\",\"name\":\"bucket\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getBucketIndexByName\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"}],\"name\":\"getBucketsByIds\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"bucketOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bucketLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileLimit\",\"type\":\"uint256\"}],\"name\":\"getBucketsByIdsWithFiles\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket[]\",\"name\":\"buckets\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getChunkByIndex\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getFileById\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"chunkCIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkSize\",\"type\":\"uint256[]\"}],\"internalType\":\"structIStorage.Chunk\",\"name\":\"chunks\",\"type\":\"tuple\"}],\"internalType\":\"structIStorage.File\",\"name\":\"file\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"getFileByName\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"chunkCIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkSize\",\"type\":\"uint256[]\"}],\"internalType\":\"structIStorage.Chunk\",\"name\":\"chunks\",\"type\":\"tuple\"}],\"internalType\":\"structIStorage.File\",\"name\":\"file\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getFileIndexById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getFileOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"bucketName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getFullFileInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"chunkCIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkSize\",\"type\":\"uint256[]\"}],\"internalType\":\"structIStorage.Chunk\",\"name\":\"chunks\",\"type\":\"tuple\"}],\"internalType\":\"structIStorage.File\",\"name\":\"file\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fileLimit\",\"type\":\"uint256\"}],\"name\":\"getOwnerBuckets\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket[]\",\"name\":\"buckets\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"getPeerBlockIndexById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"cids\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"getPeersArrayByPeerBlockCid\",\"outputs\":[{\"internalType\":\"bytes32[][]\",\"name\":\"peers\",\"type\":\"bytes32[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"getPeersByPeerBlockCid\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"peers\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"isBlockFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"isChunkFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"isFileFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"isFileFilledV2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"}],\"name\":\"isPeerBlockReplica\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accessManagerAddress\",\"type\":\"address\"}],\"name\":\"setAccessManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIAkaveToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60a060405230608052348015610013575f5ffd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516158306100f95f395f81816136150152818161363e015261377701526158305ff3fe608060405260043610610275575f3560e01c80636ce023631161014a578063d606205d116100be578063f8a3e41a11610078578063f8a3e41a146108ad578063faec0542146108ef578063fc0c546a1461090e578063fd1d3c0c1461092c578063fd21c2841461094b578063fdcb60681461096a575f5ffd5b8063d606205d146107d9578063d6d3110b14610805578063e3f787e814610824578063e4ba8a5814610843578063f83533d214610862578063f855169a1461088e575f5ffd5b80639ccd46461161010f5780639ccd4646146106f9578063ad3cb1cc1461070e578063b80777ea1461073e578063c4d66de814610750578063c95808041461076f578063d4967ce6146107ab575f5ffd5b80636ce02363146106475780637912bf681461067557806384b0196e1461069457806387c1ac07146106bb5780639a2e82b3146106da575f5ffd5b80634d7dc614116101ec5780635a4e9564116101a65780635a4e95641461052c5780635ecdfb531461057d57806362838c23146105a957806368e6408f146105c85780636a6e658b146106095780636af0f80114610628575f5ffd5b80634d7dc614146104715780634f1ef2861461049e57806351dc0c23146104b157806352d1902d146104d057806354fd4d50146104e4578063564b81ef1461051a575f5ffd5b8063308c3cf21161023d578063308c3cf21461039c57806330b91d07146103bb57806335bdb711146103e85780633f3839801461040757806348b498751461042657806349cdf6ac14610452575f5ffd5b8063018c1e9c146102795780631b475ef4146102bc5780631ed572b71461031b578063262a61ce1461033c578063287e677f14610368575b5f5ffd5b348015610284575f5ffd5b506102a7610293366004614664565b60086020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b3480156102c7575f5ffd5b506103036102d6366004614664565b5f90815260046020908152604080832060029081015484529091529020600301546001600160a01b031690565b6040516001600160a01b0390911681526020016102b3565b348015610326575f5ffd5b5061033a61033536600461467b565b610989565b005b348015610347575f5ffd5b5061035b6103563660046146cd565b610a41565b6040516102b391906147c6565b348015610373575f5ffd5b506103876103823660046148de565b610ac1565b604080519283529015156020830152016102b3565b3480156103a7575f5ffd5b5061033a6103b6366004614968565b610b53565b3480156103c6575f5ffd5b506103da6103d53660046149a6565b610d65565b6040519081526020016102b3565b3480156103f3575f5ffd5b5061035b610402366004614968565b611032565b348015610412575f5ffd5b506102a7610421366004614a7f565b6112a1565b348015610431575f5ffd5b50610445610440366004614a9f565b61135e565b6040516102b39190614ae2565b34801561045d575f5ffd5b5061038761046c366004614b24565b6113dc565b34801561047c575f5ffd5b5061049061048b366004614a7f565b611467565b6040516102b3929190614b76565b61033a6104ac366004614b97565b61155f565b3480156104bc575f5ffd5b5061033a6104cb366004614ccb565b611575565b3480156104db575f5ffd5b506103da6116fc565b3480156104ef575f5ffd5b506040805180820190915260058152640312e302e360dc1b60208201525b6040516102b39190614d74565b348015610525575f5ffd5b50466103da565b348015610537575f5ffd5b50610568610546366004614a7f565b600560209081525f928352604080842090915290825290205463ffffffff1681565b60405163ffffffff90911681526020016102b3565b348015610588575f5ffd5b5061059c610597366004614a9f565b611717565b6040516102b39190614e77565b3480156105b4575f5ffd5b5061035b6105c3366004614e89565b6119f6565b3480156105d3575f5ffd5b506102a76105e2366004614664565b5f9081526006602090815260408083205460049092529091206007015461ffff9091161490565b348015610614575f5ffd5b506102a7610623366004614ee5565b611a3b565b348015610633575f5ffd5b506103da610642366004614a9f565b611be9565b348015610652575f5ffd5b5061065d620f424081565b6040516001600160401b0390911681526020016102b3565b348015610680575f5ffd5b506102a761068f366004614a7f565b611c70565b34801561069f575f5ffd5b506106a8611cbd565b6040516102b39796959493929190614f41565b3480156106c6575f5ffd5b506103da6106d5366004614fbd565b611d6b565b3480156106e5575f5ffd5b506103da6106f436600461501b565b611e6d565b348015610704575f5ffd5b5061065d61040081565b348015610719575f5ffd5b5061050d604051806040016040528060058152602001640352e302e360dc1b81525081565b348015610749575f5ffd5b50426103da565b34801561075b575f5ffd5b5061033a61076a366004615099565b6121ca565b34801561077a575f5ffd5b5061033a610789366004615099565b600180546001600160a01b0319166001600160a01b0392909216919091179055565b3480156107b6575f5ffd5b506107ca6107c53660046150b2565b61231c565b6040516102b39392919061512e565b3480156107e4575f5ffd5b506107f86107f3366004615157565b61237a565b6040516102b391906151be565b348015610810575f5ffd5b506102a761081f366004615215565b61248d565b34801561082f575f5ffd5b506103da61083e366004615268565b6126fe565b34801561084e575f5ffd5b506102a761085d3660046152a9565b612864565b34801561086d575f5ffd5b5061088161087c3660046152dc565b6128bc565b6040516102b39190615334565b348015610899575f5ffd5b506103876108a8366004615346565b612967565b3480156108b8575f5ffd5b506108dc6108c7366004614664565b60066020525f908152604090205461ffff1681565b60405161ffff90911681526020016102b3565b3480156108fa575f5ffd5b5061059c610909366004614664565b6129e1565b348015610919575f5ffd5b505f54610303906001600160a01b031681565b348015610937575f5ffd5b506102a761094636600461537b565b612c9a565b348015610956575f5ffd5b506102a7610965366004614664565b612ed7565b348015610975575f5ffd5b50600154610303906001600160a01b031681565b610a3e813560208301356040840135606085013560808601356109b260c0880160a089016153c7565b6109bf60c08901896153e0565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250610a009250505060e08a018a6153e0565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152505050506101008a0135612f1d565b50565b6001600160a01b0385165f9081526003602090815260409182902080548351818402810184019094528084526060939192610ab692918491830182828015610aa657602002820191905f5260205f20905b815481526020019060010190808311610a92575b5050505050878787876001613254565b979650505050505050565b5f5f5f8484604051602001610ad7929190615439565b60408051601f1981840301815291815281516020928301206001600160a01b0387165f9081526003909352908220909250905b8154811015610b485782828281548110610b2657610b26615465565b905f5260205f20015403610b405780945060019350610b48565b600101610b0a565b5050505b9250929050565b5f5b81811015610d6057610d58838383818110610b7257610b72615465565b9050602002810190610b849190615479565b35848484818110610b9757610b97615465565b9050602002810190610ba99190615479565b60200135858585818110610bbf57610bbf615465565b9050602002810190610bd19190615479565b60400135868686818110610be757610be7615465565b9050602002810190610bf99190615479565b60600135878787818110610c0f57610c0f615465565b9050602002810190610c219190615479565b60800135888888818110610c3757610c37615465565b9050602002810190610c499190615479565b610c5a9060c081019060a0016153c7565b898989818110610c6c57610c6c615465565b9050602002810190610c7e9190615479565b610c8c9060c08101906153e0565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508d92508c91508b9050818110610cd457610cd4615465565b9050602002810190610ce69190615479565b610cf49060e08101906153e0565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508e92508d91508c9050818110610d3c57610d3c615465565b9050602002810190610d4e9190615479565b6101000135612f1d565b600101610b55565b505050565b5f8881526002602052604081205489908203610d945760405163938a92b760e01b815260040160405180910390fd5b5f818152600260205260409020600301546001600160a01b03163314610dda57610dbd81613572565b610dda5760405163dc64d0ad60e01b815260040160405180910390fd5b89895f8282604051602001610df0929190615498565b60405160208183030381529060405280519060200120905060045f8281526020019081526020015f20600401545f14610e3c5760405163d96b03b160e01b815260040160405180910390fd5b5f8d8d604051602001610e50929190615498565b60408051601f1981840301815291815281516020928301205f8181526004909352908220549092509003610e9757604051632abde33960e01b815260040160405180910390fd5b5f818152600460205260409020600701548714610ec7576040516301c0b3dd60e61b815260040160405180910390fd5b8988141580610ed6575060208a115b15610ef4576040516373d8ccd360e11b815260040160405180910390fd5b60208a1015610f4a575f8181526007602052604090205461ffff1615610f2d576040516355cbc83160e01b815260040160405180910390fd5b5f818152600760205260409020805461ffff191661ffff8c161790555b8b5f03610f6a57604051631b6fdfeb60e01b815260040160405180910390fd5b8c604051610f7891906154a9565b604051908190038120338252908f9083907f0d877d3656ad561e6c8224e31cc98f1b3e62b1d828fd396c47ef0daccb6838e59060200160405180910390a460045f8281526020019081526020015f206007015f018f908060018154018082558091505060019003905f5260205f20015f909190919091509081610ffb9190615537565b505f818152600460209081526040822060080180546001810182559083529120018c90559450505050509998505050505050505050565b60605f826001600160401b0381111561104d5761104d614829565b60405190808252806020026020018201604052801561108657816020015b61107361449f565b81526020019060019003908161106b5790505b5090505f5b83811015611297575f60025f8787858181106110a9576110a9615465565b9050602002013581526020019081526020015f206040518060a00160405290815f82015481526020016001820180546110e1906154b4565b80601f016020809104026020016040519081016040528092919081815260200182805461110d906154b4565b80156111585780601f1061112f57610100808354040283529160200191611158565b820191905f5260205f20905b81548152906001019060200180831161113b57829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152602001600482018054806020026020016040519081016040528092919081815260200182805480156111e657602002820191905f5260205f20905b8154815260200190600101908083116111d2575b50505050508152505090506040518060a00160405280825f01518152602001826020015181526020018260400151815260200182606001516001600160a01b031681526020015f6001600160401b0381111561124457611244614829565b60405190808252806020026020018201604052801561126d578160200160208202803683370190505b5081525083838151811061128357611283615465565b60209081029190910101525060010161108b565b5090505b92915050565b5f8281526004602052604081206007015481906112c090600190615605565b905080831480156112e157505f8481526007602052604090205461ffff1615155b15611337575f848152600760205260408120546113079060019061ffff1681901b615605565b5f868152600560209081526040808320888452909152902054811663ffffffff908116911614925061129b915050565b50505f91825260056020908152604080842092845291905290205463ffffffff9081161490565b6060600a5f8481526020019081526020015f208260405161137f91906154a9565b90815260408051918290036020908101832080548083028501830190935282845291908301828280156113cf57602002820191905f5260205f20905b8154815260200190600101908083116113bb575b5050505050905092915050565b5f5f5f85846040516020016113f2929190615439565b60408051601f1981840301815291815281516020928301205f8181526002909352908220909250600401905b815481101561145c578682828154811061143a5761143a615465565b905f5260205f20015403611454578094506001935061145c565b60010161141e565b505050935093915050565b5f82815260046020526040812060070180546060929182918590811061148f5761148f615465565b905f5260205f200180546114a2906154b4565b80601f01602080910402602001604051908101604052809291908181526020018280546114ce906154b4565b80156115195780601f106114f057610100808354040283529160200191611519565b820191905f5260205f20905b8154815290600101906020018083116114fc57829003601f168201915b5050505f88815260046020526040812060080180549495509093909250879150811061154757611547615465565b5f918252602090912001549196919550909350505050565b61156761360a565b61157182826136b0565b5050565b5f888152600260205260408120548991036115a35760405163938a92b760e01b815260040160405180910390fd5b5f818152600260205260409020600301546001600160a01b031633146115e9576115cc81613572565b6115e95760405163dc64d0ad60e01b815260040160405180910390fd5b88885f82826040516020016115ff929190615498565b60405160208183030381529060405280519060200120905060045f8281526020019081526020015f20600401545f1461164b5760405163d96b03b160e01b815260040160405180910390fd5b5f5b8d518110156116ec576116e38e828151811061166b5761166b615465565b60200260200101518e8e8e858151811061168757611687615465565b60200260200101518e8e878181106116a1576116a1615465565b90506020028101906116b39190615618565b8e8e898181106116c5576116c5615465565b90506020028101906116d79190615618565b898f6103d5919061565d565b5060010161164d565b5050505050505050505050505050565b5f61170561376c565b505f5160206157db5f395f51905f5290565b61171f6144d4565b5f8383604051602001611733929190615498565b60408051601f1981840301815282825280516020918201205f818152600483528390206101008501909352825484526001830180549195509184019190611779906154b4565b80601f01602080910402602001604051908101604052809291908181526020018280546117a5906154b4565b80156117f05780601f106117c7576101008083540402835291602001916117f0565b820191905f5260205f20905b8154815290600101906020018083116117d357829003601f168201915b5050505050815260200160028201548152602001600382018054611813906154b4565b80601f016020809104026020016040519081016040528092919081815260200182805461183f906154b4565b801561188a5780601f106118615761010080835404028352916020019161188a565b820191905f5260205f20905b81548152906001019060200180831161186d57829003601f168201915b50505050508152602001600482015481526020016005820154815260200160068201548152602001600782016040518060400160405290815f8201805480602002602001604051908101604052809291908181526020015f905b8282101561198c578382905f5260205f20018054611901906154b4565b80601f016020809104026020016040519081016040528092919081815260200182805461192d906154b4565b80156119785780601f1061194f57610100808354040283529160200191611978565b820191905f5260205f20905b81548152906001019060200180831161195b57829003601f168201915b5050505050815260200190600101906118e4565b505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156119e157602002820191905f5260205f20905b8154815260200190600101908083116119cd575b50505091909252505050905250949350505050565b6060610ab68787808060200260200160405190810160405280939291908181526020018383602002808284375f92019190915250899250889150879050866001613254565b5f8581526009602052604081205460ff16611a6957604051631512312160e01b815260040160405180910390fd5b848484604051602001611a7e93929190615670565b604051602081830303815290604052805190602001208614611ab3576040516332c83a2360e21b815260040160405180910390fd5b5f858585604051602001611ac993929190615670565b604051602081830303815290604052805190602001209050868114611b0157604051637ae9080f60e11b815260040160405180910390fd5b5f878152600960209081526040808320805460ff19169055878352600a9091528082209051611b319087906154a9565b908152602001604051809103902090508060018280549050611b539190615605565b81548110611b6357611b63615465565b905f5260205f200154818581548110611b7e57611b7e615465565b905f5260205f20018190555080805480611b9a57611b9a615687565b600190038181905f5260205f20015f9055905586887f37505c6d2cdf9e778d6110c5ad2e49c649d521a18d2da6f007d3364bd83a45ae60405160405180910390a3506001979650505050505050565b5f8281526002602052604081205483908203611c185760405163938a92b760e01b815260040160405180910390fd5b5f818152600260205260409020600301546001600160a01b03163314611c5e57611c4181613572565b611c5e5760405163dc64d0ad60e01b815260040160405180910390fd5b611c6884846137b5565b949350505050565b5f5f8284604051602001611c8e929190918252602082015260400190565b60408051808303601f1901815291815281516020928301205f908152600b90925290205460ff16949350505050565b5f60608082808083815f51602061572c5f395f51905f528054909150158015611ce857506001810154155b611d315760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064015b60405180910390fd5b611d396139d2565b611d41613a92565b604080515f80825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b5f848484604051602001611d8193929190615670565b60408051601f19818403018152828252805160209182012090830188905290820186905291505f9060600160408051601f1981840301815291815281516020928301205f85815260099093529120805460ff1916600117905590508215611dfb575f818152600b60205260409020805460ff191660011790555b5f858152600a6020526040908190209051611e179086906154a9565b908152604051908190036020908101822080546001810182555f918252918120909101889055879184917f9f91e060f6a60bc2a5e695f07d854d6a659eab9f5e1441ac24f6e62614bb8b5b91a350949350505050565b5f8581526002602052604081205486908203611e9c5760405163938a92b760e01b815260040160405180910390fd5b5f818152600260205260409020600301546001600160a01b03163314611ee257611ec581613572565b611ee25760405163dc64d0ad60e01b815260040160405180910390fd5b86865f8282604051602001611ef8929190615498565b60405160208183030381529060405280519060200120905060045f8281526020019081526020015f20600401545f14611f445760405163d96b03b160e01b815260040160405180910390fd5b5f8a8a604051602001611f58929190615498565b60408051601f1981840301815291815281516020928301205f81815260068452828120546004909452919091206007015490925061ffff90911614611fb057604051632e1b8f4960e11b815260040160405180910390fd5b885f03611fd057604051631b6fdfeb60e01b815260040160405180910390fd5b86515f03611ff157604051637f19edc960e11b815260040160405180910390fd5b5f805b5f8381526004602052604090206008015481101561204c575f83815260046020526040902060080180548290811061202e5761202e615465565b905f5260205f20015482612042919061565d565b9150600101611ff4565b5089811461206d57604051631b6fdfeb60e01b815260040160405180910390fd5b5f8281526004602052604090206001016120878982615537565b505f8281526004602081815260408084209283018e905560069092018c90556008905290205460ff1661216f575f828152600860209081526040808320805460ff191660011790556004909152812060070154906120e682600a61569b565b6120f890670de0b6b3a764000061569b565b90505f8211801561210857505f81115b1561216c575f546040516340c10f1960e01b8152336004820152602481018390526001600160a01b03909116906340c10f19906044015f604051808303815f87803b158015612155575f5ffd5b505af1158015612167573d5f5f3e3d5ffd5b505050505b50505b8a60405161217d91906154a9565b604051908190038120338252908d9084907fc208b556b9717832827b84e296c7f4bb33292ba520fed2eae2f7c0d61f1c0a619060200160405180910390a4509a9950505050505050505050565b5f6121d3613ad0565b805490915060ff600160401b82041615906001600160401b03165f811580156121f95750825b90505f826001600160401b031660011480156122145750303b155b905081158015612222575080155b156122405760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561226a57845460ff60401b1916600160401b1785555b6122ac6040518060400160405280600781526020016653746f7261676560c81b815250604051806040016040528060018152602001603160f81b815250613af8565b6122b4613b0a565b5f80546001600160a01b0319166001600160a01b038816179055831561231457845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b6123246144d4565b5f5f5f858760405160200161233a929190615498565b60405160208183030381529060405280519060200120905061235c8688611717565b93506123698882876113dc565b9093509150505b9450945094915050565b6060826001600160401b0381111561239457612394614829565b6040519080825280602002602001820160405280156123c757816020015b60608152602001906001900390816123b25790505b5090505f5b8381101561248557600a5f8686848181106123e9576123e9615465565b9050602002013581526020019081526020015f208360405161240b91906154a9565b908152604080519182900360209081018320805480830285018301909352828452919083018282801561245b57602002820191905f5260205f20905b815481526020019060010190808311612447575b505050505082828151811061247257612472615465565b60209081029190910101526001016123cc565b509392505050565b5f83815260026020526040812054849082036124bc5760405163938a92b760e01b815260040160405180910390fd5b5f818152600260205260409020600301546001600160a01b03163314612502576124e581613572565b6125025760405163dc64d0ad60e01b815260040160405180910390fd5b5f86815260046020526040812054900361252f57604051632abde33960e01b815260040160405180910390fd5b8484604051602001612542929190615498565b60405160208183030381529060405280519060200120861461257757604051630ef4797b60e31b815260040160405180910390fd5b5f86815260046020526040812081815590612595600183018261452e565b600282015f9055600382015f6125ab919061452e565b5f60048301819055600583018190556006830181905560078301906125d08282614565565b6125dd600183015f614580565b5050505f86815260026020526040902060040180549091508410158061261d57508681858154811061261157612611615465565b905f5260205f20015414155b1561263b576040516337c7f25560e01b815260040160405180910390fd5b8054819061264b90600190615605565b8154811061265b5761265b615465565b905f5260205f20015481858154811061267657612676615465565b905f5260205f2001819055508080548061269257612692615687565b600190038181905f5260205f20015f90559055846040516126b391906154a9565b60405190819003812033825290879089907f0e1bf50f5cca6659c62146db5b60160121a3752011b621d8c8953a1e0e23d0489060200160405180910390a45060019695505050505050565b5f8133604051602001612712929190615439565b60408051601f1981840301815291815281516020928301205f818152600290935291205490915015612757576040516324bf796160e11b815260040160405180910390fd5b6040805160a0810182528281526020808201858152428385015233606084015283515f80825281840186526080850191909152858152600290925292902081518155915190919060018201906127ad9082615537565b506040820151600282015560608201516003820180546001600160a01b0319166001600160a01b03909216919091179055608082015180516127f991600484019160209091019061459b565b5050335f818152600360209081526040808320805460018101825590845291909220018490555190915061282e9084906154a9565b6040519081900381209083907fb8be57bce74a717a1bbd4acf428df655720fce75c1854b02a88484388df241a8905f90a4919050565b5f60208360ff16111561288a576040516359b452ef60e01b815260040160405180910390fd5b505f838152600560209081526040808320848452909152902054600160ff84161b1663ffffffff1615155b9392505050565b6128c461449f565b5f85856040516020016128d8929190615439565b60408051601f198184030181528282528051602091820120600180855284840190935293505f929190808301908036833701905050905081815f8151811061292257612922615465565b6020026020010181815250505f61293e825f5f89896001613254565b9050805f8151811061295257612952615465565b60200260200101519350505050949350505050565b5f828152600a60205260408082209051829182916129869086906154a9565b90815260405190819003602001902090505f5b81548110156129d757868282815481106129b5576129b5615465565b905f5260205f200154036129cf57809350600192506129d7565b600101612999565b5050935093915050565b6129e96144d4565b60045f8381526020019081526020015f20604051806101000160405290815f8201548152602001600182018054612a1f906154b4565b80601f0160208091040260200160405190810160405280929190818152602001828054612a4b906154b4565b8015612a965780601f10612a6d57610100808354040283529160200191612a96565b820191905f5260205f20905b815481529060010190602001808311612a7957829003601f168201915b5050505050815260200160028201548152602001600382018054612ab9906154b4565b80601f0160208091040260200160405190810160405280929190818152602001828054612ae5906154b4565b8015612b305780601f10612b0757610100808354040283529160200191612b30565b820191905f5260205f20905b815481529060010190602001808311612b1357829003601f168201915b50505050508152602001600482015481526020016005820154815260200160068201548152602001600782016040518060400160405290815f8201805480602002602001604051908101604052809291908181526020015f905b82821015612c32578382905f5260205f20018054612ba7906154b4565b80601f0160208091040260200160405190810160405280929190818152602001828054612bd3906154b4565b8015612c1e5780601f10612bf557610100808354040283529160200191612c1e565b820191905f5260205f20905b815481529060010190602001808311612c0157829003601f168201915b505050505081526020019060010190612b8a565b50505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015612c8757602002820191905f5260205f20905b815481526020019060010190808311612c73575b5050509190925250505090525092915050565b5f8381526002602052604081205484908203612cc95760405163938a92b760e01b815260040160405180910390fd5b5f818152600260205260409020600301546001600160a01b03163314612d0f57612cf281613572565b612d0f5760405163dc64d0ad60e01b815260040160405180910390fd5b8333604051602001612d22929190615439565b604051602081830303815290604052805190602001208514612d57576040516327a5901560e11b815260040160405180910390fd5b5f8581526002602052604090206004015415612d855760405162227f7760ea1b815260040160405180910390fd5b5f85815260026020526040812081815590612da3600183018261452e565b5f600283018190556003830180546001600160a01b0319169055612dcb906004840190614580565b5050335f90815260036020526040902080548690829086908110612df157612df1615465565b905f5260205f20015414612e18576040516337c7f25560e01b815260040160405180910390fd5b80548190612e2890600190615605565b81548110612e3857612e38615465565b905f5260205f200154818581548110612e5357612e53615465565b905f5260205f20018190555080805480612e6f57612e6f615687565b600190038181905f5260205f20015f90559055336001600160a01b031685604051612e9a91906154a9565b6040519081900381209088907feda2cc8e002ead8000b1e0c1debfc9a88bd7ee6e94b8dc0763db17849fcf0389905f90a450600195945050505050565b5f805b5f83815260046020526040902060070154811015612f1457612efc83826112a1565b15155f03612f0c57505f92915050565b600101612eda565b50600192915050565b5f8784604051602001612f31929190615498565b60408051601f1981840301815291815281516020928301205f8181526004909352908220549092509003612f7857604051632abde33960e01b815260040160405180910390fd5b5f818152600460205260408120600801805489908110612f9a57612f9a615465565b905f5260205f200154905080866001612fb391906156b2565b60ff161115612fd5576040516359b452ef60e01b815260040160405180910390fd5b60408051610100810182525f848152600460205291822060070180548291908c90811061300457613004615465565b905f5260205f20018054613017906154b4565b80601f0160208091040260200160405190810160405280929190818152602001828054613043906154b4565b801561308e5780601f106130655761010080835404028352916020019161308e565b820191905f5260205f20905b81548152906001019060200180831161307157829003601f168201915b505050505081526020018d81526020018a81526020018860ff1681526020018c81526020018981526020018581526020018b81525090506130cf8186613b12565b6130da83888b612864565b156130f857604051636d680ca160e11b815260040160405180910390fd5b50505f898b8660405160200161311093929190615670565b60408051601f1981840301815291815281516020928301205f818152600990935291205490915060ff1661314c5761314a8a8c875f611d6b565b505b61315782878a613d84565b61316182896112a1565b156131a3575f82815260066020526040812080546001929061318890849061ffff166156cb565b92506101000a81548161ffff021916908361ffff1602179055505b5f546040516340c10f1960e01b8152336004820152670de0b6b3a764000060248201526001600160a01b03909116906340c10f19906044015f604051808303815f87803b1580156131f2575f5ffd5b505af1158015613204573d5f5f3e3d5ffd5b5050604080518b815260ff8a1660208201528d93508e92507feded691eefce1ca615096347126f78807fee63b86aab874574e9c733082ce867910160405180910390a35050505050505050505050565b60605f86158015613263575085155b156132705750600161327f565b61327c87895188613ded565b90505b806001600160401b0381111561329757613297614829565b6040519080825280602002602001820160405280156132d057816020015b6132bd61449f565b8152602001906001900390816132b55790505b5091505f5b81811015613566575f896132e9838b61565d565b815181106132f9576132f9615465565b602002602001015190505f60025f8381526020019081526020015f206040518060a00160405290815f8201548152602001600182018054613339906154b4565b80601f0160208091040260200160405190810160405280929190818152602001828054613365906154b4565b80156133b05780601f10613387576101008083540402835291602001916133b0565b820191905f5260205f20905b81548152906001019060200180831161339357829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020016004820180548060200260200160405190810160405280929190818152602001828054801561343e57602002820191905f5260205f20905b81548152602001906001019080831161342a575b505050505081525050905085156134be576040518060a00160405280825f01518152602001826020015181526020018260400151815260200182606001516001600160a01b0316815260200161349983608001518b8b613e3f565b8152508584815181106134ae576134ae615465565b602002602001018190525061355c565b6040518060a00160405280825f01518152602001826020015181526020018260400151815260200182606001516001600160a01b031681526020015f6001600160401b0381111561351157613511614829565b60405190808252806020026020018201604052801561353a578160200160208202803683370190505b5081525085848151811061355057613550615465565b60200260200101819052505b50506001016132d5565b50509695505050505050565b6001545f906001600160a01b0316156136035760015460405163e124bdd960e01b8152600481018490523360248201526001600160a01b039091169063e124bdd990604401602060405180830381865afa1580156135d2573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906135f691906156e5565b1561360357506001919050565b505f919050565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061369057507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166136845f5160206157db5f395f51905f52546001600160a01b031690565b6001600160a01b031614155b156136ae5760405163703e46dd60e11b815260040160405180910390fd5b565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561370a575060408051601f3d908101601f1916820190925261370791810190615700565b60015b61373257604051634c9c8ce360e01b81526001600160a01b0383166004820152602401611d28565b5f5160206157db5f395f51905f52811461376257604051632a87526960e21b815260048101829052602401611d28565b610d608383613ef1565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146136ae5760405163703e46dd60e11b815260040160405180910390fd5b5f5f83836040516020016137ca929190615498565b60408051601f1981840301815291815281516020928301205f81815260049093529120549091501561380f576040516303448eef60e51b815260040160405180910390fd5b5f84815260026020908152604080832060040180546001810182559084529190922001829055516138419084906154a9565b60405190819003812033825290859083907fb018e47bdb983351e1bee22415a8f41eef5c2bf1c43c6d3d0992e678ae762df59060200160405180910390a4604080515f81830181815260608301909352918190836138af565b606081526020019060019003908161389a5790505b5081526020015f6040519080825280602002602001820160405280156138df578160200160208202803683370190505b50905260408051610100810182528481528151602081810184525f8083528184019283528385018b9052606084018a9052608084018190524260a085015260c0840181905260e084018690528781526004909152929092208151815591519293509160018201906139509082615537565b50604082015160028201556060820151600382019061396f9082615537565b506080820151600482015560a0820151600582015560c0820151600682015560e08201518051805160078401916139ab918391602001906145e4565b5060208281015180516139c4926001850192019061459b565b509498975050505050505050565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10280546060915f51602061572c5f395f51905f5291613a10906154b4565b80601f0160208091040260200160405190810160405280929190818152602001828054613a3c906154b4565b8015613a875780601f10613a5e57610100808354040283529160200191613a87565b820191905f5260205f20905b815481529060010190602001808311613a6a57829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10380546060915f51602061572c5f395f51905f5291613a10906154b4565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0061129b565b613b00613f46565b6115718282613f6b565b6136ae613f46565b8160c00151421115613b5a5760405162461bcd60e51b815260206004820152601160248201527014da59db985d1d5c9948195e1c1a5c9959607a1b6044820152606401611d28565b60e08201515f90815260026020908152604080832060030154815160c08101909252608f8083526001600160a01b03909116939261574c9083013980519060200120845f01518051906020012085602001518660400151876060015188608001518960a001518a60c001518b60e00151604051602001613c1f9998979695949392919098895260208901979097526040880195909552606087019390935260ff91909116608086015260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012090505f613c4182613fca565b90505f613c4e8286613ff6565b9050836001600160a01b0316816001600160a01b031614613cc75760405162461bcd60e51b815260206004820152602d60248201527f496e76616c6964207369676e61747572653a204e6f74207369676e656420627960448201526c10313ab1b5b2ba1037bbb732b960991b6064820152608401611d28565b6001600160a01b0381165f908152600c6020908152604080832060e08a01518452825280832060a08a0151845290915290205460ff1615613d3f5760405162461bcd60e51b8152602060048201526012602482015271139bdb98d948185b1c9958591e481d5cd95960721b6044820152606401611d28565b6001600160a01b03165f908152600c6020908152604080832060e08901518452825280832060a090980151835296905294909420805460ff1916600117905550505050565b60208260ff1610613da8576040516359b452ef60e01b815260040160405180910390fd5b5f928352600560209081526040808520928552919052909120805463ffffffff600160ff9094169390931b83169281169290921763ffffffff19909216919091179055565b5f82841115613e0f576040516309605a0160e41b815260040160405180910390fd5b5f613e1a838661565d565b905083811115613e3657613e2e8585615605565b9150506128b5565b829150506128b5565b60605f613e4e84865185613ded565b90505f816001600160401b03811115613e6957613e69614829565b604051908082528060200260200182016040528015613e92578160200160208202803683370190505b5090505f5b82811015613ee75786613eaa828861565d565b81518110613eba57613eba615465565b6020026020010151828281518110613ed457613ed4615465565b6020908102919091010152600101613e97565b5095945050505050565b613efa8261401e565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115613f3e57610d608282614081565b6115716140f3565b613f4e614112565b6136ae57604051631afcd79f60e31b815260040160405180910390fd5b613f73613f46565b5f51602061572c5f395f51905f527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d102613fac8482615537565b5060038101613fbb8382615537565b505f8082556001909101555050565b5f61129b613fd661412b565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f6140048686614139565b9250925092506140148282614182565b5090949350505050565b806001600160a01b03163b5f0361405357604051634c9c8ce360e01b81526001600160a01b0382166004820152602401611d28565b5f5160206157db5f395f51905f5280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f5f846001600160a01b03168460405161409d91906154a9565b5f60405180830381855af49150503d805f81146140d5576040519150601f19603f3d011682016040523d82523d5f602084013e6140da565b606091505b50915091506140ea85838361423a565b95945050505050565b34156136ae5760405163b398979f60e01b815260040160405180910390fd5b5f61411b613ad0565b54600160401b900460ff16919050565b5f614134614296565b905090565b5f5f5f8351604103614170576020840151604085015160608601515f1a61416288828585614309565b95509550955050505061417b565b505081515f91506002905b9250925092565b5f82600381111561419557614195615717565b0361419e575050565b60018260038111156141b2576141b2615717565b036141d05760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156141e4576141e4615717565b036142055760405163fce698f760e01b815260048101829052602401611d28565b600382600381111561421957614219615717565b03611571576040516335e2f38360e21b815260048101829052602401611d28565b60608261424f5761424a826143cd565b6128b5565b815115801561426657506001600160a01b0384163b155b1561428f57604051639996b31560e01b81526001600160a01b0385166004820152602401611d28565b50806128b5565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6142c06143f5565b6142c861445d565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561434257505f91506003905082612370565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015614393573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166143be57505f925060019150829050612370565b975f9750879650945050505050565b8051156143dc57805160208201fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f5f51602061572c5f395f51905f528161440d6139d2565b80519091501561442557805160209091012092915050565b81548015614434579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b5f5f51602061572c5f395f51905f5281614475613a92565b80519091501561448d57805160209091012092915050565b60018201548015614434579392505050565b6040518060a001604052805f8152602001606081526020015f81526020015f6001600160a01b03168152602001606081525090565b6040518061010001604052805f8152602001606081526020015f8152602001606081526020015f81526020015f81526020015f8152602001614529604051806040016040528060608152602001606081525090565b905290565b50805461453a906154b4565b5f825580601f10614549575050565b601f0160209004905f5260205f2090810190610a3e9190614634565b5080545f8255905f5260205f2090810190610a3e9190614648565b5080545f8255905f5260205f2090810190610a3e9190614634565b828054828255905f5260205f209081019282156145d4579160200282015b828111156145d45782518255916020019190600101906145b9565b506145e0929150614634565b5090565b828054828255905f5260205f20908101928215614628579160200282015b8281111561462857825182906146189082615537565b5091602001919060010190614602565b506145e0929150614648565b5b808211156145e0575f8155600101614635565b808211156145e0575f61465b828261452e565b50600101614648565b5f60208284031215614674575f5ffd5b5035919050565b5f6020828403121561468b575f5ffd5b81356001600160401b038111156146a0575f5ffd5b820161012081850312156128b5575f5ffd5b80356001600160a01b03811681146146c8575f5ffd5b919050565b5f5f5f5f5f60a086880312156146e1575f5ffd5b6146ea866146b2565b97602087013597506040870135966060810135965060800135945092505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f8151808452602084019350602083015f5b8281101561476957815186526020958601959091019060010161474b565b5093949350505050565b805182525f602082015160a0602085015261479160a085018261470b565b90506040830151604085015260018060a01b036060840151166060850152608083015184820360808601526140ea8282614739565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561481d57603f19878603018452614808858351614773565b945060209384019391909101906001016147ec565b50929695505050505050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b038111828210171561486557614865614829565b604052919050565b5f82601f83011261487c575f5ffd5b8135602083015f5f6001600160401b0384111561489b5761489b614829565b50601f8301601f19166020016148b08161483d565b9150508281528583830111156148c4575f5ffd5b828260208301375f92810160200192909252509392505050565b5f5f604083850312156148ef575f5ffd5b82356001600160401b03811115614904575f5ffd5b6149108582860161486d565b92505061491f602084016146b2565b90509250929050565b5f5f83601f840112614938575f5ffd5b5081356001600160401b0381111561494e575f5ffd5b6020830191508360208260051b8501011115610b4c575f5ffd5b5f5f60208385031215614979575f5ffd5b82356001600160401b0381111561498e575f5ffd5b61499a85828601614928565b90969095509350505050565b5f5f5f5f5f5f5f5f5f60e08a8c0312156149be575f5ffd5b89356001600160401b038111156149d3575f5ffd5b6149df8c828d0161486d565b99505060208a0135975060408a01356001600160401b03811115614a01575f5ffd5b614a0d8c828d0161486d565b97505060608a0135955060808a01356001600160401b03811115614a2f575f5ffd5b614a3b8c828d01614928565b90965094505060a08a01356001600160401b03811115614a59575f5ffd5b614a658c828d01614928565b9a9d999c50979a9699959894979660c00135949350505050565b5f5f60408385031215614a90575f5ffd5b50508035926020909101359150565b5f5f60408385031215614ab0575f5ffd5b8235915060208301356001600160401b03811115614acc575f5ffd5b614ad88582860161486d565b9150509250929050565b602080825282518282018190525f918401906040840190835b81811015614b19578351835260209384019390920191600101614afb565b509095945050505050565b5f5f5f60608486031215614b36575f5ffd5b83356001600160401b03811115614b4b575f5ffd5b614b578682870161486d565b93505060208401359150614b6d604085016146b2565b90509250925092565b604081525f614b88604083018561470b565b90508260208301529392505050565b5f5f60408385031215614ba8575f5ffd5b614bb1836146b2565b915060208301356001600160401b03811115614acc575f5ffd5b5f6001600160401b03821115614be357614be3614829565b5060051b60200190565b5f82601f830112614bfc575f5ffd5b8135614c0f614c0a82614bcb565b61483d565b8082825260208201915060208360051b860101925085831115614c30575f5ffd5b602085015b83811015613ee75780356001600160401b03811115614c52575f5ffd5b614c61886020838a010161486d565b84525060209283019201614c35565b5f82601f830112614c7f575f5ffd5b8135614c8d614c0a82614bcb565b8082825260208201915060208360051b860101925085831115614cae575f5ffd5b602085015b83811015613ee7578035835260209283019201614cb3565b5f5f5f5f5f5f5f5f5f60e08a8c031215614ce3575f5ffd5b89356001600160401b03811115614cf8575f5ffd5b614d048c828d01614bed565b99505060208a0135975060408a01356001600160401b03811115614d26575f5ffd5b614d328c828d0161486d565b97505060608a01356001600160401b03811115614d4d575f5ffd5b614d598c828d01614c70565b96505060808a01356001600160401b03811115614a2f575f5ffd5b602081525f6128b5602083018461470b565b5f6040830182516040855281815180845260608701915060608160051b88010193506020830192505f5b81811015614de157605f19888603018352614dcc85855161470b565b94506020938401939290920191600101614db0565b50505050602083015184820360208601526140ea8282614739565b805182525f60208201516101006020850152614e1c61010085018261470b565b90506040830151604085015260608301518482036060860152614e3f828261470b565b9150506080830151608085015260a083015160a085015260c083015160c085015260e083015184820360e08601526140ea8282614d86565b602081525f6128b56020830184614dfc565b5f5f5f5f5f5f60a08789031215614e9e575f5ffd5b86356001600160401b03811115614eb3575f5ffd5b614ebf89828a01614928565b909a90995060208901359860408101359850606081013597506080013595509350505050565b5f5f5f5f5f60a08688031215614ef9575f5ffd5b85359450602086013593506040860135925060608601356001600160401b03811115614f23575f5ffd5b614f2f8882890161486d565b95989497509295608001359392505050565b60ff60f81b8816815260e060208201525f614f5f60e083018961470b565b8281036040840152614f71818961470b565b606084018890526001600160a01b038716608085015260a0840186905283810360c08501529050614fa28185614739565b9a9950505050505050505050565b8015158114610a3e575f5ffd5b5f5f5f5f60808587031215614fd0575f5ffd5b843593506020850135925060408501356001600160401b03811115614ff3575f5ffd5b614fff8782880161486d565b925050606085013561501081614fb0565b939692955090935050565b5f5f5f5f5f60a0868803121561502f575f5ffd5b8535945060208601356001600160401b0381111561504b575f5ffd5b6150578882890161486d565b945050604086013592506060860135915060808601356001600160401b03811115615080575f5ffd5b61508c8882890161486d565b9150509295509295909350565b5f602082840312156150a9575f5ffd5b6128b5826146b2565b5f5f5f5f608085870312156150c5575f5ffd5b84356001600160401b038111156150da575f5ffd5b6150e68782880161486d565b94505060208501356001600160401b03811115615101575f5ffd5b61510d8782880161486d565b93505060408501359150615123606086016146b2565b905092959194509250565b606081525f6151406060830186614dfc565b602083019490945250901515604090910152919050565b5f5f5f60408486031215615169575f5ffd5b83356001600160401b0381111561517e575f5ffd5b61518a86828701614928565b90945092505060208401356001600160401b038111156151a8575f5ffd5b6151b48682870161486d565b9150509250925092565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561481d57603f19878603018452615200858351614739565b945060209384019391909101906001016151e4565b5f5f5f5f60808587031215615228575f5ffd5b843593506020850135925060408501356001600160401b0381111561524b575f5ffd5b6152578782880161486d565b949793965093946060013593505050565b5f60208284031215615278575f5ffd5b81356001600160401b0381111561528d575f5ffd5b611c688482850161486d565b803560ff811681146146c8575f5ffd5b5f5f5f606084860312156152bb575f5ffd5b833592506152cb60208501615299565b929592945050506040919091013590565b5f5f5f5f608085870312156152ef575f5ffd5b84356001600160401b03811115615304575f5ffd5b6153108782880161486d565b94505061531f602086016146b2565b93969395505050506040820135916060013590565b602081525f6128b56020830184614773565b5f5f5f60608486031215615358575f5ffd5b833592506020840135915060408401356001600160401b038111156151a8575f5ffd5b5f5f5f6060848603121561538d575f5ffd5b8335925060208401356001600160401b038111156153a9575f5ffd5b6153b58682870161486d565b93969395505050506040919091013590565b5f602082840312156153d7575f5ffd5b6128b582615299565b5f5f8335601e198436030181126153f5575f5ffd5b8301803591506001600160401b0382111561540e575f5ffd5b602001915036819003821315610b4c575f5ffd5b5f81518060208401855e5f93019283525090919050565b5f6154448285615422565b60609390931b6bffffffffffffffffffffffff191683525050601401919050565b634e487b7160e01b5f52603260045260245ffd5b5f823561011e1983360301811261548e575f5ffd5b9190910192915050565b8281525f611c686020830184615422565b5f6128b58284615422565b600181811c908216806154c857607f821691505b6020821081036154e657634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115610d6057805f5260205f20601f840160051c810160208510156155115750805b601f840160051c820191505b81811015615530575f815560010161551d565b5050505050565b81516001600160401b0381111561555057615550614829565b6155648161555e84546154b4565b846154ec565b6020601f821160018114615596575f831561557f5750848201515b5f19600385901b1c1916600184901b178455615530565b5f84815260208120601f198516915b828110156155c557878501518255602094850194600190920191016155a5565b50848210156155e257868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561129b5761129b6155f1565b5f5f8335601e1984360301811261562d575f5ffd5b8301803591506001600160401b03821115615646575f5ffd5b6020019150600581901b3603821315610b4c575f5ffd5b8082018082111561129b5761129b6155f1565b8381528260208201525f6140ea6040830184615422565b634e487b7160e01b5f52603160045260245ffd5b808202811582820484141761129b5761129b6155f1565b60ff818116838216019081111561129b5761129b6155f1565b61ffff818116838216019081111561129b5761129b6155f1565b5f602082840312156156f5575f5ffd5b81516128b581614fb0565b5f60208284031215615710575f5ffd5b5051919050565b634e487b7160e01b5f52602160045260245ffdfea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d10053746f7261676544617461286279746573206368756e6b4349442c6279746573333220626c6f636b4349442c75696e74323536206368756e6b496e6465782c75696e743820626c6f636b496e6465782c62797465733332206e6f646549642c75696e74323536206e6f6e63652c75696e7432353620646561646c696e652c62797465733332206275636b6574496429360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbca2646970667358221220e2da43486c3100726e61a570bbde303ad49c67311dd40a065707e70edf8ccb2164736f6c634300081c0033",
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
