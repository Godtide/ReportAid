import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIActivityDateProps, IATIReportActionTypes, IATIActivityDatesReportProps, ActivitiesReportProps } from '../../../types'

import { read } from '../../actions'

interface RecordProps {
  activitiesRef: string
  activityRef: string
  dateRef: string
}

export const getActivityDateRecord = (props: RecordProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityDatesContract = state.chainContracts.data.contracts.activityDates
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef
    const dateRef = props.dateRef

    let datesData: IATIActivityDatesReportProps = {
      data: { activitiesRef: activitiesRef,
              activityRef: activityRef,
              data: [] }
    }
    let actionType = IATIReportActionTypes.ACTIVITYDATE_FAILURE
    try {

      const thisDate: IATIActivityDateProps = await activityDatesContract.getDate(activitiesRef, activityRef, dateRef)

      datesData.data.data[0] = {
        dateRef: dateRef,
        dateType: thisDate.dateType,
        date: ethers.utils.parseBytes32String(thisDate.date),
        narrative: thisDate.narrative
      }

       actionType = IATIReportActionTypes.ACTIVITYDATE_SUCCESS
    } catch (error) {
     console.log('getActivityDate error', error)
    }

    dispatch(read({data: datesData})(actionType))
  }
}

export const getActivityDates = (props: ActivitiesReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityDatesContract = state.chainContracts.data.contracts.activityDates
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef

    let datesData: IATIActivityDatesReportProps = {
      data: { activitiesRef: activitiesRef,
              activityRef: activityRef,
              data: [] }
    }

    let actionType = IATIReportActionTypes.ACTIVITYDATE_FAILURE
    try {

      const num = await activityDatesContract.getNumDates(activitiesRef, activityRef)
      const numAdditionals = num.toNumber()
      for (let i = 0; i < numAdditionals; i++) {

        const dateRef = await activityDatesContract.getReference(activitiesRef, activityRef, i.toString())
        const thisDate: IATIActivityDateProps = await activityDatesContract.getDate(activitiesRef, activityRef, dateRef)

        datesData.data.data[i] = {
           dateRef: dateRef,
           dateType: thisDate.dateType,
           date: ethers.utils.parseBytes32String(thisDate.date),
           narrative: thisDate.narrative
         }

         actionType = IATIReportActionTypes.ACTIVITYDATE_SUCCESS
      }
    } catch (error) {
      console.log('getActivityDate error', error)
    }

    dispatch(read({data: datesData})(actionType))
  }
}
