class Paths {

  static readonly home='/'
  static readonly blockchain = '/blockchain-data'
  static readonly about='/about'
  static readonly overview='/overview'
  static readonly help='/help'
  static readonly writer='/create'
  static readonly reader='/read'
  static readonly orgWriter='/create-organisation'
  static readonly orgReader='/read-organisation'
}

class Blockchain {

  // Using web3 via MetaMask, so these aren't necessary, really
  static readonly host = 'localhost'
  static readonly port = '8545'

  static readonly checkAccountInterval = 3000
}

class OrgsContract {

  static organisationsABI = [
    "event OrganisationSet(string _reference, string _name, string _type)",
    "function setOrganisation(string _reference, string _name, string _type)",
    "function getOrganisationExists(string _reference)",
    "function getNumOrganisations()",
    "function getOrganisationReference(uint256 _index)",
    "function getOrganisationName(string _reference)",
    "function getOrganisationType(string _reference)",
  ]

  /*
  organisations: 0x749f861de9e83807e0ebaadaedd88a2f645dc176
  */

  static organisationsAddress = "0x749f861de9e83807e0ebaadaedd88a2f645dc176"
}

export { Paths, Blockchain, OrgsContract }
