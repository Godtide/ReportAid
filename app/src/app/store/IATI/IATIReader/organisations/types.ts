import { PayloadProps, DictData } from '../../../types'
import { IATIOrganisationsProps } from '../../types'

export interface Data {
  [key: string]: IATIOrganisationsProps
}

export interface Organisations {
  num: number
  data: Data
}

export interface OrgData extends DictData {
  [key: string]: Organisations
}

export interface OrganisationsReaderProps extends PayloadProps {
  num: number
  data: OrgData
}

export const enum OrganisationsReaderActionTypes {
  NUM_SUCCESS = '@@OrganisationsReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrganisationsReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrganisationsReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrganisationsReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrganisationsReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrganisationsReaderAction/GETREFERENCE_FAILURE',
  NUMREP_SUCCESS = '@@OrganisationsReaderAction/GETNUMREP_SUCCESS',
  NUMREP_FAILURE = '@@OrganisationsReaderAction/GETNUMREP_FAILURE',
  REPORTREF_SUCCESS = '@@OrganisationsReaderAction/GETREPORTREF_SUCCESS',
  REPORTREF_FAILURE = '@@OrganisationsReaderAction/GETREPORTREF_FAILURE',
  REPORT_SUCCESS = '@@OrganisationsReaderAction/GETREPORT_SUCCESS',
  REPORT_FAILURE = '@@OrganisationsReaderAction/GETREPORT_FAILURE'
}
