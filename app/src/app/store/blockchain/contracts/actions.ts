import { storeAction } from '../../actions'
import { OrgContractProps } from './types'
import { ChainContractActionTypes } from './types'

export const addOrgContract = (payload: OrgContractProps) => storeAction(ChainContractActionTypes.ADD_ORG)(payload)
