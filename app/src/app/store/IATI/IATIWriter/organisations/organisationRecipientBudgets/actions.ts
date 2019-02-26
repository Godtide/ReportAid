import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../../types'
import { OrganisationRecipientBudgetProps, IATIBudgetProps } from '../../../types'
import { IATIWriterActionTypes } from '../types'

import { Transaction } from '../../../../../utils/strings'

export const setRecipientBudget = (details: OrganisationRecipientBudgetProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const recipientBudgetsContract = state.chainContracts.data.contracts.organisationRecipientBudgets

    const start = new Date(details.startYear + '/' + details.startMonth + '/' + details.startDay)
    const end = new Date(details.endYear + '/' + details.endMonth + '/' + details.endDay)

    let budgetRef = details.budgetRef
    if ( budgetRef == "" ) {
      budgetRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const recipientBudget: IATIBudgetProps = {
      budgetType: 1,
      budgetLine: ethers.utils.formatBytes32String(details.budgetLine),
      otherRef: details.recipientOrgRef,
      finance: {
        value: details.value,
        status: details.status,
        start: ethers.utils.formatBytes32String(start.toISOString()),
        end: ethers.utils.formatBytes32String(end.toISOString())
      }
    }

    let actionType = IATIWriterActionTypes.RECIPIENTORGBUDGET_FAILURE
    let txData: TxData = {}
    try {

      const tx = await recipientBudgetsContract.setRecipientBudget(details.organisationsRef,
                                                                   details.organisationRef,
                                                                   budgetRef,
                                                                   recipientBudget)
      const key = tx.hash
      txData = {
        [key]: {
          summary: `${Transaction.success}`,
          info: tx
        }
      }
      actionType = IATIWriterActionTypes.RECIPIENTORGBUDGET_SUCCESS
    } catch (error) {
      txData = {
        [-1]: {
          summary: `${Transaction.fail}`,
          info: {}
        }
      }
      console.log('setRecipientBudget error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
