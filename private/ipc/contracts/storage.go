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
	Chunks      IStorageChunk
}

// IStoragePeerBlock is an auto generated low-level Go binding around an user-defined struct.
type IStoragePeerBlock struct {
	PeerId    []byte
	IsReplica bool
}

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BlockAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockAlreadyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BlockNonexists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketInvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketNonempty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BucketNonexists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"}],\"name\":\"ChunkCIDMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileChunkDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileFullyUploaded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNameDuplicate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNonempty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileNotFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cidsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sizesLength\",\"type\":\"uint256\"}],\"name\":\"InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlockIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBlocksAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEncodedSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFileBlocksCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFileCID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLastBlockSize\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LastChunkDuplicate\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"AddFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"AddFileBlocks\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"peerId\",\"type\":\"bytes\"}],\"name\":\"AddPeerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"nodeId\",\"type\":\"bytes\"}],\"name\":\"ChunkBlockFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CreateBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CreateFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DeleteBucket\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DeleteFile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"blockId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"peerId\",\"type\":\"bytes\"}],\"name\":\"DeletePeerBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"FileUploaded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_BLOCKS_PER_FILE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BLOCK_SIZE\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"cid\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedChunkSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"cids\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkBlocksSizes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"addFileChunk\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"peerId\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isReplica\",\"type\":\"bool\"}],\"name\":\"addPeerBlock\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedFileSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"}],\"name\":\"commitFile\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedFileSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"}],\"name\":\"commitFileNew\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"createBucket\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"createFile\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deleteBucket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deleteFile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"peerId\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deletePeerBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fileFillCounter\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"fileRewardClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blockCID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"nodeId\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"fillChunkBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fulfilledBlocks\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getBucketByName\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket\",\"name\":\"bucket\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getBucketIndexByName\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"ids\",\"type\":\"bytes32[]\"}],\"name\":\"getBucketsByIds\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"files\",\"type\":\"bytes32[]\"}],\"internalType\":\"structIStorage.Bucket[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getChunkByIndex\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getFileById\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"chunkCIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkSize\",\"type\":\"uint256[]\"}],\"internalType\":\"structIStorage.Chunk\",\"name\":\"chunks\",\"type\":\"tuple\"}],\"internalType\":\"structIStorage.File\",\"name\":\"file\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getFileByName\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"fileCID\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"encodedSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"chunkCIDs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chunkSize\",\"type\":\"uint256[]\"}],\"internalType\":\"structIStorage.Chunk\",\"name\":\"chunks\",\"type\":\"tuple\"}],\"internalType\":\"structIStorage.File\",\"name\":\"file\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"getFileIndexById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getFileOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getOwnerBuckets\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"buckets\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"peerId\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"}],\"name\":\"getPeerBlockByCid\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"peerId\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isReplica\",\"type\":\"bool\"}],\"internalType\":\"structIStorage.PeerBlock\",\"name\":\"peerBlock\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getPeerBlockById\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"peerId\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"isReplica\",\"type\":\"bool\"}],\"internalType\":\"structIStorage.PeerBlock\",\"name\":\"peerBlock\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"peerId\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"}],\"name\":\"getPeerBlockIndexById\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cid\",\"type\":\"bytes32\"}],\"name\":\"getPeersByPeerBlockCid\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"peers\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"blockIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"isBlockFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"chunkIndex\",\"type\":\"uint256\"}],\"name\":\"isChunkFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"isFileFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"isFileFilledV2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIAkaveToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x610160604052348015610010575f5ffd5b5060405161486c38038061486c83398101604081905261002f916101b1565b604080518082018252600781526653746f7261676560c81b602080830191909152825180840190935260018352603160f81b9083015290610070825f610139565b6101205261007f816001610139565b61014052815160208084019190912060e052815190820120610100524660a05261010b60e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c052600280546001600160a01b0319166001600160a01b0392909216919091179055610388565b5f6020835110156101545761014d8361016b565b9050610165565b8161015f8482610276565b5060ff90505b92915050565b5f5f829050601f8151111561019e578260405163305a27a960e01b81526004016101959190610330565b60405180910390fd5b80516101a982610365565b179392505050565b5f602082840312156101c1575f5ffd5b81516001600160a01b03811681146101d7575f5ffd5b9392505050565b634e487b7160e01b5f52604160045260245ffd5b600181811c9082168061020657607f821691505b60208210810361022457634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561027157805f5260205f20601f840160051c8101602085101561024f5750805b601f840160051c820191505b8181101561026e575f815560010161025b565b50505b505050565b81516001600160401b0381111561028f5761028f6101de565b6102a38161029d84546101f2565b8461022a565b6020601f8211600181146102d5575f83156102be5750848201515b5f19600385901b1c1916600184901b17845561026e565b5f84815260208120601f198516915b8281101561030457878501518255602094850194600190920191016102e4565b508482101561032157868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80516020808301519190811015610224575f1960209190910360031b1b16919050565b60805160a05160c05160e0516101005161012051610140516144936103d95f395f612f2e01525f612efd01525f61312a01525f61310201525f61305d01525f61308701525f6130b101526144935ff3fe608060405234801561000f575f5ffd5b506004361061021e575f3560e01c80636a5d8c261161012a578063b80777ea116100b4578063f8fefaaa11610079578063f8fefaaa146105d0578063faec0542146105e3578063fc0c546a146105f6578063fd1d3c0c14610609578063fd21c2841461061c575f5ffd5b8063b80777ea1461055b578063d6d3110b14610561578063e3f787e814610574578063e4ba8a5814610587578063f8a3e41a1461059a575f5ffd5b806384b0196e116100fa57806384b0196e146104fe57806387490e181461051957806395696eb21461052c5780639a094ca21461053f5780639ccd464614610552575f5ffd5b80636a5d8c26146104945780636af0f801146104b45780636ce02363146104c757806383f77cdb146104e9575f5ffd5b80634d15ebbd116101ab578063564b81ef1161017b578063564b81ef146103d45780635a4e9564146103da5780635ecdfb531461041f5780636554cda71461043f57806368e6408f1461045f575f5ffd5b80634d15ebbd146103595780634d7dc614146103795780634ed0e6321461039a57806354fd4d50146103ad575f5ffd5b806330b91d07116101f157806330b91d07146102e0578063359b15a5146102f357806335bdb711146103065780633f383980146103265780634656b93214610339575f5ffd5b8063018c1e9c146102225780631b475ef41461025957806325ccd451146102ac578063287e677f146102cd575b5f5ffd5b610244610230366004613537565b60096020525f908152604090205460ff1681565b60405190151581526020015b60405180910390f35b610294610267366004613537565b5f90815260056020908152604080832060020154835260039182905290912001546001600160a01b031690565b6040516001600160a01b039091168152602001610250565b6102bf6102ba3660046135f1565b61062f565b604051908152602001610250565b6102bf6102db366004613681565b610974565b6102bf6102ee36600461370b565b610a00565b6102bf6103013660046137e4565b610c81565b610319610314366004613825565b610cfc565b604051610250919061391d565b610244610334366004613980565b610f15565b61034c610347366004613537565b610fd2565b60405161025091906139fa565b61036c610367366004613a50565b6110b7565b6040516102509190613a97565b61038c610387366004613980565b6111b8565b604051610250929190613aca565b61036c6103a8366004613537565b6112b2565b60408051808201825260058152640312e302e360dc1b602082015290516102509190613aeb565b466102bf565b61040a6103e8366004613980565b600660209081525f928352604080842090915290825290205463ffffffff1681565b60405163ffffffff9091168152602001610250565b61043261042d366004613afd565b611383565b6040516102509190613b7a565b61045261044d366004613c27565b611657565b6040516102509190613c60565b61024461046d366004613537565b5f9081526007602090815260408083205460059092529091206006015461ffff9091161490565b6104a76104a2366004613c72565b6117fe565b6040516102509190613c8b565b6102bf6104c2366004613afd565b611867565b6104d1620f424081565b6040516001600160401b039091168152602001610250565b6104fc6104f7366004613cdd565b611b3b565b005b610506611e34565b6040516102509796959493929190613da2565b6102bf6105273660046135f1565b611e76565b6102bf61053a366004613a50565b6120ad565b6102bf61054d366004613e11565b61212a565b6104d161040081565b426102bf565b61024461056f366004613e6f565b61223e565b6102bf610582366004613c27565b612499565b610244610595366004613ec2565b6125ff565b6105bd6105a8366004613537565b60076020525f908152604090205461ffff1681565b60405161ffff9091168152602001610250565b6102446105de366004613ef5565b612656565b6104326105f1366004613537565b6127d4565b600254610294906001600160a01b031681565b610244610617366004613f4d565b612a82565b61024461062a366004613537565b612c77565b5f84845f8282604051602001610646929190613fb0565b60408051601f1981840301815291815281516020928301205f868152600390935290822054909250900361068d5760405163938a92b760e01b815260040160405180910390fd5b5f83815260036020819052604090912001546001600160a01b031633146106c75760405163dc64d0ad60e01b815260040160405180910390fd5b5f81815260056020526040902060040154156106f65760405163d96b03b160e01b815260040160405180910390fd5b5f888860405160200161070a929190613fb0565b60408051601f1981840301815291815281516020928301205f81815260078452828120546005909452919091206006015490925061ffff9091161461076257604051632e1b8f4960e11b815260040160405180910390fd5b865f0361078257604051631b6fdfeb60e01b815260040160405180910390fd5b85515f036107a357604051637f19edc960e11b815260040160405180910390fd5b5f805b5f838152600560205260409020600701548110156107fe575f8381526005602052604090206007018054829081106107e0576107e0613fc1565b905f5260205f200154826107f49190613fe9565b91506001016107a6565b5087811461081f57604051631b6fdfeb60e01b815260040160405180910390fd5b5f8281526005602052604090206001016108398882614094565b505f8281526005602090815260408083206004018b9055600990915290205460ff1661091b575f828152600960209081526040808320805460ff1916600117905560059091528120600601549061089182600a614149565b6108a390670de0b6b3a7640000614149565b90505f821180156108b357505f81115b15610918576002546040516340c10f1960e01b8152336004820152602481018390526001600160a01b03909116906340c10f19906044015f604051808303815f87803b158015610901575f5ffd5b505af1158015610913573d5f5f3e3d5ffd5b505050505b50505b886040516109299190614160565b604051908190038120338252908b9084907fb2e54a2138908ebdd7da28708ec0bc3f1498b96fb7b0db337edef5ceeb41b16b9060200160405180910390a45098975050505050505050565b5f5f838360405160200161098992919061416b565b60408051601f1981840301815291815281516020928301206001600160a01b0386165f9081526004909352908220909250905b81548110156109f657828282815481106109d8576109d8613fc1565b905f5260205f200154036109ee578093506109f6565b6001016109bc565b5050505b92915050565b5f88885f8282604051602001610a17929190613fb0565b60408051601f1981840301815291815281516020928301205f8681526003909352908220549092509003610a5e5760405163938a92b760e01b815260040160405180910390fd5b5f83815260036020819052604090912001546001600160a01b03163314610a985760405163dc64d0ad60e01b815260040160405180910390fd5b5f8181526005602052604090206004015415610ac75760405163d96b03b160e01b815260040160405180910390fd5b5f8c8c604051602001610adb929190613fb0565b60408051601f1981840301815291815281516020928301205f8181526005909352908220549092509003610b2257604051632abde33960e01b815260040160405180910390fd5b5f818152600560205260409020600601548614610b52576040516301c0b3dd60e61b815260040160405180910390fd5b8887141580610b615750602089115b15610b7f576040516373d8ccd360e11b815260040160405180910390fd5b6020891015610bd5575f8181526008602052604090205461ffff1615610bb8576040516355cbc83160e01b815260040160405180910390fd5b5f818152600860205260409020805461ffff191661ffff8b161790555b8b604051610be39190614160565b604051908190038120338252908e9083907f01d10894cb2a39778aae51e234b669f70a74328f07e58e67a2caca4c5a3c86ff9060200160405180910390a45f81815260056020908152604082206006018054600181018255908352912001610c4b8f82614094565b505f818152600560209081526040822060070180546001810182559083529120018b905593505050509998505050505050505050565b5f5f8333604051602001610c9692919061416b565b60408051601f1981840301815291815281516020928301205f8181526003909352908220909250600401905b81548110156109f65784828281548110610cde57610cde613fc1565b905f5260205f20015403610cf4578093506109f6565b600101610cc2565b60605f826001600160401b03811115610d1757610d1761354e565b604051908082528060200260200182016040528015610d8057816020015b610d6d6040518060a001604052805f8152602001606081526020015f81526020015f6001600160a01b03168152602001606081525090565b815260200190600190039081610d355790505b5090505f5b83811015610f0d5760035f868684818110610da257610da2613fc1565b9050602002013581526020019081526020015f206040518060a00160405290815f8201548152602001600182018054610dda90613ffc565b80601f0160208091040260200160405190810160405280929190818152602001828054610e0690613ffc565b8015610e515780601f10610e2857610100808354040283529160200191610e51565b820191905f5260205f20905b815481529060010190602001808311610e3457829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b0316815260200160048201805480602002602001604051908101604052809291908181526020018280548015610edf57602002820191905f5260205f20905b815481526020019060010190808311610ecb575b505050505081525050828281518110610efa57610efa613fc1565b6020908102919091010152600101610d85565b509392505050565b5f828152600560205260408120600601548190610f3490600190614197565b90508083148015610f5557505f8481526008602052604090205461ffff1615155b15610fab575f84815260086020526040812054610f7b9060019061ffff1681901b614197565b5f868152600660209081526040808320888452909152902054811663ffffffff90811691161492506109fa915050565b50505f91825260066020908152604080842092845291905290205463ffffffff9081161490565b6060600b5f8381526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b828210156110ac578382905f5260205f2001805461102190613ffc565b80601f016020809104026020016040519081016040528092919081815260200182805461104d90613ffc565b80156110985780601f1061106f57610100808354040283529160200191611098565b820191905f5260205f20905b81548152906001019060200180831161107b57829003601f168201915b505050505081526020019060010190611004565b505050509050919050565b60408051808201909152606081525f60208201525f8484846040516020016110e1939291906141aa565b60408051601f1981840301815282825280516020918201205f818152600a9092529082902083830190925281549093508290829061111e90613ffc565b80601f016020809104026020016040519081016040528092919081815260200182805461114a90613ffc565b80156111955780601f1061116c57610100808354040283529160200191611195565b820191905f5260205f20905b81548152906001019060200180831161117857829003601f168201915b50505091835250506001919091015460ff16151560209091015295945050505050565b5f8281526005602052604081206006018054606092918291859081106111e0576111e0613fc1565b905f5260205f200180546111f390613ffc565b80601f016020809104026020016040519081016040528092919081815260200182805461121f90613ffc565b801561126a5780601f106112415761010080835404028352916020019161126a565b820191905f5260205f20905b81548152906001019060200180831161124d57829003601f168201915b5050505f88815260056020526040812060070180549495509093909250879150811061129857611298613fc1565b5f91825260209091200154919350909150505b9250929050565b60408051808201909152606081525f60208201525f828152600a60205260409081902081518083019092528054829082906112ec90613ffc565b80601f016020809104026020016040519081016040528092919081815260200182805461131890613ffc565b80156113635780601f1061133a57610100808354040283529160200191611363565b820191905f5260205f20905b81548152906001019060200180831161134657829003601f168201915b50505091835250506001919091015460ff16151560209091015292915050565b61138b6133ab565b5f838360405160200161139f929190613fb0565b60408051601f1981840301815282825280516020918201205f8181526005835283902060e085019093528254845260018301805491955091840191906113e490613ffc565b80601f016020809104026020016040519081016040528092919081815260200182805461141090613ffc565b801561145b5780601f106114325761010080835404028352916020019161145b565b820191905f5260205f20905b81548152906001019060200180831161143e57829003601f168201915b505050505081526020016002820154815260200160038201805461147e90613ffc565b80601f01602080910402602001604051908101604052809291908181526020018280546114aa90613ffc565b80156114f55780601f106114cc576101008083540402835291602001916114f5565b820191905f5260205f20905b8154815290600101906020018083116114d857829003601f168201915b505050505081526020016004820154815260200160058201548152602001600682016040518060400160405290815f8201805480602002602001604051908101604052809291908181526020015f905b828210156115ed578382905f5260205f2001805461156290613ffc565b80601f016020809104026020016040519081016040528092919081815260200182805461158e90613ffc565b80156115d95780601f106115b0576101008083540402835291602001916115d9565b820191905f5260205f20905b8154815290600101906020018083116115bc57829003601f168201915b505050505081526020019060010190611545565b5050505081526020016001820180548060200260200160405190810160405280929190818152602001828054801561164257602002820191905f5260205f20905b81548152602001906001019080831161162e575b50505091909252505050905250949350505050565b61168f6040518060a001604052805f8152602001606081526020015f81526020015f6001600160a01b03168152602001606081525090565b5f82336040516020016116a392919061416b565b60408051601f1981840301815282825280516020918201205f8181526003835283902060a085019093528254845260018301805491955091840191906116e890613ffc565b80601f016020809104026020016040519081016040528092919081815260200182805461171490613ffc565b801561175f5780601f106117365761010080835404028352916020019161175f565b820191905f5260205f20905b81548152906001019060200180831161174257829003601f168201915b5050505050815260200160028201548152602001600382015f9054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152602001600482018054806020026020016040519081016040528092919081815260200182805480156117ed57602002820191905f5260205f20905b8154815260200190600101908083116117d9575b505050505081525050915050919050565b6001600160a01b0381165f9081526004602090815260409182902080548351818402810184019094528084526060939283018282801561185b57602002820191905f5260205f20905b815481526020019060010190808311611847575b50505050509050919050565b5f82825f828260405160200161187e929190613fb0565b60408051601f1981840301815291815281516020928301205f86815260039093529082205490925090036118c55760405163938a92b760e01b815260040160405180910390fd5b5f83815260036020819052604090912001546001600160a01b031633146118ff5760405163dc64d0ad60e01b815260040160405180910390fd5b5f818152600560205260409020600401541561192e5760405163d96b03b160e01b815260040160405180910390fd5b5f8686604051602001611942929190613fb0565b60408051601f1981840301815291815281516020928301205f818152600590935291205490915015611987576040516303448eef60e51b815260040160405180910390fd5b5f87815260036020908152604080832060040180546001810182559084529190922001829055516119b9908790614160565b60405190819003812033825290889083907fb018e47bdb983351e1bee22415a8f41eef5c2bf1c43c6d3d0992e678ae762df59060200160405180910390a4604080515f8183018181526060830190935291819083611a27565b6060815260200190600190039081611a125790505b5081526020015f604051908082528060200260200182016040528015611a57578160200160208202803683370190505b5090526040805160e0810182528481528151602081810184525f8083528184019283528385018e9052606084018d9052608084018190524260a085015260c08401869052878152600590915292909220815181559151929350916001820190611ac09082614094565b506040820151600282015560608201516003820190611adf9082614094565b506080820151600482015560a0820151600582015560c0820151805180516006840191611b11918391602001906133fe565b506020828101518051611b2a9260018501920190613452565b50949b9a5050505050505050505050565b5f8683604051602001611b4f929190613fb0565b60408051601f1981840301815291815281516020928301205f8181526005909352908220549092509003611b9657604051632abde33960e01b815260040160405180910390fd5b5f8181526005602052604090206006018054611c8f919088908110611bbd57611bbd613fc1565b905f5260205f20018054611bd090613ffc565b80601f0160208091040260200160405190810160405280929190818152602001828054611bfc90613ffc565b8015611c475780601f10611c1e57610100808354040283529160200191611c47565b820191905f5260205f20905b815481529060010190602001808311611c2a57829003601f168201915b50505050508b88878d8d8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508d92508a91508f9050612cbd565b611c9a8185886125ff565b15611cb857604051636d680ca160e11b815260040160405180910390fd5b5f89898c604051602001611cce939291906141aa565b60408051601f1981840301815291815281516020928301205f818152600a9093529120805491925090611d0090613ffc565b90505f03611d1657611d148a8a8d5f61212a565b505b611d21828689612e8d565b611d2b8288610f15565b15611d6d575f828152600760205260408120805460019290611d5290849061ffff166141bc565b92506101000a81548161ffff021916908361ffff1602179055505b6002546040516340c10f1960e01b8152336004820152670de0b6b3a764000060248201526001600160a01b03909116906340c10f19906044015f604051808303815f87803b158015611dbd575f5ffd5b505af1158015611dcf573d5f5f3e3d5ffd5b505050508989604051611de39291906141d6565b6040805191829003822089835260ff88166020840152918d917f99c916ccb39a9f4db66cb0bb94ca9feafab1c68cc012e20d9907590971ad2e3c910160405180910390a35050505050505050505050565b5f6060805f5f5f6060611e45612ef6565b611e4d612f27565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b5f84845f8282604051602001611e8d929190613fb0565b60408051601f1981840301815291815281516020928301205f8681526003909352908220549092509003611ed45760405163938a92b760e01b815260040160405180910390fd5b5f83815260036020819052604090912001546001600160a01b03163314611f0e5760405163dc64d0ad60e01b815260040160405180910390fd5b5f8181526005602052604090206004015415611f3d5760405163d96b03b160e01b815260040160405180910390fd5b5f8888604051602001611f51929190613fb0565b604051602081830303815290604052805190602001209050865f03611f8957604051631b6fdfeb60e01b815260040160405180910390fd5b85515f03611faa57604051637f19edc960e11b815260040160405180910390fd5b5f805b5f83815260056020526040902060070154811015612005575f838152600560205260409020600701805482908110611fe757611fe7613fc1565b905f5260205f20015482611ffb9190613fe9565b9150600101611fad565b5087811461202657604051631b6fdfeb60e01b815260040160405180910390fd5b5f8281526005602052604090206001016120408882614094565b505f828152600560205260409081902060040189905551612062908a90614160565b604051908190038120338252908b9084907fb2e54a2138908ebdd7da28708ec0bc3f1498b96fb7b0db337edef5ceeb41b16b9060200160405180910390a45050505050949350505050565b5f818152600b60205260408120815b81548110156121215785856040516120d59291906141d6565b60405180910390208282815481106120ef576120ef613fc1565b905f5260205f200160405161210491906141e5565b60405180910390200361211957809250612121565b6001016120bc565b50509392505050565b5f848484604051602001612140939291906141aa565b60408051601f1981840301815282825280516020918201206060601f890183900490920284018201835291830187815291935082919088908890819085018382808284375f920182905250938552505050841515602092830152838152600a90915260409020815181906121b49082614094565b506020918201516001918201805460ff19169115159190911790555f858152600b83526040812080549283018155815291909120016121f4858783614256565b5084846040516122059291906141d6565b6040519081900381209082907f0d82162721d4869b33130d645e1207517c6e97d556c3981cf14a3278423be32c905f90a3949350505050565b5f83815260036020526040812054810361226b5760405163938a92b760e01b815260040160405180910390fd5b5f84815260036020819052604090912001546001600160a01b031633146122a55760405163dc64d0ad60e01b815260040160405180910390fd5b5f8581526005602052604081205490036122d257604051632abde33960e01b815260040160405180910390fd5b83836040516020016122e5929190613fb0565b60405160208183030381529060405280519060200120851461231a57604051630ef4797b60e31b815260040160405180910390fd5b5f858152600560205260408120818155906123386001830182613497565b600282015f9055600382015f61234e9190613497565b5f6004830181905560058301819055600683019061236c82826134d1565b612379600183015f6134ec565b5050505f8581526003602052604090206004018054909150831015806123b95750858184815481106123ad576123ad613fc1565b905f5260205f20015414155b156123d7576040516337c7f25560e01b815260040160405180910390fd5b805481906123e790600190614197565b815481106123f7576123f7613fc1565b905f5260205f20015481848154811061241257612412613fc1565b905f5260205f2001819055508080548061242e5761242e61430a565b600190038181905f5260205f20015f905590558360405161244f9190614160565b60405190819003812033825290869088907f0e1bf50f5cca6659c62146db5b60160121a3752011b621d8c8953a1e0e23d0489060200160405180910390a450600195945050505050565b5f81336040516020016124ad92919061416b565b60408051601f1981840301815291815281516020928301205f8181526003909352912054909150156124f2576040516324bf796160e11b815260040160405180910390fd5b6040805160a0810182528281526020808201858152428385015233606084015283515f80825281840186526080850191909152858152600390925292902081518155915190919060018201906125489082614094565b506040820151600282015560608201516003820180546001600160a01b0319166001600160a01b0390921691909117905560808201518051612594916004840191602090910190613452565b5050335f81815260046020908152604080832080546001810182559084529190922001849055519091506125c9908490614160565b6040519081900381209083907fb8be57bce74a717a1bbd4acf428df655720fce75c1854b02a88484388df241a8905f90a4919050565b5f60208360ff161115612625576040516359b452ef60e01b815260040160405180910390fd5b505f838152600660209081526040808320848452909152902054600160ff84161b1663ffffffff1615159392505050565b5f858152600a60205260408120805461266e90613ffc565b90505f0361268f57604051631512312160e01b815260040160405180910390fd5b8484846040516020016126a4939291906141aa565b6040516020818303038152906040528051906020012086146126d9576040516332c83a2360e21b815260040160405180910390fd5b5f868152600a60205260408120906126f18282613497565b506001908101805460ff191690555f848152600b6020526040902080549091829161271c9190614197565b8154811061272c5761272c613fc1565b905f5260205f200181848154811061274657612746613fc1565b905f5260205f2001908161275a919061431e565b508080548061276b5761276b61430a565b600190038181905f5260205f20015f6127849190613497565b905585856040516127969291906141d6565b6040519081900381209088907f2341a1cde752ff7f07ea331fb1668c33e2364a1262a79d78f703e622f9725400905f90a35060019695505050505050565b6127dc6133ab565b60055f8381526020019081526020015f206040518060e00160405290815f820154815260200160018201805461281190613ffc565b80601f016020809104026020016040519081016040528092919081815260200182805461283d90613ffc565b80156128885780601f1061285f57610100808354040283529160200191612888565b820191905f5260205f20905b81548152906001019060200180831161286b57829003601f168201915b50505050508152602001600282015481526020016003820180546128ab90613ffc565b80601f01602080910402602001604051908101604052809291908181526020018280546128d790613ffc565b80156129225780601f106128f957610100808354040283529160200191612922565b820191905f5260205f20905b81548152906001019060200180831161290557829003601f168201915b505050505081526020016004820154815260200160058201548152602001600682016040518060400160405290815f8201805480602002602001604051908101604052809291908181526020015f905b82821015612a1a578382905f5260205f2001805461298f90613ffc565b80601f01602080910402602001604051908101604052809291908181526020018280546129bb90613ffc565b8015612a065780601f106129dd57610100808354040283529160200191612a06565b820191905f5260205f20905b8154815290600101906020018083116129e957829003601f168201915b505050505081526020019060010190612972565b50505050815260200160018201805480602002602001604051908101604052809291908181526020018280548015612a6f57602002820191905f5260205f20905b815481526020019060010190808311612a5b575b5050509190925250505090525092915050565b5f8233604051602001612a9692919061416b565b604051602081830303815290604052805190602001208414612acb576040516327a5901560e11b815260040160405180910390fd5b5f848152600360205260408120549003612af85760405163938a92b760e01b815260040160405180910390fd5b5f8481526003602052604090206004015415612b265760405162227f7760ea1b815260040160405180910390fd5b5f84815260036020526040812081815590612b446001830182613497565b5f600283018190556003830180546001600160a01b0319169055612b6c9060048401906134ec565b5050335f90815260046020526040902080548590829085908110612b9257612b92613fc1565b905f5260205f20015414612bb9576040516337c7f25560e01b815260040160405180910390fd5b80548190612bc990600190614197565b81548110612bd957612bd9613fc1565b905f5260205f200154818481548110612bf457612bf4613fc1565b905f5260205f20018190555080805480612c1057612c1061430a565b600190038181905f5260205f20015f90559055336001600160a01b031684604051612c3b9190614160565b6040519081900381209087907feda2cc8e002ead8000b1e0c1debfc9a88bd7ee6e94b8dc0763db17849fcf0389905f90a4506001949350505050565b5f805b5f83815260056020526040902060060154811015612cb457612c9c8382610f15565b15155f03612cac57505f92915050565b600101612c7a565b50600192915050565b5f81815260036020818152604080842090920154825160a08101909352606b8084526001600160a01b039091169392916143f39083013980516020918201208b518c830120885189840120604080519485019390935291830152606082018b9052608082018a905260ff891660a083015260c082015260e08101869052610100016040516020818303038152906040528051906020012090505f612d6082612f54565b90505f612d6d8287612f80565b9050836001600160a01b0316816001600160a01b031614612deb5760405162461bcd60e51b815260206004820152602d60248201527f496e76616c6964207369676e61747572653a204e6f74207369676e656420627960448201526c10313ab1b5b2ba1037bbb732b960991b60648201526084015b60405180910390fd5b6001600160a01b0381165f908152600c602090815260408083208a845290915290205460ff1615612e535760405162461bcd60e51b8152602060048201526012602482015271139bdb98d948185b1c9958591e481d5cd95960721b6044820152606401612de2565b6001600160a01b03165f908152600c6020908152604080832098835297905295909520805460ff1916600117905550505050505050505050565b60208260ff1610612eb1576040516359b452ef60e01b815260040160405180910390fd5b5f928352600660209081526040808520928552919052909120805463ffffffff600160ff9094169390931b83169281169290921763ffffffff19909216919091179055565b6060612f227f00000000000000000000000000000000000000000000000000000000000000005f612fa8565b905090565b6060612f227f00000000000000000000000000000000000000000000000000000000000000006001612fa8565b5f6109fa612f60613051565b8360405161190160f01b8152600281019290925260228201526042902090565b5f5f5f5f612f8e868661317a565b925092509250612f9e82826131c3565b5090949350505050565b606060ff8314612fc257612fbb8361327f565b90506109fa565b818054612fce90613ffc565b80601f0160208091040260200160405190810160405280929190818152602001828054612ffa90613ffc565b80156130455780601f1061301c57610100808354040283529160200191613045565b820191905f5260205f20905b81548152906001019060200180831161302857829003601f168201915b505050505090506109fa565b5f306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480156130a957507f000000000000000000000000000000000000000000000000000000000000000046145b156130d357507f000000000000000000000000000000000000000000000000000000000000000090565b612f22604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b5f5f5f83516041036131b1576020840151604085015160608601515f1a6131a3888285856132bc565b9550955095505050506131bc565b505081515f91506002905b9250925092565b5f8260038111156131d6576131d66143de565b036131df575050565b60018260038111156131f3576131f36143de565b036132115760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115613225576132256143de565b036132465760405163fce698f760e01b815260048101829052602401612de2565b600382600381111561325a5761325a6143de565b0361327b576040516335e2f38360e21b815260048101829052602401612de2565b5050565b60605f61328b83613384565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156132f557505f9150600390508261337a565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613346573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b03811661337157505f92506001915082905061337a565b92505f91508190505b9450945094915050565b5f60ff8216601f8111156109fa57604051632cd44ac360e21b815260040160405180910390fd5b6040518060e001604052805f8152602001606081526020015f8152602001606081526020015f81526020015f81526020016133f9604051806040016040528060608152602001606081525090565b905290565b828054828255905f5260205f20908101928215613442579160200282015b8281111561344257825182906134329082614094565b509160200191906001019061341c565b5061344e929150613507565b5090565b828054828255905f5260205f2090810192821561348b579160200282015b8281111561348b578251825591602001919060010190613470565b5061344e929150613523565b5080546134a390613ffc565b5f825580601f106134b2575050565b601f0160209004905f5260205f20908101906134ce9190613523565b50565b5080545f8255905f5260205f20908101906134ce9190613507565b5080545f8255905f5260205f20908101906134ce9190613523565b8082111561344e575f61351a8282613497565b50600101613507565b5b8082111561344e575f8155600101613524565b5f60208284031215613547575f5ffd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112613571575f5ffd5b8135602083015f5f6001600160401b038411156135905761359061354e565b50604051601f19601f85018116603f011681018181106001600160401b03821117156135be576135be61354e565b6040528381529050808284018710156135d5575f5ffd5b838360208301375f602085830101528094505050505092915050565b5f5f5f5f60808587031215613604575f5ffd5b8435935060208501356001600160401b03811115613620575f5ffd5b61362c87828801613562565b9350506040850135915060608501356001600160401b0381111561364e575f5ffd5b61365a87828801613562565b91505092959194509250565b80356001600160a01b038116811461367c575f5ffd5b919050565b5f5f60408385031215613692575f5ffd5b82356001600160401b038111156136a7575f5ffd5b6136b385828601613562565b9250506136c260208401613666565b90509250929050565b5f5f83601f8401126136db575f5ffd5b5081356001600160401b038111156136f1575f5ffd5b6020830191508360208260051b85010111156112ab575f5ffd5b5f5f5f5f5f5f5f5f5f60e08a8c031215613723575f5ffd5b89356001600160401b03811115613738575f5ffd5b6137448c828d01613562565b99505060208a0135975060408a01356001600160401b03811115613766575f5ffd5b6137728c828d01613562565b97505060608a0135955060808a01356001600160401b03811115613794575f5ffd5b6137a08c828d016136cb565b90965094505060a08a01356001600160401b038111156137be575f5ffd5b6137ca8c828d016136cb565b9a9d999c50979a9699959894979660c00135949350505050565b5f5f604083850312156137f5575f5ffd5b82356001600160401b0381111561380a575f5ffd5b61381685828601613562565b95602094909401359450505050565b5f5f60208385031215613836575f5ffd5b82356001600160401b0381111561384b575f5ffd5b613857858286016136cb565b90969095509350505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b805182525f602082015160a060208501526138af60a0850182613863565b604084810151908601526060808501516001600160a01b031690860152608080850151868303918701919091528051808352602091820193505f9291909101905b8083101561391357835182526020820191506020840193506001830192506138f0565b5095945050505050565b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561397457603f1987860301845261395f858351613891565b94506020938401939190910190600101613943565b50929695505050505050565b5f5f60408385031215613991575f5ffd5b50508035926020909101359150565b5f82825180855260208501945060208160051b830101602085015f5b838110156139ee57601f198584030188526139d8838351613863565b60209889019890935091909101906001016139bc565b50909695505050505050565b602081525f613a0c60208301846139a0565b9392505050565b5f5f83601f840112613a23575f5ffd5b5081356001600160401b03811115613a39575f5ffd5b6020830191508360208285010111156112ab575f5ffd5b5f5f5f60408486031215613a62575f5ffd5b83356001600160401b03811115613a77575f5ffd5b613a8386828701613a13565b909790965060209590950135949350505050565b602081525f825160406020840152613ab26060840182613863565b90506020840151151560408401528091505092915050565b604081525f613adc6040830185613863565b90508260208301529392505050565b602081525f613a0c6020830184613863565b5f5f60408385031215613b0e575f5ffd5b8235915060208301356001600160401b03811115613b2a575f5ffd5b613b3685828601613562565b9150509250929050565b5f8151808452602084019350602083015f5b82811015613b70578151865260209586019590910190600101613b52565b5093949350505050565b60208152815160208201525f602083015160e06040840152613ba0610100840182613863565b9050604084015160608401526060840151601f19848303016080850152613bc78282613863565b915050608084015160a084015260a084015160c084015260c0840151601f198483030160e0850152805160408352613c0260408401826139a0565b9050602082015191508281036020840152613c1d8183613b40565b9695505050505050565b5f60208284031215613c37575f5ffd5b81356001600160401b03811115613c4c575f5ffd5b613c5884828501613562565b949350505050565b602081525f613a0c6020830184613891565b5f60208284031215613c82575f5ffd5b613a0c82613666565b602080825282518282018190525f918401906040840190835b81811015613cc2578351835260209384019390920191600101613ca4565b509095945050505050565b803560ff8116811461367c575f5ffd5b5f5f5f5f5f5f5f5f5f6101008a8c031215613cf6575f5ffd5b8935985060208a01356001600160401b03811115613d12575f5ffd5b613d1e8c828d01613a13565b90995097505060408a0135955060608a0135945060808a01359350613d4560a08b01613ccd565b925060c08a01356001600160401b03811115613d5f575f5ffd5b613d6b8c828d01613562565b92505060e08a01356001600160401b03811115613d86575f5ffd5b613d928c828d01613562565b9150509295985092959850929598565b60ff60f81b8816815260e060208201525f613dc060e0830189613863565b8281036040840152613dd28189613863565b606084018890526001600160a01b038716608085015260a0840186905283810360c08501529050613e038185613b40565b9a9950505050505050505050565b5f5f5f5f60608587031215613e24575f5ffd5b84356001600160401b03811115613e39575f5ffd5b613e4587828801613a13565b9095509350506020850135915060408501358015158114613e64575f5ffd5b939692955090935050565b5f5f5f5f60808587031215613e82575f5ffd5b843593506020850135925060408501356001600160401b03811115613ea5575f5ffd5b613eb187828801613562565b949793965093946060013593505050565b5f5f5f60608486031215613ed4575f5ffd5b83359250613ee460208501613ccd565b929592945050506040919091013590565b5f5f5f5f5f60808688031215613f09575f5ffd5b8535945060208601356001600160401b03811115613f25575f5ffd5b613f3188828901613a13565b9699909850959660408101359660609091013595509350505050565b5f5f5f60608486031215613f5f575f5ffd5b8335925060208401356001600160401b03811115613f7b575f5ffd5b613f8786828701613562565b93969395505050506040919091013590565b5f81518060208401855e5f93019283525090919050565b8281525f613c586020830184613f99565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b808201808211156109fa576109fa613fd5565b600181811c9082168061401057607f821691505b60208210810361402e57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561407b57805f5260205f20601f840160051c810160208510156140595750805b601f840160051c820191505b81811015614078575f8155600101614065565b50505b505050565b5f19600383901b1c191660019190911b1790565b81516001600160401b038111156140ad576140ad61354e565b6140c1816140bb8454613ffc565b84614034565b6020601f8211600181146140ee575f83156140dc5750848201515b6140e68482614080565b855550614078565b5f84815260208120601f198516915b8281101561411d57878501518255602094850194600190920191016140fd565b508482101561413a57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b80820281158282048414176109fa576109fa613fd5565b5f613a0c8284613f99565b5f6141768285613f99565b60609390931b6bffffffffffffffffffffffff191683525050601401919050565b818103818111156109fa576109fa613fd5565b82848237909101908152602001919050565b61ffff81811683821601908111156109fa576109fa613fd5565b818382375f9101908152919050565b5f5f83546141f281613ffc565b600182168015614209576001811461421e5761424b565b60ff198316865281151582028601935061424b565b865f5260205f205f5b8381101561424357815488820152600190910190602001614227565b505081860193505b509195945050505050565b6001600160401b0383111561426d5761426d61354e565b6142818361427b8354613ffc565b83614034565b5f601f8411600181146142ad575f851561429b5750838201355b6142a58682614080565b845550614078565b5f83815260208120601f198716915b828110156142dc57868501358255602094850194600190920191016142bc565b50868210156142f8575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b634e487b7160e01b5f52603160045260245ffd5b818103614329575050565b6143338254613ffc565b6001600160401b0381111561434a5761434a61354e565b614358816140bb8454613ffc565b5f601f82116001811461437b575f83156140dc5750848201546140e68482614080565b5f8581526020808220868352908220601f198616925b838110156143b15782860154825560019586019590910190602001614391565b50858310156143ce57818501545f19600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b5f52602160045260245ffdfe53746f7261676544617461286279746573206368756e6b4349442c6279746573333220626c6f636b4349442c75696e74323536206368756e6b496e6465782c75696e743820626c6f636b496e6465782c6279746573206e6f646549642c75696e74323536206e6f6e636529a2646970667358221220ce782979c5885b5aeba11679a82e343186130d6c03789fbe239774ab75acc2a264736f6c634300081c0033",
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
// Solidity: function getFileById(bytes32 id) view returns((bytes32,bytes,bytes32,string,uint256,uint256,(bytes[],uint256[])) file)
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
// Solidity: function getFileById(bytes32 id) view returns((bytes32,bytes,bytes32,string,uint256,uint256,(bytes[],uint256[])) file)
func (_Storage *StorageSession) GetFileById(id [32]byte) (IStorageFile, error) {
	return _Storage.Contract.GetFileById(&_Storage.CallOpts, id)
}

