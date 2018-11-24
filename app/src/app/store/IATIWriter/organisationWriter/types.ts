import { Action } from 'redux'

// function setOrganisation(string _reference, string _name, string _type) public;
export interface OrganisationProps {
  name: string
  reference: string
  type: string
}

export type OrganisationAddAction = AddOrgAction

export interface AddOrgAction extends Action {
    type: OrganisationActionTypes.ADD_ORG
    payload: OrganisationProps
}

export const enum OrganisationActionTypes {
  ADD_ORG = '@@OrgActionTypes/ADD_ORG'
}
