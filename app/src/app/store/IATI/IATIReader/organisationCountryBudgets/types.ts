import { PayloadProps, DictData } from '../../../types'
import { IATIOrgCountryBudgetProps } from '../../types'

export interface CountryBudgetData {
  [key: string]: IATIOrgCountryBudgetProps
}

export interface CountryBudgets {
  num: number
  data: CountryBudgetData
}

export interface OrgCountryBudgetsData extends DictData {
  [key: string]: CountryBudgets
}

export interface OrgCountryBudgetsReaderProps extends PayloadProps {
  num: number
  data: OrgCountryBudgetsData
}

export const enum OrgCountryBudgetsReaderActionTypes {
  NUM_SUCCESS = '@@OrgCountryBudgetsReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrgCountryBudgetsReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrgCountryBudgetsReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrgCountryBudgetsReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrgCountryBudgetsReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrgCountryBudgetsReaderAction/GETREFERENCE_FAILURE',
  NUMBUDGET_SUCCESS = '@@OrgCountryBudgetsReaderAction/GETNUMBUDGET_SUCCESS',
  NUMBUDGET_FAILURE = '@@OrgCountryBudgetsReaderAction/GETNUMBUDGET_FAILURE',
  BUDGETREF_SUCCESS = '@@OrgCountryBudgetsReaderAction/GETBUDGETREF_SUCCESS',
  BUDGETREF_FAILURE = '@@OrgCountryBudgetsReaderAction/GETBUDGETREF_FAILURE',
  BUDGET_SUCCESS = '@@OrgCountryBudgetsReaderAction/GETBUDGET_SUCCESS',
  BUDGET_FAILURE = '@@OrgCountryBudgetsReaderAction/GETBUDGET_FAILURE'
}
