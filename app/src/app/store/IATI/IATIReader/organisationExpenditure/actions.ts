import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'

import { storeAction } from '../../../actions'

import { ActionProps, PayloadProps } from '../../../types'
import { OrganisationExpenditureReaderActionTypes, OrganisationExpenditureReaderProps, OrganisationExpenditureData } from './types'

export const getExpenditure = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {
    await dispatch(getNums())
    await dispatch(getReferences())
    await dispatch(getNumExpenditure())
    await dispatch(getExpenditureRefs())
    dispatch(getExpenditureData())
  }
}

const get = (payload: PayloadProps): Function => {
  return (actionType: OrganisationExpenditureReaderActionTypes): OrganisationExpenditureReaderProps => {
    const getProps = storeAction(actionType)(payload) as OrganisationExpenditureReaderProps
    return getProps
  }
}

const getNums = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgExpenditureContract = state.chainContracts.data.contracts.orgExpenditureContract
    let actionType = OrganisationExpenditureReaderActionTypes.NUM_FAILURE
    let nums = { num: 0 }
    try {
      const num = await orgExpenditureContract.getNums()
      //console.log('Num orgs: ', num)
      nums.num = num.toNumber()
      actionType = OrganisationExpenditureReaderActionTypes.NUM_SUCCESS
    } catch (error) {
      console.log('getNums error', error)
    }

    dispatch(get({data: nums})(actionType))
  }
}

const getReferences = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgExpenditureContract = state.chainContracts.data.contracts.orgExpenditureContract
    const nums = state.orgExpenditureReader.num
    let reportExpenditureRefs: OrganisationExpenditureData = {}
    let actionType = OrganisationExpenditureReaderActionTypes.REF_FAILURE
    for (let i = 0; i < nums; i++) {
       try {
         const ref = await orgExpenditureContract.getReference(i.toString())
         reportExpenditureRefs[ref] = {
           num: 0,
           data: {}
         }
         actionType = OrganisationExpenditureReaderActionTypes.REF_SUCCESS
       } catch (error) {
         console.log('getReferences error', error)
       }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reportExpenditureRefs}})(actionType))
  }
}

const getNumExpenditure = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const orgExpenditureContract = state.chainContracts.data.contracts.orgExpenditureContract
    let actionType = OrganisationExpenditureReaderActionTypes.NUMEXPENDITURE_FAILURE

    const reports = state.orgExpenditureReader.data as OrganisationExpenditureData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
       try {
         //console.log(' key: ', reportKey)
         const num = await orgExpenditureContract.getNumExpenditures(reportKey)
         reports[reportKey].num = num.toNumber()
         //console.log( 'Num reports for ', thisKey, ' ; ', orgs[thisKey].num)
         actionType = OrganisationExpenditureReaderActionTypes.NUMEXPENDITURE_SUCCESS
       } catch (error) {
         console.log('getNumExpenditure error', error)
       }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getExpenditureRefs = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgExpenditureContract = state.chainContracts.data.contracts.orgExpenditureContract
    let actionType = OrganisationExpenditureReaderActionTypes.EXPENDITUREREF_FAILURE

    const reports = state.orgExpenditureReader.data as OrganisationExpenditureData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const numExpenditure = reports[reportKey].num
      for (let j = 0; j < numExpenditure; j++) {
         try {
            const ref = await orgExpenditureContract.getExpenditureReference(reportKey, j)
            //console.log (' ref: ', ref)
            reports[reportKey].data[ref] = {
              report: {
                reportRef: '',
                orgRef: ''
              },
              expenditureRef: ref,
              expenditureLine: '',
              finance: {
                value: 0,
                status: 0,
                start: '',
                end: ''
              }
            }
            actionType = OrganisationExpenditureReaderActionTypes.EXPENDITUREREF_SUCCESS
          } catch (error) {
            console.log('getExpenditureRefs error', error)
          }
      }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}

const getExpenditureData = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const orgExpenditureContract = state.chainContracts.data.contracts.orgExpenditureContract
    let actionType = OrganisationExpenditureReaderActionTypes.EXPENDITURE_FAILURE

    const reports = state.orgExpenditureReader.data as OrganisationExpenditureData
    const reportKeys = Object.keys(reports)

    for (let i = 0; i < reportKeys.length; i++) {
      const reportKey = reportKeys[i]
      const expenditureRefKeys = Object.keys(reports[reportKey].data)
      for (let j = 0; j < expenditureRefKeys.length; j++) {
       const expenditureRefKey = expenditureRefKeys[j]
       //console.log(' Data for org ', orgKey, ' ref key ', reportRefKey)
       try {
          const expenditureData = await orgExpenditureContract.getExpenditure(reportKey, expenditureRefKey)
          //const reportData = await orgsContract.getCurrency(orgKey, reportRefKey)
          //console.log(' Data for org ', orgKey, ' ref key ', reportRefKey, ' is data ', reportData)
          reports[reportKey].data[expenditureRefKey] = expenditureData
          actionType = OrganisationExpenditureReaderActionTypes.EXPENDITURE_SUCCESS
        } catch (error) {
          console.log('getExpenditureData error', error)
        }
      }
    }
    //console.log('orgRefs: ', orgRefs)
    dispatch(get({data: {data: reports}})(actionType))
  }
}
