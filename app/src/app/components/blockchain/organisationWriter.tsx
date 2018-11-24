import * as React from 'react'
import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'
import { Organisation } from '../../utils/strings'
import { OrganisationProps } from '../../store/IATIWriter/organisationWriter/types'
import {LinearProgress} from '@material-ui/core'
import Button from '@material-ui/core/Button'
import {TextField} from 'formik-material-ui'
import { WithStyles } from '@material-ui/core/styles/withStyles'
import { styles } from '../../styles/theme'

export interface OrgDispatchProps {
  handleSubmit: (values: any) => void
}

const organisationSchema = Yup.object().shape({
  name: Yup
    .string()
    .required('Required'),
  reference: Yup
    .string()
    .required('Required'),
  type: Yup
    .string()
    .required('Required'),
})

type OrgFormProps = OrganisationProps & OrgDispatchProps

const OrgForm: React.SFC<OrgFormProps> = (props: OrgFormProps) => {
  return (
    <div>
      <Formik
        initialValues={ {name: props.name, reference: props.reference, type: props.type} }
        validationSchema={organisationSchema}
        onSubmit={(values: OrganisationProps, actions: any) => {
          props.handleSubmit(values)
          actions.setSubmitting(false)
          actions.resetForm()
        }}
        render={({isSubmitting}: FormikProps<any>) => (
          <Form>
            <Field
              name='name'
              label={Organisation.orgName}
              component={TextField}
            />
            <ErrorMessage
              name='name'
            />
            <br />
            <Field
              name='reference'
              label={Organisation.reference}
              component={TextField}
            />
            <ErrorMessage
              name='reference'
            />
            <br />
            <Field
              name='type'
              label={Organisation.type}
              component={TextField}
            />
            <ErrorMessage
              name='type'
            />
            <br />
            {isSubmitting && <LinearProgress />}
            <br />
            <Button
              type='submit'
              variant="raised"
              color="primary"
              disabled={isSubmitting}
            >
              Submit
            </Button>
          </Form>
        )}
      />
    </div>
  )
}

type OrgWriterProps = WithStyles<typeof styles> & OrgDispatchProps

export const OrgWriter: React.SFC<OrgWriterProps> = (props: OrgWriterProps) => {

  return (
    <div>
      <h2>{Organisation.headingOrgWriter}</h2>
      <OrgForm
        handleSubmit={(values: any) => props.handleSubmit(values)}
        name = ''
        reference = ''
        type = ''
      />
    </div>
  )
}
