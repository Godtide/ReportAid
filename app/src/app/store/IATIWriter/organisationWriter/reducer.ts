import { OrganisationActionTypes, OrganisationAddAction, OrganisationProps } from './types'

const initialState: OrganisationProps = {
  reference: '',
  name: '',
  type: ''
}

export const reducer = (state: OrganisationProps = initialState, action: OrganisationAddAction): OrganisationProps => {
  console.log('Boom!', action.type, action.payload)
  switch (action.type) {
    case OrganisationActionTypes.ADD_ORG:
      return (<any>Object).assign({}, state, action.payload)
    default:
      return state
  }
}
