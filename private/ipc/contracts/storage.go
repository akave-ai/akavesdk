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

// IStoragePeerBlock is an auto generated low-level Go binding around an user-defined struct.
type IStoragePeerBlock struct {
	PeerId    []byte
	IsReplica bool
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BlockAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockAlreadyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockNonexists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketInvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketNonempty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketNonexists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"}],\"name\":\"ChunkCIDMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileChunkDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileFullyUploaded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNameDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNonempty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNotFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cidsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sizesLength\",\"type\":\"uint256\"}],\"name\":\"InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlocksAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEncodedSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFileBlocksCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFileCID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLastBlockSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LastChunkDuplicate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AddFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"AddFileBlocks\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"peerId\",\"type\":\"bytes\"}],\"name\":\"AddPeerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeId\",\"type\":\"bytes\"}],\"name\":\"ChunkBlockFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CreateBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CreateFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DeleteBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DeleteFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"peerId\",\"type\":\"bytes\"}],\"name\":\"DeletePeerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"FileUploaded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_BLOCKS_PER_FILE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BLOCK_SIZE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accessManager\",\"outputs\":[{\"internalType\":\"contractIAccessManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"cid\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedChunkSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"cids\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkBlocksSizes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"addFileChunk\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"peerId\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isReplica\",\"type\":\"bool\"}],\"name\":\"addPeerBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedFileSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"}],\"name\":\"commitFile\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"createBucket\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"createFile\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deleteBucket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deleteFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"peerId\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deletePeerBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fileFillCounter\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fileRewardClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockCID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"nodeId\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"fillChunkBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fulfilledBlocks\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getBucketByName\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket\",\"name\":\"bucket\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getBucketIndexByName\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"}],\"name\":\"getBucketsByIds\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getChunkByIndex\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getFileById\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"chunkCIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkSize\",\"type\":\"uint256[]\"}],\"internalType\":\"structIStorage.Chunk\",\"name\":\"chunks\",\"type\":\"tuple\"}],\"internalType\":\"structIStorage.File\",\"name\":\"file\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getFileByName\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualSize\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"chunkCIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkSize\",\"type\":\"uint256[]\"}],\"internalType\":\"structIStorage.Chunk\",\"name\":\"chunks\",\"type\":\"tuple\"}],\"internalType\":\"structIStorage.File\",\"name\":\"file\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"getFileIndexById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getFileOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getOwnerBuckets\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"buckets\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"peerId\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"}],\"name\":\"getPeerBlockByCid\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"peerId\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isReplica\",\"type\":\"bool\"}],\"internalType\":\"structIStorage.PeerBlock\",\"name\":\"peerBlock\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getPeerBlockById\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"peerId\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isReplica\",\"type\":\"bool\"}],\"internalType\":\"structIStorage.PeerBlock\",\"name\":\"peerBlock\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"peerId\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"}],\"name\":\"getPeerBlockIndexById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"}],\"name\":\"getPeersByPeerBlockCid\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"peers\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"isBlockFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"isChunkFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"isFileFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"isFileFilledV2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accessManagerAddress\",\"type\":\"address\"}],\"name\":\"setAccessManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIAkaveToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x610160604052348015610010575f5ffd5b506040516147c13803806147c183398101604081905261002f916101b1565b604080518082018252600781526653746f7261676560c81b602080830191909152825180840190935260018352603160f81b9083015290610070825f610139565b6101205261007f816001610139565b61014052815160208084019190912060e052815190820120610100524660a05261010b60e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c052600280546001600160a01b0319166001600160a01b0392909216919091179055610388565b5f6020835110156101545761014d8361016b565b9050610165565b8161015f8482610276565b5060ff90505b92915050565b5f5f829050601f8151111561019e578260405163305a27a960e01b81526004016101959190610330565b60405180910390fd5b80516101a982610365565b179392505050565b5f602082840312156101c1575f5ffd5b81516001600160a01b03811681146101d7575f5ffd5b9392505050565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061020657607f821691505b60208210810361022457634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561027157805f5260205f20601f840160051c8101602085101561024f5750805b601f840160051c820191505b8181101561026e575f815560010161025b565b50505b505050565b81516001600160401b0381111561028f5761028f6101de565b6102a38161029d84546101f2565b8461022a565b6020601f8211600181146102d5575f83156102be5750848201515b5f19600385901b1c1916600184901b17845561026e565b5f84815260208120601f198516915b8281101561030457878501518255602094850194600190920191016102e4565b508482101561032157868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610224575f1960209190910360031b1b16919050565b60805160a05160c05160e0516101005161012051610140516143e86103d95f395f612e3b01525f612e0a01525f61303701525f61300f01525f612f6a01525f612f9401525f612fbe01526143e85ff3fe608060405234801561000f575f5ffd5b5060043610610229575f3560e01c80636ce023631161012a578063d6d3110b116100b4578063faec054211610079578063faec05421461060b578063fc0c546a1461061e578063fd1d3c0c14610631578063fd21c28414610644578063fdcb606814610657575f5ffd5b8063d6d3110b14610589578063e3f787e81461059c578063e4ba8a58146105af578063f8a3e41a146105c2578063f8fefaaa146105f8575f5ffd5b80639a094ca2116100fa5780639a094ca2146105245780639a2e82b3146105375780639ccd46461461054a578063b80777ea14610553578063c958080414610559575f5ffd5b80636ce02363146104bf57806383f77cdb146104e157806384b0196e146104f657806395696eb214610511575f5ffd5b80634d7dc614116101b65780635ecdfb531161017b5780635ecdfb53146104175780636554cda71461043757806368e6408f146104575780636a5d8c261461048c5780636af0f801146104ac575f5ffd5b80634d7dc614146103715780634ed0e6321461039257806354fd4d50146103a5578063564b81ef146103cc5780635a4e9564146103d2575f5ffd5b8063359b15a5116101fc578063359b15a5146102eb57806335bdb711146102fe5780633f3839801461031e5780634656b932146103315780634d15ebbd14610351575f5ffd5b8063018c1e9c1461022d5780631b475ef414610264578063287e677f146102b757806330b91d07146102d8575b5f5ffd5b61024f61023b36600461344b565b600a6020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b61029f61027236600461344b565b5f90815260066020908152604080832060020154835260049091529020600301546001600160a01b031690565b6040516001600160a01b03909116815260200161025b565b6102ca6102c5366004613520565b61066a565b60405190815260200161025b565b6102ca6102e63660046135aa565b6106f6565b6102ca6102f9366004613683565b61099f565b61031161030c3660046136c4565b610a1b565b60405161025b91906137bc565b61024f61032c36600461381f565b610c34565b61034461033f36600461344b565b610cf1565b60405161025b9190613899565b61036461035f3660046138ef565b610dd6565b60405161025b9190613936565b61038461037f36600461381f565b610ed7565b60405161025b929190613969565b6103646103a036600461344b565b610fd1565b60408051808201825260058152640312e302e360dc1b6020820152905161025b919061398a565b466102ca565b6104026103e036600461381f565b600760209081525f928352604080842090915290825290205463ffffffff1681565b60405163ffffffff909116815260200161025b565b61042a61042536600461399c565b6110a2565b60405161025b9190613a4f565b61044a610445366004613ada565b611381565b60405161025b9190613b13565b61024f61046536600461344b565b5f9081526008602090815260408083205460069092529091206007015461ffff9091161490565b61049f61049a366004613b25565b611528565b60405161025b9190613b3e565b6102ca6104ba36600461399c565b611591565b6104c9620f424081565b6040516001600160401b03909116815260200161025b565b6104f46104ef366004613b90565b611823565b005b6104fe611b1c565b60405161025b9796959493929190613c55565b6102ca61051f3660046138ef565b611b5e565b6102ca610532366004613cd1565b611bdb565b6102ca610545366004613d2b565b611cef565b6104c961040081565b426102ca565b6104f4610567366004613b25565b600380546001600160a01b0319166001600160a01b0392909216919091179055565b61024f610597366004613da9565b612049565b6102ca6105aa366004613ada565b6122bb565b61024f6105bd366004613dfc565b612421565b6105e56105d036600461344b565b60086020525f908152604090205461ffff1681565b60405161ffff909116815260200161025b565b61024f610606366004613e2f565b612478565b61042a61061936600461344b565b6125f6565b60025461029f906001600160a01b031681565b61024f61063f366004613e87565b6128af565b61024f61065236600461344b565b612aed565b60035461029f906001600160a01b031681565b5f5f838360405160200161067f929190613eea565b60408051601f1981840301815291815281516020928301206001600160a01b0386165f9081526005909352908220909250905b81548110156106ec57828282815481106106ce576106ce613f16565b905f5260205f200154036106e4578093506106ec565b6001016106b2565b5050505b92915050565b5f88815260046020526040812054899082036107255760405163938a92b760e01b815260040160405180910390fd5b5f818152600460205260409020600301546001600160a01b0316331461076b5761074e81612b33565b61076b5760405163dc64d0ad60e01b815260040160405180910390fd5b89895f8282604051602001610781929190613f2a565b60408051601f1981840301815291815281516020928301205f8181526006909352912060040154909150156107c95760405163d96b03b160e01b815260040160405180910390fd5b5f8d8d6040516020016107dd929190613f2a565b60408051601f1981840301815291815281516020928301205f818152600690935290822054909250900361082457604051632abde33960e01b815260040160405180910390fd5b5f818152600660205260409020600701548714610854576040516301c0b3dd60e61b815260040160405180910390fd5b8988141580610863575060208a115b15610881576040516373d8ccd360e11b815260040160405180910390fd5b60208a10156108d7575f8181526009602052604090205461ffff16156108ba576040516355cbc83160e01b815260040160405180910390fd5b5f818152600960205260409020805461ffff191661ffff8c161790555b8c6040516108e59190613f3b565b604051908190038120338252908f9083907f01d10894cb2a39778aae51e234b669f70a74328f07e58e67a2caca4c5a3c86ff9060200160405180910390a460065f8281526020019081526020015f206007015f018f908060018154018082558091505060019003905f5260205f20015f9091909190915090816109689190613fde565b505f818152600660209081526040822060080180546001810182559083529120018c90559450505050509998505050505050505050565b5f5f83336040516020016109b4929190613eea565b60408051601f1981840301815291815281516020928301205f818152600493849052918220909350909101905b81548110156106ec57848282815481106109fd576109fd613f16565b905f5260205f20015403610a13578093506106ec565b6001016109e1565b60605f826001600160401b03811115610a3657610a36613462565b604051908082528060200260200182016040528015610a9f57816020015b610a8c6040518060a001604052805f8152602001606081526020015f81526020015f6001600160a01b03168152602001606081525090565b815260200190600190039081610a545790505b5090505f5b83811015610c2c5760045f868684818110610ac157610ac1613f16565b9050602002013581526020019081526020015f206040518060a00160405290815f8201548152602001600182018054610af990613f46565b80601f0160208091040260200160405190810160405280929190818152602001828054610b2590613f46565b8015610b705780601f10610b4757610100808354040283529160200191610b70565b820191905f5260205f20905b815481529060010190602001808311610b5357829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b0316815260200160048201805480602002602001604051908101604052809291908181526020018280548015610bfe57602002820191905f5260205f20905b815481526020019060010190808311610bea575b505050505081525050828281518110610c1957610c19613f16565b6020908102919091010152600101610aa4565b509392505050565b5f828152600660205260408120600701548190610c53906001906140a7565b90508083148015610c7457505f8481526009602052604090205461ffff1615155b15610cca575f84815260096020526040812054610c9a9060019061ffff1681901b6140a7565b5f868152600760209081526040808320888452909152902054811663ffffffff90811691161492506106f0915050565b50505f91825260076020908152604080842092845291905290205463ffffffff9081161490565b6060600c5f8381526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b82821015610dcb578382905f5260205f20018054610d4090613f46565b80601f0160208091040260200160405190810160405280929190818152602001828054610d6c90613f46565b8015610db75780601f10610d8e57610100808354040283529160200191610db7565b820191905f5260205f20905b815481529060010190602001808311610d9a57829003601f168201915b505050505081526020019060010190610d23565b505050509050919050565b60408051808201909152606081525f60208201525f848484604051602001610e00939291906140ba565b60408051601f1981840301815282825280516020918201205f818152600b90925290829020838301909252815490935082908290610e3d90613f46565b80601f0160208091040260200160405190810160405280929190818152602001828054610e6990613f46565b8015610eb45780601f10610e8b57610100808354040283529160200191610eb4565b820191905f5260205f20905b815481529060010190602001808311610e9757829003601f168201915b50505091835250506001919091015460ff16151560209091015295945050505050565b5f828152600660205260408120600701805460609291829185908110610eff57610eff613f16565b905f5260205f20018054610f1290613f46565b80601f0160208091040260200160405190810160405280929190818152602001828054610f3e90613f46565b8015610f895780601f10610f6057610100808354040283529160200191610f89565b820191905f5260205f20905b815481529060010190602001808311610f6c57829003601f168201915b5050505f888152600660205260408120600801805494955090939092508791508110610fb757610fb7613f16565b5f91825260209091200154919350909150505b9250929050565b60408051808201909152606081525f60208201525f828152600b602052604090819020815180830190925280548290829061100b90613f46565b80601f016020809104026020016040519081016040528092919081815260200182805461103790613f46565b80156110825780601f1061105957610100808354040283529160200191611082565b820191905f5260205f20905b81548152906001019060200180831161106557829003601f168201915b50505091835250506001919091015460ff16151560209091015292915050565b6110aa6132b8565b5f83836040516020016110be929190613f2a565b60408051601f1981840301815282825280516020918201205f81815260068352839020610100850190935282548452600183018054919550918401919061110490613f46565b80601f016020809104026020016040519081016040528092919081815260200182805461113090613f46565b801561117b5780601f106111525761010080835404028352916020019161117b565b820191905f5260205f20905b81548152906001019060200180831161115e57829003601f168201915b505050505081526020016002820154815260200160038201805461119e90613f46565b80601f01602080910402602001604051908101604052809291908181526020018280546111ca90613f46565b80156112155780601f106111ec57610100808354040283529160200191611215565b820191905f5260205f20905b8154815290600101906020018083116111f857829003601f168201915b50505050508152602001600482015481526020016005820154815260200160068201548152602001600782016040518060400160405290815f8201805480602002602001604051908101604052809291908181526020015f905b82821015611317578382905f5260205f2001805461128c90613f46565b80601f01602080910402602001604051908101604052809291908181526020018280546112b890613f46565b80156113035780601f106112da57610100808354040283529160200191611303565b820191905f5260205f20905b8154815290600101906020018083116112e657829003601f168201915b50505050508152602001906001019061126f565b5050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561136c57602002820191905f5260205f20905b815481526020019060010190808311611358575b50505091909252505050905250949350505050565b6113b96040518060a001604052805f8152602001606081526020015f81526020015f6001600160a01b03168152602001606081525090565b5f82336040516020016113cd929190613eea565b60408051601f1981840301815282825280516020918201205f8181526004835283902060a0850190935282548452600183018054919550918401919061141290613f46565b80601f016020809104026020016040519081016040528092919081815260200182805461143e90613f46565b80156114895780601f1061146057610100808354040283529160200191611489565b820191905f5260205f20905b81548152906001019060200180831161146c57829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b031681526020016004820180548060200260200160405190810160405280929190818152602001828054801561151757602002820191905f5260205f20905b815481526020019060010190808311611503575b505050505081525050915050919050565b6001600160a01b0381165f9081526005602090815260409182902080548351818402810184019094528084526060939283018282801561158557602002820191905f5260205f20905b815481526020019060010190808311611571575b50505050509050919050565b5f82815260046020526040812054839082036115c05760405163938a92b760e01b815260040160405180910390fd5b5f818152600460205260409020600301546001600160a01b03163314611606576115e981612b33565b6116065760405163dc64d0ad60e01b815260040160405180910390fd5b5f848460405160200161161a929190613f2a565b60408051601f1981840301815291815281516020928301205f81815260069093529120549091501561165f576040516303448eef60e51b815260040160405180910390fd5b5f8581526004602081815260408084209092018054600181018255908452922090910182905551611691908590613f3b565b60405190819003812033825290869083907fb018e47bdb983351e1bee22415a8f41eef5c2bf1c43c6d3d0992e678ae762df59060200160405180910390a4604080515f81830181815260608301909352918190836116ff565b60608152602001906001900390816116ea5790505b5081526020015f60405190808252806020026020018201604052801561172f578160200160208202803683370190505b50905260408051610100810182528481528151602081810184525f8083528184019283528385018c9052606084018b9052608084018190524260a085015260c0840181905260e084018690528781526006909152929092208151815591519293509160018201906117a09082613fde565b5060408201516002820155606082015160038201906117bf9082613fde565b506080820151600482015560a0820151600582015560c0820151600682015560e08201518051805160078401916117fb91839160200190613312565b5060208281015180516118149260018501920190613366565b50949998505050505050505050565b5f8683604051602001611837929190613f2a565b60408051601f1981840301815291815281516020928301205f818152600690935290822054909250900361187e57604051632abde33960e01b815260040160405180910390fd5b5f81815260066020526040902060070180546119779190889081106118a5576118a5613f16565b905f5260205f200180546118b890613f46565b80601f01602080910402602001604051908101604052809291908181526020018280546118e490613f46565b801561192f5780601f106119065761010080835404028352916020019161192f565b820191905f5260205f20905b81548152906001019060200180831161191257829003601f168201915b50505050508b88878d8d8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508d92508a91508f9050612bcb565b611982818588612421565b156119a057604051636d680ca160e11b815260040160405180910390fd5b5f89898c6040516020016119b6939291906140ba565b60408051601f1981840301815291815281516020928301205f818152600b90935291208054919250906119e890613f46565b90505f036119fe576119fc8a8a8d5f611bdb565b505b611a09828689612d9a565b611a138288610c34565b15611a55575f828152600860205260408120805460019290611a3a90849061ffff166140cc565b92506101000a81548161ffff021916908361ffff1602179055505b6002546040516340c10f1960e01b8152336004820152670de0b6b3a764000060248201526001600160a01b03909116906340c10f19906044015f604051808303815f87803b158015611aa5575f5ffd5b505af1158015611ab7573d5f5f3e3d5ffd5b505050508989604051611acb9291906140e6565b6040805191829003822089835260ff88166020840152918d917f99c916ccb39a9f4db66cb0bb94ca9feafab1c68cc012e20d9907590971ad2e3c910160405180910390a35050505050505050505050565b5f6060805f5f5f6060611b2d612e03565b611b35612e34565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b5f818152600c60205260408120815b8154811015611bd2578585604051611b869291906140e6565b6040518091039020828281548110611ba057611ba0613f16565b905f5260205f2001604051611bb591906140f5565b604051809103902003611bca57809250611bd2565b600101611b6d565b50509392505050565b5f848484604051602001611bf1939291906140ba565b60408051601f1981840301815282825280516020918201206060601f890183900490920284018201835291830187815291935082919088908890819085018382808284375f920182905250938552505050841515602092830152838152600b9091526040902081518190611c659082613fde565b506020918201516001918201805460ff19169115159190911790555f858152600c8352604081208054928301815581529190912001611ca5858783614166565b508484604051611cb69291906140e6565b6040519081900381209082907f0d82162721d4869b33130d645e1207517c6e97d556c3981cf14a3278423be32c905f90a3949350505050565b5f8581526004602052604081205486908203611d1e5760405163938a92b760e01b815260040160405180910390fd5b5f818152600460205260409020600301546001600160a01b03163314611d6457611d4781612b33565b611d645760405163dc64d0ad60e01b815260040160405180910390fd5b86865f8282604051602001611d7a929190613f2a565b60408051601f1981840301815291815281516020928301205f818152600690935291206004015490915015611dc25760405163d96b03b160e01b815260040160405180910390fd5b5f8a8a604051602001611dd6929190613f2a565b60408051601f1981840301815291815281516020928301205f81815260088452828120546006909452919091206007015490925061ffff90911614611e2e57604051632e1b8f4960e11b815260040160405180910390fd5b885f03611e4e57604051631b6fdfeb60e01b815260040160405180910390fd5b86515f03611e6f57604051637f19edc960e11b815260040160405180910390fd5b5f805b5f83815260066020526040902060080154811015611eca575f838152600660205260409020600801805482908110611eac57611eac613f16565b905f5260205f20015482611ec0919061421a565b9150600101611e72565b50898114611eeb57604051631b6fdfeb60e01b815260040160405180910390fd5b5f828152600660205260409020600101611f058982613fde565b505f828152600660208181526040808420600481018f90559092018c9055600a905290205460ff16611fee575f828152600a60208181526040808420805460ff1916600117905560069091528220600701549190611f6490839061422d565b611f7690670de0b6b3a764000061422d565b90505f82118015611f8657505f81115b15611feb576002546040516340c10f1960e01b8152336004820152602481018390526001600160a01b03909116906340c10f19906044015f604051808303815f87803b158015611fd4575f5ffd5b505af1158015611fe6573d5f5f3e3d5ffd5b505050505b50505b8a604051611ffc9190613f3b565b604051908190038120338252908d9084907fb2e54a2138908ebdd7da28708ec0bc3f1498b96fb7b0db337edef5ceeb41b16b9060200160405180910390a4509a9950505050505050505050565b5f83815260046020526040812054849082036120785760405163938a92b760e01b815260040160405180910390fd5b5f818152600460205260409020600301546001600160a01b031633146120be576120a181612b33565b6120be5760405163dc64d0ad60e01b815260040160405180910390fd5b5f8681526006602052604081205490036120eb57604051632abde33960e01b815260040160405180910390fd5b84846040516020016120fe929190613f2a565b60405160208183030381529060405280519060200120861461213357604051630ef4797b60e31b815260040160405180910390fd5b5f8681526006602052604081208181559061215160018301826133ab565b600282015f9055600382015f61216791906133ab565b5f600483018190556005830181905560068301819055600783019061218c82826133e5565b612199600183015f613400565b5050505f868152600460208190526040909120018054909150841015806121da5750868185815481106121ce576121ce613f16565b905f5260205f20015414155b156121f8576040516337c7f25560e01b815260040160405180910390fd5b80548190612208906001906140a7565b8154811061221857612218613f16565b905f5260205f20015481858154811061223357612233613f16565b905f5260205f2001819055508080548061224f5761224f614244565b600190038181905f5260205f20015f90559055846040516122709190613f3b565b60405190819003812033825290879089907f0e1bf50f5cca6659c62146db5b60160121a3752011b621d8c8953a1e0e23d0489060200160405180910390a45060019695505050505050565b5f81336040516020016122cf929190613eea565b60408051601f1981840301815291815281516020928301205f818152600490935291205490915015612314576040516324bf796160e11b815260040160405180910390fd5b6040805160a0810182528281526020808201858152428385015233606084015283515f808252818401865260808501919091528581526004909252929020815181559151909190600182019061236a9082613fde565b506040820151600282015560608201516003820180546001600160a01b0319166001600160a01b03909216919091179055608082015180516123b6916004840191602090910190613366565b5050335f81815260056020908152604080832080546001810182559084529190922001849055519091506123eb908490613f3b565b6040519081900381209083907fb8be57bce74a717a1bbd4acf428df655720fce75c1854b02a88484388df241a8905f90a4919050565b5f60208360ff161115612447576040516359b452ef60e01b815260040160405180910390fd5b505f838152600760209081526040808320848452909152902054600160ff84161b1663ffffffff1615159392505050565b5f858152600b60205260408120805461249090613f46565b90505f036124b157604051631512312160e01b815260040160405180910390fd5b8484846040516020016124c6939291906140ba565b6040516020818303038152906040528051906020012086146124fb576040516332c83a2360e21b815260040160405180910390fd5b5f868152600b602052604081209061251382826133ab565b506001908101805460ff191690555f848152600c6020526040902080549091829161253e91906140a7565b8154811061254e5761254e613f16565b905f5260205f200181848154811061256857612568613f16565b905f5260205f2001908161257c9190614258565b508080548061258d5761258d614244565b600190038181905f5260205f20015f6125a691906133ab565b905585856040516125b89291906140e6565b6040519081900381209088907f2341a1cde752ff7f07ea331fb1668c33e2364a1262a79d78f703e622f9725400905f90a35060019695505050505050565b6125fe6132b8565b60065f8381526020019081526020015f20604051806101000160405290815f820154815260200160018201805461263490613f46565b80601f016020809104026020016040519081016040528092919081815260200182805461266090613f46565b80156126ab5780601f10612682576101008083540402835291602001916126ab565b820191905f5260205f20905b81548152906001019060200180831161268e57829003601f168201915b50505050508152602001600282015481526020016003820180546126ce90613f46565b80601f01602080910402602001604051908101604052809291908181526020018280546126fa90613f46565b80156127455780601f1061271c57610100808354040283529160200191612745565b820191905f5260205f20905b81548152906001019060200180831161272857829003601f168201915b50505050508152602001600482015481526020016005820154815260200160068201548152602001600782016040518060400160405290815f8201805480602002602001604051908101604052809291908181526020015f905b82821015612847578382905f5260205f200180546127bc90613f46565b80601f01602080910402602001604051908101604052809291908181526020018280546127e890613f46565b80156128335780601f1061280a57610100808354040283529160200191612833565b820191905f5260205f20905b81548152906001019060200180831161281657829003601f168201915b50505050508152602001906001019061279f565b5050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561289c57602002820191905f5260205f20905b815481526020019060010190808311612888575b5050509190925250505090525092915050565b5f83815260046020526040812054849082036128de5760405163938a92b760e01b815260040160405180910390fd5b5f818152600460205260409020600301546001600160a01b031633146129245761290781612b33565b6129245760405163dc64d0ad60e01b815260040160405180910390fd5b8333604051602001612937929190613eea565b60405160208183030381529060405280519060200120851461296c576040516327a5901560e11b815260040160405180910390fd5b5f85815260046020819052604090912001541561299b5760405162227f7760ea1b815260040160405180910390fd5b5f858152600460205260408120818155906129b960018301826133ab565b5f600283018190556003830180546001600160a01b03191690556129e1906004840190613400565b5050335f90815260056020526040902080548690829086908110612a0757612a07613f16565b905f5260205f20015414612a2e576040516337c7f25560e01b815260040160405180910390fd5b80548190612a3e906001906140a7565b81548110612a4e57612a4e613f16565b905f5260205f200154818581548110612a6957612a69613f16565b905f5260205f20018190555080805480612a8557612a85614244565b600190038181905f5260205f20015f90559055336001600160a01b031685604051612ab09190613f3b565b6040519081900381209088907feda2cc8e002ead8000b1e0c1debfc9a88bd7ee6e94b8dc0763db17849fcf0389905f90a450600195945050505050565b5f805b5f83815260066020526040902060070154811015612b2a57612b128382610c34565b15155f03612b2257505f92915050565b600101612af0565b50600192915050565b6003545f906001600160a01b031615612bc45760035460405163e124bdd960e01b8152600481018490523360248201526001600160a01b039091169063e124bdd990604401602060405180830381865afa158015612b93573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612bb79190614318565b15612bc457506001919050565b505f919050565b5f81815260046020908152604080832060030154815160a08101909252606b8083526001600160a01b0390911693926143489083013980516020918201208b518c830120885189840120604080519485019390935291830152606082018b9052608082018a905260ff891660a083015260c082015260e08101869052610100016040516020818303038152906040528051906020012090505f612c6d82612e61565b90505f612c7a8287612e8d565b9050836001600160a01b0316816001600160a01b031614612cf85760405162461bcd60e51b815260206004820152602d60248201527f496e76616c6964207369676e61747572653a204e6f74207369676e656420627960448201526c10313ab1b5b2ba1037bbb732b960991b60648201526084015b60405180910390fd5b6001600160a01b0381165f908152600d602090815260408083208a845290915290205460ff1615612d605760405162461bcd60e51b8152602060048201526012602482015271139bdb98d948185b1c9958591e481d5cd95960721b6044820152606401612cef565b6001600160a01b03165f908152600d6020908152604080832098835297905295909520805460ff1916600117905550505050505050505050565b60208260ff1610612dbe576040516359b452ef60e01b815260040160405180910390fd5b5f928352600760209081526040808520928552919052909120805463ffffffff600160ff9094169390931b83169281169290921763ffffffff19909216919091179055565b6060612e2f7f00000000000000000000000000000000000000000000000000000000000000005f612eb5565b905090565b6060612e2f7f00000000000000000000000000000000000000000000000000000000000000006001612eb5565b5f6106f0612e6d612f5e565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f612e9b8686613087565b925092509250612eab82826130d0565b5090949350505050565b606060ff8314612ecf57612ec88361318c565b90506106f0565b818054612edb90613f46565b80601f0160208091040260200160405190810160405280929190818152602001828054612f0790613f46565b8015612f525780601f10612f2957610100808354040283529160200191612f52565b820191905f5260205f20905b815481529060010190602001808311612f3557829003601f168201915b505050505090506106f0565b5f306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015612fb657507f000000000000000000000000000000000000000000000000000000000000000046145b15612fe057507f000000000000000000000000000000000000000000000000000000000000000090565b612e2f604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b5f5f5f83516041036130be576020840151604085015160608601515f1a6130b0888285856131c9565b9550955095505050506130c9565b505081515f91506002905b9250925092565b5f8260038111156130e3576130e3614333565b036130ec575050565b600182600381111561310057613100614333565b0361311e5760405163f645eedf60e01b815260040160405180910390fd5b600282600381111561313257613132614333565b036131535760405163fce698f760e01b815260048101829052602401612cef565b600382600381111561316757613167614333565b03613188576040516335e2f38360e21b815260048101829052602401612cef565b5050565b60605f61319883613291565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561320257505f91506003905082613287565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613253573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661327e57505f925060019150829050613287565b92505f91508190505b9450945094915050565b5f60ff8216601f8111156106f057604051632cd44ac360e21b815260040160405180910390fd5b6040518061010001604052805f8152602001606081526020015f8152602001606081526020015f81526020015f81526020015f815260200161330d604051806040016040528060608152602001606081525090565b905290565b828054828255905f5260205f20908101928215613356579160200282015b8281111561335657825182906133469082613fde565b5091602001919060010190613330565b5061336292915061341b565b5090565b828054828255905f5260205f2090810192821561339f579160200282015b8281111561339f578251825591602001919060010190613384565b50613362929150613437565b5080546133b790613f46565b5f825580601f106133c6575050565b601f0160209004905f5260205f20908101906133e29190613437565b50565b5080545f8255905f5260205f20908101906133e2919061341b565b5080545f8255905f5260205f20908101906133e29190613437565b80821115613362575f61342e82826133ab565b5060010161341b565b5b80821115613362575f8155600101613438565b5f6020828403121561345b575f5ffd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112613485575f5ffd5b8135602083015f5f6001600160401b038411156134a4576134a4613462565b50604051601f19601f85018116603f011681018181106001600160401b03821117156134d2576134d2613462565b6040528381529050808284018710156134e9575f5ffd5b838360208301375f602085830101528094505050505092915050565b80356001600160a01b038116811461351b575f5ffd5b919050565b5f5f60408385031215613531575f5ffd5b82356001600160401b03811115613546575f5ffd5b61355285828601613476565b92505061356160208401613505565b90509250929050565b5f5f83601f84011261357a575f5ffd5b5081356001600160401b03811115613590575f5ffd5b6020830191508360208260051b8501011115610fca575f5ffd5b5f5f5f5f5f5f5f5f5f60e08a8c0312156135c2575f5ffd5b89356001600160401b038111156135d7575f5ffd5b6135e38c828d01613476565b99505060208a0135975060408a01356001600160401b03811115613605575f5ffd5b6136118c828d01613476565b97505060608a0135955060808a01356001600160401b03811115613633575f5ffd5b61363f8c828d0161356a565b90965094505060a08a01356001600160401b0381111561365d575f5ffd5b6136698c828d0161356a565b9a9d999c50979a9699959894979660c00135949350505050565b5f5f60408385031215613694575f5ffd5b82356001600160401b038111156136a9575f5ffd5b6136b585828601613476565b95602094909401359450505050565b5f5f602083850312156136d5575f5ffd5b82356001600160401b038111156136ea575f5ffd5b6136f68582860161356a565b90969095509350505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b805182525f602082015160a0602085015261374e60a0850182613702565b604084810151908601526060808501516001600160a01b031690860152608080850151868303918701919091528051808352602091820193505f9291909101905b808310156137b2578351825260208201915060208401935060018301925061378f565b5095945050505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561381357603f198786030184526137fe858351613730565b945060209384019391909101906001016137e2565b50929695505050505050565b5f5f60408385031215613830575f5ffd5b50508035926020909101359150565b5f82825180855260208501945060208160051b830101602085015f5b8381101561388d57601f19858403018852613877838351613702565b602098890198909350919091019060010161385b565b50909695505050505050565b602081525f6138ab602083018461383f565b9392505050565b5f5f83601f8401126138c2575f5ffd5b5081356001600160401b038111156138d8575f5ffd5b602083019150836020828501011115610fca575f5ffd5b5f5f5f60408486031215613901575f5ffd5b83356001600160401b03811115613916575f5ffd5b613922868287016138b2565b909790965060209590950135949350505050565b602081525f8251604060208401526139516060840182613702565b90506020840151151560408401528091505092915050565b604081525f61397b6040830185613702565b90508260208301529392505050565b602081525f6138ab6020830184613702565b5f5f604083850312156139ad575f5ffd5b8235915060208301356001600160401b038111156139c9575f5ffd5b6139d585828601613476565b9150509250929050565b5f8151808452602084019350602083015f5b82811015613a0f5781518652602095860195909101906001016139f1565b5093949350505050565b5f815160408452613a2d604085018261383f565b905060208301518482036020860152613a4682826139df565b95945050505050565b60208152815160208201525f60208301516101006040840152613a76610120840182613702565b9050604084015160608401526060840151601f19848303016080850152613a9d8282613702565b915050608084015160a084015260a084015160c084015260c084015160e084015260e0840151601f1984830301610100850152613a468282613a19565b5f60208284031215613aea575f5ffd5b81356001600160401b03811115613aff575f5ffd5b613b0b84828501613476565b949350505050565b602081525f6138ab6020830184613730565b5f60208284031215613b35575f5ffd5b6138ab82613505565b602080825282518282018190525f918401906040840190835b81811015613b75578351835260209384019390920191600101613b57565b509095945050505050565b803560ff8116811461351b575f5ffd5b5f5f5f5f5f5f5f5f5f6101008a8c031215613ba9575f5ffd5b8935985060208a01356001600160401b03811115613bc5575f5ffd5b613bd18c828d016138b2565b90995097505060408a0135955060608a0135945060808a01359350613bf860a08b01613b80565b925060c08a01356001600160401b03811115613c12575f5ffd5b613c1e8c828d01613476565b92505060e08a01356001600160401b03811115613c39575f5ffd5b613c458c828d01613476565b9150509295985092959850929598565b60ff60f81b8816815260e060208201525f613c7360e0830189613702565b8281036040840152613c858189613702565b606084018890526001600160a01b038716608085015260a0840186905283810360c08501529050613cb681856139df565b9a9950505050505050505050565b80151581146133e2575f5ffd5b5f5f5f5f60608587031215613ce4575f5ffd5b84356001600160401b03811115613cf9575f5ffd5b613d05878288016138b2565b909550935050602085013591506040850135613d2081613cc4565b939692955090935050565b5f5f5f5f5f60a08688031215613d3f575f5ffd5b8535945060208601356001600160401b03811115613d5b575f5ffd5b613d6788828901613476565b945050604086013592506060860135915060808601356001600160401b03811115613d90575f5ffd5b613d9c88828901613476565b9150509295509295909350565b5f5f5f5f60808587031215613dbc575f5ffd5b843593506020850135925060408501356001600160401b03811115613ddf575f5ffd5b613deb87828801613476565b949793965093946060013593505050565b5f5f5f60608486031215613e0e575f5ffd5b83359250613e1e60208501613b80565b929592945050506040919091013590565b5f5f5f5f5f60808688031215613e43575f5ffd5b8535945060208601356001600160401b03811115613e5f575f5ffd5b613e6b888289016138b2565b9699909850959660408101359660609091013595509350505050565b5f5f5f60608486031215613e99575f5ffd5b8335925060208401356001600160401b03811115613eb5575f5ffd5b613ec186828701613476565b93969395505050506040919091013590565b5f81518060208401855e5f93019283525090919050565b5f613ef58285613ed3565b60609390931b6bffffffffffffffffffffffff191683525050601401919050565b634e487b7160e01b5f52603260045260245ffd5b8281525f613b0b6020830184613ed3565b5f6138ab8284613ed3565b600181811c90821680613f5a57607f821691505b602082108103613f7857634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115613fc557805f5260205f20601f840160051c81016020851015613fa35750805b601f840160051c820191505b81811015613fc2575f8155600101613faf565b50505b505050565b5f19600383901b1c191660019190911b1790565b81516001600160401b03811115613ff757613ff7613462565b61400b816140058454613f46565b84613f7e565b6020601f821160018114614038575f83156140265750848201515b6140308482613fca565b855550613fc2565b5f84815260208120601f198516915b828110156140675787850151825560209485019460019092019101614047565b508482101561408457868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b5f52601160045260245ffd5b818103818111156106f0576106f0614093565b82848237909101908152602001919050565b61ffff81811683821601908111156106f0576106f0614093565b818382375f9101908152919050565b5f5f835461410281613f46565b600182168015614119576001811461412e5761415b565b60ff198316865281151582028601935061415b565b865f5260205f205f5b8381101561415357815488820152600190910190602001614137565b505081860193505b509195945050505050565b6001600160401b0383111561417d5761417d613462565b6141918361418b8354613f46565b83613f7e565b5f601f8411600181146141bd575f85156141ab5750838201355b6141b58682613fca565b845550613fc2565b5f83815260208120601f198716915b828110156141ec57868501358255602094850194600190920191016141cc565b5086821015614208575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b808201808211156106f0576106f0614093565b80820281158282048414176106f0576106f0614093565b634e487b7160e01b5f52603160045260245ffd5b818103614263575050565b61426d8254613f46565b6001600160401b0381111561428457614284613462565b614292816140058454613f46565b5f601f8211600181146142b5575f83156140265750848201546140308482613fca565b5f8581526020808220868352908220601f198616925b838110156142eb57828601548255600195860195909101906020016142cb565b508583101561430857818501545f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f60208284031215614328575f5ffd5b81516138ab81613cc4565b634e487b7160e01b5f52602160045260245ffdfe53746f7261676544617461286279746573206368756e6b4349442c6279746573333220626c6f636b4349442c75696e74323536206368756e6b496e6465782c75696e743820626c6f636b496e6465782c6279746573206e6f646549642c75696e74323536206e6f6e636529a2646970667358221220883e68ee3b2fef6a5be65f4d8fe86717a1ce2b3f59f9cb4ec2e9c43ba27ae3a664736f6c634300081c0033",
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

