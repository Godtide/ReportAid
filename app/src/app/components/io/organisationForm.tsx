import * as React from 'react'
import {
  Field,
  InjectedFormProps,
  reduxForm,
} from 'redux-form'
import TextField from '@material-ui/core/TextField'

import { Organisation } from '../../utils/strings'

export interface OrgFormProps {
  onSubmit: (values: any) => any
}

type AllProps = OrgFormProps & InjectedFormProps<{}, OrgFormProps>

const renderTextField = ({input}: any) => {

  console.log('here', input)
  return (
    <TextField
      {...input}
    />
  )
}

const Form: React.SFC<AllProps> = (props: AllProps) => {

  const { handleSubmit, onSubmit } = props

  return (

    <form onSubmit={handleSubmit(onSubmit)}>
      <div>
        <label htmlFor={Organisation.orgName}>{Organisation.orgName}: </label>
        <Field
          name={Organisation.orgName}
          component={renderTextField}
        />
      </div>
      <div>
        <label htmlFor={Organisation.identifier}>{Organisation.identifier}: </label>
        <Field
          name={Organisation.identifier}
          component={renderTextField}
        />
      </div>
      <div>
        <label htmlFor={Organisation.type}>{Organisation.type}: </label>
        <Field
          name={Organisation.type}
          component={renderTextField}
        />
      </div>
      <div>
        <button type="submit">
          Submit
        </button>
      </div>
    </form>

  )
}

export const OrgForm = reduxForm<{}, OrgFormProps>({
  form: "orgForm"
})(Form)
