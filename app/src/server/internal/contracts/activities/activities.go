// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package activities

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ActivitiesData is an auto generated low-level Go binding around an user-defined struct.
type ActivitiesData struct {
	Version       [32]byte
	GeneratedTime [32]byte
	LinkedData    [32]byte
}

// ActivitiesStoreABI is the input ABI used to generate the binding from.
const ActivitiesStoreABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"}],\"internalType\":\"structActivitiesData\",\"name\":\"_activities\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"}],\"internalType\":\"structActivitiesData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ActivitiesStoreFuncSigs maps the 4-byte function signature to its string representation.
var ActivitiesStoreFuncSigs = map[string]string{
	"6d4ce63c": "get()",
}

// ActivitiesStoreBin is the compiled bytecode used for deploying new contracts.
var ActivitiesStoreBin = "0x608060405234801561001057600080fd5b506040516101f03803806101f083398101604081905261002f916100a3565b805160001a60f81b6001600160f81b031916158015906100625750602081015160001a60f81b6001600160f81b03191615155b80156100815750604081015160001a60f81b6001600160f81b03191615155b61008a57600080fd5b80516000556020810151600155604001516002556100f9565b6000606082840312156100b4578081fd5b604051606081016001600160401b03811182821017156100d2578283fd5b80604052508251815260208301516020820152604083015160408201528091505092915050565b60e9806101076000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80636d4ce63c14602d575b600080fd5b60336047565b604051603e91906092565b60405180910390f35b604d6072565b5060408051606081018252600054815260015460208201526002549181019190915290565b604080516060810182526000808252602082018190529181019190915290565b8151815260208083015190820152604091820151918101919091526060019056fea26469706673582212201a9b9d320ed50be4b0237b03f97c90751ea4fe4b4d01ccb52044fbfb7da9bd9c64736f6c63430006060033"

// DeployActivitiesStore deploys a new Ethereum contract, binding an instance of ActivitiesStore to it.
func DeployActivitiesStore(auth *bind.TransactOpts, backend bind.ContractBackend, _activities ActivitiesData) (common.Address, *types.Transaction, *ActivitiesStore, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivitiesStoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ActivitiesStoreBin), backend, _activities)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ActivitiesStore{ActivitiesStoreCaller: ActivitiesStoreCaller{contract: contract}, ActivitiesStoreTransactor: ActivitiesStoreTransactor{contract: contract}, ActivitiesStoreFilterer: ActivitiesStoreFilterer{contract: contract}}, nil
}

// ActivitiesStore is an auto generated Go binding around an Ethereum contract.
type ActivitiesStore struct {
	ActivitiesStoreCaller     // Read-only binding to the contract
	ActivitiesStoreTransactor // Write-only binding to the contract
	ActivitiesStoreFilterer   // Log filterer for contract events
}

// ActivitiesStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivitiesStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivitiesStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivitiesStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivitiesStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivitiesStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivitiesStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivitiesStoreSession struct {
	Contract     *ActivitiesStore  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ActivitiesStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivitiesStoreCallerSession struct {
	Contract *ActivitiesStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ActivitiesStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivitiesStoreTransactorSession struct {
	Contract     *ActivitiesStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ActivitiesStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivitiesStoreRaw struct {
	Contract *ActivitiesStore // Generic contract binding to access the raw methods on
}

// ActivitiesStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivitiesStoreCallerRaw struct {
	Contract *ActivitiesStoreCaller // Generic read-only contract binding to access the raw methods on
}

// ActivitiesStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivitiesStoreTransactorRaw struct {
	Contract *ActivitiesStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivitiesStore creates a new instance of ActivitiesStore, bound to a specific deployed contract.
func NewActivitiesStore(address common.Address, backend bind.ContractBackend) (*ActivitiesStore, error) {
	contract, err := bindActivitiesStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ActivitiesStore{ActivitiesStoreCaller: ActivitiesStoreCaller{contract: contract}, ActivitiesStoreTransactor: ActivitiesStoreTransactor{contract: contract}, ActivitiesStoreFilterer: ActivitiesStoreFilterer{contract: contract}}, nil
}

// NewActivitiesStoreCaller creates a new read-only instance of ActivitiesStore, bound to a specific deployed contract.
func NewActivitiesStoreCaller(address common.Address, caller bind.ContractCaller) (*ActivitiesStoreCaller, error) {
	contract, err := bindActivitiesStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivitiesStoreCaller{contract: contract}, nil
}

// NewActivitiesStoreTransactor creates a new write-only instance of ActivitiesStore, bound to a specific deployed contract.
func NewActivitiesStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivitiesStoreTransactor, error) {
	contract, err := bindActivitiesStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivitiesStoreTransactor{contract: contract}, nil
}

