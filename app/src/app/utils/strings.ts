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
  static readonly orgReportWriter='Report'
  static readonly orgDocsWriter='Document'
  static readonly orgBudgetsWriter="Budget"
  static readonly orgExpenditureWriter = "Expenditure"
  static readonly orgRecipientBudgetsWriter = "Recipient Budget"
  static readonly orgRegionBudgetsWriter='Region Budget'
  static readonly orgCountryBudgetsWriter='Country Budget'
  static readonly orgReader='Organisations'
  static readonly orgReportReader='Reports'
  static readonly orgDocsReader='Documents'
  static readonly orgBudgetsReader='Budgets'
  static readonly orgExpenditureReader='Expenditures'
  static readonly orgRecipientBudgetsReader='Recipient Budgets'
  static readonly orgRegionBudgetsReader='Region Budgets'
  static readonly orgCountryBudgetsReader='Country Budgets'
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

  static headingOrgWriter = 'Create Organisation Records'
  static headingOrgReader = 'Read Organisation Records'

  static orgDetails = 'Organisation Details'
  static orgIdentifier = "Organisation Reference"

  static orgName = 'Name'
  static code = 'Code'
  static identifier = 'Identifier'

  static numOrgs = 'Number of Organisations'
}

class Organisation {

  static headingOrgWriter = 'Create Organisation'
  static headingOrgOrganisationsWriterReader = 'Organisation'

  static numOrgs = "Number of reports"
  static version = "Version"
  static orgIdentifier = "Organisation"
  static reportKey = "Key"
  static reportingOrgRef = 'Reporting Organisation'
  static reportingOrgType = "Reporting Organisation Type"
  static reportingOrgIsSecondary = "Is Reporting Organisation Secondary?"
  static language = "Language"
  static currency = "Currency"
  static lastUpdated = "Last Updated"

  static numOrganisations = 'Number of Organisations'
  static orgDetails = 'Details'
}

class OrganisationDoc {

  static headingOrgDocWriter = 'Create Document'
  static headingOrgDocReader = 'Documents'

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

  static numDocs = "Number of Documents"
  static reportDocDetails = 'Document Details'
}

class OrganisationBudget {

  static headingOrgBudgetWriter = 'Create Report Budget'
  static headingOrgBudgetReader = 'Report Budgets'

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
  static reportBudgetDetails = 'Budget Details'
}

class OrganisationExpenditure {

  static headingOrgExpenditureWriter = 'Create Report Expenditure'
  static headingOrgExpenditureReader = 'Report Expenditure'

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
  static reportExpenditureDetails = 'Expenditure Details'
}

class OrganisationRecipientBudget {

  static headingOrgRecipientBudgetWriter = 'Create Report Recipient Budget'
  static headingOrgRecipientBudgetReader = 'Report Recipient Budgets'

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
  static reportBudgetDetails = 'Recipient Budget Details'
}

class OrganisationRegionBudget {

  static headingOrgRegionBudgetWriter = 'Create Report Region Budget'
  static headingOrgRegionBudgetReader = 'Report Region Budgets'

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
  static reportBudgetDetails = 'Region Budget Details'
}

class OrganisationCountryBudget {

  static headingOrgCountryBudgetWriter = 'Create Report Country Budget'
  static headingOrgCountryBudgetReader = 'Report Country Budgets'

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
  static reportBudgetDetails = 'Country Budget Details'
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
         Organisation,
         OrganisationDoc,
         OrganisationBudget,
         OrganisationExpenditure,
         OrganisationRecipientBudget,
         OrganisationRegionBudget,
         OrganisationCountryBudget
       }
