import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIOrganisationProps } from '../../../types'
import { IATIReportActionTypes, IATIOrganisationReport, OrganisationReportProps, IATIOrganisationsData } from '../types'

import { read } from '../actions'

export const getOrganisation = (props: OrganisationReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const organisationContract = state.chainContracts.data.contracts.organisation
    const organisationsRef = props.organisationsRef

    let organisationsData: IATIOrganisationsData = state.organisationsReader.data
    //console.log('orgs stuff: ', organisationsData)

    let actionType = IATIReportActionTypes.ORGANISATION_FAILURE
    try {
      const num = await organisationContract.getNumOrganisations(organisationsRef)
      const numOrganisations = num.toNumber()
      for (let i = 0; i < numOrganisations; i++) {
         const organisationRef = await organisationContract.getOrganisationReference(organisationsRef, i.toString())
         const organisation: IATIOrganisationProps = await organisationContract.getOrganisation(organisationsRef, organisationRef)
         // console.log('Org: ', org)
         // IATIOrganisations: state.organisationsReader.data[organisationsRef].IATIOrganisations
         organisationsData[organisationsRef].data[organisationRef] = {
           IATIOrganisation: organisation,
           data: {
               totalBudget: {
                 '': {
                   budgetLine: '',
                   finance: {
                     value: 0,
                     status: 0,
                     start: '',
                     end: ''
                   }
                 }
               },
               recipientOrgBudget: {
                 '': {
                   recipientOrgRef: '',
                   budgetLine: '',
                   finance: {
                     value: 0,
                     status: 0,
                     start: '',
                     end: ''
                   }
                 }
               },
               recipientRegionBudget: {
                 '': {
                   regionRef: 0,
                   budgetLine: '',
                   finance: {
                     value: 0,
                     status: 0,
                     start: '',
                     end: ''
                   }
                 }
               },
               recipientCountryBudget: {
                 '': {
                   countryRef: '',
                   budgetLine: '',
                   finance: {
                     value: 0,
                     status: 0,
                     start: '',
                     end: ''
                   }
                 }
               },
               totalExpenditure: {
                 '': {
                   expenditureLine: '',
                   finance: {
                     value: 0,
                     status: 0,
                     start: '',
                     end: ''
                   }
                 }
               },
               document: {
                 '': {
                   title: '',
                   format: '',
                   url: '',
                   category: '',
                   countryRef: '',
                   desc: '',
                   lang: '',
                   date: ''
                 }
               }
             }
         }
         actionType = IATIReportActionTypes.ORGANISATION_SUCCESS
      }
    } catch (error) {
      console.log('getOrganisation error', error)
    }

    dispatch(read({data: organisationsData})(actionType))
  }
}
