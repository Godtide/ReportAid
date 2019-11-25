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
import { IATIActivityTerritoryReport, ActivityTerritoryProps } from '../../../../store/IATI/types'
import { FormData } from '../../../../store/helpers/forms/types'

import { setFormFunctions } from '../../../../store/helpers/forms/actions'
import { setActivityTerritory } from '../../../../store/IATI/IATIWriter/activities/activityTerritories/actions'

import { ActivitiesPicker } from '../../../../components/io/activitiesPicker'
import { ActivityPicker } from '../../../../components/io/activityPicker'
import { TransactionHelper } from '../../../io/transactionHelper'

import { ActivityTerritories as ActivityTerritoryStrings } from '../../../../utils/strings'
import { Helpers } from '../../../../utils/config'

const ActivityTerritorySchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required'),
  activityRef: Yup
    .string()
    .required('Required'),
  territoryRef: Yup
    .string()
    .required('Required'),
  countryCode: Yup
    .string(),
  regionCode: Yup
    .string(),
  percentage: Yup
    .number()
    .required('Required'),
})

interface ActivityTerritoryKeyProps {
  activitiesRef: string
  activityRef: string
  territoryRef: string
  territories: IATIActivityTerritoryReport
}

export interface ActivityTerritoryDispatchProps {
  handleSubmit: (values: any) => void
  setFormFunctions: (formProps: FormData) => void
}

type ActivityTerritoryFormProps = ActivityTerritoryKeyProps & ActivityTerritoryDispatchProps

export class ActivityTerritoryForm extends React.Component<ActivityTerritoryFormProps> {

  constructor (props: ActivityTerritoryFormProps) {
   super(props)
  }

  handleSubmit = (values: ActivityTerritoryProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  getData = (territories: IATIActivityTerritoryReport): ActivityTerritoryProps => {

    let newTerritory: ActivityTerritoryProps = {
      activitiesRef: this.props.activitiesRef,
      activityRef: this.props.activityRef,
      territoryRef: this.props.territoryRef,
      countryCode: "",
      regionCode: "",
      percentage: 100
    }
    if ( typeof territories.data != 'undefined' ) {
      let countryCode = territories.data[0].territory
      let regionCode = countryCode
      if(!isNaN(parseInt(countryCode, 10))) {
        regionCode = ""
      } else {
        countryCode = ""
      }
      newTerritory.countryCode = countryCode
      newTerritory.regionCode = regionCode
      newTerritory.percentage = territories.data[0].percentage
    }

    return newTerritory
  }

  render() {

    const thisTerritory: ActivityTerritoryProps = this.getData(this.props.territories)

    return (
      <div>
        <h2>{ActivityTerritoryStrings.headingActivityTerritoriesWriter}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: this.props.activitiesRef,
                             activityRef: this.props.activityRef,
                             territoryRef: this.props.territoryRef,
                             countryCode: thisTerritory.countryCode,
                             regionCode: thisTerritory.regionCode,
                             percentage: thisTerritory.percentage
                            }}
            enableReinitialize={true}
            validationSchema={ActivityTerritorySchema}
            onSubmit={(values: ActivityTerritoryProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivityTerritoryProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <Field
                    readOnly
                    name='territoryRef'
                    value={this.props.territoryRef}
                    label={ActivityTerritoryStrings.territoryRef}
                    component={TextField}
                  />
                  <ActivitiesPicker
                    setValue={formProps.setFieldValue}
                    name='activitiesRef'
                    label={ActivityTerritoryStrings.activitiesReference}
                  />
                  <ErrorMessage name='activitiesRef' />
                  <ActivityPicker
                    setValue={formProps.setFieldValue}
                    name='activityRef'
                    label={ActivityTerritoryStrings.activityReference}
                  />
                  <ErrorMessage name='activityRef' />
                  <Field
                    name="countryCode"
                    label={ActivityTerritoryStrings.countryCode}
                    component={Select}
                    options={Helpers.countryCodes}
                  />
                  <ErrorMessage name='countryCode' />
                  <p>Or</p>
                  <Field
                    name="regionCode"
                    label={ActivityTerritoryStrings.regionCode}
                    component={Select}
                    options={Helpers.regionCodes}
                  />
                  <ErrorMessage name='regionCode' />
                  <Field
                    name="percentage"
                    label={ActivityTerritoryStrings.percentage}
                    component={TextField}
                  />
                  <ErrorMessage name='percentage' />
                  <br />
                  {formProps.isSubmitting && <LinearProgress />}
                  <br />
                  <Button type='submit' variant="contained" color="primary" disabled={formProps.isSubmitting}>
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

const mapStateToProps = (state: ApplicationState): ActivityTerritoryKeyProps => {
  //console.log(state.orgReader)
  return {
    activitiesRef: state.keys.data.activities,
    activityRef: state.keys.data.activity,
    territoryRef: state.keys.data.activityTerritory,
    territories: state.report.data as IATIActivityTerritoryReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivityTerritoryDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setActivityTerritory(ownProps)),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const ActivityTerritory = connect<ActivityTerritoryKeyProps, ActivityTerritoryDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivityTerritoryForm)
