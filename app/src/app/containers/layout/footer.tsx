import * as React from 'react'

import Paper from '@material-ui/core/Paper'
import Grid from '@material-ui/core/Grid'
import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../styles/theme'

interface FooterProps {
  copyright: string
}

class Footer extends React.Component<WithStyles<typeof styles> & FooterProps> {

  render() {
    return (
      <React.Fragment>
        <Grid item xs={12}>
          <Paper>
            <h5>{this.props.copyright}</h5>
          </Paper>
        </Grid>
      </React.Fragment>
    )
  }
}

export default withTheme(withStyles(styles)(Footer))
