import { IATIReportActionTypes } from '../../../types'
import { ActionProps, PayloadProps } from '../../../../types'

const initialState: PayloadProps = {
  data: {}
}

export const reducer = (state: PayloadProps = initialState, action: ActionProps): PayloadProps => {

  switch (action.type) {
    case IATIReportActionTypes.ACTIVITYPICKER_SUCCESS: {
      const data = (action.payload.data as PayloadProps)
      return {...state, ...data}
    }
    default:
      return state
  }
}
