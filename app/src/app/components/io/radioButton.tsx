import * as React from 'react'
import { Tooltip, Radio } from 'antd'

interface RadioButtonInterface {
  label: string
  tip: string
  options: string[]
  optionIndex: string
}

export type RadioButtonProps = RadioButtonInterface

const RadioButton: React.SFC<RadioButtonProps> = (props) => {

  const RadGroup = Radio.Group

  return (
    <div>
    {props.label}
      <Tooltip title={props.tip}>
        <RadGroup
          options={props.options}
          onChange={}
          value={props.optionIndex}
        />
      </Tooltip>
    </div>
  )
}

  export default RadioButton
