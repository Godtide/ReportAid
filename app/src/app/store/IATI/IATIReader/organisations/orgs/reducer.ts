import { IATIOrgReportProps, IATIOrgReports } from './types'
import { IATIReportActionTypes } from '../types'
import { ActionProps, PayloadProps } from '../../../../types'

const initialState: IATIOrgReportProps = {
  data: {
    '': {
      name: '',
      identifier: ''
    }
  }
}

export const reducer = (state: IATIOrgReportProps = initialState, action: ActionProps): IATIOrgReportProps => {
  switch (action.type) {
    case IATIReportActionTypes.ORGS_INIT: {
      const data = action.payload.data as IATIOrgReportProps
      return data
    }
    case IATIReportActionTypes.ORGS_SUCCESS: {
      const data = action.payload.data as IATIOrgReportProps
      return {...state,...data}
    }
    default:
      return state
  }
}