// GetFileById is a free data retrieval call binding the contract method 0xfaec0542.
//
// Solidity: function getFileById(bytes32 id) view returns((bytes32,bytes,bytes32,string,uint256,uint256,(bytes[],uint256[])) file)
func (_Storage *StorageCallerSession) GetFileById(id [32]byte) (IStorageFile, error) {
	return _Storage.Contract.GetFileById(&_Storage.CallOpts, id)
}

// GetFileByName is a free data retrieval call binding the contract method 0x5ecdfb53.
//
// Solidity: function getFileByName(bytes32 bucketId, string name) view returns((bytes32,bytes,bytes32,string,uint256,uint256,(bytes[],uint256[])) file)
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
// Solidity: function getFileByName(bytes32 bucketId, string name) view returns((bytes32,bytes,bytes32,string,uint256,uint256,(bytes[],uint256[])) file)
func (_Storage *StorageSession) GetFileByName(bucketId [32]byte, name string) (IStorageFile, error) {
	return _Storage.Contract.GetFileByName(&_Storage.CallOpts, bucketId, name)
}

// GetFileByName is a free data retrieval call binding the contract method 0x5ecdfb53.
//
// Solidity: function getFileByName(bytes32 bucketId, string name) view returns((bytes32,bytes,bytes32,string,uint256,uint256,(bytes[],uint256[])) file)
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

