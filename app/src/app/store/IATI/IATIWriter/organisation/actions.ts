import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps, TxData } from '../../../types'
import { OrganisationProps, IATIOrgProps } from '../../types'
import { OrgWriterActionTypes, OrgWriterProps } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrgWriterActionTypes): OrgWriterProps => {
    const writerProps = storeAction(actionType)(payload) as OrgWriterProps
    return writerProps
  }
}

export const setOrganisation = (orgDetails: OrganisationProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const identifier =  orgDetails.code + '-' + orgDetails.identifier
    const reference = shortid.generate()
    let orgReference = ethers.utils.formatBytes32String(reference)
    const org: IATIOrgProps = {
        orgRef: orgReference,
        name: orgDetails.name,
        identifier: orgDetails.code + '-' + orgDetails.identifier
    }
    //console.log('Org: ', org)
    const orgContract = state.chainContracts.data.contracts.orgContract
    let actionType = OrgWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}

    /*let gasLimit = await state.chainInfo.data.provider.estimateGas(setOrganisation)
    gasLimit *= 4
    console.log('Gaslimit: ', gasLimit)*/
    try {
      //const tx = await orgContract.setOrganisation(org, {gasLimit: gasLimit})
      const tx = await orgContract.setOrganisation(org)
      const key = tx.hash
      txData[key] = tx
      actionType = OrgWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      console.log('setOrganisation error', error)
    }

    dispatch(add({data: {data: txData}})(actionType))
  }
}
