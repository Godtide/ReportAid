import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getCountryBudgets } from '../../../store/IATI/IATIReader/organisations/organisationCountryBudgets/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrgCountryBudgetsData } from '../../../store/IATI/IATIReader/organisationCountryBudgets/types'

import { OrganisationCountryBudget as OrgCountryBudgetStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgCountryBudgetProps {
  num: number
  orgCountryBudgets: OrgCountryBudgetsData
}

interface OrgCountryBudgetDispatchProps {
  getCountryBudgets: () => void
}

type OrgCountryBudgetsReaderProps =  WithStyles<typeof styles> & OrgCountryBudgetProps & OrgCountryBudgetDispatchProps

export class OrgCountryBudgets extends React.Component<OrgCountryBudgetsReaderProps> {

  constructor (props: OrgCountryBudgetsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getCountryBudgets()
  }

  render() {

    const budgetsData = Object.keys(this.props.orgCountryBudgets)
    let xs = ""
    if ( budgetsData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      budgetsData.forEach((reportKey) => {
        xs += `**${OrgCountryBudgetStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgCountryBudgets[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrgCountryBudgetStrings.numBudgets}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((budgetKey) => {
          //console.log(': ', values[1][budgetKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][budgetKey].hasOwnProperty('countryRef') && values[1][budgetKey].countryRef != "" ) {
            const countryRef = ethers.utils.parseBytes32String(values[1][budgetKey].countryRef)
            const budgetLine = ethers.utils.parseBytes32String(values[1][budgetKey].budgetLine)
            const start = ethers.utils.parseBytes32String(values[1][budgetKey].finance.start)
            const end = ethers.utils.parseBytes32String(values[1][budgetKey].finance.end)
            xs+= `**${OrgCountryBudgetStrings.reportingOrgRef}**: ${values[1][budgetKey].report.orgRef} <br />`
            xs+= `**${OrgCountryBudgetStrings.budgetReference}**: ${budgetKey} <br />`
            xs+= `**${OrgCountryBudgetStrings.countryReference}**: ${countryRef} <br />`
            xs+= `**${OrgCountryBudgetStrings.budgetLine}**: ${budgetLine} <br />`
            xs+= `**${OrgCountryBudgetStrings.value}**: ${values[1][budgetKey].finance.value} <br />`
            xs+= `**${OrgCountryBudgetStrings.status}**: ${values[1][budgetKey].finance.status} <br />`
            xs+= `**${OrgCountryBudgetStrings.budgetStart}**: ${start} <br />`
            xs+= `**${OrgCountryBudgetStrings.budgetEnd}**: ${end} <br /><br />`
          }
        })
        length += 1
        length == budgetsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }

    return (
      <div>
        <h2>{OrgCountryBudgetStrings.headingOrgCountryBudgetReader}</h2>
        <p>
          <b>{OrgCountryBudgetStrings.nums}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrgCountryBudgetStrings.reportBudgetDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgCountryBudgetProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgCountryBudgetsReader.num,
    orgCountryBudgets: state.orgCountryBudgetsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgCountryBudgetDispatchProps => {
  return {
    getCountryBudgets: () => dispatch(getCountryBudgets())
  }
}

export const OrganisationCountryBudgets = withTheme(withStyles(styles)(connect<OrgCountryBudgetProps, OrgCountryBudgetDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgCountryBudgets)))
