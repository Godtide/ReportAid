import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIOrganisationsProps } from '../../../types'
import { IATIReportActionTypes } from '../types'
import { IATIOrganisationsReportProps } from './types'

import { read } from '../actions'

export const initialise = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const initData: IATIOrganisationsReportProps = { data: {} }
    await dispatch(read({data: initData})(IATIReportActionTypes.ORGANISATIONS_INIT))
  }
}

export const getOrganisations = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const organisationsContract = state.chainContracts.data.contracts.organisations

    let organisationsData: IATIOrganisationsReportProps = {
      data: {}
    }
    let actionType = IATIReportActionTypes.ORGANISATIONS_FAILURE
    try {
      const num = await organisationsContract.getNumOrganisations()
      const numOrganisations = num.toNumber()
      for (let i = 0; i < numOrganisations; i++) {
         const ref = await organisationsContract.getOrganisationsReference(i.toString())

         const organisations: IATIOrganisationsProps = await organisationsContract.getOrganisations(ref)
         organisationsData.data[ref] = organisations
         actionType = IATIReportActionTypes.ORGANISATIONS_SUCCESS
      }
    } catch (error) {
      console.log('getOrganisations error', error)
    }

    dispatch(read({data: organisationsData})(actionType))

  }
}
