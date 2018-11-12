import * as React from 'react'
//import { bindActionCreators, Dispatch, AnyAction } from 'redux'
import { connect } from 'react-redux'
import { ApplicationState } from '../../../store'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

import { AppStrings } from '../../../utils/strings'
import { AppButton } from '../../../components/io/appButton'

import { ethers } from 'ethers'

//import { fetchRequest, RequestDataAction } from '../../../store/helpers/about/actions'
import { GetWeb3, GetProvider, SetAccount } from '../../../components/blockchain/blockchain'
import { BlockchainProps } from '../../../store/blockchain/types'

class BlockchainInfo extends React.Component<WithStyles<typeof styles> & BlockchainProps> {

  handleSubmit(event: any): void {
    console.log('boom!')
    event.preventDefault()
    const web3Props = {address: AppStrings.blockchainHost, port:  AppStrings.blockchainPort}
    const web3 = GetWeb3(web3Props)
    const providerProps = {web3: web3}
    const provider = GetProvider(providerProps)
    const account = SetAccount(providerProps)
  }

  render() {

    console.log('here')

    //const appButtonProps = { label: 'Get Blockchain', tip: 'Sets the blockchain', submit: this.handleSubmit }

    return (
      <div>
        <h2>{this.props.title}</h2>
        <p>{this.props.APIProvider}</p>
        <p>{this.props.networkName}</p>
        <p>{this.props.networkChainId}</p>
        <p>{this.props.networkENSAddress}</p>
        <p>{this.props.account}</p>
        <AppButton
          label='Get Blockchain'
          tip='Sets the blockchain'
          submit={this.handleSubmit}
        />
      </div>

    )
  }
}

const mapStateToProps = (state: ApplicationState): BlockchainProps => {
  return {
    title: state.blockchain.title,
    APIProvider: state.blockchain.APIProvider,
    networkName: state.blockchain.networkName,
    networkChainId: state.blockchain.networkChainId,
    networkENSAddress: state.blockchain.networkENSAddress,
    account: state.blockchain.account
  }
}

/* const mapDispatchToProps = (dispatch: Dispatch<AnyAction>, ownProps: AboutProps) => {
  console.log('bollox', ownProps.title, ownProps.data)
  const type = AboutActionTypes.REQ_DATA
  const payload = {
    title: ownProps.title,
    data: ownProps.data
  }
  return bindActionCreators<AboutRequestDataAction, any>({
    onSomeEvent: AboutFetchRequest(type, payload)
  },dispatch)
} */

export default withTheme(withStyles(styles)(connect(
  mapStateToProps
)(BlockchainInfo)))
