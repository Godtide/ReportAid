import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../../../../actions'

import { ActionProps, TxReport } from '../../../../types'
import { ActivityTerritoryProps, IATIActivityTerritoryProps, IATIWriterActionTypes } from '../../../types'

import { Transaction } from '../../../../../utils/strings'

export const setActivityTerritory = (details: ActivityTerritoryProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activityTerritoriesContract = state.chainContracts.data.contracts.activityTerritories

    let territoryRef = details.territoryRef
    if ( territoryRef == "" ) {
      territoryRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    // one or the other, not both
    let territory = details.countryCode
    if ( details.regionCode != "" ) {
      territory = details.regionCode
    }

    const territoryData: IATIActivityTerritoryProps = {
      percentage: details.percentage,
      territory: ethers.utils.formatBytes32String(territory),
    }

    let actionType = IATIWriterActionTypes.ACTIVITYTERRITORY_FAILURE
    let txData: TxReport = {}
    try {
      const tx = await activityTerritoriesContract.setTerritory(details.activitiesRef, details.activityRef, territoryRef, territoryData)
      const key = tx.hash
      txData = {
        [key]: {
          summary: `${Transaction.success}`,
          info: tx
        }
      }

      actionType = IATIWriterActionTypes.ACTIVITYTERRITORY_SUCCESS
    } catch (error) {
      txData = {
        [-1]: {
          summary: `${Transaction.fail}`,
          info: {}
        }
      }
      console.log('setActivityTerritory error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
