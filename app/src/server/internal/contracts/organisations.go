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

// OrganisationsOrgs is an auto generated low-level Go binding around an user-defined struct.
type OrganisationsOrgs struct {
	Version       [32]byte
	GeneratedTime [32]byte
}

// IATIOrganisationsABI is the input ABI used to generate the binding from.
const IATIOrganisationsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structOrganisations.Orgs\",\"name\":\"_organisations\",\"type\":\"tuple\"}],\"name\":\"SetOrganisations\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"}],\"name\":\"getGeneratedTime\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumOrganisations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"}],\"name\":\"getOrganisations\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"}],\"internalType\":\"structOrganisations.Orgs\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getOrganisationsReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"}],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"}],\"internalType\":\"structOrganisations.Orgs\",\"name\":\"_organisations\",\"type\":\"tuple\"}],\"name\":\"setOrganisations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IATIOrganisationsFuncSigs maps the 4-byte function signature to its string representation.
var IATIOrganisationsFuncSigs = map[string]string{
	"ab73eba0": "getGeneratedTime(bytes32)",
	"f8ef9e4b": "getNumOrganisations()",
	"cc8bd2c6": "getOrganisations(bytes32)",
	"f8165355": "getOrganisationsReference(uint256)",
	"9aaf9f08": "getVersion(bytes32)",
	"c4b66295": "setOrganisations(bytes32,(bytes32,bytes32))",
}

// IATIOrganisationsBin is the compiled bytecode used for deploying new contracts.
var IATIOrganisationsBin = "0x608060405234801561001057600080fd5b506104c3806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80639aaf9f0814610067578063ab73eba014610090578063c4b66295146100a3578063cc8bd2c6146100b8578063f8165355146100d8578063f8ef9e4b146100eb575b600080fd5b61007a61007536600461037f565b6100f3565b604051610087919061040c565b60405180910390f35b61007a61009e36600461037f565b610121565b6100b66100b1366004610397565b610153565b005b6100cb6100c636600461037f565b6102c1565b6040516100879190610479565b61007a6100e636600461037f565b61030d565b61007a61033b565b600081811a60f81b6001600160f81b03191661010e57600080fd5b5060009081526001602052604090205490565b600081811a60f81b6001600160f81b03191661013c57600080fd5b506000908152600160208190526040909120015490565b8160001a60f81b6001600160f81b031916158015906101825750805160001a60f81b6001600160f81b03191615155b80156101a15750602081015160001a60f81b6001600160f81b03191615155b6101aa57600080fd5b6000828152600160208181526040808420855181559185015191909201555163745fca7b60e01b815273__$b620240993a55615b607189176fb82e62f$__9163745fca7b916101fd918691600401610415565b60206040518083038186803b15801561021557600080fd5b505af4158015610229573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061024d9190610358565b61028457600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563018290555b7f7dba940be2bf4784e75053fca32f930608d02e9f7b47558f9513026a646c88b082826040516102b5929190610465565b60405180910390a15050565b6102c9610341565b8160001a60f81b6001600160f81b0319166102e357600080fd5b50600090815260016020818152604092839020835180850190945280548452909101549082015290565b60008054821061031c57600080fd5b6000828154811061032957fe5b90600052602060002001549050919050565b60005490565b604080518082019091526000808252602082015290565b600060208284031215610369578081fd5b81518015158114610378578182fd5b9392505050565b600060208284031215610390578081fd5b5035919050565b60008082840360608112156103aa578182fd5b833592506040601f19820112156103bf578182fd5b506040516040810181811067ffffffffffffffff821117156103df578283fd5b60409081526020858101358352940135938101939093525092909150565b80518252602090810151910152565b90815260200190565b60006040820184835260206040818501528185548084526060860191508685528285209350845b818110156104585784548352600194850194928401920161043c565b5090979650505050505050565b8281526060810161037860208301846103fd565b6040810161048782846103fd565b9291505056fea26469706673582212208f7f3c259030e7559658b3666b8e599159820df5bf068bc3ef82a9ae53fd372a64736f6c63430006060033"

