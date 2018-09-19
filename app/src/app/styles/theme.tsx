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
    primary: purple,
    secondary: green,
  },
})

const styles = (theme: Theme) =>
  createStyles({
    root: {
      textAlign: 'center',
      padding: theme.spacing.unit,
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
