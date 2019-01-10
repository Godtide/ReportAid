export interface IATIOrgProps {
  orgRef: string
  name: string
  identifier: string
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

export interface ReportProps {
  reportRef: string
  reportingOrg: ReportingOrgProps
  issuingOrgRef: string
  version: string
  lang: string
  currency: string
  generatedTime: string
  lastUpdatedTime: string
}

export interface OrgReportProps {
  reportingOrgIdentifier: string
  reportingOrgType: number
  reportingOrgIsSecondary: boolean
  issuingOrgRef: string
  version: string
  lang: string
  currency: string
}
