import * as React from 'react'
import TextField from '@material-ui/core/TextField'

interface DateProps {
  label: string
  default: string
  textField: string
  container: string
}

export const DatePicker: React.SFC<DateProps> = (props: DateProps) => {

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
