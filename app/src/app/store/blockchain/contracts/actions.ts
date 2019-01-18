import { storeAction } from '../../actions'
import { ContractProps } from './types'
import { ChainContractActionTypes } from './types'

export const addOrgContract = (payload: ContractProps) => storeAction(ChainContractActionTypes.ADD_ORG)(payload)
export const addOrgReportsContract = (payload: ContractProps) => storeAction(ChainContractActionTypes.ADD_REPORTS)(payload)
export const addOrgReportDocsContract = (payload: ContractProps) => storeAction(ChainContractActionTypes.ADD_REPORTDOCS)(payload)
export const addOrgReportBudgetsContract = (payload: ContractProps) => storeAction(ChainContractActionTypes.ADD_REPORTBUDGETS)(payload)
export const addOrgReportExpenditureContract = (payload: ContractProps) => storeAction(ChainContractActionTypes.ADD_REPORTEXPENDITURES)(payload)
export const addOrgReportRecipientBudgetsContract = (payload: ContractProps) => storeAction(ChainContractActionTypes.ADD_REPORTRECIPIENTBUDGETS)(payload)
export const addOrgReportRegionBudgetsContract = (payload: ContractProps) => storeAction(ChainContractActionTypes.ADD_REPORTREGIONBUDGETS)(payload)
export const addOrgReportCountryBudgetsContract = (payload: ContractProps) => storeAction(ChainContractActionTypes.ADD_REPORTCOUNTRYBUDGETS)(payload)
