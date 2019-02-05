export interface IATIOrgProps {
  orgRef: string
  name: string
  identifier: string
}

export interface ReportProps {
  orgRef: string
  reportRef: string
}

export interface ReportingOrgProps {
  orgRef: string
  orgType: number
  isSecondary: boolean
}

export interface IATIOrgReportProps {
  version: string
  report: ReportProps
  reportingOrg: ReportingOrgProps
  lang: string
  currency: string
  lastUpdatedTime: string
}

export interface IATIOrgReportDocProps {
  report: ReportProps
  docRef: string
  title: string
  format: string
  url: string
  category: string
  countryRef: string
  desc: string
  lang: string
  date: string
}

export interface FinanceProps {
  value: number
  status: number
  start: string
  end: string
}

export interface IATIOrgReportBudgetProps {
  report: ReportProps
  budgetRef: string
  budgetLine: string
  finance: FinanceProps
}

export interface IATIOrgReportExpenditureProps {
  report: ReportProps
  expenditureRef: string
  expenditureLine: string
  finance: FinanceProps
}

export interface IATIOrgReportRecipientBudgetProps {
  report: ReportProps
  budgetRef: string
  orgRef: string
  budgetLine: string
  finance: FinanceProps
}

export interface IATIOrgReportRegionBudgetProps {
  report: ReportProps
  budgetRef: string
  regionRef: number
  budgetLine: string
  finance: FinanceProps
}

export interface IATIOrgReportCountryBudgetProps {
  report: ReportProps
  budgetRef: string
  countryRef: string
  budgetLine: string
  finance: FinanceProps
}

export interface OrgProps {
  name: string
  code: string
  identifier: string
}



export interface OrgReportProps {
  reportingOrgRef: string
  reportingOrgType: number
  reportingOrgIsSecondary: boolean
  orgRef: string
  version: string
  lang: string
  currency: string
}

export interface OrgReportDocProps {
  report: ReportProps
  title: string
  format: string
  url: string
  category: string
  countryRef: string
  desc: string
  lang: string
  day: number
  month: number
  year: number
}

export interface OrgReportBudgetProps {
  report: ReportProps
  budgetLine: string
  value: number
  status: number
  startDay: number
  startMonth: number
  startYear: number
  endDay: number
  endMonth: number
  endYear: number
}

export interface OrgReportExpenditureProps {
  report: ReportProps
  expenditureLine: string
  value: number
  status: number
  startDay: number
  startMonth: number
  startYear: number
  endDay: number
  endMonth: number
  endYear: number
}

export interface OrgReportRecipientBudgetProps {
  report: ReportProps
  recipientOrgRef: string
  budgetLine: string
  value: number
  status: number
  startDay: number
  startMonth: number
  startYear: number
  endDay: number
  endMonth: number
  endYear: number
}

export interface OrgReportRegionBudgetProps {
  report: ReportProps
  regionRef: number
  budgetLine: string
  value: number
  status: number
  startDay: number
  startMonth: number
  startYear: number
  endDay: number
  endMonth: number
  endYear: number
}

export interface OrgReportCountryBudgetProps {
  report: ReportProps
  countryRef: string
  budgetLine: string
  value: number
  status: number
  startDay: number
  startMonth: number
  startYear: number
  endDay: number
  endMonth: number
  endYear: number
}