// DeployIATIOrganisations deploys a new Ethereum contract, binding an instance of IATIOrganisations to it.
func DeployIATIOrganisations(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IATIOrganisations, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIOrganisationsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	stringsAddr, _, _, _ := DeployStrings(auth, backend)
	IATIOrganisationsBin = strings.Replace(IATIOrganisationsBin, "__$b620240993a55615b607189176fb82e62f$__", stringsAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IATIOrganisationsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IATIOrganisations{IATIOrganisationsCaller: IATIOrganisationsCaller{contract: contract}, IATIOrganisationsTransactor: IATIOrganisationsTransactor{contract: contract}, IATIOrganisationsFilterer: IATIOrganisationsFilterer{contract: contract}}, nil
}

// IATIOrganisations is an auto generated Go binding around an Ethereum contract.
type IATIOrganisations struct {
	IATIOrganisationsCaller     // Read-only binding to the contract
	IATIOrganisationsTransactor // Write-only binding to the contract
	IATIOrganisationsFilterer   // Log filterer for contract events
}

// IATIOrganisationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IATIOrganisationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIOrganisationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IATIOrganisationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIOrganisationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IATIOrganisationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIOrganisationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IATIOrganisationsSession struct {
	Contract     *IATIOrganisations // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IATIOrganisationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IATIOrganisationsCallerSession struct {
	Contract *IATIOrganisationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IATIOrganisationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IATIOrganisationsTransactorSession struct {
	Contract     *IATIOrganisationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IATIOrganisationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IATIOrganisationsRaw struct {
	Contract *IATIOrganisations // Generic contract binding to access the raw methods on
}

// IATIOrganisationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IATIOrganisationsCallerRaw struct {
	Contract *IATIOrganisationsCaller // Generic read-only contract binding to access the raw methods on
}

// IATIOrganisationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IATIOrganisationsTransactorRaw struct {
	Contract *IATIOrganisationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIATIOrganisations creates a new instance of IATIOrganisations, bound to a specific deployed contract.
func NewIATIOrganisations(address common.Address, backend bind.ContractBackend) (*IATIOrganisations, error) {
	contract, err := bindIATIOrganisations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IATIOrganisations{IATIOrganisationsCaller: IATIOrganisationsCaller{contract: contract}, IATIOrganisationsTransactor: IATIOrganisationsTransactor{contract: contract}, IATIOrganisationsFilterer: IATIOrganisationsFilterer{contract: contract}}, nil
}

// NewIATIOrganisationsCaller creates a new read-only instance of IATIOrganisations, bound to a specific deployed contract.
func NewIATIOrganisationsCaller(address common.Address, caller bind.ContractCaller) (*IATIOrganisationsCaller, error) {
	contract, err := bindIATIOrganisations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IATIOrganisationsCaller{contract: contract}, nil
}

// NewIATIOrganisationsTransactor creates a new write-only instance of IATIOrganisations, bound to a specific deployed contract.
func NewIATIOrganisationsTransactor(address common.Address, transactor bind.ContractTransactor) (*IATIOrganisationsTransactor, error) {
	contract, err := bindIATIOrganisations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IATIOrganisationsTransactor{contract: contract}, nil
}

// NewIATIOrganisationsFilterer creates a new log filterer instance of IATIOrganisations, bound to a specific deployed contract.
func NewIATIOrganisationsFilterer(address common.Address, filterer bind.ContractFilterer) (*IATIOrganisationsFilterer, error) {
	contract, err := bindIATIOrganisations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IATIOrganisationsFilterer{contract: contract}, nil
}

// bindIATIOrganisations binds a generic wrapper to an already deployed contract.
func bindIATIOrganisations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIOrganisationsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIOrganisations *IATIOrganisationsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIOrganisations.Contract.IATIOrganisationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIOrganisations *IATIOrganisationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIOrganisations.Contract.IATIOrganisationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIOrganisations *IATIOrganisationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIOrganisations.Contract.IATIOrganisationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIOrganisations *IATIOrganisationsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIOrganisations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIOrganisations *IATIOrganisationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIOrganisations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIOrganisations *IATIOrganisationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIOrganisations.Contract.contract.Transact(opts, method, params...)
}

// GetGeneratedTime is a free data retrieval call binding the contract method 0xab73eba0.
//
// Solidity: function getGeneratedTime(bytes32 _organisationsRef) view returns(bytes32)
func (_IATIOrganisations *IATIOrganisationsCaller) GetGeneratedTime(opts *bind.CallOpts, _organisationsRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisations.contract.Call(opts, out, "getGeneratedTime", _organisationsRef)
	return *ret0, err
}

