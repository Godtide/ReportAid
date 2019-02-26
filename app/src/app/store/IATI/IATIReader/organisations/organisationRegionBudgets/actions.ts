import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIBudgetProps } from '../../../types'
import { IATIReportActionTypes, OrganisationsReportProps } from '../types'
import { IATIOrganisationRegionBudgetReportProps } from './types'

import { read } from '../actions'

export const initialise = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const initData: IATIOrganisationRegionBudgetReportProps = { data: {} }
    await dispatch(read({data: initData})(IATIReportActionTypes.RECIPIENTREGIONBUDGET_INIT))
  }
}

export const getRegionBudgets = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const regionBudgetsContract = state.chainContracts.data.contracts.organisationRegionBudgets
    const organisationsRef = props.organisationsRef
    const organisationRef = props.organisationRef

    let budgetReports: IATIOrganisationRegionBudgetReportProps = {
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
    let actionType = IATIReportActionTypes.RECIPIENTREGIONBUDGET_FAILURE
    try {
      const num = await regionBudgetsContract.getNumRegionBudgets(props.organisationsRef, props.organisationRef)
      const numBudgets = num.toNumber()
      for (let i = 0; i < numBudgets; i++) {
         const budgetRef = await regionBudgetsContract.getRegionBudgetReference(organisationsRef,
                                                                                organisationRef,
                                                                                i.toString())
         const budget: IATIBudgetProps = await regionBudgetsContract.getRegionsBudget(organisationsRef,     organisationRef, budgetRef)
         budgetReports.data[organisationsRef].data[organisationRef].data[budgetRef] = budget
         actionType = IATIReportActionTypes.RECIPIENTREGIONBUDGET_SUCCESS
      }
    } catch (error) {
      console.log('getRegionBudgets error', error)
    }

    dispatch(read({data: budgetReports})(actionType))
  }
}
