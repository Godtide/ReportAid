import { PayloadProps } from '../../../types'

export interface ActivityReportProps {
  isReport: boolean
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

export interface IATIActivityData {
  activityRef: string
  identifier: string
  reportingOrgRef: string
  reportingOrgType: number
  reportingOrgIsSecondary: boolean
  title: string
  lastUpdated: string
  lang: string
  currency: string
  humanitarian: boolean
  hierarchy: number
  linkedData: string
  budgetNotProvided: number
  status: number
  scope: number
  capitalSpend: number
  collaborationType: number
  defaultFlowType: number
  defaultFinanceType: number
  defaultAidType: string
  defaultTiedStatus: number
}

export interface IATIActivityReport {
  activitiesRef: string
  data: Array<IATIActivityData>
}

export interface IATIActivityReportProps extends PayloadProps {
  data: IATIActivityReport
}
