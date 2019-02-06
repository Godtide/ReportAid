import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'
import { OrganisationRecipientBudgetProps, IATIOrganisationRecipientBudgetProps } from '../../types'
import { OrganisationRecipientBudgetsWriterActionTypes } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrganisationRecipientBudgetsWriterActionTypes): TxProps => {
    const writerProps = storeAction(actionType)(payload) as TxProps
    return writerProps
  }
}

export const setRecipientBudget = (budgetDetails: OrganisationRecipientBudgetProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()

    const start = new Date(budgetDetails.startYear + '/' + budgetDetails.startMonth + '/' + budgetDetails.startDay)
    const startDate = start.toISOString()
    const end = new Date(budgetDetails.endYear + '/' + budgetDetails.endMonth + '/' + budgetDetails.endDay)
    const endDate = end.toISOString()
    //console.log('Start: ', startDate, ' End: ', endDate)

    const recipientBudget: IATIOrganisationRecipientBudgetProps = {
      report: {
        reportRef: budgetDetails.report.reportRef,
        orgRef: budgetDetails.report.orgRef
      },
      budgetRef: ethers.utils.formatBytes32String(shortid.generate()),
      orgRef: budgetDetails.recipientOrgRef,
      budgetLine: ethers.utils.formatBytes32String(budgetDetails.budgetLine),
      finance: {
        value: budgetDetails.value,
        status: budgetDetails.status,
        start: ethers.utils.formatBytes32String(startDate),
        end: ethers.utils.formatBytes32String(endDate)
      }
    }

    const orgRecipientBudgetsContract = state.chainContracts.data.contracts.orgRecipientBudgetsContract
    //console.log('RecipientBudget: ', orgRecipientBudget, ' Contract ', orgRecipientBudgetsContract)
    let actionType = OrganisationRecipientBudgetsWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await orgRecipientBudgetsContract.setRecipientBudget(recipientBudget)
      const key = tx.hash
      txData[key] = tx
      actionType = OrganisationRecipientBudgetsWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      txData[-1] = txData
      console.log('setRecipientBudget error', error)
    }

    dispatch(add({data: {data: txData}})(actionType))
  }
}
