import { PayloadProps, DictData } from '../../../../types'
import { IATIOrgProps } from '../../../types'

export interface IATIOrgReports extends DictData {
  [key: string]: IATIOrgProps
}

export interface IATIOrgReportProps extends PayloadProps {
  data: IATIOrgReports
}

export const enum IATIOrgsWriterActionTypes {
  ORGS_SUCCESS = '@@IATIOrgsWriterActionTypes/WRITEORGS_SUCCESS',
  ORGS_FAILURE = '@@IATIOrgsWriterActionTypes/WRITEORGS_FAILURE'
}
