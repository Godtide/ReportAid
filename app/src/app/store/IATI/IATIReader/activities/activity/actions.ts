import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIActivityProps,
         IATIReportActionTypes,
         IATIActivityReportProps,
         ActivityReportProps } from '../../../types'

import { read } from '../../actions'

interface RecordProps {
  activitiesRef: string
  activityRef: string
}

interface ReportIndexRecordProps {
  index: number
  successActionType: IATIReportActionTypes
  failureActionType: IATIReportActionTypes
}

type ActivityIndexedProps = RecordProps & ReportIndexRecordProps

const getThisActivity = (props: ActivityIndexedProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityContract = state.chainContracts.data.contracts.activity
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef
    const index = props.index

    let activityData: IATIActivityReportProps = {
      data: { activitiesRef: activitiesRef, data: [] }
    }

    let actionType = props.failureActionType
    try {

     const activity: IATIActivityProps = await activityContract.getActivity(activitiesRef, activityRef)

     activityData.data.data[index] = {
       activityRef: activityRef,
       title: ethers.utils.parseBytes32String(activity.title),
       description: activity.description,
       identifier: ethers.utils.parseBytes32String(activity.identifier),
       linkedData: ethers.utils.parseBytes32String(activity.linkedData),
       lang: ethers.utils.parseBytes32String(activity.lang),
       currency: ethers.utils.parseBytes32String(activity.currency),
       lastUpdated: ethers.utils.parseBytes32String(activity.lastUpdated),
       reportingOrgRef: activity.reportingOrg.orgRef,
       reportingOrgType: activity.reportingOrg.orgType,
       reportingOrgIsSecondary: activity.reportingOrg.isSecondary,
       status: activity.status,
       humanitarian: activity.humanitarian,
       hierarchy: activity.hierarchy,
       budgetNotProvided: activity.budgetNotProvided
     }

     actionType =  props.successActionType
    } catch (error) {
      console.log('getActivity error', error)
    }

    dispatch(read({data: activityData})(actionType))
  }
}

export const getActivityRecord = (props: RecordProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    let indexRecordProps: ActivityIndexedProps = {
      activitiesRef: props.activitiesRef,
      activityRef: props.activityRef,
      index: 0,
      successActionType: IATIReportActionTypes.ACTIVITY_FAILURE,
      failureActionType: IATIReportActionTypes.ACTIVITY_SUCCESS
    }

    getThisActivity(indexRecordProps)
  }
}

export const getActivityPicker = (props: ActivityReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityContract = state.chainContracts.data.contracts.activity

    let indexRecordProps: ActivityIndexedProps = {
      activitiesRef: props.activitiesRef,
      activityRef: "",
      index: 0,
      successActionType: IATIReportActionTypes.ACTIVITY_FAILURE,
      failureActionType: IATIReportActionTypes.ACTIVITYPICKER_SUCCESS
    }

    try {
      const num = await activityContract.getNumActivities(indexRecordProps.activitiesRef)
      const numActivitys = num.toNumber()
      for (let i = 0; i < numActivitys; i++) {
        indexRecordProps.index = i
        indexRecordProps.activityRef = await activityContract.getReference(indexRecordProps.activitiesRef, indexRecordProps.index.toString())
        getThisActivity(indexRecordProps)
      }
    } catch (error) {
      console.log('getActivityPicker error', error)
    }
  }
}

export const getActivity = (props: ActivityReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityContract = state.chainContracts.data.contracts.activity

    let indexRecordProps: ActivityIndexedProps = {
      activitiesRef: props.activitiesRef,
      activityRef: "",
      index: 0,
      successActionType: IATIReportActionTypes.ACTIVITY_FAILURE,
      failureActionType: IATIReportActionTypes.ACTIVITY_SUCCESS
    }

    try {
      const num = await activityContract.getNumActivities(indexRecordProps.activitiesRef)
      const numActivitys = num.toNumber()
      for (let i = 0; i < numActivitys; i++) {
        indexRecordProps.index = i
        indexRecordProps.activityRef = await activityContract.getReference(indexRecordProps.activitiesRef, indexRecordProps.index.toString())
        getThisActivity(indexRecordProps)
      }
    } catch (error) {
      console.log('getActivity error', error)
    }
  }
}
