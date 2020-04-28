// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package activityParticipatingOrgs

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

// ActivityParticipatingOrgsParticipatingOrg is an auto generated low-level Go binding around an user-defined struct.
type ActivityParticipatingOrgsParticipatingOrg struct {
	OrgType        uint8
	Role           uint8
	CrsChannelCode *big.Int
	Lang           [32]byte
	OrgRef         [32]byte
	Narrative      string
}

// ActivityParticipatingOrgsABI is the input ABI used to generate the binding from.
const ActivityParticipatingOrgsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_particpatingOrgRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"crsChannelCode\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"narrative\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structActivityParticipatingOrgs.ParticipatingOrg\",\"name\":\"_participatingOrg\",\"type\":\"tuple\"}],\"name\":\"SetParticipatingOrg\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_particpatingOrgRef\",\"type\":\"bytes32\"}],\"name\":\"getCrsChannelCode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_particpatingOrgRef\",\"type\":\"bytes32\"}],\"name\":\"getLang\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_particpatingOrgRef\",\"type\":\"bytes32\"}],\"name\":\"getNarrative\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getNumParticipatingOrgs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_particpatingOrgRef\",\"type\":\"bytes32\"}],\"name\":\"getOrgRef\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_particpatingOrgRef\",\"type\":\"bytes32\"}],\"name\":\"getParticipatingOrg\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"crsChannelCode\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"narrative\",\"type\":\"string\"}],\"internalType\":\"structActivityParticipatingOrgs.ParticipatingOrg\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_particpatingOrgRef\",\"type\":\"bytes32\"}],\"name\":\"getRole\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_particpatingOrgRef\",\"type\":\"bytes32\"}],\"name\":\"getType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_particpatingOrgRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"crsChannelCode\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"narrative\",\"type\":\"string\"}],\"internalType\":\"structActivityParticipatingOrgs.ParticipatingOrg\",\"name\":\"_participatingOrg\",\"type\":\"tuple\"}],\"name\":\"setParticipatingOrg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ActivityParticipatingOrgsFuncSigs maps the 4-byte function signature to its string representation.
var ActivityParticipatingOrgsFuncSigs = map[string]string{
	"0dbc539b": "getCrsChannelCode(bytes32,bytes32,bytes32)",
	"e022cb52": "getLang(bytes32,bytes32,bytes32)",
	"ae13bb7a": "getNarrative(bytes32,bytes32,bytes32)",
	"ca39101b": "getNumParticipatingOrgs(bytes32,bytes32)",
	"d4976bce": "getOrgRef(bytes32,bytes32,bytes32)",
	"4017dc0d": "getParticipatingOrg(bytes32,bytes32,bytes32)",
	"eef67d78": "getReference(bytes32,bytes32,uint256)",
	"57db6f69": "getRole(bytes32,bytes32,bytes32)",
	"b153af81": "getType(bytes32,bytes32,bytes32)",
	"a444eeda": "setParticipatingOrg(bytes32,bytes32,bytes32,(uint8,uint8,uint256,bytes32,bytes32,string))",
}

// ActivityParticipatingOrgs is an auto generated Go binding around an Ethereum contract.
type ActivityParticipatingOrgs struct {
	ActivityParticipatingOrgsCaller     // Read-only binding to the contract
	ActivityParticipatingOrgsTransactor // Write-only binding to the contract
	ActivityParticipatingOrgsFilterer   // Log filterer for contract events
}

// ActivityParticipatingOrgsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivityParticipatingOrgsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityParticipatingOrgsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivityParticipatingOrgsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityParticipatingOrgsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivityParticipatingOrgsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityParticipatingOrgsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivityParticipatingOrgsSession struct {
	Contract     *ActivityParticipatingOrgs // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ActivityParticipatingOrgsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivityParticipatingOrgsCallerSession struct {
	Contract *ActivityParticipatingOrgsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ActivityParticipatingOrgsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivityParticipatingOrgsTransactorSession struct {
	Contract     *ActivityParticipatingOrgsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ActivityParticipatingOrgsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivityParticipatingOrgsRaw struct {
	Contract *ActivityParticipatingOrgs // Generic contract binding to access the raw methods on
}

// ActivityParticipatingOrgsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivityParticipatingOrgsCallerRaw struct {
	Contract *ActivityParticipatingOrgsCaller // Generic read-only contract binding to access the raw methods on
}

// ActivityParticipatingOrgsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivityParticipatingOrgsTransactorRaw struct {
	Contract *ActivityParticipatingOrgsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivityParticipatingOrgs creates a new instance of ActivityParticipatingOrgs, bound to a specific deployed contract.
func NewActivityParticipatingOrgs(address common.Address, backend bind.ContractBackend) (*ActivityParticipatingOrgs, error) {
	contract, err := bindActivityParticipatingOrgs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ActivityParticipatingOrgs{ActivityParticipatingOrgsCaller: ActivityParticipatingOrgsCaller{contract: contract}, ActivityParticipatingOrgsTransactor: ActivityParticipatingOrgsTransactor{contract: contract}, ActivityParticipatingOrgsFilterer: ActivityParticipatingOrgsFilterer{contract: contract}}, nil
}

// NewActivityParticipatingOrgsCaller creates a new read-only instance of ActivityParticipatingOrgs, bound to a specific deployed contract.
func NewActivityParticipatingOrgsCaller(address common.Address, caller bind.ContractCaller) (*ActivityParticipatingOrgsCaller, error) {
	contract, err := bindActivityParticipatingOrgs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityParticipatingOrgsCaller{contract: contract}, nil
}

// NewActivityParticipatingOrgsTransactor creates a new write-only instance of ActivityParticipatingOrgs, bound to a specific deployed contract.
func NewActivityParticipatingOrgsTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivityParticipatingOrgsTransactor, error) {
	contract, err := bindActivityParticipatingOrgs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityParticipatingOrgsTransactor{contract: contract}, nil
}

// NewActivityParticipatingOrgsFilterer creates a new log filterer instance of ActivityParticipatingOrgs, bound to a specific deployed contract.
func NewActivityParticipatingOrgsFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivityParticipatingOrgsFilterer, error) {
	contract, err := bindActivityParticipatingOrgs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivityParticipatingOrgsFilterer{contract: contract}, nil
}

