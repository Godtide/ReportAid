import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIActivitiesProps,
         IATIReportActionTypes,
         IATIActivitiesData,
         IATIActivitiesReportProps } from '../../../types'

import { read } from '../../actions'

interface RecordProps {
  activitiesRef: string
}

interface ActivitiesDataProps {
  data: Array<IATIActivitiesData>
  successActionType: IATIReportActionTypes
  failureActionType: IATIReportActionTypes
}

type ActivitiesProps = RecordProps & ActivitiesDataProps

const getThisActivitiesRecord = (props: ActivitiesProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activitiesContract = state.chainContracts.data.contracts.activities
    const activitiesRef = props.activitiesRef
    const data =  props.data

    let activitiesData: IATIActivitiesReportProps = {
      data: { data: data }
    }

    let actionType =  props.successActionType
    try {

      const activities: IATIActivitiesProps = await activitiesContract.getActivities(props.activitiesRef)
      activitiesData.data.data[data.length] = {
        activitiesRef: props.activitiesRef,
        version: ethers.utils.parseBytes32String(activities.version),
        generatedTime:  ethers.utils.parseBytes32String(activities.generatedTime),
        linkedData: ethers.utils.parseBytes32String(activities.linkedData),
      }
    } catch (error) {
      actionType = props.failureActionType
      console.log('getActivities error', error)
    }

    //console.log('Getting activities: ', activitiesData)

    dispatch(read({data: activitiesData})(actionType))
  }
}

export const getActivitiesRecord = (props: RecordProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    let indexRecordProps: ActivitiesProps = {
      activitiesRef: props.activitiesRef,
      data: [],
      successActionType: IATIReportActionTypes.ACTIVITIES_SUCCESS,
      failureActionType: IATIReportActionTypes.ACTIVITIES_FAILURE
    }

    dispatch(getThisActivitiesRecord(indexRecordProps))
  }
}

export const getActivities = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activitiesContract = state.chainContracts.data.contracts.activities

    let indexRecordProps: ActivitiesProps = {
      activitiesRef: "",
      data: [],
      successActionType: IATIReportActionTypes.ACTIVITIES_SUCCESS,
      failureActionType: IATIReportActionTypes.ACTIVITIES_FAILURE
    }

    try {
      const num = await activitiesContract.getNumActivities()
      const numActivities = num.toNumber()
      for (let i = 0; i < numActivities; i++) {
        indexRecordProps.activitiesRef = await activitiesContract.getReference(i.toString())
        dispatch(getThisActivitiesRecord(indexRecordProps))
      }
    } catch (error) {
      console.log('getActivities error', error)
    }
  }
}

export const getActivitiesRefs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activitiesContract = state.chainContracts.data.contracts.activities

    let keysData: Array<string> = []
    let actionType = IATIReportActionTypes.ACTIVITIESPICKER_SUCCESS

    try {
      const num = await activitiesContract.getNumActivities()
      const numActivities = num.toNumber()
      for (let i = 0; i < numActivities; i++) {
        const activitiesRef = await activitiesContract.getReference(i.toString())
        keysData.push(activitiesRef)
      }
    } catch (error) {
      actionType = IATIReportActionTypes.ACTIVITIES_FAILURE
      console.log('getActivitiesRefs error', error)
    }

    //console.log('KeysData: ', keysData, actionType)
    dispatch(read({data: keysData})(actionType))
  }
}
