import Web3Handler from './web3Handler'

interface ContractHandlerProps {
  web3: Web3Handler
}

class ContractHandler {

  static organisationsABI = [
    {
      "inputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "_reference",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_name",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "_type",
          "type": "string"
        }
      ],
      "name": "OrganisationSet",
      "type": "event"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_reference",
          "type": "string"
        },
        {
          "name": "_name",
          "type": "string"
        },
        {
          "name": "_type",
          "type": "string"
        }
      ],
      "name": "setOrganisation",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_reference",
          "type": "string"
        }
      ],
      "name": "getOrganisationExists",
      "outputs": [
        {
          "name": "",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "getNumOrganisations",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_index",
          "type": "uint256"
        }
      ],
      "name": "getOrganisationReference",
      "outputs": [
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_reference",
          "type": "string"
        }
      ],
      "name": "getOrganisationName",
      "outputs": [
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_reference",
          "type": "string"
        }
      ],
      "name": "getOrganisationType",
      "outputs": [
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    }
  ]

  /*
  organisations: 0x749f861de9e83807e0ebaadaedd88a2f645dc176
  */

  static organisationsAddress = '0x749f861de9e83807e0ebaadaedd88a2f645dc176'
  organisations: {};

  constructor (_web3Handler: ContractHandlerProps) {
    const web3 = _web3Handler.web3.getWeb3()
    const organisationsContract = web3.eth.contract(ContractHandler.organisationsABI)
    this.organisations = organisationsContract.at(ContractHandler.organisationsAddress)
  }

  getOrganisations () {
    return this.organisations
  }
}

export default ContractHandler
