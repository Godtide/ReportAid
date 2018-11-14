import * as React from 'react'
import { connect } from 'react-redux'
import { ApplicationState } from '../../../store'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

import { BlockchainStrings } from '../../../utils/strings'
import { BlockchainProvider } from '../../../containers/blockchain/blockchainProvider'

//import { fetchRequest, RequestDataAction } from '../../../store/helpers/about/actions'
import { BlockchainInfoProps } from '../../../store/blockchain/types'

class BlockchainInfo extends React.Component<WithStyles<typeof styles> & BlockchainInfoProps> {

  constructor (props: any) {
    super(props)
    const providerProps = {getProvider: true}
    const blockchainProvider = new BlockchainProvider(providerProps)
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
      </div>

    )
  }
}

const mapStateToProps = (state: ApplicationState): BlockchainInfoProps => {
  return {
    APIProvider: state.blockchain.APIProvider,
    networkName: state.blockchain.networkName,
    networkChainId: state.blockchain.networkChainId,
    networkENSAddress: state.blockchain.networkENSAddress,
    account: state.blockchain.account
  }
}

export default withTheme(withStyles(styles)(connect(
  mapStateToProps
)(BlockchainInfo)))
