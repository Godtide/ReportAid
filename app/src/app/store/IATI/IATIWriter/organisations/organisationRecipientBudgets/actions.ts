import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../../types'
import { OrganisationRecipientBudgetProps, IATIOrganisationRecipientBudgetProps } from '../../../types'
import { IATIWriterActionTypes } from '../types'

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

    const recipientBudget: IATIOrganisationRecipientBudgetProps = {
      recipientOrgRef: details.recipientOrgRef,
      budgetLine: ethers.utils.formatBytes32String(details.budgetLine),
      finance: {
        value: details.value,
        status: details.status,
        start: ethers.utils.formatBytes32String(start.toISOString()),
        end: ethers.utils.formatBytes32String(end.toISOString())
      }
    }
    //console.log('RecipientBudget: ', orgRecipientBudget, ' Contract ', orgRecipientBudgetsContract)
    let actionType = IATIWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await recipientBudgetsContract.setRecipientBudget(details.organisationsRef,
                                                                   details.organisationRef,
                                                                   budgetRef,
                                                                   recipientBudget)
      const key = tx.hash
      txData[key] = tx
      actionType = IATIWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      txData[-1] = txData
      console.log('setRecipientBudget error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
