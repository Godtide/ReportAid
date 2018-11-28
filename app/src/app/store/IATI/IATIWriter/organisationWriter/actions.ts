import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'
import { IATIOrganisations } from '../../../../../blockchain/typechain/IATIOrganisations'

import { OrgWriterActionTypes, OrgAddProps, OrgAddSuccessAction, OrgAddFailureAction } from './types'
import { OrganisationProps } from '../../types'

export const addSuccess = (payload: OrgAddProps): OrgAddSuccessAction => {
  return {
    type: OrgWriterActionTypes.ADD_SUCCESS,
    payload
  }
}

export const addFailure = (payload: OrgAddProps): OrgAddFailureAction => {
  return {
    type: OrgWriterActionTypes.ADD_FAILURE,
    payload
  }
}

export const setOrganisation = (orgDetails: OrganisationProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, OrgAddSuccessAction | OrgAddFailureAction>, getState: Function) => {
    const state = getState()
    const orgContract = state.blockchain.orgContract as IATIOrganisations
    try {
      const tx = await orgContract.setOrganisation(orgDetails.name, orgDetails.reference, orgDetails.type)
      console.log('tx: ', tx)
      dispatch(addSuccess({result: tx}))
    } catch (error) {
      console.log(error)
      dispatch(addFailure({result: {}}))
    }
  }
}
