import { PayloadProps, DictData } from '../../../types'
import { OrganisationProps } from '../../types'

export interface OrgData extends DictData {
  [key: string]: OrganisationProps
}

export interface OrgGetProps extends PayloadProps {
  num: number
  data: OrgData
}

export const enum OrgGetActionTypes {
  NUM_SUCCESS = '@@OrgGetAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrgGetAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrgGetAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrgGetAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrgGetAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrgGetAction/GETREFERENCE_FAILURE',
  NAME_SUCCESS = '@@OrgGetAction/GETNAME_SUCCESS',
  NAME_FAILURE = '@@OrgGetAction/GETNAME_FAILURE'
}
