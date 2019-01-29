import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'

import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps } from '../../../types'
import { OrgReportBudgetsReaderActionTypes, OrgReportBudgetsReaderProps, OrgReportBudgetsData } from './types'

export const getReportBudgets = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNumReports())
    await dispatch(getReportReferences())
    await dispatch(getNumReportBudgets())
    await dispatch(getReportBudgetRefs())
    dispatch(getReportBudgetData())
  }
}

const get = (payload: PayloadProps): Function => {
  return (actionType: OrgReportBudgetsReaderActionTypes): OrgReportBudgetsReaderProps => {
    const getProps = storeAction(actionType)(payload) as OrgReportBudgetsReaderProps
    return getProps
  }
}

const getNumReports = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportBudgetsContract = state.chainContracts.data.contracts.orgReportBudgetsContract
    let actionType = OrgReportBudgetsReaderActionTypes.NUM_FAILURE
    let numReports = { num: 0 }
    try {
      const num = await orgReportBudgetsContract.getNumReports()
      //console.log('Num orgs: ', num)
      numReports.num = num.toNumber()
      actionType = OrgReportBudgetsReaderActionTypes.NUM_SUCCESS
    } catch (error) {
      console.log('getNumReports error', error)
    }

    dispatch(get({data: numReports})(actionType))
  }
}

const getReportReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportBudgetsContract = state.chainContracts.data.contracts.orgReportBudgetsContract
    const numReports = state.orgReportBudgetsReader.num
    let reportBudgetRefs: OrgReportBudgetsData = {}
    let actionType = OrgReportBudgetsReaderActionTypes.REF_FAILURE
    for (let i = 0; i < numReports; i++) {
       try {
         const ref = await orgReportBudgetsContract.getReportReference(i.toString())
         reportBudgetRefs[ref] = {
           num: 0,
           data: {}
         }
         actionType = OrgReportBudgetsReaderActionTypes.REF_SUCCESS
       } catch (error) {
         console.log('getReportReferences error', error)
       }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reportBudgetRefs}})(actionType))
  }
}

const getNumReportBudgets = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const orgReportBudgetsContract = state.chainContracts.data.contracts.orgReportBudgetsContract
    let actionType = OrgReportBudgetsReaderActionTypes.NUMBUDGET_FAILURE

    const reports = state.orgReportBudgetsReader.data as OrgReportBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
       try {
         //console.log('Report key: ', reportKey)
         const num = await orgReportBudgetsContract.getNumTotalBudgets(reportKey)
         reports[reportKey].num = num.toNumber()
         //console.log( 'Num reports for ', thisKey, ' ; ', orgs[thisKey].num)
         actionType = OrgReportBudgetsReaderActionTypes.NUMBUDGET_SUCCESS
       } catch (error) {
         console.log('getNumReportBudgets error', error)
       }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getReportBudgetRefs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportBudgetsContract = state.chainContracts.data.contracts.orgReportBudgetsContract
    let actionType = OrgReportBudgetsReaderActionTypes.BUDGETREF_FAILURE

    const reports = state.orgReportBudgetsReader.data as OrgReportBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const numBudgets = reports[reportKey].num
      for (let j = 0; j < numBudgets; j++) {
         try {
            const ref = await orgReportBudgetsContract.getTotalBudgetReference(reportKey, j)
            //console.log ('Report ref: ', ref)
            reports[reportKey].data[ref] = {
              report: {
                reportRef: '',
                orgRef: ''
              },
              budgetRef: ref,
              budgetLine: '',
              finance: {
                value: 0,
                status: 0,
                start: '',
                end: ''
              }
            }
            actionType = OrgReportBudgetsReaderActionTypes.BUDGETREF_SUCCESS
          } catch (error) {
            console.log('getReportBudgetRefs error', error)
          }
      }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getReportBudgetData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportBudgetsContract = state.chainContracts.data.contracts.orgReportBudgetsContract
    let actionType = OrgReportBudgetsReaderActionTypes.BUDGET_FAILURE

    const reports = state.orgReportBudgetsReader.data as OrgReportBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const budgetRefKeys = Object.keys(reports[reportKey].data)
      for (let j = 0; j < budgetRefKeys.length; j++) {
       const budgetRefKey = budgetRefKeys[j]
       //console.log('Report Data for org ', orgKey, ' ref key ', reportRefKey)
       try {
          const budgetData = await orgReportBudgetsContract.getTotalBudget(reportKey, budgetRefKey)
          //const reportData = await orgReportsContract.getCurrency(orgKey, reportRefKey)
          //console.log('Report Data for org ', orgKey, ' ref key ', reportRefKey, ' is data ', reportData)
          reports[reportKey].data[budgetRefKey] = budgetData
          actionType = OrgReportBudgetsReaderActionTypes.BUDGET_SUCCESS
        } catch (error) {
          console.log('getReportBudgetData error', error)
        }
      }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}
