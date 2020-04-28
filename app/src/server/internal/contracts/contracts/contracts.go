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

// OrganisationDocsDocument is an auto generated low-level Go binding around an user-defined struct.
type OrganisationDocsDocument struct {
	Title      string
	Format     string
	Url        string
	Category   [32]byte
	CountryRef [32]byte
	Desc       string
	Lang       [32]byte
	Date       [32]byte
}

// IATIOrganisationDocsABI is the input ABI used to generate the binding from.
const IATIOrganisationDocsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"format\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"category\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"countryRef\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structOrganisationDocs.Document\",\"name\":\"_document\",\"type\":\"tuple\"}],\"name\":\"SetDocument\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getDocReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"}],\"name\":\"getDocument\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"format\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"category\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"countryRef\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"}],\"internalType\":\"structOrganisationDocs.Document\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"}],\"name\":\"getDocumentCategory\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"}],\"name\":\"getDocumentCountry\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"}],\"name\":\"getDocumentDate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"}],\"name\":\"getDocumentDescription\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"}],\"name\":\"getDocumentFormat\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"}],\"name\":\"getDocumentLang\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"}],\"name\":\"getDocumentTitle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"}],\"name\":\"getDocumentURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getNumDocs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"format\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"category\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"countryRef\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"}],\"internalType\":\"structOrganisationDocs.Document\",\"name\":\"_document\",\"type\":\"tuple\"}],\"name\":\"setDocument\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IATIOrganisationDocsFuncSigs maps the 4-byte function signature to its string representation.
var IATIOrganisationDocsFuncSigs = map[string]string{
	"132ac373": "getDocReference(bytes32,bytes32,uint256)",
	"966c2b10": "getDocument(bytes32,bytes32,bytes32)",
	"07e6d576": "getDocumentCategory(bytes32,bytes32,bytes32)",
	"02e06d15": "getDocumentCountry(bytes32,bytes32,bytes32)",
	"efedacb6": "getDocumentDate(bytes32,bytes32,bytes32)",
	"28accf53": "getDocumentDescription(bytes32,bytes32,bytes32)",
	"3330a30b": "getDocumentFormat(bytes32,bytes32,bytes32)",
	"ffbe65db": "getDocumentLang(bytes32,bytes32,bytes32)",
	"1690c748": "getDocumentTitle(bytes32,bytes32,bytes32)",
	"98dfafc3": "getDocumentURL(bytes32,bytes32,bytes32)",
	"b11c464e": "getNumDocs(bytes32,bytes32)",
	"0ac17bad": "setDocument(bytes32,bytes32,bytes32,(string,string,string,bytes32,bytes32,string,bytes32,bytes32))",
}

