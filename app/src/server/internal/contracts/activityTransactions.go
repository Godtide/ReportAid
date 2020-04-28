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

// ActivityTransactionsOrg is an auto generated low-level Go binding around an user-defined struct.
type ActivityTransactionsOrg struct {
	OrgType     uint8
	OrgRef      [32]byte
	ActivityRef [32]byte
}

// ActivityTransactionsTransaction is an auto generated low-level Go binding around an user-defined struct.
type ActivityTransactionsTransaction struct {
	TransactionType     uint8
	DisbursementChannel uint8
	FlowType            uint8
	TiedStatus          uint8
	FinanceType         *big.Int
	AidType             [32]byte
	Date                [32]byte
	Value               ActivityTransactionsValue
	ProviderOrg         ActivityTransactionsOrg
	ReceiverOrg         ActivityTransactionsOrg
	SectorDacCode       *big.Int
	Territory           [32]byte
	Description         string
}

// ActivityTransactionsValue is an auto generated low-level Go binding around an user-defined struct.
type ActivityTransactionsValue struct {
	Value    *big.Int
	Date     [32]byte
	Currency [32]byte
}

// ActivityTransactionsABI is the input ABI used to generate the binding from.
const ActivityTransactionsABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getAidType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getDate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getDescription\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getDisbursementChannel\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getFinanceType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getFlowType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getProviderActivityRef\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getProviderOrgRef\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getProviderOrgType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getReceiverActivityRef\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getReceiverOrgRef\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getReceiverOrgType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getSectorDacCode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getTerritory\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getTiedStatus\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"transactionType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"disbursementChannel\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"flowType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tiedStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"financeType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"aidType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTransactions.Value\",\"name\":\"value\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"activityRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTransactions.Org\",\"name\":\"providerOrg\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"activityRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTransactions.Org\",\"name\":\"receiverOrg\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"sectorDacCode\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"territory\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityTransactions.Transaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getTransactionType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getValueCurrency\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getValueDate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"transactionType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"disbursementChannel\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"flowType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tiedStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"financeType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"aidType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTransactions.Value\",\"name\":\"value\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"activityRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTransactions.Org\",\"name\":\"providerOrg\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"activityRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTransactions.Org\",\"name\":\"receiverOrg\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"sectorDacCode\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"territory\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityTransactions.Transaction\",\"name\":\"_transaction\",\"type\":\"tuple\"}],\"name\":\"setTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ActivityTransactionsFuncSigs maps the 4-byte function signature to its string representation.
var ActivityTransactionsFuncSigs = map[string]string{
	"cbf851f7": "getAidType(bytes32,bytes32,bytes32)",
	"b5e2ffb4": "getDate(bytes32,bytes32,bytes32)",
	"b6f5bc3a": "getDescription(bytes32,bytes32,bytes32)",
	"35d652cd": "getDisbursementChannel(bytes32,bytes32,bytes32)",
	"5874928e": "getFinanceType(bytes32,bytes32,bytes32)",
	"fcd02026": "getFlowType(bytes32,bytes32,bytes32)",
	"47cdae01": "getNum(bytes32,bytes32)",
	"cb3f2cd9": "getProviderActivityRef(bytes32,bytes32,bytes32)",
	"48c94885": "getProviderOrgRef(bytes32,bytes32,bytes32)",
	"a57ab1c0": "getProviderOrgType(bytes32,bytes32,bytes32)",
	"0318dffa": "getReceiverActivityRef(bytes32,bytes32,bytes32)",
	"4c9ae37f": "getReceiverOrgRef(bytes32,bytes32,bytes32)",
	"5509ea42": "getReceiverOrgType(bytes32,bytes32,bytes32)",
	"eef67d78": "getReference(bytes32,bytes32,uint256)",
	"5e681587": "getSectorDacCode(bytes32,bytes32,bytes32)",
	"431c989f": "getTerritory(bytes32,bytes32,bytes32)",
	"88a34353": "getTiedStatus(bytes32,bytes32,bytes32)",
	"3f8e14a7": "getTransaction(bytes32,bytes32,bytes32)",
	"8aa632ed": "getTransactionType(bytes32,bytes32,bytes32)",
	"aa80fa45": "getValue(bytes32,bytes32,bytes32)",
	"2cce3224": "getValueCurrency(bytes32,bytes32,bytes32)",
	"ab1de21a": "getValueDate(bytes32,bytes32,bytes32)",
	"3598f212": "setTransaction(bytes32,bytes32,bytes32,(uint8,uint8,uint8,uint8,uint256,bytes32,bytes32,(uint256,bytes32,bytes32),(uint8,bytes32,bytes32),(uint8,bytes32,bytes32),uint256,bytes32,string))",
}

// ActivityTransactions is an auto generated Go binding around an Ethereum contract.
type ActivityTransactions struct {
	ActivityTransactionsCaller     // Read-only binding to the contract
	ActivityTransactionsTransactor // Write-only binding to the contract
	ActivityTransactionsFilterer   // Log filterer for contract events
}

// ActivityTransactionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivityTransactionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityTransactionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivityTransactionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityTransactionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivityTransactionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityTransactionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivityTransactionsSession struct {
	Contract     *ActivityTransactions // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ActivityTransactionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivityTransactionsCallerSession struct {
	Contract *ActivityTransactionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ActivityTransactionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivityTransactionsTransactorSession struct {
	Contract     *ActivityTransactionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ActivityTransactionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivityTransactionsRaw struct {
	Contract *ActivityTransactions // Generic contract binding to access the raw methods on
}

// ActivityTransactionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivityTransactionsCallerRaw struct {
	Contract *ActivityTransactionsCaller // Generic read-only contract binding to access the raw methods on
}

// ActivityTransactionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivityTransactionsTransactorRaw struct {
	Contract *ActivityTransactionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivityTransactions creates a new instance of ActivityTransactions, bound to a specific deployed contract.
func NewActivityTransactions(address common.Address, backend bind.ContractBackend) (*ActivityTransactions, error) {
	contract, err := bindActivityTransactions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ActivityTransactions{ActivityTransactionsCaller: ActivityTransactionsCaller{contract: contract}, ActivityTransactionsTransactor: ActivityTransactionsTransactor{contract: contract}, ActivityTransactionsFilterer: ActivityTransactionsFilterer{contract: contract}}, nil
}

// NewActivityTransactionsCaller creates a new read-only instance of ActivityTransactions, bound to a specific deployed contract.
func NewActivityTransactionsCaller(address common.Address, caller bind.ContractCaller) (*ActivityTransactionsCaller, error) {
	contract, err := bindActivityTransactions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityTransactionsCaller{contract: contract}, nil
}

// NewActivityTransactionsTransactor creates a new write-only instance of ActivityTransactions, bound to a specific deployed contract.
func NewActivityTransactionsTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivityTransactionsTransactor, error) {
	contract, err := bindActivityTransactions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityTransactionsTransactor{contract: contract}, nil
}

// NewActivityTransactionsFilterer creates a new log filterer instance of ActivityTransactions, bound to a specific deployed contract.
func NewActivityTransactionsFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivityTransactionsFilterer, error) {
	contract, err := bindActivityTransactions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivityTransactionsFilterer{contract: contract}, nil
}

// bindActivityTransactions binds a generic wrapper to an already deployed contract.
func bindActivityTransactions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivityTransactionsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityTransactions *ActivityTransactionsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivityTransactions.Contract.ActivityTransactionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityTransactions *ActivityTransactionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityTransactions.Contract.ActivityTransactionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityTransactions *ActivityTransactionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityTransactions.Contract.ActivityTransactionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityTransactions *ActivityTransactionsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivityTransactions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityTransactions *ActivityTransactionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityTransactions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityTransactions *ActivityTransactionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityTransactions.Contract.contract.Transact(opts, method, params...)
}

