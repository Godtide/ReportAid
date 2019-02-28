import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../../actions'

import { ActionProps, TxReport } from '../../../../types'
import { ActivitiesProps, IATIActivitiesProps } from '../types'
import { IATIWriterActionTypes } from '../../types'

import { Transaction } from '../../../../../utils/strings'

export const setActivities = (details: ActivitiesProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activitiesContract = state.chainContracts.data.contracts.activities

    let activitiesRef = details.activitiesRef
    if ( activitiesRef == "" ) {
      activitiesRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const organisations: IATIActivitiesProps = {
      version: ethers.utils.formatBytes32String(details.version),
      generatedTime: ethers.utils.formatBytes32String(new Date().toISOString()),
      linkedData: ethers.utils.formatBytes32String(details.linkedData),
    }

    let actionType = IATIWriterActionTypes.ACTIVITIES_FAILURE
    let txData: TxReport = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await activitiesContract.setActivities(activitiesRef, organisations)
      const key = tx.hash
      txData = {
        [key]: {
          summary: `${Transaction.success}`,
          info: tx
        }
      }
      actionType = IATIWriterActionTypes.ACTIVITIES_SUCCESS
    } catch (error) {
      txData = {
        [-1]: {
          summary: `${Transaction.fail}`,
          info: {}
        }
      }
      console.log('set error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