// IATIOrganisationDocsBin is the compiled bytecode used for deploying new contracts.
var IATIOrganisationDocsBin = "0x608060405234801561001057600080fd5b506112d3806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80633330a30b116100715780633330a30b14610150578063966c2b101461016357806398dfafc314610183578063b11c464e14610196578063efedacb6146101a9578063ffbe65db146101bc576100b4565b806302e06d15146100b957806307e6d576146100e25780630ac17bad146100f5578063132ac3731461010a5780631690c7481461011d57806328accf531461013d575b600080fd5b6100cc6100c7366004610f84565b6101cf565b6040516100d991906111c8565b60405180910390f35b6100cc6100f0366004610f84565b61024a565b610108610103366004610faf565b6102c5565b005b6100cc6101183660046110cf565b610592565b61013061012b366004610f84565b610623565b6040516100d99190611250565b61013061014b366004610f84565b610727565b61013061015e366004610f84565b6107f5565b610176610171366004610f84565b6108c2565b6040516100d99190611263565b610130610191366004610f84565b610bdc565b6100cc6101a4366004610f63565b610ca8565b6100cc6101b7366004610f84565b610cfc565b6100cc6101ca366004610f84565b610d77565b600083811a60f81b6001600160f81b031916158015906101fe57508260001a60f81b6001600160f81b03191615155b801561021957508160001a60f81b6001600160f81b03191615155b61022257600080fd5b5060009283526001602090815260408085209385529281528284209184525290206004015490565b600083811a60f81b6001600160f81b0319161580159061027957508260001a60f81b6001600160f81b03191615155b801561029457508160001a60f81b6001600160f81b03191615155b61029d57600080fd5b5060009283526001602090815260408085209385529281528284209184525290206003015490565b8360001a60f81b6001600160f81b031916158015906102f357508260001a60f81b6001600160f81b03191615155b801561030e57508160001a60f81b6001600160f81b03191615155b801561031b575080515115155b801561032c57506000816020015151115b801561033d57506000816040015151115b801561035c5750606081015160001a60f81b6001600160f81b03191615155b801561037b5750608081015160001a60f81b6001600160f81b03191615155b801561038c575060008160a0015151115b80156103ab575060c081015160001a60f81b6001600160f81b03191615155b80156103ca575060e081015160001a60f81b6001600160f81b03191615155b6103d357600080fd5b600084815260016020908152604080832086845282528083208584528252909120825180518493610408928492910190610df2565b5060208281015180516104219260018501920190610df2565b506040820151805161043d916002840191602090910190610df2565b50606082015160038201556080820151600482015560a0820151805161046d916005840191602090910190610df2565b5060c0820151600682015560e09091015160079091015560008481526020818152604080832086845290915290819020905163745fca7b60e01b815273__$b620240993a55615b607189176fb82e62f$__9163745fca7b916104d39186916004016111d1565b60206040518083038186803b1580156104eb57600080fd5b505af41580156104ff573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105239190610f3c565b61054f576000848152602081815260408083208684528252822080546001810182559083529120018290555b7fd3ebf2a2f6f8f9157c17b01687f0f13175fdcd58ac706593f24cf342fab1b024848484846040516105849493929190611221565b60405180910390a150505050565b600083811a60f81b6001600160f81b031916158015906105c157508260001a60f81b6001600160f81b03191615155b80156105e3575060008481526020818152604080832086845290915290205482105b6105ec57600080fd5b600084815260208181526040808320868452909152902080548390811061060f57fe5b906000526020600020015490509392505050565b60608360001a60f81b6001600160f81b0319161580159061065357508260001a60f81b6001600160f81b03191615155b801561066e57508160001a60f81b6001600160f81b03191615155b61067757600080fd5b6000848152600160208181526040808420878552825280842086855282529283902080548451600294821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156107195780601f106106ee57610100808354040283529160200191610719565b820191906000526020600020905b8154815290600101906020018083116106fc57829003601f168201915b505050505090509392505050565b60608360001a60f81b6001600160f81b0319161580159061075757508260001a60f81b6001600160f81b03191615155b801561077257508160001a60f81b6001600160f81b03191615155b61077b57600080fd5b6000848152600160208181526040808420878552825280842086855282529283902060050180548451600294821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156107195780601f106106ee57610100808354040283529160200191610719565b60608360001a60f81b6001600160f81b0319161580159061082557508260001a60f81b6001600160f81b03191615155b801561084057508160001a60f81b6001600160f81b03191615155b61084957600080fd5b60008481526001602081815260408084208785528252808420868552825292839020820180548451600294821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156107195780601f106106ee57610100808354040283529160200191610719565b6108ca610e70565b8360001a60f81b6001600160f81b031916158015906108f857508260001a60f81b6001600160f81b03191615155b801561091357508160001a60f81b6001600160f81b03191615155b61091c57600080fd5b600084815260016020818152604080842087855282528084208685528252928390208351815460026101009582161586026000190190911604601f8101849004909302810161012090810190955292830182815292939092849290918491908401828280156109cc5780601f106109a1576101008083540402835291602001916109cc565b820191906000526020600020905b8154815290600101906020018083116109af57829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a6e5780601f10610a4357610100808354040283529160200191610a6e565b820191906000526020600020905b815481529060010190602001808311610a5157829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f81018390048302850183019091528084529381019390830182828015610b005780601f10610ad557610100808354040283529160200191610b00565b820191906000526020600020905b815481529060010190602001808311610ae357829003601f168201915b505050505081526020016003820154815260200160048201548152602001600582018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610bb65780601f10610b8b57610100808354040283529160200191610bb6565b820191906000526020600020905b815481529060010190602001808311610b9957829003601f168201915b505050505081526020016006820154815260200160078201548152505090509392505050565b60608360001a60f81b6001600160f81b03191615801590610c0c57508260001a60f81b6001600160f81b03191615155b8015610c2757508160001a60f81b6001600160f81b03191615155b610c3057600080fd5b600084815260016020818152604080842087855282528084208685528252928390206002908101805485519481161561010002600019011691909104601f81018390048302840183019094528383529192908301828280156107195780601f106106ee57610100808354040283529160200191610719565b600082811a60f81b6001600160f81b03191615801590610cd757508160001a60f81b6001600160f81b03191615155b610ce057600080fd5b5060009182526020828152604080842092845291905290205490565b600083811a60f81b6001600160f81b03191615801590610d2b57508260001a60f81b6001600160f81b03191615155b8015610d4657508160001a60f81b6001600160f81b03191615155b610d4f57600080fd5b5060009283526001602090815260408085209385529281528284209184525290206007015490565b600083811a60f81b6001600160f81b03191615801590610da657508260001a60f81b6001600160f81b03191615155b8015610dc157508160001a60f81b6001600160f81b03191615155b610dca57600080fd5b5060009283526001602090815260408085209385529281528284209184525290206006015490565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610e3357805160ff1916838001178555610e60565b82800160010185558215610e60579182015b82811115610e60578251825591602001919060010190610e45565b50610e6c929150610eb6565b5090565b604080516101008101825260608082526020820181905291810182905260008282018190526080820181905260a082019290925260c0810182905260e081019190915290565b610ed091905b80821115610e6c5760008155600101610ebc565b90565b600082601f830112610ee3578081fd5b813567ffffffffffffffff811115610ef9578182fd5b610f0c601f8201601f1916602001611276565b9150808252836020828501011115610f2357600080fd5b8060208401602084013760009082016020015292915050565b600060208284031215610f4d578081fd5b81518015158114610f5c578182fd5b9392505050565b60008060408385031215610f75578081fd5b50508035926020909101359150565b600080600060608486031215610f98578081fd5b505081359360208301359350604090920135919050565b60008060008060808587031215610fc4578081fd5b843593506020850135925060408501359150606085013567ffffffffffffffff80821115610ff0578283fd5b610100918701808903831315611004578384fd5b61100d83611276565b813593508284111561101d578485fd5b6110298a858401610ed3565b8152602082013593508284111561103e578485fd5b61104a8a858401610ed3565b60208201526040820135935082841115611062578485fd5b61106e8a858401610ed3565b6040820152606082013560608201526080820135608082015260a082013593508284111561109a578485fd5b6110a68a858401610ed3565b60a082015260c082013560c082015260e082013560e08201528094505050505092959194509250565b600080600060608486031215610f98578283fd5b60008151808452815b81811015611108576020818501810151868301820152016110ec565b818111156111195782602083870101525b50601f01601f19169290920160200192915050565b60006101008251818552611144828601826110e3565b60208501519250858103602087015261115d81846110e3565b91505060408401519150848103604086015261117981836110e3565b606085015160608701526080850151608087015260a0850151925085810360a08701526111a681846110e3565b91505060c084015160c086015260e084015160e0860152809250505092915050565b90815260200190565b60006040820184835260206040818501528185548084526060860191508685528285209350845b81811015611214578454835260019485019492840192016111f8565b5090979650505050505050565b600085825284602083015283604083015260806060830152611246608083018461112e565b9695505050505050565b600060208252610f5c60208301846110e3565b600060208252610f5c602083018461112e565b60405181810167ffffffffffffffff8111828210171561129557600080fd5b60405291905056fea2646970667358221220a65d4d4acf25e24eb277108865276df77f4330500db6d0cc92018600dafd3c9264736f6c63430006060033"

