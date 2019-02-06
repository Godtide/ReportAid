import { PayloadProps, DictData } from '../../../types'
import { IATIOrgExpenditureProps } from '../../types'

export interface ExpenditureData {
  [key: string]: IATIOrgExpenditureProps
}

export interface Expenditure {
  num: number
  data: ExpenditureData
}

export interface OrgExpenditureData extends DictData {
  [key: string]: Expenditure
}

export interface OrgExpenditureReaderProps extends PayloadProps {
  num: number
  data: OrgExpenditureData
}

export const enum OrgExpenditureReaderActionTypes {
  NUM_SUCCESS = '@@OrgExpenditureReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrgExpenditureReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrgExpenditureReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrgExpenditureReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrgExpenditureReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrgExpenditureReaderAction/GETREFERENCE_FAILURE',
  NUMEXPENDITURE_SUCCESS = '@@OrgExpenditureReaderAction/GETNUMEXPENDITURE_SUCCESS',
  NUMEXPENDITURE_FAILURE = '@@OrgExpenditureReaderAction/GETNUMEXPENDITURE_FAILURE',
  EXPENDITUREREF_SUCCESS = '@@OrgExpenditureReaderAction/GETEXPENDITUREREF_SUCCESS',
  EXPENDITUREREF_FAILURE = '@@OrgExpenditureReaderAction/GETEXPENDITUREREF_FAILURE',
  EXPENDITURE_SUCCESS = '@@OrgExpenditureReaderAction/GETEXPENDITURE_SUCCESS',
  EXPENDITURE_FAILURE = '@@OrgExpenditureReaderAction/GETEXPENDITURE_FAILURE'
}
