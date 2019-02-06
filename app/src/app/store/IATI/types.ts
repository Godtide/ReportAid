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

export interface IATIOrganisationsProps {
  version: string
  generatedTime: string
}

export interface IATIOrganisationProps {
  reportingOrg: ReportingOrgProps
  lang: string
  currency: string
  lastUpdatedTime: string
}

export interface IATIOrganisationDocProps {
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

export interface IATIOrganisationBudgetProps {
  budgetRef: string
  budgetLine: string
  finance: FinanceProps
}

export interface IATIOrganisationExpenditureProps {
  expenditureRef: string
  expenditureLine: string
  finance: FinanceProps
}

export interface IATIOrganisationRecipientBudgetProps {
  recipientOrgRef: string
  budgetLine: string
  finance: FinanceProps
}

export interface IATIOrganisationRegionBudgetProps {
  regionRef: number
  budgetLine: string
  finance: FinanceProps
}

export interface IATIOrganisationCountryBudgetProps {
  countryRef: string
  budgetLine: string
  finance: FinanceProps
}

export interface OrgProps {
  name: string
  code: string
  identifier: string
}

export interface OrganisationsProps {
  organisationsRef: string
  version: string
}

export interface OrganisationProps {
  organisationRef: string
  reportingOrgRef: string
  reportingOrgType: number
  reportingOrgIsSecondary: boolean
  lang: string
  currency: string
}

export interface OrganisationDocProps {
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

export interface OrganisationBudgetProps {
  budgetRef: string
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

export interface OrganisationExpenditureProps {
  expenditureRef: string
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

export interface OrganisationRecipientBudgetProps {
  budgetRef: string
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

export interface OrganisationRegionBudgetProps {
  budgetRef: string
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

export interface OrganisationCountryBudgetProps {
  budgetRef: string
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
