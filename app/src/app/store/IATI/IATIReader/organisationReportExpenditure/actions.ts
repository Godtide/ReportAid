import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'

import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps } from '../../../types'
import { OrgReportExpenditureReaderActionTypes, OrgReportExpenditureReaderProps, OrgReportExpenditureData } from './types'

export const getReportExpenditure = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNumReports())
    await dispatch(getReportReferences())
    await dispatch(getNumReportExpenditure())
    await dispatch(getReportExpenditureRefs())
    dispatch(getReportExpenditureData())
  }
}

const get = (payload: PayloadProps): Function => {
  return (actionType: OrgReportExpenditureReaderActionTypes): OrgReportExpenditureReaderProps => {
    const getProps = storeAction(actionType)(payload) as OrgReportExpenditureReaderProps
    return getProps
  }
}

const getNumReports = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportExpenditureContract = state.chainContracts.data.contracts.orgReportExpenditureContract
    let actionType = OrgReportExpenditureReaderActionTypes.NUM_FAILURE
    let numReports = { num: 0 }
    try {
      const num = await orgReportExpenditureContract.getNumReports()
      //console.log('Num orgs: ', num)
      numReports.num = num.toNumber()
      actionType = OrgReportExpenditureReaderActionTypes.NUM_SUCCESS
    } catch (error) {
      console.log('getNumReports error', error)
    }

    dispatch(get({data: numReports})(actionType))
  }
}

const getReportReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportExpenditureContract = state.chainContracts.data.contracts.orgReportExpenditureContract
    const numReports = state.orgReportExpenditureReader.num
    let reportExpenditureRefs: OrgReportExpenditureData = {}
    let actionType = OrgReportExpenditureReaderActionTypes.REF_FAILURE
    for (let i = 0; i < numReports; i++) {
       try {
         const ref = await orgReportExpenditureContract.getReportReference(i.toString())
         reportExpenditureRefs[ref] = {
           num: 0,
           data: {}
         }
         actionType = OrgReportExpenditureReaderActionTypes.REF_SUCCESS
       } catch (error) {
         console.log('getReportReferences error', error)
       }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reportExpenditureRefs}})(actionType))
  }
}

const getNumReportExpenditure = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const orgReportExpenditureContract = state.chainContracts.data.contracts.orgReportExpenditureContract
    let actionType = OrgReportExpenditureReaderActionTypes.NUMEXPENDITURE_FAILURE

    const reports = state.orgReportExpenditureReader.data as OrgReportExpenditureData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
       try {
         //console.log('Report key: ', reportKey)
         const num = await orgReportExpenditureContract.getNumExpenditures(reportKey)
         reports[reportKey].num = num.toNumber()
         //console.log( 'Num reports for ', thisKey, ' ; ', orgs[thisKey].num)
         actionType = OrgReportExpenditureReaderActionTypes.NUMEXPENDITURE_SUCCESS
       } catch (error) {
         console.log('getNumReportExpenditure error', error)
       }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getReportExpenditureRefs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportExpenditureContract = state.chainContracts.data.contracts.orgReportExpenditureContract
    let actionType = OrgReportExpenditureReaderActionTypes.EXPENDITUREREF_FAILURE

    const reports = state.orgReportExpenditureReader.data as OrgReportExpenditureData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const numExpenditure = reports[reportKey].num
      for (let j = 0; j < numExpenditure; j++) {
         try {
            const ref = await orgReportExpenditureContract.getExpenditureReference(reportKey, j)
            //console.log ('Report ref: ', ref)
            reports[reportKey].data[ref] = {
              reportRef: '',
              expenditureRef: ref,
              expenditureLine: '',
              finance: {
                value: 0,
                status: '',
                start: '',
                end: ''
              }
            }
            actionType = OrgReportExpenditureReaderActionTypes.EXPENDITUREREF_SUCCESS
          } catch (error) {
            console.log('getReportExpenditureRefs error', error)
          }
      }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getReportExpenditureData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportExpenditureContract = state.chainContracts.data.contracts.orgReportExpenditureContract
    let actionType = OrgReportExpenditureReaderActionTypes.EXPENDITURE_FAILURE

    const reports = state.orgReportExpenditureReader.data as OrgReportExpenditureData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const expenditureRefKeys = Object.keys(reports[reportKey].data)
      for (let j = 0; j < expenditureRefKeys.length; j++) {
       const expenditureRefKey = expenditureRefKeys[j]
       //console.log('Report Data for org ', orgKey, ' ref key ', reportRefKey)
       try {
          const expenditureData = await orgReportExpenditureContract.getExpenditure(reportKey, expenditureRefKey)
          //const reportData = await orgReportsContract.getCurrency(orgKey, reportRefKey)
          //console.log('Report Data for org ', orgKey, ' ref key ', reportRefKey, ' is data ', reportData)
          reports[reportKey].data[expenditureRefKey] = expenditureData
          actionType = OrgReportExpenditureReaderActionTypes.EXPENDITURE_SUCCESS
        } catch (error) {
          console.log('getReportExpenditureData error', error)
        }
      }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}
