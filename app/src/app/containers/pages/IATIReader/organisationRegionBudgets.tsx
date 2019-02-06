import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getRegionBudgets } from '../../../store/IATI/IATIReader/organisationRegionBudgets/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrgRegionBudgetsData } from '../../../store/IATI/IATIReader/organisationRegionBudgets/types'

import { OrganisationRegionBudget as OrgRegionBudgetStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgRegionBudgetProps {
  num: number
  orgRegionBudgets: OrgRegionBudgetsData
}

interface OrgRegionBudgetDispatchProps {
  getRegionBudgets: () => void
}

type OrgRegionBudgetsReaderProps =  WithStyles<typeof styles> & OrgRegionBudgetProps & OrgRegionBudgetDispatchProps

export class OrgRegionBudgets extends React.Component<OrgRegionBudgetsReaderProps> {

  constructor (props: OrgRegionBudgetsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getRegionBudgets()
  }

  render() {

    const budgetsData = Object.keys(this.props.orgRegionBudgets)
    let xs = ""
    if ( budgetsData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      budgetsData.forEach((reportKey) => {
        xs += `**${OrgRegionBudgetStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgRegionBudgets[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrgRegionBudgetStrings.numBudgets}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((budgetKey) => {
          //console.log(': ', values[1][budgetKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][budgetKey].hasOwnProperty('budgetLine') && values[1][budgetKey].budgetLine != "" ) {
            const budgetLine = ethers.utils.parseBytes32String(values[1][budgetKey].budgetLine)
            const start = ethers.utils.parseBytes32String(values[1][budgetKey].finance.start)
            const end = ethers.utils.parseBytes32String(values[1][budgetKey].finance.end)
            xs+= `**${OrgRegionBudgetStrings.reportingOrgRef}**: ${values[1][budgetKey].report.orgRef} <br />`
            xs+= `**${OrgRegionBudgetStrings.budgetReference}**: ${budgetKey} <br />`
            xs+= `**${OrgRegionBudgetStrings.regionReference}**: ${values[1][budgetKey].regionRef} <br />`
            xs+= `**${OrgRegionBudgetStrings.budgetLine}**: ${budgetLine} <br />`
            xs+= `**${OrgRegionBudgetStrings.value}**: ${values[1][budgetKey].finance.value} <br />`
            xs+= `**${OrgRegionBudgetStrings.status}**: ${values[1][budgetKey].finance.status} <br />`
            xs+= `**${OrgRegionBudgetStrings.budgetStart}**: ${start} <br />`
            xs+= `**${OrgRegionBudgetStrings.budgetEnd}**: ${end} <br /><br />`
          }
        })
        length += 1
        length == budgetsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }

    return (
      <div>
        <h2>{OrgRegionBudgetStrings.headingOrgRegionBudgetReader}</h2>
        <p>
          <b>{OrgRegionBudgetStrings.nums}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrgRegionBudgetStrings.reportBudgetDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgRegionBudgetProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgRegionBudgetsReader.num,
    orgRegionBudgets: state.orgRegionBudgetsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgRegionBudgetDispatchProps => {
  return {
    getRegionBudgets: () => dispatch(getRegionBudgets())
  }
}

export const OrganisationRegionBudgets = withTheme(withStyles(styles)(connect<OrgRegionBudgetProps, OrgRegionBudgetDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgRegionBudgets)))