// DeployIATIOrganisationDocs deploys a new Ethereum contract, binding an instance of IATIOrganisationDocs to it.
func DeployIATIOrganisationDocs(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IATIOrganisationDocs, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIOrganisationDocsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	stringsAddr, _, _, _ := DeployStrings(auth, backend)
	IATIOrganisationDocsBin = strings.Replace(IATIOrganisationDocsBin, "__$b620240993a55615b607189176fb82e62f$__", stringsAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IATIOrganisationDocsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IATIOrganisationDocs{IATIOrganisationDocsCaller: IATIOrganisationDocsCaller{contract: contract}, IATIOrganisationDocsTransactor: IATIOrganisationDocsTransactor{contract: contract}, IATIOrganisationDocsFilterer: IATIOrganisationDocsFilterer{contract: contract}}, nil
}

// IATIOrganisationDocs is an auto generated Go binding around an Ethereum contract.
type IATIOrganisationDocs struct {
	IATIOrganisationDocsCaller     // Read-only binding to the contract
	IATIOrganisationDocsTransactor // Write-only binding to the contract
	IATIOrganisationDocsFilterer   // Log filterer for contract events
}

// IATIOrganisationDocsCaller is an auto generated read-only Go binding around an Ethereum contract.
type IATIOrganisationDocsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIOrganisationDocsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IATIOrganisationDocsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIOrganisationDocsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IATIOrganisationDocsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IATIOrganisationDocsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IATIOrganisationDocsSession struct {
	Contract     *IATIOrganisationDocs // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IATIOrganisationDocsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IATIOrganisationDocsCallerSession struct {
	Contract *IATIOrganisationDocsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IATIOrganisationDocsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IATIOrganisationDocsTransactorSession struct {
	Contract     *IATIOrganisationDocsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IATIOrganisationDocsRaw is an auto generated low-level Go binding around an Ethereum contract.
type IATIOrganisationDocsRaw struct {
	Contract *IATIOrganisationDocs // Generic contract binding to access the raw methods on
}

// IATIOrganisationDocsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IATIOrganisationDocsCallerRaw struct {
	Contract *IATIOrganisationDocsCaller // Generic read-only contract binding to access the raw methods on
}

// IATIOrganisationDocsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IATIOrganisationDocsTransactorRaw struct {
	Contract *IATIOrganisationDocsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIATIOrganisationDocs creates a new instance of IATIOrganisationDocs, bound to a specific deployed contract.
func NewIATIOrganisationDocs(address common.Address, backend bind.ContractBackend) (*IATIOrganisationDocs, error) {
	contract, err := bindIATIOrganisationDocs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IATIOrganisationDocs{IATIOrganisationDocsCaller: IATIOrganisationDocsCaller{contract: contract}, IATIOrganisationDocsTransactor: IATIOrganisationDocsTransactor{contract: contract}, IATIOrganisationDocsFilterer: IATIOrganisationDocsFilterer{contract: contract}}, nil
}

// NewIATIOrganisationDocsCaller creates a new read-only instance of IATIOrganisationDocs, bound to a specific deployed contract.
func NewIATIOrganisationDocsCaller(address common.Address, caller bind.ContractCaller) (*IATIOrganisationDocsCaller, error) {
	contract, err := bindIATIOrganisationDocs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IATIOrganisationDocsCaller{contract: contract}, nil
}

// NewIATIOrganisationDocsTransactor creates a new write-only instance of IATIOrganisationDocs, bound to a specific deployed contract.
func NewIATIOrganisationDocsTransactor(address common.Address, transactor bind.ContractTransactor) (*IATIOrganisationDocsTransactor, error) {
	contract, err := bindIATIOrganisationDocs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IATIOrganisationDocsTransactor{contract: contract}, nil
}

// NewIATIOrganisationDocsFilterer creates a new log filterer instance of IATIOrganisationDocs, bound to a specific deployed contract.
func NewIATIOrganisationDocsFilterer(address common.Address, filterer bind.ContractFilterer) (*IATIOrganisationDocsFilterer, error) {
	contract, err := bindIATIOrganisationDocs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IATIOrganisationDocsFilterer{contract: contract}, nil
}

// bindIATIOrganisationDocs binds a generic wrapper to an already deployed contract.
func bindIATIOrganisationDocs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IATIOrganisationDocsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIOrganisationDocs *IATIOrganisationDocsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIOrganisationDocs.Contract.IATIOrganisationDocsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIOrganisationDocs *IATIOrganisationDocsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIOrganisationDocs.Contract.IATIOrganisationDocsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIOrganisationDocs *IATIOrganisationDocsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIOrganisationDocs.Contract.IATIOrganisationDocsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IATIOrganisationDocs *IATIOrganisationDocsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IATIOrganisationDocs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IATIOrganisationDocs *IATIOrganisationDocsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IATIOrganisationDocs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IATIOrganisationDocs *IATIOrganisationDocsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IATIOrganisationDocs.Contract.contract.Transact(opts, method, params...)
}

// GetDocReference is a free data retrieval call binding the contract method 0x132ac373.
//
// Solidity: function getDocReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) view returns(bytes32)
func (_IATIOrganisationDocs *IATIOrganisationDocsCaller) GetDocReference(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisationDocs.contract.Call(opts, out, "getDocReference", _organisationsRef, _organisationRef, _index)
	return *ret0, err
}

// GetDocReference is a free data retrieval call binding the contract method 0x132ac373.
//
// Solidity: function getDocReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) view returns(bytes32)
func (_IATIOrganisationDocs *IATIOrganisationDocsSession) GetDocReference(_organisationsRef [32]byte, _organisationRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIOrganisationDocs.Contract.GetDocReference(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _index)
}

// GetDocReference is a free data retrieval call binding the contract method 0x132ac373.
//
// Solidity: function getDocReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) view returns(bytes32)
func (_IATIOrganisationDocs *IATIOrganisationDocsCallerSession) GetDocReference(_organisationsRef [32]byte, _organisationRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _IATIOrganisationDocs.Contract.GetDocReference(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _index)
}

// GetDocument is a free data retrieval call binding the contract method 0x966c2b10.
//
// Solidity: function getDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(OrganisationDocsDocument)
func (_IATIOrganisationDocs *IATIOrganisationDocsCaller) GetDocument(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (OrganisationDocsDocument, error) {
	var (
		ret0 = new(OrganisationDocsDocument)
	)
	out := ret0
	err := _IATIOrganisationDocs.contract.Call(opts, out, "getDocument", _organisationsRef, _organisationRef, _docRef)
	return *ret0, err
}

// GetDocument is a free data retrieval call binding the contract method 0x966c2b10.
//
// Solidity: function getDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(OrganisationDocsDocument)
func (_IATIOrganisationDocs *IATIOrganisationDocsSession) GetDocument(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (OrganisationDocsDocument, error) {
	return _IATIOrganisationDocs.Contract.GetDocument(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocument is a free data retrieval call binding the contract method 0x966c2b10.
//
// Solidity: function getDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(OrganisationDocsDocument)
func (_IATIOrganisationDocs *IATIOrganisationDocsCallerSession) GetDocument(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (OrganisationDocsDocument, error) {
	return _IATIOrganisationDocs.Contract.GetDocument(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentCategory is a free data retrieval call binding the contract method 0x07e6d576.
//
// Solidity: function getDocumentCategory(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_IATIOrganisationDocs *IATIOrganisationDocsCaller) GetDocumentCategory(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisationDocs.contract.Call(opts, out, "getDocumentCategory", _organisationsRef, _organisationRef, _docRef)
	return *ret0, err
}

// GetDocumentCategory is a free data retrieval call binding the contract method 0x07e6d576.
//
// Solidity: function getDocumentCategory(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_IATIOrganisationDocs *IATIOrganisationDocsSession) GetDocumentCategory(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	return _IATIOrganisationDocs.Contract.GetDocumentCategory(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentCategory is a free data retrieval call binding the contract method 0x07e6d576.
//
// Solidity: function getDocumentCategory(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_IATIOrganisationDocs *IATIOrganisationDocsCallerSession) GetDocumentCategory(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	return _IATIOrganisationDocs.Contract.GetDocumentCategory(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentCountry is a free data retrieval call binding the contract method 0x02e06d15.
//
// Solidity: function getDocumentCountry(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_IATIOrganisationDocs *IATIOrganisationDocsCaller) GetDocumentCountry(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisationDocs.contract.Call(opts, out, "getDocumentCountry", _organisationsRef, _organisationRef, _docRef)
	return *ret0, err
}

// GetDocumentCountry is a free data retrieval call binding the contract method 0x02e06d15.
//
// Solidity: function getDocumentCountry(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_IATIOrganisationDocs *IATIOrganisationDocsSession) GetDocumentCountry(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	return _IATIOrganisationDocs.Contract.GetDocumentCountry(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentCountry is a free data retrieval call binding the contract method 0x02e06d15.
//
// Solidity: function getDocumentCountry(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_IATIOrganisationDocs *IATIOrganisationDocsCallerSession) GetDocumentCountry(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	return _IATIOrganisationDocs.Contract.GetDocumentCountry(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentDate is a free data retrieval call binding the contract method 0xefedacb6.
//
// Solidity: function getDocumentDate(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_IATIOrganisationDocs *IATIOrganisationDocsCaller) GetDocumentDate(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisationDocs.contract.Call(opts, out, "getDocumentDate", _organisationsRef, _organisationRef, _docRef)
	return *ret0, err
}

// GetDocumentDate is a free data retrieval call binding the contract method 0xefedacb6.
//
// Solidity: function getDocumentDate(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_IATIOrganisationDocs *IATIOrganisationDocsSession) GetDocumentDate(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	return _IATIOrganisationDocs.Contract.GetDocumentDate(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentDate is a free data retrieval call binding the contract method 0xefedacb6.
//
// Solidity: function getDocumentDate(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_IATIOrganisationDocs *IATIOrganisationDocsCallerSession) GetDocumentDate(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	return _IATIOrganisationDocs.Contract.GetDocumentDate(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentDescription is a free data retrieval call binding the contract method 0x28accf53.
//
// Solidity: function getDocumentDescription(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_IATIOrganisationDocs *IATIOrganisationDocsCaller) GetDocumentDescription(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IATIOrganisationDocs.contract.Call(opts, out, "getDocumentDescription", _organisationsRef, _organisationRef, _docRef)
	return *ret0, err
}

// GetDocumentDescription is a free data retrieval call binding the contract method 0x28accf53.
//
// Solidity: function getDocumentDescription(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_IATIOrganisationDocs *IATIOrganisationDocsSession) GetDocumentDescription(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	return _IATIOrganisationDocs.Contract.GetDocumentDescription(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentDescription is a free data retrieval call binding the contract method 0x28accf53.
//
// Solidity: function getDocumentDescription(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_IATIOrganisationDocs *IATIOrganisationDocsCallerSession) GetDocumentDescription(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	return _IATIOrganisationDocs.Contract.GetDocumentDescription(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentFormat is a free data retrieval call binding the contract method 0x3330a30b.
//
// Solidity: function getDocumentFormat(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_IATIOrganisationDocs *IATIOrganisationDocsCaller) GetDocumentFormat(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IATIOrganisationDocs.contract.Call(opts, out, "getDocumentFormat", _organisationsRef, _organisationRef, _docRef)
	return *ret0, err
}

// GetDocumentFormat is a free data retrieval call binding the contract method 0x3330a30b.
//
// Solidity: function getDocumentFormat(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_IATIOrganisationDocs *IATIOrganisationDocsSession) GetDocumentFormat(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	return _IATIOrganisationDocs.Contract.GetDocumentFormat(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentFormat is a free data retrieval call binding the contract method 0x3330a30b.
//
// Solidity: function getDocumentFormat(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_IATIOrganisationDocs *IATIOrganisationDocsCallerSession) GetDocumentFormat(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	return _IATIOrganisationDocs.Contract.GetDocumentFormat(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentLang is a free data retrieval call binding the contract method 0xffbe65db.
//
// Solidity: function getDocumentLang(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_IATIOrganisationDocs *IATIOrganisationDocsCaller) GetDocumentLang(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IATIOrganisationDocs.contract.Call(opts, out, "getDocumentLang", _organisationsRef, _organisationRef, _docRef)
	return *ret0, err
}

// GetDocumentLang is a free data retrieval call binding the contract method 0xffbe65db.
//
// Solidity: function getDocumentLang(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_IATIOrganisationDocs *IATIOrganisationDocsSession) GetDocumentLang(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	return _IATIOrganisationDocs.Contract.GetDocumentLang(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentLang is a free data retrieval call binding the contract method 0xffbe65db.
//
// Solidity: function getDocumentLang(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_IATIOrganisationDocs *IATIOrganisationDocsCallerSession) GetDocumentLang(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	return _IATIOrganisationDocs.Contract.GetDocumentLang(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentTitle is a free data retrieval call binding the contract method 0x1690c748.
//
// Solidity: function getDocumentTitle(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_IATIOrganisationDocs *IATIOrganisationDocsCaller) GetDocumentTitle(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IATIOrganisationDocs.contract.Call(opts, out, "getDocumentTitle", _organisationsRef, _organisationRef, _docRef)
	return *ret0, err
}

// GetDocumentTitle is a free data retrieval call binding the contract method 0x1690c748.
//
// Solidity: function getDocumentTitle(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_IATIOrganisationDocs *IATIOrganisationDocsSession) GetDocumentTitle(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	return _IATIOrganisationDocs.Contract.GetDocumentTitle(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentTitle is a free data retrieval call binding the contract method 0x1690c748.
//
// Solidity: function getDocumentTitle(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_IATIOrganisationDocs *IATIOrganisationDocsCallerSession) GetDocumentTitle(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	return _IATIOrganisationDocs.Contract.GetDocumentTitle(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentURL is a free data retrieval call binding the contract method 0x98dfafc3.
//
// Solidity: function getDocumentURL(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_IATIOrganisationDocs *IATIOrganisationDocsCaller) GetDocumentURL(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IATIOrganisationDocs.contract.Call(opts, out, "getDocumentURL", _organisationsRef, _organisationRef, _docRef)
	return *ret0, err
}

