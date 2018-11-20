import { Action } from 'redux'

export type BlockchainAction = AddInfoAction | AddAccountAction | AddObjectAction
export type BlockchainProps = BlockchainInfoProps & BlockchainAccountProps & BlockchainObjectProps

export interface AddInfoAction extends Action {
    type: BlockchainActionTypes.ADD_INFO
    payload: BlockchainInfoProps
}

export interface AddAccountAction extends Action {
    type: BlockchainActionTypes.ADD_ACCOUNT
    payload: BlockchainAccountProps
}

export interface AddObjectAction extends Action {
    type: BlockchainActionTypes.ADD_OBJECT
    payload: BlockchainObjectProps
}

export interface BlockchainInfoProps {
  APIName: string
  networkName: string
  networkChainId: string
  networkENSAddress: string
}

export interface BlockchainAccountProps {
  account: string
}

export interface BlockchainObjectProps {
  provider: object
}

export const enum BlockchainActionTypes {
  ADD_INFO = '@@blockchainActionTypes/ADD_INFO',
  ADD_ACCOUNT = '@@blockchainActionTypes/ADD_ACCOUNT',
  ADD_OBJECT = '@@blockchainActionTypes/ADD_OBJECT'
}
