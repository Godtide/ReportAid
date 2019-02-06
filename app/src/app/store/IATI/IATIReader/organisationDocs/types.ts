import { PayloadProps, DictData } from '../../../types'
import { IATIOrgDocProps } from '../../types'

export interface DocsData {
  [key: string]: IATIOrgDocProps
}

export interface Docs {
  num: number
  data: DocsData
}

export interface OrgDocsData extends DictData {
  [key: string]: Docs
}

export interface OrgDocsReaderProps extends PayloadProps {
  num: number
  data: OrgDocsData
}

export const enum OrgDocsReaderActionTypes {
  NUM_SUCCESS = '@@OrgBudgetsReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrgBudgetsReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrgBudgetsReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrgBudgetsReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrgBudgetsReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrgBudgetsReaderAction/GETREFERENCE_FAILURE',
  NUMDOC_SUCCESS = '@@OrgBudgetsReaderAction/GETNUMDOC_SUCCESS',
  NUMDOC_FAILURE = '@@OrgBudgetsReaderAction/GETNUMDOC_FAILURE',
  DOCREF_SUCCESS = '@@OrgBudgetsReaderAction/GETDOCREF_SUCCESS',
  DOCREF_FAILURE = '@@OrgBudgetsReaderAction/GETDOCREF_FAILURE',
  DOC_SUCCESS = '@@OrgBudgetsReaderAction/GETDOC_SUCCESS',
  DOC_FAILURE = '@@OrgBudgetsReaderAction/GETDOC_FAILURE'
}
