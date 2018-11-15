import { OrganisationWriterActionTypes, OrganisationWriterAction, OrganisationProps } from './types'

export const addOrganisation = (payload: OrganisationProps): OrganisationWriterAction => {
  return {
    type: OrganisationWriterActionTypes.ADD_ORG,
    payload
  }
}
