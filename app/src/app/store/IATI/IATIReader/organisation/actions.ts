import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'
import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps } from '../../../types'
import { OrgReaderActionTypes, OrgReaderProps, OrgData } from './types'

export const getOrgs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNumOrganisations())
    await dispatch(getReferences())
    dispatch(getOrgData())
    /* await dispatch(getNames())
    dispatch(getIDs())*/
  }
}

const get = (payload: PayloadProps): Function => {
  return (actionType: OrgReaderActionTypes): OrgReaderProps => {
    const getProps = storeAction(actionType)(payload) as OrgReaderProps
    return getProps
  }
}

const getNumOrganisations = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgContract = state.chainContracts.data.contracts.orgContract
    let actionType = OrgReaderActionTypes.NUM_FAILURE
    let numOrgs = { num: 0 }
    try {
      const num = await orgContract.getNumOrganisations()
      numOrgs.num = num.toNumber()
      actionType = OrgReaderActionTypes.NUM_SUCCESS
    } catch (error) {
      console.log('getNumOrganisations error', error)
    }

    dispatch(get({data: numOrgs})(actionType))
  }
}

const getReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgContract = state.chainContracts.data.contracts.orgContract
    const numOrgs = state.orgReader.num
    let orgRefs: OrgData = {}
    let actionType = OrgReaderActionTypes.REF_FAILURE
    for (let i = 0; i < numOrgs; i++) {
       try {
         const ref = await orgContract.getOrganisationReference(i.toString())
         //console.log('Ref', ref)
         orgRefs[ref] = {
           orgRef: ref,
           name: '',
           identifier: ''
         }
         actionType = OrgReaderActionTypes.REF_SUCCESS
       } catch (error) {
         console.log('getReferences error', error)
       }
    }

    dispatch(get({data: {data: orgRefs}})(actionType))
  }
}

const getOrgData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgContract = state.chainContracts.data.contracts.orgContract
    let actionType = OrgReaderActionTypes.ORG_FAILURE
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
         actionType = OrgReaderActionTypes.ORG_SUCCESS
       } catch (error) {
         console.log('getOrgs error', error)
       }
    }

    dispatch(get({data: {data: orgs}})(actionType))
  }
}

const getNames = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgContract = state.chainContracts.data.contracts.orgContract
    let actionType = OrgReaderActionTypes.NAME_FAILURE
    const orgs = state.orgReader.data
    const orgKeys = Object.keys(orgs)
    //console.log('Orgkeys: ', orgKeys)
    for (let i = 0; i < orgKeys.length; i++) {
      const thisKey = orgKeys[i]
       try {
         orgs[thisKey].name = await orgContract.getOrganisationName(thisKey)
         actionType = OrgReaderActionTypes.NAME_SUCCESS
       } catch (error) {
         console.log('getNames error', error)
       }
    }

    dispatch(get({data: {data: orgs}})(actionType))
  }
}

const getIDs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgContract = state.chainContracts.data.contracts.orgContract
    let actionType = OrgReaderActionTypes.ID_FAILURE
    const orgs = state.orgReader.data
    const orgKeys = Object.keys(orgs)
    //console.log('Orgkeys: ', orgKeys)
    for (let i = 0; i < orgKeys.length; i++) {
      const thisKey = orgKeys[i]
       try {
         orgs[thisKey].identifier = await orgContract.getOrganisationIdentifier(thisKey)
         actionType = OrgReaderActionTypes.ID_SUCCESS
       } catch (error) {
         console.log('getIDs error', error)
       }
    }

    dispatch(get({data: {data: orgs}})(actionType))
  }
}
