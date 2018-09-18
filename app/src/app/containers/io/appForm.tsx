import * as React from 'react';
import { reduxForm, InjectedFormProps, Field } from 'redux-form'

interface AppFormProps {
  message: string
  pristine: boolean
  submitting: boolean
  reset: () => void
  handleSubmit: () => void
}

class AppForm extends React.Component<InjectedFormProps<AppFormProps> & AppFormProps> {

  render() {
    const { pristine, submitting, reset, handleSubmit, message } = this.props
    return (
      <form onSubmit={handleSubmit}>
        <div>{message}</div>
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
          <button type="submit" disabled={pristine || submitting}>
            Submit
          </button>
          <button type="button" disabled={pristine || submitting} onClick={reset}>
            Clear Values
          </button>
        </div>
      </form>
    )
  }
}

export default reduxForm<AppFormProps>({
  form: 'userForm',
})(AppForm);
