import { ActionTypes, AddAction, OrganisationProps } from './types'

const initialState: OrganisationProps = {
  reference: '',
  name: '',
  type: ''
}

export const reducer = (state: OrganisationProps = initialState, action: AddAction): OrganisationProps => {
  switch (action.type) {
    case ActionTypes.ADD_ORG:
      return (<any>Object).assign({}, state, action.payload)
    default:
      return state
  }
}
