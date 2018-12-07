import { OrgActionTypes, OrgProps } from './types'
import { ActionProps, TxData } from '../../../types'

const initialState: OrgProps = {
  data: {
    '': {}
  }
}

export const reducer = (state: OrgProps = initialState, action: ActionProps): OrgProps => {
  if ( (action.type == OrgActionTypes.ADD_SUCCESS ) ||
       (action.type == OrgActionTypes.ADD_FAILURE )
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
