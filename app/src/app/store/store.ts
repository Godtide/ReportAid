//import { Action } from 'redux'
import { combineReducers, Reducer, Store, createStore, applyMiddleware } from 'redux'
//import thunkMiddleware, { ThunkAction } from 'redux-thunk'
import ReduxThunk, { ThunkAction } from 'redux-thunk'

import { ActionProps, PayloadProps, TxProps } from './types'

import { KeyProps } from './helpers/keys/types'
import { FormProps } from './helpers/forms/types'

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

import { reducer as organisationsPickerReducer } from './IATI/IATIReader/organisations/organisations/reducer'
import { reducer as organisationPickerReducer } from './IATI/IATIReader/organisations/organisation/reducer'
import { reducer as orgsPickerReducer } from './IATI/IATIReader/organisations/orgs/reducer'
import { reducer as activitiesPickerReducer } from './IATI/IATIReader/activities/activities/reducer'
import { reducer as activityPickerReducer } from './IATI/IATIReader/activities/activity/reducer'
import { reducer as readerReducer } from './IATI/IATIReader/reducer'
import { reducer as writerReducer } from './IATI/IATIWriter/reducer'

export type ThunkResult<R> = ThunkAction<R, ApplicationState, null, any>

export interface ApplicationState {
  chainInfo: ChainDataProps
  chainAccount: AccountProps
  chainContracts: ContractProps
  info: InfoPageProps
  keys: KeyProps
  forms: FormProps
  writerForms: PayloadProps
  organisationsPicker: PayloadProps
  organisationPicker: PayloadProps
  orgsPicker: PayloadProps
  activitiesPicker: PayloadProps
  activityPicker: PayloadProps
  report: PayloadProps
}

export const rootReducer: Reducer<ApplicationState, ActionProps> = combineReducers<ApplicationState, ActionProps>({
  chainInfo: chainDataReducer,
  chainAccount: accountReducer,
  chainContracts: contractReducer,
  info: infoReducer,
  keys: keyReducer,
  forms: formReducer,
  writerForms: writerReducer,
  organisationsPicker: organisationsPickerReducer,
  organisationPicker: organisationPickerReducer,
  orgsPicker: orgsPickerReducer,
  activitiesPicker: activitiesPickerReducer,
  activityPicker: activityPickerReducer,
  report: readerReducer
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
