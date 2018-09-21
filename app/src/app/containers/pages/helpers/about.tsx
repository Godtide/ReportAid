import * as React from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { ApplicationState } from '../../../store'
import { AboutProps } from '../../../store/helpers/about/types'
import MarkdownText from '../../../containers/io/markdownText'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

import getData from '../../../store/helpers/about/actions'


interface AboutPropsFromState {
  about: AboutProps
}

interface AboutPropsFromDispatch {
  getData: typeof getData
}

// Combine both state + dispatch props - as well as any props we want to pass - in a union type.
type AllProps = AboutPropsFromState & AboutPropsFromDispatch

class About extends React.Component<WithStyles<typeof styles> & AllProps> {

  componentDidMount() {
    // only fetch the data if there is no data
    if (!this.props.about) this.props.getData
    console.log(this.props.about)
  }

  render() {
    const about  = this.props.about
    console.log(about)

    return (
      <div>
        <h2>{about.title}</h2>
        <MarkdownText text={about.data} />
      </div>
    )
  }
}

/*
<div>
  <h2>{data.title}</h2>
  <MarkdownText text={data.text} />
</div>
*/

const mapStateToProps = (state: ApplicationState, ownProps: AboutProps): ApplicationState => {
  return {
    about: state.about
  }
}

const mapDispatchToProps = (dispatch: any, ownProps: AboutProps) => {
  return bindActionCreators({
    getData: getData
  }, dispatch)
}

export default withTheme(withStyles(styles)(connect(
  mapStateToProps,
  mapDispatchToProps
)(About)))
