import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'

import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps } from '../../../types'
import { OrgRegionBudgetsReaderActionTypes, OrgRegionBudgetsReaderProps, OrgRegionBudgetsData } from './types'

export const getRegionBudgets = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNums())
    await dispatch(getReferences())
    await dispatch(getNumRegionBudgets())
    await dispatch(getRegionBudgetRefs())
    dispatch(getRegionBudgetData())
  }
}

const get = (payload: PayloadProps): Function => {
  return (actionType: OrgRegionBudgetsReaderActionTypes): OrgRegionBudgetsReaderProps => {
    const getProps = storeAction(actionType)(payload) as OrgRegionBudgetsReaderProps
    return getProps
  }
}

const getNums = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgRegionBudgetsContract = state.chainContracts.data.contracts.orgRegionBudgetsContract
    let actionType = OrgRegionBudgetsReaderActionTypes.NUM_FAILURE
    let nums = { num: 0 }
    try {
      const num = await orgRegionBudgetsContract.getNums()
      //console.log('Num orgs: ', num)
      nums.num = num.toNumber()
      actionType = OrgRegionBudgetsReaderActionTypes.NUM_SUCCESS
    } catch (error) {
      console.log('getNums error', error)
    }

    dispatch(get({data: nums})(actionType))
  }
}

const getReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgRegionBudgetsContract = state.chainContracts.data.contracts.orgRegionBudgetsContract
    const nums = state.orgRegionBudgetsReader.num
    let reportRegionBudgetRefs: OrgRegionBudgetsData = {}
    let actionType = OrgRegionBudgetsReaderActionTypes.REF_FAILURE
    for (let i = 0; i < nums; i++) {
       try {
         const ref = await orgRegionBudgetsContract.getReference(i.toString())
         reportRegionBudgetRefs[ref] = {
           num: 0,
           data: {}
         }
         actionType = OrgRegionBudgetsReaderActionTypes.REF_SUCCESS
       } catch (error) {
         console.log('getReferences error', error)
       }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reportRegionBudgetRefs}})(actionType))
  }
}

const getNumRegionBudgets = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const orgRegionBudgetsContract = state.chainContracts.data.contracts.orgRegionBudgetsContract
    let actionType = OrgRegionBudgetsReaderActionTypes.NUMBUDGET_FAILURE

    const reports = state.orgRegionBudgetsReader.data as OrgRegionBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
       try {
         //console.log(' key: ', reportKey)
         const num = await orgRegionBudgetsContract.getNumRegionBudgets(reportKey)
         reports[reportKey].num = num.toNumber()
         //console.log( 'Num reports for ', thisKey, ' ; ', orgs[thisKey].num)
         actionType = OrgRegionBudgetsReaderActionTypes.NUMBUDGET_SUCCESS
       } catch (error) {
         console.log('getNumRegionBudgets error', error)
       }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getRegionBudgetRefs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgRegionBudgetsContract = state.chainContracts.data.contracts.orgRegionBudgetsContract
    let actionType = OrgRegionBudgetsReaderActionTypes.BUDGETREF_FAILURE

    const reports = state.orgRegionBudgetsReader.data as OrgRegionBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const numRegionBudgets = reports[reportKey].num
      for (let j = 0; j < numRegionBudgets; j++) {
         try {
            const ref = await orgRegionBudgetsContract.getRegionBudgetReference(reportKey, j)
            //console.log (' ref: ', ref)
            reports[reportKey].data[ref] = {
              report: {
                reportRef: '',
                orgRef: ''
              },
              budgetRef: ref,
              regionRef: 0,
              budgetLine: '',
              finance: {
                value: 0,
                status: 0,
                start: '',
                end: ''
              }
            }
            actionType = OrgRegionBudgetsReaderActionTypes.BUDGETREF_SUCCESS
          } catch (error) {
            console.log('getRegionBudgetRefs error', error)
          }
      }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getRegionBudgetData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgRegionBudgetsContract = state.chainContracts.data.contracts.orgRegionBudgetsContract
    let actionType = OrgRegionBudgetsReaderActionTypes.BUDGET_FAILURE

    const reports = state.orgRegionBudgetsReader.data as OrgRegionBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const budgetRefKeys = Object.keys(reports[reportKey].data)
      for (let j = 0; j < budgetRefKeys.length; j++) {
       const budgetRefKey = budgetRefKeys[j]
       //console.log(' Data for org ', orgKey, ' ref key ', reportRefKey)
       try {
          const budgetData = await orgRegionBudgetsContract.getRegionsBudget(reportKey, budgetRefKey)
          //const reportData = await orgsContract.getCurrency(orgKey, reportRefKey)
          //console.log(' Data for org ', orgKey, ' ref key ', reportRefKey, ' is data ', reportData)
          reports[reportKey].data[budgetRefKey] = budgetData
          actionType = OrgRegionBudgetsReaderActionTypes.BUDGET_SUCCESS
        } catch (error) {
          console.log('getRegionBudgetData error', error)
        }
      }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}