// bindActivityParticipatingOrgs binds a generic wrapper to an already deployed contract.
func bindActivityParticipatingOrgs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivityParticipatingOrgsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivityParticipatingOrgs.Contract.ActivityParticipatingOrgsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityParticipatingOrgs.Contract.ActivityParticipatingOrgsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityParticipatingOrgs.Contract.ActivityParticipatingOrgsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivityParticipatingOrgs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityParticipatingOrgs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityParticipatingOrgs.Contract.contract.Transact(opts, method, params...)
}

// GetCrsChannelCode is a free data retrieval call binding the contract method 0x0dbc539b.
//
// Solidity: function getCrsChannelCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(uint256)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCaller) GetCrsChannelCode(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ActivityParticipatingOrgs.contract.Call(opts, out, "getCrsChannelCode", _activitiesRef, _activityRef, _particpatingOrgRef)
	return *ret0, err
}

// GetCrsChannelCode is a free data retrieval call binding the contract method 0x0dbc539b.
//
// Solidity: function getCrsChannelCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(uint256)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsSession) GetCrsChannelCode(_activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) (*big.Int, error) {
	return _ActivityParticipatingOrgs.Contract.GetCrsChannelCode(&_ActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _particpatingOrgRef)
}

// GetCrsChannelCode is a free data retrieval call binding the contract method 0x0dbc539b.
//
// Solidity: function getCrsChannelCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(uint256)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCallerSession) GetCrsChannelCode(_activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) (*big.Int, error) {
	return _ActivityParticipatingOrgs.Contract.GetCrsChannelCode(&_ActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _particpatingOrgRef)
}

// GetLang is a free data retrieval call binding the contract method 0xe022cb52.
//
// Solidity: function getLang(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(bytes32)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCaller) GetLang(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityParticipatingOrgs.contract.Call(opts, out, "getLang", _activitiesRef, _activityRef, _particpatingOrgRef)
	return *ret0, err
}

// GetLang is a free data retrieval call binding the contract method 0xe022cb52.
//
// Solidity: function getLang(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(bytes32)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsSession) GetLang(_activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) ([32]byte, error) {
	return _ActivityParticipatingOrgs.Contract.GetLang(&_ActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _particpatingOrgRef)
}

// GetLang is a free data retrieval call binding the contract method 0xe022cb52.
//
// Solidity: function getLang(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(bytes32)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCallerSession) GetLang(_activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) ([32]byte, error) {
	return _ActivityParticipatingOrgs.Contract.GetLang(&_ActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _particpatingOrgRef)
}

// GetNarrative is a free data retrieval call binding the contract method 0xae13bb7a.
//
// Solidity: function getNarrative(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(string)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCaller) GetNarrative(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ActivityParticipatingOrgs.contract.Call(opts, out, "getNarrative", _activitiesRef, _activityRef, _particpatingOrgRef)
	return *ret0, err
}

// GetNarrative is a free data retrieval call binding the contract method 0xae13bb7a.
//
// Solidity: function getNarrative(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(string)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsSession) GetNarrative(_activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) (string, error) {
	return _ActivityParticipatingOrgs.Contract.GetNarrative(&_ActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _particpatingOrgRef)
}

// GetNarrative is a free data retrieval call binding the contract method 0xae13bb7a.
//
// Solidity: function getNarrative(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(string)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCallerSession) GetNarrative(_activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) (string, error) {
	return _ActivityParticipatingOrgs.Contract.GetNarrative(&_ActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _particpatingOrgRef)
}

// GetNumParticipatingOrgs is a free data retrieval call binding the contract method 0xca39101b.
//
// Solidity: function getNumParticipatingOrgs(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCaller) GetNumParticipatingOrgs(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ActivityParticipatingOrgs.contract.Call(opts, out, "getNumParticipatingOrgs", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetNumParticipatingOrgs is a free data retrieval call binding the contract method 0xca39101b.
//
// Solidity: function getNumParticipatingOrgs(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsSession) GetNumParticipatingOrgs(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _ActivityParticipatingOrgs.Contract.GetNumParticipatingOrgs(&_ActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef)
}

// GetNumParticipatingOrgs is a free data retrieval call binding the contract method 0xca39101b.
//
// Solidity: function getNumParticipatingOrgs(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCallerSession) GetNumParticipatingOrgs(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _ActivityParticipatingOrgs.Contract.GetNumParticipatingOrgs(&_ActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef)
}

// GetOrgRef is a free data retrieval call binding the contract method 0xd4976bce.
//
// Solidity: function getOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(bytes32)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCaller) GetOrgRef(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityParticipatingOrgs.contract.Call(opts, out, "getOrgRef", _activitiesRef, _activityRef, _particpatingOrgRef)
	return *ret0, err
}

// GetOrgRef is a free data retrieval call binding the contract method 0xd4976bce.
//
// Solidity: function getOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(bytes32)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsSession) GetOrgRef(_activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) ([32]byte, error) {
	return _ActivityParticipatingOrgs.Contract.GetOrgRef(&_ActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _particpatingOrgRef)
}

// GetOrgRef is a free data retrieval call binding the contract method 0xd4976bce.
//
// Solidity: function getOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(bytes32)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCallerSession) GetOrgRef(_activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) ([32]byte, error) {
	return _ActivityParticipatingOrgs.Contract.GetOrgRef(&_ActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _particpatingOrgRef)
}

// GetParticipatingOrg is a free data retrieval call binding the contract method 0x4017dc0d.
//
// Solidity: function getParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(ActivityParticipatingOrgsParticipatingOrg)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCaller) GetParticipatingOrg(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) (ActivityParticipatingOrgsParticipatingOrg, error) {
	var (
		ret0 = new(ActivityParticipatingOrgsParticipatingOrg)
	)
	out := ret0
	err := _ActivityParticipatingOrgs.contract.Call(opts, out, "getParticipatingOrg", _activitiesRef, _activityRef, _particpatingOrgRef)
	return *ret0, err
}

// GetParticipatingOrg is a free data retrieval call binding the contract method 0x4017dc0d.
//
// Solidity: function getParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(ActivityParticipatingOrgsParticipatingOrg)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsSession) GetParticipatingOrg(_activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) (ActivityParticipatingOrgsParticipatingOrg, error) {
	return _ActivityParticipatingOrgs.Contract.GetParticipatingOrg(&_ActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _particpatingOrgRef)
}

