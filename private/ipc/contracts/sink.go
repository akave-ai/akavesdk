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

// SinkMetaData contains all meta data concerning the Sink contract.
var SinkMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50604680601a5f395ff3fe6080604052348015600e575f5ffd5b00fea2646970667358221220f43799cb6e28e32500f5eb3784cc07778d26ab2be04f4ee9fd27d581ad2138f464736f6c634300081c0033",
}

// SinkABI is the input ABI used to generate the binding from.
// Deprecated: Use SinkMetaData.ABI instead.
var SinkABI = SinkMetaData.ABI

// SinkBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SinkMetaData.Bin instead.
var SinkBin = SinkMetaData.Bin

// DeploySink deploys a new Ethereum contract, binding an instance of Sink to it.
func DeploySink(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Sink, error) {
	parsed, err := SinkMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SinkBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Sink{SinkCaller: SinkCaller{contract: contract}, SinkTransactor: SinkTransactor{contract: contract}, SinkFilterer: SinkFilterer{contract: contract}}, nil
}

// Sink is an auto generated Go binding around an Ethereum contract.
type Sink struct {
	SinkCaller     // Read-only binding to the contract
	SinkTransactor // Write-only binding to the contract
	SinkFilterer   // Log filterer for contract events
}

// SinkCaller is an auto generated read-only Go binding around an Ethereum contract.
type SinkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SinkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SinkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SinkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SinkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SinkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SinkSession struct {
	Contract     *Sink             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SinkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SinkCallerSession struct {
	Contract *SinkCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SinkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SinkTransactorSession struct {
	Contract     *SinkTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SinkRaw is an auto generated low-level Go binding around an Ethereum contract.
type SinkRaw struct {
	Contract *Sink // Generic contract binding to access the raw methods on
}

// SinkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SinkCallerRaw struct {
	Contract *SinkCaller // Generic read-only contract binding to access the raw methods on
}

// SinkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SinkTransactorRaw struct {
	Contract *SinkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSink creates a new instance of Sink, bound to a specific deployed contract.
func NewSink(address common.Address, backend bind.ContractBackend) (*Sink, error) {
	contract, err := bindSink(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sink{SinkCaller: SinkCaller{contract: contract}, SinkTransactor: SinkTransactor{contract: contract}, SinkFilterer: SinkFilterer{contract: contract}}, nil
}

// NewSinkCaller creates a new read-only instance of Sink, bound to a specific deployed contract.
func NewSinkCaller(address common.Address, caller bind.ContractCaller) (*SinkCaller, error) {
	contract, err := bindSink(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SinkCaller{contract: contract}, nil
}

// NewSinkTransactor creates a new write-only instance of Sink, bound to a specific deployed contract.
func NewSinkTransactor(address common.Address, transactor bind.ContractTransactor) (*SinkTransactor, error) {
	contract, err := bindSink(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SinkTransactor{contract: contract}, nil
}

// NewSinkFilterer creates a new log filterer instance of Sink, bound to a specific deployed contract.
func NewSinkFilterer(address common.Address, filterer bind.ContractFilterer) (*SinkFilterer, error) {
	contract, err := bindSink(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SinkFilterer{contract: contract}, nil
}

// bindSink binds a generic wrapper to an already deployed contract.
func bindSink(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SinkMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sink *SinkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sink.Contract.SinkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sink *SinkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sink.Contract.SinkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sink *SinkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sink.Contract.SinkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sink *SinkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sink.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sink *SinkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sink.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sink *SinkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sink.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Sink *SinkTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Sink.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Sink *SinkSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Sink.Contract.Fallback(&_Sink.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_Sink *SinkTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Sink.Contract.Fallback(&_Sink.TransactOpts, calldata)
}
