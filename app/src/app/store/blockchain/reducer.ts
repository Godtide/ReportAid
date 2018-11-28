import { BlockchainActionTypes, BlockchainProps, InfoProps, AccountProps, ObjectProps, OrgContractProps } from './types'
import { ActionProps, PayloadProps } from '../types'

const initialState: BlockchainProps = {
  data: {
    API: '',
    Name: '',
    ChainId: '',
    ENS: '',
    account: '',
    provider: {},
    orgContract: {}
  }
}

export const reducer = (state: BlockchainProps = initialState, action: ActionProps): BlockchainProps => {
  console.log('blockchain info: ', action.type, action.payload)
  switch (action.type) {
    case BlockchainActionTypes.ADD_INFO: {
      const infoData = action.payload.data as InfoProps
      let data: BlockchainProps = {
        API: infoData.API,
        Name: infoData.Name,
        ChainId: infoData.ChainId,
        ENS: infoData.ENS,
        provider: infoData.provider
      }
      return {...state, data: data }
    }
    case BlockchainActionTypes.ADD_ACCOUNT: {
      const accountData = action.payload.data as AccountProps
      let data: BlockchainProps = {
        account: accountData.account
      }
      return {...state, data: data }
    }
    case BlockchainActionTypes.ADD_ORGCONTRACT: {
      const orgContractData = action.payload.data as OrgContractProps
      let data: BlockchainProps = {
        contract: orgContractData.contract
      }
      return {...state, data: data }
    }
    default:
      return state
  }
}
