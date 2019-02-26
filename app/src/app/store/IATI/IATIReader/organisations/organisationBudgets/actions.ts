import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIBudgetProps } from '../../../types'
import { IATIReportActionTypes, OrganisationsReportProps } from '../types'
import { IATIOrganisationBudgetReportProps, IATIBudgetData } from './types'

import { read } from '../actions'

export const initialise = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const initData: any = { data: {} }
    await dispatch(read({data: initData})(IATIReportActionTypes.BUDGET_INIT))
  }
}

export const getBudgets = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const budgetsContract = state.chainContracts.data.contracts.organisationBudgets
    const organisationsRef = props.organisationsRef
    const organisationRef = props.organisationRef

    let budgetReports: IATIOrganisationBudgetReportProps = { data: {
                                                              organisationsRef: organisationsRef,
                                                              organisationRef: organisationRef,
                                                              data: []
                                                            }
                                                           }
    let actionType = IATIReportActionTypes.BUDGET_FAILURE
    try {
      const num = await budgetsContract.getNumBudgets(props.organisationsRef, props.organisationRef)
      const numBudgets = num.toNumber()
      for (let i = 0; i < numBudgets; i++) {
        const budgetRef = await budgetsContract.getBudgetReference(organisationsRef,
                                                                    organisationRef,
                                                                    i.toString())
        const budget: IATIBudgetProps = await budgetsContract.getBudget(organisationsRef,
                                                                       organisationRef,
                                                                       budgetRef)

        budgetReports.data.data[i] = {
          budgetKey: budgetRef,
          budgetType: budget.budgetType,
          budgetLine: ethers.utils.parseBytes32String(budget.budgetLine),
          otherRef: ethers.utils.parseBytes32String(budget.otherRef),
          value: ethers.utils.bigNumberify(budget.finance.value).toNumber(),
          status: budget.finance.status,
          start: ethers.utils.parseBytes32String(budget.finance.start),
          end: ethers.utils.parseBytes32String(budget.finance.end)
        }

        actionType = IATIReportActionTypes.BUDGET_SUCCESS
      }
    } catch (error) {
      console.log('getBudgets error', error)
    }

    dispatch(read({data: budgetReports})(actionType))
  }
}
