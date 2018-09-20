import * as React from 'react'

import createStyles from '@material-ui/core/styles/createStyles'
import { Theme } from '@material-ui/core/styles/createMuiTheme'
import { MuiThemeProvider, createMuiTheme } from '@material-ui/core/styles'
import purple from '@material-ui/core/colors/purple'
import green from '@material-ui/core/colors/green'


// A theme with custom primary and secondary color.
// It's optional.
const theme = createMuiTheme({
  palette: {
    background: {
      default: '#FAFAFA',
    },
    primary: {
      light: '##e1ffb1',
      main: '#c8e6c9',
      dark: '#97b498',
      contrastText: '#000000',
    },
    secondary: {
      light: '#e1ffb1',
      main: '#aed581',
      dark: '#7da453',
      contrastText: '#000000',
    },
  },
})

const styles = (theme: Theme) =>
  createStyles({
    root: {
      padding: theme.spacing.unit * 2,
      backgroundColor: theme.palette.background.default,
      flexGrow: 1,
    },
    paper: {
      padding: theme.spacing.unit * 2,
      textAlign: 'center',
      color: theme.palette.text.secondary,
    },
    header: {
      padding: theme.spacing.unit * 2,
      margin: theme.spacing.unit,
      textAlign: 'center',
      backgroundColor: theme.palette.secondary.light,
    },
    content: {
      padding: theme.spacing.unit * 2,
      margin: theme.spacing.unit,
      backgroundColor: theme.palette.primary.main,
    },
    sider: {
      padding: theme.spacing.unit * 2,
      margin: theme.spacing.unit,
      backgroundColor: theme.palette.secondary.light,
    },
    footer: {
      padding: theme.spacing.unit * 2,
      margin: theme.spacing.unit,
      textAlign: 'center',
      backgroundColor: theme.palette.secondary.light,
    },
  })

function withTheme<P>(Component: React.ComponentType<P>) {
  function WithTheme(props: P) {
    return (
      <MuiThemeProvider theme={theme}>
        <Component {...props} />
      </MuiThemeProvider>
    )
  }

  return WithTheme
}

export { withTheme, styles }