// GetGeneratedTime is a free data retrieval call binding the contract method 0xab73eba0.
//
// Solidity: function getGeneratedTime(bytes32 _organisationsRef) view returns(bytes32)
func (_IATIOrganisations *IATIOrganisationsSession) GetGeneratedTime(_organisationsRef [32]byte) ([32]byte, error) {
	return _IATIOrganisations.Contract.GetGeneratedTime(&_IATIOrganisations.CallOpts, _organisationsRef)
}

// GetGeneratedTime is a free data retrieval call binding the contract method 0xab73eba0.
//
// Solidity: function getGeneratedTime(bytes32 _organisationsRef) view returns(bytes32)
func (_IATIOrganisations *IATIOrganisationsCallerSession) GetGeneratedTime(_organisationsRef [32]byte) ([32]byte, error) {
	return _IATIOrganisations.Contract.GetGeneratedTime(&_IATIOrganisations.CallOpts, _organisationsRef)
}

// GetNumOrganisations is a free data retrieval call binding the contract method 0xf8ef9e4b.
//
// Solidity: function getNumOrganisations() view returns(uint256)
func (_IATIOrganisations *IATIOrganisationsCaller) GetNumOrganisations(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIOrganisations.contract.Call(opts, out, "getNumOrganisations")
	return *ret0, err
}

// GetNumOrganisations is a free data retrieval call binding the contract method 0xf8ef9e4b.
//
// Solidity: function getNumOrganisations() view returns(uint256)
func (_IATIOrganisations *IATIOrganisationsSession) GetNumOrganisations() (*big.Int, error) {
	return _IATIOrganisations.Contract.GetNumOrganisations(&_IATIOrganisations.CallOpts)
}

// GetNumOrganisations is a free data retrieval call binding the contract method 0xf8ef9e4b.
//
// Solidity: function getNumOrganisations() view returns(uint256)
func (_IATIOrganisations *IATIOrganisationsCallerSession) GetNumOrganisations() (*big.Int, error) {
	return _IATIOrganisations.Contract.GetNumOrganisations(&_IATIOrganisations.CallOpts)
}

// GetOrganisations is a free data retrieval call binding the contract method 0xcc8bd2c6.
//
// Solidity: function getOrganisations(bytes32 _organisationsRef) view returns(OrganisationsOrgs)
func (_IATIOrganisations *IATIOrganisationsCaller) GetOrganisations(opts *bind.CallOpts, _organisationsRef [32]byte) (OrganisationsOrgs, error) {
	var (
		ret0 = new(OrganisationsOrgs)
	)
	out := ret0
	err := _IATIOrganisations.contract.Call(opts, out, "getOrganisations", _organisationsRef)
	return *ret0, err
}

// GetOrganisations is a free data retrieval call binding the contract method 0xcc8bd2c6.
//
// Solidity: function getOrganisations(bytes32 _organisationsRef) view returns(OrganisationsOrgs)
func (_IATIOrganisations *IATIOrganisationsSession) GetOrganisations(_organisationsRef [32]byte) (OrganisationsOrgs, error) {
	return _IATIOrganisations.Contract.GetOrganisations(&_IATIOrganisations.CallOpts, _organisationsRef)
}

// GetOrganisations is a free data retrieval call binding the contract method 0xcc8bd2c6.
//
// Solidity: function getOrganisations(bytes32 _organisationsRef) view returns(OrganisationsOrgs)
func (_IATIOrganisations *IATIOrganisationsCallerSession) GetOrganisations(_organisationsRef [32]byte) (OrganisationsOrgs, error) {
	return _IATIOrganisations.Contract.GetOrganisations(&_IATIOrganisations.CallOpts, _organisationsRef)
}

// GetOrganisationsReference is a free data retrieval call binding the contract method 0xf8165355.
//
// Solidity: function getOrganisationsReference(uint256 _index) view returns(bytes32)
func (_IATIOrganisations *IATIOrganisationsCaller) GetOrganisationsReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisations.contract.Call(opts, out, "getOrganisationsReference", _index)
	return *ret0, err
}

