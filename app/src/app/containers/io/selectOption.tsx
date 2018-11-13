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

interface SelectState {
  readonly value: string
}

export type SelectProps = SelectInterface

export class SelectOption extends React.Component<SelectProps, SelectState> {

  state: SelectState = { value: '' }

  handleChange = (event: any) => {
    this.setState({ value: event.target.value })
    this.props.onChange
  }

  render() {

    const { handleChange } = this

    return (
      <Tooltip title={this.props.tip}>
        <Select
              value={this.state.value}
              onChange={handleChange}
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
}
