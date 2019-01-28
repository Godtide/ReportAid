import { PayloadProps, DictData } from '../../../types'
import { IATIOrgReportCountryBudgetProps } from '../../types'

export interface CountryBudgetData {
  [key: string]: IATIOrgReportCountryBudgetProps
}

export interface CountryBudgets {
  num: number
  data: CountryBudgetData
}

export interface OrgReportCountryBudgetsData extends DictData {
  [key: string]: CountryBudgets
}

export interface OrgReportCountryBudgetsReaderProps extends PayloadProps {
  num: number
  data: OrgReportCountryBudgetsData
}

export const enum OrgReportCountryBudgetsReaderActionTypes {
  NUM_SUCCESS = '@@OrgReportCountryBudgetsReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrgReportCountryBudgetsReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrgReportCountryBudgetsReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrgReportCountryBudgetsReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrgReportCountryBudgetsReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrgReportCountryBudgetsReaderAction/GETREFERENCE_FAILURE',
  NUMBUDGET_SUCCESS = '@@OrgReportCountryBudgetsReaderAction/GETNUMBUDGET_SUCCESS',
  NUMBUDGET_FAILURE = '@@OrgReportCountryBudgetsReaderAction/GETNUMBUDGET_FAILURE',
  BUDGETREF_SUCCESS = '@@OrgReportCountryBudgetsReaderAction/GETBUDGETREF_SUCCESS',
  BUDGETREF_FAILURE = '@@OrgReportCountryBudgetsReaderAction/GETBUDGETREF_FAILURE',
  BUDGET_SUCCESS = '@@OrgReportCountryBudgetsReaderAction/GETBUDGET_SUCCESS',
  BUDGET_FAILURE = '@@OrgReportCountryBudgetsReaderAction/GETBUDGET_FAILURE'
}
