import { OrgReportsActionTypes, ReportWriterProps  } from './types'
import { ActionProps, TxData } from '../../../types'

const initialState: ReportWriterProps = {
  data: {
    '': {}
  }
}

export const reducer = (state: ReportWriterProps = initialState, action: ActionProps): ReportWriterProps => {
  if ( (action.type == OrgReportsActionTypes.ADD_SUCCESS ) ||
       (action.type == OrgReportsActionTypes.ADD_FAILURE )
     ) {
    //console.log('Orgstate: ', state, action)
    //const data: OrgWriterProps = {...state, ...action.payload.data}
    //console.log('TxData: ', data)
    const data = (action.payload.data as TxData)
    return {...state, ...data}
  } else {
    return state
  }
}
