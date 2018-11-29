import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'
import { IATIOrganisations } from '../../../../../blockchain/typechain/IATIOrganisations'

import { storeAction } from '../../../actions'
import { ActionProps, PayloadProps } from '../../../types'
import { OrganisationProps } from '../../types'

import { OrgWriterActionTypes } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrgWriterActionTypes) => {
    storeAction(actionType)(payload)
  }
}

export const setOrganisation = (orgDetails: OrganisationProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    console.log('In setOrg')
    const state = getState()
    const orgContract = state.chainOrgContract.data.contract as IATIOrganisations
    try {
      const tx = await orgContract.setOrganisation(orgDetails.name, orgDetails.reference, orgDetails.type)
      console.log('tx: ', tx)
      dispatch(add({data: {tx}})(OrgWriterActionTypes.ADD_SUCCESS))
    } catch (error) {
      console.log('IN ERROR', error)
      dispatch(add({data: {}})(OrgWriterActionTypes.ADD_FAILURE))
    }
  }
}
