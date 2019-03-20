import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIActivityBudgetData,
         IATIBudgetProps,
         IATIReportActionTypes,
         IATIActivityBudgetReportProps,
         ActivitiesReportProps } from '../../../types'

import { read } from '../../actions'

interface RecordProps {
  activitiesRef: string
  activityRef: string
  budgetRef: string
}

interface ActivityBudgetsDataProps {
  data: Array<IATIActivityBudgetData>
  successActionType: IATIReportActionTypes
  failureActionType: IATIReportActionTypes
}

type ActivityBudgetProps = RecordProps & ActivityBudgetsDataProps

const getThisActivityBudget = (props: ActivityBudgetProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityBudgetsContract = state.chainContracts.data.contracts.activityBudgets
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef
    const budgetRef = props.budgetRef
    const data =  props.data

    let budgetData: IATIActivityBudgetReportProps = {
      data: { activitiesRef: activitiesRef,
              activityRef: activityRef,
              data: data }
    }

    let actionType =  props.successActionType

    try {

      //console.log('budget ref: ', activitiesRef, activityRef, budgetRef)

      const thisBudget: IATIBudgetProps = await activityBudgetsContract.getBudget(activitiesRef, activityRef, budgetRef)

      //console.log('budget data: ', thisBudget, data.length)

      budgetData.data.data[data.length] = {
        budgetRef: budgetRef,
        budgetType: thisBudget.budgetType,
        status: thisBudget.finance.status,
        value: ethers.utils.bigNumberify(thisBudget.finance.value).toNumber(),
        start: ethers.utils.parseBytes32String(thisBudget.finance.start),
        end: ethers.utils.parseBytes32String(thisBudget.finance.end)
      }
    } catch (error) {
      actionType = props.failureActionType
      console.log('getActivityBudgets error', error)
    }

    dispatch(read({data: budgetData})(actionType))
  }
}

export const getActivityBudgetRecord = (props: RecordProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {

    let indexRecordProps: ActivityBudgetProps = {
      activitiesRef: props.activitiesRef,
      activityRef: props.activityRef,
      budgetRef: props.budgetRef,
      data: [],
      successActionType: IATIReportActionTypes.ACTIVITYBUDGET_SUCCESS,
      failureActionType: IATIReportActionTypes.ACTIVITYBUDGET_FAILURE
    }

    dispatch(getThisActivityBudget(indexRecordProps))
  }
}

export const getActivityBudgets = (props: ActivitiesReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityBudgetsContract = state.chainContracts.data.contracts.activityBudgets

    let indexRecordProps: ActivityBudgetProps = {
      activitiesRef: props.activitiesRef,
      activityRef: props.activityRef,
      budgetRef: "",
      data: [],
      successActionType: IATIReportActionTypes.ACTIVITYBUDGET_SUCCESS,
      failureActionType: IATIReportActionTypes.ACTIVITYBUDGET_FAILURE
    }

    try {
      const num = await activityBudgetsContract.getNum(indexRecordProps.activitiesRef, indexRecordProps.activityRef)
      const nums = num.toNumber()
      for (let i = 0; i < nums; i++) {

        indexRecordProps.budgetRef = await activityBudgetsContract.getReference(indexRecordProps.activitiesRef, indexRecordProps.activityRef, i.toString())
        dispatch(getThisActivityBudget(indexRecordProps))
      }
    } catch (error) {
      console.log('getActivityBudget error', error)
    }
  }
}

export const getActivityBudgetRefs = (props: ActivitiesReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityBudgetsContract = state.chainContracts.data.contracts.activityBudgets
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef

    let keysData: Array<string> = []
    let actionType = IATIReportActionTypes.ACTIVITYBUDGETPICKER_SUCCESS

    try {
      const num = await activityBudgetsContract.getNum(activitiesRef, activityRef)
      const nums = num.toNumber()
      for (let i = 0; i < nums; i++) {
        const budgetRef = await activityBudgetsContract.getReference(activitiesRef, activityRef, i.toString())
        keysData.push(budgetRef)
      }
    } catch (error) {
      actionType = IATIReportActionTypes.ACTIVITYBUDGET_FAILURE
      console.log('getActivityBudgetRefs error', error)
    }

    //console.log('KeysData: ', keysData, actionType)
    dispatch(read({data: keysData})(actionType))
  }
}
