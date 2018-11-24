import { InfoProps,
         AccountProps,
         ObjectProps,
         BlockchainActionTypes,
         AddInfoAction,
         AddAccountAction,
         AddObjectAction } from './types'

export const addInfo = (payload: InfoProps): AddInfoAction => {
  return {
    type: BlockchainActionTypes.ADD_INFO,
    payload
  }
}

export const addAccount = (payload: AccountProps): AddAccountAction => {
  //console.log('In payload', payload)
  return {
    type: BlockchainActionTypes.ADD_ACCOUNT,
    payload
  }
}

export const addObject = (payload: ObjectProps): AddObjectAction => {
  return {
    type: BlockchainActionTypes.ADD_OBJECT,
    payload
  }
}
