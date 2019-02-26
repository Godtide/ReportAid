import { PayloadProps } from '../../types'

export interface IATIExpenditureData {
  expenditureKey: string
  expenditureLine: string,
  value: number,
  status: number,
  start: string,
  end: string
}

export interface IATIRecipientOrgBudgetData {
  budgetKey: string
  budgetLine: string,
  recipientOrgRef: string,
  value: number,
  status: number,
  start: string,
  end: string
}

export interface IATIRegionBudgetData {
  budgetKey: string
  budgetLine: string,
  regionRef: string,
  value: number,
  status: number,
  start: string,
  end: string
}

export interface IATICountryBudgetData {
  budgetKey: string
  budgetLine: string,
  countryRef: string,
  value: number,
  status: number,
  start: string,
  end: string
}

export interface IATIBudgetData {
  budgetKey: string
  budgetLine: string,
  value: number,
  status: number,
  start: string,
  end: string
}

export interface IATIBudgetReport {
  organisationsRef: string
  organisationRef: string
  data: Array<IATIBudgetData | IATIExpenditureData | IATIRecipientOrgBudgetData | IATIRegionBudgetData | IATICountryBudgetData >
}

export interface IATIBudgetReportProps extends PayloadProps {
  data: IATIBudgetReport
}