// GetPeerBlockByCid is a free data retrieval call binding the contract method 0x4d15ebbd.
//
// Solidity: function getPeerBlockByCid(bytes peerId, bytes32 cid) view returns((bytes,bool) peerBlock)
func (_Storage *StorageCaller) GetPeerBlockByCid(opts *bind.CallOpts, peerId []byte, cid [32]byte) (IStoragePeerBlock, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getPeerBlockByCid", peerId, cid)

	if err != nil {
		return *new(IStoragePeerBlock), err
	}

	out0 := *abi.ConvertType(out[0], new(IStoragePeerBlock)).(*IStoragePeerBlock)

	return out0, err

}

// GetPeerBlockByCid is a free data retrieval call binding the contract method 0x4d15ebbd.
//
// Solidity: function getPeerBlockByCid(bytes peerId, bytes32 cid) view returns((bytes,bool) peerBlock)
func (_Storage *StorageSession) GetPeerBlockByCid(peerId []byte, cid [32]byte) (IStoragePeerBlock, error) {
	return _Storage.Contract.GetPeerBlockByCid(&_Storage.CallOpts, peerId, cid)
}

// GetPeerBlockByCid is a free data retrieval call binding the contract method 0x4d15ebbd.
//
// Solidity: function getPeerBlockByCid(bytes peerId, bytes32 cid) view returns((bytes,bool) peerBlock)
func (_Storage *StorageCallerSession) GetPeerBlockByCid(peerId []byte, cid [32]byte) (IStoragePeerBlock, error) {
	return _Storage.Contract.GetPeerBlockByCid(&_Storage.CallOpts, peerId, cid)
}

