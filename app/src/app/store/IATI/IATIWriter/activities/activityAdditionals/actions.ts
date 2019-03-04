import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../../actions'

import { ActionProps, TxReport } from '../../../../types'
import { ActivityAdditionalProps, IATIActivityAdditionalProps, IATIWriterActionTypes } from '../../../types'

import { Transaction } from '../../../../../utils/strings'

export const setActivityAdditional = (details: ActivityAdditionalProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityAdditionalContract = state.chainContracts.data.contracts.activityAdditional

    let additionalRef = details.additionalRef
    if ( additionalRef == "" ) {
      additionalRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const additional: IATIActivityAdditionalProps = {
      defaultAidType: ethers.utils.formatBytes32String(details.defaultAidType),
      defaultFinanceType: details.defaultFinanceType,
      budgetNotProvided: details.budgetNotProvided,
      status: details.status,
      scope: details.scope,
      capitalSpend: details.capitalSpend,
      collaborationType: details.collaborationType,
      defaultFlowType: details.defaultFlowType,
      defaultTiedStatus: details.defaultTiedStatus
    }

    let actionType = IATIWriterActionTypes.ACTIVITYADDITIONAL_FAILURE
    let txData: TxReport = {}
    try {

      /*let gasLimit = await state.chainInfo.data.provider.estimateGas(activityAdditionalContract.setActivityAdditional)
      gasLimit *= 90
      console.log("Additional: ", details.activitiesRef, details.activityRef, additionalRef, additional)
      console.log("Gas limit: ", gasLimit)
      const tx = await activityAdditionalContract.setActivityAdditional(details.activitiesRef, details.activityRef, additionalRef, additional, {gasLimit: gasLimit})*/
      const tx = await activityAdditionalContract.setAdditional(details.activitiesRef, details.activityRef, additionalRef, additional)
      const key = tx.hash
      txData = {
        [key]: {
          summary: `${Transaction.success}`,
          info: tx
        }
      }

      actionType = IATIWriterActionTypes.ACTIVITYADDITIONAL_SUCCESS
    } catch (error) {
      txData = {
        [-1]: {
          summary: `${Transaction.fail}`,
          info: {}
        }
      }
      console.log('setActivityAdditional error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
