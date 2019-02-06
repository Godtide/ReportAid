import { PayloadProps, DictData } from '../../../types'
import { IATIOrganisationRegionBudgetProps } from '../../types'

export interface RegionBudgetData {
  [key: string]: IATIOrganisationRegionBudgetProps
}

export interface RegionBudgets {
  num: number
  data: RegionBudgetData
}

export interface OrganisationRegionBudgetsData extends DictData {
  [key: string]: RegionBudgets
}

export interface OrganisationRegionBudgetsReaderProps extends PayloadProps {
  num: number
  data: OrganisationRegionBudgetsData
}

export const enum OrganisationRegionBudgetsReaderActionTypes {
  NUM_SUCCESS = '@@OrganisationRegionBudgetsReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrganisationRegionBudgetsReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrganisationRegionBudgetsReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrganisationRegionBudgetsReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrganisationRegionBudgetsReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrganisationRegionBudgetsReaderAction/GETREFERENCE_FAILURE',
  NUMBUDGET_SUCCESS = '@@OrganisationRegionBudgetsReaderAction/GETNUMBUDGET_SUCCESS',
  NUMBUDGET_FAILURE = '@@OrganisationRegionBudgetsReaderAction/GETNUMBUDGET_FAILURE',
  BUDGETREF_SUCCESS = '@@OrganisationRegionBudgetsReaderAction/GETBUDGETREF_SUCCESS',
  BUDGETREF_FAILURE = '@@OrganisationRegionBudgetsReaderAction/GETBUDGETREF_FAILURE',
  BUDGET_SUCCESS = '@@OrganisationRegionBudgetsReaderAction/GETBUDGET_SUCCESS',
  BUDGET_FAILURE = '@@OrganisationRegionBudgetsReaderAction/GETBUDGET_FAILURE'
}
