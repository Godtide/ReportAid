// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package orgs

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

// OrgsOrg is an auto generated low-level Go binding around an user-defined struct.
type OrgsOrg struct {
	Name       string
	Identifier string
}

// IATIOrgsABI is the input ABI used to generate the binding from.
const IATIOrgsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_orgRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structOrgs.Org\",\"name\":\"_org\",\"type\":\"tuple\"}],\"name\":\"SetOrg\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getNumOrgs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgRef\",\"type\":\"bytes32\"}],\"name\":\"getOrg\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"}],\"internalType\":\"structOrgs.Org\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgRef\",\"type\":\"bytes32\"}],\"name\":\"getOrgIdentifier\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgRef\",\"type\":\"bytes32\"}],\"name\":\"getOrgName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getOrgReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"}],\"internalType\":\"structOrgs.Org\",\"name\":\"_org\",\"type\":\"tuple\"}],\"name\":\"setOrg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IATIOrgsFuncSigs maps the 4-byte function signature to its string representation.
var IATIOrgsFuncSigs = map[string]string{
	"1a10f51c": "getNumOrgs()",
	"81e84589": "getOrg(bytes32)",
	"a4abc4ce": "getOrgIdentifier(bytes32)",
	"195541cc": "getOrgName(bytes32)",
	"0dd75b22": "getOrgReference(uint256)",
	"df9a8bbe": "setOrg(bytes32,(string,string))",
}

// IATIOrgsBin is the compiled bytecode used for deploying new contracts.
var IATIOrgsBin = "0x608060405234801561001057600080fd5b506108d3806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80630dd75b2214610067578063195541cc146100905780631a10f51c146100b057806381e84589146100b8578063a4abc4ce146100d8578063df9a8bbe146100eb575b600080fd5b61007a6100753660046106a1565b610100565b60405161008791906107d6565b60405180910390f35b6100a361009e3660046106a1565b61012e565b6040516100879190610850565b61007a6101e8565b6100cb6100c63660046106a1565b6101ef565b6040516100879190610863565b6100a36100e63660046106a1565b610361565b6100fe6100f93660046106b9565b6103e6565b005b60008054821061010f57600080fd5b6000828154811061011c57fe5b90600052602060002001549050919050565b60608160001a60f81b6001600160f81b03191661014a57600080fd5b60008281526001602081815260409283902080548451600294821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156101dc5780601f106101b1576101008083540402835291602001916101dc565b820191906000526020600020905b8154815290600101906020018083116101bf57829003601f168201915b50505050509050919050565b6000545b90565b6101f761055f565b8160001a60f81b6001600160f81b03191661021157600080fd5b6000828152600160208181526040928390208351815460029481161561010002600019011693909304601f810183900490920283016060908101855293830182815292939092849290918491908401828280156102af5780601f10610284576101008083540402835291602001916102af565b820191906000526020600020905b81548152906001019060200180831161029257829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103515780601f1061032657610100808354040283529160200191610351565b820191906000526020600020905b81548152906001019060200180831161033457829003601f168201915b5050505050815250509050919050565b60608160001a60f81b6001600160f81b03191661037d57600080fd5b600082815260016020818152604092839020820180548451600294821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156101dc5780601f106101b1576101008083540402835291602001916101dc565b8160001a60f81b6001600160f81b03191615801590610406575080515115155b801561041757506000816020015151115b61042057600080fd5b6000828152600160209081526040909120825180518493610445928492910190610579565b50602082810151805161045e9260018501920190610579565b505060405163745fca7b60e01b815273__$b620240993a55615b607189176fb82e62f$__915063745fca7b9061049b9085906000906004016107df565b60206040518083038186803b1580156104b357600080fd5b505af41580156104c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104eb919061067a565b61052257600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563018290555b7f7b20403f3c5b881e1e5698bac32df980c612a3fd2632c887b2665dd91c33c465828260405161055392919061082f565b60405180910390a15050565b604051806040016040528060608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106105ba57805160ff19168380011785556105e7565b828001600101855582156105e7579182015b828111156105e75782518255916020019190600101906105cc565b506105f39291506105f7565b5090565b6101ec91905b808211156105f357600081556001016105fd565b600082601f830112610621578081fd5b813567ffffffffffffffff811115610637578182fd5b61064a601f8201601f1916602001610876565b915080825283602082850101111561066157600080fd5b8060208401602084013760009082016020015292915050565b60006020828403121561068b578081fd5b8151801515811461069a578182fd5b9392505050565b6000602082840312156106b2578081fd5b5035919050565b600080604083850312156106cb578081fd5b82359150602083013567ffffffffffffffff808211156106e9578283fd5b818501604081880312156106fb578384fd5b6107056040610876565b9250803582811115610715578485fd5b61072188828401610611565b845250602081013582811115610735578485fd5b61074188828401610611565b6020850152505050809150509250929050565b60008151808452815b818110156107795760208185018101518683018201520161075d565b8181111561078a5782602083870101525b50601f01601f19169290920160200192915050565b60008151604084526107b46040850182610754565b6020840151915084810360208601526107cd8183610754565b95945050505050565b90815260200190565b60006040820184835260206040818501528185548084526060860191508685528285209350845b8181101561082257845483526001948501949284019201610806565b5090979650505050505050565b600083825260406020830152610848604083018461079f565b949350505050565b60006020825261069a6020830184610754565b60006020825261069a602083018461079f565b60405181810167ffffffffffffffff8111828210171561089557600080fd5b60405291905056fea2646970667358221220603ba0f0e3fa28ecb0d410c7e48acb16e4a5a27a757c38d0f35f043eb4f60c2a64736f6c63430006060033"

