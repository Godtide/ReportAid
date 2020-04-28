// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ActivityAdditionalAdditional is an auto generated low-level Go binding around an user-defined struct.
type ActivityAdditionalAdditional struct {
	DefaultAidType     [32]byte
	DefaultFinanceType *big.Int
	Scope              uint8
	CapitalSpend       uint8
	CollaborationType  uint8
	DefaultFlowType    uint8
	DefaultTiedStatus  uint8
}

// ActivityAdditionalABI is the input ABI used to generate the binding from.
const ActivityAdditionalABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"}],\"name\":\"getAdditional\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"defaultAidType\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"defaultFinanceType\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"scope\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"capitalSpend\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"collaborationType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"defaultFlowType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"defaultTiedStatus\",\"type\":\"uint8\"}],\"internalType\":\"structActivityAdditional.Additional\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"}],\"name\":\"getCapitalSpend\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"}],\"name\":\"getCollaborationType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"}],\"name\":\"getDefaultAidType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"}],\"name\":\"getDefaultFinanceType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"}],\"name\":\"getDefaultFlowType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"}],\"name\":\"getDefaultTiedStatus\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getNumAdditional\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"}],\"name\":\"getScope\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"defaultAidType\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"defaultFinanceType\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"scope\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"capitalSpend\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"collaborationType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"defaultFlowType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"defaultTiedStatus\",\"type\":\"uint8\"}],\"internalType\":\"structActivityAdditional.Additional\",\"name\":\"_additional\",\"type\":\"tuple\"}],\"name\":\"setAdditional\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ActivityAdditionalFuncSigs maps the 4-byte function signature to its string representation.
var ActivityAdditionalFuncSigs = map[string]string{
	"8aaa51cb": "getAdditional(bytes32,bytes32,bytes32)",
	"74ef0ac7": "getCapitalSpend(bytes32,bytes32,bytes32)",
	"1239f64e": "getCollaborationType(bytes32,bytes32,bytes32)",
	"3cb4d8d3": "getDefaultAidType(bytes32,bytes32,bytes32)",
	"6c73ebf6": "getDefaultFinanceType(bytes32,bytes32,bytes32)",
	"f5d85a17": "getDefaultFlowType(bytes32,bytes32,bytes32)",
	"05c2e614": "getDefaultTiedStatus(bytes32,bytes32,bytes32)",
	"3818617f": "getNumAdditional(bytes32,bytes32)",
	"eef67d78": "getReference(bytes32,bytes32,uint256)",
	"3549a254": "getScope(bytes32,bytes32,bytes32)",
	"63387661": "setAdditional(bytes32,bytes32,bytes32,(bytes32,uint256,uint8,uint8,uint8,uint8,uint8))",
}

// ActivityAdditional is an auto generated Go binding around an Ethereum contract.
type ActivityAdditional struct {
	ActivityAdditionalCaller     // Read-only binding to the contract
	ActivityAdditionalTransactor // Write-only binding to the contract
	ActivityAdditionalFilterer   // Log filterer for contract events
}

// ActivityAdditionalCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivityAdditionalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityAdditionalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivityAdditionalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityAdditionalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivityAdditionalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityAdditionalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivityAdditionalSession struct {
	Contract     *ActivityAdditional // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ActivityAdditionalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivityAdditionalCallerSession struct {
	Contract *ActivityAdditionalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ActivityAdditionalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivityAdditionalTransactorSession struct {
	Contract     *ActivityAdditionalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ActivityAdditionalRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivityAdditionalRaw struct {
	Contract *ActivityAdditional // Generic contract binding to access the raw methods on
}

// ActivityAdditionalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivityAdditionalCallerRaw struct {
	Contract *ActivityAdditionalCaller // Generic read-only contract binding to access the raw methods on
}

// ActivityAdditionalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivityAdditionalTransactorRaw struct {
	Contract *ActivityAdditionalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivityAdditional creates a new instance of ActivityAdditional, bound to a specific deployed contract.
