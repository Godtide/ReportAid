import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIActivityAdditionalProps, IATIReportActionTypes, IATIActivityAdditionalReportProps, ActivitiesReportProps } from '../../../types'

import { read } from '../../actions'

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
        const additional: IATIActivityAdditionalProps = await activityAdditionalContract.getActivityAdditional(activitiesRef, activityRef, additionalRef)

        //console.log("Additional: ", additional)

        additionalData.data.data[i] = {
           additionalRef: additionalRef,
           budgetNotProvided: additional.budgetNotProvided,
           status: additional.status,
           scope: additional.scope,
           capitalSpend: additional.capitalSpend,
           collaborationType: additional.collaborationType,
           defaultFlowType: additional.defaultFlowType,
           defaultFinanceType: ethers.utils.bigNumberify(additional.defaultFinanceType).toNumber(),
           defaultAidType: ethers.utils.parseBytes32String(additional.defaultAidType),
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
