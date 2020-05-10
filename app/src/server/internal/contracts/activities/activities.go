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

// ActivitiesABI is the input ABI used to generate the binding from.
const ActivitiesABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_child\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_childType\",\"type\":\"uint8\"}],\"name\":\"AddChild\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_elementType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"Set\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_child\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_childType\",\"type\":\"uint8\"}],\"name\":\"addChild\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"getFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"}],\"internalType\":\"structActivitiesData\",\"name\":\"_activities\",\"type\":\"tuple\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setter\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ActivitiesFuncSigs maps the 4-byte function signature to its string representation.
var ActivitiesFuncSigs = map[string]string{
	"cbb2dc95": "addChild(bytes32,address,uint8)",
	"8eaa6ac0": "get(bytes32)",
	"17fc2761": "getFactory(bytes32)",
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
	"91a95fbc": "set(bytes32,(bytes32,bytes32,bytes32))",
	"3f3108f7": "setter()",
}

// ActivitiesBin is the compiled bytecode used for deploying new contracts.
var ActivitiesBin = "0x608060405234801561001057600080fd5b506000600255610a8c806100256000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638eaa6ac01161005b5780638eaa6ac0146100d557806391a95fbc146100e8578063bc599341146100fd578063cbb2dc95146101105761007d565b806317fc2761146100825780633f3108f7146100ab57806367e0badb146100c0575b600080fd5b6100956100903660046104b6565b610123565b6040516100a29190610592565b60405180910390f35b6100b3610181565b6040516100a291906105d1565b6100c861018c565b6040516100a291906105a6565b6100956100e33660046104b6565b610192565b6100fb6100f636600461051f565b6101f0565b005b6100c861010b3660046104b6565b6102d5565b6100fb61011e3660046104ce565b61030d565b600081811a60f81b6001600160f81b03191661013e57600080fd5b6000828152600360205260409020600101546001600160a01b031661016257600080fd5b506000908152600360205260409020600101546001600160a01b031690565b63246a57ef60e21b90565b60025490565b600081811a60f81b6001600160f81b0319166101ad57600080fd5b6000828152602081905260409020600101546001600160a01b03166101d157600080fd5b506000908152602081905260409020600101546001600160a01b031690565b8160001a60f81b6001600160f81b03191661020a57600080fd5b6000816040516102199061049c565b61022391906105e6565b604051809103906000f08015801561023f573d6000803e3d6000fd5b5090506102546000848363ffffffff61040516565b506000604051610263906104a9565b604051809103906000f08015801561027f573d6000803e3d6000fd5b5090506102946003858363ffffffff61040516565b507f8353a4d574992e9eb676a31cd326ce2f9f6c4829c3b84ae33ee9febbb2962e7b6002856040516102c7929190610607565b60405180910390a150505050565b60025460009082106102e657600080fd5b60018054839081106102f457fe5b9060005260206000209060020201600001549050919050565b8260001a60f81b6001600160f81b03191661032757600080fd5b6001600160a01b03821661033a57600080fd5b6000838152600360205260409020600101546001600160a01b031661035e57600080fd5b60008381526003602052604090819020600101549051630a3b0a4f60e01b81526001600160a01b03909116908190630a3b0a4f906103a0908690600401610592565b600060405180830381600087803b1580156103ba57600080fd5b505af11580156103ce573d6000803e3d6000fd5b505050507f4f2635eaa771447d282168d4763629c4b855ff7c733898eefc20b8699278d0c08484846040516102c7939291906105af565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b0385161790558015610443576001915050610495565b506001808501805491820180825560008681526020889052604090208190558591908390811061046f57fe5b600091825260208220600291820201929092559086018054600101905591506104959050565b9392505050565b61024b8061061b83390190565b6101f18061086683390190565b6000602082840312156104c7578081fd5b5035919050565b6000806000606084860312156104e2578182fd5b8335925060208401356001600160a01b03811681146104ff578283fd5b9150604084013560ff81168114610514578182fd5b809150509250925092565b6000808284036080811215610532578283fd5b833592506060601f1982011215610547578182fd5b506040516060810181811067ffffffffffffffff82111715610567578283fd5b8060405250602084013581526040840135602082015260608401356040820152809150509250929050565b6001600160a01b0391909116815260200190565b90815260200190565b9283526001600160a01b0391909116602083015260ff16604082015260600190565b6001600160e01b031991909116815260200190565b81518152602080830151908201526040918201519181019190915260600190565b60ff92909216825260208201526040019056fe608060405234801561001057600080fd5b5060405161024b38038061024b83398101604081905261002f916100a3565b805160001a60f81b6001600160f81b031916158015906100625750602081015160001a60f81b6001600160f81b03191615155b80156100815750604081015160001a60f81b6001600160f81b03191615155b61008a57600080fd5b80516000556020810151600155604001516002556100f9565b6000606082840312156100b4578081fd5b604051606081016001600160401b03811182821017156100d2578283fd5b80604052508251815260208301516020820152604083015160408201528091505092915050565b610143806101086000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636d4ce63c1461003b578063993a04b71461005a575b600080fd5b61004361006f565b6040516100519291906100e1565b60405180910390f35b6100626100a1565b60405161005191906100cc565b60006100796100ac565b5050604080516060810182526000548152600154602082015260028054928201929092529091565b631b53398f60e21b90565b604080516060810182526000808252602082018190529181019190915290565b6001600160e01b031991909116815260200190565b60ff9290921682528051602080840191909152810151604080840191909152015160608201526080019056fea2646970667358221220d395d41b4cf35e254350b710be4bc03caa8e28683ed482396e7916951afb520664736f6c63430006060033608060405234801561001057600080fd5b506101d1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630a3b0a4f146100465780634f0f4aa91461005b57806367e0badb14610084575b600080fd5b610059610054366004610138565b610099565b005b61006e610069366004610166565b6100fb565b60405161007b919061017e565b60405180910390f35b61008c610132565b60405161007b9190610192565b6001600160a01b0381166100ac57600080fd5b600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630180546001600160a01b0319166001600160a01b0392909216919091179055565b60008054821061010a57600080fd5b6000828154811061011757fe5b6000918252602090912001546001600160a01b031692915050565b60005490565b600060208284031215610149578081fd5b81356001600160a01b038116811461015f578182fd5b9392505050565b600060208284031215610177578081fd5b5035919050565b6001600160a01b0391909116815260200190565b9081526020019056fea26469706673582212200b1d0f60a813a64e8e7204f784c414c6a9bd1c546926bc9e4be49f667646b48864736f6c63430006060033a26469706673582212207ef59982ab9e1fababe4ad61ec809b40e3c40a9d6296030238c7cd83f18bdd6364736f6c63430006060033"

