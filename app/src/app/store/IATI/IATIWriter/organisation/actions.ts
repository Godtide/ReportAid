import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'
import { OrgProps, IATIOrgProps } from '../../types'
import { OrganisationWriterActionTypes } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrganisationWriterActionTypes): TxProps => {
    const writerProps = storeAction(actionType)(payload) as TxProps
    return writerProps
  }
}

export const setOrganisation = (reportDetails: OrgProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const date = new Date()
    const dateTime = date.toISOString()

    const org: IATIOrgProps = {
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

    const orgsContract = state.chainContracts.data.contracts.orgsContract
    let actionType = OrganisationWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await orgsContract.set(org)
      const key = tx.hash
      txData[key] = tx
      actionType = OrganisationWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      console.log('set error', error)
    }

    dispatch(add({data: {data: txData}})(actionType))
  }
}
