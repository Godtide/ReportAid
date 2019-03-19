import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../../actions'

import { ActionProps, PayloadProps, TxProps, TxReport } from '../../../../types'
import { OrganisationRegionBudgetProps, IATIWriterActionTypes, IATIBudgetProps } from '../../../types'

import { Transaction } from '../../../../../utils/strings'

export const setRegionBudget = (details: OrganisationRegionBudgetProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const regionBudgetsContract = state.chainContracts.data.contracts.organisationRegionBudgets

    const start = new Date(details.startYear + '/' + details.startMonth + '/' + details.startDay)
    const end = new Date(details.endYear + '/' + details.endMonth + '/' + details.endDay)

    let budgetRef = details.budgetRef
    if ( budgetRef == "" ) {
      budgetRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const regionBudget: IATIBudgetProps = {
      budgetType: 1,
      budgetLine: ethers.utils.formatBytes32String(details.budgetLine),
      otherRef: ethers.utils.formatBytes32String(details.regionRef),
      finance: {
        status: details.status,
        value: details.value,
        start: ethers.utils.formatBytes32String(start.toISOString()),
        end: ethers.utils.formatBytes32String(end.toISOString())
      }
    }

    //console.log('RegionBudget: ', regionBudget, ' Contract ', orgRegionBudgetsContract)
    let actionType = IATIWriterActionTypes.RECIPIENTREGIONBUDGET_FAILURE
    let txData: TxReport = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await regionBudgetsContract.setRegionBudget(details.organisationsRef,
                                                             details.organisationRef,
                                                             budgetRef,
                                                             regionBudget)
      const key = tx.hash
      txData = {
        [key]: {
          summary: `${Transaction.success}`,
          info: tx
        }
      }
      actionType = IATIWriterActionTypes.RECIPIENTREGIONBUDGET_SUCCESS
    } catch (error) {
      txData = {
        [-1]: {
          summary: `${Transaction.fail}`,
          info: {}
        }
      }
      console.log('setRegionBudget error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
