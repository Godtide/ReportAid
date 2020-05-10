// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package activityAdditional

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

// AdditionalData is an auto generated low-level Go binding around an user-defined struct.
type AdditionalData struct {
	Scope              uint8
	CapitalSpend       uint8
	CollaborationType  uint8
	DefaultFlowType    uint8
	DefaultTiedStatus  uint8
	DefaultAidType     [32]byte
	DefaultFinanceType *big.Int
}

// ActivityAdditionalABI is the input ABI used to generate the binding from.
const ActivityAdditionalABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_activityCon\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_elementType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"Set\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parentRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_thisRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"scope\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"capitalSpend\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"collaborationType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"defaultFlowType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"defaultTiedStatus\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"defaultAidType\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"defaultFinanceType\",\"type\":\"uint256\"}],\"internalType\":\"structAdditionalData\",\"name\":\"_additional\",\"type\":\"tuple\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setter\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ActivityAdditionalFuncSigs maps the 4-byte function signature to its string representation.
var ActivityAdditionalFuncSigs = map[string]string{
	"8eaa6ac0": "get(bytes32)",
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
	"d9a958cf": "set(bytes32,bytes32,(uint8,uint8,uint8,uint8,uint8,bytes32,uint256))",
	"3f3108f7": "setter()",
}

// ActivityAdditionalBin is the compiled bytecode used for deploying new contracts.
var ActivityAdditionalBin = "0x608060405234801561001057600080fd5b50604051610a07380380610a0783398101604081905261002f91610069565b6000546001600160a01b031661004157fe5b600080546001600160a01b0319166001600160a01b0392909216919091178155600355610097565b60006020828403121561007a578081fd5b81516001600160a01b0381168114610090578182fd5b9392505050565b610961806100a66000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80633f3108f71461005c57806367e0badb1461007a5780638eaa6ac01461008f578063bc599341146100af578063d9a958cf146100c2575b600080fd5b6100646100d7565b604051610071919061048a565b60405180910390f35b6100826100e2565b604051610071919061045f565b6100a261009d366004610367565b6100e8565b604051610071919061044b565b6100826100bd366004610367565b610148565b6100d56100d036600461037f565b610180565b005b63d9a958cf60e01b90565b60035490565b600081811a60f81b6001600160f81b03191661010357600080fd5b600082815260016020819052604090912001546001600160a01b031661012857600080fd5b50600090815260016020819052604090912001546001600160a01b031690565b600354600090821061015957600080fd5b600280548390811061016757fe5b9060005260206000209060020201600001549050919050565b8260001a60f81b6001600160f81b03191661019a57600080fd5b8160001a60f81b6001600160f81b0319166101b457600080fd5b6000816040516101c390610343565b6101cd919061049f565b604051809103906000f0801580156101e9573d6000803e3d6000fd5b5090506101fe6001848363ffffffff6102ac16565b506000546001600160a01b031663cbb2dc95858360046040518463ffffffff1660e01b815260040161023293929190610468565b600060405180830381600087803b15801561024c57600080fd5b505af1158015610260573d6000803e3d6000fd5b507f8353a4d574992e9eb676a31cd326ce2f9f6c4829c3b84ae33ee9febbb2962e7b92506004915061028f9050565b8460405161029e9291906104fc565b60405180910390a150505050565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b03851617905580156102ea57600191505061033c565b506001808501805491820180825560008681526020889052604090208190558591908390811061031657fe5b6000918252602082206002918202019290925590860180546001019055915061033c9050565b9392505050565b61041c8061051083390190565b803560ff8116811461036157600080fd5b92915050565b600060208284031215610378578081fd5b5035919050565b6000806000838503610120811215610395578283fd5b843593506020850135925060e0603f19820112156103b1578182fd5b5060405160e0810181811067ffffffffffffffff821117156103d1578283fd5b80604052506103e38660408701610350565b81526103f28660608701610350565b60208201526104048660808701610350565b60408201526104168660a08701610350565b60608201526104288660c08701610350565b608082015260e085013560a08201526101009094013560c0850152509093909250565b6001600160a01b0391909116815260200190565b90815260200190565b9283526001600160a01b0391909116602083015260ff16604082015260600190565b6001600160e01b031991909116815260200190565b600060e08201905060ff835116825260ff602084015116602083015260ff604084015116604083015260ff606084015116606083015260ff608084015116608083015260a083015160a083015260c083015160c083015292915050565b60ff92909216825260208201526040019056fe608060405234801561001057600080fd5b5060405161041c38038061041c83398101604081905261002f91610195565b805160ff161580159061004857508051600960ff909116105b801561005c57506064816020015160ff1611155b801561006e5750604081015160ff1615155b80156100845750600860ff16816040015160ff16105b80156100bb575060a081015160001a60f81b7fff000000000000000000000000000000000000000000000000000000000000001615155b80156100d15750600260ff16816080015160ff16115b80156100e75750600660ff16816080015160ff16105b6100f057600080fd5b805160008054602084015160408501516080860151606087015160ff90811663010000000263ff000000199282166401000000000260ff60201b19948316620100000262ff0000199684166101000261ff001994909a1660ff19909816979097179290921697909717939093169390931716171691909117905560a081015160015560c00151600255610239565b805160ff8116811461018f57600080fd5b92915050565b600060e082840312156101a6578081fd5b60405160e081016001600160401b03811182821017156101c4578283fd5b6040526101d1848461017e565b81526101e0846020850161017e565b60208201526101f2846040850161017e565b6040820152610204846060850161017e565b6060820152610216846080850161017e565b608082015260a083015160a082015260c083015160c08201528091505092915050565b6101d4806102486000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636d4ce63c1461003b578063993a04b71461005a575b600080fd5b61004361006f565b604051610051929190610139565b60405180910390f35b6100626100dd565b6040516100519190610124565b60006100796100e8565b50506040805160e08101825260005460ff808216835261010082048116602084015262010000820481169383019390935263010000008104831660608301526401000000009004909116608082015260015460a082015260025460c0820152600491565b631b53398f60e21b90565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c081019190915290565b6001600160e01b031991909116815260200190565b60006101008201905060ff80851683528084511660208401528060208501511660408401528060408501511660608401528060608501511660808401528060808501511660a08401525060a083015160c083015260c083015160e0830152939250505056fea2646970667358221220a10320a4c36fe579f662f08be7c65f50b1a81fcf4805a1d8cc9a0203f96ea4ae64736f6c63430006060033a2646970667358221220e57d83e5aeb26fd3dd4e8587690f9a5131f3ea9090134a8349989a60df7e5d8b64736f6c63430006060033"