// CommitFile is a paid mutator transaction binding the contract method 0x87490e18.
//
// Solidity: function commitFile(bytes32 bucketId, string name, uint256 encodedFileSize, bytes fileCID) returns(bytes32)
func (_Storage *StorageTransactor) CommitFile(opts *bind.TransactOpts, bucketId [32]byte, name string, encodedFileSize *big.Int, fileCID []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "commitFile", bucketId, name, encodedFileSize, fileCID)
}

// CommitFile is a paid mutator transaction binding the contract method 0x87490e18.
//
// Solidity: function commitFile(bytes32 bucketId, string name, uint256 encodedFileSize, bytes fileCID) returns(bytes32)
func (_Storage *StorageSession) CommitFile(bucketId [32]byte, name string, encodedFileSize *big.Int, fileCID []byte) (*types.Transaction, error) {
	return _Storage.Contract.CommitFile(&_Storage.TransactOpts, bucketId, name, encodedFileSize, fileCID)
}

// CommitFile is a paid mutator transaction binding the contract method 0x87490e18.
//
// Solidity: function commitFile(bytes32 bucketId, string name, uint256 encodedFileSize, bytes fileCID) returns(bytes32)
func (_Storage *StorageTransactorSession) CommitFile(bucketId [32]byte, name string, encodedFileSize *big.Int, fileCID []byte) (*types.Transaction, error) {
	return _Storage.Contract.CommitFile(&_Storage.TransactOpts, bucketId, name, encodedFileSize, fileCID)
}

