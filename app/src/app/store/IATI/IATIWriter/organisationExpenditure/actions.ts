import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'

import { write } from '../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'
import { OrganisationExpenditureProps, IATIOrganisationExpenditureProps } from '../../types'
import { IATIWriterActionTypes } from '../types'

export const setOrganisationExpenditure = (details: OrganisationExpenditureProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const expenditureContract = state.chainContracts.data.contracts.organisationExpenditure

    const start = new Date(details.startYear + '/' + details.startMonth + '/' + details.startDay)
    const end = new Date(details.endYear + '/' + details.endMonth + '/' + details.endDay)

    let expenditureRef = details.expenditureRef
    if ( expenditureRef == "" ) {
      expenditureRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const expenditure: IATIOrganisationExpenditureProps = {
      expenditureLine: ethers.utils.formatBytes32String(expenditureDetails.expenditureLine),
      finance: {
        value: expenditureDetails.value,
        status: expenditureDetails.status,
        start: ethers.utils.formatBytes32String(start.toISOString()),
        end: ethers.utils.formatBytes32String(end.toISOString())
      }
    }

    //console.log('Budget: ', orgBudget, ' Contract ', orgBudgetsContract)
    let actionType = IATIWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await expenditureContract.setExpenditure(details.organisationsRef,
                                                             details.organisationRef,
                                                             expenditureRef,
                                                             expenditure)
      const key = tx.hash
      txData[key] = tx
      actionType = IATIWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      txData[-1] = txData
      console.log('setBudget error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
