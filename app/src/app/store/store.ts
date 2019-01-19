//import { Action } from 'redux'
import { combineReducers, Reducer, Store, createStore, applyMiddleware } from 'redux'
//import thunkMiddleware, { ThunkAction } from 'redux-thunk'
import ReduxThunk, { ThunkAction } from 'redux-thunk'

import { ActionProps } from './types'
import { OrgWriterProps } from './IATI/IATIWriter/organisationWriter/types'
import { OrgReportWriterProps } from './IATI/IATIWriter/organisationReportsWriter/types'
import { OrgReaderProps } from './IATI/IATIReader/organisationReader/types'
import { OrgReportReaderProps } from './IATI/IATIReader/organisationReportReader/types'
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
import { reducer as orgWriterReducer } from './IATI/IATIWriter/organisationWriter/reducer'
import { reducer as orgReportsWriterReducer } from './IATI/IATIWriter/organisationReportsWriter/reducer'
import { reducer as orgReaderReducer } from './IATI/IATIReader/organisationReader/reducer'
import { reducer as orgReportsReaderReducer } from './IATI/IATIReader/organisationReportReader/reducer'

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
  orgForm: OrgWriterProps
  orgReportsForm: OrgReportWriterProps
  orgReader: OrgReaderProps
  orgReportsReader: OrgReportReaderProps
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
  orgForm: orgWriterReducer,
  orgReportsForm: orgReportsWriterReducer,
  orgReader: orgReaderReducer,
  orgReportsReader: orgReportsReaderReducer
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
