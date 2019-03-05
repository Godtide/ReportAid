import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIActivitiesProps, IATIReportActionTypes, IATIActivitiesReportProps } from '../../../types'

import { read } from '../../actions'

export const getActivitiesRecord = (props: ActivitiesProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    let actionType = IATIReportActionTypes.ACTIVITIES_FAILURE
    try {

       const activities: IATIActivitiesProps = await activitiesContract.getActivities(props.activitiesRef)
       activitiesData.data.data[i] = {
         activitiesRef: activitiesRef,
         version: ethers.utils.parseBytes32String(activities.version),
         generatedTime:  ethers.utils.parseBytes32String(activities.generatedTime),
         linkedData: ethers.utils.parseBytes32String(activities.linkedData),
       }

       actionType = IATIReportActionTypes.ACTIVITIESPICKER_SUCCESS
       if(props.isReport) {
        actionType = IATIReportActionTypes.ACTIVITIES_SUCCESS
       }

    } catch (error) {
      console.log('getActivities error', error)
    }
    dispatch(read({data: activitiesData})(actionType))
  }
}

export const getActivities = (isReport: boolean) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activitiesContract = state.chainContracts.data.contracts.activities

    let activitiesData: IATIActivitiesReportProps = {
      data: { data: [] }
    }

    let actionType = IATIReportActionTypes.ACTIVITIES_FAILURE
    try {
      const num = await activitiesContract.getNumActivities()
      const numActivities = num.toNumber()
      for (let i = 0; i < numActivities; i++) {
         const activitiesRef = await activitiesContract.getReference(i.toString())
         //console.log('Activities: ', activitiesRef, activitiesRef.length, numActivities)
         getActivitiesRecord({isReport: isReport, activitiesRef: activitiesRef})
      }
    } catch (error) {
      console.log('getActivities error', error)
    }

    dispatch(read({data: activitiesData})(actionType))

  }
}
