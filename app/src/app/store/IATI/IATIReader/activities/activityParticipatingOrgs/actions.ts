import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIActivityParticipatingOrgProps,
         IATIActivityParticipatingOrgData,
         IATIReportActionTypes,
         IATIActivityParticipatingOrgReportProps,
         ActivitiesReportProps } from '../../../types'

import { read } from '../../actions'

interface RecordProps {
  activitiesRef: string
  activityRef: string
  participatingOrgRef: string
}

interface ActivityParticipatingOrgsDataProps {
  data: Array<IATIActivityParticipatingOrgData>
  successActionType: IATIReportActionTypes
  failureActionType: IATIReportActionTypes
}

type ActivityParticipatingOrgProps = RecordProps & ActivityParticipatingOrgsDataProps

const getThisActivityParticipatingOrg = (props: ActivityParticipatingOrgProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityParticipatingOrgsContract = state.chainContracts.data.contracts.activityParticipatingOrgs
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef
    const participatingOrgRef = props.participatingOrgRef
    const data =  props.data

    let participatingOrgData: IATIActivityParticipatingOrgReportProps = {
      data: { activitiesRef: activitiesRef,
              activityRef: activityRef,
              data: data }
    }

    let actionType =  props.successActionType

    try {

      const participatingOrg: IATIActivityParticipatingOrgProps = await activityParticipatingOrgsContract.getParticipatingOrg(activitiesRef, activityRef, participatingOrgRef)

      // ethers.utils.bigNumberify(var).toNumber()
      participatingOrgData.data.data[data.length] = {
        participatingOrgRef: participatingOrgRef,
        orgRef: participatingOrg.orgRef,
        orgType: participatingOrg.orgType,
        role: participatingOrg.role,
        crsChannelCode: ethers.utils.bigNumberify(participatingOrg.crsChannelCode).toNumber(),
        lang: ethers.utils.parseBytes32String(participatingOrg.lang),
        narrative: participatingOrg.narrative
      }
    } catch (error) {
      actionType = props.failureActionType
      console.log('getActivityParticipatingOrgs error', error)
    }

    dispatch(read({data: participatingOrgData})(actionType))
  }
}

export const getActivityParticipatingOrgRecord = (props: RecordProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {

    let indexRecordProps: ActivityParticipatingOrgProps = {
      activitiesRef: props.activitiesRef,
      activityRef: props.activityRef,
      participatingOrgRef: props.participatingOrgRef,
      data: [],
      successActionType: IATIReportActionTypes.ACTIVITYPARTICIPATINGORG_SUCCESS,
      failureActionType: IATIReportActionTypes.ACTIVITYPARTICIPATINGORG_FAILURE
    }

    dispatch(getThisActivityParticipatingOrg(indexRecordProps))
  }
}

export const getActivityParticipatingOrgs = (props: ActivitiesReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityParticipatingOrgsContract = state.chainContracts.data.contracts.activityParticipatingOrgs

    let indexRecordProps: ActivityParticipatingOrgProps = {
      activitiesRef: props.activitiesRef,
      activityRef: props.activityRef,
      participatingOrgRef: "",
      data: [],
      successActionType: IATIReportActionTypes.ACTIVITYPARTICIPATINGORG_SUCCESS,
      failureActionType: IATIReportActionTypes.ACTIVITYPARTICIPATINGORG_FAILURE
    }

    try {
      const num = await activityParticipatingOrgsContract.getNumParticipatingOrgs(indexRecordProps.activitiesRef, indexRecordProps.activityRef)
      const numParticipatingOrgs = num.toNumber()
      for (let i = 0; i < numParticipatingOrgs; i++) {

        indexRecordProps.participatingOrgRef = await activityParticipatingOrgsContract.getReference(indexRecordProps.activitiesRef, indexRecordProps.activityRef, i.toString())
        dispatch(getThisActivityParticipatingOrg(indexRecordProps))
      }
    } catch (error) {
      console.log('getActivityParticipatingOrg error', error)
    }
  }
}

export const getActivityParticipatingOrgRefs = (props: ActivitiesReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityParticipatingOrgsContract = state.chainContracts.data.contracts.activityParticipatingOrgs
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef

    let keysData: Array<string> = []
    let actionType = IATIReportActionTypes.ACTIVITYPARTICIPATINGORGPICKER_SUCCESS

    try {
      const num = await activityParticipatingOrgsContract.getNumParticipatingOrgs(activitiesRef, activityRef)
      const numParticipatingOrgs = num.toNumber()
      for (let i = 0; i < numParticipatingOrgs; i++) {
        const participatingOrgRef = await activityParticipatingOrgsContract.getReference(activitiesRef, activityRef, i.toString())
        keysData.push(participatingOrgRef)
      }
    } catch (error) {
      actionType = IATIReportActionTypes.ACTIVITYPARTICIPATINGORG_FAILURE
      console.log('getActivityParticipatingOrgRefs error', error)
    }

    //console.log('KeysData: ', keysData, actionType)
    dispatch(read({data: keysData})(actionType))
  }
}
