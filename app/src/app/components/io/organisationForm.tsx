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
        initialValues={{ orgName: '', identifier: '', type: '' }}
        onSubmit={(values: OwnProps) => props.handleSubmit(values)}
        render={(values: FormikProps<OwnProps>) => (
          <Form>
            <Field
              name={Organisation.orgName}
              render={({ field, form }: FieldProps<OwnProps>) => (
                <div>
                  <input type="text" {...field} placeholder={props.orgName} />
                  {form.touched.orgName &&
                    form.errors.orgName &&
                    form.errors.orgName}
                </div>
              )}
            />
            <Field
              name={Organisation.identifier}
              render={({ field, form }: FieldProps<OwnProps>) => (
                <div>
                  <input type="text" {...field} placeholder={props.identifier} />
                  {form.touched.identifier &&
                    form.errors.identifier &&
                    form.errors.identifier}
                </div>
              )}
            />
            <Field
              name={Organisation.type}
              render={({ field, form }: FieldProps<OwnProps>) => (
                <div>
                  <input type="text" {...field} placeholder={props.type} />
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
