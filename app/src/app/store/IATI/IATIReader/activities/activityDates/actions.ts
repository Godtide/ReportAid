import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIActivityDateProps,
         IATIActivityDatesData,
         IATIReportActionTypes,
         IATIActivityDatesReportProps,
         ActivitiesReportProps } from '../../../types'

import { read } from '../../actions'

interface RecordProps {
  activitiesRef: string
  activityRef: string
  dateRef: string
}

interface ActivityDatesDataProps {
  data: Array<IATIActivityDatesData>
  successActionType: IATIReportActionTypes
  failureActionType: IATIReportActionTypes
}

type ActivityDateProps = RecordProps & ActivityDatesDataProps

const getThisActivityDate = (props: ActivityDateProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityDatesContract = state.chainContracts.data.contracts.activityDates
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef
    const dateRef = props.dateRef
    const data =  props.data

    let datesData: IATIActivityDatesReportProps = {
      data: { activitiesRef: activitiesRef,
              activityRef: activityRef,
              data: data }
    }

    let actionType =  props.successActionType

    try {

      const thisDate: IATIActivityDateProps = await activityDatesContract.getDate(activitiesRef, activityRef, dateRef)

      datesData.data.data[data.length] = {
        dateRef: dateRef,
        dateType: thisDate.dateType,
        date: ethers.utils.parseBytes32String(thisDate.date),
        narrative: thisDate.narrative
      }
    } catch (error) {
      actionType = props.failureActionType
      console.log('getActivityDates error', error)
    }

    dispatch(read({data: datesData})(actionType))
  }
}

export const getActivityDateRecord = (props: RecordProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {

    let indexRecordProps: ActivityDateProps = {
      activitiesRef: props.activitiesRef,
      activityRef: props.activityRef,
      dateRef: props.dateRef,
      data: [],
      successActionType: IATIReportActionTypes.ACTIVITYDATE_SUCCESS,
      failureActionType: IATIReportActionTypes.ACTIVITYDATE_FAILURE
    }

    dispatch(getThisActivityDate(indexRecordProps))
  }
}

export const getActivityDates = (props: ActivitiesReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityDatesContract = state.chainContracts.data.contracts.activityDates

    let indexRecordProps: ActivityDateProps = {
      activitiesRef: props.activitiesRef,
      activityRef: props.activityRef,
      dateRef: "",
      data: [],
      successActionType: IATIReportActionTypes.ACTIVITYDATE_SUCCESS,
      failureActionType: IATIReportActionTypes.ACTIVITYDATE_FAILURE
    }

    try {
      const num = await activityDatesContract.getNumDates(indexRecordProps.activitiesRef, indexRecordProps.activityRef)
      const numDates = num.toNumber()
      for (let i = 0; i < numDates; i++) {

        indexRecordProps.dateRef = await activityDatesContract.getReference(indexRecordProps.activitiesRef, indexRecordProps.activityRef, i.toString())
        dispatch(getThisActivityDate(indexRecordProps))
      }
    } catch (error) {
      console.log('getActivityDate error', error)
    }
  }
}

export const getActivityDateRefs = (props: ActivitiesReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityDatesContract = state.chainContracts.data.contracts.activityDates
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef

    let keysData: Array<string> = []
    let actionType = IATIReportActionTypes.ACTIVITYDATEPICKER_SUCCESS

    try {
      const num = await activityDatesContract.getNumDates(activitiesRef, activityRef)
      const numDates = num.toNumber()
      for (let i = 0; i < numDates; i++) {
        const dateRef = await activityDatesContract.getReference(activitiesRef, activityRef, i.toString())
        keysData.push(dateRef)
      }
    } catch (error) {
      actionType = IATIReportActionTypes.ACTIVITYDATE_FAILURE
      console.log('getActivityDateRefs error', error)
    }

    //console.log('KeysData: ', keysData, actionType)
    dispatch(read({data: keysData})(actionType))
  }
}
