import { OrgWriterActionTypes, OrgWriterProps } from './types'
import { ActionProps } from '../../../types'

const initialState: OrgWriterProps = {
  data: {
    '': {}
  }
}

export const reducer = (state: OrgWriterProps = initialState, action: ActionProps): OrgWriterProps => {
  //console.log('Boom!', action.type, action.payload)
  if ( (action.type == OrgWriterActionTypes.ADD_SUCCESS ) ||
       (action.type == OrgWriterActionTypes.ADD_FAILURE )
     ) {
    //console.log('Orgstate: ', state, action)
    const data: OrgWriterProps = {...state, ...action.payload.data}
    //console.log('TxData: ', data)
    return data
  } else {
    return state
  }
}
