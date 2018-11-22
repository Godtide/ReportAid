import * as React from 'react'
import {
  Field,
  InjectedFormProps,
  reduxForm,
} from 'redux-form'

import { Organisation } from '../../utils/strings'

export interface OrgFormProps {
  onSubmit: () => void
}

type AllProps = OrgFormProps & InjectedFormProps<{}, OrgFormProps>

const renderInput = (field: any) => {
	return (
    	<input {...field.input} type={field.type}/>
	)
}

const Form: React.SFC<AllProps> = (props: AllProps) => {

  const { handleSubmit } = props

  return (
    <form onSubmit={handleSubmit}>
      <div>
        <label htmlFor={Organisation.orgName}>{Organisation.orgName}: </label>
        <Field
          name={Organisation.orgName}
          component={renderInput}
          type="text"
          placeholder={Organisation.orgName}
        />
      </div>
      <div>
        <label htmlFor={Organisation.identifier}>{Organisation.identifier}: </label>
        <Field
          name={Organisation.identifier}
          component={renderInput}
          type="text"
          placeholder={Organisation.identifier}
        />
      </div>
      <div>
        <label htmlFor={Organisation.type}>{Organisation.type}: </label>
        <Field
          name={Organisation.type}
          component={renderInput}
          type="text"
          placeholder={Organisation.type}
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
