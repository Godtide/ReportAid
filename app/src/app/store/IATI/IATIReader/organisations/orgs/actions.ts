import { ThunkDispatch } from 'redux-thunk'

import { storeAction } from '../../../../actions'
import { ApplicationState } from '../../../../store'
import { ActionProps, PayloadProps} from '../../../../types'

import { IATIOrgProps } from '../../../IATIWriter/organisations/types'
import { IATIReportActionTypes } from '../../types'
import { IATIOrgReportProps } from '../types'

import { read } from '../../actions'

export const getOrgs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const orgsContract = state.chainContracts.data.contracts.orgs

    let orgsData: IATIOrgReportProps = {
      data: { data: [] }
    }
    let actionType = IATIReportActionTypes.ORGS_FAILURE
    try {
      const num = await orgsContract.getNumOrgs()
      const numOrgs = num.toNumber()
      //console.log("Num orgs: ", numOrgs)
      for (let i = 0; i < numOrgs; i++) {
         const orgRef = await orgsContract.getOrgReference(i.toString())
         //console.log("OrgRef: ", orgRef, orgRef.length, numOrgs)
         const org: IATIOrgProps = await orgsContract.getOrg(orgRef)
         orgsData.data.data[i] = {
           orgRef: orgRef,
           name: org.name,
           identifier: org.identifier
         }

         actionType = IATIReportActionTypes.ORGS_SUCCESS
      }
    } catch (error) {
      console.log('getOrgs error', error)
    }

    dispatch(read({data: orgsData})(actionType))

  }
}
