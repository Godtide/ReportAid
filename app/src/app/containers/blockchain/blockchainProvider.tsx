import * as React from 'react'
import { connect } from 'react-redux'

import Web3 from 'web3'
import { ethers } from 'ethers'

//import { fetchRequest, RequestDataAction } from '../../../store/helpers/about/actions'
import { BlockchainInfoProps } from '../../store/blockchain/types'
import { AddDataAction } from '../../store/blockchain/actions'
import { BlockchainStrings } from '../../utils/strings'

interface OwnProps {
  getProvider: boolean
}

interface PropsFromDispatch {
  addData: (props: BlockchainInfoProps) => any
}

type BlockchainProviderProps = OwnProps & PropsFromDispatch

export class BlockchainProvider extends React.Component<BlockchainProviderProps> {

  constructor (props: BlockchainProviderProps) {
    super(props)
    this._setProvider()
  }

  async _setProvider () {

    if(this.props.getProvider) {
      const ethereum = (window as any).ethereum
      let web3: any = (window as any).web3
      let networkName = ''
      let networkChainId = ''
      let networkENSAddress = ''
      let account = ''
      let blockchainProvider = undefined

      let providerData: BlockchainInfoProps = { APIProvider: '',
                                            networkName: '',
                                            networkChainId: '',
                                            networkENSAddress: '',
                                            account: ''
                                          }

      if (ethereum) {
        console.log('New MetaMask!')
        web3 = new Web3(ethereum)
        blockchainProvider = new ethers.providers.Web3Provider(web3.currentProvider)
        try {
            // Request account access if needed
            await ethereum.enable()
        } catch (error) {
          console.log(error)
        }
        await web3.eth.getAccounts((error: any, accounts: any) => {
            account = accounts[0]
        })
        .catch((error: any) => {
           console.log('Error!: ', error)
        })
        web3.eth.defaultAccount = account

      } else if (typeof web3 !== 'undefined') {
        console.log('In legacy web3 provider')
        blockchainProvider = new ethers.providers.Web3Provider(web3.currentProvider)
      } else {
        console.log('Running our own blockchain provider')
        const address = 'http://' + BlockchainStrings.host + ':' + BlockchainStrings.port
        web3 = new Web3(new Web3.providers.HttpProvider(address))
        blockchainProvider = new ethers.providers.Web3Provider(web3)
      }

      //console.log('Account: ', account)

      await blockchainProvider.getNetwork().then(function(chainObj: any) {
        console.log('Name: ', chainObj.name, ' ChainID: ', chainObj.chainId, 'ENS Address: ', chainObj.ensAddress)
        networkName = chainObj.name
        networkChainId = chainObj.chainId
        networkENSAddress = chainObj.ensAddress
      })

      providerData = { APIProvider: 'web3 ' + web3.version,
                       networkName: networkName,
                       networkChainId: networkChainId,
                       networkENSAddress: networkENSAddress,
                       account: account
                     }

      this.props.addData(providerData)
    }


    //console.log('Added data', providerData)
  }

}

const mapDispatchToProps = (dispatch: any, ownProps: BlockchainInfoProps): PropsFromDispatch => {
  console.log('dispatch')
  return {
    addData: (ownProps: OwnProps) => dispatch(AddDataAction(ownProps))
  }
}

export default connect(
  null,
  mapDispatchToProps
)(BlockchainProvider)
