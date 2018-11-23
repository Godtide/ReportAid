import * as React from 'react'
import { connect } from 'react-redux'
import { Dispatch, AnyAction } from 'redux'
/*import { ConfigProps } from 'redux-form'
import { ApplicationState } from '../../../store'*/

import { OrgForm } from '../../../components/io/organisationForm'
import { addOrganisation } from '../../../store/IATIWriter/organisationWriter/actions'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface IDispatchProps {
  handleSubmit: (ownProps: any) => void
}

type AllProps = WithStyles<typeof styles> & IDispatchProps

const Org: React.SFC<AllProps> = (props: AllProps) => {

  return (
    <OrgForm
      onSubmit={(values: any) => props.handleSubmit(values)}
      onSubmitSuccess={() => console.log('submit success')}
      onSubmitFail={() => console.log('submit fail')}
    />
  )
}

/*function mapStateToProps(state: ApplicationState): ConfigProps<IFormData> {
  return {
    form: "OrgForm", // Form will be handled by Redux Form using this key
    initialValues: {}
  }
}*/

function mapDispatchToProps(dispatch: Dispatch<AnyAction>): IDispatchProps {
  return {
    handleSubmit: (ownProps: any) => dispatch(addOrganisation(ownProps))
  }
}

export const OrganisationForm = withTheme(withStyles(styles)(connect(
  null,
  mapDispatchToProps
)(Org)))

/*class Organisation extends React.Component<WithStyles<typeof styles>> {

  render() {

    return (
      <OrganisationForm pristine={true} submitting={true}/>
    )
  }
}

export const OrgWriter = withTheme(withStyles(styles)(Organisation))*/