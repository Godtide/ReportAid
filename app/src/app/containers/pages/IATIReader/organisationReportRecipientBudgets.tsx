import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getReportRecipientBudgets } from '../../../store/IATI/IATIReader/organisationReportRecipientBudgets/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrgReportRecipientBudgetsData } from '../../../store/IATI/IATIReader/organisationReportRecipientBudgets/types'

import { OrganisationReportRecipientBudget as OrgReportRecipientBudgetStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgReportRecipientBudgetProps {
  num: number
  orgReportRecipientBudgets: OrgReportRecipientBudgetsData
}

interface OrgReportRecipientBudgetDispatchProps {
  getReportRecipientBudgets: () => void
}

type OrgReportRecipientBudgetsReaderProps =  WithStyles<typeof styles> & OrgReportRecipientBudgetProps & OrgReportRecipientBudgetDispatchProps

export class OrgReportRecipientBudgets extends React.Component<OrgReportRecipientBudgetsReaderProps> {

  constructor (props: OrgReportRecipientBudgetsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getReportRecipientBudgets()
  }

  render() {

    const budgetsData = Object.keys(this.props.orgReportRecipientBudgets)
    let xs = ""
    if ( budgetsData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      budgetsData.forEach((reportKey) => {
        xs += `**${OrgReportRecipientBudgetStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgReportRecipientBudgets[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrgReportRecipientBudgetStrings.numReportBudgets}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((budgetKey) => {
          //console.log('Report: ', values[1][budgetKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][budgetKey].budgetLine != "" ) {
            const budgetLine = ethers.utils.parseBytes32String(values[1][budgetKey].budgetLine)
            const start = ethers.utils.parseBytes32String(values[1][budgetKey].finance.start)
            const end = ethers.utils.parseBytes32String(values[1][budgetKey].finance.end)
            xs+= `**${OrgReportRecipientBudgetStrings.budgetReference}**: ${budgetKey} <br />`
            xs+= `**${OrgReportRecipientBudgetStrings.orgReference}**: ${values[1][budgetKey].orgRef} <br />`
            xs+= `**${OrgReportRecipientBudgetStrings.budgetLine}**: ${budgetLine} <br />`
            xs+= `**${OrgReportRecipientBudgetStrings.value}**: ${values[1][budgetKey].finance.value} <br />`
            xs+= `**${OrgReportRecipientBudgetStrings.status}**: ${values[1][budgetKey].finance.status} <br />`
            xs+= `**${OrgReportRecipientBudgetStrings.budgetStart}**: ${start} <br />`
            xs+= `**${OrgReportRecipientBudgetStrings.budgetEnd}**: ${end} <br /><br />`
          }
        })
        length += 1
        length == budgetsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }

    return (
      <div>
        <h2>{OrgReportRecipientBudgetStrings.headingOrgReportRecipientBudgetReader}</h2>
        <p>
          <b>{OrgReportRecipientBudgetStrings.numReports}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrgReportRecipientBudgetStrings.reportBudgetDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgReportRecipientBudgetProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgReportRecipientBudgetsReader.num,
    orgReportRecipientBudgets: state.orgReportRecipientBudgetsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgReportRecipientBudgetDispatchProps => {
  return {
    getReportRecipientBudgets: () => dispatch(getReportRecipientBudgets())
  }
}

export const OrganisationReportRecipientBudgets = withTheme(withStyles(styles)(connect<OrgReportRecipientBudgetProps, OrgReportRecipientBudgetDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgReportRecipientBudgets)))
