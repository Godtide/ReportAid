import { ThunkDispatch } from 'redux-thunk'

import { storeAction } from '../../../../actions'
import { ApplicationState } from '../../../../store'
import { ActionProps, PayloadProps} from '../../../../types'

import { IATIOrgProps } from '../../../types'
import { IATIOrganisationsReportProps, IATIReportActionTypes, OrganisationsReportProps } from '../types'
import { IATIOrgsWriterActionTypes, IATIOrgReports } from './types'

export const read = (payload: PayloadProps): Function => {
  return (actionType: IATIReportActionTypes): IATIOrgReports => {
    const getProps = storeAction(actionType)(payload) as IATIOrgReports
    return getProps
  }
}

export const getOrgs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const orgsContract = state.chainContracts.data.contracts.orgs

    let orgsData: IATIOrgReports = {}
    let actionType = IATIOrgsWriterActionTypes.ORGS_FAILURE
    try {
      const num = await orgsContract.getNumOrgs()
      const numOrgs = num.toNumber()
      console.log("Num orgs: ", numOrgs)
      for (let i = 0; i < numOrgs; i++) {
         const orgRef = await orgsContract.getOrgReference(i.toString())
         console.log("OrgRef: ", orgRef)
         const org: IATIOrgProps = await orgsContract.getOrg(orgRef)
         console.log("Org: ", org)
         orgsData[orgRef] = org
         actionType = IATIOrgsWriterActionTypes.ORGS_SUCCESS
      }
    } catch (error) {
      console.log('getOrgs error', error)
    }

    dispatch(read({data: orgsData})(actionType))

  }
}
