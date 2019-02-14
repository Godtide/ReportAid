import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIOrganisationProps } from '../../../types'
import { IATIReportActionTypes, OrganisationReportProps } from '../types'
import { IATIOrganisationReportProps } from './types'

import { read } from '../actions'

export const getOrganisation = (props: OrganisationReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const organisationContract = state.chainContracts.data.contracts.organisation
    const organisationsRef = props.organisationsRef

    let organisationData: IATIOrganisationReportProps = {
      data: {
        [organisationsRef]: {
          data: {}
        }
      }
    }

    let actionType = IATIReportActionTypes.ORGANISATION_FAILURE
    try {
      const num = await organisationContract.getNumOrganisations(organisationsRef)
      const numOrganisations = num.toNumber()
      for (let i = 0; i < numOrganisations; i++) {
         const organisationRef = await organisationContract.getOrganisationReference(organisationsRef, i.toString())
         const organisation: IATIOrganisationProps = await organisationContract.getOrganisation(organisationsRef, organisationRef)
         organisationData.data[organisationsRef].data[organisationRef] = organisation
         actionType = IATIReportActionTypes.ORGANISATION_SUCCESS
      }
    } catch (error) {
      console.log('getOrganisation error', error)
    }

    dispatch(read({data: organisationData})(actionType))
  }
}
