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
  static readonly orgReportReader='/read-organisation-reports'
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

    "function setOrganisation(tuple(bytes32 orgRef, string name, string identifier) _org)@150000",

    "function getOrganisationExists(bytes32 _orgRef) view returns (bool)",
    "function getNumOrganisations() view returns (uint256)",
    "function getOrganisationReference(uint256 _index) view returns (bytes32)",
    "function getOrganisation(bytes32 _orgRef) view returns (tuple(bytes32 orgRef, string name, string identifier) org)",
    "function getOrganisationName(bytes32 _orgRef) view returns (string)",
    "function getOrganisationIdentifier(bytes32 _orgRef) view returns (string)",
  ]

  static organisationReportsABI = [
    "event SetReport(bytes32 orgRef, bytes32 reportRef, tuple(bytes32 version, bytes32 orgRef, bytes32 reportRef, tuple(bytes32 orgRef, uint8 orgType, bool isSecondary), bytes32 lang, bytes32 currency, bytes32 lastUpdatedTime) report)",
    "event SetDocument(bytes32 reportRef, bytes32 docRef, tuple(bytes32 reportRef, bytes32 docRef, string title, bytes32 format, string url, bytes32 category, bytes32 countryCode, string desc, bytes32 lang, bytes32 date) doc)",

    "function setReport(tuple(bytes32 version, bytes32 orgRef, bytes32 reportRef, tuple(bytes32 orgRef, uint8 orgType, bool isSecondary) reportingOrg, bytes32 lang, bytes32 currency, bytes32 lastUpdatedTime) report)@500000",
    "function setDocument(tuple(bytes32 reportRef, bytes32 docRef, string title, bytes32 format, string url, bytes32 category, bytes32 countryCode, string desc, bytes32 lang, bytes32 date) doc)",

    "function getOrgExists(bytes32 orgRef)  view returns (bool)",
    "function getReportExists(bytes32 orgRef, bytes32 reportRef)  view returns (bool)",
    "function getReportDocExists(bytes32 reportRef, bytes32 docRef)  view returns (bool)",

    "function getNumOrgs()  view returns (uint256)",
    "function getOrganisationReference(uint256 _index) view returns (bytes32)",
    "function getNumReports(bytes32 orgRef)  view returns (uint256)",
    "function getNumReportDocs(bytes32 reportRef)  view returns (uint256)",

    "function getReportReference(bytes32 orgRef, uint256 index)  view returns (bytes32)",
    "function getReportDocReference(bytes32 docRef, uint256 index)  view returns (bytes32)",

    "function getReport(bytes32 orgRef, bytes32 reportRef)  view returns (Report memory)",

    "function getLang(bytes32 orgRef, bytes32 reportRef)  view returns (bytes32)",
    "function getCurrency(bytes32 orgRef, bytes32 reportRef)  view returns (bytes32)",
    "function getVersion(bytes32 orgRef, bytes32 reportRef)  view returns (bytes32)",
    "function getLastUpdatedTime(bytes32 orgRef, bytes32 reportRef)  view returns (bytes32)",
    "function getReportingOrg(bytes32 orgRef, bytes32 reportRef)  view returns (bytes32)",
    "function getReportingOrgType(bytes32 orgRef, bytes32 reportRef)  view returns (uint8)",
    "function getReportingOrgIsSecondary(bytes32 orgRef, bytes32 reportRef)  view returns (bool)",

    "function getDocument(bytes32 reportRef, bytes32 docRef)  view returns (Document)",
    "function getDocumentTitle(bytes32 reportRef, bytes32 docRef)  view returns (string)",
    "function getDocumentFormat(bytes32 reportRef, bytes32 docRef)  view returns (bytes32)",
    "function getDocumentURL(bytes32 reportRef, bytes32 docRef)  view returns (string memory)",
    "function getDocumentCategory(bytes32 reportRef, bytes32 docRef)  view returns (bytes32)",
    "function getDocumentCountry(bytes32 reportRef, bytes32 docRef)  view returns (bytes32)",
    "function getDocumentDescription(bytes32 reportRef, bytes32 docRef)  view returns (string)",
    "function getDocumentLang(bytes32 reportRef, bytes32 docRef)  view returns (bytes32)",
    "function getDocumentDate(bytes32 reportRef, bytes32 docRef)  view returns (bytes32)"
  ]

  /* Ropsten addresses
  static organisationsAddress = "0x2C757C6390D9186F7e35C0796256B17d650df017"
  static organisationReportsAddress = "0x81759e5B8BCB38B05A678C2b98Eb17eaf4Ae9D71"
  */

  static organisationsAddress = "0x5fd0bf6818576c58E067749821F1C41aAf97B16A"
  static organisationReportsAddress = "0x1dF54fB0bE96038017A1C0563477FabeaEf4A9E3"
}

