import { OrgGetActionTypes, OrgGetProps } from './types'
import { ActionProps } from '../../../types'

const initialState: OrgGetProps = {
  data: {
    result: {}
  }
}

export const reducer = (state: OrgGetProps = initialState, action: ActionProps): OrgGetProps => {
  //console.log('Boom!', action.type, action.payload)
  if ( (action.type == OrgGetActionTypes.GET_NUM_SUCCESS ) ||
       (action.type == OrgGetActionTypes.GET_NUM_FAILURE )
     ) {
    return (<any>Object).assign({}, state, action.payload)
  } else {
    return state
  }
}
