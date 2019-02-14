import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { getOrganisations } from '../../../store/IATI/IATIReader/organisations/organisations/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrganisationsData } from '../../../store/IATI/IATIReader/organisations/organisations/types'

import { Organisations as OrganisationsStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgProps {
  organisations: IATIOrganisationsData
}

interface OrgDispatchProps {
  getOrganisations: () => void
}

type OrgReaderProps =  WithStyles<typeof styles> & OrgProps & OrgDispatchProps

class OrganisationsReader extends React.Component<OrgReaderProps> {

  constructor (props: OrgReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getOrganisations()
  }

  render() {

    //console.log(JSON.stringify(this.props.organisations))
    //const reportsData = Object.keys(this.props.organisations)
    const reportsData = this.props.organisations
    let xs = ""
    let num = 0
    Object.keys(reportsData).forEach((organisationsKey) => {
        if ( reportsData[organisationsKey].hasOwnProperty('version') &&
             reportsData[organisationsKey].version != "" ) {
          num += 1
          const version = ethers.utils.parseBytes32String(reportsData[organisationsKey].version)
          const generatedTime = ethers.utils.parseBytes32String(reportsData[organisationsKey].generatedTime)
          xs += `**${OrganisationsStrings.reportKey}**: ${organisationsKey}<br />`
          xs += `**${OrganisationsStrings.version}**: ${version}<br />`
          xs += `**${OrganisationsStrings.generated}**: ${generatedTime}<br /><br />`
        }
    })

    return (
      <div>
        <h2>{OrganisationsStrings.headingOrganisationsReader}</h2>
        <p>
          <b>{OrganisationsStrings.numOrganisations}</b>: {num}
        </p>
        <hr />
        <h3>{OrganisationsStrings.organisationsDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgProps => {
  //console.log(state.orgReader)
  return {
    organisations: state.organisationsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    getOrganisations: () => dispatch(getOrganisations())
  }
}

export const Organisations = withTheme(withStyles(styles)(connect<OrgProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrganisationsReader)))
