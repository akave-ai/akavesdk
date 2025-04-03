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

// ListPolicyMetaData contains all meta data concerning the ListPolicy contract.
var ListPolicyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotThePolicyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotWhitelisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"assignRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"validateAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"hasAccess\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506103a78061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c80635c110a741461005957806365e88c5a1461008157806380e52e3f146100965780638da5cb5b146100a9578063c4d66de8146100d3575b5f5ffd5b61006c6100673660046102d3565b6100e6565b60405190151581526020015b60405180910390f35b61009461008f366004610351565b61012e565b005b6100946100a4366004610351565b6101bd565b5f546100bb906001600160a01b031681565b6040516001600160a01b039091168152602001610078565b6100946100e1366004610351565b610245565b5f6001600160a01b03841661010e5760405163e6c4247b60e01b815260040160405180910390fd5b5050506001600160a01b03165f9081526001602052604090205460ff1690565b5f546001600160a01b03163314610158576040516351604ff560e11b815260040160405180910390fd5b6001600160a01b0381165f9081526001602081905260409091205460ff16151590036101975760405163b73e95e160e01b815260040160405180910390fd5b6001600160a01b03165f908152600160208190526040909120805460ff19169091179055565b5f546001600160a01b031633146101e7576040516351604ff560e11b815260040160405180910390fd5b6001600160a01b0381165f9081526001602081905260409091205460ff1615151461022557604051630b094f2760e31b815260040160405180910390fd5b6001600160a01b03165f908152600160205260409020805460ff19169055565b5f546001600160a01b0316156102975760405162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b604482015260640160405180910390fd5b5f80546001600160a01b0319166001600160a01b0392909216919091179055565b80356001600160a01b03811681146102ce575f5ffd5b919050565b5f5f5f604084860312156102e5575f5ffd5b6102ee846102b8565b9250602084013567ffffffffffffffff811115610309575f5ffd5b8401601f81018613610319575f5ffd5b803567ffffffffffffffff81111561032f575f5ffd5b866020828401011115610340575f5ffd5b939660209190910195509293505050565b5f60208284031215610361575f5ffd5b61036a826102b8565b939250505056fea264697066735822122047472e4c391eccdb2c4a9389b8901ff8ee4eac8e7f69a48928e3085c565ad9aa64736f6c634300081c0033",
}

// ListPolicyABI is the input ABI used to generate the binding from.
// Deprecated: Use ListPolicyMetaData.ABI instead.
var ListPolicyABI = ListPolicyMetaData.ABI

// ListPolicyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ListPolicyMetaData.Bin instead.
var ListPolicyBin = ListPolicyMetaData.Bin

// DeployListPolicy deploys a new Ethereum contract, binding an instance of ListPolicy to it.
func DeployListPolicy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ListPolicy, error) {
	parsed, err := ListPolicyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ListPolicyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ListPolicy{ListPolicyCaller: ListPolicyCaller{contract: contract}, ListPolicyTransactor: ListPolicyTransactor{contract: contract}, ListPolicyFilterer: ListPolicyFilterer{contract: contract}}, nil
}

// ListPolicy is an auto generated Go binding around an Ethereum contract.
type ListPolicy struct {
	ListPolicyCaller     // Read-only binding to the contract
	ListPolicyTransactor // Write-only binding to the contract
	ListPolicyFilterer   // Log filterer for contract events
}

// ListPolicyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ListPolicyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ListPolicyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ListPolicyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ListPolicyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ListPolicyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ListPolicySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ListPolicySession struct {
	Contract     *ListPolicy       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ListPolicyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ListPolicyCallerSession struct {
	Contract *ListPolicyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ListPolicyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ListPolicyTransactorSession struct {
	Contract     *ListPolicyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ListPolicyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ListPolicyRaw struct {
	Contract *ListPolicy // Generic contract binding to access the raw methods on
}

// ListPolicyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ListPolicyCallerRaw struct {
	Contract *ListPolicyCaller // Generic read-only contract binding to access the raw methods on
}

// ListPolicyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ListPolicyTransactorRaw struct {
	Contract *ListPolicyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewListPolicy creates a new instance of ListPolicy, bound to a specific deployed contract.
func NewListPolicy(address common.Address, backend bind.ContractBackend) (*ListPolicy, error) {
	contract, err := bindListPolicy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ListPolicy{ListPolicyCaller: ListPolicyCaller{contract: contract}, ListPolicyTransactor: ListPolicyTransactor{contract: contract}, ListPolicyFilterer: ListPolicyFilterer{contract: contract}}, nil
}

// NewListPolicyCaller creates a new read-only instance of ListPolicy, bound to a specific deployed contract.
func NewListPolicyCaller(address common.Address, caller bind.ContractCaller) (*ListPolicyCaller, error) {
	contract, err := bindListPolicy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ListPolicyCaller{contract: contract}, nil
}

// NewListPolicyTransactor creates a new write-only instance of ListPolicy, bound to a specific deployed contract.
func NewListPolicyTransactor(address common.Address, transactor bind.ContractTransactor) (*ListPolicyTransactor, error) {
	contract, err := bindListPolicy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ListPolicyTransactor{contract: contract}, nil
}

// NewListPolicyFilterer creates a new log filterer instance of ListPolicy, bound to a specific deployed contract.
func NewListPolicyFilterer(address common.Address, filterer bind.ContractFilterer) (*ListPolicyFilterer, error) {
	contract, err := bindListPolicy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ListPolicyFilterer{contract: contract}, nil
}

// bindListPolicy binds a generic wrapper to an already deployed contract.
func bindListPolicy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ListPolicyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ListPolicy *ListPolicyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ListPolicy.Contract.ListPolicyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ListPolicy *ListPolicyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ListPolicy.Contract.ListPolicyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ListPolicy *ListPolicyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ListPolicy.Contract.ListPolicyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ListPolicy *ListPolicyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ListPolicy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ListPolicy *ListPolicyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ListPolicy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ListPolicy *ListPolicyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ListPolicy.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ListPolicy *ListPolicyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ListPolicy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ListPolicy *ListPolicySession) Owner() (common.Address, error) {
	return _ListPolicy.Contract.Owner(&_ListPolicy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ListPolicy *ListPolicyCallerSession) Owner() (common.Address, error) {
	return _ListPolicy.Contract.Owner(&_ListPolicy.CallOpts)
}

// ValidateAccess is a free data retrieval call binding the contract method 0x5c110a74.
//
// Solidity: function validateAccess(address user, bytes data) view returns(bool hasAccess)
func (_ListPolicy *ListPolicyCaller) ValidateAccess(opts *bind.CallOpts, user common.Address, data []byte) (bool, error) {
	var out []interface{}
	err := _ListPolicy.contract.Call(opts, &out, "validateAccess", user, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateAccess is a free data retrieval call binding the contract method 0x5c110a74.
//
// Solidity: function validateAccess(address user, bytes data) view returns(bool hasAccess)
func (_ListPolicy *ListPolicySession) ValidateAccess(user common.Address, data []byte) (bool, error) {
	return _ListPolicy.Contract.ValidateAccess(&_ListPolicy.CallOpts, user, data)
}

// ValidateAccess is a free data retrieval call binding the contract method 0x5c110a74.
//
// Solidity: function validateAccess(address user, bytes data) view returns(bool hasAccess)
func (_ListPolicy *ListPolicyCallerSession) ValidateAccess(user common.Address, data []byte) (bool, error) {
	return _ListPolicy.Contract.ValidateAccess(&_ListPolicy.CallOpts, user, data)
}

// AssignRole is a paid mutator transaction binding the contract method 0x65e88c5a.
//
// Solidity: function assignRole(address user) returns()
func (_ListPolicy *ListPolicyTransactor) AssignRole(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _ListPolicy.contract.Transact(opts, "assignRole", user)
}

// AssignRole is a paid mutator transaction binding the contract method 0x65e88c5a.
//
// Solidity: function assignRole(address user) returns()
func (_ListPolicy *ListPolicySession) AssignRole(user common.Address) (*types.Transaction, error) {
	return _ListPolicy.Contract.AssignRole(&_ListPolicy.TransactOpts, user)
}

// AssignRole is a paid mutator transaction binding the contract method 0x65e88c5a.
//
// Solidity: function assignRole(address user) returns()
func (_ListPolicy *ListPolicyTransactorSession) AssignRole(user common.Address) (*types.Transaction, error) {
	return _ListPolicy.Contract.AssignRole(&_ListPolicy.TransactOpts, user)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_ListPolicy *ListPolicyTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _ListPolicy.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_ListPolicy *ListPolicySession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _ListPolicy.Contract.Initialize(&_ListPolicy.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_ListPolicy *ListPolicyTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _ListPolicy.Contract.Initialize(&_ListPolicy.TransactOpts, _owner)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x80e52e3f.
//
// Solidity: function revokeRole(address user) returns()
func (_ListPolicy *ListPolicyTransactor) RevokeRole(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _ListPolicy.contract.Transact(opts, "revokeRole", user)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x80e52e3f.
//
// Solidity: function revokeRole(address user) returns()
func (_ListPolicy *ListPolicySession) RevokeRole(user common.Address) (*types.Transaction, error) {
	return _ListPolicy.Contract.RevokeRole(&_ListPolicy.TransactOpts, user)
}

// RevokeRole is a paid mutator transaction binding the contract method 0x80e52e3f.
//
// Solidity: function revokeRole(address user) returns()
func (_ListPolicy *ListPolicyTransactorSession) RevokeRole(user common.Address) (*types.Transaction, error) {
	return _ListPolicy.Contract.RevokeRole(&_ListPolicy.TransactOpts, user)
}