// DeployIATIOrgs deploys a new Ethereum contract, binding an instance of IATIOrgs to it.
func DeployIATIOrgs(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IATIOrgs, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIOrgsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	stringsAddr, _, _, _ := DeployStrings(auth, backend)
	IATIOrgsBin = strings.Replace(IATIOrgsBin, "__$b620240993a55615b607189176fb82e62f$__", stringsAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IATIOrgsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IATIOrgs{IATIOrgsCaller: IATIOrgsCaller{contract: contract}, IATIOrgsTransactor: IATIOrgsTransactor{contract: contract}, IATIOrgsFilterer: IATIOrgsFilterer{contract: contract}}, nil
}

// IATIOrgs is an auto generated Go binding around an Ethereum contract.
type IATIOrgs struct {
	IATIOrgsCaller     // Read-only binding to the contract
	IATIOrgsTransactor // Write-only binding to the contract
	IATIOrgsFilterer   // Log filterer for contract events
}

// IATIOrgsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IATIOrgsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIOrgsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IATIOrgsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIOrgsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IATIOrgsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIOrgsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IATIOrgsSession struct {
	Contract     *IATIOrgs         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IATIOrgsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IATIOrgsCallerSession struct {
	Contract *IATIOrgsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IATIOrgsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IATIOrgsTransactorSession struct {
	Contract     *IATIOrgsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IATIOrgsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IATIOrgsRaw struct {
	Contract *IATIOrgs // Generic contract binding to access the raw methods on
}

// IATIOrgsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IATIOrgsCallerRaw struct {
	Contract *IATIOrgsCaller // Generic read-only contract binding to access the raw methods on
}

// IATIOrgsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IATIOrgsTransactorRaw struct {
	Contract *IATIOrgsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIATIOrgs creates a new instance of IATIOrgs, bound to a specific deployed contract.
func NewIATIOrgs(address common.Address, backend bind.ContractBackend) (*IATIOrgs, error) {
	contract, err := bindIATIOrgs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IATIOrgs{IATIOrgsCaller: IATIOrgsCaller{contract: contract}, IATIOrgsTransactor: IATIOrgsTransactor{contract: contract}, IATIOrgsFilterer: IATIOrgsFilterer{contract: contract}}, nil
}

// NewIATIOrgsCaller creates a new read-only instance of IATIOrgs, bound to a specific deployed contract.
func NewIATIOrgsCaller(address common.Address, caller bind.ContractCaller) (*IATIOrgsCaller, error) {
	contract, err := bindIATIOrgs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IATIOrgsCaller{contract: contract}, nil
}

// NewIATIOrgsTransactor creates a new write-only instance of IATIOrgs, bound to a specific deployed contract.
func NewIATIOrgsTransactor(address common.Address, transactor bind.ContractTransactor) (*IATIOrgsTransactor, error) {
	contract, err := bindIATIOrgs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IATIOrgsTransactor{contract: contract}, nil
}

// NewIATIOrgsFilterer creates a new log filterer instance of IATIOrgs, bound to a specific deployed contract.
func NewIATIOrgsFilterer(address common.Address, filterer bind.ContractFilterer) (*IATIOrgsFilterer, error) {
	contract, err := bindIATIOrgs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IATIOrgsFilterer{contract: contract}, nil
}

// bindIATIOrgs binds a generic wrapper to an already deployed contract.
func bindIATIOrgs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIOrgsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIOrgs *IATIOrgsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIOrgs.Contract.IATIOrgsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIOrgs *IATIOrgsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIOrgs.Contract.IATIOrgsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIOrgs *IATIOrgsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIOrgs.Contract.IATIOrgsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIOrgs *IATIOrgsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIOrgs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIOrgs *IATIOrgsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIOrgs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIOrgs *IATIOrgsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIOrgs.Contract.contract.Transact(opts, method, params...)
}

// GetNumOrgs is a free data retrieval call binding the contract method 0x1a10f51c.
//
// Solidity: function getNumOrgs() view returns(uint256)
func (_IATIOrgs *IATIOrgsCaller) GetNumOrgs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIOrgs.contract.Call(opts, out, "getNumOrgs")
	return *ret0, err
}

