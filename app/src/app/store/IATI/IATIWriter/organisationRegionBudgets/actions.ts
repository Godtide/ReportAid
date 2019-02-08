import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'

import { write } from '../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'
import { OrganisationRegionBudgetProps, IATIOrganisationRegionBudgetProps } from '../../types'
import { IATIWriterActionTypes } from '../types'

export const setRegionBudget = (budgetDetails: OrganisationRegionBudgetProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const regionBudgetsContract = state.chainContracts.data.contracts.organisationRegionBudgets

    const start = new Date(budgetDetails.startYear + '/' + budgetDetails.startMonth + '/' + budgetDetails.startDay)
    const end = new Date(budgetDetails.endYear + '/' + budgetDetails.endMonth + '/' + budgetDetails.endDay)

    let budgetRef = details.budgetRef
    if ( budgetRef == "" ) {
      budgetRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const regionBudget: IATIOrganisationRegionBudgetProps = {
      regionRef: budgetDetails.regionRef,
      budgetLine: ethers.utils.formatBytes32String(budgetDetails.budgetLine),
      finance: {
        value: budgetDetails.value,
        status: budgetDetails.status,
        start: ethers.utils.formatBytes32String(start.toISOString()),
        end: ethers.utils.formatBytes32String(end.toISOString())
      }
    }

    //console.log('RegionBudget: ', regionBudget, ' Contract ', orgRegionBudgetsContract)
    let actionType = IATIWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await regionBudgetsContract.setRegionBudget(details.organisationsRef,
                                                             details.organisationRef,
                                                             budgetRef,
                                                             regionBudget)
      const key = tx.hash
      txData[key] = tx
      actionType = IATIWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      txData[-1] = txData
      console.log('setRegionBudget error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
