import { IATIActivitiesReportProps } from '../types'
import { IATIReportActionTypes } from '../../types'
import { ActionProps } from '../../../../types'

const initialState: IATIActivitiesReportProps = {
  data: {
    data: []
  }
}

export const reducer = (state: IATIActivitiesReportProps = initialState, action: ActionProps): IATIActivitiesReportProps => {

  switch (action.type) {
    case IATIReportActionTypes.ACTIVITIES_INIT: {
      const data = (action.payload.data as IATIActivitiesReportProps)
      return data
    }
    case IATIReportActionTypes.ACTIVITIES_SUCCESS: {
      const data = (action.payload.data as IATIActivitiesReportProps)
      return {...state, ...data}
    }
    default:
      return state
  }
}
