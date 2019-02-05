import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'
import { OrgProps, IATIOrgProps } from '../../types'
import { OrgWriterActionTypes } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrgWriterActionTypes): TxProps => {
    const writerProps = storeAction(actionType)(payload) as TxProps
    return writerProps
  }
}

export const setOrg = (orgDetails: OrgProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const identifier =  orgDetails.code + '-' + orgDetails.identifier

    const org: IATIOrgProps = {
        orgRef: ethers.utils.formatBytes32String(shortid.generate()),
        name: orgDetails.name,
        identifier: orgDetails.code + '-' + orgDetails.identifier
    }

    const orgContract = state.chainContracts.data.contracts.orgContract
    let actionType = OrgWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}

    try {
      const tx = await orgContract.setOrg(org)
      const key = tx.hash
      txData[key] = tx
      actionType = OrgWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      txData[-1] = txData
      console.log('setOrg error', error)
    }

    dispatch(add({data: {data: txData}})(actionType))
  }
}
