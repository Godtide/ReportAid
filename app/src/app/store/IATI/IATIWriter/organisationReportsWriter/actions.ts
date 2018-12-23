import { ThunkDispatch } from 'redux-thunk'
import { keccak256 } from 'js-sha3'

import { ApplicationState } from '../../../store'
import { IATIOrganisationReports } from '../../../../../blockchain/typechain/IATIOrganisationReports'

import { storeAction } from '../../../actions'
import { ActionProps, PayloadProps, TxData } from '../../../types'
import { OrgReportProps } from '../../types'

import { OrgReportsActionTypes, OrgReportsProps } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrgReportsActionTypes): OrgReportsProps => {
    const writerProps = storeAction(actionType)(payload) as OrgReportsProps
    return writerProps
  }
}

export const setOrganisationReport = (reportDetails: OrgReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const date = new Date()
    const dateTime = date.toISOString()
    const reference = keccak256(reportDetails.orgIdentifier + reportDetails.reportingOrgIdentifier + dateTime)
    //console.log('Ref: ', reference)
    const orgReportsContract = state.chainContracts.data.contracts.orgReportsContract as IATIOrganisationReports
    let actionType = OrgReportsActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      // setReport(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await orgReportsContract.setReport(reference, reportDetails.orgIdentifier, reportDetails.reportingOrgIdentifier, reportDetails.version, dateTime )
      const key = tx.hash
      txData[key] = tx
      actionType = OrgReportsActionTypes.ADD_SUCCESS
    } catch (error) {
      console.log('setOrgReports error', error)
    }
    //console.log('Adding tx: ', txData, actionType)
    dispatch(add({data: {data: txData}})(actionType))
  }
}
