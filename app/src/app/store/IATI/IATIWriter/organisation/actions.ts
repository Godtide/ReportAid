import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'
import { OrganisationProps, IATIOrganisationProps } from '../../types'
import { OrganisationWriterActionTypes } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrganisationWriterActionTypes): TxProps => {
    const writerProps = storeAction(actionType)(payload) as TxProps
    return writerProps
  }
}

export const setOrganisation = (details: OrganisationProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const organisationContract = state.chainContracts.data.contracts.organisation

    let organisationRef = details.organisationRef
    if ( organisationRef == "" ) {
      organisationRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const organisation: IATIOrganisationProps = {
      orgRef: details.orgRef,
      reportingOrg: {
        orgRef: details.reportingOrgRef,
        orgType: details.reportingOrgType,
        isSecondary: details.reportingOrgIsSecondary
      },
      lang: ethers.utils.formatBytes32String(details.lang),
      currency: ethers.utils.formatBytes32String(details.currency),
      lastUpdatedTime: ethers.utils.formatBytes32String(new Date().toISOString())
    }

    let actionType = OrganisationWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await organisationContract.setOrganisation(details.organisationsRef, organisationRef, organisation)
      const key = tx.hash
      txData[key] = tx
      actionType = OrganisationWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      console.log('set error', error)
    }

    dispatch(add({data: {data: txData}})(actionType))
  }
}
