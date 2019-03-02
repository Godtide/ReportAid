import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../../actions'

import { ActionProps, TxReport } from '../../../../types'
import { OrganisationProps, IATIOrganisationProps, IATIWriterActionTypes} from '../../../types'

import { Transaction } from '../../../../../utils/strings'

export const setOrganisation = (details: OrganisationProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const organisationContract = state.chainContracts.data.contracts.organisation

    //console.log('Details: ', details)

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

    //console.log ('OrgsRef: ', details.organisationsRef, 'OrgRef', organisationRef, 'Orgs: ', organisation)

    let actionType = IATIWriterActionTypes.ORGANISATION_FAILURE
    let txData: TxReport = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await organisationContract.setOrganisation(details.organisationsRef, organisationRef, organisation)
      const key = tx.hash
      txData = {
        [key]: {
          summary: `${Transaction.success}`,
          info: tx
        }
      }
      actionType = IATIWriterActionTypes.ORGANISATION_SUCCESS
    } catch (error) {
      txData = {
        [-1]: {
          summary: `${Transaction.fail}`,
          info: {}
        }
      }
      console.log('setOrganisation error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
