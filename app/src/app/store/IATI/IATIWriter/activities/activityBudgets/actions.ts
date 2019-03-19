import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../../actions'

import { ActionProps, TxReport } from '../../../../types'
import { ActivityBudgetProps, IATIWriterActionTypes, IATIBudgetProps } from '../../../types'

import { Transaction } from '../../../../../utils/strings'

export const setActivityBudget = (details: ActivityBudgetProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const budgetsContract = state.chainContracts.data.contracts.activityBudgets

    let budgetRef = details.budgetRef
    if ( budgetRef == "" ) {
      budgetRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const start = new Date(details.startYear + '/' + details.startMonth + '/' + details.startDay)
    const end = new Date(details.endYear + '/' + details.endMonth + '/' + details.endDay)

    const budget: IATIBudgetProps = {
      budgetType: details.budgetType,
      budgetLine: ethers.utils.formatBytes32String(""),
      otherRef: ethers.utils.formatBytes32String(""),
      finance: {
        status: details.status,
        value: details.value,
        start: ethers.utils.formatBytes32String(start.toISOString()),
        end: ethers.utils.formatBytes32String(end.toISOString())
      }
    }

    //console.log('Budget: ', budget)

    let actionType = IATIWriterActionTypes.BUDGET_FAILURE
    let txData: TxReport = {}
    try {
      const tx = await budgetsContract.setBudget(details.activitiesRef,
                                                 details.activityRef,
                                                 budgetRef,
                                                 budget)
      const key = tx.hash
      txData = {
        [key]: {
          summary: `${Transaction.success}`,
          info: tx
        }
      }
      actionType = IATIWriterActionTypes.BUDGET_SUCCESS
    } catch (error) {
      txData = {
        [-1]: {
          summary: `${Transaction.fail}`,
          info: {}
        }
      }
      console.log('setBudget error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
