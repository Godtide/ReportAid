import { OrgReportsWriterActionTypes, OrgReportWriterProps  } from './types'
import { ActionProps, TxData } from '../../../types'

const initialState: OrgReportWriterProps = {
  data: {
    '': {}
  }
}

export const reducer = (state: OrgReportWriterProps = initialState, action: ActionProps): OrgReportWriterProps => {
  if ( (action.type == OrgReportsWriterActionTypes.ADD_SUCCESS ) ||
       (action.type == OrgReportsWriterActionTypes.ADD_FAILURE )
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
