import * as React from 'react'
import Tooltip from '@material-ui/core/Tooltip'
import Select from '@material-ui/core/Select'
import MenuItem from '@material-ui/core/MenuItem'

interface SelectInterface {
  tip: string
  label: string
  selections: string[]
  onChange: () => void
}

export type SelectProps = SelectInterface

const SelectOption: React.SFC<SelectProps> = (props) => {

  return (
    <Tooltip title={props.tip}>
      <Select
            value={this.state.age}
            onChange={props.onChange}
            inputProps={{
              name: 'age',
              id: 'age-simple',
            }}
          >
            <MenuItem value="">
              <em>None</em>
            </MenuItem>
            <MenuItem value={10}>Ten</MenuItem>
            <MenuItem value={20}>Twenty</MenuItem>
            <MenuItem value={30}>Thirty</MenuItem>
          </Select>
    </Tooltip>
  )
}

export default SelectOption
