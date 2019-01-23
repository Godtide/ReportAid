import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'
import { OrgReportExpenditureProps, IATIOrgReportExpenditureProps } from '../../types'
import { OrgReportExpenditureWriterActionTypes } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrgReportExpenditureWriterActionTypes): TxProps => {
    const writerProps = storeAction(actionType)(payload) as TxProps
    return writerProps
  }
}

export const setOrganisationReportExpenditure = (expenditureDetails: OrgReportExpenditureProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()

    const start = new Date(expenditureDetails.startYear + '/' + expenditureDetails.startMonth + '/' + expenditureDetails.startDay)
    const startDate = start.toISOString()
    const end = new Date(expenditureDetails.endYear + '/' + expenditureDetails.endMonth + '/' + expenditureDetails.endDay)
    const endDate = end.toISOString()
    //console.log('Start: ', startDate, ' End: ', endDate)

    const orgExpenditure: IATIOrgReportExpenditureProps = {
      reportRef: expenditureDetails.reportRef,
      expenditureRef: ethers.utils.formatBytes32String(shortid.generate()),
      expenditureLine: ethers.utils.formatBytes32String(expenditureDetails.expenditureLine),
      finance: {
        value: expenditureDetails.value,
        status: expenditureDetails.status,
        start: ethers.utils.formatBytes32String(startDate),
        end: ethers.utils.formatBytes32String(endDate)
      }
    }

    const orgReportExpenditureContract = state.chainContracts.data.contracts.orgReportExpenditureContract
    //console.log('Budget: ', orgBudget, ' Contract ', orgReportBudgetsContract)
    let actionType = OrgReportExpenditureWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      // setReport(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await orgReportExpenditureContract.setExpenditure(orgExpenditure)
      const key = tx.hash
      txData[key] = tx
      actionType = OrgReportExpenditureWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      console.log('setBudget error', error)
    }

    dispatch(add({data: {data: txData}})(actionType))
  }
}
