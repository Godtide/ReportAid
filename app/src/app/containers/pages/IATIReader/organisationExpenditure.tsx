import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getExpenditure } from '../../../store/IATI/IATIReader/organisationExpenditure/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrgExpenditureData } from '../../../store/IATI/IATIReader/organisationExpenditure/types'

import { OrganisationExpenditure as OrgExpenditureStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgExpenditureProps {
  num: number
  orgExpenditure: OrgExpenditureData
}

interface OrgExpenditureDispatchProps {
  getExpenditure: () => void
}

type OrgExpenditureReaderProps =  WithStyles<typeof styles> & OrgExpenditureProps & OrgExpenditureDispatchProps

export class OrgExpenditure extends React.Component<OrgExpenditureReaderProps> {

  constructor (props: OrgExpenditureReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getExpenditure()
  }

  render() {

    const expenditureData = Object.keys(this.props.orgExpenditure)
    let xs = ""
    if ( expenditureData.length > 0 ) {
      let length = 0
      //console.log ("Orgsdata: ", orgsData, " length ", orgsData.length )
      expenditureData.forEach((reportKey) => {
        xs += `**${OrgExpenditureStrings.reportReference}**: ${reportKey}<br />`
        const values = Object.values(this.props.orgExpenditure[reportKey])
        //console.log('Values: ', values)
        xs += `**${OrgExpenditureStrings.numExpenditure}**: ${values[0]} <br /><br />`
        Object.keys(values[1]).forEach((expenditureKey) => {
          //console.log(': ', values[1][expenditureKey])
          //const version = ethers.utils.parseBytes32String(values[1][thisKey].version)
          if ( values[1][expenditureKey].hasOwnProperty('expenditureLine') && values[1][expenditureKey].expenditureLine != "" ) {
            const expenditureLine = ethers.utils.parseBytes32String(values[1][expenditureKey].expenditureLine)
            const start = ethers.utils.parseBytes32String(values[1][expenditureKey].finance.start)
            const end = ethers.utils.parseBytes32String(values[1][expenditureKey].finance.end)
            xs+= `**${OrgExpenditureStrings.reportingOrgRef}**: ${values[1][expenditureKey].report.orgRef} <br />`
            xs+= `**${OrgExpenditureStrings.expenditureReference}**: ${expenditureKey} <br />`
            xs+= `**${OrgExpenditureStrings.expenditureLine}**: ${expenditureLine} <br />`
            xs+= `**${OrgExpenditureStrings.value}**: ${values[1][expenditureKey].finance.value} <br />`
            xs+= `**${OrgExpenditureStrings.status}**: ${values[1][expenditureKey].finance.status} <br />`
            xs+= `**${OrgExpenditureStrings.expenditureStart}**: ${start} <br />`
            xs+= `**${OrgExpenditureStrings.expenditureEnd}**: ${end} <br /><br />`
          }
        })
        length += 1
        length == expenditureData.length ? xs += "" : xs += "---<br /><br />"
      })
    }

    return (
      <div>
        <h2>{OrgExpenditureStrings.headingOrgExpenditureReader}</h2>
        <p>
          <b>{OrgExpenditureStrings.nums}</b>: {this.props.num}
        </p>
        <hr />
        <h3>{OrgExpenditureStrings.reportExpenditureDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgExpenditureProps => {
  //console.log(state.orgReader)
  return {
    num: state.orgExpenditureReader.num,
    orgExpenditure: state.orgExpenditureReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgExpenditureDispatchProps => {
  return {
    getExpenditure: () => dispatch(getExpenditure())
  }
}

export const OrganisationExpenditure = withTheme(withStyles(styles)(connect<OrgExpenditureProps, OrgExpenditureDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgExpenditure)))
