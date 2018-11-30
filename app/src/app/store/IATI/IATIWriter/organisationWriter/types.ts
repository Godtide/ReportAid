import { PayloadProps } from '../../../types'

export interface TxData {
  [tx: string]: object
}

export interface OrgWriterProps extends PayloadProps {
  data: TxData
}

export const enum OrgWriterActionTypes {
  ADD_SUCCESS = '@@OrgWriterAction/ADD_SUCCESS',
  ADD_FAILURE = '@@OrgWriterAction/ADD_FAILURE'
}
