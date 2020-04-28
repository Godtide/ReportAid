// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package activityTerritories

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

// ActivityTerritoriesTerritory is an auto generated low-level Go binding around an user-defined struct.
type ActivityTerritoriesTerritory struct {
	Percentage uint8
	Territory  [32]byte
}

// ActivityTerritoriesABI is the input ABI used to generate the binding from.
const ActivityTerritoriesABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_territoryRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"percentage\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"territory\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structActivityTerritories.Territory\",\"name\":\"_territory\",\"type\":\"tuple\"}],\"name\":\"SetTerritory\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_territoryRef\",\"type\":\"bytes32\"}],\"name\":\"getDACTerritory\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_territoryRef\",\"type\":\"bytes32\"}],\"name\":\"getPercentage\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_territoryRef\",\"type\":\"bytes32\"}],\"name\":\"getTerritory\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"percentage\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"territory\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTerritories.Territory\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_territoryRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"percentage\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"territory\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTerritories.Territory\",\"name\":\"_territory\",\"type\":\"tuple\"}],\"name\":\"setTerritory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ActivityTerritoriesFuncSigs maps the 4-byte function signature to its string representation.
var ActivityTerritoriesFuncSigs = map[string]string{
	"42f952e4": "getDACTerritory(bytes32,bytes32,bytes32)",
	"47cdae01": "getNum(bytes32,bytes32)",
	"a8d77aa1": "getPercentage(bytes32,bytes32,bytes32)",
	"eef67d78": "getReference(bytes32,bytes32,uint256)",
	"431c989f": "getTerritory(bytes32,bytes32,bytes32)",
	"6feaba90": "setTerritory(bytes32,bytes32,bytes32,(uint8,bytes32))",
}

// ActivityTerritories is an auto generated Go binding around an Ethereum contract.
type ActivityTerritories struct {
	ActivityTerritoriesCaller     // Read-only binding to the contract
	ActivityTerritoriesTransactor // Write-only binding to the contract
	ActivityTerritoriesFilterer   // Log filterer for contract events
}

// ActivityTerritoriesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivityTerritoriesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityTerritoriesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivityTerritoriesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityTerritoriesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivityTerritoriesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityTerritoriesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivityTerritoriesSession struct {
	Contract     *ActivityTerritories // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ActivityTerritoriesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivityTerritoriesCallerSession struct {
	Contract *ActivityTerritoriesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ActivityTerritoriesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivityTerritoriesTransactorSession struct {
	Contract     *ActivityTerritoriesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ActivityTerritoriesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivityTerritoriesRaw struct {
	Contract *ActivityTerritories // Generic contract binding to access the raw methods on
}

// ActivityTerritoriesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivityTerritoriesCallerRaw struct {
	Contract *ActivityTerritoriesCaller // Generic read-only contract binding to access the raw methods on
}

// ActivityTerritoriesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivityTerritoriesTransactorRaw struct {
	Contract *ActivityTerritoriesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivityTerritories creates a new instance of ActivityTerritories, bound to a specific deployed contract.