// GetParticipatingOrg is a free data retrieval call binding the contract method 0x4017dc0d.
//
// Solidity: function getParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(ActivityParticipatingOrgsParticipatingOrg)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCallerSession) GetParticipatingOrg(_activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) (ActivityParticipatingOrgsParticipatingOrg, error) {
	return _ActivityParticipatingOrgs.Contract.GetParticipatingOrg(&_ActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _particpatingOrgRef)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCaller) GetReference(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityParticipatingOrgs.contract.Call(opts, out, "getReference", _activitiesRef, _activityRef, _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _ActivityParticipatingOrgs.Contract.GetReference(&_ActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCallerSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _ActivityParticipatingOrgs.Contract.GetReference(&_ActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetRole is a free data retrieval call binding the contract method 0x57db6f69.
//
// Solidity: function getRole(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(uint8)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCaller) GetRole(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ActivityParticipatingOrgs.contract.Call(opts, out, "getRole", _activitiesRef, _activityRef, _particpatingOrgRef)
	return *ret0, err
}

// GetRole is a free data retrieval call binding the contract method 0x57db6f69.
//
// Solidity: function getRole(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(uint8)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsSession) GetRole(_activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) (uint8, error) {
	return _ActivityParticipatingOrgs.Contract.GetRole(&_ActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _particpatingOrgRef)
}

// GetRole is a free data retrieval call binding the contract method 0x57db6f69.
//
// Solidity: function getRole(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(uint8)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCallerSession) GetRole(_activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) (uint8, error) {
	return _ActivityParticipatingOrgs.Contract.GetRole(&_ActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _particpatingOrgRef)
}

// GetType is a free data retrieval call binding the contract method 0xb153af81.
//
// Solidity: function getType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(uint8)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCaller) GetType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ActivityParticipatingOrgs.contract.Call(opts, out, "getType", _activitiesRef, _activityRef, _particpatingOrgRef)
	return *ret0, err
}

// GetType is a free data retrieval call binding the contract method 0xb153af81.
//
// Solidity: function getType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(uint8)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsSession) GetType(_activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) (uint8, error) {
	return _ActivityParticipatingOrgs.Contract.GetType(&_ActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _particpatingOrgRef)
}

// GetType is a free data retrieval call binding the contract method 0xb153af81.
//
// Solidity: function getType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef) view returns(uint8)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsCallerSession) GetType(_activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte) (uint8, error) {
	return _ActivityParticipatingOrgs.Contract.GetType(&_ActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _particpatingOrgRef)
}

// SetParticipatingOrg is a paid mutator transaction binding the contract method 0xa444eeda.
//
// Solidity: function setParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef, ActivityParticipatingOrgsParticipatingOrg _participatingOrg) returns()
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsTransactor) SetParticipatingOrg(opts *bind.TransactOpts, _activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte, _participatingOrg ActivityParticipatingOrgsParticipatingOrg) (*types.Transaction, error) {
	return _ActivityParticipatingOrgs.contract.Transact(opts, "setParticipatingOrg", _activitiesRef, _activityRef, _particpatingOrgRef, _participatingOrg)
}

// SetParticipatingOrg is a paid mutator transaction binding the contract method 0xa444eeda.
//
// Solidity: function setParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef, ActivityParticipatingOrgsParticipatingOrg _participatingOrg) returns()
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsSession) SetParticipatingOrg(_activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte, _participatingOrg ActivityParticipatingOrgsParticipatingOrg) (*types.Transaction, error) {
	return _ActivityParticipatingOrgs.Contract.SetParticipatingOrg(&_ActivityParticipatingOrgs.TransactOpts, _activitiesRef, _activityRef, _particpatingOrgRef, _participatingOrg)
}

// SetParticipatingOrg is a paid mutator transaction binding the contract method 0xa444eeda.
//
// Solidity: function setParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef, ActivityParticipatingOrgsParticipatingOrg _participatingOrg) returns()
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsTransactorSession) SetParticipatingOrg(_activitiesRef [32]byte, _activityRef [32]byte, _particpatingOrgRef [32]byte, _participatingOrg ActivityParticipatingOrgsParticipatingOrg) (*types.Transaction, error) {
	return _ActivityParticipatingOrgs.Contract.SetParticipatingOrg(&_ActivityParticipatingOrgs.TransactOpts, _activitiesRef, _activityRef, _particpatingOrgRef, _participatingOrg)
}

