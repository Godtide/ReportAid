import { PayloadProps } from '../../../types'

export interface OrgWriterProps extends PayloadProps {
  data: {
    result: object
  }
}

export const enum OrgWriterActionTypes {
  ADD_SUCCESS = '@@OrgWriterActionTypes/ADD_SUCCESS',
  ADD_FAILURE = '@@OrgWriterActionTypes/ADD_FAILURE'
}
