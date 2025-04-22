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

// AccessManagerMetaData contains all meta data concerning the AccessManager contract.
var AccessManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"storageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BucketNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPolicy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBucketOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"policyContract\",\"type\":\"address\"}],\"name\":\"PolicyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"name\":\"PublicAccessChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"name\":\"changePublicAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"getFileAccessInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"getPolicy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getValidateAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"policyContract\",\"type\":\"address\"}],\"name\":\"setBucketPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"policyContract\",\"type\":\"address\"}],\"name\":\"setFilePolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageContract\",\"outputs\":[{\"internalType\":\"contractIStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50604051610ee6380380610ee6833981016040819052602b91604e565b5f80546001600160a01b0319166001600160a01b03929092169190911790556079565b5f60208284031215605d575f5ffd5b81516001600160a01b03811681146072575f5ffd5b9392505050565b610e60806100865f395ff3fe608060405234801561000f575f5ffd5b506004361061007a575f3560e01c80634407880211610058578063440788021461011d578063a3f685f914610140578063c05d0d6c14610153578063e13212f814610166575f5ffd5b806311ce02671461007e57806338f5f5b2146100ad5780633a82c912146100c2575b5f5ffd5b5f54610090906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100c06100bb36600461074d565b610179565b005b6100fe6100d036600461077b565b5f908152600160209081526040808320546003909252909120546001600160a01b039091169160ff90911690565b604080516001600160a01b0390931683529015156020830152016100a4565b61013061012b3660046107a6565b610273565b60405190151581526020016100a4565b61009061014e36600461077b565b610399565b6100c0610161366004610829565b610496565b6100c0610174366004610829565b6104f9565b5f8054604051637d7602a160e11b8152600481018590528492916001600160a01b03169063faec0542906024015f60405180830381865afa1580156101c0573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526101e79190810190610add565b60408101519091505f0361020e57604051631de5d86f60e11b815260040160405180910390fd5b61021b81604001516105f5565b5f84815260036020908152604091829020805460ff1916861515908117909155915191825285917f9500a58cfb37cef230929cd9f25ce92c41374416f23b1825232d0905a7e73d5a910160405180910390a250505050565b5f8481526003602052604081205460ff168061030157505f546040516306d1d7bd60e21b8152600481018790526001600160a01b03868116921690631b475ef490602401602060405180830381865afa1580156102d2573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102f69190610bc3565b6001600160a01b0316145b1561030e57506001610391565b5f61031886610399565b604051631704429d60e21b815290915081906001600160a01b03821690635c110a749061034d90899089908990600401610be5565b602060405180830381865afa158015610368573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061038c9190610c24565b925050505b949350505050565b5f818152600160205260408120546001600160a01b031661047b575f8054604051637d7602a160e11b8152600481018590526001600160a01b039091169063faec0542906024015f60405180830381865afa1580156103fa573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526104219190810190610add565b6040808201515f908152600260205220549091506001600160a01b031661045b5760405163cefa6b0560e01b815260040160405180910390fd5b6040908101515f908152600260205220546001600160a01b031692915050565b505f908152600160205260409020546001600160a01b031690565b816104a0816105f5565b5f8381526002602052604080822080546001600160a01b0319166001600160a01b0386169081179091559051909185917fee78f4c2af2887839fdcba441de8e4d2d1b117b89ce7cc7a7f0a952871cc87d29190a3505050565b5f8054604051637d7602a160e11b8152600481018590528492916001600160a01b03169063faec0542906024015f60405180830381865afa158015610540573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526105679190810190610add565b60408101519091505f0361058e57604051631de5d86f60e11b815260040160405180910390fd5b61059b81604001516105f5565b5f8481526001602052604080822080546001600160a01b0319166001600160a01b0387169081179091559051909186917fee78f4c2af2887839fdcba441de8e4d2d1b117b89ce7cc7a7f0a952871cc87d29190a350505050565b6040805160018082528183019092525f916020808301908036833701905050905081815f8151811061062957610629610c3f565b60209081029190910101525f80546040516335bdb71160e01b81526001600160a01b03909116906335bdb71190610664908590600401610c53565b5f60405180830381865afa15801561067e573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526106a59190810190610c95565b90505f5f1b815f815181106106bc576106bc610c3f565b60200260200101515f0151036106e55760405163c4c1a0c560e01b815260040160405180910390fd5b5f815f815181106106f8576106f8610c3f565b60200260200101519050336001600160a01b031681606001516001600160a01b0316146107375760405162d6b18f60e41b815260040160405180910390fd5b50505050565b801515811461074a575f5ffd5b50565b5f5f6040838503121561075e575f5ffd5b8235915060208301356107708161073d565b809150509250929050565b5f6020828403121561078b575f5ffd5b5035919050565b6001600160a01b038116811461074a575f5ffd5b5f5f5f5f606085870312156107b9575f5ffd5b8435935060208501356107cb81610792565b925060408501356001600160401b038111156107e5575f5ffd5b8501601f810187136107f5575f5ffd5b80356001600160401b0381111561080a575f5ffd5b87602082840101111561081b575f5ffd5b949793965060200194505050565b5f5f6040838503121561083a575f5ffd5b82359150602083013561077081610792565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b03811182821017156108825761088261084c565b60405290565b60405160e081016001600160401b03811182821017156108825761088261084c565b60405160a081016001600160401b03811182821017156108825761088261084c565b604051601f8201601f191681016001600160401b03811182821017156108f4576108f461084c565b604052919050565b5f82601f83011261090b575f5ffd5b8151602083015f5f6001600160401b0384111561092a5761092a61084c565b50601f8301601f191660200161093f816108cc565b915050828152858383011115610953575f5ffd5b8282602083015e5f92810160200192909252509392505050565b5f6001600160401b038211156109855761098561084c565b5060051b60200190565b5f82601f83011261099e575f5ffd5b81516109b16109ac8261096d565b6108cc565b8082825260208201915060208360051b8601019250858311156109d2575f5ffd5b602085015b838110156109ef5780518352602092830192016109d7565b5095945050505050565b5f60408284031215610a09575f5ffd5b610a11610860565b905081516001600160401b03811115610a28575f5ffd5b8201601f81018413610a38575f5ffd5b8051610a466109ac8261096d565b8082825260208201915060208360051b850101925086831115610a67575f5ffd5b602084015b83811015610aa75780516001600160401b03811115610a89575f5ffd5b610a98896020838901016108fc565b84525060209283019201610a6c565b50845250505060208201516001600160401b03811115610ac5575f5ffd5b610ad18482850161098f565b60208301525092915050565b5f60208284031215610aed575f5ffd5b81516001600160401b03811115610b02575f5ffd5b820160e08185031215610b13575f5ffd5b610b1b610888565b8151815260208201516001600160401b03811115610b37575f5ffd5b610b43868285016108fc565b6020830152506040828101519082015260608201516001600160401b03811115610b6b575f5ffd5b610b77868285016108fc565b6060830152506080828101519082015260a0808301519082015260c08201516001600160401b03811115610ba9575f5ffd5b610bb5868285016109f9565b60c083015250949350505050565b5f60208284031215610bd3575f5ffd5b8151610bde81610792565b9392505050565b6001600160a01b03841681526040602082018190528101829052818360608301375f818301606090810191909152601f909201601f1916010192915050565b5f60208284031215610c34575f5ffd5b8151610bde8161073d565b634e487b7160e01b5f52603260045260245ffd5b602080825282518282018190525f918401906040840190835b81811015610c8a578351835260209384019390920191600101610c6c565b509095945050505050565b5f60208284031215610ca5575f5ffd5b81516001600160401b03811115610cba575f5ffd5b8201601f81018413610cca575f5ffd5b8051610cd86109ac8261096d565b8082825260208201915060208360051b850101925086831115610cf9575f5ffd5b602084015b83811015610e1f5780516001600160401b03811115610d1b575f5ffd5b850160a0818a03601f19011215610d30575f5ffd5b610d386108aa565b6020820151815260408201516001600160401b03811115610d57575f5ffd5b610d668b6020838601016108fc565b602083015250606082015160408201526080820151610d8481610792565b606082015260a08201516001600160401b03811115610da1575f5ffd5b60208184010192505089601f830112610db8575f5ffd5b8151610dc66109ac8261096d565b8082825260208201915060208360051b86010192508c831115610de7575f5ffd5b6020850194505b82851015610e09578451825260209485019490910190610dee565b6080840152505084525060209283019201610cfe565b50969550505050505056fea2646970667358221220ee4a73415c45aec05a76663df177fb6dd1e748206c9f12810802cb11527189bc64736f6c634300081c0033",
}

// AccessManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessManagerMetaData.ABI instead.
var AccessManagerABI = AccessManagerMetaData.ABI

// AccessManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccessManagerMetaData.Bin instead.
var AccessManagerBin = AccessManagerMetaData.Bin

// DeployAccessManager deploys a new Ethereum contract, binding an instance of AccessManager to it.
func DeployAccessManager(auth *bind.TransactOpts, backend bind.ContractBackend, storageAddress common.Address) (common.Address, *types.Transaction, *AccessManager, error) {
	parsed, err := AccessManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccessManagerBin), backend, storageAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccessManager{AccessManagerCaller: AccessManagerCaller{contract: contract}, AccessManagerTransactor: AccessManagerTransactor{contract: contract}, AccessManagerFilterer: AccessManagerFilterer{contract: contract}}, nil
}

// AccessManager is an auto generated Go binding around an Ethereum contract.
type AccessManager struct {
	AccessManagerCaller     // Read-only binding to the contract
	AccessManagerTransactor // Write-only binding to the contract
	AccessManagerFilterer   // Log filterer for contract events
}

// AccessManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessManagerSession struct {
	Contract     *AccessManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessManagerCallerSession struct {
	Contract *AccessManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AccessManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessManagerTransactorSession struct {
	Contract     *AccessManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccessManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessManagerRaw struct {
	Contract *AccessManager // Generic contract binding to access the raw methods on
}

// AccessManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessManagerCallerRaw struct {
	Contract *AccessManagerCaller // Generic read-only contract binding to access the raw methods on
}

// AccessManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessManagerTransactorRaw struct {
	Contract *AccessManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessManager creates a new instance of AccessManager, bound to a specific deployed contract.
func NewAccessManager(address common.Address, backend bind.ContractBackend) (*AccessManager, error) {
	contract, err := bindAccessManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessManager{AccessManagerCaller: AccessManagerCaller{contract: contract}, AccessManagerTransactor: AccessManagerTransactor{contract: contract}, AccessManagerFilterer: AccessManagerFilterer{contract: contract}}, nil
}

// NewAccessManagerCaller creates a new read-only instance of AccessManager, bound to a specific deployed contract.
func NewAccessManagerCaller(address common.Address, caller bind.ContractCaller) (*AccessManagerCaller, error) {
	contract, err := bindAccessManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessManagerCaller{contract: contract}, nil
}

// NewAccessManagerTransactor creates a new write-only instance of AccessManager, bound to a specific deployed contract.
func NewAccessManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessManagerTransactor, error) {
	contract, err := bindAccessManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessManagerTransactor{contract: contract}, nil
}

