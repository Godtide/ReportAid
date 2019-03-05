import { ActionProps } from '../../types'
import { KeyActionTypes, KeyProps, KeyData } from './types'

const initialState: KeyProps = {
  data: {
    newKey: '',
    org: '',
    organisations: '',
    organisation: '',
    organisationBudget: '',
    organisationCountryBudget: '',
    organisationDoc: '',
    organisationExpenditure: '',
    organisationRecipientBudget: '',
    organisationRegionBudget: '',
    activities: '',
    activity: '',
    activityAdditional: ''
  }
}

export const reducer = (state: KeyProps = initialState, action: ActionProps): KeyProps => {

  switch (action.type) {
    case KeyActionTypes.NEW_SUCCESS: {
      const data = action.payload as KeyProps
      return data
    }
    case KeyActionTypes.ORG_SUCCESS:
    case KeyActionTypes.ORGANISATIONS_SUCCESS:
    case KeyActionTypes.ORGANISATION_SUCCESS:
    case KeyActionTypes.ORGANISATIONBUDGET_SUCCESS:
    case KeyActionTypes.ORGANISATIONCOUNTRYBUDGET_SUCCESS:
    case KeyActionTypes.ORGANISATIONDOC_SUCCESS:
    case KeyActionTypes.ORGANISATIONEXPENDITURE_SUCCESS:
    case KeyActionTypes.ORGANISATIONRECIPIENTBUDGET_SUCCESS:
    case KeyActionTypes.ORGANISATIONREGIONBUDGET_SUCCESS:
    case KeyActionTypes.ACTIVITIES_SUCCESS:
    case KeyActionTypes.ACTIVITY_SUCCESS:
    case KeyActionTypes.ACTIVITYADDITIONAL_SUCCESS: {
      const data = (action.payload.data as KeyProps)
      //console.log('Here!: ', data, action.type)
      return {...state, ...data}
    }
    default:
      return state
  }
}
