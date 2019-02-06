import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'

import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps } from '../../../types'
import { OrgCountryBudgetsReaderActionTypes, OrgCountryBudgetsReaderProps, OrgCountryBudgetsData } from './types'

export const getCountryBudgets = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNums())
    await dispatch(getReferences())
    await dispatch(getNumCountryBudgets())
    await dispatch(getCountryBudgetRefs())
    dispatch(getCountryBudgetData())
  }
}

const get = (payload: PayloadProps): Function => {
  return (actionType: OrgCountryBudgetsReaderActionTypes): OrgCountryBudgetsReaderProps => {
    const getProps = storeAction(actionType)(payload) as OrgCountryBudgetsReaderProps
    return getProps
  }
}

const getNums = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgCountryBudgetsContract = state.chainContracts.data.contracts.orgCountryBudgetsContract
    let actionType = OrgCountryBudgetsReaderActionTypes.NUM_FAILURE
    let nums = { num: 0 }
    try {
      const num = await orgCountryBudgetsContract.getNums()
      //console.log('Num orgs: ', num)
      nums.num = num.toNumber()
      actionType = OrgCountryBudgetsReaderActionTypes.NUM_SUCCESS
    } catch (error) {
      console.log('getNums error', error)
    }

    dispatch(get({data: nums})(actionType))
  }
}

const getReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgCountryBudgetsContract = state.chainContracts.data.contracts.orgCountryBudgetsContract
    const nums = state.orgCountryBudgetsReader.num
    let reportCountryBudgetRefs: OrgCountryBudgetsData = {}
    let actionType = OrgCountryBudgetsReaderActionTypes.REF_FAILURE
    for (let i = 0; i < nums; i++) {
       try {
         const ref = await orgCountryBudgetsContract.getReference(i.toString())
         reportCountryBudgetRefs[ref] = {
           num: 0,
           data: {}
         }
         actionType = OrgCountryBudgetsReaderActionTypes.REF_SUCCESS
       } catch (error) {
         console.log('getReferences error', error)
       }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reportCountryBudgetRefs}})(actionType))
  }
}

const getNumCountryBudgets = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const orgCountryBudgetsContract = state.chainContracts.data.contracts.orgCountryBudgetsContract
    let actionType = OrgCountryBudgetsReaderActionTypes.NUMBUDGET_FAILURE

    const reports = state.orgCountryBudgetsReader.data as OrgCountryBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
       try {
         //console.log(' key: ', reportKey)
         const num = await orgCountryBudgetsContract.getNumCountryBudgets(reportKey)
         reports[reportKey].num = num.toNumber()
         //console.log( 'Num reports for ', thisKey, ' ; ', orgs[thisKey].num)
         actionType = OrgCountryBudgetsReaderActionTypes.NUMBUDGET_SUCCESS
       } catch (error) {
         console.log('getNumCountryBudgets error', error)
       }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getCountryBudgetRefs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgCountryBudgetsContract = state.chainContracts.data.contracts.orgCountryBudgetsContract
    let actionType = OrgCountryBudgetsReaderActionTypes.BUDGETREF_FAILURE

    const reports = state.orgCountryBudgetsReader.data as OrgCountryBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const numCountryBudgets = reports[reportKey].num
      for (let j = 0; j < numCountryBudgets; j++) {
         try {
            const ref = await orgCountryBudgetsContract.getCountryBudgetReference(reportKey, j)
            //console.log (' ref: ', ref)
            reports[reportKey].data[ref] = {
              report: {
                reportRef: '',
                orgRef: ''
              },
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
            actionType = OrgCountryBudgetsReaderActionTypes.BUDGETREF_SUCCESS
          } catch (error) {
            console.log('getCountryBudgetRefs error', error)
          }
      }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getCountryBudgetData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgCountryBudgetsContract = state.chainContracts.data.contracts.orgCountryBudgetsContract
    let actionType = OrgCountryBudgetsReaderActionTypes.BUDGET_FAILURE

    const reports = state.orgCountryBudgetsReader.data as OrgCountryBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const budgetRefKeys = Object.keys(reports[reportKey].data)
      for (let j = 0; j < budgetRefKeys.length; j++) {
       const budgetRefKey = budgetRefKeys[j]
       //console.log(' Data for org ', orgKey, ' ref key ', reportRefKey)
       try {
          const budgetData = await orgCountryBudgetsContract.getCountryBudget(reportKey, budgetRefKey)
          //const reportData = await orgsContract.getCurrency(orgKey, reportRefKey)
          //console.log(' Data for org ', orgKey, ' ref key ', reportRefKey, ' is data ', reportData)
          reports[reportKey].data[budgetRefKey] = budgetData
          actionType = OrgCountryBudgetsReaderActionTypes.BUDGET_SUCCESS
        } catch (error) {
          console.log('getCountryBudgetData error', error)
        }
      }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}
