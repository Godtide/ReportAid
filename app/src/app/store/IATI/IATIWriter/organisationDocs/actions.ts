import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'
import { OrgReportDocProps, IATIOrgReportDocProps } from '../../types'
import { OrgReportDocsWriterActionTypes } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrgReportDocsWriterActionTypes): TxProps => {
    const writerProps = storeAction(actionType)(payload) as TxProps
    return writerProps
  }
}

export const setOrganisationReportDoc = (reportDocDetails: OrgReportDocProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()

    const docDate = new Date(reportDocDetails.year + '/' + reportDocDetails.month + '/' + reportDocDetails.day)
    const reportDocDate = docDate.toISOString()

    //console.log('ReportRef: ', reportDocDetails.reportRef)

    const orgReportDoc: IATIOrgReportDocProps = {
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

    const orgReportDocsContract = state.chainContracts.data.contracts.orgReportDocsContract
    let actionType = OrgReportDocsWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      const tx = await orgReportDocsContract.setDocument(orgReportDoc)
      const key = tx.hash
      txData[key] = tx
      actionType = OrgReportDocsWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      console.log('setReportDoc error', error)
    }

    dispatch(add({data: {data: txData}})(actionType))
  }
}
