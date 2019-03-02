import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../../actions'

import { ActionProps, TxReport } from '../../../../types'
import { OrganisationBudgetProps, IATIWriterActionTypes, IATIBudgetProps } from '../../../types'

import { Transaction } from '../../../../../utils/strings'

export const setOrganisationBudget = (details: OrganisationBudgetProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const budgetsContract = state.chainContracts.data.contracts.organisationBudgets

    let budgetRef = details.budgetRef
    if ( budgetRef == "" ) {
      budgetRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const start = new Date(details.startYear + '/' + details.startMonth + '/' + details.startDay)
    const end = new Date(details.endYear + '/' + details.endMonth + '/' + details.endDay)

    const budget: IATIBudgetProps = {
      budgetType: 1,
      budgetLine: ethers.utils.formatBytes32String(details.budgetLine),
      otherRef: ethers.utils.formatBytes32String(""),
      finance: {
        value: details.value,
        status: details.status,
        start: ethers.utils.formatBytes32String(start.toISOString()),
        end: ethers.utils.formatBytes32String(end.toISOString())
      }
    }

    //console.log('Budget: ', budget)

    let actionType = IATIWriterActionTypes.BUDGET_FAILURE
    let txData: TxReport = {}
    try {
      const tx = await budgetsContract.setBudget(details.organisationsRef,
                                                 details.organisationRef,
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
