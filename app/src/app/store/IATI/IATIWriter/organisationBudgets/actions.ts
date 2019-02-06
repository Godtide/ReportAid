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

export const setOrganisationBudget = (budgetDetails: OrganisationBudgetProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()

    const start = new Date(budgetDetails.startYear + '/' + budgetDetails.startMonth + '/' + budgetDetails.startDay)
    const startDate = start.toISOString()
    const end = new Date(budgetDetails.endYear + '/' + budgetDetails.endMonth + '/' + budgetDetails.endDay)
    const endDate = end.toISOString()
    //console.log('Start: ', startDate, ' End: ', endDate)

    const orgBudget: IATIOrganisationBudgetProps = {
      report: {
        reportRef: budgetDetails.report.reportRef,
        orgRef: budgetDetails.report.orgRef
      },
      budgetRef: ethers.utils.formatBytes32String(shortid.generate()),
      budgetLine: ethers.utils.formatBytes32String(budgetDetails.budgetLine),
      finance: {
        value: budgetDetails.value,
        status: budgetDetails.status,
        start: ethers.utils.formatBytes32String(startDate),
        end: ethers.utils.formatBytes32String(endDate)
      }
    }

    const orgBudgetsContract = state.chainContracts.data.contracts.orgBudgetsContract
    //console.log('Budget: ', orgBudget, ' Contract ', orgBudgetsContract)
    let actionType = OrganisationBudgetsWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await orgBudgetsContract.setTotalBudget(orgBudget)
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
