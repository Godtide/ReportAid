import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'
import { IATIOrganisations } from '../../../../../blockchain/typechain/IATIOrganisations'

import { storeAction } from '../../../actions'
import { ActionProps, PayloadProps } from '../../../types'

import { OrgGetActionTypes, OrgGetProps } from './types'

const get = (payload: PayloadProps): Function => {
  return (actionType: OrgGetActionTypes): OrgGetProps => {
    const getProps = storeAction(actionType)(payload) as OrgGetProps
    return getProps
  }
}

export const getNumOrganisations = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgContract = state.chainOrgContract.data.contract as IATIOrganisations
    let actionType = OrgGetActionTypes.GET_NUM_FAILURE
    let getData = {result: {}}
    try {
      const numOrgs = await orgContract.getNumOrganisations()
      const num = numOrgs.toString()
      console.log('Num: ', num)
    } catch (error) {
      console.log(error)
    }
    dispatch(get({data: getData})(actionType))
  }
}