// ActivityParticipatingOrgsSetParticipatingOrgIterator is returned from FilterSetParticipatingOrg and is used to iterate over the raw logs and unpacked data for SetParticipatingOrg events raised by the ActivityParticipatingOrgs contract.
type ActivityParticipatingOrgsSetParticipatingOrgIterator struct {
	Event *ActivityParticipatingOrgsSetParticipatingOrg // Event containing the contract specifics and raw log

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
func (it *ActivityParticipatingOrgsSetParticipatingOrgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivityParticipatingOrgsSetParticipatingOrg)
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
		it.Event = new(ActivityParticipatingOrgsSetParticipatingOrg)
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
func (it *ActivityParticipatingOrgsSetParticipatingOrgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivityParticipatingOrgsSetParticipatingOrgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivityParticipatingOrgsSetParticipatingOrg represents a SetParticipatingOrg event raised by the ActivityParticipatingOrgs contract.
type ActivityParticipatingOrgsSetParticipatingOrg struct {
	ActivitiesRef      [32]byte
	ActivityRef        [32]byte
	ParticpatingOrgRef [32]byte
	ParticipatingOrg   ActivityParticipatingOrgsParticipatingOrg
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetParticipatingOrg is a free log retrieval operation binding the contract event 0x928f70dfd355a1b87c3df80c6373c58bf7a6e3b0acec888993915498a7f4e860.
//
// Solidity: event SetParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef, ActivityParticipatingOrgsParticipatingOrg _participatingOrg)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsFilterer) FilterSetParticipatingOrg(opts *bind.FilterOpts) (*ActivityParticipatingOrgsSetParticipatingOrgIterator, error) {

	logs, sub, err := _ActivityParticipatingOrgs.contract.FilterLogs(opts, "SetParticipatingOrg")
	if err != nil {
		return nil, err
	}
	return &ActivityParticipatingOrgsSetParticipatingOrgIterator{contract: _ActivityParticipatingOrgs.contract, event: "SetParticipatingOrg", logs: logs, sub: sub}, nil
}

// WatchSetParticipatingOrg is a free log subscription operation binding the contract event 0x928f70dfd355a1b87c3df80c6373c58bf7a6e3b0acec888993915498a7f4e860.
//
// Solidity: event SetParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef, ActivityParticipatingOrgsParticipatingOrg _participatingOrg)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsFilterer) WatchSetParticipatingOrg(opts *bind.WatchOpts, sink chan<- *ActivityParticipatingOrgsSetParticipatingOrg) (event.Subscription, error) {

	logs, sub, err := _ActivityParticipatingOrgs.contract.WatchLogs(opts, "SetParticipatingOrg")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivityParticipatingOrgsSetParticipatingOrg)
				if err := _ActivityParticipatingOrgs.contract.UnpackLog(event, "SetParticipatingOrg", log); err != nil {
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

// ParseSetParticipatingOrg is a log parse operation binding the contract event 0x928f70dfd355a1b87c3df80c6373c58bf7a6e3b0acec888993915498a7f4e860.
//
// Solidity: event SetParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef, ActivityParticipatingOrgsParticipatingOrg _participatingOrg)
func (_ActivityParticipatingOrgs *ActivityParticipatingOrgsFilterer) ParseSetParticipatingOrg(log types.Log) (*ActivityParticipatingOrgsSetParticipatingOrg, error) {
	event := new(ActivityParticipatingOrgsSetParticipatingOrg)
	if err := _ActivityParticipatingOrgs.contract.UnpackLog(event, "SetParticipatingOrg", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IATIActivityParticipatingOrgsABI is the input ABI used to generate the binding from.
const IATIActivityParticipatingOrgsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_particpatingOrgRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"crsChannelCode\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"narrative\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structActivityParticipatingOrgs.ParticipatingOrg\",\"name\":\"_participatingOrg\",\"type\":\"tuple\"}],\"name\":\"SetParticipatingOrg\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_participatingOrgRef\",\"type\":\"bytes32\"}],\"name\":\"getCrsChannelCode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_participatingOrgRef\",\"type\":\"bytes32\"}],\"name\":\"getLang\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_participatingOrgRef\",\"type\":\"bytes32\"}],\"name\":\"getNarrative\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getNumParticipatingOrgs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_participatingOrgRef\",\"type\":\"bytes32\"}],\"name\":\"getOrgRef\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_participatingOrgRef\",\"type\":\"bytes32\"}],\"name\":\"getParticipatingOrg\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"crsChannelCode\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"narrative\",\"type\":\"string\"}],\"internalType\":\"structActivityParticipatingOrgs.ParticipatingOrg\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_participatingOrgRef\",\"type\":\"bytes32\"}],\"name\":\"getRole\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_participatingOrgRef\",\"type\":\"bytes32\"}],\"name\":\"getType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_participatingOrgRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"crsChannelCode\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"narrative\",\"type\":\"string\"}],\"internalType\":\"structActivityParticipatingOrgs.ParticipatingOrg\",\"name\":\"_participatingOrg\",\"type\":\"tuple\"}],\"name\":\"setParticipatingOrg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IATIActivityParticipatingOrgsFuncSigs maps the 4-byte function signature to its string representation.
var IATIActivityParticipatingOrgsFuncSigs = map[string]string{
	"0dbc539b": "getCrsChannelCode(bytes32,bytes32,bytes32)",
	"e022cb52": "getLang(bytes32,bytes32,bytes32)",
	"ae13bb7a": "getNarrative(bytes32,bytes32,bytes32)",
	"ca39101b": "getNumParticipatingOrgs(bytes32,bytes32)",
	"d4976bce": "getOrgRef(bytes32,bytes32,bytes32)",
	"4017dc0d": "getParticipatingOrg(bytes32,bytes32,bytes32)",
	"eef67d78": "getReference(bytes32,bytes32,uint256)",
	"57db6f69": "getRole(bytes32,bytes32,bytes32)",
	"b153af81": "getType(bytes32,bytes32,bytes32)",
	"a444eeda": "setParticipatingOrg(bytes32,bytes32,bytes32,(uint8,uint8,uint256,bytes32,bytes32,string))",
}

// IATIActivityParticipatingOrgsBin is the compiled bytecode used for deploying new contracts.
var IATIActivityParticipatingOrgsBin = "0x608060405234801561001057600080fd5b50610e26806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063b153af8111610066578063b153af8114610141578063ca39101b14610154578063d4976bce14610167578063e022cb521461017a578063eef67d781461018d5761009e565b80630dbc539b146100a35780634017dc0d146100cc57806357db6f69146100ec578063a444eeda1461010c578063ae13bb7a14610121575b600080fd5b6100b66100b1366004610b62565b6101a0565b6040516100c39190610d0d565b60405180910390f35b6100df6100da366004610b62565b61021c565b6040516100c39190610da8565b6100ff6100fa366004610b62565b61036b565b6040516100c39190610dbb565b61011f61011a366004610b8d565b6103eb565b005b61013461012f366004610b62565b61066d565b6040516100c39190610d95565b6100ff61014f366004610b62565b610774565b6100b6610162366004610b41565b6107ef565b6100b6610175366004610b62565b610847565b6100b6610188366004610b62565b6108c2565b6100b661019b366004610c59565b61093d565b600083811a60f81b6001600160f81b031916158015906101cf57508260001a60f81b6001600160f81b03191615155b80156101ea57508160001a60f81b6001600160f81b03191615155b6101f357600080fd5b506000928352600160208181526040808620948652938152838520928552919091529120015490565b6102246109ce565b8360001a60f81b6001600160f81b0319161580159061025257508260001a60f81b6001600160f81b03191615155b801561026d57508160001a60f81b6001600160f81b03191615155b61027657600080fd5b60008481526001602081815260408084208785528252808420868552825292839020835160c081018552815460ff808216835261010091829004168285015282850154828701526002808401546060840152600384015460808401526004840180548851978116159093026000190190921604601f81018590048502860185019096528585529094919360a08601939092908301828280156103595780601f1061032e57610100808354040283529160200191610359565b820191906000526020600020905b81548152906001019060200180831161033c57829003601f168201915b50505050508152505090509392505050565b600083811a60f81b6001600160f81b0319161580159061039a57508260001a60f81b6001600160f81b03191615155b80156103b557508160001a60f81b6001600160f81b03191615155b6103be57600080fd5b50600092835260016020908152604080852093855292815282842091845252902054610100900460ff1690565b8360001a60f81b6001600160f81b0319161580159061041957508260001a60f81b6001600160f81b03191615155b801561043457508160001a60f81b6001600160f81b03191615155b80156104535750608081015160001a60f81b6001600160f81b03191615155b801561046657508051600a60ff90911610155b80156104785750602081015160ff1615155b801561048e5750600560ff16816020015160ff16105b80156104a05750612710816040015110155b80156104b1575060008160a0015151115b80156104d05750606081015160001a60f81b6001600160f81b03191615155b6104d957600080fd5b60008481526001602081815260408084208785528252808420868552825292839020845181548387015160ff9081166101000261ff00199190931660ff1990921691909117161781559284015191830191909155606083015160028301556080830151600383015560a0830151805184939261055c926004850192910190610a05565b50505060008481526020818152604080832086845290915290819020905163745fca7b60e01b815273__$b620240993a55615b607189176fb82e62f$__9163745fca7b916105ae918691600401610d16565b60206040518083038186803b1580156105c657600080fd5b505af41580156105da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105fe9190610b1a565b61062a576000848152602081815260408083208684528252822080546001810182559083529120018290555b7f928f70dfd355a1b87c3df80c6373c58bf7a6e3b0acec888993915498a7f4e8608484848460405161065f9493929190610d66565b60405180910390a150505050565b60608360001a60f81b6001600160f81b0319161580159061069d57508260001a60f81b6001600160f81b03191615155b80156106b857508160001a60f81b6001600160f81b03191615155b6106c157600080fd5b6000848152600160208181526040808420878552825280842086855282529283902060040180548451600294821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156107665780601f1061073b57610100808354040283529160200191610766565b820191906000526020600020905b81548152906001019060200180831161074957829003601f168201915b505050505090509392505050565b600083811a60f81b6001600160f81b031916158015906107a357508260001a60f81b6001600160f81b03191615155b80156107be57508160001a60f81b6001600160f81b03191615155b6107c757600080fd5b5060009283526001602090815260408085209385529281528284209184525290205460ff1690565b600082811a60f81b6001600160f81b0319161580159061081e57508160001a60f81b6001600160f81b03191615155b61082757600080fd5b506000828152602081815260408083208484529091529020545b92915050565b600083811a60f81b6001600160f81b0319161580159061087657508260001a60f81b6001600160f81b03191615155b801561089157508160001a60f81b6001600160f81b03191615155b61089a57600080fd5b5060009283526001602090815260408085209385529281528284209184525290206003015490565b600083811a60f81b6001600160f81b031916158015906108f157508260001a60f81b6001600160f81b03191615155b801561090c57508160001a60f81b6001600160f81b03191615155b61091557600080fd5b5060009283526001602090815260408085209385529281528284209184525290206002015490565b600083811a60f81b6001600160f81b0319161580159061096c57508260001a60f81b6001600160f81b03191615155b801561098e575060008481526020818152604080832086845290915290205482105b61099757600080fd5b60008481526020818152604080832086845290915290208054839081106109ba57fe5b906000526020600020015490509392505050565b6040805160c0810182526000808252602082018190529181018290526060808201839052608082019290925260a081019190915290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a4657805160ff1916838001178555610a73565b82800160010185558215610a73579182015b82811115610a73578251825591602001919060010190610a58565b50610a7f929150610a83565b5090565b610a9d91905b80821115610a7f5760008155600101610a89565b90565b600082601f830112610ab0578081fd5b813567ffffffffffffffff811115610ac6578182fd5b610ad9601f8201601f1916602001610dc9565b9150808252836020828501011115610af057600080fd5b8060208401602084013760009082016020015292915050565b803560ff8116811461084157600080fd5b600060208284031215610b2b578081fd5b81518015158114610b3a578182fd5b9392505050565b60008060408385031215610b53578081fd5b50508035926020909101359150565b600080600060608486031215610b76578081fd5b505081359360208301359350604090920135919050565b60008060008060808587031215610ba2578081fd5b843593506020850135925060408501359150606085013567ffffffffffffffff80821115610bce578283fd5b81870160c0818a031215610be0578384fd5b610bea60c0610dc9565b9250610bf68982610b09565b8352610c058960208301610b09565b602084015260408101356040840152606081013560608401526080810135608084015260a081013582811115610c39578485fd5b610c458a828401610aa0565b60a085015250959894975092955093505050565b600080600060608486031215610b76578283fd5b60008151808452815b81811015610c9257602081850181015186830182015201610c76565b81811115610ca35782602083870101525b50601f01601f19169290920160200192915050565b600060ff825116835260ff602083015116602084015260408201516040840152606082015160608401526080820151608084015260a082015160c060a0850152610d0560c0850182610c6d565b949350505050565b90815260200190565b60006040820184835260206040818501528185548084526060860191508685528285209350845b81811015610d5957845483526001948501949284019201610d3d565b5090979650505050505050565b600085825284602083015283604083015260806060830152610d8b6080830184610cb8565b9695505050505050565b600060208252610b3a6020830184610c6d565b600060208252610b3a6020830184610cb8565b60ff91909116815260200190565b60405181810167ffffffffffffffff81118282101715610de857600080fd5b60405291905056fea2646970667358221220751b08e680f6df3a1f26c3761d82a8845b4efc529fc3fc958ad93b53caa6657464736f6c63430006060033"

// DeployIATIActivityParticipatingOrgs deploys a new Ethereum contract, binding an instance of IATIActivityParticipatingOrgs to it.
func DeployIATIActivityParticipatingOrgs(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IATIActivityParticipatingOrgs, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivityParticipatingOrgsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	stringsAddr, _, _, _ := DeployStrings(auth, backend)
	IATIActivityParticipatingOrgsBin = strings.Replace(IATIActivityParticipatingOrgsBin, "__$b620240993a55615b607189176fb82e62f$__", stringsAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IATIActivityParticipatingOrgsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IATIActivityParticipatingOrgs{IATIActivityParticipatingOrgsCaller: IATIActivityParticipatingOrgsCaller{contract: contract}, IATIActivityParticipatingOrgsTransactor: IATIActivityParticipatingOrgsTransactor{contract: contract}, IATIActivityParticipatingOrgsFilterer: IATIActivityParticipatingOrgsFilterer{contract: contract}}, nil
}

// IATIActivityParticipatingOrgs is an auto generated Go binding around an Ethereum contract.
type IATIActivityParticipatingOrgs struct {
	IATIActivityParticipatingOrgsCaller     // Read-only binding to the contract
	IATIActivityParticipatingOrgsTransactor // Write-only binding to the contract
	IATIActivityParticipatingOrgsFilterer   // Log filterer for contract events
}

// IATIActivityParticipatingOrgsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IATIActivityParticipatingOrgsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityParticipatingOrgsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IATIActivityParticipatingOrgsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityParticipatingOrgsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IATIActivityParticipatingOrgsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityParticipatingOrgsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IATIActivityParticipatingOrgsSession struct {
	Contract     *IATIActivityParticipatingOrgs // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IATIActivityParticipatingOrgsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IATIActivityParticipatingOrgsCallerSession struct {
	Contract *IATIActivityParticipatingOrgsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// IATIActivityParticipatingOrgsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IATIActivityParticipatingOrgsTransactorSession struct {
	Contract     *IATIActivityParticipatingOrgsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// IATIActivityParticipatingOrgsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IATIActivityParticipatingOrgsRaw struct {
	Contract *IATIActivityParticipatingOrgs // Generic contract binding to access the raw methods on
}

// IATIActivityParticipatingOrgsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IATIActivityParticipatingOrgsCallerRaw struct {
	Contract *IATIActivityParticipatingOrgsCaller // Generic read-only contract binding to access the raw methods on
}

// IATIActivityParticipatingOrgsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IATIActivityParticipatingOrgsTransactorRaw struct {
	Contract *IATIActivityParticipatingOrgsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIATIActivityParticipatingOrgs creates a new instance of IATIActivityParticipatingOrgs, bound to a specific deployed contract.
func NewIATIActivityParticipatingOrgs(address common.Address, backend bind.ContractBackend) (*IATIActivityParticipatingOrgs, error) {
	contract, err := bindIATIActivityParticipatingOrgs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IATIActivityParticipatingOrgs{IATIActivityParticipatingOrgsCaller: IATIActivityParticipatingOrgsCaller{contract: contract}, IATIActivityParticipatingOrgsTransactor: IATIActivityParticipatingOrgsTransactor{contract: contract}, IATIActivityParticipatingOrgsFilterer: IATIActivityParticipatingOrgsFilterer{contract: contract}}, nil
}

// NewIATIActivityParticipatingOrgsCaller creates a new read-only instance of IATIActivityParticipatingOrgs, bound to a specific deployed contract.
func NewIATIActivityParticipatingOrgsCaller(address common.Address, caller bind.ContractCaller) (*IATIActivityParticipatingOrgsCaller, error) {
	contract, err := bindIATIActivityParticipatingOrgs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IATIActivityParticipatingOrgsCaller{contract: contract}, nil
}

// NewIATIActivityParticipatingOrgsTransactor creates a new write-only instance of IATIActivityParticipatingOrgs, bound to a specific deployed contract.
func NewIATIActivityParticipatingOrgsTransactor(address common.Address, transactor bind.ContractTransactor) (*IATIActivityParticipatingOrgsTransactor, error) {
	contract, err := bindIATIActivityParticipatingOrgs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IATIActivityParticipatingOrgsTransactor{contract: contract}, nil
}

// NewIATIActivityParticipatingOrgsFilterer creates a new log filterer instance of IATIActivityParticipatingOrgs, bound to a specific deployed contract.
func NewIATIActivityParticipatingOrgsFilterer(address common.Address, filterer bind.ContractFilterer) (*IATIActivityParticipatingOrgsFilterer, error) {
	contract, err := bindIATIActivityParticipatingOrgs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IATIActivityParticipatingOrgsFilterer{contract: contract}, nil
}

// bindIATIActivityParticipatingOrgs binds a generic wrapper to an already deployed contract.
func bindIATIActivityParticipatingOrgs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivityParticipatingOrgsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIActivityParticipatingOrgs.Contract.IATIActivityParticipatingOrgsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIActivityParticipatingOrgs.Contract.IATIActivityParticipatingOrgsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIActivityParticipatingOrgs.Contract.IATIActivityParticipatingOrgsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIActivityParticipatingOrgs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIActivityParticipatingOrgs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIActivityParticipatingOrgs.Contract.contract.Transact(opts, method, params...)
}

// GetCrsChannelCode is a free data retrieval call binding the contract method 0x0dbc539b.
//
// Solidity: function getCrsChannelCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(uint256)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCaller) GetCrsChannelCode(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIActivityParticipatingOrgs.contract.Call(opts, out, "getCrsChannelCode", _activitiesRef, _activityRef, _participatingOrgRef)
	return *ret0, err
}

// GetCrsChannelCode is a free data retrieval call binding the contract method 0x0dbc539b.
//
// Solidity: function getCrsChannelCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(uint256)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsSession) GetCrsChannelCode(_activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) (*big.Int, error) {
	return _IATIActivityParticipatingOrgs.Contract.GetCrsChannelCode(&_IATIActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _participatingOrgRef)
}

