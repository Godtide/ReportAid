import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../../actions'

import { ActionProps, TxReport } from '../../../../types'
import { ActivityProps, IATIActivityProps, IATIWriterActionTypes } from '../../../types'

import { Transaction } from '../../../../../utils/strings'

export const setActivity = (details: ActivityProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityContract = state.chainContracts.data.contracts.activity

    let activityRef = details.activityRef
    if ( activityRef == "" ) {
      activityRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const activity: IATIActivityProps = {
      humanitarian: details.humanitarian,
      hierarchy: details.hierarchy,
      status: details.status,
      budgetNotProvided: details.budgetNotProvided,
      reportingOrg: {
        orgRef: details.reportingOrgRef,
        orgType: details.reportingOrgType,
        isSecondary: details.reportingOrgIsSecondary
      },
      lastUpdated: ethers.utils.formatBytes32String(new Date().toISOString()),
      lang: ethers.utils.formatBytes32String(details.lang),
      currency: ethers.utils.formatBytes32String(details.currency),
      linkedData: ethers.utils.formatBytes32String(details.linkedData),
      identifier: ethers.utils.formatBytes32String(details.identifier),
      title: ethers.utils.formatBytes32String(details.title),
      description: details.description
    }

    let actionType = IATIWriterActionTypes.ACTIVITY_FAILURE
    let txData: TxReport = {}
    try {
      //console.log("Setting activity: ", details.activitiesRef, activityRef, activity)
      const tx = await activityContract.setActivity(details.activitiesRef, activityRef, activity)
      const key = tx.hash
      txData = {
        [key]: {
          summary: `${Transaction.success}`,
          info: tx
        }
      }

      actionType = IATIWriterActionTypes.ACTIVITY_SUCCESS
    } catch (error) {
      txData = {
        [-1]: {
          summary: `${Transaction.fail}`,
          info: {}
        }
      }
      console.log('set activity error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
