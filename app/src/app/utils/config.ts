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
    "event SetOrganisation(bytes32 _orgRef, Organisation _org)",

    "function setOrganisation(Organisation memory _org)",

    "function getOrganisationExists(bytes32 _orgRef) view returns (bool)",
    "function getNumOrganisations() view returns (uint256)",
    "function getOrganisationReference(uint256 _index) view returns (bytes32)",
    "function getOrganisation(bytes32 _orgRef) view returns (Organisation memory)",
    "function getOrganisationName(bytes32 _orgRef) view returns (string memory)",
    "function getOrganisationIdentifier(bytes32 _orgRef) view returns (string memory)",
  ]

  static organisationReportsABI = [
    "event SetReport(bytes32 _reportRef, bytes32 _issuingOrgRef, Report _report)",
    "event SetDocument(bytes32 _reportRef, bytes32 _docRef, Document _document)",

    "function setReport(Report memory _report)",
    "function setDocument(Document memory _document)",

    "function getReportExists(bytes32 _reference) view returns (bool)",
    "function getIssuingOrgExists(bytes32 _reference, bytes32 _orgRef) view returns (bool)",
    "function getReportDocExists(bytes32 _reference, bytes32 _docRef) view returns (bool)",

    "function getNumReports() view returns (uint256)",
    "function getNumReportOrgs(bytes32 _reference) view returns (uint256)",
    "function getNumReportDocs(bytes32 _reference) view returns (uint256)",

    "function getReportReference(uint256 _index) view returns (bytes32)",
    "function getReportOrgReference(bytes32 _reference, uint256 _index) view returns (bytes32)",
    "function getReportDocReference(bytes32 _reference, uint256 _index) view returns (bytes32)",

    "function getReport(bytes32 _reference, bytes32 _orgRef) view returns (Report memory)",

    "function getReportingOrg(bytes32 _reference, bytes32 _orgRef) view returns (bytes32)",
    "function getLang(bytes32 _reference, bytes32 _orgRef) view returns (bytes32)",
    "function getCurrency(bytes32 _reference,  bytes32 _orgRef) view returns (bytes32)",
    "function getVersion(bytes32 _reference,  bytes32 _orgRef) view returns (bytes32)",
    "function getGeneratedTime(bytes32 _reference,  bytes32 _orgRef) view returns (bytes32)",
    "function getLastUpdatedTime(bytes32 _reference,  bytes32 _orgRef) view returns (bytes32)",

    "function getReportingOrgType(bytes32 _reference, bytes32 _reportingOrgRef) view returns (uint8)",
    "function getReportingOrgIsSecondary(bytes32 _reference, bytes32 _reportingOrgRef) view returns (bool)",
    
    "function getDocument(bytes32 _reference, bytes32 _docRef) view returns (Document memory)",
    "function getDocumentTitle(bytes32 _reference, bytes32 _docRef) view returns (string memory)",
    "function getDocumentFormat(bytes32 _reference, bytes32 _docRef) view returns (bytes32)",
    "function getDocumentURL(bytes32 _reference, bytes32 _docRef) view returns (string memory)",
    "function getDocumentCategory(bytes32 _reference, bytes32 _docRef) view returns (bytes32)",
    "function getDocumentCountry(bytes32 _reference, bytes32 _docRef) view returns (bytes32)",
    "function getDocumentDescription(bytes32 _reference, bytes32 _docRef) view returns (string memory)",
    "function getDocumentLang(bytes32 _reference, bytes32 _docRef) view returns (bytes32)",
    "function getDocumentDate(bytes32 _reference, bytes32 _docRef) view returns (bytes32)",
  ]


  static organisationsAddress = "0x667CD537e1079044119438B0c8a62d59D557B55A"
  static organisationReportsAddress = "0xcA5784A7c10A7CFC6480072FB46f0C54dFbFB989"
}

export { Paths, Blockchain, Contract }