// GetAidType is a free data retrieval call binding the contract method 0xcbf851f7.
//
// Solidity: function getAidType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCaller) GetAidType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getAidType", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetAidType is a free data retrieval call binding the contract method 0xcbf851f7.
//
// Solidity: function getAidType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsSession) GetAidType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetAidType(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetAidType is a free data retrieval call binding the contract method 0xcbf851f7.
//
// Solidity: function getAidType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetAidType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetAidType(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetDate is a free data retrieval call binding the contract method 0xb5e2ffb4.
//
// Solidity: function getDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCaller) GetDate(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getDate", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetDate is a free data retrieval call binding the contract method 0xb5e2ffb4.
//
// Solidity: function getDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsSession) GetDate(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetDate(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetDate is a free data retrieval call binding the contract method 0xb5e2ffb4.
//
// Solidity: function getDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetDate(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetDate(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetDescription is a free data retrieval call binding the contract method 0xb6f5bc3a.
//
// Solidity: function getDescription(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(string)
func (_ActivityTransactions *ActivityTransactionsCaller) GetDescription(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getDescription", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetDescription is a free data retrieval call binding the contract method 0xb6f5bc3a.
//
// Solidity: function getDescription(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(string)
func (_ActivityTransactions *ActivityTransactionsSession) GetDescription(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (string, error) {
	return _ActivityTransactions.Contract.GetDescription(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetDescription is a free data retrieval call binding the contract method 0xb6f5bc3a.
//
// Solidity: function getDescription(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(string)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetDescription(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (string, error) {
	return _ActivityTransactions.Contract.GetDescription(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetDisbursementChannel is a free data retrieval call binding the contract method 0x35d652cd.
//
// Solidity: function getDisbursementChannel(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_ActivityTransactions *ActivityTransactionsCaller) GetDisbursementChannel(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getDisbursementChannel", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetDisbursementChannel is a free data retrieval call binding the contract method 0x35d652cd.
//
// Solidity: function getDisbursementChannel(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_ActivityTransactions *ActivityTransactionsSession) GetDisbursementChannel(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _ActivityTransactions.Contract.GetDisbursementChannel(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetDisbursementChannel is a free data retrieval call binding the contract method 0x35d652cd.
//
// Solidity: function getDisbursementChannel(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetDisbursementChannel(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _ActivityTransactions.Contract.GetDisbursementChannel(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetFinanceType is a free data retrieval call binding the contract method 0x5874928e.
//
// Solidity: function getFinanceType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint256)
func (_ActivityTransactions *ActivityTransactionsCaller) GetFinanceType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getFinanceType", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetFinanceType is a free data retrieval call binding the contract method 0x5874928e.
//
// Solidity: function getFinanceType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint256)
func (_ActivityTransactions *ActivityTransactionsSession) GetFinanceType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (*big.Int, error) {
	return _ActivityTransactions.Contract.GetFinanceType(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetFinanceType is a free data retrieval call binding the contract method 0x5874928e.
//
// Solidity: function getFinanceType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint256)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetFinanceType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (*big.Int, error) {
	return _ActivityTransactions.Contract.GetFinanceType(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetFlowType is a free data retrieval call binding the contract method 0xfcd02026.
//
// Solidity: function getFlowType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_ActivityTransactions *ActivityTransactionsCaller) GetFlowType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getFlowType", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetFlowType is a free data retrieval call binding the contract method 0xfcd02026.
//
// Solidity: function getFlowType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_ActivityTransactions *ActivityTransactionsSession) GetFlowType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _ActivityTransactions.Contract.GetFlowType(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetFlowType is a free data retrieval call binding the contract method 0xfcd02026.
//
// Solidity: function getFlowType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetFlowType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _ActivityTransactions.Contract.GetFlowType(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivityTransactions *ActivityTransactionsCaller) GetNum(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getNum", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivityTransactions *ActivityTransactionsSession) GetNum(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _ActivityTransactions.Contract.GetNum(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef)
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetNum(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _ActivityTransactions.Contract.GetNum(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef)
}

// GetProviderActivityRef is a free data retrieval call binding the contract method 0xcb3f2cd9.
//
// Solidity: function getProviderActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCaller) GetProviderActivityRef(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getProviderActivityRef", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetProviderActivityRef is a free data retrieval call binding the contract method 0xcb3f2cd9.
//
// Solidity: function getProviderActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsSession) GetProviderActivityRef(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetProviderActivityRef(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetProviderActivityRef is a free data retrieval call binding the contract method 0xcb3f2cd9.
//
// Solidity: function getProviderActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetProviderActivityRef(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetProviderActivityRef(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetProviderOrgRef is a free data retrieval call binding the contract method 0x48c94885.
//
// Solidity: function getProviderOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCaller) GetProviderOrgRef(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getProviderOrgRef", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetProviderOrgRef is a free data retrieval call binding the contract method 0x48c94885.
//
// Solidity: function getProviderOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsSession) GetProviderOrgRef(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetProviderOrgRef(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetProviderOrgRef is a free data retrieval call binding the contract method 0x48c94885.
//
// Solidity: function getProviderOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetProviderOrgRef(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetProviderOrgRef(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetProviderOrgType is a free data retrieval call binding the contract method 0xa57ab1c0.
//
// Solidity: function getProviderOrgType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_ActivityTransactions *ActivityTransactionsCaller) GetProviderOrgType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getProviderOrgType", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetProviderOrgType is a free data retrieval call binding the contract method 0xa57ab1c0.
//
// Solidity: function getProviderOrgType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_ActivityTransactions *ActivityTransactionsSession) GetProviderOrgType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _ActivityTransactions.Contract.GetProviderOrgType(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetProviderOrgType is a free data retrieval call binding the contract method 0xa57ab1c0.
//
// Solidity: function getProviderOrgType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetProviderOrgType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _ActivityTransactions.Contract.GetProviderOrgType(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetReceiverActivityRef is a free data retrieval call binding the contract method 0x0318dffa.
//
// Solidity: function getReceiverActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCaller) GetReceiverActivityRef(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getReceiverActivityRef", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetReceiverActivityRef is a free data retrieval call binding the contract method 0x0318dffa.
//
// Solidity: function getReceiverActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsSession) GetReceiverActivityRef(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetReceiverActivityRef(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetReceiverActivityRef is a free data retrieval call binding the contract method 0x0318dffa.
//
// Solidity: function getReceiverActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetReceiverActivityRef(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetReceiverActivityRef(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetReceiverOrgRef is a free data retrieval call binding the contract method 0x4c9ae37f.
//
// Solidity: function getReceiverOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCaller) GetReceiverOrgRef(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getReceiverOrgRef", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetReceiverOrgRef is a free data retrieval call binding the contract method 0x4c9ae37f.
//
// Solidity: function getReceiverOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsSession) GetReceiverOrgRef(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetReceiverOrgRef(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetReceiverOrgRef is a free data retrieval call binding the contract method 0x4c9ae37f.
//
// Solidity: function getReceiverOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetReceiverOrgRef(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetReceiverOrgRef(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetReceiverOrgType is a free data retrieval call binding the contract method 0x5509ea42.
//
// Solidity: function getReceiverOrgType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_ActivityTransactions *ActivityTransactionsCaller) GetReceiverOrgType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getReceiverOrgType", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetReceiverOrgType is a free data retrieval call binding the contract method 0x5509ea42.
//
// Solidity: function getReceiverOrgType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_ActivityTransactions *ActivityTransactionsSession) GetReceiverOrgType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _ActivityTransactions.Contract.GetReceiverOrgType(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetReceiverOrgType is a free data retrieval call binding the contract method 0x5509ea42.
//
// Solidity: function getReceiverOrgType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetReceiverOrgType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _ActivityTransactions.Contract.GetReceiverOrgType(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCaller) GetReference(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getReference", _activitiesRef, _activityRef, _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetReference(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetReference(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetSectorDacCode is a free data retrieval call binding the contract method 0x5e681587.
//
// Solidity: function getSectorDacCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint256)
func (_ActivityTransactions *ActivityTransactionsCaller) GetSectorDacCode(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getSectorDacCode", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetSectorDacCode is a free data retrieval call binding the contract method 0x5e681587.
//
// Solidity: function getSectorDacCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint256)
func (_ActivityTransactions *ActivityTransactionsSession) GetSectorDacCode(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (*big.Int, error) {
	return _ActivityTransactions.Contract.GetSectorDacCode(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetSectorDacCode is a free data retrieval call binding the contract method 0x5e681587.
//
// Solidity: function getSectorDacCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint256)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetSectorDacCode(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (*big.Int, error) {
	return _ActivityTransactions.Contract.GetSectorDacCode(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetTerritory is a free data retrieval call binding the contract method 0x431c989f.
//
// Solidity: function getTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCaller) GetTerritory(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getTerritory", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetTerritory is a free data retrieval call binding the contract method 0x431c989f.
//
// Solidity: function getTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsSession) GetTerritory(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetTerritory(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetTerritory is a free data retrieval call binding the contract method 0x431c989f.
//
// Solidity: function getTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetTerritory(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetTerritory(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetTiedStatus is a free data retrieval call binding the contract method 0x88a34353.
//
// Solidity: function getTiedStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_ActivityTransactions *ActivityTransactionsCaller) GetTiedStatus(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getTiedStatus", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetTiedStatus is a free data retrieval call binding the contract method 0x88a34353.
//
// Solidity: function getTiedStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_ActivityTransactions *ActivityTransactionsSession) GetTiedStatus(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _ActivityTransactions.Contract.GetTiedStatus(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetTiedStatus is a free data retrieval call binding the contract method 0x88a34353.
//
// Solidity: function getTiedStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetTiedStatus(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _ActivityTransactions.Contract.GetTiedStatus(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetTransaction is a free data retrieval call binding the contract method 0x3f8e14a7.
//
// Solidity: function getTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(ActivityTransactionsTransaction)
func (_ActivityTransactions *ActivityTransactionsCaller) GetTransaction(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (ActivityTransactionsTransaction, error) {
	var (
		ret0 = new(ActivityTransactionsTransaction)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getTransaction", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetTransaction is a free data retrieval call binding the contract method 0x3f8e14a7.
//
// Solidity: function getTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(ActivityTransactionsTransaction)
func (_ActivityTransactions *ActivityTransactionsSession) GetTransaction(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (ActivityTransactionsTransaction, error) {
	return _ActivityTransactions.Contract.GetTransaction(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetTransaction is a free data retrieval call binding the contract method 0x3f8e14a7.
//
// Solidity: function getTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(ActivityTransactionsTransaction)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetTransaction(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (ActivityTransactionsTransaction, error) {
	return _ActivityTransactions.Contract.GetTransaction(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetTransactionType is a free data retrieval call binding the contract method 0x8aa632ed.
//
// Solidity: function getTransactionType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_ActivityTransactions *ActivityTransactionsCaller) GetTransactionType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getTransactionType", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetTransactionType is a free data retrieval call binding the contract method 0x8aa632ed.
//
// Solidity: function getTransactionType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_ActivityTransactions *ActivityTransactionsSession) GetTransactionType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _ActivityTransactions.Contract.GetTransactionType(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetTransactionType is a free data retrieval call binding the contract method 0x8aa632ed.
//
// Solidity: function getTransactionType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetTransactionType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _ActivityTransactions.Contract.GetTransactionType(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetValue is a free data retrieval call binding the contract method 0xaa80fa45.
//
// Solidity: function getValue(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint256)
func (_ActivityTransactions *ActivityTransactionsCaller) GetValue(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getValue", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetValue is a free data retrieval call binding the contract method 0xaa80fa45.
//
// Solidity: function getValue(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint256)
func (_ActivityTransactions *ActivityTransactionsSession) GetValue(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (*big.Int, error) {
	return _ActivityTransactions.Contract.GetValue(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetValue is a free data retrieval call binding the contract method 0xaa80fa45.
//
// Solidity: function getValue(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint256)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetValue(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (*big.Int, error) {
	return _ActivityTransactions.Contract.GetValue(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetValueCurrency is a free data retrieval call binding the contract method 0x2cce3224.
//
// Solidity: function getValueCurrency(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCaller) GetValueCurrency(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getValueCurrency", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetValueCurrency is a free data retrieval call binding the contract method 0x2cce3224.
//
// Solidity: function getValueCurrency(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsSession) GetValueCurrency(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetValueCurrency(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetValueCurrency is a free data retrieval call binding the contract method 0x2cce3224.
//
// Solidity: function getValueCurrency(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetValueCurrency(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetValueCurrency(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetValueDate is a free data retrieval call binding the contract method 0xab1de21a.
//
// Solidity: function getValueDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCaller) GetValueDate(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityTransactions.contract.Call(opts, out, "getValueDate", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetValueDate is a free data retrieval call binding the contract method 0xab1de21a.
//
// Solidity: function getValueDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsSession) GetValueDate(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetValueDate(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetValueDate is a free data retrieval call binding the contract method 0xab1de21a.
//
// Solidity: function getValueDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_ActivityTransactions *ActivityTransactionsCallerSession) GetValueDate(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _ActivityTransactions.Contract.GetValueDate(&_ActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// SetTransaction is a paid mutator transaction binding the contract method 0x3598f212.
//
// Solidity: function setTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef, ActivityTransactionsTransaction _transaction) returns()
func (_ActivityTransactions *ActivityTransactionsTransactor) SetTransaction(opts *bind.TransactOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte, _transaction ActivityTransactionsTransaction) (*types.Transaction, error) {
	return _ActivityTransactions.contract.Transact(opts, "setTransaction", _activitiesRef, _activityRef, _transactionRef, _transaction)
}

// SetTransaction is a paid mutator transaction binding the contract method 0x3598f212.
//
// Solidity: function setTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef, ActivityTransactionsTransaction _transaction) returns()
func (_ActivityTransactions *ActivityTransactionsSession) SetTransaction(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte, _transaction ActivityTransactionsTransaction) (*types.Transaction, error) {
	return _ActivityTransactions.Contract.SetTransaction(&_ActivityTransactions.TransactOpts, _activitiesRef, _activityRef, _transactionRef, _transaction)
}

// SetTransaction is a paid mutator transaction binding the contract method 0x3598f212.
//
// Solidity: function setTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef, ActivityTransactionsTransaction _transaction) returns()
func (_ActivityTransactions *ActivityTransactionsTransactorSession) SetTransaction(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte, _transaction ActivityTransactionsTransaction) (*types.Transaction, error) {
	return _ActivityTransactions.Contract.SetTransaction(&_ActivityTransactions.TransactOpts, _activitiesRef, _activityRef, _transactionRef, _transaction)
}

// IATIActivityTransactionsABI is the input ABI used to generate the binding from.
const IATIActivityTransactionsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"transactionType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"disbursementChannel\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"flowType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tiedStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"financeType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"aidType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTransactions.Value\",\"name\":\"value\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"activityRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTransactions.Org\",\"name\":\"providerOrg\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"activityRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTransactions.Org\",\"name\":\"receiverOrg\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"sectorDacCode\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"territory\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structActivityTransactions.Transaction\",\"name\":\"_transaction\",\"type\":\"tuple\"}],\"name\":\"SetTransaction\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getAidType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getDate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getDescription\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getDisbursementChannel\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getFinanceType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getFlowType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"}],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getProviderActivityRef\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getProviderOrgRef\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getProviderOrgType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getReceiverActivityRef\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getReceiverOrgRef\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getReceiverOrgType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getSectorDacCode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getTerritory\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getTiedStatus\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"transactionType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"disbursementChannel\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"flowType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tiedStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"financeType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"aidType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTransactions.Value\",\"name\":\"value\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"activityRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTransactions.Org\",\"name\":\"providerOrg\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"activityRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTransactions.Org\",\"name\":\"receiverOrg\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"sectorDacCode\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"territory\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityTransactions.Transaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getTransactionType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getValueCurrency\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"}],\"name\":\"getValueDate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_activitiesRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_activityRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_transactionRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"transactionType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"disbursementChannel\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"flowType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tiedStatus\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"financeType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"aidType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTransactions.Value\",\"name\":\"value\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"activityRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTransactions.Org\",\"name\":\"providerOrg\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"activityRef\",\"type\":\"bytes32\"}],\"internalType\":\"structActivityTransactions.Org\",\"name\":\"receiverOrg\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"sectorDacCode\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"territory\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityTransactions.Transaction\",\"name\":\"_transaction\",\"type\":\"tuple\"}],\"name\":\"setTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IATIActivityTransactionsFuncSigs maps the 4-byte function signature to its string representation.
var IATIActivityTransactionsFuncSigs = map[string]string{
	"cbf851f7": "getAidType(bytes32,bytes32,bytes32)",
	"b5e2ffb4": "getDate(bytes32,bytes32,bytes32)",
	"b6f5bc3a": "getDescription(bytes32,bytes32,bytes32)",
	"35d652cd": "getDisbursementChannel(bytes32,bytes32,bytes32)",
	"5874928e": "getFinanceType(bytes32,bytes32,bytes32)",
	"fcd02026": "getFlowType(bytes32,bytes32,bytes32)",
	"47cdae01": "getNum(bytes32,bytes32)",
	"cb3f2cd9": "getProviderActivityRef(bytes32,bytes32,bytes32)",
	"48c94885": "getProviderOrgRef(bytes32,bytes32,bytes32)",
	"a57ab1c0": "getProviderOrgType(bytes32,bytes32,bytes32)",
	"0318dffa": "getReceiverActivityRef(bytes32,bytes32,bytes32)",
	"4c9ae37f": "getReceiverOrgRef(bytes32,bytes32,bytes32)",
	"5509ea42": "getReceiverOrgType(bytes32,bytes32,bytes32)",
	"eef67d78": "getReference(bytes32,bytes32,uint256)",
	"5e681587": "getSectorDacCode(bytes32,bytes32,bytes32)",
	"431c989f": "getTerritory(bytes32,bytes32,bytes32)",
	"88a34353": "getTiedStatus(bytes32,bytes32,bytes32)",
	"3f8e14a7": "getTransaction(bytes32,bytes32,bytes32)",
	"8aa632ed": "getTransactionType(bytes32,bytes32,bytes32)",
	"aa80fa45": "getValue(bytes32,bytes32,bytes32)",
	"2cce3224": "getValueCurrency(bytes32,bytes32,bytes32)",
	"ab1de21a": "getValueDate(bytes32,bytes32,bytes32)",
	"3598f212": "setTransaction(bytes32,bytes32,bytes32,(uint8,uint8,uint8,uint8,uint256,bytes32,bytes32,(uint256,bytes32,bytes32),(uint8,bytes32,bytes32),(uint8,bytes32,bytes32),uint256,bytes32,string))",
}

// IATIActivityTransactionsBin is the compiled bytecode used for deploying new contracts.
var IATIActivityTransactionsBin = "0x608060405234801561001057600080fd5b50611a83806100206000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c80635e681587116100c3578063b5e2ffb41161007c578063b5e2ffb4146102c7578063b6f5bc3a146102da578063cb3f2cd9146102fa578063cbf851f71461030d578063eef67d7814610320578063fcd02026146103335761014d565b80635e6815871461025557806388a34353146102685780638aa632ed1461027b578063a57ab1c01461028e578063aa80fa45146102a1578063ab1de21a146102b45761014d565b8063431c989f11610115578063431c989f146101e357806347cdae01146101f657806348c94885146102095780634c9ae37f1461021c5780635509ea421461022f5780635874928e146102425761014d565b80630318dffa146101525780632cce32241461017b5780633598f2121461018e57806335d652cd146101a35780633f8e14a7146101c3575b600080fd5b610165610160366004611681565b610346565b604051610172919061196a565b60405180910390f35b610165610189366004611681565b6103c1565b6101a161019c3660046116ac565b61043c565b005b6101b66101b1366004611681565b610859565b6040516101729190611a18565b6101d66101d1366004611681565b6108d9565b6040516101729190611a05565b6101656101f1366004611681565b610ad4565b610165610204366004611660565b610b4f565b610165610217366004611681565b610ba7565b61016561022a366004611681565b610c22565b6101b661023d366004611681565b610c9d565b610165610250366004611681565b610d1b565b610165610263366004611681565b610d97565b6101b6610276366004611681565b610e12565b6101b6610289366004611681565b610e94565b6101b661029c366004611681565b610f0f565b6101656102af366004611681565b610f8d565b6101656102c2366004611681565b611008565b6101656102d5366004611681565b611083565b6102ed6102e8366004611681565b6110fe565b60405161017291906119f2565b610165610308366004611681565b611205565b61016561031b366004611681565b611280565b61016561032e3660046117f1565b6112fb565b6101b6610341366004611681565b61138c565b600083811a60f81b6001600160f81b0319161580159061037557508260001a60f81b6001600160f81b03191615155b801561039057508160001a60f81b6001600160f81b03191615155b61039957600080fd5b506000928352600160209081526040808520938552928152828420918452529020600c015490565b600083811a60f81b6001600160f81b031916158015906103f057508260001a60f81b6001600160f81b03191615155b801561040b57508160001a60f81b6001600160f81b03191615155b61041457600080fd5b5060009283526001602090815260408085209385529281528284209184525290206006015490565b8360001a60f81b6001600160f81b0319161580159061046a57508260001a60f81b6001600160f81b03191615155b801561048557508160001a60f81b6001600160f81b03191615155b80156104945750805160ff1615155b80156104a657508051600e60ff909116105b80156104b85750602081015160ff1615155b80156104ce5750600560ff16816020015160ff16105b80156104ed575060c081015160001a60f81b6001600160f81b03191615155b8015610510575060e08101516020015160001a60f81b6001600160f81b03191615155b8015610533575060e08101516040015160001a60f81b6001600160f81b03191615155b801561054b575061010081015151600a60ff90911610155b801561056f57506101008101516020015160001a60f81b6001600160f81b03191615155b8015610587575061012081015151600a60ff90911610155b80156105ab57506101208101516020015160001a60f81b6001600160f81b03191615155b80156105be5750612b6681610140015110155b80156105d25750620185ec81610140015111155b80156105f2575061016081015160001a60f81b6001600160f81b03191615155b80156106045750600081610180015151115b61060d57600080fd5b60008481526001602081815260408084208785528252808420868552825292839020845181548684015187870151606089015160ff1993841660ff9586161761ff00191661010093861684021762ff0000191662010000928616929092029190911763ff000000191663010000009185169190910217845560808801519584019590955560a0870151600284015560c0870151600384015560e0870151805160048501558085015160058501558601516006840155938601518051600784018054871691841691909117905580840151600884015585015160098301556101208601518051600a8401805490961692169190911790935582820151600b8201559190920151600c820155610140830151600d820155610160830151600e8201556101808301518051849361074892600f85019291019061140d565b50505060008481526020818152604080832086845290915290819020905163745fca7b60e01b815273__$b620240993a55615b607189176fb82e62f$__9163745fca7b9161079a918691600401611973565b60206040518083038186803b1580156107b257600080fd5b505af41580156107c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107ea9190611639565b610816576000848152602081815260408083208684528252822080546001810182559083529120018290555b7f8643d357dcd4c137b20a768f0f133ca5b6b57ce5aeea49bee8e81108018c6c6c8484848460405161084b94939291906119c3565b60405180910390a150505050565b600083811a60f81b6001600160f81b0319161580159061088857508260001a60f81b6001600160f81b03191615155b80156108a357508160001a60f81b6001600160f81b03191615155b6108ac57600080fd5b50600092835260016020908152604080852093855292815282842091845252902054610100900460ff1690565b6108e161148b565b8360001a60f81b6001600160f81b0319161580159061090f57508260001a60f81b6001600160f81b03191615155b801561092a57508160001a60f81b6001600160f81b03191615155b61093357600080fd5b6000848152600160208181526040808420878552825280842086855282529283902083516101a081018552815460ff80821683526101008083048216848701526201000083048216848901526301000000909204811660608085019190915284870154608085015260028086015460a0860152600386015460c086015288518083018a52600487015481526005870154818901526006870154818b015260e086015288518083018a526007870154841681526008870154818901526009870154818b01528486015288519182018952600a8601549092168152600b85015481870152600c85015481890152610120840152600d840154610140840152600e840154610160840152600f840180548851978116159093026000190190921604601f810185900485028601850190965285855290949193610180860193909290830182828015610ac25780601f10610a9757610100808354040283529160200191610ac2565b820191906000526020600020905b815481529060010190602001808311610aa557829003601f168201915b50505050508152505090509392505050565b600083811a60f81b6001600160f81b03191615801590610b0357508260001a60f81b6001600160f81b03191615155b8015610b1e57508160001a60f81b6001600160f81b03191615155b610b2757600080fd5b506000928352600160209081526040808520938552928152828420918452529020600e015490565b600082811a60f81b6001600160f81b03191615801590610b7e57508160001a60f81b6001600160f81b03191615155b610b8757600080fd5b506000828152602081815260408083208484529091529020545b92915050565b600083811a60f81b6001600160f81b03191615801590610bd657508260001a60f81b6001600160f81b03191615155b8015610bf157508160001a60f81b6001600160f81b03191615155b610bfa57600080fd5b5060009283526001602090815260408085209385529281528284209184525290206008015490565b600083811a60f81b6001600160f81b03191615801590610c5157508260001a60f81b6001600160f81b03191615155b8015610c6c57508160001a60f81b6001600160f81b03191615155b610c7557600080fd5b506000928352600160209081526040808520938552928152828420918452529020600b015490565b600083811a60f81b6001600160f81b03191615801590610ccc57508260001a60f81b6001600160f81b03191615155b8015610ce757508160001a60f81b6001600160f81b03191615155b610cf057600080fd5b506000928352600160209081526040808520938552928152828420918452529020600a015460ff1690565b600083811a60f81b6001600160f81b03191615801590610d4a57508260001a60f81b6001600160f81b03191615155b8015610d6557508160001a60f81b6001600160f81b03191615155b610d6e57600080fd5b506000928352600160208181526040808620948652938152838520928552919091529120015490565b600083811a60f81b6001600160f81b03191615801590610dc657508260001a60f81b6001600160f81b03191615155b8015610de157508160001a60f81b6001600160f81b03191615155b610dea57600080fd5b506000928352600160209081526040808520938552928152828420918452529020600d015490565b600083811a60f81b6001600160f81b03191615801590610e4157508260001a60f81b6001600160f81b03191615155b8015610e5c57508160001a60f81b6001600160f81b03191615155b610e6557600080fd5b506000928352600160209081526040808520938552928152828420918452529020546301000000900460ff1690565b600083811a60f81b6001600160f81b03191615801590610ec357508260001a60f81b6001600160f81b03191615155b8015610ede57508160001a60f81b6001600160f81b03191615155b610ee757600080fd5b5060009283526001602090815260408085209385529281528284209184525290205460ff1690565b600083811a60f81b6001600160f81b03191615801590610f3e57508260001a60f81b6001600160f81b03191615155b8015610f5957508160001a60f81b6001600160f81b03191615155b610f6257600080fd5b5060009283526001602090815260408085209385529281528284209184525290206007015460ff1690565b600083811a60f81b6001600160f81b03191615801590610fbc57508260001a60f81b6001600160f81b03191615155b8015610fd757508160001a60f81b6001600160f81b03191615155b610fe057600080fd5b5060009283526001602090815260408085209385529281528284209184525290206004015490565b600083811a60f81b6001600160f81b0319161580159061103757508260001a60f81b6001600160f81b03191615155b801561105257508160001a60f81b6001600160f81b03191615155b61105b57600080fd5b5060009283526001602090815260408085209385529281528284209184525290206005015490565b600083811a60f81b6001600160f81b031916158015906110b257508260001a60f81b6001600160f81b03191615155b80156110cd57508160001a60f81b6001600160f81b03191615155b6110d657600080fd5b5060009283526001602090815260408085209385529281528284209184525290206003015490565b60608360001a60f81b6001600160f81b0319161580159061112e57508260001a60f81b6001600160f81b03191615155b801561114957508160001a60f81b6001600160f81b03191615155b61115257600080fd5b60008481526001602081815260408084208785528252808420868552825292839020600f0180548451600294821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156111f75780601f106111cc576101008083540402835291602001916111f7565b820191906000526020600020905b8154815290600101906020018083116111da57829003601f168201915b505050505090509392505050565b600083811a60f81b6001600160f81b0319161580159061123457508260001a60f81b6001600160f81b03191615155b801561124f57508160001a60f81b6001600160f81b03191615155b61125857600080fd5b5060009283526001602090815260408085209385529281528284209184525290206009015490565b600083811a60f81b6001600160f81b031916158015906112af57508260001a60f81b6001600160f81b03191615155b80156112ca57508160001a60f81b6001600160f81b03191615155b6112d357600080fd5b5060009283526001602090815260408085209385529281528284209184525290206002015490565b600083811a60f81b6001600160f81b0319161580159061132a57508260001a60f81b6001600160f81b03191615155b801561134c575060008481526020818152604080832086845290915290205482105b61135557600080fd5b600084815260208181526040808320868452909152902080548390811061137857fe5b906000526020600020015490509392505050565b600083811a60f81b6001600160f81b031916158015906113bb57508260001a60f81b6001600160f81b03191615155b80156113d657508160001a60f81b6001600160f81b03191615155b6113df57600080fd5b5060009283526001602090815260408085209385529281528284209184525290205462010000900460ff1690565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061144e57805160ff191683800117855561147b565b8280016001018555821561147b579182015b8281111561147b578251825591602001919060010190611460565b50611487929150611504565b5090565b604080516101a081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c081019190915260e081016114d1611521565b81526020016114de611521565b81526020016114eb611521565b8152600060208201819052604082015260609081015290565b61151e91905b80821115611487576000815560010161150a565b90565b604080516060810182526000808252602082018190529181019190915290565b600082601f830112611551578081fd5b813567ffffffffffffffff811115611567578182fd5b61157a601f8201601f1916602001611a26565b915080825283602082850101111561159157600080fd5b8060208401602084013760009082016020015292915050565b6000606082840312156115bb578081fd5b6115c56060611a26565b90506115d18383611628565b8152602082013560208201526040820135604082015292915050565b6000606082840312156115fe578081fd5b6116086060611a26565b905081358152602082013560208201526040820135604082015292915050565b803560ff81168114610ba157600080fd5b60006020828403121561164a578081fd5b81518015158114611659578182fd5b9392505050565b60008060408385031215611672578081fd5b50508035926020909101359150565b600080600060608486031215611695578081fd5b505081359360208301359350604090920135919050565b600080600080608085870312156116c1578081fd5b843593506020850135925060408501359150606085013567ffffffffffffffff808211156116ed578283fd5b818701610260818a031215611700578384fd5b6101a0925061170e83611a26565b6117188a83611628565b81526117278a60208401611628565b60208201526117398a60408401611628565b604082015261174b8a60608401611628565b60608201526080820135608082015260a082013560a082015260c082013560c082015261177b8a60e084016115ed565b60e082015261014061178f8b8285016115aa565b6101008301526117a18b8685016115aa565b610120830152610200830135908201526102208201356101608201526102408201359350828411156117d1578485fd5b6117dd8a858401611541565b610180820152969995985093965050505050565b600080600060608486031215611695578283fd5b60008151808452815b8181101561182a5760208185018101518683018201520161180e565b8181111561183b5782602083870101525b50601f01601f19169290920160200192915050565b805160ff16825260208082015190830152604090810151910152565b600061026061187c848451611963565b602083015161188e6020860182611963565b5060408301516118a16040860182611963565b5060608301516118b46060860182611963565b506080830151608085015260a083015160a085015260c083015160c085015260e08301516118e560e086018261194a565b506101008301516101406118fb81870183611850565b61012085015191506119116101a0870183611850565b84015161020086015250610160830151610220850152610180830151610240850182905261194182860182611805565b95945050505050565b8051825260208082015190830152604090810151910152565b60ff169052565b90815260200190565b60006040820184835260206040818501528185548084526060860191508685528285209350845b818110156119b65784548352600194850194928401920161199a565b5090979650505050505050565b6000858252846020830152836040830152608060608301526119e8608083018461186c565b9695505050505050565b6000602082526116596020830184611805565b600060208252611659602083018461186c565b60ff91909116815260200190565b60405181810167ffffffffffffffff81118282101715611a4557600080fd5b60405291905056fea2646970667358221220dfcf34a24aad90572676d5b6297843531148863419b1276932de14a8ee46450e64736f6c63430006060033"

// DeployIATIActivityTransactions deploys a new Ethereum contract, binding an instance of IATIActivityTransactions to it.
func DeployIATIActivityTransactions(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IATIActivityTransactions, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivityTransactionsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	stringsAddr, _, _, _ := DeployStrings(auth, backend)
	IATIActivityTransactionsBin = strings.Replace(IATIActivityTransactionsBin, "__$b620240993a55615b607189176fb82e62f$__", stringsAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IATIActivityTransactionsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IATIActivityTransactions{IATIActivityTransactionsCaller: IATIActivityTransactionsCaller{contract: contract}, IATIActivityTransactionsTransactor: IATIActivityTransactionsTransactor{contract: contract}, IATIActivityTransactionsFilterer: IATIActivityTransactionsFilterer{contract: contract}}, nil
}

// IATIActivityTransactions is an auto generated Go binding around an Ethereum contract.
type IATIActivityTransactions struct {
	IATIActivityTransactionsCaller     // Read-only binding to the contract
	IATIActivityTransactionsTransactor // Write-only binding to the contract
	IATIActivityTransactionsFilterer   // Log filterer for contract events
}

// IATIActivityTransactionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IATIActivityTransactionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityTransactionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IATIActivityTransactionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityTransactionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IATIActivityTransactionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIActivityTransactionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IATIActivityTransactionsSession struct {
	Contract     *IATIActivityTransactions // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IATIActivityTransactionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IATIActivityTransactionsCallerSession struct {
	Contract *IATIActivityTransactionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IATIActivityTransactionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IATIActivityTransactionsTransactorSession struct {
	Contract     *IATIActivityTransactionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IATIActivityTransactionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IATIActivityTransactionsRaw struct {
	Contract *IATIActivityTransactions // Generic contract binding to access the raw methods on
}

// IATIActivityTransactionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IATIActivityTransactionsCallerRaw struct {
	Contract *IATIActivityTransactionsCaller // Generic read-only contract binding to access the raw methods on
}

// IATIActivityTransactionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IATIActivityTransactionsTransactorRaw struct {
	Contract *IATIActivityTransactionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIATIActivityTransactions creates a new instance of IATIActivityTransactions, bound to a specific deployed contract.
func NewIATIActivityTransactions(address common.Address, backend bind.ContractBackend) (*IATIActivityTransactions, error) {
	contract, err := bindIATIActivityTransactions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IATIActivityTransactions{IATIActivityTransactionsCaller: IATIActivityTransactionsCaller{contract: contract}, IATIActivityTransactionsTransactor: IATIActivityTransactionsTransactor{contract: contract}, IATIActivityTransactionsFilterer: IATIActivityTransactionsFilterer{contract: contract}}, nil
}

// NewIATIActivityTransactionsCaller creates a new read-only instance of IATIActivityTransactions, bound to a specific deployed contract.
func NewIATIActivityTransactionsCaller(address common.Address, caller bind.ContractCaller) (*IATIActivityTransactionsCaller, error) {
	contract, err := bindIATIActivityTransactions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IATIActivityTransactionsCaller{contract: contract}, nil
}

// NewIATIActivityTransactionsTransactor creates a new write-only instance of IATIActivityTransactions, bound to a specific deployed contract.
func NewIATIActivityTransactionsTransactor(address common.Address, transactor bind.ContractTransactor) (*IATIActivityTransactionsTransactor, error) {
	contract, err := bindIATIActivityTransactions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IATIActivityTransactionsTransactor{contract: contract}, nil
}

// NewIATIActivityTransactionsFilterer creates a new log filterer instance of IATIActivityTransactions, bound to a specific deployed contract.
func NewIATIActivityTransactionsFilterer(address common.Address, filterer bind.ContractFilterer) (*IATIActivityTransactionsFilterer, error) {
	contract, err := bindIATIActivityTransactions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IATIActivityTransactionsFilterer{contract: contract}, nil
}

// bindIATIActivityTransactions binds a generic wrapper to an already deployed contract.
func bindIATIActivityTransactions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIActivityTransactionsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIActivityTransactions *IATIActivityTransactionsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIActivityTransactions.Contract.IATIActivityTransactionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIActivityTransactions *IATIActivityTransactionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIActivityTransactions.Contract.IATIActivityTransactionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIActivityTransactions *IATIActivityTransactionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIActivityTransactions.Contract.IATIActivityTransactionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIActivityTransactions *IATIActivityTransactionsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIActivityTransactions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIActivityTransactions *IATIActivityTransactionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIActivityTransactions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIActivityTransactions *IATIActivityTransactionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIActivityTransactions.Contract.contract.Transact(opts, method, params...)
}

// GetAidType is a free data retrieval call binding the contract method 0xcbf851f7.
//
// Solidity: function getAidType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetAidType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getAidType", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetAidType is a free data retrieval call binding the contract method 0xcbf851f7.
//
// Solidity: function getAidType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetAidType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetAidType(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetAidType is a free data retrieval call binding the contract method 0xcbf851f7.
//
// Solidity: function getAidType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetAidType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetAidType(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetDate is a free data retrieval call binding the contract method 0xb5e2ffb4.
//
// Solidity: function getDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetDate(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getDate", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetDate is a free data retrieval call binding the contract method 0xb5e2ffb4.
//
// Solidity: function getDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetDate(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetDate(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetDate is a free data retrieval call binding the contract method 0xb5e2ffb4.
//
// Solidity: function getDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetDate(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetDate(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetDescription is a free data retrieval call binding the contract method 0xb6f5bc3a.
//
// Solidity: function getDescription(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(string)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetDescription(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getDescription", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetDescription is a free data retrieval call binding the contract method 0xb6f5bc3a.
//
// Solidity: function getDescription(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(string)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetDescription(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (string, error) {
	return _IATIActivityTransactions.Contract.GetDescription(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetDescription is a free data retrieval call binding the contract method 0xb6f5bc3a.
//
// Solidity: function getDescription(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(string)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetDescription(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (string, error) {
	return _IATIActivityTransactions.Contract.GetDescription(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetDisbursementChannel is a free data retrieval call binding the contract method 0x35d652cd.
//
// Solidity: function getDisbursementChannel(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetDisbursementChannel(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getDisbursementChannel", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetDisbursementChannel is a free data retrieval call binding the contract method 0x35d652cd.
//
// Solidity: function getDisbursementChannel(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetDisbursementChannel(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _IATIActivityTransactions.Contract.GetDisbursementChannel(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetDisbursementChannel is a free data retrieval call binding the contract method 0x35d652cd.
//
// Solidity: function getDisbursementChannel(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetDisbursementChannel(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _IATIActivityTransactions.Contract.GetDisbursementChannel(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetFinanceType is a free data retrieval call binding the contract method 0x5874928e.
//
// Solidity: function getFinanceType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint256)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetFinanceType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getFinanceType", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetFinanceType is a free data retrieval call binding the contract method 0x5874928e.
//
// Solidity: function getFinanceType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint256)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetFinanceType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (*big.Int, error) {
	return _IATIActivityTransactions.Contract.GetFinanceType(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetFinanceType is a free data retrieval call binding the contract method 0x5874928e.
//
// Solidity: function getFinanceType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint256)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetFinanceType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (*big.Int, error) {
	return _IATIActivityTransactions.Contract.GetFinanceType(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetFlowType is a free data retrieval call binding the contract method 0xfcd02026.
//
// Solidity: function getFlowType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetFlowType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getFlowType", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetFlowType is a free data retrieval call binding the contract method 0xfcd02026.
//
// Solidity: function getFlowType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetFlowType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _IATIActivityTransactions.Contract.GetFlowType(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetFlowType is a free data retrieval call binding the contract method 0xfcd02026.
//
// Solidity: function getFlowType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetFlowType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _IATIActivityTransactions.Contract.GetFlowType(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetNum(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getNum", _activitiesRef, _activityRef)
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetNum(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _IATIActivityTransactions.Contract.GetNum(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef)
}

// GetNum is a free data retrieval call binding the contract method 0x47cdae01.
//
// Solidity: function getNum(bytes32 _activitiesRef, bytes32 _activityRef) view returns(uint256)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetNum(_activitiesRef [32]byte, _activityRef [32]byte) (*big.Int, error) {
	return _IATIActivityTransactions.Contract.GetNum(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef)
}

// GetProviderActivityRef is a free data retrieval call binding the contract method 0xcb3f2cd9.
//
// Solidity: function getProviderActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetProviderActivityRef(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getProviderActivityRef", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetProviderActivityRef is a free data retrieval call binding the contract method 0xcb3f2cd9.
//
// Solidity: function getProviderActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetProviderActivityRef(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetProviderActivityRef(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetProviderActivityRef is a free data retrieval call binding the contract method 0xcb3f2cd9.
//
// Solidity: function getProviderActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetProviderActivityRef(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetProviderActivityRef(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetProviderOrgRef is a free data retrieval call binding the contract method 0x48c94885.
//
// Solidity: function getProviderOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetProviderOrgRef(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getProviderOrgRef", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetProviderOrgRef is a free data retrieval call binding the contract method 0x48c94885.
//
// Solidity: function getProviderOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetProviderOrgRef(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetProviderOrgRef(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetProviderOrgRef is a free data retrieval call binding the contract method 0x48c94885.
//
// Solidity: function getProviderOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetProviderOrgRef(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetProviderOrgRef(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetProviderOrgType is a free data retrieval call binding the contract method 0xa57ab1c0.
//
// Solidity: function getProviderOrgType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetProviderOrgType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getProviderOrgType", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetProviderOrgType is a free data retrieval call binding the contract method 0xa57ab1c0.
//
// Solidity: function getProviderOrgType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetProviderOrgType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _IATIActivityTransactions.Contract.GetProviderOrgType(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetProviderOrgType is a free data retrieval call binding the contract method 0xa57ab1c0.
//
// Solidity: function getProviderOrgType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetProviderOrgType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _IATIActivityTransactions.Contract.GetProviderOrgType(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetReceiverActivityRef is a free data retrieval call binding the contract method 0x0318dffa.
//
// Solidity: function getReceiverActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetReceiverActivityRef(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getReceiverActivityRef", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetReceiverActivityRef is a free data retrieval call binding the contract method 0x0318dffa.
//
// Solidity: function getReceiverActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetReceiverActivityRef(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetReceiverActivityRef(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetReceiverActivityRef is a free data retrieval call binding the contract method 0x0318dffa.
//
// Solidity: function getReceiverActivityRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetReceiverActivityRef(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetReceiverActivityRef(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetReceiverOrgRef is a free data retrieval call binding the contract method 0x4c9ae37f.
//
// Solidity: function getReceiverOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetReceiverOrgRef(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getReceiverOrgRef", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetReceiverOrgRef is a free data retrieval call binding the contract method 0x4c9ae37f.
//
// Solidity: function getReceiverOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetReceiverOrgRef(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetReceiverOrgRef(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetReceiverOrgRef is a free data retrieval call binding the contract method 0x4c9ae37f.
//
// Solidity: function getReceiverOrgRef(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetReceiverOrgRef(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetReceiverOrgRef(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetReceiverOrgType is a free data retrieval call binding the contract method 0x5509ea42.
//
// Solidity: function getReceiverOrgType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetReceiverOrgType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getReceiverOrgType", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetReceiverOrgType is a free data retrieval call binding the contract method 0x5509ea42.
//
// Solidity: function getReceiverOrgType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetReceiverOrgType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _IATIActivityTransactions.Contract.GetReceiverOrgType(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetReceiverOrgType is a free data retrieval call binding the contract method 0x5509ea42.
//
// Solidity: function getReceiverOrgType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetReceiverOrgType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _IATIActivityTransactions.Contract.GetReceiverOrgType(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetReference(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getReference", _activitiesRef, _activityRef, _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetReference(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xeef67d78.
//
// Solidity: function getReference(bytes32 _activitiesRef, bytes32 _activityRef, uint256 _index) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetReference(_activitiesRef [32]byte, _activityRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetReference(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _index)
}

// GetSectorDacCode is a free data retrieval call binding the contract method 0x5e681587.
//
// Solidity: function getSectorDacCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint256)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetSectorDacCode(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getSectorDacCode", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetSectorDacCode is a free data retrieval call binding the contract method 0x5e681587.
//
// Solidity: function getSectorDacCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint256)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetSectorDacCode(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (*big.Int, error) {
	return _IATIActivityTransactions.Contract.GetSectorDacCode(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetSectorDacCode is a free data retrieval call binding the contract method 0x5e681587.
//
// Solidity: function getSectorDacCode(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint256)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetSectorDacCode(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (*big.Int, error) {
	return _IATIActivityTransactions.Contract.GetSectorDacCode(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetTerritory is a free data retrieval call binding the contract method 0x431c989f.
//
// Solidity: function getTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetTerritory(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getTerritory", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetTerritory is a free data retrieval call binding the contract method 0x431c989f.
//
// Solidity: function getTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetTerritory(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetTerritory(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetTerritory is a free data retrieval call binding the contract method 0x431c989f.
//
// Solidity: function getTerritory(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetTerritory(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetTerritory(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetTiedStatus is a free data retrieval call binding the contract method 0x88a34353.
//
// Solidity: function getTiedStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetTiedStatus(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getTiedStatus", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetTiedStatus is a free data retrieval call binding the contract method 0x88a34353.
//
// Solidity: function getTiedStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetTiedStatus(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _IATIActivityTransactions.Contract.GetTiedStatus(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetTiedStatus is a free data retrieval call binding the contract method 0x88a34353.
//
// Solidity: function getTiedStatus(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetTiedStatus(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _IATIActivityTransactions.Contract.GetTiedStatus(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetTransaction is a free data retrieval call binding the contract method 0x3f8e14a7.
//
// Solidity: function getTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(ActivityTransactionsTransaction)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetTransaction(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (ActivityTransactionsTransaction, error) {
	var (
		ret0 = new(ActivityTransactionsTransaction)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getTransaction", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetTransaction is a free data retrieval call binding the contract method 0x3f8e14a7.
//
// Solidity: function getTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(ActivityTransactionsTransaction)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetTransaction(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (ActivityTransactionsTransaction, error) {
	return _IATIActivityTransactions.Contract.GetTransaction(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetTransaction is a free data retrieval call binding the contract method 0x3f8e14a7.
//
// Solidity: function getTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(ActivityTransactionsTransaction)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetTransaction(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (ActivityTransactionsTransaction, error) {
	return _IATIActivityTransactions.Contract.GetTransaction(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetTransactionType is a free data retrieval call binding the contract method 0x8aa632ed.
//
// Solidity: function getTransactionType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetTransactionType(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getTransactionType", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetTransactionType is a free data retrieval call binding the contract method 0x8aa632ed.
//
// Solidity: function getTransactionType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetTransactionType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _IATIActivityTransactions.Contract.GetTransactionType(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetTransactionType is a free data retrieval call binding the contract method 0x8aa632ed.
//
// Solidity: function getTransactionType(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint8)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetTransactionType(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (uint8, error) {
	return _IATIActivityTransactions.Contract.GetTransactionType(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetValue is a free data retrieval call binding the contract method 0xaa80fa45.
//
// Solidity: function getValue(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint256)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetValue(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getValue", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetValue is a free data retrieval call binding the contract method 0xaa80fa45.
//
// Solidity: function getValue(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint256)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetValue(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (*big.Int, error) {
	return _IATIActivityTransactions.Contract.GetValue(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetValue is a free data retrieval call binding the contract method 0xaa80fa45.
//
// Solidity: function getValue(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(uint256)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetValue(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) (*big.Int, error) {
	return _IATIActivityTransactions.Contract.GetValue(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetValueCurrency is a free data retrieval call binding the contract method 0x2cce3224.
//
// Solidity: function getValueCurrency(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetValueCurrency(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getValueCurrency", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetValueCurrency is a free data retrieval call binding the contract method 0x2cce3224.
//
// Solidity: function getValueCurrency(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetValueCurrency(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetValueCurrency(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetValueCurrency is a free data retrieval call binding the contract method 0x2cce3224.
//
// Solidity: function getValueCurrency(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetValueCurrency(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetValueCurrency(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetValueDate is a free data retrieval call binding the contract method 0xab1de21a.
//
// Solidity: function getValueDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCaller) GetValueDate(opts *bind.CallOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIActivityTransactions.contract.Call(opts, out, "getValueDate", _activitiesRef, _activityRef, _transactionRef)
	return *ret0, err
}

// GetValueDate is a free data retrieval call binding the contract method 0xab1de21a.
//
// Solidity: function getValueDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsSession) GetValueDate(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetValueDate(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// GetValueDate is a free data retrieval call binding the contract method 0xab1de21a.
//
// Solidity: function getValueDate(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef) view returns(bytes32)
func (_IATIActivityTransactions *IATIActivityTransactionsCallerSession) GetValueDate(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte) ([32]byte, error) {
	return _IATIActivityTransactions.Contract.GetValueDate(&_IATIActivityTransactions.CallOpts, _activitiesRef, _activityRef, _transactionRef)
}

// SetTransaction is a paid mutator transaction binding the contract method 0x3598f212.
//
// Solidity: function setTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef, ActivityTransactionsTransaction _transaction) returns()
func (_IATIActivityTransactions *IATIActivityTransactionsTransactor) SetTransaction(opts *bind.TransactOpts, _activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte, _transaction ActivityTransactionsTransaction) (*types.Transaction, error) {
	return _IATIActivityTransactions.contract.Transact(opts, "setTransaction", _activitiesRef, _activityRef, _transactionRef, _transaction)
}

// SetTransaction is a paid mutator transaction binding the contract method 0x3598f212.
//
// Solidity: function setTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef, ActivityTransactionsTransaction _transaction) returns()
func (_IATIActivityTransactions *IATIActivityTransactionsSession) SetTransaction(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte, _transaction ActivityTransactionsTransaction) (*types.Transaction, error) {
	return _IATIActivityTransactions.Contract.SetTransaction(&_IATIActivityTransactions.TransactOpts, _activitiesRef, _activityRef, _transactionRef, _transaction)
}

// SetTransaction is a paid mutator transaction binding the contract method 0x3598f212.
//
// Solidity: function setTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef, ActivityTransactionsTransaction _transaction) returns()
func (_IATIActivityTransactions *IATIActivityTransactionsTransactorSession) SetTransaction(_activitiesRef [32]byte, _activityRef [32]byte, _transactionRef [32]byte, _transaction ActivityTransactionsTransaction) (*types.Transaction, error) {
	return _IATIActivityTransactions.Contract.SetTransaction(&_IATIActivityTransactions.TransactOpts, _activitiesRef, _activityRef, _transactionRef, _transaction)
}

// IATIActivityTransactionsSetTransactionIterator is returned from FilterSetTransaction and is used to iterate over the raw logs and unpacked data for SetTransaction events raised by the IATIActivityTransactions contract.
type IATIActivityTransactionsSetTransactionIterator struct {
	Event *IATIActivityTransactionsSetTransaction // Event containing the contract specifics and raw log

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
func (it *IATIActivityTransactionsSetTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IATIActivityTransactionsSetTransaction)
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
		it.Event = new(IATIActivityTransactionsSetTransaction)
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
func (it *IATIActivityTransactionsSetTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IATIActivityTransactionsSetTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IATIActivityTransactionsSetTransaction represents a SetTransaction event raised by the IATIActivityTransactions contract.
type IATIActivityTransactionsSetTransaction struct {
	ActivitiesRef  [32]byte
	ActivityRef    [32]byte
	TransactionRef [32]byte
	Transaction    ActivityTransactionsTransaction
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetTransaction is a free log retrieval operation binding the contract event 0x8643d357dcd4c137b20a768f0f133ca5b6b57ce5aeea49bee8e81108018c6c6c.
//
// Solidity: event SetTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef, ActivityTransactionsTransaction _transaction)
func (_IATIActivityTransactions *IATIActivityTransactionsFilterer) FilterSetTransaction(opts *bind.FilterOpts) (*IATIActivityTransactionsSetTransactionIterator, error) {

	logs, sub, err := _IATIActivityTransactions.contract.FilterLogs(opts, "SetTransaction")
	if err != nil {
		return nil, err
	}
	return &IATIActivityTransactionsSetTransactionIterator{contract: _IATIActivityTransactions.contract, event: "SetTransaction", logs: logs, sub: sub}, nil
}

// WatchSetTransaction is a free log subscription operation binding the contract event 0x8643d357dcd4c137b20a768f0f133ca5b6b57ce5aeea49bee8e81108018c6c6c.
//
// Solidity: event SetTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef, ActivityTransactionsTransaction _transaction)
func (_IATIActivityTransactions *IATIActivityTransactionsFilterer) WatchSetTransaction(opts *bind.WatchOpts, sink chan<- *IATIActivityTransactionsSetTransaction) (event.Subscription, error) {

	logs, sub, err := _IATIActivityTransactions.contract.WatchLogs(opts, "SetTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IATIActivityTransactionsSetTransaction)
				if err := _IATIActivityTransactions.contract.UnpackLog(event, "SetTransaction", log); err != nil {
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

// ParseSetTransaction is a log parse operation binding the contract event 0x8643d357dcd4c137b20a768f0f133ca5b6b57ce5aeea49bee8e81108018c6c6c.
//
// Solidity: event SetTransaction(bytes32 _activitiesRef, bytes32 _activityRef, bytes32 _transactionRef, ActivityTransactionsTransaction _transaction)
func (_IATIActivityTransactions *IATIActivityTransactionsFilterer) ParseSetTransaction(log types.Log) (*IATIActivityTransactionsSetTransaction, error) {
	event := new(IATIActivityTransactionsSetTransaction)
	if err := _IATIActivityTransactions.contract.UnpackLog(event, "SetTransaction", log); err != nil {
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
