import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'
import { OrgReportBudgetProps, IATIOrgReportBudgetProps } from '../../types'
import { OrgReportBudgetsWriterActionTypes } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrgReportBudgetsWriterActionTypes): TxProps => {
    const writerProps = storeAction(actionType)(payload) as TxProps
    return writerProps
  }
}

export const setOrganisationReportBudget = (reportDetails: OrgReportBudgetProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const date = new Date()
    const dateTime = date.toISOString()

    const orgBudget: IATIOrgReportBudgetProps = {
      reportRef: reportDetails.reportRef,
      budgetRef: ethers.utils.formatBytes32String(shortid.generate()),
      budgetLine: ethers.utils.formatBytes32String(reportDetails.budgetLine),
      finance: {
        value: reportDetails.value,
        status: ethers.utils.formatBytes32String(reportDetails.status),
        start: reportDetails.start,
        end: reportDetails.end
      }
    }

    const orgReportBudgetsContract = state.chainContracts.data.contracts.orgReportBudgetsContract
    let actionType = OrgReportBudgetsWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      // setReport(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await orgReportBudgetsContract.setTotalBudget(orgBudget)
      const key = tx.hash
      txData[key] = tx
      actionType = OrgReportBudgetsWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      console.log('setBudget error', error)
    }

    dispatch(add({data: {data: txData}})(actionType))
  }
}
