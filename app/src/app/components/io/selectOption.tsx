import * as React from 'react'
import { Tooltip, Select } from 'antd'

interface SelectInterface {
  tip: string
  label: string
  selections: string[]
  optionIndex: number
}

export type SelectProps = SelectInterface

const SelectOption: React.SFC<SelectProps> = (props) => {

  const Optn = Select.Option
  let children = []
  let size = props.selections.length
  for (let i = 0; i < size; i++) {
    children.push(<Optn key={this.props.selections[i]}>{props.selections[i]}</Optn>)
  }

  return (
    <Tooltip title={props.tip}>
      <Select placeholder={props.label}
              style={{ width: '100%' }}
              onChange={}
              tokenSeparators={[',']}
      >
        {children}
      </Select>
    </Tooltip>
  )
}

export default SelectOption
