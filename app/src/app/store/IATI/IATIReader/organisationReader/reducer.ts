import { OrganisationActionTypes, OrgAddProps, OrgAddAction } from './types'

const initialState: OrgAddProps = {
  result: {}
}

export const reducer = (state: OrgAddProps = initialState, action: OrgAddAction): OrgAddProps => {
  console.log('Boom!', action.type, action.payload)
  if ( (action.type == OrganisationActionTypes.ADD_SUCCESS ) ||
       (action.type == OrganisationActionTypes.ADD_FAILURE )
     ) {
    return (<any>Object).assign({}, state, action.payload)
  } else {
    return state
  }
}
