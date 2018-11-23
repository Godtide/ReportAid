import * as React from 'react'
import { connect } from 'react-redux'
import { Dispatch, AnyAction } from 'redux'
/*import { ConfigProps } from 'redux-form'
import { ApplicationState } from '../../../store'*/

import { Organisation } from '../../../utils/strings'

import { OrgForm, DispatchProps } from '../../../components/io/organisationForm'
import { addOrganisation } from '../../../store/IATIWriter/organisationWriter/actions'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

type AllProps = WithStyles<typeof styles> & DispatchProps

const Org: React.SFC<AllProps> = (props: AllProps) => {

  return (
    <div>
      <h2>{Organisation.headingOrgWriter}</h2>
      <OrgForm
        handleSubmit={(values: any) => props.handleSubmit(values)}
        name = ''
        reference = ''
        type = ''
      />
    </div>
  )
}

function mapDispatchToProps(dispatch: Dispatch<AnyAction>): DispatchProps {
  return {
    handleSubmit: (ownProps: any) => dispatch(addOrganisation(ownProps))
  }
}

export const OrganisationForm = withTheme(withStyles(styles)(connect(
  null,
  mapDispatchToProps
)(Org)))
