import { ThunkDispatch } from 'redux-thunk'

import shortid from 'shortid'
import { ethers } from 'ethers'

import { ApplicationState } from '../../../../store'

import { write } from '../../actions'

import { ActionProps, TxReport } from '../../../../types'
import { ActivitySectorProps, IATIActivitySectorProps, IATIWriterActionTypes } from '../../../types'

import { Transaction } from '../../../../../utils/strings'

export const setActivitySector = (details: ActivitySectorProps) => {
  return async (dispatch: ThunkDispatch<ApplicationState, null, ActionProps>, getState: Function) => {

    const state = getState()
    const activitySectorsContract = state.chainContracts.data.contracts.activitySectors

    let sectorRef = details.sectorRef
    if ( sectorRef == "" ) {
      sectorRef = ethers.utils.formatBytes32String(shortid.generate())
    }

    const sectorData: IATIActivitySectorProps = {
      percentage: details.percentage,
      dacCode: details.dacCode
    }

    let actionType = IATIWriterActionTypes.ACTIVITYSECTOR_FAILURE
    let txData: TxReport = {}
    try {
      const tx = await activitySectorsContract.setParticipatingOrg(details.activitiesRef, details.activityRef, sectorRef, sectorData)
      const key = tx.hash
      txData = {
        [key]: {
          summary: `${Transaction.success}`,
          info: tx
        }
      }

      actionType = IATIWriterActionTypes.ACTIVITYSECTOR_SUCCESS
    } catch (error) {
      txData = {
        [-1]: {
          summary: `${Transaction.fail}`,
          info: {}
        }
      }
      console.log('setActivitySector error', error)
    }

    dispatch(write({data: {data: txData}})(actionType))
  }
}