// DeployActivityAdditional deploys a new Ethereum contract, binding an instance of ActivityAdditional to it.
func DeployActivityAdditional(auth *bind.TransactOpts, backend bind.ContractBackend, _activityCon common.Address) (common.Address, *types.Transaction, *ActivityAdditional, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivityAdditionalABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ActivityAdditionalBin), backend, _activityCon)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ActivityAdditional{ActivityAdditionalCaller: ActivityAdditionalCaller{contract: contract}, ActivityAdditionalTransactor: ActivityAdditionalTransactor{contract: contract}, ActivityAdditionalFilterer: ActivityAdditionalFilterer{contract: contract}}, nil
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

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _ref) view returns(address)
func (_ActivityAdditional *ActivityAdditionalCaller) Get(opts *bind.CallOpts, _ref [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ActivityAdditional.contract.Call(opts, out, "get", _ref)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _ref) view returns(address)
func (_ActivityAdditional *ActivityAdditionalSession) Get(_ref [32]byte) (common.Address, error) {
	return _ActivityAdditional.Contract.Get(&_ActivityAdditional.CallOpts, _ref)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _ref) view returns(address)
func (_ActivityAdditional *ActivityAdditionalCallerSession) Get(_ref [32]byte) (common.Address, error) {
	return _ActivityAdditional.Contract.Get(&_ActivityAdditional.CallOpts, _ref)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_ActivityAdditional *ActivityAdditionalCaller) GetNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ActivityAdditional.contract.Call(opts, out, "getNum")
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_ActivityAdditional *ActivityAdditionalSession) GetNum() (*big.Int, error) {
	return _ActivityAdditional.Contract.GetNum(&_ActivityAdditional.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_ActivityAdditional *ActivityAdditionalCallerSession) GetNum() (*big.Int, error) {
	return _ActivityAdditional.Contract.GetNum(&_ActivityAdditional.CallOpts)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_ActivityAdditional *ActivityAdditionalCaller) GetReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityAdditional.contract.Call(opts, out, "getReference", _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_ActivityAdditional *ActivityAdditionalSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _ActivityAdditional.Contract.GetReference(&_ActivityAdditional.CallOpts, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_ActivityAdditional *ActivityAdditionalCallerSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _ActivityAdditional.Contract.GetReference(&_ActivityAdditional.CallOpts, _index)
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(bytes4)
func (_ActivityAdditional *ActivityAdditionalCaller) Setter(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _ActivityAdditional.contract.Call(opts, out, "setter")
	return *ret0, err
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(bytes4)
func (_ActivityAdditional *ActivityAdditionalSession) Setter() ([4]byte, error) {
	return _ActivityAdditional.Contract.Setter(&_ActivityAdditional.CallOpts)
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(bytes4)
func (_ActivityAdditional *ActivityAdditionalCallerSession) Setter() ([4]byte, error) {
	return _ActivityAdditional.Contract.Setter(&_ActivityAdditional.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0xd9a958cf.
//
// Solidity: function set(bytes32 _parentRef, bytes32 _thisRef, AdditionalData _additional) returns()
func (_ActivityAdditional *ActivityAdditionalTransactor) Set(opts *bind.TransactOpts, _parentRef [32]byte, _thisRef [32]byte, _additional AdditionalData) (*types.Transaction, error) {
	return _ActivityAdditional.contract.Transact(opts, "set", _parentRef, _thisRef, _additional)
}

// Set is a paid mutator transaction binding the contract method 0xd9a958cf.
//
// Solidity: function set(bytes32 _parentRef, bytes32 _thisRef, AdditionalData _additional) returns()
func (_ActivityAdditional *ActivityAdditionalSession) Set(_parentRef [32]byte, _thisRef [32]byte, _additional AdditionalData) (*types.Transaction, error) {
	return _ActivityAdditional.Contract.Set(&_ActivityAdditional.TransactOpts, _parentRef, _thisRef, _additional)
}

// Set is a paid mutator transaction binding the contract method 0xd9a958cf.
//
// Solidity: function set(bytes32 _parentRef, bytes32 _thisRef, AdditionalData _additional) returns()
func (_ActivityAdditional *ActivityAdditionalTransactorSession) Set(_parentRef [32]byte, _thisRef [32]byte, _additional AdditionalData) (*types.Transaction, error) {
	return _ActivityAdditional.Contract.Set(&_ActivityAdditional.TransactOpts, _parentRef, _thisRef, _additional)
}

// ActivityAdditionalSetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the ActivityAdditional contract.
type ActivityAdditionalSetIterator struct {
	Event *ActivityAdditionalSet // Event containing the contract specifics and raw log

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
func (it *ActivityAdditionalSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivityAdditionalSet)
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
		it.Event = new(ActivityAdditionalSet)
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
func (it *ActivityAdditionalSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivityAdditionalSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivityAdditionalSet represents a Set event raised by the ActivityAdditional contract.
type ActivityAdditionalSet struct {
	ElementType uint8
	Ref         [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0x8353a4d574992e9eb676a31cd326ce2f9f6c4829c3b84ae33ee9febbb2962e7b.
//
// Solidity: event Set(uint8 _elementType, bytes32 _ref)
func (_ActivityAdditional *ActivityAdditionalFilterer) FilterSet(opts *bind.FilterOpts) (*ActivityAdditionalSetIterator, error) {

	logs, sub, err := _ActivityAdditional.contract.FilterLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return &ActivityAdditionalSetIterator{contract: _ActivityAdditional.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0x8353a4d574992e9eb676a31cd326ce2f9f6c4829c3b84ae33ee9febbb2962e7b.
//
// Solidity: event Set(uint8 _elementType, bytes32 _ref)
func (_ActivityAdditional *ActivityAdditionalFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *ActivityAdditionalSet) (event.Subscription, error) {

	logs, sub, err := _ActivityAdditional.contract.WatchLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivityAdditionalSet)
				if err := _ActivityAdditional.contract.UnpackLog(event, "Set", log); err != nil {
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
func (_ActivityAdditional *ActivityAdditionalFilterer) ParseSet(log types.Log) (*ActivityAdditionalSet, error) {
	event := new(ActivityAdditionalSet)
	if err := _ActivityAdditional.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ActivityAdditionalNodeABI is the input ABI used to generate the binding from.
const ActivityAdditionalNodeABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"scope\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"capitalSpend\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"collaborationType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"defaultFlowType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"defaultTiedStatus\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"defaultAidType\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"defaultFinanceType\",\"type\":\"uint256\"}],\"internalType\":\"structAdditionalData\",\"name\":\"_additional\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"scope\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"capitalSpend\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"collaborationType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"defaultFlowType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"defaultTiedStatus\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"defaultAidType\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"defaultFinanceType\",\"type\":\"uint256\"}],\"internalType\":\"structAdditionalData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getter\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ActivityAdditionalNodeFuncSigs maps the 4-byte function signature to its string representation.
var ActivityAdditionalNodeFuncSigs = map[string]string{
	"6d4ce63c": "get()",
	"993a04b7": "getter()",
}

// ActivityAdditionalNodeBin is the compiled bytecode used for deploying new contracts.
var ActivityAdditionalNodeBin = "0x608060405234801561001057600080fd5b5060405161041c38038061041c83398101604081905261002f91610195565b805160ff161580159061004857508051600960ff909116105b801561005c57506064816020015160ff1611155b801561006e5750604081015160ff1615155b80156100845750600860ff16816040015160ff16105b80156100bb575060a081015160001a60f81b7fff000000000000000000000000000000000000000000000000000000000000001615155b80156100d15750600260ff16816080015160ff16115b80156100e75750600660ff16816080015160ff16105b6100f057600080fd5b805160008054602084015160408501516080860151606087015160ff90811663010000000263ff000000199282166401000000000260ff60201b19948316620100000262ff0000199684166101000261ff001994909a1660ff19909816979097179290921697909717939093169390931716171691909117905560a081015160015560c00151600255610239565b805160ff8116811461018f57600080fd5b92915050565b600060e082840312156101a6578081fd5b60405160e081016001600160401b03811182821017156101c4578283fd5b6040526101d1848461017e565b81526101e0846020850161017e565b60208201526101f2846040850161017e565b6040820152610204846060850161017e565b6060820152610216846080850161017e565b608082015260a083015160a082015260c083015160c08201528091505092915050565b6101d4806102486000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636d4ce63c1461003b578063993a04b71461005a575b600080fd5b61004361006f565b604051610051929190610139565b60405180910390f35b6100626100dd565b6040516100519190610124565b60006100796100e8565b50506040805160e08101825260005460ff808216835261010082048116602084015262010000820481169383019390935263010000008104831660608301526401000000009004909116608082015260015460a082015260025460c0820152600491565b631b53398f60e21b90565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c081019190915290565b6001600160e01b031991909116815260200190565b60006101008201905060ff80851683528084511660208401528060208501511660408401528060408501511660608401528060608501511660808401528060808501511660a08401525060a083015160c083015260c083015160e0830152939250505056fea2646970667358221220a10320a4c36fe579f662f08be7c65f50b1a81fcf4805a1d8cc9a0203f96ea4ae64736f6c63430006060033"

// DeployActivityAdditionalNode deploys a new Ethereum contract, binding an instance of ActivityAdditionalNode to it.
func DeployActivityAdditionalNode(auth *bind.TransactOpts, backend bind.ContractBackend, _additional AdditionalData) (common.Address, *types.Transaction, *ActivityAdditionalNode, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivityAdditionalNodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ActivityAdditionalNodeBin), backend, _additional)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ActivityAdditionalNode{ActivityAdditionalNodeCaller: ActivityAdditionalNodeCaller{contract: contract}, ActivityAdditionalNodeTransactor: ActivityAdditionalNodeTransactor{contract: contract}, ActivityAdditionalNodeFilterer: ActivityAdditionalNodeFilterer{contract: contract}}, nil
}

// ActivityAdditionalNode is an auto generated Go binding around an Ethereum contract.
type ActivityAdditionalNode struct {
	ActivityAdditionalNodeCaller     // Read-only binding to the contract
	ActivityAdditionalNodeTransactor // Write-only binding to the contract
	ActivityAdditionalNodeFilterer   // Log filterer for contract events
}

// ActivityAdditionalNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivityAdditionalNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityAdditionalNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivityAdditionalNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityAdditionalNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivityAdditionalNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityAdditionalNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivityAdditionalNodeSession struct {
	Contract     *ActivityAdditionalNode // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ActivityAdditionalNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivityAdditionalNodeCallerSession struct {
	Contract *ActivityAdditionalNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ActivityAdditionalNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivityAdditionalNodeTransactorSession struct {
	Contract     *ActivityAdditionalNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ActivityAdditionalNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivityAdditionalNodeRaw struct {
	Contract *ActivityAdditionalNode // Generic contract binding to access the raw methods on
}

// ActivityAdditionalNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivityAdditionalNodeCallerRaw struct {
	Contract *ActivityAdditionalNodeCaller // Generic read-only contract binding to access the raw methods on
}

// ActivityAdditionalNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivityAdditionalNodeTransactorRaw struct {
	Contract *ActivityAdditionalNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivityAdditionalNode creates a new instance of ActivityAdditionalNode, bound to a specific deployed contract.
func NewActivityAdditionalNode(address common.Address, backend bind.ContractBackend) (*ActivityAdditionalNode, error) {
	contract, err := bindActivityAdditionalNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ActivityAdditionalNode{ActivityAdditionalNodeCaller: ActivityAdditionalNodeCaller{contract: contract}, ActivityAdditionalNodeTransactor: ActivityAdditionalNodeTransactor{contract: contract}, ActivityAdditionalNodeFilterer: ActivityAdditionalNodeFilterer{contract: contract}}, nil
}

// NewActivityAdditionalNodeCaller creates a new read-only instance of ActivityAdditionalNode, bound to a specific deployed contract.
func NewActivityAdditionalNodeCaller(address common.Address, caller bind.ContractCaller) (*ActivityAdditionalNodeCaller, error) {
	contract, err := bindActivityAdditionalNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityAdditionalNodeCaller{contract: contract}, nil
}

// NewActivityAdditionalNodeTransactor creates a new write-only instance of ActivityAdditionalNode, bound to a specific deployed contract.
func NewActivityAdditionalNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivityAdditionalNodeTransactor, error) {
	contract, err := bindActivityAdditionalNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityAdditionalNodeTransactor{contract: contract}, nil
}

// NewActivityAdditionalNodeFilterer creates a new log filterer instance of ActivityAdditionalNode, bound to a specific deployed contract.
func NewActivityAdditionalNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivityAdditionalNodeFilterer, error) {
	contract, err := bindActivityAdditionalNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivityAdditionalNodeFilterer{contract: contract}, nil
}

// bindActivityAdditionalNode binds a generic wrapper to an already deployed contract.
func bindActivityAdditionalNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivityAdditionalNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityAdditionalNode *ActivityAdditionalNodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivityAdditionalNode.Contract.ActivityAdditionalNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityAdditionalNode *ActivityAdditionalNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityAdditionalNode.Contract.ActivityAdditionalNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityAdditionalNode *ActivityAdditionalNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityAdditionalNode.Contract.ActivityAdditionalNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityAdditionalNode *ActivityAdditionalNodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivityAdditionalNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityAdditionalNode *ActivityAdditionalNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityAdditionalNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityAdditionalNode *ActivityAdditionalNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityAdditionalNode.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint8, AdditionalData)
func (_ActivityAdditionalNode *ActivityAdditionalNodeCaller) Get(opts *bind.CallOpts) (uint8, AdditionalData, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(AdditionalData)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ActivityAdditionalNode.contract.Call(opts, out, "get")
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint8, AdditionalData)
func (_ActivityAdditionalNode *ActivityAdditionalNodeSession) Get() (uint8, AdditionalData, error) {
	return _ActivityAdditionalNode.Contract.Get(&_ActivityAdditionalNode.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint8, AdditionalData)
func (_ActivityAdditionalNode *ActivityAdditionalNodeCallerSession) Get() (uint8, AdditionalData, error) {
	return _ActivityAdditionalNode.Contract.Get(&_ActivityAdditionalNode.CallOpts)
}

// Getter is a free data retrieval call binding the contract method 0x993a04b7.
//
// Solidity: function getter() view returns(bytes4)
func (_ActivityAdditionalNode *ActivityAdditionalNodeCaller) Getter(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _ActivityAdditionalNode.contract.Call(opts, out, "getter")
	return *ret0, err
}

// Getter is a free data retrieval call binding the contract method 0x993a04b7.
//
// Solidity: function getter() view returns(bytes4)
func (_ActivityAdditionalNode *ActivityAdditionalNodeSession) Getter() ([4]byte, error) {
	return _ActivityAdditionalNode.Contract.Getter(&_ActivityAdditionalNode.CallOpts)
}

// Getter is a free data retrieval call binding the contract method 0x993a04b7.
//
// Solidity: function getter() view returns(bytes4)
func (_ActivityAdditionalNode *ActivityAdditionalNodeCallerSession) Getter() ([4]byte, error) {
	return _ActivityAdditionalNode.Contract.Getter(&_ActivityAdditionalNode.CallOpts)
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
