import { ThunkDispatch } from 'redux-thunk'
import { keccak256 } from 'js-sha3'

import { ApplicationState } from '../../../store'
import { IATIOrganisationReports } from '../../../../../blockchain/typechain/IATIOrganisationReports'

import { storeAction } from '../../../actions'
import { ActionProps, PayloadProps, TxData } from '../../../types'
import { OrganisationProps } from '../../types'

import { OrgReportsActionTypes, OrgReportsProps } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrgReportsActionTypes): OrgReportsProps => {
    const writerProps = storeAction(actionType)(payload) as OrgReportsProps
    return writerProps
  }
}

export const setOrganisationReport = (orgDetails: OrganisationProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const reference = keccak256(orgDetails.code + '-' + orgDetails.identifier)
    //console.log('Ref: ', reference)
    const orgReportsContract = state.chainContracts.data.contracts.orgReportsContract as IATIOrganisationReports
    let actionType = OrgReportsActionTypes.ADD_FAILURE
    let txData: TxData = {}
    try {
      //const tx = await orgContract.setOrganisation(reference, orgDetails.name, orgDetails.code, orgDetails.identifier)
      //const key = tx.hash
      //txData[key] = tx
      actionType = OrgReportsActionTypes.ADD_SUCCESS
    } catch (error) {
      console.log('setOrgReports error', error)
    }
    //console.log('Adding tx: ', txData, actionType)
    dispatch(add({data: {data: txData}})(actionType))
  }
}
