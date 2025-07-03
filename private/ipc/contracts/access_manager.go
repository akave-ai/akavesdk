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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"storageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BucketNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoDelegatedAccess\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPolicy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBucketOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"AccessDelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"AccessRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"policyContract\",\"type\":\"address\"}],\"name\":\"PolicyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"name\":\"PublicAccessChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"name\":\"changePublicAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegated\",\"type\":\"address\"}],\"name\":\"delegateAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"getFileAccessInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"getPolicy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getValidateAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isBucketOwnerOrDelegate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegated\",\"type\":\"address\"}],\"name\":\"removeAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"policyContract\",\"type\":\"address\"}],\"name\":\"setBucketPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"policyContract\",\"type\":\"address\"}],\"name\":\"setFilePolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_storageAddress\",\"type\":\"address\"}],\"name\":\"setStorageContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageContract\",\"outputs\":[{\"internalType\":\"contractIStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060405161128d38038061128d833981016040819052602b91604e565b5f80546001600160a01b0319166001600160a01b03929092169190911790556079565b5f60208284031215605d575f5ffd5b81516001600160a01b03811681146072575f5ffd5b9392505050565b611207806100865f395ff3fe608060405234801561000f575f5ffd5b50600436106100a6575f3560e01c8063a3f685f91161006e578063a3f685f91461017f578063a9b61d4d14610192578063c05d0d6c146101a5578063dc38b0a2146101b8578063e124bdd9146101e7578063e13212f8146101fa575f5ffd5b806311ce0267146100aa57806338f5f5b2146100d95780633a82c912146100ee57806344078802146101495780635e2a22de1461016c575b5f5ffd5b5f546100bc906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100ec6100e7366004610acd565b61020d565b005b61012a6100fc366004610afb565b5f908152600160209081526040808320546003909252909120546001600160a01b039091169160ff90911690565b604080516001600160a01b0390931683529015156020830152016100d0565b61015c610157366004610b26565b610307565b60405190151581526020016100d0565b6100ec61017a366004610ba9565b61042d565b6100bc61018d366004610afb565b6104fc565b6100ec6101a0366004610ba9565b6105f9565b6100ec6101b3366004610ba9565b610685565b6100ec6101c6366004610bcc565b5f80546001600160a01b0319166001600160a01b0392909216919091179055565b61015c6101f5366004610ba9565b6106e8565b6100ec610208366004610ba9565b610846565b5f8054604051637d7602a160e11b8152600481018590528492916001600160a01b03169063faec0542906024015f60405180830381865afa158015610254573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261027b9190810190610e80565b60408101519091505f036102a257604051631de5d86f60e11b815260040160405180910390fd5b6102af8160400151610942565b5f84815260036020908152604091829020805460ff1916861515908117909155915191825285917f9500a58cfb37cef230929cd9f25ce92c41374416f23b1825232d0905a7e73d5a910160405180910390a250505050565b5f8481526003602052604081205460ff168061039557505f546040516306d1d7bd60e21b8152600481018790526001600160a01b03868116921690631b475ef490602401602060405180830381865afa158015610366573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061038a9190610f71565b6001600160a01b0316145b156103a257506001610425565b5f6103ac866104fc565b604051631704429d60e21b815290915081906001600160a01b03821690635c110a74906103e190899089908990600401610f8c565b602060405180830381865afa1580156103fc573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104209190610fcb565b925050505b949350505050565b8161043781610942565b6001600160a01b03821661045e5760405163e6c4247b60e01b815260040160405180910390fd5b5f8381526004602090815260408083206001600160a01b038616845290915281205460ff16151590036104a457604051635571f9bf60e11b815260040160405180910390fd5b5f8381526004602090815260408083206001600160a01b0386168085529252808320805460ff1916905551909185917f907b4ea595c8ab4b33e5ce3b70882b068119d80536a1121f37c7ab73104e2b739190a3505050565b5f818152600160205260408120546001600160a01b03166105de575f8054604051637d7602a160e11b8152600481018590526001600160a01b039091169063faec0542906024015f60405180830381865afa15801561055d573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526105849190810190610e80565b6040808201515f908152600260205220549091506001600160a01b03166105be5760405163cefa6b0560e01b815260040160405180910390fd5b6040908101515f908152600260205220546001600160a01b031692915050565b505f908152600160205260409020546001600160a01b031690565b8161060381610942565b6001600160a01b03821661062a5760405163e6c4247b60e01b815260040160405180910390fd5b5f8381526004602090815260408083206001600160a01b0386168085529252808320805460ff1916600117905551909185917f4c71ed01d4ddaea8f2390c2feb6be34befe3571daf558f29632daebe2a1724339190a3505050565b8161068f81610942565b5f8381526002602052604080822080546001600160a01b0319166001600160a01b0386169081179091559051909185917fee78f4c2af2887839fdcba441de8e4d2d1b117b89ce7cc7a7f0a952871cc87d29190a3505050565b6040805160018082528183019092525f918291906020808301908036833701905050905083815f8151811061071f5761071f610fe6565b60209081029190910101525f80546040516335bdb71160e01b81526001600160a01b03909116906335bdb7119061075a908590600401610ffa565b5f60405180830381865afa158015610774573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261079b919081019061103c565b905080515f14806107c957505f5f1b815f815181106107bc576107bc610fe6565b60200260200101515f0151145b156107d8575f92505050610840565b5f815f815181106107eb576107eb610fe6565b6020026020010151606001519050846001600160a01b0316816001600160a01b0316148061083a57505f8681526004602090815260408083206001600160a01b038916845290915290205460ff165b93505050505b92915050565b5f8054604051637d7602a160e11b8152600481018590528492916001600160a01b03169063faec0542906024015f60405180830381865afa15801561088d573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526108b49190810190610e80565b60408101519091505f036108db57604051631de5d86f60e11b815260040160405180910390fd5b6108e88160400151610942565b5f8481526001602052604080822080546001600160a01b0319166001600160a01b0387169081179091559051909186917fee78f4c2af2887839fdcba441de8e4d2d1b117b89ce7cc7a7f0a952871cc87d29190a350505050565b6040805160018082528183019092525f916020808301908036833701905050905081815f8151811061097657610976610fe6565b60209081029190910101525f80546040516335bdb71160e01b81526001600160a01b03909116906335bdb711906109b1908590600401610ffa565b5f60405180830381865afa1580156109cb573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526109f2919081019061103c565b905080515f1480610a2057505f5f1b815f81518110610a1357610a13610fe6565b60200260200101515f0151145b15610a3e5760405163c4c1a0c560e01b815260040160405180910390fd5b5f815f81518110610a5157610a51610fe6565b60200260200101519050336001600160a01b031681606001516001600160a01b031614158015610a9a57505f84815260046020908152604080832033845290915290205460ff16155b15610ab75760405162d6b18f60e41b815260040160405180910390fd5b50505050565b8015158114610aca575f5ffd5b50565b5f5f60408385031215610ade575f5ffd5b823591506020830135610af081610abd565b809150509250929050565b5f60208284031215610b0b575f5ffd5b5035919050565b6001600160a01b0381168114610aca575f5ffd5b5f5f5f5f60608587031215610b39575f5ffd5b843593506020850135610b4b81610b12565b925060408501356001600160401b03811115610b65575f5ffd5b8501601f81018713610b75575f5ffd5b80356001600160401b03811115610b8a575f5ffd5b876020828401011115610b9b575f5ffd5b949793965060200194505050565b5f5f60408385031215610bba575f5ffd5b823591506020830135610af081610b12565b5f60208284031215610bdc575f5ffd5b8135610be781610b12565b9392505050565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715610c2457610c24610bee565b60405290565b60405161010081016001600160401b0381118282101715610c2457610c24610bee565b60405160a081016001600160401b0381118282101715610c2457610c24610bee565b604051601f8201601f191681016001600160401b0381118282101715610c9757610c97610bee565b604052919050565b5f82601f830112610cae575f5ffd5b8151602083015f5f6001600160401b03841115610ccd57610ccd610bee565b50601f8301601f1916602001610ce281610c6f565b915050828152858383011115610cf6575f5ffd5b8282602083015e5f92810160200192909252509392505050565b5f6001600160401b03821115610d2857610d28610bee565b5060051b60200190565b5f82601f830112610d41575f5ffd5b8151610d54610d4f82610d10565b610c6f565b8082825260208201915060208360051b860101925085831115610d75575f5ffd5b602085015b83811015610d92578051835260209283019201610d7a565b5095945050505050565b5f60408284031215610dac575f5ffd5b610db4610c02565b905081516001600160401b03811115610dcb575f5ffd5b8201601f81018413610ddb575f5ffd5b8051610de9610d4f82610d10565b8082825260208201915060208360051b850101925086831115610e0a575f5ffd5b602084015b83811015610e4a5780516001600160401b03811115610e2c575f5ffd5b610e3b89602083890101610c9f565b84525060209283019201610e0f565b50845250505060208201516001600160401b03811115610e68575f5ffd5b610e7484828501610d32565b60208301525092915050565b5f60208284031215610e90575f5ffd5b81516001600160401b03811115610ea5575f5ffd5b82016101008185031215610eb7575f5ffd5b610ebf610c2a565b8151815260208201516001600160401b03811115610edb575f5ffd5b610ee786828501610c9f565b6020830152506040828101519082015260608201516001600160401b03811115610f0f575f5ffd5b610f1b86828501610c9f565b6060830152506080828101519082015260a0808301519082015260c0808301519082015260e08201516001600160401b03811115610f57575f5ffd5b610f6386828501610d9c565b60e083015250949350505050565b5f60208284031215610f81575f5ffd5b8151610be781610b12565b6001600160a01b03841681526040602082018190528101829052818360608301375f818301606090810191909152601f909201601f1916010192915050565b5f60208284031215610fdb575f5ffd5b8151610be781610abd565b634e487b7160e01b5f52603260045260245ffd5b602080825282518282018190525f918401906040840190835b81811015611031578351835260209384019390920191600101611013565b509095945050505050565b5f6020828403121561104c575f5ffd5b81516001600160401b03811115611061575f5ffd5b8201601f81018413611071575f5ffd5b805161107f610d4f82610d10565b8082825260208201915060208360051b8501019250868311156110a0575f5ffd5b602084015b838110156111c65780516001600160401b038111156110c2575f5ffd5b850160a0818a03601f190112156110d7575f5ffd5b6110df610c4d565b6020820151815260408201516001600160401b038111156110fe575f5ffd5b61110d8b602083860101610c9f565b60208301525060608201516040820152608082015161112b81610b12565b606082015260a08201516001600160401b03811115611148575f5ffd5b60208184010192505089601f83011261115f575f5ffd5b815161116d610d4f82610d10565b8082825260208201915060208360051b86010192508c83111561118e575f5ffd5b6020850194505b828510156111b0578451825260209485019490910190611195565b60808401525050845250602092830192016110a5565b50969550505050505056fea2646970667358221220a64690fc14ffdad31863180e7638cdc6faf3c56531a9aa74c514c44e4b3672b364736f6c634300081c0033",
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