// GetCrsChannelCode is a free data retrieval call binding the contract method 0x0dbc539b.
//
// Solidity: function getCrsChannelCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(uint256)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCallerSession) GetCrsChannelCode(_activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) (*big.Int, error) {
	return _IATIActivityParticipatingOrgs.Contract.GetCrsChannelCode(&_IATIActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _participatingOrgRef)
}

// GetLang is a free data retrieval call binding the contract method 0xe022cb52.
//
// Solidity: function getLang(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(bytes32)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCaller) GetLang(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityParticipatingOrgs.contract.Call(opts, out, "getLang", _activitiesRef, _activityRef, _participatingOrgRef)
	return *ret0, err
}

// GetLang is a free data retrieval call binding the contract method 0xe022cb52.
//
// Solidity: function getLang(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(bytes32)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsSession) GetLang(_activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) ([32]byte, error) {
	return _IATIActivityParticipatingOrgs.Contract.GetLang(&_IATIActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _participatingOrgRef)
}

// GetLang is a free data retrieval call binding the contract method 0xe022cb52.
//
// Solidity: function getLang(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(bytes32)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCallerSession) GetLang(_activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) ([32]byte, error) {
	return _IATIActivityParticipatingOrgs.Contract.GetLang(&_IATIActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _participatingOrgRef)
}

