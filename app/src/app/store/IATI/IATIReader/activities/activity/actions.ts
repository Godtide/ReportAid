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

export const getActivityRecord = (props: RecordProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityContract = state.chainContracts.data.contracts.activity
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef

    let activityData: IATIActivityReportProps = {
      data: { activitiesRef: props.activitiesRef, data: [] }
    }
    let actionType = IATIReportActionTypes.ACTIVITIES_FAILURE
    try {

      const activity: IATIActivityProps = await activityContract.getActivity(activitiesRef, activityRef)
      activityData.data.data[0] = {
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

      actionType = IATIReportActionTypes.ACTIVITY_SUCCESS
    } catch (error) {
     console.log('getActivityRecord error', error)
    }

    dispatch(read({data: activityData})(actionType))
  }
}

export const getActivity = (props: ActivityReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityContract = state.chainContracts.data.contracts.activity
    const activitiesRef = props.activitiesRef

    let activityData: IATIActivityReportProps = {
      data: { activitiesRef: activitiesRef, data: [] }
    }

    let actionType = IATIReportActionTypes.ACTIVITY_FAILURE
    try {
      const num = await activityContract.getNumActivities(activitiesRef)
      const numActivitys = num.toNumber()
      for (let i = 0; i < numActivitys; i++) {
         const activityRef = await activityContract.getReference(activitiesRef, i.toString())
         const activity: IATIActivityProps = await activityContract.getActivity(activitiesRef, activityRef)

         activityData.data.data[i] = {
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

         actionType = IATIReportActionTypes.ACTIVITYPICKER_SUCCESS
         if(props.isReport) {
          actionType = IATIReportActionTypes.ACTIVITY_SUCCESS
         }
      }
    } catch (error) {
      console.log('getActivity error', error)
    }

    dispatch(read({data: activityData})(actionType))
  }
}
