import { Action } from 'redux'

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

export interface OrgContractProps {
  orgContract: object
}

export type BlockchainAction = InfoAddAction | AccountAddAction | ObjectAddAction | OrgContractAddAction
export type BlockchainProps = InfoProps & AccountProps & ObjectProps & OrgContractProps

export interface InfoAddAction extends Action {
    type: BlockchainActionTypes.ADD_INFO
    payload: InfoProps
}

export interface AccountAddAction extends Action {
    type: BlockchainActionTypes.ADD_ACCOUNT
    payload: AccountProps
}

export interface ObjectAddAction extends Action {
    type: BlockchainActionTypes.ADD_OBJECT
    payload: ObjectProps
}

export interface OrgContractAddAction extends Action {
    type: BlockchainActionTypes.ADD_ORGCONTRACT
    payload: OrgContractProps
}

export const enum BlockchainActionTypes {
  ADD_INFO = '@@blockchainActionTypes/ADD_INFO',
  ADD_ACCOUNT = '@@blockchainActionTypes/ADD_ACCOUNT',
  ADD_OBJECT = '@@blockchainActionTypes/ADD_OBJECT',
  ADD_ORGCONTRACT = '@@blockchainActionTypes/ADD_ORGCONTRACT'
}
