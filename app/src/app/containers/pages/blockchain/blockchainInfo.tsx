import * as React from 'react'
import { connect } from 'react-redux'
import { ApplicationState } from '../../../store'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

import { BlockchainStrings } from '../../../utils/strings'

//import { fetchRequest, RequestDataAction } from '../../../store/helpers/about/actions'
import { BlockchainInfoProps } from '../../../store/blockchain/types'

class BlockchainInfo extends React.Component<WithStyles<typeof styles> & BlockchainInfoProps> {

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

const mapStateToProps = (state: ApplicationState): BlockchainInfoProps => {
  return {
    APIName: state.blockchain.APIName,
    networkName: state.blockchain.networkName,
    networkChainId: state.blockchain.networkChainId,
    networkENSAddress: state.blockchain.networkENSAddress,
    account: state.blockchain.account
  }
}

/* export default withTheme(withStyles(styles)(connect<BlockchainInfoProps,void,{}>(
  mapStateToProps
)(BlockchainInfo)))*/

export default withTheme(withStyles(styles)(connect<BlockchainInfoProps, void, void, ApplicationState>(
  mapStateToProps
)(BlockchainInfo)))
