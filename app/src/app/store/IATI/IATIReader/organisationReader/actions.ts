import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'
import { IATIOrganisations } from '../../../../../blockchain/typechain/IATIOrganisations'

import { OrgGetActionTypes, OrgGetProps, OrgGetNumSuccessAction, OrgGetNumFailureAction } from './types'
import { OrganisationProps } from '../../types'

export const getNumSuccess = (payload: OrgGetProps): OrgGetNumSuccessAction => {
  return {
    type: OrgGetActionTypes.GET_NUM_SUCCESS,
    payload
  }
}

export const getNumFailure = (payload: OrgGetProps): OrgGetNumFailureAction => {
  return {
    type: OrgGetActionTypes.GET_NUM_FAILURE,
    payload
  }
}

export const getNumOrganisations = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, OrgGetNumSuccessAction | OrgGetNumFailureAction>, getState: Function) => {
    const state = getState()
    const orgContract = state.blockchain.orgContract as IATIOrganisations
    try {
      const numOrgs = await orgContract.getNumOrganisations()
      const num = numOrgs.toString()
      console.log('Num: ', num)
      dispatch(getNumSuccess({result: num}))
    } catch (error) {
      console.log(error)
      dispatch(getNumFailure({result: {}}))
    }
  }
}
