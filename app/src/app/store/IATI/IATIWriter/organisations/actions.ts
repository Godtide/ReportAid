import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'
import { OrgReportProps, IATIOrgReportProps } from '../../types'
import { OrgReportsWriterActionTypes } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrgReportsWriterActionTypes): TxProps => {
    const writerProps = storeAction(actionType)(payload) as TxProps
    return writerProps
  }
}

export const setOrganisationReport = (reportDetails: OrgReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const date = new Date()
    const dateTime = date.toISOString()

    const orgReport: IATIOrgReportProps = {
      version: ethers.utils.formatBytes32String(reportDetails.version),
      report: {
        reportRef: ethers.utils.formatBytes32String(shortid.generate()),
        orgRef: reportDetails.orgRef
      },
      reportingOrg: {
        orgRef: reportDetails.reportingOrgRef,
        orgType: reportDetails.reportingOrgType,
        isSecondary: reportDetails.reportingOrgIsSecondary
      },
      lang: ethers.utils.formatBytes32String(reportDetails.lang),
      currency: ethers.utils.formatBytes32String(reportDetails.currency),
      lastUpdatedTime: ethers.utils.formatBytes32String(dateTime)
    }

    const orgReportsContract = state.chainContracts.data.contracts.orgReportsContract
    let actionType = OrgReportsWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      // setReport(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await orgReportsContract.setReport(orgReport)
      const key = tx.hash
      txData[key] = tx
      actionType = OrgReportsWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      console.log('setReport error', error)
    }

    dispatch(add({data: {data: txData}})(actionType))
  }
}
