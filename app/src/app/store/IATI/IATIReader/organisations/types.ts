import { PayloadProps, DictData } from '../../../types'
import { IATIOrgReportProps } from '../../types'

export interface ReportData {
  [key: string]: IATIOrgReportProps
}

export interface Reports {
  num: number
  data: ReportData
}

export interface OrgReportData extends DictData {
  [key: string]: Reports
}

export interface OrgReportReaderProps extends PayloadProps {
  num: number
  data: OrgReportData
}

export const enum OrgReportReaderActionTypes {
  NUM_SUCCESS = '@@OrgReportReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrgReportReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrgReportReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrgReportReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrgReportReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrgReportReaderAction/GETREFERENCE_FAILURE',
  NUMREP_SUCCESS = '@@OrgReportReaderAction/GETNUMREP_SUCCESS',
  NUMREP_FAILURE = '@@OrgReportReaderAction/GETNUMREP_FAILURE',
  REPORTREF_SUCCESS = '@@OrgReportReaderAction/GETREPORTREF_SUCCESS',
  REPORTREF_FAILURE = '@@OrgReportReaderAction/GETREPORTREF_FAILURE',
  REPORT_SUCCESS = '@@OrgReportReaderAction/GETREPORT_SUCCESS',
  REPORT_FAILURE = '@@OrgReportReaderAction/GETREPORT_FAILURE'
}
