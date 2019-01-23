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
  static readonly orgReportBudgetsReader = '/read-organisation-report-budgets'
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
    "event SetTotalBudget(bytes32 reportRef, bytes32 budgetRef, tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 budgetLine, tuple(uint256 value, uint8 status, bytes32 start, bytes32 end) finance) budget)",

    "function setTotalBudget(tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 budgetLine, tuple(uint256 value, uint8 status, bytes32 start, bytes32 end) finance) budget)@500000",

    "function getReportExists(bytes32 reportRef) view returns (bool)",
    "function getTotalBudgetExists(bytes32 reportRef, bytes32 budgetRef) view returns (bool)",

    "function getNumReports() view returns (uint256)",
    "function getNumTotalBudgets(bytes32 reportRef) view returns (uint256)",

    "function getReportReference(uint256 index) view returns (bytes32)",
    "function getTotalBudgetReference(bytes32 reportRef, uint256 index) view returns (bytes32)",

    "function getTotalBudget(bytes32 reportRef, bytes32 budgetRef) view returns (tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 budgetLine, tuple(uint256 value, uint8 status, bytes32 start, bytes32 end) finance) budget)",
    "function getTotalBudgetLine(bytes32 reportRef, bytes32 budgetRef) view returns (bytes32)",
    "function getTotalBudgetValue(bytes32 reportRef, bytes32 budgetRef) view returns (uint256)",
    "function getTotalBudgetStatus(bytes32 reportRef, bytes32 budgetRef) view returns (uint8)",
    "function getTotalBudgetStart(bytes32 reportRef, bytes32 budgetRef) view returns (bytes32)",
    "function getTotalBudgetEnd(bytes32 reportRef, bytes32 budgetRef) view returns (bytes32)"
  ]

  static organisationReportExpenditureABI = [
    "event SetExpenditure(bytes32 reportRef, bytes32 expenditureRef, tuple(bytes32 reportRef, bytes32 expenditureRef, bytes32 expenditureLine, tuple(uint256 value, uint8 status, bytes32 start, bytes32 end) finance) expenditure)",

    "function setExpenditure(tuple(bytes32 reportRef, bytes32 expenditureRef, bytes32 expenditureLine, tuple(uint256 value, uint8 status, bytes32 start, bytes32 end) finance) expenditure)@500000",

    "function getReportExists(bytes32 reportRef) view returns (bool)",
    "function getExpenditureExists(bytes32 reportRef, bytes32 expenditureRef) view returns (bool)",

    "function getNumReports() view returns (uint256)",
    "function getNumExpenditures(bytes32 reportRef) view returns (uint256)",

    "function getReportReference(uint256 index) view returns (bytes32)",
    "function getExpenditureReference(bytes32 reportRef, uint256 index) view returns (bytes32)",

    "function getExpenditure(bytes32 reportRef, bytes32 expenditureRef) view returns (tuple(bytes32 reportRef, bytes32 expenditureRef, bytes32 expenditureLine, tuple(uint256 value, uint8 status, bytes32 start, bytes32 end) finance) expenditure)",
    "function getExpenditureLine(bytes32 reportRef, bytes32 budgetRef) view returns (bytes32)",
    "function getExpenditureValue(bytes32 reportRef, bytes32 expenditureRef) view returns (uint256)",
    "function getExpenditureStatus(bytes32 reportRef, bytes32 expenditureRef) view returns (uint8)",
    "function getExpenditureStart(bytes32 reportRef, bytes32 expenditureRef) view returns (bytes32)",
    "function getExpenditureEnd(bytes32 reportRef, bytes32 expenditureRef) view returns (bytes32)"
  ]

  static organisationReportRecipientBudgetsABI = [
    "event SetRecipientOrgBudget(bytes32 _reportRef, bytes32 _budgetRef, tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 orgRef, bytes32 budgetLine, tuple(uint256 value, uint8 status, bytes32 start, bytes32 end) finance) budget)",

    "function setRecipientOrgBudget(tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 orgRef, bytes32 budgetLine, tuple(uint256 value, uint8 status, bytes32 start, bytes32 end) finance) budget)@500000",

    "function getReportExists(bytes32 _reportRef) view returns (bool)",
    "function getRecipientOrgBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) view returns (bool)",

    "function getNumReports() view returns (uint256)",
    "function getNumRecipientOrgBudgets(bytes32 _reportRef) view returns (uint256)",

    "function getReportReference(uint256 _index) view returns (bytes32)",
    "function getRecipientOrgReference(bytes32 _reportRef, uint256 _index) view returns (bytes32)",

    "function getRecipientOrgBudget(bytes32 _reportRef, bytes32 _recipientOrgRef) view returns (tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 orgRef, bytes32 budgetLine, tuple(uint256 value, uint8 status, bytes32 start, bytes32 end) finance) budget)",
    "function getRecipientOrgBudgetLine(bytes32 _reportRef, bytes32 _budgetRef) view returns (bytes32)",
    "function getRecipientOrgBudgetValue(bytes32 _reportRef, bytes32 _recipientOrgRef) view returns (uint256)",
    "function getRecipientOrgBudgetStatus(bytes32 _reportRef, bytes32 _recipientOrgRef) view returns (uint8)",
    "function getRecipientOrgBudgetStart(bytes32 _reportRef, bytes32 _recipientOrgRef) view returns (bytes32)",
    "function getRecipientOrgBudgetEnd(bytes32 _reportRef, bytes32 _recipientOrgRef) view returns (bytes32)"
  ]

  static organisationReportRegionBudgetsABI = [
    "event SetRegionBudget(bytes32 _reportRef, bytes32 _budgetRef, tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 regionRef, bytes32 budgetLine, tuple(uint256 value, uint8 status, bytes32 start, bytes32 end) finance) budget)",

    "function setRegionBudget(tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 regionRef, bytes32 budgetLine, tuple(uint256 value, uint8 status, bytes32 start, bytes32 end) finance) budget)@500000",

    "function getReportExists(bytes32 _reportRef) view returns (bool)",
    "function getRegionBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) view returns (bool)",

    "function getNumReports() view returns (uint256)",
    "function getNumRegionBudgets(bytes32 _reportRef) view returns (uint256)",

    "function getReportReference(uint256 _index) view returns (bytes32)",
    "function getRegionBudgetReference(bytes32 _reportRef, uint256 _index) view returns (bytes32)",

    "function getRegionsBudget(bytes32 _reportRef, bytes32 _regionRef) view returns (tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 regionRef, bytes32 budgetLine, tuple(uint256 value, uint8 status, bytes32 start, bytes32 end) finance) budget)",
    "function getRegionsBudgetLine(bytes32 _reportRef, bytes32 _budgetRef) view returns (bytes32)",
    "function getRegionsBudgetValue(bytes32 _reportRef, bytes32 _regionRef) view returns (uint256)",
    "function getRegionsBudgetStatus(bytes32 _reportRef, bytes32 _regionRef) view returns (uint8)",
    "function getRegionsBudgetStart(bytes32 _reportRef, bytes32 _regionRef) view returns (bytes32)",
    "function getRegionsBudgetEnd(bytes32 _reportRef, bytes32 _regionRef) view returns (bytes32)"
  ]

  static organisationReportCountryBudgetsABI = [
    "event SetCountryBudget(bytes32 _reportRef, bytes32 _budgetRef, tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 regionRef, bytes32 budgetLine, tuple(uint256 value, uint8 status, bytes32 start, bytes32 end) finance) budget)",

    "function setCountryBudget(tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 countryRef, bytes32 budgetLine, tuple(uint256 value, uint8 status, bytes32 start, bytes32 end) finance) budget)@500000",

    "function getReportExists(bytes32 _reportRef) view returns (bool)",
    "function getCountryBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) view returns (bool)",

    "function getNumReports() view returns (uint256)",
    "function getNumCountryBudgets(bytes32 _reportRef) view returns (uint256)",

    "function getReportReference(uint256 _index) view returns (bytes32)",
    "function getCountryBudgetReference(bytes32 _reportRef, uint256 _index) view returns (bytes32)",

    "function getCountryBudget(bytes32 _reportRef, bytes32 _countryRef) view returns (tuple(bytes32 reportRef, bytes32 budgetRef, bytes32 countryRef, bytes32 budgetLine, tuple(uint256 value, uint8 status, bytes32 start, bytes32 end) finance) budget)",
    "function getCountryBudgetLine(bytes32 _reportRef, bytes32 _budgetRef) view returns (bytes32)",
    "function getCountryBudgetValue(bytes32 _reportRef, bytes32 _countryRef) view returns (uint256)",
    "function getCountryBudgetStatus(bytes32 _reportRef, bytes32 _countryRef) view returns (uint8)",
    "function getCountryBudgetStart(bytes32 _reportRef, bytes32 _countryRef) view returns (bytes32)",
    "function getCountryBudgetEnd(bytes32 _reportRef, bytes32 _countryRef) view returns (bytes32)"
  ]

  /* Ropsten addresses
  static organisationsAddress = "0x2C757C6390D9186F7e35C0796256B17d650df017"
  static organisationReportsAddress = "0x81759e5B8BCB38B05A678C2b98Eb17eaf4Ae9D71"
  */

  static organisationsAddress = "0x5fd0bf6818576c58E067749821F1C41aAf97B16A"
  static organisationReportsAddress = "0x1dF54fB0bE96038017A1C0563477FabeaEf4A9E3"
  static organisationReportDocsAddress = "0xFa9f7680705968660d36F34D080d5fEeD0614221"
  static organisationReportBudgetsAddress = "0xB155E22D9598cC0e635792070A888127Ae349B0c"
  static organisationReportExpenditureAddress = "0x85f1116DF7FCFE73bc511Ac1a6a69BaB0A3af1dA"
  static organisationReportRecipientBudgetsAddress = "0x7461eB577da59CBEE2618BB82c0d67311AE89960"
  static organisationReportRegionBudgetsAddress = "0x7952136EB509C59bFe8393a0BAeB17D3a5E0a400"
  static organisationReportCountryBudgetsAddress = "0x391Ef15D0640b87c6Fbaa555CaE2ed29dfd9F5c1"

  /*static organisationsAddress = "0x667CD537e1079044119438B0c8a62d59D557B55A"
  static organisationReportsAddress = "0xcA5784A7c10A7CFC6480072FB46f0C54dFbFB989"
  static organisationReportDocsAddress = "0x2C433289b8A829AA3F34287039083309F542C8C4"
  static organisationReportBudgetsAddress = "0x0A1E2F79698149724A6f02B549C9FCB216c3845C"
  static organisationReportExpenditureAddress = "0x0aD0Fb44357c3274Fbd6F47Cc60Ca60444354390"
  static organisationReportRecipientBudgetsAddress = "0x2FB258596E6359b9342E1e4d6617c72c54C893d9"
  static organisationReportRegionBudgetsAddress = "0xd3Fa22746d7d71f3386718A3F48f7b6a88073a4D"
  static organisationReportCountryBudgetsAddress = "0xe97b1Cb304DD13e64AA833061d276F3C40453325"*/
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

  static oecdDacCrs = [
    {code: 1, name: "Austria"},
    {code: 2, name: "Belgium"},
    {code: 3, name: "Denmark"},
    {code: 4, name: "France"},
    {code: 5, name: "Germany"},
    {code: 6, name: "Italy"},
    {code: 7, name: "Netherlands"},
    {code: 8, name: "Norway"},
    {code: 9, name: "Portugal"},
    {code: 10, name: "Sweden"},
    {code: 11, name: "Switzerland"},
    {code: 12, name: "United Kingdom"},
    {code: 18, name: "Finland"},
    {code: 20, name: "Iceland"},
    {code: 21, name: "Ireland"},
    {code: 22, name: "Luxembourg"},
    {code: 30, name: "Cyprus"},
    {code: 40, name: "Greece"},
    {code: 45, name: "Malta"},
    {code: 50, name: "Spain"},
    {code: 55, name: "Turkey"},
    {code: 61, name: "Slovenia"},
    {code: 62, name: "Croatia"},
    {code: 68, name: "Czech Republic"},
    {code: 69, name: "Slovak Republic"},
    {code: 70, name: "Liechtenstein"},
    {code: 72, name: "Bulgaria"},
    {code: 75, name: "Hungary"},
    {code: 76, name: "Poland"},
    {code: 77, name: "Romania"},
    {code: 82, name: "Estonia"},
    {code: 83, name: "Latvia"},
    {code: 84, name: "Lithuania"},
    {code: 87, name: "Russia"},
    {code: 104, name: "Nordic Development Fund"},
    {code: 130, name: "Algeria"},
    {code: 133, name: "Libya"},
    {code: 301, name: "Canada"},
    {code: 302, name: "United States"},
    {code: 358, name: "Mexico"},
    {code: 543, name: "Iraq"},
    {code: 546, name: "Israel"},
    {code: 552, name: "Kuwait"},
    {code: 561, name: "Qatar"},
    {code: 566, name: "Saudi Arabia"},
    {code: 576, name: "United Arab Emirates"},
    {code: 611, name: "Azerbaijan"},
    {code: 613, name: "Kazakhstan"},
    {code: 701, name: "Japan"},
    {code: 732, name: "Chinese Taipei"},
    {code: 742, name: "Korea"},
    {code: 764, name: "Thailand"},
    {code: 765, name: "Timor-Leste"},
    {code: 801, name: "Australia"},
    {code: 807, name: "UNEP"},
    {code: 811, name: "Global Environment Facility"},
    {code: 812, name: "Montreal Protocol"},
    {code: 820, name: "New Zealand"},
    {code: 901, name: "International Bank for Reconstruction and Development"},
    {code: 902, name: "Multilateral Investment Guarantee Agency"},
    {code: 903, name: "International Finance Corporation"},
    {code: 905, name: "International Development Association"},
    {code: 906, name: "Caribbean Development Bank"},
    {code: 907, name: "International Monetary Fund"},
    {code: 909, name: "Inter-American Development Bank"},
    {code: 912, name: "IDB Special Fund"},
    {code: 913, name: "African Development Bank"},
    {code: 914, name: "African Development Fund"},
    {code: 915, name: "Asian Development Bank"},
    {code: 916, name: "AsDB Special Funds"},
    {code: 918, name: "EU Institutions"},
    {code: 921, name: "Arab Fund (AFESD)"},
    {code: 923, name: "UN Peacebuilding Fund"},
    {code: 926, name: "Council of Europe"},
    {code: 928, name: "World Health Organisation"},
    {code: 932, name: "Food and Agriculture Organisation"},
    {code: 940, name: "International Labour Organisation"},
    {code: 944, name: "International Atomic Energy Agency"},
    {code: 948, name: "UNECE"},
    {code: 951, name: "OPEC Fund for International Development"},
    {code: 952, name: "OAPEC"},
    {code: 953, name: "Arab Bank for Economic Development in Africa"},
    {code: 954, name: "Special Arab Aid Fund for Africa"},
    {code: 956, name: "IMF Trust Fund"},
    {code: 958, name: "IMF (Concessional Trust Funds)"},
    {code: 959, name: "UNDP"},
    {code: 960, name: "UNTA"},
    {code: 963, name: "UNICEF"},
    {code: 964, name: "UNRWA"},
    {code: 966, name: "WFP"},
    {code: 967, name: "UNHCR"},
    {code: 971, name: "UNAIDS"},
    {code: 974, name: "UNFPA"},
    {code: 976, name: "Islamic Development Bank"},
    {code: 978, name: "OSCE"},
    {code: 979, name: "Islamic Monetary Fund"},
    {code: 980, name: "Arab Fund for Technical Assistance to African and Arab Countries"},
    {code: 981, name: "Black Sea Trade & Development Bank"},
    {code: 982, name: "GODE"},
    {code: 983, name: "Other Arab Agencies"},
    {code: 988, name: "IFAD"},
    {code: 990, name: "European Bank for Reconstruction and Development"},
    {code: 992, name: "UN AGENCIES"},
    {code: 997, name: "Global Partnership for Education"},
    {code: 1011, name: "Climate Investment Funds"},
    {code: 1012, name: "Adaptation Fund"},
    {code: 1013, name: "Council of Europe Development Bank"},
    {code: 1014, name: "Private Infrastructure Development Group"},
    {code: 1015, name: "Development Bank of Latin America"},
    {code: 1016, name: "Green Climate Fund"},
    {code: 1017, name: "Credit Guarantee and Investment Facility"},
    {code: 1018, name: "Global Energy Efficiency and Renewable Energy Fund"},
    {code: 1019, name: "IDB Invest"},
    {code: 1020, name: "Central Emergency Response Fund"},
    {code: 1023, name: "World Tourism Organisation"},
    {code: 1024, name: "Asian Infrastructure Investment Bank"},
    {code: 1025, name: "Center of Excellence in Finance"},
    {code: 1311, name: "Global Alliance for Vaccines and Immunization"},
    {code: 1312, name: "Global Fund"},
    {code: 1313, name: "Global Green Growth Institute"},
    {code: 1601, name: "Bill & Melinda Gates Foundation"},
    {code: 1602, name: "Dutch Postcode Lottery"},
    {code: 1603, name: "Swedish Postcode Lottery"},
    {code: 1604, name: "People's Postcode Lottery"},
    {code: 1605, name: "MetLife Foundation"},
    {code: 1606, name: "MasterCard Foundation"},
    {code: 1607, name: "Grameen Crédit Agricole Foundation"},
    {code: 1608, name: "IKEA Foundation"},
    {code: 1609, name: "Bernard van Leer Foundation"},
    {code: 1610, name: "MAVA Foundation"},
    {code: 1611, name: "Oak Foundation"},
    {code: 1612, name: "H&M Foundation"},
    {code: 1613, name: "C&A Foundation"},
    {code: 1614, name: "Charity Projects Ltd (Comic Relief)"},
    {code: 1615, name: "Children's Investment Fund Foundation"},
    {code: 1616, name: "Gatsby Charitable Foundation"},
    {code: 1617, name: "Conrad N. Hilton Foundation"},
    {code: 1618, name: "David & Lucile Packard Foundation"},
    {code: 1619, name: "John D. & Catherine T. MacArthur Foundation"},
    {code: 1620, name: "Carnegie Corporation of New York"},
    {code: 1621, name: "Michael & Susan Dell Foundation"},
    {code: 1622, name: "Omidyar Network Fund, Inc."},
    {code: 1623, name: "Rockefeller Foundation"},
    {code: 1624, name: "William & Flora Hewlett Foundation"},
    {code: 1625, name: "Arcus Foundation"},
    {code: 1626, name: "Gordon and Betty Moore Foundation"},
    {code: 1627, name: "Ford Foundation"},
    {code: 1628, name: "Wellcome Trust"}
  ]

  static reportVersions = [ "2.03" ]

  static budgetStatus = [
    {code: 1, name:	"Indicative"},
    {code: 2, name:	"Committed"},
  ]
}

export { Paths, Blockchain, Contract, Helpers }
