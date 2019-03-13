import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIActivityProps,
         IATIReportActionTypes,
         IATIActivityReportProps,
         IATIActivityData,
         ActivityReportProps } from '../../../types'

import { read } from '../../actions'

interface RecordProps {
  activitiesRef: string
  activityRef: string
}

interface ActivityDataProps {
  data: Array<IATIActivityData>
  successActionType: IATIReportActionTypes
  failureActionType: IATIReportActionTypes
}

type ActivityProps = RecordProps & ActivityDataProps

const getThisActivity = (props: ActivityProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityContract = state.chainContracts.data.contracts.activity
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef
    const data =  props.data

    let activityData: IATIActivityReportProps = {
      data: { activitiesRef: activitiesRef, data: data }
    }

    let actionType =  props.successActionType
    try {

     const activity: IATIActivityProps = await activityContract.getActivity(activitiesRef, activityRef)

     //console.log('Got activity: ', activitiesRef, activityRef, activity)

     activityData.data.data[activityData.data.data.length] = {
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
    } catch (error) {
      actionType = props.failureActionType
      console.log('getActivity error', error)
    }

    dispatch(read({data: activityData})(actionType))
  }
}

export const getActivityRecord = (props: RecordProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {

    let indexRecordProps: ActivityProps = {
      activitiesRef: props.activitiesRef,
      activityRef: props.activityRef,
      data: [],
      successActionType: IATIReportActionTypes.ACTIVITY_SUCCESS,
      failureActionType: IATIReportActionTypes.ACTIVITY_FAILURE
    }

    dispatch(getThisActivity(indexRecordProps))
  }
}

export const getActivity = (props: ActivityReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityContract = state.chainContracts.data.contracts.activity

    let indexRecordProps: ActivityProps = {
      activitiesRef: props.activitiesRef,
      activityRef: "",
      data: [],
      successActionType: IATIReportActionTypes.ACTIVITY_SUCCESS,
      failureActionType: IATIReportActionTypes.ACTIVITY_FAILURE
    }

    try {
      const num = await activityContract.getNumActivities(indexRecordProps.activitiesRef)
      const numActivitys = num.toNumber()
      for (let i = 0; i < numActivitys; i++) {
        indexRecordProps.activityRef = await activityContract.getReference(indexRecordProps.activitiesRef, i.toString())
        dispatch(getThisActivity(indexRecordProps))
      }
    } catch (error) {
      console.log('getActivity error', error)
    }
  }
}

export const getActivityRefs = (props: ActivityReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityContract = state.chainContracts.data.contracts.activity
    const activitiesRef = props.activitiesRef

    let keysData: Array<string> = []
    let actionType = IATIReportActionTypes.ACTIVITYPICKER_SUCCESS

    try {
      const num = await activityContract.getNumActivities(activitiesRef)
      const numActivitys = num.toNumber()
      for (let i = 0; i < numActivitys; i++) {
        const activityRef = await activityContract.getReference(activitiesRef, i.toString())
        keysData.push(activityRef)
      }
    } catch (error) {
      actionType = IATIReportActionTypes.ACTIVITY_FAILURE
      console.log('getActivityKeys error', error)
    }

    //console.log('KeysData: ', keysData, actionType)
    dispatch(read({data: keysData})(actionType))
  }
}
