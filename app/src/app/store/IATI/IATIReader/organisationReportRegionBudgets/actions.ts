import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'

import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps } from '../../../types'
import { OrgReportRegionBudgetsReaderActionTypes, OrgReportRegionBudgetsReaderProps, OrgReportRegionBudgetsData } from './types'

export const getReportRegionBudgets = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNumReports())
    await dispatch(getReportReferences())
    await dispatch(getNumReportRegionBudgets())
    await dispatch(getReportRegionBudgetRefs())
    dispatch(getReportRegionBudgetData())
  }
}

const get = (payload: PayloadProps): Function => {
  return (actionType: OrgReportRegionBudgetsReaderActionTypes): OrgReportRegionBudgetsReaderProps => {
    const getProps = storeAction(actionType)(payload) as OrgReportRegionBudgetsReaderProps
    return getProps
  }
}

const getNumReports = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportRegionBudgetsContract = state.chainContracts.data.contracts.orgReportRegionBudgetsContract
    let actionType = OrgReportRegionBudgetsReaderActionTypes.NUM_FAILURE
    let numReports = { num: 0 }
    try {
      const num = await orgReportRegionBudgetsContract.getNumReports()
      //console.log('Num orgs: ', num)
      numReports.num = num.toNumber()
      actionType = OrgReportRegionBudgetsReaderActionTypes.NUM_SUCCESS
    } catch (error) {
      console.log('getNumReports error', error)
    }

    dispatch(get({data: numReports})(actionType))
  }
}

const getReportReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportRegionBudgetsContract = state.chainContracts.data.contracts.orgReportRegionBudgetsContract
    const numReports = state.orgReportRegionBudgetsReader.num
    let reportRegionBudgetRefs: OrgReportRegionBudgetsData = {}
    let actionType = OrgReportRegionBudgetsReaderActionTypes.REF_FAILURE
    for (let i = 0; i < numReports; i++) {
       try {
         const ref = await orgReportRegionBudgetsContract.getReportReference(i.toString())
         reportRegionBudgetRefs[ref] = {
           num: 0,
           data: {}
         }
         actionType = OrgReportRegionBudgetsReaderActionTypes.REF_SUCCESS
       } catch (error) {
         console.log('getReportReferences error', error)
       }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reportRegionBudgetRefs}})(actionType))
  }
}

const getNumReportRegionBudgets = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const orgReportRegionBudgetsContract = state.chainContracts.data.contracts.orgReportRegionBudgetsContract
    let actionType = OrgReportRegionBudgetsReaderActionTypes.NUMBUDGET_FAILURE

    const reports = state.orgReportRegionBudgetsReader.data as OrgReportRegionBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
       try {
         //console.log('Report key: ', reportKey)
         const num = await orgReportRegionBudgetsContract.getNumRegionBudgets(reportKey)
         reports[reportKey].num = num.toNumber()
         //console.log( 'Num reports for ', thisKey, ' ; ', orgs[thisKey].num)
         actionType = OrgReportRegionBudgetsReaderActionTypes.NUMBUDGET_SUCCESS
       } catch (error) {
         console.log('getNumReportRegionBudgets error', error)
       }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getReportRegionBudgetRefs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportRegionBudgetsContract = state.chainContracts.data.contracts.orgReportRegionBudgetsContract
    let actionType = OrgReportRegionBudgetsReaderActionTypes.BUDGETREF_FAILURE

    const reports = state.orgReportRegionBudgetsReader.data as OrgReportRegionBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const numRegionBudgets = reports[reportKey].num
      for (let j = 0; j < numRegionBudgets; j++) {
         try {
            const ref = await orgReportRegionBudgetsContract.getRegionBudgetReference(reportKey, j)
            //console.log ('Report ref: ', ref)
            reports[reportKey].data[ref] = {
              reportRef: '',
              budgetRef: ref,
              regionRef: '',
              budgetLine: '',
              finance: {
                value: 0,
                status: 0,
                start: '',
                end: ''
              }
            }
            actionType = OrgReportRegionBudgetsReaderActionTypes.BUDGETREF_SUCCESS
          } catch (error) {
            console.log('getReportRegionBudgetRefs error', error)
          }
      }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getReportRegionBudgetData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportRegionBudgetsContract = state.chainContracts.data.contracts.orgReportRegionBudgetsContract
    let actionType = OrgReportRegionBudgetsReaderActionTypes.BUDGET_FAILURE

    const reports = state.orgReportRegionBudgetsReader.data as OrgReportRegionBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const budgetRefKeys = Object.keys(reports[reportKey].data)
      for (let j = 0; j < budgetRefKeys.length; j++) {
       const budgetRefKey = budgetRefKeys[j]
       //console.log('Report Data for org ', orgKey, ' ref key ', reportRefKey)
       try {
          const budgetData = await orgReportRegionBudgetsContract.getRegionsBudget(reportKey, budgetRefKey)
          //const reportData = await orgReportsContract.getCurrency(orgKey, reportRefKey)
          //console.log('Report Data for org ', orgKey, ' ref key ', reportRefKey, ' is data ', reportData)
          reports[reportKey].data[budgetRefKey] = budgetData
          actionType = OrgReportRegionBudgetsReaderActionTypes.BUDGET_SUCCESS
        } catch (error) {
          console.log('getReportRegionBudgetData error', error)
        }
      }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}
