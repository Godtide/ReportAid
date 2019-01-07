import { ThunkDispatch } from 'redux-thunk'
import shortid from 'shortid'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { IATIOrganisations } from '../../../../../blockchain/typechain/IATIOrganisations'

import { storeAction } from '../../../actions'
import { ActionProps, PayloadProps, TxData } from '../../../types'
import { OrganisationProps } from '../../types'

import { OrgActionTypes, OrgProps } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrgActionTypes): OrgProps => {
    const writerProps = storeAction(actionType)(payload) as OrgProps
    return writerProps
  }
}

export const setOrganisation = (orgDetails: OrganisationProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const identifier =  orgDetails.code + '-' + orgDetails.identifier
    const reference = shortid.generate()
    let orgReference = ethers.utils.formatBytes32String(reference)
    const org = {
        orgRef: orgReference,
        name: orgDetails.name,
        identifier: orgDetails.code + '-' + orgDetails.identifier
    }
    //console.log('Org: ', org)
    const orgContract = state.chainContracts.data.contracts.orgContract as IATIOrganisations
    let actionType = OrgActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      const tx = await orgContract.setOrganisation(org)
      const key = tx.hash
      txData[key] = tx
      actionType = OrgActionTypes.ADD_SUCCESS
    } catch (error) {
      console.log('setOrganisation error', error)
    }
    //console.log('Adding tx: ', txData, actionType)
    dispatch(add({data: {data: txData}})(actionType))
  }
}
