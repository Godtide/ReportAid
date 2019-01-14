import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'
import { IATIOrganisationReports } from '../../../../../blockchain/typechain/IATIOrganisationReports'

import { storeAction } from '../../../actions'
import { ActionProps, PayloadProps } from '../../../types'

import { OrgReportReaderActionTypes, OrgReportReaderProps, OrgReportData, ReportData, Report } from './types'

export const getOrgReports = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNumOrgs())
    await dispatch(getOrgReferences())
    await dispatch(getNumOrgReports())
    await dispatch(getOrgReportRefs())
    dispatch(getOrgReportData())
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

const getOrgReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportsContract = state.chainContracts.data.contracts.orgReportsContract as IATIOrganisationReports
    const numOrgs = state.orgReportsReader.num
    let orgReportRefs: OrgReportData = {}
    let actionType = OrgReportReaderActionTypes.REF_FAILURE
    for (let i = 0; i < numOrgs; i++) {
       try {
         const ref = await orgReportsContract.getOrganisationReference(i.toString())
         orgReportRefs[ref] = {
           num: 0,
           data: {}
         }
         actionType = OrgReportReaderActionTypes.REF_SUCCESS
       } catch (error) {
         console.log('getOrgReferences error', error)
       }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: orgReportRefs}})(actionType))
  }
}

const getNumOrgReports = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const orgReportsContract = state.chainContracts.data.contracts.orgReportsContract as IATIOrganisationReports
    let actionType = OrgReportReaderActionTypes.NUMREP_FAILURE

    const orgs = state.orgReportsReader.data as OrgReportData
    const orgKeys = Object.keys(orgs)

    for (let i = 0; i < orgKeys.length; i++) {
      const thisKey = orgKeys[i]
       try {
         orgs[thisKey].num = await orgReportsContract.getNumReports(thisKey)
         actionType = OrgReportReaderActionTypes.NUMREP_SUCCESS
       } catch (error) {
         console.log('getNumOrgReports error', error)
       }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: orgs}})(actionType))

  }
}

const getOrgReportRefs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportsContract = state.chainContracts.data.contracts.orgReportsContract as IATIOrganisationReports
    let actionType = OrgReportReaderActionTypes.REPORT_FAILURE

    const orgs = state.orgReportsReader.data as OrgReportData
    const orgKeys = Object.keys(orgs)

    for (let i = 0; i < orgKeys.length; i++) {
      const thisKey = orgKeys[i]
      const numReports = orgs[thisKey].num
      for (let j = 0; j < numReports; i++) {
       try {
          const ref  = await orgReportsContract.getReportReference(thisKey, j)
          orgs[thisKey].data[ref].reportRef = ref as string
          actionType = OrgReportReaderActionTypes.REPORT_SUCCESS
        } catch (error) {
          console.log('getNumOrgReports error', error)
        }
      }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: orgs}})(actionType))

  }
}

const getOrgReportData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
  }
}
