import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'

import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps } from '../../../types'
import { OrgBudgetsReaderActionTypes, OrgBudgetsReaderProps, OrgBudgetsData } from './types'

export const getBudgets = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNums())
    await dispatch(getReferences())
    await dispatch(getNumBudgets())
    await dispatch(getBudgetRefs())
    dispatch(getBudgetData())
  }
}

const get = (payload: PayloadProps): Function => {
  return (actionType: OrgBudgetsReaderActionTypes): OrgBudgetsReaderProps => {
    const getProps = storeAction(actionType)(payload) as OrgBudgetsReaderProps
    return getProps
  }
}

const getNums = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgBudgetsContract = state.chainContracts.data.contracts.orgBudgetsContract
    let actionType = OrgBudgetsReaderActionTypes.NUM_FAILURE
    let nums = { num: 0 }
    try {
      const num = await orgBudgetsContract.getNums()
      //console.log('Num orgs: ', num)
      nums.num = num.toNumber()
      actionType = OrgBudgetsReaderActionTypes.NUM_SUCCESS
    } catch (error) {
      console.log('getNums error', error)
    }

    dispatch(get({data: nums})(actionType))
  }
}

const getReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgBudgetsContract = state.chainContracts.data.contracts.orgBudgetsContract
    const nums = state.orgBudgetsReader.num
    let reportBudgetRefs: OrgBudgetsData = {}
    let actionType = OrgBudgetsReaderActionTypes.REF_FAILURE
    for (let i = 0; i < nums; i++) {
       try {
         const ref = await orgBudgetsContract.getReference(i.toString())
         reportBudgetRefs[ref] = {
           num: 0,
           data: {}
         }
         actionType = OrgBudgetsReaderActionTypes.REF_SUCCESS
       } catch (error) {
         console.log('getReferences error', error)
       }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reportBudgetRefs}})(actionType))
  }
}

const getNumBudgets = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const orgBudgetsContract = state.chainContracts.data.contracts.orgBudgetsContract
    let actionType = OrgBudgetsReaderActionTypes.NUMBUDGET_FAILURE

    const reports = state.orgBudgetsReader.data as OrgBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
       try {
         //console.log(' key: ', reportKey)
         const num = await orgBudgetsContract.getNumTotalBudgets(reportKey)
         reports[reportKey].num = num.toNumber()
         //console.log( 'Num reports for ', thisKey, ' ; ', orgs[thisKey].num)
         actionType = OrgBudgetsReaderActionTypes.NUMBUDGET_SUCCESS
       } catch (error) {
         console.log('getNumBudgets error', error)
       }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getBudgetRefs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgBudgetsContract = state.chainContracts.data.contracts.orgBudgetsContract
    let actionType = OrgBudgetsReaderActionTypes.BUDGETREF_FAILURE

    const reports = state.orgBudgetsReader.data as OrgBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const numBudgets = reports[reportKey].num
      for (let j = 0; j < numBudgets; j++) {
         try {
            const ref = await orgBudgetsContract.getTotalBudgetReference(reportKey, j)
            //console.log (' ref: ', ref)
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
            actionType = OrgBudgetsReaderActionTypes.BUDGETREF_SUCCESS
          } catch (error) {
            console.log('getBudgetRefs error', error)
          }
      }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getBudgetData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgBudgetsContract = state.chainContracts.data.contracts.orgBudgetsContract
    let actionType = OrgBudgetsReaderActionTypes.BUDGET_FAILURE

    const reports = state.orgBudgetsReader.data as OrgBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const budgetRefKeys = Object.keys(reports[reportKey].data)
      for (let j = 0; j < budgetRefKeys.length; j++) {
       const budgetRefKey = budgetRefKeys[j]
       //console.log(' Data for org ', orgKey, ' ref key ', reportRefKey)
       try {
          const budgetData = await orgBudgetsContract.getTotalBudget(reportKey, budgetRefKey)
          //const reportData = await orgsContract.getCurrency(orgKey, reportRefKey)
          //console.log(' Data for org ', orgKey, ' ref key ', reportRefKey, ' is data ', reportData)
          reports[reportKey].data[budgetRefKey] = budgetData
          actionType = OrgBudgetsReaderActionTypes.BUDGET_SUCCESS
        } catch (error) {
          console.log('getBudgetData error', error)
        }
      }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}
