import { IATIOrgsWriterActionTypes, IATIOrgReportProps } from './types'
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

  const payload = action.payload as PayloadProps
  if ( (action.type == IATIOrgsWriterActionTypes.ORGS_SUCCESS ) ) {
    const data = (action.payload.data as IATIOrgReportProps)
    return {...state, ...data}
  } else {
    return state
  }
}
