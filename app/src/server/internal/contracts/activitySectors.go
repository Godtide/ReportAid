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

// ActivitySectorsSector is an auto generated low-level Go binding around an user-defined struct.
type ActivitySectorsSector struct {
	Percentage uint8
	DacCode    *big.Int
}

// ActivitySectorsABI is the input ABI used to generate the binding from.
const ActivitySectorsABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sectorRef\",\"type\":\"bytes32\"}],\"name\":\"getDACCode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sectorRef\",\"type\":\"bytes32\"}],\"name\":\"getPercentage\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sectorRef\",\"type\":\"bytes32\"}],\"name\":\"getSector\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"percentage\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dacCode\",\"type\":\"uint256\"}],\"internalType\":\"structActivitySectors.Sector\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sectorRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"percentage\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dacCode\",\"type\":\"uint256\"}],\"internalType\":\"structActivitySectors.Sector\",\"name\":\"_sector\",\"type\":\"tuple\"}],\"name\":\"setSector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ActivitySectorsFuncSigs maps the 4-byte function signature to its string representation.
var ActivitySectorsFuncSigs = map[string]string{
	"fa00fc2a": "getDACCode(bytes32,bytes32,bytes32)",
	"47cdae01": "getNum(bytes32,bytes32)",
	"a8d77aa1": "getPercentage(bytes32,bytes32,bytes32)",
	"eef67d78": "getReference(bytes32,bytes32,uint256)",
	"19eb36ad": "getSector(bytes32,bytes32,bytes32)",
	"65251f11": "setSector(bytes32,bytes32,bytes32,(uint8,uint256))",
}

// ActivitySectors is an auto generated Go binding around an Ethereum contract.
type ActivitySectors struct {
	ActivitySectorsCaller     // Read-only binding to the contract
	ActivitySectorsTransactor // Write-only binding to the contract
	ActivitySectorsFilterer   // Log filterer for contract events
}

// ActivitySectorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivitySectorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivitySectorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivitySectorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivitySectorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivitySectorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivitySectorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivitySectorsSession struct {
	Contract     *ActivitySectors  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ActivitySectorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivitySectorsCallerSession struct {
	Contract *ActivitySectorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ActivitySectorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivitySectorsTransactorSession struct {
	Contract     *ActivitySectorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ActivitySectorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivitySectorsRaw struct {
	Contract *ActivitySectors // Generic contract binding to access the raw methods on
}

// ActivitySectorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivitySectorsCallerRaw struct {
	Contract *ActivitySectorsCaller // Generic read-only contract binding to access the raw methods on
}

// ActivitySectorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivitySectorsTransactorRaw struct {
	Contract *ActivitySectorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivitySectors creates a new instance of ActivitySectors, bound to a specific deployed contract.
func NewActivitySectors(address common.Address, backend bind.ContractBackend) (*ActivitySectors, error) {
	contract, err := bindActivitySectors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ActivitySectors{ActivitySectorsCaller: ActivitySectorsCaller{contract: contract}, ActivitySectorsTransactor: ActivitySectorsTransactor{contract: contract}, ActivitySectorsFilterer: ActivitySectorsFilterer{contract: contract}}, nil
}

// NewActivitySectorsCaller creates a new read-only instance of ActivitySectors, bound to a specific deployed contract.
func NewActivitySectorsCaller(address common.Address, caller bind.ContractCaller) (*ActivitySectorsCaller, error) {
	contract, err := bindActivitySectors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivitySectorsCaller{contract: contract}, nil
}

// NewActivitySectorsTransactor creates a new write-only instance of ActivitySectors, bound to a specific deployed contract.
func NewActivitySectorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivitySectorsTransactor, error) {
	contract, err := bindActivitySectors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivitySectorsTransactor{contract: contract}, nil
}

// NewActivitySectorsFilterer creates a new log filterer instance of ActivitySectors, bound to a specific deployed contract.
func NewActivitySectorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivitySectorsFilterer, error) {
	contract, err := bindActivitySectors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivitySectorsFilterer{contract: contract}, nil
}

