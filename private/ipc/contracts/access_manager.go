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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"storageAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BucketNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FileDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoDelegatedAccess\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPolicy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBucketOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"AccessDelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"AccessRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"policyContract\",\"type\":\"address\"}],\"name\":\"PolicyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"name\":\"PublicAccessChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"isPublic\",\"type\":\"bool\"}],\"name\":\"changePublicAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegated\",\"type\":\"address\"}],\"name\":\"delegateAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"getFileAccessInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"}],\"name\":\"getPolicy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getValidateAccess\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getValidateAccessToBucket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isBucketOwnerOrDelegate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"delegated\",\"type\":\"address\"}],\"name\":\"removeAccess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bucketId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"policyContract\",\"type\":\"address\"}],\"name\":\"setBucketPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"fileId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"policyContract\",\"type\":\"address\"}],\"name\":\"setFilePolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_storageAddress\",\"type\":\"address\"}],\"name\":\"setStorageContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageContract\",\"outputs\":[{\"internalType\":\"contractIStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50604051611483380380611483833981016040819052602b91604e565b5f80546001600160a01b0319166001600160a01b03929092169190911790556079565b5f60208284031215605d575f5ffd5b81516001600160a01b03811681146072575f5ffd5b9392505050565b6113fd806100865f395ff3fe608060405234801561000f575f5ffd5b50600436106100b1575f3560e01c8063a3f685f91161006e578063a3f685f91461019d578063a9b61d4d146101b0578063c05d0d6c146101c3578063dc38b0a2146101d6578063e124bdd914610205578063e13212f814610218575f5ffd5b806311ce0267146100b557806338f5f5b2146100e45780633a82c912146100f957806344078802146101545780635e2a22de14610177578063a20e18df1461018a575b5f5ffd5b5f546100c7906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100f76100f2366004610c2d565b61022b565b005b610135610107366004610c5b565b5f908152600160209081526040808320546003909252909120546001600160a01b039091169160ff90911690565b604080516001600160a01b0390931683529015156020830152016100db565b610167610162366004610c86565b610325565b60405190151581526020016100db565b6100f7610185366004610d09565b610461565b610167610198366004610c86565b610530565b6100c76101ab366004610c5b565b61062d565b6100f76101be366004610d09565b610719565b6100f76101d1366004610d09565b6107a5565b6100f76101e4366004610d2c565b5f80546001600160a01b0319166001600160a01b0392909216919091179055565b610167610213366004610d09565b610808565b6100f7610226366004610d09565b610966565b5f8054604051637d7602a160e11b8152600481018590528492916001600160a01b03169063faec0542906024015f60405180830381865afa158015610272573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526102999190810190611076565b60408101519091505f036102c057604051631de5d86f60e11b815260040160405180910390fd5b6102cd8160400151610a62565b5f84815260036020908152604091829020805460ff1916861515908117909155915191825285917f9500a58cfb37cef230929cd9f25ce92c41374416f23b1825232d0905a7e73d5a910160405180910390a250505050565b5f8481526003602052604081205460ff16806103b357505f546040516306d1d7bd60e21b8152600481018790526001600160a01b03868116921690631b475ef490602401602060405180830381865afa158015610384573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103a89190611167565b6001600160a01b0316145b156103c057506001610459565b5f6103ca8661062d565b90506001600160a01b0381166103e3575f915050610459565b604051631704429d60e21b815281906001600160a01b03821690635c110a749061041590899089908990600401611182565b602060405180830381865afa158015610430573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061045491906111c1565b925050505b949350505050565b8161046b81610a62565b6001600160a01b0382166104925760405163e6c4247b60e01b815260040160405180910390fd5b5f8381526004602090815260408083206001600160a01b038616845290915281205460ff16151590036104d857604051635571f9bf60e11b815260040160405180910390fd5b5f8381526004602090815260408083206001600160a01b0386168085529252808320805460ff1916905551909185917f907b4ea595c8ab4b33e5ce3b70882b068119d80536a1121f37c7ab73104e2b739190a3505050565b5f5f61053b86610ac6565b60608101519091506001600160a01b031633148061057157505f86815260046020908152604080832033845290915290205460ff165b15610580576001915050610459565b61058986610ac6565b505f868152600260205260409020546001600160a01b0316806105b0575f92505050610459565b604051631704429d60e21b815281906001600160a01b03821690635c110a74906105e2908a908a908a90600401611182565b602060405180830381865afa1580156105fd573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061062191906111c1565b98975050505050505050565b5f818152600160205260408120546001600160a01b03166106fe575f8054604051637d7602a160e11b8152600481018590526001600160a01b039091169063faec0542906024015f60405180830381865afa15801561068e573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526106b59190810190611076565b6040808201515f908152600260205220549091506001600160a01b03166106de57505f92915050565b6040908101515f908152600260205220546001600160a01b031692915050565b505f908152600160205260409020546001600160a01b031690565b8161072381610a62565b6001600160a01b03821661074a5760405163e6c4247b60e01b815260040160405180910390fd5b5f8381526004602090815260408083206001600160a01b0386168085529252808320805460ff1916600117905551909185917f4c71ed01d4ddaea8f2390c2feb6be34befe3571daf558f29632daebe2a1724339190a3505050565b816107af81610a62565b5f8381526002602052604080822080546001600160a01b0319166001600160a01b0386169081179091559051909185917fee78f4c2af2887839fdcba441de8e4d2d1b117b89ce7cc7a7f0a952871cc87d29190a3505050565b6040805160018082528183019092525f918291906020808301908036833701905050905083815f8151811061083f5761083f6111dc565b60209081029190910101525f80546040516335bdb71160e01b81526001600160a01b03909116906335bdb7119061087a9085906004016111f0565b5f60405180830381865afa158015610894573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526108bb9190810190611232565b905080515f14806108e957505f5f1b815f815181106108dc576108dc6111dc565b60200260200101515f0151145b156108f8575f92505050610960565b5f815f8151811061090b5761090b6111dc565b6020026020010151606001519050846001600160a01b0316816001600160a01b0316148061095a57505f8681526004602090815260408083206001600160a01b038916845290915290205460ff165b93505050505b92915050565b5f8054604051637d7602a160e11b8152600481018590528492916001600160a01b03169063faec0542906024015f60405180830381865afa1580156109ad573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526109d49190810190611076565b60408101519091505f036109fb57604051631de5d86f60e11b815260040160405180910390fd5b610a088160400151610a62565b5f8481526001602052604080822080546001600160a01b0319166001600160a01b0387169081179091559051909186917fee78f4c2af2887839fdcba441de8e4d2d1b117b89ce7cc7a7f0a952871cc87d29190a350505050565b5f610a6c82610ac6565b60608101519091506001600160a01b03163314801590610aa557505f82815260046020908152604080832033845290915290205460ff16155b15610ac25760405162d6b18f60e41b815260040160405180910390fd5b5050565b610afe6040518060a001604052805f8152602001606081526020015f81526020015f6001600160a01b03168152602001606081525090565b6040805160018082528183019092525f916020808301908036833701905050905082815f81518110610b3257610b326111dc565b60209081029190910101525f80546040516335bdb71160e01b81526001600160a01b03909116906335bdb71190610b6d9085906004016111f0565b5f60405180830381865afa158015610b87573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610bae9190810190611232565b905080515f1480610bdc57505f5f1b815f81518110610bcf57610bcf6111dc565b60200260200101515f0151145b15610bfa5760405163c4c1a0c560e01b815260040160405180910390fd5b805f81518110610c0c57610c0c6111dc565b602002602001015192505050919050565b8015158114610c2a575f5ffd5b50565b5f5f60408385031215610c3e575f5ffd5b823591506020830135610c5081610c1d565b809150509250929050565b5f60208284031215610c6b575f5ffd5b5035919050565b6001600160a01b0381168114610c2a575f5ffd5b5f5f5f5f60608587031215610c99575f5ffd5b843593506020850135610cab81610c72565b925060408501356001600160401b03811115610cc5575f5ffd5b8501601f81018713610cd5575f5ffd5b80356001600160401b03811115610cea575f5ffd5b876020828401011115610cfb575f5ffd5b949793965060200194505050565b5f5f60408385031215610d1a575f5ffd5b823591506020830135610c5081610c72565b5f60208284031215610d3c575f5ffd5b8135610d4781610c72565b9392505050565b634e487b7160e01b5f52604160045260245ffd5b604051606081016001600160401b0381118282101715610d8457610d84610d4e565b60405290565b60405161010081016001600160401b0381118282101715610d8457610d84610d4e565b60405160a081016001600160401b0381118282101715610d8457610d84610d4e565b604051601f8201601f191681016001600160401b0381118282101715610df757610df7610d4e565b604052919050565b5f82601f830112610e0e575f5ffd5b8151602083015f5f6001600160401b03841115610e2d57610e2d610d4e565b50601f8301601f1916602001610e4281610dcf565b915050828152858383011115610e56575f5ffd5b8282602083015e5f92810160200192909252509392505050565b5f6001600160401b03821115610e8857610e88610d4e565b5060051b60200190565b5f82601f830112610ea1575f5ffd5b8151610eb4610eaf82610e70565b610dcf565b8082825260208201915060208360051b860101925085831115610ed5575f5ffd5b602085015b83811015610ef2578051835260209283019201610eda565b5095945050505050565b5f82601f830112610f0b575f5ffd5b8151610f19610eaf82610e70565b8082825260208201915060208360051b860101925085831115610f3a575f5ffd5b602085015b83811015610ef257805163ffffffff81168114610f5a575f5ffd5b835260209283019201610f3f565b5f60608284031215610f78575f5ffd5b610f80610d62565b905081516001600160401b03811115610f97575f5ffd5b8201601f81018413610fa7575f5ffd5b8051610fb5610eaf82610e70565b8082825260208201915060208360051b850101925086831115610fd6575f5ffd5b602084015b838110156110165780516001600160401b03811115610ff8575f5ffd5b61100789602083890101610dff565b84525060209283019201610fdb565b50845250505060208201516001600160401b03811115611034575f5ffd5b61104084828501610e92565b60208301525060408201516001600160401b0381111561105e575f5ffd5b61106a84828501610efc565b60408301525092915050565b5f60208284031215611086575f5ffd5b81516001600160401b0381111561109b575f5ffd5b820161010081850312156110ad575f5ffd5b6110b5610d8a565b8151815260208201516001600160401b038111156110d1575f5ffd5b6110dd86828501610dff565b6020830152506040828101519082015260608201516001600160401b03811115611105575f5ffd5b61111186828501610dff565b6060830152506080828101519082015260a0808301519082015260c0808301519082015260e08201516001600160401b0381111561114d575f5ffd5b61115986828501610f68565b60e083015250949350505050565b5f60208284031215611177575f5ffd5b8151610d4781610c72565b6001600160a01b03841681526040602082018190528101829052818360608301375f818301606090810191909152601f909201601f1916010192915050565b5f602082840312156111d1575f5ffd5b8151610d4781610c1d565b634e487b7160e01b5f52603260045260245ffd5b602080825282518282018190525f918401906040840190835b81811015611227578351835260209384019390920191600101611209565b509095945050505050565b5f60208284031215611242575f5ffd5b81516001600160401b03811115611257575f5ffd5b8201601f81018413611267575f5ffd5b8051611275610eaf82610e70565b8082825260208201915060208360051b850101925086831115611296575f5ffd5b602084015b838110156113bc5780516001600160401b038111156112b8575f5ffd5b850160a0818a03601f190112156112cd575f5ffd5b6112d5610dad565b6020820151815260408201516001600160401b038111156112f4575f5ffd5b6113038b602083860101610dff565b60208301525060608201516040820152608082015161132181610c72565b606082015260a08201516001600160401b0381111561133e575f5ffd5b60208184010192505089601f830112611355575f5ffd5b8151611363610eaf82610e70565b8082825260208201915060208360051b86010192508c831115611384575f5ffd5b6020850194505b828510156113a657845182526020948501949091019061138b565b608084015250508452506020928301920161129b565b50969550505050505056fea264697066735822122008f473355b5f57cd9ab144637dab3ff483ba17f8f6904099b1dbfd761dc1cdf164736f6c634300081c0033",
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

