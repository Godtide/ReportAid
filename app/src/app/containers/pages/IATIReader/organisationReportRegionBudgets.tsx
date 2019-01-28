import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getReportRegionBudgets } from '../../../store/IATI/IATIReader/organisationReportRegionBudgets/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrgReportRegionBudgetsData } from '../../../store/IATI/IATIReader/organisationReportRegionBudgets/types'

import { OrganisationReportRegionBudget as OrgReportRegionBudgetStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgReportRegionBudgetProps {
  num: number
  orgReportRegionBudgets: OrgReportRegionBudgetsData
}

interface OrgReportRegionBudgetDispatchProps {
  getReportRegionBudgets: () => void
}

type OrgReportRegionBudgetsReaderProps =  WithStyles<typeof styles> & OrgReportRegionBudgetProps & OrgReportRegionBudgetDispatchProps

export class OrgReportRegionBudgets extends React.Component<OrgReportRegionBudgetsReaderProps> {

  constructor (props: OrgReportRegionBudgetsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getReportRegionBudgets()
  }

  render() {

    const budgetsData = Object.keys(this.props.orgReportRegionBudgets)
    let xs = ""
    if ( budgetsData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      budgetsData.forEach((reportKey) => {
        xs += `**${OrgReportRegionBudgetStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgReportRegionBudgets[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrgReportRegionBudgetStrings.numReportBudgets}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((budgetKey) => {
          //console.log('Report: ', values[1][budgetKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][budgetKey].budgetLine != "" ) {
            const budgetLine = ethers.utils.parseBytes32String(values[1][budgetKey].budgetLine)
            const start = ethers.utils.parseBytes32String(values[1][budgetKey].finance.start)
            const end = ethers.utils.parseBytes32String(values[1][budgetKey].finance.end)
            xs+= `**${OrgReportRegionBudgetStrings.budgetReference}**: ${budgetKey} <br />`
            xs+= `**${OrgReportRegionBudgetStrings.regionReference}**: ${values[1][budgetKey].regionRef} <br />`
            xs+= `**${OrgReportRegionBudgetStrings.budgetLine}**: ${budgetLine} <br />`
            xs+= `**${OrgReportRegionBudgetStrings.value}**: ${values[1][budgetKey].finance.value} <br />`
            xs+= `**${OrgReportRegionBudgetStrings.status}**: ${values[1][budgetKey].finance.status} <br />`
            xs+= `**${OrgReportRegionBudgetStrings.budgetStart}**: ${start} <br />`
            xs+= `**${OrgReportRegionBudgetStrings.budgetEnd}**: ${end} <br /><br />`
          }
        })
        length += 1
        length == budgetsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }

    return (
      <div>
        <h2>{OrgReportRegionBudgetStrings.headingOrgReportRegionBudgetReader}</h2>
        <p>
          <b>{OrgReportRegionBudgetStrings.numReports}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrgReportRegionBudgetStrings.reportBudgetDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgReportRegionBudgetProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgReportRegionBudgetsReader.num,
    orgReportRegionBudgets: state.orgReportRegionBudgetsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgReportRegionBudgetDispatchProps => {
  return {
    getReportRegionBudgets: () => dispatch(getReportRegionBudgets())
  }
}

export const OrganisationReportRegionBudgets = withTheme(withStyles(styles)(connect<OrgReportRegionBudgetProps, OrgReportRegionBudgetDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgReportRegionBudgets)))
