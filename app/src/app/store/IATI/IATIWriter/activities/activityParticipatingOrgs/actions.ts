import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../../../../actions'

import { ActionProps, TxReport } from '../../../../types'
import { ActivityParticipatingOrgProps, IATIActivityParticipatingOrgProps, IATIWriterActionTypes } from '../../../types'

import { Transaction } from '../../../../../utils/strings'

export const setActivityParticipatingOrg = (details: ActivityParticipatingOrgProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityParticipatingOrgsContract = state.chainContracts.data.contracts.activityParticipatingOrgs

    let participatingOrgRef = details.participatingOrgRef
    if ( participatingOrgRef == "" ) {
      participatingOrgRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const participatingOrgData: IATIActivityParticipatingOrgProps = {
      orgType: details.orgType,
      role: details.role,
      crsChannelCode: details.crsChannelCode,
      lang: ethers.utils.formatBytes32String(details.lang),
      orgRef: details.orgRef,
      narrative:  details.narrative
    }

    let actionType = IATIWriterActionTypes.ACTIVITYPARTICIPATINGORG_FAILURE
    let txData: TxReport = {}
    try {
      const tx = await activityParticipatingOrgsContract.setParticipatingOrg(details.activitiesRef, details.activityRef, participatingOrgRef, participatingOrgData)
      const key = tx.hash
      txData = {
        [key]: {
          summary: `${Transaction.success}`,
          info: tx
        }
      }

      actionType = IATIWriterActionTypes.ACTIVITYPARTICIPATINGORG_SUCCESS
    } catch (error) {
      txData = {
        [-1]: {
          summary: `${Transaction.fail}`,
          info: {}
        }
      }
      console.log('setActivityParticipatingOrg error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
