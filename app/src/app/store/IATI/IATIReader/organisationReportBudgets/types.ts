import { PayloadProps, DictData } from '../../../types'
import { IATIOrgReportBudgetProps } from '../../types'

export interface BudgetData {
  [key: string]: IATIOrgReportBudgetProps
}

export interface Budgets {
  num: number
  data: BudgetData
}

export interface OrgReportBudgetsData extends DictData {
  [key: string]: Budgets
}

export interface OrgReportBudgetsReaderProps extends PayloadProps {
  num: number
  data: OrgReportBudgetsData
}

export const enum OrgReportBudgetsReaderActionTypes {
  NUM_SUCCESS = '@@OrgReportBudgetsReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrgReportBudgetsReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrgReportBudgetsReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrgReportBudgetsReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrgReportBudgetsReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrgReportBudgetsReaderAction/GETREFERENCE_FAILURE',
  NUMBUDGET_SUCCESS = '@@OrgReportBudgetsReaderAction/GETNUMBUDGET_SUCCESS',
  NUMBUDGET_FAILURE = '@@OrgReportBudgetsReaderAction/GETNUMBUDGET_FAILURE',
  BUDGETREF_SUCCESS = '@@OrgReportBudgetsReaderAction/GETBUDGETREF_SUCCESS',
  BUDGETREF_FAILURE = '@@OrgReportBudgetsReaderAction/GETBUDGETREF_FAILURE',
  BUDGET_SUCCESS = '@@OrgReportBudgetsReaderAction/GETBUDGET_SUCCESS',
  BUDGET_FAILURE = '@@OrgReportBudgetsReaderAction/GETBUDGET_FAILURE'
}
