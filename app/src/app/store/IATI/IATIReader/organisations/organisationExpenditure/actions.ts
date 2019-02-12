import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { IATIOrganisationExpenditureProps } from '../../../types'
import { IATIReportActionTypes,
         IATITotalExpenditureReportProps,
         OrganisationsReportProps } from '../types'

import { read } from '../actions'

export const getExpenditure = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const expenditureContract = state.chainContracts.data.contracts.organisationExpenditure
    const organisationsRef = props.organisationsRef
    const organisationRef = props.organisationRef

    const expenditureReports = state.organisationsReport.data[organisationsRef].data[organisationRef].data.totalExpenditure as IATITotalExpenditureReportProps

    let actionType = IATIReportActionTypes.TOTALEXPENDITURE_FAILURE
    try {
      const num = await expenditureContract.getNumExpenditures(props.organisationsRef, props.organisationRef)
      const numExpenditure = num.toNumber()
      for (let i = 0; i < numExpenditure; i++) {
         const expenditureRef = await expenditureContract.getExpenditureReference(organisationsRef,
                                                                                  organisationRef,
                                                                                  i.toString())
         const expenditure: IATIOrganisationExpenditureProps = await expenditureContract.getExpenditure(organisationsRef,
                                                                                                        organisationRef,
                                                                                                        expenditureRef)
         expenditureReports[expenditureRef] = expenditure
         actionType = IATIReportActionTypes.TOTALEXPENDITURE_SUCCESS
      }
    } catch (error) {
      console.log('getExpenditure error', error)
    }

    dispatch(read({data: expenditureReports})(actionType))
  }
}
