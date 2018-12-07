import { PayloadProps, TxData } from '../../../types'

export interface OrgProps extends PayloadProps {
  data: TxData
}

export const enum OrgActionTypes {
  ADD_SUCCESS = '@@OrgWriterAction/ADD_SUCCESS',
  ADD_FAILURE = '@@OrgWriterAction/ADD_FAILURE'
}
