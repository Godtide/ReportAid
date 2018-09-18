import * as React from 'react'
import Tooltip from '@material-ui/core/Tooltip'
import Textfield from '@material-ui/core/Textfield'

interface TextAreaInputProps {
  tip: string
  label: string
  submit: () => void
}

const TextAreaInput: React.SFC<TextAreaInputProps> = (props) => {

  return (
    <Tooltip title={this.props.tip}>
      <p>{this.props.label}
        <Textfield
          label={props.label}
          multiline={true}
          onChange={props.submit}
        />
      </p>
    </Tooltip>
  )
}

export default TextAreaInput
