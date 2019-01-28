import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'
import { OrgReportRegionBudgetProps, IATIOrgReportRegionBudgetProps } from '../../types'
import { OrgReportRegionBudgetsWriterActionTypes } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrgReportRegionBudgetsWriterActionTypes): TxProps => {
    const writerProps = storeAction(actionType)(payload) as TxProps
    return writerProps
  }
}

export const setRegionBudget = (budgetDetails: OrgReportRegionBudgetProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()

    const start = new Date(budgetDetails.startYear + '/' + budgetDetails.startMonth + '/' + budgetDetails.startDay)
    const startDate = start.toISOString()
    const end = new Date(budgetDetails.endYear + '/' + budgetDetails.endMonth + '/' + budgetDetails.endDay)
    const endDate = end.toISOString()
    //console.log('Start: ', startDate, ' End: ', endDate)

    const regionBudget: IATIOrgReportRegionBudgetProps = {
      reportRef: budgetDetails.reportRef,
      budgetRef: ethers.utils.formatBytes32String(shortid.generate()),
      regionRef: budgetDetails.regionRef,
      budgetLine: ethers.utils.formatBytes32String(budgetDetails.budgetLine),
      finance: {
        value: budgetDetails.value,
        status: budgetDetails.status,
        start: ethers.utils.formatBytes32String(startDate),
        end: ethers.utils.formatBytes32String(endDate)
      }
    }

    const orgReportRegionBudgetsContract = state.chainContracts.data.contracts.orgReportRegionBudgetsContract
    //console.log('RegionBudget: ', regionBudget, ' Contract ', orgReportRegionBudgetsContract)
    let actionType = OrgReportRegionBudgetsWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      // setReport(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await orgReportRegionBudgetsContract.setRegionBudget(regionBudget)
      const key = tx.hash
      txData[key] = tx
      actionType = OrgReportRegionBudgetsWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      txData[-1] = txData
      console.log('setRegionBudget error', error)
    }

    dispatch(add({data: {data: txData}})(actionType))
  }
}
