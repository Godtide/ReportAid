import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { getOverview } from '../../../store/IATI/IATIReader/organisationReader/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrgData } from '../../../store/IATI/IATIReader/organisationReader/types'

import { Organisation as OrgStrings } from '../../../utils/strings'

import { MarkdownText } from '../../../components/io/markdownText'
import { getDictEntries } from '../../../components/io/dict'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgProps {
  num: number
  orgs: OrgData
}

interface OrgDispatchProps {
  getOverview: () => void
}

type OrgReaderProps =  WithStyles<typeof styles> & OrgProps & OrgDispatchProps

export class OrgReader extends React.Component<OrgReaderProps> {

  constructor (props: OrgReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getOverview()
  }

  render() {

    const orgs = getDictEntries(this.props.orgs)

    return (
      <div>
        <h2>{OrgStrings.headingOrgReader}</h2>
        <p>
          <b>{OrgStrings.numOrgs}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrgStrings.orgDetails}</h3>
        <MarkdownText text={orgs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgReader.num,
    orgs: state.orgReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    getOverview: () => dispatch(getOverview())
  }
}

export const OrganisationReader = withTheme(withStyles(styles)(connect<OrgProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgReader)))
