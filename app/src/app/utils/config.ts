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
  static readonly orgReportBudgetsWriter = '/create-organisation-report-budget'
  static readonly orgReader='/read-organisations'
  static readonly orgReportReader='/read-organisation-reports'
}

class Blockchain {

  // Using web3 via MetaMask, so these aren't necessary, really
  static readonly host = 'http://127.0.0.1'
  static readonly port = '8545'
  static readonly network = 'ropsten'

  static readonly checkInterval = 3000
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
    "function getOrganisationIdentifier(bytes32 _orgRef) view returns (string)"
  ]

  static organisationReportsABI = [
    "event SetReport(bytes32 orgRef, bytes32 reportRef, tuple(bytes32 version, bytes32 orgRef, bytes32 reportRef, tuple(bytes32 orgRef, uint8 orgType, bool isSecondary) reportingOrg, bytes32 lang, bytes32 currency, bytes32 lastUpdatedTime) report)",

    "function setReport(tuple(bytes32 version, bytes32 orgRef, bytes32 reportRef, tuple(bytes32 orgRef, uint8 orgType, bool isSecondary) reportingOrg, bytes32 lang, bytes32 currency, bytes32 lastUpdatedTime) report)@500000",

    "function getOrgExists(bytes32 orgRef) view returns (bool)",
    "function getReportExists(bytes32 orgRef, bytes32 reportRef) view returns (bool)",

    "function getNumOrgs() view returns (uint256)",
    "function getNumReports(bytes32 orgRef) view returns (uint256)",

    "function getOrganisationReference(uint256 _index) view returns (bytes32)",
    "function getReportReference(bytes32 orgRef, uint256 index) view returns (bytes32)",

    "function getReport(bytes32 orgRef, bytes32 reportRef) view returns (tuple(bytes32 version, bytes32 orgRef, bytes32 reportRef, tuple(bytes32 orgRef, uint8 orgType, bool isSecondary) reportingOrg, bytes32 lang, bytes32 currency, bytes32 lastUpdatedTime) report)",

    "function getLang(bytes32 orgRef, bytes32 reportRef) view returns (bytes32)",
    "function getCurrency(bytes32 orgRef, bytes32 reportRef) view returns (bytes32)",
    "function getVersion(bytes32 orgRef, bytes32 reportRef) view returns (bytes32)",
    "function getLastUpdatedTime(bytes32 orgRef, bytes32 reportRef) view returns (bytes32)",
    "function getReportingOrg(bytes32 orgRef, bytes32 reportRef) view returns (bytes32)",
    "function getReportingOrgType(bytes32 orgRef, bytes32 reportRef) view returns (uint8)",
    "function getReportingOrgIsSecondary(bytes32 orgRef, bytes32 reportRef) view returns (bool)"
  ]

  static organisationReportDocsABI = [
    "event SetDocument(bytes32 reportRef, bytes32 docRef, tuple(bytes32 reportRef, bytes32 docRef, string title, bytes32 format, string url, bytes32 category, bytes32 countryCode, string desc, bytes32 lang, bytes32 date) doc)",

    "function setDocument(tuple(bytes32 reportRef, bytes32 docRef, string title, bytes32 format, string url, bytes32 category, bytes32 countryCode, string desc, bytes32 lang, bytes32 date) doc)@500000",

    "function getReportDocExists(bytes32 reportRef, bytes32 docRef) view returns (bool)",

    "function getNumReportDocs(bytes32 reportRef) view returns (uint256)",

    "function getReportDocReference(bytes32 docRef, uint256 index) view returns (bytes32)",

    "function getDocument(bytes32 reportRef, bytes32 docRef) view returns (tuple(bytes32 reportRef, bytes32 docRef, string title, bytes32 format, string url, bytes32 category, bytes32 countryCode, string desc, bytes32 lang, bytes32 date) doc)",

    "function getDocumentTitle(bytes32 reportRef, bytes32 docRef) view returns (string)",
    "function getDocumentFormat(bytes32 reportRef, bytes32 docRef) view returns (bytes32)",
    "function getDocumentURL(bytes32 reportRef, bytes32 docRef) view returns (string)",
    "function getDocumentCategory(bytes32 reportRef, bytes32 docRef) view returns (bytes32)",
    "function getDocumentCountry(bytes32 reportRef, bytes32 docRef) view returns (bytes32)",
    "function getDocumentDescription(bytes32 reportRef, bytes32 docRef) view returns (string)",
    "function getDocumentLang(bytes32 reportRef, bytes32 docRef) view returns (bytes32)",
    "function getDocumentDate(bytes32 reportRef, bytes32 docRef) view returns (bytes32)"
  ]

  static organisationReportBudgetsABI = [
    "event SetTotalBudget(bytes32 reportRef, bytes32 budgetRef, tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)",

    "function setTotalBudget(tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)@500000",

    "function getReportExists(bytes32 reportRef) view returns (bool)",
    "function getTotalBudgetExists(bytes32 reportRef, bytes32 budgetRef) view returns (bool)",

    "function getNumReports() view returns (uint256)",
    "function getNumTotalBudgets(bytes32 reportRef) view returns (uint256)",

    "function getReportReference(uint256 index) view returns (bytes32)",
    "function getTotalBudgetReference(bytes32 reportRef, uint256 index) view returns (bytes32)",

    "function getTotalBudget(bytes32 reportRef, bytes32 budgetRef) view returns (tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)",
    "function getTotalBudgetLine(bytes32 reportRef, bytes32 budgetRef) view returns (bytes32)",
    "function getTotalBudgetValue(bytes32 reportRef, bytes32 budgetRef) view returns (uint256)",
    "function getTotalBudgetStatus(bytes32 reportRef, bytes32 budgetRef) view returns (bytes32)",
    "function getTotalBudgetStart(bytes32 reportRef, bytes32 budgetRef) view returns (bytes32)",
    "function getTotalBudgetEnd(bytes32 reportRef, bytes32 budgetRef) view returns (bytes32)"
  ]

  static organisationReportExpenditureABI = [
    "event SetExpenditure(bytes32 reportRef, bytes32 expenditureRef, tuple(bytes32 reportRef, bytes32 expenditureRef, bytes32 expenditureLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) expenditure)",

    "function setExpenditure(tuple(bytes32 reportRef, bytes32 expenditureRef, bytes32 expenditureLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) expenditure)@500000",

    "function getReportExists(bytes32 reportRef) view returns (bool)",
    "function getExpenditureExists(bytes32 reportRef, bytes32 expenditureRef) view returns (bool)",

    "function getNumReports() view returns (uint256)",
    "function getNumExpenditures(bytes32 reportRef) view returns (uint256)",

    "function getReportReference(uint256 index) view returns (bytes32)",
    "function getExpenditureReference(bytes32 reportRef, uint256 index) view returns (bytes32)",

    "function getExpenditure(bytes32 reportRef, bytes32 expenditureRef) view returns (tuple(bytes32 reportRef, bytes32 expenditureRef, bytes32 expenditureLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) expenditure)",
    "function getExpenditureLine(bytes32 reportRef, bytes32 budgetRef) view returns (bytes32)",
    "function getExpenditureValue(bytes32 reportRef, bytes32 expenditureRef) view returns (uint256)",
    "function getExpenditureStatus(bytes32 reportRef, bytes32 expenditureRef) view returns (bytes32)",
    "function getExpenditureStart(bytes32 reportRef, bytes32 expenditureRef) view returns (bytes32)",
    "function getExpenditureEnd(bytes32 reportRef, bytes32 expenditureRef) view returns (bytes32)"
  ]

  static organisationReportRecipientBudgetsABI = [
    "event SetRecipientOrgBudget(bytes32 _reportRef, bytes32 _budgetRef, tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 orgRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)",

    "function setRecipientOrgBudget(tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 orgRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)@500000",

    "function getReportExists(bytes32 _reportRef) view returns (bool)",
    "function getRecipientOrgBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) view returns (bool)",

    "function getNumReports() view returns (uint256)",
    "function getNumRecipientOrgBudgets(bytes32 _reportRef) view returns (uint256)",

    "function getReportReference(uint256 _index) view returns (bytes32)",
    "function getRecipientOrgReference(bytes32 _reportRef, uint256 _index) view returns (bytes32)",

    "function getRecipientOrgBudget(bytes32 _reportRef, bytes32 _recipientOrgRef) view returns (tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 orgRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)",
    "function getRecipientOrgBudgetLine(bytes32 _reportRef, bytes32 _budgetRef) view returns (bytes32)",
    "function getRecipientOrgBudgetValue(bytes32 _reportRef, bytes32 _recipientOrgRef) view returns (uint256)",
    "function getRecipientOrgBudgetStatus(bytes32 _reportRef, bytes32 _recipientOrgRef) view returns (bytes32)",
    "function getRecipientOrgBudgetStart(bytes32 _reportRef, bytes32 _recipientOrgRef) view returns (bytes32)",
    "function getRecipientOrgBudgetEnd(bytes32 _reportRef, bytes32 _recipientOrgRef) view returns (bytes32)"
  ]

  static organisationReportRegionBudgetsABI = [
    "event SetRegionBudget(bytes32 _reportRef, bytes32 _budgetRef, tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 regionRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)",

    "function setRegionBudget(tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 regionRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)@500000",

    "function getReportExists(bytes32 _reportRef) view returns (bool)",
    "function getRegionBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) view returns (bool)",

    "function getNumReports() view returns (uint256)",
    "function getNumRegionBudgets(bytes32 _reportRef) view returns (uint256)",

    "function getReportReference(uint256 _index) view returns (bytes32)",
    "function getRegionBudgetReference(bytes32 _reportRef, uint256 _index) view returns (bytes32)",

    "function getRegionsBudget(bytes32 _reportRef, bytes32 _regionRef) view returns (tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 regionRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)",
    "function getRegionsBudgetLine(bytes32 _reportRef, bytes32 _budgetRef) view returns (bytes32)",
    "function getRegionsBudgetValue(bytes32 _reportRef, bytes32 _regionRef) view returns (uint256)",
    "function getRegionsBudgetStatus(bytes32 _reportRef, bytes32 _regionRef) view returns (bytes32)",
    "function getRegionsBudgetStart(bytes32 _reportRef, bytes32 _regionRef) view returns (bytes32)",
    "function getRegionsBudgetEnd(bytes32 _reportRef, bytes32 _regionRef) view returns (bytes32)"
  ]

  static organisationReportCountryBudgetsABI = [
    "event SetCountryBudget(bytes32 _reportRef, bytes32 _budgetRef, tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 regionRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)",

    "function setCountryBudget(tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 countryRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)@500000",

    "function getReportExists(bytes32 _reportRef) view returns (bool)",
    "function getCountryBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) view returns (bool)",

    "function getNumReports() view returns (uint256)",
    "function getNumCountryBudgets(bytes32 _reportRef) view returns (uint256)",

    "function getReportReference(uint256 _index) view returns (bytes32)",
    "function getCountryBudgetReference(bytes32 _reportRef, uint256 _index) view returns (bytes32)",

    "function getCountryBudget(bytes32 _reportRef, bytes32 _countryRef) view returns (tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 countryRef, bytes32 budgetLine, tuple(uint256 value, bytes32 status, bytes32 start, bytes32 end) finance) budget)",
    "function getCountryBudgetLine(bytes32 _reportRef, bytes32 _budgetRef) view returns (bytes32)",
    "function getCountryBudgetValue(bytes32 _reportRef, bytes32 _countryRef) view returns (uint256)",
    "function getCountryBudgetStatus(bytes32 _reportRef, bytes32 _countryRef) view returns (bytes32)",
    "function getCountryBudgetStart(bytes32 _reportRef, bytes32 _countryRef) view returns (bytes32)",
    "function getCountryBudgetEnd(bytes32 _reportRef, bytes32 _countryRef) view returns (bytes32)"
  ]

  /* Ropsten addresses
  static organisationsAddress = "0x2C757C6390D9186F7e35C0796256B17d650df017"
  static organisationReportsAddress = "0x81759e5B8BCB38B05A678C2b98Eb17eaf4Ae9D71"
  */

  static organisationsAddress = "0x667CD537e1079044119438B0c8a62d59D557B55A"
  static organisationReportsAddress = "0xcA5784A7c10A7CFC6480072FB46f0C54dFbFB989"
  static organisationReportDocsAddress = "0x2C433289b8A829AA3F34287039083309F542C8C4"
  static organisationReportBudgetsAddress = "0x0A1E2F79698149724A6f02B549C9FCB216c3845C"
  static organisationReportExpenditureAddress = "0x0aD0Fb44357c3274Fbd6F47Cc60Ca60444354390"
  static organisationReportRecipientBudgetsAddress = "0x2FB258596E6359b9342E1e4d6617c72c54C893d9"
  static organisationReportRegionBudgetsAddress = "0xd3Fa22746d7d71f3386718A3F48f7b6a88073a4D"
  static organisationReportCountryBudgetsAddress = "0xe97b1Cb304DD13e64AA833061d276F3C40453325"
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