// GetDocumentURL is a free data retrieval call binding the contract method 0x98dfafc3.
//
// Solidity: function getDocumentURL(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_IATIOrganisationDocs *IATIOrganisationDocsSession) GetDocumentURL(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	return _IATIOrganisationDocs.Contract.GetDocumentURL(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentURL is a free data retrieval call binding the contract method 0x98dfafc3.
//
// Solidity: function getDocumentURL(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_IATIOrganisationDocs *IATIOrganisationDocsCallerSession) GetDocumentURL(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	return _IATIOrganisationDocs.Contract.GetDocumentURL(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetNumDocs is a free data retrieval call binding the contract method 0xb11c464e.
//
// Solidity: function getNumDocs(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(uint256)
func (_IATIOrganisationDocs *IATIOrganisationDocsCaller) GetNumDocs(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IATIOrganisationDocs.contract.Call(opts, out, "getNumDocs", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetNumDocs is a free data retrieval call binding the contract method 0xb11c464e.
//
// Solidity: function getNumDocs(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(uint256)
func (_IATIOrganisationDocs *IATIOrganisationDocsSession) GetNumDocs(_organisationsRef [32]byte, _organisationRef [32]byte) (*big.Int, error) {
	return _IATIOrganisationDocs.Contract.GetNumDocs(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef)
}

// GetNumDocs is a free data retrieval call binding the contract method 0xb11c464e.
//
// Solidity: function getNumDocs(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(uint256)
func (_IATIOrganisationDocs *IATIOrganisationDocsCallerSession) GetNumDocs(_organisationsRef [32]byte, _organisationRef [32]byte) (*big.Int, error) {
	return _IATIOrganisationDocs.Contract.GetNumDocs(&_IATIOrganisationDocs.CallOpts, _organisationsRef, _organisationRef)
}

// SetDocument is a paid mutator transaction binding the contract method 0x0ac17bad.
//
// Solidity: function setDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef, OrganisationDocsDocument _document) returns()
func (_IATIOrganisationDocs *IATIOrganisationDocsTransactor) SetDocument(opts *bind.TransactOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte, _document OrganisationDocsDocument) (*types.Transaction, error) {
	return _IATIOrganisationDocs.contract.Transact(opts, "setDocument", _organisationsRef, _organisationRef, _docRef, _document)
}

// SetDocument is a paid mutator transaction binding the contract method 0x0ac17bad.
//
// Solidity: function setDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef, OrganisationDocsDocument _document) returns()
func (_IATIOrganisationDocs *IATIOrganisationDocsSession) SetDocument(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte, _document OrganisationDocsDocument) (*types.Transaction, error) {
	return _IATIOrganisationDocs.Contract.SetDocument(&_IATIOrganisationDocs.TransactOpts, _organisationsRef, _organisationRef, _docRef, _document)
}

// SetDocument is a paid mutator transaction binding the contract method 0x0ac17bad.
//
// Solidity: function setDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef, OrganisationDocsDocument _document) returns()
func (_IATIOrganisationDocs *IATIOrganisationDocsTransactorSession) SetDocument(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte, _document OrganisationDocsDocument) (*types.Transaction, error) {
	return _IATIOrganisationDocs.Contract.SetDocument(&_IATIOrganisationDocs.TransactOpts, _organisationsRef, _organisationRef, _docRef, _document)
}

// IATIOrganisationDocsSetDocumentIterator is returned from FilterSetDocument and is used to iterate over the raw logs and unpacked data for SetDocument events raised by the IATIOrganisationDocs contract.
type IATIOrganisationDocsSetDocumentIterator struct {
	Event *IATIOrganisationDocsSetDocument // Event containing the contract specifics and raw log

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
func (it *IATIOrganisationDocsSetDocumentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IATIOrganisationDocsSetDocument)
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
		it.Event = new(IATIOrganisationDocsSetDocument)
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
func (it *IATIOrganisationDocsSetDocumentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IATIOrganisationDocsSetDocumentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IATIOrganisationDocsSetDocument represents a SetDocument event raised by the IATIOrganisationDocs contract.
type IATIOrganisationDocsSetDocument struct {
	OrganisationsRef [32]byte
	OrganisationRef  [32]byte
	DocRef           [32]byte
	Document         OrganisationDocsDocument
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetDocument is a free log retrieval operation binding the contract event 0xd3ebf2a2f6f8f9157c17b01687f0f13175fdcd58ac706593f24cf342fab1b024.
//
// Solidity: event SetDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef, OrganisationDocsDocument _document)
func (_IATIOrganisationDocs *IATIOrganisationDocsFilterer) FilterSetDocument(opts *bind.FilterOpts) (*IATIOrganisationDocsSetDocumentIterator, error) {

	logs, sub, err := _IATIOrganisationDocs.contract.FilterLogs(opts, "SetDocument")
	if err != nil {
		return nil, err
	}
	return &IATIOrganisationDocsSetDocumentIterator{contract: _IATIOrganisationDocs.contract, event: "SetDocument", logs: logs, sub: sub}, nil
}

// WatchSetDocument is a free log subscription operation binding the contract event 0xd3ebf2a2f6f8f9157c17b01687f0f13175fdcd58ac706593f24cf342fab1b024.
//
// Solidity: event SetDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef, OrganisationDocsDocument _document)
func (_IATIOrganisationDocs *IATIOrganisationDocsFilterer) WatchSetDocument(opts *bind.WatchOpts, sink chan<- *IATIOrganisationDocsSetDocument) (event.Subscription, error) {

	logs, sub, err := _IATIOrganisationDocs.contract.WatchLogs(opts, "SetDocument")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IATIOrganisationDocsSetDocument)
				if err := _IATIOrganisationDocs.contract.UnpackLog(event, "SetDocument", log); err != nil {
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

// ParseSetDocument is a log parse operation binding the contract event 0xd3ebf2a2f6f8f9157c17b01687f0f13175fdcd58ac706593f24cf342fab1b024.
//
// Solidity: event SetDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef, OrganisationDocsDocument _document)
func (_IATIOrganisationDocs *IATIOrganisationDocsFilterer) ParseSetDocument(log types.Log) (*IATIOrganisationDocsSetDocument, error) {
	event := new(IATIOrganisationDocsSetDocument)
	if err := _IATIOrganisationDocs.contract.UnpackLog(event, "SetDocument", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OrganisationDocsABI is the input ABI used to generate the binding from.
const OrganisationDocsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"format\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"category\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"countryRef\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structOrganisationDocs.Document\",\"name\":\"_document\",\"type\":\"tuple\"}],\"name\":\"SetDocument\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getDocReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"}],\"name\":\"getDocument\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"format\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"category\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"countryRef\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"}],\"internalType\":\"structOrganisationDocs.Document\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"}],\"name\":\"getDocumentCategory\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"}],\"name\":\"getDocumentCountry\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"}],\"name\":\"getDocumentDate\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"}],\"name\":\"getDocumentDescription\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"}],\"name\":\"getDocumentFormat\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"}],\"name\":\"getDocumentLang\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"}],\"name\":\"getDocumentTitle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"}],\"name\":\"getDocumentURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"}],\"name\":\"getNumDocs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_organisationsRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_organisationRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_docRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"format\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"category\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"countryRef\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"desc\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"date\",\"type\":\"bytes32\"}],\"internalType\":\"structOrganisationDocs.Document\",\"name\":\"_document\",\"type\":\"tuple\"}],\"name\":\"setDocument\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OrganisationDocsFuncSigs maps the 4-byte function signature to its string representation.
var OrganisationDocsFuncSigs = map[string]string{
	"132ac373": "getDocReference(bytes32,bytes32,uint256)",
	"966c2b10": "getDocument(bytes32,bytes32,bytes32)",
	"07e6d576": "getDocumentCategory(bytes32,bytes32,bytes32)",
	"02e06d15": "getDocumentCountry(bytes32,bytes32,bytes32)",
	"efedacb6": "getDocumentDate(bytes32,bytes32,bytes32)",
	"28accf53": "getDocumentDescription(bytes32,bytes32,bytes32)",
	"3330a30b": "getDocumentFormat(bytes32,bytes32,bytes32)",
	"ffbe65db": "getDocumentLang(bytes32,bytes32,bytes32)",
	"1690c748": "getDocumentTitle(bytes32,bytes32,bytes32)",
	"98dfafc3": "getDocumentURL(bytes32,bytes32,bytes32)",
	"b11c464e": "getNumDocs(bytes32,bytes32)",
	"0ac17bad": "setDocument(bytes32,bytes32,bytes32,(string,string,string,bytes32,bytes32,string,bytes32,bytes32))",
}

