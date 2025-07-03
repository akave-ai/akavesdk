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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BlockAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockAlreadyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockNonexists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketInvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketNonempty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketNonexists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"}],\"name\":\"ChunkCIDMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileChunkDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileFullyUploaded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNameDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNonempty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNotFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cidsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sizesLength\",\"type\":\"uint256\"}],\"name\":\"InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlocksAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEncodedSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFileBlocksCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFileCID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLastBlockSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LastChunkDuplicate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AddFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"AddFileBlocks\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"}],\"name\":\"AddPeerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"}],\"name\":\"ChunkBlockFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CreateBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CreateFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DeleteBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DeleteFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"}],\"name\":\"DeletePeerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"FileUploaded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_BLOCKS_PER_FILE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BLOCK_SIZE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessManager\",\"outputs\":[{\"internalType\":\"contractIAccessManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"cid\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedChunkSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"cids\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkBlocksSizes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"addFileChunk\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isReplica\",\"type\":\"bool\"}],\"name\":\"addPeerBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedFileSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"}],\"name\":\"commitFile\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"createBucket\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"createFile\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deleteBucket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deleteFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deletePeerBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fileFillCounter\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fileRewardClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockCID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nodeId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"fillChunkBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fulfilledBlocks\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getBucketByName\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket\",\"name\":\"bucket\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getBucketIndexByName\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"}],\"name\":\"getBucketsByIds\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getChunkByIndex\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getFileById\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"chunkCIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkSize\",\"type\":\"uint256[]\"}],\"internalType\":\"structIStorage.Chunk\",\"name\":\"chunks\",\"type\":\"tuple\"}],\"internalType\":\"structIStorage.File\",\"name\":\"file\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getFileByName\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"chunkCIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkSize\",\"type\":\"uint256[]\"}],\"internalType\":\"structIStorage.Chunk\",\"name\":\"chunks\",\"type\":\"tuple\"}],\"internalType\":\"structIStorage.File\",\"name\":\"file\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"getFileIndexById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getFileOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getOwnerBuckets\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"buckets\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getPeerBlockIndexById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"cids\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getPeersArrayByPeerBlockCid\",\"outputs\":[{\"internalType\":\"bytes32[][]\",\"name\":\"peers\",\"type\":\"bytes32[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getPeersByPeerBlockCid\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"peers\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"isBlockFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"isChunkFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"isFileFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"isFileFilledV2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"peerId\",\"type\":\"bytes32\"}],\"name\":\"isPeerBlockReplica\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accessManagerAddress\",\"type\":\"address\"}],\"name\":\"setAccessManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIAkaveToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x610160604052348015610010575f5ffd5b506040516145fa3803806145fa83398101604081905261002f916101b1565b604080518082018252600781526653746f7261676560c81b602080830191909152825180840190935260018352603160f81b9083015290610070825f610139565b6101205261007f816001610139565b61014052815160208084019190912060e052815190820120610100524660a05261010b60e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c052600280546001600160a01b0319166001600160a01b0392909216919091179055610388565b5f6020835110156101545761014d8361016b565b9050610165565b8161015f8482610276565b5060ff90505b92915050565b5f5f829050601f8151111561019e578260405163305a27a960e01b81526004016101959190610330565b60405180910390fd5b80516101a982610365565b179392505050565b5f602082840312156101c1575f5ffd5b81516001600160a01b03811681146101d7575f5ffd5b9392505050565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061020657607f821691505b60208210810361022457634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561027157805f5260205f20601f840160051c8101602085101561024f5750805b601f840160051c820191505b8181101561026e575f815560010161025b565b50505b505050565b81516001600160401b0381111561028f5761028f6101de565b6102a38161029d84546101f2565b8461022a565b6020601f8211600181146102d5575f83156102be5750848201515b5f19600385901b1c1916600184901b17845561026e565b5f84815260208120601f198516915b8281101561030457878501518255602094850194600190920191016102e4565b508482101561032157868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610224575f1960209190910360031b1b16919050565b60805160a05160c05160e0516101005161012051610140516142216103d95f395f612b8501525f612b5401525f61309e01525f61307601525f612fd101525f612ffb01525f61302501526142215ff3fe608060405234801561000f575f5ffd5b5060043610610229575f3560e01c80636ce023631161012a578063d6d3110b116100b4578063faec054211610079578063faec0542146105fe578063fc0c546a14610611578063fd1d3c0c14610624578063fd21c28414610637578063fdcb60681461064a575f5ffd5b8063d6d3110b1461057c578063e3f787e81461058f578063e4ba8a58146105a2578063f855169a146105b5578063f8a3e41a146105c8575f5ffd5b80639a2e82b3116100fa5780639a2e82b31461050a5780639ccd46461461051d578063b80777ea14610526578063c95808041461052c578063d606205d1461055c575f5ffd5b80636ce02363146104a75780637912bf68146104c957806384b0196e146104dc57806387c1ac07146104f7575f5ffd5b80634f55ba82116101b65780636554cda71161017b5780636554cda71461041957806368e6408f146104395780636a5d8c261461046e5780636a6e658b146104815780636af0f80114610494575f5ffd5b80634f55ba821461037257806354fd4d5014610387578063564b81ef146103ae5780635a4e9564146103b45780635ecdfb53146103f9575f5ffd5b8063359b15a5116101fc578063359b15a5146102eb57806335bdb711146102fe5780633f3839801461031e57806348b49875146103315780634d7dc61414610351575f5ffd5b8063018c1e9c1461022d5780631b475ef414610264578063287e677f146102b757806330b91d07146102d8575b5f5ffd5b61024f61023b366004613475565b600a6020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61029f610272366004613475565b5f90815260066020908152604080832060020154835260049091529020600301546001600160a01b031690565b6040516001600160a01b03909116815260200161025b565b6102ca6102c536600461354a565b61065d565b60405190815260200161025b565b6102ca6102e63660046135d4565b6106e9565b6102ca6102f93660046136ad565b610992565b61031161030c3660046136ee565b610a0e565b60405161025b91906137f0565b61024f61032c366004613853565b610c27565b61034461033f366004613873565b610ce4565b60405161025b91906138b6565b61036461035f366004613853565b610d62565b60405161025b9291906138f8565b610385610380366004613929565b610e5c565b005b60408051808201825260058152640312e302e360dc1b6020820152905161025b91906139d6565b466102ca565b6103e46103c2366004613853565b600760209081525f928352604080842090915290825290205463ffffffff1681565b60405163ffffffff909116815260200161025b565b61040c610407366004613873565b610e78565b60405161025b9190613a65565b61042c610427366004613af0565b611157565b60405161025b9190613b29565b61024f610447366004613475565b5f9081526008602090815260408083205460069092529091206007015461ffff9091161490565b61034461047c366004613b3b565b6112fe565b61024f61048f366004613b54565b611367565b6102ca6104a2366004613873565b611515565b6104b1620f424081565b6040516001600160401b03909116815260200161025b565b61024f6104d7366004613853565b6117a7565b6104e46117f4565b60405161025b9796959493929190613bb0565b6102ca610505366004613c2c565b611836565b6102ca610518366004613c8a565b611938565b6104b161040081565b426102ca565b61038561053a366004613b3b565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b61056f61056a366004613d08565b611c92565b60405161025b9190613d6f565b61024f61058a366004613dc6565b611d9d565b6102ca61059d366004613af0565b61200f565b61024f6105b0366004613e19565b612175565b6102ca6105c3366004613e4c565b6121cc565b6105eb6105d6366004613475565b60086020525f908152604090205461ffff1681565b60405161ffff909116815260200161025b565b61040c61060c366004613475565b612240565b60025461029f906001600160a01b031681565b61024f610632366004613e81565b6124f9565b61024f610645366004613475565b612737565b60035461029f906001600160a01b031681565b5f5f8383604051602001610672929190613ee4565b60408051601f1981840301815291815281516020928301206001600160a01b0386165f9081526005909352908220909250905b81548110156106df57828282815481106106c1576106c1613f10565b905f5260205f200154036106d7578093506106df565b6001016106a5565b5050505b92915050565b5f88815260046020526040812054899082036107185760405163938a92b760e01b815260040160405180910390fd5b5f818152600460205260409020600301546001600160a01b0316331461075e576107418161277d565b61075e5760405163dc64d0ad60e01b815260040160405180910390fd5b89895f8282604051602001610774929190613f24565b60408051601f1981840301815291815281516020928301205f8181526006909352912060040154909150156107bc5760405163d96b03b160e01b815260040160405180910390fd5b5f8d8d6040516020016107d0929190613f24565b60408051601f1981840301815291815281516020928301205f818152600690935290822054909250900361081757604051632abde33960e01b815260040160405180910390fd5b5f818152600660205260409020600701548714610847576040516301c0b3dd60e61b815260040160405180910390fd5b8988141580610856575060208a115b15610874576040516373d8ccd360e11b815260040160405180910390fd5b60208a10156108ca575f8181526009602052604090205461ffff16156108ad576040516355cbc83160e01b815260040160405180910390fd5b5f818152600960205260409020805461ffff191661ffff8c161790555b8c6040516108d89190613f35565b604051908190038120338252908f9083907f01d10894cb2a39778aae51e234b669f70a74328f07e58e67a2caca4c5a3c86ff9060200160405180910390a460065f8281526020019081526020015f206007015f018f908060018154018082558091505060019003905f5260205f20015f90919091909150908161095b9190613fc4565b505f818152600660209081526040822060080180546001810182559083529120018c90559450505050509998505050505050505050565b5f5f83336040516020016109a7929190613ee4565b60408051601f1981840301815291815281516020928301205f818152600493849052918220909350909101905b81548110156106df57848282815481106109f0576109f0613f10565b905f5260205f20015403610a06578093506106df565b6001016109d4565b60605f826001600160401b03811115610a2957610a2961348c565b604051908082528060200260200182016040528015610a9257816020015b610a7f6040518060a001604052805f8152602001606081526020015f81526020015f6001600160a01b03168152602001606081525090565b815260200190600190039081610a475790505b5090505f5b83811015610c1f5760045f868684818110610ab457610ab4613f10565b9050602002013581526020019081526020015f206040518060a00160405290815f8201548152602001600182018054610aec90613f40565b80601f0160208091040260200160405190810160405280929190818152602001828054610b1890613f40565b8015610b635780601f10610b3a57610100808354040283529160200191610b63565b820191905f5260205f20905b815481529060010190602001808311610b4657829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b0316815260200160048201805480602002602001604051908101604052809291908181526020018280548015610bf157602002820191905f5260205f20905b815481526020019060010190808311610bdd575b505050505081525050828281518110610c0c57610c0c613f10565b6020908102919091010152600101610a97565b509392505050565b5f828152600660205260408120600701548190610c4690600190614092565b90508083148015610c6757505f8481526009602052604090205461ffff1615155b15610cbd575f84815260096020526040812054610c8d9060019061ffff1681901b614092565b5f868152600760209081526040808320888452909152902054811663ffffffff90811691161492506106e3915050565b50505f91825260076020908152604080842092845291905290205463ffffffff9081161490565b6060600c5f8481526020019081526020015f2082604051610d059190613f35565b9081526040805191829003602090810183208054808302850183019093528284529190830182828015610d5557602002820191905f5260205f20905b815481526020019060010190808311610d41575b5050505050905092915050565b5f828152600660205260408120600701805460609291829185908110610d8a57610d8a613f10565b905f5260205f20018054610d9d90613f40565b80601f0160208091040260200160405190810160405280929190818152602001828054610dc990613f40565b8015610e145780601f10610deb57610100808354040283529160200191610e14565b820191905f5260205f20905b815481529060010190602001808311610df757829003601f168201915b5050505f888152600660205260408120600801805494955090939092508791508110610e4257610e42613f10565b5f91825260209091200154919350909150505b9250929050565b610e6d898989898989898989612815565b505050505050505050565b610e806132e2565b5f8383604051602001610e94929190613f24565b60408051601f1981840301815282825280516020918201205f818152600683528390206101008501909352825484526001830180549195509184019190610eda90613f40565b80601f0160208091040260200160405190810160405280929190818152602001828054610f0690613f40565b8015610f515780601f10610f2857610100808354040283529160200191610f51565b820191905f5260205f20905b815481529060010190602001808311610f3457829003601f168201915b5050505050815260200160028201548152602001600382018054610f7490613f40565b80601f0160208091040260200160405190810160405280929190818152602001828054610fa090613f40565b8015610feb5780601f10610fc257610100808354040283529160200191610feb565b820191905f5260205f20905b815481529060010190602001808311610fce57829003601f168201915b50505050508152602001600482015481526020016005820154815260200160068201548152602001600782016040518060400160405290815f8201805480602002602001604051908101604052809291908181526020015f905b828210156110ed578382905f5260205f2001805461106290613f40565b80601f016020809104026020016040519081016040528092919081815260200182805461108e90613f40565b80156110d95780601f106110b0576101008083540402835291602001916110d9565b820191905f5260205f20905b8154815290600101906020018083116110bc57829003601f168201915b505050505081526020019060010190611045565b5050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561114257602002820191905f5260205f20905b81548152602001906001019080831161112e575b50505091909252505050905250949350505050565b61118f6040518060a001604052805f8152602001606081526020015f81526020015f6001600160a01b03168152602001606081525090565b5f82336040516020016111a3929190613ee4565b60408051601f1981840301815282825280516020918201205f8181526004835283902060a085019093528254845260018301805491955091840191906111e890613f40565b80601f016020809104026020016040519081016040528092919081815260200182805461121490613f40565b801561125f5780601f106112365761010080835404028352916020019161125f565b820191905f5260205f20905b81548152906001019060200180831161124257829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152602001600482018054806020026020016040519081016040528092919081815260200182805480156112ed57602002820191905f5260205f20905b8154815260200190600101908083116112d9575b505050505081525050915050919050565b6001600160a01b0381165f9081526005602090815260409182902080548351818402810184019094528084526060939283018282801561135b57602002820191905f5260205f20905b815481526020019060010190808311611347575b50505050509050919050565b5f858152600b602052604081205460ff1661139557604051631512312160e01b815260040160405180910390fd5b8484846040516020016113aa939291906140a5565b6040516020818303038152906040528051906020012086146113df576040516332c83a2360e21b815260040160405180910390fd5b5f8585856040516020016113f5939291906140a5565b60405160208183030381529060405280519060200120905086811461142d57604051637ae9080f60e11b815260040160405180910390fd5b5f878152600b60209081526040808320805460ff19169055878352600c909152808220905161145d908790613f35565b90815260200160405180910390209050806001828054905061147f9190614092565b8154811061148f5761148f613f10565b905f5260205f2001548185815481106114aa576114aa613f10565b905f5260205f200181905550808054806114c6576114c66140bc565b600190038181905f5260205f20015f9055905586887f37505c6d2cdf9e778d6110c5ad2e49c649d521a18d2da6f007d3364bd83a45ae60405160405180910390a3506001979650505050505050565b5f82815260046020526040812054839082036115445760405163938a92b760e01b815260040160405180910390fd5b5f818152600460205260409020600301546001600160a01b0316331461158a5761156d8161277d565b61158a5760405163dc64d0ad60e01b815260040160405180910390fd5b5f848460405160200161159e929190613f24565b60408051601f1981840301815291815281516020928301205f8181526006909352912054909150156115e3576040516303448eef60e51b815260040160405180910390fd5b5f8581526004602081815260408084209092018054600181018255908452922090910182905551611615908590613f35565b60405190819003812033825290869083907fb018e47bdb983351e1bee22415a8f41eef5c2bf1c43c6d3d0992e678ae762df59060200160405180910390a4604080515f8183018181526060830190935291819083611683565b606081526020019060019003908161166e5790505b5081526020015f6040519080825280602002602001820160405280156116b3578160200160208202803683370190505b50905260408051610100810182528481528151602081810184525f8083528184019283528385018c9052606084018b9052608084018190524260a085015260c0840181905260e084018690528781526006909152929092208151815591519293509160018201906117249082613fc4565b5060408201516002820155606082015160038201906117439082613fc4565b506080820151600482015560a0820151600582015560c0820151600682015560e082015180518051600784019161177f9183916020019061333c565b5060208281015180516117989260018501920190613390565b50949998505050505050505050565b5f5f82846040516020016117c5929190918252602082015260400190565b60408051808303601f1901815291815281516020928301205f908152600d90925290205460ff16949350505050565b5f6060805f5f5f6060611805612b4d565b61180d612b7e565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b5f84848460405160200161184c939291906140a5565b60408051601f19818403018152828252805160209182012090830188905290820186905291505f9060600160408051601f1981840301815291815281516020928301205f858152600b9093529120805460ff19166001179055905082156118c6575f818152600d60205260409020805460ff191660011790555b5f858152600c60205260409081902090516118e2908690613f35565b908152604051908190036020908101822080546001810182555f918252918120909101889055879184917f9f91e060f6a60bc2a5e695f07d854d6a659eab9f5e1441ac24f6e62614bb8b5b91a350949350505050565b5f85815260046020526040812054869082036119675760405163938a92b760e01b815260040160405180910390fd5b5f818152600460205260409020600301546001600160a01b031633146119ad576119908161277d565b6119ad5760405163dc64d0ad60e01b815260040160405180910390fd5b86865f82826040516020016119c3929190613f24565b60408051601f1981840301815291815281516020928301205f818152600690935291206004015490915015611a0b5760405163d96b03b160e01b815260040160405180910390fd5b5f8a8a604051602001611a1f929190613f24565b60408051601f1981840301815291815281516020928301205f81815260088452828120546006909452919091206007015490925061ffff90911614611a7757604051632e1b8f4960e11b815260040160405180910390fd5b885f03611a9757604051631b6fdfeb60e01b815260040160405180910390fd5b86515f03611ab857604051637f19edc960e11b815260040160405180910390fd5b5f805b5f83815260066020526040902060080154811015611b13575f838152600660205260409020600801805482908110611af557611af5613f10565b905f5260205f20015482611b0991906140d0565b9150600101611abb565b50898114611b3457604051631b6fdfeb60e01b815260040160405180910390fd5b5f828152600660205260409020600101611b4e8982613fc4565b505f828152600660208181526040808420600481018f90559092018c9055600a905290205460ff16611c37575f828152600a60208181526040808420805460ff1916600117905560069091528220600701549190611bad9083906140e3565b611bbf90670de0b6b3a76400006140e3565b90505f82118015611bcf57505f81115b15611c34576002546040516340c10f1960e01b8152336004820152602481018390526001600160a01b03909116906340c10f19906044015f604051808303815f87803b158015611c1d575f5ffd5b505af1158015611c2f573d5f5f3e3d5ffd5b505050505b50505b8a604051611c459190613f35565b604051908190038120338252908d9084907fb2e54a2138908ebdd7da28708ec0bc3f1498b96fb7b0db337edef5ceeb41b16b9060200160405180910390a4509a9950505050505050505050565b6060826001600160401b03811115611cac57611cac61348c565b604051908082528060200260200182016040528015611cdf57816020015b6060815260200190600190039081611cca5790505b5090505f5b83811015610c1f57600c5f868684818110611d0157611d01613f10565b9050602002013581526020019081526020015f2083604051611d239190613f35565b9081526040805191829003602090810183208054808302850183019093528284529190830182828015611d7357602002820191905f5260205f20905b815481526020019060010190808311611d5f575b5050505050828281518110611d8a57611d8a613f10565b6020908102919091010152600101611ce4565b5f8381526004602052604081205484908203611dcc5760405163938a92b760e01b815260040160405180910390fd5b5f818152600460205260409020600301546001600160a01b03163314611e1257611df58161277d565b611e125760405163dc64d0ad60e01b815260040160405180910390fd5b5f868152600660205260408120549003611e3f57604051632abde33960e01b815260040160405180910390fd5b8484604051602001611e52929190613f24565b604051602081830303815290604052805190602001208614611e8757604051630ef4797b60e31b815260040160405180910390fd5b5f86815260066020526040812081815590611ea560018301826133d5565b600282015f9055600382015f611ebb91906133d5565b5f6004830181905560058301819055600683018190556007830190611ee0828261340f565b611eed600183015f61342a565b5050505f86815260046020819052604090912001805490915084101580611f2e575086818581548110611f2257611f22613f10565b905f5260205f20015414155b15611f4c576040516337c7f25560e01b815260040160405180910390fd5b80548190611f5c90600190614092565b81548110611f6c57611f6c613f10565b905f5260205f200154818581548110611f8757611f87613f10565b905f5260205f20018190555080805480611fa357611fa36140bc565b600190038181905f5260205f20015f9055905584604051611fc49190613f35565b60405190819003812033825290879089907f0e1bf50f5cca6659c62146db5b60160121a3752011b621d8c8953a1e0e23d0489060200160405180910390a45060019695505050505050565b5f8133604051602001612023929190613ee4565b60408051601f1981840301815291815281516020928301205f818152600490935291205490915015612068576040516324bf796160e11b815260040160405180910390fd5b6040805160a0810182528281526020808201858152428385015233606084015283515f80825281840186526080850191909152858152600490925292902081518155915190919060018201906120be9082613fc4565b506040820151600282015560608201516003820180546001600160a01b0319166001600160a01b039092169190911790556080820151805161210a916004840191602090910190613390565b5050335f818152600560209081526040808320805460018101825590845291909220018490555190915061213f908490613f35565b6040519081900381209083907fb8be57bce74a717a1bbd4acf428df655720fce75c1854b02a88484388df241a8905f90a4919050565b5f60208360ff16111561219b576040516359b452ef60e01b815260040160405180910390fd5b505f838152600760209081526040808320848452909152902054600160ff84161b1663ffffffff1615159392505050565b5f828152600c602052604080822090518291906121ea908590613f35565b90815260405190819003602001902090505f5b8154811015612237578582828154811061221957612219613f10565b905f5260205f2001540361222f57809250612237565b6001016121fd565b50509392505050565b6122486132e2565b60065f8381526020019081526020015f20604051806101000160405290815f820154815260200160018201805461227e90613f40565b80601f01602080910402602001604051908101604052809291908181526020018280546122aa90613f40565b80156122f55780601f106122cc576101008083540402835291602001916122f5565b820191905f5260205f20905b8154815290600101906020018083116122d857829003601f168201915b505050505081526020016002820154815260200160038201805461231890613f40565b80601f016020809104026020016040519081016040528092919081815260200182805461234490613f40565b801561238f5780601f106123665761010080835404028352916020019161238f565b820191905f5260205f20905b81548152906001019060200180831161237257829003601f168201915b50505050508152602001600482015481526020016005820154815260200160068201548152602001600782016040518060400160405290815f8201805480602002602001604051908101604052809291908181526020015f905b82821015612491578382905f5260205f2001805461240690613f40565b80601f016020809104026020016040519081016040528092919081815260200182805461243290613f40565b801561247d5780601f106124545761010080835404028352916020019161247d565b820191905f5260205f20905b81548152906001019060200180831161246057829003601f168201915b5050505050815260200190600101906123e9565b505050508152602001600182018054806020026020016040519081016040528092919081815260200182805480156124e657602002820191905f5260205f20905b8154815260200190600101908083116124d2575b5050509190925250505090525092915050565b5f83815260046020526040812054849082036125285760405163938a92b760e01b815260040160405180910390fd5b5f818152600460205260409020600301546001600160a01b0316331461256e576125518161277d565b61256e5760405163dc64d0ad60e01b815260040160405180910390fd5b8333604051602001612581929190613ee4565b6040516020818303038152906040528051906020012085146125b6576040516327a5901560e11b815260040160405180910390fd5b5f8581526004602081905260409091200154156125e55760405162227f7760ea1b815260040160405180910390fd5b5f8581526004602052604081208181559061260360018301826133d5565b5f600283018190556003830180546001600160a01b031916905561262b90600484019061342a565b5050335f9081526005602052604090208054869082908690811061265157612651613f10565b905f5260205f20015414612678576040516337c7f25560e01b815260040160405180910390fd5b8054819061268890600190614092565b8154811061269857612698613f10565b905f5260205f2001548185815481106126b3576126b3613f10565b905f5260205f200181905550808054806126cf576126cf6140bc565b600190038181905f5260205f20015f90559055336001600160a01b0316856040516126fa9190613f35565b6040519081900381209088907feda2cc8e002ead8000b1e0c1debfc9a88bd7ee6e94b8dc0763db17849fcf0389905f90a450600195945050505050565b5f805b5f838152600660205260409020600701548110156127745761275c8382610c27565b15155f0361276c57505f92915050565b60010161273a565b50600192915050565b6003545f906001600160a01b03161561280e5760035460405163e124bdd960e01b8152600481018490523360248201526001600160a01b039091169063e124bdd990604401602060405180830381865afa1580156127dd573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061280191906140fa565b1561280e57506001919050565b505f919050565b5f8784604051602001612829929190613f24565b60408051601f1981840301815291815281516020928301205f818152600690935290822054909250900361287057604051632abde33960e01b815260040160405180910390fd5b5f81815260066020526040812060080180548990811061289257612892613f10565b905f5260205f2001549050808660016128ab9190614115565b60ff1611156128cd576040516359b452ef60e01b815260040160405180910390fd5b60408051610100810182525f848152600660205291822060070180548291908c9081106128fc576128fc613f10565b905f5260205f2001805461290f90613f40565b80601f016020809104026020016040519081016040528092919081815260200182805461293b90613f40565b80156129865780601f1061295d57610100808354040283529160200191612986565b820191905f5260205f20905b81548152906001019060200180831161296957829003601f168201915b505050505081526020018d81526020018a81526020018860ff1681526020018c81526020018981526020018581526020018b81525090506129c78186612bab565b6129d283888b612175565b156129f057604051636d680ca160e11b815260040160405180910390fd5b50505f898b86604051602001612a08939291906140a5565b60408051601f1981840301815291815281516020928301205f818152600b90935291205490915060ff16612a4457612a428a8c875f611836565b505b612a4f82878a612e22565b612a598289610c27565b15612a9b575f828152600860205260408120805460019290612a8090849061ffff1661412e565b92506101000a81548161ffff021916908361ffff1602179055505b6002546040516340c10f1960e01b8152336004820152670de0b6b3a764000060248201526001600160a01b03909116906340c10f19906044015f604051808303815f87803b158015612aeb575f5ffd5b505af1158015612afd573d5f5f3e3d5ffd5b5050604080518b815260ff8a1660208201528d93508e92507f2d22365fce1f74f329413a52494921c82b25abcc4b7d40efb5fc6ad9098be526910160405180910390a35050505050505050505050565b6060612b797f00000000000000000000000000000000000000000000000000000000000000005f612e8b565b905090565b6060612b797f00000000000000000000000000000000000000000000000000000000000000006001612e8b565b8160c00151421115612bf85760405162461bcd60e51b815260206004820152601160248201527014da59db985d1d5c9948195e1c1a5c9959607a1b60448201526064015b60405180910390fd5b60e08201515f90815260046020908152604080832060030154815160c08101909252608f8083526001600160a01b03909116939261415d9083013980519060200120845f01518051906020012085602001518660400151876060015188608001518960a001518a60c001518b60e00151604051602001612cbd9998979695949392919098895260208901979097526040880195909552606087019390935260ff91909116608086015260a085015260c084015260e08301526101008201526101200190565b6040516020818303038152906040528051906020012090505f612cdf82612f34565b90505f612cec8286612f60565b9050836001600160a01b0316816001600160a01b031614612d655760405162461bcd60e51b815260206004820152602d60248201527f496e76616c6964207369676e61747572653a204e6f74207369676e656420627960448201526c10313ab1b5b2ba1037bbb732b960991b6064820152608401612bef565b6001600160a01b0381165f908152600e6020908152604080832060e08a01518452825280832060a08a0151845290915290205460ff1615612ddd5760405162461bcd60e51b8152602060048201526012602482015271139bdb98d948185b1c9958591e481d5cd95960721b6044820152606401612bef565b6001600160a01b03165f908152600e6020908152604080832060e08901518452825280832060a090980151835296905294909420805460ff1916600117905550505050565b60208260ff1610612e46576040516359b452ef60e01b815260040160405180910390fd5b5f928352600760209081526040808520928552919052909120805463ffffffff600160ff9094169390931b83169281169290921763ffffffff19909216919091179055565b606060ff8314612ea557612e9e83612f88565b90506106e3565b818054612eb190613f40565b80601f0160208091040260200160405190810160405280929190818152602001828054612edd90613f40565b8015612f285780601f10612eff57610100808354040283529160200191612f28565b820191905f5260205f20905b815481529060010190602001808311612f0b57829003601f168201915b505050505090506106e3565b5f6106e3612f40612fc5565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f612f6e86866130ee565b925092509250612f7e8282613137565b5090949350505050565b60605f612f94836131f3565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614801561301d57507f000000000000000000000000000000000000000000000000000000000000000046145b1561304757507f000000000000000000000000000000000000000000000000000000000000000090565b612b79604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b5f5f5f8351604103613125576020840151604085015160608601515f1a6131178882858561321a565b955095509550505050613130565b505081515f91506002905b9250925092565b5f82600381111561314a5761314a614148565b03613153575050565b600182600381111561316757613167614148565b036131855760405163f645eedf60e01b815260040160405180910390fd5b600282600381111561319957613199614148565b036131ba5760405163fce698f760e01b815260048101829052602401612bef565b60038260038111156131ce576131ce614148565b036131ef576040516335e2f38360e21b815260048101829052602401612bef565b5050565b5f60ff8216601f8111156106e357604051632cd44ac360e21b815260040160405180910390fd5b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561325357505f915060039050826132d8565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156132a4573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166132cf57505f9250600191508290506132d8565b92505f91508190505b9450945094915050565b6040518061010001604052805f8152602001606081526020015f8152602001606081526020015f81526020015f81526020015f8152602001613337604051806040016040528060608152602001606081525090565b905290565b828054828255905f5260205f20908101928215613380579160200282015b8281111561338057825182906133709082613fc4565b509160200191906001019061335a565b5061338c929150613445565b5090565b828054828255905f5260205f209081019282156133c9579160200282015b828111156133c95782518255916020019190600101906133ae565b5061338c929150613461565b5080546133e190613f40565b5f825580601f106133f0575050565b601f0160209004905f5260205f209081019061340c9190613461565b50565b5080545f8255905f5260205f209081019061340c9190613445565b5080545f8255905f5260205f209081019061340c9190613461565b8082111561338c575f61345882826133d5565b50600101613445565b5b8082111561338c575f8155600101613462565b5f60208284031215613485575f5ffd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f8301126134af575f5ffd5b8135602083015f5f6001600160401b038411156134ce576134ce61348c565b50604051601f19601f85018116603f011681018181106001600160401b03821117156134fc576134fc61348c565b604052838152905080828401871015613513575f5ffd5b838360208301375f602085830101528094505050505092915050565b80356001600160a01b0381168114613545575f5ffd5b919050565b5f5f6040838503121561355b575f5ffd5b82356001600160401b03811115613570575f5ffd5b61357c858286016134a0565b92505061358b6020840161352f565b90509250929050565b5f5f83601f8401126135a4575f5ffd5b5081356001600160401b038111156135ba575f5ffd5b6020830191508360208260051b8501011115610e55575f5ffd5b5f5f5f5f5f5f5f5f5f60e08a8c0312156135ec575f5ffd5b89356001600160401b03811115613601575f5ffd5b61360d8c828d016134a0565b99505060208a0135975060408a01356001600160401b0381111561362f575f5ffd5b61363b8c828d016134a0565b97505060608a0135955060808a01356001600160401b0381111561365d575f5ffd5b6136698c828d01613594565b90965094505060a08a01356001600160401b03811115613687575f5ffd5b6136938c828d01613594565b9a9d999c50979a9699959894979660c00135949350505050565b5f5f604083850312156136be575f5ffd5b82356001600160401b038111156136d3575f5ffd5b6136df858286016134a0565b95602094909401359450505050565b5f5f602083850312156136ff575f5ffd5b82356001600160401b03811115613714575f5ffd5b61372085828601613594565b90969095509350505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b5f8151808452602084019350602083015f5b8281101561378a57815186526020958601959091019060010161376c565b5093949350505050565b805182525f602082015160a060208501526137b260a085018261372c565b90506040830151604085015260018060a01b036060840151166060850152608083015184820360808601526137e7828261375a565b95945050505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561384757603f19878603018452613832858351613794565b94506020938401939190910190600101613816565b50929695505050505050565b5f5f60408385031215613864575f5ffd5b50508035926020909101359150565b5f5f60408385031215613884575f5ffd5b8235915060208301356001600160401b038111156138a0575f5ffd5b6138ac858286016134a0565b9150509250929050565b602080825282518282018190525f918401906040840190835b818110156138ed5783518352602093840193909201916001016138cf565b509095945050505050565b604081525f61390a604083018561372c565b90508260208301529392505050565b803560ff81168114613545575f5ffd5b5f5f5f5f5f5f5f5f5f6101208a8c031215613942575f5ffd5b8935985060208a0135975060408a0135965060608a0135955060808a0135945061396e60a08b01613919565b935060c08a01356001600160401b03811115613988575f5ffd5b6139948c828d016134a0565b93505060e08a01356001600160401b038111156139af575f5ffd5b6139bb8c828d016134a0565b999c989b509699959894979396509194610100013592915050565b602081525f6139e8602083018461372c565b9392505050565b5f6040830182516040855281815180845260608701915060608160051b88010193506020830192505f5b81811015613a4a57605f19888603018352613a3585855161372c565b94506020938401939290920191600101613a19565b50505050602083015184820360208601526137e7828261375a565b60208152815160208201525f60208301516101006040840152613a8c61012084018261372c565b9050604084015160608401526060840151601f19848303016080850152613ab3828261372c565b915050608084015160a084015260a084015160c084015260c084015160e084015260e0840151601f19848303016101008501526137e782826139ef565b5f60208284031215613b00575f5ffd5b81356001600160401b03811115613b15575f5ffd5b613b21848285016134a0565b949350505050565b602081525f6139e86020830184613794565b5f60208284031215613b4b575f5ffd5b6139e88261352f565b5f5f5f5f5f60a08688031215613b68575f5ffd5b85359450602086013593506040860135925060608601356001600160401b03811115613b92575f5ffd5b613b9e888289016134a0565b95989497509295608001359392505050565b60ff60f81b8816815260e060208201525f613bce60e083018961372c565b8281036040840152613be0818961372c565b606084018890526001600160a01b038716608085015260a0840186905283810360c08501529050613c11818561375a565b9a9950505050505050505050565b801515811461340c575f5ffd5b5f5f5f5f60808587031215613c3f575f5ffd5b843593506020850135925060408501356001600160401b03811115613c62575f5ffd5b613c6e878288016134a0565b9250506060850135613c7f81613c1f565b939692955090935050565b5f5f5f5f5f60a08688031215613c9e575f5ffd5b8535945060208601356001600160401b03811115613cba575f5ffd5b613cc6888289016134a0565b945050604086013592506060860135915060808601356001600160401b03811115613cef575f5ffd5b613cfb888289016134a0565b9150509295509295909350565b5f5f5f60408486031215613d1a575f5ffd5b83356001600160401b03811115613d2f575f5ffd5b613d3b86828701613594565b90945092505060208401356001600160401b03811115613d59575f5ffd5b613d65868287016134a0565b9150509250925092565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561384757603f19878603018452613db185835161375a565b94506020938401939190910190600101613d95565b5f5f5f5f60808587031215613dd9575f5ffd5b843593506020850135925060408501356001600160401b03811115613dfc575f5ffd5b613e08878288016134a0565b949793965093946060013593505050565b5f5f5f60608486031215613e2b575f5ffd5b83359250613e3b60208501613919565b929592945050506040919091013590565b5f5f5f60608486031215613e5e575f5ffd5b833592506020840135915060408401356001600160401b03811115613d59575f5ffd5b5f5f5f60608486031215613e93575f5ffd5b8335925060208401356001600160401b03811115613eaf575f5ffd5b613ebb868287016134a0565b93969395505050506040919091013590565b5f81518060208401855e5f93019283525090919050565b5f613eef8285613ecd565b60609390931b6bffffffffffffffffffffffff191683525050601401919050565b634e487b7160e01b5f52603260045260245ffd5b8281525f613b216020830184613ecd565b5f6139e88284613ecd565b600181811c90821680613f5457607f821691505b602082108103613f7257634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115613fbf57805f5260205f20601f840160051c81016020851015613f9d5750805b601f840160051c820191505b81811015613fbc575f8155600101613fa9565b50505b505050565b81516001600160401b03811115613fdd57613fdd61348c565b613ff181613feb8454613f40565b84613f78565b6020601f821160018114614023575f831561400c5750848201515b5f19600385901b1c1916600184901b178455613fbc565b5f84815260208120601f198516915b828110156140525787850151825560209485019460019092019101614032565b508482101561406f57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b818103818111156106e3576106e361407e565b8381528260208201525f6137e76040830184613ecd565b634e487b7160e01b5f52603160045260245ffd5b808201808211156106e3576106e361407e565b80820281158282048414176106e3576106e361407e565b5f6020828403121561410a575f5ffd5b81516139e881613c1f565b60ff81811683821601908111156106e3576106e361407e565b61ffff81811683821601908111156106e3576106e361407e565b634e487b7160e01b5f52602160045260245ffdfe53746f7261676544617461286279746573206368756e6b4349442c6279746573333220626c6f636b4349442c75696e74323536206368756e6b496e6465782c75696e743820626c6f636b496e6465782c62797465733332206e6f646549642c75696e74323536206e6f6e63652c75696e7432353620646561646c696e652c62797465733332206275636b6574496429a2646970667358221220f5e086b53a7628ebd72ce8fcb467f86655c16a748186554d7202205660f677ed64736f6c634300081c0033",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// StorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageMetaData.Bin instead.
var StorageBin = StorageMetaData.Bin

