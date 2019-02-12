import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIOrganisationDocProps } from '../../../types'
import { IATIReportActionTypes,
         IATIDocumentReportProps,
         OrganisationsReportProps } from '../types'

import { read } from '../actions'

export const getDocs = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const docsContract = state.chainContracts.data.contracts.organisationDocs
    const organisationsRef = props.organisationsRef
    const organisationRef = props.organisationRef

    const docReports = state.organisationsReport.data[organisationsRef].data[organisationRef].data.document as IATIDocumentReportProps

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
         docReports[docRef] = doc
         actionType = IATIReportActionTypes.DOCUMENT_SUCCESS
      }
    } catch (error) {
      console.log('getDocs error', error)
    }

    dispatch(read({data: docReports})(actionType))
  }
}
