import { PayloadProps, TxData } from '../../../types'

export interface OrgReportWriterProps extends PayloadProps {
  data: TxData
}

export const enum OrgReportsWriterActionTypes {
  ADD_SUCCESS = '@@OrgReportsWriterAction/ADD_SUCCESS',
  ADD_FAILURE = '@@OrgReportsWriterAction/ADD_FAILURE'
}
