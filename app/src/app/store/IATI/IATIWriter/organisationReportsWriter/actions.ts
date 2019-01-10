import { ThunkDispatch } from 'redux-thunk'
import shortid from 'shortid'

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
    const reportReference = shortid.generate()

    const orgReport: ReportProps = {
      reportRef: reportReference,
      reportingOrg: {
        orgIdentifier: reportDetails.reportingOrgIdentifier,
        orgType: reportDetails.reportingOrgType,
        isSecondary: reportDetails.reportingOrgIsSecondary
      },
      issuingOrgRef: reportDetails.issuingOrgRef,
      version: reportDetails.version,
      lang: reportDetails.lang,
      currency: reportDetails.currency,
      generatedTime: dateTime,
      lastUpdatedTime: dateTime
    }

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
