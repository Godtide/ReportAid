import { PayloadProps, DictData } from '../../../types'
import { IATIOrgReportDocProps } from '../../types'

export interface ReportDocsData {
  [key: string]: IATIOrgReportDocProps
}

export interface ReportDocs {
  num: number
  data: ReportDocsData
}

export interface OrgReportDocsData extends DictData {
  [key: string]: ReportDocs
}

export interface OrgReportDocsReaderProps extends PayloadProps {
  num: number
  data: OrgReportDocsData
}

export const enum OrgReportDocsReaderActionTypes {
  NUM_SUCCESS = '@@OrgReportBudgetsReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrgReportBudgetsReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrgReportBudgetsReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrgReportBudgetsReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrgReportBudgetsReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrgReportBudgetsReaderAction/GETREFERENCE_FAILURE',
  NUMDOC_SUCCESS = '@@OrgReportBudgetsReaderAction/GETNUMDOC_SUCCESS',
  NUMDOC_FAILURE = '@@OrgReportBudgetsReaderAction/GETNUMDOC_FAILURE',
  DOCREF_SUCCESS = '@@OrgReportBudgetsReaderAction/GETDOCREF_SUCCESS',
  DOCREF_FAILURE = '@@OrgReportBudgetsReaderAction/GETDOCREF_FAILURE',
  DOC_SUCCESS = '@@OrgReportBudgetsReaderAction/GETDOC_SUCCESS',
  DOC_FAILURE = '@@OrgReportBudgetsReaderAction/GETDOC_FAILURE'
}
