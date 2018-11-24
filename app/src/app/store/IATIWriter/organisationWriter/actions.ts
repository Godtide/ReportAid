import { OrganisationActionTypes, OrganisationAddAction, OrganisationProps } from './types'

export const addOrganisation = (payload: OrganisationProps): OrganisationAddAction => {
  return {
    type: OrganisationActionTypes.ADD_ORG,
    payload
  }
}
