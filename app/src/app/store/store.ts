//import { Action } from 'redux'
import { combineReducers, Reducer, Store, createStore, applyMiddleware } from 'redux'
//import thunkMiddleware, { ThunkAction } from 'redux-thunk'
import ReduxThunk, { ThunkAction } from 'redux-thunk'

import { ActionProps, TxProps } from './types'

import { IATIOrganisationsReportProps } from './IATI/IATIReader/organisations/types'

/*import { OrgReaderProps } from './IATI/IATIReader/orgs/types'
import { OrganisationReaderProps } from './IATI/IATIReader/organisations/types'
import { OrganisationDocsReaderProps } from './IATI/IATIReader/organisationDocs/types'
import { OrganisationBudgetsReaderProps } from './IATI/IATIReader/organisationBudgets/types'
import { OrganisationExpenditureReaderProps } from './IATI/IATIReader/organisationExpenditure/types'
import { OrganisationRecipientBudgetsReaderProps } from './IATI/IATIReader/organisationRecipientBudgets/types'
import { OrganisationRegionBudgetsReaderProps } from './IATI/IATIReader/organisationRegionBudgets/types'
import { OrganisationCountryBudgetsReaderProps } from './IATI/IATIReader/organisationCountryBudgets/types'*/

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

import { reducer as organisationsReaderReducer } from './IATI/IATIReader/organisations/reducer'

/*
import { reducer as orgReaderReducer } from './IATI/IATIReader/orgs/reducer'
import { reducer as organisationsReaderReducer } from './IATI/IATIReader/organisations/reducer'
import { reducer as organisationReaderReducer } from './IATI/IATIReader/organisation/reducer'
import { reducer as organisationDocsReaderReducer } from './IATI/IATIReader/organisationDocs/reducer'
import { reducer as organisationBudgetsReaderReducer } from './IATI/IATIReader/organisationBudgets/reducer'
import { reducer as organisationExpenditureReaderReducer } from './IATI/IATIReader/organisationExpenditure/reducer'
import { reducer as organisationRecipientBudgetsReaderReducer } from './IATI/IATIReader/organisationRecipientBudgets/reducer'
import { reducer as organisationRegionBudgetsReaderReducer } from './IATI/IATIReader/organisationRegionBudgets/reducer'
import { reducer as organisationCountryBudgetsReaderReducer } from './IATI/IATIReader/organisationCountryBudgets/reducer'
*/

import { reducer as organisationsWriterReducer } from './IATI/IATIWriter/organisations/reducer'

/*import { reducer as organisationsWriterReducer } from './IATI/IATIWriter/organisations/reducer'
import { reducer as organisationWriterReducer } from './IATI/IATIWriter/organisation/reducer'
import { reducer as organisationDocsWriterReducer } from './IATI/IATIWriter/organisationDocs/reducer'
import { reducer as organisationBudgetsWriterReducer } from './IATI/IATIWriter/organisationBudgets/reducer'
import { reducer as organisationExpenditureWriterReducer } from './IATI/IATIWriter/organisationExpenditure/reducer'
import { reducer as organisationRecipientBudgetsWriterReducer } from './IATI/IATIWriter/organisationRecipientBudgets/reducer'
import { reducer as organisationRegionBudgetsWriterReducer } from './IATI/IATIWriter/organisationRegionBudgets/reducer'
import { reducer as organisationCountryBudgetsWriterReducer } from './IATI/IATIWriter/organisationCountryBudgets/reducer'*/

export type ThunkResult<R> = ThunkAction<R, ApplicationState, null, any>

// The top-level state object
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
  orgForm: TxProps
  organisationsForm: TxProps
  organisationForm: TxProps
  organisationDocsForm: TxProps
  organisationBudgetsForm: TxProps
  organisationExpenditureForm: TxProps
  organisationRecipientBudgetsForm: TxProps
  organisationRegionBudgetsForm: TxProps
  organisationCountryBudgetsForm: TxProps
  organisationsReport: IATIOrganisationsReportProps
  /*orgReader: OrgReaderProps
  organisationsReader: OrganisationsReaderProps
  organisationReader: OrganisationReaderProps
  organisationDocsReader: OrganisationDocsReaderProps
  organisationBudgetsReader: OrganisationBudgetsReaderProps
  organisationExpenditureReader: OrganisationExpenditureReaderProps
  organisationRecipientBudgetsReader: OrganisationRecipientBudgetsReaderProps
  organisationRegionBudgetsReader: OrganisationRegionBudgetsReaderProps
  organisationCountryBudgetsReader: OrganisationCountryBudgetsReaderProps*/
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
  orgForm: organisationsWriterReducer,
  organisationsForm: organisationsWriterReducer,
  organisationForm: organisationsWriterReducer,
  organisationDocsForm: organisationsWriterReducer,
  organisationBudgetsForm: organisationsWriterReducer,
  organisationExpenditureForm: organisationsWriterReducer,
  organisationRecipientBudgetsForm: organisationsWriterReducer,
  organisationRegionBudgetsForm: organisationsWriterReducer,
  organisationCountryBudgetsForm: organisationsWriterReducer,
  organisationsReport: organisationsReaderReducer
  /* orgReader: orgReaderReducer,
  organisationsReader: organisationsReaderReducer,
  organisationReader: organisationReaderReducer,
  organisationDocsReader: organisationDocsReaderReducer,
  organisationBudgetsReader: organisationBudgetsReaderReducer,
  organisationExpenditureReader: organisationExpenditureReaderReducer,
  organisationRecipientBudgetsReader: organisationRecipientBudgetsReaderReducer,
  organisationRegionBudgetsReader: organisationRegionBudgetsReaderReducer,
  organisationCountryBudgetsReader: organisationCountryBudgetsReaderReducer,*/
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
