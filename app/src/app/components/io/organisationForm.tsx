import * as React from 'react'
import { Formik, FormikProps, Form, Field, FieldProps } from 'formik'
import { Organisation } from '../../utils/strings'

interface OwnProps {
  orgName: string
  identifier: string
  type: string
}

export interface DispatchProps {
  handleSubmit: (values: any) => void
}

//type AllProps = OrgFormProps & InjectedFormProps<{}, OrgFormProps>

type AllProps = OwnProps & DispatchProps

export const OrgForm: React.SFC<AllProps> = (props: AllProps) => {
  return (
    <div>
      <Formik
        initialValues={ {orgName: '', identifier: '', type: ''} }
        onSubmit={(values: OwnProps) => props.handleSubmit(values)}
        render={(values: FormikProps<OwnProps>) => (
          <Form>
            <Field
              name='orgName'
              render={({ field, form }: FieldProps<OwnProps>) => (
                <div>
                  <span>{Organisation.orgName}: </span>
                  <input
                    type="text"
                    value={props.orgName}
                    placeholder={props.orgName}
                    {...field}
                  />
                  {form.touched.orgName &&
                    form.errors.orgName &&
                    form.errors.orgName}
                </div>
              )}
            />
            <Field
              name='identifier'
              render={({ field, form }: FieldProps<OwnProps>) => (
                <div>
                  <span>{Organisation.identifier}: </span>
                  <input
                    type="text"
                    value={props.identifier}
                    placeholder={props.identifier}
                    {...field}
                  />
                  {form.touched.identifier &&
                    form.errors.identifier &&
                    form.errors.identifier}
                </div>
              )}
            />
            <Field
              name='type'
              render={({ field, form }: FieldProps<OwnProps>) => (
                <div>
                  <span>{Organisation.type}: </span>
                  <input
                    type="text"
                    value={props.type}
                    placeholder={props.type}
                    {...field}  />
                  {form.touched.type &&
                    form.errors.type &&
                    form.errors.type}
                </div>
              )}
            />
            <button type="submit">
              Submit
            </button>
          </Form>
        )}
      />
    </div>
  )
}
