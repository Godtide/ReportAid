import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'
import { OrgExpenditureProps, IATIOrgExpenditureProps } from '../../types'
import { OrgExpenditureWriterActionTypes } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrgExpenditureWriterActionTypes): TxProps => {
    const writerProps = storeAction(actionType)(payload) as TxProps
    return writerProps
  }
}

export const setOrganisationExpenditure = (expenditureDetails: OrgExpenditureProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()

    const start = new Date(expenditureDetails.startYear + '/' + expenditureDetails.startMonth + '/' + expenditureDetails.startDay)
    const startDate = start.toISOString()
    const end = new Date(expenditureDetails.endYear + '/' + expenditureDetails.endMonth + '/' + expenditureDetails.endDay)
    const endDate = end.toISOString()
    //console.log('Start: ', startDate, ' End: ', endDate)

    const orgExpenditure: IATIOrgExpenditureProps = {
      report: {
        reportRef: expenditureDetails.report.reportRef,
        orgRef: expenditureDetails.report.orgRef
      },
      expenditureRef: ethers.utils.formatBytes32String(shortid.generate()),
      expenditureLine: ethers.utils.formatBytes32String(expenditureDetails.expenditureLine),
      finance: {
        value: expenditureDetails.value,
        status: expenditureDetails.status,
        start: ethers.utils.formatBytes32String(startDate),
        end: ethers.utils.formatBytes32String(endDate)
      }
    }

    const orgExpenditureContract = state.chainContracts.data.contracts.orgExpenditureContract
    //console.log('Budget: ', orgBudget, ' Contract ', orgBudgetsContract)
    let actionType = OrgExpenditureWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await orgExpenditureContract.setExpenditure(orgExpenditure)
      const key = tx.hash
      txData[key] = tx
      actionType = OrgExpenditureWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      txData[-1] = txData
      console.log('setBudget error', error)
    }

    dispatch(add({data: {data: txData}})(actionType))
  }
}
