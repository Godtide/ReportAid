import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIOrganisationDocProps } from '../../../IATIWriter/organisations/types'
import { IATIReportActionTypes } from '../../types'
import { IATIOrganisationDocReportProps, OrganisationsReportProps } from '../types'

import { read } from '../actions'

export const initialise = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const initData: any = { data: {} }
    await dispatch(read({data: initData})(IATIReportActionTypes.DOCUMENT_INIT))
  }
}

export const getDocs = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const docsContract = state.chainContracts.data.contracts.organisationDocs
    const organisationsRef = props.organisationsRef
    const organisationRef = props.organisationRef

    let docReports: IATIOrganisationDocReportProps = {
      data: { organisationsRef: organisationsRef,
              organisationRef: organisationRef,
              data: []
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

        docReports.data.data[i] = {
          title: doc.title,
          format: doc.format,
          url: doc.url,
          category: ethers.utils.parseBytes32String(doc.category),
          countryRef: ethers.utils.parseBytes32String(doc.countryRef),
          description: doc.desc,
          language: ethers.utils.parseBytes32String(doc.lang),
          date: ethers.utils.parseBytes32String(doc.date)
        }

        actionType = IATIReportActionTypes.DOCUMENT_SUCCESS
      }
    } catch (error) {
      console.log('getDocs error', error)
    }

    dispatch(read({data: docReports})(actionType))
  }
}
