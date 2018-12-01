import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { OrganisationProps } from '../../../store/IATI/types'
import { TxData } from '../../../store/IATI/IATIWriter/organisationWriter/types'

import { setOrganisation } from '../../../store/IATI/IATIWriter/organisationWriter/actions'

import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import { TextField } from 'formik-material-ui'

import { Organisation } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgTXProps {
  tx: TxData
}

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

type OrgWriterFormProps = WithStyles<typeof styles> & OrgTXProps & OrgDispatchProps

export class OrgForm extends React.Component<OrgWriterFormProps> {

  state = {
    txKey: '',
    toggleSubmitting: false,
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgWriterFormProps) {
   super(props)
  }

  componentDidUpdate(previousProps: OrgWriterFormProps) {
    //console.log(this.props.tx)
    if(previousProps.tx != this.props.tx) {
      //console.log('Prvious ', previousProps.tx, 'New ', this.props.tx, 'SUBMITTING ', this.state.toggleSubmitting)
      const submitting = !this.state.toggleSubmitting
      const txKey = Object.keys(this.props.tx)[0]
      this.setState({txKey: txKey, toggleSubmitting: submitting})
      console.log(submitting)
      this.state.submitFunc(submitting)
      this.state.resetFunc()
    }
  }

  handleSubmit = (values: OrganisationProps, setSubmitting: Function, reset: Function) => {
    const submitting = !this.state.toggleSubmitting
    this.setState({txKey: '', toggleSubmitting: submitting, submitFunc: setSubmitting, resetFunc: reset})
    setSubmitting(submitting)
    this.props.handleSubmit(values)
  }

  render() {

    return (
      <div>
        <h2>{Organisation.headingOrgWriter}</h2>

        <div>
          <Formik
            initialValues={ {name: '', reference: '', type: ''} }
            validationSchema={organisationSchema}
            onSubmit={(values: OrganisationProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<any>) => (
              <Form>
                <Field name='name' label={Organisation.orgName} component={TextField} />
                <ErrorMessage name='name' />
                <br />
                <Field name='reference' label={Organisation.reference} component={TextField} />
                <ErrorMessage name='reference' />
                <br />
                <Field name='type' label={Organisation.type} component={TextField} />
                <ErrorMessage name='type' />
                <br />
                {formProps.isSubmitting && <LinearProgress />}
                <br />
                <Button type='submit' variant="raised" color="primary" disabled={formProps.isSubmitting}>
                  Submit
                </Button>
              </Form>
            )}
          />
        </div>
        <hr />
        <h3>{Organisation.orgTXHeader}</h3>
        <p>
          <b>Transaction Key</b>: {this.state.txKey}
        </p>
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgTXProps => {
  return {
    tx: state.orgForm.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisation(ownProps))
  }
}

export const OrganisationWriter = withTheme(withStyles(styles)(connect<OrgTXProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgForm)))
