import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'
import { OrganisationsReportProps, IATIReportActionTypes, IATIBudgetReportProps, IATIBudgetProps } from '../../../types'

import { read } from '../../actions'

export const getExpenditure = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const expenditureContract = state.chainContracts.data.contracts.organisationExpenditure
    const organisationsRef = props.organisationsRef
    const organisationRef = props.organisationRef

    let expenditureReports: IATIBudgetReportProps = {
      data: { organisationsRef: organisationsRef,
              organisationRef: organisationRef,
              data: []
            }
    }

    let actionType = IATIReportActionTypes.TOTALEXPENDITURE_FAILURE
    try {
      const num = await expenditureContract.getNumExpenditures(props.organisationsRef, props.organisationRef)
      const numExpenditure = num.toNumber()
      for (let i = 0; i < numExpenditure; i++) {
        const expenditureRef = await expenditureContract.getExpenditureReference(organisationsRef,
                                                                                  organisationRef,
                                                                                  i.toString())
        const expenditure: IATIBudgetProps = await expenditureContract.getExpenditure(organisationsRef, organisationRef, expenditureRef)

        expenditureReports.data.data[i] = {
          expenditureKey: expenditureRef,
          expenditureLine: ethers.utils.parseBytes32String(expenditure.budgetLine),
          value: ethers.utils.bigNumberify(expenditure.finance.value).toNumber(),
          status: expenditure.finance.status,
          start: ethers.utils.parseBytes32String(expenditure.finance.start),
          end: ethers.utils.parseBytes32String(expenditure.finance.end)
        }

        actionType = IATIReportActionTypes.TOTALEXPENDITURE_SUCCESS
      }
    } catch (error) {
      console.log('getExpenditure error', error)
    }

    dispatch(read({data: expenditureReports})(actionType))
  }
}
