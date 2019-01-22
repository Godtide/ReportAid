import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'

import { ApplicationState } from '../../../store'
import { ActionProps, TxData } from '../../../store/types'
import { IATIOrgReportBudgetProps, OrgReportBudgetProps} from '../../../store/IATI/types'

import { getOrgReports } from '../../../store/IATI/IATIReader/organisationReports/actions'
import { OrgReportData } from '../../../store/IATI/IATIReader/organisationReports/types'

import { setOrganisationReportBudget } from '../../../store/IATI/IATIWriter/organisationReportBudgets/actions'

import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import { Select, TextField } from "material-ui-formik-components"
//import { Date } from 'formik-material-ui'

import { OrganisationReportBudget, Transaction } from '../../../utils/strings'
import { Helpers } from '../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  reportRef: Yup
    .string()
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

interface BudgetProps {
  tx: TxData,
  orgReports: OrgReportData
}

export interface OrgReportBudgetsDispatchProps {
  handleSubmit: (values: any) => void
  getOrgReports: () => void
}

type OrgReportBudgetsFormProps = WithStyles<typeof styles> & BudgetProps & OrgReportBudgetsDispatchProps

export class OrgReportBudgetsForm extends React.Component<OrgReportBudgetsFormProps> {

  state = {
    txKey: '',
    txSummary: '',
    toggleSubmitting: false,
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgReportBudgetsFormProps) {
   super(props)
  }

  componentDidMount() {
    this.props.getOrgReports()
  }

  componentDidUpdate(previousProps: OrgReportBudgetsFormProps) {
    //console.log(this.props.tx)
    if(previousProps.tx != this.props.tx) {
      const submitting = !this.state.toggleSubmitting
      let txKey = ''
      let txSummary = `${Transaction.fail}`
      const txKeys = Object.keys(this.props.tx)
      if (txKeys.length > 0 ) {
        txKey = Object.keys(this.props.tx)[0]
        txSummary = `${Transaction.success}`
      }
      this.setState({txKey: txKey, txSummary: txSummary, toggleSubmitting: submitting})
      this.state.submitFunc(submitting)
      this.state.resetFunc()
    }
  }

  handleSubmit = (values: OrgReportBudgetProps, setSubmitting: Function, reset: Function) => {
    //console.log('Values: ', values)
    const submitting = !this.state.toggleSubmitting
    this.setState({txKey: '', txSummary: '', toggleSubmitting: submitting, submitFunc: setSubmitting, resetFunc: reset})
    setSubmitting(submitting)
    this.props.handleSubmit(values)
  }

  render() {

    let reportRefs: any[] = [{ value: "", label: "" }]
    Object.keys(this.props.orgReports).forEach((orgKey) => {
      //console.log(orgKey)
      const values = Object.values(this.props.orgReports[orgKey])
      //console.log(values)
      Object.keys(values[1]).forEach((reportKey) => {
        //console.log('Key: ', reportKey)
        reportRefs.push({ value: reportKey, label: reportKey })
      })
    })

    let status: any[] = [{ value: 0, label: "" }]
    Helpers.budgetStatus.forEach( (value: any) => {
      //console.log(value, value.code)
      status.push({ value: value.code, label: value.name })
    })

    //console.log(reportRefs)

    let dayRefs: any[] = []
    Array.from({ length: 31 }, (v: number, i: number) => {
      const value = ++i
      dayRefs.push({ value: value, label: value.toString() })
    })

    let monthRefs: any[] = []
    Array.from({ length: 12 }, (v: number, i: number) => {
      const value = ++i
      monthRefs.push({ value: value, label: value.toString() })
    })

    const startYear = 1990
    const stopYear = 2030
    const step = 1
    let yearRefs: any[] = []
    Array.from({ length: (stopYear - startYear) / step }, (_, i: number) => {
      const year = startYear + (i * step)
      yearRefs.push({ value: year, label: year.toString() })
    })

    return (
      <div>
        <h2>{OrganisationReportBudget.headingOrgReportBudgetWriter}</h2>
        <div>
          <Formik
            initialValues={ {reportRef: reportRefs[0].value,
                             budgetLine: "",
                             value: 0,
                             status: status[0].values,
                             startDay: dayRefs[0].value,
                             startMonth: monthRefs[0].value,
                             startYear: yearRefs[0].value,
                             endDay: dayRefs[0].value,
                             endMonth: monthRefs[0].value,
                             endYear: yearRefs[0].value,
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: OrgReportBudgetProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgReportBudgetProps>) => (
              <Form>
                <Field
                  name="reportRef"
                  label={OrganisationReportBudget.reportReference}
                  component={Select}
                  options={reportRefs}
                />
                <ErrorMessage name='reportRef' />
                <br />
                <Field
                  name='budgetLine'
                  label={OrganisationReportBudget.budgetLine}
                  component={TextField}
                />
                <ErrorMessage name='budgetLine' />
                <br />
                <Field
                  name='value'
                  label={OrganisationReportBudget.value}
                  component={TextField}
                />
                <ErrorMessage name='value' />
                <Field
                  name='status'
                  label={OrganisationReportBudget.status}
                  component={Select}
                  options={status}
                />
                <ErrorMessage name='status' />
                <Field
                  name='startDay'
                  label={OrganisationReportBudget.budgetStartDay}
                  component={Select}
                  options={dayRefs}
                />
                <ErrorMessage name='startDay' />
                <Field
                  name='startMonth'
                  label={OrganisationReportBudget.budgetStartMonth}
                  component={Select}
                  options={monthRefs}
                />
                <ErrorMessage name='startMonth' />
                <Field
                  name='startYear'
                  label={OrganisationReportBudget.budgetStartYear}
                  component={Select}
                  options={yearRefs}
                />
                <ErrorMessage name='startYear' />
                <br />
                <Field
                  name='endDay'
                  label={OrganisationReportBudget.budgetEndDay}
                  component={Select}
                  options={dayRefs}
                />
                <ErrorMessage name='endDay' />
                <Field
                  name='endMonth'
                  label={OrganisationReportBudget.budgetEndMonth}
                  component={Select}
                  options={monthRefs}
                />
                <ErrorMessage name='endMonth' />
                <Field
                  name='endYear'
                  label={OrganisationReportBudget.budgetEndYear}
                  component={Select}
                  options={yearRefs}
                />
                <ErrorMessage name='endYear' />
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
        <h3>{Transaction.heading}</h3>
        <p>
          <b>{Transaction.summary}</b>: {this.state.txSummary}<br />
          <b>{Transaction.key}</b>: {this.state.txKey}
        </p>
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): BudgetProps => {
  //console.log(state.orgReader)
  return {
    tx: state.orgReportBudgetsForm.data,
    orgReports: state.orgReportsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgReportBudgetsDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisationReportBudget(ownProps)),
    getOrgReports: () => dispatch(getOrgReports())
  }
}

export const OrganisationReportBudgets = withTheme(withStyles(styles)(connect<BudgetProps, OrgReportBudgetsDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgReportBudgetsForm)))
