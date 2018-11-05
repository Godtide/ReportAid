import * as React from 'react'
import Tooltip from '@material-ui/core/Tooltip'
import TextField from '@material-ui/core/TextField'

interface TextAreaInputProps {
  tip: string
  label: string
  submit: () => void
}

export const TextAreaInput: React.SFC<TextAreaInputProps> = (props: TextAreaInputProps) => {

  return (
    <Tooltip title={props.tip}>
      <p>{props.label}
        <TextField
          label={props.label}
          multiline={true}
          onChange={props.submit}
        />
      </p>
    </Tooltip>
  )
}
