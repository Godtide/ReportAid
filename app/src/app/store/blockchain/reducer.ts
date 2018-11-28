import { BlockchainActionTypes, InfoProps, AccountProps, ObjectProps, OrgContractProps } from './types'
import { ActionProps, PayloadProps } from '../types'

const initialState: PayloadProps = {
  data: {}
}

export const reducer = (state: PayloadProps = initialState, action: ActionProps): PayloadProps => {
  console.log('blockchain info: ', action.type, action.payload)
  if ( (action.type == BlockchainActionTypes.ADD_INFO ) {}
       (action.type == BlockchainActionTypes.ADD_ACCOUNT ) ||
       (action.type == BlockchainActionTypes.ADD_OBJECT ) ||
       (action.type == BlockchainActionTypes.ADD_ORGCONTRACT )
     ) {
    //console.log(action.payload)
    //return JSON.parse(JSON.stringify(state, action.payload));
    //return (<any>Object).assign({}, state, action.payload)
    const objData = (<any>Object).assign({}, action.payload.data)
    //const objData = {...action.payload.data}
    return {...state, data: objData}
  } else {
    return state
  }
}
