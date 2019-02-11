import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'

import { read } from '../actions'

import { IATIOrganisationsReportProps, IATIReportActionTypes, OrganisationsReportProps } from '../types'

export const getBudgets = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const budgetsContract = state.chainContracts.data.contracts.organisationBudgets
    const organisationsRef = props.organisationsRef
    const organisationRef = props.organisationRef

    const report = state.organisationsReport.data as IATIOrganisationsReportProps
    let actionType = IATIReportActionTypes.BUDGET_FAILURE
    try {
      const num = await budgetsContract.getNumBudgets(props.organisationsRef, props.organisationRef)
      const numBudgets = num.toNumber()
      for (let i = 0; i < numBudgets; i++) {
         const budgetRef = await budgetsContract.getBudgetReference(organisationsRef,
                                                                    organisationRef,
                                                                    i.toString())
         const budget: IATIBudgetReportProps = await budgetsContract.getBudget(organisationsRef,
                                                        organisationRef,
                                                        budgetRef)
         report[organisationsRef][organisationRef].totalBudget[budgetRef] = budget
         actionType = IATIReportActionTypes.BUDGET_SUCCESS
      }
    } catch (error) {
      console.log('getBudgets error', error)
    }

    dispatch(read({data: report})(actionType))
  }
}
