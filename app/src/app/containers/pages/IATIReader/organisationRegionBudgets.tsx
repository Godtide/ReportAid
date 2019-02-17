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

import { initialise, getRegionBudgets } from '../../../store/IATI/IATIReader/organisations/organisationRegionBudgets/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrganisationRegionBudgetReport } from '../../../store/IATI/IATIReader/organisations/organisationRegionBudgets/types'
import { OrganisationsReportProps } from '../../../store/IATI/IATIReader/organisations/types'

import { OrganisationRegionBudget as OrganisationRegionBudgetStrings } from '../../../utils/strings'

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

interface OrganisationRegionBudgetProps {
  organisationsRef: string,
  organisationRef: string,
  budgets: IATIOrganisationRegionBudgetReport
}

interface OrganisationRegionBudgetDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
}

type OrganisationRegionBudgetsReaderProps =  WithStyles<typeof styles> & OrganisationRegionBudgetProps & OrganisationRegionBudgetDispatchProps

class RegionBudgets extends React.Component<OrganisationRegionBudgetsReaderProps> {

  state = {
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrganisationRegionBudgetsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  componentDidUpdate(previousProps: OrganisationRegionBudgetsReaderProps) {
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

  render() {

    const budgetsData = this.props.budgets
    let xs = ""
    let num = 0
    Object.keys(budgetsData).forEach((organisationsKey) => {
      //numOrganisations += 1
      xs += `**${OrganisationRegionBudgetStrings.organisationsReference}**: ${organisationsKey}<br />`
      Object.keys(budgetsData[organisationsKey].data).forEach((organisationKey) => {

        xs += `**${OrganisationRegionBudgetStrings.organisationReference}**: ${organisationKey}<br />`
        Object.keys(budgetsData[organisationsKey].data[organisationKey].data).forEach((budgetKey) => {
          if ( budgetsData[organisationsKey].data[organisationKey].data[budgetKey].hasOwnProperty('budgetLine') &&
               budgetsData[organisationsKey].data[organisationKey].data[budgetKey].budgetLine != "" ) {

            num += 1
            const thisbudgetData =  budgetsData[organisationsKey].data[organisationKey].data[budgetKey]

            const budgetLine = ethers.utils.parseBytes32String(thisbudgetData.budgetLine)
            const start = ethers.utils.parseBytes32String(thisbudgetData.finance.start)
            const end = ethers.utils.parseBytes32String(thisbudgetData.finance.end)

            xs += `**${OrganisationRegionBudgetStrings.budgetReference}**: ${budgetKey} <br />`
            xs += `**${OrganisationRegionBudgetStrings.regionReference}**: ${thisbudgetData.regionRef} <br />`
            xs += `**${OrganisationRegionBudgetStrings.budgetLine}**: ${budgetLine} <br />`
            xs += `**${OrganisationRegionBudgetStrings.value}**: ${thisbudgetData.finance.value} <br />`
            xs += `**${OrganisationRegionBudgetStrings.status}**: ${thisbudgetData.finance.status} <br />`
            xs += `**${OrganisationRegionBudgetStrings.budgetStart}**: ${start} <br />`
            xs += `**${OrganisationRegionBudgetStrings.budgetEnd}**: ${end} <br /><br />`
          }
        })
      })
    })

    return (
      <div>
        <h2>{OrganisationRegionBudgetStrings.headingOrganisationRegionBudgetReader}</h2>
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
                    label={OrganisationRegionBudgetStrings.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
                  <OrganisationPicker
                    setValue={formProps.setFieldValue}
                    name='organisationRef'
                    label={OrganisationRegionBudgetStrings.organisationReference}
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
          <b>{OrganisationRegionBudgetStrings.numBudgets}</b>: {num}
        </p>
        <h3>{OrganisationRegionBudgetStrings.organisationRegionBudgetDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrganisationRegionBudgetProps => {
  //console.log(state.orgReader)
  return {
    organisationsRef: state.keys.data.organisations,
    organisationRef: state.keys.data.organisation,
    budgets: state.organisationRegionBudgetsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrganisationRegionBudgetDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(getRegionBudgets(ownProps)),
    initialise: () => dispatch(initialise())
  }
}

export const OrganisationRegionBudgets = withTheme(withStyles(styles)(connect<OrganisationRegionBudgetProps, OrganisationRegionBudgetDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(RegionBudgets)))
