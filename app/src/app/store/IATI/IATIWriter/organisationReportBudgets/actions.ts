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

export const setOrganisationReportBudget = (budgetDetails: OrgReportBudgetProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()

    const start = new Date(budgetDetails.startYear + '/' + budgetDetails.startMonth + '/' + budgetDetails.startDay)
    const startDate = start.toISOString()
    const end = new Date(budgetDetails.endYear + '/' + budgetDetails.endMonth + '/' + budgetDetails.endDay)
    const endDate = end.toISOString()
    //console.log('Start: ', startDate, ' End: ', endDate)

    const orgBudget: IATIOrgReportBudgetProps = {
      reportRef: budgetDetails.reportRef,
      budgetRef: ethers.utils.formatBytes32String(shortid.generate()),
      budgetLine: ethers.utils.formatBytes32String(budgetDetails.budgetLine),
      finance: {
        value: budgetDetails.value,
        status: ethers.utils.formatBytes32String(budgetDetails.status),
        start: ethers.utils.formatBytes32String(startDate),
        end: ethers.utils.formatBytes32String(endDate)
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
