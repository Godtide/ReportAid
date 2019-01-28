import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'
import { OrgReportRecipientBudgetProps, IATIOrgReportRecipientBudgetProps } from '../../types'
import { OrgReportRecipientBudgetsWriterActionTypes } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrgReportRecipientBudgetsWriterActionTypes): TxProps => {
    const writerProps = storeAction(actionType)(payload) as TxProps
    return writerProps
  }
}

export const setRecipientBudget = (budgetDetails: OrgReportRecipientBudgetProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()

    const start = new Date(budgetDetails.startYear + '/' + budgetDetails.startMonth + '/' + budgetDetails.startDay)
    const startDate = start.toISOString()
    const end = new Date(budgetDetails.endYear + '/' + budgetDetails.endMonth + '/' + budgetDetails.endDay)
    const endDate = end.toISOString()
    //console.log('Start: ', startDate, ' End: ', endDate)

    const recipientBudget: IATIOrgReportRecipientBudgetProps = {
      reportRef: budgetDetails.reportRef,
      budgetRef: ethers.utils.formatBytes32String(shortid.generate()),
      orgRef: budgetDetails.orgRef,
      budgetLine: ethers.utils.formatBytes32String(budgetDetails.budgetLine),
      finance: {
        value: budgetDetails.value,
        status: budgetDetails.status,
        start: ethers.utils.formatBytes32String(startDate),
        end: ethers.utils.formatBytes32String(endDate)
      }
    }

    const orgReportRecipientBudgetsContract = state.chainContracts.data.contracts.orgReportRecipientBudgetsContract
    //console.log('RecipientBudget: ', orgRecipientBudget, ' Contract ', orgReportRecipientBudgetsContract)
    let actionType = OrgReportRecipientBudgetsWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      // setReport(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await orgReportRecipientBudgetsContract.setRecipientBudget(recipientBudget)
      const key = tx.hash
      txData[key] = tx
      actionType = OrgReportRecipientBudgetsWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      txData[-1] = txData
      console.log('setRecipientBudget error', error)
    }

    dispatch(add({data: {data: txData}})(actionType))
  }
}
