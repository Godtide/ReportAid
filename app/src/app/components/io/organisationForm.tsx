import * as React from 'react'
import {
  Field as FormField,
  InjectedFormProps,
  reduxForm,
} from 'redux-form'

import { Organisation } from '../../utils/strings'

export interface IFormData {
  bah: string
}

export interface IOwnProps {
  foo: string
}

export interface IDispatchProps {
  handleSubmit: (ownProps: any) => void
}

type AllProps = IOwnProps & IDispatchProps & InjectedFormProps<IFormData, IOwnProps>

/*
export interface DispatchProps {
  onSubmit: (data: IFormData, dispatch: Dispatch<any>, props: IOwnProps) => void;
}*/


const Form: React.SFC<AllProps> = (props: AllProps) => {

  return (
    <form onSubmit={props.handleSubmit}>
      <div>
        <label>{Organisation.orgName}</label>
        <FormField
          name={Organisation.orgName}
          component="input"
          type="text"
          placeholder={Organisation.orgName}
        />
      </div>
      <div>
        <label>{Organisation.identifier}</label>
        <FormField
          name={Organisation.identifier}
          component="input"
          type="text"
          placeholder={Organisation.identifier}
        />
      </div>
      <div>
        <label>{Organisation.type}</label>
        <FormField
          name={Organisation.type}
          component="input"
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

export const OrgForm = reduxForm<IFormData, IOwnProps>({
  form: "dummy-form"
})(Form)
