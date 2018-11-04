import * as React from 'react'
import Web3 from 'web3'
import { ethers } from 'ethers'

interface BlockchainProps {
  readonly address: string
}

export class Blockchain extends React.Component<BlockchainProps> {

  currentProvider: any
  blockchainProvider: any

  constructor (props: BlockchainProps) {
    super(props)
    this.currentProvider = new Web3.providers.HttpProvider(this.props.address)
    this.blockchainProvider = new ethers.providers.Web3Provider(this.currentProvider)
  }

  getCurrentProvider () {
    return this.blockchainProvider
  }

  getBlockchainProvider () {
    return this.blockchainProvider
  }
}
