import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'

import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps } from '../../../types'
import { OrganisationsReaderActionTypes, OrganisationsReaderProps, OrgData } from './types'

export const getOrgs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNumOrgs())
    await dispatch(getOrgReferences())
    await dispatch(getNumOrgs())
    await dispatch(getOrgRefs())
    dispatch(getOrgData())
  }
}

const get = (payload: PayloadProps): Function => {
  return (actionType: OrganisationsReaderActionTypes): OrganisationsReaderProps => {
    const getProps = storeAction(actionType)(payload) as OrganisationsReaderProps
    return getProps
  }
}

const getNumOrgs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgsContract = state.chainContracts.data.contracts.orgsContract
    let actionType = OrganisationsReaderActionTypes.NUM_FAILURE
    let numOrgs = { num: 0 }
    try {
      const num = await orgsContract.getNumOrgs()
      //console.log('Num orgs: ', num)
      numOrgs.num = num.toNumber()
      actionType = OrganisationsReaderActionTypes.NUM_SUCCESS
    } catch (error) {
      console.log('getNumOrganisations error', error)
    }

    dispatch(get({data: numOrgs})(actionType))
  }
}

const getOrgReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgsContract = state.chainContracts.data.contracts.orgsContract
    const numOrgs = state.orgsReader.num
    let orgRefs: OrgData = {}
    let actionType = OrganisationsReaderActionTypes.REF_FAILURE
    for (let i = 0; i < numOrgs; i++) {
       try {
         const ref = await orgsContract.getOrganisationReference(i.toString())
         orgRefs[ref] = {
           num: 0,
           data: {}
         }
         actionType = OrganisationsReaderActionTypes.REF_SUCCESS
       } catch (error) {
         console.log('getOrgReferences error', error)
       }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: orgRefs}})(actionType))
  }
}

const getNumOrgs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const orgsContract = state.chainContracts.data.contracts.orgsContract
    let actionType = OrganisationsReaderActionTypes.NUMREP_FAILURE

    const orgs = state.orgsReader.data as OrgData
    const orgKeys = Object.keys(orgs)

    for (let i = 0; i < orgKeys.length; i++) {
      const thisKey = orgKeys[i]
       try {
         const num = await orgsContract.getNums(thisKey)
         orgs[thisKey].num = num.toNumber()
         //console.log( 'Num reports for ', thisKey, ' ; ', orgs[thisKey].num)
         actionType = OrganisationsReaderActionTypes.NUMREP_SUCCESS
       } catch (error) {
         console.log('getNumOrgs error', error)
       }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: orgs}})(actionType))
  }
}

const getOrgRefs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgsContract = state.chainContracts.data.contracts.orgsContract
    let actionType = OrganisationsReaderActionTypes.REPORTREF_FAILURE

    const orgs = state.orgsReader.data as OrgData
    const orgKeys = Object.keys(orgs)

    for (let i = 0; i < orgKeys.length; i++) {
      const orgKey = orgKeys[i]
      const nums = orgs[orgKey].num
      for (let j = 0; j < nums; j++) {
         try {
            const ref = await orgsContract.getReference(orgKey, j)
            //console.log (' ref: ', ref)
            orgs[orgKey].data[ref] = {
              version: '',
              report: {
                reportRef: ref,
                orgRef: ''
              },
              reportingOrg: {
                orgRef: '',
                orgType: 0,
                isSecondary: false
              },
              lang: '',
              currency: '',
              lastUpdatedTime: ''
            }
            actionType = OrganisationsReaderActionTypes.REPORTREF_SUCCESS
          } catch (error) {
            console.log('getOrgRefs error', error)
          }
      }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: orgs}})(actionType))
  }
}

const getOrgData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgsContract = state.chainContracts.data.contracts.orgsContract
    let actionType = OrganisationsReaderActionTypes.REPORT_FAILURE

    const orgs = state.orgsReader.data as OrgData
    const orgKeys = Object.keys(orgs)

    for (let i = 0; i < orgKeys.length; i++) {
      const orgKey = orgKeys[i]
      const orgRefKeys = Object.keys(orgs[orgKey].data)
      for (let j = 0; j < orgRefKeys.length; j++) {
       const reportRefKey = orgRefKeys[j]
       //console.log(' Data for org ', orgKey, ' ref key ', reportRefKey)
       try {
          const reportData = await orgsContract.get(orgKey, reportRefKey)
          //const reportData = await orgsContract.getCurrency(orgKey, reportRefKey)
          //console.log(' Data for org ', orgKey, ' ref key ', reportRefKey, ' is data ', reportData)
          orgs[orgKey].data[reportRefKey] = reportData
          actionType = OrganisationsReaderActionTypes.REPORT_SUCCESS
        } catch (error) {
          console.log('getOrgData error', error)
        }
      }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: orgs}})(actionType))
  }
}
