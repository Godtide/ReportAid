import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIActivitySectorProps,
         IATIActivitySectorData,
         IATIReportActionTypes,
         IATIActivitySectorReportProps,
         ActivitiesReportProps } from '../../../types'

import { read } from '../../../../actions'

interface RecordProps {
  activitiesRef: string
  activityRef: string
  sectorRef: string
}

interface ActivitySectorsDataProps {
  data: Array<IATIActivitySectorData>
  successActionType: IATIReportActionTypes
  failureActionType: IATIReportActionTypes
}

type ActivitySectorProps = RecordProps & ActivitySectorsDataProps

const getThisActivitySector = (props: ActivitySectorProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activitySectorsContract = state.chainContracts.data.contracts.activitySectors
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef
    const sectorRef = props.sectorRef
    const data =  props.data

    let sectorData: IATIActivitySectorReportProps = {
      data: { activitiesRef: activitiesRef,
              activityRef: activityRef,
              data: data }
    }

    let actionType =  props.successActionType

    try {

      const thisSector: IATIActivitySectorProps = await activitySectorsContract.getSector(activitiesRef, activityRef, sectorRef)

      sectorData.data.data[data.length] = {
        sectorRef: sectorRef,
        percentage: thisSector.percentage,
        dacCode: ethers.utils.bigNumberify(thisSector.dacCode).toNumber()
      }
    } catch (error) {
      actionType = props.failureActionType
      console.log('getActivitySectors error', error)
    }

    dispatch(read({data: sectorData})(actionType))
  }
}

export const getActivitySectorRecord = (props: RecordProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {

    let indexRecordProps: ActivitySectorProps = {
      activitiesRef: props.activitiesRef,
      activityRef: props.activityRef,
      sectorRef: props.sectorRef,
      data: [],
      successActionType: IATIReportActionTypes.ACTIVITYSECTOR_SUCCESS,
      failureActionType: IATIReportActionTypes.ACTIVITYSECTOR_FAILURE
    }

    dispatch(getThisActivitySector(indexRecordProps))
  }
}

export const getActivitySectors = (props: ActivitiesReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activitySectorsContract = state.chainContracts.data.contracts.activitySectors

    let indexRecordProps: ActivitySectorProps = {
      activitiesRef: props.activitiesRef,
      activityRef: props.activityRef,
      sectorRef: "",
      data: [],
      successActionType: IATIReportActionTypes.ACTIVITYSECTOR_SUCCESS,
      failureActionType: IATIReportActionTypes.ACTIVITYSECTOR_FAILURE
    }

    try {
      const num = await activitySectorsContract.getNum(indexRecordProps.activitiesRef, indexRecordProps.activityRef)
      const nums = num.toNumber()
      for (let i = 0; i < nums; i++) {

        indexRecordProps.sectorRef = await activitySectorsContract.getReference(indexRecordProps.activitiesRef, indexRecordProps.activityRef, i.toString())
        dispatch(getThisActivitySector(indexRecordProps))
      }
    } catch (error) {
      console.log('getActivitySector error', error)
    }
  }
}

export const getActivitySectorRefs = (props: ActivitiesReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activitySectorsContract = state.chainContracts.data.contracts.activitySectors
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef

    let keysData: Array<string> = []
    let actionType = IATIReportActionTypes.ACTIVITYSECTORPICKER_SUCCESS

    try {
      const num = await activitySectorsContract.getNum(activitiesRef, activityRef)
      const nums = num.toNumber()
      for (let i = 0; i < nums; i++) {
        const sectorRef = await activitySectorsContract.getReference(activitiesRef, activityRef, i.toString())
        keysData.push(sectorRef)
      }
    } catch (error) {
      actionType = IATIReportActionTypes.ACTIVITYSECTOR_FAILURE
      console.log('getActivitySectorRefs error', error)
    }

    //console.log('KeysData: ', keysData, actionType)
    dispatch(read({data: keysData})(actionType))
  }
}
