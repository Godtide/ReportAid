import { Action } from 'redux'

export type BlockchainAction = AddDataAction

export interface AddDataAction extends Action {
    type: BlockchainActionTypes.ADD_DATA
    payload: BlockchainInfoProps
}

export interface BlockchainInfoProps {
  APIProvider: string
  networkName: string
  networkChainId: string
  networkENSAddress: string
  account: string
}

export const enum BlockchainActionTypes {
  ADD_DATA = '@@blockchainActionTypes/ADD_DATA'
}
