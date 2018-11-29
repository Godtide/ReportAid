import { ChainContractActionTypes, OrgContractProps } from './types'
import { ActionProps } from '../../types'

 const initialOrgContractState: OrgContractProps = {
    data: {
      contract: {}
    }
  }

export const orgContractReducer = (state: OrgContractProps = initialOrgContractState, action: ActionProps): OrgContractProps => {
  //console.log('Org Contract info: ', action.type, action.payload)
  if ( action.type == ChainContractActionTypes.ADD_ORG ) {
    //console.log('Orgstate: ', state)
    return (<any>Object).assign({}, state, action.payload)
  } else {
    return state
  }
}
