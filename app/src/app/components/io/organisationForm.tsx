import * as React from 'react'
import { Formik, FormikProps, Form, Field, FieldProps } from 'formik'
import { Organisation } from '../../utils/strings'
import { OrganisationProps } from '../../store/IATIWriter/organisationWriter/types'

export interface DispatchProps {
  handleSubmit: (values: any) => void
}

type AllProps = OrganisationProps & DispatchProps

export const OrgForm: React.SFC<AllProps> = (props: AllProps) => {
  return (
    <div>
      <Formik
        initialValues={ {name: props.name, reference: props.reference, type: props.type} }
        onSubmit={(values: OrganisationProps) => props.handleSubmit(values)}
        render={() => (
          <Form>
            <Field
              name='name'
              render={({ field, form }: FieldProps<OrganisationProps>) => (
                <div>
                  <span>{Organisation.orgName}: </span>
                  <input
                    type="text"
                    value={props.name}
                    placeholder={props.name}
                    {...field}
                  />
                  {form.touched.name &&
                    form.errors.name &&
                    form.errors.name}
                </div>
              )}
            />
            <Field
              name='reference'
              render={({ field, form }: FieldProps<OrganisationProps>) => (
                <div>
                  <span>{Organisation.reference}: </span>
                  <input
                    type="text"
                    value={props.reference}
                    placeholder={props.reference}
                    {...field}
                  />
                  {form.touched.reference &&
                    form.errors.reference &&
                    form.errors.reference}
                </div>
              )}
            />
            <Field
              name='type'
              render={({ field, form }: FieldProps<OrganisationProps>) => (
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