// GetValidateAccessToBucket is a free data retrieval call binding the contract method 0xa20e18df.
//
// Solidity: function getValidateAccessToBucket(bytes32 bucketId, address user, bytes data) view returns(bool)
func (_AccessManager *AccessManagerCaller) GetValidateAccessToBucket(opts *bind.CallOpts, bucketId [32]byte, user common.Address, data []byte) (bool, error) {
	var out []interface{}
	err := _AccessManager.contract.Call(opts, &out, "getValidateAccessToBucket", bucketId, user, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetValidateAccessToBucket is a free data retrieval call binding the contract method 0xa20e18df.
//
// Solidity: function getValidateAccessToBucket(bytes32 bucketId, address user, bytes data) view returns(bool)
func (_AccessManager *AccessManagerSession) GetValidateAccessToBucket(bucketId [32]byte, user common.Address, data []byte) (bool, error) {
	return _AccessManager.Contract.GetValidateAccessToBucket(&_AccessManager.CallOpts, bucketId, user, data)
}

// GetValidateAccessToBucket is a free data retrieval call binding the contract method 0xa20e18df.
//
// Solidity: function getValidateAccessToBucket(bytes32 bucketId, address user, bytes data) view returns(bool)
func (_AccessManager *AccessManagerCallerSession) GetValidateAccessToBucket(bucketId [32]byte, user common.Address, data []byte) (bool, error) {
	return _AccessManager.Contract.GetValidateAccessToBucket(&_AccessManager.CallOpts, bucketId, user, data)
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
