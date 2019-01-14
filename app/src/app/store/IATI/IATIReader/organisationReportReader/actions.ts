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
         const num = await orgReportsContract.getNumReports(thisKey)
         orgs[thisKey].num = num.toNumber()
         console.log( 'Num reports for ', thisKey, ' ; ', orgs[thisKey].num)
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
    let actionType = OrgReportReaderActionTypes.REPORTREF_FAILURE

    const orgs = state.orgReportsReader.data as OrgReportData
    const orgKeys = Object.keys(orgs)

    for (let i = 0; i < orgKeys.length; i++) {
      const orgKey = orgKeys[i]
      const numReports = orgs[orgKey].num
      for (let j = 0; j < numReports; j++) {
         try {
            const ref  = await orgReportsContract.getReportReference(orgKey, j)
            console.log ('Report ref: ', ref)
            orgs[orgKey].data[ref] = {
              version: '',
              orgRef: '',
              reportRef: ref,
              reportingOrg: {
                orgRef: '',
                orgType: 0,
                isSecondary: false
              },
              lang: '',
              currency: '',
              lastUpdatedTime: ''
            }
            actionType = OrgReportReaderActionTypes.REPORTREF_SUCCESS
          } catch (error) {
            console.log('getOrgReportRefs error', error)
          }
      }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: orgs}})(actionType))
  }
}

const getOrgReportData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportsContract = state.chainContracts.data.contracts.orgReportsContract as IATIOrganisationReports
    let actionType = OrgReportReaderActionTypes.REPORT_FAILURE

    const orgs = state.orgReportsReader.data as OrgReportData
    const orgKeys = Object.keys(orgs)

    for (let i = 0; i < orgKeys.length; i++) {
      const orgKey = orgKeys[i]
      const orgReportRefKeys = Object.keys(orgs[orgKey].data)
      for (let j = 0; j < orgReportRefKeys.length; j++) {
       const reportRefKey = orgReportRefKeys[j]
        console.log('Report Data for org ', orgKey, ' ref key ', reportRefKey)
       try {
          const reportData = await orgReportsContract.getReport(orgKey, reportRefKey)
          console.log('Report Data for org ', orgKey, ' ref key ', reportRefKey, ' is data ', reportData)
          orgs[orgKey].data[reportRefKey] = reportData
          actionType = OrgReportReaderActionTypes.REPORT_SUCCESS
        } catch (error) {
          console.log('getOrgReportData error', error)
        }
      }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: orgs}})(actionType))
  }
}