// DeployStorage deploys a new Ethereum contract, binding an instance of Storage to it.
func DeployStorage(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddress common.Address) (common.Address, *types.Transaction, *Storage, error) {
	parsed, err := StorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageBin), backend, tokenAddress)
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

// GetBucketByName is a free data retrieval call binding the contract method 0x6554cda7.
//
// Solidity: function getBucketByName(string name) view returns((bytes32,string,uint256,address,bytes32[]) bucket)
func (_Storage *StorageCaller) GetBucketByName(opts *bind.CallOpts, name string) (IStorageBucket, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getBucketByName", name)

	if err != nil {
		return *new(IStorageBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(IStorageBucket)).(*IStorageBucket)

	return out0, err

}

// GetBucketByName is a free data retrieval call binding the contract method 0x6554cda7.
//
// Solidity: function getBucketByName(string name) view returns((bytes32,string,uint256,address,bytes32[]) bucket)
func (_Storage *StorageSession) GetBucketByName(name string) (IStorageBucket, error) {
	return _Storage.Contract.GetBucketByName(&_Storage.CallOpts, name)
}

// GetBucketByName is a free data retrieval call binding the contract method 0x6554cda7.
//
// Solidity: function getBucketByName(string name) view returns((bytes32,string,uint256,address,bytes32[]) bucket)
func (_Storage *StorageCallerSession) GetBucketByName(name string) (IStorageBucket, error) {
	return _Storage.Contract.GetBucketByName(&_Storage.CallOpts, name)
}

// GetBucketIndexByName is a free data retrieval call binding the contract method 0x287e677f.
//
// Solidity: function getBucketIndexByName(string name, address owner) view returns(uint256 index)
func (_Storage *StorageCaller) GetBucketIndexByName(opts *bind.CallOpts, name string, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getBucketIndexByName", name, owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBucketIndexByName is a free data retrieval call binding the contract method 0x287e677f.
//
// Solidity: function getBucketIndexByName(string name, address owner) view returns(uint256 index)
func (_Storage *StorageSession) GetBucketIndexByName(name string, owner common.Address) (*big.Int, error) {
	return _Storage.Contract.GetBucketIndexByName(&_Storage.CallOpts, name, owner)
}

// GetBucketIndexByName is a free data retrieval call binding the contract method 0x287e677f.
//
// Solidity: function getBucketIndexByName(string name, address owner) view returns(uint256 index)
func (_Storage *StorageCallerSession) GetBucketIndexByName(name string, owner common.Address) (*big.Int, error) {
	return _Storage.Contract.GetBucketIndexByName(&_Storage.CallOpts, name, owner)
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
// Solidity: function getFileByName(bytes32 bucketId, string name) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[])) file)
func (_Storage *StorageCaller) GetFileByName(opts *bind.CallOpts, bucketId [32]byte, name string) (IStorageFile, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getFileByName", bucketId, name)

	if err != nil {
		return *new(IStorageFile), err
	}

	out0 := *abi.ConvertType(out[0], new(IStorageFile)).(*IStorageFile)

	return out0, err

}

// GetFileByName is a free data retrieval call binding the contract method 0x5ecdfb53.
//
// Solidity: function getFileByName(bytes32 bucketId, string name) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[])) file)
func (_Storage *StorageSession) GetFileByName(bucketId [32]byte, name string) (IStorageFile, error) {
	return _Storage.Contract.GetFileByName(&_Storage.CallOpts, bucketId, name)
}

// GetFileByName is a free data retrieval call binding the contract method 0x5ecdfb53.
//
// Solidity: function getFileByName(bytes32 bucketId, string name) view returns((bytes32,bytes,bytes32,string,uint256,uint256,uint256,(bytes[],uint256[])) file)
func (_Storage *StorageCallerSession) GetFileByName(bucketId [32]byte, name string) (IStorageFile, error) {
	return _Storage.Contract.GetFileByName(&_Storage.CallOpts, bucketId, name)
}

// GetFileIndexById is a free data retrieval call binding the contract method 0x359b15a5.
//
// Solidity: function getFileIndexById(string name, bytes32 fileId) view returns(uint256 index)
func (_Storage *StorageCaller) GetFileIndexById(opts *bind.CallOpts, name string, fileId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getFileIndexById", name, fileId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFileIndexById is a free data retrieval call binding the contract method 0x359b15a5.
//
// Solidity: function getFileIndexById(string name, bytes32 fileId) view returns(uint256 index)
func (_Storage *StorageSession) GetFileIndexById(name string, fileId [32]byte) (*big.Int, error) {
	return _Storage.Contract.GetFileIndexById(&_Storage.CallOpts, name, fileId)
}

// GetFileIndexById is a free data retrieval call binding the contract method 0x359b15a5.
//
// Solidity: function getFileIndexById(string name, bytes32 fileId) view returns(uint256 index)
func (_Storage *StorageCallerSession) GetFileIndexById(name string, fileId [32]byte) (*big.Int, error) {
	return _Storage.Contract.GetFileIndexById(&_Storage.CallOpts, name, fileId)
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

// GetOwnerBuckets is a free data retrieval call binding the contract method 0x6a5d8c26.
//
// Solidity: function getOwnerBuckets(address owner) view returns(bytes32[] buckets)
func (_Storage *StorageCaller) GetOwnerBuckets(opts *bind.CallOpts, owner common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getOwnerBuckets", owner)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetOwnerBuckets is a free data retrieval call binding the contract method 0x6a5d8c26.
//
// Solidity: function getOwnerBuckets(address owner) view returns(bytes32[] buckets)
func (_Storage *StorageSession) GetOwnerBuckets(owner common.Address) ([][32]byte, error) {
	return _Storage.Contract.GetOwnerBuckets(&_Storage.CallOpts, owner)
}

// GetOwnerBuckets is a free data retrieval call binding the contract method 0x6a5d8c26.
//
// Solidity: function getOwnerBuckets(address owner) view returns(bytes32[] buckets)
func (_Storage *StorageCallerSession) GetOwnerBuckets(owner common.Address) ([][32]byte, error) {
	return _Storage.Contract.GetOwnerBuckets(&_Storage.CallOpts, owner)
}

// GetPeerBlockIndexById is a free data retrieval call binding the contract method 0xf855169a.
//
// Solidity: function getPeerBlockIndexById(bytes32 peerId, bytes32 cid, string name) view returns(uint256 index)
func (_Storage *StorageCaller) GetPeerBlockIndexById(opts *bind.CallOpts, peerId [32]byte, cid [32]byte, name string) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getPeerBlockIndexById", peerId, cid, name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPeerBlockIndexById is a free data retrieval call binding the contract method 0xf855169a.
//
// Solidity: function getPeerBlockIndexById(bytes32 peerId, bytes32 cid, string name) view returns(uint256 index)
func (_Storage *StorageSession) GetPeerBlockIndexById(peerId [32]byte, cid [32]byte, name string) (*big.Int, error) {
	return _Storage.Contract.GetPeerBlockIndexById(&_Storage.CallOpts, peerId, cid, name)
}

// GetPeerBlockIndexById is a free data retrieval call binding the contract method 0xf855169a.
//
// Solidity: function getPeerBlockIndexById(bytes32 peerId, bytes32 cid, string name) view returns(uint256 index)
func (_Storage *StorageCallerSession) GetPeerBlockIndexById(peerId [32]byte, cid [32]byte, name string) (*big.Int, error) {
	return _Storage.Contract.GetPeerBlockIndexById(&_Storage.CallOpts, peerId, cid, name)
}

// GetPeersArrayByPeerBlockCid is a free data retrieval call binding the contract method 0xd606205d.
//
// Solidity: function getPeersArrayByPeerBlockCid(bytes32[] cids, string name) view returns(bytes32[][] peers)
func (_Storage *StorageCaller) GetPeersArrayByPeerBlockCid(opts *bind.CallOpts, cids [][32]byte, name string) ([][][32]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getPeersArrayByPeerBlockCid", cids, name)

	if err != nil {
		return *new([][][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][][32]byte)).(*[][][32]byte)

	return out0, err

}

// GetPeersArrayByPeerBlockCid is a free data retrieval call binding the contract method 0xd606205d.
//
// Solidity: function getPeersArrayByPeerBlockCid(bytes32[] cids, string name) view returns(bytes32[][] peers)
func (_Storage *StorageSession) GetPeersArrayByPeerBlockCid(cids [][32]byte, name string) ([][][32]byte, error) {
	return _Storage.Contract.GetPeersArrayByPeerBlockCid(&_Storage.CallOpts, cids, name)
}

// GetPeersArrayByPeerBlockCid is a free data retrieval call binding the contract method 0xd606205d.
//
// Solidity: function getPeersArrayByPeerBlockCid(bytes32[] cids, string name) view returns(bytes32[][] peers)
func (_Storage *StorageCallerSession) GetPeersArrayByPeerBlockCid(cids [][32]byte, name string) ([][][32]byte, error) {
	return _Storage.Contract.GetPeersArrayByPeerBlockCid(&_Storage.CallOpts, cids, name)
}

// GetPeersByPeerBlockCid is a free data retrieval call binding the contract method 0x48b49875.
//
// Solidity: function getPeersByPeerBlockCid(bytes32 cid, string name) view returns(bytes32[] peers)
func (_Storage *StorageCaller) GetPeersByPeerBlockCid(opts *bind.CallOpts, cid [32]byte, name string) ([][32]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getPeersByPeerBlockCid", cid, name)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetPeersByPeerBlockCid is a free data retrieval call binding the contract method 0x48b49875.
//
// Solidity: function getPeersByPeerBlockCid(bytes32 cid, string name) view returns(bytes32[] peers)
func (_Storage *StorageSession) GetPeersByPeerBlockCid(cid [32]byte, name string) ([][32]byte, error) {
	return _Storage.Contract.GetPeersByPeerBlockCid(&_Storage.CallOpts, cid, name)
}

// GetPeersByPeerBlockCid is a free data retrieval call binding the contract method 0x48b49875.
//
// Solidity: function getPeersByPeerBlockCid(bytes32 cid, string name) view returns(bytes32[] peers)
func (_Storage *StorageCallerSession) GetPeersByPeerBlockCid(cid [32]byte, name string) ([][32]byte, error) {
	return _Storage.Contract.GetPeersByPeerBlockCid(&_Storage.CallOpts, cid, name)
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
// Solidity: function addFileChunk(bytes cid, bytes32 bucketId, string name, uint256 encodedChunkSize, bytes32[] cids, uint256[] chunkBlocksSizes, uint256 chunkIndex) returns(bytes32)
func (_Storage *StorageTransactor) AddFileChunk(opts *bind.TransactOpts, cid []byte, bucketId [32]byte, name string, encodedChunkSize *big.Int, cids [][32]byte, chunkBlocksSizes []*big.Int, chunkIndex *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addFileChunk", cid, bucketId, name, encodedChunkSize, cids, chunkBlocksSizes, chunkIndex)
}

// AddFileChunk is a paid mutator transaction binding the contract method 0x30b91d07.
//
// Solidity: function addFileChunk(bytes cid, bytes32 bucketId, string name, uint256 encodedChunkSize, bytes32[] cids, uint256[] chunkBlocksSizes, uint256 chunkIndex) returns(bytes32)
func (_Storage *StorageSession) AddFileChunk(cid []byte, bucketId [32]byte, name string, encodedChunkSize *big.Int, cids [][32]byte, chunkBlocksSizes []*big.Int, chunkIndex *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddFileChunk(&_Storage.TransactOpts, cid, bucketId, name, encodedChunkSize, cids, chunkBlocksSizes, chunkIndex)
}

// AddFileChunk is a paid mutator transaction binding the contract method 0x30b91d07.
//
// Solidity: function addFileChunk(bytes cid, bytes32 bucketId, string name, uint256 encodedChunkSize, bytes32[] cids, uint256[] chunkBlocksSizes, uint256 chunkIndex) returns(bytes32)
func (_Storage *StorageTransactorSession) AddFileChunk(cid []byte, bucketId [32]byte, name string, encodedChunkSize *big.Int, cids [][32]byte, chunkBlocksSizes []*big.Int, chunkIndex *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.AddFileChunk(&_Storage.TransactOpts, cid, bucketId, name, encodedChunkSize, cids, chunkBlocksSizes, chunkIndex)
}

// AddPeerBlock is a paid mutator transaction binding the contract method 0x87c1ac07.
//
// Solidity: function addPeerBlock(bytes32 peerId, bytes32 cid, string name, bool isReplica) returns(bytes32 id)
func (_Storage *StorageTransactor) AddPeerBlock(opts *bind.TransactOpts, peerId [32]byte, cid [32]byte, name string, isReplica bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addPeerBlock", peerId, cid, name, isReplica)
}

// AddPeerBlock is a paid mutator transaction binding the contract method 0x87c1ac07.
//
// Solidity: function addPeerBlock(bytes32 peerId, bytes32 cid, string name, bool isReplica) returns(bytes32 id)
func (_Storage *StorageSession) AddPeerBlock(peerId [32]byte, cid [32]byte, name string, isReplica bool) (*types.Transaction, error) {
	return _Storage.Contract.AddPeerBlock(&_Storage.TransactOpts, peerId, cid, name, isReplica)
}

// AddPeerBlock is a paid mutator transaction binding the contract method 0x87c1ac07.
//
// Solidity: function addPeerBlock(bytes32 peerId, bytes32 cid, string name, bool isReplica) returns(bytes32 id)
func (_Storage *StorageTransactorSession) AddPeerBlock(peerId [32]byte, cid [32]byte, name string, isReplica bool) (*types.Transaction, error) {
	return _Storage.Contract.AddPeerBlock(&_Storage.TransactOpts, peerId, cid, name, isReplica)
}

// CommitFile is a paid mutator transaction binding the contract method 0x9a2e82b3.
//
// Solidity: function commitFile(bytes32 bucketId, string name, uint256 encodedFileSize, uint256 actualSize, bytes fileCID) returns(bytes32)
func (_Storage *StorageTransactor) CommitFile(opts *bind.TransactOpts, bucketId [32]byte, name string, encodedFileSize *big.Int, actualSize *big.Int, fileCID []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "commitFile", bucketId, name, encodedFileSize, actualSize, fileCID)
}

// CommitFile is a paid mutator transaction binding the contract method 0x9a2e82b3.
//
// Solidity: function commitFile(bytes32 bucketId, string name, uint256 encodedFileSize, uint256 actualSize, bytes fileCID) returns(bytes32)
func (_Storage *StorageSession) CommitFile(bucketId [32]byte, name string, encodedFileSize *big.Int, actualSize *big.Int, fileCID []byte) (*types.Transaction, error) {
	return _Storage.Contract.CommitFile(&_Storage.TransactOpts, bucketId, name, encodedFileSize, actualSize, fileCID)
}

// CommitFile is a paid mutator transaction binding the contract method 0x9a2e82b3.
//
// Solidity: function commitFile(bytes32 bucketId, string name, uint256 encodedFileSize, uint256 actualSize, bytes fileCID) returns(bytes32)
func (_Storage *StorageTransactorSession) CommitFile(bucketId [32]byte, name string, encodedFileSize *big.Int, actualSize *big.Int, fileCID []byte) (*types.Transaction, error) {
	return _Storage.Contract.CommitFile(&_Storage.TransactOpts, bucketId, name, encodedFileSize, actualSize, fileCID)
}

// CreateBucket is a paid mutator transaction binding the contract method 0xe3f787e8.
//
// Solidity: function createBucket(string name) returns(bytes32 id)
func (_Storage *StorageTransactor) CreateBucket(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "createBucket", name)
}

// CreateBucket is a paid mutator transaction binding the contract method 0xe3f787e8.
//
// Solidity: function createBucket(string name) returns(bytes32 id)
func (_Storage *StorageSession) CreateBucket(name string) (*types.Transaction, error) {
	return _Storage.Contract.CreateBucket(&_Storage.TransactOpts, name)
}

// CreateBucket is a paid mutator transaction binding the contract method 0xe3f787e8.
//
// Solidity: function createBucket(string name) returns(bytes32 id)
func (_Storage *StorageTransactorSession) CreateBucket(name string) (*types.Transaction, error) {
	return _Storage.Contract.CreateBucket(&_Storage.TransactOpts, name)
}

// CreateFile is a paid mutator transaction binding the contract method 0x6af0f801.
//
// Solidity: function createFile(bytes32 bucketId, string name) returns(bytes32)
func (_Storage *StorageTransactor) CreateFile(opts *bind.TransactOpts, bucketId [32]byte, name string) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "createFile", bucketId, name)
}

// CreateFile is a paid mutator transaction binding the contract method 0x6af0f801.
//
// Solidity: function createFile(bytes32 bucketId, string name) returns(bytes32)
func (_Storage *StorageSession) CreateFile(bucketId [32]byte, name string) (*types.Transaction, error) {
	return _Storage.Contract.CreateFile(&_Storage.TransactOpts, bucketId, name)
}

// CreateFile is a paid mutator transaction binding the contract method 0x6af0f801.
//
// Solidity: function createFile(bytes32 bucketId, string name) returns(bytes32)
func (_Storage *StorageTransactorSession) CreateFile(bucketId [32]byte, name string) (*types.Transaction, error) {
	return _Storage.Contract.CreateFile(&_Storage.TransactOpts, bucketId, name)
}

// DeleteBucket is a paid mutator transaction binding the contract method 0xfd1d3c0c.
//
// Solidity: function deleteBucket(bytes32 id, string name, uint256 index) returns(bool)
func (_Storage *StorageTransactor) DeleteBucket(opts *bind.TransactOpts, id [32]byte, name string, index *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteBucket", id, name, index)
}

// DeleteBucket is a paid mutator transaction binding the contract method 0xfd1d3c0c.
//
// Solidity: function deleteBucket(bytes32 id, string name, uint256 index) returns(bool)
func (_Storage *StorageSession) DeleteBucket(id [32]byte, name string, index *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DeleteBucket(&_Storage.TransactOpts, id, name, index)
}

// DeleteBucket is a paid mutator transaction binding the contract method 0xfd1d3c0c.
//
// Solidity: function deleteBucket(bytes32 id, string name, uint256 index) returns(bool)
func (_Storage *StorageTransactorSession) DeleteBucket(id [32]byte, name string, index *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DeleteBucket(&_Storage.TransactOpts, id, name, index)
}

// DeleteFile is a paid mutator transaction binding the contract method 0xd6d3110b.
//
// Solidity: function deleteFile(bytes32 fileID, bytes32 bucketId, string name, uint256 index) returns(bool)
func (_Storage *StorageTransactor) DeleteFile(opts *bind.TransactOpts, fileID [32]byte, bucketId [32]byte, name string, index *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deleteFile", fileID, bucketId, name, index)
}

// DeleteFile is a paid mutator transaction binding the contract method 0xd6d3110b.
//
// Solidity: function deleteFile(bytes32 fileID, bytes32 bucketId, string name, uint256 index) returns(bool)
func (_Storage *StorageSession) DeleteFile(fileID [32]byte, bucketId [32]byte, name string, index *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DeleteFile(&_Storage.TransactOpts, fileID, bucketId, name, index)
}

// DeleteFile is a paid mutator transaction binding the contract method 0xd6d3110b.
//
// Solidity: function deleteFile(bytes32 fileID, bytes32 bucketId, string name, uint256 index) returns(bool)
func (_Storage *StorageTransactorSession) DeleteFile(fileID [32]byte, bucketId [32]byte, name string, index *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DeleteFile(&_Storage.TransactOpts, fileID, bucketId, name, index)
}

// DeletePeerBlock is a paid mutator transaction binding the contract method 0x6a6e658b.
//
// Solidity: function deletePeerBlock(bytes32 id, bytes32 peerId, bytes32 cid, string name, uint256 index) returns(bool)
func (_Storage *StorageTransactor) DeletePeerBlock(opts *bind.TransactOpts, id [32]byte, peerId [32]byte, cid [32]byte, name string, index *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deletePeerBlock", id, peerId, cid, name, index)
}

// DeletePeerBlock is a paid mutator transaction binding the contract method 0x6a6e658b.
//
// Solidity: function deletePeerBlock(bytes32 id, bytes32 peerId, bytes32 cid, string name, uint256 index) returns(bool)
func (_Storage *StorageSession) DeletePeerBlock(id [32]byte, peerId [32]byte, cid [32]byte, name string, index *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DeletePeerBlock(&_Storage.TransactOpts, id, peerId, cid, name, index)
}

// DeletePeerBlock is a paid mutator transaction binding the contract method 0x6a6e658b.
//
// Solidity: function deletePeerBlock(bytes32 id, bytes32 peerId, bytes32 cid, string name, uint256 index) returns(bool)
func (_Storage *StorageTransactorSession) DeletePeerBlock(id [32]byte, peerId [32]byte, cid [32]byte, name string, index *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DeletePeerBlock(&_Storage.TransactOpts, id, peerId, cid, name, index)
}

// FillChunkBlock is a paid mutator transaction binding the contract method 0x4f55ba82.
//
// Solidity: function fillChunkBlock(bytes32 blockCID, bytes32 nodeId, bytes32 bucketId, uint256 chunkIndex, uint256 nonce, uint8 blockIndex, string name, bytes signature, uint256 deadline) returns()
func (_Storage *StorageTransactor) FillChunkBlock(opts *bind.TransactOpts, blockCID [32]byte, nodeId [32]byte, bucketId [32]byte, chunkIndex *big.Int, nonce *big.Int, blockIndex uint8, name string, signature []byte, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "fillChunkBlock", blockCID, nodeId, bucketId, chunkIndex, nonce, blockIndex, name, signature, deadline)
}

// FillChunkBlock is a paid mutator transaction binding the contract method 0x4f55ba82.
//
// Solidity: function fillChunkBlock(bytes32 blockCID, bytes32 nodeId, bytes32 bucketId, uint256 chunkIndex, uint256 nonce, uint8 blockIndex, string name, bytes signature, uint256 deadline) returns()
func (_Storage *StorageSession) FillChunkBlock(blockCID [32]byte, nodeId [32]byte, bucketId [32]byte, chunkIndex *big.Int, nonce *big.Int, blockIndex uint8, name string, signature []byte, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.FillChunkBlock(&_Storage.TransactOpts, blockCID, nodeId, bucketId, chunkIndex, nonce, blockIndex, name, signature, deadline)
}

// FillChunkBlock is a paid mutator transaction binding the contract method 0x4f55ba82.
//
// Solidity: function fillChunkBlock(bytes32 blockCID, bytes32 nodeId, bytes32 bucketId, uint256 chunkIndex, uint256 nonce, uint8 blockIndex, string name, bytes signature, uint256 deadline) returns()
func (_Storage *StorageTransactorSession) FillChunkBlock(blockCID [32]byte, nodeId [32]byte, bucketId [32]byte, chunkIndex *big.Int, nonce *big.Int, blockIndex uint8, name string, signature []byte, deadline *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.FillChunkBlock(&_Storage.TransactOpts, blockCID, nodeId, bucketId, chunkIndex, nonce, blockIndex, name, signature, deadline)
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

// StorageAddFileIterator is returned from FilterAddFile and is used to iterate over the raw logs and unpacked data for AddFile events raised by the Storage contract.
type StorageAddFileIterator struct {
	Event *StorageAddFile // Event containing the contract specifics and raw log

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
func (it *StorageAddFileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageAddFile)
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
		it.Event = new(StorageAddFile)
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
func (it *StorageAddFileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageAddFileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageAddFile represents a AddFile event raised by the Storage contract.
type StorageAddFile struct {
	Id       [32]byte
	BucketId [32]byte
	Name     common.Hash
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddFile is a free log retrieval operation binding the contract event 0x01d10894cb2a39778aae51e234b669f70a74328f07e58e67a2caca4c5a3c86ff.
//
// Solidity: event AddFile(bytes32 indexed id, bytes32 indexed bucketId, string indexed name, address owner)
func (_Storage *StorageFilterer) FilterAddFile(opts *bind.FilterOpts, id [][32]byte, bucketId [][32]byte, name []string) (*StorageAddFileIterator, error) {

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

	logs, sub, err := _Storage.contract.FilterLogs(opts, "AddFile", idRule, bucketIdRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &StorageAddFileIterator{contract: _Storage.contract, event: "AddFile", logs: logs, sub: sub}, nil
}

// WatchAddFile is a free log subscription operation binding the contract event 0x01d10894cb2a39778aae51e234b669f70a74328f07e58e67a2caca4c5a3c86ff.
//
// Solidity: event AddFile(bytes32 indexed id, bytes32 indexed bucketId, string indexed name, address owner)
func (_Storage *StorageFilterer) WatchAddFile(opts *bind.WatchOpts, sink chan<- *StorageAddFile, id [][32]byte, bucketId [][32]byte, name []string) (event.Subscription, error) {

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

	logs, sub, err := _Storage.contract.WatchLogs(opts, "AddFile", idRule, bucketIdRule, nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageAddFile)
				if err := _Storage.contract.UnpackLog(event, "AddFile", log); err != nil {
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

// ParseAddFile is a log parse operation binding the contract event 0x01d10894cb2a39778aae51e234b669f70a74328f07e58e67a2caca4c5a3c86ff.
//
// Solidity: event AddFile(bytes32 indexed id, bytes32 indexed bucketId, string indexed name, address owner)
func (_Storage *StorageFilterer) ParseAddFile(log types.Log) (*StorageAddFile, error) {
	event := new(StorageAddFile)
	if err := _Storage.contract.UnpackLog(event, "AddFile", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// StorageChunkBlockFilledIterator is returned from FilterChunkBlockFilled and is used to iterate over the raw logs and unpacked data for ChunkBlockFilled events raised by the Storage contract.
type StorageChunkBlockFilledIterator struct {
	Event *StorageChunkBlockFilled // Event containing the contract specifics and raw log

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
func (it *StorageChunkBlockFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageChunkBlockFilled)
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
		it.Event = new(StorageChunkBlockFilled)
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
func (it *StorageChunkBlockFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageChunkBlockFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageChunkBlockFilled represents a ChunkBlockFilled event raised by the Storage contract.
type StorageChunkBlockFilled struct {
	Cid        [32]byte
	ChunkIndex *big.Int
	BlockIndex uint8
	NodeId     [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChunkBlockFilled is a free log retrieval operation binding the contract event 0x2d22365fce1f74f329413a52494921c82b25abcc4b7d40efb5fc6ad9098be526.
//
// Solidity: event ChunkBlockFilled(bytes32 indexed cid, uint256 chunkIndex, uint8 blockIndex, bytes32 indexed nodeId)
func (_Storage *StorageFilterer) FilterChunkBlockFilled(opts *bind.FilterOpts, cid [][32]byte, nodeId [][32]byte) (*StorageChunkBlockFilledIterator, error) {

	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "ChunkBlockFilled", cidRule, nodeIdRule)
	if err != nil {
		return nil, err
	}
	return &StorageChunkBlockFilledIterator{contract: _Storage.contract, event: "ChunkBlockFilled", logs: logs, sub: sub}, nil
}

// WatchChunkBlockFilled is a free log subscription operation binding the contract event 0x2d22365fce1f74f329413a52494921c82b25abcc4b7d40efb5fc6ad9098be526.
//
// Solidity: event ChunkBlockFilled(bytes32 indexed cid, uint256 chunkIndex, uint8 blockIndex, bytes32 indexed nodeId)
func (_Storage *StorageFilterer) WatchChunkBlockFilled(opts *bind.WatchOpts, sink chan<- *StorageChunkBlockFilled, cid [][32]byte, nodeId [][32]byte) (event.Subscription, error) {

	var cidRule []interface{}
	for _, cidItem := range cid {
		cidRule = append(cidRule, cidItem)
	}

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "ChunkBlockFilled", cidRule, nodeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageChunkBlockFilled)
				if err := _Storage.contract.UnpackLog(event, "ChunkBlockFilled", log); err != nil {
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

// ParseChunkBlockFilled is a log parse operation binding the contract event 0x2d22365fce1f74f329413a52494921c82b25abcc4b7d40efb5fc6ad9098be526.
//
// Solidity: event ChunkBlockFilled(bytes32 indexed cid, uint256 chunkIndex, uint8 blockIndex, bytes32 indexed nodeId)
func (_Storage *StorageFilterer) ParseChunkBlockFilled(log types.Log) (*StorageChunkBlockFilled, error) {
	event := new(StorageChunkBlockFilled)
	if err := _Storage.contract.UnpackLog(event, "ChunkBlockFilled", log); err != nil {
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

// StorageFileUploadedIterator is returned from FilterFileUploaded and is used to iterate over the raw logs and unpacked data for FileUploaded events raised by the Storage contract.
type StorageFileUploadedIterator struct {
	Event *StorageFileUploaded // Event containing the contract specifics and raw log

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
func (it *StorageFileUploadedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageFileUploaded)
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
		it.Event = new(StorageFileUploaded)
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
func (it *StorageFileUploadedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageFileUploadedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageFileUploaded represents a FileUploaded event raised by the Storage contract.
type StorageFileUploaded struct {
	Id       [32]byte
	BucketId [32]byte
	Name     common.Hash
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFileUploaded is a free log retrieval operation binding the contract event 0xb2e54a2138908ebdd7da28708ec0bc3f1498b96fb7b0db337edef5ceeb41b16b.
//
// Solidity: event FileUploaded(bytes32 indexed id, bytes32 indexed bucketId, string indexed name, address owner)
func (_Storage *StorageFilterer) FilterFileUploaded(opts *bind.FilterOpts, id [][32]byte, bucketId [][32]byte, name []string) (*StorageFileUploadedIterator, error) {

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

	logs, sub, err := _Storage.contract.FilterLogs(opts, "FileUploaded", idRule, bucketIdRule, nameRule)
	if err != nil {
		return nil, err
	}
	return &StorageFileUploadedIterator{contract: _Storage.contract, event: "FileUploaded", logs: logs, sub: sub}, nil
}

// WatchFileUploaded is a free log subscription operation binding the contract event 0xb2e54a2138908ebdd7da28708ec0bc3f1498b96fb7b0db337edef5ceeb41b16b.
//
// Solidity: event FileUploaded(bytes32 indexed id, bytes32 indexed bucketId, string indexed name, address owner)
func (_Storage *StorageFilterer) WatchFileUploaded(opts *bind.WatchOpts, sink chan<- *StorageFileUploaded, id [][32]byte, bucketId [][32]byte, name []string) (event.Subscription, error) {

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

	logs, sub, err := _Storage.contract.WatchLogs(opts, "FileUploaded", idRule, bucketIdRule, nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageFileUploaded)
				if err := _Storage.contract.UnpackLog(event, "FileUploaded", log); err != nil {
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

// ParseFileUploaded is a log parse operation binding the contract event 0xb2e54a2138908ebdd7da28708ec0bc3f1498b96fb7b0db337edef5ceeb41b16b.
//
// Solidity: event FileUploaded(bytes32 indexed id, bytes32 indexed bucketId, string indexed name, address owner)
func (_Storage *StorageFilterer) ParseFileUploaded(log types.Log) (*StorageFileUploaded, error) {
	event := new(StorageFileUploaded)
	if err := _Storage.contract.UnpackLog(event, "FileUploaded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
