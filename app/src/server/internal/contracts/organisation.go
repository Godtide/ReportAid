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

// OrganisationOrg is an auto generated low-level Go binding around an user-defined struct.
type OrganisationOrg struct {
	OrgRef          [32]byte
	ReportingOrg    OrganisationReportingOrg
	Lang            [32]byte
	Currency        [32]byte
	LastUpdatedTime [32]byte
}

// OrganisationReportingOrg is an auto generated low-level Go binding around an user-defined struct.
type OrganisationReportingOrg struct {
	OrgType     uint8
	IsSecondary bool
	OrgRef      [32]byte
}

// IATIOrganisationABI is the input ABI used to generate the binding from.
const IATIOrganisationABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSecondary\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"}],\"internalType\":\"structOrganisation.ReportingOrg\",\"name\":\"reportingOrg\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastUpdatedTime\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structOrganisation.Org\",\"name\":\"_org\",\"type\":\"tuple\"}],\"name\":\"SetOrganisation\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getCurrency\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getLang\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getLastUpdatedTime\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"}],\"name\":\"getNumOrganisations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getOrganisation\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSecondary\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"}],\"internalType\":\"structOrganisation.ReportingOrg\",\"name\":\"reportingOrg\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastUpdatedTime\",\"type\":\"bytes32\"}],\"internalType\":\"structOrganisation.Org\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getOrganisationOrg\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getOrganisationReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getReportingOrg\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getReportingOrgIsSecondary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getReportingOrgType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSecondary\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"}],\"internalType\":\"structOrganisation.ReportingOrg\",\"name\":\"reportingOrg\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastUpdatedTime\",\"type\":\"bytes32\"}],\"internalType\":\"structOrganisation.Org\",\"name\":\"_org\",\"type\":\"tuple\"}],\"name\":\"setOrganisation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IATIOrganisationFuncSigs maps the 4-byte function signature to its string representation.
var IATIOrganisationFuncSigs = map[string]string{
	"eef45eae": "getCurrency(bytes32,bytes32)",
	"aad88871": "getLang(bytes32,bytes32)",
	"7d78621b": "getLastUpdatedTime(bytes32,bytes32)",
	"29b537f7": "getNumOrganisations(bytes32)",
	"f302d1f3": "getOrganisation(bytes32,bytes32)",
	"654183e5": "getOrganisationOrg(bytes32,bytes32)",
	"f483d30e": "getOrganisationReference(bytes32,uint256)",
	"40c21069": "getReportingOrg(bytes32,bytes32)",
	"9748f6a9": "getReportingOrgIsSecondary(bytes32,bytes32)",
	"b871fbf2": "getReportingOrgType(bytes32,bytes32)",
	"8835cac2": "setOrganisation(bytes32,bytes32,(bytes32,(uint8,bool,bytes32),bytes32,bytes32,bytes32))",
}

// IATIOrganisationBin is the compiled bytecode used for deploying new contracts.
var IATIOrganisationBin = "0x608060405234801561001057600080fd5b50610ad1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80639748f6a9116100715780639748f6a914610125578063aad8887114610145578063b871fbf214610158578063eef45eae14610178578063f302d1f31461018b578063f483d30e146101ab576100a9565b806329b537f7146100ae57806340c21069146100d7578063654183e5146100ea5780637d78621b146100fd5780638835cac214610110575b600080fd5b6100c16100bc36600461086c565b6101be565b6040516100ce91906109ca565b60405180910390f35b6100c16100e5366004610884565b6101ec565b6100c16100f8366004610884565b610249565b6100c161010b366004610884565b61029f565b61012361011e3660046108a5565b6102f8565b005b610138610133366004610884565b610550565b6040516100ce91906109bf565b6100c1610153366004610884565b6105af565b61016b610166366004610884565b610608565b6040516100ce9190610a55565b6100c1610186366004610884565b610662565b61019e610199366004610884565b6106bb565b6040516100ce9190610a47565b6100c16101b9366004610961565b610774565b600081811a60f81b6001600160f81b0319166101d957600080fd5b5060009081526020819052604090205490565b600082811a60f81b6001600160f81b0319161580159061021b57508160001a60f81b6001600160f81b03191615155b61022457600080fd5b5060008281526001602090815260408083208484529091529020600201545b92915050565b600082811a60f81b6001600160f81b0319161580159061027857508160001a60f81b6001600160f81b03191615155b61028157600080fd5b50600091825260016020908152604080842092845291905290205490565b600082811a60f81b6001600160f81b031916158015906102ce57508160001a60f81b6001600160f81b03191615155b6102d757600080fd5b50600091825260016020908152604080842092845291905290206005015490565b8260001a60f81b6001600160f81b0319161580159061032657508160001a60f81b6001600160f81b03191615155b80156103425750805160001a60f81b6001600160f81b03191615155b8015610365575060208101516040015160001a60f81b6001600160f81b03191615155b8015610378575060208101515160ff1615155b80156103975750604081015160001a60f81b6001600160f81b03191615155b80156103b65750606081015160001a60f81b6001600160f81b03191615155b80156103d55750608081015160001a60f81b6001600160f81b03191615155b6103de57600080fd5b600083815260016020818152604080842086855282528084208551815582860151805194820180548286015115156101000261ff001960ff90981660ff19909216919091179690961695909517909455928101516002840155808501516003840155606085015160048085019190915560808601516005909401939093558684529083905291829020915163745fca7b60e01b815273__$b620240993a55615b607189176fb82e62f$__9263745fca7b9261049c92879291016109d3565b60206040518083038186803b1580156104b457600080fd5b505af41580156104c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ec9190610849565b61051057600083815260208181526040822080546001810182559083529120018290555b7fc17edce2da9d70d3db0e9fb63b714ef85de4c80ffb6729283a1f229cc40afb2083838360405161054393929190610a23565b60405180910390a1505050565b600082811a60f81b6001600160f81b0319161580159061057f57508160001a60f81b6001600160f81b03191615155b61058857600080fd5b50600091825260016020818152604080852093855292905291200154610100900460ff1690565b600082811a60f81b6001600160f81b031916158015906105de57508160001a60f81b6001600160f81b03191615155b6105e757600080fd5b50600091825260016020908152604080842092845291905290206003015490565b600082811a60f81b6001600160f81b0319161580159061063757508160001a60f81b6001600160f81b03191615155b61064057600080fd5b5060009182526001602081815260408085209385529290529120015460ff1690565b600082811a60f81b6001600160f81b0319161580159061069157508160001a60f81b6001600160f81b03191615155b61069a57600080fd5b50600091825260016020908152604080842092845291905290206004015490565b6106c36107d7565b8260001a60f81b6001600160f81b031916158015906106f157508160001a60f81b6001600160f81b03191615155b6106fa57600080fd5b50600091825260016020818152604080852093855292815292829020825160a081018452815481528351606080820186529383015460ff8082168352610100909104161515818701526002830154818601529481019490945260038101549284019290925260048201549083015260050154608082015290565b600082811a60f81b6001600160f81b031916158015906107a1575060008381526020819052604090205482105b6107aa57600080fd5b60008381526020819052604090208054839081106107c457fe5b9060005260206000200154905092915050565b6040805160a0810190915260008152602081016107f261080d565b81526000602082018190526040820181905260609091015290565b604080516060810182526000808252602082018190529181019190915290565b803561024381610a8a565b803560ff8116811461024357600080fd5b60006020828403121561085a578081fd5b815161086581610a8a565b9392505050565b60006020828403121561087d578081fd5b5035919050565b60008060408385031215610896578081fd5b50508035926020909101359150565b60008060008385036101208112156108bb578182fd5b843593506020850135925060e0603f19820112156108d7578182fd5b6108e160a0610a63565b604086013581526060605f19830112156108f9578283fd5b6109036060610a63565b91506109128760608801610838565b8252610921876080880161082d565b602083015260a0860135604083015281602082015260c0860135604082015260e08601356060820152610100860135608082015280925050509250925092565b60008060408385031215610896578182fd5b80518252602081015160ff8151166020840152602081015115156040840152604081015160608401525060408101516080830152606081015160a0830152608081015160c08301525050565b901515815260200190565b90815260200190565b60006040820184835260206040818501528185548084526060860191508685528285209350845b81811015610a16578454835260019485019492840192016109fa565b5090979650505050505050565b838152602081018390526101208101610a3f6040830184610973565b949350505050565b60e081016102438284610973565b60ff91909116815260200190565b60405181810167ffffffffffffffff81118282101715610a8257600080fd5b604052919050565b8015158114610a9857600080fd5b5056fea26469706673582212206ab40e8952a60f2c5244b378cb8b67821726564a6e0bde9404513766f8205af764736f6c63430006060033"

