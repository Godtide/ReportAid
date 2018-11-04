import * as React from 'react'
import logo from '../images/logo.png'
import Grid from '@material-ui/core/Grid'
import Paper from '@material-ui/core/Paper'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles';
import { withTheme, styles } from '../styles/theme'

// import Content from './content'
import AppMenu from './appMenu'
import Content from './content'
import { AppStrings } from '../utils/strings'

import { Blockchain } from './blockchain/blockchain'
import { OrgsContract } from './blockchain/orgsContract'

class Main extends React.Component<WithStyles<typeof styles>> {

  blockchain: any
  orgsContract: any

  constructor (props: any) {
    super(props)
    const thisAddress = {address: "http://localhost:8545"}
    this.blockchain = new Blockchain(thisAddress)
    const thisBlockchain = {blockchain: this.blockchain}
    this.orgsContract = new OrgsContract(thisBlockchain)
  }

  render() {

    return (
      <div>
        <Paper className={this.props.classes.root}>
          <Paper className={this.props.classes.header}>
            <Grid container justify='center' spacing={0}>
              <Grid item xs={12} sm={2}>
                <img className={this.props.classes.button} src={logo}/>
              </Grid>
              <Grid item xs={12} sm={10}>
                <Paper className={this.props.classes.title}>
                  <h1>{AppStrings.appTitle}</h1>
                  <h3>{AppStrings.appTagline}</h3>
                </Paper>
              </Grid>
            </Grid>
          </Paper>
          <Paper className={this.props.classes.content}>
            <Grid container spacing={0}>
              <Grid item xs={12} sm={2}>
                <Paper className={this.props.classes.sider}>
                  <AppMenu />
                </Paper>
              </Grid>
              <Grid item xs={12} sm={10}>
                <Paper className={this.props.classes.content}>
                  <Content />
                </Paper>
              </Grid>
            </Grid>
          </Paper>
          <Paper className={this.props.classes.footer}>
            <Grid container spacing={0}>
              <Grid item xs={2}>
                <Paper className={this.props.classes.footer}>
                  <p>{AppStrings.author}</p>
                </Paper>
              </Grid>
              <Grid item xs={10}>
                <Paper className={this.props.classes.footer}>
                  <p>{AppStrings.copyright}</p>
                </Paper>
              </Grid>
            </Grid>
          </Paper>
        </Paper>
      </div>
    )
  }
}

export default withTheme(withStyles(styles)(Main))
