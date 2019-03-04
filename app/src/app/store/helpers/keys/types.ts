import { PayloadProps } from '../../types'

export const enum KeyTypes {
  ORG = 'org',
  ORGANISATIONS = 'organisations',
  ORGANISATION = 'organisation',
  ORGANISATIONBUDGET = 'organisationBudget',
  ORGANISATIONCOUNTRYBUDGET = 'organisationCountryBudget',
  ORGANISATIONDOC = 'organisationDoc',
  ORGANISATIONEXPENDITURE = 'organisationExpenditure',
  ORGANISATIONRECIPIENTBUDGET = 'organisationRecipientBudget',
  ORGANISATIONREGIONBUDGET = 'organisationRegionBudget',
  ACTIVITIES = 'activities',
  ACTIVITY = 'activity',
  ACTIVITYADDITIONAL = 'activityAdditional'
}

export interface Keys {
  key: string
  keyType: KeyTypes
}

export interface KeyData {
  newKey: string
  org: string
  organisations: string
  organisation: string
  organisationBudget: string
  organisationCountryBudget: string
  organisationDoc: string
  organisationExpenditure: string
  organisationRecipientBudget: string
  organisationRegionBudget: string
  activities: string
  activity: string
  activityAdditional: string
}

export interface KeyProps extends PayloadProps {
  data: KeyData
}

export const enum KeyActionTypes {
  NEW_SUCCESS = '@@KeyActionTypes/NEW_SUCCESS',
  NEW_FAILURE = '@@KeyActionTypes/NEW_FAILURE',
  ORG_SUCCESS = '@@KeyActionTypes/ORG_SUCCESS',
  ORG_FAILURE = '@@KeyActionTypes/ORG_FAILURE',
  ORGANISATIONS_SUCCESS = '@@KeyActionTypes/ORGANISATIONS_SUCCESS',
  ORGANISATIONS_FAILURE = '@@KeyActionTypes/ORGANISATIONS_FAILURE',
  ORGANISATION_SUCCESS = '@@KeyActionTypes/ORGANISATION_SUCCESS',
  ORGANISATION_FAILURE = '@@KeyActionTypes/ORGANISATION_FAILURE',
  ORGANISATIONBUDGET_SUCCESS = '@@KeyActionTypes/ORGANISATIONBUDGET_SUCCESS',
  ORGANISATIONBUDGET_FAILURE = '@@KeyActionTypes/ORGANISATIONBUDGET_FAILURE',
  ORGANISATIONCOUNTRYBUDGET_SUCCESS = '@@KeyActionTypes/ORGANISATIONCOUNTRYBUDGET_SUCCESS',
  ORGANISATIONCOUNTRYBUDGET_FAILURE = '@@KeyActionTypes/ORGANISATIONCOUNTRYBUDGET_FAILURE',
  ORGANISATIONDOC_SUCCESS = '@@KeyActionTypes/ORGANISATIONDOC_SUCCESS',
  ORGANISATIONDOC_FAILURE = '@@KeyActionTypes/ORGANISATIONDOC_FAILURE',
  ORGANISATIONEXPENDITURE_SUCCESS = '@@KeyActionTypes/ORGANISATIONEXPENDITURE_SUCCESS',
  ORGANISATIONEXPENDITURE_FAILURE = '@@KeyActionTypes/ORGANISATIONEXPENDITURE_FAILURE',
  ORGANISATIONRECIPIENTBUDGET_SUCCESS = '@@KeyActionTypes/ORGANISATIONRECIPIENTBUDGET_SUCCESS',
  ORGANISATIONRECIPIENTBUDGET_FAILURE = '@@KeyActionTypes/ORGANISATIONRECIPIENTBUDGET_FAILURE',
  ORGANISATIONREGIONBUDGET_SUCCESS = '@@KeyActionTypes/ORGANISATIONREGIONBUDGET_SUCCESS',
  ORGANISATIONREGIONBUDGET_FAILURE = '@@KeyActionTypes/ORGANISATIONREGIONBUDGET_FAILURE',
  ACTIVITIES_SUCCESS = '@@KeyActionTypes/ACTIVITIES_SUCCESS',
  ACTIVITIES_FAILURE = '@@KeyActionTypes/ACTIVITIES_FAILURE',
  ACTIVITY_SUCCESS = '@@KeyActionTypes/ACTIVITY_SUCCESS',
  ACTIVITY_FAILURE = '@@KeyActionTypes/ACTIVITY_FAILURE',
  ACTIVITYADDITIONAL_SUCCESS = '@@KeyActionTypes/ACTIVITYADDITIONAL_SUCCESS',
  ACTIVITYADDITIONAL_FAILURE = '@@KeyActionTypes/ACTIVITYADDITIONAL_FAILURE'
}
