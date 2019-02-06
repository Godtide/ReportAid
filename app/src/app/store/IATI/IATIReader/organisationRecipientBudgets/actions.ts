import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'

import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps } from '../../../types'
import { OrgRecipientBudgetsReaderActionTypes, OrgRecipientBudgetsReaderProps, OrgRecipientBudgetsData } from './types'

export const getRecipientBudgets = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNums())
    await dispatch(getReferences())
    await dispatch(getNumRecipientBudgets())
    await dispatch(getRecipientBudgetRefs())
    dispatch(getRecipientBudgetData())
  }
}

const get = (payload: PayloadProps): Function => {
  return (actionType: OrgRecipientBudgetsReaderActionTypes): OrgRecipientBudgetsReaderProps => {
    const getProps = storeAction(actionType)(payload) as OrgRecipientBudgetsReaderProps
    return getProps
  }
}

const getNums = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgRecipientBudgetsContract = state.chainContracts.data.contracts.orgRecipientBudgetsContract
    let actionType = OrgRecipientBudgetsReaderActionTypes.NUM_FAILURE
    let nums = { num: 0 }
    try {
      const num = await orgRecipientBudgetsContract.getNums()
      //console.log('Num orgs: ', num)
      nums.num = num.toNumber()
      actionType = OrgRecipientBudgetsReaderActionTypes.NUM_SUCCESS
    } catch (error) {
      console.log('getNums error', error)
    }

    dispatch(get({data: nums})(actionType))
  }
}

const getReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgRecipientBudgetsContract = state.chainContracts.data.contracts.orgRecipientBudgetsContract
    const nums = state.orgRecipientBudgetsReader.num
    let reportRecipientBudgetRefs: OrgRecipientBudgetsData = {}
    let actionType = OrgRecipientBudgetsReaderActionTypes.REF_FAILURE
    for (let i = 0; i < nums; i++) {
       try {
         const ref = await orgRecipientBudgetsContract.getReference(i.toString())
         reportRecipientBudgetRefs[ref] = {
           num: 0,
           data: {}
         }
         actionType = OrgRecipientBudgetsReaderActionTypes.REF_SUCCESS
       } catch (error) {
         console.log('getReferences error', error)
       }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reportRecipientBudgetRefs}})(actionType))
  }
}

const getNumRecipientBudgets = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const orgRecipientBudgetsContract = state.chainContracts.data.contracts.orgRecipientBudgetsContract
    let actionType = OrgRecipientBudgetsReaderActionTypes.NUMBUDGET_FAILURE

    const reports = state.orgRecipientBudgetsReader.data as OrgRecipientBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
       try {
         //console.log(' key: ', reportKey)
         const num = await orgRecipientBudgetsContract.getNumRecipientBudgets(reportKey)
         reports[reportKey].num = num.toNumber()
         //console.log( 'Num reports for ', thisKey, ' ; ', orgs[thisKey].num)
         actionType = OrgRecipientBudgetsReaderActionTypes.NUMBUDGET_SUCCESS
       } catch (error) {
         console.log('getNumRecipientBudgets error', error)
       }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getRecipientBudgetRefs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgRecipientBudgetsContract = state.chainContracts.data.contracts.orgRecipientBudgetsContract
    let actionType = OrgRecipientBudgetsReaderActionTypes.BUDGETREF_FAILURE

    const reports = state.orgRecipientBudgetsReader.data as OrgRecipientBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const numRecipientBudgets = reports[reportKey].num
      for (let j = 0; j < numRecipientBudgets; j++) {
         try {
            const ref = await orgRecipientBudgetsContract.getRecipientBudgetReference(reportKey, j)
            //console.log (' ref: ', ref)
            reports[reportKey].data[ref] = {
              report: {
                reportRef: '',
                orgRef: ''
              },
              budgetRef: ref,
              orgRef: '',
              budgetLine: '',
              finance: {
                value: 0,
                status: 0,
                start: '',
                end: ''
              }
            }
            actionType = OrgRecipientBudgetsReaderActionTypes.BUDGETREF_SUCCESS
          } catch (error) {
            console.log('getRecipientBudgetRefs error', error)
          }
      }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getRecipientBudgetData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgRecipientBudgetsContract = state.chainContracts.data.contracts.orgRecipientBudgetsContract
    let actionType = OrgRecipientBudgetsReaderActionTypes.BUDGET_FAILURE

    const reports = state.orgRecipientBudgetsReader.data as OrgRecipientBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const budgetRefKeys = Object.keys(reports[reportKey].data)
      for (let j = 0; j < budgetRefKeys.length; j++) {
       const budgetRefKey = budgetRefKeys[j]
       //console.log(' Data for org ', orgKey, ' ref key ', reportRefKey)
       try {
          const budgetData = await orgRecipientBudgetsContract.getRecipientBudget(reportKey, budgetRefKey)
          //const reportData = await orgsContract.getCurrency(orgKey, reportRefKey)
          //console.log(' Data for org ', orgKey, ' ref key ', reportRefKey, ' is data ', reportData)
          reports[reportKey].data[budgetRefKey] = budgetData
          actionType = OrgRecipientBudgetsReaderActionTypes.BUDGET_SUCCESS
        } catch (error) {
          console.log('getRecipientBudgetData error', error)
        }
      }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}