class Paths {

  static readonly home='/'
  static readonly blockchain = '/blockchain-data'
  static readonly about='/about'
  static readonly overview='/overview'
  static readonly help='/help'
  static readonly writer='/create'
  static readonly reader='/read'
  static readonly orgWriter='/create-organisation'
  static readonly orgReportWriter='/create-organisation-report'
  static readonly orgReader='/read-organisations'
  static readonly orgReportReader='/read-organisation-report'
}

class Blockchain {

  // Using web3 via MetaMask, so these aren't necessary, really
  static readonly host = 'localhost'
  static readonly port = '8545'

  static readonly checkAccountInterval = 3000
}

class Contract {

  static organisationsABI = [
    "event SetOrganisation(string memory _reference, string memory _name, string memory _identifier)",

    "function setOrganisation(string memory _reference, string memory _name, string memory _identifier)",
    "function getOrganisationExists(string memory _reference) view returns (bool)",

    "function getNumOrganisations() view returns (uint256)",

    "function getOrganisationReference(uint256 _index) view returns (string memory)",
    "function getOrganisationName(string memory _reference) view returns (string memory)",
    "function getOrganisationIdentifier(string memory _reference) view returns (string memory)"
  ]

  static organisationReportsABI = [
    "event SetReport(string memory _reference, string memory _orgRef, string memory _reportingOrgRef, string memory _version, string memory _generatedTime)",
    "event SetDefaults(string memory _reference, string memory _orgRef, string memory _defaultLang, string memory _defaultCurrency)",
    "event SetReportingOrgType(string memory _reference, string memory _orgRef, string memory _reportingOrgRef, uint8 _type, bool _isSecondary)",
    "event SetAssociatedDocument(string memory _reference, string memory _docRef)",

    "function setReport(string memory _reference, string memory _orgRef, string memory _reportingOrgRef, string memory _version, string memory _generatedTime) public",
    "function setDefaults(string memory _reference, string memory _orgRef, string memory _defaultLang, string memory _defaultCurrency) public",
    "function setReportingOrgType(string memory _reference, string memory _orgRef, string memory _reportingOrgRef, uint8 _type, bool _isSecondary) public",
    "function setAssociatedDocument(string memory _reference, string memory _docRef, bytes32[] _attributes) public",

    "function getReportExists(string memory _reference) public view returns (bool)",
    "function getReportOrgExists(string memory _reference, string memory _orgRef) public view returns (bool)",
    "function getReportDocExists(string memory _reference, string memory _docRef) public view returns (bool)",

    "function getNumReports() public view returns (uint256)",
    "function getNumReportOrgs(string memory _reference) public view returns (uint256)",
    "function getNumReportDocs(string memory _reference) public view returns (uint256)",

    "function getReportReference(uint256 _index) public view returns (string memory)",
    "function getReportOrgReference(string memory _reference, uint256 _index) public view returns (string memory)",
    "function getReportDocReference(string memory _reference, uint256 _index) public view returns (string memory)",

    "function getReportingOrg(string memory _reference, string memory _orgRef) public view returns (string memory)",
    "function getLang(string memory _reference, string memory _orgRef) public view returns (string memory)",
    "function getCurrency(string memory _reference,  string memory _orgRef) public view returns (string memory)",
    "function getVersion(string memory _reference,  string memory _orgRef) public view returns (string memory)",
    "function getGeneratedTime(string memory _reference,  string memory _orgRef) public view returns (string memory)",
    "function getLastUpdatedTime(string memory _reference,  string memory _orgRef) public view returns (string memory)",

    "function getReportingOrgType(string memory _reference, string memory _orgRef) public view returns (uint8)",
    "function getReportingOrgIsSecondary(string memory _reference, string memory _orgRef) public view returns (bool)",

    "function getDocumentTitle(string memory _reference, string memory _docRef) public view returns (bytes32)",
    "function getDocumentFormat(string memory _reference, string memory _docRef) public view returns (bytes32)",
    "function getDocumentURL(string memory _reference, string memory _docRef) public view returns (bytes32)",
    "function getDocumentCategory(string memory _reference, string memory _docRef) public view returns (bytes32)",
    "function getDocumentCountry(string memory _reference, string memory _docRef) public view returns (bytes32)",
    "function getDocumentDescription(string memory _reference, string memory _docRef) public view returns (bytes32)",
    "function getDocumentLang(string memory _reference, string memory _docRef) public view returns (bytes32)",
    "function getDocumentDate(string memory _reference, string memory _docRef) public view returns (bytes32)"
  ]


  static organisationsAddress = "0x78dbfca6d18f0f4dec22bdccd4c3f90295ff1831"
  static organisationReportsAddress = "0x00ac4b12eb39de040860ddae3cbba3f5e7145c5f"
}

export { Paths, Blockchain, Contract }
