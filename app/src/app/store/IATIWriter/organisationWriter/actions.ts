import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../store'
import { IATIOrganisations } from '../../../../blockchain/typechain/IATIOrganisations'

import { OrganisationActionTypes, OrganisationProps, OrgAddProps, OrgAddSuccessAction, OrgAddFailAction } from './types'

export const addSuccess = (payload: OrgAddProps): OrgAddSuccessAction => {
  return {
    type: OrganisationActionTypes.ADD_SUCCESS,
    payload
  }
}

export const addFailure = (payload: OrgAddProps): OrgAddFailAction => {
  return {
    type: OrganisationActionTypes.ADD_FAILURE,
    payload
  }
}

export const setOrganisation = (orgDetails: OrganisationProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, OrgAddSuccessAction | OrgAddFailAction>, getState: Function) => {
    const state = getState()
    const orgContract = state.blockchain.orgContract as IATIOrganisations
    try {
      const tx = await orgContract.setOrganisation(orgDetails.name, orgDetails.reference, orgDetails.type)
      console.log(tx)
      dispatch(addSuccess({result: tx}))
    } catch (error) {
      console.log(error)
      dispatch(addFailure({result: 'Fail'}))
    }
  }
}
