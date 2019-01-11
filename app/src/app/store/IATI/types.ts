export interface IATIOrgProps {
  orgRef: string
  name: string
  identifier: string
}

export interface IATIOrgReportProps {
  version: string
  orgRef: string
  reportRef: string
  reportingOrg: ReportingOrgProps
  lang: string
  currency: string
  lastUpdatedTime: string
}

export interface OrganisationProps {
  name: string
  code: string
  identifier: string
}

export interface ReportingOrgProps {
  orgRef: string
  orgType: number
  isSecondary: boolean
}

export interface OrgReportProps {
  reportingOrgRef: string
  reportingOrgType: number
  reportingOrgIsSecondary: boolean
  orgRef: string
  version: string
  lang: string
  currency: string
}
