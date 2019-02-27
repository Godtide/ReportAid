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
import { OrganisationExpenditureProps } from '../../../store/IATI/IATIWriter/organisations/types'
import { FormData } from '../../../store/helpers/forms/types'

import { setFormFunctions } from '../../../store/helpers/forms/actions'
import { initialise } from '../../../store/IATI/IATIWriter/organisations/actions'
import { newKey } from '../../../store/helpers/keys/actions'
import { setOrganisationExpenditure } from '../../../store/IATI/IATIWriter/organisations/organisationExpenditure/actions'

import { FormikDatePicker } from '../../../components/io/datePicker'
import { OrganisationsPicker } from '../../../components/io/organisationsPicker'
import { OrganisationPicker } from '../../../components/io/organisationPicker'
import { OrgPicker } from '../../../components/io/orgPicker'
import { TransactionHelper } from '../../io/transactionHelper'

import { OrganisationExpenditure as OrganisationExpenditureStrings } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  organisationsRef: Yup
    .string()
    .required('Required'),
  organisationRef: Yup
    .string()
    .required('Required'),
  expenditureLine: Yup
    .string()
    .required('Required'),
  value: Yup
    .number()
    .required('Required'),
  status: Yup
    .number()
    .required('Required')
})

const StartDatePickerProps = {
  day: {
    name: 'startDay',
    label: OrganisationExpenditureStrings.expenditureStartDay
  },
  month: {
    name: 'startMonth',
    label: OrganisationExpenditureStrings.expenditureStartMonth
  },
  year: {
    name: 'startYear',
    label: OrganisationExpenditureStrings.expenditureStartYear
  }
}

const EndDatePickerProps = {
  day: {
    name: 'endDay',
    label: OrganisationExpenditureStrings.expenditureEndDay
  },
  month: {
    name: 'endMonth',
    label: OrganisationExpenditureStrings.expenditureEndMonth
  },
  year: {
    name: 'endYear',
    label: OrganisationExpenditureStrings.expenditureEndYear
  }
}

interface OrganisationExpenditureKeyProps {
  organisationsRef: string
  organisationRef: string
}

export interface OrganisationExpenditureDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
  setFormFunctions: (formProps: FormData) => void
}

type OrganisationExpenditureFormProps = WithStyles<typeof styles> & OrganisationExpenditureKeyProps & OrganisationExpenditureDispatchProps

export class OrganisationExpenditureForm extends React.Component<OrganisationExpenditureFormProps> {

  constructor (props: OrganisationExpenditureFormProps) {
    super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  handleSubmit = (values: OrganisationExpenditureProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.initialise()
    this.props.handleSubmit(values)
  }

  render() {

    return (
      <div>
        <h2>{OrganisationExpenditureStrings.headingOrganisationExpenditureWriter}</h2>
        <div>
          <Formik
            initialValues={ {organisationsRef: "",
                             organisationRef: "",
                             expenditureRef: "",
                             expenditureLine: "",
                             value: 0,
                             status: 1,
                             startDay: 0,
                             startMonth: 0,
                             startYear: 0,
                             endDay: 0,
                             endMonth: 0,
                             endYear: 0
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: OrganisationExpenditureProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrganisationExpenditureProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrganisationsPicker
                    setValue={formProps.setFieldValue}
                    name='organisationsRef'
                    label={OrganisationExpenditureStrings.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
                  <OrganisationPicker
                    setValue={formProps.setFieldValue}
                    name='organisationRef'
                    label={OrganisationExpenditureStrings.organisationReference}
                  />
                  <ErrorMessage name='organisationRef' />
                  <Field
                    name='expenditureLine'
                    label={OrganisationExpenditureStrings.expenditureLine}
                    component={TextField}
                  />
                  <ErrorMessage name='expenditureLine' />
                  <Field
                    name='value'
                    label={OrganisationExpenditureStrings.value}
                    component={TextField}
                  />
                  <ErrorMessage name='value' />
                  <Field
                    name="status"
                    label={OrganisationExpenditureStrings.status}
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
        <TransactionHelper/>
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationExpenditureKeyProps => {
  return {
    organisationsRef: state.keys.data.organisations,
    organisationRef: state.keys.data.organisation
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationExpenditureDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisationExpenditure(ownProps)),
    initialise: () => dispatch(initialise()),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const OrganisationExpenditure = withTheme(withStyles(styles)(connect<OrganisationExpenditureKeyProps, OrganisationExpenditureDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrganisationExpenditureForm)))
