import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getRegionBudgets } from '../../../store/IATI/IATIReader/organisations/organisationRegionBudgets/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrganisationRegionBudgetsData } from '../../../store/IATI/IATIReader/organisations/types'

import { OrganisationRegionBudget as OrganisationRegionBudgetStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrganisationRegionBudgetProps {
  num: number
  orgRegionBudgets: OrganisationRegionBudgetsData
}

interface OrganisationRegionBudgetDispatchProps {
  getRegionBudgets: () => void
}

type OrganisationRegionBudgetsReaderProps =  WithStyles<typeof styles> & OrganisationRegionBudgetProps & OrganisationRegionBudgetDispatchProps

class RegionBudgets extends React.Component<OrganisationRegionBudgetsReaderProps> {

  constructor (props: OrganisationRegionBudgetsReaderProps) {
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
        xs += `**${OrganisationRegionBudgetStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgRegionBudgets[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrganisationRegionBudgetStrings.numBudgets}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((budgetKey) => {
          //console.log(': ', values[1][budgetKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][budgetKey].hasOwnProperty('budgetLine') && values[1][budgetKey].budgetLine != "" ) {
            const budgetLine = ethers.utils.parseBytes32String(values[1][budgetKey].budgetLine)
            const start = ethers.utils.parseBytes32String(values[1][budgetKey].finance.start)
            const end = ethers.utils.parseBytes32String(values[1][budgetKey].finance.end)
            xs+= `**${OrganisationRegionBudgetStrings.reportingOrgRef}**: ${values[1][budgetKey].report.orgRef} <br />`
            xs+= `**${OrganisationRegionBudgetStrings.budgetReference}**: ${budgetKey} <br />`
            xs+= `**${OrganisationRegionBudgetStrings.regionReference}**: ${values[1][budgetKey].regionRef} <br />`
            xs+= `**${OrganisationRegionBudgetStrings.budgetLine}**: ${budgetLine} <br />`
            xs+= `**${OrganisationRegionBudgetStrings.value}**: ${values[1][budgetKey].finance.value} <br />`
            xs+= `**${OrganisationRegionBudgetStrings.status}**: ${values[1][budgetKey].finance.status} <br />`
            xs+= `**${OrganisationRegionBudgetStrings.budgetStart}**: ${start} <br />`
            xs+= `**${OrganisationRegionBudgetStrings.budgetEnd}**: ${end} <br /><br />`
          }
        })
        length += 1
        length == budgetsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }

    return (
      <div>
        <h2>{OrganisationRegionBudgetStrings.headingOrganisationRegionBudgetReader}</h2>
        <p>
          <b>{OrganisationRegionBudgetStrings.nums}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrganisationRegionBudgetStrings.reportBudgetDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationRegionBudgetProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgRegionBudgetsReader.num,
    orgRegionBudgets: state.orgRegionBudgetsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationRegionBudgetDispatchProps => {
  return {
    getRegionBudgets: () => dispatch(getRegionBudgets())
  }
}

export const OrganisationRegionBudgets = withTheme(withStyles(styles)(connect<OrganisationRegionBudgetProps, OrganisationRegionBudgetDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(RegionBudgets)))
