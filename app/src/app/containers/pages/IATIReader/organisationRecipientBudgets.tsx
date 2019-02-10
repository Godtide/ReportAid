import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getRecipientBudgets } from '../../../store/IATI/IATIReader/organisations/organisationRecipientBudgets/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrgRecipientBudgetsData } from '../../../store/IATI/IATIReader/organisationRecipientBudgets/types'

import { OrganisationRecipientBudget as OrgRecipientBudgetStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgRecipientBudgetProps {
  num: number
  orgRecipientBudgets: OrgRecipientBudgetsData
}

interface OrgRecipientBudgetDispatchProps {
  getRecipientBudgets: () => void
}

type OrgRecipientBudgetsReaderProps =  WithStyles<typeof styles> & OrgRecipientBudgetProps & OrgRecipientBudgetDispatchProps

export class OrgRecipientBudgets extends React.Component<OrgRecipientBudgetsReaderProps> {

  constructor (props: OrgRecipientBudgetsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getRecipientBudgets()
  }

  render() {

    const budgetsData = Object.keys(this.props.orgRecipientBudgets)
    let xs = ""
    if ( budgetsData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      budgetsData.forEach((reportKey) => {
        xs += `**${OrgRecipientBudgetStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgRecipientBudgets[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrgRecipientBudgetStrings.numBudgets}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((budgetKey) => {
          //console.log(': ', values[1][budgetKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][budgetKey].hasOwnProperty('budgetLine') && values[1][budgetKey].budgetLine != "" ) {
            const budgetLine = ethers.utils.parseBytes32String(values[1][budgetKey].budgetLine)
            const start = ethers.utils.parseBytes32String(values[1][budgetKey].finance.start)
            const end = ethers.utils.parseBytes32String(values[1][budgetKey].finance.end)
            xs+= `**${OrgRecipientBudgetStrings.reportingOrgRef}**: ${values[1][budgetKey].report.orgRef} <br />`
            xs+= `**${OrgRecipientBudgetStrings.budgetReference}**: ${budgetKey} <br />`
            xs+= `**${OrgRecipientBudgetStrings.orgReference}**: ${values[1][budgetKey].orgRef} <br />`
            xs+= `**${OrgRecipientBudgetStrings.budgetLine}**: ${budgetLine} <br />`
            xs+= `**${OrgRecipientBudgetStrings.value}**: ${values[1][budgetKey].finance.value} <br />`
            xs+= `**${OrgRecipientBudgetStrings.status}**: ${values[1][budgetKey].finance.status} <br />`
            xs+= `**${OrgRecipientBudgetStrings.budgetStart}**: ${start} <br />`
            xs+= `**${OrgRecipientBudgetStrings.budgetEnd}**: ${end} <br /><br />`
          }
        })
        length += 1
        length == budgetsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }

    return (
      <div>
        <h2>{OrgRecipientBudgetStrings.headingOrgRecipientBudgetReader}</h2>
        <p>
          <b>{OrgRecipientBudgetStrings.nums}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrgRecipientBudgetStrings.reportBudgetDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgRecipientBudgetProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgRecipientBudgetsReader.num,
    orgRecipientBudgets: state.orgRecipientBudgetsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgRecipientBudgetDispatchProps => {
  return {
    getRecipientBudgets: () => dispatch(getRecipientBudgets())
  }
}

export const OrganisationRecipientBudgets = withTheme(withStyles(styles)(connect<OrgRecipientBudgetProps, OrgRecipientBudgetDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgRecipientBudgets)))
