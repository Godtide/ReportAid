import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getExpenditure } from '../../../store/IATI/IATIReader/organisations/organisationExpenditure/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrganisationExpenditureData } from '../../../store/IATI/IATIReader/types'

import { OrganisationExpenditure as OrganisationExpenditureStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrganisationExpenditureProps {
  num: number
  orgExpenditure: OrganisationExpenditureData
}

interface OrganisationExpenditureDispatchProps {
  getExpenditure: () => void
}

type OrganisationExpenditureReaderProps =  WithStyles<typeof styles> & OrganisationExpenditureProps & OrganisationExpenditureDispatchProps

class Expenditure extends React.Component<OrganisationExpenditureReaderProps> {

  constructor (props: OrganisationExpenditureReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getExpenditure()
  }

  render() {

    const reportsData = Object.keys(this.props.organisations)
    let xs = ""
    let num = 0
    /*let xs = ""
    if ( expenditureData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      expenditureData.forEach((reportKey) => {
        xs += `**${OrganisationExpenditureStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgExpenditure[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrganisationExpenditureStrings.numExpenditure}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((expenditureKey) => {
          //console.log(': ', values[1][expenditureKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][expenditureKey].hasOwnProperty('expenditureLine') && values[1][expenditureKey].expenditureLine != "" ) {
            const expenditureLine = ethers.utils.parseBytes32String(values[1][expenditureKey].expenditureLine)
            const start = ethers.utils.parseBytes32String(values[1][expenditureKey].finance.start)
            const end = ethers.utils.parseBytes32String(values[1][expenditureKey].finance.end)
            xs+= `**${OrganisationExpenditureStrings.reportingOrgRef}**: ${values[1][expenditureKey].report.orgRef} <br />`
            xs+= `**${OrganisationExpenditureStrings.expenditureReference}**: ${expenditureKey} <br />`
            xs+= `**${OrganisationExpenditureStrings.expenditureLine}**: ${expenditureLine} <br />`
            xs+= `**${OrganisationExpenditureStrings.value}**: ${values[1][expenditureKey].finance.value} <br />`
            xs+= `**${OrganisationExpenditureStrings.status}**: ${values[1][expenditureKey].finance.status} <br />`
            xs+= `**${OrganisationExpenditureStrings.expenditureStart}**: ${start} <br />`
            xs+= `**${OrganisationExpenditureStrings.expenditureEnd}**: ${end} <br /><br />`
          }
        })
        length += 1
        length == expenditureData.length ? xs += "" : xs += "---<br /><br />"
      })
    }*/

    return (
      <div>
        <h2>{OrganisationExpenditureStrings.headingOrganisationExpenditureReader}</h2>
        <p>
          <b>{OrganisationExpenditureStrings.nums}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrganisationExpenditureStrings.reportExpenditureDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationExpenditureProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgExpenditureReader.num,
    orgExpenditure: state.orgExpenditureReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationExpenditureDispatchProps => {
  return {
    getExpenditure: () => dispatch(getExpenditure())
  }
}

export const OrganisationExpenditure = withTheme(withStyles(styles)(connect<OrganisationExpenditureProps, OrganisationExpenditureDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Expenditure)))