// OrganisationDocs is an auto generated Go binding around an Ethereum contract.
type OrganisationDocs struct {
	OrganisationDocsCaller     // Read-only binding to the contract
	OrganisationDocsTransactor // Write-only binding to the contract
	OrganisationDocsFilterer   // Log filterer for contract events
}

// OrganisationDocsCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrganisationDocsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganisationDocsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrganisationDocsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganisationDocsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrganisationDocsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrganisationDocsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrganisationDocsSession struct {
	Contract     *OrganisationDocs // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrganisationDocsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrganisationDocsCallerSession struct {
	Contract *OrganisationDocsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// OrganisationDocsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrganisationDocsTransactorSession struct {
	Contract     *OrganisationDocsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// OrganisationDocsRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrganisationDocsRaw struct {
	Contract *OrganisationDocs // Generic contract binding to access the raw methods on
}

// OrganisationDocsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrganisationDocsCallerRaw struct {
	Contract *OrganisationDocsCaller // Generic read-only contract binding to access the raw methods on
}

// OrganisationDocsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrganisationDocsTransactorRaw struct {
	Contract *OrganisationDocsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrganisationDocs creates a new instance of OrganisationDocs, bound to a specific deployed contract.
func NewOrganisationDocs(address common.Address, backend bind.ContractBackend) (*OrganisationDocs, error) {
	contract, err := bindOrganisationDocs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrganisationDocs{OrganisationDocsCaller: OrganisationDocsCaller{contract: contract}, OrganisationDocsTransactor: OrganisationDocsTransactor{contract: contract}, OrganisationDocsFilterer: OrganisationDocsFilterer{contract: contract}}, nil
}

