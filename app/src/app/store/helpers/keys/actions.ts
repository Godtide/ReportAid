import { ThunkDispatch } from 'redux-thunk'
import { ApplicationState } from '../../../store'

import { storeAction } from '../../actions'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ActionProps, PayloadProps } from '../../types'
import { KeyActionTypes, KeyData } from './types'

export const write = (payload: PayloadProps): Function => {
  return (actionType: KeyActionTypes): PayloadProps => {
    const getProps = storeAction(actionType)(payload) as PayloadProps
    return getProps
  }
}

export const newKey = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const keyData: KeyData = state.keys.data
    keyData.newKey = ethers.utils.formatBytes32String(shortid.generate())
    //console.log('New Key! ', keyData.newKey)
    await dispatch(write({data: keyData})(KeyActionTypes.NEWKEY_SUCCESS))
  }
}

export const setOrganisationsKey = (key: string) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const keyData: KeyData = state.keys.data
    keyData.organisations = key
    console.log('Set Organisations Key! ', keyData.organisations)
    await dispatch(write({data: keyData})(KeyActionTypes.ORGANISATIONS_SUCCESS))
  }
}

export const setOrganisationKey = (key: string) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const keyData: KeyData = state.keys.data
    keyData.organisation = key
    console.log('Set Organisation Key! ', keyData.organisations)
    await dispatch(write({data: keyData})(KeyActionTypes.ORGANISATION_SUCCESS))
  }
}
