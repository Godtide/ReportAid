import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getRecipientBudgets } from '../../../store/IATI/IATIReader/organisations/organisationRecipientBudgets/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrganisationRecipientBudgetsData } from '../../../store/IATI/IATIReader/types'

import { OrganisationRecipientBudget as OrganisationRecipientBudgetStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrganisationRecipientBudgetProps {
  num: number
  orgRecipientBudgets: OrganisationRecipientBudgetsData
}

interface OrganisationRecipientBudgetDispatchProps {
  getRecipientBudgets: () => void
}

type OrganisationRecipientBudgetsReaderProps =  WithStyles<typeof styles> & OrganisationRecipientBudgetProps & OrganisationRecipientBudgetDispatchProps

class RecipientBudgets extends React.Component<OrganisationRecipientBudgetsReaderProps> {

  constructor (props: OrganisationRecipientBudgetsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getRecipientBudgets()
  }

  render() {

    const reportsData = Object.keys(this.props.organisations)
    let xs = ""
    let num = 0
    /*let xs = ""
    if ( budgetsData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      budgetsData.forEach((reportKey) => {
        xs += `**${OrganisationRecipientBudgetStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgRecipientBudgets[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrganisationRecipientBudgetStrings.numBudgets}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((budgetKey) => {
          //console.log(': ', values[1][budgetKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][budgetKey].hasOwnProperty('budgetLine') && values[1][budgetKey].budgetLine != "" ) {
            const budgetLine = ethers.utils.parseBytes32String(values[1][budgetKey].budgetLine)
            const start = ethers.utils.parseBytes32String(values[1][budgetKey].finance.start)
            const end = ethers.utils.parseBytes32String(values[1][budgetKey].finance.end)
            xs+= `**${OrganisationRecipientBudgetStrings.reportingOrgRef}**: ${values[1][budgetKey].report.orgRef} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.budgetReference}**: ${budgetKey} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.orgReference}**: ${values[1][budgetKey].orgRef} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.budgetLine}**: ${budgetLine} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.value}**: ${values[1][budgetKey].finance.value} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.status}**: ${values[1][budgetKey].finance.status} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.budgetStart}**: ${start} <br />`
            xs+= `**${OrganisationRecipientBudgetStrings.budgetEnd}**: ${end} <br /><br />`
          }
        })
        length += 1
        length == budgetsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }*/

    return (
      <div>
        <h2>{OrganisationRecipientBudgetStrings.headingOrganisationRecipientBudgetReader}</h2>
        <p>
          <b>{OrganisationRecipientBudgetStrings.nums}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrganisationRecipientBudgetStrings.reportBudgetDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationRecipientBudgetProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgRecipientBudgetsReader.num,
    orgRecipientBudgets: state.orgRecipientBudgetsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationRecipientBudgetDispatchProps => {
  return {
    getRecipientBudgets: () => dispatch(getRecipientBudgets())
  }
}

export const OrganisationRecipientBudgets = withTheme(withStyles(styles)(connect<OrganisationRecipientBudgetProps, OrganisationRecipientBudgetDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(RecipientBudgets)))
