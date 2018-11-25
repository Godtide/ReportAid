import { InfoProps,
         AccountProps,
         ObjectProps,
         OrgContractProps,
         BlockchainActionTypes,
         AddInfoAction,
         AddAccountAction,
         AddObjectAction,
         AddOrgContractAction
       } from './types'

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

export const addOrgContract = (payload: OrgContractProps): AddOrgContractAction => {
  return {
    type: BlockchainActionTypes.ADD_ORGCONTRACT,
    payload
  }
}