// GetOrganisationsReference is a free data retrieval call binding the contract method 0xf8165355.
//
// Solidity: function getOrganisationsReference(uint256 _index) view returns(bytes32)
func (_IATIOrganisations *IATIOrganisationsSession) GetOrganisationsReference(_index *big.Int) ([32]byte, error) {
	return _IATIOrganisations.Contract.GetOrganisationsReference(&_IATIOrganisations.CallOpts, _index)
}

// GetOrganisationsReference is a free data retrieval call binding the contract method 0xf8165355.
//
// Solidity: function getOrganisationsReference(uint256 _index) view returns(bytes32)
func (_IATIOrganisations *IATIOrganisationsCallerSession) GetOrganisationsReference(_index *big.Int) ([32]byte, error) {
	return _IATIOrganisations.Contract.GetOrganisationsReference(&_IATIOrganisations.CallOpts, _index)
}

// GetVersion is a free data retrieval call binding the contract method 0x9aaf9f08.
//
// Solidity: function getVersion(bytes32 _organisationsRef) view returns(bytes32)
func (_IATIOrganisations *IATIOrganisationsCaller) GetVersion(opts *bind.CallOpts, _organisationsRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisations.contract.Call(opts, out, "getVersion", _organisationsRef)
	return *ret0, err
}

// GetVersion is a free data retrieval call binding the contract method 0x9aaf9f08.
//
// Solidity: function getVersion(bytes32 _organisationsRef) view returns(bytes32)
func (_IATIOrganisations *IATIOrganisationsSession) GetVersion(_organisationsRef [32]byte) ([32]byte, error) {
	return _IATIOrganisations.Contract.GetVersion(&_IATIOrganisations.CallOpts, _organisationsRef)
}

// GetVersion is a free data retrieval call binding the contract method 0x9aaf9f08.
//
// Solidity: function getVersion(bytes32 _organisationsRef) view returns(bytes32)
func (_IATIOrganisations *IATIOrganisationsCallerSession) GetVersion(_organisationsRef [32]byte) ([32]byte, error) {
	return _IATIOrganisations.Contract.GetVersion(&_IATIOrganisations.CallOpts, _organisationsRef)
}

// SetOrganisations is a paid mutator transaction binding the contract method 0xc4b66295.
//
// Solidity: function setOrganisations(bytes32 _organisationsRef, OrganisationsOrgs _organisations) returns()
func (_IATIOrganisations *IATIOrganisationsTransactor) SetOrganisations(opts *bind.TransactOpts, _organisationsRef [32]byte, _organisations OrganisationsOrgs) (*types.Transaction, error) {
	return _IATIOrganisations.contract.Transact(opts, "setOrganisations", _organisationsRef, _organisations)
}

// SetOrganisations is a paid mutator transaction binding the contract method 0xc4b66295.
//
// Solidity: function setOrganisations(bytes32 _organisationsRef, OrganisationsOrgs _organisations) returns()
func (_IATIOrganisations *IATIOrganisationsSession) SetOrganisations(_organisationsRef [32]byte, _organisations OrganisationsOrgs) (*types.Transaction, error) {
	return _IATIOrganisations.Contract.SetOrganisations(&_IATIOrganisations.TransactOpts, _organisationsRef, _organisations)
}

// SetOrganisations is a paid mutator transaction binding the contract method 0xc4b66295.
//
// Solidity: function setOrganisations(bytes32 _organisationsRef, OrganisationsOrgs _organisations) returns()
func (_IATIOrganisations *IATIOrganisationsTransactorSession) SetOrganisations(_organisationsRef [32]byte, _organisations OrganisationsOrgs) (*types.Transaction, error) {
	return _IATIOrganisations.Contract.SetOrganisations(&_IATIOrganisations.TransactOpts, _organisationsRef, _organisations)
}

