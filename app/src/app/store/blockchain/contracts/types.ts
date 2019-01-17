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
  ADD_ORG = '@@ChainContractAction/ADD_ORG',
  ADD_REPORTS = '@@ChainContractAction/ADD_ORGREPORTS',
  ADD_REPORTDOCS = '@@ChainContractAction/ADD_REPORTDOCS',
  ADD_REPORTBUDGETS = '@@ChainContractAction/ADD_REPORTBUDGETS',
  ADD_REPORTEXPENDITURES = '@@ChainContractAction/ADD_REPORTEXPENDITURES',
  ADD_REPORTRECIPIENTBUDGETS = '@@ChainContractAction/ADD_REPORTRECIPIENTBUDGETS',
  ADD_REPORTREGIONBUDGETS = '@@ChainContractAction/ADD_REPORTREGIONBUDGETS',
  ADD_REPORTCOUNTRYBUDGETS = '@@ChainContractAction/ADD_REPORTCOUNTRYBUDGETS'
}
