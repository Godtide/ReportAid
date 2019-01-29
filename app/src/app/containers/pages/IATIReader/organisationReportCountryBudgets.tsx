import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getReportCountryBudgets } from '../../../store/IATI/IATIReader/organisationReportCountryBudgets/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrgReportCountryBudgetsData } from '../../../store/IATI/IATIReader/organisationReportCountryBudgets/types'

import { OrganisationReportCountryBudget as OrgReportCountryBudgetStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgReportCountryBudgetProps {
  num: number
  orgReportCountryBudgets: OrgReportCountryBudgetsData
}

interface OrgReportCountryBudgetDispatchProps {
  getReportCountryBudgets: () => void
}

type OrgReportCountryBudgetsReaderProps =  WithStyles<typeof styles> & OrgReportCountryBudgetProps & OrgReportCountryBudgetDispatchProps

export class OrgReportCountryBudgets extends React.Component<OrgReportCountryBudgetsReaderProps> {

  constructor (props: OrgReportCountryBudgetsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getReportCountryBudgets()
  }

  render() {

    const budgetsData = Object.keys(this.props.orgReportCountryBudgets)
    let xs = ""
    if ( budgetsData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      budgetsData.forEach((reportKey) => {
        xs += `**${OrgReportCountryBudgetStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgReportCountryBudgets[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrgReportCountryBudgetStrings.numReportBudgets}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((budgetKey) => {
          //console.log('Report: ', values[1][budgetKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][budgetKey].countryRef != "" ) {
            const countryRef = ethers.utils.parseBytes32String(values[1][budgetKey].countryRef)
            const budgetLine = ethers.utils.parseBytes32String(values[1][budgetKey].budgetLine)
            const start = ethers.utils.parseBytes32String(values[1][budgetKey].finance.start)
            const end = ethers.utils.parseBytes32String(values[1][budgetKey].finance.end)
            xs+= `**${OrgReportCountryBudgetStrings.reportingOrgRef}**: ${values[1][budgetKey].report.orgRef} <br />`
            xs+= `**${OrgReportCountryBudgetStrings.budgetReference}**: ${budgetKey} <br />`
            xs+= `**${OrgReportCountryBudgetStrings.countryReference}**: ${countryRef} <br />`
            xs+= `**${OrgReportCountryBudgetStrings.budgetLine}**: ${budgetLine} <br />`
            xs+= `**${OrgReportCountryBudgetStrings.value}**: ${values[1][budgetKey].finance.value} <br />`
            xs+= `**${OrgReportCountryBudgetStrings.status}**: ${values[1][budgetKey].finance.status} <br />`
            xs+= `**${OrgReportCountryBudgetStrings.budgetStart}**: ${start} <br />`
            xs+= `**${OrgReportCountryBudgetStrings.budgetEnd}**: ${end} <br /><br />`
          }
        })
        length += 1
        length == budgetsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }

    return (
      <div>
        <h2>{OrgReportCountryBudgetStrings.headingOrgReportCountryBudgetReader}</h2>
        <p>
          <b>{OrgReportCountryBudgetStrings.numReports}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrgReportCountryBudgetStrings.reportBudgetDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgReportCountryBudgetProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgReportCountryBudgetsReader.num,
    orgReportCountryBudgets: state.orgReportCountryBudgetsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgReportCountryBudgetDispatchProps => {
  return {
    getReportCountryBudgets: () => dispatch(getReportCountryBudgets())
  }
}

export const OrganisationReportCountryBudgets = withTheme(withStyles(styles)(connect<OrgReportCountryBudgetProps, OrgReportCountryBudgetDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgReportCountryBudgets)))
