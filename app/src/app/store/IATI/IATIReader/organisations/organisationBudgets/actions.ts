import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIOrganisationBudgetProps } from '../../../types'
import { IATIReportActionTypes,
         IATIBudgetReportProps,
         OrganisationsReportProps } from '../types'

import { read } from '../actions'

export const getBudgets = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const budgetsContract = state.chainContracts.data.contracts.organisationBudgets
    const organisationsRef = props.organisationsRef
    const organisationRef = props.organisationRef

    const budgetReports = state.organisationsReport.data[organisationsRef].data[organisationRef].data.totalBudget as IATIBudgetReportProps

    let actionType = IATIReportActionTypes.BUDGET_FAILURE
    try {
      const num = await budgetsContract.getNumBudgets(props.organisationsRef, props.organisationRef)
      const numBudgets = num.toNumber()
      for (let i = 0; i < numBudgets; i++) {
         const budgetRef = await budgetsContract.getBudgetReference(organisationsRef,
                                                                    organisationRef,
                                                                    i.toString())
         const budget: IATIOrganisationBudgetProps = await budgetsContract.getBudget(organisationsRef,
                                                                               organisationRef,
                                                                               budgetRef)
         budgetReports[budgetRef] = budget
         actionType = IATIReportActionTypes.BUDGET_SUCCESS
      }
    } catch (error) {
      console.log('getBudgets error', error)
    }

    dispatch(read({data: budgetReports})(actionType))
  }
}
