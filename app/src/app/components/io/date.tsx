import * as React from 'react'
import TextField from '@material-ui/core/TextField'

interface DateInterface {
  label: string
  default: string
  textField: string
  container: string
}

export type DateProps = DateInterface

const DatePicker: React.SFC<DateProps> = (props) => {

  return (
    <form className={props.container} noValidate>
      <TextField
        id="date"
        label={props.label}
        type={props.default}
        className={props.textField}
        InputLabelProps={{
          shrink: true,
        }}
      />
    </form>
  );
}

export default DatePicker
