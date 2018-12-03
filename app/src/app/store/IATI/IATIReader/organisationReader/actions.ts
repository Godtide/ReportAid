import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'
import { IATIOrganisations } from '../../../../../blockchain/typechain/IATIOrganisations'

import { storeAction } from '../../../actions'
import { ActionProps, PayloadProps } from '../../../types'

import { OrgGetActionTypes, OrgGetProps, OrgData } from './types'

const get = (payload: PayloadProps): Function => {
  return (actionType: OrgGetActionTypes): OrgGetProps => {
    const getProps = storeAction(actionType)(payload) as OrgGetProps
    return getProps
  }
}

export const getOverview = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNumOrganisations())
    await dispatch(getReferences())
    await dispatch(getNames())
    dispatch(getTypes())
  }
}

export const getNumOrganisations = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgContract = state.chainOrgContract.data.contract as IATIOrganisations
    let actionType = OrgGetActionTypes.NUM_FAILURE
    let numOrgs = { num: 0 }
    try {
      const num = await orgContract.getNumOrganisations()
      numOrgs.num = num.toNumber()
      actionType = OrgGetActionTypes.NUM_SUCCESS
    } catch (error) {
      console.log('getNumOrganisations error', error)
    }
    dispatch(get({data: numOrgs})(actionType))
  }
}

export const getReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgContract = state.chainOrgContract.data.contract as IATIOrganisations
    const numOrgs = state.orgReader.num
    let orgRefs: OrgData = {}
    let actionType = OrgGetActionTypes.REF_FAILURE
    for (let i = 0; i < numOrgs; i++) {
       try {
         const ref = await orgContract.getOrganisationReference(i.toString())
         //console.log('Ref', ref)
         orgRefs[ref] = {
           name: '',
           reference: ref,
           type: ''
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

export const getNames = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgContract = state.chainOrgContract.data.contract as IATIOrganisations
    let actionType = OrgGetActionTypes.NAME_FAILURE
    const orgs = state.orgReader.data
    const orgKeys = Object.keys(orgs)
    //console.log('Orgkeys: ', orgKeys)
    for (let i = 0; i < orgKeys.length; i++) {
      const thisKey = orgKeys[i]
       try {
         orgs[thisKey].name = await orgContract.getOrganisationName(thisKey)
         actionType = OrgGetActionTypes.NAME_SUCCESS
       } catch (error) {
         console.log('getNames error', error)
       }
    }
    //console.log('New Orgs; ', orgs)
    dispatch(get({data: {data: orgs}})(actionType))
  }
}

export const getTypes = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgContract = state.chainOrgContract.data.contract as IATIOrganisations
    let actionType = OrgGetActionTypes.TYPE_FAILURE
    const orgs = state.orgReader.data
    const orgKeys = Object.keys(orgs)
    for (let i = 0; i < orgKeys.length; i++) {
      const thisKey = orgKeys[i]
       try {
         orgs[thisKey].type = await orgContract.getOrganisationType(thisKey)
         actionType = OrgGetActionTypes.TYPE_SUCCESS
       } catch (error) {
         console.log('getTypes error', error)
       }
    }
    //console.log('New Type; ', orgs)
    dispatch(get({data: {data: orgs}})(actionType))
  }
}
