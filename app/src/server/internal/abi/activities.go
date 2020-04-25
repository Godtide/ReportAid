// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// ActivitiesOrgActivities is an auto generated low-level Go binding around an user-defined struct.
type ActivitiesOrgActivities struct {
	Version       [32]byte
	GeneratedTime [32]byte
	LinkedData    [32]byte
}

// ActivitiesABI is the input ABI used to generate the binding from.
const ActivitiesABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structActivities.OrgActivities\",\"name\":\"_activities\",\"type\":\"tuple\"}],\"name\":\"SetActivities\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"}],\"name\":\"getActivities\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"}],\"internalType\":\"structActivities.OrgActivities\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"}],\"name\":\"getGeneratedTime\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"}],\"name\":\"getLinkedData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumActivities\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"}],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"}],\"internalType\":\"structActivities.OrgActivities\",\"name\":\"_activities\",\"type\":\"tuple\"}],\"name\":\"setActivities\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ActivitiesFuncSigs maps the 4-byte function signature to its string representation.
var ActivitiesFuncSigs = map[string]string{
	"81a695a4": "getActivities(bytes32)",
	"ab73eba0": "getGeneratedTime(bytes32)",
	"053cd86f": "getLinkedData(bytes32)",
	"1f5771d9": "getNumActivities()",
	"bc599341": "getReference(uint256)",
	"9aaf9f08": "getVersion(bytes32)",
	"b828c537": "setActivities(bytes32,(bytes32,bytes32,bytes32))",
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

// GetActivities is a free data retrieval call binding the contract method 0x81a695a4.
//
// Solidity: function getActivities(bytes32 _activitiesRef) view returns(ActivitiesOrgActivities)
func (_Activities *ActivitiesCaller) GetActivities(opts *bind.CallOpts, _activitiesRef [32]byte) (ActivitiesOrgActivities, error) {
	var (
		ret0 = new(ActivitiesOrgActivities)
	)
	out := ret0
	err := _Activities.contract.Call(opts, out, "getActivities", _activitiesRef)
	return *ret0, err
}

// GetActivities is a free data retrieval call binding the contract method 0x81a695a4.
//
// Solidity: function getActivities(bytes32 _activitiesRef) view returns(ActivitiesOrgActivities)
func (_Activities *ActivitiesSession) GetActivities(_activitiesRef [32]byte) (ActivitiesOrgActivities, error) {
	return _Activities.Contract.GetActivities(&_Activities.CallOpts, _activitiesRef)
}

// GetActivities is a free data retrieval call binding the contract method 0x81a695a4.
//
// Solidity: function getActivities(bytes32 _activitiesRef) view returns(ActivitiesOrgActivities)
func (_Activities *ActivitiesCallerSession) GetActivities(_activitiesRef [32]byte) (ActivitiesOrgActivities, error) {
	return _Activities.Contract.GetActivities(&_Activities.CallOpts, _activitiesRef)
}

// GetGeneratedTime is a free data retrieval call binding the contract method 0xab73eba0.
//
// Solidity: function getGeneratedTime(bytes32 _activitiesRef) view returns(bytes32)
func (_Activities *ActivitiesCaller) GetGeneratedTime(opts *bind.CallOpts, _activitiesRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Activities.contract.Call(opts, out, "getGeneratedTime", _activitiesRef)
	return *ret0, err
}

// GetGeneratedTime is a free data retrieval call binding the contract method 0xab73eba0.
//
// Solidity: function getGeneratedTime(bytes32 _activitiesRef) view returns(bytes32)
func (_Activities *ActivitiesSession) GetGeneratedTime(_activitiesRef [32]byte) ([32]byte, error) {
	return _Activities.Contract.GetGeneratedTime(&_Activities.CallOpts, _activitiesRef)
}

// GetGeneratedTime is a free data retrieval call binding the contract method 0xab73eba0.
//
// Solidity: function getGeneratedTime(bytes32 _activitiesRef) view returns(bytes32)
func (_Activities *ActivitiesCallerSession) GetGeneratedTime(_activitiesRef [32]byte) ([32]byte, error) {
	return _Activities.Contract.GetGeneratedTime(&_Activities.CallOpts, _activitiesRef)
}

// GetLinkedData is a free data retrieval call binding the contract method 0x053cd86f.
//
// Solidity: function getLinkedData(bytes32 _activitiesRef) view returns(bytes32)
func (_Activities *ActivitiesCaller) GetLinkedData(opts *bind.CallOpts, _activitiesRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Activities.contract.Call(opts, out, "getLinkedData", _activitiesRef)
	return *ret0, err
}

// GetLinkedData is a free data retrieval call binding the contract method 0x053cd86f.
//
// Solidity: function getLinkedData(bytes32 _activitiesRef) view returns(bytes32)
func (_Activities *ActivitiesSession) GetLinkedData(_activitiesRef [32]byte) ([32]byte, error) {
	return _Activities.Contract.GetLinkedData(&_Activities.CallOpts, _activitiesRef)
}

// GetLinkedData is a free data retrieval call binding the contract method 0x053cd86f.
//
// Solidity: function getLinkedData(bytes32 _activitiesRef) view returns(bytes32)
func (_Activities *ActivitiesCallerSession) GetLinkedData(_activitiesRef [32]byte) ([32]byte, error) {
	return _Activities.Contract.GetLinkedData(&_Activities.CallOpts, _activitiesRef)
}

// GetNumActivities is a free data retrieval call binding the contract method 0x1f5771d9.
//
// Solidity: function getNumActivities() view returns(uint256)
func (_Activities *ActivitiesCaller) GetNumActivities(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Activities.contract.Call(opts, out, "getNumActivities")
	return *ret0, err
}

// GetNumActivities is a free data retrieval call binding the contract method 0x1f5771d9.
//
// Solidity: function getNumActivities() view returns(uint256)
func (_Activities *ActivitiesSession) GetNumActivities() (*big.Int, error) {
	return _Activities.Contract.GetNumActivities(&_Activities.CallOpts)
}

// GetNumActivities is a free data retrieval call binding the contract method 0x1f5771d9.
//
// Solidity: function getNumActivities() view returns(uint256)
func (_Activities *ActivitiesCallerSession) GetNumActivities() (*big.Int, error) {
	return _Activities.Contract.GetNumActivities(&_Activities.CallOpts)
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

// GetVersion is a free data retrieval call binding the contract method 0x9aaf9f08.
//
// Solidity: function getVersion(bytes32 _activitiesRef) view returns(bytes32)
func (_Activities *ActivitiesCaller) GetVersion(opts *bind.CallOpts, _activitiesRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Activities.contract.Call(opts, out, "getVersion", _activitiesRef)
	return *ret0, err
}

// GetVersion is a free data retrieval call binding the contract method 0x9aaf9f08.
//
// Solidity: function getVersion(bytes32 _activitiesRef) view returns(bytes32)
func (_Activities *ActivitiesSession) GetVersion(_activitiesRef [32]byte) ([32]byte, error) {
	return _Activities.Contract.GetVersion(&_Activities.CallOpts, _activitiesRef)
}

// GetVersion is a free data retrieval call binding the contract method 0x9aaf9f08.
//
// Solidity: function getVersion(bytes32 _activitiesRef) view returns(bytes32)
func (_Activities *ActivitiesCallerSession) GetVersion(_activitiesRef [32]byte) ([32]byte, error) {
	return _Activities.Contract.GetVersion(&_Activities.CallOpts, _activitiesRef)
}

// SetActivities is a paid mutator transaction binding the contract method 0xb828c537.
//
// Solidity: function setActivities(bytes32 _activitiesRef, ActivitiesOrgActivities _activities) returns()
func (_Activities *ActivitiesTransactor) SetActivities(opts *bind.TransactOpts, _activitiesRef [32]byte, _activities ActivitiesOrgActivities) (*types.Transaction, error) {
	return _Activities.contract.Transact(opts, "setActivities", _activitiesRef, _activities)
}

// SetActivities is a paid mutator transaction binding the contract method 0xb828c537.
//
// Solidity: function setActivities(bytes32 _activitiesRef, ActivitiesOrgActivities _activities) returns()
func (_Activities *ActivitiesSession) SetActivities(_activitiesRef [32]byte, _activities ActivitiesOrgActivities) (*types.Transaction, error) {
	return _Activities.Contract.SetActivities(&_Activities.TransactOpts, _activitiesRef, _activities)
}

// SetActivities is a paid mutator transaction binding the contract method 0xb828c537.
//
// Solidity: function setActivities(bytes32 _activitiesRef, ActivitiesOrgActivities _activities) returns()
func (_Activities *ActivitiesTransactorSession) SetActivities(_activitiesRef [32]byte, _activities ActivitiesOrgActivities) (*types.Transaction, error) {
	return _Activities.Contract.SetActivities(&_Activities.TransactOpts, _activitiesRef, _activities)
}

// ActivitiesSetActivitiesIterator is returned from FilterSetActivities and is used to iterate over the raw logs and unpacked data for SetActivities events raised by the Activities contract.
type ActivitiesSetActivitiesIterator struct {
	Event *ActivitiesSetActivities // Event containing the contract specifics and raw log

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
func (it *ActivitiesSetActivitiesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivitiesSetActivities)
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
		it.Event = new(ActivitiesSetActivities)
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
func (it *ActivitiesSetActivitiesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivitiesSetActivitiesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivitiesSetActivities represents a SetActivities event raised by the Activities contract.
type ActivitiesSetActivities struct {
	ActivitiesRef [32]byte
	Activities    ActivitiesOrgActivities
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetActivities is a free log retrieval operation binding the contract event 0xc6fb585a72eee30aaa08dd8cc66620d02cc0b6238a4d1ad4d9dc6abffdbaf609.
//
// Solidity: event SetActivities(bytes32 _activitiesRef, ActivitiesOrgActivities _activities)
func (_Activities *ActivitiesFilterer) FilterSetActivities(opts *bind.FilterOpts) (*ActivitiesSetActivitiesIterator, error) {

	logs, sub, err := _Activities.contract.FilterLogs(opts, "SetActivities")
	if err != nil {
		return nil, err
	}
	return &ActivitiesSetActivitiesIterator{contract: _Activities.contract, event: "SetActivities", logs: logs, sub: sub}, nil
}

// WatchSetActivities is a free log subscription operation binding the contract event 0xc6fb585a72eee30aaa08dd8cc66620d02cc0b6238a4d1ad4d9dc6abffdbaf609.
//
// Solidity: event SetActivities(bytes32 _activitiesRef, ActivitiesOrgActivities _activities)
func (_Activities *ActivitiesFilterer) WatchSetActivities(opts *bind.WatchOpts, sink chan<- *ActivitiesSetActivities) (event.Subscription, error) {

	logs, sub, err := _Activities.contract.WatchLogs(opts, "SetActivities")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivitiesSetActivities)
				if err := _Activities.contract.UnpackLog(event, "SetActivities", log); err != nil {
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

// ParseSetActivities is a log parse operation binding the contract event 0xc6fb585a72eee30aaa08dd8cc66620d02cc0b6238a4d1ad4d9dc6abffdbaf609.
//
// Solidity: event SetActivities(bytes32 _activitiesRef, ActivitiesOrgActivities _activities)
func (_Activities *ActivitiesFilterer) ParseSetActivities(log types.Log) (*ActivitiesSetActivities, error) {
	event := new(ActivitiesSetActivities)
	if err := _Activities.contract.UnpackLog(event, "SetActivities", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IATIActivitiesABI is the input ABI used to generate the binding from.
const IATIActivitiesABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structActivities.OrgActivities\",\"name\":\"_activities\",\"type\":\"tuple\"}],\"name\":\"SetActivities\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"}],\"name\":\"getActivities\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"}],\"internalType\":\"structActivities.OrgActivities\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"}],\"name\":\"getGeneratedTime\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"}],\"name\":\"getLinkedData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumActivities\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"}],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"}],\"internalType\":\"structActivities.OrgActivities\",\"name\":\"_activities\",\"type\":\"tuple\"}],\"name\":\"setActivities\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IATIActivitiesFuncSigs maps the 4-byte function signature to its string representation.
var IATIActivitiesFuncSigs = map[string]string{
	"81a695a4": "getActivities(bytes32)",
	"ab73eba0": "getGeneratedTime(bytes32)",
	"053cd86f": "getLinkedData(bytes32)",
	"1f5771d9": "getNumActivities()",
	"bc599341": "getReference(uint256)",
	"9aaf9f08": "getVersion(bytes32)",
	"b828c537": "setActivities(bytes32,(bytes32,bytes32,bytes32))",
}

