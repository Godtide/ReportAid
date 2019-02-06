import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'

import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps } from '../../../types'
import { OrganisationRecipientBudgetsReaderActionTypes, OrganisationRecipientBudgetsReaderProps, OrganisationRecipientBudgetsData } from './types'

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
  return (actionType: OrganisationRecipientBudgetsReaderActionTypes): OrganisationRecipientBudgetsReaderProps => {
    const getProps = storeAction(actionType)(payload) as OrganisationRecipientBudgetsReaderProps
    return getProps
  }
}

const getNums = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgRecipientBudgetsContract = state.chainContracts.data.contracts.orgRecipientBudgetsContract
    let actionType = OrganisationRecipientBudgetsReaderActionTypes.NUM_FAILURE
    let nums = { num: 0 }
    try {
      const num = await orgRecipientBudgetsContract.getNums()
      //console.log('Num orgs: ', num)
      nums.num = num.toNumber()
      actionType = OrganisationRecipientBudgetsReaderActionTypes.NUM_SUCCESS
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
    let reportRecipientBudgetRefs: OrganisationRecipientBudgetsData = {}
    let actionType = OrganisationRecipientBudgetsReaderActionTypes.REF_FAILURE
    for (let i = 0; i < nums; i++) {
       try {
         const ref = await orgRecipientBudgetsContract.getReference(i.toString())
         reportRecipientBudgetRefs[ref] = {
           num: 0,
           data: {}
         }
         actionType = OrganisationRecipientBudgetsReaderActionTypes.REF_SUCCESS
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
    let actionType = OrganisationRecipientBudgetsReaderActionTypes.NUMBUDGET_FAILURE

    const reports = state.orgRecipientBudgetsReader.data as OrganisationRecipientBudgetsData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
       try {
         //console.log(' key: ', reportKey)
         const num = await orgRecipientBudgetsContract.getNumRecipientBudgets(reportKey)
         reports[reportKey].num = num.toNumber()
         //console.log( 'Num reports for ', thisKey, ' ; ', orgs[thisKey].num)
         actionType = OrganisationRecipientBudgetsReaderActionTypes.NUMBUDGET_SUCCESS
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
    let actionType = OrganisationRecipientBudgetsReaderActionTypes.BUDGETREF_FAILURE

    const reports = state.orgRecipientBudgetsReader.data as OrganisationRecipientBudgetsData
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
            actionType = OrganisationRecipientBudgetsReaderActionTypes.BUDGETREF_SUCCESS
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
    let actionType = OrganisationRecipientBudgetsReaderActionTypes.BUDGET_FAILURE

    const reports = state.orgRecipientBudgetsReader.data as OrganisationRecipientBudgetsData
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
          actionType = OrganisationRecipientBudgetsReaderActionTypes.BUDGET_SUCCESS
        } catch (error) {
          console.log('getRecipientBudgetData error', error)
        }
      }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}
