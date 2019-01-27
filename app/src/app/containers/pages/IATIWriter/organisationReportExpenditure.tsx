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
import { IATIOrgReportExpenditureProps, OrgReportExpenditureProps} from '../../../store/IATI/types'

import { setOrganisationReportExpenditure } from '../../../store/IATI/IATIWriter/organisationReportExpenditure/actions'

import { FormikDatePicker } from '../../../components/io/datePicker'
import { OrganisationPicker } from '../../../components/io/orgPicker'
import { FormikStatusPicker } from '../../../components/io/statusPicker'
import { OrgReportPicker } from '../../../components/io/reportPicker'
import { TransactionHelper, TransactionTypes } from '../../io/transactionHelper'

import { OrganisationReportExpenditure as OrgReportExpenditure } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  reportRef: Yup
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
    label: OrgReportExpenditure.expenditureStartDay
  },
  month: {
    name: 'startMonth',
    label: OrgReportExpenditure.expenditureStartMonth
  },
  year: {
    name: 'startYear',
    label: OrgReportExpenditure.expenditureStartYear
  }
}

const EndDatePickerProps = {
  day: {
    name: 'endDay',
    label: OrgReportExpenditure.expenditureEndDay
  },
  month: {
    name: 'endMonth',
    label: OrgReportExpenditure.expenditureEndMonth
  },
  year: {
    name: 'endYear',
    label: OrgReportExpenditure.expenditureEndYear
  }
}

export interface OrgReportExpenditureDispatchProps {
  handleSubmit: (values: any) => void
}

type OrgReportExpenditureFormProps = WithStyles<typeof styles> & OrgReportExpenditureDispatchProps

export class OrgReportExpenditureForm extends React.Component<OrgReportExpenditureFormProps> {

  state = {
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgReportExpenditureFormProps) {
   super(props)
  }

  handleSubmit = (values: OrgReportExpenditureProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  render() {

    return (
      <div>
        <h2>{OrgReportExpenditure.headingOrgReportExpenditureWriter}</h2>
        <div>
          <Formik
            initialValues={ {reportRef: "",
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
            onSubmit={(values: OrgReportExpenditureProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgReportExpenditureProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrgReportPicker name='reportRef' label={OrgReportExpenditure.reportReference} />
                  <Field
                    name='expenditureLine'
                    label={OrgReportExpenditure.expenditureLine}
                    component={TextField}
                  />
                  <ErrorMessage name='expenditureLine' />
                  <Field
                    name='value'
                    label={OrgReportExpenditure.value}
                    component={TextField}
                  />
                  <ErrorMessage name='value' />
                  <Field
                    name="status"
                    label={OrgReportExpenditure.status}
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
          type={TransactionTypes.ORGREPORTEXPENDITURE}
          submitFunc={this.state.submitFunc}
          resetFunc={this.state.resetFunc}
        />
      </div>
    )
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgReportExpenditureDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisationReportExpenditure(ownProps))
  }
}

export const OrganisationReportExpenditure = withTheme(withStyles(styles)(connect<void, OrgReportExpenditureDispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(OrgReportExpenditureForm)))
