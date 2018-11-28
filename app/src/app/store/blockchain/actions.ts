import { storeAction } from '../actions'
import { InfoProps, AccountProps, ObjectProps, OrgContractProps } from './types'
import { BlockchainActionTypes } from './types'

export const addInfo = (payload: InfoProps) => storeAction(BlockchainActionTypes.ADD_INFO)(payload)
export const addAccount = (payload: AccountProps) => storeAction(BlockchainActionTypes.ADD_ACCOUNT)(payload)
export const addOrgContract = (payload: OrgContractProps) => storeAction(BlockchainActionTypes.ADD_ORGCONTRACT)(payload)
