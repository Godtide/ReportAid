import { Action } from 'redux'

export type BlockchainAction = AddInfoAction | AddObjectsAction
export type BlockchainProps = BlockchainObjectProps & BlockchainInfoProps

export interface AddObjectsAction extends Action {
    type: BlockchainActionTypes.ADD_OBJECTS
    payload: BlockchainObjectProps
}

export interface AddInfoAction extends Action {
    type: BlockchainActionTypes.ADD_INFO
    payload: BlockchainInfoProps
}

export interface BlockchainObjectProps {
  web3: object
  ethers: object
}

export interface BlockchainInfoProps {
  APIName: string
  networkName: string
  networkChainId: string
  networkENSAddress: string
  account: string
}

export const enum BlockchainActionTypes {
  ADD_INFO = '@@blockchainActionTypes/ADD_INFO',
  ADD_OBJECTS = '@@blockchainActionTypes/ADD_OBJECTS'
}