func NewActivityAdditional(address common.Address, backend bind.ContractBackend) (*ActivityAdditional, error) {
	contract, err := bindActivityAdditional(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ActivityAdditional{ActivityAdditionalCaller: ActivityAdditionalCaller{contract: contract}, ActivityAdditionalTransactor: ActivityAdditionalTransactor{contract: contract}, ActivityAdditionalFilterer: ActivityAdditionalFilterer{contract: contract}}, nil
}

// NewActivityAdditionalCaller creates a new read-only instance of ActivityAdditional, bound to a specific deployed contract.
func NewActivityAdditionalCaller(address common.Address, caller bind.ContractCaller) (*ActivityAdditionalCaller, error) {
	contract, err := bindActivityAdditional(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityAdditionalCaller{contract: contract}, nil
}

// NewActivityAdditionalTransactor creates a new write-only instance of ActivityAdditional, bound to a specific deployed contract.
func NewActivityAdditionalTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivityAdditionalTransactor, error) {
	contract, err := bindActivityAdditional(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityAdditionalTransactor{contract: contract}, nil
}

// NewActivityAdditionalFilterer creates a new log filterer instance of ActivityAdditional, bound to a specific deployed contract.
func NewActivityAdditionalFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivityAdditionalFilterer, error) {
	contract, err := bindActivityAdditional(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivityAdditionalFilterer{contract: contract}, nil
}

// bindActivityAdditional binds a generic wrapper to an already deployed contract.
func bindActivityAdditional(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivityAdditionalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityAdditional *ActivityAdditionalRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivityAdditional.Contract.ActivityAdditionalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityAdditional *ActivityAdditionalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityAdditional.Contract.ActivityAdditionalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityAdditional *ActivityAdditionalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityAdditional.Contract.ActivityAdditionalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityAdditional *ActivityAdditionalCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivityAdditional.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityAdditional *ActivityAdditionalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityAdditional.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityAdditional *ActivityAdditionalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityAdditional.Contract.contract.Transact(opts, method, params...)
}

// GetAdditional is a free data retrieval call binding the contract method 0x8aaa51cb.
//
// Solidity: function getAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(ActivityAdditionalAdditional)
func (_ActivityAdditional *ActivityAdditionalCaller) GetAdditional(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (ActivityAdditionalAdditional, error) {
	var (
		ret0 = new(ActivityAdditionalAdditional)
	)
	out := ret0
	err := _ActivityAdditional.contract.Call(opts, out, "getAdditional", _activitiesRef, _activityRef, _additionalRef)
	return *ret0, err
}

// GetAdditional is a free data retrieval call binding the contract method 0x8aaa51cb.
//
// Solidity: function getAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(ActivityAdditionalAdditional)
func (_ActivityAdditional *ActivityAdditionalSession) GetAdditional(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (ActivityAdditionalAdditional, error) {
	return _ActivityAdditional.Contract.GetAdditional(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetAdditional is a free data retrieval call binding the contract method 0x8aaa51cb.
//
// Solidity: function getAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(ActivityAdditionalAdditional)
func (_ActivityAdditional *ActivityAdditionalCallerSession) GetAdditional(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (ActivityAdditionalAdditional, error) {
	return _ActivityAdditional.Contract.GetAdditional(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetCapitalSpend is a free data retrieval call binding the contract method 0x74ef0ac7.
//
// Solidity: function getCapitalSpend(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_ActivityAdditional *ActivityAdditionalCaller) GetCapitalSpend(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ActivityAdditional.contract.Call(opts, out, "getCapitalSpend", _activitiesRef, _activityRef, _additionalRef)
	return *ret0, err
}

// GetCapitalSpend is a free data retrieval call binding the contract method 0x74ef0ac7.
//
// Solidity: function getCapitalSpend(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_ActivityAdditional *ActivityAdditionalSession) GetCapitalSpend(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _ActivityAdditional.Contract.GetCapitalSpend(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetCapitalSpend is a free data retrieval call binding the contract method 0x74ef0ac7.
//
// Solidity: function getCapitalSpend(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_ActivityAdditional *ActivityAdditionalCallerSession) GetCapitalSpend(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _ActivityAdditional.Contract.GetCapitalSpend(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetCollaborationType is a free data retrieval call binding the contract method 0x1239f64e.
//
// Solidity: function getCollaborationType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_ActivityAdditional *ActivityAdditionalCaller) GetCollaborationType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ActivityAdditional.contract.Call(opts, out, "getCollaborationType", _activitiesRef, _activityRef, _additionalRef)
	return *ret0, err
}

// GetCollaborationType is a free data retrieval call binding the contract method 0x1239f64e.
//
// Solidity: function getCollaborationType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_ActivityAdditional *ActivityAdditionalSession) GetCollaborationType(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _ActivityAdditional.Contract.GetCollaborationType(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetCollaborationType is a free data retrieval call binding the contract method 0x1239f64e.
//
// Solidity: function getCollaborationType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_ActivityAdditional *ActivityAdditionalCallerSession) GetCollaborationType(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _ActivityAdditional.Contract.GetCollaborationType(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetDefaultAidType is a free data retrieval call binding the contract method 0x3cb4d8d3.
//
// Solidity: function getDefaultAidType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(bytes32)
func (_ActivityAdditional *ActivityAdditionalCaller) GetDefaultAidType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityAdditional.contract.Call(opts, out, "getDefaultAidType", _activitiesRef, _activityRef, _additionalRef)
	return *ret0, err
}

// GetDefaultAidType is a free data retrieval call binding the contract method 0x3cb4d8d3.
//
// Solidity: function getDefaultAidType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(bytes32)
func (_ActivityAdditional *ActivityAdditionalSession) GetDefaultAidType(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) ([32]byte, error) {
	return _ActivityAdditional.Contract.GetDefaultAidType(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetDefaultAidType is a free data retrieval call binding the contract method 0x3cb4d8d3.
//
// Solidity: function getDefaultAidType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(bytes32)
func (_ActivityAdditional *ActivityAdditionalCallerSession) GetDefaultAidType(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) ([32]byte, error) {
	return _ActivityAdditional.Contract.GetDefaultAidType(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetDefaultFinanceType is a free data retrieval call binding the contract method 0x6c73ebf6.
//
// Solidity: function getDefaultFinanceType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint256)
func (_ActivityAdditional *ActivityAdditionalCaller) GetDefaultFinanceType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ActivityAdditional.contract.Call(opts, out, "getDefaultFinanceType", _activitiesRef, _activityRef, _additionalRef)
	return *ret0, err
}

// GetDefaultFinanceType is a free data retrieval call binding the contract method 0x6c73ebf6.
//
// Solidity: function getDefaultFinanceType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint256)
func (_ActivityAdditional *ActivityAdditionalSession) GetDefaultFinanceType(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (*big.Int, error) {
	return _ActivityAdditional.Contract.GetDefaultFinanceType(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetDefaultFinanceType is a free data retrieval call binding the contract method 0x6c73ebf6.
//
// Solidity: function getDefaultFinanceType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint256)
func (_ActivityAdditional *ActivityAdditionalCallerSession) GetDefaultFinanceType(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (*big.Int, error) {
	return _ActivityAdditional.Contract.GetDefaultFinanceType(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetDefaultFlowType is a free data retrieval call binding the contract method 0xf5d85a17.
//
// Solidity: function getDefaultFlowType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_ActivityAdditional *ActivityAdditionalCaller) GetDefaultFlowType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ActivityAdditional.contract.Call(opts, out, "getDefaultFlowType", _activitiesRef, _activityRef, _additionalRef)
	return *ret0, err
}

// GetDefaultFlowType is a free data retrieval call binding the contract method 0xf5d85a17.
//
// Solidity: function getDefaultFlowType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_ActivityAdditional *ActivityAdditionalSession) GetDefaultFlowType(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _ActivityAdditional.Contract.GetDefaultFlowType(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetDefaultFlowType is a free data retrieval call binding the contract method 0xf5d85a17.
//
// Solidity: function getDefaultFlowType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_ActivityAdditional *ActivityAdditionalCallerSession) GetDefaultFlowType(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _ActivityAdditional.Contract.GetDefaultFlowType(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetDefaultTiedStatus is a free data retrieval call binding the contract method 0x05c2e614.
//
// Solidity: function getDefaultTiedStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_ActivityAdditional *ActivityAdditionalCaller) GetDefaultTiedStatus(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ActivityAdditional.contract.Call(opts, out, "getDefaultTiedStatus", _activitiesRef, _activityRef, _additionalRef)
	return *ret0, err
}

// GetDefaultTiedStatus is a free data retrieval call binding the contract method 0x05c2e614.
//
// Solidity: function getDefaultTiedStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_ActivityAdditional *ActivityAdditionalSession) GetDefaultTiedStatus(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _ActivityAdditional.Contract.GetDefaultTiedStatus(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetDefaultTiedStatus is a free data retrieval call binding the contract method 0x05c2e614.
//
// Solidity: function getDefaultTiedStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_ActivityAdditional *ActivityAdditionalCallerSession) GetDefaultTiedStatus(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _ActivityAdditional.Contract.GetDefaultTiedStatus(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetNumAdditional is a free data retrieval call binding the contract method 0x3818617f.
//
// Solidity: function getNumAdditional(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivityAdditional *ActivityAdditionalCaller) GetNumAdditional(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ActivityAdditional.contract.Call(opts, out, "getNumAdditional", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetNumAdditional is a free data retrieval call binding the contract method 0x3818617f.
//
// Solidity: function getNumAdditional(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivityAdditional *ActivityAdditionalSession) GetNumAdditional(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _ActivityAdditional.Contract.GetNumAdditional(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef)
}

// GetNumAdditional is a free data retrieval call binding the contract method 0x3818617f.
//
// Solidity: function getNumAdditional(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivityAdditional *ActivityAdditionalCallerSession) GetNumAdditional(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _ActivityAdditional.Contract.GetNumAdditional(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivityAdditional *ActivityAdditionalCaller) GetReference(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityAdditional.contract.Call(opts, out, "getReference", _activitiesRef, _activityRef, _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivityAdditional *ActivityAdditionalSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _ActivityAdditional.Contract.GetReference(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivityAdditional *ActivityAdditionalCallerSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _ActivityAdditional.Contract.GetReference(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetScope is a free data retrieval call binding the contract method 0x3549a254.
//
// Solidity: function getScope(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_ActivityAdditional *ActivityAdditionalCaller) GetScope(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ActivityAdditional.contract.Call(opts, out, "getScope", _activitiesRef, _activityRef, _additionalRef)
	return *ret0, err
}

// GetScope is a free data retrieval call binding the contract method 0x3549a254.
//
// Solidity: function getScope(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_ActivityAdditional *ActivityAdditionalSession) GetScope(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _ActivityAdditional.Contract.GetScope(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetScope is a free data retrieval call binding the contract method 0x3549a254.
//
// Solidity: function getScope(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_ActivityAdditional *ActivityAdditionalCallerSession) GetScope(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _ActivityAdditional.Contract.GetScope(&_ActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// SetAdditional is a paid mutator transaction binding the contract method 0x63387661.
//
// Solidity: function setAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef, ActivityAdditionalAdditional _additional) returns()
func (_ActivityAdditional *ActivityAdditionalTransactor) SetAdditional(opts *bind.TransactOpts, _activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte, _additional ActivityAdditionalAdditional) (*types.Transaction, error) {
	return _ActivityAdditional.contract.Transact(opts, "setAdditional", _activitiesRef, _activityRef, _additionalRef, _additional)
}

// SetAdditional is a paid mutator transaction binding the contract method 0x63387661.
//
// Solidity: function setAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef, ActivityAdditionalAdditional _additional) returns()
func (_ActivityAdditional *ActivityAdditionalSession) SetAdditional(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte, _additional ActivityAdditionalAdditional) (*types.Transaction, error) {
	return _ActivityAdditional.Contract.SetAdditional(&_ActivityAdditional.TransactOpts, _activitiesRef, _activityRef, _additionalRef, _additional)
}

// SetAdditional is a paid mutator transaction binding the contract method 0x63387661.
//
// Solidity: function setAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef, ActivityAdditionalAdditional _additional) returns()
func (_ActivityAdditional *ActivityAdditionalTransactorSession) SetAdditional(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte, _additional ActivityAdditionalAdditional) (*types.Transaction, error) {
	return _ActivityAdditional.Contract.SetAdditional(&_ActivityAdditional.TransactOpts, _activitiesRef, _activityRef, _additionalRef, _additional)
}

// IATIActivityAdditionalABI is the input ABI used to generate the binding from.
const IATIActivityAdditionalABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"defaultAidType\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"defaultFinanceType\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"scope\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"capitalSpend\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"collaborationType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"defaultFlowType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"defaultTiedStatus\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structActivityAdditional.Additional\",\"name\":\"_additional\",\"type\":\"tuple\"}],\"name\":\"SetAdditional\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"}],\"name\":\"getAdditional\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"defaultAidType\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"defaultFinanceType\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"scope\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"capitalSpend\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"collaborationType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"defaultFlowType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"defaultTiedStatus\",\"type\":\"uint8\"}],\"internalType\":\"structActivityAdditional.Additional\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"}],\"name\":\"getCapitalSpend\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"}],\"name\":\"getCollaborationType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"}],\"name\":\"getDefaultAidType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"}],\"name\":\"getDefaultFinanceType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"}],\"name\":\"getDefaultFlowType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"}],\"name\":\"getDefaultTiedStatus\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getNumAdditional\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"}],\"name\":\"getScope\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_additionalRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"defaultAidType\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"defaultFinanceType\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"scope\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"capitalSpend\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"collaborationType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"defaultFlowType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"defaultTiedStatus\",\"type\":\"uint8\"}],\"internalType\":\"structActivityAdditional.Additional\",\"name\":\"_additional\",\"type\":\"tuple\"}],\"name\":\"setAdditional\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IATIActivityAdditionalFuncSigs maps the 4-byte function signature to its string representation.
var IATIActivityAdditionalFuncSigs = map[string]string{
	"8aaa51cb": "getAdditional(bytes32,bytes32,bytes32)",
	"74ef0ac7": "getCapitalSpend(bytes32,bytes32,bytes32)",
	"1239f64e": "getCollaborationType(bytes32,bytes32,bytes32)",
	"3cb4d8d3": "getDefaultAidType(bytes32,bytes32,bytes32)",
	"6c73ebf6": "getDefaultFinanceType(bytes32,bytes32,bytes32)",
	"f5d85a17": "getDefaultFlowType(bytes32,bytes32,bytes32)",
	"05c2e614": "getDefaultTiedStatus(bytes32,bytes32,bytes32)",
	"3818617f": "getNumAdditional(bytes32,bytes32)",
	"eef67d78": "getReference(bytes32,bytes32,uint256)",
	"3549a254": "getScope(bytes32,bytes32,bytes32)",
	"63387661": "setAdditional(bytes32,bytes32,bytes32,(bytes32,uint256,uint8,uint8,uint8,uint8,uint8))",
}

