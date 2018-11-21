import { ActionTypes, AddAction, OrganisationProps } from './types'

export const addOrganisation = (payload: OrganisationProps): AddAction => {
  return {
    type: ActionTypes.ADD_ORG,
    payload
  }
}
