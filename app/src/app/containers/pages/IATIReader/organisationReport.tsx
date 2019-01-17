import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
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

    const orgsData = Object.keys(this.props.orgReports)
    let xs = ""
    if ( orgsData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      orgsData.forEach((key) => {
        xs += `**${OrgReportStrings.orgIdentifier}**: ${key}<br />`
        const values = Object.values(this.props.orgReports[key])
        //console.log('Values: ', values)
        xs += `**${OrgReportStrings.numOrgReports}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((thisKey) => {
          //console.log('Report: ', values[1][thisKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][thisKey].Report.version != "" ) {
            const version = ethers.utils.parseBytes32String(values[1][thisKey].Report.version)
            const language =  ethers.utils.parseBytes32String(values[1][thisKey].Report.lang)
            const currency =  ethers.utils.parseBytes32String(values[1][thisKey].Report.currency)
            const lastUpdated =  ethers.utils.parseBytes32String(values[1][thisKey].Report.lastUpdatedTime)
            xs+= `**${OrgReportStrings.reportKey}**: ${thisKey} <br />`
            xs+= `**${OrgReportStrings.version}**:  ${version} <br />`
            xs+= `**${OrgReportStrings.reportingOrgRef}**:  ${values[1][thisKey].Report.reportingOrg.orgRef} <br />`
            xs+= `**${OrgReportStrings.reportingOrgType}**:  ${values[1][thisKey].Report.reportingOrg.orgType} <br />`
            xs+= `**${OrgReportStrings.reportingOrgIsSecondary}**:  ${values[1][thisKey].Report.reportingOrg.isSecondary} <br />`
            xs+= `**${OrgReportStrings.language}**:  ${language} <br />`
            xs+= `**${OrgReportStrings.currency}**:  ${currency} <br />`
            xs+= `**${OrgReportStrings.lastUpdated}**:  ${lastUpdated} <br /><br />`
          }
        })
        length += 1
        length == orgsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }

    return (
      <div>
        <h2>{OrgReportStrings.headingOrgReportReader}</h2>
        <p>
          <b>{OrgReportStrings.numOrganisations}</b>: {this.props.num}
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
