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
    let getData = {num: 0}
    try {
      const numOrgs = await orgContract.getNumOrganisations()
      const num = numOrgs.toString()
      getData = {num: num}
      actionType = OrgGetActionTypes.NUM_SUCCESS
    } catch (error) {
      console.log(error)
    }
    dispatch(get({data: getData})(actionType))
  }
}

export const getReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgContract = state.chainOrgContract.data.contract as IATIOrganisations
    const numOrgs = state.orgReader.data.num
    //console.log('Num Orgs is: ', numOrgs)
    let actionType = OrgGetActionTypes.REF_FAILURE
    let refs = []
    for (let i = 0; i < numOrgs; i++) {
       try {
         const ref = await orgContract.getOrganisationReference(i.toString())
         refs.push(ref)
         //console.log('ref: ', ref)
         actionType = OrgGetActionTypes.REF_SUCCESS
       } catch (error) {
         OrgGetActionTypes.REF_FAILURE
         console.log(error)
       }
    }
    const getData = {refs: refs}
    dispatch(get({data: getData})(actionType))
  }
}

export const getNames = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgContract = state.chainOrgContract.data.contract as IATIOrganisations
    const refs = state.orgReader.data.refs
    const refsLength = refs.length
    //console.log('Num Orgs is: ', numOrgs)
    let actionType = OrgGetActionTypes.NAME_FAILURE
    let names = []
    for (let i = 0; i < refsLength; i++) {
       try {
         const name = await orgContract.getOrganisationName(refs[i])
         names.push(name)
         //console.log('blah: ', name)
         actionType = OrgGetActionTypes.NAME_SUCCESS
       } catch (error) {
         OrgGetActionTypes.NAME_FAILURE
         console.log(error)
       }
    }
    const getData = {names: names}
    dispatch(get({data: getData})(actionType))
  }
}

export const getTypes = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgContract = state.chainOrgContract.data.contract as IATIOrganisations
    const refs = state.orgReader.data.refs
    const refsLength = refs.length
    //console.log('Num Orgs is: ', numOrgs)
    let actionType = OrgGetActionTypes.TYPE_FAILURE
    let types = []
    for (let i = 0; i < refsLength; i++) {
       try {
         const type = await orgContract.getOrganisationType(refs[i])
         types.push(type)
         //console.log('type: ', type)
         actionType = OrgGetActionTypes.TYPE_SUCCESS
       } catch (error) {
         OrgGetActionTypes.TYPE_FAILURE
         console.log(error)
       }
    }
    const getData = {types: types}
    dispatch(get({data: getData})(actionType))
  }
}
