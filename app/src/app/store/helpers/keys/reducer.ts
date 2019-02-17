import { ActionProps } from '../../types'
import { KeyActionTypes, KeyProps, KeyData } from './types'

const initialState: KeyProps = {
  data: {
    newKey: '',
    org: '',
    organisations: '',
    organisation: '',
    activities: '',
    activity: ''
  }
}

export const reducer = (state: KeyProps = initialState, action: ActionProps): KeyProps => {

  switch (action.type) {
    case KeyActionTypes.NEWKEY_SUCCESS:
    case KeyActionTypes.ORGS_SUCCESS:
    case KeyActionTypes.ORGANISATIONS_SUCCESS:
    case KeyActionTypes.ORGANISATION_SUCCESS:
    case KeyActionTypes.ACTIVITIES_SUCCESS:
    case KeyActionTypes.ACTIVITY_SUCCESS: {
      const data = (action.payload.data as KeyProps)
      const newData = {...state, ...data}
      console.log('New key Data! ', newData)
      return {...state, ...data}
    }
    default:
      return state
  }
}