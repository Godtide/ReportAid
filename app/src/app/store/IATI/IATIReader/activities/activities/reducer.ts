import { IATIReportActionTypes } from '../../../types'
import { ActionProps, PayloadProps } from '../../../../types'

const initialState: PayloadProps = {
  data: {}
}

export const reducer = (state: PayloadProps = initialState, action: ActionProps): PayloadProps => {

  switch (action.type) {
    case IATIReportActionTypes.ACTIVITIESPICKER_SUCCESS: {
      const data = (action.payload as PayloadProps)
      return {...state, ...data}
    }
    default:
      return state
  }
}
