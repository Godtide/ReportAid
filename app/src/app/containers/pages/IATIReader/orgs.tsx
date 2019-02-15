import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'
import Markdown from 'react-markdown'

import { getOrgs } from '../../../store/IATI/IATIReader/organisations/orgs/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrgReports } from '../../../store/IATI/IATIReader/organisations/orgs/types'

import { Org as OrgStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgProps {
  orgs: IATIOrgReports
}

interface OrgDispatchProps {
  getOrgs: () => void
}

type OrgReaderProps =  WithStyles<typeof styles> & OrgProps & OrgDispatchProps

class OrgsReader extends React.Component<OrgReaderProps> {

  constructor (props: OrgReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getOrgs()
  }

  render() {

    let xs: string = ""
    let num = 0
    //console.log("Orgs data: ", this.props.orgs)
    Object.keys(this.props.orgs).forEach((key) => {
      if ( this.props.orgs[key].hasOwnProperty('name') &&
           this.props.orgs[key].name != "" ) {
      //console.log(key)
        num += 1
        xs += `**${OrgStrings.orgIdentifier}**: ${key}, `
        xs += `**${OrgStrings.orgName}**: ${this.props.orgs[key].name}, `
        xs += `**${OrgStrings.identifier}**: ${this.props.orgs[key].identifier}<br />`
      }
    })

    return (
      <div>
        <h2>{OrgStrings.headingOrgReader}</h2>
        <p>
          <b>{OrgStrings.numOrgs}</b>: {num}
        </p>
        <hr />
        <h3>{OrgStrings.orgDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgProps => {
  //console.log(state.orgReader)
  return {
    orgs: state.orgsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    getOrgs: () => dispatch(getOrgs())
  }
}

export const Orgs = withTheme(withStyles(styles)(connect<OrgProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgsReader)))
