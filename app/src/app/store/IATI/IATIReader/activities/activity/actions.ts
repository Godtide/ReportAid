import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIActivityProps, IATIReportActionTypes, IATIActivityReportProps, ActivityReportProps } from '../../../types'

import { read } from '../../actions'

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
         const activityRef = await activityContract.getActivityReference(activitiesRef, i.toString())
         const activity: IATIActivityProps = await activityContract.getActivity(activitiesRef, activityRef)

         activityData.data.data[i] = {
           activityRef: activityRef,
           title: ethers.utils.parseBytes32String(activity.title),
           identifier: ethers.utils.parseBytes32String(activity.identifier),
           linkedData: activity.linkedData,
           lang: ethers.utils.parseBytes32String(activity.lang),
           currency: ethers.utils.parseBytes32String(activity.currency),
           lastUpdated: ethers.utils.parseBytes32String(activity.lastUpdated),
           reportingOrgRef: activity.reportingOrg.orgRef,
           reportingOrgType: activity.reportingOrg.orgType,
           reportingOrgIsSecondary: activity.reportingOrg.isSecondary,
           humanitarian: activity.humanitarian,
           hierarchy: activity.hierarchy,
           budgetNotProvided: activity.budgetNotProvided,
           status: activity.status,
           scope: activity.scope,
           capitalSpend: activity.capitalSpend,
           collaborationType: activity.collaborationType,
           defaultFlowType: activity.defaultFlowType,
           defaultFinanceType: activity.defaultFinanceType,
           defaultAidType: ethers.utils.parseBytes32String(activity.defaultAidType),
           defaultTiedStatus: activity.defaultTiedStatus
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
