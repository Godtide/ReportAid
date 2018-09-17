import * as React from 'react'
import Tooltip from '@material-ui/core/Tooltip'
import Button from '@material-ui/core/Button'

interface AppButtonInterface {
  label: string
  tip: string
  submit: () => void
}

export type AppButtonProps = AppButtonInterface

const AppButton: React.SFC<AppButtonProps> = (props) => {

  return (
    <Tooltip title={props.tip}>
      <Button
        aria-label={props.label}
        onClick={props.submit}>
        >
      </Button>
    </Tooltip>
  )
}

export default AppButton