// GetNumOrgs is a free data retrieval call binding the contract method 0x1a10f51c.
//
// Solidity: function getNumOrgs() view returns(uint256)
func (_IATIOrgs *IATIOrgsSession) GetNumOrgs() (*big.Int, error) {
	return _IATIOrgs.Contract.GetNumOrgs(&_IATIOrgs.CallOpts)
}

// GetNumOrgs is a free data retrieval call binding the contract method 0x1a10f51c.
//
// Solidity: function getNumOrgs() view returns(uint256)
func (_IATIOrgs *IATIOrgsCallerSession) GetNumOrgs() (*big.Int, error) {
	return _IATIOrgs.Contract.GetNumOrgs(&_IATIOrgs.CallOpts)
}

// GetOrg is a free data retrieval call binding the contract method 0x81e84589.
//
// Solidity: function getOrg(bytes32 _orgRef) view returns(OrgsOrg)
func (_IATIOrgs *IATIOrgsCaller) GetOrg(opts *bind.CallOpts, _orgRef [32]byte) (OrgsOrg, error) {
	var (
		ret0 = new(OrgsOrg)
	)
	out := ret0
	err := _IATIOrgs.contract.Call(opts, out, "getOrg", _orgRef)
	return *ret0, err
}

// GetOrg is a free data retrieval call binding the contract method 0x81e84589.
//
// Solidity: function getOrg(bytes32 _orgRef) view returns(OrgsOrg)
func (_IATIOrgs *IATIOrgsSession) GetOrg(_orgRef [32]byte) (OrgsOrg, error) {
	return _IATIOrgs.Contract.GetOrg(&_IATIOrgs.CallOpts, _orgRef)
}

// GetOrg is a free data retrieval call binding the contract method 0x81e84589.
//
// Solidity: function getOrg(bytes32 _orgRef) view returns(OrgsOrg)
func (_IATIOrgs *IATIOrgsCallerSession) GetOrg(_orgRef [32]byte) (OrgsOrg, error) {
	return _IATIOrgs.Contract.GetOrg(&_IATIOrgs.CallOpts, _orgRef)
}

// GetOrgIdentifier is a free data retrieval call binding the contract method 0xa4abc4ce.
//
// Solidity: function getOrgIdentifier(bytes32 _orgRef) view returns(string)
func (_IATIOrgs *IATIOrgsCaller) GetOrgIdentifier(opts *bind.CallOpts, _orgRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IATIOrgs.contract.Call(opts, out, "getOrgIdentifier", _orgRef)
	return *ret0, err
}