class Helpers {

  static countryCodes = ['AD', 'AE', 'AF', 'AG', 'AI', 'AL', 'AM', 'AO', 'AQ', 'AR', 'AS', 'AT', 'AU', 'AW', 'AX', 'AZ', 'BA', 'BB', 'BD', 'BE', 'BF', 'BG', 'BH', 'BI', 'BJ', 'BL', 'BM', 'BN', 'BO', 'BQ', 'BR', 'BS', 'BT', 'BV', 'BW', 'BY', 'BZ', 'CA', 'CC', 'CD', 'CF', 'CG', 'CH', 'CI', 'CK', 'CL', 'CM', 'CN', 'CO', 'CR', 'CU', 'CV', 'CW', 'CX', 'CY', 'CZ', 'DE', 'DJ', 'DK', 'DM', 'DO', 'DZ', 'EC', 'EE', 'EG', 'EH', 'ER', 'ES', 'ET', 'FI', 'FJ', 'FK', 'FM', 'FO', 'FR', 'GA', 'GB', 'GD', 'GE', 'GF', 'GG', 'GH', 'GI', 'GL', 'GM', 'GN', 'GP', 'GQ', 'GR', 'GS', 'GT', 'GU', 'GW', 'GY', 'HK', 'HM', 'HN', 'HR', 'HT', 'HU', 'ID', 'IE', 'IL', 'IM', 'IN', 'IO', 'IQ', 'IR', 'IS', 'IT', 'JE', 'JM', 'JO', 'JP', 'KE', 'KG', 'KH', 'KI', 'KM', 'KN', 'KP', 'KR', 'KW', 'KY', 'KZ', 'LA', 'LB', 'LC', 'LI', 'LK', 'LR', 'LS', 'LT', 'LU', 'LV', 'LY', 'MA', 'MC', 'MD', 'ME', 'MF', 'MG', 'MH', 'MK', 'ML', 'MM', 'MN', 'MO', 'MP', 'MQ', 'MR', 'MS', 'MT', 'MU', 'MV', 'MW', 'MX', 'MY', 'MZ', 'NA', 'NC', 'NE', 'NF', 'NG', 'NI', 'NL', 'NO', 'NP', 'NR', 'NU', 'NZ', 'OM', 'PA', 'PE', 'PF', 'PG', 'PH', 'PK', 'PL', 'PM', 'PN', 'PR', 'PS', 'PT', 'PW', 'PY', 'QA', 'RE', 'RO', 'RS', 'RU', 'RW', 'SA', 'SB', 'SC', 'SD', 'SE', 'SG', 'SH', 'SI', 'SJ', 'SK', 'SL', 'SM', 'SN', 'SO', 'SR', 'SS', 'ST', 'SV', 'SX', 'SY', 'SZ', 'TC', 'TD', 'TF', 'TG', 'TH', 'TJ', 'TK', 'TL', 'TM', 'TN', 'TO', 'TR', 'TT', 'TV', 'TW', 'TZ', 'UA', 'UG', 'UM', 'US', 'UY', 'UZ', 'VA', 'VC', 'VE', 'VG', 'VI', 'VN', 'VU', 'WF', 'WS', 'YE', 'YT', 'ZA', 'ZM', 'ZW']

