import { PayloadProps } from '../types'

export type BlockchainProps = InfoProps & AccountProps & OrgContractProps

export interface InfoProps extends PayloadProps {
  data: {
    API: string
    Name: string
    ChainId: string
    ENS: string
    provider: object
  }
}

export interface AccountProps extends PayloadProps {
  data: {
    account: string
  }
}

export interface OrgContractProps extends PayloadProps {
  data: {
    orgContract: object
  }
}

export const enum BlockchainActionTypes {
  ADD_INFO = '@@blockchainActionTypes/ADD_INFO',
  ADD_ACCOUNT = '@@blockchainActionTypes/ADD_ACCOUNT',
  ADD_OBJECT = '@@blockchainActionTypes/ADD_OBJECT',
  ADD_ORGCONTRACT = '@@blockchainActionTypes/ADD_ORGCONTRACT'
}
