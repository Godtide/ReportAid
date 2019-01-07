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
  static readonly host = 'http://127.0.0.1'
  static readonly port = '8545'
  static readonly network = 'ropsten'

  static readonly checkAccountInterval = 3000
}

class Contract {

  static organisationsABI = [
    "event SetOrganisation(bytes32 _orgRef, tuple(bytes32 orgRef, string name, string identifier) _org)",

    "function setOrganisation(tuple(bytes32 orgRef, string name, string identifier) _org)",

    "function getOrganisationExists(bytes32 _orgRef) view returns (bool)",
    "function getNumOrganisations() view returns (uint256)",
    "function getOrganisationReference(uint256 _index) view returns (bytes32)",
    "function getOrganisation(bytes32 _orgRef) view returns (tuple(bytes32 orgRef, string name, string identifier) org)",
    "function getOrganisationName(bytes32 _orgRef) view returns (string)",
    "function getOrganisationIdentifier(bytes32 _orgRef) view returns (string)",
  ]

  static organisationReportsABI = [
    "event SetReport(bytes32 _reportRef, bytes32 _issuingOrgRef, tuple(bytes32 reportRef, tuple(bytes32 orgRef, uint8 orgType, bool isSecondary), bytes32 issuingOrgRef, bytes32 version, bytes32 lang, bytes32 currency, bytes32 generatedTime, bytes32 lastUpdatedTime) _report)",
    "event SetDocument(bytes32 _reportRef, bytes32 _docRef, tuple(bytes32 reportRef, bytes32 docRef, string title, bytes32 format, string url, bytes32 category, bytes32 countryCode, string desc, bytes32 lang, bytes32 date) _doc)",

    "function setReport(tuple(bytes32 reportRef, tuple(bytes32 orgRef, uint8 orgType, bool isSecondary), bytes32 issuingOrgRef, bytes32 version, bytes32 lang, bytes32 currency, bytes32 generatedTime, bytes32 lastUpdatedTime) _report)",

    "function setDocument(tuple(bytes32 reportRef, bytes32 docRef, string title, bytes32 format, string url, bytes32 category, bytes32 countryCode, string desc, bytes32 lang, bytes32 date) _doc)",

    "function getReportExists(bytes32 _reference) view returns (bool)",
    "function getIssuingOrgExists(bytes32 _reference, bytes32 _orgRef) view returns (bool)",
    "function getReportDocExists(bytes32 _reference, bytes32 _docRef) view returns (bool)",

    "function getNumReports() view returns (uint256)",
    "function getNumReportOrgs(bytes32 _reference) view returns (uint256)",
    "function getNumReportDocs(bytes32 _reference) view returns (uint256)",

    "function getReportReference(uint256 _index) view returns (bytes32)",
    "function getReportOrgReference(bytes32 _reference, uint256 _index) view returns (bytes32)",
    "function getReportDocReference(bytes32 _reference, uint256 _index) view returns (bytes32)",

    "function getReport(bytes32 _reference, bytes32 _orgRef) view returns (tuple(bytes32 reportRef, tuple (bytes32 orgRef, uint8 orgType, bool isSecondary), bytes32 issuingOrgRef, bytes32 version, bytes32 lang, bytes32 currency, bytes32 generatedTime, bytes32 lastUpdatedTime) _report)",

    "function getReportingOrg(bytes32 _reference, bytes32 _orgRef) view returns (bytes32)",
    "function getLang(bytes32 _reference, bytes32 _orgRef) view returns (bytes32)",
    "function getCurrency(bytes32 _reference,  bytes32 _orgRef) view returns (bytes32)",
    "function getVersion(bytes32 _reference,  bytes32 _orgRef) view returns (bytes32)",
    "function getGeneratedTime(bytes32 _reference,  bytes32 _orgRef) view returns (bytes32)",
    "function getLastUpdatedTime(bytes32 _reference,  bytes32 _orgRef) view returns (bytes32)",

    "function getReportingOrgType(bytes32 _reference, bytes32 _reportingOrgRef) view returns (uint8)",
    "function getReportingOrgIsSecondary(bytes32 _reference, bytes32 _reportingOrgRef) view returns (bool)",

    "function getDocument(bytes32 _reference, bytes32 _docRef) view returns (tuple(bytes32 reportRef, bytes32 docRef, string title, bytes32 format, string url, bytes32 category, bytes32 countryCode, string desc, bytes32 lang, bytes32 date) _doc)",
    "function getDocumentTitle(bytes32 _reference, bytes32 _docRef) view returns (string)",
    "function getDocumentFormat(bytes32 _reference, bytes32 _docRef) view returns (bytes32)",
    "function getDocumentURL(bytes32 _reference, bytes32 _docRef) view returns (string)",
    "function getDocumentCategory(bytes32 _reference, bytes32 _docRef) view returns (bytes32)",
    "function getDocumentCountry(bytes32 _reference, bytes32 _docRef) view returns (bytes32)",
    "function getDocumentDescription(bytes32 _reference, bytes32 _docRef) view returns (string)",
    "function getDocumentLang(bytes32 _reference, bytes32 _docRef) view returns (bytes32)",
    "function getDocumentDate(bytes32 _reference, bytes32 _docRef) view returns (bytes32)",
  ]

  /* Ropsten addresses
  static organisationsAddress = "0x4A7cB305792eda1b7110d5484E15326ac7876398"
  static organisationReportsAddress = "0x0Cd0274C0FAa1b77fCc634313b499EFE5E457113"
  */

  static organisationsAddress = "0x4A7cB305792eda1b7110d5484E15326ac7876398"
  static organisationReportsAddress = "0x0Cd0274C0FAa1b77fCc634313b499EFE5E457113"
}

export { Paths, Blockchain, Contract }