// GetOrgIdentifier is a free data retrieval call binding the contract method 0xa4abc4ce.
//
// Solidity: function getOrgIdentifier(bytes32 _orgRef) view returns(string)
func (_IATIOrgs *IATIOrgsSession) GetOrgIdentifier(_orgRef [32]byte) (string, error) {
	return _IATIOrgs.Contract.GetOrgIdentifier(&_IATIOrgs.CallOpts, _orgRef)
}

// GetOrgIdentifier is a free data retrieval call binding the contract method 0xa4abc4ce.
//
// Solidity: function getOrgIdentifier(bytes32 _orgRef) view returns(string)
func (_IATIOrgs *IATIOrgsCallerSession) GetOrgIdentifier(_orgRef [32]byte) (string, error) {
	return _IATIOrgs.Contract.GetOrgIdentifier(&_IATIOrgs.CallOpts, _orgRef)
}

// GetOrgName is a free data retrieval call binding the contract method 0x195541cc.
//
// Solidity: function getOrgName(bytes32 _orgRef) view returns(string)
func (_IATIOrgs *IATIOrgsCaller) GetOrgName(opts *bind.CallOpts, _orgRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IATIOrgs.contract.Call(opts, out, "getOrgName", _orgRef)
	return *ret0, err
}

// GetOrgName is a free data retrieval call binding the contract method 0x195541cc.
//
// Solidity: function getOrgName(bytes32 _orgRef) view returns(string)
func (_IATIOrgs *IATIOrgsSession) GetOrgName(_orgRef [32]byte) (string, error) {
	return _IATIOrgs.Contract.GetOrgName(&_IATIOrgs.CallOpts, _orgRef)
}

// GetOrgName is a free data retrieval call binding the contract method 0x195541cc.
//
// Solidity: function getOrgName(bytes32 _orgRef) view returns(string)
func (_IATIOrgs *IATIOrgsCallerSession) GetOrgName(_orgRef [32]byte) (string, error) {
	return _IATIOrgs.Contract.GetOrgName(&_IATIOrgs.CallOpts, _orgRef)
}

// GetOrgReference is a free data retrieval call binding the contract method 0x0dd75b22.
//
// Solidity: function getOrgReference(uint256 _index) view returns(bytes32)
func (_IATIOrgs *IATIOrgsCaller) GetOrgReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrgs.contract.Call(opts, out, "getOrgReference", _index)
	return *ret0, err
}

// GetOrgReference is a free data retrieval call binding the contract method 0x0dd75b22.
//
// Solidity: function getOrgReference(uint256 _index) view returns(bytes32)
func (_IATIOrgs *IATIOrgsSession) GetOrgReference(_index *big.Int) ([32]byte, error) {
	return _IATIOrgs.Contract.GetOrgReference(&_IATIOrgs.CallOpts, _index)
}

// GetOrgReference is a free data retrieval call binding the contract method 0x0dd75b22.
//
// Solidity: function getOrgReference(uint256 _index) view returns(bytes32)
func (_IATIOrgs *IATIOrgsCallerSession) GetOrgReference(_index *big.Int) ([32]byte, error) {
	return _IATIOrgs.Contract.GetOrgReference(&_IATIOrgs.CallOpts, _index)
}

// SetOrg is a paid mutator transaction binding the contract method 0xdf9a8bbe.
//
// Solidity: function setOrg(bytes32 _orgRef, OrgsOrg _org) returns()
func (_IATIOrgs *IATIOrgsTransactor) SetOrg(opts *bind.TransactOpts, _orgRef [32]byte, _org OrgsOrg) (*types.Transaction, error) {
	return _IATIOrgs.contract.Transact(opts, "setOrg", _orgRef, _org)
}

// SetOrg is a paid mutator transaction binding the contract method 0xdf9a8bbe.
//
// Solidity: function setOrg(bytes32 _orgRef, OrgsOrg _org) returns()
func (_IATIOrgs *IATIOrgsSession) SetOrg(_orgRef [32]byte, _org OrgsOrg) (*types.Transaction, error) {
	return _IATIOrgs.Contract.SetOrg(&_IATIOrgs.TransactOpts, _orgRef, _org)
}

