import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import FormControl from '@material-ui/core/FormControl'
//import { Select } from 'formik-material-ui'
import { TextField, Select } from "material-ui-formik-components"

import { ApplicationState } from '../../../../store'
import { ActionProps } from '../../../../store/types'
import { IATIActivityTransactionReport, ActivityTransactionProps } from '../../../../store/IATI/types'
import { FormData } from '../../../../store/helpers/forms/types'

import { setFormFunctions } from '../../../../store/helpers/forms/actions'
import { setActivityTransaction } from '../../../../store/IATI/IATIWriter/activities/activityTransactions/actions'

import { FormikDatePicker } from '../../../../components/io/datePicker'
import { ActivitiesPicker } from '../../../../components/io/activitiesPicker'
import { ActivityPicker } from '../../../../components/io/activityPicker'
import { OrgPicker } from '../../../../components/io/orgPicker'
import { TransactionHelper } from '../../../io/transactionHelper'

import { ActivityTransactions as ActivityTransactionStrings } from '../../../../utils/strings'
import { Helpers } from '../../../../utils/config'

const activityTransactionSchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required'),
  activityRef: Yup
    .string()
    .required('Required'),
  transactionRef: Yup
    .string()
    .required('Required'),
  transactionType: Yup
    .number()
    .required('Required'),
  disbursementChannel: Yup
    .number()
    .required('Required'),
  flowType: Yup
    .number(),
  tiedStatus: Yup
    .number(),
  financeType: Yup
    .number(),
  aidType: Yup
    .string(),
  day: Yup
    .number()
    .required('Required'),
  month: Yup
    .number()
    .required('Required'),
  year: Yup
    .number()
    .required('Required'),
  value: Yup
    .number()
    .required('Required'),
  valueDay: Yup
    .number()
    .required('Required'),
  valueMonth: Yup
    .number()
    .required('Required'),
  valueYear: Yup
    .number()
    .required('Required'),
  currency: Yup
    .string()
    .required('Required'),
  providerOrgType: Yup
    .number()
    .required('Required'),
  providerOrgRef: Yup
    .string()
    .required('Required'),
  providerActivityRef: Yup
    .string(),
  receiverOrgType: Yup
    .number()
    .required('Required'),
  receiverOrgRef: Yup
    .string()
    .required('Required'),
  receiverActivityRef: Yup
    .string(),
  sectorDacCode: Yup
    .number()
    .required('Required'),
  countryCode: Yup
    .string(),
  regionCode: Yup
    .string(),
  description: Yup
    .string()
    .required('Required')
})

const DatePickerProps = {
  day: {
    name: 'day',
    label: ActivityTransactionStrings.day
  },
  month: {
    name: 'month',
    label: ActivityTransactionStrings.month
  },
  year: {
    name: 'year',
    label: ActivityTransactionStrings.year
  }
}

const ValueDatePickerProps = {
  day: {
    name: 'valueDay',
    label: ActivityTransactionStrings.valueDay
  },
  month: {
    name: 'valueMonth',
    label: ActivityTransactionStrings.valueMonth
  },
  year: {
    name: 'valueYear',
    label: ActivityTransactionStrings.valueYear
  }
}

interface ActivityTransactionKeyProps {
  activitiesRef: string
  activityRef: string
  transactionRef: string
  transactions: IATIActivityTransactionReport
}

export interface ActivityTransactionDispatchProps {
  handleSubmit: (values: any) => void
  setFormFunctions: (formProps: FormData) => void
}

type ActivityTransactionFormProps = ActivityTransactionKeyProps & ActivityTransactionDispatchProps

export class ActivityTransactionForm extends React.Component<ActivityTransactionFormProps> {

  constructor (props: ActivityTransactionFormProps) {
   super(props)
  }

