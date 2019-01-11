import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'
import { IATIOrganisationReports } from '../../../../../blockchain/typechain/IATIOrganisationReports'

import { storeAction } from '../../../actions'
import { ActionProps, PayloadProps } from '../../../types'

import { OrgReportReaderActionTypes, OrgReportReaderProps, OrgReportData } from './types'

export const getOrgReports = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNumOrgs())
    await dispatch(getReferences())
    dispatch(getOrgReportData())
    /* await dispatch(getNames())
    dispatch(getIDs())*/
  }
}

const get = (payload: PayloadProps): Function => {
  return (actionType: OrgReportReaderActionTypes): OrgReportReaderProps => {
    const getProps = storeAction(actionType)(payload) as OrgReportReaderProps
    return getProps
  }
}

const getNumOrgs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportsContract = state.chainContracts.data.contracts.orgReportsContract as IATIOrganisationReports
    let actionType = OrgReportReaderActionTypes.NUM_FAILURE
    let numOrgs = { num: 0 }
    try {
      const num = await orgReportsContract.getNumOrgs()
      //console.log('Num orgs: ', num)
      numOrgs.num = num.toNumber()
      actionType = OrgReportReaderActionTypes.NUM_SUCCESS
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
    let orgReportRefs: OrgReportData = {}
    let actionType = OrgReportReaderActionTypes.REF_FAILURE
    for (let i = 0; i < numOrgs; i++) {
       try {
         const ref = await orgReportsContract.getOrganisationReference(i.toString())
         //console.log('Ref', ref)
         orgReportRefs[ref] = {
           '': {
             version: '',
             orgRef: ref,
             reportRef: '',
             reportingOrg: {
               orgRef: '',
               orgType: 0,
               isSecondary: false
             },
             lang: '',
             currency: '',
             lastUpdatedTime: ''
           }
         }
         actionType = OrgReportReaderActionTypes.REF_SUCCESS
       } catch (error) {
         console.log('getReferences error', error)
       }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: orgReportRefs}})(actionType))
  }
}

const getOrgReportData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    /*const state = getState()
    const orgReportsContract = state.chainContracts.data.contracts.orgReportsContract as IATIOrganisationReports
    let actionType = OrgGetActionTypes.ORG_FAILURE
    const orgs = state.orgReader.data
    const orgKeys = Object.keys(orgs)
    //console.log('Orgkeys: ', orgKeys)
    for (let i = 0; i < orgKeys.length; i++) {
      const thisKey = orgKeys[i]
       try {
         const data = await orgReportsContract.getOrganisation(thisKey)
         //console.log ('Org stuff: ', data)
         orgs[thisKey].name = data[1]
         orgs[thisKey].identifier = data[2]
         actionType = OrgGetActionTypes.ORG_SUCCESS
       } catch (error) {
         console.log('getOrgs error', error)
       }
    }
    //console.log('New Orgs; ', orgs)
    dispatch(get({data: {data: orgs}})(actionType))*/
  }
}
