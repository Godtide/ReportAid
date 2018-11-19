import { Action } from 'redux'

export type BlockchainAction = AddInfoAction | AddAccountAction | AddObjectsAction
export type BlockchainProps = BlockchainInfoProps & BlockchainAccountProps & BlockchainObjectProps

export interface AddInfoAction extends Action {
    type: BlockchainActionTypes.ADD_INFO
    payload: BlockchainInfoProps
}

export interface AddAccountAction extends Action {
    type: BlockchainActionTypes.ADD_ACCOUNT
    payload: BlockchainAccountProps
}

export interface AddObjectsAction extends Action {
    type: BlockchainActionTypes.ADD_OBJECTS
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
  ADD_OBJECTS = '@@blockchainActionTypes/ADD_OBJECTS'
}