// IsBucketOwnerOrDelegate is a free data retrieval call binding the contract method 0xe124bdd9.
//
// Solidity: function isBucketOwnerOrDelegate(bytes32 bucketId, address user) view returns(bool)
func (_AccessManager *AccessManagerCaller) IsBucketOwnerOrDelegate(opts *bind.CallOpts, bucketId [32]byte, user common.Address) (bool, error) {
	var out []interface{}
	err := _AccessManager.contract.Call(opts, &out, "isBucketOwnerOrDelegate", bucketId, user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBucketOwnerOrDelegate is a free data retrieval call binding the contract method 0xe124bdd9.
//
// Solidity: function isBucketOwnerOrDelegate(bytes32 bucketId, address user) view returns(bool)
func (_AccessManager *AccessManagerSession) IsBucketOwnerOrDelegate(bucketId [32]byte, user common.Address) (bool, error) {
	return _AccessManager.Contract.IsBucketOwnerOrDelegate(&_AccessManager.CallOpts, bucketId, user)
}

// IsBucketOwnerOrDelegate is a free data retrieval call binding the contract method 0xe124bdd9.
//
// Solidity: function isBucketOwnerOrDelegate(bytes32 bucketId, address user) view returns(bool)
func (_AccessManager *AccessManagerCallerSession) IsBucketOwnerOrDelegate(bucketId [32]byte, user common.Address) (bool, error) {
	return _AccessManager.Contract.IsBucketOwnerOrDelegate(&_AccessManager.CallOpts, bucketId, user)
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

// DelegateAccess is a paid mutator transaction binding the contract method 0xa9b61d4d.
//
// Solidity: function delegateAccess(bytes32 bucketId, address delegated) returns()
func (_AccessManager *AccessManagerTransactor) DelegateAccess(opts *bind.TransactOpts, bucketId [32]byte, delegated common.Address) (*types.Transaction, error) {
	return _AccessManager.contract.Transact(opts, "delegateAccess", bucketId, delegated)
}

// DelegateAccess is a paid mutator transaction binding the contract method 0xa9b61d4d.
//
// Solidity: function delegateAccess(bytes32 bucketId, address delegated) returns()
func (_AccessManager *AccessManagerSession) DelegateAccess(bucketId [32]byte, delegated common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.DelegateAccess(&_AccessManager.TransactOpts, bucketId, delegated)
}

// DelegateAccess is a paid mutator transaction binding the contract method 0xa9b61d4d.
//
// Solidity: function delegateAccess(bytes32 bucketId, address delegated) returns()
func (_AccessManager *AccessManagerTransactorSession) DelegateAccess(bucketId [32]byte, delegated common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.DelegateAccess(&_AccessManager.TransactOpts, bucketId, delegated)
}

// RemoveAccess is a paid mutator transaction binding the contract method 0x5e2a22de.
//
// Solidity: function removeAccess(bytes32 bucketId, address delegated) returns()
func (_AccessManager *AccessManagerTransactor) RemoveAccess(opts *bind.TransactOpts, bucketId [32]byte, delegated common.Address) (*types.Transaction, error) {
	return _AccessManager.contract.Transact(opts, "removeAccess", bucketId, delegated)
}

// RemoveAccess is a paid mutator transaction binding the contract method 0x5e2a22de.
//
// Solidity: function removeAccess(bytes32 bucketId, address delegated) returns()
func (_AccessManager *AccessManagerSession) RemoveAccess(bucketId [32]byte, delegated common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.RemoveAccess(&_AccessManager.TransactOpts, bucketId, delegated)
}

// RemoveAccess is a paid mutator transaction binding the contract method 0x5e2a22de.
//
// Solidity: function removeAccess(bytes32 bucketId, address delegated) returns()
func (_AccessManager *AccessManagerTransactorSession) RemoveAccess(bucketId [32]byte, delegated common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.RemoveAccess(&_AccessManager.TransactOpts, bucketId, delegated)
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

// SetStorageContract is a paid mutator transaction binding the contract method 0xdc38b0a2.
//
// Solidity: function setStorageContract(address _storageAddress) returns()
func (_AccessManager *AccessManagerTransactor) SetStorageContract(opts *bind.TransactOpts, _storageAddress common.Address) (*types.Transaction, error) {
	return _AccessManager.contract.Transact(opts, "setStorageContract", _storageAddress)
}

// SetStorageContract is a paid mutator transaction binding the contract method 0xdc38b0a2.
//
// Solidity: function setStorageContract(address _storageAddress) returns()
func (_AccessManager *AccessManagerSession) SetStorageContract(_storageAddress common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.SetStorageContract(&_AccessManager.TransactOpts, _storageAddress)
}

// SetStorageContract is a paid mutator transaction binding the contract method 0xdc38b0a2.
//
// Solidity: function setStorageContract(address _storageAddress) returns()
func (_AccessManager *AccessManagerTransactorSession) SetStorageContract(_storageAddress common.Address) (*types.Transaction, error) {
	return _AccessManager.Contract.SetStorageContract(&_AccessManager.TransactOpts, _storageAddress)
}

// AccessManagerAccessDelegatedIterator is returned from FilterAccessDelegated and is used to iterate over the raw logs and unpacked data for AccessDelegated events raised by the AccessManager contract.
type AccessManagerAccessDelegatedIterator struct {
	Event *AccessManagerAccessDelegated // Event containing the contract specifics and raw log

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
func (it *AccessManagerAccessDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessManagerAccessDelegated)
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
		it.Event = new(AccessManagerAccessDelegated)
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
func (it *AccessManagerAccessDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessManagerAccessDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessManagerAccessDelegated represents a AccessDelegated event raised by the AccessManager contract.
type AccessManagerAccessDelegated struct {
	BucketId [32]byte
	Delegate common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAccessDelegated is a free log retrieval operation binding the contract event 0x4c71ed01d4ddaea8f2390c2feb6be34befe3571daf558f29632daebe2a172433.
//
// Solidity: event AccessDelegated(bytes32 indexed bucketId, address indexed delegate)
func (_AccessManager *AccessManagerFilterer) FilterAccessDelegated(opts *bind.FilterOpts, bucketId [][32]byte, delegate []common.Address) (*AccessManagerAccessDelegatedIterator, error) {

	var bucketIdRule []interface{}
	for _, bucketIdItem := range bucketId {
		bucketIdRule = append(bucketIdRule, bucketIdItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _AccessManager.contract.FilterLogs(opts, "AccessDelegated", bucketIdRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return &AccessManagerAccessDelegatedIterator{contract: _AccessManager.contract, event: "AccessDelegated", logs: logs, sub: sub}, nil
}

// WatchAccessDelegated is a free log subscription operation binding the contract event 0x4c71ed01d4ddaea8f2390c2feb6be34befe3571daf558f29632daebe2a172433.
//
// Solidity: event AccessDelegated(bytes32 indexed bucketId, address indexed delegate)
func (_AccessManager *AccessManagerFilterer) WatchAccessDelegated(opts *bind.WatchOpts, sink chan<- *AccessManagerAccessDelegated, bucketId [][32]byte, delegate []common.Address) (event.Subscription, error) {

	var bucketIdRule []interface{}
	for _, bucketIdItem := range bucketId {
		bucketIdRule = append(bucketIdRule, bucketIdItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _AccessManager.contract.WatchLogs(opts, "AccessDelegated", bucketIdRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessManagerAccessDelegated)
				if err := _AccessManager.contract.UnpackLog(event, "AccessDelegated", log); err != nil {
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

// ParseAccessDelegated is a log parse operation binding the contract event 0x4c71ed01d4ddaea8f2390c2feb6be34befe3571daf558f29632daebe2a172433.
//
// Solidity: event AccessDelegated(bytes32 indexed bucketId, address indexed delegate)
func (_AccessManager *AccessManagerFilterer) ParseAccessDelegated(log types.Log) (*AccessManagerAccessDelegated, error) {
	event := new(AccessManagerAccessDelegated)
	if err := _AccessManager.contract.UnpackLog(event, "AccessDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessManagerAccessRemovedIterator is returned from FilterAccessRemoved and is used to iterate over the raw logs and unpacked data for AccessRemoved events raised by the AccessManager contract.
type AccessManagerAccessRemovedIterator struct {
	Event *AccessManagerAccessRemoved // Event containing the contract specifics and raw log

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
func (it *AccessManagerAccessRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessManagerAccessRemoved)
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
		it.Event = new(AccessManagerAccessRemoved)
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
func (it *AccessManagerAccessRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessManagerAccessRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessManagerAccessRemoved represents a AccessRemoved event raised by the AccessManager contract.
type AccessManagerAccessRemoved struct {
	BucketId [32]byte
	Delegate common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAccessRemoved is a free log retrieval operation binding the contract event 0x907b4ea595c8ab4b33e5ce3b70882b068119d80536a1121f37c7ab73104e2b73.
//
// Solidity: event AccessRemoved(bytes32 indexed bucketId, address indexed delegate)
func (_AccessManager *AccessManagerFilterer) FilterAccessRemoved(opts *bind.FilterOpts, bucketId [][32]byte, delegate []common.Address) (*AccessManagerAccessRemovedIterator, error) {

	var bucketIdRule []interface{}
	for _, bucketIdItem := range bucketId {
		bucketIdRule = append(bucketIdRule, bucketIdItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _AccessManager.contract.FilterLogs(opts, "AccessRemoved", bucketIdRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return &AccessManagerAccessRemovedIterator{contract: _AccessManager.contract, event: "AccessRemoved", logs: logs, sub: sub}, nil
}

// WatchAccessRemoved is a free log subscription operation binding the contract event 0x907b4ea595c8ab4b33e5ce3b70882b068119d80536a1121f37c7ab73104e2b73.
//
// Solidity: event AccessRemoved(bytes32 indexed bucketId, address indexed delegate)
func (_AccessManager *AccessManagerFilterer) WatchAccessRemoved(opts *bind.WatchOpts, sink chan<- *AccessManagerAccessRemoved, bucketId [][32]byte, delegate []common.Address) (event.Subscription, error) {

	var bucketIdRule []interface{}
	for _, bucketIdItem := range bucketId {
		bucketIdRule = append(bucketIdRule, bucketIdItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _AccessManager.contract.WatchLogs(opts, "AccessRemoved", bucketIdRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessManagerAccessRemoved)
				if err := _AccessManager.contract.UnpackLog(event, "AccessRemoved", log); err != nil {
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

// ParseAccessRemoved is a log parse operation binding the contract event 0x907b4ea595c8ab4b33e5ce3b70882b068119d80536a1121f37c7ab73104e2b73.
//
// Solidity: event AccessRemoved(bytes32 indexed bucketId, address indexed delegate)
func (_AccessManager *AccessManagerFilterer) ParseAccessRemoved(log types.Log) (*AccessManagerAccessRemoved, error) {
	event := new(AccessManagerAccessRemoved)
	if err := _AccessManager.contract.UnpackLog(event, "AccessRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
