import { PayloadProps, DictData } from '../../../types'
import { IATIOrgRegionBudgetProps } from '../../types'

export interface RegionBudgetData {
  [key: string]: IATIOrgRegionBudgetProps
}

export interface RegionBudgets {
  num: number
  data: RegionBudgetData
}

export interface OrgRegionBudgetsData extends DictData {
  [key: string]: RegionBudgets
}

export interface OrgRegionBudgetsReaderProps extends PayloadProps {
  num: number
  data: OrgRegionBudgetsData
}

export const enum OrgRegionBudgetsReaderActionTypes {
  NUM_SUCCESS = '@@OrgRegionBudgetsReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrgRegionBudgetsReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrgRegionBudgetsReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrgRegionBudgetsReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrgRegionBudgetsReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrgRegionBudgetsReaderAction/GETREFERENCE_FAILURE',
  NUMBUDGET_SUCCESS = '@@OrgRegionBudgetsReaderAction/GETNUMBUDGET_SUCCESS',
  NUMBUDGET_FAILURE = '@@OrgRegionBudgetsReaderAction/GETNUMBUDGET_FAILURE',
  BUDGETREF_SUCCESS = '@@OrgRegionBudgetsReaderAction/GETBUDGETREF_SUCCESS',
  BUDGETREF_FAILURE = '@@OrgRegionBudgetsReaderAction/GETBUDGETREF_FAILURE',
  BUDGET_SUCCESS = '@@OrgRegionBudgetsReaderAction/GETBUDGET_SUCCESS',
  BUDGET_FAILURE = '@@OrgRegionBudgetsReaderAction/GETBUDGET_FAILURE'
}
