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

export const enum IATIReportActionTypes {
  ORGS_INIT = '@@IATIReportActionTypes/READORGS_INIT',
  ORGS_SUCCESS = '@@IATIReportActionTypes/READORGS_SUCCESS',
  ORGS_FAILURE = '@@IATIReportActionTypes/READORGS_FAILURE',
  ORGANISATIONS_INIT = '@@IATIReportActionTypes/READORGANISATIONS_INIT',
  ORGANISATIONS_SUCCESS = '@@IATIReportActionTypes/READORGANISATIONS_SUCCESS',
  ORGANISATIONS_FAILURE = '@@IATIReportActionTypes/READORGANISATIONS_FAILURE',
  ORGANISATION_INIT = '@@IATIReportActionTypes/READORGANISATION_INIT',
  ORGANISATION_SUCCESS = '@@IATIReportActionTypes/READORGANISATION_SUCCESS',
  ORGANISATION_FAILURE = '@@IATIReportActionTypes/READORGANISATION_FAILURE',
  BUDGET_INIT = '@@IATIReportActionTypes/READBUDGET_INIT',
  BUDGET_SUCCESS = '@@IATIReportActionTypes/READBUDGET_SUCCESS',
  BUDGET_FAILURE = '@@IATIReportActionTypes/READBUDGET_FAILURE',
  RECIPIENTORGBUDGET_INIT = '@@IATIReportActionTypes/READRECIPIENTORGBUDGET_INIT',
  RECIPIENTORGBUDGET_SUCCESS = '@@IATIReportActionTypes/READRECIPIENTORGBUDGET_SUCCESS',
  RECIPIENTORGBUDGET_FAILURE = '@@IATIReportActionTypes/READRECIPIENTORGBUDGET_FAILURE',
  RECIPIENTREGIONBUDGET_INIT = '@@IATIReportActionTypes/READRECIPIENTREGIONBUDGET_INIT',
  RECIPIENTREGIONBUDGET_SUCCESS = '@@IATIReportActionTypes/READRECIPIENTREGIONBUDGET_SUCCESS',
  RECIPIENTREGIONBUDGET_FAILURE = '@@IATIReportActionTypes/READRECIPIENTREGIONBUDGET_FAILURE',
  RECIPIENTCOUNTRYBUDGET_INIT = '@@IATIReportActionTypes/READRECIPIENTCOUNTRYBUDGET_INIT',
  RECIPIENTCOUNTRYBUDGET_SUCCESS = '@@IATIReportActionTypes/READRECIPIENTCOUNTRYBUDGET_SUCCESS',
  RECIPIENTCOUNTRYBUDGET_FAILURE = '@@IATIReportActionTypes/READRECIPIENTCOUNTRYBUDGET_FAILURE',
  TOTALEXPENDITURE_INIT = '@@IATIReportActionTypes/READTOTALEXPENDITURE_INIT',
  TOTALEXPENDITURE_SUCCESS = '@@IATIReportActionTypes/READTOTALEXPENDITURE_SUCCESS',
  TOTALEXPENDITURE_FAILURE = '@@IATIReportActionTypes/READTOTALEXPENDITURE_FAILURE',
  DOCUMENT_INIT = '@@IATIReportActionTypes/READDOCUMENT_INIT',
  DOCUMENT_SUCCESS = '@@IATIReportActionTypes/READDOCUMENT_SUCCESS',
  DOCUMENT_FAILURE = '@@IATIReportActionTypes/READDOCUMENT_FAILURE'
}
