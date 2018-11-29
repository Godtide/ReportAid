import { PayloadProps } from '../../types'

export interface OrgContractProps extends PayloadProps {
  data: {
    contract: object
  }
}

export const enum ChainContractActionTypes {
  ADD_ORG = '@@ChainContractAction/ADD_ORG'
}
