import { PayloadProps, DictData } from '../../../types'
import { IATIOrganisationRecipientBudgetProps } from '../../types'

export interface RecipientBudgetData {
  [key: string]: IATIOrganisationRecipientBudgetProps
}

export interface RecipientBudgets {
  num: number
  data: RecipientBudgetData
}

export interface OrganisationRecipientBudgetsData extends DictData {
  [key: string]: RecipientBudgets
}

export interface OrganisationRecipientBudgetsReaderProps extends PayloadProps {
  num: number
  data: OrganisationRecipientBudgetsData
}

export const enum OrganisationRecipientBudgetsReaderActionTypes {
  NUM_SUCCESS = '@@OrganisationRecipientBudgetsReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrganisationRecipientBudgetsReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrganisationRecipientBudgetsReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrganisationRecipientBudgetsReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrganisationRecipientBudgetsReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrganisationRecipientBudgetsReaderAction/GETREFERENCE_FAILURE',
  NUMBUDGET_SUCCESS = '@@OrganisationRecipientBudgetsReaderAction/GETNUMBUDGET_SUCCESS',
  NUMBUDGET_FAILURE = '@@OrganisationRecipientBudgetsReaderAction/GETNUMBUDGET_FAILURE',
  BUDGETREF_SUCCESS = '@@OrganisationRecipientBudgetsReaderAction/GETBUDGETREF_SUCCESS',
  BUDGETREF_FAILURE = '@@OrganisationRecipientBudgetsReaderAction/GETBUDGETREF_FAILURE',
  BUDGET_SUCCESS = '@@OrganisationRecipientBudgetsReaderAction/GETBUDGET_SUCCESS',
  BUDGET_FAILURE = '@@OrganisationRecipientBudgetsReaderAction/GETBUDGET_FAILURE'
}