// IATIActivityAdditionalBin is the compiled bytecode used for deploying new contracts.
var IATIActivityAdditionalBin = "0x608060405234801561001057600080fd5b50610c94806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063633876611161007157806363387661146101305780636c73ebf61461014557806374ef0ac7146101585780638aaa51cb1461016b578063eef67d781461018b578063f5d85a171461019e576100a9565b806305c2e614146100ae5780631239f64e146100d75780633549a254146100ea5780633818617f146100fd5780633cb4d8d31461011d575b600080fd5b6100c16100bc366004610a47565b6101b1565b6040516100ce9190610c29565b60405180910390f35b6100c16100e5366004610a47565b610237565b6100c16100f8366004610a47565b6102bb565b61011061010b366004610a26565b610339565b6040516100ce9190610b96565b61011061012b366004610a47565b610391565b61014361013e366004610a72565b610409565b005b610110610153366004610a47565b6106bf565b6100c1610166366004610a47565b61073b565b61017e610179366004610a47565b6107be565b6040516100ce9190610c1b565b610110610199366004610b2f565b61089c565b6100c16101ac366004610a47565b61092d565b600083811a60f81b6001600160f81b031916158015906101e057508260001a60f81b6001600160f81b03191615155b80156101fb57508160001a60f81b6001600160f81b03191615155b61020457600080fd5b50600092835260016020908152604080852093855292815282842091845252902060020154640100000000900460ff1690565b600083811a60f81b6001600160f81b0319161580159061026657508260001a60f81b6001600160f81b03191615155b801561028157508160001a60f81b6001600160f81b03191615155b61028a57600080fd5b5060009283526001602090815260408085209385529281528284209184525290206002015462010000900460ff1690565b600083811a60f81b6001600160f81b031916158015906102ea57508260001a60f81b6001600160f81b03191615155b801561030557508160001a60f81b6001600160f81b03191615155b61030e57600080fd5b5060009283526001602090815260408085209385529281528284209184525290206002015460ff1690565b600082811a60f81b6001600160f81b0319161580159061036857508160001a60f81b6001600160f81b03191615155b61037157600080fd5b506000828152602081815260408083208484529091529020545b92915050565b600083811a60f81b6001600160f81b031916158015906103c057508260001a60f81b6001600160f81b03191615155b80156103db57508160001a60f81b6001600160f81b03191615155b6103e457600080fd5b5060009283526001602090815260408085209385529281528284209184525290205490565b8360001a60f81b6001600160f81b0319161580159061043757508260001a60f81b6001600160f81b03191615155b801561045257508160001a60f81b6001600160f81b03191615155b80156104645750604081015160ff1615155b801561047a5750600960ff16816040015160ff16105b801561048e57506064816060015160ff1611155b80156104a05750608081015160ff1615155b80156104b65750600860ff16816080015160ff16105b80156104d25750805160001a60f81b6001600160f81b03191615155b80156104e85750600260ff168160c0015160ff16115b80156104fe5750600660ff168160c0015160ff16105b61050757600080fd5b60008481526001602081815260408084208785528252808420868552825280842085518155828601519381019390935580850151600290930180546060870151608088015160a089015160c08a015160ff9081166401000000000264ff000000001992821663010000000263ff00000019948316620100000262ff0000199684166101000261ff001994909c1660ff19909816979097179290921699909917939093169390931716171693909317909255868352828152818320868452905290819020905163745fca7b60e01b815273__$b620240993a55615b607189176fb82e62f$__9163745fca7b91610600918691600401610b9f565b60206040518083038186803b15801561061857600080fd5b505af415801561062c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061065091906109ff565b61067c576000848152602081815260408083208684528252822080546001810182559083529120018290555b7f868373f61c5202d8e8d3a78da20efd58f3b7e9c6e44e93860b04799cecce4755848484846040516106b19493929190610bef565b60405180910390a150505050565b600083811a60f81b6001600160f81b031916158015906106ee57508260001a60f81b6001600160f81b03191615155b801561070957508160001a60f81b6001600160f81b03191615155b61071257600080fd5b506000928352600160208181526040808620948652938152838520928552919091529120015490565b600083811a60f81b6001600160f81b0319161580159061076a57508260001a60f81b6001600160f81b03191615155b801561078557508160001a60f81b6001600160f81b03191615155b61078e57600080fd5b50600092835260016020908152604080852093855292815282842091845252902060020154610100900460ff1690565b6107c66109b2565b8360001a60f81b6001600160f81b031916158015906107f457508260001a60f81b6001600160f81b03191615155b801561080f57508160001a60f81b6001600160f81b03191615155b61081857600080fd5b50600092835260016020818152604080862094865293815283852092855291825292829020825160e0810184528154815293810154918401919091526002015460ff808216928401929092526101008104821660608401526201000081048216608084015263010000008104821660a084015264010000000090041660c082015290565b600083811a60f81b6001600160f81b031916158015906108cb57508260001a60f81b6001600160f81b03191615155b80156108ed575060008481526020818152604080832086845290915290205482105b6108f657600080fd5b600084815260208181526040808320868452909152902080548390811061091957fe5b906000526020600020015490509392505050565b600083811a60f81b6001600160f81b0319161580159061095c57508260001a60f81b6001600160f81b03191615155b801561097757508160001a60f81b6001600160f81b03191615155b61098057600080fd5b506000928352600160209081526040808520938552928152828420918452529020600201546301000000900460ff1690565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c081019190915290565b803560ff8116811461038b57600080fd5b600060208284031215610a10578081fd5b81518015158114610a1f578182fd5b9392505050565b60008060408385031215610a38578081fd5b50508035926020909101359150565b600080600060608486031215610a5b578081fd5b505081359360208301359350604090920135919050565b600080600080848603610140811215610a89578182fd5b85359450602086013593506040860135925060e0605f1982011215610aac578182fd5b50610ab760e0610c37565b6060860135815260808601356020820152610ad58760a088016109ee565b6040820152610ae78760c088016109ee565b6060820152610af98760e088016109ee565b6080820152610b0c8761010088016109ee565b60a0820152610b1f8761012088016109ee565b60c0820152939692955090935050565b600080600060608486031215610a5b578283fd5b805182526020810151602083015260ff604082015116604083015260ff606082015116606083015260ff608082015116608083015260ff60a08201511660a083015260ff60c08201511660c08301525050565b90815260200190565b60006040820184835260206040818501528185548084526060860191508685528285209350845b81811015610be257845483526001948501949284019201610bc6565b5090979650505050505050565b84815260208101849052604081018390526101408101610c126060830184610b43565b95945050505050565b60e0810161038b8284610b43565b60ff91909116815260200190565b60405181810167ffffffffffffffff81118282101715610c5657600080fd5b60405291905056fea26469706673582212201062873c4851c182bf2b32fcb027813484da9d9e58e3cc71cc784ddbeec5b15d64736f6c63430006060033"

