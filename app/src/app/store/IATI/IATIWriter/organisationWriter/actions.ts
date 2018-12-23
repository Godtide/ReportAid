import { ThunkDispatch } from 'redux-thunk'
import { keccak256 } from 'js-sha3'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../store'
import { IATIOrganisations } from '../../../../../blockchain/typechain/IATIOrganisations'

import { storeAction } from '../../../actions'
import { ActionProps, PayloadProps, TxData } from '../../../types'
import { OrganisationProps } from '../../types'

import { OrgActionTypes, OrgProps } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrgActionTypes): OrgProps => {
    const writerProps = storeAction(actionType)(payload) as OrgProps
    return writerProps
  }
}

export const setOrganisation = (orgDetails: OrganisationProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const stringIdentifier =  ethers.utils.formatBytes32String(orgDetails.code + '-' + orgDetails.identifier)
    const reference = ethers.utils.soliditySha256('string', stringIdentifier) as string
    const identifier = ethers.utils.formatBytes32String(stringIdentifier)
    const name = ethers.utils.formatBytes32String(orgDetails.name)
    console.log('Ref: ', reference, 'name: ', name, 'ident: ', identifier)
    const orgContract = state.chainContracts.data.contracts.orgContract as IATIOrganisations
    let actionType = OrgActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      const tx = await orgContract.setOrganisation(reference, name, identifier)
      const key = tx.hash
      txData[key] = tx
      actionType = OrgActionTypes.ADD_SUCCESS
    } catch (error) {
      console.log('setOrganisation error', error)
    }
    //console.log('Adding tx: ', txData, actionType)
    dispatch(add({data: {data: txData}})(actionType))
  }
}
