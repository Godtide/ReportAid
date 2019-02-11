import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../actions'

import { ActionProps, TxData } from '../../../../types'
import { OrganisationBudgetProps, IATIOrganisationBudgetProps } from '../../../types'
import { IATIWriterActionTypes } from '../types'

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

    const budget: IATIOrganisationBudgetProps = {
      budgetLine: ethers.utils.formatBytes32String(details.budgetLine),
      finance: {
        value: details.value,
        status: details.status,
        start: ethers.utils.formatBytes32String(start.toISOString()),
        end: ethers.utils.formatBytes32String(end.toISOString())
      }
    }

    let actionType = IATIWriterActionTypes.BUDGET_FAILURE
    let txData: TxData = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await budgetsContract.setBudget(details.organisationsRef,
                                                 details.organisationRef,
                                                 budgetRef,
                                                 budget)
      const key = tx.hash
      txData[key] = tx
      actionType = IATIWriterActionTypes.BUDGET_SUCCESS
    } catch (error) {
      txData[-1] = txData
      console.log('setBudget error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
