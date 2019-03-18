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
import { IATIActivitySectorReport, ActivitySectorProps } from '../../../../store/IATI/types'
import { FormData } from '../../../../store/helpers/forms/types'

import { setFormFunctions } from '../../../../store/helpers/forms/actions'
import { setActivitySector } from '../../../../store/IATI/IATIWriter/activities/activitySectors/actions'

import { ActivitiesPicker } from '../../../../components/io/activitiesPicker'
import { ActivityPicker } from '../../../../components/io/activityPicker'
import { TransactionHelper } from '../../../io/transactionHelper'

import { ActivitySectors as ActivitySectorStrings } from '../../../../utils/strings'
import { Helpers } from '../../../../utils/config'

import withStyles, { WithStyles } from '@material-ui/core/styles/withStyles'
import { withTheme, styles } from '../../../../styles/theme'

const activitySectorSchema = Yup.object().shape({
  activitiesRef: Yup
    .string()
    .required('Required'),
  activityRef: Yup
    .string()
    .required('Required'),
  sectorRef: Yup
    .string()
    .required('Required'),
  dacCode: Yup
    .number()
    .required('Required'),
  percentage: Yup
    .number()
    .required('Required'),
})

interface ActivitySectorKeyProps {
  activitiesRef: string
  activityRef: string
  sectorRef: string
  sectors: IATIActivitySectorReport
}

export interface ActivitySectorDispatchProps {
  handleSubmit: (values: any) => void
  setFormFunctions: (formProps: FormData) => void
}

type ActivitySectorFormProps = WithStyles<typeof styles> & ActivitySectorKeyProps & ActivitySectorDispatchProps

export class ActivitySectorForm extends React.Component<ActivitySectorFormProps> {

  constructor (props: ActivitySectorFormProps) {
   super(props)
  }

  handleSubmit = (values: ActivitySectorProps, setSubmitting: Function, reset: Function) => {
    this.props.setFormFunctions({submitFunc: setSubmitting, resetFunc: reset})
    this.props.handleSubmit(values)
  }

  getData = (sectors: IATIActivitySectorReport): ActivitySectorProps => {

    let newSector: ActivitySectorProps = {
      activitiesRef: this.props.activitiesRef,
      activityRef: this.props.activityRef,
      sectorRef: this.props.sectorRef,
      dacCode: 0,
      percentage: 100
    }
    if ( typeof sectors.data != 'undefined' ) {
      newSector.dacCode = sectors.data[0].dacCode
      newSector.percentage = sectors.data[0].percentage
    }

    return newSector
  }

  render() {

    const thisSector: ActivitySectorProps = this.getData(this.props.sectors)

    return (
      <div>
        <h2>{ActivitySectorStrings.headingActivitySectorsWriter}</h2>
        <div>
          <Formik
            initialValues={ {activitiesRef: this.props.activitiesRef,
                             activityRef: this.props.activityRef,
                             sectorRef: this.props.sectorRef,
                             dacCode: thisSector.dacCode,
                             percentage: thisSector.percentage
                            }}
            enableReinitialize={true}
            validationSchema={activitySectorSchema}
            onSubmit={(values: ActivitySectorProps, actions: any) => {
              this.handleSubmit(values, actions.setSubmitting, actions.resetForm)
            }}
            render={(formProps: FormikProps<ActivitySectorProps>) => (
              <Form>
                <FormControl fullWidth={true}>
                  <Field
                    readOnly
                    name='sectorRef'
                    value={this.props.sectorRef}
                    label={ActivitySectorStrings.sectorRef}
                    component={TextField}
                  />
                  <ActivitiesPicker
                    setValue={formProps.setFieldValue}
                    name='activitiesRef'
                    label={ActivitySectorStrings.activitiesReference}
                  />
                  <ErrorMessage name='activitiesRef' />
                  <ActivityPicker
                    setValue={formProps.setFieldValue}
                    name='activityRef'
                    label={ActivitySectorStrings.activityReference}
                  />
                  <ErrorMessage name='activityRef' />
                  <Field
                    name="dacCode"
                    label={ActivitySectorStrings.dacCode}
                    component={Select}
                    options={Helpers.DACSector}
                  />
                  <ErrorMessage name='dacCode' />
                  <Field
                    name="percentage"
                    label={ActivitySectorStrings.percentage}
                    component={TextField}
                  />
                  <ErrorMessage name='percentage' />
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
        <TransactionHelper/>
      </div>
    )
  }
}

const mapStateToProps = (state: ApplicationState): ActivitySectorKeyProps => {
  //console.log(state.orgReader)
  return {
    activitiesRef: state.keys.data.activities,
    activityRef: state.keys.data.activity,
    sectorRef: state.keys.data.activitySector,
    sectors: state.report.data as IATIActivitySectorReport
  }
}

const mapDispatchToProps = (dispatch: ThunkDispatch<ApplicationState, any, ActionProps>): ActivitySectorDispatchProps => {
  return {
    handleSubmit: (ownProps: any) => dispatch(setActivitySector(ownProps)),
    setFormFunctions: (formProps: FormData) => dispatch(setFormFunctions(formProps))
  }
}

export const ActivitySector = withTheme(withStyles(styles)(connect<ActivitySectorKeyProps, ActivitySectorDispatchProps, {}, ApplicationState>(
  mapStateToProps,
  mapDispatchToProps
)(ActivitySectorForm)))