// IATIOrganisationsSetOrganisationsIterator is returned from FilterSetOrganisations and is used to iterate over the raw logs and unpacked data for SetOrganisations events raised by the IATIOrganisations contract.
type IATIOrganisationsSetOrganisationsIterator struct {
	Event *IATIOrganisationsSetOrganisations // Event containing the contract specifics and raw log

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
func (it *IATIOrganisationsSetOrganisationsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IATIOrganisationsSetOrganisations)
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
		it.Event = new(IATIOrganisationsSetOrganisations)
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
func (it *IATIOrganisationsSetOrganisationsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IATIOrganisationsSetOrganisationsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IATIOrganisationsSetOrganisations represents a SetOrganisations event raised by the IATIOrganisations contract.
type IATIOrganisationsSetOrganisations struct {
	OrganisationsRef [32]byte
	Organisations    OrganisationsOrgs
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetOrganisations is a free log retrieval operation binding the contract event 0x7dba940be2bf4784e75053fca32f930608d02e9f7b47558f9513026a646c88b0.
//
// Solidity: event SetOrganisations(bytes32 _organisationsRef, OrganisationsOrgs _organisations)
func (_IATIOrganisations *IATIOrganisationsFilterer) FilterSetOrganisations(opts *bind.FilterOpts) (*IATIOrganisationsSetOrganisationsIterator, error) {

	logs, sub, err := _IATIOrganisations.contract.FilterLogs(opts, "SetOrganisations")
	if err != nil {
		return nil, err
	}
	return &IATIOrganisationsSetOrganisationsIterator{contract: _IATIOrganisations.contract, event: "SetOrganisations", logs: logs, sub: sub}, nil
}

// WatchSetOrganisations is a free log subscription operation binding the contract event 0x7dba940be2bf4784e75053fca32f930608d02e9f7b47558f9513026a646c88b0.
//
// Solidity: event SetOrganisations(bytes32 _organisationsRef, OrganisationsOrgs _organisations)
func (_IATIOrganisations *IATIOrganisationsFilterer) WatchSetOrganisations(opts *bind.WatchOpts, sink chan<- *IATIOrganisationsSetOrganisations) (event.Subscription, error) {

	logs, sub, err := _IATIOrganisations.contract.WatchLogs(opts, "SetOrganisations")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IATIOrganisationsSetOrganisations)
				if err := _IATIOrganisations.contract.UnpackLog(event, "SetOrganisations", log); err != nil {
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

// ParseSetOrganisations is a log parse operation binding the contract event 0x7dba940be2bf4784e75053fca32f930608d02e9f7b47558f9513026a646c88b0.
//
// Solidity: event SetOrganisations(bytes32 _organisationsRef, OrganisationsOrgs _organisations)
func (_IATIOrganisations *IATIOrganisationsFilterer) ParseSetOrganisations(log types.Log) (*IATIOrganisationsSetOrganisations, error) {
	event := new(IATIOrganisationsSetOrganisations)
	if err := _IATIOrganisations.contract.UnpackLog(event, "SetOrganisations", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OrganisationsABI is the input ABI used to generate the binding from.
const OrganisationsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structOrganisations.Orgs\",\"name\":\"_organisations\",\"type\":\"tuple\"}],\"name\":\"SetOrganisations\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"}],\"name\":\"getGeneratedTime\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumOrganisations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"}],\"name\":\"getOrganisations\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"}],\"internalType\":\"structOrganisations.Orgs\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getOrganisationsReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"}],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"version\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"generatedTime\",\"type\":\"bytes32\"}],\"internalType\":\"structOrganisations.Orgs\",\"name\":\"_organisations\",\"type\":\"tuple\"}],\"name\":\"setOrganisations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OrganisationsFuncSigs maps the 4-byte function signature to its string representation.
var OrganisationsFuncSigs = map[string]string{
	"ab73eba0": "getGeneratedTime(bytes32)",
	"f8ef9e4b": "getNumOrganisations()",
	"cc8bd2c6": "getOrganisations(bytes32)",
	"f8165355": "getOrganisationsReference(uint256)",
	"9aaf9f08": "getVersion(bytes32)",
	"c4b66295": "setOrganisations(bytes32,(bytes32,bytes32))",
}

// Organisations is an auto generated Go binding around an Ethereum contract.
type Organisations struct {
	OrganisationsCaller     // Read-only binding to the contract
	OrganisationsTransactor // Write-only binding to the contract
	OrganisationsFilterer   // Log filterer for contract events
}

// OrganisationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrganisationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganisationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrganisationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganisationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrganisationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganisationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrganisationsSession struct {
	Contract     *Organisations    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrganisationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrganisationsCallerSession struct {
	Contract *OrganisationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OrganisationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrganisationsTransactorSession struct {
	Contract     *OrganisationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OrganisationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrganisationsRaw struct {
	Contract *Organisations // Generic contract binding to access the raw methods on
}

// OrganisationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrganisationsCallerRaw struct {
	Contract *OrganisationsCaller // Generic read-only contract binding to access the raw methods on
}

// OrganisationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrganisationsTransactorRaw struct {
	Contract *OrganisationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrganisations creates a new instance of Organisations, bound to a specific deployed contract.
func NewOrganisations(address common.Address, backend bind.ContractBackend) (*Organisations, error) {
	contract, err := bindOrganisations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Organisations{OrganisationsCaller: OrganisationsCaller{contract: contract}, OrganisationsTransactor: OrganisationsTransactor{contract: contract}, OrganisationsFilterer: OrganisationsFilterer{contract: contract}}, nil
}

// NewOrganisationsCaller creates a new read-only instance of Organisations, bound to a specific deployed contract.
func NewOrganisationsCaller(address common.Address, caller bind.ContractCaller) (*OrganisationsCaller, error) {
	contract, err := bindOrganisations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrganisationsCaller{contract: contract}, nil
}

// NewOrganisationsTransactor creates a new write-only instance of Organisations, bound to a specific deployed contract.
func NewOrganisationsTransactor(address common.Address, transactor bind.ContractTransactor) (*OrganisationsTransactor, error) {
	contract, err := bindOrganisations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrganisationsTransactor{contract: contract}, nil
}

// NewOrganisationsFilterer creates a new log filterer instance of Organisations, bound to a specific deployed contract.
func NewOrganisationsFilterer(address common.Address, filterer bind.ContractFilterer) (*OrganisationsFilterer, error) {
	contract, err := bindOrganisations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrganisationsFilterer{contract: contract}, nil
}

// bindOrganisations binds a generic wrapper to an already deployed contract.
func bindOrganisations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrganisationsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Organisations *OrganisationsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Organisations.Contract.OrganisationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Organisations *OrganisationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Organisations.Contract.OrganisationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Organisations *OrganisationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Organisations.Contract.OrganisationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Organisations *OrganisationsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Organisations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Organisations *OrganisationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Organisations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Organisations *OrganisationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Organisations.Contract.contract.Transact(opts, method, params...)
}

// GetGeneratedTime is a free data retrieval call binding the contract method 0xab73eba0.
//
// Solidity: function getGeneratedTime(bytes32 _organisationsRef) view returns(bytes32)
func (_Organisations *OrganisationsCaller) GetGeneratedTime(opts *bind.CallOpts, _organisationsRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Organisations.contract.Call(opts, out, "getGeneratedTime", _organisationsRef)
	return *ret0, err
}

// GetGeneratedTime is a free data retrieval call binding the contract method 0xab73eba0.
//
// Solidity: function getGeneratedTime(bytes32 _organisationsRef) view returns(bytes32)
func (_Organisations *OrganisationsSession) GetGeneratedTime(_organisationsRef [32]byte) ([32]byte, error) {
	return _Organisations.Contract.GetGeneratedTime(&_Organisations.CallOpts, _organisationsRef)
}

// GetGeneratedTime is a free data retrieval call binding the contract method 0xab73eba0.
//
// Solidity: function getGeneratedTime(bytes32 _organisationsRef) view returns(bytes32)
func (_Organisations *OrganisationsCallerSession) GetGeneratedTime(_organisationsRef [32]byte) ([32]byte, error) {
	return _Organisations.Contract.GetGeneratedTime(&_Organisations.CallOpts, _organisationsRef)
}

// GetNumOrganisations is a free data retrieval call binding the contract method 0xf8ef9e4b.
//
// Solidity: function getNumOrganisations() view returns(uint256)
func (_Organisations *OrganisationsCaller) GetNumOrganisations(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Organisations.contract.Call(opts, out, "getNumOrganisations")
	return *ret0, err
}

// GetNumOrganisations is a free data retrieval call binding the contract method 0xf8ef9e4b.
//
// Solidity: function getNumOrganisations() view returns(uint256)
func (_Organisations *OrganisationsSession) GetNumOrganisations() (*big.Int, error) {
	return _Organisations.Contract.GetNumOrganisations(&_Organisations.CallOpts)
}

