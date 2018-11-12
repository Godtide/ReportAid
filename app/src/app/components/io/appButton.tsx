import * as React from 'react'
import Tooltip from '@material-ui/core/Tooltip'
import Button from '@material-ui/core/Button'

interface AppButtonProps {
  label: string
  tip: string
  submit(event: any): void
}

export const AppButton: React.SFC<AppButtonProps> = (props: AppButtonProps) => {

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
