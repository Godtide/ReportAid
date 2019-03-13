import { IATIReportActionTypes } from '../../../types'
import { ActionProps, PayloadProps } from '../../../../types'

const initialState: PayloadProps = {
  data: {}
}

export const reducer = (state: PayloadProps = initialState, action: ActionProps): PayloadProps => {

  switch (action.type) {
    case IATIReportActionTypes.ORGANISATIONPICKER_SUCCESS: {
      const data = (action.payload as PayloadProps)
      //console.log('New organisation: ', action.type, data)
      return {...state, ...data}
    }
    default:
      return state
  }
}
