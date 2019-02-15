import { IATIOrganisationsReportProps } from './types'
import { IATIReportActionTypes } from '../types'
import { ActionProps } from '../../../../types'

const initialState: IATIOrganisationsReportProps = {
  data: {
    '': {
      version: '',
      generatedTime: ''
    }
  }
}

export const reducer = (state: IATIOrganisationsReportProps = initialState, action: ActionProps): IATIOrganisationsReportProps => {

  switch (action.type) {
    case IATIReportActionTypes.ORGANISATIONS_INIT: {
      const data = (action.payload.data as IATIOrganisationsReportProps)
      return data
    }
    case IATIReportActionTypes.ORGANISATIONS_SUCCESS: {
      const data = (action.payload.data as IATIOrganisationsReportProps)
      return {...state, ...data}
    }
    default:
      return state
  }
}
