import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'

import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps } from '../../../types'
import { OrganisationDocsReaderActionTypes, OrganisationDocsReaderProps, OrganisationDocsData } from './types'

export const getDocs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNums())
    await dispatch(getReferences())
    await dispatch(getNumDocs())
    await dispatch(getDocRefs())
    dispatch(getDocsData())
  }
}

const get = (payload: PayloadProps): Function => {
  return (actionType: OrganisationDocsReaderActionTypes): OrganisationDocsReaderProps => {
    const getProps = storeAction(actionType)(payload) as OrganisationDocsReaderProps
    return getProps
  }
}

const getNums = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgDocsContract = state.chainContracts.data.contracts.orgDocsContract
    let actionType = OrganisationDocsReaderActionTypes.NUM_FAILURE
    let nums = { num: 0 }
    try {
      const num = await orgDocsContract.getNums()
      //console.log('Num orgs: ', num)
      nums.num = num.toNumber()
      actionType = OrganisationDocsReaderActionTypes.NUM_SUCCESS
    } catch (error) {
      console.log('getNums error', error)
    }

    dispatch(get({data: nums})(actionType))
  }
}

const getReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgDocsContract = state.chainContracts.data.contracts.orgDocsContract
    const nums = state.orgDocsReader.num
    let reportDocRefs: OrganisationDocsData = {}
    let actionType = OrganisationDocsReaderActionTypes.REF_FAILURE
    for (let i = 0; i < nums; i++) {
       try {
         const ref = await orgDocsContract.getReference(i.toString())
         reportDocRefs[ref] = {
           num: 0,
           data: {}
         }
         actionType = OrganisationDocsReaderActionTypes.REF_SUCCESS
       } catch (error) {
         console.log('getReferences error', error)
       }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reportDocRefs}})(actionType))
  }
}

const getNumDocs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const orgDocsContract = state.chainContracts.data.contracts.orgDocsContract
    let actionType = OrganisationDocsReaderActionTypes.NUMDOC_FAILURE

    const reports = state.orgDocsReader.data as OrganisationDocsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
       try {
         //console.log(' key: ', reportKey)
         const num = await orgDocsContract.getNumDocs(reportKey)
         reports[reportKey].num = num.toNumber()
         //console.log( 'Num reports for ', thisKey, ' ; ', orgs[thisKey].num)
         actionType = OrganisationDocsReaderActionTypes.NUMDOC_SUCCESS
       } catch (error) {
         console.log('getNumDocs error', error)
       }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getDocRefs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgDocsContract = state.chainContracts.data.contracts.orgDocsContract
    let actionType = OrganisationDocsReaderActionTypes.DOCREF_FAILURE

    const reports = state.orgDocsReader.data as OrganisationDocsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const numDocs = reports[reportKey].num
      for (let j = 0; j < numDocs; j++) {
         try {
            const ref = await orgDocsContract.getDocReference(reportKey, j)
            //console.log (' ref: ', ref)
            reports[reportKey].data[ref] = {
              report: {
                reportRef: '',
                orgRef: ''
              },
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
            actionType = OrganisationDocsReaderActionTypes.DOCREF_SUCCESS
          } catch (error) {
            console.log('getDocRefs error', error)
          }
      }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getDocsData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgDocsContract = state.chainContracts.data.contracts.orgDocsContract
    let actionType = OrganisationDocsReaderActionTypes.DOC_FAILURE

    const reports = state.orgDocsReader.data as OrganisationDocsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const docRefKeys = Object.keys(reports[reportKey].data)
      for (let j = 0; j < docRefKeys.length; j++) {
       const docRefKey = docRefKeys[j]
       //console.log(' Data for org ', orgKey, ' ref key ', reportRefKey)
       try {
          const docData = await orgDocsContract.getDocument(reportKey, docRefKey)
          //const reportData = await orgsContract.getCurrency(orgKey, reportRefKey)
          //console.log(' Data for org ', orgKey, ' ref key ', reportRefKey, ' is data ', reportData)
          reports[reportKey].data[docRefKey] = docData
          actionType = OrganisationDocsReaderActionTypes.DOC_SUCCESS
        } catch (error) {
          console.log('getDocData error', error)
        }
      }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}
