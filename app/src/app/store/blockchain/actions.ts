import { BlockchainInfoProps, BlockchainAccountProps, BlockchainObjectProps, BlockchainActionTypes, AddInfoAction, AddAccountAction, AddObjectsAction } from './types'

export const addInfo = (payload: BlockchainInfoProps): AddInfoAction => {
  return {
    type: BlockchainActionTypes.ADD_INFO,
    payload
  }
}

export const addAccount = (payload: BlockchainAccountProps): AddAccountAction => {
  return {
    type: BlockchainActionTypes.ADD_ACCOUNT,
    payload
  }
}

export const addObjects = (payload: BlockchainObjectProps): AddObjectsAction => {
  return {
    type: BlockchainActionTypes.ADD_OBJECTS,
    payload
  }
}
