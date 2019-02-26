//import { Action } from 'redux'
import { combineReducers, Reducer, Store, createStore, applyMiddleware } from 'redux'
//import thunkMiddleware, { ThunkAction } from 'redux-thunk'
import ReduxThunk, { ThunkAction } from 'redux-thunk'

import { ActionProps, TxProps } from './types'

import { KeyProps } from './helpers/keys/types'
import { FormProps } from './helpers/forms/types'

import { IATIOrgReportProps } from './IATI/IATIReader/organisations/types'
import { IATIOrganisationsReportProps } from './IATI/IATIReader/organisations/types'
import { IATIOrganisationReportProps } from './IATI/IATIReader/organisations/types'
import { IATIOrganisationDocReportProps } from './IATI/IATIReader/organisations/organisationDocs/types'
import { IATIBudgetReportProps } from './IATI/IATIReader/types'

import { InfoPageProps } from './info/types'
import { ChainDataProps } from  './blockchain/data/types'
import { AccountProps } from  './blockchain/account/types'
import { ContractProps } from  './blockchain/contracts/types'

import { reducer as keyReducer } from './helpers/keys/reducer'
import { reducer as formReducer } from './helpers/forms/reducer'

import { reducer as chainDataReducer } from './blockchain/data/reducer'
import { reducer as accountReducer } from './blockchain/account/reducer'
import { reducer as contractReducer } from './blockchain/contracts/reducer'
import { reducer as infoReducer } from './info/reducer'

import { reducer as orgsReaderReducer } from './IATI/IATIReader/organisations/orgs/reducer'
import { reducer as organisationsReaderReducer } from './IATI/IATIReader/organisations/organisations/reducer'
import { reducer as organisationReaderReducer } from './IATI/IATIReader/organisations/organisation/reducer'
import { reducer as organisationDocsReaderReducer } from './IATI/IATIReader/organisations/organisationDocs/reducer'
import { reducer as organisationBudgetsReaderReducer } from './IATI/IATIReader/organisations/organisationBudgets/reducer'
import { reducer as organisationExpenditureReaderReducer } from './IATI/IATIReader/organisations/organisationExpenditure/reducer'
import { reducer as organisationRecipientBudgetsReaderReducer } from './IATI/IATIReader/organisations/organisationRecipientBudgets/reducer'
import { reducer as organisationRegionBudgetsReaderReducer } from './IATI/IATIReader/organisations/organisationRegionBudgets/reducer'
import { reducer as organisationCountryBudgetsReaderReducer } from './IATI/IATIReader/organisations/organisationCountryBudgets/reducer'

import { reducer as organisationsWriterReducer } from './IATI/IATIWriter/organisations/reducer'

export type ThunkResult<R> = ThunkAction<R, ApplicationState, null, any>

export interface ApplicationState {
  chainInfo: ChainDataProps
  chainAccount: AccountProps
  chainContracts: ContractProps
  info: InfoPageProps
  keys: KeyProps
  forms: FormProps
  organisationsWriterForms: TxProps
  orgsReader: IATIOrgReportProps
  organisationsReader: IATIOrganisationsReportProps
  organisationReader: IATIOrganisationReportProps
  organisationDocsReader: IATIOrganisationDocReportProps
  organisationBudgetsReader: IATIBudgetReportProps
  organisationExpenditureReader: IATIBudgetReportProps
  organisationRecipientBudgetsReader: IATIBudgetReportProps
  organisationRegionBudgetsReader: IATIBudgetReportProps
  organisationCountryBudgetsReader: IATIBudgetReportProps
}

export const rootReducer: Reducer<ApplicationState, ActionProps> = combineReducers<ApplicationState, ActionProps>({
  chainInfo: chainDataReducer,
  chainAccount: accountReducer,
  chainContracts: contractReducer,
  info: infoReducer,
  keys: keyReducer,
  forms: formReducer,
  organisationsWriterForms: organisationsWriterReducer,
  orgsReader: orgsReaderReducer,
  organisationsReader: organisationsReaderReducer,
  organisationReader: organisationReaderReducer,
  organisationDocsReader: organisationDocsReaderReducer,
  organisationBudgetsReader: organisationBudgetsReaderReducer,
  organisationExpenditureReader: organisationExpenditureReaderReducer,
  organisationRecipientBudgetsReader: organisationRecipientBudgetsReaderReducer,
  organisationRegionBudgetsReader: organisationRegionBudgetsReaderReducer,
  organisationCountryBudgetsReader: organisationCountryBudgetsReaderReducer
})

export function configureStore(
  initialState: ApplicationState
): Store<ApplicationState> {

  // create the redux-saga middleware
  const store = createStore(
    rootReducer,
    initialState,
    applyMiddleware(ReduxThunk)
  )

  return store
}
