import { Action } from 'redux'

export type AddAction = AddInfoAction | AddAccountAction | AddObjectAction
export type BlockchainProps = InfoProps & AccountProps & ObjectProps

export interface AddInfoAction extends Action {
    type:ActionTypes.ADD_INFO
    payload: InfoProps
}

export interface AddAccountAction extends Action {
    type: ActionTypes.ADD_ACCOUNT
    payload: AccountProps
}

export interface AddObjectAction extends Action {
    type: ActionTypes.ADD_OBJECT
    payload: ObjectProps
}

export interface InfoProps {
  API: string
  Name: string
  ChainId: string
  ENS: string
}

export interface AccountProps {
  account: string
}

export interface ObjectProps {
  provider: object
}

export const enum ActionTypes {
  ADD_INFO = '@@blockchainActionTypes/ADD_INFO',
  ADD_ACCOUNT = '@@blockchainActionTypes/ADD_ACCOUNT',
  ADD_OBJECT = '@@blockchainActionTypes/ADD_OBJECT'
}
