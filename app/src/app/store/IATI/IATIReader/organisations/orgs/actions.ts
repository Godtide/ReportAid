import { ThunkDispatch } from 'redux-thunk'

import { storeAction } from '../../../../actions'
import { ApplicationState } from '../../../../store'
import { ActionProps, PayloadProps} from '../../../../types'

import { IATIOrgProps } from '../../../types'
import { IATIOrgReports } from '../types'

import { IATIOrganisationsReportProps, IATIReportActionTypes, OrganisationsReportProps } from '../types'

import { IATIOrgsWriterActionTypes } from './types'

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

    const report = state.orgsData.data as IATIOrgReports
    let actionType = IATIOrgsWriterActionTypes.ORGS_FAILURE
    try {
      const num = await orgsContract.getNumOrgs()
      const numOrgs = num.toNumber()
      for (let i = 0; i < numOrgs; i++) {
         const orgRef = await orgsContract.getOrgReference(i.toString())
         const org: IATIOrgProps = await orgsContract.getOrg(orgRef)
         report[orgRef] = org
         actionType = IATIOrgsWriterActionTypes.ORGS_SUCCESS
      }
    } catch (error) {
      console.log('getBudgets error', error)
    }

    dispatch(read({data: report})(actionType))

  }
}
