import { PayloadProps } from '../types'

export interface InfoProps extends PayloadProps {
  data: {
    API: string
    Name: string
    ChainId: string
    ENS: string
  }
}

export interface AccountProps extends PayloadProps {
  data: {
    account: string
  }
}

export interface ObjectProps extends PayloadProps {
  data: {
    provider: object
  }
}

export interface OrgContractProps extends PayloadProps {
  data: {
    orgContract: object
  }
}

export type BlockchainProps = InfoProps & AccountProps & ObjectProps & OrgContractProps

/*
export type BlockchainAction = InfoAddAction | AccountAddAction | ObjectAddAction | OrgContractAddAction
export interface InfoAddAction extends ActionProps {
    type: BlockchainActionTypes.ADD_INFO
    payload: InfoProps
}

export interface AccountAddAction extends ActionProps {
    type: BlockchainActionTypes.ADD_ACCOUNT
    payload: AccountProps
}

export interface ObjectAddAction extends ActionProps {
    type: BlockchainActionTypes.ADD_OBJECT
    payload: ObjectProps
}

export interface OrgContractAddAction extends ActionProps {
    type: BlockchainActionTypes.ADD_ORGCONTRACT
    payload: OrgContractProps
}
*/

export const enum BlockchainActionTypes {
  ADD_INFO = '@@blockchainActionTypes/ADD_INFO',
  ADD_ACCOUNT = '@@blockchainActionTypes/ADD_ACCOUNT',
  ADD_OBJECT = '@@blockchainActionTypes/ADD_OBJECT',
  ADD_ORGCONTRACT = '@@blockchainActionTypes/ADD_ORGCONTRACT'
}
