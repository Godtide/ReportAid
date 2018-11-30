import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'
import { IATIOrganisations } from '../../../../../blockchain/typechain/IATIOrganisations'

import { storeAction } from '../../../actions'
import { ActionProps, PayloadProps } from '../../../types'
import { OrganisationProps } from '../../types'

import { OrgWriterActionTypes, OrgWriterProps, TxData } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrgWriterActionTypes): OrgWriterProps => {
    const writerProps = storeAction(actionType)(payload) as OrgWriterProps
    return writerProps
  }
}

export const setOrganisation = (orgDetails: OrganisationProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgContract = state.chainOrgContract.data.contract as IATIOrganisations
    let actionType = OrgWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      const tx = await orgContract.setOrganisation(orgDetails.name, orgDetails.reference, orgDetails.type)
      const key = tx.hash
      txData[key] = tx
      actionType = OrgWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      console.log('In error: ', error)
    }
    console.log('Adding tx: ', txData, actionType)
    dispatch(add({data: txData})(actionType))
  }
}
