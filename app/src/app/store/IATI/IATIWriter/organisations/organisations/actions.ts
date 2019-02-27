import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../actions'

import { ActionProps, PayloadProps, TxProps, TxReport } from '../../../../types'
import { OrganisationsProps, IATIOrganisationsProps } from '../types'
import { IATIWriterActionTypes } from '../../types'

import { Transaction } from '../../../../../utils/strings'

export const setOrganisations = (details: OrganisationsProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const organisationsContract = state.chainContracts.data.contracts.organisations
    const organisationsRef = ethers.utils.formatBytes32String(shortid.generate())

    const organisations: IATIOrganisationsProps = {
      version: ethers.utils.formatBytes32String(details.version),
      generatedTime: ethers.utils.formatBytes32String(new Date().toISOString())
    }

    let actionType = IATIWriterActionTypes.ORGANISATIONS_FAILURE
    let txData: TxReport = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await organisationsContract.setOrganisations(organisationsRef, organisations)
      const key = tx.hash
      txData = {
        [key]: {
          summary: `${Transaction.success}`,
          info: tx
        }
      }
      actionType = IATIWriterActionTypes.ORGANISATIONS_SUCCESS
    } catch (error) {
      txData = {
        [-1]: {
          summary: `${Transaction.fail}`,
          info: {}
        }
      }
      console.log('set error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
