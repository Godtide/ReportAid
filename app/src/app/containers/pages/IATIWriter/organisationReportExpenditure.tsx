import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'

import { ApplicationState } from '../../../store'
import { ActionProps, TxData } from '../../../store/types'
import { IATIOrgReportExpenditureProps, OrgReportExpenditureProps} from '../../../store/IATI/types'

import { getOrgReports } from '../../../store/IATI/IATIReader/organisationReports/actions'
import { OrgReportData } from '../../../store/IATI/IATIReader/organisationReports/types'

import { setOrganisationReportExpenditure } from '../../../store/IATI/IATIWriter/organisationReportExpenditure/actions'

import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import Grid from '@material-ui/core/Grid'
import FormControl from '@material-ui/core/FormControl'
import { Select, TextField } from "material-ui-formik-components"
//import { Date } from 'formik-material-ui'

import { OrganisationReportExpenditure as OrgReportExpenditure, Transaction } from '../../../utils/strings'
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

interface ExpenditureProps {
  tx: TxData,
  orgReports: OrgReportData
}

export interface OrgReportExpenditureDispatchProps {
  handleSubmit: (values: any) => void
  getOrgReports: () => void
}

type OrgReportExpenditureFormProps = WithStyles<typeof styles> & ExpenditureProps & OrgReportExpenditureDispatchProps

export class OrgReportExpenditureForm extends React.Component<OrgReportExpenditureFormProps> {

  state = {
    txKey: '',
    txSummary: '',
    toggleSubmitting: false,
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgReportExpenditureFormProps) {
   super(props)
  }

  componentDidMount() {
    this.props.getOrgReports()
  }

  componentDidUpdate(previousProps: OrgReportExpenditureFormProps) {
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

  handleSubmit = (values: OrgReportExpenditureProps, setSubmitting: Function, reset: Function) => {
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

    let status: any[] = []
    Helpers.financeStatus.forEach( (value: any) => {
      //console.log(value)
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
        <h2>{OrgReportExpenditure.headingOrgReportExpenditureWriter}</h2>
        <div>
          <Formik
            initialValues={ {reportRef: reportRefs[0].value,
                             expenditureLine: "",
                             value: 0,
                             status: status[0].value,
                             startDay: dayRefs[0].value,
                             startMonth: monthRefs[0].value,
                             startYear: yearRefs[0].value,
                             endDay: dayRefs[0].value,
                             endMonth: monthRefs[0].value,
                             endYear: yearRefs[0].value,
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: OrgReportExpenditureProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgReportExpenditureProps>) => (
              <Form>
                <FormControl fullWidth={false}>
                  <Field
                    name="reportRef"
                    label={OrgReportExpenditure.reportReference}
                    component={Select}
                    options={reportRefs}
                  />
                  <ErrorMessage name='reportRef' />
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
                    name='status'
                    label={OrgReportExpenditure.status}
                    component={Select}
                    options={status}
                  />
                  <ErrorMessage name='status' />
                  <Grid container>
                    <Grid item xs={12} sm={3}>
                      <Field
                        name='startDay'
                        label={OrgReportExpenditure.expenditureStartDay}
                        //style={{ width: '10%' }}
                        component={Select}
                        options={dayRefs}
                      />
                      <ErrorMessage name='startDay' />
                    </Grid>
                    <Grid item xs={12} sm={3}>
                      <Field
                        name='startMonth'
                        label={OrgReportExpenditure.expenditureStartMonth}
                        //style={{ width: '10%' }}
                        component={Select}
                        options={monthRefs}
                      />
                      <ErrorMessage name='startMonth' />
                    </Grid>
                    <Grid item xs={12} sm={6}>
                      <Field
                        name='startYear'
                        label={OrgReportExpenditure.expenditureStartYear}
                        //style={{ width: '10%' }}
                        component={Select}
                        options={yearRefs}
                      />
                      <ErrorMessage name='startYear' />
                    </Grid>
                  </Grid>
                  <Grid container>
                    <Grid item xs={12} sm={3}>
                      <Field
                        name='endDay'
                        label={OrgReportExpenditure.expenditureEndDay}
                        //style={{ width: '10%' }}
                        component={Select}
                        options={dayRefs}
                      />
                      <ErrorMessage name='endDay' />
                    </Grid>
                    <Grid item xs={12} sm={3}>
                      <Field
                        name='endMonth'
                        label={OrgReportExpenditure.expenditureEndMonth}
                        //style={{ width: '10%' }}
                        component={Select}
                        options={monthRefs}
                      />
                      <ErrorMessage name='endMonth' />
                    </Grid>
                    <Grid item xs={12} sm={6}>
                      <Field
                        name='endYear'
                        label={OrgReportExpenditure.expenditureEndYear}
                        //style={{ width: '10%' }}
                        component={Select}
                        options={yearRefs}
                      />
                      <ErrorMessage name='endYear' />
                    </Grid>
                  </Grid>
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

const mapStateToProps = (state: ApplicationState): ExpenditureProps => {
  //console.log(state.orgReader)
  return {
    tx: state.orgReportExpenditureForm.data,
    orgReports: state.orgReportsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgReportExpenditureDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisationReportExpenditure(ownProps)),
    getOrgReports: () => dispatch(getOrgReports())
  }
}

export const OrganisationReportExpenditure = withTheme(withStyles(styles)(connect<ExpenditureProps, OrgReportExpenditureDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgReportExpenditureForm)))
