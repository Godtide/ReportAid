import { OrganisationWriterActionTypes, OrganisationWriterAction, OrganisationProps } from './types'

const initialState: OrganisationProps = {
  reference: '',
  name: '',
  type: ''
}

export const organisationWriterReducer = (state: OrganisationProps = initialState, action: OrganisationWriterAction): OrganisationProps => {
  switch (action.type) {
    case OrganisationWriterActionTypes.ADD_ORG:
      return (<any>Object).assign({}, state, action.payload)
    default:
      return state
  }
}
