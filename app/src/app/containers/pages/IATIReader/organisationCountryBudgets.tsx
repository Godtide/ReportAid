import * as React from 'react'
import { connect } from 'react-redux'
import { ThunkDispatch } from 'redux-thunk'

import { ethers } from 'ethers'
import Markdown from 'react-markdown'

import { Formik, Form, Field, FormikProps, ErrorMessage} from 'formik'
import * as Yup from 'yup'
import { LinearProgress } from '@material-ui/core'
import Button from '@material-ui/core/Button'
//import { Date } from 'formik-material-ui'
import FormControl from '@material-ui/core/FormControl'

import { OrganisationsPicker } from '../../../components/io/organisationsPicker'
import { OrganisationPicker } from '../../../components/io/organisationPicker'

import { initialise, getCountryBudgets } from '../../../store/IATI/IATIReader/organisations/organisationCountryBudgets/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrganisationCountryBudgetReport } from '../../../store/IATI/IATIReader/organisations/organisationCountryBudgets/types'
import { OrganisationsReportProps } from '../../../store/IATI/IATIReader/organisations/types'

import { OrganisationCountryBudget as OrganisationCountryBudgetStrings } from '../../../utils/strings'

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

interface OrganisationCountryBudgetProps {
  budgets: IATIOrganisationCountryBudgetReport
}

interface OrganisationCountryBudgetDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
}

type OrganisationCountryBudgetsReaderProps =  WithStyles<typeof styles> & OrganisationCountryBudgetProps & OrganisationCountryBudgetDispatchProps

class CountryBudgets extends React.Component<OrganisationCountryBudgetsReaderProps> {

  state = {
    organisationsRef: "",
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrganisationCountryBudgetsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  componentDidUpdate(previousProps: OrganisationCountryBudgetsReaderProps) {
    if(previousProps.budgets != this.props.budgets) {
      this.state.submitFunc(false)
      this.state.resetFunc()
    }
  }

  handleSubmit = (values: OrganisationsReportProps, setSubmitting: Function, reset: Function) => {
    this.setState({submitFunc: setSubmitting, resetFunc: reset})
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

    const budgetsData = this.props.budgets
    let xs = ""
    let num = 0
    Object.keys(budgetsData).forEach((organisationsKey) => {
      //numOrganisations += 1
      xs += `**${OrganisationCountryBudgetStrings.organisationsReference}**: ${organisationsKey}<br />`
      Object.keys(budgetsData[organisationsKey].data).forEach((organisationKey) => {

        xs += `**${OrganisationCountryBudgetStrings.organisationReference}**: ${organisationKey}<br />`
        Object.keys(budgetsData[organisationsKey].data[organisationKey].data).forEach((budgetKey) => {
          if ( budgetsData[organisationsKey].data[organisationKey].data[budgetKey].hasOwnProperty('budgetLine') &&
               budgetsData[organisationsKey].data[organisationKey].data[budgetKey].budgetLine != "" ) {

            num += 1
            const thisbudgetData =  budgetsData[organisationsKey].data[organisationKey].data[budgetKey]

            const countryRef = ethers.utils.parseBytes32String(thisbudgetData.countryRef)
            const budgetLine = ethers.utils.parseBytes32String(thisbudgetData.budgetLine)
            const start = ethers.utils.parseBytes32String(thisbudgetData.finance.start)
            const end = ethers.utils.parseBytes32String(thisbudgetData.finance.end)
            xs+= `**${OrganisationCountryBudgetStrings.budgetReference}**: ${budgetKey} <br />`
            xs+= `**${OrganisationCountryBudgetStrings.countryReference}**: ${countryRef} <br />`
            xs+= `**${OrganisationCountryBudgetStrings.budgetLine}**: ${budgetLine} <br />`
            xs+= `**${OrganisationCountryBudgetStrings.value}**: ${thisbudgetData.finance.value} <br />`
            xs+= `**${OrganisationCountryBudgetStrings.status}**: ${thisbudgetData.finance.status} <br />`
            xs+= `**${OrganisationCountryBudgetStrings.budgetStart}**: ${start} <br />`
            xs+= `**${OrganisationCountryBudgetStrings.budgetEnd}**: ${end} <br /><br />`
          }
        })
      })
    })

    return (
      <div>
        <h2>{OrganisationCountryBudgetStrings.headingOrganisationCountryBudgetReader}</h2>
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
                    changeFunction={this.handleOrganisationsChange}
                    setValue={formProps.setFieldValue}
                    name='organisationsRef'
                    label={OrganisationCountryBudgetStrings.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
                  <OrganisationPicker
                    organisationsRef={this.state.organisationsRef}
                    changeFunction={this.handleOrganisationChange}
                    setValue={formProps.setFieldValue}
                    name='organisationRef'
                    label={OrganisationCountryBudgetStrings.organisationReference}
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
        <p>
          <b>{OrganisationCountryBudgetStrings.numBudgets}</b>: {num}
        </p>
        <h3>{OrganisationCountryBudgetStrings.organisationCountryBudgetDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationCountryBudgetProps => {
  //console.log(state.orgReader)
  return {
    budgets: state.organisationCountryBudgetsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationCountryBudgetDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(getCountryBudgets(ownProps)),
    initialise: () => dispatch(initialise())
  }
}

export const OrganisationCountryBudgets = withTheme(withStyles(styles)(connect<OrganisationCountryBudgetProps, OrganisationCountryBudgetDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(CountryBudgets)))
