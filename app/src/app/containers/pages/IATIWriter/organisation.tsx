import * as React from 'react'

import { Organisation as OrgStrings } from '../../../utils/strings'
import { OrgWriterForm } from '../../../components/blockchain/orgWriterForm'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

export class OrgWriter extends React.Component<WithStyles<typeof styles>> {

  render() {

    return (
      <div>
        <h2>{OrgStrings.headingOrgWriter}</h2>
        <OrgWriterForm />
      </div>
    )
  }
}

export const OrganisationWriter = withTheme(withStyles(styles)(OrgWriter))