// NewActivitiesStoreFilterer creates a new log filterer instance of ActivitiesStore, bound to a specific deployed contract.
func NewActivitiesStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivitiesStoreFilterer, error) {
	contract, err := bindActivitiesStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivitiesStoreFilterer{contract: contract}, nil
}

// bindActivitiesStore binds a generic wrapper to an already deployed contract.
func bindActivitiesStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivitiesStoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivitiesStore *ActivitiesStoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivitiesStore.Contract.ActivitiesStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivitiesStore *ActivitiesStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivitiesStore.Contract.ActivitiesStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivitiesStore *ActivitiesStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivitiesStore.Contract.ActivitiesStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivitiesStore *ActivitiesStoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivitiesStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivitiesStore *ActivitiesStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivitiesStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivitiesStore *ActivitiesStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivitiesStore.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(ActivitiesData)
func (_ActivitiesStore *ActivitiesStoreCaller) Get(opts *bind.CallOpts) (ActivitiesData, error) {
	var (
		ret0 = new(ActivitiesData)
	)
	out := ret0
	err := _ActivitiesStore.contract.Call(opts, out, "get")
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(ActivitiesData)
func (_ActivitiesStore *ActivitiesStoreSession) Get() (ActivitiesData, error) {
	return _ActivitiesStore.Contract.Get(&_ActivitiesStore.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(ActivitiesData)
func (_ActivitiesStore *ActivitiesStoreCallerSession) Get() (ActivitiesData, error) {
	return _ActivitiesStore.Contract.Get(&_ActivitiesStore.CallOpts)
}

// IATIActivitiesABI is the input ABI used to generate the binding from.
const IATIActivitiesABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"contractName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"Set\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"}],\"internalType\":\"structActivitiesData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getter\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"}],\"internalType\":\"structActivitiesData\",\"name\":\"_activities\",\"type\":\"tuple\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setter\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IATIActivitiesFuncSigs maps the 4-byte function signature to its string representation.
var IATIActivitiesFuncSigs = map[string]string{
	"8eaa6ac0": "get(bytes32)",
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
	"993a04b7": "getter()",
	"91a95fbc": "set(bytes32,(bytes32,bytes32,bytes32))",
	"3f3108f7": "setter()",
}

// IATIActivitiesBin is the compiled bytecode used for deploying new contracts.
var IATIActivitiesBin = "0x608060405234801561001057600080fd5b5060006002556106c3806100256000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80633f3108f71461006757806367e0badb146100855780638eaa6ac01461009a57806391a95fbc146100ba578063993a04b7146100cf578063bc599341146100d7575b600080fd5b61006f6100ea565b60405161007c9190610414565b60405180910390f35b61008d6100f5565b60405161007c919061040b565b6100ad6100a836600461035d565b6100fb565b60405161007c9190610455565b6100cd6100c8366004610375565b6101d6565b005b61006f61025d565b61008d6100e536600461035d565b610268565b63246a57ef60e21b90565b60025490565b610103610330565b8160001a60f81b6001600160f81b03191661011d57600080fd5b6000828152602081905260409020600101546001600160a01b031661014157600080fd5b60008281526020819052604090819020600101548151631b53398f60e21b815291516001600160a01b03909116918291636d4ce63c91600480820192606092909190829003018186803b15801561019757600080fd5b505afa1580156101ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101cf91906103ce565b9392505050565b6000816040516101e590610350565b6101ef9190610455565b604051809103906000f08015801561020b573d6000803e3d6000fd5b5090506102206000848363ffffffff6102a016565b507f89752cf887704c720fdae38088909ed147c755f39bd110525728448ec4cefc8b836040516102509190610429565b60405180910390a1505050565b63023aa9ab60e61b90565b600254600090821061027957600080fd5b600180548390811061028757fe5b9060005260206000209060020201600001549050919050565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b03851617905580156102de5760019150506101cf565b506001808501805491820180825560008681526020889052604090208190558591908390811061030a57fe5b600091825260208220600291820201929092559086018054600101905591506101cf9050565b604080516060810182526000808252602082018190529181019190915290565b6101f08061049e83390190565b60006020828403121561036e578081fd5b5035919050565b6000808284036080811215610388578182fd5b833592506060601f198201121561039d578182fd5b506103a86060610476565b602084013581526040840135602082015260608401356040820152809150509250929050565b6000606082840312156103df578081fd5b6103e96060610476565b8251815260208301516020820152604083015160408201528091505092915050565b90815260200190565b6001600160e01b031991909116815260200190565b6040808252600a90820152694163746976697469657360b01b6060820152602081019190915260800190565b81518152602080830151908201526040918201519181019190915260600190565b60405181810167ffffffffffffffff8111828210171561049557600080fd5b60405291905056fe608060405234801561001057600080fd5b506040516101f03803806101f083398101604081905261002f916100a3565b805160001a60f81b6001600160f81b031916158015906100625750602081015160001a60f81b6001600160f81b03191615155b80156100815750604081015160001a60f81b6001600160f81b03191615155b61008a57600080fd5b80516000556020810151600155604001516002556100f9565b6000606082840312156100b4578081fd5b604051606081016001600160401b03811182821017156100d2578283fd5b80604052508251815260208301516020820152604083015160408201528091505092915050565b60e9806101076000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80636d4ce63c14602d575b600080fd5b60336047565b604051603e91906092565b60405180910390f35b604d6072565b5060408051606081018252600054815260015460208201526002549181019190915290565b604080516060810182526000808252602082018190529181019190915290565b8151815260208083015190820152604091820151918101919091526060019056fea26469706673582212201a9b9d320ed50be4b0237b03f97c90751ea4fe4b4d01ccb52044fbfb7da9bd9c64736f6c63430006060033a2646970667358221220c3c208085041b489ff848dcbde7f3285922005441d5a1945c480284beac7e4e764736f6c63430006060033"

// DeployIATIActivities deploys a new Ethereum contract, binding an instance of IATIActivities to it.
func DeployIATIActivities(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IATIActivities, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivitiesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IATIActivitiesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IATIActivities{IATIActivitiesCaller: IATIActivitiesCaller{contract: contract}, IATIActivitiesTransactor: IATIActivitiesTransactor{contract: contract}, IATIActivitiesFilterer: IATIActivitiesFilterer{contract: contract}}, nil
}

// IATIActivities is an auto generated Go binding around an Ethereum contract.
type IATIActivities struct {
	IATIActivitiesCaller     // Read-only binding to the contract
	IATIActivitiesTransactor // Write-only binding to the contract
	IATIActivitiesFilterer   // Log filterer for contract events
}

// IATIActivitiesCaller is an auto generated read-only Go binding around an Ethereum contract.
type IATIActivitiesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivitiesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IATIActivitiesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivitiesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IATIActivitiesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivitiesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IATIActivitiesSession struct {
	Contract     *IATIActivities   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IATIActivitiesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IATIActivitiesCallerSession struct {
	Contract *IATIActivitiesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IATIActivitiesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IATIActivitiesTransactorSession struct {
	Contract     *IATIActivitiesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IATIActivitiesRaw is an auto generated low-level Go binding around an Ethereum contract.
type IATIActivitiesRaw struct {
	Contract *IATIActivities // Generic contract binding to access the raw methods on
}

// IATIActivitiesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IATIActivitiesCallerRaw struct {
	Contract *IATIActivitiesCaller // Generic read-only contract binding to access the raw methods on
}

// IATIActivitiesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IATIActivitiesTransactorRaw struct {
	Contract *IATIActivitiesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIATIActivities creates a new instance of IATIActivities, bound to a specific deployed contract.
func NewIATIActivities(address common.Address, backend bind.ContractBackend) (*IATIActivities, error) {
	contract, err := bindIATIActivities(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IATIActivities{IATIActivitiesCaller: IATIActivitiesCaller{contract: contract}, IATIActivitiesTransactor: IATIActivitiesTransactor{contract: contract}, IATIActivitiesFilterer: IATIActivitiesFilterer{contract: contract}}, nil
}

// NewIATIActivitiesCaller creates a new read-only instance of IATIActivities, bound to a specific deployed contract.
func NewIATIActivitiesCaller(address common.Address, caller bind.ContractCaller) (*IATIActivitiesCaller, error) {
	contract, err := bindIATIActivities(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IATIActivitiesCaller{contract: contract}, nil
}

// NewIATIActivitiesTransactor creates a new write-only instance of IATIActivities, bound to a specific deployed contract.
func NewIATIActivitiesTransactor(address common.Address, transactor bind.ContractTransactor) (*IATIActivitiesTransactor, error) {
	contract, err := bindIATIActivities(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IATIActivitiesTransactor{contract: contract}, nil
}

// NewIATIActivitiesFilterer creates a new log filterer instance of IATIActivities, bound to a specific deployed contract.
func NewIATIActivitiesFilterer(address common.Address, filterer bind.ContractFilterer) (*IATIActivitiesFilterer, error) {
	contract, err := bindIATIActivities(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IATIActivitiesFilterer{contract: contract}, nil
}

// bindIATIActivities binds a generic wrapper to an already deployed contract.
func bindIATIActivities(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivitiesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIActivities *IATIActivitiesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIActivities.Contract.IATIActivitiesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIActivities *IATIActivitiesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIActivities.Contract.IATIActivitiesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIActivities *IATIActivitiesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIActivities.Contract.IATIActivitiesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIActivities *IATIActivitiesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIActivities.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIActivities *IATIActivitiesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIActivities.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIActivities *IATIActivitiesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIActivities.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _ref) view returns(ActivitiesData)
func (_IATIActivities *IATIActivitiesCaller) Get(opts *bind.CallOpts, _ref [32]byte) (ActivitiesData, error) {
	var (
		ret0 = new(ActivitiesData)
	)
	out := ret0
	err := _IATIActivities.contract.Call(opts, out, "get", _ref)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _ref) view returns(ActivitiesData)
func (_IATIActivities *IATIActivitiesSession) Get(_ref [32]byte) (ActivitiesData, error) {
	return _IATIActivities.Contract.Get(&_IATIActivities.CallOpts, _ref)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _ref) view returns(ActivitiesData)
func (_IATIActivities *IATIActivitiesCallerSession) Get(_ref [32]byte) (ActivitiesData, error) {
	return _IATIActivities.Contract.Get(&_IATIActivities.CallOpts, _ref)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_IATIActivities *IATIActivitiesCaller) GetNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIActivities.contract.Call(opts, out, "getNum")
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_IATIActivities *IATIActivitiesSession) GetNum() (*big.Int, error) {
	return _IATIActivities.Contract.GetNum(&_IATIActivities.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_IATIActivities *IATIActivitiesCallerSession) GetNum() (*big.Int, error) {
	return _IATIActivities.Contract.GetNum(&_IATIActivities.CallOpts)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_IATIActivities *IATIActivitiesCaller) GetReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivities.contract.Call(opts, out, "getReference", _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_IATIActivities *IATIActivitiesSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _IATIActivities.Contract.GetReference(&_IATIActivities.CallOpts, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_IATIActivities *IATIActivitiesCallerSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _IATIActivities.Contract.GetReference(&_IATIActivities.CallOpts, _index)
}

// Getter is a free data retrieval call binding the contract method 0x993a04b7.
//
// Solidity: function getter() view returns(bytes4)
func (_IATIActivities *IATIActivitiesCaller) Getter(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _IATIActivities.contract.Call(opts, out, "getter")
	return *ret0, err
}

// Getter is a free data retrieval call binding the contract method 0x993a04b7.
//
// Solidity: function getter() view returns(bytes4)
func (_IATIActivities *IATIActivitiesSession) Getter() ([4]byte, error) {
	return _IATIActivities.Contract.Getter(&_IATIActivities.CallOpts)
}

// Getter is a free data retrieval call binding the contract method 0x993a04b7.
//
// Solidity: function getter() view returns(bytes4)
func (_IATIActivities *IATIActivitiesCallerSession) Getter() ([4]byte, error) {
	return _IATIActivities.Contract.Getter(&_IATIActivities.CallOpts)
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(bytes4)
func (_IATIActivities *IATIActivitiesCaller) Setter(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _IATIActivities.contract.Call(opts, out, "setter")
	return *ret0, err
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(bytes4)
func (_IATIActivities *IATIActivitiesSession) Setter() ([4]byte, error) {
	return _IATIActivities.Contract.Setter(&_IATIActivities.CallOpts)
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(bytes4)
func (_IATIActivities *IATIActivitiesCallerSession) Setter() ([4]byte, error) {
	return _IATIActivities.Contract.Setter(&_IATIActivities.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x91a95fbc.
//
// Solidity: function set(bytes32 _ref, ActivitiesData _activities) returns()
func (_IATIActivities *IATIActivitiesTransactor) Set(opts *bind.TransactOpts, _ref [32]byte, _activities ActivitiesData) (*types.Transaction, error) {
	return _IATIActivities.contract.Transact(opts, "set", _ref, _activities)
}

// Set is a paid mutator transaction binding the contract method 0x91a95fbc.
//
// Solidity: function set(bytes32 _ref, ActivitiesData _activities) returns()
func (_IATIActivities *IATIActivitiesSession) Set(_ref [32]byte, _activities ActivitiesData) (*types.Transaction, error) {
	return _IATIActivities.Contract.Set(&_IATIActivities.TransactOpts, _ref, _activities)
}

// Set is a paid mutator transaction binding the contract method 0x91a95fbc.
//
// Solidity: function set(bytes32 _ref, ActivitiesData _activities) returns()
func (_IATIActivities *IATIActivitiesTransactorSession) Set(_ref [32]byte, _activities ActivitiesData) (*types.Transaction, error) {
	return _IATIActivities.Contract.Set(&_IATIActivities.TransactOpts, _ref, _activities)
}

// IATIActivitiesSetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the IATIActivities contract.
type IATIActivitiesSetIterator struct {
	Event *IATIActivitiesSet // Event containing the contract specifics and raw log

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
func (it *IATIActivitiesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IATIActivitiesSet)
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
		it.Event = new(IATIActivitiesSet)
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
func (it *IATIActivitiesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IATIActivitiesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IATIActivitiesSet represents a Set event raised by the IATIActivities contract.
type IATIActivitiesSet struct {
	ContractName string
	Ref          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0x89752cf887704c720fdae38088909ed147c755f39bd110525728448ec4cefc8b.
//
// Solidity: event Set(string contractName, bytes32 _ref)
func (_IATIActivities *IATIActivitiesFilterer) FilterSet(opts *bind.FilterOpts) (*IATIActivitiesSetIterator, error) {

	logs, sub, err := _IATIActivities.contract.FilterLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return &IATIActivitiesSetIterator{contract: _IATIActivities.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0x89752cf887704c720fdae38088909ed147c755f39bd110525728448ec4cefc8b.
//
// Solidity: event Set(string contractName, bytes32 _ref)
func (_IATIActivities *IATIActivitiesFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *IATIActivitiesSet) (event.Subscription, error) {

	logs, sub, err := _IATIActivities.contract.WatchLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IATIActivitiesSet)
				if err := _IATIActivities.contract.UnpackLog(event, "Set", log); err != nil {
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

// ParseSet is a log parse operation binding the contract event 0x89752cf887704c720fdae38088909ed147c755f39bd110525728448ec4cefc8b.
//
// Solidity: event Set(string contractName, bytes32 _ref)
func (_IATIActivities *IATIActivitiesFilterer) ParseSet(log types.Log) (*IATIActivitiesSet, error) {
	event := new(IATIActivitiesSet)
	if err := _IATIActivities.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IDataABI is the input ABI used to generate the binding from.
const IDataABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"contractName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"Set\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getter\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setter\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IDataFuncSigs maps the 4-byte function signature to its string representation.
var IDataFuncSigs = map[string]string{
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
	"993a04b7": "getter()",
	"3f3108f7": "setter()",
}

// IData is an auto generated Go binding around an Ethereum contract.
type IData struct {
	IDataCaller     // Read-only binding to the contract
	IDataTransactor // Write-only binding to the contract
	IDataFilterer   // Log filterer for contract events
}

// IDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDataSession struct {
	Contract     *IData            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDataCallerSession struct {
	Contract *IDataCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDataTransactorSession struct {
	Contract     *IDataTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDataRaw struct {
	Contract *IData // Generic contract binding to access the raw methods on
}

// IDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDataCallerRaw struct {
	Contract *IDataCaller // Generic read-only contract binding to access the raw methods on
}

// IDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDataTransactorRaw struct {
	Contract *IDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIData creates a new instance of IData, bound to a specific deployed contract.
func NewIData(address common.Address, backend bind.ContractBackend) (*IData, error) {
	contract, err := bindIData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IData{IDataCaller: IDataCaller{contract: contract}, IDataTransactor: IDataTransactor{contract: contract}, IDataFilterer: IDataFilterer{contract: contract}}, nil
}

// NewIDataCaller creates a new read-only instance of IData, bound to a specific deployed contract.
func NewIDataCaller(address common.Address, caller bind.ContractCaller) (*IDataCaller, error) {
	contract, err := bindIData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDataCaller{contract: contract}, nil
}

// NewIDataTransactor creates a new write-only instance of IData, bound to a specific deployed contract.
func NewIDataTransactor(address common.Address, transactor bind.ContractTransactor) (*IDataTransactor, error) {
	contract, err := bindIData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDataTransactor{contract: contract}, nil
}

// NewIDataFilterer creates a new log filterer instance of IData, bound to a specific deployed contract.
func NewIDataFilterer(address common.Address, filterer bind.ContractFilterer) (*IDataFilterer, error) {
	contract, err := bindIData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDataFilterer{contract: contract}, nil
}

// bindIData binds a generic wrapper to an already deployed contract.
func bindIData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IData *IDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IData.Contract.IDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IData *IDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IData.Contract.IDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IData *IDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IData.Contract.IDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IData *IDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IData *IDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IData *IDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IData.Contract.contract.Transact(opts, method, params...)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_IData *IDataCaller) GetNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IData.contract.Call(opts, out, "getNum")
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_IData *IDataSession) GetNum() (*big.Int, error) {
	return _IData.Contract.GetNum(&_IData.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_IData *IDataCallerSession) GetNum() (*big.Int, error) {
	return _IData.Contract.GetNum(&_IData.CallOpts)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_IData *IDataCaller) GetReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IData.contract.Call(opts, out, "getReference", _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_IData *IDataSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _IData.Contract.GetReference(&_IData.CallOpts, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_IData *IDataCallerSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _IData.Contract.GetReference(&_IData.CallOpts, _index)
}

// Getter is a free data retrieval call binding the contract method 0x993a04b7.
//
// Solidity: function getter() view returns(bytes4)
func (_IData *IDataCaller) Getter(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _IData.contract.Call(opts, out, "getter")
	return *ret0, err
}

// Getter is a free data retrieval call binding the contract method 0x993a04b7.
//
// Solidity: function getter() view returns(bytes4)
func (_IData *IDataSession) Getter() ([4]byte, error) {
	return _IData.Contract.Getter(&_IData.CallOpts)
}

// Getter is a free data retrieval call binding the contract method 0x993a04b7.
//
// Solidity: function getter() view returns(bytes4)
func (_IData *IDataCallerSession) Getter() ([4]byte, error) {
	return _IData.Contract.Getter(&_IData.CallOpts)
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(bytes4)
func (_IData *IDataCaller) Setter(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _IData.contract.Call(opts, out, "setter")
	return *ret0, err
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(bytes4)
func (_IData *IDataSession) Setter() ([4]byte, error) {
	return _IData.Contract.Setter(&_IData.CallOpts)
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(bytes4)
func (_IData *IDataCallerSession) Setter() ([4]byte, error) {
	return _IData.Contract.Setter(&_IData.CallOpts)
}

// IDataSetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the IData contract.
type IDataSetIterator struct {
	Event *IDataSet // Event containing the contract specifics and raw log

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
func (it *IDataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDataSet)
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
		it.Event = new(IDataSet)
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
func (it *IDataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDataSet represents a Set event raised by the IData contract.
type IDataSet struct {
	ContractName string
	Ref          [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0x89752cf887704c720fdae38088909ed147c755f39bd110525728448ec4cefc8b.
//
// Solidity: event Set(string contractName, bytes32 _ref)
func (_IData *IDataFilterer) FilterSet(opts *bind.FilterOpts) (*IDataSetIterator, error) {

	logs, sub, err := _IData.contract.FilterLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return &IDataSetIterator{contract: _IData.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0x89752cf887704c720fdae38088909ed147c755f39bd110525728448ec4cefc8b.
//
// Solidity: event Set(string contractName, bytes32 _ref)
func (_IData *IDataFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *IDataSet) (event.Subscription, error) {

	logs, sub, err := _IData.contract.WatchLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDataSet)
				if err := _IData.contract.UnpackLog(event, "Set", log); err != nil {
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

// ParseSet is a log parse operation binding the contract event 0x89752cf887704c720fdae38088909ed147c755f39bd110525728448ec4cefc8b.
//
// Solidity: event Set(string contractName, bytes32 _ref)
func (_IData *IDataFilterer) ParseSet(log types.Log) (*IDataSet, error) {
	event := new(IDataSet)
	if err := _IData.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IterableDataABI is the input ABI used to generate the binding from.
const IterableDataABI = "[]"

// IterableDataBin is the compiled bytecode used for deploying new contracts.
var IterableDataBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204712e050da4fa5d3eca663a59ab1b467a9e95c3f95440c8fe45549d6ef13b59664736f6c63430006060033"

// DeployIterableData deploys a new Ethereum contract, binding an instance of IterableData to it.
func DeployIterableData(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IterableData, error) {
	parsed, err := abi.JSON(strings.NewReader(IterableDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IterableDataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IterableData{IterableDataCaller: IterableDataCaller{contract: contract}, IterableDataTransactor: IterableDataTransactor{contract: contract}, IterableDataFilterer: IterableDataFilterer{contract: contract}}, nil
}

// IterableData is an auto generated Go binding around an Ethereum contract.
type IterableData struct {
	IterableDataCaller     // Read-only binding to the contract
	IterableDataTransactor // Write-only binding to the contract
	IterableDataFilterer   // Log filterer for contract events
}

// IterableDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IterableDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IterableDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IterableDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IterableDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IterableDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IterableDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IterableDataSession struct {
	Contract     *IterableData     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IterableDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IterableDataCallerSession struct {
	Contract *IterableDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IterableDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IterableDataTransactorSession struct {
	Contract     *IterableDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IterableDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IterableDataRaw struct {
	Contract *IterableData // Generic contract binding to access the raw methods on
}

// IterableDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IterableDataCallerRaw struct {
	Contract *IterableDataCaller // Generic read-only contract binding to access the raw methods on
}

// IterableDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IterableDataTransactorRaw struct {
	Contract *IterableDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIterableData creates a new instance of IterableData, bound to a specific deployed contract.
func NewIterableData(address common.Address, backend bind.ContractBackend) (*IterableData, error) {
	contract, err := bindIterableData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IterableData{IterableDataCaller: IterableDataCaller{contract: contract}, IterableDataTransactor: IterableDataTransactor{contract: contract}, IterableDataFilterer: IterableDataFilterer{contract: contract}}, nil
}

// NewIterableDataCaller creates a new read-only instance of IterableData, bound to a specific deployed contract.
func NewIterableDataCaller(address common.Address, caller bind.ContractCaller) (*IterableDataCaller, error) {
	contract, err := bindIterableData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IterableDataCaller{contract: contract}, nil
}

// NewIterableDataTransactor creates a new write-only instance of IterableData, bound to a specific deployed contract.
func NewIterableDataTransactor(address common.Address, transactor bind.ContractTransactor) (*IterableDataTransactor, error) {
	contract, err := bindIterableData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IterableDataTransactor{contract: contract}, nil
}

// NewIterableDataFilterer creates a new log filterer instance of IterableData, bound to a specific deployed contract.
func NewIterableDataFilterer(address common.Address, filterer bind.ContractFilterer) (*IterableDataFilterer, error) {
	contract, err := bindIterableData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IterableDataFilterer{contract: contract}, nil
}

// bindIterableData binds a generic wrapper to an already deployed contract.
func bindIterableData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IterableDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IterableData *IterableDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IterableData.Contract.IterableDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IterableData *IterableDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IterableData.Contract.IterableDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IterableData *IterableDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IterableData.Contract.IterableDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IterableData *IterableDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IterableData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IterableData *IterableDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IterableData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IterableData *IterableDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IterableData.Contract.contract.Transact(opts, method, params...)
}
