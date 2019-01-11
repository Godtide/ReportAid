import { PayloadProps, DictData } from '../../../types'
import { IATIOrgProps } from '../../types'

export interface OrgData extends DictData {
  [key: string]: IATIOrgProps
}

export interface OrgReaderProps extends PayloadProps {
  num: number
  data: OrgData
}

export const enum OrgReaderActionTypes {
  NUM_SUCCESS = '@@OrgReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrgReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrgReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrgReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrgReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrgReaderAction/GETREFERENCE_FAILURE',
  ORG_SUCCESS = '@@OrgReaderAction/GETORG_SUCCESS',
  ORG_FAILURE = '@@OrgReaderAction/GETORG_FAILURE',
  NAME_SUCCESS = '@@OrgReaderAction/GETNAME_SUCCESS',
  NAME_FAILURE = '@@OrgReaderAction/GETNAME_FAILURE',
  ID_SUCCESS = '@@OrgReaderAction/GETID_SUCCESS',
  ID_FAILURE = '@@OrgReaderAction/GETID_FAILURE'
}
