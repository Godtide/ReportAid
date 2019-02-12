import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIOrganisationCountryBudgetProps } from '../../../types'
import { IATIReportActionTypes,
         IATIRecipientCountryBudgetReportProps,
         OrganisationsReportProps } from '../types'

import { read } from '../actions'

export const getCountryBudgets = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const countryBudgetsContract = state.chainContracts.data.contracts.organisationCountryBudgets
    const organisationsRef = props.organisationsRef
    const organisationRef = props.organisationRef

    const budgetReports = state.organisationsReport.data[organisationsRef].data[organisationRef].data.recipientCountryBudget as IATIRecipientCountryBudgetReportProps

    let actionType = IATIReportActionTypes.RECIPIENTCOUNTRYBUDGET_FAILURE
    try {
      const num = await countryBudgetsContract.getNumCountryBudgets(props.organisationsRef, props.organisationRef)
      const numBudgets = num.toNumber()
      for (let i = 0; i < numBudgets; i++) {
         const budgetRef = await countryBudgetsContract.getCountryBudgetReference(organisationsRef,
                                                                                  organisationRef,
                                                                                  i.toString())
         const budget: IATIOrganisationCountryBudgetProps = await countryBudgetsContract.getCountryBudget(organisationsRef,
                                                                                                          organisationRef,
                                                                                                          budgetRef)
         budgetReports[budgetRef] = budget
         actionType = IATIReportActionTypes.RECIPIENTCOUNTRYBUDGET_SUCCESS
      }
    } catch (error) {
      console.log('getCountryBudgets error', error)
    }

    dispatch(read({data: budgetReports})(actionType))
  }
}