// GetNarrative is a free data retrieval call binding the contract method 0xae13bb7a.
//
// Solidity: function getNarrative(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(string)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCaller) GetNarrative(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IATIActivityParticipatingOrgs.contract.Call(opts, out, "getNarrative", _activitiesRef, _activityRef, _participatingOrgRef)
	return *ret0, err
}

// GetNarrative is a free data retrieval call binding the contract method 0xae13bb7a.
//
// Solidity: function getNarrative(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(string)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsSession) GetNarrative(_activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) (string, error) {
	return _IATIActivityParticipatingOrgs.Contract.GetNarrative(&_IATIActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _participatingOrgRef)
}

// GetNarrative is a free data retrieval call binding the contract method 0xae13bb7a.
//
// Solidity: function getNarrative(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(string)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCallerSession) GetNarrative(_activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) (string, error) {
	return _IATIActivityParticipatingOrgs.Contract.GetNarrative(&_IATIActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _participatingOrgRef)
}

// GetNumParticipatingOrgs is a free data retrieval call binding the contract method 0xca39101b.
//
// Solidity: function getNumParticipatingOrgs(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCaller) GetNumParticipatingOrgs(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIActivityParticipatingOrgs.contract.Call(opts, out, "getNumParticipatingOrgs", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetNumParticipatingOrgs is a free data retrieval call binding the contract method 0xca39101b.
//
// Solidity: function getNumParticipatingOrgs(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsSession) GetNumParticipatingOrgs(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _IATIActivityParticipatingOrgs.Contract.GetNumParticipatingOrgs(&_IATIActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef)
}

