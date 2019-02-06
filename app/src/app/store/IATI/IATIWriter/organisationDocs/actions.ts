import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'
import { OrgDocProps, IATIOrgDocProps } from '../../types'
import { OrgDocsWriterActionTypes } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrgDocsWriterActionTypes): TxProps => {
    const writerProps = storeAction(actionType)(payload) as TxProps
    return writerProps
  }
}

export const setOrganisationDoc = (reportDocDetails: OrgDocProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()

    const docDate = new Date(reportDocDetails.year + '/' + reportDocDetails.month + '/' + reportDocDetails.day)
    const reportDocDate = docDate.toISOString()

    //console.log('Ref: ', reportDocDetails.reportRef)

    const orgDoc: IATIOrgDocProps = {
      report: {
        reportRef: reportDocDetails.report.reportRef,
        orgRef: reportDocDetails.report.orgRef
      },
      docRef: ethers.utils.formatBytes32String(shortid.generate()),
      title: reportDocDetails.title,
      format: reportDocDetails.format,
      url: reportDocDetails.url,
      category: ethers.utils.formatBytes32String(reportDocDetails.category),
      countryRef: ethers.utils.formatBytes32String(reportDocDetails.countryRef),
      desc: reportDocDetails.desc,
      lang: ethers.utils.formatBytes32String(reportDocDetails.lang),
      date: ethers.utils.formatBytes32String(reportDocDate)
    }

    const orgDocsContract = state.chainContracts.data.contracts.orgDocsContract
    let actionType = OrgDocsWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      const tx = await orgDocsContract.setDocument(orgDoc)
      const key = tx.hash
      txData[key] = tx
      actionType = OrgDocsWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      console.log('setDoc error', error)
    }

    dispatch(add({data: {data: txData}})(actionType))
  }
}
