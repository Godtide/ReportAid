import { PayloadProps } from '../../../types'

export interface OrganisationReportProps {
  isReport: boolean
  organisationsRef: string
}

export interface OrganisationsReportProps {
 organisationsRef: string
 organisationRef: string
}

export interface IATIOrganisationsData {
  organisationsRef: string
  version: string
  generatedTime: string
}

export interface IATIOrganisationsReport {
  data: Array<IATIOrganisationsData>
}

export interface IATIOrganisationsReportProps extends PayloadProps {
  data: IATIOrganisationsReport
}

export interface IATIOrganisationData {
    organisationRef: string
    issuingOrgRef: string
    reportingOrgRef: string
    reportingOrgType: number
    reportingOrgIsSecondary: boolean
    lang: string
    currency: string
    lastUpdatedTime: string
}

export interface IATIOrganisationReport {
  organisationsRef: string
  data: Array<IATIOrganisationData>
}

export interface IATIOrganisationReportProps extends PayloadProps {
  data: IATIOrganisationReport
}

export interface IATIOrgData {
  orgRef: string
  name: string
  identifier: string
}

export interface IATIOrgReport {
  data: Array<IATIOrgData>
}

export interface IATIOrgReportProps extends PayloadProps {
  data: IATIOrgReport
}

export interface IATIOrganisationDocData {
  title: string
  format: string
  url: string
  category: string
  countryRef: string
  description: string
  language: string
  date: string
}

export interface IATIOrganisationDocReport {
  organisationsRef: string
  organisationRef: string
  data: Array<IATIOrganisationDocData>
}

export interface IATIOrganisationDocReportProps extends PayloadProps {
  data: IATIOrganisationDocReport
}