// GetNumParticipatingOrgs is a free data retrieval call binding the contract method 0xca39101b.
//
// Solidity: function getNumParticipatingOrgs(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCallerSession) GetNumParticipatingOrgs(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _IATIActivityParticipatingOrgs.Contract.GetNumParticipatingOrgs(&_IATIActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef)
}

// GetOrgRef is a free data retrieval call binding the contract method 0xd4976bce.
//
// Solidity: function getOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(bytes32)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCaller) GetOrgRef(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityParticipatingOrgs.contract.Call(opts, out, "getOrgRef", _activitiesRef, _activityRef, _participatingOrgRef)
	return *ret0, err
}

// GetOrgRef is a free data retrieval call binding the contract method 0xd4976bce.
//
// Solidity: function getOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(bytes32)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsSession) GetOrgRef(_activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) ([32]byte, error) {
	return _IATIActivityParticipatingOrgs.Contract.GetOrgRef(&_IATIActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _participatingOrgRef)
}

// GetOrgRef is a free data retrieval call binding the contract method 0xd4976bce.
//
// Solidity: function getOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(bytes32)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCallerSession) GetOrgRef(_activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) ([32]byte, error) {
	return _IATIActivityParticipatingOrgs.Contract.GetOrgRef(&_IATIActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _participatingOrgRef)
}

// GetParticipatingOrg is a free data retrieval call binding the contract method 0x4017dc0d.
//
// Solidity: function getParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(ActivityParticipatingOrgsParticipatingOrg)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCaller) GetParticipatingOrg(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) (ActivityParticipatingOrgsParticipatingOrg, error) {
	var (
		ret0 = new(ActivityParticipatingOrgsParticipatingOrg)
	)
	out := ret0
	err := _IATIActivityParticipatingOrgs.contract.Call(opts, out, "getParticipatingOrg", _activitiesRef, _activityRef, _participatingOrgRef)
	return *ret0, err
}

// GetParticipatingOrg is a free data retrieval call binding the contract method 0x4017dc0d.
//
// Solidity: function getParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(ActivityParticipatingOrgsParticipatingOrg)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsSession) GetParticipatingOrg(_activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) (ActivityParticipatingOrgsParticipatingOrg, error) {
	return _IATIActivityParticipatingOrgs.Contract.GetParticipatingOrg(&_IATIActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _participatingOrgRef)
}

// GetParticipatingOrg is a free data retrieval call binding the contract method 0x4017dc0d.
//
// Solidity: function getParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(ActivityParticipatingOrgsParticipatingOrg)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCallerSession) GetParticipatingOrg(_activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) (ActivityParticipatingOrgsParticipatingOrg, error) {
	return _IATIActivityParticipatingOrgs.Contract.GetParticipatingOrg(&_IATIActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _participatingOrgRef)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCaller) GetReference(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityParticipatingOrgs.contract.Call(opts, out, "getReference", _activitiesRef, _activityRef, _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIActivityParticipatingOrgs.Contract.GetReference(&_IATIActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCallerSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIActivityParticipatingOrgs.Contract.GetReference(&_IATIActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetRole is a free data retrieval call binding the contract method 0x57db6f69.
//
// Solidity: function getRole(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(uint8)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCaller) GetRole(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivityParticipatingOrgs.contract.Call(opts, out, "getRole", _activitiesRef, _activityRef, _participatingOrgRef)
	return *ret0, err
}

// GetRole is a free data retrieval call binding the contract method 0x57db6f69.
//
// Solidity: function getRole(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(uint8)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsSession) GetRole(_activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) (uint8, error) {
	return _IATIActivityParticipatingOrgs.Contract.GetRole(&_IATIActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _participatingOrgRef)
}

// GetRole is a free data retrieval call binding the contract method 0x57db6f69.
//
// Solidity: function getRole(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(uint8)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCallerSession) GetRole(_activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) (uint8, error) {
	return _IATIActivityParticipatingOrgs.Contract.GetRole(&_IATIActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _participatingOrgRef)
}

// GetType is a free data retrieval call binding the contract method 0xb153af81.
//
// Solidity: function getType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(uint8)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCaller) GetType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivityParticipatingOrgs.contract.Call(opts, out, "getType", _activitiesRef, _activityRef, _participatingOrgRef)
	return *ret0, err
}

