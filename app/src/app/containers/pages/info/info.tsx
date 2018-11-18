import * as React from 'react'
//import { bindActionCreators, Dispatch, AnyAction } from 'redux'
import { connect } from 'react-redux'
import { ApplicationState } from '../../../store'
import { MarkdownText } from '../../../components/io/markdownText'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

//import { fetchRequest, RequestDataAction } from '../../../store/helpers/about/actions'
import { InfoProps } from '../../../store/info/types'
import { InfoTypes } from './types'

interface StateProps {
  type: InfoTypes
}

type AllProps = InfoProps & StateProps

class AppInfo extends React.Component<WithStyles<typeof styles> & AllProps> {

  render() {

    return (
      <div>
        <h2>{this.props.title}</h2>
        <MarkdownText text={this.props.data} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState, ownProps: StateProps): InfoProps => {
  switch (ownProps.type) {
    case InfoTypes.HOME:
      return { title: state.home.title, data: state.home.data }
    case InfoTypes.ABOUT:
      return { title: state.about.title, data: state.about.data }
    case InfoTypes.OVERVIEW:
      return { title: state.overview.title, data: state.overview.data }
    case InfoTypes.HELP:
      return { title: state.help.title, data: state.help.data }
    default:
      return { title: state.home.title, data: state.home.data }
  }
}

export const Info = withTheme(withStyles(styles)(connect<InfoProps, {}, StateProps, ApplicationState>(
  mapStateToProps
)(AppInfo)))
