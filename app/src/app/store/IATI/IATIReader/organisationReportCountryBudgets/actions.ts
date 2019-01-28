import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'

import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps } from '../../../types'
import { OrgReportCountryBudgetsReaderActionTypes, OrgReportCountryBudgetsReaderProps, OrgReportCountryBudgetsData } from './types'

export const getReportCountryBudgets = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNumReports())
    await dispatch(getReportReferences())
    await dispatch(getNumReportCountryBudgets())
    await dispatch(getReportCountryBudgetRefs())
    dispatch(getReportCountryBudgetData())
  }
}

const get = (payload: PayloadProps): Function => {
  return (actionType: OrgReportCountryBudgetsReaderActionTypes): OrgReportCountryBudgetsReaderProps => {
    const getProps = storeAction(actionType)(payload) as OrgReportCountryBudgetsReaderProps
    return getProps
  }
}

const getNumReports = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportCountryBudgetsContract = state.chainContracts.data.contracts.orgReportCountryBudgetsContract
    let actionType = OrgReportCountryBudgetsReaderActionTypes.NUM_FAILURE
    let numReports = { num: 0 }
    try {
      const num = await orgReportCountryBudgetsContract.getNumReports()
      //console.log('Num orgs: ', num)
      numReports.num = num.toNumber()
      actionType = OrgReportCountryBudgetsReaderActionTypes.NUM_SUCCESS
    } catch (error) {
      console.log('getNumReports error', error)
    }

    dispatch(get({data: numReports})(actionType))
  }
}

const getReportReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportCountryBudgetsContract = state.chainContracts.data.contracts.orgReportCountryBudgetsContract
    const numReports = state.orgReportCountryBudgetsReader.num
    let reportCountryBudgetRefs: OrgReportCountryBudgetsData = {}
    let actionType = OrgReportCountryBudgetsReaderActionTypes.REF_FAILURE
    for (let i = 0; i < numReports; i++) {
       try {
         const ref = await orgReportCountryBudgetsContract.getReportReference(i.toString())
         reportCountryBudgetRefs[ref] = {
           num: 0,
           data: {}
         }
         actionType = OrgReportCountryBudgetsReaderActionTypes.REF_SUCCESS
       } catch (error) {
         console.log('getReportReferences error', error)
       }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reportCountryBudgetRefs}})(actionType))
  }
}

const getNumReportCountryBudgets = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const orgReportCountryBudgetsContract = state.chainContracts.data.contracts.orgReportCountryBudgetsContract
    let actionType = OrgReportCountryBudgetsReaderActionTypes.NUMBUDGET_FAILURE

    const reports = state.orgReportCountryBudgetsReader.data as OrgReportCountryBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
       try {
         //console.log('Report key: ', reportKey)
         const num = await orgReportCountryBudgetsContract.getNumCountryBudgets(reportKey)
         reports[reportKey].num = num.toNumber()
         //console.log( 'Num reports for ', thisKey, ' ; ', orgs[thisKey].num)
         actionType = OrgReportCountryBudgetsReaderActionTypes.NUMBUDGET_SUCCESS
       } catch (error) {
         console.log('getNumReportCountryBudgets error', error)
       }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getReportCountryBudgetRefs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportCountryBudgetsContract = state.chainContracts.data.contracts.orgReportCountryBudgetsContract
    let actionType = OrgReportCountryBudgetsReaderActionTypes.BUDGETREF_FAILURE

    const reports = state.orgReportCountryBudgetsReader.data as OrgReportCountryBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const numCountryBudgets = reports[reportKey].num
      for (let j = 0; j < numCountryBudgets; j++) {
         try {
            const ref = await orgReportCountryBudgetsContract.getCountryBudgetReference(reportKey, j)
            //console.log ('Report ref: ', ref)
            reports[reportKey].data[ref] = {
              reportRef: '',
              budgetRef: ref,
              countryRef: '',
              budgetLine: '',
              finance: {
                value: 0,
                status: 0,
                start: '',
                end: ''
              }
            }
            actionType = OrgReportCountryBudgetsReaderActionTypes.BUDGETREF_SUCCESS
          } catch (error) {
            console.log('getReportCountryBudgetRefs error', error)
          }
      }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getReportCountryBudgetData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportCountryBudgetsContract = state.chainContracts.data.contracts.orgReportCountryBudgetsContract
    let actionType = OrgReportCountryBudgetsReaderActionTypes.BUDGET_FAILURE

    const reports = state.orgReportCountryBudgetsReader.data as OrgReportCountryBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const budgetRefKeys = Object.keys(reports[reportKey].data)
      for (let j = 0; j < budgetRefKeys.length; j++) {
       const budgetRefKey = budgetRefKeys[j]
       //console.log('Report Data for org ', orgKey, ' ref key ', reportRefKey)
       try {
          const budgetData = await orgReportCountryBudgetsContract.getCountryBudget(reportKey, budgetRefKey)
          //const reportData = await orgReportsContract.getCurrency(orgKey, reportRefKey)
          //console.log('Report Data for org ', orgKey, ' ref key ', reportRefKey, ' is data ', reportData)
          reports[reportKey].data[budgetRefKey] = budgetData
          actionType = OrgReportCountryBudgetsReaderActionTypes.BUDGET_SUCCESS
        } catch (error) {
          console.log('getReportCountryBudgetData error', error)
        }
      }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}
