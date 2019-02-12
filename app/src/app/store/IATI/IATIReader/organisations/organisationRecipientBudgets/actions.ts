import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIOrganisationRecipientBudgetProps } from '../../../types'
import { IATIReportActionTypes,
         IATIRecipientOrgBudgetReportProps,
         OrganisationsReportProps } from '../types'

import { read } from '../actions'

export const getRecipientBudgets = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const recipientBudgetsContract = state.chainContracts.data.contracts.organisationRecipientBudgets
    const organisationsRef = props.organisationsRef
    const organisationRef = props.organisationRef

    const budgetReports = state.organisationsReport.data[organisationsRef].data[organisationRef].data.recipientOrgBudget as IATIRecipientOrgBudgetReportProps

    let actionType = IATIReportActionTypes.RECIPIENTORGBUDGET_FAILURE
    try {
      const num = await recipientBudgetsContract.getNumRecipientBudgets(props.organisationsRef, props.organisationRef)
      const numBudgets = num.toNumber()
      for (let i = 0; i < numBudgets; i++) {
         const budgetRef = await recipientBudgetsContract.getRecipientBudgetReference(organisationsRef,
                                                                                      organisationRef,
                                                                                      i.toString())
         const budget: IATIOrganisationRecipientBudgetProps = await recipientBudgetsContract.getRecipientBudget(organisationsRef,
                                                                                                                organisationRef,
                                                                                                                budgetRef)
         budgetReports[budgetRef] = budget
         actionType = IATIReportActionTypes.RECIPIENTORGBUDGET_SUCCESS
      }
    } catch (error) {
      console.log('getRecipientBudgets error', error)
    }

    dispatch(read({data: budgetReports})(actionType))

  }
}
