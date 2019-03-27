import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIActivityTransactionProps,
         IATIActivityTransactionData,
         IATIReportActionTypes,
         IATIActivityTransactionReportProps,
         ActivitiesReportProps } from '../../../types'

import { read } from '../../../../actions'

interface RecordProps {
  activitiesRef: string
  activityRef: string
  transactionRef: string
}

interface ActivityTransactionsDataProps {
  data: Array<IATIActivityTransactionData>
  successActionType: IATIReportActionTypes
  failureActionType: IATIReportActionTypes
}

type ActivityTransactionProps = RecordProps & ActivityTransactionsDataProps

const getThisActivityTransaction = (props: ActivityTransactionProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityTransactionsContract = state.chainContracts.data.contracts.activityTransactions
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef
    const transactionRef = props.transactionRef
    const data =  props.data

    let transactionData: IATIActivityTransactionReportProps = {
      data: { activitiesRef: activitiesRef,
              activityRef: activityRef,
              data: data }
    }

    let actionType =  props.successActionType

    try {

      const thisTransaction: IATIActivityTransactionProps = await activityTransactionsContract.getTransaction(activitiesRef, activityRef, transactionRef)

      //console.log('Transaction: ', thisTransaction)

      transactionData.data.data[data.length] = {
        transactionRef: transactionRef,
        transactionType: thisTransaction.transactionType,
        disbursementChannel: thisTransaction.disbursementChannel,
        flowType: thisTransaction.flowType,
        tiedStatus: thisTransaction.tiedStatus,
        financeType: ethers.utils.bigNumberify(thisTransaction.financeType).toNumber(),
        aidType: ethers.utils.parseBytes32String(thisTransaction.aidType),
        date: ethers.utils.parseBytes32String(thisTransaction.date),
        value: ethers.utils.bigNumberify(thisTransaction.value.value).toNumber(),
        valueDate: ethers.utils.parseBytes32String(thisTransaction.value.date),
        currency: ethers.utils.parseBytes32String(thisTransaction.value.currency),
        providerOrgType: thisTransaction.providerOrg.orgType,
        providerOrgRef: thisTransaction.providerOrg.orgRef,
        providerActivityRef: thisTransaction.providerOrg.activityRef,
        receiverOrgType: thisTransaction.receiverOrg.orgType,
        receiverOrgRef: thisTransaction.receiverOrg.orgRef,
        receiverActivityRef: thisTransaction.receiverOrg.activityRef,
        sectorDacCode: ethers.utils.bigNumberify(thisTransaction.sectorDacCode).toNumber(),
        territory: ethers.utils.parseBytes32String(thisTransaction.territory),
        description: thisTransaction.description
      }
    } catch (error) {
      actionType = props.failureActionType
      console.log('getActivityTransactions error', error)
    }

    dispatch(read({data: transactionData})(actionType))
  }
}

export const getActivityTransactionRecord = (props: RecordProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>) => {

    let indexRecordProps: ActivityTransactionProps = {
      activitiesRef: props.activitiesRef,
      activityRef: props.activityRef,
      transactionRef: props.transactionRef,
      data: [],
      successActionType: IATIReportActionTypes.ACTIVITYTRANSACTION_SUCCESS,
      failureActionType: IATIReportActionTypes.ACTIVITYTRANSACTION_FAILURE
    }

    dispatch(getThisActivityTransaction(indexRecordProps))
  }
}

export const getActivityTransactions = (props: ActivitiesReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityTransactionsContract = state.chainContracts.data.contracts.activityTransactions

    let indexRecordProps: ActivityTransactionProps = {
      activitiesRef: props.activitiesRef,
      activityRef: props.activityRef,
      transactionRef: "",
      data: [],
      successActionType: IATIReportActionTypes.ACTIVITYTRANSACTION_SUCCESS,
      failureActionType: IATIReportActionTypes.ACTIVITYTRANSACTION_FAILURE
    }

    try {
      const num = await activityTransactionsContract.getNum(indexRecordProps.activitiesRef, indexRecordProps.activityRef)
      const nums = num.toNumber()
      for (let i = 0; i < nums; i++) {

        indexRecordProps.transactionRef = await activityTransactionsContract.getReference(indexRecordProps.activitiesRef, indexRecordProps.activityRef, i.toString())
        dispatch(getThisActivityTransaction(indexRecordProps))
      }
    } catch (error) {
      console.log('getActivityTransaction error', error)
    }
  }
}

export const getActivityTransactionRefs = (props: ActivitiesReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityTransactionsContract = state.chainContracts.data.contracts.activityTransactions
    const activitiesRef = props.activitiesRef
    const activityRef = props.activityRef

    let keysData: Array<string> = []
    let actionType = IATIReportActionTypes.ACTIVITYTRANSACTIONPICKER_SUCCESS

    try {
      const num = await activityTransactionsContract.getNum(activitiesRef, activityRef)
      const nums = num.toNumber()
      for (let i = 0; i < nums; i++) {
        const transactionRef = await activityTransactionsContract.getReference(activitiesRef, activityRef, i.toString())
        keysData.push(transactionRef)
      }
    } catch (error) {
      actionType = IATIReportActionTypes.ACTIVITYTRANSACTION_FAILURE
      console.log('getActivityTransactionRefs error', error)
    }

    //console.log('KeysData: ', keysData, actionType)
    dispatch(read({data: keysData})(actionType))
  }
}
