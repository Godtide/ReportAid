import * as React from 'react'
import Tooltip from '@material-ui/core/Tooltip'
import TextField from '@material-ui/core/TextField'

interface TextAreaInputProps {
  tip: string
  label: string
  submit: () => void
}

const TextInput: React.SFC<TextAreaInputProps> = (props: TextAreaInputProps) => {

  return (
    <Tooltip title={props.tip}>
      <p>{props.label}
        <TextField
          label={props.label}
          multiline={false}
          onChange={props.submit}
        />
      </p>
    </Tooltip>
  )
}

export default TextInput