// CommitFileNew is a paid mutator transaction binding the contract method 0x25ccd451.
//
// Solidity: function commitFileNew(bytes32 bucketId, string name, uint256 encodedFileSize, bytes fileCID) returns(bytes32)
func (_Storage *StorageTransactor) CommitFileNew(opts *bind.TransactOpts, bucketId [32]byte, name string, encodedFileSize *big.Int, fileCID []byte) (*types.Transaction, error) {
	return _Storage.contract.Transact(opts, "commitFileNew", bucketId, name, encodedFileSize, fileCID)
}

// CommitFileNew is a paid mutator transaction binding the contract method 0x25ccd451.
//
// Solidity: function commitFileNew(bytes32 bucketId, string name, uint256 encodedFileSize, bytes fileCID) returns(bytes32)
func (_Storage *StorageSession) CommitFileNew(bucketId [32]byte, name string, encodedFileSize *big.Int, fileCID []byte) (*types.Transaction, error) {
	return _Storage.Contract.CommitFileNew(&_Storage.TransactOpts, bucketId, name, encodedFileSize, fileCID)
}

// CommitFileNew is a paid mutator transaction binding the contract method 0x25ccd451.
//
// Solidity: function commitFileNew(bytes32 bucketId, string name, uint256 encodedFileSize, bytes fileCID) returns(bytes32)
func (_Storage *StorageTransactorSession) CommitFileNew(bucketId [32]byte, name string, encodedFileSize *big.Int, fileCID []byte) (*types.Transaction, error) {
	return _Storage.Contract.CommitFileNew(&_Storage.TransactOpts, bucketId, name, encodedFileSize, fileCID)
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
