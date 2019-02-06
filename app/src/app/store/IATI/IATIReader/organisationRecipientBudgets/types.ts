import { PayloadProps, DictData } from '../../../types'
import { IATIOrgRecipientBudgetProps } from '../../types'

export interface RecipientBudgetData {
  [key: string]: IATIOrgRecipientBudgetProps
}

export interface RecipientBudgets {
  num: number
  data: RecipientBudgetData
}

export interface OrgRecipientBudgetsData extends DictData {
  [key: string]: RecipientBudgets
}

export interface OrgRecipientBudgetsReaderProps extends PayloadProps {
  num: number
  data: OrgRecipientBudgetsData
}

export const enum OrgRecipientBudgetsReaderActionTypes {
  NUM_SUCCESS = '@@OrgRecipientBudgetsReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrgRecipientBudgetsReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrgRecipientBudgetsReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrgRecipientBudgetsReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrgRecipientBudgetsReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrgRecipientBudgetsReaderAction/GETREFERENCE_FAILURE',
  NUMBUDGET_SUCCESS = '@@OrgRecipientBudgetsReaderAction/GETNUMBUDGET_SUCCESS',
  NUMBUDGET_FAILURE = '@@OrgRecipientBudgetsReaderAction/GETNUMBUDGET_FAILURE',
  BUDGETREF_SUCCESS = '@@OrgRecipientBudgetsReaderAction/GETBUDGETREF_SUCCESS',
  BUDGETREF_FAILURE = '@@OrgRecipientBudgetsReaderAction/GETBUDGETREF_FAILURE',
  BUDGET_SUCCESS = '@@OrgRecipientBudgetsReaderAction/GETBUDGET_SUCCESS',
  BUDGET_FAILURE = '@@OrgRecipientBudgetsReaderAction/GETBUDGET_FAILURE'
}
