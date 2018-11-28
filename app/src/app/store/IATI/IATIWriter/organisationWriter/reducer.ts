import { OrgWriterActionTypes, OrgWriterProps } from './types'
import { ActionProps } from '../../../types'

const initialState: OrgWriterProps = {
  data: {
    result: {}
  }
}

export const reducer = (state: OrgWriterProps = initialState, action: ActionProps): OrgWriterProps => {
  console.log('Boom!', action.type, action.payload)
  if ( (action.type == OrgWriterActionTypes.ADD_SUCCESS ) ||
       (action.type == OrgWriterActionTypes.ADD_FAILURE )
     ) {
    console.log('Orgstate: ', state)
    return (<any>Object).assign({}, state, action.payload)
  } else {
    return state
  }
}
