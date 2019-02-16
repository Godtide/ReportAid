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

import { OrganisationsPicker } from '../../../components/io/organisationsPicker'
import { OrganisationPicker } from '../../../components/io/organisationPicker'

import { initialise, getBudgets } from '../../../store/IATI/IATIReader/organisations/organisationBudgets/actions'

import { ApplicationState } from '../../../store'
import { ActionProps } from '../../../store/types'
import { IATIOrganisationBudgetReport } from '../../../store/IATI/IATIReader/organisations/organisationBudgets/types'
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
  budgets: IATIOrganisationBudgetReport
}

interface OrgBudgetDispatchProps {
  handleSubmit: (values: any) => void
  initialise: () => void
}

type OrganisationBudgetsReaderProps =  WithStyles<typeof styles> & OrgBudgetProps & OrgBudgetDispatchProps

class Budgets extends React.Component<OrganisationBudgetsReaderProps> {

  state = {
    organisationsRef: "",
    submitFunc: (function(submit: boolean) { return submit }),
    resetFunc: (function() { return null })
  }

  constructor (props: OrganisationBudgetsReaderProps) {
    super(props)
  }

  componentDidMount() {
    this.props.initialise()
  }

  componentDidUpdate(previousProps: OrganisationBudgetsReaderProps) {
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
      xs += `**${OrgBudgetStrings.organisationsReference}**: ${organisationsKey}<br />`
      Object.keys(budgetsData[organisationsKey].data).forEach((organisationKey) => {

        xs += `**${OrgBudgetStrings.organisationReference}**: ${organisationKey}<br />`
        Object.keys(budgetsData[organisationsKey].data[organisationKey].data).forEach((budgetKey) => {
          if ( budgetsData[organisationsKey].data[organisationKey].data[budgetKey].hasOwnProperty('budgetLine') &&
               budgetsData[organisationsKey].data[organisationKey].data[budgetKey].budgetLine != "" ) {

            num += 1
            const thisbudgetData =  budgetsData[organisationsKey].data[organisationKey].data[budgetKey]

            const budgetLine = ethers.utils.parseBytes32String(thisbudgetData.budgetLine)
            const start = ethers.utils.parseBytes32String(thisbudgetData.finance.start)
            const end = ethers.utils.parseBytes32String(thisbudgetData.finance.end)
            xs+= `**${OrgBudgetStrings.budgetReference}**: ${budgetKey} <br />`
            xs+= `**${OrgBudgetStrings.budgetLine}**: ${budgetLine} <br />`
            xs+= `**${OrgBudgetStrings.value}**: ${thisbudgetData.finance.value} <br />`
            xs+= `**${OrgBudgetStrings.status}**: ${thisbudgetData.finance.status} <br />`
            xs+= `**${OrgBudgetStrings.budgetStart}**: ${start} <br />`
            xs+= `**${OrgBudgetStrings.budgetEnd}**: ${end} <br /><br />`
          }
        })
      })
    })

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
                    changeFunction={this.handleOrganisationsChange}
                    setValue={formProps.setFieldValue}
                    name='organisationsRef'
                    label={OrgBudgetStrings.organisationsReference}
                  />
                  <ErrorMessage name='organisationsRef' />
                  <OrganisationPicker
                    organisationsRef={this.state.organisationsRef}
                    changeFunction={this.handleOrganisationChange}
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
        <p>
          <b>{OrgBudgetStrings.numBudgets}</b>: {num}
        </p>
        <h3>{OrgBudgetStrings.organisationBudgetDetails}</h3>
        <Markdown escapeHtml={false} source={xs} />
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): OrgBudgetProps => {
  //console.log(state.orgReader)
  return {
    budgets: state.organisationBudgetsReader.data
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): OrgBudgetDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(getBudgets(ownProps)),
    initialise: () => dispatch(initialise())
  }
}

export const OrganisationBudgets = withTheme(withStyles(styles)(connect<OrgBudgetProps, OrgBudgetDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(Budgets)))
