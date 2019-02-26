import { PayloadProps } from '../../../types'

export interface OrganisationReportProps {
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

export const enum IATIReportActionTypes {
  ORGS_INIT = '@@IATIReportActionTypes/READORGS_INIT',
  ORGS_SUCCESS = '@@IATIReportActionTypes/READORGS_SUCCESS',
  ORGS_FAILURE = '@@IATIReportActionTypes/READORGS_FAILURE',
  ORGANISATIONS_INIT = '@@IATIReportActionTypes/READORGANISATIONS_INIT',
  ORGANISATIONS_SUCCESS = '@@IATIReportActionTypes/READORGANISATIONS_SUCCESS',
  ORGANISATIONS_FAILURE = '@@IATIReportActionTypes/READORGANISATIONS_FAILURE',
  ORGANISATION_INIT = '@@IATIReportActionTypes/READORGANISATION_INIT',
  ORGANISATION_SUCCESS = '@@IATIReportActionTypes/READORGANISATION_SUCCESS',
  ORGANISATION_FAILURE = '@@IATIReportActionTypes/READORGANISATION_FAILURE',
  BUDGET_INIT = '@@IATIReportActionTypes/READBUDGET_INIT',
  BUDGET_SUCCESS = '@@IATIReportActionTypes/READBUDGET_SUCCESS',
  BUDGET_FAILURE = '@@IATIReportActionTypes/READBUDGET_FAILURE',
  RECIPIENTORGBUDGET_INIT = '@@IATIReportActionTypes/READRECIPIENTORGBUDGET_INIT',
  RECIPIENTORGBUDGET_SUCCESS = '@@IATIReportActionTypes/READRECIPIENTORGBUDGET_SUCCESS',
  RECIPIENTORGBUDGET_FAILURE = '@@IATIReportActionTypes/READRECIPIENTORGBUDGET_FAILURE',
  RECIPIENTREGIONBUDGET_INIT = '@@IATIReportActionTypes/READRECIPIENTREGIONBUDGET_INIT',
  RECIPIENTREGIONBUDGET_SUCCESS = '@@IATIReportActionTypes/READRECIPIENTREGIONBUDGET_SUCCESS',
  RECIPIENTREGIONBUDGET_FAILURE = '@@IATIReportActionTypes/READRECIPIENTREGIONBUDGET_FAILURE',
  RECIPIENTCOUNTRYBUDGET_INIT = '@@IATIReportActionTypes/READRECIPIENTCOUNTRYBUDGET_INIT',
  RECIPIENTCOUNTRYBUDGET_SUCCESS = '@@IATIReportActionTypes/READRECIPIENTCOUNTRYBUDGET_SUCCESS',
  RECIPIENTCOUNTRYBUDGET_FAILURE = '@@IATIReportActionTypes/READRECIPIENTCOUNTRYBUDGET_FAILURE',
  TOTALEXPENDITURE_INIT = '@@IATIReportActionTypes/READTOTALEXPENDITURE_INIT',
  TOTALEXPENDITURE_SUCCESS = '@@IATIReportActionTypes/READTOTALEXPENDITURE_SUCCESS',
  TOTALEXPENDITURE_FAILURE = '@@IATIReportActionTypes/READTOTALEXPENDITURE_FAILURE',
  DOCUMENT_INIT = '@@IATIReportActionTypes/READDOCUMENT_INIT',
  DOCUMENT_SUCCESS = '@@IATIReportActionTypes/READDOCUMENT_SUCCESS',
  DOCUMENT_FAILURE = '@@IATIReportActionTypes/READDOCUMENT_FAILURE'
}
