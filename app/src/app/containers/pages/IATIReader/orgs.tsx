import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'
import Markdown from 'react-markdown'

import { getDictEntries } from '../../../components/io/dict'
import { getOrgs } from '../../../store/IATI/IATIReader/organisations/orgs/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrgReport } from '../../../store/IATI/IATIReader/organisations/types'

import { Org as OrgStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgProps {
  orgs: IATIOrgReport
}

interface OrgDispatchProps {
  getOrgs: (isReport: boolean) => void
}

type OrgReaderProps =  WithStyles<typeof styles> & OrgProps & OrgDispatchProps

class OrgsReader extends React.Component<OrgReaderProps> {

  constructor (props: OrgReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getOrgs(true)
  }

  render() {

    const xs = getDictEntries(this.props.orgs)

    return (
      <div>
        <h2>{OrgStrings.headingOrgReader}</h2>
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
    orgs: state.report.data as IATIOrgReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    getOrgs: (isReport: boolean) => dispatch(getOrgs(isReport))
  }
}

export const Orgs = withTheme(withStyles(styles)(connect<OrgProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgsReader)))
