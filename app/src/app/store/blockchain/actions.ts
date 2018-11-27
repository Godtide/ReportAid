import { InfoProps,
         AccountProps,
         ObjectProps,
         OrgContractProps,
         BlockchainActionTypes,
         InfoAddAction,
         AccountAddAction,
         ObjectAddAction,
         OrgContractAddAction
       } from './types'

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
