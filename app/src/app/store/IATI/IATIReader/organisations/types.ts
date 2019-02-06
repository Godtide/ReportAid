import { PayloadProps, DictData } from '../../../types'
import { IATIOrgProps } from '../../types'

export interface Data {
  [key: string]: IATIOrgProps
}

export interface s {
  num: number
  data: Data
}

export interface OrgData extends DictData {
  [key: string]: s
}

export interface OrgReaderProps extends PayloadProps {
  num: number
  data: OrgData
}

export const enum OrgReaderActionTypes {
  NUM_SUCCESS = '@@OrgReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrgReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrgReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrgReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrgReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrgReaderAction/GETREFERENCE_FAILURE',
  NUMREP_SUCCESS = '@@OrgReaderAction/GETNUMREP_SUCCESS',
  NUMREP_FAILURE = '@@OrgReaderAction/GETNUMREP_FAILURE',
  REPORTREF_SUCCESS = '@@OrgReaderAction/GETREPORTREF_SUCCESS',
  REPORTREF_FAILURE = '@@OrgReaderAction/GETREPORTREF_FAILURE',
  REPORT_SUCCESS = '@@OrgReaderAction/GETREPORT_SUCCESS',
  REPORT_FAILURE = '@@OrgReaderAction/GETREPORT_FAILURE'
}
