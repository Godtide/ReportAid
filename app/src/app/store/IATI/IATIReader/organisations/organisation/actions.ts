import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIOrganisationProps,
         IATIOrganisationData,
         IATIReportActionTypes,
         IATIOrganisationReportProps,
         OrganisationReportProps } from '../../../types'

import { read } from '../../actions'

interface RecordProps {
  organisationsRef: string
  organisationRef: string
}

interface OrganisationDataProps {
  data: Array<IATIOrganisationData>
  successActionType: IATIReportActionTypes
  failureActionType: IATIReportActionTypes
}

type OrganisationProps = RecordProps & OrganisationDataProps

const getThisOrganisation = (props: OrganisationProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const organisationContract = state.chainContracts.data.contracts.organisation
    const organisationsRef = props.organisationsRef
    const organisationRef = props.organisationRef
    const data =  props.data

    //let organisationData: IATIOrganisationReportProps = {...state.}
    let organisationData: IATIOrganisationReportProps = {
      data: { organisationsRef: organisationsRef, data: data }
    }

    let actionType = props.failureActionType

    try {
      const organisation: IATIOrganisationProps = await organisationContract.getOrganisation(organisationsRef, organisationRef)

      //console.log(organisationData.data.data.length)
      organisationData.data.data[data.length] = {
        organisationRef: organisationRef,
        issuingOrgRef: organisation.orgRef,
        reportingOrgRef: organisation.reportingOrg.orgRef,
        reportingOrgType: organisation.reportingOrg.orgType,
        reportingOrgIsSecondary: organisation.reportingOrg.isSecondary,
        lang: ethers.utils.parseBytes32String(organisation.lang),
        currency: ethers.utils.parseBytes32String(organisation.currency),
        lastUpdatedTime: ethers.utils.parseBytes32String(organisation.lastUpdatedTime)
      }

      actionType = props.successActionType
    } catch (error) {
      console.log('getOrganisation error', error)
    }

    dispatch(read({data: organisationData})(actionType))
  }
}

export const getOrganisationRecord = (props: RecordProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {

    let indexRecordProps: OrganisationProps = {
      organisationsRef: props.organisationsRef,
      organisationRef: props.organisationRef,
      data: [],
      successActionType: IATIReportActionTypes.ORGANISATION_SUCCESS,
      failureActionType: IATIReportActionTypes.ORGANISATION_FAILURE
    }

    dispatch(getThisOrganisation(indexRecordProps))
  }
}

export const getOrganisation = (props: OrganisationReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const organisationContract = state.chainContracts.data.contracts.organisation

    let indexRecordProps: OrganisationProps = {
      organisationsRef: props.organisationsRef,
      organisationRef: "",
      data: [],
      successActionType: IATIReportActionTypes.ORGANISATION_SUCCESS,
      failureActionType: IATIReportActionTypes.ORGANISATION_FAILURE
    }

    try {
      const num = await organisationContract.getNumOrganisations(indexRecordProps.organisationsRef)
      const numOrganisations = num.toNumber()
      for (let i = 0; i < numOrganisations; i++) {
        indexRecordProps.organisationRef = await organisationContract.getOrganisationReference(indexRecordProps.organisationsRef, i.toString())
        dispatch(getThisOrganisation(indexRecordProps))
      }
    } catch (error) {
      console.log('getOrganisation error', error)
    }
  }
}

export const getOrganisationRefs = (props: OrganisationReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const organisationContract = state.chainContracts.data.contracts.organisation
    const organisationsRef = props.organisationsRef

    let keysData: Array<string> = []
    let actionType = IATIReportActionTypes.ORGANISATIONPICKER_SUCCESS
    try {
      const num = await organisationContract.getNumOrganisations(organisationsRef)
      const numOrganisations = num.toNumber()
      for (let i = 0; i < numOrganisations; i++) {
         const organisationRef = await organisationContract.getOrganisationReference(organisationsRef, i.toString())
         keysData.push(organisationRef)
      }
    } catch (error) {
      actionType = IATIReportActionTypes.ORGANISATION_FAILURE
      console.log('getOrganisationRefs error', error)
    }

    dispatch(read({data: keysData})(actionType))
  }
}
