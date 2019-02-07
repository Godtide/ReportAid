import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps, TxProps, TxData } from '../../../types'
import { OrganisationsProps, IATIOrganisationsProps } from '../../types'
import { OrganisationsWriterActionTypes } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrganisationsWriterActionTypes): TxProps => {
    const writerProps = storeAction(actionType)(payload) as TxProps
    return writerProps
  }
}

export const setOrganisations = (details: OrgProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const organisationsContract = state.chainContracts.data.contracts.organisations
    const organisationsRef = ethers.utils.formatBytes32String(shortid.generate())

    const organisations: IATIOrganisationsProps = {
      version: ethers.utils.formatBytes32String(details.version),
      generatedTime: ethers.utils.formatBytes32String(new Date().toISOString())
    }

    let actionType = OrganisationsWriterActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      // set(bytes32 _reference, bytes32 _orgRef, bytes32 _reportingOrgRef, bytes32 _version, bytes32 _generatedTime)
      const tx = await organisationsContract.setOrganisations(organisationsRef, organisations)
      const key = tx.hash
      txData[key] = tx
      actionType = OrganisationsWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      console.log('set error', error)
    }

    dispatch(add({data: {data: txData}})(actionType))
  }
}
