import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIOrganisationRegionBudgetProps } from '../../../types'
import { IATIReportActionTypes,
         IATIRecipientRegionBudgetReportProps,
         OrganisationsReportProps } from '../types'

import { read } from '../actions'

export const getRegionBudgets = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const regionBudgetsContract = state.chainContracts.data.contracts.organisationRegionBudgets
    const organisationsRef = props.organisationsRef
    const organisationRef = props.organisationRef

    const budgetReports = state.organisationsReport.data[organisationsRef].data[organisationRef].data.recipientRegionBudget as IATIRecipientRegionBudgetReportProps

    let actionType = IATIReportActionTypes.RECIPIENTREGIONBUDGET_FAILURE
    try {
      const num = await regionBudgetsContract.getNumRegionBudgets(props.organisationsRef, props.organisationRef)
      const numBudgets = num.toNumber()
      for (let i = 0; i < numBudgets; i++) {
         const budgetRef = await regionBudgetsContract.getRegionBudgetReference(organisationsRef,
                                                                                organisationRef,
                                                                                i.toString())
         const budget: IATIOrganisationRegionBudgetProps = await regionBudgetsContract.getRegionsBudget(organisationsRef,
                                                                                                        organisationRef,
                                                                                                        budgetRef)
         budgetReports[budgetRef] = budget
         actionType = IATIReportActionTypes.RECIPIENTREGIONBUDGET_SUCCESS
      }
    } catch (error) {
      console.log('getRegionBudgets error', error)
    }

    dispatch(read({data: budgetReports})(actionType))
  }
}
