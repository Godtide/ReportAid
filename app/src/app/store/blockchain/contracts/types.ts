import { PayloadProps } from '../../types'

export interface OrgContractProps extends PayloadProps {
  data: {
    contract: object
  }
}

export const enum ChainContractActionTypes {
  ADD_ORG = '@@ChainContractActionTypes/ADD_ORG_CONTRACT'
}
