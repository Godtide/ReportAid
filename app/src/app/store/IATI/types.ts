export interface IATIOrgProps {
  orgRef: string
  name: string
  identifier: string
}

export interface ReportingOrgProps {
  orgRef: string
  orgType: number
  isSecondary: boolean
}

export interface IATIOrgReportProps {
  version: string
  orgRef: string
  reportRef: string
  reportingOrg: ReportingOrgProps
  lang: string
  currency: string
  lastUpdatedTime: string
}

export interface IATIOrgReportDocProps {
  reportRef: string
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
  reportRef: string
  budgetRef: string
  budgetLine: string
  finance: FinanceProps
}

export interface IATIOrgReportExpenditureProps {
  reportRef: string
  expenditureRef: string
  expenditureLine: string
  finance: FinanceProps
}

export interface IATIOrgReportRecipientBudgetProps {
  reportRef: string
  budgetRef: string
  orgRef: string
  budgetLine: string
  finance: FinanceProps
}

export interface IATIOrgReportRegionBudgetProps {
  reportRef: string
  budgetRef: string
  regionRef: number
  budgetLine: string
  finance: FinanceProps
}

export interface IATIOrgReportCountryBudgetProps {
  reportRef: string
  budgetRef: string
  countryRef: string
  budgetLine: string
  finance: FinanceProps
}

export interface OrganisationProps {
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
  reportRef: string
  docRef: string
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
  reportRef: string
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
  reportRef: string
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
  reportRef: string
  orgRef: string
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
  reportRef: string
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
  reportRef: string
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
