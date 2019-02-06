import { Paths as configPaths } from './config'

class App {

  static readonly title='ReportAid'
  static readonly tagline='Using Blockchain Technology to Add Trust to Humanitarian Aid Reporing'
  static readonly copyright='[University of Sussex](https://www.sussex.ac.uk/) © 2018'
  static readonly author='Created by [Steven Huckle](https://glowkeeper.github.io/)'

}

class Paths {

  // AppBar
  static readonly home='Home'
  static readonly blockchain = 'Blockchain Info'
  static readonly about='About'
  static readonly overview='Overview'
  static readonly help='Help'
  static readonly writer='Create Records'
  static readonly reader='Read Records'

  static readonly orgWriter='Organisation'
  static readonly organisationsWriter='Report'
  static readonly organisationWriter='Organisation Report'
  static readonly organisationDocsWriter='Document'
  static readonly organisationBudgetsWriter="Budget"
  static readonly organisationExpenditureWriter = "Expenditure"
  static readonly organisationRecipientBudgetsWriter = "Recipient Budget"
  static readonly organisationRegionBudgetsWriter='Region Budget'
  static readonly organisationCountryBudgetsWriter='Country Budget'
  static readonly orgReader='Organisation'
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

  static info = '**ReportAid** is the result of an academic paper titled: _Humanitarian Aid - a Blockchain Application That Adds Trust to Aid Provisioning_. The article discusses how the trust mechanisms of blockchain technology might be used to promote transparanecy in humanitarian aid provisioning. The general idea is that blockchains can add trust to the reporting of humanitarian aid funding.<br /><br />For more information about **ReportAid**, please contact s dot huckle at sussex dot ac dot uk.'
}

class Help {

  static heading = 'ReportAid Help'

  static info = `**ReportAid** allows humanitarian aid organisations to record information about funding.<br /><br />Have a read of the [Overview](#${configPaths.overview}) section, which gives some background about the app\'. <br /><br />To store a humanitarian aid record, click on the [Create Records](#${configPaths.writer}) link. To retrieve information, click on the [Read Records](#${configPaths.reader}) link. This app' relies on [MetaMask](https://metamask.io/).`
}

class IATIWriter {

  static heading = 'IATI Writer'

  static info = `The [Create Organisation](#${configPaths.orgWriter}) link, lets you create a top-level record for an IATI reporting organsation.`
}

class IATIReader {

  static heading = 'IATI Reader'

  static info = `The [Read Organisations](#${configPaths.orgReader}) link, lets you read a top-level record for an IATI reporting organsation.`
}

class Transaction {

  static heading = "Transaction Details"

  static key = 'Transaction Receipt'
  static summary = 'Transaction Summary'
  static success = 'Submitted Successfully!'
  static fail = 'Submission Failed!'
}

class Org {

  static headingOrgWriter = 'Create Organisation'
  static headingOrgReader = 'Read Organisation'

  static orgDetails = 'Organisation Details'
  static orgIdentifier = "Organisation Reference"

  static orgName = 'Name'
  static code = 'Code'
  static identifier = 'Identifier'

  static numOrgs = 'Number of Organisations'
}

class Organisations {

  static headingOrganisationsWriter = 'Write Organisations'
  static headingOrganisationsReader = 'Read Organisations'

  static version = "Version"
  static reportKey = "Key"
  static generated = "Generated"

  static numOrganisations = 'Number of Organisations'
  static organisationsDetails = 'Details'
}

class Organisation {

  static headingOrganisationWriter = 'Create Organisation'
  static headingOrganisationReader = 'Organisation'

  static numOrganisation = "Number of reports"
  static version = "Version"
  static orgIdentifier = "Organisation"
  static reportKey = "Key"
  static reportingOrgRef = 'Reporting Organisation'
  static reportingOrgType = "Reporting Organisation Type"
  static reportingOrgIsSecondary = "Is Reporting Organisation Secondary?"
  static language = "Language"
  static currency = "Currency"
  static lastUpdated = "Last Updated"

  static numOrganisation = 'Number of Organisation Records'
  static organisationDetails = 'Details'
}

class OrganisationDoc {

  static headingOrganisationDocWriter = 'Create Document'
  static headingOrganisationDocReader = 'Documents'

  static num = "Number of Organisations Reporting"

  static reportReference = "Report Reference"
  static reportingOrgRef = 'Reporting Organisation'
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

  static numOrganisationDocs = "Number of Documents"
  static organisationDocDetails = 'Document Details'
}

class OrganisationBudget {

  static headingOrganisationBudgetWriter = 'Create Report Budget'
  static headingOrganisationBudgetReader = 'Report Budgets'

  static num = "Number of Organisations Reporting"

  static reportingOrgRef = 'Reporting Organisation'
  static budgetReference = "Budget Reference"
  static reportReference = "Report Reference"
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

  static num = "Number of Organisations Reporting"

  static reportingOrgRef = 'Reporting Organisation'
  static expenditureReference = "Expenditure Reference"
  static reportReference = "Report Reference"
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

  static num = "Number of Organisations Reporting"

  static reportingOrgRef = 'Reporting Organisation'
  static budgetReference = "Budget Reference"
  static reportReference = "Report Reference"
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
  static organisationBudgetDetails = 'Recipient Budget Details'
}

class OrganisationRegionBudget {

  static headingOrganisationRegionBudgetWriter = 'Create Report Region Budget'
  static headingOrganisationRegionBudgetReader = 'Report Region Budgets'

  static num = "Number of Organisations Reporting"

  static reportingOrgRef = 'Reporting Organisation'
  static budgetReference = "Budget Reference"
  static reportReference = "Report Reference"
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
  static organisationBudgetDetails = 'Region Budget Details'
}

class OrganisationCountryBudget {

  static headingOrganisationCountryBudgetWriter = 'Create Report Country Budget'
  static headingOrganisationCountryBudgetReader = 'Report Country Budgets'

  static num = "Number of Organisations Reporting"

  static reportingOrgRef = 'Reporting Organisation'
  static budgetReference = "Budget Reference"
  static reportReference = "Report Reference"
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
  static organisationBudgetDetails = 'Country Budget Details'
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
