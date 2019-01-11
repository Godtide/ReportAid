import { OrgWriterActionTypes, OrgWriterProps } from './types'
import { ActionProps, TxData } from '../../../types'

const initialState: OrgWriterProps = {
  data: {
    '': {}
  }
}

export const reducer = (state: OrgWriterProps = initialState, action: ActionProps): OrgWriterProps => {
  if ( (action.type == OrgWriterActionTypes.ADD_SUCCESS ) ||
       (action.type == OrgWriterActionTypes.ADD_FAILURE )
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
