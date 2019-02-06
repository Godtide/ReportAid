import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'
import { OrganisationCountryBudgetProps, IATIOrganisationCountryBudgetProps } from '../../types'
import { OrganisationCountryBudgetsWriterActionTypes } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrganisationCountryBudgetsWriterActionTypes): TxProps => {
    const writerProps = storeAction(actionType)(payload) as TxProps
    return writerProps
  }
}

export const setCountryBudget = (budgetDetails: OrganisationCountryBudgetProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()

    const start = new Date(budgetDetails.startYear + '/' + budgetDetails.startMonth + '/' + budgetDetails.startDay)
    const startDate = start.toISOString()
    const end = new Date(budgetDetails.endYear + '/' + budgetDetails.endMonth + '/' + budgetDetails.endDay)
    const endDate = end.toISOString()
    //console.log('Start: ', startDate, ' End: ', endDate)

    const countryBudget: IATIOrganisationCountryBudgetProps = {
      report: {
        reportRef: budgetDetails.report.reportRef,
        orgRef: budgetDetails.report.orgRef
      },
      budgetRef: ethers.utils.formatBytes32String(shortid.generate()),
      countryRef: ethers.utils.formatBytes32String(budgetDetails.countryRef),
      budgetLine: ethers.utils.formatBytes32String(budgetDetails.budgetLine),
      finance: {
        value: budgetDetails.value,
        status: budgetDetails.status,
        start: ethers.utils.formatBytes32String(startDate),
        end: ethers.utils.formatBytes32String(endDate)
      }
    }

    const orgCountryBudgetsContract = state.chainContracts.data.contracts.orgCountryBudgetsContract
    //console.log('CountryBudget: ', countryBudget, ' Contract ', orgCountryBudgetsContract)
    let actionType = OrganisationCountryBudgetsWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await orgCountryBudgetsContract.setCountryBudget(countryBudget)
      const key = tx.hash
      txData[key] = tx
      actionType = OrganisationCountryBudgetsWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      txData[-1] = txData
      console.log('setCountryBudget error', error)
    }

    dispatch(add({data: {data: txData}})(actionType))
  }
}
