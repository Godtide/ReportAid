//import { Action } from 'redux'
import { combineReducers, Reducer, Store, createStore, applyMiddleware } from 'redux'
//import thunkMiddleware, { ThunkAction } from 'redux-thunk'
import ReduxThunk, { ThunkAction } from 'redux-thunk'

import { ActionProps, TxProps } from './types'

import { IATIOrgReportProps } from './IATI/IATIReader/organisations/orgs/types'
import { IATIOrganisationsReportProps } from './IATI/IATIReader/organisations/organisations/types'
import { IATIOrganisationReportProps } from './IATI/IATIReader/organisations/organisation/types'
import { IATIOrganisationDocReportProps } from './IATI/IATIReader/organisations/organisationDocs/types'
import { IATIOrganisationBudgetReportProps } from './IATI/IATIReader/organisations/organisationBudgets/types'
import { IATIOrganisationExpenditureReportProps } from './IATI/IATIReader/organisations/organisationExpenditure/types'
import { IATIOrganisationRecipientBudgetReportProps } from './IATI/IATIReader/organisations/organisationRecipientBudgets/types'
import { IATIOrganisationRegionBudgetReportProps } from './IATI/IATIReader/organisations/organisationRegionBudgets/types'
import { IATIOrganisationCountryBudgetReportProps } from './IATI/IATIReader/organisations/organisationCountryBudgets/types'

import { InfoProps } from './info/types'
import { InfoProps as BlockchainInfoProps } from  './blockchain/info/types'
import { AccountProps } from  './blockchain/account/types'
import { ContractProps } from  './blockchain/contracts/types'

import { infoReducer } from './blockchain/info/reducer'
import { reducer as accountReducer } from './blockchain/account/reducer'
import { reducer as contractReducer } from './blockchain/contracts/reducer'
import { reducer as aboutReducer } from './info/about/reducer'
import { reducer as homeReducer } from './info/home/reducer'
import { reducer as helpReducer } from './info/help/reducer'
import { reducer as overviewReducer } from './info/overview/reducer'
import { reducer as IATIWriterInfoReducer } from './info/IATIWriter/reducer'
import { reducer as IATIReaderInfoReducer } from './info/IATIReader/reducer'

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
  chainInfo: BlockchainInfoProps
  chainAccount: AccountProps
  chainContracts: ContractProps
  about: InfoProps
  home: InfoProps
  help: InfoProps
  overview: InfoProps
  writer: InfoProps
  reader: InfoProps
  organisationsWriterForms: TxProps
  orgsReader: IATIOrgReportProps
  organisationsReader: IATIOrganisationsReportProps
  organisationReader: IATIOrganisationReportProps
  organisationDocsReader: IATIOrganisationDocReportProps
  organisationBudgetsReader: IATIOrganisationBudgetReportProps
  organisationExpenditureReader: IATIOrganisationExpenditureReportProps
  organisationRecipientBudgetsReader: IATIOrganisationRecipientBudgetReportProps
  organisationRegionBudgetsReader: IATIOrganisationRegionBudgetReportProps
  organisationCountryBudgetsReader: IATIOrganisationCountryBudgetReportProps
}

export const rootReducer: Reducer<ApplicationState, ActionProps> = combineReducers<ApplicationState, ActionProps>({
  chainInfo: infoReducer,
  chainAccount: accountReducer,
  chainContracts: contractReducer,
  about: aboutReducer,
  home: homeReducer,
  help: helpReducer,
  overview: overviewReducer,
  writer: IATIWriterInfoReducer,
  reader: IATIReaderInfoReducer,
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
