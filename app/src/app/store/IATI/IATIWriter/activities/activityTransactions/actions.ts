import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../../../../actions'

import { ActionProps, TxReport } from '../../../../types'
import { ActivityTransactionProps, IATIActivityTransactionProps, IATIWriterActionTypes } from '../../../types'

import { Transaction } from '../../../../../utils/strings'

export const setActivityTransaction = (details: ActivityTransactionProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityTransactionsContract = state.chainContracts.data.contracts.activityTransactions

    let transactionRef = details.transactionRef
    if ( transactionRef == "" ) {
      transactionRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    // one or the other, not both
    let territory = details.countryCode
    if ( details.regionCode != "" ) {
      territory = details.regionCode
    }

    let providerActivityRef = details.providerActivityRef
    if ( providerActivityRef == "" ) {
      providerActivityRef = ethers.utils.formatBytes32String("")
    }

    let receiverActivityRef = details.receiverActivityRef
    if ( receiverActivityRef == "" ) {
      receiverActivityRef = ethers.utils.formatBytes32String("")
    }

    const thisDate = new Date(details.year + '/' + details.month + '/' + details.day)
    const valueDate = new Date(details.valueYear + '/' + details.valueMonth + '/' + details.valueDay)

    const transactionData: IATIActivityTransactionProps = {
      transactionType: details.transactionType,
      disbursementChannel: details.disbursementChannel,
      flowType: details.flowType,
      tiedStatus: details.tiedStatus,
      financeType: details.financeType,
      aidType: ethers.utils.formatBytes32String(details.aidType),
      date: ethers.utils.formatBytes32String(thisDate.toISOString()),
      value: {
        value: details.value,
        date: ethers.utils.formatBytes32String(valueDate.toISOString()),
        currency: ethers.utils.formatBytes32String(details.currency)
      },
      providerOrg: {
        orgType: details.providerOrgType,
        orgRef: details.providerOrgRef,
        activityRef: providerActivityRef,
      },
      receiverOrg: {
        orgType: details.providerOrgType,
        orgRef: details.providerOrgRef,
        activityRef: receiverActivityRef,
      },
      sectorDacCode: details.sectorDacCode,
      territory: ethers.utils.formatBytes32String(territory),
      description: details.description
    }

    //console.log('Activity transaction: ', details.activitiesRef, details.activityRef, transactionRef, transactionData)

    let actionType = IATIWriterActionTypes.ACTIVITYTRANSACTION_FAILURE
    let txData: TxReport = {}
    try {
      const tx = await activityTransactionsContract.setTransaction(details.activitiesRef, details.activityRef, transactionRef, transactionData)
      const key = tx.hash
      txData = {
        [key]: {
          summary: `${Transaction.success}`,
          info: tx
        }
      }

      actionType = IATIWriterActionTypes.ACTIVITYTRANSACTION_SUCCESS
    } catch (error) {
      txData = {
        [-1]: {
          summary: `${Transaction.fail}`,
          info: {}
        }
      }
      console.log('setActivityTransaction error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
