import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIBudgetProps, OrganisationsReportProps, IATIReportActionTypes, IATIBudgetReportProps } from '../../../types'

import { read } from '../../actions'

export const getRecipientBudgets = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const recipientBudgetsContract = state.chainContracts.data.contracts.organisationRecipientBudgets
    const organisationsRef = props.organisationsRef
    const organisationRef = props.organisationRef

    let budgetReports: IATIBudgetReportProps = {
      data: { organisationsRef: organisationsRef,
              organisationRef: organisationRef,
              data: []
            }
    }

    let actionType = IATIReportActionTypes.RECIPIENTORGBUDGET_FAILURE
    try {
      const num = await recipientBudgetsContract.getNumRecipientBudgets(props.organisationsRef, props.organisationRef)
      const numBudgets = num.toNumber()
      for (let i = 0; i < numBudgets; i++) {
         const budgetRef = await recipientBudgetsContract.getRecipientBudgetReference(organisationsRef,
                                                                                      organisationRef,
                                                                                      i.toString())
         const budget: IATIBudgetProps = await recipientBudgetsContract.getRecipientBudget(organisationsRef,
                                                                                           organisationRef, budgetRef)

         budgetReports.data.data[i] = {
           budgetKey: budgetRef,
           budgetLine: ethers.utils.parseBytes32String(budget.budgetLine),
           recipientOrgRef: budget.otherRef,
           value: ethers.utils.bigNumberify(budget.finance.value).toNumber(),
           status: budget.finance.status,
           start: ethers.utils.parseBytes32String(budget.finance.start),
           end: ethers.utils.parseBytes32String(budget.finance.end)
         }

         actionType = IATIReportActionTypes.RECIPIENTORGBUDGET_SUCCESS
      }
    } catch (error) {
      console.log('getRecipientBudgets error', error)
    }

    dispatch(read({data: budgetReports})(actionType))

  }
}