// IATIActivitiesBin is the compiled bytecode used for deploying new contracts.
var IATIActivitiesBin = "0x608060405234801561001057600080fd5b5061055c806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80639aaf9f081161005b5780639aaf9f08146100d3578063ab73eba0146100e6578063b828c537146100f9578063bc5993411461010e5761007d565b8063053cd86f146100825780631f5771d9146100ab57806381a695a4146100b3575b600080fd5b610095610090366004610401565b610121565b6040516100a291906104a5565b60405180910390f35b610095610152565b6100c66100c1366004610401565b610158565b6040516100a29190610512565b6100956100e1366004610401565b6101b2565b6100956100f4366004610401565b6101e0565b61010c610107366004610419565b610212565b005b61009561011c366004610401565b61038c565b600081811a60f81b6001600160f81b03191661013c57600080fd5b5060009081526001602052604090206002015490565b60005490565b6101606103ba565b8160001a60f81b6001600160f81b03191661017a57600080fd5b506000908152600160208181526040928390208351606081018552815481529281015491830191909152600201549181019190915290565b600081811a60f81b6001600160f81b0319166101cd57600080fd5b5060009081526001602052604090205490565b600081811a60f81b6001600160f81b0319166101fb57600080fd5b506000908152600160208190526040909120015490565b8160001a60f81b6001600160f81b031916158015906102415750805160001a60f81b6001600160f81b03191615155b80156102605750602081015160001a60f81b6001600160f81b03191615155b61026957600080fd5b6000828152600160208181526040808420855181559185015192820192909255818401516002909101555163745fca7b60e01b815273__$3ebfd403dd5e707cd4ab529244fde782a8$__9163745fca7b916102c89186916004016104ae565b60206040518083038186803b1580156102e057600080fd5b505af41580156102f4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061031891906103da565b61034f57600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563018290555b7fc6fb585a72eee30aaa08dd8cc66620d02cc0b6238a4d1ad4d9dc6abffdbaf60982826040516103809291906104fe565b60405180910390a15050565b60008054821061039b57600080fd5b600082815481106103a857fe5b90600052602060002001549050919050565b604080516060810182526000808252602082018190529181019190915290565b6000602082840312156103eb578081fd5b815180151581146103fa578182fd5b9392505050565b600060208284031215610412578081fd5b5035919050565b600080828403608081121561042c578182fd5b833592506060601f1982011215610441578182fd5b506040516060810181811067ffffffffffffffff82111715610461578283fd5b8060405250602084013581526040840135602082015260608401356040820152809150509250929050565b8051825260208082015190830152604090810151910152565b90815260200190565b60006040820184835260206040818501528185548084526060860191508685528285209350845b818110156104f1578454835260019485019492840192016104d5565b5090979650505050505050565b828152608081016103fa602083018461048c565b60608101610520828461048c565b9291505056fea2646970667358221220feba22347713ae45874abe2efea054d2c7693271866234d7c3982ff59626519a64736f6c63430006060033"

