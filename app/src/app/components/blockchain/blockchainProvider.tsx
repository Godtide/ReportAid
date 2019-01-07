import { ethers } from 'ethers'
import { Blockchain } from '../../utils/config'

export const getProvider = () => {

  const host = Blockchain.host
  const port = Blockchain.port
  const provider = host + ":" + port
  const network = Blockchain.network
  const blockchainProvider = new ethers.providers.JsonRpcProvider(provider, network)
  return blockchainProvider
}
