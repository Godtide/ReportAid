import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import FormControl from '@material-ui/core/FormControl'
import { Select, TextField } from "material-ui-formik-components"
//import { Date } from 'formik-material-ui'

import { ApplicationState } from '../../../store'
import { ActionProps, TxData } from '../../../store/types'
import { OrganisationExpenditureProps } from '../../../store/IATI/types'

import { setOrganisationExpenditure } from '../../../store/IATI/IATIWriter/organisations/organisationExpenditure/actions'

import { FormikDatePicker } from '../../../components/io/datePicker'
import { OrganisationsPicker } from '../../../components/io/organisationsPicker'
import { OrganisationPicker } from '../../../components/io/orgPicker'
import { OrgPicker } from '../../../components/io/reportPicker'
import { TransactionHelper } from '../../io/transactionHelper'

import { OrganisationExpenditure as OrganisationExpenditure } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  organisations: Yup
    .object()
    .required('Required'),
  expenditureLine: Yup
    .string()
    .required('Required'),
  value: Yup
    .number()
    .required('Required'),
  status: Yup
    .number()
    .required('Required'),
  startDay: Yup
    .number()
    .required('Required'),
  startMonth: Yup
    .number()
    .required('Required'),
  startYear: Yup
    .number()
    .required('Required'),
  endDay: Yup
    .number()
    .required('Required'),
  endMonth: Yup
    .number()
    .required('Required'),
  endYear: Yup
    .number()
    .required('Required'),
})

const StartDatePickerProps = {
  day: {
    name: 'startDay',
    label: OrganisationExpenditure.expenditureStartDay
  },
  month: {
    name: 'startMonth',
    label: OrganisationExpenditure.expenditureStartMonth
  },
  year: {
    name: 'startYear',
    label: OrganisationExpenditure.expenditureStartYear
  }
}

const EndDatePickerProps = {
  day: {
    name: 'endDay',
    label: OrganisationExpenditure.expenditureEndDay
  },
  month: {
    name: 'endMonth',
    label: OrganisationExpenditure.expenditureEndMonth
  },
  year: {
    name: 'endYear',
    label: OrganisationExpenditure.expenditureEndYear
  }
}

export interface OrganisationExpenditureDispatchProps {
  handleSubmit: (values: any) => void
}

type OrganisationExpenditureFormProps = WithStyles<typeof styles> & OrganisationExpenditureDispatchProps

export class OrganisationExpenditureForm extends React.Component<OrganisationExpenditureFormProps> {

  state = {
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrganisationExpenditureFormProps) {
   super(props)
  }

  handleSubmit = (values: OrganisationExpenditureProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  render() {

    return (
      <div>
        <h2>{OrganisationExpenditure.headingOrganisationExpenditureWriter}</h2>
        <div>
          <Formik
            initialValues={ {organisations: {} as Props,
                             expenditureLine: "",
                             value: 0,
                             status: 1,
                             startDay: 1,
                             startMonth: 1,
                             startYear: 2000,
                             endDay: 1,
                             endMonth: 1,
                             endYear: 2000
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: OrganisationExpenditureProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrganisationExpenditureProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrganisationsPicker name='organisations' label={OrganisationExpenditure.organisationsReference} />
                  <Field
                    name='expenditureLine'
                    label={OrganisationExpenditure.expenditureLine}
                    component={TextField}
                  />
                  <ErrorMessage name='expenditureLine' />
                  <Field
                    name='value'
                    label={OrganisationExpenditure.value}
                    component={TextField}
                  />
                  <ErrorMessage name='value' />
                  <Field
                    name="status"
                    label={OrganisationExpenditure.status}
                    component={Select}
                    options={Helpers.financeStatus}
                  />
                  <ErrorMessage name='status' />
                  <FormikDatePicker dates={StartDatePickerProps} />
                  <FormikDatePicker dates={EndDatePickerProps} />
                  <br />
                  {formProps.isSubmitting && <LinearProgress />}
                  <br />
                  <Button type='submit' variant="raised" color="primary" disabled={formProps.isSubmitting}>
                    Submit
                  </Button>
                </FormControl>
              </Form>
            )}
          />
        </div>
        <TransactionHelper
          submitFunc={this.state.submitFunc}
          resetFunc={this.state.resetFunc}
        />
      </div>
    )
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationExpenditureDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisationExpenditure(ownProps))
  }
}

export const OrganisationExpenditure = withTheme(withStyles(styles)(connect<void, OrganisationExpenditureDispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(OrganisationExpenditureForm)))
