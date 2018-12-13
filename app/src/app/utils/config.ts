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
    "event SetOrganisation(string _reference, string _name, string _identifier)",

    "function setOrganisation(string _reference, string _name, string _identifier)",
    "function getOrganisationExists(string _reference) view returns (bool)",

    "function getNumOrganisations() view returns (uint256)",

    "function getOrganisationReference(uint256 _index) view returns (string)",
    "function getOrganisationName(string _reference) view returns (string)",
    "function getOrganisationIdentifier(string _reference) view returns (string)"
  ]

  static organisationReportsABI = [
    "event SetReport(string _reference, string _orgRef, string _reportingOrgRef, string _version, string _generatedTime)",
    "event SetDefaults(string _reference, string _orgRef, string _defaultLang, string _defaultCurrency)",
    "event SetReportingOrgType(string _reference, string _orgRef, string _reportingOrgRef, uint8 _type, bool _isSecondary)",
    "event SetAssociatedDocument(string _reference, string _docRef)",

    "function setReport(string _reference, string _orgRef, string _reportingOrgRef, string _version, string _generatedTime) public",
    "function setDefaults(string _reference, string _orgRef, string _defaultLang, string _defaultCurrency) public",
    "function setReportingOrgType(string _reference, string _orgRef, string _reportingOrgRef, uint8 _type, bool _isSecondary) public",
    "function setAssociatedDocument(string _reference, string _docRef, bytes32[] _attributes) public",

    "function getReportExists(string _reference) public view returns (bool)",
    "function getReportOrgExists(string _reference, string _orgRef) public view returns (bool)",
    "function getReportDocExists(string _reference, string _docRef) public view returns (bool)",

    "function getNumReports() public view returns (uint256)",
    "function getNumReportOrgs(string _reference) public view returns (uint256)",
    "function getNumReportDocs(string _reference) public view returns (uint256)",

    "function getReportReference(uint256 _index) public view returns (string)",
    "function getReportOrgReference(string _reference, uint256 _index) public view returns (string)",
    "function getReportDocReference(string _reference, uint256 _index) public view returns (string)",

    "function getReportingOrg(string _reference, string _orgRef) public view returns (string)",
    "function getLang(string _reference, string _orgRef) public view returns (string)",
    "function getCurrency(string _reference,  string _orgRef) public view returns (string)",
    "function getVersion(string _reference,  string _orgRef) public view returns (string)",
    "function getGeneratedTime(string _reference,  string _orgRef) public view returns (string)",
    "function getLastUpdatedTime(string _reference,  string _orgRef) public view returns (string)",

    "function getReportingOrgType(string _reference, string _orgRef) public view returns (uint8)",
    "function getReportingOrgIsSecondary(string _reference, string _orgRef) public view returns (bool)",

    "function getDocumentTitle(string _reference, string _docRef) public view returns (bytes32)",
    "function getDocumentFormat(string _reference, string _docRef) public view returns (bytes32)",
    "function getDocumentURL(string _reference, string _docRef) public view returns (bytes32)",
    "function getDocumentCategory(string _reference, string _docRef) public view returns (bytes32)",
    "function getDocumentCountry(string _reference, string _docRef) public view returns (bytes32)",
    "function getDocumentDescription(string _reference, string _docRef) public view returns (bytes32)",
    "function getDocumentLang(string _reference, string _docRef) public view returns (bytes32)",
    "function getDocumentDate(string _reference, string _docRef) public view returns (bytes32)"
  ]


  static organisationsAddress = "0x4a43ac2b58d0986bc528c7a13d52c07a53d9d9b0"
  static organisationReportsAddress = "0x2bf82369e89d1ce24796a6f242b2dc1b65d1bf18"
}

export { Paths, Blockchain, Contract }
