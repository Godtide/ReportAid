import * as React from 'react'
import logo from '../images/logo.png'

import Grid from '@material-ui/core/Grid'
import Paper from '@material-ui/core/Paper'
import { SiderCreateMenu } from './siderCreateMenu'
import { SiderReadMenu } from './siderReadMenu'
import { ApplicationBar } from './appBar'
import { Content } from './content'
import { App } from '../utils/strings'


import { MarkdownText } from '../components/io/markdownText'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles';
import { withTheme, styles } from '../styles/theme'

class MainLayout extends React.Component<WithStyles<typeof styles>> {

  render() {

    return (
      <div>
        <Paper className={this.props.classes.root}>
          <Paper className={this.props.classes.header}>
            <Grid container>
              <Grid item xs={12} sm={2}>
                <img className={this.props.classes.button} src={logo}/>
              </Grid>
              <Grid item xs={12} sm={8}>
                <Paper className={this.props.classes.appBar}>
                  <ApplicationBar />
                </Paper>
              </Grid>
              <Grid item xs={12} sm={2}>
                <Paper className={this.props.classes.header}>
                  <h1>{App.title}</h1>
                </Paper>
              </Grid>
            </Grid>
          </Paper>
          <Paper className={this.props.classes.content}>
            <Grid container spacing={0}>
              <Grid item xs={12} sm={2}>
                <Paper className={this.props.classes.sider}>
                  <SiderCreateMenu />
                </Paper>
                <Paper className={this.props.classes.sider}>
                  <SiderReadMenu />
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
                  <MarkdownText text={App.author} />
                </Paper>
              </Grid>
              <Grid item xs={8}>
              <Paper className={this.props.classes.title}>
                <h3>{App.tagline}</h3>
              </Paper>
              </Grid>
              <Grid item xs={2}>
                <Paper className={this.props.classes.footer}>
                  <MarkdownText text={App.copyright} />
                </Paper>
              </Grid>
            </Grid>
          </Paper>
        </Paper>
      </div>
    )
  }
}

export const Main = withTheme(withStyles(styles)(MainLayout))
