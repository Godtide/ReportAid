import * as React from 'react'
import { Tooltip, Input } from 'antd'

interface TextAreaInputInterface {
  tip: string
  label: string
  numRows: number
}

export type TextAreaInputProps = TextAreaInputInterface

const TextAreaInput: React.SFC<TextAreaInputProps> = (props) => {

  const { TextArea } = Input

  return (
    <Tooltip title={this.props.tip}>
      <p>{this.props.label}
        <TextArea
          rows={props.numRows}
          onChange={}
        />
      </p>
    </Tooltip>
  )
}

export default TextAreaInput
