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

// ActivityOrgActivity is an auto generated low-level Go binding around an user-defined struct.
type ActivityOrgActivity struct {
	Humanitarian      bool
	Hierarchy         uint8
	Status            uint8
	BudgetNotProvided uint8
	ReportingOrg      ActivityReportingOrg
	LastUpdated       [32]byte
	Lang              [32]byte
	Currency          [32]byte
	LinkedData        [32]byte
	Identifier        [32]byte
	Title             string
	Description       string
}

// ActivityReportingOrg is an auto generated low-level Go binding around an user-defined struct.
type ActivityReportingOrg struct {
	OrgType     uint8
	IsSecondary bool
	OrgRef      [32]byte
}

// ActivityABI is the input ABI used to generate the binding from.
const ActivityABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getActivity\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"humanitarian\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"hierarchy\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"budgetNotProvided\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSecondary\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivity.ReportingOrg\",\"name\":\"reportingOrg\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"lastUpdated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivity.OrgActivity\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getBudgetNotProvided\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getCurrency\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getDescription\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getHierarchy\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getHumanitarian\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getIdentifier\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getLang\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getLastUpdatedTime\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getLinkedData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"}],\"name\":\"getNumActivities\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getReportingOrg\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSecondary\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivity.ReportingOrg\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getTitle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"activityRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"humanitarian\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"hierarchy\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"budgetNotProvided\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSecondary\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivity.ReportingOrg\",\"name\":\"reportingOrg\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"lastUpdated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivity.OrgActivity\",\"name\":\"_activity\",\"type\":\"tuple\"}],\"name\":\"setActivity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ActivityFuncSigs maps the 4-byte function signature to its string representation.
var ActivityFuncSigs = map[string]string{
	"5aaf1a2d": "getActivity(bytes32,bytes32)",
	"a2f640a5": "getBudgetNotProvided(bytes32,bytes32)",
	"eef45eae": "getCurrency(bytes32,bytes32)",
	"87f9ba4a": "getDescription(bytes32,bytes32)",
	"b863fc70": "getHierarchy(bytes32,bytes32)",
	"1a95155e": "getHumanitarian(bytes32,bytes32)",
	"2b1cb7d5": "getIdentifier(bytes32,bytes32)",
	"aad88871": "getLang(bytes32,bytes32)",
	"7d78621b": "getLastUpdatedTime(bytes32,bytes32)",
	"10b132d5": "getLinkedData(bytes32,bytes32)",
	"db200f3f": "getNumActivities(bytes32)",
	"9db9b0f6": "getReference(bytes32,uint256)",
	"40c21069": "getReportingOrg(bytes32,bytes32)",
	"8e56ee2c": "getStatus(bytes32,bytes32)",
	"8adc94fb": "getTitle(bytes32,bytes32)",
	"95b7378b": "setActivity(bytes32,bytes32,(bool,uint8,uint8,uint8,(uint8,bool,bytes32),bytes32,bytes32,bytes32,bytes32,bytes32,string,string))",
}

// Activity is an auto generated Go binding around an Ethereum contract.
type Activity struct {
	ActivityCaller     // Read-only binding to the contract
	ActivityTransactor // Write-only binding to the contract
	ActivityFilterer   // Log filterer for contract events
}

// ActivityCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivitySession struct {
	Contract     *Activity         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ActivityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivityCallerSession struct {
	Contract *ActivityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ActivityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivityTransactorSession struct {
	Contract     *ActivityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ActivityRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivityRaw struct {
	Contract *Activity // Generic contract binding to access the raw methods on
}

// ActivityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivityCallerRaw struct {
	Contract *ActivityCaller // Generic read-only contract binding to access the raw methods on
}

// ActivityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivityTransactorRaw struct {
	Contract *ActivityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivity creates a new instance of Activity, bound to a specific deployed contract.
func NewActivity(address common.Address, backend bind.ContractBackend) (*Activity, error) {
	contract, err := bindActivity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Activity{ActivityCaller: ActivityCaller{contract: contract}, ActivityTransactor: ActivityTransactor{contract: contract}, ActivityFilterer: ActivityFilterer{contract: contract}}, nil
}

// NewActivityCaller creates a new read-only instance of Activity, bound to a specific deployed contract.
func NewActivityCaller(address common.Address, caller bind.ContractCaller) (*ActivityCaller, error) {
	contract, err := bindActivity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityCaller{contract: contract}, nil
}

// NewActivityTransactor creates a new write-only instance of Activity, bound to a specific deployed contract.
func NewActivityTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivityTransactor, error) {
	contract, err := bindActivity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityTransactor{contract: contract}, nil
}

// NewActivityFilterer creates a new log filterer instance of Activity, bound to a specific deployed contract.
func NewActivityFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivityFilterer, error) {
	contract, err := bindActivity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivityFilterer{contract: contract}, nil
}

// bindActivity binds a generic wrapper to an already deployed contract.
func bindActivity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Activity *ActivityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Activity.Contract.ActivityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Activity *ActivityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activity.Contract.ActivityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Activity *ActivityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Activity.Contract.ActivityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Activity *ActivityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Activity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Activity *ActivityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Activity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Activity *ActivityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Activity.Contract.contract.Transact(opts, method, params...)
}