// GetPeerBlockById is a free data retrieval call binding the contract method 0x4ed0e632.
//
// Solidity: function getPeerBlockById(bytes32 id) view returns((bytes,bool) peerBlock)
func (_Storage *StorageCaller) GetPeerBlockById(opts *bind.CallOpts, id [32]byte) (IStoragePeerBlock, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getPeerBlockById", id)

	if err != nil {
		return *new(IStoragePeerBlock), err
	}

	out0 := *abi.ConvertType(out[0], new(IStoragePeerBlock)).(*IStoragePeerBlock)

	return out0, err

}

// GetPeerBlockById is a free data retrieval call binding the contract method 0x4ed0e632.
//
// Solidity: function getPeerBlockById(bytes32 id) view returns((bytes,bool) peerBlock)
func (_Storage *StorageSession) GetPeerBlockById(id [32]byte) (IStoragePeerBlock, error) {
	return _Storage.Contract.GetPeerBlockById(&_Storage.CallOpts, id)
}

// GetPeerBlockById is a free data retrieval call binding the contract method 0x4ed0e632.
//
// Solidity: function getPeerBlockById(bytes32 id) view returns((bytes,bool) peerBlock)
func (_Storage *StorageCallerSession) GetPeerBlockById(id [32]byte) (IStoragePeerBlock, error) {
	return _Storage.Contract.GetPeerBlockById(&_Storage.CallOpts, id)
}

