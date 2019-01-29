import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import FormControl from '@material-ui/core/FormControl'
import { Select, TextField } from "material-ui-formik-components"

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrgReportRegionBudgetProps, OrgReportRegionBudgetProps, ReportProps } from '../../../store/IATI/types'

import { setRegionBudget } from '../../../store/IATI/IATIWriter/organisationReportRegionBudgets/actions'

import { FormikDatePicker } from '../../../components/io/datePicker'
import { OrgReportPicker } from '../../../components/io/reportPicker'
import { TransactionHelper, TransactionTypes } from '../../io/transactionHelper'

import { OrganisationReportRegionBudget } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  reportRef: Yup
    .string()
    .required('Required'),
  regionRef: Yup
    .number()
    .required('Required'),
  budgetLine: Yup
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
    label: OrganisationReportRegionBudget.budgetStartDay
  },
  month: {
    name: 'startMonth',
    label: OrganisationReportRegionBudget.budgetStartMonth
  },
  year: {
    name: 'startYear',
    label: OrganisationReportRegionBudget.budgetStartYear
  }
}

const EndDatePickerProps = {
  day: {
    name: 'endDay',
    label: OrganisationReportRegionBudget.budgetEndDay
  },
  month: {
    name: 'endMonth',
    label: OrganisationReportRegionBudget.budgetEndMonth
  },
  year: {
    name: 'endYear',
    label: OrganisationReportRegionBudget.budgetEndYear
  }
}

interface OrgReportRegionBudgetsDispatchProps {
  handleSubmit: (values: any) => void
}

type OrgReportRegionBudgetsFormProps = WithStyles<typeof styles> & OrgReportRegionBudgetsDispatchProps

export class OrgReportRegionBudgetsForm extends React.Component<OrgReportRegionBudgetsFormProps> {

  state = {
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgReportRegionBudgetsFormProps) {
   super(props)
  }

  handleSubmit = (values: OrgReportRegionBudgetProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  render() {

    return (
      <div>
        <h2>{OrganisationReportRegionBudget.headingOrgReportRegionBudgetWriter}</h2>
        <div>
          <Formik
            initialValues={ {report: {} as ReportProps,
                             regionRef: 1,
                             budgetLine: "",
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
            onSubmit={(values: OrgReportRegionBudgetProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgReportRegionBudgetProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrgReportPicker name='reportRef' label={OrganisationReportRegionBudget.reportReference} />
                  <Field
                    name='budgetLine'
                    label={OrganisationReportRegionBudget.budgetLine}
                    component={TextField}
                  />
                  <ErrorMessage name='budgetLine' />
                  <Field
                    name="regionRef"
                    label={OrganisationReportRegionBudget.regionReference}
                    component={Select}
                    options={Helpers.oecdDacCrs}
                  />
                  <ErrorMessage name='regionRef' />
                  <Field
                    name='value'
                    label={OrganisationReportRegionBudget.value}
                    component={TextField}
                  />
                  <ErrorMessage name='value' />
                  <Field
                    name="status"
                    label={OrganisationReportRegionBudget.status}
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
          type={TransactionTypes.ORGREPORTREGIONBUDGET}
          submitFunc={this.state.submitFunc}
          resetFunc={this.state.resetFunc}
        />
      </div>
    )
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgReportRegionBudgetsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setRegionBudget(ownProps))
  }
}

export const OrganisationReportRegionBudgets = withTheme(withStyles(styles)(connect<void, OrgReportRegionBudgetsDispatchProps, {}, ApplicationState>(
  null,
  mapDispatchToProps
)(OrgReportRegionBudgetsForm)))
