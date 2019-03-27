import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../../../../actions'

import { ActionProps, TxReport } from '../../../../types'
import { ActivityRelatedActivityProps, IATIActivityRelatedActivityProps, IATIWriterActionTypes } from '../../../types'

import { Transaction } from '../../../../../utils/strings'

export const setRelatedActivity = (details: ActivityRelatedActivityProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const  activityRelatedActivitiesContract = state.chainContracts.data.contracts. activityRelatedActivities

    let activityRelatedActivityRef = details.activityRelatedActivityRef
    if ( activityRelatedActivityRef == "" ) {
      activityRelatedActivityRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const relatedActivityData: IATIActivityRelatedActivityProps = {
      relationType: details.relationType,
      activityRef: details.relatedActivityRef
    }

    console.log('Activity Related: ', details.activitiesRef, details.activityRef, activityRelatedActivityRef, relatedActivityData)

    let actionType = IATIWriterActionTypes.ACTIVITYRELATEDACTIVITY_FAILURE
    let txData: TxReport = {}
    try {
      const tx = await  activityRelatedActivitiesContract.setRelatedActivity(details.activitiesRef, details.activityRef, activityRelatedActivityRef, relatedActivityData)
      const key = tx.hash
      txData = {
        [key]: {
          summary: `${Transaction.success}`,
          info: tx
        }
      }

      actionType = IATIWriterActionTypes.ACTIVITYRELATEDACTIVITY_SUCCESS
    } catch (error) {
      txData = {
        [-1]: {
          summary: `${Transaction.fail}`,
          info: {}
        }
      }
      console.log('setActivityRelatedActivity error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
