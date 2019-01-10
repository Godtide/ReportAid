import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'

import { ApplicationState } from '../../../store'
import { ActionProps, TxData } from '../../../store/types'
import { IATIOrgProps, OrgReportProps } from '../../../store/IATI/types'

import { getOrgs } from '../../../store/IATI/IATIReader/organisationReader/actions'
import { OrgData } from '../../../store/IATI/IATIReader/organisationReader/types'

//import { getDictEntries } from '../../../utils/dict'

import { setOrganisationReport } from '../../../store/IATI/IATIWriter/organisationReportsWriter/actions'

import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import MenuItem from '@material-ui/core/MenuItem'
import InputLabel from '@material-ui/core/InputLabel'
import FormControl from '@material-ui/core/FormControl'
//import { Select } from 'formik-material-ui'
import { Select } from "material-ui-formik-components"

import { Helpers, OrganisationReport, Transaction } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

interface OrgProps {
  tx: TxData,
  orgs: OrgData
}

export interface OrgDispatchProps {
  handleSubmit: (values: any) => void
  getOrgs: () => void
}

const organisationSchema = Yup.object().shape({
  orgIdentifier: Yup
    .string()
    .required('Required')
})

type OrgReportWriterFormProps = WithStyles<typeof styles> & OrgProps & OrgDispatchProps

export class OrgReportForm extends React.Component<OrgReportWriterFormProps> {

  state = {
    txKey: '',
    txSummary: '',
    toggleSubmitting: false,
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrgReportWriterFormProps) {
   super(props)
   console.log(props.orgs)
  }

  componentDidMount() {
    this.props.getOrgs()
  }

  componentDidUpdate(previousProps: OrgReportWriterFormProps) {
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

  handleSubmit = (values: OrgReportProps, setSubmitting: Function, reset: Function) => {
    //console.log('Values: ', values)
    const submitting = !this.state.toggleSubmitting
    this.setState({txKey: '', txSummary: '', toggleSubmitting: submitting, submitFunc: setSubmitting, resetFunc: reset})
    setSubmitting(submitting)
    this.props.handleSubmit(values)
  }

  render() {

    //console.log('Props orgs: ', this.props.orgs)
    let orgRefs: any = []
    Object.keys(this.props.orgs).forEach((key) => {
      orgRefs.push({ value: key, label: this.props.orgs[key].name })
    })

    /*
    function setReport(tuple(bytes32 reportRef,
                        tuple(bytes32 orgRef, uint8 orgType, bool isSecondary),
                        bytes32 issuingOrgRef,
                        bytes32 version,
                        bytes32 lang,
                        bytes32 currency,
                        bytes32 generatedTime,
                        bytes32 lastUpdatedTime) _report)",

                        export interface OrgReportProps {
                          reportRef: string
                          reportingOrgIdentifier: string
                          reportingOrgType: number
                          reportingOrgisSecondary: boolean
                          issuingOrgRef: string
                          version: string
                          lang: string
                          currency: string
                          generatedTime: string
                          lastUpdatedTime: string
                        }
                        */


    return (
      <div>
        <h2>{OrganisationReport.headingOrgReportWriter}</h2>
        <div>
          <Formik
            initialValues={ {reportingOrgIdentifier: '',
                             reportingOrgType: 0,
                             reportingOrgIsSecondary: false,
                             issuingOrgRef: '',
                             version: '',
                             lang: '',
                             currency: '',
                             generatedTime: '',
                             lastUpdatedTime: ''}
                           }
            validationSchema={organisationSchema}
            onSubmit={(values: OrgReportProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrgReportProps>) => (
              <Form>
                <Field
                  required
                  name="reportingOrgIdentifier"
                  label={OrganisationReport.reportingOrgIdentifier}
                  component={Select}
                  options={orgRefs}
                />
                <ErrorMessage name='orgIdentifier' />
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

const mapStateToProps = (state: ApplicationState): OrgProps => {
  //console.log(state.orgReader)
  return {
    tx: state.orgReportsForm.data,
    orgs: state.orgReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setOrganisationReport(ownProps)),
    getOrgs: () => dispatch(getOrgs())
  }
}

export const OrgReportWriter = withTheme(withStyles(styles)(connect<OrgProps, OrgDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(OrgReportForm)))
