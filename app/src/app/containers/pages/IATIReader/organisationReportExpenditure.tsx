import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getReportExpenditure } from '../../../store/IATI/IATIReader/organisationReportExpenditure/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrgReportExpenditureData } from '../../../store/IATI/IATIReader/organisationReportExpenditure/types'

import { OrganisationReportExpenditure as OrgReportExpenditureStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgReportExpenditureProps {
  num: number
  orgReportExpenditure: OrgReportExpenditureData
}

interface OrgReportExpenditureDispatchProps {
  getReportExpenditure: () => void
}

type OrgReportExpenditureReaderProps =  WithStyles<typeof styles> & OrgReportExpenditureProps & OrgReportExpenditureDispatchProps

export class OrgReportExpenditure extends React.Component<OrgReportExpenditureReaderProps> {

  constructor (props: OrgReportExpenditureReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getReportExpenditure()
  }

  render() {

    const expenditureData = Object.keys(this.props.orgReportExpenditure)
    let xs = ""
    if ( expenditureData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      expenditureData.forEach((reportKey) => {
        xs += `**${OrgReportExpenditureStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgReportExpenditure[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrgReportExpenditureStrings.numReportExpenditure}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((expenditureKey) => {
          //console.log('Report: ', values[1][expenditureKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][expenditureKey].expenditureLine != "" ) {
            const expenditureLine = ethers.utils.parseBytes32String(values[1][expenditureKey].expenditureLine)
            const start = ethers.utils.parseBytes32String(values[1][expenditureKey].finance.start)
            const end = ethers.utils.parseBytes32String(values[1][expenditureKey].finance.end)
            xs+= `**${OrgReportExpenditureStrings.reportingOrgRef}**: ${values[1][expenditureKey].report.orgRef} <br />`
            xs+= `**${OrgReportExpenditureStrings.expenditureReference}**: ${expenditureKey} <br />`
            xs+= `**${OrgReportExpenditureStrings.expenditureLine}**: ${expenditureLine} <br />`
            xs+= `**${OrgReportExpenditureStrings.value}**: ${values[1][expenditureKey].finance.value} <br />`
            xs+= `**${OrgReportExpenditureStrings.status}**: ${values[1][expenditureKey].finance.status} <br />`
            xs+= `**${OrgReportExpenditureStrings.expenditureStart}**: ${start} <br />`
            xs+= `**${OrgReportExpenditureStrings.expenditureEnd}**: ${end} <br /><br />`
          }
        })
        length += 1
        length == expenditureData.length ? xs += "" : xs += "---<br /><br />"
      })
    }

    return (
      <div>
        <h2>{OrgReportExpenditureStrings.headingOrgReportExpenditureReader}</h2>
        <p>
          <b>{OrgReportExpenditureStrings.numReports}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrgReportExpenditureStrings.reportExpenditureDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgReportExpenditureProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgReportExpenditureReader.num,
    orgReportExpenditure: state.orgReportExpenditureReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgReportExpenditureDispatchProps => {
  return {
    getReportExpenditure: () => dispatch(getReportExpenditure())
  }
}

export const OrganisationReportExpenditure = withTheme(withStyles(styles)(connect<OrgReportExpenditureProps, OrgReportExpenditureDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgReportExpenditure)))
