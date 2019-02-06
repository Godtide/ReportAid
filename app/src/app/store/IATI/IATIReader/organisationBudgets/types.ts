import { PayloadProps, DictData } from '../../../types'
import { IATIOrgBudgetProps } from '../../types'

export interface BudgetData {
  [key: string]: IATIOrgBudgetProps
}

export interface Budgets {
  num: number
  data: BudgetData
}

export interface OrgBudgetsData extends DictData {
  [key: string]: Budgets
}

export interface OrgBudgetsReaderProps extends PayloadProps {
  num: number
  data: OrgBudgetsData
}

export const enum OrgBudgetsReaderActionTypes {
  NUM_SUCCESS = '@@OrgBudgetsReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrgBudgetsReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrgBudgetsReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrgBudgetsReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrgBudgetsReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrgBudgetsReaderAction/GETREFERENCE_FAILURE',
  NUMBUDGET_SUCCESS = '@@OrgBudgetsReaderAction/GETNUMBUDGET_SUCCESS',
  NUMBUDGET_FAILURE = '@@OrgBudgetsReaderAction/GETNUMBUDGET_FAILURE',
  BUDGETREF_SUCCESS = '@@OrgBudgetsReaderAction/GETBUDGETREF_SUCCESS',
  BUDGETREF_FAILURE = '@@OrgBudgetsReaderAction/GETBUDGETREF_FAILURE',
  BUDGET_SUCCESS = '@@OrgBudgetsReaderAction/GETBUDGET_SUCCESS',
  BUDGET_FAILURE = '@@OrgBudgetsReaderAction/GETBUDGET_FAILURE'
}
