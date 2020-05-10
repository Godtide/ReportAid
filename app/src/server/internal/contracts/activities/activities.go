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

// ActivitiesFactoryABI is the input ABI used to generate the binding from.
const ActivitiesFactoryABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_elementType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"Set\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"}],\"internalType\":\"structActivitiesData\",\"name\":\"_activities\",\"type\":\"tuple\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setter\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ActivitiesFactoryFuncSigs maps the 4-byte function signature to its string representation.
var ActivitiesFactoryFuncSigs = map[string]string{
	"8eaa6ac0": "get(bytes32)",
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
	"91a95fbc": "set(bytes32,(bytes32,bytes32,bytes32))",
	"3f3108f7": "setter()",
}

// ActivitiesFactoryBin is the compiled bytecode used for deploying new contracts.
var ActivitiesFactoryBin = "0x608060405234801561001057600080fd5b506000600255610638806100256000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80633f3108f71461005c57806367e0badb1461007a5780638eaa6ac01461008f57806391a95fbc146100af578063bc599341146100c4575b600080fd5b6100646100d7565b604051610071919061036e565b60405180910390f35b6100826100e2565b6040516100719190610365565b6100a261009d3660046102c6565b6100e8565b6040516100719190610351565b6100c26100bd3660046102de565b610146565b005b6100826100d23660046102c6565b6101ea565b63246a57ef60e21b90565b60025490565b600081811a60f81b6001600160f81b03191661010357600080fd5b6000828152602081905260409020600101546001600160a01b031661012757600080fd5b506000908152602081905260409020600101546001600160a01b031690565b8160001a60f81b6001600160f81b03191661016057600080fd5b60008160405161016f906102b9565b6101799190610383565b604051809103906000f080158015610195573d6000803e3d6000fd5b5090506101aa6000848363ffffffff61022216565b507f8353a4d574992e9eb676a31cd326ce2f9f6c4829c3b84ae33ee9febbb2962e7b6002846040516101dd9291906103a4565b60405180910390a1505050565b60025460009082106101fb57600080fd5b600180548390811061020957fe5b9060005260206000209060020201600001549050919050565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b03851617905580156102605760019150506102b2565b506001808501805491820180825560008681526020889052604090208190558591908390811061028c57fe5b600091825260208220600291820201929092559086018054600101905591506102b29050565b9392505050565b61024b806103b883390190565b6000602082840312156102d7578081fd5b5035919050565b60008082840360808112156102f1578182fd5b833592506060601f1982011215610306578182fd5b506040516060810181811067ffffffffffffffff82111715610326578283fd5b8060405250602084013581526040840135602082015260608401356040820152809150509250929050565b6001600160a01b0391909116815260200190565b90815260200190565b6001600160e01b031991909116815260200190565b81518152602080830151908201526040918201519181019190915260600190565b60ff92909216825260208201526040019056fe608060405234801561001057600080fd5b5060405161024b38038061024b83398101604081905261002f916100a3565b805160001a60f81b6001600160f81b031916158015906100625750602081015160001a60f81b6001600160f81b03191615155b80156100815750604081015160001a60f81b6001600160f81b03191615155b61008a57600080fd5b80516000556020810151600155604001516002556100f9565b6000606082840312156100b4578081fd5b604051606081016001600160401b03811182821017156100d2578283fd5b80604052508251815260208301516020820152604083015160408201528091505092915050565b610143806101086000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636d4ce63c1461003b578063993a04b71461005a575b600080fd5b61004361006f565b6040516100519291906100e1565b60405180910390f35b6100626100a1565b60405161005191906100cc565b60006100796100ac565b5050604080516060810182526000548152600154602082015260028054928201929092529091565b631b53398f60e21b90565b604080516060810182526000808252602082018190529181019190915290565b6001600160e01b031991909116815260200190565b60ff9290921682528051602080840191909152810151604080840191909152015160608201526080019056fea2646970667358221220912857c82dd00a076280548bba4d1c7684b646da08b97b4425bcfb99c1d9c16764736f6c63430006060033a2646970667358221220865ffab2b3490cb23ad7394a9ba081700df089fca047c3234aca4a6c2ac74a9164736f6c63430006060033"

// DeployActivitiesFactory deploys a new Ethereum contract, binding an instance of ActivitiesFactory to it.
func DeployActivitiesFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ActivitiesFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivitiesFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ActivitiesFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ActivitiesFactory{ActivitiesFactoryCaller: ActivitiesFactoryCaller{contract: contract}, ActivitiesFactoryTransactor: ActivitiesFactoryTransactor{contract: contract}, ActivitiesFactoryFilterer: ActivitiesFactoryFilterer{contract: contract}}, nil
}

// ActivitiesFactory is an auto generated Go binding around an Ethereum contract.
type ActivitiesFactory struct {
	ActivitiesFactoryCaller     // Read-only binding to the contract
	ActivitiesFactoryTransactor // Write-only binding to the contract
	ActivitiesFactoryFilterer   // Log filterer for contract events
}

// ActivitiesFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivitiesFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivitiesFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivitiesFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivitiesFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivitiesFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivitiesFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivitiesFactorySession struct {
	Contract     *ActivitiesFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ActivitiesFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivitiesFactoryCallerSession struct {
	Contract *ActivitiesFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ActivitiesFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivitiesFactoryTransactorSession struct {
	Contract     *ActivitiesFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ActivitiesFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivitiesFactoryRaw struct {
	Contract *ActivitiesFactory // Generic contract binding to access the raw methods on
}

// ActivitiesFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivitiesFactoryCallerRaw struct {
	Contract *ActivitiesFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ActivitiesFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivitiesFactoryTransactorRaw struct {
	Contract *ActivitiesFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivitiesFactory creates a new instance of ActivitiesFactory, bound to a specific deployed contract.
func NewActivitiesFactory(address common.Address, backend bind.ContractBackend) (*ActivitiesFactory, error) {
	contract, err := bindActivitiesFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ActivitiesFactory{ActivitiesFactoryCaller: ActivitiesFactoryCaller{contract: contract}, ActivitiesFactoryTransactor: ActivitiesFactoryTransactor{contract: contract}, ActivitiesFactoryFilterer: ActivitiesFactoryFilterer{contract: contract}}, nil
}

// NewActivitiesFactoryCaller creates a new read-only instance of ActivitiesFactory, bound to a specific deployed contract.
func NewActivitiesFactoryCaller(address common.Address, caller bind.ContractCaller) (*ActivitiesFactoryCaller, error) {
	contract, err := bindActivitiesFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivitiesFactoryCaller{contract: contract}, nil
}

// NewActivitiesFactoryTransactor creates a new write-only instance of ActivitiesFactory, bound to a specific deployed contract.
func NewActivitiesFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivitiesFactoryTransactor, error) {
	contract, err := bindActivitiesFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivitiesFactoryTransactor{contract: contract}, nil
}

// NewActivitiesFactoryFilterer creates a new log filterer instance of ActivitiesFactory, bound to a specific deployed contract.
func NewActivitiesFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivitiesFactoryFilterer, error) {
	contract, err := bindActivitiesFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivitiesFactoryFilterer{contract: contract}, nil
}

// bindActivitiesFactory binds a generic wrapper to an already deployed contract.
func bindActivitiesFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivitiesFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivitiesFactory *ActivitiesFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivitiesFactory.Contract.ActivitiesFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivitiesFactory *ActivitiesFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivitiesFactory.Contract.ActivitiesFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivitiesFactory *ActivitiesFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivitiesFactory.Contract.ActivitiesFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivitiesFactory *ActivitiesFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivitiesFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivitiesFactory *ActivitiesFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivitiesFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivitiesFactory *ActivitiesFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivitiesFactory.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _ref) view returns(address)
func (_ActivitiesFactory *ActivitiesFactoryCaller) Get(opts *bind.CallOpts, _ref [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ActivitiesFactory.contract.Call(opts, out, "get", _ref)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _ref) view returns(address)
func (_ActivitiesFactory *ActivitiesFactorySession) Get(_ref [32]byte) (common.Address, error) {
	return _ActivitiesFactory.Contract.Get(&_ActivitiesFactory.CallOpts, _ref)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _ref) view returns(address)
func (_ActivitiesFactory *ActivitiesFactoryCallerSession) Get(_ref [32]byte) (common.Address, error) {
	return _ActivitiesFactory.Contract.Get(&_ActivitiesFactory.CallOpts, _ref)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_ActivitiesFactory *ActivitiesFactoryCaller) GetNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ActivitiesFactory.contract.Call(opts, out, "getNum")
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_ActivitiesFactory *ActivitiesFactorySession) GetNum() (*big.Int, error) {
	return _ActivitiesFactory.Contract.GetNum(&_ActivitiesFactory.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_ActivitiesFactory *ActivitiesFactoryCallerSession) GetNum() (*big.Int, error) {
	return _ActivitiesFactory.Contract.GetNum(&_ActivitiesFactory.CallOpts)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_ActivitiesFactory *ActivitiesFactoryCaller) GetReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivitiesFactory.contract.Call(opts, out, "getReference", _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_ActivitiesFactory *ActivitiesFactorySession) GetReference(_index *big.Int) ([32]byte, error) {
	return _ActivitiesFactory.Contract.GetReference(&_ActivitiesFactory.CallOpts, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_ActivitiesFactory *ActivitiesFactoryCallerSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _ActivitiesFactory.Contract.GetReference(&_ActivitiesFactory.CallOpts, _index)
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(bytes4)
func (_ActivitiesFactory *ActivitiesFactoryCaller) Setter(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _ActivitiesFactory.contract.Call(opts, out, "setter")
	return *ret0, err
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(bytes4)
func (_ActivitiesFactory *ActivitiesFactorySession) Setter() ([4]byte, error) {
	return _ActivitiesFactory.Contract.Setter(&_ActivitiesFactory.CallOpts)
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(bytes4)
func (_ActivitiesFactory *ActivitiesFactoryCallerSession) Setter() ([4]byte, error) {
	return _ActivitiesFactory.Contract.Setter(&_ActivitiesFactory.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x91a95fbc.
//
// Solidity: function set(bytes32 _ref, ActivitiesData _activities) returns()
func (_ActivitiesFactory *ActivitiesFactoryTransactor) Set(opts *bind.TransactOpts, _ref [32]byte, _activities ActivitiesData) (*types.Transaction, error) {
	return _ActivitiesFactory.contract.Transact(opts, "set", _ref, _activities)
}

// Set is a paid mutator transaction binding the contract method 0x91a95fbc.
//
// Solidity: function set(bytes32 _ref, ActivitiesData _activities) returns()
func (_ActivitiesFactory *ActivitiesFactorySession) Set(_ref [32]byte, _activities ActivitiesData) (*types.Transaction, error) {
	return _ActivitiesFactory.Contract.Set(&_ActivitiesFactory.TransactOpts, _ref, _activities)
}

// Set is a paid mutator transaction binding the contract method 0x91a95fbc.
//
// Solidity: function set(bytes32 _ref, ActivitiesData _activities) returns()
func (_ActivitiesFactory *ActivitiesFactoryTransactorSession) Set(_ref [32]byte, _activities ActivitiesData) (*types.Transaction, error) {
	return _ActivitiesFactory.Contract.Set(&_ActivitiesFactory.TransactOpts, _ref, _activities)
}

// ActivitiesFactorySetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the ActivitiesFactory contract.
type ActivitiesFactorySetIterator struct {
	Event *ActivitiesFactorySet // Event containing the contract specifics and raw log

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
func (it *ActivitiesFactorySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivitiesFactorySet)
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
		it.Event = new(ActivitiesFactorySet)
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
func (it *ActivitiesFactorySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivitiesFactorySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivitiesFactorySet represents a Set event raised by the ActivitiesFactory contract.
type ActivitiesFactorySet struct {
	ElementType uint8
	Ref         [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0x8353a4d574992e9eb676a31cd326ce2f9f6c4829c3b84ae33ee9febbb2962e7b.
//
// Solidity: event Set(uint8 _elementType, bytes32 _ref)
func (_ActivitiesFactory *ActivitiesFactoryFilterer) FilterSet(opts *bind.FilterOpts) (*ActivitiesFactorySetIterator, error) {

	logs, sub, err := _ActivitiesFactory.contract.FilterLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return &ActivitiesFactorySetIterator{contract: _ActivitiesFactory.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0x8353a4d574992e9eb676a31cd326ce2f9f6c4829c3b84ae33ee9febbb2962e7b.
//
// Solidity: event Set(uint8 _elementType, bytes32 _ref)
func (_ActivitiesFactory *ActivitiesFactoryFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *ActivitiesFactorySet) (event.Subscription, error) {

	logs, sub, err := _ActivitiesFactory.contract.WatchLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivitiesFactorySet)
				if err := _ActivitiesFactory.contract.UnpackLog(event, "Set", log); err != nil {
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
func (_ActivitiesFactory *ActivitiesFactoryFilterer) ParseSet(log types.Log) (*ActivitiesFactorySet, error) {
	event := new(ActivitiesFactorySet)
	if err := _ActivitiesFactory.contract.UnpackLog(event, "Set", log); err != nil {
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
var ActivitiesNodeBin = "0x608060405234801561001057600080fd5b5060405161024b38038061024b83398101604081905261002f916100a3565b805160001a60f81b6001600160f81b031916158015906100625750602081015160001a60f81b6001600160f81b03191615155b80156100815750604081015160001a60f81b6001600160f81b03191615155b61008a57600080fd5b80516000556020810151600155604001516002556100f9565b6000606082840312156100b4578081fd5b604051606081016001600160401b03811182821017156100d2578283fd5b80604052508251815260208301516020820152604083015160408201528091505092915050565b610143806101086000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636d4ce63c1461003b578063993a04b71461005a575b600080fd5b61004361006f565b6040516100519291906100e1565b60405180910390f35b6100626100a1565b60405161005191906100cc565b60006100796100ac565b5050604080516060810182526000548152600154602082015260028054928201929092529091565b631b53398f60e21b90565b604080516060810182526000808252602082018190529181019190915290565b6001600160e01b031991909116815260200190565b60ff9290921682528051602080840191909152810151604080840191909152015160608201526080019056fea2646970667358221220912857c82dd00a076280548bba4d1c7684b646da08b97b4425bcfb99c1d9c16764736f6c63430006060033"

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

// IterableDataABI is the input ABI used to generate the binding from.
const IterableDataABI = "[]"

// IterableDataBin is the compiled bytecode used for deploying new contracts.
var IterableDataBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e52e8c30ae389f681d7d596e03e2846b093c9e36b778b6817bcbce4e0903c46564736f6c63430006060033"

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