// DeployIATIOrganisation deploys a new Ethereum contract, binding an instance of IATIOrganisation to it.
func DeployIATIOrganisation(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IATIOrganisation, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIOrganisationABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	stringsAddr, _, _, _ := DeployStrings(auth, backend)
	IATIOrganisationBin = strings.Replace(IATIOrganisationBin, "__$b620240993a55615b607189176fb82e62f$__", stringsAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IATIOrganisationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IATIOrganisation{IATIOrganisationCaller: IATIOrganisationCaller{contract: contract}, IATIOrganisationTransactor: IATIOrganisationTransactor{contract: contract}, IATIOrganisationFilterer: IATIOrganisationFilterer{contract: contract}}, nil
}

// IATIOrganisation is an auto generated Go binding around an Ethereum contract.
type IATIOrganisation struct {
	IATIOrganisationCaller     // Read-only binding to the contract
	IATIOrganisationTransactor // Write-only binding to the contract
	IATIOrganisationFilterer   // Log filterer for contract events
}

// IATIOrganisationCaller is an auto generated read-only Go binding around an Ethereum contract.
type IATIOrganisationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIOrganisationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IATIOrganisationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIOrganisationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IATIOrganisationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIOrganisationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IATIOrganisationSession struct {
	Contract     *IATIOrganisation // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IATIOrganisationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IATIOrganisationCallerSession struct {
	Contract *IATIOrganisationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IATIOrganisationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IATIOrganisationTransactorSession struct {
	Contract     *IATIOrganisationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IATIOrganisationRaw is an auto generated low-level Go binding around an Ethereum contract.
type IATIOrganisationRaw struct {
	Contract *IATIOrganisation // Generic contract binding to access the raw methods on
}

// IATIOrganisationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IATIOrganisationCallerRaw struct {
	Contract *IATIOrganisationCaller // Generic read-only contract binding to access the raw methods on
}

// IATIOrganisationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IATIOrganisationTransactorRaw struct {
	Contract *IATIOrganisationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIATIOrganisation creates a new instance of IATIOrganisation, bound to a specific deployed contract.
func NewIATIOrganisation(address common.Address, backend bind.ContractBackend) (*IATIOrganisation, error) {
	contract, err := bindIATIOrganisation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IATIOrganisation{IATIOrganisationCaller: IATIOrganisationCaller{contract: contract}, IATIOrganisationTransactor: IATIOrganisationTransactor{contract: contract}, IATIOrganisationFilterer: IATIOrganisationFilterer{contract: contract}}, nil
}

// NewIATIOrganisationCaller creates a new read-only instance of IATIOrganisation, bound to a specific deployed contract.
func NewIATIOrganisationCaller(address common.Address, caller bind.ContractCaller) (*IATIOrganisationCaller, error) {
	contract, err := bindIATIOrganisation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IATIOrganisationCaller{contract: contract}, nil
}

// NewIATIOrganisationTransactor creates a new write-only instance of IATIOrganisation, bound to a specific deployed contract.
func NewIATIOrganisationTransactor(address common.Address, transactor bind.ContractTransactor) (*IATIOrganisationTransactor, error) {
	contract, err := bindIATIOrganisation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IATIOrganisationTransactor{contract: contract}, nil
}

// NewIATIOrganisationFilterer creates a new log filterer instance of IATIOrganisation, bound to a specific deployed contract.
func NewIATIOrganisationFilterer(address common.Address, filterer bind.ContractFilterer) (*IATIOrganisationFilterer, error) {
	contract, err := bindIATIOrganisation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IATIOrganisationFilterer{contract: contract}, nil
}

// bindIATIOrganisation binds a generic wrapper to an already deployed contract.
func bindIATIOrganisation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIOrganisationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIOrganisation *IATIOrganisationRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIOrganisation.Contract.IATIOrganisationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIOrganisation *IATIOrganisationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIOrganisation.Contract.IATIOrganisationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIOrganisation *IATIOrganisationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIOrganisation.Contract.IATIOrganisationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIOrganisation *IATIOrganisationCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIOrganisation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIOrganisation *IATIOrganisationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIOrganisation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIOrganisation *IATIOrganisationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIOrganisation.Contract.contract.Transact(opts, method, params...)
}

// GetCurrency is a free data retrieval call binding the contract method 0xeef45eae.
//
// Solidity: function getCurrency(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_IATIOrganisation *IATIOrganisationCaller) GetCurrency(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisation.contract.Call(opts, out, "getCurrency", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetCurrency is a free data retrieval call binding the contract method 0xeef45eae.
//
// Solidity: function getCurrency(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_IATIOrganisation *IATIOrganisationSession) GetCurrency(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _IATIOrganisation.Contract.GetCurrency(&_IATIOrganisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetCurrency is a free data retrieval call binding the contract method 0xeef45eae.
//
// Solidity: function getCurrency(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_IATIOrganisation *IATIOrganisationCallerSession) GetCurrency(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _IATIOrganisation.Contract.GetCurrency(&_IATIOrganisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetLang is a free data retrieval call binding the contract method 0xaad88871.
//
// Solidity: function getLang(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_IATIOrganisation *IATIOrganisationCaller) GetLang(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisation.contract.Call(opts, out, "getLang", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetLang is a free data retrieval call binding the contract method 0xaad88871.
//
// Solidity: function getLang(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_IATIOrganisation *IATIOrganisationSession) GetLang(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _IATIOrganisation.Contract.GetLang(&_IATIOrganisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetLang is a free data retrieval call binding the contract method 0xaad88871.
//
// Solidity: function getLang(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_IATIOrganisation *IATIOrganisationCallerSession) GetLang(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _IATIOrganisation.Contract.GetLang(&_IATIOrganisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetLastUpdatedTime is a free data retrieval call binding the contract method 0x7d78621b.
//
// Solidity: function getLastUpdatedTime(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_IATIOrganisation *IATIOrganisationCaller) GetLastUpdatedTime(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisation.contract.Call(opts, out, "getLastUpdatedTime", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetLastUpdatedTime is a free data retrieval call binding the contract method 0x7d78621b.
//
// Solidity: function getLastUpdatedTime(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_IATIOrganisation *IATIOrganisationSession) GetLastUpdatedTime(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _IATIOrganisation.Contract.GetLastUpdatedTime(&_IATIOrganisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetLastUpdatedTime is a free data retrieval call binding the contract method 0x7d78621b.
//
// Solidity: function getLastUpdatedTime(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_IATIOrganisation *IATIOrganisationCallerSession) GetLastUpdatedTime(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _IATIOrganisation.Contract.GetLastUpdatedTime(&_IATIOrganisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetNumOrganisations is a free data retrieval call binding the contract method 0x29b537f7.
//
// Solidity: function getNumOrganisations(bytes32 _organisationsRef) view returns(uint256)
func (_IATIOrganisation *IATIOrganisationCaller) GetNumOrganisations(opts *bind.CallOpts, _organisationsRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIOrganisation.contract.Call(opts, out, "getNumOrganisations", _organisationsRef)
	return *ret0, err
}

// GetNumOrganisations is a free data retrieval call binding the contract method 0x29b537f7.
//
// Solidity: function getNumOrganisations(bytes32 _organisationsRef) view returns(uint256)
func (_IATIOrganisation *IATIOrganisationSession) GetNumOrganisations(_organisationsRef [32]byte) (*big.Int, error) {
	return _IATIOrganisation.Contract.GetNumOrganisations(&_IATIOrganisation.CallOpts, _organisationsRef)
}

// GetNumOrganisations is a free data retrieval call binding the contract method 0x29b537f7.
//
// Solidity: function getNumOrganisations(bytes32 _organisationsRef) view returns(uint256)
func (_IATIOrganisation *IATIOrganisationCallerSession) GetNumOrganisations(_organisationsRef [32]byte) (*big.Int, error) {
	return _IATIOrganisation.Contract.GetNumOrganisations(&_IATIOrganisation.CallOpts, _organisationsRef)
}

// GetOrganisation is a free data retrieval call binding the contract method 0xf302d1f3.
//
// Solidity: function getOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(OrganisationOrg)
func (_IATIOrganisation *IATIOrganisationCaller) GetOrganisation(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) (OrganisationOrg, error) {
	var (
		ret0 = new(OrganisationOrg)
	)
	out := ret0
	err := _IATIOrganisation.contract.Call(opts, out, "getOrganisation", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetOrganisation is a free data retrieval call binding the contract method 0xf302d1f3.
//
// Solidity: function getOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(OrganisationOrg)
func (_IATIOrganisation *IATIOrganisationSession) GetOrganisation(_organisationsRef [32]byte, _organisationRef [32]byte) (OrganisationOrg, error) {
	return _IATIOrganisation.Contract.GetOrganisation(&_IATIOrganisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetOrganisation is a free data retrieval call binding the contract method 0xf302d1f3.
//
// Solidity: function getOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(OrganisationOrg)
func (_IATIOrganisation *IATIOrganisationCallerSession) GetOrganisation(_organisationsRef [32]byte, _organisationRef [32]byte) (OrganisationOrg, error) {
	return _IATIOrganisation.Contract.GetOrganisation(&_IATIOrganisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetOrganisationOrg is a free data retrieval call binding the contract method 0x654183e5.
//
// Solidity: function getOrganisationOrg(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_IATIOrganisation *IATIOrganisationCaller) GetOrganisationOrg(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisation.contract.Call(opts, out, "getOrganisationOrg", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetOrganisationOrg is a free data retrieval call binding the contract method 0x654183e5.
//
// Solidity: function getOrganisationOrg(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_IATIOrganisation *IATIOrganisationSession) GetOrganisationOrg(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _IATIOrganisation.Contract.GetOrganisationOrg(&_IATIOrganisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetOrganisationOrg is a free data retrieval call binding the contract method 0x654183e5.
//
// Solidity: function getOrganisationOrg(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_IATIOrganisation *IATIOrganisationCallerSession) GetOrganisationOrg(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _IATIOrganisation.Contract.GetOrganisationOrg(&_IATIOrganisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetOrganisationReference is a free data retrieval call binding the contract method 0xf483d30e.
//
// Solidity: function getOrganisationReference(bytes32 _organisationsRef, uint256 _index) view returns(bytes32)
func (_IATIOrganisation *IATIOrganisationCaller) GetOrganisationReference(opts *bind.CallOpts, _organisationsRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisation.contract.Call(opts, out, "getOrganisationReference", _organisationsRef, _index)
	return *ret0, err
}

// GetOrganisationReference is a free data retrieval call binding the contract method 0xf483d30e.
//
// Solidity: function getOrganisationReference(bytes32 _organisationsRef, uint256 _index) view returns(bytes32)
func (_IATIOrganisation *IATIOrganisationSession) GetOrganisationReference(_organisationsRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIOrganisation.Contract.GetOrganisationReference(&_IATIOrganisation.CallOpts, _organisationsRef, _index)
}

// GetOrganisationReference is a free data retrieval call binding the contract method 0xf483d30e.
//
// Solidity: function getOrganisationReference(bytes32 _organisationsRef, uint256 _index) view returns(bytes32)
func (_IATIOrganisation *IATIOrganisationCallerSession) GetOrganisationReference(_organisationsRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIOrganisation.Contract.GetOrganisationReference(&_IATIOrganisation.CallOpts, _organisationsRef, _index)
}

// GetReportingOrg is a free data retrieval call binding the contract method 0x40c21069.
//
// Solidity: function getReportingOrg(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_IATIOrganisation *IATIOrganisationCaller) GetReportingOrg(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisation.contract.Call(opts, out, "getReportingOrg", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetReportingOrg is a free data retrieval call binding the contract method 0x40c21069.
//
// Solidity: function getReportingOrg(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_IATIOrganisation *IATIOrganisationSession) GetReportingOrg(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _IATIOrganisation.Contract.GetReportingOrg(&_IATIOrganisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetReportingOrg is a free data retrieval call binding the contract method 0x40c21069.
//
// Solidity: function getReportingOrg(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_IATIOrganisation *IATIOrganisationCallerSession) GetReportingOrg(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _IATIOrganisation.Contract.GetReportingOrg(&_IATIOrganisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetReportingOrgIsSecondary is a free data retrieval call binding the contract method 0x9748f6a9.
//
// Solidity: function getReportingOrgIsSecondary(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bool)
func (_IATIOrganisation *IATIOrganisationCaller) GetReportingOrgIsSecondary(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IATIOrganisation.contract.Call(opts, out, "getReportingOrgIsSecondary", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetReportingOrgIsSecondary is a free data retrieval call binding the contract method 0x9748f6a9.
//
// Solidity: function getReportingOrgIsSecondary(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bool)
func (_IATIOrganisation *IATIOrganisationSession) GetReportingOrgIsSecondary(_organisationsRef [32]byte, _organisationRef [32]byte) (bool, error) {
	return _IATIOrganisation.Contract.GetReportingOrgIsSecondary(&_IATIOrganisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetReportingOrgIsSecondary is a free data retrieval call binding the contract method 0x9748f6a9.
//
// Solidity: function getReportingOrgIsSecondary(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bool)
func (_IATIOrganisation *IATIOrganisationCallerSession) GetReportingOrgIsSecondary(_organisationsRef [32]byte, _organisationRef [32]byte) (bool, error) {
	return _IATIOrganisation.Contract.GetReportingOrgIsSecondary(&_IATIOrganisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetReportingOrgType is a free data retrieval call binding the contract method 0xb871fbf2.
//
// Solidity: function getReportingOrgType(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(uint8)
func (_IATIOrganisation *IATIOrganisationCaller) GetReportingOrgType(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IATIOrganisation.contract.Call(opts, out, "getReportingOrgType", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetReportingOrgType is a free data retrieval call binding the contract method 0xb871fbf2.
//
// Solidity: function getReportingOrgType(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(uint8)
func (_IATIOrganisation *IATIOrganisationSession) GetReportingOrgType(_organisationsRef [32]byte, _organisationRef [32]byte) (uint8, error) {
	return _IATIOrganisation.Contract.GetReportingOrgType(&_IATIOrganisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetReportingOrgType is a free data retrieval call binding the contract method 0xb871fbf2.
//
// Solidity: function getReportingOrgType(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(uint8)
func (_IATIOrganisation *IATIOrganisationCallerSession) GetReportingOrgType(_organisationsRef [32]byte, _organisationRef [32]byte) (uint8, error) {
	return _IATIOrganisation.Contract.GetReportingOrgType(&_IATIOrganisation.CallOpts, _organisationsRef, _organisationRef)
}

// SetOrganisation is a paid mutator transaction binding the contract method 0x8835cac2.
//
// Solidity: function setOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef, OrganisationOrg _org) returns()
func (_IATIOrganisation *IATIOrganisationTransactor) SetOrganisation(opts *bind.TransactOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _org OrganisationOrg) (*types.Transaction, error) {
	return _IATIOrganisation.contract.Transact(opts, "setOrganisation", _organisationsRef, _organisationRef, _org)
}

// SetOrganisation is a paid mutator transaction binding the contract method 0x8835cac2.
//
// Solidity: function setOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef, OrganisationOrg _org) returns()
func (_IATIOrganisation *IATIOrganisationSession) SetOrganisation(_organisationsRef [32]byte, _organisationRef [32]byte, _org OrganisationOrg) (*types.Transaction, error) {
	return _IATIOrganisation.Contract.SetOrganisation(&_IATIOrganisation.TransactOpts, _organisationsRef, _organisationRef, _org)
}

// SetOrganisation is a paid mutator transaction binding the contract method 0x8835cac2.
//
// Solidity: function setOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef, OrganisationOrg _org) returns()
func (_IATIOrganisation *IATIOrganisationTransactorSession) SetOrganisation(_organisationsRef [32]byte, _organisationRef [32]byte, _org OrganisationOrg) (*types.Transaction, error) {
	return _IATIOrganisation.Contract.SetOrganisation(&_IATIOrganisation.TransactOpts, _organisationsRef, _organisationRef, _org)
}

// IATIOrganisationSetOrganisationIterator is returned from FilterSetOrganisation and is used to iterate over the raw logs and unpacked data for SetOrganisation events raised by the IATIOrganisation contract.
type IATIOrganisationSetOrganisationIterator struct {
	Event *IATIOrganisationSetOrganisation // Event containing the contract specifics and raw log

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
func (it *IATIOrganisationSetOrganisationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IATIOrganisationSetOrganisation)
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
		it.Event = new(IATIOrganisationSetOrganisation)
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
func (it *IATIOrganisationSetOrganisationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IATIOrganisationSetOrganisationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IATIOrganisationSetOrganisation represents a SetOrganisation event raised by the IATIOrganisation contract.
type IATIOrganisationSetOrganisation struct {
	OrganisationsRef [32]byte
	OrganisationRef  [32]byte
	Org              OrganisationOrg
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetOrganisation is a free log retrieval operation binding the contract event 0xc17edce2da9d70d3db0e9fb63b714ef85de4c80ffb6729283a1f229cc40afb20.
//
// Solidity: event SetOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef, OrganisationOrg _org)
func (_IATIOrganisation *IATIOrganisationFilterer) FilterSetOrganisation(opts *bind.FilterOpts) (*IATIOrganisationSetOrganisationIterator, error) {

	logs, sub, err := _IATIOrganisation.contract.FilterLogs(opts, "SetOrganisation")
	if err != nil {
		return nil, err
	}
	return &IATIOrganisationSetOrganisationIterator{contract: _IATIOrganisation.contract, event: "SetOrganisation", logs: logs, sub: sub}, nil
}

// WatchSetOrganisation is a free log subscription operation binding the contract event 0xc17edce2da9d70d3db0e9fb63b714ef85de4c80ffb6729283a1f229cc40afb20.
//
// Solidity: event SetOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef, OrganisationOrg _org)
func (_IATIOrganisation *IATIOrganisationFilterer) WatchSetOrganisation(opts *bind.WatchOpts, sink chan<- *IATIOrganisationSetOrganisation) (event.Subscription, error) {

	logs, sub, err := _IATIOrganisation.contract.WatchLogs(opts, "SetOrganisation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IATIOrganisationSetOrganisation)
				if err := _IATIOrganisation.contract.UnpackLog(event, "SetOrganisation", log); err != nil {
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

// ParseSetOrganisation is a log parse operation binding the contract event 0xc17edce2da9d70d3db0e9fb63b714ef85de4c80ffb6729283a1f229cc40afb20.
//
// Solidity: event SetOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef, OrganisationOrg _org)
func (_IATIOrganisation *IATIOrganisationFilterer) ParseSetOrganisation(log types.Log) (*IATIOrganisationSetOrganisation, error) {
	event := new(IATIOrganisationSetOrganisation)
	if err := _IATIOrganisation.contract.UnpackLog(event, "SetOrganisation", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OrganisationABI is the input ABI used to generate the binding from.
const OrganisationABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSecondary\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"}],\"internalType\":\"structOrganisation.ReportingOrg\",\"name\":\"reportingOrg\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastUpdatedTime\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structOrganisation.Org\",\"name\":\"_org\",\"type\":\"tuple\"}],\"name\":\"SetOrganisation\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getCurrency\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getLang\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getLastUpdatedTime\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"}],\"name\":\"getNumOrganisations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getOrganisation\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSecondary\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"}],\"internalType\":\"structOrganisation.ReportingOrg\",\"name\":\"reportingOrg\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastUpdatedTime\",\"type\":\"bytes32\"}],\"internalType\":\"structOrganisation.Org\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getOrganisationOrg\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getOrganisationReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getReportingOrg\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getReportingOrgIsSecondary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getReportingOrgType\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSecondary\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"}],\"internalType\":\"structOrganisation.ReportingOrg\",\"name\":\"reportingOrg\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastUpdatedTime\",\"type\":\"bytes32\"}],\"internalType\":\"structOrganisation.Org\",\"name\":\"_org\",\"type\":\"tuple\"}],\"name\":\"setOrganisation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OrganisationFuncSigs maps the 4-byte function signature to its string representation.
var OrganisationFuncSigs = map[string]string{
	"eef45eae": "getCurrency(bytes32,bytes32)",
	"aad88871": "getLang(bytes32,bytes32)",
	"7d78621b": "getLastUpdatedTime(bytes32,bytes32)",
	"29b537f7": "getNumOrganisations(bytes32)",
	"f302d1f3": "getOrganisation(bytes32,bytes32)",
	"654183e5": "getOrganisationOrg(bytes32,bytes32)",
	"f483d30e": "getOrganisationReference(bytes32,uint256)",
	"40c21069": "getReportingOrg(bytes32,bytes32)",
	"9748f6a9": "getReportingOrgIsSecondary(bytes32,bytes32)",
	"b871fbf2": "getReportingOrgType(bytes32,bytes32)",
	"8835cac2": "setOrganisation(bytes32,bytes32,(bytes32,(uint8,bool,bytes32),bytes32,bytes32,bytes32))",
}

// Organisation is an auto generated Go binding around an Ethereum contract.
type Organisation struct {
	OrganisationCaller     // Read-only binding to the contract
	OrganisationTransactor // Write-only binding to the contract
	OrganisationFilterer   // Log filterer for contract events
}

// OrganisationCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrganisationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganisationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrganisationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganisationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrganisationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganisationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrganisationSession struct {
	Contract     *Organisation     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrganisationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrganisationCallerSession struct {
	Contract *OrganisationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OrganisationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrganisationTransactorSession struct {
	Contract     *OrganisationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OrganisationRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrganisationRaw struct {
	Contract *Organisation // Generic contract binding to access the raw methods on
}

// OrganisationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrganisationCallerRaw struct {
	Contract *OrganisationCaller // Generic read-only contract binding to access the raw methods on
}

// OrganisationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrganisationTransactorRaw struct {
	Contract *OrganisationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrganisation creates a new instance of Organisation, bound to a specific deployed contract.
func NewOrganisation(address common.Address, backend bind.ContractBackend) (*Organisation, error) {
	contract, err := bindOrganisation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Organisation{OrganisationCaller: OrganisationCaller{contract: contract}, OrganisationTransactor: OrganisationTransactor{contract: contract}, OrganisationFilterer: OrganisationFilterer{contract: contract}}, nil
}

// NewOrganisationCaller creates a new read-only instance of Organisation, bound to a specific deployed contract.
func NewOrganisationCaller(address common.Address, caller bind.ContractCaller) (*OrganisationCaller, error) {
	contract, err := bindOrganisation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrganisationCaller{contract: contract}, nil
}

// NewOrganisationTransactor creates a new write-only instance of Organisation, bound to a specific deployed contract.
func NewOrganisationTransactor(address common.Address, transactor bind.ContractTransactor) (*OrganisationTransactor, error) {
	contract, err := bindOrganisation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrganisationTransactor{contract: contract}, nil
}

// NewOrganisationFilterer creates a new log filterer instance of Organisation, bound to a specific deployed contract.
func NewOrganisationFilterer(address common.Address, filterer bind.ContractFilterer) (*OrganisationFilterer, error) {
	contract, err := bindOrganisation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrganisationFilterer{contract: contract}, nil
}

// bindOrganisation binds a generic wrapper to an already deployed contract.
func bindOrganisation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrganisationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Organisation *OrganisationRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Organisation.Contract.OrganisationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Organisation *OrganisationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Organisation.Contract.OrganisationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Organisation *OrganisationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Organisation.Contract.OrganisationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Organisation *OrganisationCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Organisation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Organisation *OrganisationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Organisation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Organisation *OrganisationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Organisation.Contract.contract.Transact(opts, method, params...)
}

// GetCurrency is a free data retrieval call binding the contract method 0xeef45eae.
//
// Solidity: function getCurrency(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_Organisation *OrganisationCaller) GetCurrency(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Organisation.contract.Call(opts, out, "getCurrency", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetCurrency is a free data retrieval call binding the contract method 0xeef45eae.
//
// Solidity: function getCurrency(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_Organisation *OrganisationSession) GetCurrency(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _Organisation.Contract.GetCurrency(&_Organisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetCurrency is a free data retrieval call binding the contract method 0xeef45eae.
//
// Solidity: function getCurrency(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_Organisation *OrganisationCallerSession) GetCurrency(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _Organisation.Contract.GetCurrency(&_Organisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetLang is a free data retrieval call binding the contract method 0xaad88871.
//
// Solidity: function getLang(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_Organisation *OrganisationCaller) GetLang(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Organisation.contract.Call(opts, out, "getLang", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetLang is a free data retrieval call binding the contract method 0xaad88871.
//
// Solidity: function getLang(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_Organisation *OrganisationSession) GetLang(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _Organisation.Contract.GetLang(&_Organisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetLang is a free data retrieval call binding the contract method 0xaad88871.
//
// Solidity: function getLang(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_Organisation *OrganisationCallerSession) GetLang(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _Organisation.Contract.GetLang(&_Organisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetLastUpdatedTime is a free data retrieval call binding the contract method 0x7d78621b.
//
// Solidity: function getLastUpdatedTime(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_Organisation *OrganisationCaller) GetLastUpdatedTime(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Organisation.contract.Call(opts, out, "getLastUpdatedTime", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetLastUpdatedTime is a free data retrieval call binding the contract method 0x7d78621b.
//
// Solidity: function getLastUpdatedTime(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_Organisation *OrganisationSession) GetLastUpdatedTime(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _Organisation.Contract.GetLastUpdatedTime(&_Organisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetLastUpdatedTime is a free data retrieval call binding the contract method 0x7d78621b.
//
// Solidity: function getLastUpdatedTime(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_Organisation *OrganisationCallerSession) GetLastUpdatedTime(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _Organisation.Contract.GetLastUpdatedTime(&_Organisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetNumOrganisations is a free data retrieval call binding the contract method 0x29b537f7.
//
// Solidity: function getNumOrganisations(bytes32 _organisationsRef) view returns(uint256)
func (_Organisation *OrganisationCaller) GetNumOrganisations(opts *bind.CallOpts, _organisationsRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Organisation.contract.Call(opts, out, "getNumOrganisations", _organisationsRef)
	return *ret0, err
}

// GetNumOrganisations is a free data retrieval call binding the contract method 0x29b537f7.
//
// Solidity: function getNumOrganisations(bytes32 _organisationsRef) view returns(uint256)
func (_Organisation *OrganisationSession) GetNumOrganisations(_organisationsRef [32]byte) (*big.Int, error) {
	return _Organisation.Contract.GetNumOrganisations(&_Organisation.CallOpts, _organisationsRef)
}

// GetNumOrganisations is a free data retrieval call binding the contract method 0x29b537f7.
//
// Solidity: function getNumOrganisations(bytes32 _organisationsRef) view returns(uint256)
func (_Organisation *OrganisationCallerSession) GetNumOrganisations(_organisationsRef [32]byte) (*big.Int, error) {
	return _Organisation.Contract.GetNumOrganisations(&_Organisation.CallOpts, _organisationsRef)
}

// GetOrganisation is a free data retrieval call binding the contract method 0xf302d1f3.
//
// Solidity: function getOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(OrganisationOrg)
func (_Organisation *OrganisationCaller) GetOrganisation(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) (OrganisationOrg, error) {
	var (
		ret0 = new(OrganisationOrg)
	)
	out := ret0
	err := _Organisation.contract.Call(opts, out, "getOrganisation", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetOrganisation is a free data retrieval call binding the contract method 0xf302d1f3.
//
// Solidity: function getOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(OrganisationOrg)
func (_Organisation *OrganisationSession) GetOrganisation(_organisationsRef [32]byte, _organisationRef [32]byte) (OrganisationOrg, error) {
	return _Organisation.Contract.GetOrganisation(&_Organisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetOrganisation is a free data retrieval call binding the contract method 0xf302d1f3.
//
// Solidity: function getOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(OrganisationOrg)
func (_Organisation *OrganisationCallerSession) GetOrganisation(_organisationsRef [32]byte, _organisationRef [32]byte) (OrganisationOrg, error) {
	return _Organisation.Contract.GetOrganisation(&_Organisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetOrganisationOrg is a free data retrieval call binding the contract method 0x654183e5.
//
// Solidity: function getOrganisationOrg(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_Organisation *OrganisationCaller) GetOrganisationOrg(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Organisation.contract.Call(opts, out, "getOrganisationOrg", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetOrganisationOrg is a free data retrieval call binding the contract method 0x654183e5.
//
// Solidity: function getOrganisationOrg(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_Organisation *OrganisationSession) GetOrganisationOrg(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _Organisation.Contract.GetOrganisationOrg(&_Organisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetOrganisationOrg is a free data retrieval call binding the contract method 0x654183e5.
//
// Solidity: function getOrganisationOrg(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_Organisation *OrganisationCallerSession) GetOrganisationOrg(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _Organisation.Contract.GetOrganisationOrg(&_Organisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetOrganisationReference is a free data retrieval call binding the contract method 0xf483d30e.
//
// Solidity: function getOrganisationReference(bytes32 _organisationsRef, uint256 _index) view returns(bytes32)
func (_Organisation *OrganisationCaller) GetOrganisationReference(opts *bind.CallOpts, _organisationsRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Organisation.contract.Call(opts, out, "getOrganisationReference", _organisationsRef, _index)
	return *ret0, err
}

// GetOrganisationReference is a free data retrieval call binding the contract method 0xf483d30e.
//
// Solidity: function getOrganisationReference(bytes32 _organisationsRef, uint256 _index) view returns(bytes32)
func (_Organisation *OrganisationSession) GetOrganisationReference(_organisationsRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _Organisation.Contract.GetOrganisationReference(&_Organisation.CallOpts, _organisationsRef, _index)
}

// GetOrganisationReference is a free data retrieval call binding the contract method 0xf483d30e.
//
// Solidity: function getOrganisationReference(bytes32 _organisationsRef, uint256 _index) view returns(bytes32)
func (_Organisation *OrganisationCallerSession) GetOrganisationReference(_organisationsRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _Organisation.Contract.GetOrganisationReference(&_Organisation.CallOpts, _organisationsRef, _index)
}

// GetReportingOrg is a free data retrieval call binding the contract method 0x40c21069.
//
// Solidity: function getReportingOrg(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_Organisation *OrganisationCaller) GetReportingOrg(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Organisation.contract.Call(opts, out, "getReportingOrg", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetReportingOrg is a free data retrieval call binding the contract method 0x40c21069.
//
// Solidity: function getReportingOrg(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_Organisation *OrganisationSession) GetReportingOrg(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _Organisation.Contract.GetReportingOrg(&_Organisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetReportingOrg is a free data retrieval call binding the contract method 0x40c21069.
//
// Solidity: function getReportingOrg(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bytes32)
func (_Organisation *OrganisationCallerSession) GetReportingOrg(_organisationsRef [32]byte, _organisationRef [32]byte) ([32]byte, error) {
	return _Organisation.Contract.GetReportingOrg(&_Organisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetReportingOrgIsSecondary is a free data retrieval call binding the contract method 0x9748f6a9.
//
// Solidity: function getReportingOrgIsSecondary(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bool)
func (_Organisation *OrganisationCaller) GetReportingOrgIsSecondary(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Organisation.contract.Call(opts, out, "getReportingOrgIsSecondary", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetReportingOrgIsSecondary is a free data retrieval call binding the contract method 0x9748f6a9.
//
// Solidity: function getReportingOrgIsSecondary(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bool)
func (_Organisation *OrganisationSession) GetReportingOrgIsSecondary(_organisationsRef [32]byte, _organisationRef [32]byte) (bool, error) {
	return _Organisation.Contract.GetReportingOrgIsSecondary(&_Organisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetReportingOrgIsSecondary is a free data retrieval call binding the contract method 0x9748f6a9.
//
// Solidity: function getReportingOrgIsSecondary(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(bool)
func (_Organisation *OrganisationCallerSession) GetReportingOrgIsSecondary(_organisationsRef [32]byte, _organisationRef [32]byte) (bool, error) {
	return _Organisation.Contract.GetReportingOrgIsSecondary(&_Organisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetReportingOrgType is a free data retrieval call binding the contract method 0xb871fbf2.
//
// Solidity: function getReportingOrgType(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(uint8)
func (_Organisation *OrganisationCaller) GetReportingOrgType(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Organisation.contract.Call(opts, out, "getReportingOrgType", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetReportingOrgType is a free data retrieval call binding the contract method 0xb871fbf2.
//
// Solidity: function getReportingOrgType(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(uint8)
func (_Organisation *OrganisationSession) GetReportingOrgType(_organisationsRef [32]byte, _organisationRef [32]byte) (uint8, error) {
	return _Organisation.Contract.GetReportingOrgType(&_Organisation.CallOpts, _organisationsRef, _organisationRef)
}

// GetReportingOrgType is a free data retrieval call binding the contract method 0xb871fbf2.
//
// Solidity: function getReportingOrgType(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(uint8)
func (_Organisation *OrganisationCallerSession) GetReportingOrgType(_organisationsRef [32]byte, _organisationRef [32]byte) (uint8, error) {
	return _Organisation.Contract.GetReportingOrgType(&_Organisation.CallOpts, _organisationsRef, _organisationRef)
}

// SetOrganisation is a paid mutator transaction binding the contract method 0x8835cac2.
//
// Solidity: function setOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef, OrganisationOrg _org) returns()
func (_Organisation *OrganisationTransactor) SetOrganisation(opts *bind.TransactOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _org OrganisationOrg) (*types.Transaction, error) {
	return _Organisation.contract.Transact(opts, "setOrganisation", _organisationsRef, _organisationRef, _org)
}

// SetOrganisation is a paid mutator transaction binding the contract method 0x8835cac2.
//
// Solidity: function setOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef, OrganisationOrg _org) returns()
func (_Organisation *OrganisationSession) SetOrganisation(_organisationsRef [32]byte, _organisationRef [32]byte, _org OrganisationOrg) (*types.Transaction, error) {
	return _Organisation.Contract.SetOrganisation(&_Organisation.TransactOpts, _organisationsRef, _organisationRef, _org)
}

// SetOrganisation is a paid mutator transaction binding the contract method 0x8835cac2.
//
// Solidity: function setOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef, OrganisationOrg _org) returns()
func (_Organisation *OrganisationTransactorSession) SetOrganisation(_organisationsRef [32]byte, _organisationRef [32]byte, _org OrganisationOrg) (*types.Transaction, error) {
	return _Organisation.Contract.SetOrganisation(&_Organisation.TransactOpts, _organisationsRef, _organisationRef, _org)
}

// OrganisationSetOrganisationIterator is returned from FilterSetOrganisation and is used to iterate over the raw logs and unpacked data for SetOrganisation events raised by the Organisation contract.
type OrganisationSetOrganisationIterator struct {
	Event *OrganisationSetOrganisation // Event containing the contract specifics and raw log

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
func (it *OrganisationSetOrganisationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrganisationSetOrganisation)
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
		it.Event = new(OrganisationSetOrganisation)
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
func (it *OrganisationSetOrganisationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrganisationSetOrganisationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrganisationSetOrganisation represents a SetOrganisation event raised by the Organisation contract.
type OrganisationSetOrganisation struct {
	OrganisationsRef [32]byte
	OrganisationRef  [32]byte
	Org              OrganisationOrg
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetOrganisation is a free log retrieval operation binding the contract event 0xc17edce2da9d70d3db0e9fb63b714ef85de4c80ffb6729283a1f229cc40afb20.
//
// Solidity: event SetOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef, OrganisationOrg _org)
func (_Organisation *OrganisationFilterer) FilterSetOrganisation(opts *bind.FilterOpts) (*OrganisationSetOrganisationIterator, error) {

	logs, sub, err := _Organisation.contract.FilterLogs(opts, "SetOrganisation")
	if err != nil {
		return nil, err
	}
	return &OrganisationSetOrganisationIterator{contract: _Organisation.contract, event: "SetOrganisation", logs: logs, sub: sub}, nil
}

// WatchSetOrganisation is a free log subscription operation binding the contract event 0xc17edce2da9d70d3db0e9fb63b714ef85de4c80ffb6729283a1f229cc40afb20.
//
// Solidity: event SetOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef, OrganisationOrg _org)
func (_Organisation *OrganisationFilterer) WatchSetOrganisation(opts *bind.WatchOpts, sink chan<- *OrganisationSetOrganisation) (event.Subscription, error) {

	logs, sub, err := _Organisation.contract.WatchLogs(opts, "SetOrganisation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrganisationSetOrganisation)
				if err := _Organisation.contract.UnpackLog(event, "SetOrganisation", log); err != nil {
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

// ParseSetOrganisation is a log parse operation binding the contract event 0xc17edce2da9d70d3db0e9fb63b714ef85de4c80ffb6729283a1f229cc40afb20.
//
// Solidity: event SetOrganisation(bytes32 _organisationsRef, bytes32 _organisationRef, OrganisationOrg _org)
func (_Organisation *OrganisationFilterer) ParseSetOrganisation(log types.Log) (*OrganisationSetOrganisation, error) {
	event := new(OrganisationSetOrganisation)
	if err := _Organisation.contract.UnpackLog(event, "SetOrganisation", log); err != nil {
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
