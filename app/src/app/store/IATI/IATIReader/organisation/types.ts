import { PayloadProps, DictData } from '../../../types'
import { IATIOrganisationProps } from '../../types'

export interface Data {
  [key: string]: IATIOrganisationProps
}

export interface Organisation {
  data: Data
}

export interface OrganisationData extends DictData {
  [key: string]: Organisation
}

export interface OrganisationReaderProps extends PayloadProps {
  data: OrganisationData
}

export const enum OrganisationReaderActionTypes {
  NUM_SUCCESS = '@@OrganisationReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrganisationReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrganisationReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrganisationReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrganisationReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrganisationReaderAction/GETREFERENCE_FAILURE',
  NUMREP_SUCCESS = '@@OrganisationReaderAction/GETNUMREP_SUCCESS',
  NUMREP_FAILURE = '@@OrganisationReaderAction/GETNUMREP_FAILURE',
  REPORTREF_SUCCESS = '@@OrganisationReaderAction/GETREPORTREF_SUCCESS',
  REPORTREF_FAILURE = '@@OrganisationReaderAction/GETREPORTREF_FAILURE',
  REPORT_SUCCESS = '@@OrganisationReaderAction/GETREPORT_SUCCESS',
  REPORT_FAILURE = '@@OrganisationReaderAction/GETREPORT_FAILURE'
}