// GetNumOrganisations is a free data retrieval call binding the contract method 0xf8ef9e4b.
//
// Solidity: function getNumOrganisations() view returns(uint256)
func (_Organisations *OrganisationsCallerSession) GetNumOrganisations() (*big.Int, error) {
	return _Organisations.Contract.GetNumOrganisations(&_Organisations.CallOpts)
}

// GetOrganisations is a free data retrieval call binding the contract method 0xcc8bd2c6.
//
// Solidity: function getOrganisations(bytes32 _organisationsRef) view returns(OrganisationsOrgs)
func (_Organisations *OrganisationsCaller) GetOrganisations(opts *bind.CallOpts, _organisationsRef [32]byte) (OrganisationsOrgs, error) {
	var (
		ret0 = new(OrganisationsOrgs)
	)
	out := ret0
	err := _Organisations.contract.Call(opts, out, "getOrganisations", _organisationsRef)
	return *ret0, err
}

// GetOrganisations is a free data retrieval call binding the contract method 0xcc8bd2c6.
//
// Solidity: function getOrganisations(bytes32 _organisationsRef) view returns(OrganisationsOrgs)
func (_Organisations *OrganisationsSession) GetOrganisations(_organisationsRef [32]byte) (OrganisationsOrgs, error) {
	return _Organisations.Contract.GetOrganisations(&_Organisations.CallOpts, _organisationsRef)
}

// GetOrganisations is a free data retrieval call binding the contract method 0xcc8bd2c6.
//
// Solidity: function getOrganisations(bytes32 _organisationsRef) view returns(OrganisationsOrgs)
func (_Organisations *OrganisationsCallerSession) GetOrganisations(_organisationsRef [32]byte) (OrganisationsOrgs, error) {
	return _Organisations.Contract.GetOrganisations(&_Organisations.CallOpts, _organisationsRef)
}

// GetOrganisationsReference is a free data retrieval call binding the contract method 0xf8165355.
//
// Solidity: function getOrganisationsReference(uint256 _index) view returns(bytes32)
func (_Organisations *OrganisationsCaller) GetOrganisationsReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Organisations.contract.Call(opts, out, "getOrganisationsReference", _index)
	return *ret0, err
}

// GetOrganisationsReference is a free data retrieval call binding the contract method 0xf8165355.
//
// Solidity: function getOrganisationsReference(uint256 _index) view returns(bytes32)
func (_Organisations *OrganisationsSession) GetOrganisationsReference(_index *big.Int) ([32]byte, error) {
	return _Organisations.Contract.GetOrganisationsReference(&_Organisations.CallOpts, _index)
}

// GetOrganisationsReference is a free data retrieval call binding the contract method 0xf8165355.
//
// Solidity: function getOrganisationsReference(uint256 _index) view returns(bytes32)
func (_Organisations *OrganisationsCallerSession) GetOrganisationsReference(_index *big.Int) ([32]byte, error) {
	return _Organisations.Contract.GetOrganisationsReference(&_Organisations.CallOpts, _index)
}

// GetVersion is a free data retrieval call binding the contract method 0x9aaf9f08.
//
// Solidity: function getVersion(bytes32 _organisationsRef) view returns(bytes32)
func (_Organisations *OrganisationsCaller) GetVersion(opts *bind.CallOpts, _organisationsRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Organisations.contract.Call(opts, out, "getVersion", _organisationsRef)
	return *ret0, err
}

// GetVersion is a free data retrieval call binding the contract method 0x9aaf9f08.
//
// Solidity: function getVersion(bytes32 _organisationsRef) view returns(bytes32)
func (_Organisations *OrganisationsSession) GetVersion(_organisationsRef [32]byte) ([32]byte, error) {
	return _Organisations.Contract.GetVersion(&_Organisations.CallOpts, _organisationsRef)
}

// GetVersion is a free data retrieval call binding the contract method 0x9aaf9f08.
//
// Solidity: function getVersion(bytes32 _organisationsRef) view returns(bytes32)
func (_Organisations *OrganisationsCallerSession) GetVersion(_organisationsRef [32]byte) ([32]byte, error) {
	return _Organisations.Contract.GetVersion(&_Organisations.CallOpts, _organisationsRef)
}

