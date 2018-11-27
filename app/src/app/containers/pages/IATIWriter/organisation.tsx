import * as React from 'react'
//import { connect } from 'react-redux'
//import { Dispatch, AnyAction } from 'redux'

//import { ApplicationState } from '../../../store'

import { Organisation as OrgStrings } from '../../../utils/strings'
//import { OrgWriterForm, OrgDispatchProps } from '../../../components/blockchain/OrgWriterForm'
import { OrgWriterForm } from '../../../components/blockchain/OrgWriterForm'
//import { addOrganisation } from '../../../store/IATIWriter/organisationWriter/actions'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

//type OrgWriterProps =  WithStyles<typeof styles> & OrgContractProps & OrgDispatchProps

export class OrgWriter extends React.Component<WithStyles<typeof styles>> {

  /*constructor (props: OrgWriterProps) {
    super(props)
    console.log('Org Contract', props.orgContract)
  }*/

  render() {

    return (
      <div>
        <h2>{OrgStrings.headingOrgWriter}</h2>
        <OrgWriterForm />
      </div>
    )
  }
}

/*const mapStateToProps = (state: ApplicationState): OrgContractProps => {
  return {
    name: state.orgForm.name,
    reference: state.orgForm.reference,
    type: state.orgForm.type
  }
}

const mapDispatchToProps = (dispatch: Dispatch<AnyAction>): OrgDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(addOrganisation(ownProps))
  }
}

export const Organisation = withTheme(withStyles(styles)(connect(
  mapStateToProps,
  mapDispatchToProps
)(OrgWriter)))

export const Organisation = withTheme(withStyles(styles)(connect(
  mapStateToProps
)(OrgWriter)))*/

export const Organisation = withTheme(withStyles(styles)(OrgWriter))
