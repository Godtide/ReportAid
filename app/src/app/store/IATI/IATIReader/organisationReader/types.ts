import { PayloadProps } from '../../../types'

export interface OrgGetProps extends PayloadProps {
  data: {
    num: number,
    refs: Array<string>,
    names: Array<string>,
    types: Array<string>
  }
}

export const enum OrgGetActionTypes {
  NUM_SUCCESS = '@@OrgGetAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrgGetAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrgGetAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrgGetAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrgGetAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrgGetAction/GETREFERENCE_FAILURE',
  NAME_SUCCESS = '@@OrgGetAction/GETNAME_SUCCESS',
  NAME_FAILURE = '@@OrgGetAction/GETNAME_FAILURE',
  TYPE_SUCCESS = '@@OrgGetAction/GETTYPE_SUCCESS',
  TYPE_FAILURE = '@@OrgGetAction/GETTYPE_FAILURE'
}
