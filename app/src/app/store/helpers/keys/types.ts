import { PayloadProps, DictData } from '../../types'

export interface KeyData {
  newKey: string
  org: string
  organisations: string
  organisation: string
  activities: string
  activity: string
}

export interface KeyProps extends PayloadProps {
  data: KeyData
}

export const enum KeyActionTypes {
  NEWKEY_SUCCESS = '@@IATIReportActionTypes/KEYBUDGET_SUCCESS',
  NEWKEY_FAILURE = '@@IATIReportActionTypes/KEYBUDGET_FAILURE',
  ORGS_SUCCESS = '@@IATIReportActionTypes/KEYORGS_SUCCESS',
  ORGS_FAILURE = '@@IATIReportActionTypes/KEYORGS_FAILURE',
  ORGANISATIONS_SUCCESS = '@@IATIReportActionTypes/KEYORGANISATIONS_SUCCESS',
  ORGANISATIONS_FAILURE = '@@IATIReportActionTypes/KEYORGANISATIONS_FAILURE',
  ORGANISATION_SUCCESS = '@@IATIReportActionTypes/KEYORGANISATION_SUCCESS',
  ORGANISATION_FAILURE = '@@IATIReportActionTypes/KEYORGANISATION_FAILURE',
  ACTIVITIES_SUCCESS = '@@IATIReportActionTypes/KEYACTIVITIES_SUCCESS',
  ACTIVITIES_FAILURE = '@@IATIReportActionTypes/KEYACTIVITIES_FAILURE',
  ACTIVITY_SUCCESS = '@@IATIReportActionTypes/KEYACTIVITY_SUCCESS',
  ACTIVITY_FAILURE = '@@IATIReportActionTypes/KEYACTIVITY_FAILURE'
}
