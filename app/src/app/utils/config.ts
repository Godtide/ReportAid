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

    "function getOrganisationExists(bytes32 _orgRef) returns (bool)",
    "function getNumOrganisations() returns (uint256)",
    "function getOrganisationReference(uint256 _index) returns (bytes32)",
    "function getOrganisation(bytes32 _orgRef) returns (tuple(bytes32 orgRef, string name, string identifier) org)",
    "function getOrganisationName(bytes32 _orgRef) returns (string)",
    "function getOrganisationIdentifier(bytes32 _orgRef) returns (string)",
  ]

  static organisationReportsABI = [
    "event SetReport(bytes32 orgRef, bytes32 reportRef, tuple(bytes32 version, bytes32 orgRef, bytes32 reportRef, tuple(bytes32 orgRef, uint8 orgType, bool isSecondary) reportingOrg, bytes32 lang, bytes32 currency, bytes32 lastUpdatedTime) report)",

    "function setReport(tuple(bytes32 version, bytes32 orgRef, bytes32 reportRef, tuple(bytes32 orgRef, uint8 orgType, bool isSecondary) reportingOrg, bytes32 lang, bytes32 currency, bytes32 lastUpdatedTime) report)@500000",

    "function getOrgExists(bytes32 orgRef)  returns (bool)",
    "function getReportExists(bytes32 orgRef, bytes32 reportRef)  returns (bool)",

    "function getNumOrgs()  returns (uint256)",
    "function getNumReports(bytes32 orgRef)  returns (uint256)",

    "function getOrganisationReference(uint256 _index) returns (bytes32)",
    "function getReportReference(bytes32 orgRef, uint256 index)  returns (bytes32)",

    "function getReport(bytes32 orgRef, bytes32 reportRef)  returns (tuple(bytes32 version, bytes32 orgRef, bytes32 reportRef, tuple(bytes32 orgRef, uint8 orgType, bool isSecondary) reportingOrg, bytes32 lang, bytes32 currency, bytes32 lastUpdatedTime) report)",

    "function getLang(bytes32 orgRef, bytes32 reportRef)  returns (bytes32)",
    "function getCurrency(bytes32 orgRef, bytes32 reportRef)  returns (bytes32)",
    "function getVersion(bytes32 orgRef, bytes32 reportRef)  returns (bytes32)",
    "function getLastUpdatedTime(bytes32 orgRef, bytes32 reportRef)  returns (bytes32)",
    "function getReportingOrg(bytes32 orgRef, bytes32 reportRef)  returns (bytes32)",
    "function getReportingOrgType(bytes32 orgRef, bytes32 reportRef)  returns (uint8)",
    "function getReportingOrgIsSecondary(bytes32 orgRef, bytes32 reportRef)  returns (bool)",
  ]

  static organisationReportDocsABI = [
    "event SetDocument(bytes32 reportRef, bytes32 docRef, tuple(bytes32 reportRef, bytes32 docRef, string title, bytes32 format, string url, bytes32 category, bytes32 countryCode, string desc, bytes32 lang, bytes32 date) doc)",

    "function setDocument(tuple(bytes32 reportRef, bytes32 docRef, string title, bytes32 format, string url, bytes32 category, bytes32 countryCode, string desc, bytes32 lang, bytes32 date) doc)@500000",

    "function getReportDocExists(bytes32 reportRef, bytes32 docRef)  returns (bool)",

    "function getNumReportDocs(bytes32 reportRef)  returns (uint256)",

    "function getReportDocReference(bytes32 docRef, uint256 index)  returns (bytes32)",

    "function getDocument(bytes32 reportRef, bytes32 docRef)  returns (tuple(bytes32 reportRef, bytes32 docRef, string title, bytes32 format, string url, bytes32 category, bytes32 countryCode, string desc, bytes32 lang, bytes32 date) doc)",

    "function getDocumentTitle(bytes32 reportRef, bytes32 docRef)  returns (string)",
    "function getDocumentFormat(bytes32 reportRef, bytes32 docRef)  returns (bytes32)",
    "function getDocumentURL(bytes32 reportRef, bytes32 docRef)  returns (string)",
    "function getDocumentCategory(bytes32 reportRef, bytes32 docRef)  returns (bytes32)",
    "function getDocumentCountry(bytes32 reportRef, bytes32 docRef)  returns (bytes32)",
    "function getDocumentDescription(bytes32 reportRef, bytes32 docRef)  returns (string)",
    "function getDocumentLang(bytes32 reportRef, bytes32 docRef)  returns (bytes32)",
    "function getDocumentDate(bytes32 reportRef, bytes32 docRef)  returns (bytes32)"
  ]

  static organisationReportBudgetsABI = [
    "event SetTotalBudget(bytes32 reportRef, bytes32 budgetRef, tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)",

    "function setTotalBudget(tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)@500000",

    "function getReportExists(bytes32 reportRef) returns (bool)",
    "function getTotalBudgetExists(bytes32 reportRef, bytes32 budgetRef) returns (bool)",

    "function getNumReports() returns (uint256)",
    "function getNumTotalBudgets(bytes32 reportRef) returns (uint256)",

    "function getReportReference(uint256 index) returns (bytes32)",
    "function getTotalBudgetReference(bytes32 reportRef, uint256 index) returns (bytes32)",

    "function getTotalBudget(bytes32 reportRef, bytes32 budgetRef) returns (tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)",
    "function getTotalBudgetLine(bytes32 reportRef, bytes32 budgetRef) returns (bytes32)",
    "function getTotalBudgetValue(bytes32 reportRef, bytes32 budgetRef) returns (uint256)",
    "function getTotalBudgetStatus(bytes32 reportRef, bytes32 budgetRef) returns (bytes32)",
    "function getTotalBudgetStart(bytes32 reportRef, bytes32 budgetRef) returns (bytes32)",
    "function getTotalBudgetEnd(bytes32 reportRef, bytes32 budgetRef) returns (bytes32)"
  ]

  static organisationReportExpenditureABI = [
    "event SetExpenditure(bytes32 reportRef, bytes32 expenditureRef, tuple(bytes32 reportRef, bytes32 expenditureRef, bytes32 expenditureLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) expenditure)",

    "function setExpenditure(tuple(bytes32 reportRef, bytes32 expenditureRef, bytes32 expenditureLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) expenditure)@500000",

    "function getReportExists(bytes32 reportRef) returns (bool)",
    "function getExpenditureExists(bytes32 reportRef, bytes32 expenditureRef) returns (bool)",

    "function getNumReports() returns (uint256)",
    "function getNumExpenditures(bytes32 reportRef) returns (uint256)",

    "function getReportReference(uint256 index) returns (bytes32)",
    "function getExpenditureReference(bytes32 reportRef, uint256 index) returns (bytes32)",

    "function getExpenditure(bytes32 reportRef, bytes32 expenditureRef) returns (tuple(bytes32 reportRef, bytes32 expenditureRef, bytes32 expenditureLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) expenditure)",
    "function getExpenditureLine(bytes32 reportRef, bytes32 budgetRef) returns (bytes32)",
    "function getExpenditureValue(bytes32 reportRef, bytes32 expenditureRef) returns (uint256)",
    "function getExpenditureStatus(bytes32 reportRef, bytes32 expenditureRef) returns (bytes32)",
    "function getExpenditureStart(bytes32 reportRef, bytes32 expenditureRef) returns (bytes32)",
    "function getExpenditureEnd(bytes32 reportRef, bytes32 expenditureRef) returns (bytes32)"
  ]

  static organisationReportRecipientBudgetsABI = [
    "event SetRecipientOrgBudget(bytes32 _reportRef, bytes32 _budgetRef, tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 orgRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)",

    "function setRecipientOrgBudget(tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 orgRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)@500000",

    "function getReportExists(bytes32 _reportRef) returns (bool)",
    "function getRecipientOrgBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) returns (bool)",

    "function getNumReports() returns (uint256)",
    "function getNumRecipientOrgBudgets(bytes32 _reportRef) returns (uint256)",

    "function getReportReference(uint256 _index) returns (bytes32)",
    "function getRecipientOrgReference(bytes32 _reportRef, uint256 _index) returns (bytes32)",

    "function getRecipientOrgBudget(bytes32 _reportRef, bytes32 _recipientOrgRef) returns (tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 orgRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)",
    "function getRecipientOrgBudgetLine(bytes32 _reportRef, bytes32 _budgetRef) returns (bytes32)",
    "function getRecipientOrgBudgetValue(bytes32 _reportRef, bytes32 _recipientOrgRef) returns (uint256)",
    "function getRecipientOrgBudgetStatus(bytes32 _reportRef, bytes32 _recipientOrgRef) returns (bytes32)",
    "function getRecipientOrgBudgetStart(bytes32 _reportRef, bytes32 _recipientOrgRef) returns (bytes32)",
    "function getRecipientOrgBudgetEnd(bytes32 _reportRef, bytes32 _recipientOrgRef) returns (bytes32)"
  ]

  static organisationReportRegionBudgetsABI = [
    "event SetRegionBudget(bytes32 _reportRef, bytes32 _budgetRef, tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 regionRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)",

    "function setRegionBudget(tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 regionRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)@500000",

    "function getReportExists(bytes32 _reportRef) returns (bool)",
    "function getRegionBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) returns (bool)",

    "function getNumReports() returns (uint256)",
    "function getNumRegionBudgets(bytes32 _reportRef) returns (uint256)",

    "function getReportReference(uint256 _index) returns (bytes32)",
    "function getRegionBudgetReference(bytes32 _reportRef, uint256 _index) returns (bytes32)",

    "function getRegionsBudget(bytes32 _reportRef, bytes32 _regionRef) returns (tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 regionRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)",
    "function getRegionsBudgetLine(bytes32 _reportRef, bytes32 _budgetRef) returns (bytes32)",
    "function getRegionsBudgetValue(bytes32 _reportRef, bytes32 _regionRef) returns (uint256)",
    "function getRegionsBudgetStatus(bytes32 _reportRef, bytes32 _regionRef) returns (bytes32)",
    "function getRegionsBudgetStart(bytes32 _reportRef, bytes32 _regionRef) returns (bytes32)",
    "function getRegionsBudgetEnd(bytes32 _reportRef, bytes32 _regionRef) returns (bytes32)"

  ]

  static organisationReportCountryBudgetsABI = [
    "event SetCountryBudget(bytes32 _reportRef, bytes32 _budgetRef, tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 regionRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)",

    "function setCountryBudget(tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 countryRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)@500000",

    "function getReportExists(bytes32 _reportRef) returns (bool)",
    "function getCountryBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) returns (bool)",

    "function getNumReports() returns (uint256)",
    "function getNumCountryBudgets(bytes32 _reportRef) returns (uint256)",

    "function getReportReference(uint256 _index) returns (bytes32)",
    "function getCountryBudgetReference(bytes32 _reportRef, uint256 _index) returns (bytes32)",

    "function getCountryBudget(bytes32 _reportRef, bytes32 _countryRef) returns (tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 countryRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)",
    "function getCountryBudgetLine(bytes32 _reportRef, bytes32 _budgetRef) returns (bytes32)",
    "function getCountryBudgetValue(bytes32 _reportRef, bytes32 _countryRef) returns (uint256)",
    "function getCountryBudgetStatus(bytes32 _reportRef, bytes32 _countryRef) returns (bytes32)",
    "function getCountryBudgetStart(bytes32 _reportRef, bytes32 _countryRef) returns (bytes32)",
    "function getCountryBudgetEnd(bytes32 _reportRef, bytes32 _countryRef) returns (bytes32)",
  ]

  /* Ropsten addresses
  static organisationsAddress = "0x2C757C6390D9186F7e35C0796256B17d650df017"
  static organisationReportsAddress = "0x81759e5B8BCB38B05A678C2b98Eb17eaf4Ae9D71"
  */

  static organisationsAddress = "0x379abC7EDF25A9D0aA8401713657207f56CbEe13"
  static organisationReportsAddress = "0x65038dCf7547ed3a288DBB0c36061c92365E1d80"
  static organisationReportDocsAddress = "0xDb65296eA3bf2aFd42b644169C0FdaD046393bd8"
  static organisationReportBudgetsAddress = "0xb5d2646873c8AeC60AD44A504Bb4710E21218743"
  static organisationReportExpenditureAddress = "0x3E900F26a6f632B96745e2ED782067854299027a"
  static organisationReportRecipientBudgetsAddress = "0x15D6181565677Da4355201D99EcF262667Ae801A"
  static organisationReportRegionBudgetsAddress = "0xcC1b5DFea0764d5B501ED705e1Bc65076C2fE21d"
  static organisationReportCountryBudgetsAddress = "0xc85DdD83fBb2e5581A5525E085F9D5c08aF0E3C1"
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
