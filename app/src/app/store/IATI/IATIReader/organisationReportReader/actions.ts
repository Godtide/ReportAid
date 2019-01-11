import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'
import { IATIOrganisations } from '../../../../../blockchain/typechain/IATIOrganisations'

import { storeAction } from '../../../actions'
import { ActionProps, PayloadProps } from '../../../types'

import { OrgGetActionTypes, OrgReportGetProps, OrgReportData } from './types'

export const getOrgReports = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNumOrgReports())
    await dispatch(getReferences())
    dispatch(getOrgReportData())
    /* await dispatch(getNames())
    dispatch(getIDs())*/
  }
}

const get = (payload: PayloadProps): Function => {
  return (actionType: OrgReportGetActionTypes): OrgReportGetProps => {
    const getProps = storeAction(actionType)(payload) as OrgReportGetProps
    return getProps
  }
}

const getNumOrgReports = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportsContract = state.chainContracts.data.contracts.orgReportsContract as IATIOrganisationReports
    let actionType = OrgReportGetActionTypes.NUM_FAILURE
    let numOrgs = { num: 0 }
    try {
      const num = await orgContract.getNumOrganisations()
      //console.log('Num orgs: ', num)
      numOrgs.num = num.toNumber()
      actionType = OrgReportGetActionTypes.NUM_SUCCESS
    } catch (error) {
      console.log('getNumOrganisations error', error)
    }
    dispatch(get({data: numOrgs})(actionType))
  }
}

const getReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportsContract = state.chainContracts.data.contracts.orgReportsContract as IATIOrganisationReports
    const numOrgs = state.orgReader.num
    let orgRefs: OrgData = {}
    let actionType = OrgGetActionTypes.REF_FAILURE
    for (let i = 0; i < numOrgs; i++) {
       try {
         const ref = await orgContract.getOrganisationReference(i.toString())
         //console.log('Ref', ref)
         orgRefs[ref] = {
           orgRef: ref,
           name: '',
           identifier: ''
         }
         actionType = OrgGetActionTypes.REF_SUCCESS
       } catch (error) {
         console.log('getReferences error', error)
       }
    }
    //console.log('OrgRefs: ', orgRefs)
    dispatch(get({data: {data: orgRefs}})(actionType))
  }
}

const getOrgData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgContract = state.chainContracts.data.contracts.orgContract as IATIOrganisations
    let actionType = OrgGetActionTypes.ORG_FAILURE
    const orgs = state.orgReader.data
    const orgKeys = Object.keys(orgs)
    //console.log('Orgkeys: ', orgKeys)
    for (let i = 0; i < orgKeys.length; i++) {
      const thisKey = orgKeys[i]
       try {
         const data = await orgContract.getOrganisation(thisKey)
         //console.log ('Org stuff: ', data)
         orgs[thisKey].name = data[1]
         orgs[thisKey].identifier = data[2]
         actionType = OrgGetActionTypes.ORG_SUCCESS
       } catch (error) {
         console.log('getOrgs error', error)
       }
    }
    //console.log('New Orgs; ', orgs)
    dispatch(get({data: {data: orgs}})(actionType))
  }
}
