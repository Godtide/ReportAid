import { Action } from 'redux'

export type AddAction = AddOrgAction

export interface AddOrgAction extends Action {
    type: ActionTypes.ADD_ORG
    payload: OrganisationProps
}

// function setOrganisation(string _reference, string _name, string _type) public;
export interface OrganisationProps {
  reference: string
  name: string
  type: string
}

export const enum ActionTypes {
  ADD_ORG = '@@OrgActionTypes/ADD_ORG'
}
