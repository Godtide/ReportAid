import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'

import { ActionProps } from '../../../types'

import { read } from '../actions'

import { IATIReportActionTypes, OrganisationsReportProps } from '../types'

export const getRecipientBudgets = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>,, getState: Function) => {
  }
}
