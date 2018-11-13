import * as React from 'react'
import { connect } from 'react-redux'
import { ApplicationState } from '../../../store'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

import { BlockchainStrings } from '../../../utils/strings'
import { AppButton } from '../../../components/io/appButton'

import Web3 from 'web3'
import { ethers } from 'ethers'

//import { fetchRequest, RequestDataAction } from '../../../store/helpers/about/actions'
import { BlockchainProps } from '../../../store/blockchain/types'
import { addData } from '../../../store/blockchain/actions'

interface PropsFromDispatch {
  addData: (providerData: BlockchainProps) => void
}

export type AllProps = BlockchainProps & PropsFromDispatch

class BlockchainInfo extends React.Component<WithStyles<typeof styles> & AllProps> {

  providerData: BlockchainProps = { APIProvider: {},
                                    networkName: '',
                                    networkChainId: '',
                                    networkENSAddress: '',
                                    account: ''
                                  }

  addToStore() {
    console.log(this.props.addData)
    this.props.addData(this.providerData)
  }

  async _setProvider() {

    const ethereum = (window as any).ethereum
    let web3: any = (window as any).web3
    let networkName = ''
    let networkChainId = ''
    let networkENSAddress = ''
    let account = ''
    let blockchainProvider = undefined

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

    console.log('Account: ', account)

    await blockchainProvider.getNetwork().then(function(chainObj: any) {
      console.log('Name: ', chainObj.name, ' ChainID: ', chainObj.chainId, 'ENS Address: ', chainObj.ensAddress)
      networkName = chainObj.name
      networkChainId = chainObj.chainId
      networkENSAddress = chainObj.ensAddress
    })

    this.providerData = { APIProvider: web3,
                          networkName: networkName,
                          networkChainId: networkChainId,
                          networkENSAddress: networkENSAddress,
                          account: account
                        }
    //console.log(this.providerData)
  }

  handleSubmit(event: any) {
    const setProvider = this._setProvider.bind(this)
    event.preventDefault()
    setProvider()
  }

  render() {

    //const appButtonProps = { label: 'Get Blockchain', tip: 'Sets the blockchain', submit: this.handleSubmit }

    return (
      <div>
        <h2>{BlockchainStrings.heading}</h2>
        <p><b>{BlockchainStrings.APIProvider}</b>{this.props.APIProvider}</p>
        <p><b>{BlockchainStrings.networkName}</b>{this.props.networkName}</p>
        <p><b>{BlockchainStrings.networkChainId}</b>{this.props.networkChainId}</p>
        <p><b>{BlockchainStrings.ENSAddress}</b>{this.props.networkENSAddress}</p>
        <p><b>{BlockchainStrings.networkAccount}</b>{this.props.account}</p>
        <AppButton
          label={BlockchainStrings.setButton}
          tip={BlockchainStrings.setButtonTip}
          submit={this.handleSubmit.bind(this)}
        />
        <AppButton
          label='fek'
          tip='fek'
          submit={this.addToStore.bind(this)}
        />
      </div>

    )
  }
}

const mapStateToProps = (state: ApplicationState): BlockchainProps => {
  return {
    APIProvider: state.blockchain.APIProvider,
    networkName: state.blockchain.networkName,
    networkChainId: state.blockchain.networkChainId,
    networkENSAddress: state.blockchain.networkENSAddress,
    account: state.blockchain.account
  }
}

const mapDispatchToProps = (dispatch: any, ownProps: BlockchainProps): PropsFromDispatch => {
  console.log('mapDispatch')
  return {
    addData: dispatch(addData(ownProps))
  }
}

export default withTheme(withStyles(styles)(connect(
  mapStateToProps,
  mapDispatchToProps
)(BlockchainInfo)))
