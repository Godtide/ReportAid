import { PayloadProps, DictData } from '../../../types'
import { IATIOrgReportExpenditureProps } from '../../types'

export interface ExpenditureData {
  [key: string]: IATIOrgReportExpenditureProps
}

export interface Expenditure {
  num: number
  data: ExpenditureData
}

export interface OrgReportExpenditureData extends DictData {
  [key: string]: Expenditure
}

export interface OrgReportExpenditureReaderProps extends PayloadProps {
  num: number
  data: OrgReportExpenditureData
}

export const enum OrgReportExpenditureReaderActionTypes {
  NUM_SUCCESS = '@@OrgReportExpenditureReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrgReportExpenditureReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrgReportExpenditureReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrgReportExpenditureReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrgReportExpenditureReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrgReportExpenditureReaderAction/GETREFERENCE_FAILURE',
  NUMEXPENDITURE_SUCCESS = '@@OrgReportExpenditureReaderAction/GETNUMEXPENDITURE_SUCCESS',
  NUMEXPENDITURE_FAILURE = '@@OrgReportExpenditureReaderAction/GETNUMEXPENDITURE_FAILURE',
  EXPENDITUREREF_SUCCESS = '@@OrgReportExpenditureReaderAction/GETEXPENDITUREREF_SUCCESS',
  EXPENDITUREREF_FAILURE = '@@OrgReportExpenditureReaderAction/GETEXPENDITUREREF_FAILURE',
  EXPENDITURE_SUCCESS = '@@OrgReportExpenditureReaderAction/GETEXPENDITURE_SUCCESS',
  EXPENDITURE_FAILURE = '@@OrgReportExpenditureReaderAction/GETEXPENDITURE_FAILURE'
}
