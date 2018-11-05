import * as React from 'react'
import Web3 from 'web3'
import { ethers } from 'ethers'

interface BlockchainProps {
  readonly address: string
}

export class Blockchain extends React.Component<BlockchainProps> {

  address: string
  blockchainProvider: any

  constructor (props: BlockchainProps) {
    super(props)
    this.address = props.address
    this.setProvider()
  }

  async setProvider () {
    const ethereum = (window as any).ethereum
    let web3 = (window as any).web3
    if (ethereum) {
      web3 = new Web3(ethereum)
      this.blockchainProvider = new ethers.providers.Web3Provider(web3.currentProvider)
      try {
          // Request account access if needed
          await ethereum.enable()
      } catch (error) {
        console.log(error)
      }
    } else if (typeof web3 !== 'undefined') {
     this.blockchainProvider = new ethers.providers.Web3Provider(web3.currentProvider)
    } else {
        // Set the provider you want from Web3.providers
        web3 = new Web3(new Web3.providers.HttpProvider(this.address))
        this.blockchainProvider = new ethers.providers.Web3Provider(web3)
    }
  }

  getProvider () {
    return this.blockchainProvider
  }
}

/*
if (window.ethereum) {
        window.web3 = new Web3(ethereum);
        try {
            // Request account access if needed
            await ethereum.enable();
            // Acccounts now exposed
            web3.eth.sendTransaction({});
        } catch (error) {
            // User denied account access...
        }
    }
    // Legacy dapp browsers...
    else if (window.web3) {
        window.web3 = new Web3(web3.currentProvider);
        // Acccounts always exposed
        web3.eth.sendTransaction({});
    }
    // Non-dapp browsers...
    else {
        console.log('Non-Ethereum browser detected. You should consider trying MetaMask!');
    }
*/
