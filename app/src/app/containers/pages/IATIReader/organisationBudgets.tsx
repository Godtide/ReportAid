import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getBudgets } from '../../../store/IATI/IATIReader/organisations/organisationBudgets/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrganisationBudgetsData } from '../../../store/IATI/IATIReader/types'

import { OrganisationBudget as OrgBudgetStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgBudgetProps {
  num: number
  orgBudgets: OrganisationBudgetsData
}

interface OrgBudgetDispatchProps {
  getBudgets: () => void
}

type OrganisationBudgetsReaderProps =  WithStyles<typeof styles> & OrgBudgetProps & OrgBudgetDispatchProps

class Budgets extends React.Component<OrganisationBudgetsReaderProps> {

  constructor (props: OrganisationBudgetsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getBudgets()
  }

  render() {

    const budgetsData = Object.keys(this.props.orgBudgets)
    let xs = ""
    if ( budgetsData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      budgetsData.forEach((reportKey) => {
        xs += `**${OrgBudgetStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgBudgets[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrgBudgetStrings.numBudgets}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((budgetKey) => {
          //console.log('Budget: ', values[1][budgetKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][budgetKey].hasOwnProperty('budgetLine') && values[1][budgetKey].budgetLine != "" ) {
            const budgetLine = ethers.utils.parseBytes32String(values[1][budgetKey].budgetLine)
            const start = ethers.utils.parseBytes32String(values[1][budgetKey].finance.start)
            const end = ethers.utils.parseBytes32String(values[1][budgetKey].finance.end)
            xs+= `**${OrgBudgetStrings.reportingOrgRef}**: ${values[1][budgetKey].report.orgRef} <br />`
            xs+= `**${OrgBudgetStrings.budgetReference}**: ${budgetKey} <br />`
            xs+= `**${OrgBudgetStrings.budgetLine}**: ${budgetLine} <br />`
            xs+= `**${OrgBudgetStrings.value}**: ${values[1][budgetKey].finance.value} <br />`
            xs+= `**${OrgBudgetStrings.status}**: ${values[1][budgetKey].finance.status} <br />`
            xs+= `**${OrgBudgetStrings.budgetStart}**: ${start} <br />`
            xs+= `**${OrgBudgetStrings.budgetEnd}**: ${end} <br /><br />`
          }
        })
        length += 1
        length == budgetsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }

    return (
      <div>
        <h2>{OrgBudgetStrings.headingOrganisationBudgetReader}</h2>
        <p>
          <b>{OrgBudgetStrings.num}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrgBudgetStrings.organisationBudgetDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgBudgetProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgBudgetsReader.num,
    orgBudgets: state.orgBudgetsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgBudgetDispatchProps => {
  return {
    getBudgets: () => dispatch(getBudgets())
  }
}

export const OrganisationBudgets = withTheme(withStyles(styles)(connect<OrgBudgetProps, OrgBudgetDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Budgets)))
