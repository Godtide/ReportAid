import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'

import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps } from '../../../types'
import { OrgReportRecipientBudgetsReaderActionTypes, OrgReportRecipientBudgetsReaderProps, OrgReportRecipientBudgetsData } from './types'

export const getReportRecipientBudgets = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNumReports())
    await dispatch(getReportReferences())
    await dispatch(getNumReportRecipientBudgets())
    await dispatch(getReportRecipientBudgetRefs())
    dispatch(getReportRecipientBudgetData())
  }
}

const get = (payload: PayloadProps): Function => {
  return (actionType: OrgReportRecipientBudgetsReaderActionTypes): OrgReportRecipientBudgetsReaderProps => {
    const getProps = storeAction(actionType)(payload) as OrgReportRecipientBudgetsReaderProps
    return getProps
  }
}

const getNumReports = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportRecipientBudgetsContract = state.chainContracts.data.contracts.orgReportRecipientBudgetsContract
    let actionType = OrgReportRecipientBudgetsReaderActionTypes.NUM_FAILURE
    let numReports = { num: 0 }
    try {
      const num = await orgReportRecipientBudgetsContract.getNumReports()
      //console.log('Num orgs: ', num)
      numReports.num = num.toNumber()
      actionType = OrgReportRecipientBudgetsReaderActionTypes.NUM_SUCCESS
    } catch (error) {
      console.log('getNumReports error', error)
    }

    dispatch(get({data: numReports})(actionType))
  }
}

const getReportReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportRecipientBudgetsContract = state.chainContracts.data.contracts.orgReportRecipientBudgetsContract
    const numReports = state.orgReportRecipientBudgetsReader.num
    let reportRecipientBudgetRefs: OrgReportRecipientBudgetsData = {}
    let actionType = OrgReportRecipientBudgetsReaderActionTypes.REF_FAILURE
    for (let i = 0; i < numReports; i++) {
       try {
         const ref = await orgReportRecipientBudgetsContract.getReportReference(i.toString())
         reportRecipientBudgetRefs[ref] = {
           num: 0,
           data: {}
         }
         actionType = OrgReportRecipientBudgetsReaderActionTypes.REF_SUCCESS
       } catch (error) {
         console.log('getReportReferences error', error)
       }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reportRecipientBudgetRefs}})(actionType))
  }
}

const getNumReportRecipientBudgets = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const orgReportRecipientBudgetsContract = state.chainContracts.data.contracts.orgReportRecipientBudgetsContract
    let actionType = OrgReportRecipientBudgetsReaderActionTypes.NUMBUDGET_FAILURE

    const reports = state.orgReportRecipientBudgetsReader.data as OrgReportRecipientBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
       try {
         //console.log('Report key: ', reportKey)
         const num = await orgReportRecipientBudgetsContract.getNumRecipientBudgets(reportKey)
         reports[reportKey].num = num.toNumber()
         //console.log( 'Num reports for ', thisKey, ' ; ', orgs[thisKey].num)
         actionType = OrgReportRecipientBudgetsReaderActionTypes.NUMBUDGET_SUCCESS
       } catch (error) {
         console.log('getNumReportRecipientBudgets error', error)
       }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getReportRecipientBudgetRefs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportRecipientBudgetsContract = state.chainContracts.data.contracts.orgReportRecipientBudgetsContract
    let actionType = OrgReportRecipientBudgetsReaderActionTypes.BUDGETREF_FAILURE

    const reports = state.orgReportRecipientBudgetsReader.data as OrgReportRecipientBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const numRecipientBudgets = reports[reportKey].num
      for (let j = 0; j < numRecipientBudgets; j++) {
         try {
            const ref = await orgReportRecipientBudgetsContract.getRecipientBudgetReference(reportKey, j)
            //console.log ('Report ref: ', ref)
            reports[reportKey].data[ref] = {
              reportRef: '',
              budgetRef: ref,
              orgRef: '',
              budgetLine: '',
              finance: {
                value: 0,
                status: '',
                start: '',
                end: ''
              }
            }
            actionType = OrgReportRecipientBudgetsReaderActionTypes.BUDGETREF_SUCCESS
          } catch (error) {
            console.log('getReportRecipientBudgetRefs error', error)
          }
      }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getReportRecipientBudgetData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgReportRecipientBudgetsContract = state.chainContracts.data.contracts.orgReportRecipientBudgetsContract
    let actionType = OrgReportRecipientBudgetsReaderActionTypes.BUDGET_FAILURE

    const reports = state.orgReportRecipientBudgetsReader.data as OrgReportRecipientBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const budgetRefKeys = Object.keys(reports[reportKey].data)
      for (let j = 0; j < budgetRefKeys.length; j++) {
       const budgetRefKey = budgetRefKeys[j]
       //console.log('Report Data for org ', orgKey, ' ref key ', reportRefKey)
       try {
          const budgetData = await orgReportRecipientBudgetsContract.getRecipientBudget(reportKey, budgetRefKey)
          //const reportData = await orgReportsContract.getCurrency(orgKey, reportRefKey)
          //console.log('Report Data for org ', orgKey, ' ref key ', reportRefKey, ' is data ', reportData)
          reports[reportKey].data[budgetRefKey] = budgetData
          actionType = OrgReportRecipientBudgetsReaderActionTypes.BUDGET_SUCCESS
        } catch (error) {
          console.log('getReportRecipientBudgetData error', error)
        }
      }
    }
    //console.log('orgReportRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}