// bindActivitySectors binds a generic wrapper to an already deployed contract.
func bindActivitySectors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivitySectorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivitySectors *ActivitySectorsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivitySectors.Contract.ActivitySectorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivitySectors *ActivitySectorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivitySectors.Contract.ActivitySectorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivitySectors *ActivitySectorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivitySectors.Contract.ActivitySectorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivitySectors *ActivitySectorsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivitySectors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivitySectors *ActivitySectorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivitySectors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivitySectors *ActivitySectorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivitySectors.Contract.contract.Transact(opts, method, params...)
}

// GetDACCode is a free data retrieval call binding the contract method 0xfa00fc2a.
//
// Solidity: function getDACCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) view returns(uint256)
func (_ActivitySectors *ActivitySectorsCaller) GetDACCode(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ActivitySectors.contract.Call(opts, out, "getDACCode", _activitiesRef, _activityRef, _sectorRef)
	return *ret0, err
}

// GetDACCode is a free data retrieval call binding the contract method 0xfa00fc2a.
//
// Solidity: function getDACCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) view returns(uint256)
func (_ActivitySectors *ActivitySectorsSession) GetDACCode(_activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte) (*big.Int, error) {
	return _ActivitySectors.Contract.GetDACCode(&_ActivitySectors.CallOpts, _activitiesRef, _activityRef, _sectorRef)
}

