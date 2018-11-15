import { Action } from 'redux'

export type OrganisationWriterAction = AddOrganisationAction

export interface AddOrganisationAction extends Action {
    type: OrganisationWriterActionTypes.ADD_ORG
    payload: OrganisationProps
}

// function setOrganisation(string _reference, string _name, string _type) public;
export interface OrganisationProps {
  reference: string
  name: string
  type: string
}

export const enum OrganisationWriterActionTypes {
  ADD_ORG = '@@blockchainActionTypes/ADD_ORG'
}
