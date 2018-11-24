import { connect } from 'react-redux'
import { Dispatch, AnyAction } from 'redux'

import { OrgWriter, OrgDispatchProps } from '../../../components/blockchain/organisationWriter'
import { addOrganisation } from '../../../store/IATIWriter/organisationWriter/actions'

import withStyles from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

function mapDispatchToProps(dispatch: Dispatch<AnyAction>): OrgDispatchProps {
  return {
    handleSubmit: (ownProps: any) => dispatch(addOrganisation(ownProps))
  }
}

export const Organisation = withTheme(withStyles(styles)(connect(
  null,
  mapDispatchToProps
)(OrgWriter)))
