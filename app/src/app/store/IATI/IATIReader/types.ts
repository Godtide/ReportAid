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
  REPORT_INIT =  '@@IATIReportActionTypes/REPORT_INIT',
  ORGS_SUCCESS = '@@IATIReportActionTypes/ORGS_SUCCESS',
  ORGSPICKER_SUCCESS = '@@IATIReportActionTypes/ORGSPICKER_SUCCESS',
  ORGS_FAILURE = '@@IATIReportActionTypes/ORGS_FAILURE',
  ORGANISATIONS_SUCCESS = '@@IATIReportActionTypes/ORGANISATIONS_SUCCESS',
  ORGANISATIONSPICKER_SUCCESS = '@@IATIReportActionTypes/ORGANISATIONSPICKER_SUCCESS',
  ORGANISATIONS_FAILURE = '@@IATIReportActionTypes/ORGANISATIONS_FAILURE',
  ORGANISATION_SUCCESS = '@@IATIReportActionTypes/ORGANISATION_SUCCESS',
  ORGANISATIONPICKER_SUCCESS = '@@IATIReportActionTypes/ORGANISATIONPICKER_SUCCESS',
  ORGANISATION_FAILURE = '@@IATIReportActionTypes/ORGANISATION_FAILURE',
  BUDGET_SUCCESS = '@@IATIReportActionTypes/BUDGET_SUCCESS',
  BUDGET_FAILURE = '@@IATIReportActionTypes/BUDGET_FAILURE',
  RECIPIENTORGBUDGET_SUCCESS = '@@IATIReportActionTypes/RECIPIENTORGBUDGET_SUCCESS',
  RECIPIENTORGBUDGET_FAILURE = '@@IATIReportActionTypes/RECIPIENTORGBUDGET_FAILURE',
  RECIPIENTREGIONBUDGET_SUCCESS = '@@IATIReportActionTypes/RECIPIENTREGIONBUDGET_SUCCESS',
  RECIPIENTREGIONBUDGET_FAILURE = '@@IATIReportActionTypes/RECIPIENTREGIONBUDGET_FAILURE',
  RECIPIENTCOUNTRYBUDGET_SUCCESS = '@@IATIReportActionTypes/RECIPIENTCOUNTRYBUDGET_SUCCESS',
  RECIPIENTCOUNTRYBUDGET_FAILURE = '@@IATIReportActionTypes/RECIPIENTCOUNTRYBUDGET_FAILURE',
  TOTALEXPENDITURE_SUCCESS = '@@IATIReportActionTypes/TOTALEXPENDITURE_SUCCESS',
  TOTALEXPENDITURE_FAILURE = '@@IATIReportActionTypes/TOTALEXPENDITURE_FAILURE',
  DOCUMENT_SUCCESS = '@@IATIReportActionTypes/DOCUMENT_SUCCESS',
  DOCUMENT_FAILURE = '@@IATIReportActionTypes/DOCUMENT_FAILURE',
  ACTIVITIES_SUCCESS = '@@IATIReportActionTypes/ACTIVITIES_SUCCESS',
  ACTIVITIES_FAILURE = '@@IATIReportActionTypes/ACTIVITIES_FAILURE'
}
