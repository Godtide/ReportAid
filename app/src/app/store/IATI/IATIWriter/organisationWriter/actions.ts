import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'
import { IATIOrganisations } from '../../../../../blockchain/typechain/IATIOrganisations'

import { storeAction } from '../../../actions'
import { ActionProps } from '../../../types'
import { OrganisationProps } from '../../types'

import { OrgWriterActionTypes } from './types'

const add = (payload: object): Function => {
  return (actionType: OrgWriterActionTypes) => {
    storeAction(actionType)(payload)
  }
}

export const setOrganisation = (orgDetails: OrganisationProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgContract = state.blockchain.orgContract as IATIOrganisations
    try {
      const tx = await orgContract.setOrganisation(orgDetails.name, orgDetails.reference, orgDetails.type)
      console.log('tx: ', tx)
      dispatch(add({result: tx})(OrgWriterActionTypes.ADD_SUCCESS))
    } catch (error) {
      console.log(error)
      dispatch(add({result: {}})(OrgWriterActionTypes.ADD_FAILURE))
    }
  }
}