// SetOrg is a paid mutator transaction binding the contract method 0xdf9a8bbe.
//
// Solidity: function setOrg(bytes32 _orgRef, OrgsOrg _org) returns()
func (_IATIOrgs *IATIOrgsTransactorSession) SetOrg(_orgRef [32]byte, _org OrgsOrg) (*types.Transaction, error) {
	return _IATIOrgs.Contract.SetOrg(&_IATIOrgs.TransactOpts, _orgRef, _org)
}

// IATIOrgsSetOrgIterator is returned from FilterSetOrg and is used to iterate over the raw logs and unpacked data for SetOrg events raised by the IATIOrgs contract.
type IATIOrgsSetOrgIterator struct {
	Event *IATIOrgsSetOrg // Event containing the contract specifics and raw log

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
func (it *IATIOrgsSetOrgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IATIOrgsSetOrg)
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
		it.Event = new(IATIOrgsSetOrg)
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
func (it *IATIOrgsSetOrgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IATIOrgsSetOrgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IATIOrgsSetOrg represents a SetOrg event raised by the IATIOrgs contract.
type IATIOrgsSetOrg struct {
	OrgRef [32]byte
	Org    OrgsOrg
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetOrg is a free log retrieval operation binding the contract event 0x7b20403f3c5b881e1e5698bac32df980c612a3fd2632c887b2665dd91c33c465.
//
// Solidity: event SetOrg(bytes32 _orgRef, OrgsOrg _org)
func (_IATIOrgs *IATIOrgsFilterer) FilterSetOrg(opts *bind.FilterOpts) (*IATIOrgsSetOrgIterator, error) {

	logs, sub, err := _IATIOrgs.contract.FilterLogs(opts, "SetOrg")
	if err != nil {
		return nil, err
	}
	return &IATIOrgsSetOrgIterator{contract: _IATIOrgs.contract, event: "SetOrg", logs: logs, sub: sub}, nil
}

// WatchSetOrg is a free log subscription operation binding the contract event 0x7b20403f3c5b881e1e5698bac32df980c612a3fd2632c887b2665dd91c33c465.
//
// Solidity: event SetOrg(bytes32 _orgRef, OrgsOrg _org)
func (_IATIOrgs *IATIOrgsFilterer) WatchSetOrg(opts *bind.WatchOpts, sink chan<- *IATIOrgsSetOrg) (event.Subscription, error) {

	logs, sub, err := _IATIOrgs.contract.WatchLogs(opts, "SetOrg")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IATIOrgsSetOrg)
				if err := _IATIOrgs.contract.UnpackLog(event, "SetOrg", log); err != nil {
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

// ParseSetOrg is a log parse operation binding the contract event 0x7b20403f3c5b881e1e5698bac32df980c612a3fd2632c887b2665dd91c33c465.
//
// Solidity: event SetOrg(bytes32 _orgRef, OrgsOrg _org)
func (_IATIOrgs *IATIOrgsFilterer) ParseSetOrg(log types.Log) (*IATIOrgsSetOrg, error) {
	event := new(IATIOrgsSetOrg)
	if err := _IATIOrgs.contract.UnpackLog(event, "SetOrg", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OrgsABI is the input ABI used to generate the binding from.
const OrgsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_orgRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structOrgs.Org\",\"name\":\"_org\",\"type\":\"tuple\"}],\"name\":\"SetOrg\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getNumOrgs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgRef\",\"type\":\"bytes32\"}],\"name\":\"getOrg\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"}],\"internalType\":\"structOrgs.Org\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgRef\",\"type\":\"bytes32\"}],\"name\":\"getOrgIdentifier\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgRef\",\"type\":\"bytes32\"}],\"name\":\"getOrgName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getOrgReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orgRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"}],\"internalType\":\"structOrgs.Org\",\"name\":\"_org\",\"type\":\"tuple\"}],\"name\":\"setOrg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OrgsFuncSigs maps the 4-byte function signature to its string representation.
var OrgsFuncSigs = map[string]string{
	"1a10f51c": "getNumOrgs()",
	"81e84589": "getOrg(bytes32)",
	"a4abc4ce": "getOrgIdentifier(bytes32)",
	"195541cc": "getOrgName(bytes32)",
	"0dd75b22": "getOrgReference(uint256)",
	"df9a8bbe": "setOrg(bytes32,(string,string))",
}

// Orgs is an auto generated Go binding around an Ethereum contract.
type Orgs struct {
	OrgsCaller     // Read-only binding to the contract
	OrgsTransactor // Write-only binding to the contract
	OrgsFilterer   // Log filterer for contract events
}

// OrgsCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrgsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrgsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrgsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrgsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrgsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrgsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrgsSession struct {
	Contract     *Orgs             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrgsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrgsCallerSession struct {
	Contract *OrgsCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OrgsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrgsTransactorSession struct {
	Contract     *OrgsTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrgsRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrgsRaw struct {
	Contract *Orgs // Generic contract binding to access the raw methods on
}

// OrgsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrgsCallerRaw struct {
	Contract *OrgsCaller // Generic read-only contract binding to access the raw methods on
}

// OrgsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrgsTransactorRaw struct {
	Contract *OrgsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrgs creates a new instance of Orgs, bound to a specific deployed contract.
func NewOrgs(address common.Address, backend bind.ContractBackend) (*Orgs, error) {
	contract, err := bindOrgs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Orgs{OrgsCaller: OrgsCaller{contract: contract}, OrgsTransactor: OrgsTransactor{contract: contract}, OrgsFilterer: OrgsFilterer{contract: contract}}, nil
}

// NewOrgsCaller creates a new read-only instance of Orgs, bound to a specific deployed contract.
func NewOrgsCaller(address common.Address, caller bind.ContractCaller) (*OrgsCaller, error) {
	contract, err := bindOrgs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrgsCaller{contract: contract}, nil
}

// NewOrgsTransactor creates a new write-only instance of Orgs, bound to a specific deployed contract.
func NewOrgsTransactor(address common.Address, transactor bind.ContractTransactor) (*OrgsTransactor, error) {
	contract, err := bindOrgs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrgsTransactor{contract: contract}, nil
}

// NewOrgsFilterer creates a new log filterer instance of Orgs, bound to a specific deployed contract.
func NewOrgsFilterer(address common.Address, filterer bind.ContractFilterer) (*OrgsFilterer, error) {
	contract, err := bindOrgs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrgsFilterer{contract: contract}, nil
}

// bindOrgs binds a generic wrapper to an already deployed contract.
func bindOrgs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrgsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orgs *OrgsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Orgs.Contract.OrgsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orgs *OrgsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orgs.Contract.OrgsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orgs *OrgsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orgs.Contract.OrgsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orgs *OrgsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Orgs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orgs *OrgsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orgs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orgs *OrgsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orgs.Contract.contract.Transact(opts, method, params...)
}

// GetNumOrgs is a free data retrieval call binding the contract method 0x1a10f51c.
//
// Solidity: function getNumOrgs() view returns(uint256)
func (_Orgs *OrgsCaller) GetNumOrgs(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Orgs.contract.Call(opts, out, "getNumOrgs")
	return *ret0, err
}

// GetNumOrgs is a free data retrieval call binding the contract method 0x1a10f51c.
//
// Solidity: function getNumOrgs() view returns(uint256)
func (_Orgs *OrgsSession) GetNumOrgs() (*big.Int, error) {
	return _Orgs.Contract.GetNumOrgs(&_Orgs.CallOpts)
}

// GetNumOrgs is a free data retrieval call binding the contract method 0x1a10f51c.
//
// Solidity: function getNumOrgs() view returns(uint256)
func (_Orgs *OrgsCallerSession) GetNumOrgs() (*big.Int, error) {
	return _Orgs.Contract.GetNumOrgs(&_Orgs.CallOpts)
}

// GetOrg is a free data retrieval call binding the contract method 0x81e84589.
//
// Solidity: function getOrg(bytes32 _orgRef) view returns(OrgsOrg)
func (_Orgs *OrgsCaller) GetOrg(opts *bind.CallOpts, _orgRef [32]byte) (OrgsOrg, error) {
	var (
		ret0 = new(OrgsOrg)
	)
	out := ret0
	err := _Orgs.contract.Call(opts, out, "getOrg", _orgRef)
	return *ret0, err
}

// GetOrg is a free data retrieval call binding the contract method 0x81e84589.
//
// Solidity: function getOrg(bytes32 _orgRef) view returns(OrgsOrg)
func (_Orgs *OrgsSession) GetOrg(_orgRef [32]byte) (OrgsOrg, error) {
	return _Orgs.Contract.GetOrg(&_Orgs.CallOpts, _orgRef)
}

// GetOrg is a free data retrieval call binding the contract method 0x81e84589.
//
// Solidity: function getOrg(bytes32 _orgRef) view returns(OrgsOrg)
func (_Orgs *OrgsCallerSession) GetOrg(_orgRef [32]byte) (OrgsOrg, error) {
	return _Orgs.Contract.GetOrg(&_Orgs.CallOpts, _orgRef)
}

// GetOrgIdentifier is a free data retrieval call binding the contract method 0xa4abc4ce.
//
// Solidity: function getOrgIdentifier(bytes32 _orgRef) view returns(string)
func (_Orgs *OrgsCaller) GetOrgIdentifier(opts *bind.CallOpts, _orgRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Orgs.contract.Call(opts, out, "getOrgIdentifier", _orgRef)
	return *ret0, err
}

// GetOrgIdentifier is a free data retrieval call binding the contract method 0xa4abc4ce.
//
// Solidity: function getOrgIdentifier(bytes32 _orgRef) view returns(string)
func (_Orgs *OrgsSession) GetOrgIdentifier(_orgRef [32]byte) (string, error) {
	return _Orgs.Contract.GetOrgIdentifier(&_Orgs.CallOpts, _orgRef)
}

// GetOrgIdentifier is a free data retrieval call binding the contract method 0xa4abc4ce.
//
// Solidity: function getOrgIdentifier(bytes32 _orgRef) view returns(string)
func (_Orgs *OrgsCallerSession) GetOrgIdentifier(_orgRef [32]byte) (string, error) {
	return _Orgs.Contract.GetOrgIdentifier(&_Orgs.CallOpts, _orgRef)
}

// GetOrgName is a free data retrieval call binding the contract method 0x195541cc.
//
// Solidity: function getOrgName(bytes32 _orgRef) view returns(string)
func (_Orgs *OrgsCaller) GetOrgName(opts *bind.CallOpts, _orgRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Orgs.contract.Call(opts, out, "getOrgName", _orgRef)
	return *ret0, err
}

// GetOrgName is a free data retrieval call binding the contract method 0x195541cc.
//
// Solidity: function getOrgName(bytes32 _orgRef) view returns(string)
func (_Orgs *OrgsSession) GetOrgName(_orgRef [32]byte) (string, error) {
	return _Orgs.Contract.GetOrgName(&_Orgs.CallOpts, _orgRef)
}

// GetOrgName is a free data retrieval call binding the contract method 0x195541cc.
//
// Solidity: function getOrgName(bytes32 _orgRef) view returns(string)
func (_Orgs *OrgsCallerSession) GetOrgName(_orgRef [32]byte) (string, error) {
	return _Orgs.Contract.GetOrgName(&_Orgs.CallOpts, _orgRef)
}

// GetOrgReference is a free data retrieval call binding the contract method 0x0dd75b22.
//
// Solidity: function getOrgReference(uint256 _index) view returns(bytes32)
func (_Orgs *OrgsCaller) GetOrgReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Orgs.contract.Call(opts, out, "getOrgReference", _index)
	return *ret0, err
}

// GetOrgReference is a free data retrieval call binding the contract method 0x0dd75b22.
//
// Solidity: function getOrgReference(uint256 _index) view returns(bytes32)
func (_Orgs *OrgsSession) GetOrgReference(_index *big.Int) ([32]byte, error) {
	return _Orgs.Contract.GetOrgReference(&_Orgs.CallOpts, _index)
}

// GetOrgReference is a free data retrieval call binding the contract method 0x0dd75b22.
//
// Solidity: function getOrgReference(uint256 _index) view returns(bytes32)
func (_Orgs *OrgsCallerSession) GetOrgReference(_index *big.Int) ([32]byte, error) {
	return _Orgs.Contract.GetOrgReference(&_Orgs.CallOpts, _index)
}

// SetOrg is a paid mutator transaction binding the contract method 0xdf9a8bbe.
//
// Solidity: function setOrg(bytes32 _orgRef, OrgsOrg _org) returns()
func (_Orgs *OrgsTransactor) SetOrg(opts *bind.TransactOpts, _orgRef [32]byte, _org OrgsOrg) (*types.Transaction, error) {
	return _Orgs.contract.Transact(opts, "setOrg", _orgRef, _org)
}

// SetOrg is a paid mutator transaction binding the contract method 0xdf9a8bbe.
//
// Solidity: function setOrg(bytes32 _orgRef, OrgsOrg _org) returns()
func (_Orgs *OrgsSession) SetOrg(_orgRef [32]byte, _org OrgsOrg) (*types.Transaction, error) {
	return _Orgs.Contract.SetOrg(&_Orgs.TransactOpts, _orgRef, _org)
}

// SetOrg is a paid mutator transaction binding the contract method 0xdf9a8bbe.
//
// Solidity: function setOrg(bytes32 _orgRef, OrgsOrg _org) returns()
func (_Orgs *OrgsTransactorSession) SetOrg(_orgRef [32]byte, _org OrgsOrg) (*types.Transaction, error) {
	return _Orgs.Contract.SetOrg(&_Orgs.TransactOpts, _orgRef, _org)
}

// OrgsSetOrgIterator is returned from FilterSetOrg and is used to iterate over the raw logs and unpacked data for SetOrg events raised by the Orgs contract.
type OrgsSetOrgIterator struct {
	Event *OrgsSetOrg // Event containing the contract specifics and raw log

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
func (it *OrgsSetOrgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrgsSetOrg)
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
		it.Event = new(OrgsSetOrg)
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
func (it *OrgsSetOrgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrgsSetOrgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrgsSetOrg represents a SetOrg event raised by the Orgs contract.
type OrgsSetOrg struct {
	OrgRef [32]byte
	Org    OrgsOrg
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetOrg is a free log retrieval operation binding the contract event 0x7b20403f3c5b881e1e5698bac32df980c612a3fd2632c887b2665dd91c33c465.
//
// Solidity: event SetOrg(bytes32 _orgRef, OrgsOrg _org)
func (_Orgs *OrgsFilterer) FilterSetOrg(opts *bind.FilterOpts) (*OrgsSetOrgIterator, error) {

	logs, sub, err := _Orgs.contract.FilterLogs(opts, "SetOrg")
	if err != nil {
		return nil, err
	}
	return &OrgsSetOrgIterator{contract: _Orgs.contract, event: "SetOrg", logs: logs, sub: sub}, nil
}

// WatchSetOrg is a free log subscription operation binding the contract event 0x7b20403f3c5b881e1e5698bac32df980c612a3fd2632c887b2665dd91c33c465.
//
// Solidity: event SetOrg(bytes32 _orgRef, OrgsOrg _org)
func (_Orgs *OrgsFilterer) WatchSetOrg(opts *bind.WatchOpts, sink chan<- *OrgsSetOrg) (event.Subscription, error) {

	logs, sub, err := _Orgs.contract.WatchLogs(opts, "SetOrg")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrgsSetOrg)
				if err := _Orgs.contract.UnpackLog(event, "SetOrg", log); err != nil {
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

// ParseSetOrg is a log parse operation binding the contract event 0x7b20403f3c5b881e1e5698bac32df980c612a3fd2632c887b2665dd91c33c465.
//
// Solidity: event SetOrg(bytes32 _orgRef, OrgsOrg _org)
func (_Orgs *OrgsFilterer) ParseSetOrg(log types.Log) (*OrgsSetOrg, error) {
	event := new(OrgsSetOrg)
	if err := _Orgs.contract.UnpackLog(event, "SetOrg", log); err != nil {
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
