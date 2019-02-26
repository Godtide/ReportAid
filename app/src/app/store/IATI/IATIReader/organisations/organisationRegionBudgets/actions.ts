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
    await dispatch(read({data: initData})(IATIReportActionTypes.RECIPIENTREGIONBUDGET_INIT))
  }
}

export const getRegionBudgets = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const regionBudgetsContract = state.chainContracts.data.contracts.organisationRegionBudgets
    const organisationsRef = props.organisationsRef
    const organisationRef = props.organisationRef

    let budgetReports: IATIBudgetReportProps = {
      data: { organisationsRef: organisationsRef,
              organisationRef: organisationRef,
              data: []
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

         budgetReports.data.data[i] = {
           budgetKey: budgetRef,
           budgetLine: ethers.utils.parseBytes32String(budget.budgetLine),
           regionRef: ethers.utils.parseBytes32String(budget.otherRef),
           value: ethers.utils.bigNumberify(budget.finance.value).toNumber(),
           status: budget.finance.status,
           start: ethers.utils.parseBytes32String(budget.finance.start),
           end: ethers.utils.parseBytes32String(budget.finance.end)
         }

         actionType = IATIReportActionTypes.RECIPIENTREGIONBUDGET_SUCCESS
      }
    } catch (error) {
      console.log('getRegionBudgets error', error)
    }

    dispatch(read({data: budgetReports})(actionType))
  }
}
