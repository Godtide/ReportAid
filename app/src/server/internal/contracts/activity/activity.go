// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package activity

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

// ActivityData is an auto generated low-level Go binding around an user-defined struct.
type ActivityData struct {
	Humanitarian      bool
	Hierarchy         uint8
	Status            uint8
	BudgetNotProvided uint8
	ReportingOrg      ReportingOrgData
	LastUpdated       [32]byte
	Lang              [32]byte
	Currency          [32]byte
	LinkedData        [32]byte
	Identifier        [32]byte
	Title             string
	Description       string
}

// ReportingOrgData is an auto generated low-level Go binding around an user-defined struct.
type ReportingOrgData struct {
	OrgType     uint8
	IsSecondary bool
	OrgRef      [32]byte
}

// ActivityFactoryABI is the input ABI used to generate the binding from.
const ActivityFactoryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_activitiesCon\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_child\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_childType\",\"type\":\"uint8\"}],\"name\":\"AddChild\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_elementType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"Set\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_child\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_childType\",\"type\":\"uint8\"}],\"name\":\"addChild\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ref\",\"type\":\"bytes32\"}],\"name\":\"getFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getReference\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_parentRef\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_thisRef\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"humanitarian\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"hierarchy\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"budgetNotProvided\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSecondary\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"}],\"internalType\":\"structReportingOrgData\",\"name\":\"reportingOrg\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"lastUpdated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityData\",\"name\":\"_activity\",\"type\":\"tuple\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setter\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ActivityFactoryFuncSigs maps the 4-byte function signature to its string representation.
var ActivityFactoryFuncSigs = map[string]string{
	"cbb2dc95": "addChild(bytes32,address,uint8)",
	"8eaa6ac0": "get(bytes32)",
	"17fc2761": "getFactory(bytes32)",
	"67e0badb": "getNum()",
	"bc599341": "getReference(uint256)",
	"1006fbc7": "set(bytes32,bytes32,(bool,uint8,uint8,uint8,(uint8,bool,bytes32),bytes32,bytes32,bytes32,bytes32,bytes32,string,string))",
	"3f3108f7": "setter()",
}

