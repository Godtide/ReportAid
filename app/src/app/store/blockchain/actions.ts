import { BlockchainInfoProps,
         BlockchainAccountProps,
         BlockchainObjectProps,
         BlockchainActionTypes,
         AddInfoAction,
         AddAccountAction,
         AddObjectAction } from './types'

export const addInfo = (payload: BlockchainInfoProps): AddInfoAction => {
  return {
    type: BlockchainActionTypes.ADD_INFO,
    payload
  }
}

export const addAccount = (payload: BlockchainAccountProps): AddAccountAction => {
  //console.log('In payload', payload)
  return {
    type: BlockchainActionTypes.ADD_ACCOUNT,
    payload
  }
}

export const addObject = (payload: BlockchainObjectProps): AddObjectAction => {
  return {
    type: BlockchainActionTypes.ADD_OBJECT,
    payload
  }
}