// NewAccessManagerFilterer creates a new log filterer instance of AccessManager, bound to a specific deployed contract.
func NewAccessManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessManagerFilterer, error) {
	contract, err := bindAccessManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessManagerFilterer{contract: contract}, nil
}

// bindAccessManager binds a generic wrapper to an already deployed contract.
func bindAccessManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccessManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessManager *AccessManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessManager.Contract.AccessManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessManager *AccessManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessManager.Contract.AccessManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessManager *AccessManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessManager.Contract.AccessManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessManager *AccessManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessManager *AccessManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessManager *AccessManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessManager.Contract.contract.Transact(opts, method, params...)
}

// GetFileAccessInfo is a free data retrieval call binding the contract method 0x3a82c912.
//
// Solidity: function getFileAccessInfo(bytes32 fileId) view returns(address, bool)
func (_AccessManager *AccessManagerCaller) GetFileAccessInfo(opts *bind.CallOpts, fileId [32]byte) (common.Address, bool, error) {
	var out []interface{}
	err := _AccessManager.contract.Call(opts, &out, "getFileAccessInfo", fileId)

	if err != nil {
		return *new(common.Address), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetFileAccessInfo is a free data retrieval call binding the contract method 0x3a82c912.
//
// Solidity: function getFileAccessInfo(bytes32 fileId) view returns(address, bool)
func (_AccessManager *AccessManagerSession) GetFileAccessInfo(fileId [32]byte) (common.Address, bool, error) {
	return _AccessManager.Contract.GetFileAccessInfo(&_AccessManager.CallOpts, fileId)
}

// GetFileAccessInfo is a free data retrieval call binding the contract method 0x3a82c912.
//
// Solidity: function getFileAccessInfo(bytes32 fileId) view returns(address, bool)
func (_AccessManager *AccessManagerCallerSession) GetFileAccessInfo(fileId [32]byte) (common.Address, bool, error) {
	return _AccessManager.Contract.GetFileAccessInfo(&_AccessManager.CallOpts, fileId)
}

// GetPolicy is a free data retrieval call binding the contract method 0xa3f685f9.
//
// Solidity: function getPolicy(bytes32 fileId) view returns(address)
func (_AccessManager *AccessManagerCaller) GetPolicy(opts *bind.CallOpts, fileId [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AccessManager.contract.Call(opts, &out, "getPolicy", fileId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPolicy is a free data retrieval call binding the contract method 0xa3f685f9.
//
// Solidity: function getPolicy(bytes32 fileId) view returns(address)
func (_AccessManager *AccessManagerSession) GetPolicy(fileId [32]byte) (common.Address, error) {
	return _AccessManager.Contract.GetPolicy(&_AccessManager.CallOpts, fileId)
}

// GetPolicy is a free data retrieval call binding the contract method 0xa3f685f9.
//
// Solidity: function getPolicy(bytes32 fileId) view returns(address)
func (_AccessManager *AccessManagerCallerSession) GetPolicy(fileId [32]byte) (common.Address, error) {
	return _AccessManager.Contract.GetPolicy(&_AccessManager.CallOpts, fileId)
}

// GetValidateAccess is a free data retrieval call binding the contract method 0x44078802.
//
// Solidity: function getValidateAccess(bytes32 fileId, address user, bytes data) view returns(bool)
func (_AccessManager *AccessManagerCaller) GetValidateAccess(opts *bind.CallOpts, fileId [32]byte, user common.Address, data []byte) (bool, error) {
	var out []interface{}
	err := _AccessManager.contract.Call(opts, &out, "getValidateAccess", fileId, user, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetValidateAccess is a free data retrieval call binding the contract method 0x44078802.
//
// Solidity: function getValidateAccess(bytes32 fileId, address user, bytes data) view returns(bool)
func (_AccessManager *AccessManagerSession) GetValidateAccess(fileId [32]byte, user common.Address, data []byte) (bool, error) {
	return _AccessManager.Contract.GetValidateAccess(&_AccessManager.CallOpts, fileId, user, data)
}

// GetValidateAccess is a free data retrieval call binding the contract method 0x44078802.
//
// Solidity: function getValidateAccess(bytes32 fileId, address user, bytes data) view returns(bool)
func (_AccessManager *AccessManagerCallerSession) GetValidateAccess(fileId [32]byte, user common.Address, data []byte) (bool, error) {
	return _AccessManager.Contract.GetValidateAccess(&_AccessManager.CallOpts, fileId, user, data)
}

// StorageContract is a free data retrieval call binding the contract method 0x11ce0267.
//
// Solidity: function storageContract() view returns(address)
func (_AccessManager *AccessManagerCaller) StorageContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AccessManager.contract.Call(opts, &out, "storageContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StorageContract is a free data retrieval call binding the contract method 0x11ce0267.
//
// Solidity: function storageContract() view returns(address)
func (_AccessManager *AccessManagerSession) StorageContract() (common.Address, error) {
	return _AccessManager.Contract.StorageContract(&_AccessManager.CallOpts)
}

// StorageContract is a free data retrieval call binding the contract method 0x11ce0267.
//
// Solidity: function storageContract() view returns(address)
func (_AccessManager *AccessManagerCallerSession) StorageContract() (common.Address, error) {
	return _AccessManager.Contract.StorageContract(&_AccessManager.CallOpts)
}

// ChangePublicAccess is a paid mutator transaction binding the contract method 0x38f5f5b2.
//
// Solidity: function changePublicAccess(bytes32 fileId, bool isPublic) returns()
func (_AccessManager *AccessManagerTransactor) ChangePublicAccess(opts *bind.TransactOpts, fileId [32]byte, isPublic bool) (*types.Transaction, error) {
	return _AccessManager.contract.Transact(opts, "changePublicAccess", fileId, isPublic)
}

// ChangePublicAccess is a paid mutator transaction binding the contract method 0x38f5f5b2.
//
// Solidity: function changePublicAccess(bytes32 fileId, bool isPublic) returns()
func (_AccessManager *AccessManagerSession) ChangePublicAccess(fileId [32]byte, isPublic bool) (*types.Transaction, error) {
	return _AccessManager.Contract.ChangePublicAccess(&_AccessManager.TransactOpts, fileId, isPublic)
}

// ChangePublicAccess is a paid mutator transaction binding the contract method 0x38f5f5b2.
//
// Solidity: function changePublicAccess(bytes32 fileId, bool isPublic) returns()
func (_AccessManager *AccessManagerTransactorSession) ChangePublicAccess(fileId [32]byte, isPublic bool) (*types.Transaction, error) {
	return _AccessManager.Contract.ChangePublicAccess(&_AccessManager.TransactOpts, fileId, isPublic)
}

// SetBucketPolicy is a paid mutator transaction binding the contract method 0xc05d0d6c.
//
// Solidity: function setBucketPolicy(bytes32 bucketId, address policyContract) returns()
func (_AccessManager *AccessManagerTransactor) SetBucketPolicy(opts *bind.TransactOpts, bucketId [32]byte, policyContract common.Address) (*types.Transaction, error) {
	return _AccessManager.contract.Transact(opts, "setBucketPolicy", bucketId, policyContract)
}

// SetBucketPolicy is a paid mutator transaction binding the contract method 0xc05d0d6c.
//
// Solidity: function setBucketPolicy(bytes32 bucketId, address policyContract) returns()
func (_AccessManager *AccessManagerSession) SetBucketPolicy(bucketId [32]byte, policyContract common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.SetBucketPolicy(&_AccessManager.TransactOpts, bucketId, policyContract)
}

// SetBucketPolicy is a paid mutator transaction binding the contract method 0xc05d0d6c.
//
// Solidity: function setBucketPolicy(bytes32 bucketId, address policyContract) returns()
func (_AccessManager *AccessManagerTransactorSession) SetBucketPolicy(bucketId [32]byte, policyContract common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.SetBucketPolicy(&_AccessManager.TransactOpts, bucketId, policyContract)
}

// SetFilePolicy is a paid mutator transaction binding the contract method 0xe13212f8.
//
// Solidity: function setFilePolicy(bytes32 fileId, address policyContract) returns()
func (_AccessManager *AccessManagerTransactor) SetFilePolicy(opts *bind.TransactOpts, fileId [32]byte, policyContract common.Address) (*types.Transaction, error) {
	return _AccessManager.contract.Transact(opts, "setFilePolicy", fileId, policyContract)
}

// SetFilePolicy is a paid mutator transaction binding the contract method 0xe13212f8.
//
// Solidity: function setFilePolicy(bytes32 fileId, address policyContract) returns()
func (_AccessManager *AccessManagerSession) SetFilePolicy(fileId [32]byte, policyContract common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.SetFilePolicy(&_AccessManager.TransactOpts, fileId, policyContract)
}

// SetFilePolicy is a paid mutator transaction binding the contract method 0xe13212f8.
//
// Solidity: function setFilePolicy(bytes32 fileId, address policyContract) returns()
func (_AccessManager *AccessManagerTransactorSession) SetFilePolicy(fileId [32]byte, policyContract common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.SetFilePolicy(&_AccessManager.TransactOpts, fileId, policyContract)
}

// AccessManagerPolicyChangedIterator is returned from FilterPolicyChanged and is used to iterate over the raw logs and unpacked data for PolicyChanged events raised by the AccessManager contract.
type AccessManagerPolicyChangedIterator struct {
	Event *AccessManagerPolicyChanged // Event containing the contract specifics and raw log

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
func (it *AccessManagerPolicyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessManagerPolicyChanged)
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
		it.Event = new(AccessManagerPolicyChanged)
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
func (it *AccessManagerPolicyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessManagerPolicyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessManagerPolicyChanged represents a PolicyChanged event raised by the AccessManager contract.
type AccessManagerPolicyChanged struct {
	FileId         [32]byte
	PolicyContract common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPolicyChanged is a free log retrieval operation binding the contract event 0xee78f4c2af2887839fdcba441de8e4d2d1b117b89ce7cc7a7f0a952871cc87d2.
//
// Solidity: event PolicyChanged(bytes32 indexed fileId, address indexed policyContract)
func (_AccessManager *AccessManagerFilterer) FilterPolicyChanged(opts *bind.FilterOpts, fileId [][32]byte, policyContract []common.Address) (*AccessManagerPolicyChangedIterator, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}
	var policyContractRule []interface{}
	for _, policyContractItem := range policyContract {
		policyContractRule = append(policyContractRule, policyContractItem)
	}

	logs, sub, err := _AccessManager.contract.FilterLogs(opts, "PolicyChanged", fileIdRule, policyContractRule)
	if err != nil {
		return nil, err
	}
	return &AccessManagerPolicyChangedIterator{contract: _AccessManager.contract, event: "PolicyChanged", logs: logs, sub: sub}, nil
}

// WatchPolicyChanged is a free log subscription operation binding the contract event 0xee78f4c2af2887839fdcba441de8e4d2d1b117b89ce7cc7a7f0a952871cc87d2.
//
// Solidity: event PolicyChanged(bytes32 indexed fileId, address indexed policyContract)
func (_AccessManager *AccessManagerFilterer) WatchPolicyChanged(opts *bind.WatchOpts, sink chan<- *AccessManagerPolicyChanged, fileId [][32]byte, policyContract []common.Address) (event.Subscription, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}
	var policyContractRule []interface{}
	for _, policyContractItem := range policyContract {
		policyContractRule = append(policyContractRule, policyContractItem)
	}

	logs, sub, err := _AccessManager.contract.WatchLogs(opts, "PolicyChanged", fileIdRule, policyContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessManagerPolicyChanged)
				if err := _AccessManager.contract.UnpackLog(event, "PolicyChanged", log); err != nil {
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

// ParsePolicyChanged is a log parse operation binding the contract event 0xee78f4c2af2887839fdcba441de8e4d2d1b117b89ce7cc7a7f0a952871cc87d2.
//
// Solidity: event PolicyChanged(bytes32 indexed fileId, address indexed policyContract)
func (_AccessManager *AccessManagerFilterer) ParsePolicyChanged(log types.Log) (*AccessManagerPolicyChanged, error) {
	event := new(AccessManagerPolicyChanged)
	if err := _AccessManager.contract.UnpackLog(event, "PolicyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessManagerPublicAccessChangedIterator is returned from FilterPublicAccessChanged and is used to iterate over the raw logs and unpacked data for PublicAccessChanged events raised by the AccessManager contract.
type AccessManagerPublicAccessChangedIterator struct {
	Event *AccessManagerPublicAccessChanged // Event containing the contract specifics and raw log

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
func (it *AccessManagerPublicAccessChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessManagerPublicAccessChanged)
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
		it.Event = new(AccessManagerPublicAccessChanged)
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
func (it *AccessManagerPublicAccessChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessManagerPublicAccessChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessManagerPublicAccessChanged represents a PublicAccessChanged event raised by the AccessManager contract.
type AccessManagerPublicAccessChanged struct {
	FileId   [32]byte
	IsPublic bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPublicAccessChanged is a free log retrieval operation binding the contract event 0x9500a58cfb37cef230929cd9f25ce92c41374416f23b1825232d0905a7e73d5a.
//
// Solidity: event PublicAccessChanged(bytes32 indexed fileId, bool isPublic)
func (_AccessManager *AccessManagerFilterer) FilterPublicAccessChanged(opts *bind.FilterOpts, fileId [][32]byte) (*AccessManagerPublicAccessChangedIterator, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}

	logs, sub, err := _AccessManager.contract.FilterLogs(opts, "PublicAccessChanged", fileIdRule)
	if err != nil {
		return nil, err
	}
	return &AccessManagerPublicAccessChangedIterator{contract: _AccessManager.contract, event: "PublicAccessChanged", logs: logs, sub: sub}, nil
}

// WatchPublicAccessChanged is a free log subscription operation binding the contract event 0x9500a58cfb37cef230929cd9f25ce92c41374416f23b1825232d0905a7e73d5a.
//
// Solidity: event PublicAccessChanged(bytes32 indexed fileId, bool isPublic)
func (_AccessManager *AccessManagerFilterer) WatchPublicAccessChanged(opts *bind.WatchOpts, sink chan<- *AccessManagerPublicAccessChanged, fileId [][32]byte) (event.Subscription, error) {

	var fileIdRule []interface{}
	for _, fileIdItem := range fileId {
		fileIdRule = append(fileIdRule, fileIdItem)
	}

	logs, sub, err := _AccessManager.contract.WatchLogs(opts, "PublicAccessChanged", fileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessManagerPublicAccessChanged)
				if err := _AccessManager.contract.UnpackLog(event, "PublicAccessChanged", log); err != nil {
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

// ParsePublicAccessChanged is a log parse operation binding the contract event 0x9500a58cfb37cef230929cd9f25ce92c41374416f23b1825232d0905a7e73d5a.
//
// Solidity: event PublicAccessChanged(bytes32 indexed fileId, bool isPublic)
func (_AccessManager *AccessManagerFilterer) ParsePublicAccessChanged(log types.Log) (*AccessManagerPublicAccessChanged, error) {
	event := new(AccessManagerPublicAccessChanged)
	if err := _AccessManager.contract.UnpackLog(event, "PublicAccessChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
