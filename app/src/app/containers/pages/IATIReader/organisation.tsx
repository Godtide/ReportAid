import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { getNumOrganisations } from '../../../store/IATI/IATIReader/organisationReader/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { Organisation as OrgStrings } from '../../../utils/strings'

//import { PlainTextWithTitle, PlainTextKeyedList } from '../../../components/io/plainText'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgProps {
  numOrgs: object
}

interface OrgDispatchProps {
  handleSubmit: () => void
}

type OrgReaderProps =  WithStyles<typeof styles> & OrgProps & OrgDispatchProps

export class OrgReader extends React.Component<OrgReaderProps> {

  constructor (props: OrgReaderProps) {
    super(props)
    props.handleSubmit()
  }

  render() {

    const num = this.props.numOrgs.toString()

    return (
      <div>
        <h2>{OrgStrings.headingOrgReader}</h2>
        <p>
          <b>{OrgStrings.numOrgs}</b>: {num}
        </p>
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgProps => {
  return {
    numOrgs: state.orgReader.data.result
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    handleSubmit: () => dispatch(getNumOrganisations())
  }
}

export const OrganisationReader = withTheme(withStyles(styles)(connect<OrgProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgReader)))