// DeployIATIActivities deploys a new Ethereum contract, binding an instance of IATIActivities to it.
func DeployIATIActivities(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IATIActivities, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivitiesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	stringsAddr, _, _, _ := DeployStrings(auth, backend)
	IATIActivitiesBin = strings.Replace(IATIActivitiesBin, "__$3ebfd403dd5e707cd4ab529244fde782a8$__", stringsAddr.String()[2:], -1)

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

// GetActivities is a free data retrieval call binding the contract method 0x81a695a4.
//
// Solidity: function getActivities(bytes32 _activitiesRef) view returns(ActivitiesOrgActivities)
func (_IATIActivities *IATIActivitiesCaller) GetActivities(opts *bind.CallOpts, _activitiesRef [32]byte) (ActivitiesOrgActivities, error) {
	var (
		ret0 = new(ActivitiesOrgActivities)
	)
	out := ret0
	err := _IATIActivities.contract.Call(opts, out, "getActivities", _activitiesRef)
	return *ret0, err
}

// GetActivities is a free data retrieval call binding the contract method 0x81a695a4.
//
// Solidity: function getActivities(bytes32 _activitiesRef) view returns(ActivitiesOrgActivities)
func (_IATIActivities *IATIActivitiesSession) GetActivities(_activitiesRef [32]byte) (ActivitiesOrgActivities, error) {
	return _IATIActivities.Contract.GetActivities(&_IATIActivities.CallOpts, _activitiesRef)
}

// GetActivities is a free data retrieval call binding the contract method 0x81a695a4.
//
// Solidity: function getActivities(bytes32 _activitiesRef) view returns(ActivitiesOrgActivities)
func (_IATIActivities *IATIActivitiesCallerSession) GetActivities(_activitiesRef [32]byte) (ActivitiesOrgActivities, error) {
	return _IATIActivities.Contract.GetActivities(&_IATIActivities.CallOpts, _activitiesRef)
}

// GetGeneratedTime is a free data retrieval call binding the contract method 0xab73eba0.
//
// Solidity: function getGeneratedTime(bytes32 _activitiesRef) view returns(bytes32)
func (_IATIActivities *IATIActivitiesCaller) GetGeneratedTime(opts *bind.CallOpts, _activitiesRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivities.contract.Call(opts, out, "getGeneratedTime", _activitiesRef)
	return *ret0, err
}

// GetGeneratedTime is a free data retrieval call binding the contract method 0xab73eba0.
//
// Solidity: function getGeneratedTime(bytes32 _activitiesRef) view returns(bytes32)
func (_IATIActivities *IATIActivitiesSession) GetGeneratedTime(_activitiesRef [32]byte) ([32]byte, error) {
	return _IATIActivities.Contract.GetGeneratedTime(&_IATIActivities.CallOpts, _activitiesRef)
}

// GetGeneratedTime is a free data retrieval call binding the contract method 0xab73eba0.
//
// Solidity: function getGeneratedTime(bytes32 _activitiesRef) view returns(bytes32)
func (_IATIActivities *IATIActivitiesCallerSession) GetGeneratedTime(_activitiesRef [32]byte) ([32]byte, error) {
	return _IATIActivities.Contract.GetGeneratedTime(&_IATIActivities.CallOpts, _activitiesRef)
}

// GetLinkedData is a free data retrieval call binding the contract method 0x053cd86f.
//
// Solidity: function getLinkedData(bytes32 _activitiesRef) view returns(bytes32)
func (_IATIActivities *IATIActivitiesCaller) GetLinkedData(opts *bind.CallOpts, _activitiesRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivities.contract.Call(opts, out, "getLinkedData", _activitiesRef)
	return *ret0, err
}

// GetLinkedData is a free data retrieval call binding the contract method 0x053cd86f.
//
// Solidity: function getLinkedData(bytes32 _activitiesRef) view returns(bytes32)
func (_IATIActivities *IATIActivitiesSession) GetLinkedData(_activitiesRef [32]byte) ([32]byte, error) {
	return _IATIActivities.Contract.GetLinkedData(&_IATIActivities.CallOpts, _activitiesRef)
}

// GetLinkedData is a free data retrieval call binding the contract method 0x053cd86f.
//
// Solidity: function getLinkedData(bytes32 _activitiesRef) view returns(bytes32)
func (_IATIActivities *IATIActivitiesCallerSession) GetLinkedData(_activitiesRef [32]byte) ([32]byte, error) {
	return _IATIActivities.Contract.GetLinkedData(&_IATIActivities.CallOpts, _activitiesRef)
}

// GetNumActivities is a free data retrieval call binding the contract method 0x1f5771d9.
//
// Solidity: function getNumActivities() view returns(uint256)
func (_IATIActivities *IATIActivitiesCaller) GetNumActivities(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIActivities.contract.Call(opts, out, "getNumActivities")
	return *ret0, err
}

// GetNumActivities is a free data retrieval call binding the contract method 0x1f5771d9.
//
// Solidity: function getNumActivities() view returns(uint256)
func (_IATIActivities *IATIActivitiesSession) GetNumActivities() (*big.Int, error) {
	return _IATIActivities.Contract.GetNumActivities(&_IATIActivities.CallOpts)
}

// GetNumActivities is a free data retrieval call binding the contract method 0x1f5771d9.
//
// Solidity: function getNumActivities() view returns(uint256)
func (_IATIActivities *IATIActivitiesCallerSession) GetNumActivities() (*big.Int, error) {
	return _IATIActivities.Contract.GetNumActivities(&_IATIActivities.CallOpts)
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

// GetVersion is a free data retrieval call binding the contract method 0x9aaf9f08.
//
// Solidity: function getVersion(bytes32 _activitiesRef) view returns(bytes32)
func (_IATIActivities *IATIActivitiesCaller) GetVersion(opts *bind.CallOpts, _activitiesRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivities.contract.Call(opts, out, "getVersion", _activitiesRef)
	return *ret0, err
}

// GetVersion is a free data retrieval call binding the contract method 0x9aaf9f08.
//
// Solidity: function getVersion(bytes32 _activitiesRef) view returns(bytes32)
func (_IATIActivities *IATIActivitiesSession) GetVersion(_activitiesRef [32]byte) ([32]byte, error) {
	return _IATIActivities.Contract.GetVersion(&_IATIActivities.CallOpts, _activitiesRef)
}

// GetVersion is a free data retrieval call binding the contract method 0x9aaf9f08.
//
// Solidity: function getVersion(bytes32 _activitiesRef) view returns(bytes32)
func (_IATIActivities *IATIActivitiesCallerSession) GetVersion(_activitiesRef [32]byte) ([32]byte, error) {
	return _IATIActivities.Contract.GetVersion(&_IATIActivities.CallOpts, _activitiesRef)
}

// SetActivities is a paid mutator transaction binding the contract method 0xb828c537.
//
// Solidity: function setActivities(bytes32 _activitiesRef, ActivitiesOrgActivities _activities) returns()
func (_IATIActivities *IATIActivitiesTransactor) SetActivities(opts *bind.TransactOpts, _activitiesRef [32]byte, _activities ActivitiesOrgActivities) (*types.Transaction, error) {
	return _IATIActivities.contract.Transact(opts, "setActivities", _activitiesRef, _activities)
}

// SetActivities is a paid mutator transaction binding the contract method 0xb828c537.
//
// Solidity: function setActivities(bytes32 _activitiesRef, ActivitiesOrgActivities _activities) returns()
func (_IATIActivities *IATIActivitiesSession) SetActivities(_activitiesRef [32]byte, _activities ActivitiesOrgActivities) (*types.Transaction, error) {
	return _IATIActivities.Contract.SetActivities(&_IATIActivities.TransactOpts, _activitiesRef, _activities)
}

// SetActivities is a paid mutator transaction binding the contract method 0xb828c537.
//
// Solidity: function setActivities(bytes32 _activitiesRef, ActivitiesOrgActivities _activities) returns()
func (_IATIActivities *IATIActivitiesTransactorSession) SetActivities(_activitiesRef [32]byte, _activities ActivitiesOrgActivities) (*types.Transaction, error) {
	return _IATIActivities.Contract.SetActivities(&_IATIActivities.TransactOpts, _activitiesRef, _activities)
}

// IATIActivitiesSetActivitiesIterator is returned from FilterSetActivities and is used to iterate over the raw logs and unpacked data for SetActivities events raised by the IATIActivities contract.
type IATIActivitiesSetActivitiesIterator struct {
	Event *IATIActivitiesSetActivities // Event containing the contract specifics and raw log

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
func (it *IATIActivitiesSetActivitiesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IATIActivitiesSetActivities)
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
		it.Event = new(IATIActivitiesSetActivities)
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
func (it *IATIActivitiesSetActivitiesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IATIActivitiesSetActivitiesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IATIActivitiesSetActivities represents a SetActivities event raised by the IATIActivities contract.
type IATIActivitiesSetActivities struct {
	ActivitiesRef [32]byte
	Activities    ActivitiesOrgActivities
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetActivities is a free log retrieval operation binding the contract event 0xc6fb585a72eee30aaa08dd8cc66620d02cc0b6238a4d1ad4d9dc6abffdbaf609.
//
// Solidity: event SetActivities(bytes32 _activitiesRef, ActivitiesOrgActivities _activities)
func (_IATIActivities *IATIActivitiesFilterer) FilterSetActivities(opts *bind.FilterOpts) (*IATIActivitiesSetActivitiesIterator, error) {

	logs, sub, err := _IATIActivities.contract.FilterLogs(opts, "SetActivities")
	if err != nil {
		return nil, err
	}
	return &IATIActivitiesSetActivitiesIterator{contract: _IATIActivities.contract, event: "SetActivities", logs: logs, sub: sub}, nil
}

// WatchSetActivities is a free log subscription operation binding the contract event 0xc6fb585a72eee30aaa08dd8cc66620d02cc0b6238a4d1ad4d9dc6abffdbaf609.
//
// Solidity: event SetActivities(bytes32 _activitiesRef, ActivitiesOrgActivities _activities)
func (_IATIActivities *IATIActivitiesFilterer) WatchSetActivities(opts *bind.WatchOpts, sink chan<- *IATIActivitiesSetActivities) (event.Subscription, error) {

	logs, sub, err := _IATIActivities.contract.WatchLogs(opts, "SetActivities")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IATIActivitiesSetActivities)
				if err := _IATIActivities.contract.UnpackLog(event, "SetActivities", log); err != nil {
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

// ParseSetActivities is a log parse operation binding the contract event 0xc6fb585a72eee30aaa08dd8cc66620d02cc0b6238a4d1ad4d9dc6abffdbaf609.
//
// Solidity: event SetActivities(bytes32 _activitiesRef, ActivitiesOrgActivities _activities)
func (_IATIActivities *IATIActivitiesFilterer) ParseSetActivities(log types.Log) (*IATIActivitiesSetActivities, error) {
	event := new(IATIActivitiesSetActivities)
	if err := _IATIActivities.contract.UnpackLog(event, "SetActivities", log); err != nil {
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
var StringsBin = "0x610662610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061006c5760003560e01c806302a334641461007157806328602ab21461009a5780633a96fdd7146100ba578063745fca7b146100cd578063c55f207f146100e0578063ef0b2368146100f3575b600080fd5b61008461007f36600461046d565b610113565b60405161009191906105dc565b60405180910390f35b6100ad6100a836600461044c565b610180565b604051610091919061057e565b6100ad6100c836600461051d565b6101bd565b6100ad6100db3660046103a8565b61020b565b6100846100ee3660046103a8565b61024b565b610106610101366004610390565b6102b5565b6040516100919190610589565b6000808351118015610126575060008251115b61012f57600080fd5b815160005b835181101561017657606084828151811061014b57fe5b6020026020010151905061015f86826101bd565b1561016d5781925050610176565b50600101610134565b5090505b92915050565b600082811a60f81b6001600160f81b031916158015906101af57508160001a60f81b6001600160f81b03191615155b6101b857600080fd5b501490565b60008083511180156101d0575060008251115b6101d957600080fd5b8151835184918491146101f15760009250505061017a565b80805190602001208280519060200120149250505061017a565b600082811a60f81b6001600160f81b03191661022657600080fd5b81516000901561024457600061023c858561024b565b845114159150505b9392505050565b600082811a60f81b6001600160f81b0319161580159061026c575060008251115b61027557600080fd5b815160005b8351811015610176576102a08585838151811061029357fe5b6020026020010151610180565b156102ad57809150610176565b60010161027a565b60408051602080825281830190925260609182919060208201818036833701905050905060005b6020811015610320578381602081106102f157fe5b1a60f81b82828151811061030157fe5b60200101906001600160f81b031916908160001a9053506001016102dc565b5092915050565b600082601f830112610337578081fd5b813567ffffffffffffffff81111561034d578182fd5b610360601f8201601f19166020016105e5565b915080825283602082850101111561037757600080fd5b8060208401602084013760009082016020015292915050565b6000602082840312156103a1578081fd5b5035919050565b600080604083850312156103ba578081fd5b8235915060208084013567ffffffffffffffff8111156103d8578283fd5b80850186601f8201126103e9578384fd5b803591506103fe6103f98361060c565b6105e5565b82815283810190828501858502840186018a101561041a578687fd5b8693505b8484101561043c57803583526001939093019291850191850161041e565b5080955050505050509250929050565b6000806040838503121561045e578182fd5b50508035926020909101359150565b6000806040838503121561047f578182fd5b823567ffffffffffffffff80821115610496578384fd5b6104a286838701610327565b93506020915081850135818111156104b8578384fd5b85019050601f810186136104ca578283fd5b80356104d86103f98261060c565b81815283810190838501865b8481101561050d576104fb8b888435890101610327565b845292860192908601906001016104e4565b5096999098509650505050505050565b6000806040838503121561052f578182fd5b823567ffffffffffffffff80821115610546578384fd5b61055286838701610327565b93506020850135915080821115610567578283fd5b5061057485828601610327565b9150509250929050565b901515815260200190565b6000602080835283518082850152825b818110156105b557858101830151858201604001528201610599565b818111156105c65783604083870101525b50601f01601f1916929092016040019392505050565b90815260200190565b60405181810167ffffffffffffffff8111828210171561060457600080fd5b604052919050565b600067ffffffffffffffff821115610622578081fd5b506020908102019056fea264697066735822122098efb921c3b0845906c7afa8b8d85f591c0df4271dd9189851a9b6efa8b0a98a64736f6c63430006060033"

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
