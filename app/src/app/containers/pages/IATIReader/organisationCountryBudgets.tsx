import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getCountryBudgets } from '../../../store/IATI/IATIReader/organisations/organisationCountryBudgets/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrganisationsData } from '../../../store/IATI/IATIReader/organisations/types'

import { OrganisationCountryBudget as OrganisationCountryBudgetStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrganisationCountryBudgetProps {
  organisations: IATIOrganisationsData
}

interface OrganisationCountryBudgetDispatchProps {
  getCountryBudgets: () => void
}

type OrganisationCountryBudgetsReaderProps =  WithStyles<typeof styles> & OrganisationCountryBudgetProps & OrganisationCountryBudgetDispatchProps

class CountryBudgets extends React.Component<OrganisationCountryBudgetsReaderProps> {

  constructor (props: OrganisationCountryBudgetsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getCountryBudgets()
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
        xs += `**${OrganisationCountryBudgetStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgCountryBudgets[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrganisationCountryBudgetStrings.numBudgets}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((budgetKey) => {
          //console.log(': ', values[1][budgetKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][budgetKey].hasOwnProperty('countryRef') && values[1][budgetKey].countryRef != "" ) {
            const countryRef = ethers.utils.parseBytes32String(values[1][budgetKey].countryRef)
            const budgetLine = ethers.utils.parseBytes32String(values[1][budgetKey].budgetLine)
            const start = ethers.utils.parseBytes32String(values[1][budgetKey].finance.start)
            const end = ethers.utils.parseBytes32String(values[1][budgetKey].finance.end)
            xs+= `**${OrganisationCountryBudgetStrings.reportingOrgRef}**: ${values[1][budgetKey].report.orgRef} <br />`
            xs+= `**${OrganisationCountryBudgetStrings.budgetReference}**: ${budgetKey} <br />`
            xs+= `**${OrganisationCountryBudgetStrings.countryReference}**: ${countryRef} <br />`
            xs+= `**${OrganisationCountryBudgetStrings.budgetLine}**: ${budgetLine} <br />`
            xs+= `**${OrganisationCountryBudgetStrings.value}**: ${values[1][budgetKey].finance.value} <br />`
            xs+= `**${OrganisationCountryBudgetStrings.status}**: ${values[1][budgetKey].finance.status} <br />`
            xs+= `**${OrganisationCountryBudgetStrings.budgetStart}**: ${start} <br />`
            xs+= `**${OrganisationCountryBudgetStrings.budgetEnd}**: ${end} <br /><br />`
          }
        })
        length += 1
        length == budgetsData.length ? xs += "" : xs += "---<br /><br />"
      })
    }*/

    return (
      <div>
        <h2>{OrganisationCountryBudgetStrings.headingOrganisationCountryBudgetReader}</h2>
        <p>
          <b>{OrganisationCountryBudgetStrings.numBudgets}</b>: {num}
        </p>
        <hr />
        <h3>{OrganisationCountryBudgetStrings.organisationCountryBudgetDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationCountryBudgetProps => {
  //console.log(state.orgReader)
  return {
    organisations: state.organisationsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationCountryBudgetDispatchProps => {
  return {
    getCountryBudgets: () => dispatch(getCountryBudgets())
  }
}

export const OrganisationCountryBudgets = withTheme(withStyles(styles)(connect<OrganisationCountryBudgetProps, OrganisationCountryBudgetDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(CountryBudgets)))
