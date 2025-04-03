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

// PolicyFactoryMetaData contains all meta data concerning the PolicyFactory contract.
var PolicyFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_basePolicyImplementation\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"policyInstance\",\"type\":\"address\"}],\"name\":\"PolicyDeployed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"basePolicyImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"}],\"name\":\"deployPolicy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"policyInstance\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a0604052348015600e575f5ffd5b506040516103d83803806103d8833981016040819052602b91603b565b6001600160a01b03166080526066565b5f60208284031215604a575f5ffd5b81516001600160a01b0381168114605f575f5ffd5b9392505050565b6080516103556100835f395f8181603d0152608f01526103555ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c8063200afae814610038578063b8dc780f1461007b575b5f5ffd5b61005f7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200160405180910390f35b61005f610089366004610256565b5f6100b37f000000000000000000000000000000000000000000000000000000000000000061019d565b90505f816001600160a01b0316836040516100ce9190610309565b5f604051808303815f865af19150503d805f8114610107576040519150601f19603f3d011682016040523d82523d5f602084013e61010c565b606091505b50509050806101625760405162461bcd60e51b815260206004820152601c60248201527f506f6c69637920696e697469616c697a6174696f6e206661696c65640000000060448201526064015b60405180910390fd5b6040516001600160a01b0383169033907f87ba47a73518e5c03313f0d265288539fb71194e940ca6698184d22ae045ef95905f90a350919050565b5f6101a8825f6101ae565b92915050565b5f814710156101d95760405163cf47918160e01b815247600482015260248101839052604401610159565b763d602d80600a3d3981f3363d3d373d3d3d363d730000008360601b60e81c175f526e5af43d82803e903d91602b57fd5bf38360781b176020526037600983f090506001600160a01b0381166101a85760405163b06ebf3d60e01b815260040160405180910390fd5b634e487b7160e01b5f52604160045260245ffd5b5f60208284031215610266575f5ffd5b813567ffffffffffffffff81111561027c575f5ffd5b8201601f8101841361028c575f5ffd5b803567ffffffffffffffff8111156102a6576102a6610242565b604051601f8201601f19908116603f0116810167ffffffffffffffff811182821017156102d5576102d5610242565b6040528181528282016020018610156102ec575f5ffd5b816020840160208301375f91810160200191909152949350505050565b5f82518060208501845e5f92019182525091905056fea2646970667358221220b923c782f24902acfa267b9c35fce6555b73cbd5d7ca06e05f4aaad439029f9164736f6c634300081c0033",
}

// PolicyFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use PolicyFactoryMetaData.ABI instead.
var PolicyFactoryABI = PolicyFactoryMetaData.ABI

// PolicyFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PolicyFactoryMetaData.Bin instead.
var PolicyFactoryBin = PolicyFactoryMetaData.Bin

// DeployPolicyFactory deploys a new Ethereum contract, binding an instance of PolicyFactory to it.
func DeployPolicyFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _basePolicyImplementation common.Address) (common.Address, *types.Transaction, *PolicyFactory, error) {
	parsed, err := PolicyFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PolicyFactoryBin), backend, _basePolicyImplementation)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PolicyFactory{PolicyFactoryCaller: PolicyFactoryCaller{contract: contract}, PolicyFactoryTransactor: PolicyFactoryTransactor{contract: contract}, PolicyFactoryFilterer: PolicyFactoryFilterer{contract: contract}}, nil
}

// PolicyFactory is an auto generated Go binding around an Ethereum contract.
type PolicyFactory struct {
	PolicyFactoryCaller     // Read-only binding to the contract
	PolicyFactoryTransactor // Write-only binding to the contract
	PolicyFactoryFilterer   // Log filterer for contract events
}

// PolicyFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PolicyFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicyFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PolicyFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicyFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PolicyFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PolicyFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PolicyFactorySession struct {
	Contract     *PolicyFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PolicyFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PolicyFactoryCallerSession struct {
	Contract *PolicyFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PolicyFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PolicyFactoryTransactorSession struct {
	Contract     *PolicyFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PolicyFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PolicyFactoryRaw struct {
	Contract *PolicyFactory // Generic contract binding to access the raw methods on
}

// PolicyFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PolicyFactoryCallerRaw struct {
	Contract *PolicyFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// PolicyFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PolicyFactoryTransactorRaw struct {
	Contract *PolicyFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPolicyFactory creates a new instance of PolicyFactory, bound to a specific deployed contract.
func NewPolicyFactory(address common.Address, backend bind.ContractBackend) (*PolicyFactory, error) {
	contract, err := bindPolicyFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PolicyFactory{PolicyFactoryCaller: PolicyFactoryCaller{contract: contract}, PolicyFactoryTransactor: PolicyFactoryTransactor{contract: contract}, PolicyFactoryFilterer: PolicyFactoryFilterer{contract: contract}}, nil
}

// NewPolicyFactoryCaller creates a new read-only instance of PolicyFactory, bound to a specific deployed contract.
func NewPolicyFactoryCaller(address common.Address, caller bind.ContractCaller) (*PolicyFactoryCaller, error) {
	contract, err := bindPolicyFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PolicyFactoryCaller{contract: contract}, nil
}

// NewPolicyFactoryTransactor creates a new write-only instance of PolicyFactory, bound to a specific deployed contract.
func NewPolicyFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PolicyFactoryTransactor, error) {
	contract, err := bindPolicyFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PolicyFactoryTransactor{contract: contract}, nil
}

// NewPolicyFactoryFilterer creates a new log filterer instance of PolicyFactory, bound to a specific deployed contract.
func NewPolicyFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PolicyFactoryFilterer, error) {
	contract, err := bindPolicyFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PolicyFactoryFilterer{contract: contract}, nil
}

// bindPolicyFactory binds a generic wrapper to an already deployed contract.
func bindPolicyFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PolicyFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolicyFactory *PolicyFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolicyFactory.Contract.PolicyFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolicyFactory *PolicyFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolicyFactory.Contract.PolicyFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolicyFactory *PolicyFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolicyFactory.Contract.PolicyFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PolicyFactory *PolicyFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PolicyFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PolicyFactory *PolicyFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PolicyFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PolicyFactory *PolicyFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PolicyFactory.Contract.contract.Transact(opts, method, params...)
}

