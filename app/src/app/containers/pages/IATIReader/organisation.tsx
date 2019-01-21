import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'
import Markdown from 'react-markdown'

import { getOrgs } from '../../../store/IATI/IATIReader/organisation/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrgData } from '../../../store/IATI/IATIReader/organisation/types'

import { Organisation as OrgStrings } from '../../../utils/strings'

import { getDictEntries } from '../../../components/io/dict'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgProps {
  num: number
  orgs: OrgData
}

interface OrgDispatchProps {
  getOrgs: () => void
}

type OrgReaderProps =  WithStyles<typeof styles> & OrgProps & OrgDispatchProps

export class Orgs extends React.Component<OrgReaderProps> {

  constructor (props: OrgReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getOrgs()
  }

  render() {

    let xs: string = ""
    //console.log('Orgs props: ', this.props.orgs)
    Object.keys(this.props.orgs).forEach((key) => {
      xs += `**${OrgStrings.orgIdentifier}**: ${key}, `
      xs += `**${OrgStrings.orgName}**: ${this.props.orgs[key].name}, `
      xs += `**${OrgStrings.identifier}**: ${this.props.orgs[key].identifier}<br />`
    })

    const date = new Date('12/31/2009')
    console.log('Date is, ', date.toISOString())

    return (
      <div>
        <h2>{OrgStrings.headingOrgReader}</h2>
        <p>
          <b>{OrgStrings.numOrgs}</b>: {this.props.num}
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
    num: state.orgReader.num,
    orgs: state.orgReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    getOrgs: () => dispatch(getOrgs())
  }
}

export const Organisations = withTheme(withStyles(styles)(connect<OrgProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Orgs)))
