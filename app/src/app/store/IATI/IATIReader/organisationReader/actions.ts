import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'
import { IATIOrganisations } from '../../../../../blockchain/typechain/IATIOrganisations'

import { storeAction } from '../../../actions'
import { ActionProps, PayloadProps } from '../../../types'

import { OrgGetActionTypes } from './types'

const get = (payload: PayloadProps): Function => {
  return (actionType: OrgGetActionTypes) => {
    storeAction(actionType)(payload)
  }
}

export const getNumOrganisations = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgContract = state.chainOrgContract.data.contract as IATIOrganisations
    try {
      const numOrgs = await orgContract.getNumOrganisations()
      const num = numOrgs.toString()
      console.log('Num: ', num)
      dispatch(get({data: {num}})(OrgGetActionTypes.GET_NUM_SUCCESS))
    } catch (error) {
      console.log(error)
      dispatch(get({data: {}})(OrgGetActionTypes.GET_NUM_FAILURE))
    }
  }
}