// GetType is a free data retrieval call binding the contract method 0xb153af81.
//
// Solidity: function getType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(uint8)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsSession) GetType(_activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) (uint8, error) {
	return _IATIActivityParticipatingOrgs.Contract.GetType(&_IATIActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _participatingOrgRef)
}

// GetType is a free data retrieval call binding the contract method 0xb153af81.
//
// Solidity: function getType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef) view returns(uint8)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsCallerSession) GetType(_activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte) (uint8, error) {
	return _IATIActivityParticipatingOrgs.Contract.GetType(&_IATIActivityParticipatingOrgs.CallOpts, _activitiesRef, _activityRef, _participatingOrgRef)
}

// SetParticipatingOrg is a paid mutator transaction binding the contract method 0xa444eeda.
//
// Solidity: function setParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef, ActivityParticipatingOrgsParticipatingOrg _participatingOrg) returns()
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsTransactor) SetParticipatingOrg(opts *bind.TransactOpts, _activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte, _participatingOrg ActivityParticipatingOrgsParticipatingOrg) (*types.Transaction, error) {
	return _IATIActivityParticipatingOrgs.contract.Transact(opts, "setParticipatingOrg", _activitiesRef, _activityRef, _participatingOrgRef, _participatingOrg)
}

// SetParticipatingOrg is a paid mutator transaction binding the contract method 0xa444eeda.
//
// Solidity: function setParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef, ActivityParticipatingOrgsParticipatingOrg _participatingOrg) returns()
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsSession) SetParticipatingOrg(_activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte, _participatingOrg ActivityParticipatingOrgsParticipatingOrg) (*types.Transaction, error) {
	return _IATIActivityParticipatingOrgs.Contract.SetParticipatingOrg(&_IATIActivityParticipatingOrgs.TransactOpts, _activitiesRef, _activityRef, _participatingOrgRef, _participatingOrg)
}

// SetParticipatingOrg is a paid mutator transaction binding the contract method 0xa444eeda.
//
// Solidity: function setParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _participatingOrgRef, ActivityParticipatingOrgsParticipatingOrg _participatingOrg) returns()
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsTransactorSession) SetParticipatingOrg(_activitiesRef [32]byte, _activityRef [32]byte, _participatingOrgRef [32]byte, _participatingOrg ActivityParticipatingOrgsParticipatingOrg) (*types.Transaction, error) {
	return _IATIActivityParticipatingOrgs.Contract.SetParticipatingOrg(&_IATIActivityParticipatingOrgs.TransactOpts, _activitiesRef, _activityRef, _participatingOrgRef, _participatingOrg)
}

// IATIActivityParticipatingOrgsSetParticipatingOrgIterator is returned from FilterSetParticipatingOrg and is used to iterate over the raw logs and unpacked data for SetParticipatingOrg events raised by the IATIActivityParticipatingOrgs contract.
type IATIActivityParticipatingOrgsSetParticipatingOrgIterator struct {
	Event *IATIActivityParticipatingOrgsSetParticipatingOrg // Event containing the contract specifics and raw log

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
func (it *IATIActivityParticipatingOrgsSetParticipatingOrgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IATIActivityParticipatingOrgsSetParticipatingOrg)
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
		it.Event = new(IATIActivityParticipatingOrgsSetParticipatingOrg)
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
func (it *IATIActivityParticipatingOrgsSetParticipatingOrgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IATIActivityParticipatingOrgsSetParticipatingOrgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IATIActivityParticipatingOrgsSetParticipatingOrg represents a SetParticipatingOrg event raised by the IATIActivityParticipatingOrgs contract.
type IATIActivityParticipatingOrgsSetParticipatingOrg struct {
	ActivitiesRef      [32]byte
	ActivityRef        [32]byte
	ParticpatingOrgRef [32]byte
	ParticipatingOrg   ActivityParticipatingOrgsParticipatingOrg
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSetParticipatingOrg is a free log retrieval operation binding the contract event 0x928f70dfd355a1b87c3df80c6373c58bf7a6e3b0acec888993915498a7f4e860.
//
// Solidity: event SetParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef, ActivityParticipatingOrgsParticipatingOrg _participatingOrg)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsFilterer) FilterSetParticipatingOrg(opts *bind.FilterOpts) (*IATIActivityParticipatingOrgsSetParticipatingOrgIterator, error) {

	logs, sub, err := _IATIActivityParticipatingOrgs.contract.FilterLogs(opts, "SetParticipatingOrg")
	if err != nil {
		return nil, err
	}
	return &IATIActivityParticipatingOrgsSetParticipatingOrgIterator{contract: _IATIActivityParticipatingOrgs.contract, event: "SetParticipatingOrg", logs: logs, sub: sub}, nil
}

// WatchSetParticipatingOrg is a free log subscription operation binding the contract event 0x928f70dfd355a1b87c3df80c6373c58bf7a6e3b0acec888993915498a7f4e860.
//
// Solidity: event SetParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef, ActivityParticipatingOrgsParticipatingOrg _participatingOrg)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsFilterer) WatchSetParticipatingOrg(opts *bind.WatchOpts, sink chan<- *IATIActivityParticipatingOrgsSetParticipatingOrg) (event.Subscription, error) {

	logs, sub, err := _IATIActivityParticipatingOrgs.contract.WatchLogs(opts, "SetParticipatingOrg")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IATIActivityParticipatingOrgsSetParticipatingOrg)
				if err := _IATIActivityParticipatingOrgs.contract.UnpackLog(event, "SetParticipatingOrg", log); err != nil {
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

// ParseSetParticipatingOrg is a log parse operation binding the contract event 0x928f70dfd355a1b87c3df80c6373c58bf7a6e3b0acec888993915498a7f4e860.
//
// Solidity: event SetParticipatingOrg(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _particpatingOrgRef, ActivityParticipatingOrgsParticipatingOrg _participatingOrg)
func (_IATIActivityParticipatingOrgs *IATIActivityParticipatingOrgsFilterer) ParseSetParticipatingOrg(log types.Log) (*IATIActivityParticipatingOrgsSetParticipatingOrg, error) {
	event := new(IATIActivityParticipatingOrgsSetParticipatingOrg)
	if err := _IATIActivityParticipatingOrgs.contract.UnpackLog(event, "SetParticipatingOrg", log); err != nil {
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
