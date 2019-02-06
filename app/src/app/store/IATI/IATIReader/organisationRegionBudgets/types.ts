import { PayloadProps, DictData } from '../../../types'
import { IATIOrgReportRegionBudgetProps } from '../../types'

export interface RegionBudgetData {
  [key: string]: IATIOrgReportRegionBudgetProps
}

export interface RegionBudgets {
  num: number
  data: RegionBudgetData
}

export interface OrgReportRegionBudgetsData extends DictData {
  [key: string]: RegionBudgets
}

export interface OrgReportRegionBudgetsReaderProps extends PayloadProps {
  num: number
  data: OrgReportRegionBudgetsData
}

export const enum OrgReportRegionBudgetsReaderActionTypes {
  NUM_SUCCESS = '@@OrgReportRegionBudgetsReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrgReportRegionBudgetsReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrgReportRegionBudgetsReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrgReportRegionBudgetsReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrgReportRegionBudgetsReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrgReportRegionBudgetsReaderAction/GETREFERENCE_FAILURE',
  NUMBUDGET_SUCCESS = '@@OrgReportRegionBudgetsReaderAction/GETNUMBUDGET_SUCCESS',
  NUMBUDGET_FAILURE = '@@OrgReportRegionBudgetsReaderAction/GETNUMBUDGET_FAILURE',
  BUDGETREF_SUCCESS = '@@OrgReportRegionBudgetsReaderAction/GETBUDGETREF_SUCCESS',
  BUDGETREF_FAILURE = '@@OrgReportRegionBudgetsReaderAction/GETBUDGETREF_FAILURE',
  BUDGET_SUCCESS = '@@OrgReportRegionBudgetsReaderAction/GETBUDGET_SUCCESS',
  BUDGET_FAILURE = '@@OrgReportRegionBudgetsReaderAction/GETBUDGET_FAILURE'
}
