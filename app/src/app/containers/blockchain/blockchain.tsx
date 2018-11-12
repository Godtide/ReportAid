import * as React from 'react'
import Web3 from 'web3'
import { ethers } from 'ethers'

interface BlockchainProps {
  readonly address: string
  readonly port: string
}

export class Blockchain extends React.Component<BlockchainProps> {

  address: string
  port: string
  web3: any
  blockchainProvider: any
  account: string

  constructor (props: BlockchainProps) {
    super(props)
    this.address = props.address
    this.port = props.port
    this.account = ''
    this._setBlockchainProvider()
    this._setAccount()
  }

  async _setBlockchainProvider () {

    const ethereum = (window as any).ethereum
    this.web3 = (window as any).web3
    if (ethereum) {
      console.log('New MetaMask!')
      this.web3 = new Web3(ethereum)
      this.blockchainProvider = new ethers.providers.Web3Provider(this.web3.currentProvider)
      try {
          // Request account access if needed
          await ethereum.enable()
      } catch (error) {
        console.log(error)
      }
    } else if (typeof this.web3 !== 'undefined') {
      console.log('In legacy web3 provider')
      this.blockchainProvider = new ethers.providers.Web3Provider(this.web3.currentProvider)
    } else {
      console.log('Running our own blockchain provider')
      const address = 'http://' + this.address + ':' + this.port
      this.web3 = new Web3(new Web3.providers.HttpProvider(address))
      this.blockchainProvider = new ethers.providers.Web3Provider(this.web3)
    }

    this.blockchainProvider.getNetwork().then(function(chainObj: any) {
      console.log('Name: ', chainObj.name, ' ChainID: ', chainObj.chainId, 'ENS Address: ', chainObj.ensAddress)
    })
  }

  // metamask sets its account to web3.eth.accounts[0]
  async _setAccount () {
    if (this.web3.eth.accounts[0] !== this.account) {
      console.log('Setting Account')
      await this.web3.eth.getAccounts((error: any, accounts: any) => {
        this.account = accounts[0]
      })
      this.web3.eth.defaultAccount = this.account
      console.log('Account: ', this.account)
    }
  }

  getProvider () {
    return this.blockchainProvider
  }
}