// ActivityFactoryBin is the compiled bytecode used for deploying new contracts.
var ActivityFactoryBin = "0x608060405234801561001057600080fd5b5060405161174138038061174183398101604081905261002f9161006c565b6001600160a01b03811661003f57fe5b600080546001600160a01b0319166001600160a01b0392909216919091178155600381905560065561009a565b60006020828403121561007d578081fd5b81516001600160a01b0381168114610093578182fd5b9392505050565b611698806100a96000396000f3fe60806040523480156200001157600080fd5b5060043610620000885760003560e01c806367e0badb116200006357806367e0badb14620000ee5780638eaa6ac01462000107578063bc599341146200011e578063cbb2dc9514620001355762000088565b80631006fbc7146200008d57806317fc276114620000a65780633f3108f714620000d5575b600080fd5b620000a46200009e366004620006ee565b6200014c565b005b620000bd620000b736600462000688565b620002c5565b604051620000cc9190620008ba565b60405180910390f35b620000df62000325565b604051620000cc9190620008f9565b620000f862000330565b604051620000cc9190620008ce565b620000bd6200011836600462000688565b62000336565b620000f86200012f36600462000688565b62000398565b620000a462000146366004620006a1565b620003d2565b8260001a60f81b6001600160f81b0319166200016757600080fd5b8160001a60f81b6001600160f81b0319166200018257600080fd5b60008160405162000193906200057c565b6200019f91906200090e565b604051809103906000f080158015620001bc573d6000803e3d6000fd5b509050620001d36001848363ffffffff620004e116565b506000546001600160a01b031663cbb2dc95858360036040518463ffffffff1660e01b81526004016200020993929190620008d7565b600060405180830381600087803b1580156200022457600080fd5b505af115801562000239573d6000803e3d6000fd5b5050505060006040516200024d906200058a565b604051809103906000f0801580156200026a573d6000803e3d6000fd5b509050620002816004858363ffffffff620004e116565b507f8353a4d574992e9eb676a31cd326ce2f9f6c4829c3b84ae33ee9febbb2962e7b600385604051620002b692919062000a07565b60405180910390a15050505050565b600081811a60f81b6001600160f81b031916620002e157600080fd5b6000828152600460205260409020600101546001600160a01b03166200030657600080fd5b506000908152600460205260409020600101546001600160a01b031690565b631006fbc760e01b90565b60035490565b600081811a60f81b6001600160f81b0319166200035257600080fd5b600082815260016020819052604090912001546001600160a01b03166200037857600080fd5b50600090815260016020819052604090912001546001600160a01b031690565b6003546000908210620003aa57600080fd5b6002805483908110620003b957fe5b9060005260206000209060020201600001549050919050565b8260001a60f81b6001600160f81b031916620003ed57600080fd5b6001600160a01b0382166200040157600080fd5b6000838152600460205260409020600101546001600160a01b03166200042657600080fd5b600083815260046020819052604091829020600101549151630a3b0a4f60e01b81526001600160a01b03909216918291630a3b0a4f916200046a91879101620008ba565b600060405180830381600087803b1580156200048557600080fd5b505af11580156200049a573d6000803e3d6000fd5b505050507f4f2635eaa771447d282168d4763629c4b855ff7c733898eefc20b8699278d0c0848484604051620004d393929190620008d7565b60405180910390a150505050565b60008281526020849052604081208054600190910180546001600160a01b0319166001600160a01b03851617905580156200052157600191505062000575565b50600180850180549182018082556000868152602088905260409020819055859190839081106200054e57fe5b60009182526020822060029182020192909255908601805460010190559150620005759050565b9392505050565b610a2f8062000a4383390190565b6101f1806200147283390190565b80358015158114620005a957600080fd5b92915050565b600082601f830112620005c0578081fd5b813567ffffffffffffffff811115620005d7578182fd5b620005ec601f8201601f191660200162000a1a565b91508082528360208285010111156200060457600080fd5b8060208401602084013760009082016020015292915050565b6000606082840312156200062f578081fd5b6200063b606062000a1a565b905062000649838362000676565b8152602082013580151581146200065f57600080fd5b806020830152506040820135604082015292915050565b803560ff81168114620005a957600080fd5b6000602082840312156200069a578081fd5b5035919050565b600080600060608486031215620006b6578182fd5b8335925060208401356001600160a01b0381168114620006d4578283fd5b9150620006e5856040860162000676565b90509250925092565b60008060006060848603121562000703578283fd5b8335925060208401359150604084013567ffffffffffffffff8082111562000729578283fd5b8186016101c081890312156200073d578384fd5b61018092506200074d8362000a1a565b62000759898362000598565b81526200076a896020840162000676565b60208201526200077e896040840162000676565b604082015262000792896060840162000676565b6060820152620007a689608084016200061d565b608082015260e082013560a08201526101008083013560c08301526101208084013560e084015261014080850135838501526101609250828501358285015286850135915085821115620007f8578788fd5b620008068c838701620005af565b90840152506101a083013594508385111562000820578586fd5b6200082e8a868501620005af565b8183015250809450505050509250925092565b15159052565b60008151808452815b818110156200086e5760208185018101518683018201520162000850565b81811115620008805782602083870101525b50601f01601f19169290920160200192915050565b805160ff168252602080820151151590830152604090810151910152565b60ff169052565b6001600160a01b0391909116815260200190565b90815260200190565b9283526001600160a01b0391909116602083015260ff16604082015260600190565b6001600160e01b031991909116815260200190565b6000602082526200092460208301845162000841565b6020830151620009386040840182620008b3565b5060408301516200094d6060840182620008b3565b506060830151620009626080840182620008b3565b5060808301516200097760a084018262000895565b5060a0830151610100818185015260c08501519150610120828186015260e086015192506101408381870152828701519350610160925083838701528187015161018087015280870151935050506101c0806101a0860152620009df6101e086018462000847565b86830151868203601f1901838801529350620009fc818562000847565b979650505050505050565b60ff929092168252602082015260400190565b60405181810167ffffffffffffffff8111828210171562000a3a57600080fd5b60405291905056fe60806040523480156200001157600080fd5b5060405162000a2f38038062000a2f83398101604081905262000034916200042f565b61012081015160001a60f81b6001600160f81b0319161580159062000070575060808101516040015160001a60f81b6001600160f81b03191615155b801562000084575060808101515160ff1615155b8015620000975750600081610140015151115b8015620000b7575060a081015160001a60f81b6001600160f81b03191615155b8015620000d7575060c081015160001a60f81b6001600160f81b03191615155b8015620000f7575060e081015160001a60f81b6001600160f81b03191615155b80156200010a5750602081015160ff1615155b8015620001215750600460ff16816020015160ff16105b8015620001345750604081015160ff1615155b80156200014b5750600760ff16816040015160ff16105b80156200015e5750606081015160ff1615155b8015620001755750600460ff16816060015160ff16105b8015620001885750600081610160015151115b6200019257600080fd5b6101208101516007556080810151602080820151600180546040850151600255935160ff1660ff199215156101000261ff00199095169490941791909116929092179091556101408201518051620001ef92600892019062000279565b5060a081015160035560c081015160045560e0810151600555602080820151600080546040850151606086015161ff001990921661010060ff958616021762ff0000191662010000918516919091021763ff00000019166301000000939091169290920291909117905561016082015180516200027192600992019062000279565b505062000592565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002bc57805160ff1916838001178555620002ec565b82800160010185558215620002ec579182015b82811115620002ec578251825591602001919060010190620002cf565b50620002fa929150620002fe565b5090565b6200031b91905b80821115620002fa576000815560010162000305565b90565b805180151581146200032f57600080fd5b92915050565b600082601f83011262000346578081fd5b81516001600160401b038111156200035c578182fd5b602062000372601f8301601f191682016200056b565b925081835284818386010111156200038957600080fd5b60005b82811015620003a95784810182015184820183015281016200038c565b82811115620003bb5760008284860101525b50505092915050565b600060608284031215620003d6578081fd5b620003e260606200056b565b9050620003f083836200041d565b8152602082015180151581146200040657600080fd5b806020830152506040820151604082015292915050565b805160ff811681146200032f57600080fd5b60006020828403121562000441578081fd5b81516001600160401b038082111562000458578283fd5b8184016101c081870312156200046c578384fd5b61018092506200047c836200056b565b6200048887836200031e565b81526200049987602084016200041d565b6020820152620004ad87604084016200041d565b6040820152620004c187606084016200041d565b6060820152620004d58760808401620003c4565b608082015260e082015160a08201526101008083015160c08301526101208084015160e08401526101408085015183850152610160925082850151828501528685015191508582111562000527578788fd5b620005358a83870162000335565b90840152506101a08301519450838511156200054f578586fd5b6200055d8886850162000335565b908201529695505050505050565b6040518181016001600160401b03811182821017156200058a57600080fd5b604052919050565b61048d80620005a26000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636d4ce63c1461003b578063993a04b71461005a575b600080fd5b61004361006f565b604051610051929190610364565b60405180910390f35b610062610244565b604051610051919061034f565b600061007961024f565b600360408051610180810182526000805460ff80821615158452610100808304821660208087019190915262010000840483168688015263010000009093048216606080870191909152865190810187526001805480851683528390049093161515818501526002805482890152608087019190915260035460a087015260045460c087015260055460e087015260065482870152600754610120870152600880548851948116159093026000190190921604601f81018490048402830184019096528582529294859361014086019390919083018282801561019d5780601f106101725761010080835404028352916020019161019d565b820191906000526020600020905b81548152906001019060200180831161018057829003601f168201915b505050918352505060098201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156102315780601f1061020657610100808354040283529160200191610231565b820191906000526020600020905b81548152906001019060200180831161021457829003601f168201915b5050505050815250509050915091509091565b631b53398f60e21b90565b60408051610180810182526000808252602082018190529181018290526060810191909152608081016102806102b9565b81526000602082018190526040820181905260608083018290526080830182905260a083019190915260c0820181905260e09091015290565b604080516060810182526000808252602082018190529181019190915290565b15159052565b60008151808452815b81811015610304576020818501810151868301820152016102e8565b818111156103155782602083870101525b50601f01601f19169290920160200192915050565b805160ff168252602080820151151590830152604090810151910152565b60ff169052565b6001600160e01b031991909116815260200190565b600060ff84168252604060208301526103816040830184516102d9565b60208301516103936060840182610348565b5060408301516103a66080840182610348565b5060608301516103b960a0840182610348565b5060808301516103cc60c084018261032a565b5060a0830151610120818185015260c08501519150610140828186015260e086015192506101608381870152610100870151610180870152828701516101a08701528187015193506101c09250828387015261042c6102008701856102df565b90870151868203603f19016101e08801529350915061044d905081836102df565b969550505050505056fea26469706673582212203693742a72cc4a3724bb08c439c78a997d5dffd3cc64f289026744bc7f92170864736f6c63430006060033608060405234801561001057600080fd5b506101d1806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630a3b0a4f146100465780634f0f4aa91461005b57806367e0badb14610084575b600080fd5b610059610054366004610138565b610099565b005b61006e610069366004610166565b6100fb565b60405161007b919061017e565b60405180910390f35b61008c610132565b60405161007b9190610192565b6001600160a01b0381166100ac57600080fd5b600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630180546001600160a01b0319166001600160a01b0392909216919091179055565b60008054821061010a57600080fd5b6000828154811061011757fe5b6000918252602090912001546001600160a01b031692915050565b60005490565b600060208284031215610149578081fd5b81356001600160a01b038116811461015f578182fd5b9392505050565b600060208284031215610177578081fd5b5035919050565b6001600160a01b0391909116815260200190565b9081526020019056fea26469706673582212200b1d0f60a813a64e8e7204f784c414c6a9bd1c546926bc9e4be49f667646b48864736f6c63430006060033a26469706673582212200ef0199127f03d93d6b26bfc3f14c1bdf7f9ccdc64a1c638318a0c38309387fd64736f6c63430006060033"

