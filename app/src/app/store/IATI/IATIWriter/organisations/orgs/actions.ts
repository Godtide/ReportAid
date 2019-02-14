import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../../types'
import { OrgProps, IATIOrgProps } from '../../../types'
import { IATIWriterActionTypes } from '../types'

export const setOrg = (details: OrgProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgsContract = state.chainContracts.data.contracts.orgs
    const identifier =  details.code + '-' + details.identifier

    let orgRef = details.orgRef
    if ( orgRef == "" ) {
      orgRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const org: IATIOrgProps = {
        name: details.name,
        identifier: details.code + '-' + details.identifier
    }

    let actionType = IATIWriterActionTypes.ORGS_FAILURE
    let txData: TxData = {}

    try {
      const tx = await orgsContract.setOrg(orgRef, org)
      const key = tx.hash
      txData[key] = tx
      actionType = IATIWriterActionTypes.ORGS_SUCCESS
    } catch (error) {
      txData[-1] = txData
      console.log('setOrg error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}