// DeployActivities deploys a new Ethereum contract, binding an instance of Activities to it.
func DeployActivities(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Activities, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivitiesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ActivitiesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Activities{ActivitiesCaller: ActivitiesCaller{contract: contract}, ActivitiesTransactor: ActivitiesTransactor{contract: contract}, ActivitiesFilterer: ActivitiesFilterer{contract: contract}}, nil
}

// Activities is an auto generated Go binding around an Ethereum contract.
type Activities struct {
	ActivitiesCaller     // Read-only binding to the contract
	ActivitiesTransactor // Write-only binding to the contract
	ActivitiesFilterer   // Log filterer for contract events
}

// ActivitiesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivitiesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivitiesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivitiesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivitiesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivitiesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivitiesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivitiesSession struct {
	Contract     *Activities       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ActivitiesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivitiesCallerSession struct {
	Contract *ActivitiesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ActivitiesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivitiesTransactorSession struct {
	Contract     *ActivitiesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ActivitiesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivitiesRaw struct {
	Contract *Activities // Generic contract binding to access the raw methods on
}

// ActivitiesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivitiesCallerRaw struct {
	Contract *ActivitiesCaller // Generic read-only contract binding to access the raw methods on
}

// ActivitiesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivitiesTransactorRaw struct {
	Contract *ActivitiesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivities creates a new instance of Activities, bound to a specific deployed contract.
func NewActivities(address common.Address, backend bind.ContractBackend) (*Activities, error) {
	contract, err := bindActivities(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Activities{ActivitiesCaller: ActivitiesCaller{contract: contract}, ActivitiesTransactor: ActivitiesTransactor{contract: contract}, ActivitiesFilterer: ActivitiesFilterer{contract: contract}}, nil
}

// NewActivitiesCaller creates a new read-only instance of Activities, bound to a specific deployed contract.
func NewActivitiesCaller(address common.Address, caller bind.ContractCaller) (*ActivitiesCaller, error) {
	contract, err := bindActivities(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivitiesCaller{contract: contract}, nil
}

// NewActivitiesTransactor creates a new write-only instance of Activities, bound to a specific deployed contract.
func NewActivitiesTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivitiesTransactor, error) {
	contract, err := bindActivities(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivitiesTransactor{contract: contract}, nil
}

// NewActivitiesFilterer creates a new log filterer instance of Activities, bound to a specific deployed contract.
func NewActivitiesFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivitiesFilterer, error) {
	contract, err := bindActivities(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivitiesFilterer{contract: contract}, nil
}

// bindActivities binds a generic wrapper to an already deployed contract.
func bindActivities(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivitiesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Activities *ActivitiesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Activities.Contract.ActivitiesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Activities *ActivitiesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activities.Contract.ActivitiesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Activities *ActivitiesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Activities.Contract.ActivitiesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Activities *ActivitiesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Activities.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Activities *ActivitiesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activities.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Activities *ActivitiesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Activities.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _ref) view returns(address)
func (_Activities *ActivitiesCaller) Get(opts *bind.CallOpts, _ref [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Activities.contract.Call(opts, out, "get", _ref)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _ref) view returns(address)
func (_Activities *ActivitiesSession) Get(_ref [32]byte) (common.Address, error) {
	return _Activities.Contract.Get(&_Activities.CallOpts, _ref)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _ref) view returns(address)
func (_Activities *ActivitiesCallerSession) Get(_ref [32]byte) (common.Address, error) {
	return _Activities.Contract.Get(&_Activities.CallOpts, _ref)
}

// GetFactory is a free data retrieval call binding the contract method 0x17fc2761.
//
// Solidity: function getFactory(bytes32 _ref) view returns(address)
func (_Activities *ActivitiesCaller) GetFactory(opts *bind.CallOpts, _ref [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Activities.contract.Call(opts, out, "getFactory", _ref)
	return *ret0, err
}

// GetFactory is a free data retrieval call binding the contract method 0x17fc2761.
//
// Solidity: function getFactory(bytes32 _ref) view returns(address)
func (_Activities *ActivitiesSession) GetFactory(_ref [32]byte) (common.Address, error) {
	return _Activities.Contract.GetFactory(&_Activities.CallOpts, _ref)
}

// GetFactory is a free data retrieval call binding the contract method 0x17fc2761.
//
// Solidity: function getFactory(bytes32 _ref) view returns(address)
func (_Activities *ActivitiesCallerSession) GetFactory(_ref [32]byte) (common.Address, error) {
	return _Activities.Contract.GetFactory(&_Activities.CallOpts, _ref)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Activities *ActivitiesCaller) GetNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Activities.contract.Call(opts, out, "getNum")
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Activities *ActivitiesSession) GetNum() (*big.Int, error) {
	return _Activities.Contract.GetNum(&_Activities.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Activities *ActivitiesCallerSession) GetNum() (*big.Int, error) {
	return _Activities.Contract.GetNum(&_Activities.CallOpts)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_Activities *ActivitiesCaller) GetReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Activities.contract.Call(opts, out, "getReference", _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_Activities *ActivitiesSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _Activities.Contract.GetReference(&_Activities.CallOpts, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_Activities *ActivitiesCallerSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _Activities.Contract.GetReference(&_Activities.CallOpts, _index)
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(bytes4)
func (_Activities *ActivitiesCaller) Setter(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _Activities.contract.Call(opts, out, "setter")
	return *ret0, err
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(bytes4)
func (_Activities *ActivitiesSession) Setter() ([4]byte, error) {
	return _Activities.Contract.Setter(&_Activities.CallOpts)
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(bytes4)
func (_Activities *ActivitiesCallerSession) Setter() ([4]byte, error) {
	return _Activities.Contract.Setter(&_Activities.CallOpts)
}

// AddChild is a paid mutator transaction binding the contract method 0xcbb2dc95.
//
// Solidity: function addChild(bytes32 _ref, address _child, uint8 _childType) returns()
func (_Activities *ActivitiesTransactor) AddChild(opts *bind.TransactOpts, _ref [32]byte, _child common.Address, _childType uint8) (*types.Transaction, error) {
	return _Activities.contract.Transact(opts, "addChild", _ref, _child, _childType)
}

// AddChild is a paid mutator transaction binding the contract method 0xcbb2dc95.
//
// Solidity: function addChild(bytes32 _ref, address _child, uint8 _childType) returns()
func (_Activities *ActivitiesSession) AddChild(_ref [32]byte, _child common.Address, _childType uint8) (*types.Transaction, error) {
	return _Activities.Contract.AddChild(&_Activities.TransactOpts, _ref, _child, _childType)
}

// AddChild is a paid mutator transaction binding the contract method 0xcbb2dc95.
//
// Solidity: function addChild(bytes32 _ref, address _child, uint8 _childType) returns()
func (_Activities *ActivitiesTransactorSession) AddChild(_ref [32]byte, _child common.Address, _childType uint8) (*types.Transaction, error) {
	return _Activities.Contract.AddChild(&_Activities.TransactOpts, _ref, _child, _childType)
}

// Set is a paid mutator transaction binding the contract method 0x91a95fbc.
//
// Solidity: function set(bytes32 _ref, ActivitiesData _activities) returns()
func (_Activities *ActivitiesTransactor) Set(opts *bind.TransactOpts, _ref [32]byte, _activities ActivitiesData) (*types.Transaction, error) {
	return _Activities.contract.Transact(opts, "set", _ref, _activities)
}

// Set is a paid mutator transaction binding the contract method 0x91a95fbc.
//
// Solidity: function set(bytes32 _ref, ActivitiesData _activities) returns()
func (_Activities *ActivitiesSession) Set(_ref [32]byte, _activities ActivitiesData) (*types.Transaction, error) {
	return _Activities.Contract.Set(&_Activities.TransactOpts, _ref, _activities)
}

// Set is a paid mutator transaction binding the contract method 0x91a95fbc.
//
// Solidity: function set(bytes32 _ref, ActivitiesData _activities) returns()
func (_Activities *ActivitiesTransactorSession) Set(_ref [32]byte, _activities ActivitiesData) (*types.Transaction, error) {
	return _Activities.Contract.Set(&_Activities.TransactOpts, _ref, _activities)
}

// ActivitiesAddChildIterator is returned from FilterAddChild and is used to iterate over the raw logs and unpacked data for AddChild events raised by the Activities contract.
type ActivitiesAddChildIterator struct {
	Event *ActivitiesAddChild // Event containing the contract specifics and raw log

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
func (it *ActivitiesAddChildIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivitiesAddChild)
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
		it.Event = new(ActivitiesAddChild)
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
func (it *ActivitiesAddChildIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivitiesAddChildIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivitiesAddChild represents a AddChild event raised by the Activities contract.
type ActivitiesAddChild struct {
	Ref       [32]byte
	Child     common.Address
	ChildType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddChild is a free log retrieval operation binding the contract event 0x4f2635eaa771447d282168d4763629c4b855ff7c733898eefc20b8699278d0c0.
//
// Solidity: event AddChild(bytes32 _ref, address _child, uint8 _childType)
func (_Activities *ActivitiesFilterer) FilterAddChild(opts *bind.FilterOpts) (*ActivitiesAddChildIterator, error) {

	logs, sub, err := _Activities.contract.FilterLogs(opts, "AddChild")
	if err != nil {
		return nil, err
	}
	return &ActivitiesAddChildIterator{contract: _Activities.contract, event: "AddChild", logs: logs, sub: sub}, nil
}

// WatchAddChild is a free log subscription operation binding the contract event 0x4f2635eaa771447d282168d4763629c4b855ff7c733898eefc20b8699278d0c0.
//
// Solidity: event AddChild(bytes32 _ref, address _child, uint8 _childType)
func (_Activities *ActivitiesFilterer) WatchAddChild(opts *bind.WatchOpts, sink chan<- *ActivitiesAddChild) (event.Subscription, error) {

	logs, sub, err := _Activities.contract.WatchLogs(opts, "AddChild")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivitiesAddChild)
				if err := _Activities.contract.UnpackLog(event, "AddChild", log); err != nil {
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

// ParseAddChild is a log parse operation binding the contract event 0x4f2635eaa771447d282168d4763629c4b855ff7c733898eefc20b8699278d0c0.
//
// Solidity: event AddChild(bytes32 _ref, address _child, uint8 _childType)
func (_Activities *ActivitiesFilterer) ParseAddChild(log types.Log) (*ActivitiesAddChild, error) {
	event := new(ActivitiesAddChild)
	if err := _Activities.contract.UnpackLog(event, "AddChild", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ActivitiesSetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the Activities contract.
type ActivitiesSetIterator struct {
	Event *ActivitiesSet // Event containing the contract specifics and raw log

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
func (it *ActivitiesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivitiesSet)
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
		it.Event = new(ActivitiesSet)
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
func (it *ActivitiesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivitiesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivitiesSet represents a Set event raised by the Activities contract.
type ActivitiesSet struct {
	ElementType uint8
	Ref         [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0x8353a4d574992e9eb676a31cd326ce2f9f6c4829c3b84ae33ee9febbb2962e7b.
//
// Solidity: event Set(uint8 _elementType, bytes32 _ref)
func (_Activities *ActivitiesFilterer) FilterSet(opts *bind.FilterOpts) (*ActivitiesSetIterator, error) {

	logs, sub, err := _Activities.contract.FilterLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return &ActivitiesSetIterator{contract: _Activities.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0x8353a4d574992e9eb676a31cd326ce2f9f6c4829c3b84ae33ee9febbb2962e7b.
//
// Solidity: event Set(uint8 _elementType, bytes32 _ref)
func (_Activities *ActivitiesFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *ActivitiesSet) (event.Subscription, error) {

	logs, sub, err := _Activities.contract.WatchLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivitiesSet)
				if err := _Activities.contract.UnpackLog(event, "Set", log); err != nil {
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

// ParseSet is a log parse operation binding the contract event 0x8353a4d574992e9eb676a31cd326ce2f9f6c4829c3b84ae33ee9febbb2962e7b.
//
// Solidity: event Set(uint8 _elementType, bytes32 _ref)
func (_Activities *ActivitiesFilterer) ParseSet(log types.Log) (*ActivitiesSet, error) {
	event := new(ActivitiesSet)
	if err := _Activities.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ActivitiesNodeABI is the input ABI used to generate the binding from.
const ActivitiesNodeABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"}],\"internalType\":\"structActivitiesData\",\"name\":\"_activities\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"}],\"internalType\":\"structActivitiesData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getter\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ActivitiesNodeFuncSigs maps the 4-byte function signature to its string representation.
var ActivitiesNodeFuncSigs = map[string]string{
	"6d4ce63c": "get()",
	"993a04b7": "getter()",
}

// ActivitiesNodeBin is the compiled bytecode used for deploying new contracts.
var ActivitiesNodeBin = "0x608060405234801561001057600080fd5b5060405161024b38038061024b83398101604081905261002f916100a3565b805160001a60f81b6001600160f81b031916158015906100625750602081015160001a60f81b6001600160f81b03191615155b80156100815750604081015160001a60f81b6001600160f81b03191615155b61008a57600080fd5b80516000556020810151600155604001516002556100f9565b6000606082840312156100b4578081fd5b604051606081016001600160401b03811182821017156100d2578283fd5b80604052508251815260208301516020820152604083015160408201528091505092915050565b610143806101086000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636d4ce63c1461003b578063993a04b71461005a575b600080fd5b61004361006f565b6040516100519291906100e1565b60405180910390f35b6100626100a1565b60405161005191906100cc565b60006100796100ac565b5050604080516060810182526000548152600154602082015260028054928201929092529091565b631b53398f60e21b90565b604080516060810182526000808252602082018190529181019190915290565b6001600160e01b031991909116815260200190565b60ff9290921682528051602080840191909152810151604080840191909152015160608201526080019056fea2646970667358221220d395d41b4cf35e254350b710be4bc03caa8e28683ed482396e7916951afb520664736f6c63430006060033"

// DeployActivitiesNode deploys a new Ethereum contract, binding an instance of ActivitiesNode to it.
func DeployActivitiesNode(auth *bind.TransactOpts, backend bind.ContractBackend, _activities ActivitiesData) (common.Address, *types.Transaction, *ActivitiesNode, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivitiesNodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ActivitiesNodeBin), backend, _activities)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ActivitiesNode{ActivitiesNodeCaller: ActivitiesNodeCaller{contract: contract}, ActivitiesNodeTransactor: ActivitiesNodeTransactor{contract: contract}, ActivitiesNodeFilterer: ActivitiesNodeFilterer{contract: contract}}, nil
}

// ActivitiesNode is an auto generated Go binding around an Ethereum contract.
type ActivitiesNode struct {
	ActivitiesNodeCaller     // Read-only binding to the contract
	ActivitiesNodeTransactor // Write-only binding to the contract
	ActivitiesNodeFilterer   // Log filterer for contract events
}

// ActivitiesNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivitiesNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivitiesNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivitiesNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivitiesNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivitiesNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivitiesNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivitiesNodeSession struct {
	Contract     *ActivitiesNode   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ActivitiesNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivitiesNodeCallerSession struct {
	Contract *ActivitiesNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ActivitiesNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivitiesNodeTransactorSession struct {
	Contract     *ActivitiesNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ActivitiesNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivitiesNodeRaw struct {
	Contract *ActivitiesNode // Generic contract binding to access the raw methods on
}

// ActivitiesNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivitiesNodeCallerRaw struct {
	Contract *ActivitiesNodeCaller // Generic read-only contract binding to access the raw methods on
}

// ActivitiesNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivitiesNodeTransactorRaw struct {
	Contract *ActivitiesNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivitiesNode creates a new instance of ActivitiesNode, bound to a specific deployed contract.
func NewActivitiesNode(address common.Address, backend bind.ContractBackend) (*ActivitiesNode, error) {
	contract, err := bindActivitiesNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ActivitiesNode{ActivitiesNodeCaller: ActivitiesNodeCaller{contract: contract}, ActivitiesNodeTransactor: ActivitiesNodeTransactor{contract: contract}, ActivitiesNodeFilterer: ActivitiesNodeFilterer{contract: contract}}, nil
}

// NewActivitiesNodeCaller creates a new read-only instance of ActivitiesNode, bound to a specific deployed contract.
func NewActivitiesNodeCaller(address common.Address, caller bind.ContractCaller) (*ActivitiesNodeCaller, error) {
	contract, err := bindActivitiesNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivitiesNodeCaller{contract: contract}, nil
}

// NewActivitiesNodeTransactor creates a new write-only instance of ActivitiesNode, bound to a specific deployed contract.
func NewActivitiesNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivitiesNodeTransactor, error) {
	contract, err := bindActivitiesNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivitiesNodeTransactor{contract: contract}, nil
}

// NewActivitiesNodeFilterer creates a new log filterer instance of ActivitiesNode, bound to a specific deployed contract.
func NewActivitiesNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivitiesNodeFilterer, error) {
	contract, err := bindActivitiesNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivitiesNodeFilterer{contract: contract}, nil
}

// bindActivitiesNode binds a generic wrapper to an already deployed contract.
func bindActivitiesNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivitiesNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivitiesNode *ActivitiesNodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivitiesNode.Contract.ActivitiesNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivitiesNode *ActivitiesNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivitiesNode.Contract.ActivitiesNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivitiesNode *ActivitiesNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivitiesNode.Contract.ActivitiesNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivitiesNode *ActivitiesNodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivitiesNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivitiesNode *ActivitiesNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivitiesNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivitiesNode *ActivitiesNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivitiesNode.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint8, ActivitiesData)
func (_ActivitiesNode *ActivitiesNodeCaller) Get(opts *bind.CallOpts) (uint8, ActivitiesData, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(ActivitiesData)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ActivitiesNode.contract.Call(opts, out, "get")
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint8, ActivitiesData)
func (_ActivitiesNode *ActivitiesNodeSession) Get() (uint8, ActivitiesData, error) {
	return _ActivitiesNode.Contract.Get(&_ActivitiesNode.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint8, ActivitiesData)
func (_ActivitiesNode *ActivitiesNodeCallerSession) Get() (uint8, ActivitiesData, error) {
	return _ActivitiesNode.Contract.Get(&_ActivitiesNode.CallOpts)
}

// Getter is a free data retrieval call binding the contract method 0x993a04b7.
//
// Solidity: function getter() view returns(bytes4)
func (_ActivitiesNode *ActivitiesNodeCaller) Getter(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _ActivitiesNode.contract.Call(opts, out, "getter")
	return *ret0, err
}

// Getter is a free data retrieval call binding the contract method 0x993a04b7.
//
// Solidity: function getter() view returns(bytes4)
func (_ActivitiesNode *ActivitiesNodeSession) Getter() ([4]byte, error) {
	return _ActivitiesNode.Contract.Getter(&_ActivitiesNode.CallOpts)
}

// Getter is a free data retrieval call binding the contract method 0x993a04b7.
//
// Solidity: function getter() view returns(bytes4)
func (_ActivitiesNode *ActivitiesNodeCallerSession) Getter() ([4]byte, error) {
	return _ActivitiesNode.Contract.Getter(&_ActivitiesNode.CallOpts)
}

// FactoryABI is the input ABI used to generate the binding from.
const FactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FactoryFuncSigs maps the 4-byte function signature to its string representation.
var FactoryFuncSigs = map[string]string{
	"0a3b0a4f": "add(address)",
	"4f0f4aa9": "getNode(uint256)",
	"67e0badb": "getNum()",
}

// FactoryBin is the compiled bytecode used for deploying new contracts.
var FactoryBin = "0x608060405234801561001057600080fd5b506101d1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630a3b0a4f146100465780634f0f4aa91461005b57806367e0badb14610084575b600080fd5b610059610054366004610138565b610099565b005b61006e610069366004610166565b6100fb565b60405161007b919061017e565b60405180910390f35b61008c610132565b60405161007b9190610192565b6001600160a01b0381166100ac57600080fd5b600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630180546001600160a01b0319166001600160a01b0392909216919091179055565b60008054821061010a57600080fd5b6000828154811061011757fe5b6000918252602090912001546001600160a01b031692915050565b60005490565b600060208284031215610149578081fd5b81356001600160a01b038116811461015f578182fd5b9392505050565b600060208284031215610177578081fd5b5035919050565b6001600160a01b0391909116815260200190565b9081526020019056fea26469706673582212200b1d0f60a813a64e8e7204f784c414c6a9bd1c546926bc9e4be49f667646b48864736f6c63430006060033"

// DeployFactory deploys a new Ethereum contract, binding an instance of Factory to it.
func DeployFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Factory, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller     // Read-only binding to the contract
	FactoryTransactor // Write-only binding to the contract
	FactoryFilterer   // Log filterer for contract events
}

// FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorySession struct {
	Contract     *Factory          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryCallerSession struct {
	Contract *FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryTransactorSession struct {
	Contract     *FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryRaw struct {
	Contract *Factory // Generic contract binding to access the raw methods on
}

// FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryCallerRaw struct {
	Contract *FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryTransactorRaw struct {
	Contract *FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	contract, err := bindFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// NewFactoryCaller creates a new read-only instance of Factory, bound to a specific deployed contract.
func NewFactoryCaller(address common.Address, caller bind.ContractCaller) (*FactoryCaller, error) {
	contract, err := bindFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryCaller{contract: contract}, nil
}

// NewFactoryTransactor creates a new write-only instance of Factory, bound to a specific deployed contract.
func NewFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTransactor, error) {
	contract, err := bindFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTransactor{contract: contract}, nil
}

// NewFactoryFilterer creates a new log filterer instance of Factory, bound to a specific deployed contract.
func NewFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryFilterer, error) {
	contract, err := bindFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryFilterer{contract: contract}, nil
}

// bindFactory binds a generic wrapper to an already deployed contract.
func bindFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transact(opts, method, params...)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 _index) view returns(address)
func (_Factory *FactoryCaller) GetNode(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "getNode", _index)
	return *ret0, err
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 _index) view returns(address)
func (_Factory *FactorySession) GetNode(_index *big.Int) (common.Address, error) {
	return _Factory.Contract.GetNode(&_Factory.CallOpts, _index)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 _index) view returns(address)
func (_Factory *FactoryCallerSession) GetNode(_index *big.Int) (common.Address, error) {
	return _Factory.Contract.GetNode(&_Factory.CallOpts, _index)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Factory *FactoryCaller) GetNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Factory.contract.Call(opts, out, "getNum")
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Factory *FactorySession) GetNum() (*big.Int, error) {
	return _Factory.Contract.GetNum(&_Factory.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_Factory *FactoryCallerSession) GetNum() (*big.Int, error) {
	return _Factory.Contract.GetNum(&_Factory.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address _node) returns()
func (_Factory *FactoryTransactor) Add(opts *bind.TransactOpts, _node common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "add", _node)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address _node) returns()
func (_Factory *FactorySession) Add(_node common.Address) (*types.Transaction, error) {
	return _Factory.Contract.Add(&_Factory.TransactOpts, _node)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address _node) returns()
func (_Factory *FactoryTransactorSession) Add(_node common.Address) (*types.Transaction, error) {
	return _Factory.Contract.Add(&_Factory.TransactOpts, _node)
}

// IDataABI is the input ABI used to generate the binding from.
const IDataABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_elementType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"Set\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setter\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IDataFuncSigs maps the 4-byte function signature to its string representation.
var IDataFuncSigs = map[string]string{
	"8eaa6ac0": "get(bytes32)",
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
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

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _ref) view returns(address)
func (_IData *IDataCaller) Get(opts *bind.CallOpts, _ref [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IData.contract.Call(opts, out, "get", _ref)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _ref) view returns(address)
func (_IData *IDataSession) Get(_ref [32]byte) (common.Address, error) {
	return _IData.Contract.Get(&_IData.CallOpts, _ref)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _ref) view returns(address)
func (_IData *IDataCallerSession) Get(_ref [32]byte) (common.Address, error) {
	return _IData.Contract.Get(&_IData.CallOpts, _ref)
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
	ElementType uint8
	Ref         [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0x8353a4d574992e9eb676a31cd326ce2f9f6c4829c3b84ae33ee9febbb2962e7b.
//
// Solidity: event Set(uint8 _elementType, bytes32 _ref)
func (_IData *IDataFilterer) FilterSet(opts *bind.FilterOpts) (*IDataSetIterator, error) {

	logs, sub, err := _IData.contract.FilterLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return &IDataSetIterator{contract: _IData.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0x8353a4d574992e9eb676a31cd326ce2f9f6c4829c3b84ae33ee9febbb2962e7b.
//
// Solidity: event Set(uint8 _elementType, bytes32 _ref)
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

// ParseSet is a log parse operation binding the contract event 0x8353a4d574992e9eb676a31cd326ce2f9f6c4829c3b84ae33ee9febbb2962e7b.
//
// Solidity: event Set(uint8 _elementType, bytes32 _ref)
func (_IData *IDataFilterer) ParseSet(log types.Log) (*IDataSet, error) {
	event := new(IDataSet)
	if err := _IData.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IFactoryABI is the input ABI used to generate the binding from.
const IFactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_node\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IFactoryFuncSigs maps the 4-byte function signature to its string representation.
var IFactoryFuncSigs = map[string]string{
	"0a3b0a4f": "add(address)",
	"4f0f4aa9": "getNode(uint256)",
	"67e0badb": "getNum()",
}

// IFactory is an auto generated Go binding around an Ethereum contract.
type IFactory struct {
	IFactoryCaller     // Read-only binding to the contract
	IFactoryTransactor // Write-only binding to the contract
	IFactoryFilterer   // Log filterer for contract events
}

// IFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFactorySession struct {
	Contract     *IFactory         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFactoryCallerSession struct {
	Contract *IFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFactoryTransactorSession struct {
	Contract     *IFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFactoryRaw struct {
	Contract *IFactory // Generic contract binding to access the raw methods on
}

// IFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFactoryCallerRaw struct {
	Contract *IFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFactoryTransactorRaw struct {
	Contract *IFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFactory creates a new instance of IFactory, bound to a specific deployed contract.
func NewIFactory(address common.Address, backend bind.ContractBackend) (*IFactory, error) {
	contract, err := bindIFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFactory{IFactoryCaller: IFactoryCaller{contract: contract}, IFactoryTransactor: IFactoryTransactor{contract: contract}, IFactoryFilterer: IFactoryFilterer{contract: contract}}, nil
}

// NewIFactoryCaller creates a new read-only instance of IFactory, bound to a specific deployed contract.
func NewIFactoryCaller(address common.Address, caller bind.ContractCaller) (*IFactoryCaller, error) {
	contract, err := bindIFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFactoryCaller{contract: contract}, nil
}

// NewIFactoryTransactor creates a new write-only instance of IFactory, bound to a specific deployed contract.
func NewIFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IFactoryTransactor, error) {
	contract, err := bindIFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFactoryTransactor{contract: contract}, nil
}

// NewIFactoryFilterer creates a new log filterer instance of IFactory, bound to a specific deployed contract.
func NewIFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IFactoryFilterer, error) {
	contract, err := bindIFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFactoryFilterer{contract: contract}, nil
}

// bindIFactory binds a generic wrapper to an already deployed contract.
func bindIFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFactory *IFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IFactory.Contract.IFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFactory *IFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFactory.Contract.IFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFactory *IFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFactory.Contract.IFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFactory *IFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFactory *IFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFactory *IFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFactory.Contract.contract.Transact(opts, method, params...)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 _index) view returns(address)
func (_IFactory *IFactoryCaller) GetNode(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IFactory.contract.Call(opts, out, "getNode", _index)
	return *ret0, err
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 _index) view returns(address)
func (_IFactory *IFactorySession) GetNode(_index *big.Int) (common.Address, error) {
	return _IFactory.Contract.GetNode(&_IFactory.CallOpts, _index)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 _index) view returns(address)
func (_IFactory *IFactoryCallerSession) GetNode(_index *big.Int) (common.Address, error) {
	return _IFactory.Contract.GetNode(&_IFactory.CallOpts, _index)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_IFactory *IFactoryCaller) GetNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IFactory.contract.Call(opts, out, "getNum")
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_IFactory *IFactorySession) GetNum() (*big.Int, error) {
	return _IFactory.Contract.GetNum(&_IFactory.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_IFactory *IFactoryCallerSession) GetNum() (*big.Int, error) {
	return _IFactory.Contract.GetNum(&_IFactory.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address _node) returns()
func (_IFactory *IFactoryTransactor) Add(opts *bind.TransactOpts, _node common.Address) (*types.Transaction, error) {
	return _IFactory.contract.Transact(opts, "add", _node)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address _node) returns()
func (_IFactory *IFactorySession) Add(_node common.Address) (*types.Transaction, error) {
	return _IFactory.Contract.Add(&_IFactory.TransactOpts, _node)
}

// Add is a paid mutator transaction binding the contract method 0x0a3b0a4f.
//
// Solidity: function add(address _node) returns()
func (_IFactory *IFactoryTransactorSession) Add(_node common.Address) (*types.Transaction, error) {
	return _IFactory.Contract.Add(&_IFactory.TransactOpts, _node)
}

// INodeABI is the input ABI used to generate the binding from.
const INodeABI = "[{\"inputs\":[],\"name\":\"getter\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// INodeFuncSigs maps the 4-byte function signature to its string representation.
var INodeFuncSigs = map[string]string{
	"993a04b7": "getter()",
}

// INode is an auto generated Go binding around an Ethereum contract.
type INode struct {
	INodeCaller     // Read-only binding to the contract
	INodeTransactor // Write-only binding to the contract
	INodeFilterer   // Log filterer for contract events
}

// INodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type INodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INodeSession struct {
	Contract     *INode            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INodeCallerSession struct {
	Contract *INodeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// INodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INodeTransactorSession struct {
	Contract     *INodeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// INodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type INodeRaw struct {
	Contract *INode // Generic contract binding to access the raw methods on
}

// INodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INodeCallerRaw struct {
	Contract *INodeCaller // Generic read-only contract binding to access the raw methods on
}

// INodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INodeTransactorRaw struct {
	Contract *INodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINode creates a new instance of INode, bound to a specific deployed contract.
func NewINode(address common.Address, backend bind.ContractBackend) (*INode, error) {
	contract, err := bindINode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INode{INodeCaller: INodeCaller{contract: contract}, INodeTransactor: INodeTransactor{contract: contract}, INodeFilterer: INodeFilterer{contract: contract}}, nil
}

// NewINodeCaller creates a new read-only instance of INode, bound to a specific deployed contract.
func NewINodeCaller(address common.Address, caller bind.ContractCaller) (*INodeCaller, error) {
	contract, err := bindINode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INodeCaller{contract: contract}, nil
}

// NewINodeTransactor creates a new write-only instance of INode, bound to a specific deployed contract.
func NewINodeTransactor(address common.Address, transactor bind.ContractTransactor) (*INodeTransactor, error) {
	contract, err := bindINode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INodeTransactor{contract: contract}, nil
}

// NewINodeFilterer creates a new log filterer instance of INode, bound to a specific deployed contract.
func NewINodeFilterer(address common.Address, filterer bind.ContractFilterer) (*INodeFilterer, error) {
	contract, err := bindINode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INodeFilterer{contract: contract}, nil
}

// bindINode binds a generic wrapper to an already deployed contract.
func bindINode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(INodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INode *INodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _INode.Contract.INodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INode *INodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INode.Contract.INodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INode *INodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INode.Contract.INodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INode *INodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _INode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INode *INodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INode *INodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INode.Contract.contract.Transact(opts, method, params...)
}

// Getter is a free data retrieval call binding the contract method 0x993a04b7.
//
// Solidity: function getter() view returns(bytes4)
func (_INode *INodeCaller) Getter(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _INode.contract.Call(opts, out, "getter")
	return *ret0, err
}

// Getter is a free data retrieval call binding the contract method 0x993a04b7.
//
// Solidity: function getter() view returns(bytes4)
func (_INode *INodeSession) Getter() ([4]byte, error) {
	return _INode.Contract.Getter(&_INode.CallOpts)
}

// Getter is a free data retrieval call binding the contract method 0x993a04b7.
//
// Solidity: function getter() view returns(bytes4)
func (_INode *INodeCallerSession) Getter() ([4]byte, error) {
	return _INode.Contract.Getter(&_INode.CallOpts)
}

// ITreeABI is the input ABI used to generate the binding from.
const ITreeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_child\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_childType\",\"type\":\"uint8\"}],\"name\":\"AddChild\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_child\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_childType\",\"type\":\"uint8\"}],\"name\":\"addChild\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"getFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ITreeFuncSigs maps the 4-byte function signature to its string representation.
var ITreeFuncSigs = map[string]string{
	"cbb2dc95": "addChild(bytes32,address,uint8)",
	"17fc2761": "getFactory(bytes32)",
}

// ITree is an auto generated Go binding around an Ethereum contract.
type ITree struct {
	ITreeCaller     // Read-only binding to the contract
	ITreeTransactor // Write-only binding to the contract
	ITreeFilterer   // Log filterer for contract events
}

// ITreeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITreeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITreeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITreeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITreeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITreeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITreeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITreeSession struct {
	Contract     *ITree            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITreeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITreeCallerSession struct {
	Contract *ITreeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ITreeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITreeTransactorSession struct {
	Contract     *ITreeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITreeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITreeRaw struct {
	Contract *ITree // Generic contract binding to access the raw methods on
}

// ITreeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITreeCallerRaw struct {
	Contract *ITreeCaller // Generic read-only contract binding to access the raw methods on
}

// ITreeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITreeTransactorRaw struct {
	Contract *ITreeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITree creates a new instance of ITree, bound to a specific deployed contract.
func NewITree(address common.Address, backend bind.ContractBackend) (*ITree, error) {
	contract, err := bindITree(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITree{ITreeCaller: ITreeCaller{contract: contract}, ITreeTransactor: ITreeTransactor{contract: contract}, ITreeFilterer: ITreeFilterer{contract: contract}}, nil
}

// NewITreeCaller creates a new read-only instance of ITree, bound to a specific deployed contract.
func NewITreeCaller(address common.Address, caller bind.ContractCaller) (*ITreeCaller, error) {
	contract, err := bindITree(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITreeCaller{contract: contract}, nil
}

// NewITreeTransactor creates a new write-only instance of ITree, bound to a specific deployed contract.
func NewITreeTransactor(address common.Address, transactor bind.ContractTransactor) (*ITreeTransactor, error) {
	contract, err := bindITree(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITreeTransactor{contract: contract}, nil
}

// NewITreeFilterer creates a new log filterer instance of ITree, bound to a specific deployed contract.
func NewITreeFilterer(address common.Address, filterer bind.ContractFilterer) (*ITreeFilterer, error) {
	contract, err := bindITree(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITreeFilterer{contract: contract}, nil
}

// bindITree binds a generic wrapper to an already deployed contract.
func bindITree(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITreeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITree *ITreeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ITree.Contract.ITreeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITree *ITreeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITree.Contract.ITreeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITree *ITreeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITree.Contract.ITreeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITree *ITreeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ITree.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITree *ITreeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITree.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITree *ITreeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITree.Contract.contract.Transact(opts, method, params...)
}

// GetFactory is a free data retrieval call binding the contract method 0x17fc2761.
//
// Solidity: function getFactory(bytes32 _ref) view returns(address)
func (_ITree *ITreeCaller) GetFactory(opts *bind.CallOpts, _ref [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ITree.contract.Call(opts, out, "getFactory", _ref)
	return *ret0, err
}

// GetFactory is a free data retrieval call binding the contract method 0x17fc2761.
//
// Solidity: function getFactory(bytes32 _ref) view returns(address)
func (_ITree *ITreeSession) GetFactory(_ref [32]byte) (common.Address, error) {
	return _ITree.Contract.GetFactory(&_ITree.CallOpts, _ref)
}

// GetFactory is a free data retrieval call binding the contract method 0x17fc2761.
//
// Solidity: function getFactory(bytes32 _ref) view returns(address)
func (_ITree *ITreeCallerSession) GetFactory(_ref [32]byte) (common.Address, error) {
	return _ITree.Contract.GetFactory(&_ITree.CallOpts, _ref)
}

// AddChild is a paid mutator transaction binding the contract method 0xcbb2dc95.
//
// Solidity: function addChild(bytes32 _ref, address _child, uint8 _childType) returns()
func (_ITree *ITreeTransactor) AddChild(opts *bind.TransactOpts, _ref [32]byte, _child common.Address, _childType uint8) (*types.Transaction, error) {
	return _ITree.contract.Transact(opts, "addChild", _ref, _child, _childType)
}

// AddChild is a paid mutator transaction binding the contract method 0xcbb2dc95.
//
// Solidity: function addChild(bytes32 _ref, address _child, uint8 _childType) returns()
func (_ITree *ITreeSession) AddChild(_ref [32]byte, _child common.Address, _childType uint8) (*types.Transaction, error) {
	return _ITree.Contract.AddChild(&_ITree.TransactOpts, _ref, _child, _childType)
}

// AddChild is a paid mutator transaction binding the contract method 0xcbb2dc95.
//
// Solidity: function addChild(bytes32 _ref, address _child, uint8 _childType) returns()
func (_ITree *ITreeTransactorSession) AddChild(_ref [32]byte, _child common.Address, _childType uint8) (*types.Transaction, error) {
	return _ITree.Contract.AddChild(&_ITree.TransactOpts, _ref, _child, _childType)
}

// ITreeAddChildIterator is returned from FilterAddChild and is used to iterate over the raw logs and unpacked data for AddChild events raised by the ITree contract.
type ITreeAddChildIterator struct {
	Event *ITreeAddChild // Event containing the contract specifics and raw log

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
func (it *ITreeAddChildIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITreeAddChild)
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
		it.Event = new(ITreeAddChild)
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
func (it *ITreeAddChildIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITreeAddChildIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITreeAddChild represents a AddChild event raised by the ITree contract.
type ITreeAddChild struct {
	Ref       [32]byte
	Child     common.Address
	ChildType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddChild is a free log retrieval operation binding the contract event 0x4f2635eaa771447d282168d4763629c4b855ff7c733898eefc20b8699278d0c0.
//
// Solidity: event AddChild(bytes32 _ref, address _child, uint8 _childType)
func (_ITree *ITreeFilterer) FilterAddChild(opts *bind.FilterOpts) (*ITreeAddChildIterator, error) {

	logs, sub, err := _ITree.contract.FilterLogs(opts, "AddChild")
	if err != nil {
		return nil, err
	}
	return &ITreeAddChildIterator{contract: _ITree.contract, event: "AddChild", logs: logs, sub: sub}, nil
}

// WatchAddChild is a free log subscription operation binding the contract event 0x4f2635eaa771447d282168d4763629c4b855ff7c733898eefc20b8699278d0c0.
//
// Solidity: event AddChild(bytes32 _ref, address _child, uint8 _childType)
func (_ITree *ITreeFilterer) WatchAddChild(opts *bind.WatchOpts, sink chan<- *ITreeAddChild) (event.Subscription, error) {

	logs, sub, err := _ITree.contract.WatchLogs(opts, "AddChild")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITreeAddChild)
				if err := _ITree.contract.UnpackLog(event, "AddChild", log); err != nil {
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

// ParseAddChild is a log parse operation binding the contract event 0x4f2635eaa771447d282168d4763629c4b855ff7c733898eefc20b8699278d0c0.
//
// Solidity: event AddChild(bytes32 _ref, address _child, uint8 _childType)
func (_ITree *ITreeFilterer) ParseAddChild(log types.Log) (*ITreeAddChild, error) {
	event := new(ITreeAddChild)
	if err := _ITree.contract.UnpackLog(event, "AddChild", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IterableDataABI is the input ABI used to generate the binding from.
const IterableDataABI = "[]"

// IterableDataBin is the compiled bytecode used for deploying new contracts.
var IterableDataBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122079cdbd1314ea6e222fa565c37c6a42d359afb4d017dabcb7081a6ef3690c3a3e64736f6c63430006060033"

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
