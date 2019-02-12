import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'

import { read } from '../actions'

import { IATIOrganisationsProps } from '../../../types'
import { IATIReportActionTypes, IATIOrganisationsData } from '../types'

export const getOrganisations = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const organisationsContract = state.chainContracts.data.contracts.organisations

    const report = state.organisationsReport.data as IATIOrganisationsData
    let actionType = IATIReportActionTypes.ORGANISATIONS_FAILURE
    try {
      const num = await organisationsContract.getNumOrganisations()
      const numOrganisations = num.toNumber()
      for (let i = 0; i < numOrganisations; i++) {
         const ref = await organisationsContract.getOrganisationsReference(i.toString())
         const organisations: IATIOrganisationsProps = await organisationsContract.getOrganisations(ref)
         report[ref].IATIOrganisations = organisations
         actionType = IATIReportActionTypes.ORGANISATIONS_SUCCESS
      }
    } catch (error) {
      console.log('getBudgets error', error)
    }

    dispatch(read({data: report})(actionType))

  }
}
