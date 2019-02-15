import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIOrganisationCountryBudgetProps } from '../../../types'
import { IATIReportActionTypes, OrganisationsReportProps } from '../types'
import { IATIOrganisationCountryBudgetReportProps } from './types'

import { read } from '../actions'

export const initialise = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const initData: IATIOrganisationCountryBudgetReportProps = { data: {} }
    dispatch(read({data: initData})(IATIReportActionTypes.RECIPIENTCOUNTRYBUDGET_INIT))
  }
}

export const getCountryBudgets = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const countryBudgetsContract = state.chainContracts.data.contracts.organisationCountryBudgets
    const organisationsRef = props.organisationsRef
    const organisationRef = props.organisationRef

    let budgetReports: IATIOrganisationCountryBudgetReportProps = {
      data: {
        [organisationsRef]: {
          data: {
            [organisationRef]: {
              data: {}
            }
          }
        }
      }
    }

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
         budgetReports.data[organisationsRef].data[organisationRef].data[budgetRef] = budget
         actionType = IATIReportActionTypes.RECIPIENTCOUNTRYBUDGET_SUCCESS
      }
    } catch (error) {
      console.log('getCountryBudgets error', error)
    }

    dispatch(read({data: budgetReports})(actionType))
  }
}
