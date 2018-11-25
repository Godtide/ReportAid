import * as React from 'react'
import { connect } from 'react-redux'
import { Dispatch, AnyAction } from 'redux'

import { ApplicationState } from '../../../store'
import { OrgContractProps } from '../../../store/blockchain/types'

import { Organisation as OrgStrings } from '../../../utils/strings'
import { OrgWriterForm, OrgDispatchProps } from '../../../components/blockchain/OrgWriterForm'
import { addOrganisation } from '../../../store/IATIWriter/organisationWriter/actions'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

type OrgWriterProps =  WithStyles<typeof styles> & OrgContractProps & OrgDispatchProps

class OrgWriter extends React.Component<OrgWriterProps> {

  constructor (props: OrgWriterProps) {
    super(props)
    console.log('Org Contract', props.orgContract)
  }

  render() {

    return (
      <div>
        <h2>{OrgStrings.headingOrgWriter}</h2>
        <OrgWriterForm
          handleSubmit={(values: any) => this.props.handleSubmit(values)}
          name = ''
          reference = ''
          type = ''
        />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgContractProps => {
  return {
    orgContract: state.blockchain.orgContract
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
