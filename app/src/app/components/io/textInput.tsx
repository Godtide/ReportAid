import * as React from 'react'
import Tooltip from '@material-ui/core/Tooltip'
import TextField from '@material-ui/core/TextField'

interface TextAreaInputInterface {
  tip: string
  label: string
  submit: () => void
}

export type TextAreaInputProps = TextAreaInputInterface

const TextInput: React.SFC<TextAreaInputProps> = (props) => {

  return (
    <Tooltip title={this.props.tip}>
      <p>{this.props.label}
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
