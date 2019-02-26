import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIBudgetProps } from '../../../types'
import { IATIReportActionTypes, OrganisationsReportProps } from '../types'
import { IATIBudgetReportProps, IATIBudgetData } from '../../types'

import { read } from '../actions'

export const initialise = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const initData: any = { data: {} }
    await dispatch(read({data: initData})(IATIReportActionTypes.RECIPIENTCOUNTRYBUDGET_INIT))
  }
}

export const getCountryBudgets = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const countryBudgetsContract = state.chainContracts.data.contracts.organisationCountryBudgets
    const organisationsRef = props.organisationsRef
    const organisationRef = props.organisationRef

    let budgetReports: IATIBudgetReportProps = {
      data: { organisationsRef: organisationsRef,
              organisationRef: organisationRef,
              data: []
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
         const budget: IATIBudgetProps = await countryBudgetsContract.getCountryBudget(organisationsRef,
                                                                                       organisationRef,
                                                                                       budgetRef)

         budgetReports.data.data[i] = {
           budgetKey: budgetRef,
           budgetLine: ethers.utils.parseBytes32String(budget.budgetLine),
           countryRef: ethers.utils.parseBytes32String(budget.otherRef),
           value: ethers.utils.bigNumberify(budget.finance.value).toNumber(),
           status: budget.finance.status,
           start: ethers.utils.parseBytes32String(budget.finance.start),
           end: ethers.utils.parseBytes32String(budget.finance.end)
         }
     actionType = IATIReportActionTypes.RECIPIENTCOUNTRYBUDGET_SUCCESS
      }
    } catch (error) {
      console.log('getCountryBudgets error', error)
    }

    dispatch(read({data: budgetReports})(actionType))
  }
}
