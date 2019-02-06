import { PayloadProps, DictData } from '../../../types'
import { IATIOrganisationBudgetProps } from '../../types'

export interface BudgetData {
  [key: string]: IATIOrganisationBudgetProps
}

export interface Budgets {
  num: number
  data: BudgetData
}

export interface OrganisationBudgetsData extends DictData {
  [key: string]: Budgets
}

export interface OrganisationBudgetsReaderProps extends PayloadProps {
  num: number
  data: OrganisationBudgetsData
}

export const enum OrganisationBudgetsReaderActionTypes {
  NUM_SUCCESS = '@@OrganisationBudgetsReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrganisationBudgetsReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrganisationBudgetsReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrganisationBudgetsReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrganisationBudgetsReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrganisationBudgetsReaderAction/GETREFERENCE_FAILURE',
  NUMBUDGET_SUCCESS = '@@OrganisationBudgetsReaderAction/GETNUMBUDGET_SUCCESS',
  NUMBUDGET_FAILURE = '@@OrganisationBudgetsReaderAction/GETNUMBUDGET_FAILURE',
  BUDGETREF_SUCCESS = '@@OrganisationBudgetsReaderAction/GETBUDGETREF_SUCCESS',
  BUDGETREF_FAILURE = '@@OrganisationBudgetsReaderAction/GETBUDGETREF_FAILURE',
  BUDGET_SUCCESS = '@@OrganisationBudgetsReaderAction/GETBUDGET_SUCCESS',
  BUDGET_FAILURE = '@@OrganisationBudgetsReaderAction/GETBUDGET_FAILURE'
}
