import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../../actions'

import { ActionProps, TxReport } from '../../../../types'
import { ActivityDateProps, IATIActivityDateProps, IATIWriterActionTypes } from '../../../types'

import { Transaction } from '../../../../../utils/strings'

export const setActivityDate = (details: ActivityDateProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityDatesContract = state.chainContracts.data.contracts.activityDates

    let dateRef = details.dateRef
    if ( dateRef == "" ) {
      dateRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const thisDate = new Date(details.year + '/' + details.month + '/' + details.day)

    //console.log(details, dateRef, thisDate)

    const dateData: IATIActivityDateProps = {
      dateType: details.dateType,
      date: ethers.utils.formatBytes32String(thisDate.toISOString()),
      narrative: details.narrative
    }

    let actionType = IATIWriterActionTypes.ACTIVITYDATE_FAILURE
    let txData: TxReport = {}
    try {
      const tx = await activityDatesContract.setDate(details.activitiesRef, details.activityRef, dateRef, dateData)
      const key = tx.hash
      txData = {
        [key]: {
          summary: `${Transaction.success}`,
          info: tx
        }
      }

      actionType = IATIWriterActionTypes.ACTIVITYDATE_SUCCESS
    } catch (error) {
      txData = {
        [-1]: {
          summary: `${Transaction.fail}`,
          info: {}
        }
      }
      console.log('setActivityDate error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
