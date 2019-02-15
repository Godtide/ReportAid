import { ThunkDispatch } from 'redux-thunk'

import { storeAction } from '../../../../actions'
import { ApplicationState } from '../../../../store'
import { ActionProps, PayloadProps} from '../../../../types'

import { IATIOrgProps } from '../../../types'
import { IATIReportActionTypes } from '../types'
import { IATIOrgReportProps } from './types'

const read = (payload: PayloadProps): Function => {
  return (actionType: IATIReportActionTypes): PayloadProps => {
    const getProps = storeAction(actionType)(payload) as PayloadProps
    return getProps
  }
}

export const initialise = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    await dispatch(read({data: { data: {} }})(IATIReportActionTypes.ORGS_INIT))
  }
}

export const getOrgs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const orgsContract = state.chainContracts.data.contracts.orgs

    let orgsData: IATIOrgReportProps = {
      data: {}
    }
    let actionType = IATIReportActionTypes.ORGS_FAILURE
    try {
      const num = await orgsContract.getNumOrgs()
      const numOrgs = num.toNumber()
      //console.log("Num orgs: ", numOrgs)
      for (let i = 0; i < numOrgs; i++) {
         const orgRef = await orgsContract.getOrgReference(i.toString())
         //console.log("OrgRef: ", orgRef)
         const org: IATIOrgProps = await orgsContract.getOrg(orgRef)
         //console.log("Org: ", org)
         orgsData.data[orgRef] = org
         actionType = IATIReportActionTypes.ORGS_SUCCESS
      }
    } catch (error) {
      console.log('getOrgs error', error)
    }

    dispatch(read({data: orgsData})(actionType))

  }
}
