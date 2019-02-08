import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'

import { write } from '../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'
import { OrganisationDocProps, IATIOrganisationDocProps } from '../../types'
import { IATIWriterActionTypes } from '../types'

export const setOrganisationDoc = (details: OrganisationDocProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const docsContract = state.chainContracts.data.contracts.organisactionDocs

    const docDate = new Date(details.year + '/' + details.month + '/' + details.day)

    let docRef = details.docRef
    if ( docRef == "" ) {
      docRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const doc: IATIOrganisationDocProps = {
      title: details.title,
      format: details.format,
      url: details.url,
      category: ethers.utils.formatBytes32String(details.category),
      countryRef: ethers.utils.formatBytes32String(details.countryRef),
      desc: details.desc,
      lang: ethers.utils.formatBytes32String(details.lang),
      date: ethers.utils.formatBytes32String(docDate.toISOString())
    }

    let actionType = IATIWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      const tx = await orgDocsContract.setDocument(details.organisationsRef,
                                                   details.organisationRef,
                                                   docRef,
                                                   doc)
      const key = tx.hash
      txData[key] = tx
      actionType = IATIWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      console.log('setDoc error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
