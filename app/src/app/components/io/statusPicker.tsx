import * as React from 'react'
import { Field, ErrorMessage } from 'formik'
import { Select } from "material-ui-formik-components"

import { Helpers } from '../../utils/config'

interface StatusPickerProps {
  name: string
  label: string
}

export const FormikStatusPicker = (props: StatusPickerProps) => {

  let status: any[] = []
  Helpers.financeStatus.forEach( (value: any) => {
    //console.log(value)
    status.push({ value: value.code, label: value.name })
  })

  return (
    <React.Fragment>
      <Field
        name={props.name}
        label={props.label}
        component={Select}
        options={status}
      />
      <ErrorMessage name={props.name} />
    </React.Fragment>
  )
}