// GetActivity is a free data retrieval call binding the contract method 0x5aaf1a2d.
//
// Solidity: function getActivity(bytes32 _activitiesRef, bytes32 _activityRef) view returns(ActivityOrgActivity)
func (_Activity *ActivityCaller) GetActivity(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (ActivityOrgActivity, error) {
	var (
		ret0 = new(ActivityOrgActivity)
	)
	out := ret0
	err := _Activity.contract.Call(opts, out, "getActivity", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetActivity is a free data retrieval call binding the contract method 0x5aaf1a2d.
//
// Solidity: function getActivity(bytes32 _activitiesRef, bytes32 _activityRef) view returns(ActivityOrgActivity)
func (_Activity *ActivitySession) GetActivity(_activitiesRef [32]byte, _activityRef [32]byte) (ActivityOrgActivity, error) {
	return _Activity.Contract.GetActivity(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetActivity is a free data retrieval call binding the contract method 0x5aaf1a2d.
//
// Solidity: function getActivity(bytes32 _activitiesRef, bytes32 _activityRef) view returns(ActivityOrgActivity)
func (_Activity *ActivityCallerSession) GetActivity(_activitiesRef [32]byte, _activityRef [32]byte) (ActivityOrgActivity, error) {
	return _Activity.Contract.GetActivity(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetBudgetNotProvided is a free data retrieval call binding the contract method 0xa2f640a5.
//
// Solidity: function getBudgetNotProvided(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint8)
func (_Activity *ActivityCaller) GetBudgetNotProvided(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Activity.contract.Call(opts, out, "getBudgetNotProvided", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetBudgetNotProvided is a free data retrieval call binding the contract method 0xa2f640a5.
//
// Solidity: function getBudgetNotProvided(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint8)
func (_Activity *ActivitySession) GetBudgetNotProvided(_activitiesRef [32]byte, _activityRef [32]byte) (uint8, error) {
	return _Activity.Contract.GetBudgetNotProvided(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetBudgetNotProvided is a free data retrieval call binding the contract method 0xa2f640a5.
//
// Solidity: function getBudgetNotProvided(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint8)
func (_Activity *ActivityCallerSession) GetBudgetNotProvided(_activitiesRef [32]byte, _activityRef [32]byte) (uint8, error) {
	return _Activity.Contract.GetBudgetNotProvided(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetCurrency is a free data retrieval call binding the contract method 0xeef45eae.
//
// Solidity: function getCurrency(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_Activity *ActivityCaller) GetCurrency(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Activity.contract.Call(opts, out, "getCurrency", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetCurrency is a free data retrieval call binding the contract method 0xeef45eae.
//
// Solidity: function getCurrency(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_Activity *ActivitySession) GetCurrency(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _Activity.Contract.GetCurrency(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetCurrency is a free data retrieval call binding the contract method 0xeef45eae.
//
// Solidity: function getCurrency(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_Activity *ActivityCallerSession) GetCurrency(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _Activity.Contract.GetCurrency(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetDescription is a free data retrieval call binding the contract method 0x87f9ba4a.
//
// Solidity: function getDescription(bytes32 _activitiesRef, bytes32 _activityRef) view returns(string)
func (_Activity *ActivityCaller) GetDescription(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Activity.contract.Call(opts, out, "getDescription", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetDescription is a free data retrieval call binding the contract method 0x87f9ba4a.
//
// Solidity: function getDescription(bytes32 _activitiesRef, bytes32 _activityRef) view returns(string)
func (_Activity *ActivitySession) GetDescription(_activitiesRef [32]byte, _activityRef [32]byte) (string, error) {
	return _Activity.Contract.GetDescription(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetDescription is a free data retrieval call binding the contract method 0x87f9ba4a.
//
// Solidity: function getDescription(bytes32 _activitiesRef, bytes32 _activityRef) view returns(string)
func (_Activity *ActivityCallerSession) GetDescription(_activitiesRef [32]byte, _activityRef [32]byte) (string, error) {
	return _Activity.Contract.GetDescription(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetHierarchy is a free data retrieval call binding the contract method 0xb863fc70.
//
// Solidity: function getHierarchy(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint8)
func (_Activity *ActivityCaller) GetHierarchy(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Activity.contract.Call(opts, out, "getHierarchy", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetHierarchy is a free data retrieval call binding the contract method 0xb863fc70.
//
// Solidity: function getHierarchy(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint8)
func (_Activity *ActivitySession) GetHierarchy(_activitiesRef [32]byte, _activityRef [32]byte) (uint8, error) {
	return _Activity.Contract.GetHierarchy(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetHierarchy is a free data retrieval call binding the contract method 0xb863fc70.
//
// Solidity: function getHierarchy(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint8)
func (_Activity *ActivityCallerSession) GetHierarchy(_activitiesRef [32]byte, _activityRef [32]byte) (uint8, error) {
	return _Activity.Contract.GetHierarchy(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetHumanitarian is a free data retrieval call binding the contract method 0x1a95155e.
//
// Solidity: function getHumanitarian(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bool)
func (_Activity *ActivityCaller) GetHumanitarian(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Activity.contract.Call(opts, out, "getHumanitarian", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetHumanitarian is a free data retrieval call binding the contract method 0x1a95155e.
//
// Solidity: function getHumanitarian(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bool)
func (_Activity *ActivitySession) GetHumanitarian(_activitiesRef [32]byte, _activityRef [32]byte) (bool, error) {
	return _Activity.Contract.GetHumanitarian(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetHumanitarian is a free data retrieval call binding the contract method 0x1a95155e.
//
// Solidity: function getHumanitarian(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bool)
func (_Activity *ActivityCallerSession) GetHumanitarian(_activitiesRef [32]byte, _activityRef [32]byte) (bool, error) {
	return _Activity.Contract.GetHumanitarian(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetIdentifier is a free data retrieval call binding the contract method 0x2b1cb7d5.
//
// Solidity: function getIdentifier(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_Activity *ActivityCaller) GetIdentifier(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Activity.contract.Call(opts, out, "getIdentifier", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetIdentifier is a free data retrieval call binding the contract method 0x2b1cb7d5.
//
// Solidity: function getIdentifier(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_Activity *ActivitySession) GetIdentifier(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _Activity.Contract.GetIdentifier(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetIdentifier is a free data retrieval call binding the contract method 0x2b1cb7d5.
//
// Solidity: function getIdentifier(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_Activity *ActivityCallerSession) GetIdentifier(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _Activity.Contract.GetIdentifier(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetLang is a free data retrieval call binding the contract method 0xaad88871.
//
// Solidity: function getLang(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_Activity *ActivityCaller) GetLang(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Activity.contract.Call(opts, out, "getLang", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetLang is a free data retrieval call binding the contract method 0xaad88871.
//
// Solidity: function getLang(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_Activity *ActivitySession) GetLang(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _Activity.Contract.GetLang(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetLang is a free data retrieval call binding the contract method 0xaad88871.
//
// Solidity: function getLang(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_Activity *ActivityCallerSession) GetLang(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _Activity.Contract.GetLang(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetLastUpdatedTime is a free data retrieval call binding the contract method 0x7d78621b.
//
// Solidity: function getLastUpdatedTime(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_Activity *ActivityCaller) GetLastUpdatedTime(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Activity.contract.Call(opts, out, "getLastUpdatedTime", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetLastUpdatedTime is a free data retrieval call binding the contract method 0x7d78621b.
//
// Solidity: function getLastUpdatedTime(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_Activity *ActivitySession) GetLastUpdatedTime(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _Activity.Contract.GetLastUpdatedTime(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetLastUpdatedTime is a free data retrieval call binding the contract method 0x7d78621b.
//
// Solidity: function getLastUpdatedTime(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_Activity *ActivityCallerSession) GetLastUpdatedTime(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _Activity.Contract.GetLastUpdatedTime(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetLinkedData is a free data retrieval call binding the contract method 0x10b132d5.
//
// Solidity: function getLinkedData(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_Activity *ActivityCaller) GetLinkedData(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Activity.contract.Call(opts, out, "getLinkedData", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetLinkedData is a free data retrieval call binding the contract method 0x10b132d5.
//
// Solidity: function getLinkedData(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_Activity *ActivitySession) GetLinkedData(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _Activity.Contract.GetLinkedData(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetLinkedData is a free data retrieval call binding the contract method 0x10b132d5.
//
// Solidity: function getLinkedData(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_Activity *ActivityCallerSession) GetLinkedData(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _Activity.Contract.GetLinkedData(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetNumActivities is a free data retrieval call binding the contract method 0xdb200f3f.
//
// Solidity: function getNumActivities(bytes32 _activitiesRef) view returns(uint256)
func (_Activity *ActivityCaller) GetNumActivities(opts *bind.CallOpts, _activitiesRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Activity.contract.Call(opts, out, "getNumActivities", _activitiesRef)
	return *ret0, err
}

// GetNumActivities is a free data retrieval call binding the contract method 0xdb200f3f.
//
// Solidity: function getNumActivities(bytes32 _activitiesRef) view returns(uint256)
func (_Activity *ActivitySession) GetNumActivities(_activitiesRef [32]byte) (*big.Int, error) {
	return _Activity.Contract.GetNumActivities(&_Activity.CallOpts, _activitiesRef)
}

// GetNumActivities is a free data retrieval call binding the contract method 0xdb200f3f.
//
// Solidity: function getNumActivities(bytes32 _activitiesRef) view returns(uint256)
func (_Activity *ActivityCallerSession) GetNumActivities(_activitiesRef [32]byte) (*big.Int, error) {
	return _Activity.Contract.GetNumActivities(&_Activity.CallOpts, _activitiesRef)
}

// GetReference is a free data retrieval call binding the contract method 0x9db9b0f6.
//
// Solidity: function getReference(bytes32 _activitiesRef, uint256 _index) view returns(bytes32)
func (_Activity *ActivityCaller) GetReference(opts *bind.CallOpts, _activitiesRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Activity.contract.Call(opts, out, "getReference", _activitiesRef, _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0x9db9b0f6.
//
// Solidity: function getReference(bytes32 _activitiesRef, uint256 _index) view returns(bytes32)
func (_Activity *ActivitySession) GetReference(_activitiesRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _Activity.Contract.GetReference(&_Activity.CallOpts, _activitiesRef, _index)
}

// GetReference is a free data retrieval call binding the contract method 0x9db9b0f6.
//
// Solidity: function getReference(bytes32 _activitiesRef, uint256 _index) view returns(bytes32)
func (_Activity *ActivityCallerSession) GetReference(_activitiesRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _Activity.Contract.GetReference(&_Activity.CallOpts, _activitiesRef, _index)
}

// GetReportingOrg is a free data retrieval call binding the contract method 0x40c21069.
//
// Solidity: function getReportingOrg(bytes32 _activitiesRef, bytes32 _activityRef) view returns(ActivityReportingOrg)
func (_Activity *ActivityCaller) GetReportingOrg(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (ActivityReportingOrg, error) {
	var (
		ret0 = new(ActivityReportingOrg)
	)
	out := ret0
	err := _Activity.contract.Call(opts, out, "getReportingOrg", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetReportingOrg is a free data retrieval call binding the contract method 0x40c21069.
//
// Solidity: function getReportingOrg(bytes32 _activitiesRef, bytes32 _activityRef) view returns(ActivityReportingOrg)
func (_Activity *ActivitySession) GetReportingOrg(_activitiesRef [32]byte, _activityRef [32]byte) (ActivityReportingOrg, error) {
	return _Activity.Contract.GetReportingOrg(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetReportingOrg is a free data retrieval call binding the contract method 0x40c21069.
//
// Solidity: function getReportingOrg(bytes32 _activitiesRef, bytes32 _activityRef) view returns(ActivityReportingOrg)
func (_Activity *ActivityCallerSession) GetReportingOrg(_activitiesRef [32]byte, _activityRef [32]byte) (ActivityReportingOrg, error) {
	return _Activity.Contract.GetReportingOrg(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetStatus is a free data retrieval call binding the contract method 0x8e56ee2c.
//
// Solidity: function getStatus(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint8)
func (_Activity *ActivityCaller) GetStatus(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Activity.contract.Call(opts, out, "getStatus", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x8e56ee2c.
//
// Solidity: function getStatus(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint8)
func (_Activity *ActivitySession) GetStatus(_activitiesRef [32]byte, _activityRef [32]byte) (uint8, error) {
	return _Activity.Contract.GetStatus(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetStatus is a free data retrieval call binding the contract method 0x8e56ee2c.
//
// Solidity: function getStatus(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint8)
func (_Activity *ActivityCallerSession) GetStatus(_activitiesRef [32]byte, _activityRef [32]byte) (uint8, error) {
	return _Activity.Contract.GetStatus(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetTitle is a free data retrieval call binding the contract method 0x8adc94fb.
//
// Solidity: function getTitle(bytes32 _activitiesRef, bytes32 _activityRef) view returns(string)
func (_Activity *ActivityCaller) GetTitle(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Activity.contract.Call(opts, out, "getTitle", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetTitle is a free data retrieval call binding the contract method 0x8adc94fb.
//
// Solidity: function getTitle(bytes32 _activitiesRef, bytes32 _activityRef) view returns(string)
func (_Activity *ActivitySession) GetTitle(_activitiesRef [32]byte, _activityRef [32]byte) (string, error) {
	return _Activity.Contract.GetTitle(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// GetTitle is a free data retrieval call binding the contract method 0x8adc94fb.
//
// Solidity: function getTitle(bytes32 _activitiesRef, bytes32 _activityRef) view returns(string)
func (_Activity *ActivityCallerSession) GetTitle(_activitiesRef [32]byte, _activityRef [32]byte) (string, error) {
	return _Activity.Contract.GetTitle(&_Activity.CallOpts, _activitiesRef, _activityRef)
}

// SetActivity is a paid mutator transaction binding the contract method 0x95b7378b.
//
// Solidity: function setActivity(bytes32 _activitiesRef, bytes32 activityRef, ActivityOrgActivity _activity) returns()
func (_Activity *ActivityTransactor) SetActivity(opts *bind.TransactOpts, _activitiesRef [32]byte, activityRef [32]byte, _activity ActivityOrgActivity) (*types.Transaction, error) {
	return _Activity.contract.Transact(opts, "setActivity", _activitiesRef, activityRef, _activity)
}

// SetActivity is a paid mutator transaction binding the contract method 0x95b7378b.
//
// Solidity: function setActivity(bytes32 _activitiesRef, bytes32 activityRef, ActivityOrgActivity _activity) returns()
func (_Activity *ActivitySession) SetActivity(_activitiesRef [32]byte, activityRef [32]byte, _activity ActivityOrgActivity) (*types.Transaction, error) {
	return _Activity.Contract.SetActivity(&_Activity.TransactOpts, _activitiesRef, activityRef, _activity)
}

// SetActivity is a paid mutator transaction binding the contract method 0x95b7378b.
//
// Solidity: function setActivity(bytes32 _activitiesRef, bytes32 activityRef, ActivityOrgActivity _activity) returns()
func (_Activity *ActivityTransactorSession) SetActivity(_activitiesRef [32]byte, activityRef [32]byte, _activity ActivityOrgActivity) (*types.Transaction, error) {
	return _Activity.Contract.SetActivity(&_Activity.TransactOpts, _activitiesRef, activityRef, _activity)
}

// IATIActivityABI is the input ABI used to generate the binding from.
const IATIActivityABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"activityRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"humanitarian\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"hierarchy\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"budgetNotProvided\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSecondary\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivity.ReportingOrg\",\"name\":\"reportingOrg\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"lastUpdated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structActivity.OrgActivity\",\"name\":\"_activity\",\"type\":\"tuple\"}],\"name\":\"SetActivity\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getActivity\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"humanitarian\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"hierarchy\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"budgetNotProvided\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSecondary\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivity.ReportingOrg\",\"name\":\"reportingOrg\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"lastUpdated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivity.OrgActivity\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getBudgetNotProvided\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getCurrency\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getDescription\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getHierarchy\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getHumanitarian\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getIdentifier\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getLang\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getLastUpdatedTime\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getLinkedData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"}],\"name\":\"getNumActivities\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getReportingOrg\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSecondary\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivity.ReportingOrg\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getTitle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"activityRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"humanitarian\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"hierarchy\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"budgetNotProvided\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSecondary\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivity.ReportingOrg\",\"name\":\"reportingOrg\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"lastUpdated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivity.OrgActivity\",\"name\":\"_activity\",\"type\":\"tuple\"}],\"name\":\"setActivity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IATIActivityFuncSigs maps the 4-byte function signature to its string representation.
var IATIActivityFuncSigs = map[string]string{
	"5aaf1a2d": "getActivity(bytes32,bytes32)",
	"a2f640a5": "getBudgetNotProvided(bytes32,bytes32)",
	"eef45eae": "getCurrency(bytes32,bytes32)",
	"87f9ba4a": "getDescription(bytes32,bytes32)",
	"b863fc70": "getHierarchy(bytes32,bytes32)",
	"1a95155e": "getHumanitarian(bytes32,bytes32)",
	"2b1cb7d5": "getIdentifier(bytes32,bytes32)",
	"aad88871": "getLang(bytes32,bytes32)",
	"7d78621b": "getLastUpdatedTime(bytes32,bytes32)",
	"10b132d5": "getLinkedData(bytes32,bytes32)",
	"db200f3f": "getNumActivities(bytes32)",
	"9db9b0f6": "getReference(bytes32,uint256)",
	"40c21069": "getReportingOrg(bytes32,bytes32)",
	"8e56ee2c": "getStatus(bytes32,bytes32)",
	"8adc94fb": "getTitle(bytes32,bytes32)",
	"95b7378b": "setActivity(bytes32,bytes32,(bool,uint8,uint8,uint8,(uint8,bool,bytes32),bytes32,bytes32,bytes32,bytes32,bytes32,string,string))",
}

// IATIActivityBin is the compiled bytecode used for deploying new contracts.
var IATIActivityBin = "0x608060405234801561001057600080fd5b5061145b806100206000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80638e56ee2c11610097578063aad8887111610066578063aad8887114610242578063b863fc7014610255578063db200f3f14610268578063eef45eae1461027b57610100565b80638e56ee2c146101e757806395b7378b146102075780639db9b0f61461021c578063a2f640a51461022f57610100565b80635aaf1a2d116100d35780635aaf1a2d146101815780637d78621b146101a157806387f9ba4a146101b45780638adc94fb146101d457610100565b806310b132d5146101055780631a95155e1461012e5780632b1cb7d51461014e57806340c2106914610161575b600080fd5b610118610113366004611059565b61028e565b604051610125919061132a565b60405180910390f35b61014161013c366004611059565b6102eb565b604051610125919061131f565b61011861015c366004611059565b610344565b61017461016f366004611059565b61039d565b60405161012591906113d1565b61019461018f366004611059565b610428565b60405161012591906113be565b6101186101af366004611059565b61064d565b6101c76101c2366004611059565b6106a6565b60405161012591906113ab565b6101c76101e2366004611059565b610789565b6101fa6101f5366004611059565b610834565b60405161012591906113df565b61021a61021536600461107a565b610893565b005b61011861022a3660046111b8565b610c25565b6101fa61023d366004611059565b610c88565b610118610250366004611059565b610ce8565b6101fa610263366004611059565b610d41565b610118610276366004611041565b610d9f565b610118610289366004611059565b610dcd565b600082811a60f81b6001600160f81b031916158015906102bd57508160001a60f81b6001600160f81b03191615155b6102c657600080fd5b5060008281526001602090815260408083208484529091529020600601545b92915050565b600082811a60f81b6001600160f81b0319161580159061031a57508160001a60f81b6001600160f81b03191615155b61032357600080fd5b50600091825260016020908152604080842092845291905290205460ff1690565b600082811a60f81b6001600160f81b0319161580159061037357508160001a60f81b6001600160f81b03191615155b61037c57600080fd5b50600091825260016020908152604080842092845291905290206007015490565b6103a5610e26565b8260001a60f81b6001600160f81b031916158015906103d357508160001a60f81b6001600160f81b03191615155b6103dc57600080fd5b5060009182526001602081815260408085209385529281529282902082516060810184529181015460ff8082168452610100909104161515938201939093526002909201549082015290565b610430610e46565b8260001a60f81b6001600160f81b0319161580159061045e57508160001a60f81b6001600160f81b03191615155b61046757600080fd5b6000838152600160208181526040808420868552825292839020835161018081018552815460ff808216151583526101008083048216848701526201000083048216848901526301000000909204811660608085019190915287519081018852848701548083168252839004909116151581860152600280850154828901526080840191909152600384015460a0840152600484015460c0840152600584015460e084015260068401548284015260078401546101208401526008840180548851978116159093026000190190921604601f8101859004850286018501909652858552909491936101408601939092908301828280156105a85780601f1061057d576101008083540402835291602001916105a8565b820191906000526020600020905b81548152906001019060200180831161058b57829003601f168201915b505050918352505060098201805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815293820193929183018282801561063c5780601f106106115761010080835404028352916020019161063c565b820191906000526020600020905b81548152906001019060200180831161061f57829003601f168201915b505050505081525050905092915050565b600082811a60f81b6001600160f81b0319161580159061067c57508160001a60f81b6001600160f81b03191615155b61068557600080fd5b50600091825260016020908152604080842092845291905290206003015490565b60608260001a60f81b6001600160f81b031916158015906106d657508160001a60f81b6001600160f81b03191615155b6106df57600080fd5b600083815260016020818152604080842086855282529283902060090180548451600294821615610100026000190190911693909304601f810183900483028401830190945283835291929083018282801561077c5780601f106107515761010080835404028352916020019161077c565b820191906000526020600020905b81548152906001019060200180831161075f57829003601f168201915b5050505050905092915050565b60608260001a60f81b6001600160f81b031916158015906107b957508160001a60f81b6001600160f81b03191615155b6107c257600080fd5b600083815260016020818152604080842086855282529283902060080180548451600294821615610100026000190190911693909304601f810183900483028401830190945283835291929083018282801561077c5780601f106107515761010080835404028352916020019161077c565b600082811a60f81b6001600160f81b0319161580159061086357508160001a60f81b6001600160f81b03191615155b61086c57600080fd5b50600091825260016020908152604080842092845291905290205462010000900460ff1690565b8260001a60f81b6001600160f81b031916158015906108c157508160001a60f81b6001600160f81b03191615155b80156108e1575061012081015160001a60f81b6001600160f81b03191615155b8015610904575060808101516040015160001a60f81b6001600160f81b03191615155b8015610917575060808101515160ff1615155b80156109295750600081610140015151115b8015610948575060a081015160001a60f81b6001600160f81b03191615155b8015610967575060c081015160001a60f81b6001600160f81b03191615155b8015610986575060e081015160001a60f81b6001600160f81b03191615155b80156109985750602081015160ff1615155b80156109ae5750600460ff16816020015160ff16105b80156109c05750604081015160ff1615155b80156109d65750600760ff16816040015160ff16105b80156109e85750606081015160ff1615155b80156109fe5750600460ff16816060015160ff16105b8015610a105750600081610160015151115b610a1957600080fd5b6000838152600160208181526040808420868552825292839020845181548387015186880151606089015160ff90811663010000000263ff00000019928216620100000262ff00001994831661010090810261ff001998151560ff1998891617891617959095161792909216919091178555608089015180519786018054828901511515850299909316929094169190911790931695909517905590930151600284015560a0840151600384015560c0840151600484015560e0840151600584015590830151600683015561012083015160078301556101408301518051849392610b0b926008850192910190610eb0565b506101608201518051610b28916009840191602090910190610eb0565b50505060008381526020819052604090819020905163745fca7b60e01b815273__$b620240993a55615b607189176fb82e62f$__9163745fca7b91610b71918691600401611333565b60206040518083038186803b158015610b8957600080fd5b505af4158015610b9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bc1919061101e565b610be557600083815260208181526040822080546001810182559083529120018290555b7f2947d521f8b2a845fbf77122191ac5933cbe8c8cf9dd9998c3d810c170d0bdda838383604051610c1893929190611383565b60405180910390a1505050565b600082811a60f81b6001600160f81b03191615801590610c52575060008381526020819052604090205482105b610c5b57600080fd5b6000838152602081905260409020805483908110610c7557fe5b9060005260206000200154905092915050565b600082811a60f81b6001600160f81b03191615801590610cb757508160001a60f81b6001600160f81b03191615155b610cc057600080fd5b5060009182526001602090815260408084209284529190529020546301000000900460ff1690565b600082811a60f81b6001600160f81b03191615801590610d1757508160001a60f81b6001600160f81b03191615155b610d2057600080fd5b50600091825260016020908152604080842092845291905290206004015490565b600082811a60f81b6001600160f81b03191615801590610d7057508160001a60f81b6001600160f81b03191615155b610d7957600080fd5b506000918252600160209081526040808420928452919052902054610100900460ff1690565b600081811a60f81b6001600160f81b031916610dba57600080fd5b5060009081526020819052604090205490565b600082811a60f81b6001600160f81b03191615801590610dfc57508160001a60f81b6001600160f81b03191615155b610e0557600080fd5b50600091825260016020908152604080842092845291905290206005015490565b604080516060810182526000808252602082018190529181019190915290565b6040805161018081018252600080825260208201819052918101829052606081019190915260808101610e77610e26565b81526000602082018190526040820181905260608083018290526080830182905260a083019190915260c0820181905260e09091015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610ef157805160ff1916838001178555610f1e565b82800160010185558215610f1e579182015b82811115610f1e578251825591602001919060010190610f03565b50610f2a929150610f2e565b5090565b610f4891905b80821115610f2a5760008155600101610f34565b90565b80356102e581611414565b600082601f830112610f66578081fd5b813567ffffffffffffffff811115610f7c578182fd5b610f8f601f8201601f19166020016113ed565b9150808252836020828501011115610fa657600080fd5b8060208401602084013760009082016020015292915050565b600060608284031215610fd0578081fd5b610fda60606113ed565b9050610fe6838361100d565b81526020820135610ff681611414565b806020830152506040820135604082015292915050565b803560ff811681146102e557600080fd5b60006020828403121561102f578081fd5b815161103a81611414565b9392505050565b600060208284031215611052578081fd5b5035919050565b6000806040838503121561106b578081fd5b50508035926020909101359150565b60008060006060848603121561108e578081fd5b8335925060208401359150604084013567ffffffffffffffff808211156110b3578283fd5b8186016101c081890312156110c6578384fd5b61018092506110d4836113ed565b6110de8983610f4b565b81526110ed896020840161100d565b60208201526110ff896040840161100d565b6040820152611111896060840161100d565b60608201526111238960808401610fbf565b608082015260e082013560a08201526101008083013560c08301526101208084013560e084015261014080850135838501526101609250828501358285015286850135915085821115611174578788fd5b6111808c838701610f56565b90840152506101a0830135945083851115611199578586fd5b6111a58a868501610f56565b8183015250809450505050509250925092565b6000806040838503121561106b578182fd5b15159052565b60008151808452815b818110156111f5576020818501810151868301820152016111d9565b818111156112065782602083870101525b50601f01601f19169290920160200192915050565b60006101c061122b8484516111ca565b602083015161123d6020860182611318565b5060408301516112506040860182611318565b5060608301516112636060860182611318565b50608083015161127660808601826112fa565b5060a083015160e085015260c0830151610100818187015260e08501519150610120828188015281860151925061014091508282880152808601519250506101608281880152818601519250836101808801526112d5848801846111d0565b908601518782036101a0890152935090506112f081846111d0565b9695505050505050565b805160ff168252602080820151151590830152604090810151910152565b60ff169052565b901515815260200190565b90815260200190565b60006040820184835260206040818501528185548084526060860191508685528285209350845b818110156113765784548352600194850194928401920161135a565b5090979650505050505050565b6000848252836020830152606060408301526113a2606083018461121b565b95945050505050565b60006020825261103a60208301846111d0565b60006020825261103a602083018461121b565b606081016102e582846112fa565b60ff91909116815260200190565b60405181810167ffffffffffffffff8111828210171561140c57600080fd5b604052919050565b801515811461142257600080fd5b5056fea2646970667358221220dadb1e7b0deffb5fce3a57f0e8005568e583c110574f23d5936354d681c7fa3764736f6c63430006060033"

// DeployIATIActivity deploys a new Ethereum contract, binding an instance of IATIActivity to it.
func DeployIATIActivity(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IATIActivity, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	stringsAddr, _, _, _ := DeployStrings(auth, backend)
	IATIActivityBin = strings.Replace(IATIActivityBin, "__$b620240993a55615b607189176fb82e62f$__", stringsAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IATIActivityBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IATIActivity{IATIActivityCaller: IATIActivityCaller{contract: contract}, IATIActivityTransactor: IATIActivityTransactor{contract: contract}, IATIActivityFilterer: IATIActivityFilterer{contract: contract}}, nil
}

// IATIActivity is an auto generated Go binding around an Ethereum contract.
type IATIActivity struct {
	IATIActivityCaller     // Read-only binding to the contract
	IATIActivityTransactor // Write-only binding to the contract
	IATIActivityFilterer   // Log filterer for contract events
}

// IATIActivityCaller is an auto generated read-only Go binding around an Ethereum contract.
type IATIActivityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IATIActivityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IATIActivityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IATIActivitySession struct {
	Contract     *IATIActivity     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IATIActivityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IATIActivityCallerSession struct {
	Contract *IATIActivityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IATIActivityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IATIActivityTransactorSession struct {
	Contract     *IATIActivityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IATIActivityRaw is an auto generated low-level Go binding around an Ethereum contract.
type IATIActivityRaw struct {
	Contract *IATIActivity // Generic contract binding to access the raw methods on
}

// IATIActivityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IATIActivityCallerRaw struct {
	Contract *IATIActivityCaller // Generic read-only contract binding to access the raw methods on
}

// IATIActivityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IATIActivityTransactorRaw struct {
	Contract *IATIActivityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIATIActivity creates a new instance of IATIActivity, bound to a specific deployed contract.
func NewIATIActivity(address common.Address, backend bind.ContractBackend) (*IATIActivity, error) {
	contract, err := bindIATIActivity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IATIActivity{IATIActivityCaller: IATIActivityCaller{contract: contract}, IATIActivityTransactor: IATIActivityTransactor{contract: contract}, IATIActivityFilterer: IATIActivityFilterer{contract: contract}}, nil
}

// NewIATIActivityCaller creates a new read-only instance of IATIActivity, bound to a specific deployed contract.
func NewIATIActivityCaller(address common.Address, caller bind.ContractCaller) (*IATIActivityCaller, error) {
	contract, err := bindIATIActivity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IATIActivityCaller{contract: contract}, nil
}

// NewIATIActivityTransactor creates a new write-only instance of IATIActivity, bound to a specific deployed contract.
func NewIATIActivityTransactor(address common.Address, transactor bind.ContractTransactor) (*IATIActivityTransactor, error) {
	contract, err := bindIATIActivity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IATIActivityTransactor{contract: contract}, nil
}

// NewIATIActivityFilterer creates a new log filterer instance of IATIActivity, bound to a specific deployed contract.
func NewIATIActivityFilterer(address common.Address, filterer bind.ContractFilterer) (*IATIActivityFilterer, error) {
	contract, err := bindIATIActivity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IATIActivityFilterer{contract: contract}, nil
}

// bindIATIActivity binds a generic wrapper to an already deployed contract.
func bindIATIActivity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIActivity *IATIActivityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIActivity.Contract.IATIActivityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIActivity *IATIActivityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIActivity.Contract.IATIActivityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIActivity *IATIActivityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIActivity.Contract.IATIActivityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIActivity *IATIActivityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIActivity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIActivity *IATIActivityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIActivity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIActivity *IATIActivityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIActivity.Contract.contract.Transact(opts, method, params...)
}

// GetActivity is a free data retrieval call binding the contract method 0x5aaf1a2d.
//
// Solidity: function getActivity(bytes32 _activitiesRef, bytes32 _activityRef) view returns(ActivityOrgActivity)
func (_IATIActivity *IATIActivityCaller) GetActivity(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (ActivityOrgActivity, error) {
	var (
		ret0 = new(ActivityOrgActivity)
	)
	out := ret0
	err := _IATIActivity.contract.Call(opts, out, "getActivity", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetActivity is a free data retrieval call binding the contract method 0x5aaf1a2d.
//
// Solidity: function getActivity(bytes32 _activitiesRef, bytes32 _activityRef) view returns(ActivityOrgActivity)
func (_IATIActivity *IATIActivitySession) GetActivity(_activitiesRef [32]byte, _activityRef [32]byte) (ActivityOrgActivity, error) {
	return _IATIActivity.Contract.GetActivity(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetActivity is a free data retrieval call binding the contract method 0x5aaf1a2d.
//
// Solidity: function getActivity(bytes32 _activitiesRef, bytes32 _activityRef) view returns(ActivityOrgActivity)
func (_IATIActivity *IATIActivityCallerSession) GetActivity(_activitiesRef [32]byte, _activityRef [32]byte) (ActivityOrgActivity, error) {
	return _IATIActivity.Contract.GetActivity(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetBudgetNotProvided is a free data retrieval call binding the contract method 0xa2f640a5.
//
// Solidity: function getBudgetNotProvided(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint8)
func (_IATIActivity *IATIActivityCaller) GetBudgetNotProvided(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivity.contract.Call(opts, out, "getBudgetNotProvided", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetBudgetNotProvided is a free data retrieval call binding the contract method 0xa2f640a5.
//
// Solidity: function getBudgetNotProvided(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint8)
func (_IATIActivity *IATIActivitySession) GetBudgetNotProvided(_activitiesRef [32]byte, _activityRef [32]byte) (uint8, error) {
	return _IATIActivity.Contract.GetBudgetNotProvided(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetBudgetNotProvided is a free data retrieval call binding the contract method 0xa2f640a5.
//
// Solidity: function getBudgetNotProvided(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint8)
func (_IATIActivity *IATIActivityCallerSession) GetBudgetNotProvided(_activitiesRef [32]byte, _activityRef [32]byte) (uint8, error) {
	return _IATIActivity.Contract.GetBudgetNotProvided(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetCurrency is a free data retrieval call binding the contract method 0xeef45eae.
//
// Solidity: function getCurrency(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_IATIActivity *IATIActivityCaller) GetCurrency(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivity.contract.Call(opts, out, "getCurrency", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetCurrency is a free data retrieval call binding the contract method 0xeef45eae.
//
// Solidity: function getCurrency(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_IATIActivity *IATIActivitySession) GetCurrency(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _IATIActivity.Contract.GetCurrency(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetCurrency is a free data retrieval call binding the contract method 0xeef45eae.
//
// Solidity: function getCurrency(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_IATIActivity *IATIActivityCallerSession) GetCurrency(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _IATIActivity.Contract.GetCurrency(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetDescription is a free data retrieval call binding the contract method 0x87f9ba4a.
//
// Solidity: function getDescription(bytes32 _activitiesRef, bytes32 _activityRef) view returns(string)
func (_IATIActivity *IATIActivityCaller) GetDescription(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IATIActivity.contract.Call(opts, out, "getDescription", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetDescription is a free data retrieval call binding the contract method 0x87f9ba4a.
//
// Solidity: function getDescription(bytes32 _activitiesRef, bytes32 _activityRef) view returns(string)
func (_IATIActivity *IATIActivitySession) GetDescription(_activitiesRef [32]byte, _activityRef [32]byte) (string, error) {
	return _IATIActivity.Contract.GetDescription(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetDescription is a free data retrieval call binding the contract method 0x87f9ba4a.
//
// Solidity: function getDescription(bytes32 _activitiesRef, bytes32 _activityRef) view returns(string)
func (_IATIActivity *IATIActivityCallerSession) GetDescription(_activitiesRef [32]byte, _activityRef [32]byte) (string, error) {
	return _IATIActivity.Contract.GetDescription(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetHierarchy is a free data retrieval call binding the contract method 0xb863fc70.
//
// Solidity: function getHierarchy(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint8)
func (_IATIActivity *IATIActivityCaller) GetHierarchy(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivity.contract.Call(opts, out, "getHierarchy", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetHierarchy is a free data retrieval call binding the contract method 0xb863fc70.
//
// Solidity: function getHierarchy(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint8)
func (_IATIActivity *IATIActivitySession) GetHierarchy(_activitiesRef [32]byte, _activityRef [32]byte) (uint8, error) {
	return _IATIActivity.Contract.GetHierarchy(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetHierarchy is a free data retrieval call binding the contract method 0xb863fc70.
//
// Solidity: function getHierarchy(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint8)
func (_IATIActivity *IATIActivityCallerSession) GetHierarchy(_activitiesRef [32]byte, _activityRef [32]byte) (uint8, error) {
	return _IATIActivity.Contract.GetHierarchy(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetHumanitarian is a free data retrieval call binding the contract method 0x1a95155e.
//
// Solidity: function getHumanitarian(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bool)
func (_IATIActivity *IATIActivityCaller) GetHumanitarian(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IATIActivity.contract.Call(opts, out, "getHumanitarian", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetHumanitarian is a free data retrieval call binding the contract method 0x1a95155e.
//
// Solidity: function getHumanitarian(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bool)
func (_IATIActivity *IATIActivitySession) GetHumanitarian(_activitiesRef [32]byte, _activityRef [32]byte) (bool, error) {
	return _IATIActivity.Contract.GetHumanitarian(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetHumanitarian is a free data retrieval call binding the contract method 0x1a95155e.
//
// Solidity: function getHumanitarian(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bool)
func (_IATIActivity *IATIActivityCallerSession) GetHumanitarian(_activitiesRef [32]byte, _activityRef [32]byte) (bool, error) {
	return _IATIActivity.Contract.GetHumanitarian(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetIdentifier is a free data retrieval call binding the contract method 0x2b1cb7d5.
//
// Solidity: function getIdentifier(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_IATIActivity *IATIActivityCaller) GetIdentifier(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivity.contract.Call(opts, out, "getIdentifier", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetIdentifier is a free data retrieval call binding the contract method 0x2b1cb7d5.
//
// Solidity: function getIdentifier(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_IATIActivity *IATIActivitySession) GetIdentifier(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _IATIActivity.Contract.GetIdentifier(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetIdentifier is a free data retrieval call binding the contract method 0x2b1cb7d5.
//
// Solidity: function getIdentifier(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_IATIActivity *IATIActivityCallerSession) GetIdentifier(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _IATIActivity.Contract.GetIdentifier(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetLang is a free data retrieval call binding the contract method 0xaad88871.
//
// Solidity: function getLang(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_IATIActivity *IATIActivityCaller) GetLang(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivity.contract.Call(opts, out, "getLang", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetLang is a free data retrieval call binding the contract method 0xaad88871.
//
// Solidity: function getLang(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_IATIActivity *IATIActivitySession) GetLang(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _IATIActivity.Contract.GetLang(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetLang is a free data retrieval call binding the contract method 0xaad88871.
//
// Solidity: function getLang(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_IATIActivity *IATIActivityCallerSession) GetLang(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _IATIActivity.Contract.GetLang(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetLastUpdatedTime is a free data retrieval call binding the contract method 0x7d78621b.
//
// Solidity: function getLastUpdatedTime(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_IATIActivity *IATIActivityCaller) GetLastUpdatedTime(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivity.contract.Call(opts, out, "getLastUpdatedTime", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetLastUpdatedTime is a free data retrieval call binding the contract method 0x7d78621b.
//
// Solidity: function getLastUpdatedTime(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_IATIActivity *IATIActivitySession) GetLastUpdatedTime(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _IATIActivity.Contract.GetLastUpdatedTime(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetLastUpdatedTime is a free data retrieval call binding the contract method 0x7d78621b.
//
// Solidity: function getLastUpdatedTime(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_IATIActivity *IATIActivityCallerSession) GetLastUpdatedTime(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _IATIActivity.Contract.GetLastUpdatedTime(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetLinkedData is a free data retrieval call binding the contract method 0x10b132d5.
//
// Solidity: function getLinkedData(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_IATIActivity *IATIActivityCaller) GetLinkedData(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivity.contract.Call(opts, out, "getLinkedData", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetLinkedData is a free data retrieval call binding the contract method 0x10b132d5.
//
// Solidity: function getLinkedData(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_IATIActivity *IATIActivitySession) GetLinkedData(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _IATIActivity.Contract.GetLinkedData(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetLinkedData is a free data retrieval call binding the contract method 0x10b132d5.
//
// Solidity: function getLinkedData(bytes32 _activitiesRef, bytes32 _activityRef) view returns(bytes32)
func (_IATIActivity *IATIActivityCallerSession) GetLinkedData(_activitiesRef [32]byte, _activityRef [32]byte) ([32]byte, error) {
	return _IATIActivity.Contract.GetLinkedData(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetNumActivities is a free data retrieval call binding the contract method 0xdb200f3f.
//
// Solidity: function getNumActivities(bytes32 _activitiesRef) view returns(uint256)
func (_IATIActivity *IATIActivityCaller) GetNumActivities(opts *bind.CallOpts, _activitiesRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIActivity.contract.Call(opts, out, "getNumActivities", _activitiesRef)
	return *ret0, err
}

// GetNumActivities is a free data retrieval call binding the contract method 0xdb200f3f.
//
// Solidity: function getNumActivities(bytes32 _activitiesRef) view returns(uint256)
func (_IATIActivity *IATIActivitySession) GetNumActivities(_activitiesRef [32]byte) (*big.Int, error) {
	return _IATIActivity.Contract.GetNumActivities(&_IATIActivity.CallOpts, _activitiesRef)
}

// GetNumActivities is a free data retrieval call binding the contract method 0xdb200f3f.
//
// Solidity: function getNumActivities(bytes32 _activitiesRef) view returns(uint256)
func (_IATIActivity *IATIActivityCallerSession) GetNumActivities(_activitiesRef [32]byte) (*big.Int, error) {
	return _IATIActivity.Contract.GetNumActivities(&_IATIActivity.CallOpts, _activitiesRef)
}

// GetReference is a free data retrieval call binding the contract method 0x9db9b0f6.
//
// Solidity: function getReference(bytes32 _activitiesRef, uint256 _index) view returns(bytes32)
func (_IATIActivity *IATIActivityCaller) GetReference(opts *bind.CallOpts, _activitiesRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivity.contract.Call(opts, out, "getReference", _activitiesRef, _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0x9db9b0f6.
//
// Solidity: function getReference(bytes32 _activitiesRef, uint256 _index) view returns(bytes32)
func (_IATIActivity *IATIActivitySession) GetReference(_activitiesRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIActivity.Contract.GetReference(&_IATIActivity.CallOpts, _activitiesRef, _index)
}

// GetReference is a free data retrieval call binding the contract method 0x9db9b0f6.
//
// Solidity: function getReference(bytes32 _activitiesRef, uint256 _index) view returns(bytes32)
func (_IATIActivity *IATIActivityCallerSession) GetReference(_activitiesRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIActivity.Contract.GetReference(&_IATIActivity.CallOpts, _activitiesRef, _index)
}

// GetReportingOrg is a free data retrieval call binding the contract method 0x40c21069.
//
// Solidity: function getReportingOrg(bytes32 _activitiesRef, bytes32 _activityRef) view returns(ActivityReportingOrg)
func (_IATIActivity *IATIActivityCaller) GetReportingOrg(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (ActivityReportingOrg, error) {
	var (
		ret0 = new(ActivityReportingOrg)
	)
	out := ret0
	err := _IATIActivity.contract.Call(opts, out, "getReportingOrg", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetReportingOrg is a free data retrieval call binding the contract method 0x40c21069.
//
// Solidity: function getReportingOrg(bytes32 _activitiesRef, bytes32 _activityRef) view returns(ActivityReportingOrg)
func (_IATIActivity *IATIActivitySession) GetReportingOrg(_activitiesRef [32]byte, _activityRef [32]byte) (ActivityReportingOrg, error) {
	return _IATIActivity.Contract.GetReportingOrg(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetReportingOrg is a free data retrieval call binding the contract method 0x40c21069.
//
// Solidity: function getReportingOrg(bytes32 _activitiesRef, bytes32 _activityRef) view returns(ActivityReportingOrg)
func (_IATIActivity *IATIActivityCallerSession) GetReportingOrg(_activitiesRef [32]byte, _activityRef [32]byte) (ActivityReportingOrg, error) {
	return _IATIActivity.Contract.GetReportingOrg(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetStatus is a free data retrieval call binding the contract method 0x8e56ee2c.
//
// Solidity: function getStatus(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint8)
func (_IATIActivity *IATIActivityCaller) GetStatus(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivity.contract.Call(opts, out, "getStatus", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetStatus is a free data retrieval call binding the contract method 0x8e56ee2c.
//
// Solidity: function getStatus(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint8)
func (_IATIActivity *IATIActivitySession) GetStatus(_activitiesRef [32]byte, _activityRef [32]byte) (uint8, error) {
	return _IATIActivity.Contract.GetStatus(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetStatus is a free data retrieval call binding the contract method 0x8e56ee2c.
//
// Solidity: function getStatus(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint8)
func (_IATIActivity *IATIActivityCallerSession) GetStatus(_activitiesRef [32]byte, _activityRef [32]byte) (uint8, error) {
	return _IATIActivity.Contract.GetStatus(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetTitle is a free data retrieval call binding the contract method 0x8adc94fb.
//
// Solidity: function getTitle(bytes32 _activitiesRef, bytes32 _activityRef) view returns(string)
func (_IATIActivity *IATIActivityCaller) GetTitle(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IATIActivity.contract.Call(opts, out, "getTitle", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetTitle is a free data retrieval call binding the contract method 0x8adc94fb.
//
// Solidity: function getTitle(bytes32 _activitiesRef, bytes32 _activityRef) view returns(string)
func (_IATIActivity *IATIActivitySession) GetTitle(_activitiesRef [32]byte, _activityRef [32]byte) (string, error) {
	return _IATIActivity.Contract.GetTitle(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// GetTitle is a free data retrieval call binding the contract method 0x8adc94fb.
//
// Solidity: function getTitle(bytes32 _activitiesRef, bytes32 _activityRef) view returns(string)
func (_IATIActivity *IATIActivityCallerSession) GetTitle(_activitiesRef [32]byte, _activityRef [32]byte) (string, error) {
	return _IATIActivity.Contract.GetTitle(&_IATIActivity.CallOpts, _activitiesRef, _activityRef)
}

// SetActivity is a paid mutator transaction binding the contract method 0x95b7378b.
//
// Solidity: function setActivity(bytes32 _activitiesRef, bytes32 activityRef, ActivityOrgActivity _activity) returns()
func (_IATIActivity *IATIActivityTransactor) SetActivity(opts *bind.TransactOpts, _activitiesRef [32]byte, activityRef [32]byte, _activity ActivityOrgActivity) (*types.Transaction, error) {
	return _IATIActivity.contract.Transact(opts, "setActivity", _activitiesRef, activityRef, _activity)
}

// SetActivity is a paid mutator transaction binding the contract method 0x95b7378b.
//
// Solidity: function setActivity(bytes32 _activitiesRef, bytes32 activityRef, ActivityOrgActivity _activity) returns()
func (_IATIActivity *IATIActivitySession) SetActivity(_activitiesRef [32]byte, activityRef [32]byte, _activity ActivityOrgActivity) (*types.Transaction, error) {
	return _IATIActivity.Contract.SetActivity(&_IATIActivity.TransactOpts, _activitiesRef, activityRef, _activity)
}

// SetActivity is a paid mutator transaction binding the contract method 0x95b7378b.
//
// Solidity: function setActivity(bytes32 _activitiesRef, bytes32 activityRef, ActivityOrgActivity _activity) returns()
func (_IATIActivity *IATIActivityTransactorSession) SetActivity(_activitiesRef [32]byte, activityRef [32]byte, _activity ActivityOrgActivity) (*types.Transaction, error) {
	return _IATIActivity.Contract.SetActivity(&_IATIActivity.TransactOpts, _activitiesRef, activityRef, _activity)
}

// IATIActivitySetActivityIterator is returned from FilterSetActivity and is used to iterate over the raw logs and unpacked data for SetActivity events raised by the IATIActivity contract.
type IATIActivitySetActivityIterator struct {
	Event *IATIActivitySetActivity // Event containing the contract specifics and raw log

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
func (it *IATIActivitySetActivityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IATIActivitySetActivity)
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
		it.Event = new(IATIActivitySetActivity)
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
func (it *IATIActivitySetActivityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IATIActivitySetActivityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IATIActivitySetActivity represents a SetActivity event raised by the IATIActivity contract.
type IATIActivitySetActivity struct {
	ActivitiesRef [32]byte
	ActivityRef   [32]byte
	Activity      ActivityOrgActivity
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetActivity is a free log retrieval operation binding the contract event 0x2947d521f8b2a845fbf77122191ac5933cbe8c8cf9dd9998c3d810c170d0bdda.
//
// Solidity: event SetActivity(bytes32 _activitiesRef, bytes32 activityRef, ActivityOrgActivity _activity)
func (_IATIActivity *IATIActivityFilterer) FilterSetActivity(opts *bind.FilterOpts) (*IATIActivitySetActivityIterator, error) {

	logs, sub, err := _IATIActivity.contract.FilterLogs(opts, "SetActivity")
	if err != nil {
		return nil, err
	}
	return &IATIActivitySetActivityIterator{contract: _IATIActivity.contract, event: "SetActivity", logs: logs, sub: sub}, nil
}

// WatchSetActivity is a free log subscription operation binding the contract event 0x2947d521f8b2a845fbf77122191ac5933cbe8c8cf9dd9998c3d810c170d0bdda.
//
// Solidity: event SetActivity(bytes32 _activitiesRef, bytes32 activityRef, ActivityOrgActivity _activity)
func (_IATIActivity *IATIActivityFilterer) WatchSetActivity(opts *bind.WatchOpts, sink chan<- *IATIActivitySetActivity) (event.Subscription, error) {

	logs, sub, err := _IATIActivity.contract.WatchLogs(opts, "SetActivity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IATIActivitySetActivity)
				if err := _IATIActivity.contract.UnpackLog(event, "SetActivity", log); err != nil {
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

// ParseSetActivity is a log parse operation binding the contract event 0x2947d521f8b2a845fbf77122191ac5933cbe8c8cf9dd9998c3d810c170d0bdda.
//
// Solidity: event SetActivity(bytes32 _activitiesRef, bytes32 activityRef, ActivityOrgActivity _activity)
func (_IATIActivity *IATIActivityFilterer) ParseSetActivity(log types.Log) (*IATIActivitySetActivity, error) {
	event := new(IATIActivitySetActivity)
	if err := _IATIActivity.contract.UnpackLog(event, "SetActivity", log); err != nil {
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
