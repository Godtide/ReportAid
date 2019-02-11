import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../../store'

import { ActionProps } from '../../../../types'

import { read } from '../actions'

import { IATIOrganisationsReportProps, IATIReportActionTypes, OrganisationsReportProps } from '../types'

export const getCountryBudgets = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
  }
}
