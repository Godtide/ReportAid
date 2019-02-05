import { PayloadProps } from '../../types'

export interface ContractProps extends PayloadProps {
  data: {
    contracts: {
      orgs: object
      organisations: object
      organisation: object
      orgReportDocs: object
      orgReportBudgets: object
      orgReportExpenditure: object
      orgReportRecipientBudgets: object
      orgReportRegionBudgets: object
      orgReportCountryBudgets: object
      activities: object
      activity: object
    }
  }
}

export const enum ChainContractActionTypes {
  ADD_CONTRACTS = '@@ChainContractAction/ADD_CONTRACTS'
}
