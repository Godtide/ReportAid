import * as React from 'react'

import Markdown from 'react-markdown'
import Grid from '@material-ui/core/Grid'
import Paper from '@material-ui/core/Paper'

import { useStyles } from '../styles/theme';

import { SiderOrganisationMenu } from './siderOrganisationMenu'
import { ApplicationBar } from './appBar'
import { Content } from './content'
import { App } from '../utils/strings'

import logo from '../images/logo.png'


export const Main = () => {

  const classes = useStyles()

  return (
    <div>
      <Paper className={classes.root}>
        <Paper className={classes.header}>
          <Grid container>
            <Grid item xs={12} sm={1}>
              <img className={classes.button} src={logo}/>
            </Grid>
            <Grid item xs={12} sm={9}>
              <Paper className={classes.appBar}>
                <ApplicationBar />
              </Paper>
            </Grid>
            <Grid item xs={12} sm={2}>
              <Paper className={classes.header}>
                <h1>{App.title}</h1>
              </Paper>
            </Grid>
          </Grid>
        </Paper>
        <Paper className={classes.content}>
          <Grid container spacing={0}>
            <Grid item xs={12} sm={3}>
              <Paper className={classes.sider}>
                <SiderOrganisationMenu />
              </Paper>
            </Grid>
            <Grid item xs={12} sm={9}>
              <Paper className={classes.content}>
                <Content />
              </Paper>
            </Grid>
          </Grid>
        </Paper>
        <Paper className={classes.footer}>
          <Grid container spacing={0}>
            <Grid item xs={2}>
              <Paper className={classes.footer}>
                <Markdown escapeHtml={false} source={App.author} />
              </Paper>
            </Grid>
            <Grid item xs={8}>
            <Paper className={classes.title}>
              <h3>{App.tagline}</h3>
            </Paper>
            </Grid>
            <Grid item xs={2}>
              <Paper className={classes.footer}>
                <Markdown escapeHtml={false} source={App.copyright} />
              </Paper>
            </Grid>
          </Grid>
        </Paper>
      </Paper>
    </div>
  )
}
