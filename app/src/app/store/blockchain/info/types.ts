import { PayloadProps } from '../../types'

export interface InfoProps extends PayloadProps {
  data: {
    Name: string
    ChainId: string
    ENS: string
    provider: object
  }
}

export const enum ChainInfoActionTypes {
  ADD_INFO = '@@ChainInfoAction/ADD_INFO'
}
