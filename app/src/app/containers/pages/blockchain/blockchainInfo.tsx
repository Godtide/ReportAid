import * as React from 'react'
import { connect } from 'react-redux'
import { ApplicationState } from '../../../store'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

import { BlockchainStrings } from '../../../utils/strings'

import { BlockchainProps } from '../../../store/blockchain/types'

class Info extends React.Component<WithStyles<typeof styles> & BlockchainProps> {

  constructor (props: any) {
    super(props)
  }

  render() {

    return (
      <div>
        <h2>{BlockchainStrings.heading}</h2>
        <p><b>{BlockchainStrings.APIName}</b>{this.props.APIName}</p>
        <p><b>{BlockchainStrings.networkName}</b>{this.props.networkName}</p>
        <p><b>{BlockchainStrings.networkChainId}</b>{this.props.networkChainId}</p>
        <p><b>{BlockchainStrings.ENSAddress}</b>{this.props.networkENSAddress}</p>
        <p><b>{BlockchainStrings.networkAccount}</b>{this.props.account}</p>
      </div>

    )
  }
}

const mapStateToProps = (state: ApplicationState): BlockchainProps => {
  return {
    APIName: state.blockchain.APIName,
    networkName: state.blockchain.networkName,
    networkChainId: state.blockchain.networkChainId,
    networkENSAddress: state.blockchain.networkENSAddress,
    account: state.blockchain.account,
    provider: state.blockchain.provider
  }
}

export const BlockchainInfo = withTheme(withStyles(styles)(connect<BlockchainProps, void, void, ApplicationState>(
  mapStateToProps
)(Info)))