// BasePolicyImplementation is a free data retrieval call binding the contract method 0x200afae8.
//
// Solidity: function basePolicyImplementation() view returns(address)
func (_PolicyFactory *PolicyFactoryCaller) BasePolicyImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PolicyFactory.contract.Call(opts, &out, "basePolicyImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BasePolicyImplementation is a free data retrieval call binding the contract method 0x200afae8.
//
// Solidity: function basePolicyImplementation() view returns(address)
func (_PolicyFactory *PolicyFactorySession) BasePolicyImplementation() (common.Address, error) {
	return _PolicyFactory.Contract.BasePolicyImplementation(&_PolicyFactory.CallOpts)
}

// BasePolicyImplementation is a free data retrieval call binding the contract method 0x200afae8.
//
// Solidity: function basePolicyImplementation() view returns(address)
func (_PolicyFactory *PolicyFactoryCallerSession) BasePolicyImplementation() (common.Address, error) {
	return _PolicyFactory.Contract.BasePolicyImplementation(&_PolicyFactory.CallOpts)
}

// DeployPolicy is a paid mutator transaction binding the contract method 0xb8dc780f.
//
// Solidity: function deployPolicy(bytes initData) returns(address policyInstance)
func (_PolicyFactory *PolicyFactoryTransactor) DeployPolicy(opts *bind.TransactOpts, initData []byte) (*types.Transaction, error) {
	return _PolicyFactory.contract.Transact(opts, "deployPolicy", initData)
}

// DeployPolicy is a paid mutator transaction binding the contract method 0xb8dc780f.
//
// Solidity: function deployPolicy(bytes initData) returns(address policyInstance)
func (_PolicyFactory *PolicyFactorySession) DeployPolicy(initData []byte) (*types.Transaction, error) {
	return _PolicyFactory.Contract.DeployPolicy(&_PolicyFactory.TransactOpts, initData)
}

// DeployPolicy is a paid mutator transaction binding the contract method 0xb8dc780f.
//
// Solidity: function deployPolicy(bytes initData) returns(address policyInstance)
func (_PolicyFactory *PolicyFactoryTransactorSession) DeployPolicy(initData []byte) (*types.Transaction, error) {
	return _PolicyFactory.Contract.DeployPolicy(&_PolicyFactory.TransactOpts, initData)
}

// PolicyFactoryPolicyDeployedIterator is returned from FilterPolicyDeployed and is used to iterate over the raw logs and unpacked data for PolicyDeployed events raised by the PolicyFactory contract.
type PolicyFactoryPolicyDeployedIterator struct {
	Event *PolicyFactoryPolicyDeployed // Event containing the contract specifics and raw log

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
func (it *PolicyFactoryPolicyDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PolicyFactoryPolicyDeployed)
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
		it.Event = new(PolicyFactoryPolicyDeployed)
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
func (it *PolicyFactoryPolicyDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PolicyFactoryPolicyDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PolicyFactoryPolicyDeployed represents a PolicyDeployed event raised by the PolicyFactory contract.
type PolicyFactoryPolicyDeployed struct {
	Owner          common.Address
	PolicyInstance common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPolicyDeployed is a free log retrieval operation binding the contract event 0x87ba47a73518e5c03313f0d265288539fb71194e940ca6698184d22ae045ef95.
//
// Solidity: event PolicyDeployed(address indexed owner, address indexed policyInstance)
func (_PolicyFactory *PolicyFactoryFilterer) FilterPolicyDeployed(opts *bind.FilterOpts, owner []common.Address, policyInstance []common.Address) (*PolicyFactoryPolicyDeployedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var policyInstanceRule []interface{}
	for _, policyInstanceItem := range policyInstance {
		policyInstanceRule = append(policyInstanceRule, policyInstanceItem)
	}

	logs, sub, err := _PolicyFactory.contract.FilterLogs(opts, "PolicyDeployed", ownerRule, policyInstanceRule)
	if err != nil {
		return nil, err
	}
	return &PolicyFactoryPolicyDeployedIterator{contract: _PolicyFactory.contract, event: "PolicyDeployed", logs: logs, sub: sub}, nil
}

// WatchPolicyDeployed is a free log subscription operation binding the contract event 0x87ba47a73518e5c03313f0d265288539fb71194e940ca6698184d22ae045ef95.
//
// Solidity: event PolicyDeployed(address indexed owner, address indexed policyInstance)
func (_PolicyFactory *PolicyFactoryFilterer) WatchPolicyDeployed(opts *bind.WatchOpts, sink chan<- *PolicyFactoryPolicyDeployed, owner []common.Address, policyInstance []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var policyInstanceRule []interface{}
	for _, policyInstanceItem := range policyInstance {
		policyInstanceRule = append(policyInstanceRule, policyInstanceItem)
	}

	logs, sub, err := _PolicyFactory.contract.WatchLogs(opts, "PolicyDeployed", ownerRule, policyInstanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PolicyFactoryPolicyDeployed)
				if err := _PolicyFactory.contract.UnpackLog(event, "PolicyDeployed", log); err != nil {
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

// ParsePolicyDeployed is a log parse operation binding the contract event 0x87ba47a73518e5c03313f0d265288539fb71194e940ca6698184d22ae045ef95.
//
// Solidity: event PolicyDeployed(address indexed owner, address indexed policyInstance)
func (_PolicyFactory *PolicyFactoryFilterer) ParsePolicyDeployed(log types.Log) (*PolicyFactoryPolicyDeployed, error) {
	event := new(PolicyFactoryPolicyDeployed)
	if err := _PolicyFactory.contract.UnpackLog(event, "PolicyDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
