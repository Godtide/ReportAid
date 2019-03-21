import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIOrganisationsProps,
         IATIOrganisationsData,
         IATIReportActionTypes,
         IATIOrganisationsReportProps } from '../../../types'

import { read } from '../../../../actions'

interface RecordProps {
  organisationsRef: string
}

interface OrganisationsDataProps {
  data: Array<IATIOrganisationsData>
  successActionType: IATIReportActionTypes
  failureActionType: IATIReportActionTypes
}

type OrganisationsProps = RecordProps & OrganisationsDataProps

const getThisOrganisationsRecord = (props: OrganisationsProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const organisationsContract = state.chainContracts.data.contracts.organisations
    const organisationsRef = props.organisationsRef
    const data =  props.data

    let organisationsData: IATIOrganisationsReportProps = {
      data: { data: data }
    }

    let actionType =  props.successActionType
    try {
      const organisations: IATIOrganisationsProps = await organisationsContract.getOrganisations(organisationsRef)
      organisationsData.data.data[data.length] = {
        organisationsRef: organisationsRef,
        version: ethers.utils.parseBytes32String(organisations.version),
        generatedTime:  ethers.utils.parseBytes32String(organisations.generatedTime)
      }
    } catch (error) {
      actionType = props.failureActionType
      console.log('getOrganisations error', error)
    }

    //console.log('Getting activities: ', activitiesData)

    dispatch(read({data: organisationsData})(actionType))
  }
}

export const getOrganisationsRecord = (props: RecordProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    let indexRecordProps: OrganisationsProps = {
      organisationsRef: props.organisationsRef,
      data: [],
      successActionType: IATIReportActionTypes.ORGANISATIONS_SUCCESS,
      failureActionType: IATIReportActionTypes.ORGANISATIONS_FAILURE
    }

    dispatch(getThisOrganisationsRecord(indexRecordProps))
  }
}

export const getOrganisations = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const organisationsContract = state.chainContracts.data.contracts.organisations

    let indexRecordProps: OrganisationsProps = {
      organisationsRef: "",
      data: [],
      successActionType: IATIReportActionTypes.ORGANISATIONS_SUCCESS,
      failureActionType: IATIReportActionTypes.ORGANISATIONS_FAILURE
    }

    try {
      const num = await organisationsContract.getNumOrganisations()
      const numOrganisations = num.toNumber()
      for (let i = 0; i < numOrganisations; i++) {
        indexRecordProps.organisationsRef = await organisationsContract.getOrganisationsReference(i.toString())
        dispatch(getThisOrganisationsRecord(indexRecordProps))
      }
    } catch (error) {
      console.log('getOrganisations error', error)
    }
  }
}

export const getOrganisationsRefs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const organisationsContract = state.chainContracts.data.contracts.organisations

    let keysData: Array<string> = []
    let actionType = IATIReportActionTypes.ORGANISATIONSPICKER_SUCCESS

    try {
      const num = await organisationsContract.getNumOrganisations()
      const numOrganisations = num.toNumber()
      for (let i = 0; i < numOrganisations; i++) {
        const organisationsRef = await organisationsContract.getOrganisationsReference(i.toString())
        keysData.push(organisationsRef)
      }
    } catch (error) {
      actionType = IATIReportActionTypes.ORGANISATIONS_FAILURE
      console.log('getOrganisationsRefs error', error)
    }

    //console.log('KeysData: ', keysData, actionType)
    dispatch(read({data: keysData})(actionType))
  }
}