// NewOrganisationDocsCaller creates a new read-only instance of OrganisationDocs, bound to a specific deployed contract.
func NewOrganisationDocsCaller(address common.Address, caller bind.ContractCaller) (*OrganisationDocsCaller, error) {
	contract, err := bindOrganisationDocs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrganisationDocsCaller{contract: contract}, nil
}

// NewOrganisationDocsTransactor creates a new write-only instance of OrganisationDocs, bound to a specific deployed contract.
func NewOrganisationDocsTransactor(address common.Address, transactor bind.ContractTransactor) (*OrganisationDocsTransactor, error) {
	contract, err := bindOrganisationDocs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrganisationDocsTransactor{contract: contract}, nil
}

// NewOrganisationDocsFilterer creates a new log filterer instance of OrganisationDocs, bound to a specific deployed contract.
func NewOrganisationDocsFilterer(address common.Address, filterer bind.ContractFilterer) (*OrganisationDocsFilterer, error) {
	contract, err := bindOrganisationDocs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrganisationDocsFilterer{contract: contract}, nil
}

// bindOrganisationDocs binds a generic wrapper to an already deployed contract.
func bindOrganisationDocs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrganisationDocsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrganisationDocs *OrganisationDocsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OrganisationDocs.Contract.OrganisationDocsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrganisationDocs *OrganisationDocsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrganisationDocs.Contract.OrganisationDocsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrganisationDocs *OrganisationDocsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrganisationDocs.Contract.OrganisationDocsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrganisationDocs *OrganisationDocsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OrganisationDocs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrganisationDocs *OrganisationDocsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrganisationDocs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrganisationDocs *OrganisationDocsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrganisationDocs.Contract.contract.Transact(opts, method, params...)
}

// GetDocReference is a free data retrieval call binding the contract method 0x132ac373.
//
// Solidity: function getDocReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) view returns(bytes32)
func (_OrganisationDocs *OrganisationDocsCaller) GetDocReference(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OrganisationDocs.contract.Call(opts, out, "getDocReference", _organisationsRef, _organisationRef, _index)
	return *ret0, err
}

// GetDocReference is a free data retrieval call binding the contract method 0x132ac373.
//
// Solidity: function getDocReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) view returns(bytes32)
func (_OrganisationDocs *OrganisationDocsSession) GetDocReference(_organisationsRef [32]byte, _organisationRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _OrganisationDocs.Contract.GetDocReference(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _index)
}

// GetDocReference is a free data retrieval call binding the contract method 0x132ac373.
//
// Solidity: function getDocReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) view returns(bytes32)
func (_OrganisationDocs *OrganisationDocsCallerSession) GetDocReference(_organisationsRef [32]byte, _organisationRef [32]byte, _index *big.Int) ([32]byte, error) {
	return _OrganisationDocs.Contract.GetDocReference(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _index)
}

