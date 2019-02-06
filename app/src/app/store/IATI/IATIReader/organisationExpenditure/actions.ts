import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'

import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps } from '../../../types'
import { OrgExpenditureReaderActionTypes, OrgExpenditureReaderProps, OrgExpenditureData } from './types'

export const getExpenditure = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNums())
    await dispatch(getReferences())
    await dispatch(getNumExpenditure())
    await dispatch(getExpenditureRefs())
    dispatch(getExpenditureData())
  }
}

const get = (payload: PayloadProps): Function => {
  return (actionType: OrgExpenditureReaderActionTypes): OrgExpenditureReaderProps => {
    const getProps = storeAction(actionType)(payload) as OrgExpenditureReaderProps
    return getProps
  }
}

const getNums = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgExpenditureContract = state.chainContracts.data.contracts.orgExpenditureContract
    let actionType = OrgExpenditureReaderActionTypes.NUM_FAILURE
    let nums = { num: 0 }
    try {
      const num = await orgExpenditureContract.getNums()
      //console.log('Num orgs: ', num)
      nums.num = num.toNumber()
      actionType = OrgExpenditureReaderActionTypes.NUM_SUCCESS
    } catch (error) {
      console.log('getNums error', error)
    }

    dispatch(get({data: nums})(actionType))
  }
}

const getReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgExpenditureContract = state.chainContracts.data.contracts.orgExpenditureContract
    const nums = state.orgExpenditureReader.num
    let reportExpenditureRefs: OrgExpenditureData = {}
    let actionType = OrgExpenditureReaderActionTypes.REF_FAILURE
    for (let i = 0; i < nums; i++) {
       try {
         const ref = await orgExpenditureContract.getReference(i.toString())
         reportExpenditureRefs[ref] = {
           num: 0,
           data: {}
         }
         actionType = OrgExpenditureReaderActionTypes.REF_SUCCESS
       } catch (error) {
         console.log('getReferences error', error)
       }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reportExpenditureRefs}})(actionType))
  }
}

const getNumExpenditure = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const orgExpenditureContract = state.chainContracts.data.contracts.orgExpenditureContract
    let actionType = OrgExpenditureReaderActionTypes.NUMEXPENDITURE_FAILURE

    const reports = state.orgExpenditureReader.data as OrgExpenditureData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
       try {
         //console.log(' key: ', reportKey)
         const num = await orgExpenditureContract.getNumExpenditures(reportKey)
         reports[reportKey].num = num.toNumber()
         //console.log( 'Num reports for ', thisKey, ' ; ', orgs[thisKey].num)
         actionType = OrgExpenditureReaderActionTypes.NUMEXPENDITURE_SUCCESS
       } catch (error) {
         console.log('getNumExpenditure error', error)
       }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getExpenditureRefs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgExpenditureContract = state.chainContracts.data.contracts.orgExpenditureContract
    let actionType = OrgExpenditureReaderActionTypes.EXPENDITUREREF_FAILURE

    const reports = state.orgExpenditureReader.data as OrgExpenditureData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const numExpenditure = reports[reportKey].num
      for (let j = 0; j < numExpenditure; j++) {
         try {
            const ref = await orgExpenditureContract.getExpenditureReference(reportKey, j)
            //console.log (' ref: ', ref)
            reports[reportKey].data[ref] = {
              report: {
                reportRef: '',
                orgRef: ''
              },
              expenditureRef: ref,
              expenditureLine: '',
              finance: {
                value: 0,
                status: 0,
                start: '',
                end: ''
              }
            }
            actionType = OrgExpenditureReaderActionTypes.EXPENDITUREREF_SUCCESS
          } catch (error) {
            console.log('getExpenditureRefs error', error)
          }
      }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getExpenditureData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgExpenditureContract = state.chainContracts.data.contracts.orgExpenditureContract
    let actionType = OrgExpenditureReaderActionTypes.EXPENDITURE_FAILURE

    const reports = state.orgExpenditureReader.data as OrgExpenditureData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const expenditureRefKeys = Object.keys(reports[reportKey].data)
      for (let j = 0; j < expenditureRefKeys.length; j++) {
       const expenditureRefKey = expenditureRefKeys[j]
       //console.log(' Data for org ', orgKey, ' ref key ', reportRefKey)
       try {
          const expenditureData = await orgExpenditureContract.getExpenditure(reportKey, expenditureRefKey)
          //const reportData = await orgsContract.getCurrency(orgKey, reportRefKey)
          //console.log(' Data for org ', orgKey, ' ref key ', reportRefKey, ' is data ', reportData)
          reports[reportKey].data[expenditureRefKey] = expenditureData
          actionType = OrgExpenditureReaderActionTypes.EXPENDITURE_SUCCESS
        } catch (error) {
          console.log('getExpenditureData error', error)
        }
      }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}
