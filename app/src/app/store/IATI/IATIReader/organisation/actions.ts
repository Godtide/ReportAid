import { ThunkDispatch } from 'redux-thunk'

import { ApplicationState } from '../../../store'

import { ActionProps } from '../../../types'

import { get } from '../actions'

import { OrganisationBudgetsReaderActionTypes, OrganisationsReportProps } from '../types'

export const getOrganisation = (props: OrganisationsReportProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>,, getState: Function) => {
  }
}