// GetDocument is a free data retrieval call binding the contract method 0x966c2b10.
//
// Solidity: function getDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(OrganisationDocsDocument)
func (_OrganisationDocs *OrganisationDocsCaller) GetDocument(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (OrganisationDocsDocument, error) {
	var (
		ret0 = new(OrganisationDocsDocument)
	)
	out := ret0
	err := _OrganisationDocs.contract.Call(opts, out, "getDocument", _organisationsRef, _organisationRef, _docRef)
	return *ret0, err
}

// GetDocument is a free data retrieval call binding the contract method 0x966c2b10.
//
// Solidity: function getDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(OrganisationDocsDocument)
func (_OrganisationDocs *OrganisationDocsSession) GetDocument(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (OrganisationDocsDocument, error) {
	return _OrganisationDocs.Contract.GetDocument(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocument is a free data retrieval call binding the contract method 0x966c2b10.
//
// Solidity: function getDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(OrganisationDocsDocument)
func (_OrganisationDocs *OrganisationDocsCallerSession) GetDocument(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (OrganisationDocsDocument, error) {
	return _OrganisationDocs.Contract.GetDocument(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentCategory is a free data retrieval call binding the contract method 0x07e6d576.
//
// Solidity: function getDocumentCategory(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_OrganisationDocs *OrganisationDocsCaller) GetDocumentCategory(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OrganisationDocs.contract.Call(opts, out, "getDocumentCategory", _organisationsRef, _organisationRef, _docRef)
	return *ret0, err
}

// GetDocumentCategory is a free data retrieval call binding the contract method 0x07e6d576.
//
// Solidity: function getDocumentCategory(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_OrganisationDocs *OrganisationDocsSession) GetDocumentCategory(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	return _OrganisationDocs.Contract.GetDocumentCategory(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentCategory is a free data retrieval call binding the contract method 0x07e6d576.
//
// Solidity: function getDocumentCategory(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_OrganisationDocs *OrganisationDocsCallerSession) GetDocumentCategory(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	return _OrganisationDocs.Contract.GetDocumentCategory(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentCountry is a free data retrieval call binding the contract method 0x02e06d15.
//
// Solidity: function getDocumentCountry(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_OrganisationDocs *OrganisationDocsCaller) GetDocumentCountry(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OrganisationDocs.contract.Call(opts, out, "getDocumentCountry", _organisationsRef, _organisationRef, _docRef)
	return *ret0, err
}

// GetDocumentCountry is a free data retrieval call binding the contract method 0x02e06d15.
//
// Solidity: function getDocumentCountry(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_OrganisationDocs *OrganisationDocsSession) GetDocumentCountry(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	return _OrganisationDocs.Contract.GetDocumentCountry(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentCountry is a free data retrieval call binding the contract method 0x02e06d15.
//
// Solidity: function getDocumentCountry(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_OrganisationDocs *OrganisationDocsCallerSession) GetDocumentCountry(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	return _OrganisationDocs.Contract.GetDocumentCountry(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentDate is a free data retrieval call binding the contract method 0xefedacb6.
//
// Solidity: function getDocumentDate(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_OrganisationDocs *OrganisationDocsCaller) GetDocumentDate(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OrganisationDocs.contract.Call(opts, out, "getDocumentDate", _organisationsRef, _organisationRef, _docRef)
	return *ret0, err
}

// GetDocumentDate is a free data retrieval call binding the contract method 0xefedacb6.
//
// Solidity: function getDocumentDate(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_OrganisationDocs *OrganisationDocsSession) GetDocumentDate(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	return _OrganisationDocs.Contract.GetDocumentDate(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentDate is a free data retrieval call binding the contract method 0xefedacb6.
//
// Solidity: function getDocumentDate(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_OrganisationDocs *OrganisationDocsCallerSession) GetDocumentDate(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	return _OrganisationDocs.Contract.GetDocumentDate(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentDescription is a free data retrieval call binding the contract method 0x28accf53.
//
// Solidity: function getDocumentDescription(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_OrganisationDocs *OrganisationDocsCaller) GetDocumentDescription(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _OrganisationDocs.contract.Call(opts, out, "getDocumentDescription", _organisationsRef, _organisationRef, _docRef)
	return *ret0, err
}

// GetDocumentDescription is a free data retrieval call binding the contract method 0x28accf53.
//
// Solidity: function getDocumentDescription(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_OrganisationDocs *OrganisationDocsSession) GetDocumentDescription(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	return _OrganisationDocs.Contract.GetDocumentDescription(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentDescription is a free data retrieval call binding the contract method 0x28accf53.
//
// Solidity: function getDocumentDescription(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_OrganisationDocs *OrganisationDocsCallerSession) GetDocumentDescription(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	return _OrganisationDocs.Contract.GetDocumentDescription(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentFormat is a free data retrieval call binding the contract method 0x3330a30b.
//
// Solidity: function getDocumentFormat(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_OrganisationDocs *OrganisationDocsCaller) GetDocumentFormat(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _OrganisationDocs.contract.Call(opts, out, "getDocumentFormat", _organisationsRef, _organisationRef, _docRef)
	return *ret0, err
}

// GetDocumentFormat is a free data retrieval call binding the contract method 0x3330a30b.
//
// Solidity: function getDocumentFormat(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_OrganisationDocs *OrganisationDocsSession) GetDocumentFormat(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	return _OrganisationDocs.Contract.GetDocumentFormat(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentFormat is a free data retrieval call binding the contract method 0x3330a30b.
//
// Solidity: function getDocumentFormat(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_OrganisationDocs *OrganisationDocsCallerSession) GetDocumentFormat(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	return _OrganisationDocs.Contract.GetDocumentFormat(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentLang is a free data retrieval call binding the contract method 0xffbe65db.
//
// Solidity: function getDocumentLang(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_OrganisationDocs *OrganisationDocsCaller) GetDocumentLang(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _OrganisationDocs.contract.Call(opts, out, "getDocumentLang", _organisationsRef, _organisationRef, _docRef)
	return *ret0, err
}

// GetDocumentLang is a free data retrieval call binding the contract method 0xffbe65db.
//
// Solidity: function getDocumentLang(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_OrganisationDocs *OrganisationDocsSession) GetDocumentLang(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	return _OrganisationDocs.Contract.GetDocumentLang(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentLang is a free data retrieval call binding the contract method 0xffbe65db.
//
// Solidity: function getDocumentLang(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(bytes32)
func (_OrganisationDocs *OrganisationDocsCallerSession) GetDocumentLang(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) ([32]byte, error) {
	return _OrganisationDocs.Contract.GetDocumentLang(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentTitle is a free data retrieval call binding the contract method 0x1690c748.
//
// Solidity: function getDocumentTitle(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_OrganisationDocs *OrganisationDocsCaller) GetDocumentTitle(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _OrganisationDocs.contract.Call(opts, out, "getDocumentTitle", _organisationsRef, _organisationRef, _docRef)
	return *ret0, err
}

// GetDocumentTitle is a free data retrieval call binding the contract method 0x1690c748.
//
// Solidity: function getDocumentTitle(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_OrganisationDocs *OrganisationDocsSession) GetDocumentTitle(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	return _OrganisationDocs.Contract.GetDocumentTitle(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentTitle is a free data retrieval call binding the contract method 0x1690c748.
//
// Solidity: function getDocumentTitle(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_OrganisationDocs *OrganisationDocsCallerSession) GetDocumentTitle(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	return _OrganisationDocs.Contract.GetDocumentTitle(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentURL is a free data retrieval call binding the contract method 0x98dfafc3.
//
// Solidity: function getDocumentURL(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_OrganisationDocs *OrganisationDocsCaller) GetDocumentURL(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _OrganisationDocs.contract.Call(opts, out, "getDocumentURL", _organisationsRef, _organisationRef, _docRef)
	return *ret0, err
}

// GetDocumentURL is a free data retrieval call binding the contract method 0x98dfafc3.
//
// Solidity: function getDocumentURL(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_OrganisationDocs *OrganisationDocsSession) GetDocumentURL(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	return _OrganisationDocs.Contract.GetDocumentURL(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetDocumentURL is a free data retrieval call binding the contract method 0x98dfafc3.
//
// Solidity: function getDocumentURL(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef) view returns(string)
func (_OrganisationDocs *OrganisationDocsCallerSession) GetDocumentURL(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte) (string, error) {
	return _OrganisationDocs.Contract.GetDocumentURL(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef, _docRef)
}

// GetNumDocs is a free data retrieval call binding the contract method 0xb11c464e.
//
// Solidity: function getNumDocs(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(uint256)
func (_OrganisationDocs *OrganisationDocsCaller) GetNumDocs(opts *bind.CallOpts, _organisationsRef [32]byte, _organisationRef [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _OrganisationDocs.contract.Call(opts, out, "getNumDocs", _organisationsRef, _organisationRef)
	return *ret0, err
}

// GetNumDocs is a free data retrieval call binding the contract method 0xb11c464e.
//
// Solidity: function getNumDocs(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(uint256)
func (_OrganisationDocs *OrganisationDocsSession) GetNumDocs(_organisationsRef [32]byte, _organisationRef [32]byte) (*big.Int, error) {
	return _OrganisationDocs.Contract.GetNumDocs(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef)
}

// GetNumDocs is a free data retrieval call binding the contract method 0xb11c464e.
//
// Solidity: function getNumDocs(bytes32 _organisationsRef, bytes32 _organisationRef) view returns(uint256)
func (_OrganisationDocs *OrganisationDocsCallerSession) GetNumDocs(_organisationsRef [32]byte, _organisationRef [32]byte) (*big.Int, error) {
	return _OrganisationDocs.Contract.GetNumDocs(&_OrganisationDocs.CallOpts, _organisationsRef, _organisationRef)
}

// SetDocument is a paid mutator transaction binding the contract method 0x0ac17bad.
//
// Solidity: function setDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef, OrganisationDocsDocument _document) returns()
func (_OrganisationDocs *OrganisationDocsTransactor) SetDocument(opts *bind.TransactOpts, _organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte, _document OrganisationDocsDocument) (*types.Transaction, error) {
	return _OrganisationDocs.contract.Transact(opts, "setDocument", _organisationsRef, _organisationRef, _docRef, _document)
}

// SetDocument is a paid mutator transaction binding the contract method 0x0ac17bad.
//
// Solidity: function setDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef, OrganisationDocsDocument _document) returns()
func (_OrganisationDocs *OrganisationDocsSession) SetDocument(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte, _document OrganisationDocsDocument) (*types.Transaction, error) {
	return _OrganisationDocs.Contract.SetDocument(&_OrganisationDocs.TransactOpts, _organisationsRef, _organisationRef, _docRef, _document)
}

// SetDocument is a paid mutator transaction binding the contract method 0x0ac17bad.
//
// Solidity: function setDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef, OrganisationDocsDocument _document) returns()
func (_OrganisationDocs *OrganisationDocsTransactorSession) SetDocument(_organisationsRef [32]byte, _organisationRef [32]byte, _docRef [32]byte, _document OrganisationDocsDocument) (*types.Transaction, error) {
	return _OrganisationDocs.Contract.SetDocument(&_OrganisationDocs.TransactOpts, _organisationsRef, _organisationRef, _docRef, _document)
}

// OrganisationDocsSetDocumentIterator is returned from FilterSetDocument and is used to iterate over the raw logs and unpacked data for SetDocument events raised by the OrganisationDocs contract.
type OrganisationDocsSetDocumentIterator struct {
	Event *OrganisationDocsSetDocument // Event containing the contract specifics and raw log

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
func (it *OrganisationDocsSetDocumentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrganisationDocsSetDocument)
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
		it.Event = new(OrganisationDocsSetDocument)
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
func (it *OrganisationDocsSetDocumentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrganisationDocsSetDocumentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrganisationDocsSetDocument represents a SetDocument event raised by the OrganisationDocs contract.
type OrganisationDocsSetDocument struct {
	OrganisationsRef [32]byte
	OrganisationRef  [32]byte
	DocRef           [32]byte
	Document         OrganisationDocsDocument
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetDocument is a free log retrieval operation binding the contract event 0xd3ebf2a2f6f8f9157c17b01687f0f13175fdcd58ac706593f24cf342fab1b024.
//
// Solidity: event SetDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef, OrganisationDocsDocument _document)
func (_OrganisationDocs *OrganisationDocsFilterer) FilterSetDocument(opts *bind.FilterOpts) (*OrganisationDocsSetDocumentIterator, error) {

	logs, sub, err := _OrganisationDocs.contract.FilterLogs(opts, "SetDocument")
	if err != nil {
		return nil, err
	}
	return &OrganisationDocsSetDocumentIterator{contract: _OrganisationDocs.contract, event: "SetDocument", logs: logs, sub: sub}, nil
}

// WatchSetDocument is a free log subscription operation binding the contract event 0xd3ebf2a2f6f8f9157c17b01687f0f13175fdcd58ac706593f24cf342fab1b024.
//
// Solidity: event SetDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef, OrganisationDocsDocument _document)
func (_OrganisationDocs *OrganisationDocsFilterer) WatchSetDocument(opts *bind.WatchOpts, sink chan<- *OrganisationDocsSetDocument) (event.Subscription, error) {

	logs, sub, err := _OrganisationDocs.contract.WatchLogs(opts, "SetDocument")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrganisationDocsSetDocument)
				if err := _OrganisationDocs.contract.UnpackLog(event, "SetDocument", log); err != nil {
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

// ParseSetDocument is a log parse operation binding the contract event 0xd3ebf2a2f6f8f9157c17b01687f0f13175fdcd58ac706593f24cf342fab1b024.
//
// Solidity: event SetDocument(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _docRef, OrganisationDocsDocument _document)
func (_OrganisationDocs *OrganisationDocsFilterer) ParseSetDocument(log types.Log) (*OrganisationDocsSetDocument, error) {
	event := new(OrganisationDocsSetDocument)
	if err := _OrganisationDocs.contract.UnpackLog(event, "SetDocument", log); err != nil {
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
