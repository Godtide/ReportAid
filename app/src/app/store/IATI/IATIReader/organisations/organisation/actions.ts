import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'

import { read } from '../actions'

import { IATIOrganisationProps } from '../../../types'
import { IATIReportActionTypes, IATIOrganisationReportProps, OrganisationReportProps } from '../types'

export const getOrganisation = (props: OrganisationReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const organisationContract = state.chainContracts.data.contracts.organisation
    const organisationsRef = props.organisationsRef
    const report = state.organisationsReport.data[organisationsRef].data as IATIOrganisationReportProps

    let actionType = IATIReportActionTypes.ORGANISATION_FAILURE
    try {
      const num = await organisationContract.getNumOrganisations(organisationsRef)
      const numOrganisations = num.toNumber()
      for (let i = 0; i < numOrganisations; i++) {
         const orgRef = await organisationContract.getOrganisationReference(organisationsRef, i.toString())
         const org: IATIOrganisationProps = await organisationContract.getOrganisation(organisationsRef, orgRef)
         report[orgRef].IATIOrganisation = org
         actionType = IATIReportActionTypes.ORGANISATION_SUCCESS
      }
    } catch (error) {
      console.log('getOrganisation error', error)
    }

    dispatch(read({data: report})(actionType))
  }
}
