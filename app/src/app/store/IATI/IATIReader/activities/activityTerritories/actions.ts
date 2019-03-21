import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIActivityTerritoryProps,
         IATIActivityTerritoryData,
         IATIReportActionTypes,
         IATIActivityTerritoryReportProps,
         ActivitiesReportProps } from '../../../types'

import { read } from '../../../../actions'

interface RecordProps {
  activitiesRef: string
  activityRef: string
  territoryRef: string
}

interface ActivityTerritorysDataProps {
  data: Array<IATIActivityTerritoryData>
  successActionType: IATIReportActionTypes
  failureActionType: IATIReportActionTypes
}

type ActivityTerritoryProps = RecordProps & ActivityTerritorysDataProps

const getThisActivityTerritory = (props: ActivityTerritoryProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityTerritoriesContract = state.chainContracts.data.contracts.activityTerritories
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef
    const territoryRef = props.territoryRef
    const data =  props.data

    let territoryData: IATIActivityTerritoryReportProps = {
      data: { activitiesRef: activitiesRef,
              activityRef: activityRef,
              data: data }
    }

    let actionType =  props.successActionType

    try {

      const thisTerritory: IATIActivityTerritoryProps = await activityTerritoriesContract.getTerritory(activitiesRef, activityRef, territoryRef)

      territoryData.data.data[data.length] = {
        territoryRef: territoryRef,
        territory: ethers.utils.parseBytes32String(thisTerritory.territory),
        percentage: thisTerritory.percentage,
      }
    } catch (error) {
      actionType = props.failureActionType
      console.log('getActivityTerritories error', error)
    }

    dispatch(read({data: territoryData})(actionType))
  }
}

export const getActivityTerritoryRecord = (props: RecordProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {

    let indexRecordProps: ActivityTerritoryProps = {
      activitiesRef: props.activitiesRef,
      activityRef: props.activityRef,
      territoryRef: props.territoryRef,
      data: [],
      successActionType: IATIReportActionTypes.ACTIVITYTERRITORY_SUCCESS,
      failureActionType: IATIReportActionTypes.ACTIVITYTERRITORY_FAILURE
    }

    dispatch(getThisActivityTerritory(indexRecordProps))
  }
}

export const getActivityTerritories = (props: ActivitiesReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityTerritoriesContract = state.chainContracts.data.contracts.activityTerritories

    let indexRecordProps: ActivityTerritoryProps = {
      activitiesRef: props.activitiesRef,
      activityRef: props.activityRef,
      territoryRef: "",
      data: [],
      successActionType: IATIReportActionTypes.ACTIVITYTERRITORY_SUCCESS,
      failureActionType: IATIReportActionTypes.ACTIVITYTERRITORY_FAILURE
    }

    try {
      const num = await activityTerritoriesContract.getNum(indexRecordProps.activitiesRef, indexRecordProps.activityRef)
      const nums = num.toNumber()
      for (let i = 0; i < nums; i++) {

        indexRecordProps.territoryRef = await activityTerritoriesContract.getReference(indexRecordProps.activitiesRef, indexRecordProps.activityRef, i.toString())
        dispatch(getThisActivityTerritory(indexRecordProps))
      }
    } catch (error) {
      console.log('getActivityTerritory error', error)
    }
  }
}

export const getActivityTerritoryRefs = (props: ActivitiesReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityTerritoriesContract = state.chainContracts.data.contracts.activityTerritories
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef

    let keysData: Array<string> = []
    let actionType = IATIReportActionTypes.ACTIVITYTERRITORYPICKER_SUCCESS

    try {
      const num = await activityTerritoriesContract.getNum(activitiesRef, activityRef)
      const nums = num.toNumber()
      for (let i = 0; i < nums; i++) {
        const territoryRef = await activityTerritoriesContract.getReference(activitiesRef, activityRef, i.toString())
        keysData.push(territoryRef)
      }
    } catch (error) {
      actionType = IATIReportActionTypes.ACTIVITYTERRITORY_FAILURE
      console.log('getActivityTerritoryRefs error', error)
    }

    //console.log('KeysData: ', keysData, actionType)
    dispatch(read({data: keysData})(actionType))
  }
}