  static currencyCodes = ["AED" ,"AFN" ,"ALL" ,"AMD" ,"ANG" ,"AOA" ,"ARS" ,"AUD" ,"AWG" ,"AZN" ,"BAM" ,"BBD" ,"BDT" ,"BGN" ,"BHD" ,"BIF" ,"BMD" ,"BND" ,"BOB" ,"BRL" ,"BSD" ,"BTN" ,"BWP" ,"BYN" ,"BZD" ,"CAD" ,"CDF" ,"CHF" ,"CLP" ,"CNY" ,"COP" ,"CRC" ,"CUC" ,"CUP" ,"CVE" ,"CZK" ,"DJF" ,"DKK" ,"DOP" ,"DZD" ,"EGP" ,"ERN" ,"ETB" ,"EUR" ,"FJD" ,"FKP" ,"GBP" ,"GEL" ,"GGP" ,"GHS" ,"GIP" ,"GMD" ,"GNF" ,"GTQ" ,"GYD" ,"HKD" ,"HNL" ,"HRK" ,"HTG" ,"HUF" ,"IDR" ,"ILS" ,"IMP" ,"INR" ,"IQD" ,"IRR" ,"ISK" ,"JEP" ,"JMD" ,"JOD" ,"JPY" ,"KES" ,"KGS" ,"KHR" ,"KMF" ,"KPW" ,"KRW" ,"KWD" ,"KYD" ,"KZT" ,"LAK" ,"LBP" ,"LKR" ,"LRD" ,"LSL" ,"LYD" ,"MAD" ,"MDL" ,"MGA" ,"MKD" ,"MMK" ,"MNT" ,"MOP" ,"MRU" ,"MUR" ,"MVR" ,"MWK" ,"MXN" ,"MYR" ,"MZN" ,"NAD" ,"NGN" ,"NIO" ,"NOK" ,"NPR" ,"NZD" ,"OMR" ,"PAB" ,"PEN" ,"PGK" ,"PHP" ,"PKR" ,"PLN" ,"PYG" ,"QAR" ,"RON" ,"RSD" ,"RUB" ,"RWF" ,"SAR" ,"SBD" ,"SCR" ,"SDG" ,"SEK" ,"SGD" ,"SHP" ,"SLL" ,"SOS" ,"SPL*" ,"SRD" ,"STN" ,"SVC" ,"SYP" ,"SZL" ,"THB" ,"TJS" ,"TMT" ,"TND" ,"TOP" ,"TRY" ,"TTD" ,"TVD" ,"TWD" ,"TZS" ,"UAH" ,"UGX" ,"USD" ,"UYU" ,"UZS" ,"VEF" ,"VND" ,"VUV" ,"WST" ,"XAF" ,"XCD" ,"XDR" ,"XOF" ,"XPF" ,"YER" ,"ZAR" ,"ZMW" ,"ZWD"]

  static organisationCodes = [
    {code: 10, type: "Government"},
    {code: 11, type: "Local Government"},
    {code: 15, type: "Other Public Sector"},
    {code: 21, type: "International NGO"},
    {code: 22, type: "National NGO"},
    {code: 23, type: "Regional NGO"},
    {code: 24, type: "Partner Country based NGO"},
    {code: 30, type: "Public Private Partnership"},
    {code: 40, type: "Multilateral"},
    {code: 60, type: "Foundation"},
    {code: 70, type: "Private Sector"},
    {code: 71, type: "Private Sector in Provider Country"},
    {code: 72, type: "Private Sector in Aid Recipient Country"},
    {code: 73, type: "Private Sector in Third Country"},
    {code: 80, type: "Academic, Training and Research"},
    {code: 90, type: "Other"}
  ]

  static reportVersions = [ "2.03" ]
}

export { Paths, Blockchain, Contract, Helpers }