// SetOrganisations is a paid mutator transaction binding the contract method 0xc4b66295.
//
// Solidity: function setOrganisations(bytes32 _organisationsRef, OrganisationsOrgs _organisations) returns()
func (_Organisations *OrganisationsTransactor) SetOrganisations(opts *bind.TransactOpts, _organisationsRef [32]byte, _organisations OrganisationsOrgs) (*types.Transaction, error) {
	return _Organisations.contract.Transact(opts, "setOrganisations", _organisationsRef, _organisations)
}

// SetOrganisations is a paid mutator transaction binding the contract method 0xc4b66295.
//
// Solidity: function setOrganisations(bytes32 _organisationsRef, OrganisationsOrgs _organisations) returns()
func (_Organisations *OrganisationsSession) SetOrganisations(_organisationsRef [32]byte, _organisations OrganisationsOrgs) (*types.Transaction, error) {
	return _Organisations.Contract.SetOrganisations(&_Organisations.TransactOpts, _organisationsRef, _organisations)
}

// SetOrganisations is a paid mutator transaction binding the contract method 0xc4b66295.
//
// Solidity: function setOrganisations(bytes32 _organisationsRef, OrganisationsOrgs _organisations) returns()
func (_Organisations *OrganisationsTransactorSession) SetOrganisations(_organisationsRef [32]byte, _organisations OrganisationsOrgs) (*types.Transaction, error) {
	return _Organisations.Contract.SetOrganisations(&_Organisations.TransactOpts, _organisationsRef, _organisations)
}

// OrganisationsSetOrganisationsIterator is returned from FilterSetOrganisations and is used to iterate over the raw logs and unpacked data for SetOrganisations events raised by the Organisations contract.
type OrganisationsSetOrganisationsIterator struct {
	Event *OrganisationsSetOrganisations // Event containing the contract specifics and raw log

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
func (it *OrganisationsSetOrganisationsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrganisationsSetOrganisations)
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
		it.Event = new(OrganisationsSetOrganisations)
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
func (it *OrganisationsSetOrganisationsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrganisationsSetOrganisationsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrganisationsSetOrganisations represents a SetOrganisations event raised by the Organisations contract.
type OrganisationsSetOrganisations struct {
	OrganisationsRef [32]byte
	Organisations    OrganisationsOrgs
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetOrganisations is a free log retrieval operation binding the contract event 0x7dba940be2bf4784e75053fca32f930608d02e9f7b47558f9513026a646c88b0.
//
// Solidity: event SetOrganisations(bytes32 _organisationsRef, OrganisationsOrgs _organisations)
func (_Organisations *OrganisationsFilterer) FilterSetOrganisations(opts *bind.FilterOpts) (*OrganisationsSetOrganisationsIterator, error) {

	logs, sub, err := _Organisations.contract.FilterLogs(opts, "SetOrganisations")
	if err != nil {
		return nil, err
	}
	return &OrganisationsSetOrganisationsIterator{contract: _Organisations.contract, event: "SetOrganisations", logs: logs, sub: sub}, nil
}

// WatchSetOrganisations is a free log subscription operation binding the contract event 0x7dba940be2bf4784e75053fca32f930608d02e9f7b47558f9513026a646c88b0.
//
// Solidity: event SetOrganisations(bytes32 _organisationsRef, OrganisationsOrgs _organisations)
func (_Organisations *OrganisationsFilterer) WatchSetOrganisations(opts *bind.WatchOpts, sink chan<- *OrganisationsSetOrganisations) (event.Subscription, error) {

	logs, sub, err := _Organisations.contract.WatchLogs(opts, "SetOrganisations")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrganisationsSetOrganisations)
				if err := _Organisations.contract.UnpackLog(event, "SetOrganisations", log); err != nil {
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

// ParseSetOrganisations is a log parse operation binding the contract event 0x7dba940be2bf4784e75053fca32f930608d02e9f7b47558f9513026a646c88b0.
//
// Solidity: event SetOrganisations(bytes32 _organisationsRef, OrganisationsOrgs _organisations)
func (_Organisations *OrganisationsFilterer) ParseSetOrganisations(log types.Log) (*OrganisationsSetOrganisations, error) {
	event := new(OrganisationsSetOrganisations)
	if err := _Organisations.contract.UnpackLog(event, "SetOrganisations", log); err != nil {
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
