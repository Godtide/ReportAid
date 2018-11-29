import { PayloadProps } from '../../../types'

export interface OrgWriterProps extends PayloadProps {
  data: {
    result: object
  }
}

export const enum OrgWriterActionTypes {
  ADD_SUCCESS = '@@OrgWriterAction/ADD_SUCCESS',
  ADD_FAILURE = '@@OrgWriterAction/ADD_FAILURE'
}
