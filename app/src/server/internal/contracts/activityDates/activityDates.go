// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package activityDates

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

// ActivityDatesDate is an auto generated low-level Go binding around an user-defined struct.
type ActivityDatesDate struct {
	DateType  uint8
	Date      [32]byte
	Narrative string
}

// ActivityDatesABI is the input ABI used to generate the binding from.
const ActivityDatesABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_dateRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"dateType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"narrative\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structActivityDates.Date\",\"name\":\"_date\",\"type\":\"tuple\"}],\"name\":\"SetDate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dateRef\",\"type\":\"bytes32\"}],\"name\":\"getDate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"dateType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"narrative\",\"type\":\"string\"}],\"internalType\":\"structActivityDates.Date\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dateRef\",\"type\":\"bytes32\"}],\"name\":\"getISODate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dateRef\",\"type\":\"bytes32\"}],\"name\":\"getNarrative\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getNumDates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dateRef\",\"type\":\"bytes32\"}],\"name\":\"getType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dateRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"dateType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"narrative\",\"type\":\"string\"}],\"internalType\":\"structActivityDates.Date\",\"name\":\"_date\",\"type\":\"tuple\"}],\"name\":\"setDate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ActivityDatesFuncSigs maps the 4-byte function signature to its string representation.
var ActivityDatesFuncSigs = map[string]string{
	"b5e2ffb4": "getDate(bytes32,bytes32,bytes32)",
	"5c77487f": "getISODate(bytes32,bytes32,bytes32)",
	"ae13bb7a": "getNarrative(bytes32,bytes32,bytes32)",
	"ffd55577": "getNumDates(bytes32,bytes32)",
	"eef67d78": "getReference(bytes32,bytes32,uint256)",
	"b153af81": "getType(bytes32,bytes32,bytes32)",
	"934fe599": "setDate(bytes32,bytes32,bytes32,(uint8,bytes32,string))",
}

// ActivityDates is an auto generated Go binding around an Ethereum contract.
type ActivityDates struct {
	ActivityDatesCaller     // Read-only binding to the contract
	ActivityDatesTransactor // Write-only binding to the contract
	ActivityDatesFilterer   // Log filterer for contract events
}

// ActivityDatesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivityDatesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityDatesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivityDatesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityDatesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivityDatesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityDatesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivityDatesSession struct {
	Contract     *ActivityDates    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ActivityDatesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivityDatesCallerSession struct {
	Contract *ActivityDatesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ActivityDatesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivityDatesTransactorSession struct {
	Contract     *ActivityDatesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ActivityDatesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivityDatesRaw struct {
	Contract *ActivityDates // Generic contract binding to access the raw methods on
}

// ActivityDatesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivityDatesCallerRaw struct {
	Contract *ActivityDatesCaller // Generic read-only contract binding to access the raw methods on
}

// ActivityDatesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivityDatesTransactorRaw struct {
	Contract *ActivityDatesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivityDates creates a new instance of ActivityDates, bound to a specific deployed contract.
func NewActivityDates(address common.Address, backend bind.ContractBackend) (*ActivityDates, error) {
	contract, err := bindActivityDates(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ActivityDates{ActivityDatesCaller: ActivityDatesCaller{contract: contract}, ActivityDatesTransactor: ActivityDatesTransactor{contract: contract}, ActivityDatesFilterer: ActivityDatesFilterer{contract: contract}}, nil
}

// NewActivityDatesCaller creates a new read-only instance of ActivityDates, bound to a specific deployed contract.
func NewActivityDatesCaller(address common.Address, caller bind.ContractCaller) (*ActivityDatesCaller, error) {
	contract, err := bindActivityDates(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityDatesCaller{contract: contract}, nil
}

// NewActivityDatesTransactor creates a new write-only instance of ActivityDates, bound to a specific deployed contract.
func NewActivityDatesTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivityDatesTransactor, error) {
	contract, err := bindActivityDates(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityDatesTransactor{contract: contract}, nil
}

// NewActivityDatesFilterer creates a new log filterer instance of ActivityDates, bound to a specific deployed contract.
func NewActivityDatesFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivityDatesFilterer, error) {
	contract, err := bindActivityDates(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivityDatesFilterer{contract: contract}, nil
}

// bindActivityDates binds a generic wrapper to an already deployed contract.
func bindActivityDates(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivityDatesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityDates *ActivityDatesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivityDates.Contract.ActivityDatesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityDates *ActivityDatesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityDates.Contract.ActivityDatesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityDates *ActivityDatesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityDates.Contract.ActivityDatesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityDates *ActivityDatesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivityDates.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityDates *ActivityDatesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityDates.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityDates *ActivityDatesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityDates.Contract.contract.Transact(opts, method, params...)
}

// GetDate is a free data retrieval call binding the contract method 0xb5e2ffb4.
//
// Solidity: function getDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(ActivityDatesDate)
func (_ActivityDates *ActivityDatesCaller) GetDate(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) (ActivityDatesDate, error) {
	var (
		ret0 = new(ActivityDatesDate)
	)
	out := ret0
	err := _ActivityDates.contract.Call(opts, out, "getDate", _activitiesRef, _activityRef, _dateRef)
	return *ret0, err
}

// GetDate is a free data retrieval call binding the contract method 0xb5e2ffb4.
//
// Solidity: function getDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(ActivityDatesDate)
func (_ActivityDates *ActivityDatesSession) GetDate(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) (ActivityDatesDate, error) {
	return _ActivityDates.Contract.GetDate(&_ActivityDates.CallOpts, _activitiesRef, _activityRef, _dateRef)
}

// GetDate is a free data retrieval call binding the contract method 0xb5e2ffb4.
//
// Solidity: function getDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(ActivityDatesDate)
func (_ActivityDates *ActivityDatesCallerSession) GetDate(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) (ActivityDatesDate, error) {
	return _ActivityDates.Contract.GetDate(&_ActivityDates.CallOpts, _activitiesRef, _activityRef, _dateRef)
}

// GetISODate is a free data retrieval call binding the contract method 0x5c77487f.
//
// Solidity: function getISODate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(bytes32)
func (_ActivityDates *ActivityDatesCaller) GetISODate(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityDates.contract.Call(opts, out, "getISODate", _activitiesRef, _activityRef, _dateRef)
	return *ret0, err
}

// GetISODate is a free data retrieval call binding the contract method 0x5c77487f.
//
// Solidity: function getISODate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(bytes32)
func (_ActivityDates *ActivityDatesSession) GetISODate(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) ([32]byte, error) {
	return _ActivityDates.Contract.GetISODate(&_ActivityDates.CallOpts, _activitiesRef, _activityRef, _dateRef)
}

// GetISODate is a free data retrieval call binding the contract method 0x5c77487f.
//
// Solidity: function getISODate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(bytes32)
func (_ActivityDates *ActivityDatesCallerSession) GetISODate(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) ([32]byte, error) {
	return _ActivityDates.Contract.GetISODate(&_ActivityDates.CallOpts, _activitiesRef, _activityRef, _dateRef)
}

// GetNarrative is a free data retrieval call binding the contract method 0xae13bb7a.
//
// Solidity: function getNarrative(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(string)
func (_ActivityDates *ActivityDatesCaller) GetNarrative(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ActivityDates.contract.Call(opts, out, "getNarrative", _activitiesRef, _activityRef, _dateRef)
	return *ret0, err
}

// GetNarrative is a free data retrieval call binding the contract method 0xae13bb7a.
//
// Solidity: function getNarrative(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(string)
func (_ActivityDates *ActivityDatesSession) GetNarrative(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) (string, error) {
	return _ActivityDates.Contract.GetNarrative(&_ActivityDates.CallOpts, _activitiesRef, _activityRef, _dateRef)
}

// GetNarrative is a free data retrieval call binding the contract method 0xae13bb7a.
//
// Solidity: function getNarrative(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(string)
func (_ActivityDates *ActivityDatesCallerSession) GetNarrative(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) (string, error) {
	return _ActivityDates.Contract.GetNarrative(&_ActivityDates.CallOpts, _activitiesRef, _activityRef, _dateRef)
}

// GetNumDates is a free data retrieval call binding the contract method 0xffd55577.
//
// Solidity: function getNumDates(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivityDates *ActivityDatesCaller) GetNumDates(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ActivityDates.contract.Call(opts, out, "getNumDates", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetNumDates is a free data retrieval call binding the contract method 0xffd55577.
//
// Solidity: function getNumDates(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivityDates *ActivityDatesSession) GetNumDates(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _ActivityDates.Contract.GetNumDates(&_ActivityDates.CallOpts, _activitiesRef, _activityRef)
}

// GetNumDates is a free data retrieval call binding the contract method 0xffd55577.
//
// Solidity: function getNumDates(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivityDates *ActivityDatesCallerSession) GetNumDates(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _ActivityDates.Contract.GetNumDates(&_ActivityDates.CallOpts, _activitiesRef, _activityRef)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivityDates *ActivityDatesCaller) GetReference(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityDates.contract.Call(opts, out, "getReference", _activitiesRef, _activityRef, _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivityDates *ActivityDatesSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _ActivityDates.Contract.GetReference(&_ActivityDates.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivityDates *ActivityDatesCallerSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _ActivityDates.Contract.GetReference(&_ActivityDates.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetType is a free data retrieval call binding the contract method 0xb153af81.
//
// Solidity: function getType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(uint8)
func (_ActivityDates *ActivityDatesCaller) GetType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ActivityDates.contract.Call(opts, out, "getType", _activitiesRef, _activityRef, _dateRef)
	return *ret0, err
}

// GetType is a free data retrieval call binding the contract method 0xb153af81.
//
// Solidity: function getType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(uint8)
func (_ActivityDates *ActivityDatesSession) GetType(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) (uint8, error) {
	return _ActivityDates.Contract.GetType(&_ActivityDates.CallOpts, _activitiesRef, _activityRef, _dateRef)
}

// GetType is a free data retrieval call binding the contract method 0xb153af81.
//
// Solidity: function getType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(uint8)
func (_ActivityDates *ActivityDatesCallerSession) GetType(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) (uint8, error) {
	return _ActivityDates.Contract.GetType(&_ActivityDates.CallOpts, _activitiesRef, _activityRef, _dateRef)
}

// SetDate is a paid mutator transaction binding the contract method 0x934fe599.
//
// Solidity: function setDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef, ActivityDatesDate _date) returns()
func (_ActivityDates *ActivityDatesTransactor) SetDate(opts *bind.TransactOpts, _activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte, _date ActivityDatesDate) (*types.Transaction, error) {
	return _ActivityDates.contract.Transact(opts, "setDate", _activitiesRef, _activityRef, _dateRef, _date)
}

// SetDate is a paid mutator transaction binding the contract method 0x934fe599.
//
// Solidity: function setDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef, ActivityDatesDate _date) returns()
func (_ActivityDates *ActivityDatesSession) SetDate(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte, _date ActivityDatesDate) (*types.Transaction, error) {
	return _ActivityDates.Contract.SetDate(&_ActivityDates.TransactOpts, _activitiesRef, _activityRef, _dateRef, _date)
}

// SetDate is a paid mutator transaction binding the contract method 0x934fe599.
//
// Solidity: function setDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef, ActivityDatesDate _date) returns()
func (_ActivityDates *ActivityDatesTransactorSession) SetDate(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte, _date ActivityDatesDate) (*types.Transaction, error) {
	return _ActivityDates.Contract.SetDate(&_ActivityDates.TransactOpts, _activitiesRef, _activityRef, _dateRef, _date)
}

// ActivityDatesSetDateIterator is returned from FilterSetDate and is used to iterate over the raw logs and unpacked data for SetDate events raised by the ActivityDates contract.
type ActivityDatesSetDateIterator struct {
	Event *ActivityDatesSetDate // Event containing the contract specifics and raw log

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
func (it *ActivityDatesSetDateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivityDatesSetDate)
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
		it.Event = new(ActivityDatesSetDate)
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
func (it *ActivityDatesSetDateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivityDatesSetDateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivityDatesSetDate represents a SetDate event raised by the ActivityDates contract.
type ActivityDatesSetDate struct {
	ActivitiesRef [32]byte
	ActivityRef   [32]byte
	DateRef       [32]byte
	Date          ActivityDatesDate
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetDate is a free log retrieval operation binding the contract event 0x1f1087f8c393b2db10215ffd640a5ba2cda450169f714629afd0d02330e3d102.
//
// Solidity: event SetDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef, ActivityDatesDate _date)
func (_ActivityDates *ActivityDatesFilterer) FilterSetDate(opts *bind.FilterOpts) (*ActivityDatesSetDateIterator, error) {

	logs, sub, err := _ActivityDates.contract.FilterLogs(opts, "SetDate")
	if err != nil {
		return nil, err
	}
	return &ActivityDatesSetDateIterator{contract: _ActivityDates.contract, event: "SetDate", logs: logs, sub: sub}, nil
}

// WatchSetDate is a free log subscription operation binding the contract event 0x1f1087f8c393b2db10215ffd640a5ba2cda450169f714629afd0d02330e3d102.
//
// Solidity: event SetDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef, ActivityDatesDate _date)
func (_ActivityDates *ActivityDatesFilterer) WatchSetDate(opts *bind.WatchOpts, sink chan<- *ActivityDatesSetDate) (event.Subscription, error) {

	logs, sub, err := _ActivityDates.contract.WatchLogs(opts, "SetDate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivityDatesSetDate)
				if err := _ActivityDates.contract.UnpackLog(event, "SetDate", log); err != nil {
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

// ParseSetDate is a log parse operation binding the contract event 0x1f1087f8c393b2db10215ffd640a5ba2cda450169f714629afd0d02330e3d102.
//
// Solidity: event SetDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef, ActivityDatesDate _date)
func (_ActivityDates *ActivityDatesFilterer) ParseSetDate(log types.Log) (*ActivityDatesSetDate, error) {
	event := new(ActivityDatesSetDate)
	if err := _ActivityDates.contract.UnpackLog(event, "SetDate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IATIActivityDatesABI is the input ABI used to generate the binding from.
const IATIActivityDatesABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_dateRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"dateType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"narrative\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structActivityDates.Date\",\"name\":\"_date\",\"type\":\"tuple\"}],\"name\":\"SetDate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dateRef\",\"type\":\"bytes32\"}],\"name\":\"getDate\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"dateType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"narrative\",\"type\":\"string\"}],\"internalType\":\"structActivityDates.Date\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dateRef\",\"type\":\"bytes32\"}],\"name\":\"getISODate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dateRef\",\"type\":\"bytes32\"}],\"name\":\"getNarrative\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getNumDates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dateRef\",\"type\":\"bytes32\"}],\"name\":\"getType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dateRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"dateType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"narrative\",\"type\":\"string\"}],\"internalType\":\"structActivityDates.Date\",\"name\":\"_date\",\"type\":\"tuple\"}],\"name\":\"setDate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IATIActivityDatesFuncSigs maps the 4-byte function signature to its string representation.
var IATIActivityDatesFuncSigs = map[string]string{
	"b5e2ffb4": "getDate(bytes32,bytes32,bytes32)",
	"5c77487f": "getISODate(bytes32,bytes32,bytes32)",
	"ae13bb7a": "getNarrative(bytes32,bytes32,bytes32)",
	"ffd55577": "getNumDates(bytes32,bytes32)",
	"eef67d78": "getReference(bytes32,bytes32,uint256)",
	"b153af81": "getType(bytes32,bytes32,bytes32)",
	"934fe599": "setDate(bytes32,bytes32,bytes32,(uint8,bytes32,string))",
}

// IATIActivityDatesBin is the compiled bytecode used for deploying new contracts.
var IATIActivityDatesBin = "0x608060405234801561001057600080fd5b50610b3b806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063b153af811161005b578063b153af81146100e0578063b5e2ffb414610100578063eef67d7814610120578063ffd55577146101335761007d565b80635c77487f14610082578063934fe599146100ab578063ae13bb7a146100c0575b600080fd5b610095610090366004610863565b610146565b6040516100a29190610a10565b60405180910390f35b6100be6100b936600461088e565b6101c2565b005b6100d36100ce366004610863565b6103cc565b6040516100a29190610a98565b6100f36100ee366004610863565b6104d1565b6040516100a29190610abe565b61011361010e366004610863565b61054c565b6040516100a29190610aab565b61009561012e36600461097d565b61067c565b610095610141366004610842565b61070d565b600083811a60f81b6001600160f81b0319161580159061017557508260001a60f81b6001600160f81b03191615155b801561019057508160001a60f81b6001600160f81b03191615155b61019957600080fd5b506000928352600160208181526040808620948652938152838520928552919091529120015490565b8360001a60f81b6001600160f81b031916158015906101f057508260001a60f81b6001600160f81b03191615155b801561020b57508160001a60f81b6001600160f81b03191615155b801561021a5750805160ff1615155b801561022c57508051600560ff909116105b801561024b5750602081015160001a60f81b6001600160f81b03191615155b801561025c57506000816040015151115b61026557600080fd5b600084815260016020818152604080842087855282528084208685528252928390208451815460ff191660ff909116178155848201519281019290925591830151805184936102bb926002850192910190610761565b50505060008481526020818152604080832086845290915290819020905163745fca7b60e01b815273__$b620240993a55615b607189176fb82e62f$__9163745fca7b9161030d918691600401610a19565b60206040518083038186803b15801561032557600080fd5b505af4158015610339573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061035d919061081b565b610389576000848152602081815260408083208684528252822080546001810182559083529120018290555b7f1f1087f8c393b2db10215ffd640a5ba2cda450169f714629afd0d02330e3d102848484846040516103be9493929190610a69565b60405180910390a150505050565b60608360001a60f81b6001600160f81b031916158015906103fc57508260001a60f81b6001600160f81b03191615155b801561041757508160001a60f81b6001600160f81b03191615155b61042057600080fd5b600084815260016020818152604080842087855282528084208685528252928390206002908101805485519481161561010002600019011691909104601f81018390048302840183019094528383529192908301828280156104c35780601f10610498576101008083540402835291602001916104c3565b820191906000526020600020905b8154815290600101906020018083116104a657829003601f168201915b505050505090509392505050565b600083811a60f81b6001600160f81b0319161580159061050057508260001a60f81b6001600160f81b03191615155b801561051b57508160001a60f81b6001600160f81b03191615155b61052457600080fd5b5060009283526001602090815260408085209385529281528284209184525290205460ff1690565b6105546107df565b8360001a60f81b6001600160f81b0319161580159061058257508260001a60f81b6001600160f81b03191615155b801561059d57508160001a60f81b6001600160f81b03191615155b6105a657600080fd5b600084815260016020818152604080842087855282528084208685528252928390208351606081018552815460ff16815281840154818401526002808301805487516101009782161597909702600019011691909104601f81018590048502860185018752808652919592949286019390919083018282801561066a5780601f1061063f5761010080835404028352916020019161066a565b820191906000526020600020905b81548152906001019060200180831161064d57829003601f168201915b50505050508152505090509392505050565b600083811a60f81b6001600160f81b031916158015906106ab57508260001a60f81b6001600160f81b03191615155b80156106cd575060008481526020818152604080832086845290915290205482105b6106d657600080fd5b60008481526020818152604080832086845290915290208054839081106106f957fe5b906000526020600020015490509392505050565b600082811a60f81b6001600160f81b0319161580159061073c57508160001a60f81b6001600160f81b03191615155b61074557600080fd5b5060009182526020828152604080842092845291905290205490565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106107a257805160ff19168380011785556107cf565b828001600101855582156107cf579182015b828111156107cf5782518255916020019190600101906107b4565b506107db9291506107fe565b5090565b6040805160608082018352600080835260208301529181019190915290565b61081891905b808211156107db5760008155600101610804565b90565b60006020828403121561082c578081fd5b8151801515811461083b578182fd5b9392505050565b60008060408385031215610854578081fd5b50508035926020909101359150565b600080600060608486031215610877578081fd5b505081359360208301359350604090920135919050565b600080600080608085870312156108a3578081fd5b84359350602080860135935060408601359250606086013567ffffffffffffffff808211156108d0578384fd5b8188016060818b0312156108e2578485fd5b6108ec6060610acc565b925080356108f981610af3565b83528084013584840152604081013582811115610914578586fd5b8082018b601f820112610925578687fd5b8035925083831115610935578687fd5b610947601f8401601f19168701610acc565b93508284528b8684830101111561095c578687fd5b82868201878601375050810190920192909252604082015292959194509250565b600080600060608486031215610877578283fd5b60008151808452815b818110156109b65760208185018101518683018201520161099a565b818111156109c75782602083870101525b50601f01601f19169290920160200192915050565b600060ff825116835260208201516020840152604082015160606040850152610a086060850182610991565b949350505050565b90815260200190565b60006040820184835260206040818501528185548084526060860191508685528285209350845b81811015610a5c57845483526001948501949284019201610a40565b5090979650505050505050565b600085825284602083015283604083015260806060830152610a8e60808301846109dc565b9695505050505050565b60006020825261083b6020830184610991565b60006020825261083b60208301846109dc565b60ff91909116815260200190565b60405181810167ffffffffffffffff81118282101715610aeb57600080fd5b604052919050565b60ff81168114610b0257600080fd5b5056fea2646970667358221220dfe67811cf6e21d0091469f5f60bee187194ad4762fc42f4e55f939e7c89ac5364736f6c63430006060033"

// DeployIATIActivityDates deploys a new Ethereum contract, binding an instance of IATIActivityDates to it.
func DeployIATIActivityDates(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IATIActivityDates, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivityDatesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	stringsAddr, _, _, _ := DeployStrings(auth, backend)
	IATIActivityDatesBin = strings.Replace(IATIActivityDatesBin, "__$b620240993a55615b607189176fb82e62f$__", stringsAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IATIActivityDatesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IATIActivityDates{IATIActivityDatesCaller: IATIActivityDatesCaller{contract: contract}, IATIActivityDatesTransactor: IATIActivityDatesTransactor{contract: contract}, IATIActivityDatesFilterer: IATIActivityDatesFilterer{contract: contract}}, nil
}

// IATIActivityDates is an auto generated Go binding around an Ethereum contract.
type IATIActivityDates struct {
	IATIActivityDatesCaller     // Read-only binding to the contract
	IATIActivityDatesTransactor // Write-only binding to the contract
	IATIActivityDatesFilterer   // Log filterer for contract events
}

// IATIActivityDatesCaller is an auto generated read-only Go binding around an Ethereum contract.
type IATIActivityDatesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityDatesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IATIActivityDatesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityDatesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IATIActivityDatesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityDatesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IATIActivityDatesSession struct {
	Contract     *IATIActivityDates // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IATIActivityDatesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IATIActivityDatesCallerSession struct {
	Contract *IATIActivityDatesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IATIActivityDatesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IATIActivityDatesTransactorSession struct {
	Contract     *IATIActivityDatesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IATIActivityDatesRaw is an auto generated low-level Go binding around an Ethereum contract.
type IATIActivityDatesRaw struct {
	Contract *IATIActivityDates // Generic contract binding to access the raw methods on
}

// IATIActivityDatesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IATIActivityDatesCallerRaw struct {
	Contract *IATIActivityDatesCaller // Generic read-only contract binding to access the raw methods on
}

// IATIActivityDatesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IATIActivityDatesTransactorRaw struct {
	Contract *IATIActivityDatesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIATIActivityDates creates a new instance of IATIActivityDates, bound to a specific deployed contract.
func NewIATIActivityDates(address common.Address, backend bind.ContractBackend) (*IATIActivityDates, error) {
	contract, err := bindIATIActivityDates(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IATIActivityDates{IATIActivityDatesCaller: IATIActivityDatesCaller{contract: contract}, IATIActivityDatesTransactor: IATIActivityDatesTransactor{contract: contract}, IATIActivityDatesFilterer: IATIActivityDatesFilterer{contract: contract}}, nil
}

// NewIATIActivityDatesCaller creates a new read-only instance of IATIActivityDates, bound to a specific deployed contract.
func NewIATIActivityDatesCaller(address common.Address, caller bind.ContractCaller) (*IATIActivityDatesCaller, error) {
	contract, err := bindIATIActivityDates(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IATIActivityDatesCaller{contract: contract}, nil
}

// NewIATIActivityDatesTransactor creates a new write-only instance of IATIActivityDates, bound to a specific deployed contract.
func NewIATIActivityDatesTransactor(address common.Address, transactor bind.ContractTransactor) (*IATIActivityDatesTransactor, error) {
	contract, err := bindIATIActivityDates(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IATIActivityDatesTransactor{contract: contract}, nil
}

// NewIATIActivityDatesFilterer creates a new log filterer instance of IATIActivityDates, bound to a specific deployed contract.
func NewIATIActivityDatesFilterer(address common.Address, filterer bind.ContractFilterer) (*IATIActivityDatesFilterer, error) {
	contract, err := bindIATIActivityDates(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IATIActivityDatesFilterer{contract: contract}, nil
}

// bindIATIActivityDates binds a generic wrapper to an already deployed contract.
func bindIATIActivityDates(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivityDatesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIActivityDates *IATIActivityDatesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIActivityDates.Contract.IATIActivityDatesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIActivityDates *IATIActivityDatesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIActivityDates.Contract.IATIActivityDatesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIActivityDates *IATIActivityDatesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIActivityDates.Contract.IATIActivityDatesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIActivityDates *IATIActivityDatesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIActivityDates.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIActivityDates *IATIActivityDatesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIActivityDates.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIActivityDates *IATIActivityDatesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIActivityDates.Contract.contract.Transact(opts, method, params...)
}

// GetDate is a free data retrieval call binding the contract method 0xb5e2ffb4.
//
// Solidity: function getDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(ActivityDatesDate)
func (_IATIActivityDates *IATIActivityDatesCaller) GetDate(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) (ActivityDatesDate, error) {
	var (
		ret0 = new(ActivityDatesDate)
	)
	out := ret0
	err := _IATIActivityDates.contract.Call(opts, out, "getDate", _activitiesRef, _activityRef, _dateRef)
	return *ret0, err
}

// GetDate is a free data retrieval call binding the contract method 0xb5e2ffb4.
//
// Solidity: function getDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(ActivityDatesDate)
func (_IATIActivityDates *IATIActivityDatesSession) GetDate(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) (ActivityDatesDate, error) {
	return _IATIActivityDates.Contract.GetDate(&_IATIActivityDates.CallOpts, _activitiesRef, _activityRef, _dateRef)
}

// GetDate is a free data retrieval call binding the contract method 0xb5e2ffb4.
//
// Solidity: function getDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(ActivityDatesDate)
func (_IATIActivityDates *IATIActivityDatesCallerSession) GetDate(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) (ActivityDatesDate, error) {
	return _IATIActivityDates.Contract.GetDate(&_IATIActivityDates.CallOpts, _activitiesRef, _activityRef, _dateRef)
}

// GetISODate is a free data retrieval call binding the contract method 0x5c77487f.
//
// Solidity: function getISODate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(bytes32)
func (_IATIActivityDates *IATIActivityDatesCaller) GetISODate(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityDates.contract.Call(opts, out, "getISODate", _activitiesRef, _activityRef, _dateRef)
	return *ret0, err
}

// GetISODate is a free data retrieval call binding the contract method 0x5c77487f.
//
// Solidity: function getISODate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(bytes32)
func (_IATIActivityDates *IATIActivityDatesSession) GetISODate(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) ([32]byte, error) {
	return _IATIActivityDates.Contract.GetISODate(&_IATIActivityDates.CallOpts, _activitiesRef, _activityRef, _dateRef)
}

// GetISODate is a free data retrieval call binding the contract method 0x5c77487f.
//
// Solidity: function getISODate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(bytes32)
func (_IATIActivityDates *IATIActivityDatesCallerSession) GetISODate(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) ([32]byte, error) {
	return _IATIActivityDates.Contract.GetISODate(&_IATIActivityDates.CallOpts, _activitiesRef, _activityRef, _dateRef)
}

// GetNarrative is a free data retrieval call binding the contract method 0xae13bb7a.
//
// Solidity: function getNarrative(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(string)
func (_IATIActivityDates *IATIActivityDatesCaller) GetNarrative(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IATIActivityDates.contract.Call(opts, out, "getNarrative", _activitiesRef, _activityRef, _dateRef)
	return *ret0, err
}

// GetNarrative is a free data retrieval call binding the contract method 0xae13bb7a.
//
// Solidity: function getNarrative(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(string)
func (_IATIActivityDates *IATIActivityDatesSession) GetNarrative(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) (string, error) {
	return _IATIActivityDates.Contract.GetNarrative(&_IATIActivityDates.CallOpts, _activitiesRef, _activityRef, _dateRef)
}

// GetNarrative is a free data retrieval call binding the contract method 0xae13bb7a.
//
// Solidity: function getNarrative(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(string)
func (_IATIActivityDates *IATIActivityDatesCallerSession) GetNarrative(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) (string, error) {
	return _IATIActivityDates.Contract.GetNarrative(&_IATIActivityDates.CallOpts, _activitiesRef, _activityRef, _dateRef)
}

// GetNumDates is a free data retrieval call binding the contract method 0xffd55577.
//
// Solidity: function getNumDates(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivityDates *IATIActivityDatesCaller) GetNumDates(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIActivityDates.contract.Call(opts, out, "getNumDates", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetNumDates is a free data retrieval call binding the contract method 0xffd55577.
//
// Solidity: function getNumDates(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivityDates *IATIActivityDatesSession) GetNumDates(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _IATIActivityDates.Contract.GetNumDates(&_IATIActivityDates.CallOpts, _activitiesRef, _activityRef)
}

// GetNumDates is a free data retrieval call binding the contract method 0xffd55577.
//
// Solidity: function getNumDates(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivityDates *IATIActivityDatesCallerSession) GetNumDates(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _IATIActivityDates.Contract.GetNumDates(&_IATIActivityDates.CallOpts, _activitiesRef, _activityRef)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivityDates *IATIActivityDatesCaller) GetReference(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityDates.contract.Call(opts, out, "getReference", _activitiesRef, _activityRef, _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivityDates *IATIActivityDatesSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIActivityDates.Contract.GetReference(&_IATIActivityDates.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivityDates *IATIActivityDatesCallerSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIActivityDates.Contract.GetReference(&_IATIActivityDates.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetType is a free data retrieval call binding the contract method 0xb153af81.
//
// Solidity: function getType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(uint8)
func (_IATIActivityDates *IATIActivityDatesCaller) GetType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivityDates.contract.Call(opts, out, "getType", _activitiesRef, _activityRef, _dateRef)
	return *ret0, err
}

// GetType is a free data retrieval call binding the contract method 0xb153af81.
//
// Solidity: function getType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(uint8)
func (_IATIActivityDates *IATIActivityDatesSession) GetType(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) (uint8, error) {
	return _IATIActivityDates.Contract.GetType(&_IATIActivityDates.CallOpts, _activitiesRef, _activityRef, _dateRef)
}

// GetType is a free data retrieval call binding the contract method 0xb153af81.
//
// Solidity: function getType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef) view returns(uint8)
func (_IATIActivityDates *IATIActivityDatesCallerSession) GetType(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte) (uint8, error) {
	return _IATIActivityDates.Contract.GetType(&_IATIActivityDates.CallOpts, _activitiesRef, _activityRef, _dateRef)
}

// SetDate is a paid mutator transaction binding the contract method 0x934fe599.
//
// Solidity: function setDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef, ActivityDatesDate _date) returns()
func (_IATIActivityDates *IATIActivityDatesTransactor) SetDate(opts *bind.TransactOpts, _activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte, _date ActivityDatesDate) (*types.Transaction, error) {
	return _IATIActivityDates.contract.Transact(opts, "setDate", _activitiesRef, _activityRef, _dateRef, _date)
}

// SetDate is a paid mutator transaction binding the contract method 0x934fe599.
//
// Solidity: function setDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef, ActivityDatesDate _date) returns()
func (_IATIActivityDates *IATIActivityDatesSession) SetDate(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte, _date ActivityDatesDate) (*types.Transaction, error) {
	return _IATIActivityDates.Contract.SetDate(&_IATIActivityDates.TransactOpts, _activitiesRef, _activityRef, _dateRef, _date)
}

// SetDate is a paid mutator transaction binding the contract method 0x934fe599.
//
// Solidity: function setDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef, ActivityDatesDate _date) returns()
func (_IATIActivityDates *IATIActivityDatesTransactorSession) SetDate(_activitiesRef [32]byte, _activityRef [32]byte, _dateRef [32]byte, _date ActivityDatesDate) (*types.Transaction, error) {
	return _IATIActivityDates.Contract.SetDate(&_IATIActivityDates.TransactOpts, _activitiesRef, _activityRef, _dateRef, _date)
}

// IATIActivityDatesSetDateIterator is returned from FilterSetDate and is used to iterate over the raw logs and unpacked data for SetDate events raised by the IATIActivityDates contract.
type IATIActivityDatesSetDateIterator struct {
	Event *IATIActivityDatesSetDate // Event containing the contract specifics and raw log

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
func (it *IATIActivityDatesSetDateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IATIActivityDatesSetDate)
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
		it.Event = new(IATIActivityDatesSetDate)
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
func (it *IATIActivityDatesSetDateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IATIActivityDatesSetDateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IATIActivityDatesSetDate represents a SetDate event raised by the IATIActivityDates contract.
type IATIActivityDatesSetDate struct {
	ActivitiesRef [32]byte
	ActivityRef   [32]byte
	DateRef       [32]byte
	Date          ActivityDatesDate
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetDate is a free log retrieval operation binding the contract event 0x1f1087f8c393b2db10215ffd640a5ba2cda450169f714629afd0d02330e3d102.
//
// Solidity: event SetDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef, ActivityDatesDate _date)
func (_IATIActivityDates *IATIActivityDatesFilterer) FilterSetDate(opts *bind.FilterOpts) (*IATIActivityDatesSetDateIterator, error) {

	logs, sub, err := _IATIActivityDates.contract.FilterLogs(opts, "SetDate")
	if err != nil {
		return nil, err
	}
	return &IATIActivityDatesSetDateIterator{contract: _IATIActivityDates.contract, event: "SetDate", logs: logs, sub: sub}, nil
}

// WatchSetDate is a free log subscription operation binding the contract event 0x1f1087f8c393b2db10215ffd640a5ba2cda450169f714629afd0d02330e3d102.
//
// Solidity: event SetDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef, ActivityDatesDate _date)
func (_IATIActivityDates *IATIActivityDatesFilterer) WatchSetDate(opts *bind.WatchOpts, sink chan<- *IATIActivityDatesSetDate) (event.Subscription, error) {

	logs, sub, err := _IATIActivityDates.contract.WatchLogs(opts, "SetDate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IATIActivityDatesSetDate)
				if err := _IATIActivityDates.contract.UnpackLog(event, "SetDate", log); err != nil {
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

// ParseSetDate is a log parse operation binding the contract event 0x1f1087f8c393b2db10215ffd640a5ba2cda450169f714629afd0d02330e3d102.
//
// Solidity: event SetDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _dateRef, ActivityDatesDate _date)
func (_IATIActivityDates *IATIActivityDatesFilterer) ParseSetDate(log types.Log) (*IATIActivityDatesSetDate, error) {
	event := new(IATIActivityDatesSetDate)
	if err := _IATIActivityDates.contract.UnpackLog(event, "SetDate", log); err != nil {
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
