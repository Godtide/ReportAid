import { Action } from 'redux'
//import { ThunkResult } from '../../store'

export interface OrganisationProps {
  name: string
  reference: string
  type: string
}

export interface OrgAddProps {
  result: object
}

export type OrgAddAction = OrgAddSuccessAction | OrgAddFailAction

export interface OrgAddSuccessAction extends Action {
    type: OrganisationActionTypes.ADD_SUCCESS
    payload: OrgAddProps
}

export interface OrgAddFailAction extends Action {
    type: OrganisationActionTypes.ADD_FAILURE
    payload: OrgAddProps
}

export const enum OrganisationActionTypes {
  ADD_SUCCESS = '@@OrgActionTypes/ADD_SUCCESS',
  ADD_FAILURE = '@@OrgActionTypes/ADD_FAILURE'
}