  handleSubmit = (values: ActivityTransactionProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  getData = (transactions: IATIActivityTransactionReport): ActivityTransactionProps => {

    let newTransaction: ActivityTransactionProps = {
      activitiesRef: this.props.activitiesRef,
      activityRef: this.props.activityRef,
      transactionRef: this.props.transactionRef,
      transactionType: 0,
      disbursementChannel: 0,
      flowType: 0,
      tiedStatus: 0,
      financeType: 0,
      aidType: "",
      day: 0,
      month: 0,
      year: 0,
      value: 0,
      valueDay: 0,
      valueMonth: 0,
      valueYear: 0,
      currency: "",
      providerOrgType: 0,
      providerOrgRef: "",
      providerActivityRef: "",
      receiverOrgType: 0,
      receiverOrgRef: "",
      receiverActivityRef: "",
      sectorDacCode: 0,
      countryCode: "",
      regionCode: "",
      description: ""
    }
    if ( typeof transactions.data != 'undefined' ) {
      let countryCode = transactions.data[0].territory
      let regionCode = countryCode
      if(!isNaN(parseInt(countryCode, 10))) {
        regionCode = ""
      } else {
        countryCode = ""
      }
      const thisDate = new Date(transactions.data[0].date)
      const valueDate = new Date(transactions.data[0].valueDate)
      newTransaction.transactionType = transactions.data[0].transactionType,
      newTransaction.disbursementChannel = transactions.data[0].disbursementChannel,
      newTransaction.flowType = transactions.data[0].flowType,
      newTransaction.tiedStatus = transactions.data[0].tiedStatus,
      newTransaction.financeType = transactions.data[0].financeType,
      newTransaction.aidType = transactions.data[0].aidType,
      newTransaction.day = thisDate.getDay(),
      newTransaction.month = thisDate.getMonth(),
      newTransaction.year = thisDate.getFullYear(),
      newTransaction.value = transactions.data[0].value,
      newTransaction.valueDay = valueDate.getDay(),
      newTransaction.valueMonth = valueDate.getMonth(),
      newTransaction.valueYear = valueDate.getFullYear(),
      newTransaction.currency = transactions.data[0].currency,
      newTransaction.providerOrgType = transactions.data[0].providerOrgType,
      newTransaction.providerOrgRef = transactions.data[0].providerOrgRef,
      newTransaction.providerActivityRef = transactions.data[0].providerActivityRef,
      newTransaction.receiverOrgType = transactions.data[0].receiverOrgType,
      newTransaction.receiverOrgRef = transactions.data[0].receiverOrgRef,
      newTransaction.receiverActivityRef = transactions.data[0].receiverActivityRef,
      newTransaction.sectorDacCode = transactions.data[0].sectorDacCode,
      newTransaction.countryCode = countryCode,
      newTransaction.regionCode = regionCode,
      newTransaction.description = transactions.data[0].description
    }

    return newTransaction
  }

  render() {

    const thisTransaction: ActivityTransactionProps = this.getData(this.props.transactions)

    return (
      <div>
        <h2>{ActivityTransactionStrings.headingActivityTransactionsWriter}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: this.props.activitiesRef,
                             activityRef: this.props.activityRef,
                             transactionRef: this.props.transactionRef,
                             transactionType: thisTransaction.transactionType,
                             disbursementChannel: thisTransaction.disbursementChannel,
                             flowType: thisTransaction.flowType,
                             tiedStatus: thisTransaction.tiedStatus,
                             financeType: thisTransaction.financeType,
                             aidType: thisTransaction.aidType,
                             day: thisTransaction.day,
                             month: thisTransaction.month,
                             year: thisTransaction.year,
                             value: thisTransaction.value,
                             valueDay: thisTransaction.valueDay,
                             valueMonth: thisTransaction.valueMonth,
                             valueYear: thisTransaction.valueYear,
                             currency: thisTransaction.currency,
                             providerOrgType: thisTransaction.providerOrgType,
                             providerOrgRef: thisTransaction.providerOrgRef,
                             providerActivityRef: thisTransaction.providerActivityRef,
                             receiverOrgType: thisTransaction.receiverOrgType,
                             receiverOrgRef: thisTransaction.receiverOrgRef,
                             receiverActivityRef: thisTransaction.receiverActivityRef,
                             sectorDacCode: thisTransaction.sectorDacCode,
                             countryCode: thisTransaction.countryCode,
                             regionCode: thisTransaction.regionCode,
                             description: thisTransaction.description
                            }}
            enableReinitialize={true}
            validationSchema={activityTransactionSchema}
            onSubmit={(values: ActivityTransactionProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivityTransactionProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <Field
                    readOnly
                    name='transactionRef'
                    value={this.props.transactionRef}
                    label={ActivityTransactionStrings.transactionRef}
                    component={TextField}
                  />
                  <ActivitiesPicker
                    setValue={formProps.setFieldValue}
                    name='activitiesRef'
                    label={ActivityTransactionStrings.activitiesReference}
                  />
                  <ErrorMessage name='activitiesRef' />
                  <ActivityPicker
                    setValue={formProps.setFieldValue}
                    name='activityRef'
                    label={ActivityTransactionStrings.activityReference}
                  />
                  <ErrorMessage name='activityRef' />
                  <Field
                    name="transactionType"
                    label={ActivityTransactionStrings.transactionType}
                    component={Select}
                    options={Helpers.transactionType}
                  />
                  <ErrorMessage name='transactionType' />
                  <Field
                    name="disbursementChannel"
                    label={ActivityTransactionStrings.disbursementChannel}
                    component={Select}
                    options={Helpers.disbursementChannel}
                  />
                  <ErrorMessage name='disbursementChannel' />
                  <FormikDatePicker dates={DatePickerProps} />
                  <Field
                    name='value'
                    label={ActivityTransactionStrings.value}
                    component={TextField}
                  />
                  <ErrorMessage name='value' />
                  <FormikDatePicker dates={ValueDatePickerProps} />
                  <Field
                    name="currency"
                    label={ActivityTransactionStrings.currency}
                    component={Select}
                    options={Helpers.currencyCodes}
                  />
                  <OrgPicker
                    setValue={formProps.setFieldValue}
                    name='providerOrgRef'
                    label={ActivityTransactionStrings.providerOrgRef}
                  />
                  <ErrorMessage name='providerOrgRef' />
                  <Field
                    name="providerOrgType"
                    label={ActivityTransactionStrings.providerOrgType}
                    component={Select}
                    options={Helpers.organisationCodes}
                  />
                  <ErrorMessage name='providerOrgType' />
                  <ActivityPicker
                    setValue={formProps.setFieldValue}
                    name='providerActivityRef'
                    label={ActivityTransactionStrings.providerActivityRef}
                  />
                  <ErrorMessage name='providerActivityRef' />
                  <OrgPicker
                    setValue={formProps.setFieldValue}
                    name='receiverOrgRef'
                    label={ActivityTransactionStrings.receiverOrgRef}
                  />
                  <ErrorMessage name='receiverOrgRef' />
                  <Field
                    name="receiverOrgType"
                    label={ActivityTransactionStrings.receiverOrgType}
                    component={Select}
                    options={Helpers.organisationCodes}
                  />
                  <ErrorMessage name='receiverOrgType' />
                  <ActivityPicker
                    setValue={formProps.setFieldValue}
                    name='receiverActivityRef'
                    label={ActivityTransactionStrings.receiverActivityRef}
                  />
                  <ErrorMessage name='receiverActivityRef' />

                  <Field
                    name="sectorDacCode"
                    label={ActivityTransactionStrings.sectorDacCode}
                    component={Select}
                    options={Helpers.DACSector}
                  />
                  <ErrorMessage name='sectorDacCode' />
                  <Field
                    name="countryCode"
                    label={ActivityTransactionStrings.countryCode}
                    component={Select}
                    options={Helpers.countryCodes}
                  />
                  <ErrorMessage name='countryCode' />
                  <p>Or</p>
                  <Field
                    name="regionCode"
                    label={ActivityTransactionStrings.regionCode}
                    component={Select}
                    options={Helpers.regionCodes}
                  />
                  <ErrorMessage name='regionCode' />
                  <Field
                    name='description'
                    label={ActivityTransactionStrings.description}
                    component={TextField}
                  />
                  <ErrorMessage name='description' />
                  <Field
                    name="flowType"
                    label={ActivityTransactionStrings.flowType}
                    component={Select}
                    options={Helpers.defaultFlowType}
                  />
                  <ErrorMessage name='flowType' />
                  <Field
                    name="tiedStatus"
                    label={ActivityTransactionStrings.tiedStatus}
                    component={Select}
                    options={Helpers.defaultTiedStatus}
                  />
                  <ErrorMessage name='tiedStatus' />
                  <Field
                    name="financeType"
                    label={ActivityTransactionStrings.financeType}
                    component={Select}
                    options={Helpers.defaultFinanceType}
                  />
                  <ErrorMessage name='financeType' />
                  <Field
                    name="aidType"
                    label={ActivityTransactionStrings.aidType}
                    component={Select}
                    options={Helpers.defaultAidType}
                  />
                  <ErrorMessage name='aidType' />
                  <br />
                  {formProps.isSubmitting && <LinearProgress />}
                  <br />
                  <Button type='submit' color="primary" disabled={formProps.isSubmitting}>
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

const mapStateToProps = (state: ApplicationState): ActivityTransactionKeyProps => {
  //console.log(state.orgReader)
  return {
    activitiesRef: state.keys.data.activities,
    activityRef: state.keys.data.activity,
    transactionRef: state.keys.data.activityTransaction,
    transactions: state.report.data as IATIActivityTransactionReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityTransactionDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setActivityTransaction(ownProps)),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const ActivityTransaction = connect<ActivityTransactionKeyProps, ActivityTransactionDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivityTransactionForm)
