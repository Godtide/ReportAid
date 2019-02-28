import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIOrganisationProps } from '../../../IATIWriter/organisations/types'
import { IATIReportActionTypes } from '../../types'
import { IATIOrganisationReportProps, OrganisationReportProps } from '../types'

import { read } from '../../actions'

export const initialise = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const initData: any = { data: {} }
    await dispatch(read({data: initData})(IATIReportActionTypes.ORGANISATION_INIT))
  }
}

export const getOrganisation = (props: OrganisationReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const organisationContract = state.chainContracts.data.contracts.organisation
    const organisationsRef = props.organisationsRef

    let organisationData: IATIOrganisationReportProps = {
      data: { organisationsRef: organisationsRef, data: [] }
    }

    let actionType = IATIReportActionTypes.ORGANISATION_FAILURE
    try {
      const num = await organisationContract.getNumOrganisations(organisationsRef)
      const numOrganisations = num.toNumber()
      for (let i = 0; i < numOrganisations; i++) {
         const organisationRef = await organisationContract.getOrganisationReference(organisationsRef, i.toString())
         const organisation: IATIOrganisationProps = await organisationContract.getOrganisation(organisationsRef, organisationRef)

         organisationData.data.data[i] = {
           organisationRef: organisationRef,
           issuingOrgRef: organisation.orgRef,
           reportingOrgRef: organisation.reportingOrg.orgRef,
           reportingOrgType: organisation.reportingOrg.orgType,
           reportingOrgIsSecondary: organisation.reportingOrg.isSecondary,
           lang: ethers.utils.parseBytes32String(organisation.lang),
           currency: ethers.utils.parseBytes32String(organisation.currency),
           lastUpdatedTime: ethers.utils.parseBytes32String(organisation.lastUpdatedTime)
         }
         //console.log(organisationsRef, organisationData.data[organisationsRef].data[organisationRef] )
         actionType = IATIReportActionTypes.ORGANISATION_SUCCESS
      }
    } catch (error) {
      console.log('getOrganisation error', error)
    }

    dispatch(read({data: organisationData})(actionType))
  }
}
