import { ThunkDispatch } from 'redux-thunk'
import { ApplicationState } from '../../../store'

import { storeAction } from '../../actions'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ActionProps, PayloadProps } from '../../types'
import { KeyActionTypes, KeyTypes, KeyData, Keys } from './types'

export const write = (payload: PayloadProps): Function => {
  return (actionType: KeyActionTypes): PayloadProps => {
    const getProps = storeAction(actionType)(payload) as PayloadProps
    return getProps
  }
}

export const setKey = (props: Keys) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    let keyData: KeyData = state.keys.data
    let actionType = KeyActionTypes.NEW_FAILURE

    let key = props.key
    if (props.key == '') {
      key = ethers.utils.formatBytes32String(shortid.generate())
    }

    switch (props.keyType) {
      case KeyTypes.NEW: {
        keyData = {
          newKey: key,
          org: key,
          organisations:  key,
          organisation:  key,
          organisationBudget:  key,
          organisationCountryBudget:  key,
          organisationDoc:  key,
          organisationExpenditure:  key,
          organisationRecipientBudget:  key,
          organisationRegionBudget:  key,
          activities:  key,
          activity:  key,
          activityAdditional:  key
        }
        actionType = KeyActionTypes.NEW_SUCCESS
        break
      }
      case KeyTypes.ORG: {
        keyData.org = key
        actionType = KeyActionTypes.ORG_SUCCESS
        break
      }
      case KeyTypes.ORGANISATIONS: {
        keyData.organisations = key
        actionType = KeyActionTypes.ORGANISATIONS_SUCCESS
        break
      }
      case KeyTypes.ORGANISATION: {
        keyData.organisation = key
        actionType = KeyActionTypes.ORGANISATION_SUCCESS
        break
      }
      case KeyTypes.ORGANISATIONBUDGET: {
        keyData.organisationBudget = key
        actionType = KeyActionTypes.ORGANISATIONBUDGET_SUCCESS
        break
      }
      case KeyTypes.ORGANISATIONCOUNTRYBUDGET: {
        keyData.organisationCountryBudget = key
        actionType = KeyActionTypes.ORGANISATIONCOUNTRYBUDGET_SUCCESS
        break
      }
      case KeyTypes.ORGANISATIONDOC: {
        keyData.organisationDoc = key
        actionType = KeyActionTypes.ORGANISATIONDOC_SUCCESS
        break
      }
      case KeyTypes.ORGANISATIONEXPENDITURE: {
        keyData.organisationExpenditure = key
        actionType = KeyActionTypes.ORGANISATIONEXPENDITURE_SUCCESS
        break
      }
      case KeyTypes.ORGANISATIONRECIPIENTBUDGET: {
        keyData.organisationRecipientBudget = key
        actionType = KeyActionTypes.ORGANISATIONRECIPIENTBUDGET_SUCCESS
        break
      }
      case KeyTypes.ORGANISATIONREGIONBUDGET: {
        keyData.organisationRegionBudget = key
        actionType = KeyActionTypes.ORGANISATIONREGIONBUDGET_SUCCESS
        break
      }
      case KeyTypes.ACTIVITIES: {
        keyData.activities = key
        actionType = KeyActionTypes.ACTIVITIES_SUCCESS
        break
      }
      case KeyTypes.ACTIVITY: {
        keyData.activity = key
        actionType = KeyActionTypes.ACTIVITY_SUCCESS
        break
      }
      case KeyTypes.ACTIVITYADDITIONAL: {
        keyData.activityAdditional = key
        actionType = KeyActionTypes.ACTIVITYADDITIONAL_SUCCESS
        break
      }
      default:
        break
    }

    console.log('dispatch: ', keyData, actionType)

    await dispatch(write({data: keyData})(actionType))
  }
}

export const newKey = () => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const keyData: KeyData = state.keys.data
    keyData.newKey = ethers.utils.formatBytes32String(shortid.generate())
    console.log('New Key from here! ', keyData.newKey)
    await dispatch(write({data: keyData})(KeyActionTypes.NEW_SUCCESS))
  }
}

export const setOrganisationsKey = (key: string) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const keyData: KeyData = state.keys.data
    keyData.organisations = key
    //console.log('Set Organisations Key! ', keyData.organisations)
    await dispatch(write({data: keyData})(KeyActionTypes.ORGANISATIONS_SUCCESS))
  }
}

export const setOrganisationKey = (key: string) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const keyData: KeyData = state.keys.data
    keyData.organisation = key
    //console.log('Set Organisation Key! ', keyData.organisations)
    await dispatch(write({data: keyData})(KeyActionTypes.ORGANISATION_SUCCESS))
  }
}

export const setActivitiesKey = (key: string) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const keyData: KeyData = state.keys.data
    keyData.activities = key
    //console.log('Set Organisations Key! ', keyData.organisations)
    await dispatch(write({data: keyData})(KeyActionTypes.ACTIVITIES_SUCCESS))
  }
}

export const setActivityKey = (key: string) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {
    const state = getState()
    const keyData: KeyData = state.keys.data
    keyData.activity = key
    //console.log('Set Organisation Key! ', keyData.organisations)
    await dispatch(write({data: keyData})(KeyActionTypes.ACTIVITY_SUCCESS))
  }
}
