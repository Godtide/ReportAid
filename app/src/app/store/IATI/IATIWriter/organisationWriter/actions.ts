import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'
import { IATIOrganisations } from '../../../../../blockchain/typechain/IATIOrganisations'

import { storeAction } from '../../../actions'
import { ActionProps, PayloadProps } from '../../../types'
import { OrganisationProps } from '../../types'

import { OrgWriterActionTypes, OrgWriterProps } from './types'

const add = (payload: PayloadProps): Function => {
  return (actionType: OrgWriterActionTypes): OrgWriterProps => {
    const writerProps = storeAction(actionType)(payload) as OrgWriterProps
    return writerProps
  }
}

export const setOrganisation = (orgDetails: OrganisationProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    //console.log('In setOrg')
    const state = getState()
    const orgContract = state.chainOrgContract.data.contract as IATIOrganisations
    let actionType = OrgWriterActionTypes.ADD_FAILURE
    let addData = {result: {}}
    try {
      const tx = await orgContract.setOrganisation(orgDetails.name, orgDetails.reference, orgDetails.type)
      addData = {result: {tx}}
      actionType = OrgWriterActionTypes.ADD_SUCCESS
    } catch (error) {
      console.log('In error: ', error)
    }
    //console.log('Adding tx: ', addData, actionType)
    dispatch(add({data: addData})(actionType))
  }
}