// GetPeerBlockIndexById is a free data retrieval call binding the contract method 0x95696eb2.
//
// Solidity: function getPeerBlockIndexById(bytes peerId, bytes32 cid) view returns(uint256 index)
func (_Storage *StorageCaller) GetPeerBlockIndexById(opts *bind.CallOpts, peerId []byte, cid [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getPeerBlockIndexById", peerId, cid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPeerBlockIndexById is a free data retrieval call binding the contract method 0x95696eb2.
//
// Solidity: function getPeerBlockIndexById(bytes peerId, bytes32 cid) view returns(uint256 index)
func (_Storage *StorageSession) GetPeerBlockIndexById(peerId []byte, cid [32]byte) (*big.Int, error) {
	return _Storage.Contract.GetPeerBlockIndexById(&_Storage.CallOpts, peerId, cid)
}

// GetPeerBlockIndexById is a free data retrieval call binding the contract method 0x95696eb2.
//
// Solidity: function getPeerBlockIndexById(bytes peerId, bytes32 cid) view returns(uint256 index)
func (_Storage *StorageCallerSession) GetPeerBlockIndexById(peerId []byte, cid [32]byte) (*big.Int, error) {
	return _Storage.Contract.GetPeerBlockIndexById(&_Storage.CallOpts, peerId, cid)
}

// GetPeersByPeerBlockCid is a free data retrieval call binding the contract method 0x4656b932.
//
// Solidity: function getPeersByPeerBlockCid(bytes32 cid) view returns(bytes[] peers)
func (_Storage *StorageCaller) GetPeersByPeerBlockCid(opts *bind.CallOpts, cid [32]byte) ([][]byte, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "getPeersByPeerBlockCid", cid)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetPeersByPeerBlockCid is a free data retrieval call binding the contract method 0x4656b932.
//
// Solidity: function getPeersByPeerBlockCid(bytes32 cid) view returns(bytes[] peers)
func (_Storage *StorageSession) GetPeersByPeerBlockCid(cid [32]byte) ([][]byte, error) {
	return _Storage.Contract.GetPeersByPeerBlockCid(&_Storage.CallOpts, cid)
}

// GetPeersByPeerBlockCid is a free data retrieval call binding the contract method 0x4656b932.
//
// Solidity: function getPeersByPeerBlockCid(bytes32 cid) view returns(bytes[] peers)
func (_Storage *StorageCallerSession) GetPeersByPeerBlockCid(cid [32]byte) ([][]byte, error) {
	return _Storage.Contract.GetPeersByPeerBlockCid(&_Storage.CallOpts, cid)
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

// AddPeerBlock is a paid mutator transaction binding the contract method 0x9a094ca2.
//
// Solidity: function addPeerBlock(bytes peerId, bytes32 cid, bool isReplica) returns(bytes32 id)
func (_Storage *StorageTransactor) AddPeerBlock(opts *bind.TransactOpts, peerId []byte, cid [32]byte, isReplica bool) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "addPeerBlock", peerId, cid, isReplica)
}

// AddPeerBlock is a paid mutator transaction binding the contract method 0x9a094ca2.
//
// Solidity: function addPeerBlock(bytes peerId, bytes32 cid, bool isReplica) returns(bytes32 id)
func (_Storage *StorageSession) AddPeerBlock(peerId []byte, cid [32]byte, isReplica bool) (*types.Transaction, error) {
	return _Storage.Contract.AddPeerBlock(&_Storage.TransactOpts, peerId, cid, isReplica)
}

// AddPeerBlock is a paid mutator transaction binding the contract method 0x9a094ca2.
//
// Solidity: function addPeerBlock(bytes peerId, bytes32 cid, bool isReplica) returns(bytes32 id)
func (_Storage *StorageTransactorSession) AddPeerBlock(peerId []byte, cid [32]byte, isReplica bool) (*types.Transaction, error) {
	return _Storage.Contract.AddPeerBlock(&_Storage.TransactOpts, peerId, cid, isReplica)
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

// DeletePeerBlock is a paid mutator transaction binding the contract method 0xf8fefaaa.
//
// Solidity: function deletePeerBlock(bytes32 id, bytes peerId, bytes32 cid, uint256 index) returns(bool)
func (_Storage *StorageTransactor) DeletePeerBlock(opts *bind.TransactOpts, id [32]byte, peerId []byte, cid [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "deletePeerBlock", id, peerId, cid, index)
}

// DeletePeerBlock is a paid mutator transaction binding the contract method 0xf8fefaaa.
//
// Solidity: function deletePeerBlock(bytes32 id, bytes peerId, bytes32 cid, uint256 index) returns(bool)
func (_Storage *StorageSession) DeletePeerBlock(id [32]byte, peerId []byte, cid [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DeletePeerBlock(&_Storage.TransactOpts, id, peerId, cid, index)
}

// DeletePeerBlock is a paid mutator transaction binding the contract method 0xf8fefaaa.
//
// Solidity: function deletePeerBlock(bytes32 id, bytes peerId, bytes32 cid, uint256 index) returns(bool)
func (_Storage *StorageTransactorSession) DeletePeerBlock(id [32]byte, peerId []byte, cid [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Storage.Contract.DeletePeerBlock(&_Storage.TransactOpts, id, peerId, cid, index)
}

// FillChunkBlock is a paid mutator transaction binding the contract method 0x83f77cdb.
//
// Solidity: function fillChunkBlock(bytes32 blockCID, bytes nodeId, bytes32 bucketId, uint256 chunkIndex, uint256 nonce, uint8 blockIndex, string name, bytes signature) returns()
func (_Storage *StorageTransactor) FillChunkBlock(opts *bind.TransactOpts, blockCID [32]byte, nodeId []byte, bucketId [32]byte, chunkIndex *big.Int, nonce *big.Int, blockIndex uint8, name string, signature []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "fillChunkBlock", blockCID, nodeId, bucketId, chunkIndex, nonce, blockIndex, name, signature)
}

// FillChunkBlock is a paid mutator transaction binding the contract method 0x83f77cdb.
//
// Solidity: function fillChunkBlock(bytes32 blockCID, bytes nodeId, bytes32 bucketId, uint256 chunkIndex, uint256 nonce, uint8 blockIndex, string name, bytes signature) returns()
func (_Storage *StorageSession) FillChunkBlock(blockCID [32]byte, nodeId []byte, bucketId [32]byte, chunkIndex *big.Int, nonce *big.Int, blockIndex uint8, name string, signature []byte) (*types.Transaction, error) {
	return _Storage.Contract.FillChunkBlock(&_Storage.TransactOpts, blockCID, nodeId, bucketId, chunkIndex, nonce, blockIndex, name, signature)
}

// FillChunkBlock is a paid mutator transaction binding the contract method 0x83f77cdb.
//
// Solidity: function fillChunkBlock(bytes32 blockCID, bytes nodeId, bytes32 bucketId, uint256 chunkIndex, uint256 nonce, uint8 blockIndex, string name, bytes signature) returns()
func (_Storage *StorageTransactorSession) FillChunkBlock(blockCID [32]byte, nodeId []byte, bucketId [32]byte, chunkIndex *big.Int, nonce *big.Int, blockIndex uint8, name string, signature []byte) (*types.Transaction, error) {
	return _Storage.Contract.FillChunkBlock(&_Storage.TransactOpts, blockCID, nodeId, bucketId, chunkIndex, nonce, blockIndex, name, signature)
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
	PeerId  common.Hash
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddPeerBlock is a free log retrieval operation binding the contract event 0x0d82162721d4869b33130d645e1207517c6e97d556c3981cf14a3278423be32c.
//
// Solidity: event AddPeerBlock(bytes32 indexed blockId, bytes indexed peerId)
func (_Storage *StorageFilterer) FilterAddPeerBlock(opts *bind.FilterOpts, blockId [][32]byte, peerId [][]byte) (*StorageAddPeerBlockIterator, error) {

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

// WatchAddPeerBlock is a free log subscription operation binding the contract event 0x0d82162721d4869b33130d645e1207517c6e97d556c3981cf14a3278423be32c.
//
// Solidity: event AddPeerBlock(bytes32 indexed blockId, bytes indexed peerId)
func (_Storage *StorageFilterer) WatchAddPeerBlock(opts *bind.WatchOpts, sink chan<- *StorageAddPeerBlock, blockId [][32]byte, peerId [][]byte) (event.Subscription, error) {

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

// ParseAddPeerBlock is a log parse operation binding the contract event 0x0d82162721d4869b33130d645e1207517c6e97d556c3981cf14a3278423be32c.
//
// Solidity: event AddPeerBlock(bytes32 indexed blockId, bytes indexed peerId)
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
	NodeId     common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChunkBlockFilled is a free log retrieval operation binding the contract event 0x99c916ccb39a9f4db66cb0bb94ca9feafab1c68cc012e20d9907590971ad2e3c.
//
// Solidity: event ChunkBlockFilled(bytes32 indexed cid, uint256 chunkIndex, uint8 blockIndex, bytes indexed nodeId)
func (_Storage *StorageFilterer) FilterChunkBlockFilled(opts *bind.FilterOpts, cid [][32]byte, nodeId [][]byte) (*StorageChunkBlockFilledIterator, error) {

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

// WatchChunkBlockFilled is a free log subscription operation binding the contract event 0x99c916ccb39a9f4db66cb0bb94ca9feafab1c68cc012e20d9907590971ad2e3c.
//
// Solidity: event ChunkBlockFilled(bytes32 indexed cid, uint256 chunkIndex, uint8 blockIndex, bytes indexed nodeId)
func (_Storage *StorageFilterer) WatchChunkBlockFilled(opts *bind.WatchOpts, sink chan<- *StorageChunkBlockFilled, cid [][32]byte, nodeId [][]byte) (event.Subscription, error) {

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

// ParseChunkBlockFilled is a log parse operation binding the contract event 0x99c916ccb39a9f4db66cb0bb94ca9feafab1c68cc012e20d9907590971ad2e3c.
//
// Solidity: event ChunkBlockFilled(bytes32 indexed cid, uint256 chunkIndex, uint8 blockIndex, bytes indexed nodeId)
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
	PeerId  common.Hash
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeletePeerBlock is a free log retrieval operation binding the contract event 0x2341a1cde752ff7f07ea331fb1668c33e2364a1262a79d78f703e622f9725400.
//
// Solidity: event DeletePeerBlock(bytes32 indexed blockId, bytes indexed peerId)
func (_Storage *StorageFilterer) FilterDeletePeerBlock(opts *bind.FilterOpts, blockId [][32]byte, peerId [][]byte) (*StorageDeletePeerBlockIterator, error) {

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

// WatchDeletePeerBlock is a free log subscription operation binding the contract event 0x2341a1cde752ff7f07ea331fb1668c33e2364a1262a79d78f703e622f9725400.
//
// Solidity: event DeletePeerBlock(bytes32 indexed blockId, bytes indexed peerId)
func (_Storage *StorageFilterer) WatchDeletePeerBlock(opts *bind.WatchOpts, sink chan<- *StorageDeletePeerBlock, blockId [][32]byte, peerId [][]byte) (event.Subscription, error) {

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

// ParseDeletePeerBlock is a log parse operation binding the contract event 0x2341a1cde752ff7f07ea331fb1668c33e2364a1262a79d78f703e622f9725400.
//
// Solidity: event DeletePeerBlock(bytes32 indexed blockId, bytes indexed peerId)
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
