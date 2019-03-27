import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIActivityRelatedActivityProps,
         IATIActivityRelatedActivityData,
         IATIReportActionTypes,
         IATIActivityRelatedActivityReportProps,
         ActivitiesReportProps } from '../../../types'

import { read } from '../../../../actions'

interface RecordProps {
  activitiesRef: string
  activityRef: string
  activityRelatedActivityRef: string
}

interface ActivityRelatedActivitysDataProps {
  data: Array<IATIActivityRelatedActivityData>
  successActionType: IATIReportActionTypes
  failureActionType: IATIReportActionTypes
}

type ActivityRelatedActivityProps = RecordProps & ActivityRelatedActivitysDataProps

const getThisRelatedActivity = (props: ActivityRelatedActivityProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityRelatedActivitiesContract = state.chainContracts.data.contracts.activityRelatedActivities
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef
    const activityRelatedActivityRef = props.activityRelatedActivityRef
    const data =  props.data

    let activityRelatedActivityData: IATIActivityRelatedActivityReportProps = {
      data: { activitiesRef: activitiesRef,
              activityRef: activityRef,
              data: data }
    }

    let actionType =  props.successActionType

    try {

      const activityRelatedActivity: IATIActivityRelatedActivityProps = await activityRelatedActivitiesContract.getRelatedActivity(activitiesRef, activityRef, activityRelatedActivityRef)

      // ethers.utils.bigNumberify(var).toNumber()
      activityRelatedActivityData.data.data[data.length] = {
        activityRelatedActivityRef: activityRelatedActivityRef,
        relationType: activityRelatedActivity.relationType,
        relatedActivityRef: activityRelatedActivity.activityRef
      }
    } catch (error) {
      actionType = props.failureActionType
      console.log('getActivityRelatedActivitys error', error)
    }

    dispatch(read({data: activityRelatedActivityData})(actionType))
  }
}

export const getRelatedActivityRecord = (props: RecordProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {

    let indexRecordProps: ActivityRelatedActivityProps = {
      activitiesRef: props.activitiesRef,
      activityRef: props.activityRef,
      activityRelatedActivityRef: props.activityRelatedActivityRef,
      data: [],
      successActionType: IATIReportActionTypes.ACTIVITYRELATEDACTIVITY_SUCCESS,
      failureActionType: IATIReportActionTypes.ACTIVITYRELATEDACTIVITY_FAILURE
    }

    dispatch(getThisRelatedActivity(indexRecordProps))
  }
}

export const getRelatedActivities = (props: ActivitiesReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityRelatedActivitiesContract = state.chainContracts.data.contracts.activityRelatedActivities

    let indexRecordProps: ActivityRelatedActivityProps = {
      activitiesRef: props.activitiesRef,
      activityRef: props.activityRef,
      activityRelatedActivityRef: "",
      data: [],
      successActionType: IATIReportActionTypes.ACTIVITYRELATEDACTIVITY_SUCCESS,
      failureActionType: IATIReportActionTypes.ACTIVITYRELATEDACTIVITY_FAILURE
    }

    try {
      const num = await activityRelatedActivitiesContract.getNum(indexRecordProps.activitiesRef, indexRecordProps.activityRef)
      const numParticipatingOrgs = num.toNumber()
      for (let i = 0; i < numParticipatingOrgs; i++) {

        indexRecordProps.activityRelatedActivityRef = await activityRelatedActivitiesContract.getReference(indexRecordProps.activitiesRef, indexRecordProps.activityRef, i.toString())
        dispatch(getThisRelatedActivity(indexRecordProps))
      }
    } catch (error) {
      console.log('getActivityRelatedActivity error', error)
    }
  }
}

export const getRelatedActivityRefs = (props: ActivitiesReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityRelatedActivitiesContract = state.chainContracts.data.contracts.activityRelatedActivities
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef

    let keysData: Array<string> = []
    let actionType = IATIReportActionTypes.ACTIVITYRELATEDACTIVITYPICKER_SUCCESS

    try {
      const num = await activityRelatedActivitiesContract.getNum(activitiesRef, activityRef)
      const numParticipatingOrgs = num.toNumber()
      for (let i = 0; i < numParticipatingOrgs; i++) {
        const activityRelatedActivityRef = await activityRelatedActivitiesContract.getReference(activitiesRef, activityRef, i.toString())
        keysData.push(activityRelatedActivityRef)
      }
    } catch (error) {
      actionType = IATIReportActionTypes.ACTIVITYRELATEDACTIVITY_FAILURE
      console.log('getActivityRelatedActivityRefs error', error)
    }

    //console.log('KeysData: ', keysData, actionType)
    dispatch(read({data: keysData})(actionType))
  }
}
