import { PayloadProps } from '../../types'

export interface AccountProps extends PayloadProps {
  data: {
    account: string
  }
}

export const enum ChainAccountActionTypes {
  ADD_ACCOUNT = '@@ChainAccountActionTypes/ADD_ACCOUNT'
}
