import { AccountProps, ChainAccountActionTypes } from './types'
import { ActionProps } from '../../types'

const initialAccountState: AccountProps = {
   data: {
     account: ''
   }
 }

export const accountReducer = (state: AccountProps = initialAccountState, action: ActionProps): AccountProps => {
  console.log('Account info: ', action.type, action.payload)
  if ( action.type == ChainAccountActionTypes.ADD_ACCOUNT ) {
    //console.log('Orgstate: ', state)
    return (<any>Object).assign({}, state, action.payload)
  } else {
    return state
  }
}
