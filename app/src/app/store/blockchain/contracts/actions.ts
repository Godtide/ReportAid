import { storeAction } from '../../actions'
import { ContractProps } from './types'
import { ChainContractActionTypes } from './types'

export const addOrgContract = (payload: ContractProps) => storeAction(ChainContractActionTypes.ADD_ORG)(payload)
export const addOrgReportsContract = (payload: ContractProps) => storeAction(ChainContractActionTypes.ADD_REPORTS)(payload)
