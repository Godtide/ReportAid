import { PayloadProps, DictData } from '../../../types'
import { IATIOrgReportRecipientBudgetProps } from '../../types'

export interface RecipientBudgetData {
  [key: string]: IATIOrgReportRecipientBudgetProps
}

export interface RecipientBudgets {
  num: number
  data: RecipientBudgetData
}

export interface OrgReportRecipientBudgetsData extends DictData {
  [key: string]: RecipientBudgets
}

export interface OrgReportRecipientBudgetsReaderProps extends PayloadProps {
  num: number
  data: OrgReportRecipientBudgetsData
}

export const enum OrgReportRecipientBudgetsReaderActionTypes {
  NUM_SUCCESS = '@@OrgReportRecipientBudgetsReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrgReportRecipientBudgetsReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrgReportRecipientBudgetsReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrgReportRecipientBudgetsReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrgReportRecipientBudgetsReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrgReportRecipientBudgetsReaderAction/GETREFERENCE_FAILURE',
  NUMBUDGET_SUCCESS = '@@OrgReportRecipientBudgetsReaderAction/GETNUMBUDGET_SUCCESS',
  NUMBUDGET_FAILURE = '@@OrgReportRecipientBudgetsReaderAction/GETNUMBUDGET_FAILURE',
  BUDGETREF_SUCCESS = '@@OrgReportRecipientBudgetsReaderAction/GETBUDGETREF_SUCCESS',
  BUDGETREF_FAILURE = '@@OrgReportRecipientBudgetsReaderAction/GETBUDGETREF_FAILURE',
  BUDGET_SUCCESS = '@@OrgReportRecipientBudgetsReaderAction/GETBUDGET_SUCCESS',
  BUDGET_FAILURE = '@@OrgReportRecipientBudgetsReaderAction/GETBUDGET_FAILURE'
}
