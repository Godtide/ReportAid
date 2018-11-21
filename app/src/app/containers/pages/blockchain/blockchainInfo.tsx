import * as React from 'react'
import { connect } from 'react-redux'
import { ApplicationState } from '../../../store'

import { PlainTextKeyedWithTitleList } from '../../../components/io/plainText'
import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

import { Blockchain } from '../../../utils/strings'

interface InfoProps {
  propertiesList: object
}

class Info extends React.Component<WithStyles<typeof styles> & InfoProps> {

  render() {

    return (
      <PlainTextKeyedWithTitleList title={Blockchain.heading} list={this.props.propertiesList} />
    )
  }
}

const mapStateToProps = (state: ApplicationState): InfoProps => {
  const propertiesList = {
      API: state.blockchain.API,
      Network: state.blockchain.Name,
      ChainId: state.blockchain.ChainId,
      ENS: state.blockchain.ENS,
      Account: state.blockchain.account
  }
  const properties = {
    propertiesList: propertiesList
  }
  return properties
}

export const BlockchainInfo = withTheme(withStyles(styles)(connect<InfoProps, void, void, ApplicationState>(
  mapStateToProps
)(Info)))
