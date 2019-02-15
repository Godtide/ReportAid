import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIOrganisationDocProps } from '../../../types'
import { IATIReportActionTypes, OrganisationsReportProps } from '../types'
import { IATIOrganisationDocReportProps } from './types'

import { read } from '../actions'

export const initialise = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const initData: IATIOrganisationDocReportProps = { data: {} }
    dispatch(read({data: initData})(IATIReportActionTypes.DOCUMENT_INIT))
  }
}

export const getDocs = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const docsContract = state.chainContracts.data.contracts.organisationDocs
    const organisationsRef = props.organisationsRef
    const organisationRef = props.organisationRef

    let docReports: IATIOrganisationDocReportProps = {
      data: {
        [organisationsRef]: {
          data: {
            [organisationRef]: {
              data: {}
            }
          }
        }
      }
    }
    let actionType = IATIReportActionTypes.DOCUMENT_FAILURE
    try {
      const num = await docsContract.getNumDocs(props.organisationsRef, props.organisationRef)
      const numDocs = num.toNumber()
      for (let i = 0; i < numDocs; i++) {
        const docRef = await docsContract.getDocReference(organisationsRef,
                                                           organisationRef,
                                                           i.toString())
        const doc: IATIOrganisationDocProps = await docsContract.getDocument(organisationsRef,
                                                                                organisationRef,
                                                                                docRef)

        docReports.data[organisationsRef].data[organisationRef].data[docRef] = doc
        actionType = IATIReportActionTypes.DOCUMENT_SUCCESS
      }
    } catch (error) {
      console.log('getDocs error', error)
    }

    dispatch(read({data: docReports})(actionType))
  }
}
