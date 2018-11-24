import * as React from 'react'
import { Formik, Form, Field, FormikProps } from 'formik'
import { Organisation } from '../../utils/strings'
import { OrganisationProps } from '../../store/IATIWriter/organisationWriter/types'
import {LinearProgress} from '@material-ui/core'
import Button from '@material-ui/core/Button'
import {TextField} from 'formik-material-ui'

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
              label={Organisation.orgName}
              component={TextField}
            />
            <br />
            <Field
              name='reference'
              label={Organisation.reference}
              component={TextField}
            />
            <br />
            <Field
              name='type'
              label={Organisation.type}
              component={TextField}
            />
            <br />
            <br />
            <Button
              type='submit'
              variant="raised"
              color="primary"
            >
              Submit
            </Button>
          </Form>
        )}
      />
    </div>
  )
}
