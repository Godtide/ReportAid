export interface IATIActivitiesProps {
  version: string
  generatedTime: string
  linkedData: string
}

export interface ActivitiesProps {
  activitiesRef: string
  version: string
  linkedData: string
}

export interface ReportingOrgProps {
  orgRef: string
  orgType: number
  isSecondary: boolean
}

export interface IATIActivityProps {
  identifier: string
  reportingOrg: ReportingOrgProps
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
  defaulTiedStatus: number
}

export interface ActivityProps {
  activitiesRef: string
  activityRef: string
  identifier: string
  reportingOrgRef: string
  reportingOrgType: number
  reportingOrgIsSecondary: boolean
  title: string
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
  defaulTiedStatus: number
}
