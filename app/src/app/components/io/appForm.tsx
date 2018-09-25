import * as React from 'react'
import { Field } from 'redux-form'

interface AppFormProps {
  pristine: boolean
  submitting: boolean
  reset: () => void
  handleSubmit: () => void
}

const AppForm: React.SFC<AppFormProps> = (props: AppFormProps) => {

  return (
    <div>
      <form onSubmit={props.handleSubmit}>
        <div>
          <label>First Name </label>
          <Field
            name="firstName"
            component="input"
            type="text"
            placeholder="First Name"
          />
        </div>
        <div>
          <label>Last Name </label>
          <Field
            name="lastName"
            component="input"
            type="text"
            placeholder="Last Name"
          />
        </div>
        <div>
          <button type="submit" disabled={props.pristine || props.submitting}>
            Submit
          </button>
          <button type="button" disabled={props.pristine || props.submitting} onClick={props.reset}>
            Clear Values
          </button>
        </div>
      </form>
    </div>
  )
}

export default AppForm
