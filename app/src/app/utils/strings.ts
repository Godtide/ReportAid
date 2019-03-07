import { Paths as configPaths } from './config'

class App {

  static readonly title='ReportAid'
  static readonly tagline='Using Blockchain Technology to Add Trust to Humanitarian Aid Reporing'
  static readonly copyright='[University of Sussex](https://www.sussex.ac.uk/) © 2018'
  static readonly author='Created by [Steven Huckle](https://glowkeeper.github.io/)'

  static headingOrganisationsWriter = 'Create Organisations Data'
  static headingOrganisationsReader = 'Read Organisations Data'
  static headingActivitiesWriter = 'Create Activities Data'
  static headingActivitiesUpdater = 'Update Activities Data'
  static headingActivitiesReader = 'Read Activities Data'
}

class Paths {

  // AppBar
  static readonly home='Home'
  static readonly blockchain = 'Blockchain Data'
  static readonly about='About'
  static readonly overview='Overview'
  static readonly help='Help'
  static readonly writer='Create Records'
  static readonly reader='Read Records'

  static readonly activitiesWriter='Activities'
  static readonly activityWriter='Activity'
  static readonly activityAdditionalWriter='Activity Additional Info\''
  static readonly orgWriter='Org\''
  static readonly organisationsWriter='Organisations'
  static readonly organisationWriter='Organisation Report'
  static readonly organisationDocsWriter='Document'
  static readonly organisationBudgetsWriter="Budget"
  static readonly organisationExpenditureWriter = "Expenditure"
  static readonly organisationRecipientBudgetsWriter = "Recipient Budget"
  static readonly organisationRegionBudgetsWriter='Region Budget'
  static readonly organisationCountryBudgetsWriter='Country Budget'

  static readonly activitiesUpdater='Activities'
  static readonly activityUpdater='Activity'

  static readonly activitiesReader='Activities'
  static readonly activityReader='Activity'
  static readonly activityAdditionalReader='Activity Additional Info\''
  static readonly orgsReader='Org\'s'
  static readonly organisationsReader='Organisations'
  static readonly organisationReader='Organisation Reports'
  static readonly organisationDocsReader='Documents'
  static readonly organisationBudgetsReader='Budgets'
  static readonly organisationExpenditureReader='Expenditures'
  static readonly organisationRecipientBudgetsReader='Recipient Budgets'
  static readonly organisationRegionBudgetsReader='Region Budgets'
  static readonly organisationCountryBudgetsReader='Country Budgets'
}

class Blockchain {

  static heading = 'Blockchain Data'
}

class Home {

  static heading = 'Home'

  static info = `Use this application to record humanitarian aid data and to get humanitarian aid information.<br /><br />Read the [Overview](#${configPaths.overview}) section to learn about the origins of **ReportAid**.<br /><br />The [Help](#${configPaths.help}) section gives brief instructions as to how to use **ReportAid** - in essence, to create a humanitarian aid record, click on the [Create Record](#${configPaths.writer}) link and fill in all fields. To retrieve a humanitarian aid record, click on the [Read Record](#${configPaths.reader}) link. You must have [MetaMask](https://metamask.io/) installed.`
}

class About {

  static heading = 'About ReportAid'

  static info = '**ReportAid** version 0.0.2.<br /><br />Created by [Steven Huckle](https://glowkeeper.github.io/) at the [University of Sussex](https://www.sussex.ac.uk/). '

}

class Overview {

  static heading = 'Overview of ReportAid'

  static info = '**ReportAid** iorgRefs the result of an academic paper titled: _Humanitarian Aid - a Blockchain Application That Adds Trust to Aid Provisioning_. The article discusses how the trust mechanisms of blockchain technology might be used to promote transparanecy in humanitarian aid provisioning. The general idea is that blockchains can add trust to the reporting of humanitarian aid funding.<br /><br />For more information about **ReportAid**, please contact s dot huckle at sussex dot ac dot uk.'
}

class Help {

  static heading = 'ReportAid Help'

  static info = `**ReportAid** allows humanitarian aid organisations to record information about funding.<br /><br />Have a read of the [Overview](#${configPaths.overview}) section, which gives some background about the app\'. <br /><br />The _${App.headingOrganisationsWriter}_ menu lets you create IATI Organsations records. The _${App.headingActivitiesWriter}_ menu lets you create IATI Activities records. The _${App.headingOrganisationsReader}_ menu lets you read IATI Organsations records. The _${App.headingActivitiesReader}_ menu lets you read IATI Activities records.<br /><br />This app' relies on [MetaMask](https://metamask.io/).`
}

class IATIWriter {

  static heading = 'IATI Writer'

  static info = `The _${App.headingOrganisationsWriter}_ menu lets you create IATI Organisations records. The _${App.headingActivitiesWriter}_ menu lets you create IATI Activities records.`
}

class IATIReader {

  static heading = 'IATI Reader'

