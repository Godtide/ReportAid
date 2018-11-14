export interface BlockchainProps {
  APIProvider: string
  networkName: string
  networkChainId: string
  networkENSAddress: string
  account: string
}

export const enum BlockchainActionTypes {
  ADD_DATA = '@@blockchainActionTypes/ADD_DATA'
}
