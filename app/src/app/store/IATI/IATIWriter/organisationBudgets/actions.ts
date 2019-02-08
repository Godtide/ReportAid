import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'
import { OrganisationBudgetProps, IATIOrganisationBudgetProps } from '../../types'
import { OrganisationBudgetsWriterActionTypes } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrganisationBudgetsWriterActionTypes): TxProps => {
    const writerProps = storeAction(actionType)(payload) as TxProps
    return writerProps
  }
}

export const setOrganisationBudget = (details: OrganisationBudgetProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const budgetsContract = state.chainContracts.data.contracts.organisationBudgets

    let budgetRef = details.budgetRef
    if ( budgetRef == "" ) {
      budgetRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const start = new Date(budgetDetails.startYear + '/' + budgetDetails.startMonth + '/' + budgetDetails.startDay)
    const end = new Date(budgetDetails.endYear + '/' + budgetDetails.endMonth + '/' + budgetDetails.endDay)

    const budget: IATIOrganisationBudgetProps = {
      budgetLine: ethers.utils.formatBytes32String(budgetDetails.budgetLine),
      finance: {
        value: budgetDetails.value,
        status: budgetDetails.status,
        start: ethers.utils.formatBytes32String(start.toISOString()),
        end: ethers.utils.formatBytes32String(end.toISOString())
      }
    }

    let actionType = OrganisationBudgetsWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await budgetsContract.setBudget(details.organisationsRef,
                                                 details.organisationRef,
                                                 budgetRef,
                                                 budget)
      const key = tx.hash
      txData[key] = tx
      actionType = OrganisationBudgetsWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      txData[-1] = txData
      console.log('setBudget error', error)
    }

    dispatch(add({data: {data: txData}})(actionType))
  }
}
