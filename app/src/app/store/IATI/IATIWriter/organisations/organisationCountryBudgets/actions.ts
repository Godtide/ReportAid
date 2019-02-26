import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../../types'
import { OrganisationCountryBudgetProps, IATIBudgetProps } from '../../../types'
import { IATIWriterActionTypes } from '../types'

import { Transaction } from '../../../../../utils/strings'

export const setCountryBudget = (details: OrganisationCountryBudgetProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const countryBudgetsContract = state.chainContracts.data.contracts.organisationCountryBudgets

    let budgetRef = details.budgetRef
    if ( budgetRef == "" ) {
      budgetRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const start = new Date(details.startYear + '/' + details.startMonth + '/' + details.startDay)
    const end = new Date(details.endYear + '/' + details.endMonth + '/' + details.endDay)

    const countryBudget: IATIBudgetProps = {
      budgetType: 1,
      budgetLine: ethers.utils.formatBytes32String(details.budgetLine),
      otherRef: ethers.utils.formatBytes32String(details.countryRef),
      finance: {
        value: details.value,
        status: details.status,
        start: ethers.utils.formatBytes32String(start.toISOString()),
        end: ethers.utils.formatBytes32String(end.toISOString())
      }
    }
    //console.log('CountryBudget: ', countryBudget, ' Contract ', orgCountryBudgetsContract)
    let actionType = IATIWriterActionTypes.RECIPIENTCOUNTRYBUDGET_FAILURE
    let txData: TxData = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await countryBudgetsContract.setCountryBudget(details.organisationsRef,
                                                               details.organisationRef,
                                                               budgetRef,
                                                               countryBudget)
      const key = tx.hash
      txData = {
        [key]: {
          summary: `${Transaction.success}`,
          info: tx
        }
      }
      actionType = IATIWriterActionTypes.RECIPIENTCOUNTRYBUDGET_SUCCESS
    } catch (error) {
      txData = {
        [-1]: {
          summary: `${Transaction.fail}`,
          info: {}
        }
      }
      console.log('setCountryBudget error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
