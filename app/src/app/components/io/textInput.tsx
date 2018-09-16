import * as React from 'react'
import { Tooltip, Input } from 'antd'

interface TextInputInterface {
  tip: string
  label: string
  placeHolder: string
}

export type TextInputProps = TextInputInterface

const TextInput: React.SFC<TextInputProps> = (props) => {

  return (
    <Tooltip title={props.tip}>
      <Input
        addonBefore={props.label}
        defaultValue={props.placeHolder}
        onChange={}
      />
    </Tooltip>
  )
}

export default TextInput
