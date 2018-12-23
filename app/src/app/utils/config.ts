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
    "event SetOrganisation(bytes32 _reference, bytes32 _name, bytes32 _identifier)",

    "function setOrganisation(bytes32 _reference, bytes32 _name, bytes32 _identifier)",
    "function getOrganisationExists(bytes32 _reference) view returns (bool)",

    "function getNumOrganisations() view returns (uint256)",

    "function getOrganisationReference(uint256 _index) view returns (bytes32)",
    "function getOrganisationName(bytes32 _reference) view returns (bytes32)",
    "function getOrganisationIdentifier(bytes32 _reference) view returns (bytes32)"
  ]

  static organisationReportsABI = [
    "event SetReport(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)",
    "event SetDefaults(bytes32 _reference, bytes32 _orgRef, bytes32 _defaultLang, bytes32 _defaultCurrency)",
    "event SetReportingOrgType(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, uint8 _type, bool _isSecondary)",
    "event SetAssociatedDocument(bytes32 _reference, bytes32 _docRef)",

    "function setReport(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime) public",
    "function setDefaults(bytes32 _reference, bytes32 _orgRef, bytes32 _defaultLang, bytes32 _defaultCurrency) public",
    "function setReportingOrgType(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, uint8 _type, bool _isSecondary) public",
    "function setAssociatedDocument(bytes32 _reference, bytes32 _docRef, bytes32[] _attributes) public",

    "function getReportExists(bytes32 _reference) public view returns (bool)",
    "function getReportOrgExists(bytes32 _reference, bytes32 _orgRef) public view returns (bool)",
    "function getReportDocExists(bytes32 _reference, bytes32 _docRef) public view returns (bool)",

    "function getNumReports() public view returns (uint256)",
    "function getNumReportOrgs(bytes32 _reference) public view returns (uint256)",
    "function getNumReportDocs(bytes32 _reference) public view returns (uint256)",

    "function getReportReference(uint256 _index) public view returns (bytes32)",
    "function getReportOrgReference(bytes32 _reference, uint256 _index) public view returns (bytes32)",
    "function getReportDocReference(bytes32 _reference, uint256 _index) public view returns (bytes32)",

    "function getReportingOrg(bytes32 _reference, bytes32 _orgRef) public view returns (bytes32)",
    "function getLang(bytes32 _reference, bytes32 _orgRef) public view returns (bytes32)",
    "function getCurrency(bytes32 _reference,  bytes32 _orgRef) public view returns (bytes32)",
    "function getVersion(bytes32 _reference,  bytes32 _orgRef) public view returns (bytes32)",
    "function getGeneratedTime(bytes32 _reference,  bytes32 _orgRef) public view returns (bytes32)",
    "function getLastUpdatedTime(bytes32 _reference,  bytes32 _orgRef) public view returns (bytes32)",

    "function getReportingOrgType(bytes32 _reference, bytes32 _orgRef) public view returns (uint8)",
    "function getReportingOrgIsSecondary(bytes32 _reference, bytes32 _orgRef) public view returns (bool)",

    "function getDocumentTitle(bytes32 _reference, bytes32 _docRef) public view returns (bytes32)",
    "function getDocumentFormat(bytes32 _reference, bytes32 _docRef) public view returns (bytes32)",
    "function getDocumentURL(bytes32 _reference, bytes32 _docRef) public view returns (bytes32)",
    "function getDocumentCategory(bytes32 _reference, bytes32 _docRef) public view returns (bytes32)",
    "function getDocumentCountry(bytes32 _reference, bytes32 _docRef) public view returns (bytes32)",
    "function getDocumentDescription(bytes32 _reference, bytes32 _docRef) public view returns (bytes32)",
    "function getDocumentLang(bytes32 _reference, bytes32 _docRef) public view returns (bytes32)",
    "function getDocumentDate(bytes32 _reference, bytes32 _docRef) public view returns (bytes32)"
  ]


  static organisationsAddress = "0x78dbfca6d18f0f4dec22bdccd4c3f90295ff1831"
  static organisationReportsAddress = "0x00ac4b12eb39de040860ddae3cbba3f5e7145c5f"
}

export { Paths, Blockchain, Contract }
