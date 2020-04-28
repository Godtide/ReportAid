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

// ActivityRelatedActivitiesRelatedActivity is an auto generated low-level Go binding around an user-defined struct.
type ActivityRelatedActivitiesRelatedActivity struct {
	RelationType uint8
	ActivityRef  [32]byte
}

// ActivityRelatedActivitiesABI is the input ABI used to generate the binding from.
const ActivityRelatedActivitiesABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_relatedActivityRef\",\"type\":\"bytes32\"}],\"name\":\"getActivityRef\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_relatedActivityRef\",\"type\":\"bytes32\"}],\"name\":\"getRelatedActivity\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"relationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"activityRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityRelatedActivities.RelatedActivity\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_relatedActivityRef\",\"type\":\"bytes32\"}],\"name\":\"getRelationType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_relatedActivityRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"relationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"activityRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityRelatedActivities.RelatedActivity\",\"name\":\"_relatedActivity\",\"type\":\"tuple\"}],\"name\":\"setRelatedActivity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ActivityRelatedActivitiesFuncSigs maps the 4-byte function signature to its string representation.
var ActivityRelatedActivitiesFuncSigs = map[string]string{
	"769956ad": "getActivityRef(bytes32,bytes32,bytes32)",
	"47cdae01": "getNum(bytes32,bytes32)",
	"eef67d78": "getReference(bytes32,bytes32,uint256)",
	"c58e6abc": "getRelatedActivity(bytes32,bytes32,bytes32)",
	"8f11cb91": "getRelationType(bytes32,bytes32,bytes32)",
	"8825b26a": "setRelatedActivity(bytes32,bytes32,bytes32,(uint8,bytes32))",
}

// ActivityRelatedActivities is an auto generated Go binding around an Ethereum contract.
type ActivityRelatedActivities struct {
	ActivityRelatedActivitiesCaller     // Read-only binding to the contract
	ActivityRelatedActivitiesTransactor // Write-only binding to the contract
	ActivityRelatedActivitiesFilterer   // Log filterer for contract events
}

// ActivityRelatedActivitiesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivityRelatedActivitiesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityRelatedActivitiesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivityRelatedActivitiesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityRelatedActivitiesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivityRelatedActivitiesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityRelatedActivitiesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivityRelatedActivitiesSession struct {
	Contract     *ActivityRelatedActivities // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ActivityRelatedActivitiesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivityRelatedActivitiesCallerSession struct {
	Contract *ActivityRelatedActivitiesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ActivityRelatedActivitiesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivityRelatedActivitiesTransactorSession struct {
	Contract     *ActivityRelatedActivitiesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ActivityRelatedActivitiesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivityRelatedActivitiesRaw struct {
	Contract *ActivityRelatedActivities // Generic contract binding to access the raw methods on
}

// ActivityRelatedActivitiesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivityRelatedActivitiesCallerRaw struct {
	Contract *ActivityRelatedActivitiesCaller // Generic read-only contract binding to access the raw methods on
}

// ActivityRelatedActivitiesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivityRelatedActivitiesTransactorRaw struct {
	Contract *ActivityRelatedActivitiesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivityRelatedActivities creates a new instance of ActivityRelatedActivities, bound to a specific deployed contract.
