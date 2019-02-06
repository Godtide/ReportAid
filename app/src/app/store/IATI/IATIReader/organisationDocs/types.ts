import { PayloadProps, DictData } from '../../../types'
import { IATIOrgDocProps } from '../../types'

export interface DocsData {
  [key: string]: IATIOrgDocProps
}

export interface Docs {
  num: number
  data: DocsData
}

export interface OrganisationDocsData extends DictData {
  [key: string]: Docs
}

export interface OrganisationDocsReaderProps extends PayloadProps {
  num: number
  data: OrganisationDocsData
}

export const enum OrganisationDocsReaderActionTypes {
  NUM_SUCCESS = '@@OrganisationDocsReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrganisationDocsReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrganisationDocsReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrganisationDocsReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrganisationDocsReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrganisationDocsReaderAction/GETREFERENCE_FAILURE',
  NUMDOC_SUCCESS = '@@OrganisationDocsReaderAction/GETNUMDOC_SUCCESS',
  NUMDOC_FAILURE = '@@OrganisationDocsReaderAction/GETNUMDOC_FAILURE',
  DOCREF_SUCCESS = '@@OrganisationDocsReaderAction/GETDOCREF_SUCCESS',
  DOCREF_FAILURE = '@@OrganisationDocsReaderAction/GETDOCREF_FAILURE',
  DOC_SUCCESS = '@@OrganisationDocsReaderAction/GETDOC_SUCCESS',
  DOC_FAILURE = '@@OrganisationDocsReaderAction/GETDOC_FAILURE'
}
