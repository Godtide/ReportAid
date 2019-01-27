import * as React from 'react'
import { Field, ErrorMessage} from 'formik'
import Grid from '@material-ui/core/Grid'
import { Select } from "material-ui-formik-components"

import { Helpers } from '../../utils/config'

interface DatePickerProps {
  dates: {
    day: {
      name: string,
      label: string
    }
    month: {
      name: string,
      label: string
    }
    year: {
      name: string,
      label: string
    }
  }
}

export const FormikDatePicker = (props: DatePickerProps) => {

  /*let dayRefs: any[] = []
  Array.from({ length: 31 }, (v: number, i: number) => {
    const value = ++i
    dayRefs.push({ value: value, label: value.toString() })
  })

  let monthRefs: any[] = []
  Array.from({ length: 12 }, (v: number, i: number) => {
    const value = ++i
    monthRefs.push({ value: value, label: value.toString() })
  })

  const startYear = 2000
  const stopYear = 2040
  const step = 1
  let yearRefs: any[] = []
  Array.from({ length: (stopYear - startYear) / step }, (_, i: number) => {
    const year = startYear + (i * step)
    yearRefs.push({ value: year, label: year.toString() })
  })*/

  return (
    <Grid container>
      <Grid item xs={12} sm={3}>
        <Field
          name={props.dates.day.name}
          label={props.dates.day.label}
          //style={{ width: '10%' }}
          component={Select}
          options={Helpers.days}
        />
        <ErrorMessage name={props.dates.day.name} />
      </Grid>
      <Grid item xs={12} sm={3}>
        <Field
          name={props.dates.month.name}
          label={props.dates.month.label}
          //style={{ width: '10%' }}
          component={Select}
          options={Helpers.months}
        />
        <ErrorMessage name={props.dates.month.name} />
      </Grid>
      <Grid item xs={12} sm={6}>
        <Field
          name={props.dates.year.name}
          label={props.dates.year.label}
          //style={{ width: '10%' }}
          component={Select}
          options={Helpers.years}
        />
        <ErrorMessage name={props.dates.year.name} />
      </Grid>
    </Grid>
  )
}
