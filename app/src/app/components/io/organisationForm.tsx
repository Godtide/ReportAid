import * as React from 'react'
import {
  Field,
  InjectedFormProps,
  reduxForm,
} from 'redux-form'

import { Organisation } from '../../utils/strings'

export interface OrgFormProps {
  onSubmit: (values: any) => any
}

type AllProps = OrgFormProps & InjectedFormProps<{}, OrgFormProps>

const renderInput = (field: any) => {
  console.log(field.input)
	return (
    	<input {...field.input} type={field.type}/>
	)
}

const Form: React.SFC<AllProps> = (props: AllProps) => {

  const { handleSubmit, onSubmit } = props
  console.log('props', props)

  //{handleSubmit(data => onSubmit({...data, versionNumber: version}))}

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <div>
        <label htmlFor={Organisation.orgName}>{Organisation.orgName}: </label>
        <Field
          name={Organisation.orgName}
          component="input"
          type="text"
        />
      </div>
      <div>
        <label htmlFor={Organisation.identifier}>{Organisation.identifier}: </label>
        <Field
          name={Organisation.identifier}
          component="input"
          type="text"
        />
      </div>
      <div>
        <label htmlFor={Organisation.type}>{Organisation.type}: </label>
        <Field
          name={Organisation.type}
          component="input"
          type="text"
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
