class Paths {

  static readonly home='/'
  static readonly blockchain = '/blockchain-data'
  static readonly about='/about'
  static readonly overview='/overview'
  static readonly help='/help'
  static readonly writer='/create'
  static readonly reader='/read'
  static readonly orgWriter='/create-organisation'
  static readonly orgReader='/read-organisations'
}

class Blockchain {

  // Using web3 via MetaMask, so these aren't necessary, really
  static readonly host = 'localhost'
  static readonly port = '8545'

  static readonly checkAccountInterval = 3000
}

class Contract {

  static organisationsABI = [
    "event SetOrganisation(string _reference, string _name, string _namespaceCode, string _baseIdentifier)",

    "function setOrganisation(string _reference, string _name, string _namespaceCode, string _baseIdentifier)",
    "function getOrganisationExists(string _reference) view returns (bool)",

    "function getNumOrganisations() view returns (uint256)",

    "function getOrganisationReference(uint256 _index) view returns (string)",
    "function getOrganisationName(string _reference) view returns (string)",
    "function getOrganisationNamespaceCode(string _reference) view returns (string)",
    "function getOrganisationBaseIdentifier(string _reference) view returns (string)"
  ]

  static organisationReportsABI = [
    "function setReport(string _reference, string _version) public",
    "function setOrganisation(string _reference, string _orgRef, string _defaultLang, string _defaultCurrency) public",
    "function setReportingOrganisation(string _reference, string _reportingOrgRef, string _type, bool _isSecondary) public",
    "function setDocument(string _reference, string _docRef, bytes32[] _attributes) public",

    "function getReportExists(string _reference) public view returns (bool)",
    "function getReportDocExists(string _reference, string _docRef) public view returns (bool)",

    "function getVersion(string _reference) public view returns (string)",

    "function getNumReports() public view returns (uint256)",
    "function getNumReportDocs(string _reference) public view returns (uint256)",

    "function getReportReference(uint256 _index) public view returns (string)",
    "function getReportDocReference(string _reference, uint256 _index) public view returns (string)",

    "function getOrganisation(string _reference) public view returns (string)",
    "function getOrganisationDefaultLang(string _reference) public view returns (string)",
    "function getOrganisationDefaultCurrency(string _reference) public view returns (string)",

    "function getReportingOrganisation(string _reference) public view returns (string)",
    "function getReportingOrganisationType(string _reference) public view returns (string)",
    "function getReportingOrganisationIsSecondary(string _reference) public view returns (bool)",

    "function getDocumentTitle(string _reference, string _docRef) public view returns (bytes32)",
    "function getDocumentFormat(string _reference, string _docRef) public view returns (bytes32)",
    "function getDocumentURL(string _reference, string _docRef) public view returns (bytes32)",
    "function getDocumentCategory(string _reference, string _docRef) public view returns (bytes32)",
    "function getDocumentCountry(string _reference, string _docRef) public view returns (bytes32)",
    "function getDocumentDescription(string _reference, string _docRef) public view returns (bytes32)",
    "function getDocumentLang(string _reference, string _docRef) public view returns (bytes32)",
    "function getDocumentDate(string _reference, string _docRef) public view returns (bytes32)"
  ]

  /*
  organisations: 0x749f861de9e83807e0ebaadaedd88a2f645dc176
  */

  static organisationsAddress = "0xbc3bc6b7e02fdef11017d6b2ca1d63ec841f0ae6"
  static organisationReportsAddress = "0xc5fd6b07e25f3c33868cef0e36a0c1992ac3c254"
}

export { Paths, Blockchain, Contract }
