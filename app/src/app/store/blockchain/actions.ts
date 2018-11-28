import { storeAction } from '../actions'
import { InfoProps, AccountProps, ObjectProps, OrgContractProps } from './types'
import { BlockchainActionTypes } from './types'

export const addInfo = (payload: InfoProps) => storeAction(BlockchainActionTypes.ADD_INFO)(payload)
export const addAccount = (payload: AccountProps) => storeAction(BlockchainActionTypes.ADD_ACCOUNT)(payload)
export const addObject = (payload: ObjectProps) => storeAction(BlockchainActionTypes.ADD_OBJECT)(payload)
export const addOrgContract = (payload: OrgContractProps) => storeAction(BlockchainActionTypes.ADD_ORGCONTRACT)(payload)


 /*
export const addInfo = (payload: InfoProps): InfoAddAction => {
  return {
    type: BlockchainActionTypes.ADD_INFO,
    payload
  }
}

export const addAccount = (payload: AccountProps): AccountAddAction => {
  //console.log('In payload', payload)
  return {
    type: BlockchainActionTypes.ADD_ACCOUNT,
    payload
  }
}

export const addObject = (payload: ObjectProps): ObjectAddAction => {
  return {
    type: BlockchainActionTypes.ADD_OBJECT,
    payload
  }
}

export const addOrgContract = (payload: OrgContractProps): OrgContractAddAction => {
  return {
    type: BlockchainActionTypes.ADD_ORGCONTRACT,
    payload
  }
}
*/
