import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
import FormControl from '@material-ui/core/FormControl'

import { getDictEntries } from '../../../components/io/dict'
import { OrganisationsPicker } from '../../../components/io/organisationsPicker'
import { OrganisationPicker } from '../../../components/io/organisationPicker'
import { FormData } from '../../../store/helpers/forms/types'

import { setFormFunctions } from '../../../store/helpers/forms/actions'
import { initialise, getBudgets } from '../../../store/IATI/IATIReader/organisations/organisationBudgets/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIBudgetReport } from '../../../store/IATI/IATIReader/organisations/organisationBudgets/types'
import { OrganisationsReportProps } from '../../../store/IATI/IATIReader/organisations/types'

import { OrganisationBudget as OrgBudgetStrings } from '../../../utils/strings'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../styles/theme'

const reportSchema = Yup.object().shape({
  organisationsRef: Yup
    .string()
    .required('Required'),
  organisationRef: Yup
    .string()
    .required('Required')
})

interface OrgBudgetProps {
  submittingFunc: Function,
  resettingFunc: Function
  organisationsRef: string,
  organisationRef: string,
  budgets: IATIBudgetReport
}

interface OrgBudgetDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
  setFormFunctions: (formProps: FormData) => void
}

type OrganisationBudgetsReaderProps =  WithStyles<typeof styles> & OrgBudgetProps & OrgBudgetDispatchProps

class Budgets extends React.Component<OrganisationBudgetsReaderProps> {

  constructor (props: OrganisationBudgetsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  componentDidUpdate(previousProps: OrganisationBudgetsReaderProps) {
    if(previousProps.budgets != this.props.budgets) {
      this.props.submittingFunc(false)
      this.props.resettingFunc()
    }
  }

  handleSubmit = (values: OrganisationsReportProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.initialise()
    this.props.handleSubmit(values)
  }

  handleOrganisationsChange = (value: string) => {
    this.setState({organisationsRef: value})
  }

  handleOrganisationChange = (value: string) => {
    console.log(value)
  }

  render() {

    const xs = getDictEntries(this.props.budgets)

    return (
      <div>
        <h2>{OrgBudgetStrings.headingOrganisationBudgetReader}</h2>
        <div>
          <Formik
            initialValues={ {organisationsRef: "",
                             organisationRef: ""
                            }}
            validationSchema={reportSchema}
            onSubmit={(values: OrganisationsReportProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<OrganisationsReportProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <OrganisationsPicker
                    setValue={formProps.setFieldValue}
                    name='organisationsRef'
                    label={OrgBudgetStrings.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
                  <OrganisationPicker
                    setValue={formProps.setFieldValue}
                    name='organisationRef'
                    label={OrgBudgetStrings.organisationReference}
                  />
                  <ErrorMessage name='organisationRef' />
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
        <hr />
        <h3>{OrgBudgetStrings.organisationBudgetDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgBudgetProps => {
  //console.log(state.orgReader)
  return {
    submittingFunc: state.forms.data.submitFunc,
    resettingFunc: state.forms.data.resetFunc,
    organisationsRef: state.keys.data.organisations,
    organisationRef: state.keys.data.organisation,
    budgets: state.organisationBudgetsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgBudgetDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(getBudgets(ownProps)),
    initialise: () => dispatch(initialise()),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const OrganisationBudgets = withTheme(withStyles(styles)(connect<OrgBudgetProps, OrgBudgetDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Budgets)))
