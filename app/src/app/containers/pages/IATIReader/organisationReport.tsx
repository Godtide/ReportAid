import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'
import Markdown from 'react-markdown'

import { getOrgReports } from '../../../store/IATI/IATIReader/organisationReader/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrgReportData } from '../../../store/IATI/IATIReader/organisationReportReader/types'

import { OrganisationReport as OrgReportStrings } from '../../../utils/strings'

import { getDictEntries } from '../../../components/io/dict'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgReportProps {
  num: number
  orgReports: OrgReportData
}

interface OrgReportDispatchProps {
  getOrgReports: () => void
}

type OrgReportReaderProps =  WithStyles<typeof styles> & OrgReportProps & OrgReportDispatchProps

export class OrgReportReader extends React.Component<OrgReportReaderProps> {

  constructor (props: OrgReportReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getOrgReports()
  }

  render() {

    const orgs = getDictEntries(this.props.orgReports)

    return (
      <div>
        <h2>{OrgReportStrings.headingOrgReportReader}</h2>
        <p>
          <b>{OrgReportStrings.numOrgReports}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrgReportStrings.orgDetails}</h3>
        <Markdown escapeHtml={false} source={orgs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgReportProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgReader.num,
    orgs: state.orgReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgReportDispatchProps => {
  return {
    getOrgReports: () => dispatch(getOrgReports())
  }
}

export const OrganisationReportReader = withTheme(withStyles(styles)(connect<OrgReportProps, OrgReportDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgReportReader)))
