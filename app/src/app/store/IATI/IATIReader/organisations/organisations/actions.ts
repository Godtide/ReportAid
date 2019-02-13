import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'

import { read } from '../actions'

import { IATIOrganisationsProps } from '../../../types'
import { IATIReportActionTypes, IATIOrganisationsReport, IATIOrganisationsData } from '../types'

export const getOrganisations = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const organisationsContract = state.chainContracts.data.contracts.organisations

    let organisationsData: IATIOrganisationsData = {}
    let actionType = IATIReportActionTypes.ORGANISATIONS_FAILURE
    try {
      const num = await organisationsContract.getNumOrganisations()
      const numOrganisations = num.toNumber()
      for (let i = 0; i < numOrganisations; i++) {
         const ref = await organisationsContract.getOrganisationsReference(i.toString())

         const organisations: IATIOrganisationsProps = await organisationsContract.getOrganisations(ref)
         const thisData: IATIOrganisationsReport = {
           IATIOrganisations: organisations,
           data: {}
         }
         organisationsData[ref] = thisData
         actionType = IATIReportActionTypes.ORGANISATIONS_SUCCESS
      }
    } catch (error) {
      console.log('getOrganisations error', error)
    }

    dispatch(read({data: organisationsData})(actionType))

  }
}