  static info = `The _${App.headingOrganisationsReader}_ menu lets you read IATI Organisations records. The _${App.headingActivitiesReader}_ menu lets you read IATI Activities records.`
}

class Transaction {

  static heading = "Transaction Details"

  static hash = 'Transaction Hash'
  static summary = 'Transaction Summary'
  static success = 'Submitted Successfully!'
  static fail = 'Submission Failed!'
}

class Activities {

  static headingActivitiesWriter = 'Write Activities'
  static headingActivitiesReader = 'Read Activities'
  static headingActivitiesUpdater = 'Update Activities'

  static activitiesReference = "Activities Reference"
  static reportKey = "Activities Reference"
  static version = "Version"
  static generated = "Generated"
  static linkedData = "Linked Data URL"

  static numActivities = 'Number of Activities Reports'
  static activitiesDetails = 'Activities Reports Details'
}

class Activity {

  static headingActivityWriter = 'Write Activity'
  static headingActivityReader = 'Read Activity'
  static headingActivityUpdater = 'Update Activity'

  static activitiesReference = "Activities Reference"
  static activityReference = "Activity Reference"

  static identifier = "Activity Identifier"
  static reportingOrgRef = 'Reporting Organisation'
  static reportingOrgType = "Reporting Organisation Type"
  static reportingOrgIsSecondary = "Is Reporting Organisation Secondary?"
  static title = "Title"
  static lastUpdated = "Last Updated"
  static language = "Language"
  static currency = "Currency"
  static budgetNotProvided = "Budget Not Provided"
  static status = "Status"
  static humanitarian = "Is Humanitarian?"
  static hierarchy = "Hierarchy"
  static description = "Description"
  static linkedData = "Linked Data URL"

  static activityDetails = 'Activity Reports Details'
}

class ActivityAdditional {

  static headingActivityAdditionalWriter = 'Write Activity Additional Info\''
  static headingActivityAdditionalReader = 'Read Activity Additional Info\''

  static activitiesReference = "Activities Reference"
  static activityReference = "Activity Reference"
  static additionalRef = "Activity Additional Info\' Reference"

  static scope = "Scope"
  static capitalSpend = "Capital Spend"
  static collaborationType = "Collaboration Type"
  static defaultFlowType = "Default Flow Type"
  static defaultFinanceType = "Default Finance Type"
  static defaultAidType = "Default Aid Type"
  static defaultTiedStatus = "Defaul Tied Status"

  static activityAdditionalDetails = "Activity Additional Info\' Details"
}

class Org {

  static headingOrgWriter = 'Create Organisation'
  static headingOrgReader = 'Organisations'

  static orgIdentifier = "Org\' Reference"

  static orgName = 'Name'
  static code = 'Code'
  static identifier = 'Identifier'

  static numOrgs = 'Number of Organisations'
  static orgDetails = 'Organisation Details'
}

class Organisations {

  static headingOrganisationsWriter = 'Write Organisations'
  static headingOrganisationsReader = 'Read Organisations'

  static organisationsReference = "Organisations Reference"
  static reportKey = "Organisations Reference"
  static version = "Version"
  static generated = "Generated"

  static numOrganisations = 'Number of Organisations Reports'
  static organisationsDetails = 'Organisations Reports Details'
}

class Organisation {

  static headingOrganisationWriter = 'Create Organisation'
  static headingOrganisationReader = 'Organisation'

  static organisationsReference = "Organisations Reference"
  static organisationReference = "Organisation Reference"
  static orgRef = "Issuing Organisation"
  static reportingOrgRef = 'Reporting Organisation'
  static reportingOrgType = "Reporting Organisation Type"
  static reportingOrgIsSecondary = "Is Reporting Organisation Secondary?"
  static language = "Language"
  static currency = "Currency"
  static lastUpdated = "Last Updated"

  static numOrganisations = "Number of Organisations Reports"
  static numOrganisation = 'Number of Organisation Reports'
  static organisationDetails = 'Organisation Details'
}

class OrganisationDoc {

  static headingOrganisationDocWriter = 'Create Document'
  static headingOrganisationDocReader = 'Documents'

  static organisationsReference = "Organisations Reference"
  static organisationReference = "Organisation Reference"
  static docReference = "Document Reference"
  static documentTitle = "Document Title"
  static documentFormat = "Format"
  static documentURL = "URL"
  static documentCategory = "Category"
  static documentCountryRef = "Country Reference"
  static documentDesc = "Description"
  static documentLang = "Language"
  static documentDay = "Day"
  static documentMonth = "Month"
  static documentYear = "Year"
  static documentDate = "Date"

  static numDocs = "Number of Documents"
  static organisationDocDetails = 'Document Details'
}

class OrganisationBudget {

  static headingOrganisationBudgetWriter = 'Create Report Budget'
  static headingOrganisationBudgetReader = 'Report Budgets'

