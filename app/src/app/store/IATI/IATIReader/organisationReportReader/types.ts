import { PayloadProps, DictData } from '../../../types'
import { IATIOrgReportProps } from '../../types'

export interface OrgReportData extends DictData {
  [key: string]: IATIOrgReportProps
}

export interface OrgReportGetProps extends PayloadProps {
  num: number
  data: OrgReportData
}

export const enum OrgGetActionTypes {
  NUM_SUCCESS = '@@OrgReportsGetAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrgReportsGetAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrgReportsGetAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrgReportsGetAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrgReportsGetAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrgReportsGetAction/GETREFERENCE_FAILURE',
  REPORT_SUCCESS = '@@OrgReportsGetAction/GETORG_SUCCESS',
  REPORT_FAILURE = '@@OrgReportsGetAction/GETORG_FAILURE'
}
