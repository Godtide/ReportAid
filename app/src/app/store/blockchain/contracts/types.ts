import { PayloadProps } from '../../types'

export interface ContractProps extends PayloadProps {
  data: {
    contracts: {
      orgContract: {},
      orgReportsContract: {}
    }
  }
}

export const enum ChainContractActionTypes {
  ADD_ORG = '@@ChainContractAction/ADD_ORG',
  ADD_REPORTS = '@@ChainContractAction/ADD_ORGREPORTS'
}