func NewActivityRelatedActivities(address common.Address, backend bind.ContractBackend) (*ActivityRelatedActivities, error) {
	contract, err := bindActivityRelatedActivities(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ActivityRelatedActivities{ActivityRelatedActivitiesCaller: ActivityRelatedActivitiesCaller{contract: contract}, ActivityRelatedActivitiesTransactor: ActivityRelatedActivitiesTransactor{contract: contract}, ActivityRelatedActivitiesFilterer: ActivityRelatedActivitiesFilterer{contract: contract}}, nil
}

// NewActivityRelatedActivitiesCaller creates a new read-only instance of ActivityRelatedActivities, bound to a specific deployed contract.
func NewActivityRelatedActivitiesCaller(address common.Address, caller bind.ContractCaller) (*ActivityRelatedActivitiesCaller, error) {
	contract, err := bindActivityRelatedActivities(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityRelatedActivitiesCaller{contract: contract}, nil
}

// NewActivityRelatedActivitiesTransactor creates a new write-only instance of ActivityRelatedActivities, bound to a specific deployed contract.
func NewActivityRelatedActivitiesTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivityRelatedActivitiesTransactor, error) {
	contract, err := bindActivityRelatedActivities(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityRelatedActivitiesTransactor{contract: contract}, nil
}

// NewActivityRelatedActivitiesFilterer creates a new log filterer instance of ActivityRelatedActivities, bound to a specific deployed contract.
func NewActivityRelatedActivitiesFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivityRelatedActivitiesFilterer, error) {
	contract, err := bindActivityRelatedActivities(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivityRelatedActivitiesFilterer{contract: contract}, nil
}

// bindActivityRelatedActivities binds a generic wrapper to an already deployed contract.
func bindActivityRelatedActivities(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivityRelatedActivitiesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityRelatedActivities *ActivityRelatedActivitiesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivityRelatedActivities.Contract.ActivityRelatedActivitiesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityRelatedActivities *ActivityRelatedActivitiesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityRelatedActivities.Contract.ActivityRelatedActivitiesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityRelatedActivities *ActivityRelatedActivitiesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityRelatedActivities.Contract.ActivityRelatedActivitiesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityRelatedActivities *ActivityRelatedActivitiesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivityRelatedActivities.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityRelatedActivities *ActivityRelatedActivitiesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityRelatedActivities.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityRelatedActivities *ActivityRelatedActivitiesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityRelatedActivities.Contract.contract.Transact(opts, method, params...)
}

// GetActivityRef is a free data retrieval call binding the contract method 0x769956ad.
//
// Solidity: function getActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) view returns(bytes32)
func (_ActivityRelatedActivities *ActivityRelatedActivitiesCaller) GetActivityRef(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityRelatedActivities.contract.Call(opts, out, "getActivityRef", _activitiesRef, _activityRef, _relatedActivityRef)
	return *ret0, err
}

// GetActivityRef is a free data retrieval call binding the contract method 0x769956ad.
//
// Solidity: function getActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) view returns(bytes32)
func (_ActivityRelatedActivities *ActivityRelatedActivitiesSession) GetActivityRef(_activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte) ([32]byte, error) {
	return _ActivityRelatedActivities.Contract.GetActivityRef(&_ActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef, _relatedActivityRef)
}

// GetActivityRef is a free data retrieval call binding the contract method 0x769956ad.
//
// Solidity: function getActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) view returns(bytes32)
func (_ActivityRelatedActivities *ActivityRelatedActivitiesCallerSession) GetActivityRef(_activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte) ([32]byte, error) {
	return _ActivityRelatedActivities.Contract.GetActivityRef(&_ActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef, _relatedActivityRef)
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivityRelatedActivities *ActivityRelatedActivitiesCaller) GetNum(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ActivityRelatedActivities.contract.Call(opts, out, "getNum", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivityRelatedActivities *ActivityRelatedActivitiesSession) GetNum(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _ActivityRelatedActivities.Contract.GetNum(&_ActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef)
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivityRelatedActivities *ActivityRelatedActivitiesCallerSession) GetNum(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _ActivityRelatedActivities.Contract.GetNum(&_ActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivityRelatedActivities *ActivityRelatedActivitiesCaller) GetReference(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityRelatedActivities.contract.Call(opts, out, "getReference", _activitiesRef, _activityRef, _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivityRelatedActivities *ActivityRelatedActivitiesSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _ActivityRelatedActivities.Contract.GetReference(&_ActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivityRelatedActivities *ActivityRelatedActivitiesCallerSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _ActivityRelatedActivities.Contract.GetReference(&_ActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetRelatedActivity is a free data retrieval call binding the contract method 0xc58e6abc.
//
// Solidity: function getRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) view returns(ActivityRelatedActivitiesRelatedActivity)
func (_ActivityRelatedActivities *ActivityRelatedActivitiesCaller) GetRelatedActivity(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte) (ActivityRelatedActivitiesRelatedActivity, error) {
	var (
		ret0 = new(ActivityRelatedActivitiesRelatedActivity)
	)
	out := ret0
	err := _ActivityRelatedActivities.contract.Call(opts, out, "getRelatedActivity", _activitiesRef, _activityRef, _relatedActivityRef)
	return *ret0, err
}

// GetRelatedActivity is a free data retrieval call binding the contract method 0xc58e6abc.
//
// Solidity: function getRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) view returns(ActivityRelatedActivitiesRelatedActivity)
func (_ActivityRelatedActivities *ActivityRelatedActivitiesSession) GetRelatedActivity(_activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte) (ActivityRelatedActivitiesRelatedActivity, error) {
	return _ActivityRelatedActivities.Contract.GetRelatedActivity(&_ActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef, _relatedActivityRef)
}

// GetRelatedActivity is a free data retrieval call binding the contract method 0xc58e6abc.
//
// Solidity: function getRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) view returns(ActivityRelatedActivitiesRelatedActivity)
func (_ActivityRelatedActivities *ActivityRelatedActivitiesCallerSession) GetRelatedActivity(_activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte) (ActivityRelatedActivitiesRelatedActivity, error) {
	return _ActivityRelatedActivities.Contract.GetRelatedActivity(&_ActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef, _relatedActivityRef)
}

// GetRelationType is a free data retrieval call binding the contract method 0x8f11cb91.
//
// Solidity: function getRelationType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) view returns(uint8)
func (_ActivityRelatedActivities *ActivityRelatedActivitiesCaller) GetRelationType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ActivityRelatedActivities.contract.Call(opts, out, "getRelationType", _activitiesRef, _activityRef, _relatedActivityRef)
	return *ret0, err
}

// GetRelationType is a free data retrieval call binding the contract method 0x8f11cb91.
//
// Solidity: function getRelationType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) view returns(uint8)
func (_ActivityRelatedActivities *ActivityRelatedActivitiesSession) GetRelationType(_activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte) (uint8, error) {
	return _ActivityRelatedActivities.Contract.GetRelationType(&_ActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef, _relatedActivityRef)
}

// GetRelationType is a free data retrieval call binding the contract method 0x8f11cb91.
//
// Solidity: function getRelationType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) view returns(uint8)
func (_ActivityRelatedActivities *ActivityRelatedActivitiesCallerSession) GetRelationType(_activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte) (uint8, error) {
	return _ActivityRelatedActivities.Contract.GetRelationType(&_ActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef, _relatedActivityRef)
}

// SetRelatedActivity is a paid mutator transaction binding the contract method 0x8825b26a.
//
// Solidity: function setRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef, ActivityRelatedActivitiesRelatedActivity _relatedActivity) returns()
func (_ActivityRelatedActivities *ActivityRelatedActivitiesTransactor) SetRelatedActivity(opts *bind.TransactOpts, _activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte, _relatedActivity ActivityRelatedActivitiesRelatedActivity) (*types.Transaction, error) {
	return _ActivityRelatedActivities.contract.Transact(opts, "setRelatedActivity", _activitiesRef, _activityRef, _relatedActivityRef, _relatedActivity)
}

// SetRelatedActivity is a paid mutator transaction binding the contract method 0x8825b26a.
//
// Solidity: function setRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef, ActivityRelatedActivitiesRelatedActivity _relatedActivity) returns()
func (_ActivityRelatedActivities *ActivityRelatedActivitiesSession) SetRelatedActivity(_activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte, _relatedActivity ActivityRelatedActivitiesRelatedActivity) (*types.Transaction, error) {
	return _ActivityRelatedActivities.Contract.SetRelatedActivity(&_ActivityRelatedActivities.TransactOpts, _activitiesRef, _activityRef, _relatedActivityRef, _relatedActivity)
}

// SetRelatedActivity is a paid mutator transaction binding the contract method 0x8825b26a.
//
// Solidity: function setRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef, ActivityRelatedActivitiesRelatedActivity _relatedActivity) returns()
func (_ActivityRelatedActivities *ActivityRelatedActivitiesTransactorSession) SetRelatedActivity(_activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte, _relatedActivity ActivityRelatedActivitiesRelatedActivity) (*types.Transaction, error) {
	return _ActivityRelatedActivities.Contract.SetRelatedActivity(&_ActivityRelatedActivities.TransactOpts, _activitiesRef, _activityRef, _relatedActivityRef, _relatedActivity)
}

// IATIActivityRelatedActivitiesABI is the input ABI used to generate the binding from.
const IATIActivityRelatedActivitiesABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_relatedActivityRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"relationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"activityRef\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structActivityRelatedActivities.RelatedActivity\",\"name\":\"_relatedActivity\",\"type\":\"tuple\"}],\"name\":\"SetRelatedActivity\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_relatedActivityRef\",\"type\":\"bytes32\"}],\"name\":\"getActivityRef\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_relatedActivityRef\",\"type\":\"bytes32\"}],\"name\":\"getRelatedActivity\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"relationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"activityRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityRelatedActivities.RelatedActivity\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_relatedActivityRef\",\"type\":\"bytes32\"}],\"name\":\"getRelationType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_relatedActivityRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"relationType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"activityRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityRelatedActivities.RelatedActivity\",\"name\":\"_relatedActivity\",\"type\":\"tuple\"}],\"name\":\"setRelatedActivity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IATIActivityRelatedActivitiesFuncSigs maps the 4-byte function signature to its string representation.
var IATIActivityRelatedActivitiesFuncSigs = map[string]string{
	"769956ad": "getActivityRef(bytes32,bytes32,bytes32)",
	"47cdae01": "getNum(bytes32,bytes32)",
	"eef67d78": "getReference(bytes32,bytes32,uint256)",
	"c58e6abc": "getRelatedActivity(bytes32,bytes32,bytes32)",
	"8f11cb91": "getRelationType(bytes32,bytes32,bytes32)",
	"8825b26a": "setRelatedActivity(bytes32,bytes32,bytes32,(uint8,bytes32))",
}

// IATIActivityRelatedActivitiesBin is the compiled bytecode used for deploying new contracts.
var IATIActivityRelatedActivitiesBin = "0x608060405234801561001057600080fd5b50610767806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806347cdae0114610067578063769956ad146100905780638825b26a146100a35780638f11cb91146100b8578063c58e6abc146100d8578063eef67d78146100f8575b600080fd5b61007a610075366004610598565b61010b565b6040516100879190610691565b60405180910390f35b61007a61009e3660046105b9565b610163565b6100b66100b13660046105e4565b6101df565b005b6100cb6100c63660046105b9565b6103b3565b6040516100879190610723565b6100eb6100e63660046105b9565b61042e565b6040516100879190610715565b61007a61010636600461066b565b6104c9565b600082811a60f81b6001600160f81b0319161580159061013a57508160001a60f81b6001600160f81b03191615155b61014357600080fd5b506000828152602081815260408083208484529091529020545b92915050565b600083811a60f81b6001600160f81b0319161580159061019257508260001a60f81b6001600160f81b03191615155b80156101ad57508160001a60f81b6001600160f81b03191615155b6101b657600080fd5b506000928352600160208181526040808620948652938152838520928552919091529120015490565b8360001a60f81b6001600160f81b0319161580159061020d57508260001a60f81b6001600160f81b03191615155b801561022857508160001a60f81b6001600160f81b03191615155b80156102475750602081015160001a60f81b6001600160f81b03191615155b80156102565750805160ff1615155b801561026857508051600660ff909116105b61027157600080fd5b6000848152600160208181526040808420878552825280842086855282528084208551815460ff191660ff90911617815585830151930192909255868352828152818320868452905290819020905163745fca7b60e01b815273__$b620240993a55615b607189176fb82e62f$__9163745fca7b916102f491869160040161069a565b60206040518083038186803b15801561030c57600080fd5b505af4158015610320573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103449190610571565b610370576000848152602081815260408083208684528252822080546001810182559083529120018290555b7fd4dfb4963ae0de7dd6d1dd9eb94d64c7956bb84b940a01667894134a0d7f08fa848484846040516103a594939291906106ea565b60405180910390a150505050565b600083811a60f81b6001600160f81b031916158015906103e257508260001a60f81b6001600160f81b03191615155b80156103fd57508160001a60f81b6001600160f81b03191615155b61040657600080fd5b5060009283526001602090815260408085209385529281528284209184525290205460ff1690565b61043661055a565b8360001a60f81b6001600160f81b0319161580159061046457508260001a60f81b6001600160f81b03191615155b801561047f57508160001a60f81b6001600160f81b03191615155b61048857600080fd5b50600083815260016020818152604080842086855282528084208585528252928390208351808501909452805460ff16845290910154908201529392505050565b600083811a60f81b6001600160f81b031916158015906104f857508260001a60f81b6001600160f81b03191615155b801561051a575060008481526020818152604080832086845290915290205482105b61052357600080fd5b600084815260208181526040808320868452909152902080548390811061054657fe5b906000526020600020015490509392505050565b604080518082019091526000808252602082015290565b600060208284031215610582578081fd5b81518015158114610591578182fd5b9392505050565b600080604083850312156105aa578081fd5b50508035926020909101359150565b6000806000606084860312156105cd578081fd5b505081359360208301359350604090920135919050565b60008060008084860360a08112156105fa578182fd5b85359450602086013593506040808701359350605f198201121561061c578182fd5b506040516040810181811067ffffffffffffffff8211171561063c578283fd5b604052606086013560ff81168114610652578283fd5b8152608095909501356020860152509194909350909190565b6000806000606084860312156105cd578283fd5b805160ff168252602090810151910152565b90815260200190565b60006040820184835260206040818501528185548084526060860191508685528285209350845b818110156106dd578454835260019485019492840192016106c1565b5090979650505050505050565b848152602081018490526040810183905260a0810161070c606083018461067f565b95945050505050565b6040810161015d828461067f565b60ff9190911681526020019056fea2646970667358221220b1a069a35a766657aa4598ca77073d591927148a8075e4a543ed46ad4d84309064736f6c63430006060033"

// DeployIATIActivityRelatedActivities deploys a new Ethereum contract, binding an instance of IATIActivityRelatedActivities to it.
func DeployIATIActivityRelatedActivities(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IATIActivityRelatedActivities, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivityRelatedActivitiesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	stringsAddr, _, _, _ := DeployStrings(auth, backend)
	IATIActivityRelatedActivitiesBin = strings.Replace(IATIActivityRelatedActivitiesBin, "__$b620240993a55615b607189176fb82e62f$__", stringsAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IATIActivityRelatedActivitiesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IATIActivityRelatedActivities{IATIActivityRelatedActivitiesCaller: IATIActivityRelatedActivitiesCaller{contract: contract}, IATIActivityRelatedActivitiesTransactor: IATIActivityRelatedActivitiesTransactor{contract: contract}, IATIActivityRelatedActivitiesFilterer: IATIActivityRelatedActivitiesFilterer{contract: contract}}, nil
}

// IATIActivityRelatedActivities is an auto generated Go binding around an Ethereum contract.
type IATIActivityRelatedActivities struct {
	IATIActivityRelatedActivitiesCaller     // Read-only binding to the contract
	IATIActivityRelatedActivitiesTransactor // Write-only binding to the contract
	IATIActivityRelatedActivitiesFilterer   // Log filterer for contract events
}

// IATIActivityRelatedActivitiesCaller is an auto generated read-only Go binding around an Ethereum contract.
type IATIActivityRelatedActivitiesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityRelatedActivitiesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IATIActivityRelatedActivitiesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityRelatedActivitiesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IATIActivityRelatedActivitiesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityRelatedActivitiesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IATIActivityRelatedActivitiesSession struct {
	Contract     *IATIActivityRelatedActivities // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IATIActivityRelatedActivitiesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IATIActivityRelatedActivitiesCallerSession struct {
	Contract *IATIActivityRelatedActivitiesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// IATIActivityRelatedActivitiesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IATIActivityRelatedActivitiesTransactorSession struct {
	Contract     *IATIActivityRelatedActivitiesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// IATIActivityRelatedActivitiesRaw is an auto generated low-level Go binding around an Ethereum contract.
type IATIActivityRelatedActivitiesRaw struct {
	Contract *IATIActivityRelatedActivities // Generic contract binding to access the raw methods on
}

// IATIActivityRelatedActivitiesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IATIActivityRelatedActivitiesCallerRaw struct {
	Contract *IATIActivityRelatedActivitiesCaller // Generic read-only contract binding to access the raw methods on
}

// IATIActivityRelatedActivitiesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IATIActivityRelatedActivitiesTransactorRaw struct {
	Contract *IATIActivityRelatedActivitiesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIATIActivityRelatedActivities creates a new instance of IATIActivityRelatedActivities, bound to a specific deployed contract.
func NewIATIActivityRelatedActivities(address common.Address, backend bind.ContractBackend) (*IATIActivityRelatedActivities, error) {
	contract, err := bindIATIActivityRelatedActivities(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IATIActivityRelatedActivities{IATIActivityRelatedActivitiesCaller: IATIActivityRelatedActivitiesCaller{contract: contract}, IATIActivityRelatedActivitiesTransactor: IATIActivityRelatedActivitiesTransactor{contract: contract}, IATIActivityRelatedActivitiesFilterer: IATIActivityRelatedActivitiesFilterer{contract: contract}}, nil
}

// NewIATIActivityRelatedActivitiesCaller creates a new read-only instance of IATIActivityRelatedActivities, bound to a specific deployed contract.
func NewIATIActivityRelatedActivitiesCaller(address common.Address, caller bind.ContractCaller) (*IATIActivityRelatedActivitiesCaller, error) {
	contract, err := bindIATIActivityRelatedActivities(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IATIActivityRelatedActivitiesCaller{contract: contract}, nil
}

// NewIATIActivityRelatedActivitiesTransactor creates a new write-only instance of IATIActivityRelatedActivities, bound to a specific deployed contract.
func NewIATIActivityRelatedActivitiesTransactor(address common.Address, transactor bind.ContractTransactor) (*IATIActivityRelatedActivitiesTransactor, error) {
	contract, err := bindIATIActivityRelatedActivities(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IATIActivityRelatedActivitiesTransactor{contract: contract}, nil
}

// NewIATIActivityRelatedActivitiesFilterer creates a new log filterer instance of IATIActivityRelatedActivities, bound to a specific deployed contract.
func NewIATIActivityRelatedActivitiesFilterer(address common.Address, filterer bind.ContractFilterer) (*IATIActivityRelatedActivitiesFilterer, error) {
	contract, err := bindIATIActivityRelatedActivities(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IATIActivityRelatedActivitiesFilterer{contract: contract}, nil
}

// bindIATIActivityRelatedActivities binds a generic wrapper to an already deployed contract.
func bindIATIActivityRelatedActivities(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivityRelatedActivitiesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIActivityRelatedActivities.Contract.IATIActivityRelatedActivitiesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIActivityRelatedActivities.Contract.IATIActivityRelatedActivitiesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIActivityRelatedActivities.Contract.IATIActivityRelatedActivitiesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIActivityRelatedActivities.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIActivityRelatedActivities.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIActivityRelatedActivities.Contract.contract.Transact(opts, method, params...)
}

// GetActivityRef is a free data retrieval call binding the contract method 0x769956ad.
//
// Solidity: function getActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) view returns(bytes32)
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesCaller) GetActivityRef(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityRelatedActivities.contract.Call(opts, out, "getActivityRef", _activitiesRef, _activityRef, _relatedActivityRef)
	return *ret0, err
}

// GetActivityRef is a free data retrieval call binding the contract method 0x769956ad.
//
// Solidity: function getActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) view returns(bytes32)
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesSession) GetActivityRef(_activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte) ([32]byte, error) {
	return _IATIActivityRelatedActivities.Contract.GetActivityRef(&_IATIActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef, _relatedActivityRef)
}

// GetActivityRef is a free data retrieval call binding the contract method 0x769956ad.
//
// Solidity: function getActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) view returns(bytes32)
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesCallerSession) GetActivityRef(_activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte) ([32]byte, error) {
	return _IATIActivityRelatedActivities.Contract.GetActivityRef(&_IATIActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef, _relatedActivityRef)
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesCaller) GetNum(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIActivityRelatedActivities.contract.Call(opts, out, "getNum", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesSession) GetNum(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _IATIActivityRelatedActivities.Contract.GetNum(&_IATIActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef)
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesCallerSession) GetNum(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _IATIActivityRelatedActivities.Contract.GetNum(&_IATIActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesCaller) GetReference(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityRelatedActivities.contract.Call(opts, out, "getReference", _activitiesRef, _activityRef, _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIActivityRelatedActivities.Contract.GetReference(&_IATIActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesCallerSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIActivityRelatedActivities.Contract.GetReference(&_IATIActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetRelatedActivity is a free data retrieval call binding the contract method 0xc58e6abc.
//
// Solidity: function getRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) view returns(ActivityRelatedActivitiesRelatedActivity)
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesCaller) GetRelatedActivity(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte) (ActivityRelatedActivitiesRelatedActivity, error) {
	var (
		ret0 = new(ActivityRelatedActivitiesRelatedActivity)
	)
	out := ret0
	err := _IATIActivityRelatedActivities.contract.Call(opts, out, "getRelatedActivity", _activitiesRef, _activityRef, _relatedActivityRef)
	return *ret0, err
}

// GetRelatedActivity is a free data retrieval call binding the contract method 0xc58e6abc.
//
// Solidity: function getRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) view returns(ActivityRelatedActivitiesRelatedActivity)
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesSession) GetRelatedActivity(_activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte) (ActivityRelatedActivitiesRelatedActivity, error) {
	return _IATIActivityRelatedActivities.Contract.GetRelatedActivity(&_IATIActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef, _relatedActivityRef)
}

// GetRelatedActivity is a free data retrieval call binding the contract method 0xc58e6abc.
//
// Solidity: function getRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) view returns(ActivityRelatedActivitiesRelatedActivity)
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesCallerSession) GetRelatedActivity(_activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte) (ActivityRelatedActivitiesRelatedActivity, error) {
	return _IATIActivityRelatedActivities.Contract.GetRelatedActivity(&_IATIActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef, _relatedActivityRef)
}

// GetRelationType is a free data retrieval call binding the contract method 0x8f11cb91.
//
// Solidity: function getRelationType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) view returns(uint8)
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesCaller) GetRelationType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivityRelatedActivities.contract.Call(opts, out, "getRelationType", _activitiesRef, _activityRef, _relatedActivityRef)
	return *ret0, err
}

// GetRelationType is a free data retrieval call binding the contract method 0x8f11cb91.
//
// Solidity: function getRelationType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) view returns(uint8)
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesSession) GetRelationType(_activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte) (uint8, error) {
	return _IATIActivityRelatedActivities.Contract.GetRelationType(&_IATIActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef, _relatedActivityRef)
}

// GetRelationType is a free data retrieval call binding the contract method 0x8f11cb91.
//
// Solidity: function getRelationType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef) view returns(uint8)
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesCallerSession) GetRelationType(_activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte) (uint8, error) {
	return _IATIActivityRelatedActivities.Contract.GetRelationType(&_IATIActivityRelatedActivities.CallOpts, _activitiesRef, _activityRef, _relatedActivityRef)
}

// SetRelatedActivity is a paid mutator transaction binding the contract method 0x8825b26a.
//
// Solidity: function setRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef, ActivityRelatedActivitiesRelatedActivity _relatedActivity) returns()
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesTransactor) SetRelatedActivity(opts *bind.TransactOpts, _activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte, _relatedActivity ActivityRelatedActivitiesRelatedActivity) (*types.Transaction, error) {
	return _IATIActivityRelatedActivities.contract.Transact(opts, "setRelatedActivity", _activitiesRef, _activityRef, _relatedActivityRef, _relatedActivity)
}

// SetRelatedActivity is a paid mutator transaction binding the contract method 0x8825b26a.
//
// Solidity: function setRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef, ActivityRelatedActivitiesRelatedActivity _relatedActivity) returns()
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesSession) SetRelatedActivity(_activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte, _relatedActivity ActivityRelatedActivitiesRelatedActivity) (*types.Transaction, error) {
	return _IATIActivityRelatedActivities.Contract.SetRelatedActivity(&_IATIActivityRelatedActivities.TransactOpts, _activitiesRef, _activityRef, _relatedActivityRef, _relatedActivity)
}

// SetRelatedActivity is a paid mutator transaction binding the contract method 0x8825b26a.
//
// Solidity: function setRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef, ActivityRelatedActivitiesRelatedActivity _relatedActivity) returns()
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesTransactorSession) SetRelatedActivity(_activitiesRef [32]byte, _activityRef [32]byte, _relatedActivityRef [32]byte, _relatedActivity ActivityRelatedActivitiesRelatedActivity) (*types.Transaction, error) {
	return _IATIActivityRelatedActivities.Contract.SetRelatedActivity(&_IATIActivityRelatedActivities.TransactOpts, _activitiesRef, _activityRef, _relatedActivityRef, _relatedActivity)
}

// IATIActivityRelatedActivitiesSetRelatedActivityIterator is returned from FilterSetRelatedActivity and is used to iterate over the raw logs and unpacked data for SetRelatedActivity events raised by the IATIActivityRelatedActivities contract.
type IATIActivityRelatedActivitiesSetRelatedActivityIterator struct {
	Event *IATIActivityRelatedActivitiesSetRelatedActivity // Event containing the contract specifics and raw log

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
func (it *IATIActivityRelatedActivitiesSetRelatedActivityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IATIActivityRelatedActivitiesSetRelatedActivity)
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
		it.Event = new(IATIActivityRelatedActivitiesSetRelatedActivity)
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
func (it *IATIActivityRelatedActivitiesSetRelatedActivityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IATIActivityRelatedActivitiesSetRelatedActivityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IATIActivityRelatedActivitiesSetRelatedActivity represents a SetRelatedActivity event raised by the IATIActivityRelatedActivities contract.
type IATIActivityRelatedActivitiesSetRelatedActivity struct {
	ActivitiesRef      [32]byte
	ActivityRef        [32]byte
	RelatedActivityRef [32]byte
	RelatedActivity    ActivityRelatedActivitiesRelatedActivity
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetRelatedActivity is a free log retrieval operation binding the contract event 0xd4dfb4963ae0de7dd6d1dd9eb94d64c7956bb84b940a01667894134a0d7f08fa.
//
// Solidity: event SetRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef, ActivityRelatedActivitiesRelatedActivity _relatedActivity)
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesFilterer) FilterSetRelatedActivity(opts *bind.FilterOpts) (*IATIActivityRelatedActivitiesSetRelatedActivityIterator, error) {

	logs, sub, err := _IATIActivityRelatedActivities.contract.FilterLogs(opts, "SetRelatedActivity")
	if err != nil {
		return nil, err
	}
	return &IATIActivityRelatedActivitiesSetRelatedActivityIterator{contract: _IATIActivityRelatedActivities.contract, event: "SetRelatedActivity", logs: logs, sub: sub}, nil
}

// WatchSetRelatedActivity is a free log subscription operation binding the contract event 0xd4dfb4963ae0de7dd6d1dd9eb94d64c7956bb84b940a01667894134a0d7f08fa.
//
// Solidity: event SetRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef, ActivityRelatedActivitiesRelatedActivity _relatedActivity)
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesFilterer) WatchSetRelatedActivity(opts *bind.WatchOpts, sink chan<- *IATIActivityRelatedActivitiesSetRelatedActivity) (event.Subscription, error) {

	logs, sub, err := _IATIActivityRelatedActivities.contract.WatchLogs(opts, "SetRelatedActivity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IATIActivityRelatedActivitiesSetRelatedActivity)
				if err := _IATIActivityRelatedActivities.contract.UnpackLog(event, "SetRelatedActivity", log); err != nil {
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

// ParseSetRelatedActivity is a log parse operation binding the contract event 0xd4dfb4963ae0de7dd6d1dd9eb94d64c7956bb84b940a01667894134a0d7f08fa.
//
// Solidity: event SetRelatedActivity(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _relatedActivityRef, ActivityRelatedActivitiesRelatedActivity _relatedActivity)
func (_IATIActivityRelatedActivities *IATIActivityRelatedActivitiesFilterer) ParseSetRelatedActivity(log types.Log) (*IATIActivityRelatedActivitiesSetRelatedActivity, error) {
	event := new(IATIActivityRelatedActivitiesSetRelatedActivity)
	if err := _IATIActivityRelatedActivities.contract.UnpackLog(event, "SetRelatedActivity", log); err != nil {
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
