import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getReportBudgets } from '../../../store/IATI/IATIReader/organisationReportBudgets/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrgReportBudgetsData } from '../../../store/IATI/IATIReader/organisationReportBudgets/types'

import { OrganisationReportBudget as OrgReportBudgetStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgReportBudgetProps {
  num: number
  orgReportBudgets: OrgReportBudgetsData
}

interface OrgReportBudgetDispatchProps {
  getReportBudgets: () => void
}

type OrgReportBudgetsReaderProps =  WithStyles<typeof styles> & OrgReportBudgetProps & OrgReportBudgetDispatchProps

export class OrgReportBudgets extends React.Component<OrgReportBudgetsReaderProps> {

  constructor (props: OrgReportBudgetsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getReportBudgets()
  }

  render() {

    const budgetsData = Object.keys(this.props.orgReportBudgets)
    let xs = ""
    if ( budgetsData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      budgetsData.forEach((reportKey) => {
        xs += `**${OrgReportBudgetStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgReportBudgets[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrgReportBudgetStrings.numReportBudgets}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((budgetKey) => {
          //console.log('Report: ', values[1][budgetKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][budgetKey].budgetLine != "" ) {
            const budgetLine = ethers.utils.parseBytes32String(values[1][budgetKey].budgetLine)
            const status = ethers.utils.parseBytes32String(values[1][budgetKey].finance.status)
            const start = ethers.utils.parseBytes32String(values[1][budgetKey].finance.start)
            const end = ethers.utils.parseBytes32String(values[1][budgetKey].finance.end)
            xs+= `**${OrgReportBudgetStrings.budgetReference}**: ${budgetKey} <br />`
            xs+= `**${OrgReportBudgetStrings.budgetLine}**: ${budgetLine} <br />`
            xs+= `**${OrgReportBudgetStrings.value}**: ${values[1][budgetKey].finance.value} <br />`
            xs+= `**${OrgReportBudgetStrings.status}**: ${status} <br />`
            xs+= `**${OrgReportBudgetStrings.budgetStart}**: ${start} <br />`
            xs+= `**${OrgReportBudgetStrings.budgetEnd}**: ${end} <br /><br />`
          }
        })
        length += 1
        length == budgetsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }

    return (
      <div>
        <h2>{OrgReportBudgetStrings.headingOrgReportBudgetReader}</h2>
        <p>
          <b>{OrgReportBudgetStrings.numReports}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrgReportBudgetStrings.reportBudgetDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgReportBudgetProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgReportBudgetsReader.num,
    orgReportBudgets: state.orgReportBudgetsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgReportBudgetDispatchProps => {
  return {
    getReportBudgets: () => dispatch(getReportBudgets())
  }
}

export const OrganisationReportBudgets = withTheme(withStyles(styles)(connect<OrgReportBudgetProps, OrgReportBudgetDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgReportBudgets)))
