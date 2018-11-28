import { Action } from 'redux'

export interface OrgAddProps {
  result: object
}

export type OrgAddAction = OrgAddSuccessAction | OrgAddFailureAction

export interface OrgAddSuccessAction extends Action {
    type: OrgWriterActionTypes.ADD_SUCCESS
    payload: OrgAddProps
}

export interface OrgAddFailureAction extends Action {
    type: OrgWriterActionTypes.ADD_FAILURE
    payload: OrgAddProps
}

export const enum OrgWriterActionTypes {
  ADD_SUCCESS = '@@OrgWriterActionTypes/ADD_SUCCESS',
  ADD_FAILURE = '@@OrgWriterActionTypes/ADD_FAILURE'
}
