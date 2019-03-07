import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIActivityAdditionalProps, IATIReportActionTypes, IATIActivityAdditionalReportProps, ActivitiesReportProps } from '../../../types'

import { read } from '../../actions'

interface RecordProps {
  activitiesRef: string
  activityRef: string
  additionalRef: string
}

export const getActivityAdditionalRecord = (props: RecordProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityAdditionalContract = state.chainContracts.data.contracts.activityAdditional
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef
    const additionalRef = props.additionalRef

    let additionalData: IATIActivityAdditionalReportProps = {
      data: { activitiesRef: activitiesRef,
              activityRef: activityRef,
              data: [] }
    }
    let actionType = IATIReportActionTypes.ACTIVITYADDITIONAL_FAILURE
    try {

      const additional: IATIActivityAdditionalProps = await activityAdditionalContract.getAdditional(activitiesRef, activityRef, additionalRef)
      additionalData.data.data[0] = {
         additionalRef: additionalRef,
         defaultAidType: ethers.utils.parseBytes32String(additional.defaultAidType),
         defaultFinanceType: ethers.utils.bigNumberify(additional.defaultFinanceType).toNumber(),
         scope: additional.scope,
         capitalSpend: additional.capitalSpend,
         collaborationType: additional.collaborationType,
         defaultFlowType: additional.defaultFlowType,
         defaultTiedStatus: additional.defaultTiedStatus
       }

       actionType = IATIReportActionTypes.ACTIVITYADDITIONAL_SUCCESS
    } catch (error) {
     console.log('getActivityAdditional error', error)
    }

    dispatch(read({data: additionalData})(actionType))
  }
}

export const getActivityAdditionals = (props: ActivitiesReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityAdditionalContract = state.chainContracts.data.contracts.activityAdditional
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef

    let additionalData: IATIActivityAdditionalReportProps = {
      data: { activitiesRef: activitiesRef,
              activityRef: activityRef,
              data: [] }
    }

    let actionType = IATIReportActionTypes.ACTIVITYADDITIONAL_FAILURE
    try {

      const num = await activityAdditionalContract.getNumAdditional(activitiesRef, activityRef)
      const numAdditionals = num.toNumber()
      for (let i = 0; i < numAdditionals; i++) {

        const additionalRef = await activityAdditionalContract.getReference(activitiesRef, activityRef, i.toString())
        const additional: IATIActivityAdditionalProps = await activityAdditionalContract.getAdditional(activitiesRef, activityRef, additionalRef)

        additionalData.data.data[i] = {
           additionalRef: additionalRef,
           defaultAidType: ethers.utils.parseBytes32String(additional.defaultAidType),
           defaultFinanceType: ethers.utils.bigNumberify(additional.defaultFinanceType).toNumber(),
           scope: additional.scope,
           capitalSpend: additional.capitalSpend,
           collaborationType: additional.collaborationType,
           defaultFlowType: additional.defaultFlowType,
           defaultTiedStatus: additional.defaultTiedStatus
         }

         actionType = IATIReportActionTypes.ACTIVITYADDITIONAL_SUCCESS
      }
    } catch (error) {
      console.log('getActivityAdditional error', error)
    }

    dispatch(read({data: additionalData})(actionType))
  }
}
