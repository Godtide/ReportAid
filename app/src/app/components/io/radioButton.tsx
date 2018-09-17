import * as React from 'react'
import Tooltip from '@material-ui/core/Tooltip'
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup'
import FormControlLabel from '@material-ui/core/FormControlLabel';
import FormControl from '@material-ui/core/FormControl';
import FormLabel from '@material-ui/core/FormLabel';

interface RadioButtonInterface {
  class: string
  submit: (event: {}, value: string) => void
  tip: string
  options: string[]
  optionIndex: string
}

export type RadioButtonProps = RadioButtonInterface

const RadioButton: React.SFC<RadioButtonProps> = (props) => {

  return (
    <Tooltip title={props.tip}>
      <FormControl component="fieldset" className={'blah'}>
          <FormLabel component="legend">Gender</FormLabel>
          <RadioGroup
            aria-label="Gender"
            name="gender1"
            className={'blah'}
            value={this.state.value}
            onChange={this.handleChange}
          >
            <FormControlLabel value="female" control={<Radio />} label="Female" />
            <FormControlLabel value="male" control={<Radio />} label="Male" />
            <FormControlLabel value="other" control={<Radio />} label="Other" />
            <FormControlLabel
              value="disabled"
              disabled
              control={<Radio />}
              label="(Disabled option)"
            />
          </RadioGroup>
        </FormControl>
    </Tooltip>
  )
}

  export default RadioButton
