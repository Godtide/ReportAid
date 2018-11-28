import { OrgWriterActionTypes, OrgAddProps, OrgAddAction } from './types'

const initialState: OrgAddProps = {
  result: {}
}

export const reducer = (state: OrgAddProps = initialState, action: OrgAddAction): OrgAddProps => {
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
