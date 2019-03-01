import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../../actions'

import { ActionProps, TxReport } from '../../../../types'
import { ActivityProps, IATIActivityProps } from '../types'
import { IATIWriterActionTypes } from '../../types'

import { Transaction } from '../../../../../utils/strings'

export const setActivity = (details: ActivityProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activitiesContract = state.chainContracts.data.contracts.activities

    let activityRef = details.activityRef
    if ( activityRef == "" ) {
      activityRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const activity: IATIActivityProps = {
      identifier: details.identifier,
      reportingOrg: {
        orgRef: details.reportingOrgRef,
        orgType: details.reportingOrgType,
        isSecondary: details.reportingOrgIsSecondary
      },
      title: details.title,
      lastUpdated: ethers.utils.formatBytes32String(new Date().toISOString()),
      lang: ethers.utils.formatBytes32String(details.lang),
      currency: ethers.utils.formatBytes32String(details.currency),
      humanitarian: details.humanitarian,
      hierarchy: details.hierarchy,
      linkedData: ethers.utils.formatBytes32String(details.linkedData),
      budgetNotProvided: details.budgetNotProvided,
      status: details.status,
      scope: details.scope,
      capitalSpend: details.capitalSpend,
      collaborationType: details.collaborationType,
      defaultFlowType: details.defaultFlowType,
      defaultFinanceType: details.defaultFinanceType,
      defaultAidType: string,
      defaulTiedStatus: details.defaulTiedStatus,
    }

    let actionType = IATIWriterActionTypes.ACTIVITY_FAILURE
    let txData: TxReport = {}
    try {
      //console.log("Setting activities: ", activitiesRef, activitiesRef.length, activities)
      const tx = await activitiesContract.setActivity(details.activitiesRef, activityRef, activity)
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
