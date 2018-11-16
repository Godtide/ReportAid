import { BlockchainInfoProps, BlockchainObjectProps, BlockchainActionTypes, AddInfoAction, AddObjectsAction } from './types'

export const addObjects = (payload: BlockchainObjectProps): AddObjectsAction => {
  return {
    type: BlockchainActionTypes.ADD_OBJECTS,
    payload
  }
}

export const addInfo = (payload: BlockchainInfoProps): AddInfoAction => {
  return {
    type: BlockchainActionTypes.ADD_INFO,
    payload
  }
}