// GetDACCode is a free data retrieval call binding the contract method 0xfa00fc2a.
//
// Solidity: function getDACCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) view returns(uint256)
func (_ActivitySectors *ActivitySectorsCallerSession) GetDACCode(_activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte) (*big.Int, error) {
	return _ActivitySectors.Contract.GetDACCode(&_ActivitySectors.CallOpts, _activitiesRef, _activityRef, _sectorRef)
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivitySectors *ActivitySectorsCaller) GetNum(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ActivitySectors.contract.Call(opts, out, "getNum", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivitySectors *ActivitySectorsSession) GetNum(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _ActivitySectors.Contract.GetNum(&_ActivitySectors.CallOpts, _activitiesRef, _activityRef)
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivitySectors *ActivitySectorsCallerSession) GetNum(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _ActivitySectors.Contract.GetNum(&_ActivitySectors.CallOpts, _activitiesRef, _activityRef)
}

// GetPercentage is a free data retrieval call binding the contract method 0xa8d77aa1.
//
// Solidity: function getPercentage(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) view returns(uint8)
func (_ActivitySectors *ActivitySectorsCaller) GetPercentage(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ActivitySectors.contract.Call(opts, out, "getPercentage", _activitiesRef, _activityRef, _sectorRef)
	return *ret0, err
}

// GetPercentage is a free data retrieval call binding the contract method 0xa8d77aa1.
//
// Solidity: function getPercentage(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) view returns(uint8)
func (_ActivitySectors *ActivitySectorsSession) GetPercentage(_activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte) (uint8, error) {
	return _ActivitySectors.Contract.GetPercentage(&_ActivitySectors.CallOpts, _activitiesRef, _activityRef, _sectorRef)
}

// GetPercentage is a free data retrieval call binding the contract method 0xa8d77aa1.
//
// Solidity: function getPercentage(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) view returns(uint8)
func (_ActivitySectors *ActivitySectorsCallerSession) GetPercentage(_activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte) (uint8, error) {
	return _ActivitySectors.Contract.GetPercentage(&_ActivitySectors.CallOpts, _activitiesRef, _activityRef, _sectorRef)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivitySectors *ActivitySectorsCaller) GetReference(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivitySectors.contract.Call(opts, out, "getReference", _activitiesRef, _activityRef, _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivitySectors *ActivitySectorsSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _ActivitySectors.Contract.GetReference(&_ActivitySectors.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivitySectors *ActivitySectorsCallerSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _ActivitySectors.Contract.GetReference(&_ActivitySectors.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetSector is a free data retrieval call binding the contract method 0x19eb36ad.
//
// Solidity: function getSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) view returns(ActivitySectorsSector)
func (_ActivitySectors *ActivitySectorsCaller) GetSector(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte) (ActivitySectorsSector, error) {
	var (
		ret0 = new(ActivitySectorsSector)
	)
	out := ret0
	err := _ActivitySectors.contract.Call(opts, out, "getSector", _activitiesRef, _activityRef, _sectorRef)
	return *ret0, err
}

// GetSector is a free data retrieval call binding the contract method 0x19eb36ad.
//
// Solidity: function getSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) view returns(ActivitySectorsSector)
func (_ActivitySectors *ActivitySectorsSession) GetSector(_activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte) (ActivitySectorsSector, error) {
	return _ActivitySectors.Contract.GetSector(&_ActivitySectors.CallOpts, _activitiesRef, _activityRef, _sectorRef)
}

// GetSector is a free data retrieval call binding the contract method 0x19eb36ad.
//
// Solidity: function getSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) view returns(ActivitySectorsSector)
func (_ActivitySectors *ActivitySectorsCallerSession) GetSector(_activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte) (ActivitySectorsSector, error) {
	return _ActivitySectors.Contract.GetSector(&_ActivitySectors.CallOpts, _activitiesRef, _activityRef, _sectorRef)
}

// SetSector is a paid mutator transaction binding the contract method 0x65251f11.
//
// Solidity: function setSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef, ActivitySectorsSector _sector) returns()
func (_ActivitySectors *ActivitySectorsTransactor) SetSector(opts *bind.TransactOpts, _activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte, _sector ActivitySectorsSector) (*types.Transaction, error) {
	return _ActivitySectors.contract.Transact(opts, "setSector", _activitiesRef, _activityRef, _sectorRef, _sector)
}

// SetSector is a paid mutator transaction binding the contract method 0x65251f11.
//
// Solidity: function setSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef, ActivitySectorsSector _sector) returns()
func (_ActivitySectors *ActivitySectorsSession) SetSector(_activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte, _sector ActivitySectorsSector) (*types.Transaction, error) {
	return _ActivitySectors.Contract.SetSector(&_ActivitySectors.TransactOpts, _activitiesRef, _activityRef, _sectorRef, _sector)
}

// SetSector is a paid mutator transaction binding the contract method 0x65251f11.
//
// Solidity: function setSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef, ActivitySectorsSector _sector) returns()
func (_ActivitySectors *ActivitySectorsTransactorSession) SetSector(_activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte, _sector ActivitySectorsSector) (*types.Transaction, error) {
	return _ActivitySectors.Contract.SetSector(&_ActivitySectors.TransactOpts, _activitiesRef, _activityRef, _sectorRef, _sector)
}

// IATIActivitySectorsABI is the input ABI used to generate the binding from.
const IATIActivitySectorsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_sectorRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"percentage\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dacCode\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structActivitySectors.Sector\",\"name\":\"_sector\",\"type\":\"tuple\"}],\"name\":\"SetSector\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sectorRef\",\"type\":\"bytes32\"}],\"name\":\"getDACCode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sectorRef\",\"type\":\"bytes32\"}],\"name\":\"getPercentage\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sectorRef\",\"type\":\"bytes32\"}],\"name\":\"getSector\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"percentage\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dacCode\",\"type\":\"uint256\"}],\"internalType\":\"structActivitySectors.Sector\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sectorRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"percentage\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dacCode\",\"type\":\"uint256\"}],\"internalType\":\"structActivitySectors.Sector\",\"name\":\"_sector\",\"type\":\"tuple\"}],\"name\":\"setSector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IATIActivitySectorsFuncSigs maps the 4-byte function signature to its string representation.
var IATIActivitySectorsFuncSigs = map[string]string{
	"fa00fc2a": "getDACCode(bytes32,bytes32,bytes32)",
	"47cdae01": "getNum(bytes32,bytes32)",
	"a8d77aa1": "getPercentage(bytes32,bytes32,bytes32)",
	"eef67d78": "getReference(bytes32,bytes32,uint256)",
	"19eb36ad": "getSector(bytes32,bytes32,bytes32)",
	"65251f11": "setSector(bytes32,bytes32,bytes32,(uint8,uint256))",
}

// IATIActivitySectorsBin is the compiled bytecode used for deploying new contracts.
var IATIActivitySectorsBin = "0x608060405234801561001057600080fd5b50610760806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806319eb36ad1461006757806347cdae011461009057806365251f11146100b0578063a8d77aa1146100c5578063eef67d78146100e5578063fa00fc2a146100f8575b600080fd5b61007a6100753660046105b2565b61010b565b604051610087919061070e565b60405180910390f35b6100a361009e366004610591565b6101a6565b604051610087919061068a565b6100c36100be3660046105dd565b6101fe565b005b6100d86100d33660046105b2565b6103cb565b604051610087919061071c565b6100a36100f3366004610664565b610446565b6100a36101063660046105b2565b6104d7565b610113610553565b8360001a60f81b6001600160f81b0319161580159061014157508260001a60f81b6001600160f81b03191615155b801561015c57508160001a60f81b6001600160f81b03191615155b61016557600080fd5b50600083815260016020818152604080842086855282528084208585528252928390208351808501909452805460ff16845290910154908201529392505050565b600082811a60f81b6001600160f81b031916158015906101d557508160001a60f81b6001600160f81b03191615155b6101de57600080fd5b506000828152602081815260408083208484529091529020545b92915050565b8360001a60f81b6001600160f81b0319161580159061022c57508260001a60f81b6001600160f81b03191615155b801561024757508160001a60f81b6001600160f81b03191615155b80156102595750612b66816020015110155b801561026c5750620185ec816020015111155b801561028057506064816000015160ff1611155b61028957600080fd5b6000848152600160208181526040808420878552825280842086855282528084208551815460ff191660ff90911617815585830151930192909255868352828152818320868452905290819020905163745fca7b60e01b815273__$b620240993a55615b607189176fb82e62f$__9163745fca7b9161030c918691600401610693565b60206040518083038186803b15801561032457600080fd5b505af4158015610338573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061035c919061056a565b610388576000848152602081815260408083208684528252822080546001810182559083529120018290555b7fda3cd6206e5d5a9ca79d812c149f6987f7332ab438d704c76c5c3cc28a0b5a21848484846040516103bd94939291906106e3565b60405180910390a150505050565b600083811a60f81b6001600160f81b031916158015906103fa57508260001a60f81b6001600160f81b03191615155b801561041557508160001a60f81b6001600160f81b03191615155b61041e57600080fd5b5060009283526001602090815260408085209385529281528284209184525290205460ff1690565b600083811a60f81b6001600160f81b0319161580159061047557508260001a60f81b6001600160f81b03191615155b8015610497575060008481526020818152604080832086845290915290205482105b6104a057600080fd5b60008481526020818152604080832086845290915290208054839081106104c357fe5b906000526020600020015490509392505050565b600083811a60f81b6001600160f81b0319161580159061050657508260001a60f81b6001600160f81b03191615155b801561052157508160001a60f81b6001600160f81b03191615155b61052a57600080fd5b506000928352600160208181526040808620948652938152838520928552919091529120015490565b604080518082019091526000808252602082015290565b60006020828403121561057b578081fd5b8151801515811461058a578182fd5b9392505050565b600080604083850312156105a3578081fd5b50508035926020909101359150565b6000806000606084860312156105c6578081fd5b505081359360208301359350604090920135919050565b60008060008084860360a08112156105f3578182fd5b85359450602086013593506040808701359350605f1982011215610615578182fd5b506040516040810181811067ffffffffffffffff82111715610635578283fd5b604052606086013560ff8116811461064b578283fd5b8152608095909501356020860152509194909350909190565b6000806000606084860312156105c6578283fd5b805160ff168252602090810151910152565b90815260200190565b60006040820184835260206040818501528185548084526060860191508685528285209350845b818110156106d6578454835260019485019492840192016106ba565b5090979650505050505050565b848152602081018490526040810183905260a081016107056060830184610678565b95945050505050565b604081016101f88284610678565b60ff9190911681526020019056fea2646970667358221220285c1e0c37fcc17d7b1f0bb8a1c1305c5ba98c20d8dfa79ba9342cf3ce50aefd64736f6c63430006060033"

// DeployIATIActivitySectors deploys a new Ethereum contract, binding an instance of IATIActivitySectors to it.
func DeployIATIActivitySectors(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IATIActivitySectors, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivitySectorsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	stringsAddr, _, _, _ := DeployStrings(auth, backend)
	IATIActivitySectorsBin = strings.Replace(IATIActivitySectorsBin, "__$b620240993a55615b607189176fb82e62f$__", stringsAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IATIActivitySectorsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IATIActivitySectors{IATIActivitySectorsCaller: IATIActivitySectorsCaller{contract: contract}, IATIActivitySectorsTransactor: IATIActivitySectorsTransactor{contract: contract}, IATIActivitySectorsFilterer: IATIActivitySectorsFilterer{contract: contract}}, nil
}

// IATIActivitySectors is an auto generated Go binding around an Ethereum contract.
type IATIActivitySectors struct {
	IATIActivitySectorsCaller     // Read-only binding to the contract
	IATIActivitySectorsTransactor // Write-only binding to the contract
	IATIActivitySectorsFilterer   // Log filterer for contract events
}

// IATIActivitySectorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IATIActivitySectorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivitySectorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IATIActivitySectorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivitySectorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IATIActivitySectorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivitySectorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IATIActivitySectorsSession struct {
	Contract     *IATIActivitySectors // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IATIActivitySectorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IATIActivitySectorsCallerSession struct {
	Contract *IATIActivitySectorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IATIActivitySectorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IATIActivitySectorsTransactorSession struct {
	Contract     *IATIActivitySectorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IATIActivitySectorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IATIActivitySectorsRaw struct {
	Contract *IATIActivitySectors // Generic contract binding to access the raw methods on
}

// IATIActivitySectorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IATIActivitySectorsCallerRaw struct {
	Contract *IATIActivitySectorsCaller // Generic read-only contract binding to access the raw methods on
}

// IATIActivitySectorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IATIActivitySectorsTransactorRaw struct {
	Contract *IATIActivitySectorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIATIActivitySectors creates a new instance of IATIActivitySectors, bound to a specific deployed contract.
func NewIATIActivitySectors(address common.Address, backend bind.ContractBackend) (*IATIActivitySectors, error) {
	contract, err := bindIATIActivitySectors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IATIActivitySectors{IATIActivitySectorsCaller: IATIActivitySectorsCaller{contract: contract}, IATIActivitySectorsTransactor: IATIActivitySectorsTransactor{contract: contract}, IATIActivitySectorsFilterer: IATIActivitySectorsFilterer{contract: contract}}, nil
}

// NewIATIActivitySectorsCaller creates a new read-only instance of IATIActivitySectors, bound to a specific deployed contract.
func NewIATIActivitySectorsCaller(address common.Address, caller bind.ContractCaller) (*IATIActivitySectorsCaller, error) {
	contract, err := bindIATIActivitySectors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IATIActivitySectorsCaller{contract: contract}, nil
}

// NewIATIActivitySectorsTransactor creates a new write-only instance of IATIActivitySectors, bound to a specific deployed contract.
func NewIATIActivitySectorsTransactor(address common.Address, transactor bind.ContractTransactor) (*IATIActivitySectorsTransactor, error) {
	contract, err := bindIATIActivitySectors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IATIActivitySectorsTransactor{contract: contract}, nil
}

// NewIATIActivitySectorsFilterer creates a new log filterer instance of IATIActivitySectors, bound to a specific deployed contract.
func NewIATIActivitySectorsFilterer(address common.Address, filterer bind.ContractFilterer) (*IATIActivitySectorsFilterer, error) {
	contract, err := bindIATIActivitySectors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IATIActivitySectorsFilterer{contract: contract}, nil
}

// bindIATIActivitySectors binds a generic wrapper to an already deployed contract.
func bindIATIActivitySectors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivitySectorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIActivitySectors *IATIActivitySectorsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIActivitySectors.Contract.IATIActivitySectorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIActivitySectors *IATIActivitySectorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIActivitySectors.Contract.IATIActivitySectorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIActivitySectors *IATIActivitySectorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIActivitySectors.Contract.IATIActivitySectorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIActivitySectors *IATIActivitySectorsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIActivitySectors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIActivitySectors *IATIActivitySectorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIActivitySectors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIActivitySectors *IATIActivitySectorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIActivitySectors.Contract.contract.Transact(opts, method, params...)
}

// GetDACCode is a free data retrieval call binding the contract method 0xfa00fc2a.
//
// Solidity: function getDACCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) view returns(uint256)
func (_IATIActivitySectors *IATIActivitySectorsCaller) GetDACCode(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIActivitySectors.contract.Call(opts, out, "getDACCode", _activitiesRef, _activityRef, _sectorRef)
	return *ret0, err
}

// GetDACCode is a free data retrieval call binding the contract method 0xfa00fc2a.
//
// Solidity: function getDACCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) view returns(uint256)
func (_IATIActivitySectors *IATIActivitySectorsSession) GetDACCode(_activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte) (*big.Int, error) {
	return _IATIActivitySectors.Contract.GetDACCode(&_IATIActivitySectors.CallOpts, _activitiesRef, _activityRef, _sectorRef)
}

// GetDACCode is a free data retrieval call binding the contract method 0xfa00fc2a.
//
// Solidity: function getDACCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) view returns(uint256)
func (_IATIActivitySectors *IATIActivitySectorsCallerSession) GetDACCode(_activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte) (*big.Int, error) {
	return _IATIActivitySectors.Contract.GetDACCode(&_IATIActivitySectors.CallOpts, _activitiesRef, _activityRef, _sectorRef)
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivitySectors *IATIActivitySectorsCaller) GetNum(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIActivitySectors.contract.Call(opts, out, "getNum", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivitySectors *IATIActivitySectorsSession) GetNum(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _IATIActivitySectors.Contract.GetNum(&_IATIActivitySectors.CallOpts, _activitiesRef, _activityRef)
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivitySectors *IATIActivitySectorsCallerSession) GetNum(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _IATIActivitySectors.Contract.GetNum(&_IATIActivitySectors.CallOpts, _activitiesRef, _activityRef)
}

// GetPercentage is a free data retrieval call binding the contract method 0xa8d77aa1.
//
// Solidity: function getPercentage(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) view returns(uint8)
func (_IATIActivitySectors *IATIActivitySectorsCaller) GetPercentage(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivitySectors.contract.Call(opts, out, "getPercentage", _activitiesRef, _activityRef, _sectorRef)
	return *ret0, err
}

// GetPercentage is a free data retrieval call binding the contract method 0xa8d77aa1.
//
// Solidity: function getPercentage(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) view returns(uint8)
func (_IATIActivitySectors *IATIActivitySectorsSession) GetPercentage(_activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte) (uint8, error) {
	return _IATIActivitySectors.Contract.GetPercentage(&_IATIActivitySectors.CallOpts, _activitiesRef, _activityRef, _sectorRef)
}

// GetPercentage is a free data retrieval call binding the contract method 0xa8d77aa1.
//
// Solidity: function getPercentage(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) view returns(uint8)
func (_IATIActivitySectors *IATIActivitySectorsCallerSession) GetPercentage(_activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte) (uint8, error) {
	return _IATIActivitySectors.Contract.GetPercentage(&_IATIActivitySectors.CallOpts, _activitiesRef, _activityRef, _sectorRef)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivitySectors *IATIActivitySectorsCaller) GetReference(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivitySectors.contract.Call(opts, out, "getReference", _activitiesRef, _activityRef, _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivitySectors *IATIActivitySectorsSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIActivitySectors.Contract.GetReference(&_IATIActivitySectors.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivitySectors *IATIActivitySectorsCallerSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIActivitySectors.Contract.GetReference(&_IATIActivitySectors.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetSector is a free data retrieval call binding the contract method 0x19eb36ad.
//
// Solidity: function getSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) view returns(ActivitySectorsSector)
func (_IATIActivitySectors *IATIActivitySectorsCaller) GetSector(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte) (ActivitySectorsSector, error) {
	var (
		ret0 = new(ActivitySectorsSector)
	)
	out := ret0
	err := _IATIActivitySectors.contract.Call(opts, out, "getSector", _activitiesRef, _activityRef, _sectorRef)
	return *ret0, err
}

// GetSector is a free data retrieval call binding the contract method 0x19eb36ad.
//
// Solidity: function getSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) view returns(ActivitySectorsSector)
func (_IATIActivitySectors *IATIActivitySectorsSession) GetSector(_activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte) (ActivitySectorsSector, error) {
	return _IATIActivitySectors.Contract.GetSector(&_IATIActivitySectors.CallOpts, _activitiesRef, _activityRef, _sectorRef)
}

// GetSector is a free data retrieval call binding the contract method 0x19eb36ad.
//
// Solidity: function getSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef) view returns(ActivitySectorsSector)
func (_IATIActivitySectors *IATIActivitySectorsCallerSession) GetSector(_activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte) (ActivitySectorsSector, error) {
	return _IATIActivitySectors.Contract.GetSector(&_IATIActivitySectors.CallOpts, _activitiesRef, _activityRef, _sectorRef)
}

// SetSector is a paid mutator transaction binding the contract method 0x65251f11.
//
// Solidity: function setSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef, ActivitySectorsSector _sector) returns()
func (_IATIActivitySectors *IATIActivitySectorsTransactor) SetSector(opts *bind.TransactOpts, _activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte, _sector ActivitySectorsSector) (*types.Transaction, error) {
	return _IATIActivitySectors.contract.Transact(opts, "setSector", _activitiesRef, _activityRef, _sectorRef, _sector)
}

// SetSector is a paid mutator transaction binding the contract method 0x65251f11.
//
// Solidity: function setSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef, ActivitySectorsSector _sector) returns()
func (_IATIActivitySectors *IATIActivitySectorsSession) SetSector(_activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte, _sector ActivitySectorsSector) (*types.Transaction, error) {
	return _IATIActivitySectors.Contract.SetSector(&_IATIActivitySectors.TransactOpts, _activitiesRef, _activityRef, _sectorRef, _sector)
}

// SetSector is a paid mutator transaction binding the contract method 0x65251f11.
//
// Solidity: function setSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef, ActivitySectorsSector _sector) returns()
func (_IATIActivitySectors *IATIActivitySectorsTransactorSession) SetSector(_activitiesRef [32]byte, _activityRef [32]byte, _sectorRef [32]byte, _sector ActivitySectorsSector) (*types.Transaction, error) {
	return _IATIActivitySectors.Contract.SetSector(&_IATIActivitySectors.TransactOpts, _activitiesRef, _activityRef, _sectorRef, _sector)
}

// IATIActivitySectorsSetSectorIterator is returned from FilterSetSector and is used to iterate over the raw logs and unpacked data for SetSector events raised by the IATIActivitySectors contract.
type IATIActivitySectorsSetSectorIterator struct {
	Event *IATIActivitySectorsSetSector // Event containing the contract specifics and raw log

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
func (it *IATIActivitySectorsSetSectorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IATIActivitySectorsSetSector)
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
		it.Event = new(IATIActivitySectorsSetSector)
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
func (it *IATIActivitySectorsSetSectorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IATIActivitySectorsSetSectorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IATIActivitySectorsSetSector represents a SetSector event raised by the IATIActivitySectors contract.
type IATIActivitySectorsSetSector struct {
	ActivitiesRef [32]byte
	ActivityRef   [32]byte
	SectorRef     [32]byte
	Sector        ActivitySectorsSector
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetSector is a free log retrieval operation binding the contract event 0xda3cd6206e5d5a9ca79d812c149f6987f7332ab438d704c76c5c3cc28a0b5a21.
//
// Solidity: event SetSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef, ActivitySectorsSector _sector)
func (_IATIActivitySectors *IATIActivitySectorsFilterer) FilterSetSector(opts *bind.FilterOpts) (*IATIActivitySectorsSetSectorIterator, error) {

	logs, sub, err := _IATIActivitySectors.contract.FilterLogs(opts, "SetSector")
	if err != nil {
		return nil, err
	}
	return &IATIActivitySectorsSetSectorIterator{contract: _IATIActivitySectors.contract, event: "SetSector", logs: logs, sub: sub}, nil
}

// WatchSetSector is a free log subscription operation binding the contract event 0xda3cd6206e5d5a9ca79d812c149f6987f7332ab438d704c76c5c3cc28a0b5a21.
//
// Solidity: event SetSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef, ActivitySectorsSector _sector)
func (_IATIActivitySectors *IATIActivitySectorsFilterer) WatchSetSector(opts *bind.WatchOpts, sink chan<- *IATIActivitySectorsSetSector) (event.Subscription, error) {

	logs, sub, err := _IATIActivitySectors.contract.WatchLogs(opts, "SetSector")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IATIActivitySectorsSetSector)
				if err := _IATIActivitySectors.contract.UnpackLog(event, "SetSector", log); err != nil {
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

// ParseSetSector is a log parse operation binding the contract event 0xda3cd6206e5d5a9ca79d812c149f6987f7332ab438d704c76c5c3cc28a0b5a21.
//
// Solidity: event SetSector(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _sectorRef, ActivitySectorsSector _sector)
func (_IATIActivitySectors *IATIActivitySectorsFilterer) ParseSetSector(log types.Log) (*IATIActivitySectorsSetSector, error) {
	event := new(IATIActivitySectorsSetSector)
	if err := _IATIActivitySectors.contract.UnpackLog(event, "SetSector", log); err != nil {
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
