import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'
import Markdown from 'react-markdown'

import { getOrgReports } from '../../../store/IATI/IATIReader/organisationReportReader/actions'

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

    let xs = ""
    Object.keys(this.props.orgReports).forEach((key) => {
      xs += `**Organisation Key**: ${key}, `
      const values = Object.values(this.props.orgReports[key])
      //console.log('Values: ', values)
      xs += `**Num reports**: ${values[0]} <br /><br />`
      Object.keys(values[1]).forEach((thisKey) => {
        //console.log('Blah', values[1][thisKey])
        xs+= `**Report Key**: ${thisKey} <br />`
        xs+= `**Report Version**:  ${values[1][thisKey].version} <br />`
        xs+= `**Report Organisation Reference**:  ${values[1][thisKey].orgRef} <br />`
        xs+= `**Report Reference**:  ${values[1][thisKey].reportRef} <br />`
        xs+= `**Report Reporting Organisation Reference**:  ${values[1][thisKey].reportingOrg.orgRef} <br />`
        xs+= `**Report Reporting Organisation Type**:  ${values[1][thisKey].reportingOrg.orgType} <br />`
        xs+= `**Report Reporting Organisation is Secondary?**:  ${values[1][thisKey].reportingOrg.isSecondary} <br />`
        xs+= `**Report Language**:  ${values[1][thisKey].lang} <br />`
        xs+= `**Report Currency**:  ${values[1][thisKey].currency} <br />`
        xs+= `**Report Last Updated**:  ${values[1][thisKey].lastUpdatedTime} <br /><br />`
      })
    })
    //console.log('Org Reports: ', orgReports)

    return (
      <div>
        <h2>{OrgReportStrings.headingOrgReportReader}</h2>
        <p>
          <b>{OrgReportStrings.numOrgReports}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrgReportStrings.orgReportDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgReportProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgReportsReader.num,
    orgReports: state.orgReportsReader.data
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
