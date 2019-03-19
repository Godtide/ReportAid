import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../../actions'

import { ActionProps, PayloadProps, TxProps, TxReport } from '../../../../types'
import { OrganisationExpenditureProps, IATIWriterActionTypes, IATIBudgetProps } from '../../../types'

import { Transaction } from '../../../../../utils/strings'

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

    const expenditure: IATIBudgetProps = {
      budgetType: 1,
      budgetLine: ethers.utils.formatBytes32String(details.expenditureLine),
      otherRef: ethers.utils.formatBytes32String(""),
      finance: {
        status: details.status,
        value: details.value,
        start: ethers.utils.formatBytes32String(start.toISOString()),
        end: ethers.utils.formatBytes32String(end.toISOString())
      }
    }

    //console.log('Budget: ', orgBudget, ' Contract ', orgBudgetsContract)
    let actionType = IATIWriterActionTypes.TOTALEXPENDITURE_FAILURE
    let txData: TxReport = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await expenditureContract.setExpenditure(details.organisationsRef,
                                                             details.organisationRef,
                                                             expenditureRef,
                                                             expenditure)
      const key = tx.hash
      txData = {
        [key]: {
          summary: `${Transaction.success}`,
          info: tx
        }
      }
      actionType = IATIWriterActionTypes.TOTALEXPENDITURE_SUCCESS
    } catch (error) {
      txData = {
        [-1]: {
          summary: `${Transaction.fail}`,
          info: {}
        }
      }
      console.log('setBudget error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
