import { IATIOrganisationReportProps } from './types'
import { IATIReportActionTypes } from '../types'
import { ActionProps } from '../../../../types'

const initialState: IATIOrganisationReportProps = {
  data: {
    '': {
      data: {
        '': {
          orgRef: '',
          reportingOrg: {
            orgRef: '',
            orgType: 0,
            isSecondary: false
          },
          lang: '',
          currency: '',
          lastUpdatedTime: ''
        }
      }
    }
  }
}

export const reducer = (state: IATIOrganisationReportProps = initialState, action: ActionProps): IATIOrganisationReportProps => {

  switch (action.type) {
    case IATIReportActionTypes.ORGANISATION_INIT: {
      const data = (action.payload.data as IATIOrganisationReportProps)
      return data
    }
    case IATIReportActionTypes.ORGANISATION_SUCCESS: {
      const data = (action.payload.data as IATIOrganisationReportProps)
      return {...state, ...data}
    }
    default:
      return state
  }
}