// DeployIATIActivityAdditional deploys a new Ethereum contract, binding an instance of IATIActivityAdditional to it.
func DeployIATIActivityAdditional(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IATIActivityAdditional, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivityAdditionalABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	stringsAddr, _, _, _ := DeployStrings(auth, backend)
	IATIActivityAdditionalBin = strings.Replace(IATIActivityAdditionalBin, "__$b620240993a55615b607189176fb82e62f$__", stringsAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IATIActivityAdditionalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IATIActivityAdditional{IATIActivityAdditionalCaller: IATIActivityAdditionalCaller{contract: contract}, IATIActivityAdditionalTransactor: IATIActivityAdditionalTransactor{contract: contract}, IATIActivityAdditionalFilterer: IATIActivityAdditionalFilterer{contract: contract}}, nil
}

// IATIActivityAdditional is an auto generated Go binding around an Ethereum contract.
type IATIActivityAdditional struct {
	IATIActivityAdditionalCaller     // Read-only binding to the contract
	IATIActivityAdditionalTransactor // Write-only binding to the contract
	IATIActivityAdditionalFilterer   // Log filterer for contract events
}

// IATIActivityAdditionalCaller is an auto generated read-only Go binding around an Ethereum contract.
type IATIActivityAdditionalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityAdditionalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IATIActivityAdditionalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityAdditionalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IATIActivityAdditionalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityAdditionalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IATIActivityAdditionalSession struct {
	Contract     *IATIActivityAdditional // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IATIActivityAdditionalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IATIActivityAdditionalCallerSession struct {
	Contract *IATIActivityAdditionalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// IATIActivityAdditionalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IATIActivityAdditionalTransactorSession struct {
	Contract     *IATIActivityAdditionalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// IATIActivityAdditionalRaw is an auto generated low-level Go binding around an Ethereum contract.
type IATIActivityAdditionalRaw struct {
	Contract *IATIActivityAdditional // Generic contract binding to access the raw methods on
}

// IATIActivityAdditionalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IATIActivityAdditionalCallerRaw struct {
	Contract *IATIActivityAdditionalCaller // Generic read-only contract binding to access the raw methods on
}

// IATIActivityAdditionalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IATIActivityAdditionalTransactorRaw struct {
	Contract *IATIActivityAdditionalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIATIActivityAdditional creates a new instance of IATIActivityAdditional, bound to a specific deployed contract.
func NewIATIActivityAdditional(address common.Address, backend bind.ContractBackend) (*IATIActivityAdditional, error) {
	contract, err := bindIATIActivityAdditional(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IATIActivityAdditional{IATIActivityAdditionalCaller: IATIActivityAdditionalCaller{contract: contract}, IATIActivityAdditionalTransactor: IATIActivityAdditionalTransactor{contract: contract}, IATIActivityAdditionalFilterer: IATIActivityAdditionalFilterer{contract: contract}}, nil
}

// NewIATIActivityAdditionalCaller creates a new read-only instance of IATIActivityAdditional, bound to a specific deployed contract.
func NewIATIActivityAdditionalCaller(address common.Address, caller bind.ContractCaller) (*IATIActivityAdditionalCaller, error) {
	contract, err := bindIATIActivityAdditional(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IATIActivityAdditionalCaller{contract: contract}, nil
}

// NewIATIActivityAdditionalTransactor creates a new write-only instance of IATIActivityAdditional, bound to a specific deployed contract.
func NewIATIActivityAdditionalTransactor(address common.Address, transactor bind.ContractTransactor) (*IATIActivityAdditionalTransactor, error) {
	contract, err := bindIATIActivityAdditional(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IATIActivityAdditionalTransactor{contract: contract}, nil
}

// NewIATIActivityAdditionalFilterer creates a new log filterer instance of IATIActivityAdditional, bound to a specific deployed contract.
func NewIATIActivityAdditionalFilterer(address common.Address, filterer bind.ContractFilterer) (*IATIActivityAdditionalFilterer, error) {
	contract, err := bindIATIActivityAdditional(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IATIActivityAdditionalFilterer{contract: contract}, nil
}

// bindIATIActivityAdditional binds a generic wrapper to an already deployed contract.
func bindIATIActivityAdditional(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivityAdditionalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIActivityAdditional *IATIActivityAdditionalRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIActivityAdditional.Contract.IATIActivityAdditionalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIActivityAdditional *IATIActivityAdditionalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIActivityAdditional.Contract.IATIActivityAdditionalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIActivityAdditional *IATIActivityAdditionalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIActivityAdditional.Contract.IATIActivityAdditionalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIActivityAdditional *IATIActivityAdditionalCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIActivityAdditional.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIActivityAdditional *IATIActivityAdditionalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIActivityAdditional.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIActivityAdditional *IATIActivityAdditionalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIActivityAdditional.Contract.contract.Transact(opts, method, params...)
}

// GetAdditional is a free data retrieval call binding the contract method 0x8aaa51cb.
//
// Solidity: function getAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(ActivityAdditionalAdditional)
func (_IATIActivityAdditional *IATIActivityAdditionalCaller) GetAdditional(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (ActivityAdditionalAdditional, error) {
	var (
		ret0 = new(ActivityAdditionalAdditional)
	)
	out := ret0
	err := _IATIActivityAdditional.contract.Call(opts, out, "getAdditional", _activitiesRef, _activityRef, _additionalRef)
	return *ret0, err
}

// GetAdditional is a free data retrieval call binding the contract method 0x8aaa51cb.
//
// Solidity: function getAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(ActivityAdditionalAdditional)
func (_IATIActivityAdditional *IATIActivityAdditionalSession) GetAdditional(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (ActivityAdditionalAdditional, error) {
	return _IATIActivityAdditional.Contract.GetAdditional(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetAdditional is a free data retrieval call binding the contract method 0x8aaa51cb.
//
// Solidity: function getAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(ActivityAdditionalAdditional)
func (_IATIActivityAdditional *IATIActivityAdditionalCallerSession) GetAdditional(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (ActivityAdditionalAdditional, error) {
	return _IATIActivityAdditional.Contract.GetAdditional(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetCapitalSpend is a free data retrieval call binding the contract method 0x74ef0ac7.
//
// Solidity: function getCapitalSpend(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_IATIActivityAdditional *IATIActivityAdditionalCaller) GetCapitalSpend(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivityAdditional.contract.Call(opts, out, "getCapitalSpend", _activitiesRef, _activityRef, _additionalRef)
	return *ret0, err
}

// GetCapitalSpend is a free data retrieval call binding the contract method 0x74ef0ac7.
//
// Solidity: function getCapitalSpend(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_IATIActivityAdditional *IATIActivityAdditionalSession) GetCapitalSpend(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _IATIActivityAdditional.Contract.GetCapitalSpend(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetCapitalSpend is a free data retrieval call binding the contract method 0x74ef0ac7.
//
// Solidity: function getCapitalSpend(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_IATIActivityAdditional *IATIActivityAdditionalCallerSession) GetCapitalSpend(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _IATIActivityAdditional.Contract.GetCapitalSpend(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetCollaborationType is a free data retrieval call binding the contract method 0x1239f64e.
//
// Solidity: function getCollaborationType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_IATIActivityAdditional *IATIActivityAdditionalCaller) GetCollaborationType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivityAdditional.contract.Call(opts, out, "getCollaborationType", _activitiesRef, _activityRef, _additionalRef)
	return *ret0, err
}

// GetCollaborationType is a free data retrieval call binding the contract method 0x1239f64e.
//
// Solidity: function getCollaborationType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_IATIActivityAdditional *IATIActivityAdditionalSession) GetCollaborationType(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _IATIActivityAdditional.Contract.GetCollaborationType(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetCollaborationType is a free data retrieval call binding the contract method 0x1239f64e.
//
// Solidity: function getCollaborationType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_IATIActivityAdditional *IATIActivityAdditionalCallerSession) GetCollaborationType(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _IATIActivityAdditional.Contract.GetCollaborationType(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetDefaultAidType is a free data retrieval call binding the contract method 0x3cb4d8d3.
//
// Solidity: function getDefaultAidType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(bytes32)
func (_IATIActivityAdditional *IATIActivityAdditionalCaller) GetDefaultAidType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityAdditional.contract.Call(opts, out, "getDefaultAidType", _activitiesRef, _activityRef, _additionalRef)
	return *ret0, err
}

// GetDefaultAidType is a free data retrieval call binding the contract method 0x3cb4d8d3.
//
// Solidity: function getDefaultAidType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(bytes32)
func (_IATIActivityAdditional *IATIActivityAdditionalSession) GetDefaultAidType(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) ([32]byte, error) {
	return _IATIActivityAdditional.Contract.GetDefaultAidType(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetDefaultAidType is a free data retrieval call binding the contract method 0x3cb4d8d3.
//
// Solidity: function getDefaultAidType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(bytes32)
func (_IATIActivityAdditional *IATIActivityAdditionalCallerSession) GetDefaultAidType(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) ([32]byte, error) {
	return _IATIActivityAdditional.Contract.GetDefaultAidType(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetDefaultFinanceType is a free data retrieval call binding the contract method 0x6c73ebf6.
//
// Solidity: function getDefaultFinanceType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint256)
func (_IATIActivityAdditional *IATIActivityAdditionalCaller) GetDefaultFinanceType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIActivityAdditional.contract.Call(opts, out, "getDefaultFinanceType", _activitiesRef, _activityRef, _additionalRef)
	return *ret0, err
}

// GetDefaultFinanceType is a free data retrieval call binding the contract method 0x6c73ebf6.
//
// Solidity: function getDefaultFinanceType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint256)
func (_IATIActivityAdditional *IATIActivityAdditionalSession) GetDefaultFinanceType(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (*big.Int, error) {
	return _IATIActivityAdditional.Contract.GetDefaultFinanceType(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetDefaultFinanceType is a free data retrieval call binding the contract method 0x6c73ebf6.
//
// Solidity: function getDefaultFinanceType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint256)
func (_IATIActivityAdditional *IATIActivityAdditionalCallerSession) GetDefaultFinanceType(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (*big.Int, error) {
	return _IATIActivityAdditional.Contract.GetDefaultFinanceType(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetDefaultFlowType is a free data retrieval call binding the contract method 0xf5d85a17.
//
// Solidity: function getDefaultFlowType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_IATIActivityAdditional *IATIActivityAdditionalCaller) GetDefaultFlowType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivityAdditional.contract.Call(opts, out, "getDefaultFlowType", _activitiesRef, _activityRef, _additionalRef)
	return *ret0, err
}

// GetDefaultFlowType is a free data retrieval call binding the contract method 0xf5d85a17.
//
// Solidity: function getDefaultFlowType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_IATIActivityAdditional *IATIActivityAdditionalSession) GetDefaultFlowType(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _IATIActivityAdditional.Contract.GetDefaultFlowType(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetDefaultFlowType is a free data retrieval call binding the contract method 0xf5d85a17.
//
// Solidity: function getDefaultFlowType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_IATIActivityAdditional *IATIActivityAdditionalCallerSession) GetDefaultFlowType(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _IATIActivityAdditional.Contract.GetDefaultFlowType(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetDefaultTiedStatus is a free data retrieval call binding the contract method 0x05c2e614.
//
// Solidity: function getDefaultTiedStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_IATIActivityAdditional *IATIActivityAdditionalCaller) GetDefaultTiedStatus(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivityAdditional.contract.Call(opts, out, "getDefaultTiedStatus", _activitiesRef, _activityRef, _additionalRef)
	return *ret0, err
}

// GetDefaultTiedStatus is a free data retrieval call binding the contract method 0x05c2e614.
//
// Solidity: function getDefaultTiedStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_IATIActivityAdditional *IATIActivityAdditionalSession) GetDefaultTiedStatus(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _IATIActivityAdditional.Contract.GetDefaultTiedStatus(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetDefaultTiedStatus is a free data retrieval call binding the contract method 0x05c2e614.
//
// Solidity: function getDefaultTiedStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_IATIActivityAdditional *IATIActivityAdditionalCallerSession) GetDefaultTiedStatus(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _IATIActivityAdditional.Contract.GetDefaultTiedStatus(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetNumAdditional is a free data retrieval call binding the contract method 0x3818617f.
//
// Solidity: function getNumAdditional(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivityAdditional *IATIActivityAdditionalCaller) GetNumAdditional(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIActivityAdditional.contract.Call(opts, out, "getNumAdditional", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetNumAdditional is a free data retrieval call binding the contract method 0x3818617f.
//
// Solidity: function getNumAdditional(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivityAdditional *IATIActivityAdditionalSession) GetNumAdditional(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _IATIActivityAdditional.Contract.GetNumAdditional(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef)
}

// GetNumAdditional is a free data retrieval call binding the contract method 0x3818617f.
//
// Solidity: function getNumAdditional(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivityAdditional *IATIActivityAdditionalCallerSession) GetNumAdditional(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _IATIActivityAdditional.Contract.GetNumAdditional(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivityAdditional *IATIActivityAdditionalCaller) GetReference(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityAdditional.contract.Call(opts, out, "getReference", _activitiesRef, _activityRef, _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivityAdditional *IATIActivityAdditionalSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIActivityAdditional.Contract.GetReference(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivityAdditional *IATIActivityAdditionalCallerSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIActivityAdditional.Contract.GetReference(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetScope is a free data retrieval call binding the contract method 0x3549a254.
//
// Solidity: function getScope(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_IATIActivityAdditional *IATIActivityAdditionalCaller) GetScope(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivityAdditional.contract.Call(opts, out, "getScope", _activitiesRef, _activityRef, _additionalRef)
	return *ret0, err
}

// GetScope is a free data retrieval call binding the contract method 0x3549a254.
//
// Solidity: function getScope(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_IATIActivityAdditional *IATIActivityAdditionalSession) GetScope(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _IATIActivityAdditional.Contract.GetScope(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// GetScope is a free data retrieval call binding the contract method 0x3549a254.
//
// Solidity: function getScope(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef) view returns(uint8)
func (_IATIActivityAdditional *IATIActivityAdditionalCallerSession) GetScope(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte) (uint8, error) {
	return _IATIActivityAdditional.Contract.GetScope(&_IATIActivityAdditional.CallOpts, _activitiesRef, _activityRef, _additionalRef)
}

// SetAdditional is a paid mutator transaction binding the contract method 0x63387661.
//
// Solidity: function setAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef, ActivityAdditionalAdditional _additional) returns()
func (_IATIActivityAdditional *IATIActivityAdditionalTransactor) SetAdditional(opts *bind.TransactOpts, _activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte, _additional ActivityAdditionalAdditional) (*types.Transaction, error) {
	return _IATIActivityAdditional.contract.Transact(opts, "setAdditional", _activitiesRef, _activityRef, _additionalRef, _additional)
}

// SetAdditional is a paid mutator transaction binding the contract method 0x63387661.
//
// Solidity: function setAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef, ActivityAdditionalAdditional _additional) returns()
func (_IATIActivityAdditional *IATIActivityAdditionalSession) SetAdditional(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte, _additional ActivityAdditionalAdditional) (*types.Transaction, error) {
	return _IATIActivityAdditional.Contract.SetAdditional(&_IATIActivityAdditional.TransactOpts, _activitiesRef, _activityRef, _additionalRef, _additional)
}

// SetAdditional is a paid mutator transaction binding the contract method 0x63387661.
//
// Solidity: function setAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef, ActivityAdditionalAdditional _additional) returns()
func (_IATIActivityAdditional *IATIActivityAdditionalTransactorSession) SetAdditional(_activitiesRef [32]byte, _activityRef [32]byte, _additionalRef [32]byte, _additional ActivityAdditionalAdditional) (*types.Transaction, error) {
	return _IATIActivityAdditional.Contract.SetAdditional(&_IATIActivityAdditional.TransactOpts, _activitiesRef, _activityRef, _additionalRef, _additional)
}

// IATIActivityAdditionalSetAdditionalIterator is returned from FilterSetAdditional and is used to iterate over the raw logs and unpacked data for SetAdditional events raised by the IATIActivityAdditional contract.
type IATIActivityAdditionalSetAdditionalIterator struct {
	Event *IATIActivityAdditionalSetAdditional // Event containing the contract specifics and raw log

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
func (it *IATIActivityAdditionalSetAdditionalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IATIActivityAdditionalSetAdditional)
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
		it.Event = new(IATIActivityAdditionalSetAdditional)
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
func (it *IATIActivityAdditionalSetAdditionalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IATIActivityAdditionalSetAdditionalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IATIActivityAdditionalSetAdditional represents a SetAdditional event raised by the IATIActivityAdditional contract.
type IATIActivityAdditionalSetAdditional struct {
	ActivitiesRef [32]byte
	ActivityRef   [32]byte
	AdditionalRef [32]byte
	Additional    ActivityAdditionalAdditional
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetAdditional is a free log retrieval operation binding the contract event 0x868373f61c5202d8e8d3a78da20efd58f3b7e9c6e44e93860b04799cecce4755.
//
// Solidity: event SetAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef, ActivityAdditionalAdditional _additional)
func (_IATIActivityAdditional *IATIActivityAdditionalFilterer) FilterSetAdditional(opts *bind.FilterOpts) (*IATIActivityAdditionalSetAdditionalIterator, error) {

	logs, sub, err := _IATIActivityAdditional.contract.FilterLogs(opts, "SetAdditional")
	if err != nil {
		return nil, err
	}
	return &IATIActivityAdditionalSetAdditionalIterator{contract: _IATIActivityAdditional.contract, event: "SetAdditional", logs: logs, sub: sub}, nil
}

// WatchSetAdditional is a free log subscription operation binding the contract event 0x868373f61c5202d8e8d3a78da20efd58f3b7e9c6e44e93860b04799cecce4755.
//
// Solidity: event SetAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef, ActivityAdditionalAdditional _additional)
func (_IATIActivityAdditional *IATIActivityAdditionalFilterer) WatchSetAdditional(opts *bind.WatchOpts, sink chan<- *IATIActivityAdditionalSetAdditional) (event.Subscription, error) {

	logs, sub, err := _IATIActivityAdditional.contract.WatchLogs(opts, "SetAdditional")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IATIActivityAdditionalSetAdditional)
				if err := _IATIActivityAdditional.contract.UnpackLog(event, "SetAdditional", log); err != nil {
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

// ParseSetAdditional is a log parse operation binding the contract event 0x868373f61c5202d8e8d3a78da20efd58f3b7e9c6e44e93860b04799cecce4755.
//
// Solidity: event SetAdditional(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _additionalRef, ActivityAdditionalAdditional _additional)
func (_IATIActivityAdditional *IATIActivityAdditionalFilterer) ParseSetAdditional(log types.Log) (*IATIActivityAdditionalSetAdditional, error) {
	event := new(IATIActivityAdditionalSetAdditional)
	if err := _IATIActivityAdditional.contract.UnpackLog(event, "SetAdditional", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StringsABI is the input ABI used to generate the binding from.
const StringsABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_x\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToStr\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_a\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_b\",\"type\":\"bytes32\"}],\"name\":\"compare\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_a\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_b\",\"type\":\"string\"}],\"name\":\"compare\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_xs\",\"type\":\"bytes32[]\"}],\"name\":\"getExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_a\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_xs\",\"type\":\"string[]\"}],\"name\":\"getIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"_xs\",\"type\":\"bytes32[]\"}],\"name\":\"getIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// StringsFuncSigs maps the 4-byte function signature to its string representation.
var StringsFuncSigs = map[string]string{
	"ef0b2368": "bytes32ToStr(bytes32)",
	"28602ab2": "compare(bytes32,bytes32)",
	"3a96fdd7": "compare(string,string)",
	"745fca7b": "getExists(bytes32,bytes32[])",
	"c55f207f": "getIndex(bytes32,bytes32[])",
	"02a33464": "getIndex(string,string[])",
}

// StringsBin is the compiled bytecode used for deploying new contracts.
var StringsBin = "0x610662610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061006c5760003560e01c806302a334641461007157806328602ab21461009a5780633a96fdd7146100ba578063745fca7b146100cd578063c55f207f146100e0578063ef0b2368146100f3575b600080fd5b61008461007f36600461046d565b610113565b60405161009191906105dc565b60405180910390f35b6100ad6100a836600461044c565b610180565b604051610091919061057e565b6100ad6100c836600461051d565b6101bd565b6100ad6100db3660046103a8565b61020b565b6100846100ee3660046103a8565b61024b565b610106610101366004610390565b6102b5565b6040516100919190610589565b6000808351118015610126575060008251115b61012f57600080fd5b815160005b835181101561017657606084828151811061014b57fe5b6020026020010151905061015f86826101bd565b1561016d5781925050610176565b50600101610134565b5090505b92915050565b600082811a60f81b6001600160f81b031916158015906101af57508160001a60f81b6001600160f81b03191615155b6101b857600080fd5b501490565b60008083511180156101d0575060008251115b6101d957600080fd5b8151835184918491146101f15760009250505061017a565b80805190602001208280519060200120149250505061017a565b600082811a60f81b6001600160f81b03191661022657600080fd5b81516000901561024457600061023c858561024b565b845114159150505b9392505050565b600082811a60f81b6001600160f81b0319161580159061026c575060008251115b61027557600080fd5b815160005b8351811015610176576102a08585838151811061029357fe5b6020026020010151610180565b156102ad57809150610176565b60010161027a565b60408051602080825281830190925260609182919060208201818036833701905050905060005b6020811015610320578381602081106102f157fe5b1a60f81b82828151811061030157fe5b60200101906001600160f81b031916908160001a9053506001016102dc565b5092915050565b600082601f830112610337578081fd5b813567ffffffffffffffff81111561034d578182fd5b610360601f8201601f19166020016105e5565b915080825283602082850101111561037757600080fd5b8060208401602084013760009082016020015292915050565b6000602082840312156103a1578081fd5b5035919050565b600080604083850312156103ba578081fd5b8235915060208084013567ffffffffffffffff8111156103d8578283fd5b80850186601f8201126103e9578384fd5b803591506103fe6103f98361060c565b6105e5565b82815283810190828501858502840186018a101561041a578687fd5b8693505b8484101561043c57803583526001939093019291850191850161041e565b5080955050505050509250929050565b6000806040838503121561045e578182fd5b50508035926020909101359150565b6000806040838503121561047f578182fd5b823567ffffffffffffffff80821115610496578384fd5b6104a286838701610327565b93506020915081850135818111156104b8578384fd5b85019050601f810186136104ca578283fd5b80356104d86103f98261060c565b81815283810190838501865b8481101561050d576104fb8b888435890101610327565b845292860192908601906001016104e4565b5096999098509650505050505050565b6000806040838503121561052f578182fd5b823567ffffffffffffffff80821115610546578384fd5b61055286838701610327565b93506020850135915080821115610567578283fd5b5061057485828601610327565b9150509250929050565b901515815260200190565b6000602080835283518082850152825b818110156105b557858101830151858201604001528201610599565b818111156105c65783604083870101525b50601f01601f1916929092016040019392505050565b90815260200190565b60405181810167ffffffffffffffff8111828210171561060457600080fd5b604052919050565b600067ffffffffffffffff821115610622578081fd5b506020908102019056fea2646970667358221220afe80f7b79e21558518987b567fad792efadb78d8c923a4d1825b3998ada86f464736f6c63430006060033"

// DeployStrings deploys a new Ethereum contract, binding an instance of Strings to it.
func DeployStrings(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Strings, error) {
	parsed, err := abi.JSON(strings.NewReader(StringsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StringsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// Strings is an auto generated Go binding around an Ethereum contract.
type Strings struct {
	StringsCaller     // Read-only binding to the contract
	StringsTransactor // Write-only binding to the contract
	StringsFilterer   // Log filterer for contract events
}

// StringsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StringsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StringsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StringsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StringsSession struct {
	Contract     *Strings          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StringsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StringsCallerSession struct {
	Contract *StringsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StringsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StringsTransactorSession struct {
	Contract     *StringsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StringsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StringsRaw struct {
	Contract *Strings // Generic contract binding to access the raw methods on
}

// StringsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StringsCallerRaw struct {
	Contract *StringsCaller // Generic read-only contract binding to access the raw methods on
}

// StringsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StringsTransactorRaw struct {
	Contract *StringsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrings creates a new instance of Strings, bound to a specific deployed contract.
func NewStrings(address common.Address, backend bind.ContractBackend) (*Strings, error) {
	contract, err := bindStrings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// NewStringsCaller creates a new read-only instance of Strings, bound to a specific deployed contract.
func NewStringsCaller(address common.Address, caller bind.ContractCaller) (*StringsCaller, error) {
	contract, err := bindStrings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StringsCaller{contract: contract}, nil
}

// NewStringsTransactor creates a new write-only instance of Strings, bound to a specific deployed contract.
func NewStringsTransactor(address common.Address, transactor bind.ContractTransactor) (*StringsTransactor, error) {
	contract, err := bindStrings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StringsTransactor{contract: contract}, nil
}

// NewStringsFilterer creates a new log filterer instance of Strings, bound to a specific deployed contract.
func NewStringsFilterer(address common.Address, filterer bind.ContractFilterer) (*StringsFilterer, error) {
	contract, err := bindStrings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StringsFilterer{contract: contract}, nil
}

// bindStrings binds a generic wrapper to an already deployed contract.
func bindStrings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StringsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.StringsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transact(opts, method, params...)
}

// Bytes32ToStr is a free data retrieval call binding the contract method 0xef0b2368.
//
// Solidity: function bytes32ToStr(bytes32 _x) pure returns(string)
func (_Strings *StringsCaller) Bytes32ToStr(opts *bind.CallOpts, _x [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Strings.contract.Call(opts, out, "bytes32ToStr", _x)
	return *ret0, err
}

// Bytes32ToStr is a free data retrieval call binding the contract method 0xef0b2368.
//
// Solidity: function bytes32ToStr(bytes32 _x) pure returns(string)
func (_Strings *StringsSession) Bytes32ToStr(_x [32]byte) (string, error) {
	return _Strings.Contract.Bytes32ToStr(&_Strings.CallOpts, _x)
}

// Bytes32ToStr is a free data retrieval call binding the contract method 0xef0b2368.
//
// Solidity: function bytes32ToStr(bytes32 _x) pure returns(string)
func (_Strings *StringsCallerSession) Bytes32ToStr(_x [32]byte) (string, error) {
	return _Strings.Contract.Bytes32ToStr(&_Strings.CallOpts, _x)
}

// Compare is a free data retrieval call binding the contract method 0x28602ab2.
//
// Solidity: function compare(bytes32 _a, bytes32 _b) pure returns(bool)
func (_Strings *StringsCaller) Compare(opts *bind.CallOpts, _a [32]byte, _b [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Strings.contract.Call(opts, out, "compare", _a, _b)
	return *ret0, err
}

// Compare is a free data retrieval call binding the contract method 0x28602ab2.
//
// Solidity: function compare(bytes32 _a, bytes32 _b) pure returns(bool)
func (_Strings *StringsSession) Compare(_a [32]byte, _b [32]byte) (bool, error) {
	return _Strings.Contract.Compare(&_Strings.CallOpts, _a, _b)
}

// Compare is a free data retrieval call binding the contract method 0x28602ab2.
//
// Solidity: function compare(bytes32 _a, bytes32 _b) pure returns(bool)
func (_Strings *StringsCallerSession) Compare(_a [32]byte, _b [32]byte) (bool, error) {
	return _Strings.Contract.Compare(&_Strings.CallOpts, _a, _b)
}

// Compare0 is a free data retrieval call binding the contract method 0x3a96fdd7.
//
// Solidity: function compare(string _a, string _b) pure returns(bool)
func (_Strings *StringsCaller) Compare0(opts *bind.CallOpts, _a string, _b string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Strings.contract.Call(opts, out, "compare0", _a, _b)
	return *ret0, err
}

// Compare0 is a free data retrieval call binding the contract method 0x3a96fdd7.
//
// Solidity: function compare(string _a, string _b) pure returns(bool)
func (_Strings *StringsSession) Compare0(_a string, _b string) (bool, error) {
	return _Strings.Contract.Compare0(&_Strings.CallOpts, _a, _b)
}

// Compare0 is a free data retrieval call binding the contract method 0x3a96fdd7.
//
// Solidity: function compare(string _a, string _b) pure returns(bool)
func (_Strings *StringsCallerSession) Compare0(_a string, _b string) (bool, error) {
	return _Strings.Contract.Compare0(&_Strings.CallOpts, _a, _b)
}

// GetExists is a free data retrieval call binding the contract method 0x745fca7b.
//
// Solidity: function getExists(bytes32 _x, bytes32[] _xs) pure returns(bool)
func (_Strings *StringsCaller) GetExists(opts *bind.CallOpts, _x [32]byte, _xs [][32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Strings.contract.Call(opts, out, "getExists", _x, _xs)
	return *ret0, err
}

// GetExists is a free data retrieval call binding the contract method 0x745fca7b.
//
// Solidity: function getExists(bytes32 _x, bytes32[] _xs) pure returns(bool)
func (_Strings *StringsSession) GetExists(_x [32]byte, _xs [][32]byte) (bool, error) {
	return _Strings.Contract.GetExists(&_Strings.CallOpts, _x, _xs)
}

// GetExists is a free data retrieval call binding the contract method 0x745fca7b.
//
// Solidity: function getExists(bytes32 _x, bytes32[] _xs) pure returns(bool)
func (_Strings *StringsCallerSession) GetExists(_x [32]byte, _xs [][32]byte) (bool, error) {
	return _Strings.Contract.GetExists(&_Strings.CallOpts, _x, _xs)
}

// GetIndex is a free data retrieval call binding the contract method 0x02a33464.
//
// Solidity: function getIndex(string _a, string[] _xs) pure returns(uint256)
func (_Strings *StringsCaller) GetIndex(opts *bind.CallOpts, _a string, _xs []string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Strings.contract.Call(opts, out, "getIndex", _a, _xs)
	return *ret0, err
}

// GetIndex is a free data retrieval call binding the contract method 0x02a33464.
//
// Solidity: function getIndex(string _a, string[] _xs) pure returns(uint256)
func (_Strings *StringsSession) GetIndex(_a string, _xs []string) (*big.Int, error) {
	return _Strings.Contract.GetIndex(&_Strings.CallOpts, _a, _xs)
}

// GetIndex is a free data retrieval call binding the contract method 0x02a33464.
//
// Solidity: function getIndex(string _a, string[] _xs) pure returns(uint256)
func (_Strings *StringsCallerSession) GetIndex(_a string, _xs []string) (*big.Int, error) {
	return _Strings.Contract.GetIndex(&_Strings.CallOpts, _a, _xs)
}

// GetIndex0 is a free data retrieval call binding the contract method 0xc55f207f.
//
// Solidity: function getIndex(bytes32 _x, bytes32[] _xs) pure returns(uint256)
func (_Strings *StringsCaller) GetIndex0(opts *bind.CallOpts, _x [32]byte, _xs [][32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Strings.contract.Call(opts, out, "getIndex0", _x, _xs)
	return *ret0, err
}

// GetIndex0 is a free data retrieval call binding the contract method 0xc55f207f.
//
// Solidity: function getIndex(bytes32 _x, bytes32[] _xs) pure returns(uint256)
func (_Strings *StringsSession) GetIndex0(_x [32]byte, _xs [][32]byte) (*big.Int, error) {
	return _Strings.Contract.GetIndex0(&_Strings.CallOpts, _x, _xs)
}

// GetIndex0 is a free data retrieval call binding the contract method 0xc55f207f.
//
// Solidity: function getIndex(bytes32 _x, bytes32[] _xs) pure returns(uint256)
func (_Strings *StringsCallerSession) GetIndex0(_x [32]byte, _xs [][32]byte) (*big.Int, error) {
	return _Strings.Contract.GetIndex0(&_Strings.CallOpts, _x, _xs)
}
