import { PayloadProps, DictData } from '../../../types'
import { IATIOrgReportProps } from '../../types'

export interface ReportData extends DictData {
  [key: string]: IATIOrgReportProps
}

export interface OrgReportData extends DictData {
  [key: string]: ReportData
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
  REPORT_SUCCESS = '@@OrgReportReaderAction/GETORG_SUCCESS',
  REPORT_FAILURE = '@@OrgReportReaderAction/GETORG_FAILURE'
}