  static reportingOrgRef = 'Reporting Organisation'
  static budgetReference = "Budget Reference"
  static organisationsReference = "Organisations Reference"
  static organisationReference = "Organisation Reference"
  static budgetLine = "Budget Line"
  static value = "Value"
  static status = "Status"
  static budgetStart = "Start Date"
  static budgetStartDay = "Start Day"
  static budgetStartMonth = "Start Month"
  static budgetStartYear = "Start Year"
  static budgetEnd = "End Date"
  static budgetEndDay = "End Day"
  static budgetEndMonth = "End Month"
  static budgetEndYear = "End Year"

  static numBudgets = "Number of Budgets"
  static organisationBudgetDetails = 'Budget Details'
}

class OrganisationExpenditure {

  static headingOrganisationExpenditureWriter = 'Create Report Expenditure'
  static headingOrganisationExpenditureReader = 'Report Expenditure'

  static reportingOrgRef = 'Reporting Organisation'
  static expenditureReference = "Expenditure Reference"
  static organisationsReference = "Organisations Reference"
  static organisationReference = "Organisation Reference"
  static expenditureLine = "Expenditure Line"
  static value = "Value"
  static status = "Status"
  static expenditureStart = "Start Date"
  static expenditureStartDay = "Start Day"
  static expenditureStartMonth = "Start Month"
  static expenditureStartYear = "Start Year"
  static expenditureEnd = "End Date"
  static expenditureEndDay = "End Day"
  static expenditureEndMonth = "End Month"
  static expenditureEndYear = "End Year"

  static numExpenditure = "Number of Expenditure"
  static organisationExpenditureDetails = 'Expenditure Details'
}

class OrganisationRecipientBudget {

  static headingOrganisationRecipientBudgetWriter = 'Create Report Recipient Budget'
  static headingOrganisationRecipientBudgetReader = 'Report Recipient Budgets'

  static reportingOrgRef = 'Reporting Organisation'
  static budgetReference = "Budget Reference"
  static organisationsReference = "Organisations Reference"
  static organisationReference = "Organisation Reference"
  static orgReference = "Recipient Reference"
  static budgetLine = "Budget Line"
  static value = "Value"
  static status = "Status"
  static budgetStart = "Start Date"
  static budgetStartDay = "Start Day"
  static budgetStartMonth = "Start Month"
  static budgetStartYear = "Start Year"
  static budgetEnd = "End Date"
  static budgetEndDay = "End Day"
  static budgetEndMonth = "End Month"
  static budgetEndYear = "End Year"

  static numBudgets = "Number of Recipient Budgets"
  static organisationRecipientBudgetDetails = 'Recipient Budget Details'
}

class OrganisationRegionBudget {

  static headingOrganisationRegionBudgetWriter = 'Create Report Region Budget'
  static headingOrganisationRegionBudgetReader = 'Report Region Budgets'

  static reportingOrgRef = 'Reporting Organisation'
  static budgetReference = "Budget Reference"
  static organisationsReference = "Organisations Reference"
  static organisationReference = "Organisation Reference"
  static regionReference = "Region Reference"
  static budgetLine = "Budget Line"
  static value = "Value"
  static status = "Status"
  static budgetStart = "Start Date"
  static budgetStartDay = "Start Day"
  static budgetStartMonth = "Start Month"
  static budgetStartYear = "Start Year"
  static budgetEnd = "End Date"
  static budgetEndDay = "End Day"
  static budgetEndMonth = "End Month"
  static budgetEndYear = "End Year"

  static numBudgets = "Number of Region Budgets"
  static organisationRegionBudgetDetails = 'Region Budget Details'
}

class OrganisationCountryBudget {

  static headingOrganisationCountryBudgetWriter = 'Create Report Country Budget'
  static headingOrganisationCountryBudgetReader = 'Report Country Budgets'

  static reportingOrgRef = 'Reporting Organisation'
  static budgetReference = "Budget Reference"
  static organisationsReference = "Organisations Reference"
  static organisationReference = "Organisation Reference"
  static countryReference = "Country Reference"
  static budgetLine = "Budget Line"
  static value = "Value"
  static status = "Status"
  static budgetStart = "Start Date"
  static budgetStartDay = "Start Day"
  static budgetStartMonth = "Start Month"
  static budgetStartYear = "Start Year"
  static budgetEnd = "End Date"
  static budgetEndDay = "End Day"
  static budgetEndMonth = "End Month"
  static budgetEndYear = "End Year"

  static numBudgets = "Number of Country Budgets"
  static organisationCountryBudgetDetails = 'Country Budget Details'
}


export { App,
         Paths,
         Blockchain,
         Home,
         About,
         Overview,
         Help,
         IATIWriter,
         IATIReader,
         Transaction,
         Activities,
         Activity,
         ActivityAdditional,
         Org,
         Organisations,
         Organisation,
         OrganisationDoc,
         OrganisationBudget,
         OrganisationExpenditure,
         OrganisationRecipientBudget,
         OrganisationRegionBudget,
         OrganisationCountryBudget
       }
