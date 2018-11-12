export interface BlockchainProps {
  title: string
  APIProvider: string
  networkName: string
  networkChainId: string
  networkENSAddress: string
  account: string
}

export const enum BlockchainActionTypes {
  REQ_DATA = '@@blockchainTypes/REQ_DATA'
}
