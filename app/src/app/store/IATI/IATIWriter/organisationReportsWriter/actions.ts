import { ThunkDispatch } from 'redux-thunk'
import shortid from 'shortid'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { IATIOrganisationReports } from '../../../../../blockchain/typechain/IATIOrganisationReports'

import { storeAction } from '../../../actions'
import { ActionProps, PayloadProps, TxData } from '../../../types'
import { OrgReportProps, ReportProps } from '../../types'

import { OrgReportsActionTypes, ReportWriterProps } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrgReportsActionTypes): ReportWriterProps => {
    const writerProps = storeAction(actionType)(payload) as ReportWriterProps
    return writerProps
  }
}

export const setOrganisationReport = (reportDetails: OrgReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const date = new Date()
    const dateTime = date.toISOString()

    const orgReport: ReportProps = {
      version: ethers.utils.formatBytes32String(reportDetails.version),
      orgRef: reportDetails.orgRef,
      reportRef:  ethers.utils.formatBytes32String(shortid.generate()),
      reportingOrg: {
        orgRef: reportDetails.reportingOrgRef,
        orgType: reportDetails.reportingOrgType,
        isSecondary: reportDetails.reportingOrgIsSecondary
      },
      lang: ethers.utils.formatBytes32String(reportDetails.lang),
      currency: ethers.utils.formatBytes32String(reportDetails.currency),
      lastUpdatedTime: ethers.utils.formatBytes32String(dateTime)
    }

    console.log('OrgReport: ', orgReport)

    const orgReportsContract = state.chainContracts.data.contracts.orgReportsContract as IATIOrganisationReports
    let actionType = OrgReportsActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      // setReport(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await orgReportsContract.setReport(orgReport)
      const key = tx.hash
      txData[key] = tx
      actionType = OrgReportsActionTypes.ADD_SUCCESS
    } catch (error) {
      console.log('setReport error', error)
    }
    //console.log('Adding tx: ', txData, actionType)
    dispatch(add({data: {data: txData}})(actionType))
  }
}
