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
