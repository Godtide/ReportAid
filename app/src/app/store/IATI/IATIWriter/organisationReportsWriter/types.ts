import { PayloadProps, TxData } from '../../../types'

export interface OrgReportsProps extends PayloadProps {
  data: TxData
}

export const enum OrgReportsActionTypes {
  ADD_SUCCESS = '@@OrgReportsWriterAction/ADD_SUCCESS',
  ADD_FAILURE = '@@OrgReportsWriterAction/ADD_FAILURE'
}