// DeployActivityFactory deploys a new Ethereum contract, binding an instance of ActivityFactory to it.
func DeployActivityFactory(auth *bind.TransactOpts, backend bind.ContractBackend, _activitiesCon common.Address) (common.Address, *types.Transaction, *ActivityFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivityFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ActivityFactoryBin), backend, _activitiesCon)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ActivityFactory{ActivityFactoryCaller: ActivityFactoryCaller{contract: contract}, ActivityFactoryTransactor: ActivityFactoryTransactor{contract: contract}, ActivityFactoryFilterer: ActivityFactoryFilterer{contract: contract}}, nil
}

// ActivityFactory is an auto generated Go binding around an Ethereum contract.
type ActivityFactory struct {
	ActivityFactoryCaller     // Read-only binding to the contract
	ActivityFactoryTransactor // Write-only binding to the contract
	ActivityFactoryFilterer   // Log filterer for contract events
}

// ActivityFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivityFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivityFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivityFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivityFactorySession struct {
	Contract     *ActivityFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ActivityFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivityFactoryCallerSession struct {
	Contract *ActivityFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ActivityFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivityFactoryTransactorSession struct {
	Contract     *ActivityFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ActivityFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivityFactoryRaw struct {
	Contract *ActivityFactory // Generic contract binding to access the raw methods on
}

// ActivityFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivityFactoryCallerRaw struct {
	Contract *ActivityFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ActivityFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivityFactoryTransactorRaw struct {
	Contract *ActivityFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivityFactory creates a new instance of ActivityFactory, bound to a specific deployed contract.
func NewActivityFactory(address common.Address, backend bind.ContractBackend) (*ActivityFactory, error) {
	contract, err := bindActivityFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ActivityFactory{ActivityFactoryCaller: ActivityFactoryCaller{contract: contract}, ActivityFactoryTransactor: ActivityFactoryTransactor{contract: contract}, ActivityFactoryFilterer: ActivityFactoryFilterer{contract: contract}}, nil
}

// NewActivityFactoryCaller creates a new read-only instance of ActivityFactory, bound to a specific deployed contract.
func NewActivityFactoryCaller(address common.Address, caller bind.ContractCaller) (*ActivityFactoryCaller, error) {
	contract, err := bindActivityFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityFactoryCaller{contract: contract}, nil
}

// NewActivityFactoryTransactor creates a new write-only instance of ActivityFactory, bound to a specific deployed contract.
func NewActivityFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivityFactoryTransactor, error) {
	contract, err := bindActivityFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityFactoryTransactor{contract: contract}, nil
}

// NewActivityFactoryFilterer creates a new log filterer instance of ActivityFactory, bound to a specific deployed contract.
func NewActivityFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivityFactoryFilterer, error) {
	contract, err := bindActivityFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivityFactoryFilterer{contract: contract}, nil
}

// bindActivityFactory binds a generic wrapper to an already deployed contract.
func bindActivityFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivityFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityFactory *ActivityFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivityFactory.Contract.ActivityFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityFactory *ActivityFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityFactory.Contract.ActivityFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityFactory *ActivityFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityFactory.Contract.ActivityFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityFactory *ActivityFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivityFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityFactory *ActivityFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityFactory *ActivityFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityFactory.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _ref) view returns(address)
func (_ActivityFactory *ActivityFactoryCaller) Get(opts *bind.CallOpts, _ref [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ActivityFactory.contract.Call(opts, out, "get", _ref)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _ref) view returns(address)
func (_ActivityFactory *ActivityFactorySession) Get(_ref [32]byte) (common.Address, error) {
	return _ActivityFactory.Contract.Get(&_ActivityFactory.CallOpts, _ref)
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _ref) view returns(address)
func (_ActivityFactory *ActivityFactoryCallerSession) Get(_ref [32]byte) (common.Address, error) {
	return _ActivityFactory.Contract.Get(&_ActivityFactory.CallOpts, _ref)
}

// GetFactory is a free data retrieval call binding the contract method 0x17fc2761.
//
// Solidity: function getFactory(bytes32 _ref) view returns(address)
func (_ActivityFactory *ActivityFactoryCaller) GetFactory(opts *bind.CallOpts, _ref [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ActivityFactory.contract.Call(opts, out, "getFactory", _ref)
	return *ret0, err
}

// GetFactory is a free data retrieval call binding the contract method 0x17fc2761.
//
// Solidity: function getFactory(bytes32 _ref) view returns(address)
func (_ActivityFactory *ActivityFactorySession) GetFactory(_ref [32]byte) (common.Address, error) {
	return _ActivityFactory.Contract.GetFactory(&_ActivityFactory.CallOpts, _ref)
}

// GetFactory is a free data retrieval call binding the contract method 0x17fc2761.
//
// Solidity: function getFactory(bytes32 _ref) view returns(address)
func (_ActivityFactory *ActivityFactoryCallerSession) GetFactory(_ref [32]byte) (common.Address, error) {
	return _ActivityFactory.Contract.GetFactory(&_ActivityFactory.CallOpts, _ref)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_ActivityFactory *ActivityFactoryCaller) GetNum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ActivityFactory.contract.Call(opts, out, "getNum")
	return *ret0, err
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_ActivityFactory *ActivityFactorySession) GetNum() (*big.Int, error) {
	return _ActivityFactory.Contract.GetNum(&_ActivityFactory.CallOpts)
}

// GetNum is a free data retrieval call binding the contract method 0x67e0badb.
//
// Solidity: function getNum() view returns(uint256)
func (_ActivityFactory *ActivityFactoryCallerSession) GetNum() (*big.Int, error) {
	return _ActivityFactory.Contract.GetNum(&_ActivityFactory.CallOpts)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_ActivityFactory *ActivityFactoryCaller) GetReference(opts *bind.CallOpts, _index *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ActivityFactory.contract.Call(opts, out, "getReference", _index)
	return *ret0, err
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_ActivityFactory *ActivityFactorySession) GetReference(_index *big.Int) ([32]byte, error) {
	return _ActivityFactory.Contract.GetReference(&_ActivityFactory.CallOpts, _index)
}

// GetReference is a free data retrieval call binding the contract method 0xbc599341.
//
// Solidity: function getReference(uint256 _index) view returns(bytes32)
func (_ActivityFactory *ActivityFactoryCallerSession) GetReference(_index *big.Int) ([32]byte, error) {
	return _ActivityFactory.Contract.GetReference(&_ActivityFactory.CallOpts, _index)
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(bytes4)
func (_ActivityFactory *ActivityFactoryCaller) Setter(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _ActivityFactory.contract.Call(opts, out, "setter")
	return *ret0, err
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(bytes4)
func (_ActivityFactory *ActivityFactorySession) Setter() ([4]byte, error) {
	return _ActivityFactory.Contract.Setter(&_ActivityFactory.CallOpts)
}

// Setter is a free data retrieval call binding the contract method 0x3f3108f7.
//
// Solidity: function setter() view returns(bytes4)
func (_ActivityFactory *ActivityFactoryCallerSession) Setter() ([4]byte, error) {
	return _ActivityFactory.Contract.Setter(&_ActivityFactory.CallOpts)
}

// AddChild is a paid mutator transaction binding the contract method 0xcbb2dc95.
//
// Solidity: function addChild(bytes32 _ref, address _child, uint8 _childType) returns()
func (_ActivityFactory *ActivityFactoryTransactor) AddChild(opts *bind.TransactOpts, _ref [32]byte, _child common.Address, _childType uint8) (*types.Transaction, error) {
	return _ActivityFactory.contract.Transact(opts, "addChild", _ref, _child, _childType)
}

// AddChild is a paid mutator transaction binding the contract method 0xcbb2dc95.
//
// Solidity: function addChild(bytes32 _ref, address _child, uint8 _childType) returns()
func (_ActivityFactory *ActivityFactorySession) AddChild(_ref [32]byte, _child common.Address, _childType uint8) (*types.Transaction, error) {
	return _ActivityFactory.Contract.AddChild(&_ActivityFactory.TransactOpts, _ref, _child, _childType)
}

// AddChild is a paid mutator transaction binding the contract method 0xcbb2dc95.
//
// Solidity: function addChild(bytes32 _ref, address _child, uint8 _childType) returns()
func (_ActivityFactory *ActivityFactoryTransactorSession) AddChild(_ref [32]byte, _child common.Address, _childType uint8) (*types.Transaction, error) {
	return _ActivityFactory.Contract.AddChild(&_ActivityFactory.TransactOpts, _ref, _child, _childType)
}

// Set is a paid mutator transaction binding the contract method 0x1006fbc7.
//
// Solidity: function set(bytes32 _parentRef, bytes32 _thisRef, ActivityData _activity) returns()
func (_ActivityFactory *ActivityFactoryTransactor) Set(opts *bind.TransactOpts, _parentRef [32]byte, _thisRef [32]byte, _activity ActivityData) (*types.Transaction, error) {
	return _ActivityFactory.contract.Transact(opts, "set", _parentRef, _thisRef, _activity)
}

// Set is a paid mutator transaction binding the contract method 0x1006fbc7.
//
// Solidity: function set(bytes32 _parentRef, bytes32 _thisRef, ActivityData _activity) returns()
func (_ActivityFactory *ActivityFactorySession) Set(_parentRef [32]byte, _thisRef [32]byte, _activity ActivityData) (*types.Transaction, error) {
	return _ActivityFactory.Contract.Set(&_ActivityFactory.TransactOpts, _parentRef, _thisRef, _activity)
}

// Set is a paid mutator transaction binding the contract method 0x1006fbc7.
//
// Solidity: function set(bytes32 _parentRef, bytes32 _thisRef, ActivityData _activity) returns()
func (_ActivityFactory *ActivityFactoryTransactorSession) Set(_parentRef [32]byte, _thisRef [32]byte, _activity ActivityData) (*types.Transaction, error) {
	return _ActivityFactory.Contract.Set(&_ActivityFactory.TransactOpts, _parentRef, _thisRef, _activity)
}

// ActivityFactoryAddChildIterator is returned from FilterAddChild and is used to iterate over the raw logs and unpacked data for AddChild events raised by the ActivityFactory contract.
type ActivityFactoryAddChildIterator struct {
	Event *ActivityFactoryAddChild // Event containing the contract specifics and raw log

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
func (it *ActivityFactoryAddChildIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivityFactoryAddChild)
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
		it.Event = new(ActivityFactoryAddChild)
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
func (it *ActivityFactoryAddChildIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivityFactoryAddChildIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivityFactoryAddChild represents a AddChild event raised by the ActivityFactory contract.
type ActivityFactoryAddChild struct {
	Ref       [32]byte
	Child     common.Address
	ChildType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddChild is a free log retrieval operation binding the contract event 0x4f2635eaa771447d282168d4763629c4b855ff7c733898eefc20b8699278d0c0.
//
// Solidity: event AddChild(bytes32 _ref, address _child, uint8 _childType)
func (_ActivityFactory *ActivityFactoryFilterer) FilterAddChild(opts *bind.FilterOpts) (*ActivityFactoryAddChildIterator, error) {

	logs, sub, err := _ActivityFactory.contract.FilterLogs(opts, "AddChild")
	if err != nil {
		return nil, err
	}
	return &ActivityFactoryAddChildIterator{contract: _ActivityFactory.contract, event: "AddChild", logs: logs, sub: sub}, nil
}

// WatchAddChild is a free log subscription operation binding the contract event 0x4f2635eaa771447d282168d4763629c4b855ff7c733898eefc20b8699278d0c0.
//
// Solidity: event AddChild(bytes32 _ref, address _child, uint8 _childType)
func (_ActivityFactory *ActivityFactoryFilterer) WatchAddChild(opts *bind.WatchOpts, sink chan<- *ActivityFactoryAddChild) (event.Subscription, error) {

	logs, sub, err := _ActivityFactory.contract.WatchLogs(opts, "AddChild")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivityFactoryAddChild)
				if err := _ActivityFactory.contract.UnpackLog(event, "AddChild", log); err != nil {
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
func (_ActivityFactory *ActivityFactoryFilterer) ParseAddChild(log types.Log) (*ActivityFactoryAddChild, error) {
	event := new(ActivityFactoryAddChild)
	if err := _ActivityFactory.contract.UnpackLog(event, "AddChild", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ActivityFactorySetIterator is returned from FilterSet and is used to iterate over the raw logs and unpacked data for Set events raised by the ActivityFactory contract.
type ActivityFactorySetIterator struct {
	Event *ActivityFactorySet // Event containing the contract specifics and raw log

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
func (it *ActivityFactorySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ActivityFactorySet)
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
		it.Event = new(ActivityFactorySet)
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
func (it *ActivityFactorySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ActivityFactorySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ActivityFactorySet represents a Set event raised by the ActivityFactory contract.
type ActivityFactorySet struct {
	ElementType uint8
	Ref         [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSet is a free log retrieval operation binding the contract event 0x8353a4d574992e9eb676a31cd326ce2f9f6c4829c3b84ae33ee9febbb2962e7b.
//
// Solidity: event Set(uint8 _elementType, bytes32 _ref)
func (_ActivityFactory *ActivityFactoryFilterer) FilterSet(opts *bind.FilterOpts) (*ActivityFactorySetIterator, error) {

	logs, sub, err := _ActivityFactory.contract.FilterLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return &ActivityFactorySetIterator{contract: _ActivityFactory.contract, event: "Set", logs: logs, sub: sub}, nil
}

// WatchSet is a free log subscription operation binding the contract event 0x8353a4d574992e9eb676a31cd326ce2f9f6c4829c3b84ae33ee9febbb2962e7b.
//
// Solidity: event Set(uint8 _elementType, bytes32 _ref)
func (_ActivityFactory *ActivityFactoryFilterer) WatchSet(opts *bind.WatchOpts, sink chan<- *ActivityFactorySet) (event.Subscription, error) {

	logs, sub, err := _ActivityFactory.contract.WatchLogs(opts, "Set")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ActivityFactorySet)
				if err := _ActivityFactory.contract.UnpackLog(event, "Set", log); err != nil {
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
func (_ActivityFactory *ActivityFactoryFilterer) ParseSet(log types.Log) (*ActivityFactorySet, error) {
	event := new(ActivityFactorySet)
	if err := _ActivityFactory.contract.UnpackLog(event, "Set", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ActivityNodeABI is the input ABI used to generate the binding from.
const ActivityNodeABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"humanitarian\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"hierarchy\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"budgetNotProvided\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSecondary\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"}],\"internalType\":\"structReportingOrgData\",\"name\":\"reportingOrg\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"lastUpdated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityData\",\"name\":\"_activity\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"humanitarian\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"hierarchy\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"budgetNotProvided\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orgType\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isSecondary\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"orgRef\",\"type\":\"bytes32\"}],\"internalType\":\"structReportingOrgData\",\"name\":\"reportingOrg\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"lastUpdated\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lang\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"currency\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkedData\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"internalType\":\"structActivityData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getter\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ActivityNodeFuncSigs maps the 4-byte function signature to its string representation.
var ActivityNodeFuncSigs = map[string]string{
	"6d4ce63c": "get()",
	"993a04b7": "getter()",
}

// ActivityNodeBin is the compiled bytecode used for deploying new contracts.
var ActivityNodeBin = "0x60806040523480156200001157600080fd5b5060405162000a2f38038062000a2f83398101604081905262000034916200042f565b61012081015160001a60f81b6001600160f81b0319161580159062000070575060808101516040015160001a60f81b6001600160f81b03191615155b801562000084575060808101515160ff1615155b8015620000975750600081610140015151115b8015620000b7575060a081015160001a60f81b6001600160f81b03191615155b8015620000d7575060c081015160001a60f81b6001600160f81b03191615155b8015620000f7575060e081015160001a60f81b6001600160f81b03191615155b80156200010a5750602081015160ff1615155b8015620001215750600460ff16816020015160ff16105b8015620001345750604081015160ff1615155b80156200014b5750600760ff16816040015160ff16105b80156200015e5750606081015160ff1615155b8015620001755750600460ff16816060015160ff16105b8015620001885750600081610160015151115b6200019257600080fd5b6101208101516007556080810151602080820151600180546040850151600255935160ff1660ff199215156101000261ff00199095169490941791909116929092179091556101408201518051620001ef92600892019062000279565b5060a081015160035560c081015160045560e0810151600555602080820151600080546040850151606086015161ff001990921661010060ff958616021762ff0000191662010000918516919091021763ff00000019166301000000939091169290920291909117905561016082015180516200027192600992019062000279565b505062000592565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002bc57805160ff1916838001178555620002ec565b82800160010185558215620002ec579182015b82811115620002ec578251825591602001919060010190620002cf565b50620002fa929150620002fe565b5090565b6200031b91905b80821115620002fa576000815560010162000305565b90565b805180151581146200032f57600080fd5b92915050565b600082601f83011262000346578081fd5b81516001600160401b038111156200035c578182fd5b602062000372601f8301601f191682016200056b565b925081835284818386010111156200038957600080fd5b60005b82811015620003a95784810182015184820183015281016200038c565b82811115620003bb5760008284860101525b50505092915050565b600060608284031215620003d6578081fd5b620003e260606200056b565b9050620003f083836200041d565b8152602082015180151581146200040657600080fd5b806020830152506040820151604082015292915050565b805160ff811681146200032f57600080fd5b60006020828403121562000441578081fd5b81516001600160401b038082111562000458578283fd5b8184016101c081870312156200046c578384fd5b61018092506200047c836200056b565b6200048887836200031e565b81526200049987602084016200041d565b6020820152620004ad87604084016200041d565b6040820152620004c187606084016200041d565b6060820152620004d58760808401620003c4565b608082015260e082015160a08201526101008083015160c08301526101208084015160e08401526101408085015183850152610160925082850151828501528685015191508582111562000527578788fd5b620005358a83870162000335565b90840152506101a08301519450838511156200054f578586fd5b6200055d8886850162000335565b908201529695505050505050565b6040518181016001600160401b03811182821017156200058a57600080fd5b604052919050565b61048d80620005a26000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636d4ce63c1461003b578063993a04b71461005a575b600080fd5b61004361006f565b604051610051929190610364565b60405180910390f35b610062610244565b604051610051919061034f565b600061007961024f565b600360408051610180810182526000805460ff80821615158452610100808304821660208087019190915262010000840483168688015263010000009093048216606080870191909152865190810187526001805480851683528390049093161515818501526002805482890152608087019190915260035460a087015260045460c087015260055460e087015260065482870152600754610120870152600880548851948116159093026000190190921604601f81018490048402830184019096528582529294859361014086019390919083018282801561019d5780601f106101725761010080835404028352916020019161019d565b820191906000526020600020905b81548152906001019060200180831161018057829003601f168201915b505050918352505060098201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156102315780601f1061020657610100808354040283529160200191610231565b820191906000526020600020905b81548152906001019060200180831161021457829003601f168201915b5050505050815250509050915091509091565b631b53398f60e21b90565b60408051610180810182526000808252602082018190529181018290526060810191909152608081016102806102b9565b81526000602082018190526040820181905260608083018290526080830182905260a083019190915260c0820181905260e09091015290565b604080516060810182526000808252602082018190529181019190915290565b15159052565b60008151808452815b81811015610304576020818501810151868301820152016102e8565b818111156103155782602083870101525b50601f01601f19169290920160200192915050565b805160ff168252602080820151151590830152604090810151910152565b60ff169052565b6001600160e01b031991909116815260200190565b600060ff84168252604060208301526103816040830184516102d9565b60208301516103936060840182610348565b5060408301516103a66080840182610348565b5060608301516103b960a0840182610348565b5060808301516103cc60c084018261032a565b5060a0830151610120818185015260c08501519150610140828186015260e086015192506101608381870152610100870151610180870152828701516101a08701528187015193506101c09250828387015261042c6102008701856102df565b90870151868203603f19016101e08801529350915061044d905081836102df565b969550505050505056fea26469706673582212203693742a72cc4a3724bb08c439c78a997d5dffd3cc64f289026744bc7f92170864736f6c63430006060033"

// DeployActivityNode deploys a new Ethereum contract, binding an instance of ActivityNode to it.
func DeployActivityNode(auth *bind.TransactOpts, backend bind.ContractBackend, _activity ActivityData) (common.Address, *types.Transaction, *ActivityNode, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivityNodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ActivityNodeBin), backend, _activity)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ActivityNode{ActivityNodeCaller: ActivityNodeCaller{contract: contract}, ActivityNodeTransactor: ActivityNodeTransactor{contract: contract}, ActivityNodeFilterer: ActivityNodeFilterer{contract: contract}}, nil
}

// ActivityNode is an auto generated Go binding around an Ethereum contract.
type ActivityNode struct {
	ActivityNodeCaller     // Read-only binding to the contract
	ActivityNodeTransactor // Write-only binding to the contract
	ActivityNodeFilterer   // Log filterer for contract events
}

// ActivityNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ActivityNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ActivityNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ActivityNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ActivityNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ActivityNodeSession struct {
	Contract     *ActivityNode     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ActivityNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ActivityNodeCallerSession struct {
	Contract *ActivityNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ActivityNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ActivityNodeTransactorSession struct {
	Contract     *ActivityNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ActivityNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ActivityNodeRaw struct {
	Contract *ActivityNode // Generic contract binding to access the raw methods on
}

// ActivityNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ActivityNodeCallerRaw struct {
	Contract *ActivityNodeCaller // Generic read-only contract binding to access the raw methods on
}

// ActivityNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ActivityNodeTransactorRaw struct {
	Contract *ActivityNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewActivityNode creates a new instance of ActivityNode, bound to a specific deployed contract.
func NewActivityNode(address common.Address, backend bind.ContractBackend) (*ActivityNode, error) {
	contract, err := bindActivityNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ActivityNode{ActivityNodeCaller: ActivityNodeCaller{contract: contract}, ActivityNodeTransactor: ActivityNodeTransactor{contract: contract}, ActivityNodeFilterer: ActivityNodeFilterer{contract: contract}}, nil
}

// NewActivityNodeCaller creates a new read-only instance of ActivityNode, bound to a specific deployed contract.
func NewActivityNodeCaller(address common.Address, caller bind.ContractCaller) (*ActivityNodeCaller, error) {
	contract, err := bindActivityNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityNodeCaller{contract: contract}, nil
}

// NewActivityNodeTransactor creates a new write-only instance of ActivityNode, bound to a specific deployed contract.
func NewActivityNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*ActivityNodeTransactor, error) {
	contract, err := bindActivityNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ActivityNodeTransactor{contract: contract}, nil
}

// NewActivityNodeFilterer creates a new log filterer instance of ActivityNode, bound to a specific deployed contract.
func NewActivityNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*ActivityNodeFilterer, error) {
	contract, err := bindActivityNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ActivityNodeFilterer{contract: contract}, nil
}

// bindActivityNode binds a generic wrapper to an already deployed contract.
func bindActivityNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ActivityNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityNode *ActivityNodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivityNode.Contract.ActivityNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityNode *ActivityNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityNode.Contract.ActivityNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityNode *ActivityNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityNode.Contract.ActivityNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ActivityNode *ActivityNodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ActivityNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ActivityNode *ActivityNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ActivityNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ActivityNode *ActivityNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ActivityNode.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint8, ActivityData)
func (_ActivityNode *ActivityNodeCaller) Get(opts *bind.CallOpts) (uint8, ActivityData, error) {
	var (
		ret0 = new(uint8)
		ret1 = new(ActivityData)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ActivityNode.contract.Call(opts, out, "get")
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint8, ActivityData)
func (_ActivityNode *ActivityNodeSession) Get() (uint8, ActivityData, error) {
	return _ActivityNode.Contract.Get(&_ActivityNode.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint8, ActivityData)
func (_ActivityNode *ActivityNodeCallerSession) Get() (uint8, ActivityData, error) {
	return _ActivityNode.Contract.Get(&_ActivityNode.CallOpts)
}

// Getter is a free data retrieval call binding the contract method 0x993a04b7.
//
// Solidity: function getter() view returns(bytes4)
func (_ActivityNode *ActivityNodeCaller) Getter(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _ActivityNode.contract.Call(opts, out, "getter")
	return *ret0, err
}

// Getter is a free data retrieval call binding the contract method 0x993a04b7.
//
// Solidity: function getter() view returns(bytes4)
func (_ActivityNode *ActivityNodeSession) Getter() ([4]byte, error) {
	return _ActivityNode.Contract.Getter(&_ActivityNode.CallOpts)
}

// Getter is a free data retrieval call binding the contract method 0x993a04b7.
//
// Solidity: function getter() view returns(bytes4)
func (_ActivityNode *ActivityNodeCallerSession) Getter() ([4]byte, error) {
	return _ActivityNode.Contract.Getter(&_ActivityNode.CallOpts)
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
