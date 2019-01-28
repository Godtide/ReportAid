import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'

import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps } from '../../../types'
import { OrgReportDocsReaderActionTypes, OrgReportDocsReaderProps, OrgReportDocsData } from './types'

export const getReportDocs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNumReports())
    await dispatch(getReportReferences())
    await dispatch(getNumReportDocs())
    await dispatch(getReportDocRefs())
    dispatch(getReportDocsData())
  }
}

const get = (payload: PayloadProps): Function => {
  return (actionType: OrgReportDocsReaderActionTypes): OrgReportDocsReaderProps => {
    const getProps = storeAction(actionType)(payload) as OrgReportDocsReaderProps
    return getProps
  }
}

const getNumReports = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportDocsContract = state.chainContracts.data.contracts.orgReportDocsContract
    let actionType = OrgReportDocsReaderActionTypes.NUM_FAILURE
    let numReports = { num: 0 }
    try {
      const num = await orgReportDocsContract.getNumReports()
      //console.log('Num orgs: ', num)
      numReports.num = num.toNumber()
      actionType = OrgReportDocsReaderActionTypes.NUM_SUCCESS
    } catch (error) {
      console.log('getNumReports error', error)
    }

    dispatch(get({data: numReports})(actionType))
  }
}

const getReportReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportDocsContract = state.chainContracts.data.contracts.orgReportDocsContract
    const numReports = state.orgReportDocsReader.num
    let reportDocRefs: OrgReportDocsData = {}
    let actionType = OrgReportDocsReaderActionTypes.REF_FAILURE
    for (let i = 0; i < numReports; i++) {
       try {
         const ref = await orgReportDocsContract.getReportReference(i.toString())
         reportDocRefs[ref] = {
           num: 0,
           data: {}
         }
         actionType = OrgReportDocsReaderActionTypes.REF_SUCCESS
       } catch (error) {
         console.log('getReportReferences error', error)
       }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reportDocRefs}})(actionType))
  }
}

const getNumReportDocs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const orgReportDocsContract = state.chainContracts.data.contracts.orgReportDocsContract
    let actionType = OrgReportDocsReaderActionTypes.NUMDOC_FAILURE

    const reports = state.orgReportDocsReader.data as OrgReportDocsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
       try {
         //console.log('Report key: ', reportKey)
         const num = await orgReportDocsContract.getNumReportDocs(reportKey)
         reports[reportKey].num = num.toNumber()
         //console.log( 'Num reports for ', thisKey, ' ; ', orgs[thisKey].num)
         actionType = OrgReportDocsReaderActionTypes.NUMDOC_SUCCESS
       } catch (error) {
         console.log('getNumReportDocs error', error)
       }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getReportDocRefs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportDocsContract = state.chainContracts.data.contracts.orgReportDocsContract
    let actionType = OrgReportDocsReaderActionTypes.DOCREF_FAILURE

    const reports = state.orgReportDocsReader.data as OrgReportDocsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const numDocs = reports[reportKey].num
      for (let j = 0; j < numDocs; j++) {
         try {
            const ref = await orgReportDocsContract.getReportDocReference(reportKey, j)
            //console.log ('Report ref: ', ref)
            reports[reportKey].data[ref] = {
              reportRef: '',
              docRef: ref,
              title: '',
              format: '',
              url: '',
              category: '',
              countryRef: '',
              desc: '',
              lang: '',
              date: ''
            }
            actionType = OrgReportDocsReaderActionTypes.DOCREF_SUCCESS
          } catch (error) {
            console.log('getReportDocRefs error', error)
          }
      }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getReportDocsData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportDocsContract = state.chainContracts.data.contracts.orgReportDocsContract
    let actionType = OrgReportDocsReaderActionTypes.DOC_FAILURE

    const reports = state.orgReportDocsReader.data as OrgReportDocsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const docRefKeys = Object.keys(reports[reportKey].data)
      for (let j = 0; j < docRefKeys.length; j++) {
       const docRefKey = docRefKeys[j]
       //console.log('Report Data for org ', orgKey, ' ref key ', reportRefKey)
       try {
          const docData = await orgReportDocsContract.getDocument(reportKey, docRefKey)
          //const reportData = await orgReportsContract.getCurrency(orgKey, reportRefKey)
          //console.log('Report Data for org ', orgKey, ' ref key ', reportRefKey, ' is data ', reportData)
          reports[reportKey].data[docRefKey] = docData
          actionType = OrgReportDocsReaderActionTypes.DOC_SUCCESS
        } catch (error) {
          console.log('getReportDocData error', error)
        }
      }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}
