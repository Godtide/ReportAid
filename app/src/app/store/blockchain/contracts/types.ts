import { PayloadProps } from '../../types'

export interface ContractProps extends PayloadProps {
  data: {
    contracts: {
      orgContract: object
      orgReportsContract: object
      orgReportDocsContract: object
      orgReportBudgetsContract: object
      orgReportExpenditureContract: object
      orgReportRecipientBudgetsContract: object
      orgReportRegionBudgetsContract: object
      orgReportCountryBudgetsContract: object
    }
  }
}

export const enum ChainContractActionTypes {
  ADD_CONTRACTS = '@@ChainContractAction/ADD_CONTRACTS'
}
