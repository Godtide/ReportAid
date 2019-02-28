import { PayloadProps } from '../../../types'

export interface ActivityReportProps {
  activitiesRef: string
}

export interface ActivitiesReportProps {
 activitiesRef: string
 activityRef: string
}

export interface IATIActivitiesData {
  activitiesRef: string
  version: string
  generatedTime: string
  linkedData: string
}

export interface IATIActivitiesReport {
  data: Array<IATIActivitiesData>
}

export interface IATIActivitiesReportProps extends PayloadProps {
  data: IATIActivitiesReport
}