func NewActivityTerritories(address common.Address, backend bind.ContractBackend) (*ActivityTerritories, error) {
	contract, err := bindActivityTerritories(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ActivityTerritories{ActivityTerritoriesCaller: ActivityTerritoriesCaller{contract: contract}, ActivityTerritoriesTransactor: ActivityTerritoriesTransactor{contract: contract}, ActivityTerritoriesFilterer: ActivityTerritoriesFilterer{contract: contract}}, nil
}

// NewActivityTerritoriesCaller creates a new read-only instance of ActivityTerritories, bound to a specific deployed contract.
func NewActivityTerritoriesCaller(address common.Address, caller bind.ContractCaller) (*ActivityTerritoriesCaller, error) {
	contract, err := bindActivityTerritories(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityTerritoriesCaller{contract: contract}, nil
}

// NewActivityTerritoriesTransactor creates a new write-only instance of ActivityTerritories, bound to a specific deployed contract.
func NewActivityTerritoriesTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivityTerritoriesTransactor, error) {
	contract, err := bindActivityTerritories(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityTerritoriesTransactor{contract: contract}, nil
}

// NewActivityTerritoriesFilterer creates a new log filterer instance of ActivityTerritories, bound to a specific deployed contract.
func NewActivityTerritoriesFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivityTerritoriesFilterer, error) {
	contract, err := bindActivityTerritories(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivityTerritoriesFilterer{contract: contract}, nil
}

// bindActivityTerritories binds a generic wrapper to an already deployed contract.
func bindActivityTerritories(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivityTerritoriesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityTerritories *ActivityTerritoriesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivityTerritories.Contract.ActivityTerritoriesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityTerritories *ActivityTerritoriesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityTerritories.Contract.ActivityTerritoriesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityTerritories *ActivityTerritoriesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityTerritories.Contract.ActivityTerritoriesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityTerritories *ActivityTerritoriesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivityTerritories.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityTerritories *ActivityTerritoriesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityTerritories.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityTerritories *ActivityTerritoriesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityTerritories.Contract.contract.Transact(opts, method, params...)
}

// GetDACTerritory is a free data retrieval call binding the contract method 0x42f952e4.
//
// Solidity: function getDACTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) view returns(bytes32)
func (_ActivityTerritories *ActivityTerritoriesCaller) GetDACTerritory(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityTerritories.contract.Call(opts, out, "getDACTerritory", _activitiesRef, _activityRef, _territoryRef)
	return *ret0, err
}

// GetDACTerritory is a free data retrieval call binding the contract method 0x42f952e4.
//
// Solidity: function getDACTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) view returns(bytes32)
func (_ActivityTerritories *ActivityTerritoriesSession) GetDACTerritory(_activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte) ([32]byte, error) {
	return _ActivityTerritories.Contract.GetDACTerritory(&_ActivityTerritories.CallOpts, _activitiesRef, _activityRef, _territoryRef)
}

// GetDACTerritory is a free data retrieval call binding the contract method 0x42f952e4.
//
// Solidity: function getDACTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) view returns(bytes32)
func (_ActivityTerritories *ActivityTerritoriesCallerSession) GetDACTerritory(_activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte) ([32]byte, error) {
	return _ActivityTerritories.Contract.GetDACTerritory(&_ActivityTerritories.CallOpts, _activitiesRef, _activityRef, _territoryRef)
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivityTerritories *ActivityTerritoriesCaller) GetNum(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ActivityTerritories.contract.Call(opts, out, "getNum", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivityTerritories *ActivityTerritoriesSession) GetNum(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _ActivityTerritories.Contract.GetNum(&_ActivityTerritories.CallOpts, _activitiesRef, _activityRef)
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivityTerritories *ActivityTerritoriesCallerSession) GetNum(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _ActivityTerritories.Contract.GetNum(&_ActivityTerritories.CallOpts, _activitiesRef, _activityRef)
}

// GetPercentage is a free data retrieval call binding the contract method 0xa8d77aa1.
//
// Solidity: function getPercentage(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) view returns(uint8)
func (_ActivityTerritories *ActivityTerritoriesCaller) GetPercentage(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ActivityTerritories.contract.Call(opts, out, "getPercentage", _activitiesRef, _activityRef, _territoryRef)
	return *ret0, err
}

// GetPercentage is a free data retrieval call binding the contract method 0xa8d77aa1.
//
// Solidity: function getPercentage(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) view returns(uint8)
func (_ActivityTerritories *ActivityTerritoriesSession) GetPercentage(_activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte) (uint8, error) {
	return _ActivityTerritories.Contract.GetPercentage(&_ActivityTerritories.CallOpts, _activitiesRef, _activityRef, _territoryRef)
}

// GetPercentage is a free data retrieval call binding the contract method 0xa8d77aa1.
//
// Solidity: function getPercentage(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) view returns(uint8)
func (_ActivityTerritories *ActivityTerritoriesCallerSession) GetPercentage(_activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte) (uint8, error) {
	return _ActivityTerritories.Contract.GetPercentage(&_ActivityTerritories.CallOpts, _activitiesRef, _activityRef, _territoryRef)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivityTerritories *ActivityTerritoriesCaller) GetReference(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityTerritories.contract.Call(opts, out, "getReference", _activitiesRef, _activityRef, _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivityTerritories *ActivityTerritoriesSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _ActivityTerritories.Contract.GetReference(&_ActivityTerritories.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivityTerritories *ActivityTerritoriesCallerSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _ActivityTerritories.Contract.GetReference(&_ActivityTerritories.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetTerritory is a free data retrieval call binding the contract method 0x431c989f.
//
// Solidity: function getTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) view returns(ActivityTerritoriesTerritory)
func (_ActivityTerritories *ActivityTerritoriesCaller) GetTerritory(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte) (ActivityTerritoriesTerritory, error) {
	var (
		ret0 = new(ActivityTerritoriesTerritory)
	)
	out := ret0
	err := _ActivityTerritories.contract.Call(opts, out, "getTerritory", _activitiesRef, _activityRef, _territoryRef)
	return *ret0, err
}

// GetTerritory is a free data retrieval call binding the contract method 0x431c989f.
//
// Solidity: function getTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) view returns(ActivityTerritoriesTerritory)
func (_ActivityTerritories *ActivityTerritoriesSession) GetTerritory(_activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte) (ActivityTerritoriesTerritory, error) {
	return _ActivityTerritories.Contract.GetTerritory(&_ActivityTerritories.CallOpts, _activitiesRef, _activityRef, _territoryRef)
}

// GetTerritory is a free data retrieval call binding the contract method 0x431c989f.
//
// Solidity: function getTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) view returns(ActivityTerritoriesTerritory)
func (_ActivityTerritories *ActivityTerritoriesCallerSession) GetTerritory(_activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte) (ActivityTerritoriesTerritory, error) {
	return _ActivityTerritories.Contract.GetTerritory(&_ActivityTerritories.CallOpts, _activitiesRef, _activityRef, _territoryRef)
}

// SetTerritory is a paid mutator transaction binding the contract method 0x6feaba90.
//
// Solidity: function setTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef, ActivityTerritoriesTerritory _territory) returns()
func (_ActivityTerritories *ActivityTerritoriesTransactor) SetTerritory(opts *bind.TransactOpts, _activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte, _territory ActivityTerritoriesTerritory) (*types.Transaction, error) {
	return _ActivityTerritories.contract.Transact(opts, "setTerritory", _activitiesRef, _activityRef, _territoryRef, _territory)
}

// SetTerritory is a paid mutator transaction binding the contract method 0x6feaba90.
//
// Solidity: function setTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef, ActivityTerritoriesTerritory _territory) returns()
func (_ActivityTerritories *ActivityTerritoriesSession) SetTerritory(_activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte, _territory ActivityTerritoriesTerritory) (*types.Transaction, error) {
	return _ActivityTerritories.Contract.SetTerritory(&_ActivityTerritories.TransactOpts, _activitiesRef, _activityRef, _territoryRef, _territory)
}

// SetTerritory is a paid mutator transaction binding the contract method 0x6feaba90.
//
// Solidity: function setTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef, ActivityTerritoriesTerritory _territory) returns()
func (_ActivityTerritories *ActivityTerritoriesTransactorSession) SetTerritory(_activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte, _territory ActivityTerritoriesTerritory) (*types.Transaction, error) {
	return _ActivityTerritories.Contract.SetTerritory(&_ActivityTerritories.TransactOpts, _activitiesRef, _activityRef, _territoryRef, _territory)
}

// ActivityTerritoriesSetTerritoryIterator is returned from FilterSetTerritory and is used to iterate over the raw logs and unpacked data for SetTerritory events raised by the ActivityTerritories contract.
type ActivityTerritoriesSetTerritoryIterator struct {
	Event *ActivityTerritoriesSetTerritory // Event containing the contract specifics and raw log

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
func (it *ActivityTerritoriesSetTerritoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivityTerritoriesSetTerritory)
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
		it.Event = new(ActivityTerritoriesSetTerritory)
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
func (it *ActivityTerritoriesSetTerritoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivityTerritoriesSetTerritoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivityTerritoriesSetTerritory represents a SetTerritory event raised by the ActivityTerritories contract.
type ActivityTerritoriesSetTerritory struct {
	ActivitiesRef [32]byte
	ActivityRef   [32]byte
	TerritoryRef  [32]byte
	Territory     ActivityTerritoriesTerritory
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTerritory is a free log retrieval operation binding the contract event 0xc26957b4ab39c91f14b8e349e368f3a59dada03f580bd816ef5ac8b65656aa51.
//
// Solidity: event SetTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef, ActivityTerritoriesTerritory _territory)
func (_ActivityTerritories *ActivityTerritoriesFilterer) FilterSetTerritory(opts *bind.FilterOpts) (*ActivityTerritoriesSetTerritoryIterator, error) {

	logs, sub, err := _ActivityTerritories.contract.FilterLogs(opts, "SetTerritory")
	if err != nil {
		return nil, err
	}
	return &ActivityTerritoriesSetTerritoryIterator{contract: _ActivityTerritories.contract, event: "SetTerritory", logs: logs, sub: sub}, nil
}

// WatchSetTerritory is a free log subscription operation binding the contract event 0xc26957b4ab39c91f14b8e349e368f3a59dada03f580bd816ef5ac8b65656aa51.
//
// Solidity: event SetTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef, ActivityTerritoriesTerritory _territory)
func (_ActivityTerritories *ActivityTerritoriesFilterer) WatchSetTerritory(opts *bind.WatchOpts, sink chan<- *ActivityTerritoriesSetTerritory) (event.Subscription, error) {

	logs, sub, err := _ActivityTerritories.contract.WatchLogs(opts, "SetTerritory")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivityTerritoriesSetTerritory)
				if err := _ActivityTerritories.contract.UnpackLog(event, "SetTerritory", log); err != nil {
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

// ParseSetTerritory is a log parse operation binding the contract event 0xc26957b4ab39c91f14b8e349e368f3a59dada03f580bd816ef5ac8b65656aa51.
//
// Solidity: event SetTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef, ActivityTerritoriesTerritory _territory)
func (_ActivityTerritories *ActivityTerritoriesFilterer) ParseSetTerritory(log types.Log) (*ActivityTerritoriesSetTerritory, error) {
	event := new(ActivityTerritoriesSetTerritory)
	if err := _ActivityTerritories.contract.UnpackLog(event, "SetTerritory", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IATIActivityTerritoriesABI is the input ABI used to generate the binding from.
const IATIActivityTerritoriesABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_territoryRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"percentage\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"territory\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structActivityTerritories.Territory\",\"name\":\"_territory\",\"type\":\"tuple\"}],\"name\":\"SetTerritory\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_territoryRef\",\"type\":\"bytes32\"}],\"name\":\"getDACTerritory\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_territoryRef\",\"type\":\"bytes32\"}],\"name\":\"getPercentage\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_territoryRef\",\"type\":\"bytes32\"}],\"name\":\"getTerritory\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"percentage\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"territory\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTerritories.Territory\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_territoryRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"percentage\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"territory\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTerritories.Territory\",\"name\":\"_territory\",\"type\":\"tuple\"}],\"name\":\"setTerritory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IATIActivityTerritoriesFuncSigs maps the 4-byte function signature to its string representation.
var IATIActivityTerritoriesFuncSigs = map[string]string{
	"42f952e4": "getDACTerritory(bytes32,bytes32,bytes32)",
	"47cdae01": "getNum(bytes32,bytes32)",
	"a8d77aa1": "getPercentage(bytes32,bytes32,bytes32)",
	"eef67d78": "getReference(bytes32,bytes32,uint256)",
	"431c989f": "getTerritory(bytes32,bytes32,bytes32)",
	"6feaba90": "setTerritory(bytes32,bytes32,bytes32,(uint8,bytes32))",
}

// IATIActivityTerritoriesBin is the compiled bytecode used for deploying new contracts.
var IATIActivityTerritoriesBin = "0x608060405234801561001057600080fd5b5061075a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806342f952e414610067578063431c989f1461009057806347cdae01146100b05780636feaba90146100c3578063a8d77aa1146100d8578063eef67d78146100f8575b600080fd5b61007a6100753660046105ac565b61010b565b6040516100879190610684565b60405180910390f35b6100a361009e3660046105ac565b610187565b6040516100879190610708565b61007a6100be36600461058b565b610222565b6100d66100d13660046105d7565b61027a565b005b6100eb6100e63660046105ac565b610441565b6040516100879190610716565b61007a61010636600461065e565b6104bc565b600083811a60f81b6001600160f81b0319161580159061013a57508260001a60f81b6001600160f81b03191615155b801561015557508160001a60f81b6001600160f81b03191615155b61015e57600080fd5b506000928352600160208181526040808620948652938152838520928552919091529120015490565b61018f61054d565b8360001a60f81b6001600160f81b031916158015906101bd57508260001a60f81b6001600160f81b03191615155b80156101d857508160001a60f81b6001600160f81b03191615155b6101e157600080fd5b50600083815260016020818152604080842086855282528084208585528252928390208351808501909452805460ff16845290910154908201529392505050565b600082811a60f81b6001600160f81b0319161580159061025157508160001a60f81b6001600160f81b03191615155b61025a57600080fd5b506000828152602081815260408083208484529091529020545b92915050565b8360001a60f81b6001600160f81b031916158015906102a857508260001a60f81b6001600160f81b03191615155b80156102c357508160001a60f81b6001600160f81b03191615155b80156102e25750602081015160001a60f81b6001600160f81b03191615155b80156102f657506064816000015160ff1611155b6102ff57600080fd5b6000848152600160208181526040808420878552825280842086855282528084208551815460ff191660ff90911617815585830151930192909255868352828152818320868452905290819020905163745fca7b60e01b815273__$b620240993a55615b607189176fb82e62f$__9163745fca7b9161038291869160040161068d565b60206040518083038186803b15801561039a57600080fd5b505af41580156103ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103d29190610564565b6103fe576000848152602081815260408083208684528252822080546001810182559083529120018290555b7fc26957b4ab39c91f14b8e349e368f3a59dada03f580bd816ef5ac8b65656aa518484848460405161043394939291906106dd565b60405180910390a150505050565b600083811a60f81b6001600160f81b0319161580159061047057508260001a60f81b6001600160f81b03191615155b801561048b57508160001a60f81b6001600160f81b03191615155b61049457600080fd5b5060009283526001602090815260408085209385529281528284209184525290205460ff1690565b600083811a60f81b6001600160f81b031916158015906104eb57508260001a60f81b6001600160f81b03191615155b801561050d575060008481526020818152604080832086845290915290205482105b61051657600080fd5b600084815260208181526040808320868452909152902080548390811061053957fe5b906000526020600020015490509392505050565b604080518082019091526000808252602082015290565b600060208284031215610575578081fd5b81518015158114610584578182fd5b9392505050565b6000806040838503121561059d578081fd5b50508035926020909101359150565b6000806000606084860312156105c0578081fd5b505081359360208301359350604090920135919050565b60008060008084860360a08112156105ed578182fd5b85359450602086013593506040808701359350605f198201121561060f578182fd5b506040516040810181811067ffffffffffffffff8211171561062f578283fd5b604052606086013560ff81168114610645578283fd5b8152608095909501356020860152509194909350909190565b6000806000606084860312156105c0578283fd5b805160ff168252602090810151910152565b90815260200190565b60006040820184835260206040818501528185548084526060860191508685528285209350845b818110156106d0578454835260019485019492840192016106b4565b5090979650505050505050565b848152602081018490526040810183905260a081016106ff6060830184610672565b95945050505050565b604081016102748284610672565b60ff9190911681526020019056fea2646970667358221220dc45ad27834f1da4c1325efc80947b00b9e848fc13c34b638ab694980b762e7664736f6c63430006060033"

// DeployIATIActivityTerritories deploys a new Ethereum contract, binding an instance of IATIActivityTerritories to it.
func DeployIATIActivityTerritories(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IATIActivityTerritories, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivityTerritoriesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	stringsAddr, _, _, _ := DeployStrings(auth, backend)
	IATIActivityTerritoriesBin = strings.Replace(IATIActivityTerritoriesBin, "__$b620240993a55615b607189176fb82e62f$__", stringsAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IATIActivityTerritoriesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IATIActivityTerritories{IATIActivityTerritoriesCaller: IATIActivityTerritoriesCaller{contract: contract}, IATIActivityTerritoriesTransactor: IATIActivityTerritoriesTransactor{contract: contract}, IATIActivityTerritoriesFilterer: IATIActivityTerritoriesFilterer{contract: contract}}, nil
}

// IATIActivityTerritories is an auto generated Go binding around an Ethereum contract.
type IATIActivityTerritories struct {
	IATIActivityTerritoriesCaller     // Read-only binding to the contract
	IATIActivityTerritoriesTransactor // Write-only binding to the contract
	IATIActivityTerritoriesFilterer   // Log filterer for contract events
}

// IATIActivityTerritoriesCaller is an auto generated read-only Go binding around an Ethereum contract.
type IATIActivityTerritoriesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityTerritoriesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IATIActivityTerritoriesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityTerritoriesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IATIActivityTerritoriesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityTerritoriesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IATIActivityTerritoriesSession struct {
	Contract     *IATIActivityTerritories // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IATIActivityTerritoriesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IATIActivityTerritoriesCallerSession struct {
	Contract *IATIActivityTerritoriesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// IATIActivityTerritoriesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IATIActivityTerritoriesTransactorSession struct {
	Contract     *IATIActivityTerritoriesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// IATIActivityTerritoriesRaw is an auto generated low-level Go binding around an Ethereum contract.
type IATIActivityTerritoriesRaw struct {
	Contract *IATIActivityTerritories // Generic contract binding to access the raw methods on
}

// IATIActivityTerritoriesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IATIActivityTerritoriesCallerRaw struct {
	Contract *IATIActivityTerritoriesCaller // Generic read-only contract binding to access the raw methods on
}

// IATIActivityTerritoriesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IATIActivityTerritoriesTransactorRaw struct {
	Contract *IATIActivityTerritoriesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIATIActivityTerritories creates a new instance of IATIActivityTerritories, bound to a specific deployed contract.
func NewIATIActivityTerritories(address common.Address, backend bind.ContractBackend) (*IATIActivityTerritories, error) {
	contract, err := bindIATIActivityTerritories(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IATIActivityTerritories{IATIActivityTerritoriesCaller: IATIActivityTerritoriesCaller{contract: contract}, IATIActivityTerritoriesTransactor: IATIActivityTerritoriesTransactor{contract: contract}, IATIActivityTerritoriesFilterer: IATIActivityTerritoriesFilterer{contract: contract}}, nil
}

// NewIATIActivityTerritoriesCaller creates a new read-only instance of IATIActivityTerritories, bound to a specific deployed contract.
func NewIATIActivityTerritoriesCaller(address common.Address, caller bind.ContractCaller) (*IATIActivityTerritoriesCaller, error) {
	contract, err := bindIATIActivityTerritories(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IATIActivityTerritoriesCaller{contract: contract}, nil
}

// NewIATIActivityTerritoriesTransactor creates a new write-only instance of IATIActivityTerritories, bound to a specific deployed contract.
func NewIATIActivityTerritoriesTransactor(address common.Address, transactor bind.ContractTransactor) (*IATIActivityTerritoriesTransactor, error) {
	contract, err := bindIATIActivityTerritories(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IATIActivityTerritoriesTransactor{contract: contract}, nil
}

// NewIATIActivityTerritoriesFilterer creates a new log filterer instance of IATIActivityTerritories, bound to a specific deployed contract.
func NewIATIActivityTerritoriesFilterer(address common.Address, filterer bind.ContractFilterer) (*IATIActivityTerritoriesFilterer, error) {
	contract, err := bindIATIActivityTerritories(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IATIActivityTerritoriesFilterer{contract: contract}, nil
}

// bindIATIActivityTerritories binds a generic wrapper to an already deployed contract.
func bindIATIActivityTerritories(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivityTerritoriesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIActivityTerritories *IATIActivityTerritoriesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIActivityTerritories.Contract.IATIActivityTerritoriesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIActivityTerritories *IATIActivityTerritoriesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIActivityTerritories.Contract.IATIActivityTerritoriesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIActivityTerritories *IATIActivityTerritoriesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIActivityTerritories.Contract.IATIActivityTerritoriesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIActivityTerritories *IATIActivityTerritoriesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIActivityTerritories.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIActivityTerritories *IATIActivityTerritoriesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIActivityTerritories.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIActivityTerritories *IATIActivityTerritoriesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIActivityTerritories.Contract.contract.Transact(opts, method, params...)
}

// GetDACTerritory is a free data retrieval call binding the contract method 0x42f952e4.
//
// Solidity: function getDACTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) view returns(bytes32)
func (_IATIActivityTerritories *IATIActivityTerritoriesCaller) GetDACTerritory(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityTerritories.contract.Call(opts, out, "getDACTerritory", _activitiesRef, _activityRef, _territoryRef)
	return *ret0, err
}

// GetDACTerritory is a free data retrieval call binding the contract method 0x42f952e4.
//
// Solidity: function getDACTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) view returns(bytes32)
func (_IATIActivityTerritories *IATIActivityTerritoriesSession) GetDACTerritory(_activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte) ([32]byte, error) {
	return _IATIActivityTerritories.Contract.GetDACTerritory(&_IATIActivityTerritories.CallOpts, _activitiesRef, _activityRef, _territoryRef)
}

// GetDACTerritory is a free data retrieval call binding the contract method 0x42f952e4.
//
// Solidity: function getDACTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) view returns(bytes32)
func (_IATIActivityTerritories *IATIActivityTerritoriesCallerSession) GetDACTerritory(_activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte) ([32]byte, error) {
	return _IATIActivityTerritories.Contract.GetDACTerritory(&_IATIActivityTerritories.CallOpts, _activitiesRef, _activityRef, _territoryRef)
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivityTerritories *IATIActivityTerritoriesCaller) GetNum(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIActivityTerritories.contract.Call(opts, out, "getNum", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivityTerritories *IATIActivityTerritoriesSession) GetNum(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _IATIActivityTerritories.Contract.GetNum(&_IATIActivityTerritories.CallOpts, _activitiesRef, _activityRef)
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivityTerritories *IATIActivityTerritoriesCallerSession) GetNum(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _IATIActivityTerritories.Contract.GetNum(&_IATIActivityTerritories.CallOpts, _activitiesRef, _activityRef)
}

// GetPercentage is a free data retrieval call binding the contract method 0xa8d77aa1.
//
// Solidity: function getPercentage(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) view returns(uint8)
func (_IATIActivityTerritories *IATIActivityTerritoriesCaller) GetPercentage(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivityTerritories.contract.Call(opts, out, "getPercentage", _activitiesRef, _activityRef, _territoryRef)
	return *ret0, err
}

// GetPercentage is a free data retrieval call binding the contract method 0xa8d77aa1.
//
// Solidity: function getPercentage(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) view returns(uint8)
func (_IATIActivityTerritories *IATIActivityTerritoriesSession) GetPercentage(_activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte) (uint8, error) {
	return _IATIActivityTerritories.Contract.GetPercentage(&_IATIActivityTerritories.CallOpts, _activitiesRef, _activityRef, _territoryRef)
}

// GetPercentage is a free data retrieval call binding the contract method 0xa8d77aa1.
//
// Solidity: function getPercentage(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) view returns(uint8)
func (_IATIActivityTerritories *IATIActivityTerritoriesCallerSession) GetPercentage(_activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte) (uint8, error) {
	return _IATIActivityTerritories.Contract.GetPercentage(&_IATIActivityTerritories.CallOpts, _activitiesRef, _activityRef, _territoryRef)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivityTerritories *IATIActivityTerritoriesCaller) GetReference(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityTerritories.contract.Call(opts, out, "getReference", _activitiesRef, _activityRef, _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivityTerritories *IATIActivityTerritoriesSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIActivityTerritories.Contract.GetReference(&_IATIActivityTerritories.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivityTerritories *IATIActivityTerritoriesCallerSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIActivityTerritories.Contract.GetReference(&_IATIActivityTerritories.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetTerritory is a free data retrieval call binding the contract method 0x431c989f.
//
// Solidity: function getTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) view returns(ActivityTerritoriesTerritory)
func (_IATIActivityTerritories *IATIActivityTerritoriesCaller) GetTerritory(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte) (ActivityTerritoriesTerritory, error) {
	var (
		ret0 = new(ActivityTerritoriesTerritory)
	)
	out := ret0
	err := _IATIActivityTerritories.contract.Call(opts, out, "getTerritory", _activitiesRef, _activityRef, _territoryRef)
	return *ret0, err
}

// GetTerritory is a free data retrieval call binding the contract method 0x431c989f.
//
// Solidity: function getTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) view returns(ActivityTerritoriesTerritory)
func (_IATIActivityTerritories *IATIActivityTerritoriesSession) GetTerritory(_activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte) (ActivityTerritoriesTerritory, error) {
	return _IATIActivityTerritories.Contract.GetTerritory(&_IATIActivityTerritories.CallOpts, _activitiesRef, _activityRef, _territoryRef)
}

// GetTerritory is a free data retrieval call binding the contract method 0x431c989f.
//
// Solidity: function getTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef) view returns(ActivityTerritoriesTerritory)
func (_IATIActivityTerritories *IATIActivityTerritoriesCallerSession) GetTerritory(_activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte) (ActivityTerritoriesTerritory, error) {
	return _IATIActivityTerritories.Contract.GetTerritory(&_IATIActivityTerritories.CallOpts, _activitiesRef, _activityRef, _territoryRef)
}

// SetTerritory is a paid mutator transaction binding the contract method 0x6feaba90.
//
// Solidity: function setTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef, ActivityTerritoriesTerritory _territory) returns()
func (_IATIActivityTerritories *IATIActivityTerritoriesTransactor) SetTerritory(opts *bind.TransactOpts, _activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte, _territory ActivityTerritoriesTerritory) (*types.Transaction, error) {
	return _IATIActivityTerritories.contract.Transact(opts, "setTerritory", _activitiesRef, _activityRef, _territoryRef, _territory)
}

// SetTerritory is a paid mutator transaction binding the contract method 0x6feaba90.
//
// Solidity: function setTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef, ActivityTerritoriesTerritory _territory) returns()
func (_IATIActivityTerritories *IATIActivityTerritoriesSession) SetTerritory(_activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte, _territory ActivityTerritoriesTerritory) (*types.Transaction, error) {
	return _IATIActivityTerritories.Contract.SetTerritory(&_IATIActivityTerritories.TransactOpts, _activitiesRef, _activityRef, _territoryRef, _territory)
}

// SetTerritory is a paid mutator transaction binding the contract method 0x6feaba90.
//
// Solidity: function setTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef, ActivityTerritoriesTerritory _territory) returns()
func (_IATIActivityTerritories *IATIActivityTerritoriesTransactorSession) SetTerritory(_activitiesRef [32]byte, _activityRef [32]byte, _territoryRef [32]byte, _territory ActivityTerritoriesTerritory) (*types.Transaction, error) {
	return _IATIActivityTerritories.Contract.SetTerritory(&_IATIActivityTerritories.TransactOpts, _activitiesRef, _activityRef, _territoryRef, _territory)
}

// IATIActivityTerritoriesSetTerritoryIterator is returned from FilterSetTerritory and is used to iterate over the raw logs and unpacked data for SetTerritory events raised by the IATIActivityTerritories contract.
type IATIActivityTerritoriesSetTerritoryIterator struct {
	Event *IATIActivityTerritoriesSetTerritory // Event containing the contract specifics and raw log

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
func (it *IATIActivityTerritoriesSetTerritoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IATIActivityTerritoriesSetTerritory)
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
		it.Event = new(IATIActivityTerritoriesSetTerritory)
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
func (it *IATIActivityTerritoriesSetTerritoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IATIActivityTerritoriesSetTerritoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IATIActivityTerritoriesSetTerritory represents a SetTerritory event raised by the IATIActivityTerritories contract.
type IATIActivityTerritoriesSetTerritory struct {
	ActivitiesRef [32]byte
	ActivityRef   [32]byte
	TerritoryRef  [32]byte
	Territory     ActivityTerritoriesTerritory
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTerritory is a free log retrieval operation binding the contract event 0xc26957b4ab39c91f14b8e349e368f3a59dada03f580bd816ef5ac8b65656aa51.
//
// Solidity: event SetTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef, ActivityTerritoriesTerritory _territory)
func (_IATIActivityTerritories *IATIActivityTerritoriesFilterer) FilterSetTerritory(opts *bind.FilterOpts) (*IATIActivityTerritoriesSetTerritoryIterator, error) {

	logs, sub, err := _IATIActivityTerritories.contract.FilterLogs(opts, "SetTerritory")
	if err != nil {
		return nil, err
	}
	return &IATIActivityTerritoriesSetTerritoryIterator{contract: _IATIActivityTerritories.contract, event: "SetTerritory", logs: logs, sub: sub}, nil
}

// WatchSetTerritory is a free log subscription operation binding the contract event 0xc26957b4ab39c91f14b8e349e368f3a59dada03f580bd816ef5ac8b65656aa51.
//
// Solidity: event SetTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef, ActivityTerritoriesTerritory _territory)
func (_IATIActivityTerritories *IATIActivityTerritoriesFilterer) WatchSetTerritory(opts *bind.WatchOpts, sink chan<- *IATIActivityTerritoriesSetTerritory) (event.Subscription, error) {

	logs, sub, err := _IATIActivityTerritories.contract.WatchLogs(opts, "SetTerritory")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IATIActivityTerritoriesSetTerritory)
				if err := _IATIActivityTerritories.contract.UnpackLog(event, "SetTerritory", log); err != nil {
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

// ParseSetTerritory is a log parse operation binding the contract event 0xc26957b4ab39c91f14b8e349e368f3a59dada03f580bd816ef5ac8b65656aa51.
//
// Solidity: event SetTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _territoryRef, ActivityTerritoriesTerritory _territory)
func (_IATIActivityTerritories *IATIActivityTerritoriesFilterer) ParseSetTerritory(log types.Log) (*IATIActivityTerritoriesSetTerritory, error) {
	event := new(IATIActivityTerritoriesSetTerritory)
	if err := _IATIActivityTerritories.contract.UnpackLog(event, "SetTerritory", log); err != nil {
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
