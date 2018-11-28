import { PayloadProps } from '../../../types'

export interface OrgGetProps extends PayloadProps {
  data: {
    result: object
  }
}

/* export type OrgGetAction = OrgGetNumSuccessAction | OrgGetNumFailureAction

export interface OrgGetNumSuccessAction extends Action {
    type: OrgGetActionTypes.GET_NUM_SUCCESS
    payload: OrgGetProps
}

export interface OrgGetNumFailureAction extends Action {
    type: OrgGetActionTypes.GET_NUM_FAILURE
    payload: OrgGetProps
} */

/*function getOrganisationExists(string _reference) constant returns (bool)",
"function getNumOrganisations() constant returns (uint256)",
"function getOrganisationReference(uint256 _index) constant returns (string)",
"function getOrganisationName(string _reference) constant returns (string)",
"function getOrganisationType(string _reference) constant returns (string)",*/

export const enum OrgGetActionTypes {
  GET_NUM_SUCCESS = '@@OrgGetActionTypes/GET_NUM_SUCCESS',
  GET_NUM_FAILURE = '@@OrgGetActionTypes/GET_NUM_FAILURE',
  GET_EXISTS_SUCCESS = '@@OrgGetActionTypes/GET_EXISTS_SUCCESS',
  GET_EXISTS_FAILURE = '@@OrgGetActionTypes/GET_EXISTS_FAILURE',
  GET_REFERENCE_SUCCESS = '@@OrgGetActionTypes/GET_REFERENCE_SUCCESS',
  GET_REFERENCE_FAILURE = '@@OrgGetActionTypes/GET_REFERENCE_FAILURE',
  GET_NAME_SUCCESS = '@@OrgGetActionTypes/GET_NAME_SUCCESS',
  GET_NAME_FAILURE = '@@OrgGetActionTypes/GET_NAME_FAILURE',
  GET_TYPE_SUCCESS = '@@OrgGetActionTypes/GET_TYPE_SUCCESS',
  GET_TYPE_FAILURE = '@@OrgGetActionTypes/GET_TYPE_FAILURE',
}
