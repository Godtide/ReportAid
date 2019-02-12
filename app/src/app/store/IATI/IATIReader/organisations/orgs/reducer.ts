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

  if ( (action.type == IATIOrgsWriterActionTypes.ORGS_SUCCESS ) ) {
    const data = (action.payload.data as IATIOrgReportProps)
    //console.log("Data: ", data)
    state.data = data
  }
  return state
}
