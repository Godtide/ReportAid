import { PayloadProps } from '../types'

/* IATI Organisation Reader */

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

export interface IATIExpenditureData {
  expenditureKey: string
  expenditureLine: string,
  value: number,
  status: number,
  start: string,
  end: string
}

export interface IATIRecipientOrgBudgetData {
  budgetKey: string
  budgetLine: string,
  recipientOrgRef: string,
  value: number,
  status: number,
  start: string,
  end: string
}

export interface IATIRegionBudgetData {
  budgetKey: string
  budgetLine: string,
  regionRef: string,
  value: number,
  status: number,
  start: string,
  end: string
}

export interface IATICountryBudgetData {
  budgetKey: string
  budgetLine: string,
  countryRef: string,
  value: number,
  status: number,
  start: string,
  end: string
}

export interface IATIBudgetData {
  budgetKey: string
  budgetLine: string,
  value: number,
  status: number,
  start: string,
  end: string
}

export interface IATIBudgetReport {
  organisationsRef: string
  organisationRef: string
  data: Array<IATIBudgetData | IATIExpenditureData | IATIRecipientOrgBudgetData | IATIRegionBudgetData | IATICountryBudgetData >
}

export interface IATIBudgetReportProps extends PayloadProps {
  data: IATIBudgetReport
}

/* IATI Organisation Writer */

export interface IATIOrgProps {
  name: string
  identifier: string
}

export interface OrgProps {
  orgRef: string
  name: string
  code: string
  identifier: string
}

export interface IATIOrganisationsProps {
  version: string
  generatedTime: string
}

export interface OrganisationsProps {
  organisationsRef: string
  version: string
}

export interface ReportingOrgProps {
  orgType: number
  isSecondary: boolean
  orgRef: string
}

export interface IATIOrganisationProps {
  orgRef: string
  reportingOrg: ReportingOrgProps
  lang: string
  currency: string
  lastUpdatedTime: string
}

export interface OrganisationProps {
  organisationsRef: string
  organisationRef: string
  orgRef: string
  reportingOrgRef: string
  reportingOrgType: number
  reportingOrgIsSecondary: boolean
  lang: string
  currency: string
}

export interface IATIOrganisationDocProps {
  title: string
  format: string
  url: string
  category: string
  countryRef: string
  desc: string
  lang: string
  date: string
}

export interface OrganisationDocProps {
  organisationsRef: string
  organisationRef: string
  docRef: string
  title: string
  format: string
  url: string
  category: string
  countryRef: string
  desc: string
  lang: string
  day: number
  month: number
  year: number
}

interface FinanceProps {
  value: number
  status: number
  start: string
  end: string
}

export interface IATIBudgetProps {
  budgetType: number
  budgetLine: string
  otherRef: string
  finance: FinanceProps
}

export interface OrganisationBudgetProps {
  organisationsRef: string
  organisationRef: string
  budgetRef: string
  budgetLine: string
  value: number
  status: number
  startDay: number
  startMonth: number
  startYear: number
  endDay: number
  endMonth: number
  endYear: number
}

export interface OrganisationExpenditureProps {
  organisationsRef: string
  organisationRef: string
  expenditureRef: string
  expenditureLine: string
  value: number
  status: number
  startDay: number
  startMonth: number
  startYear: number
  endDay: number
  endMonth: number
  endYear: number
}

export interface OrganisationRecipientBudgetProps {
  organisationsRef: string
  organisationRef: string
  budgetRef: string
  recipientOrgRef: string
  budgetLine: string
  value: number
  status: number
  startDay: number
  startMonth: number
  startYear: number
  endDay: number
  endMonth: number
  endYear: number
}

export interface OrganisationRegionBudgetProps {
  organisationsRef: string
  organisationRef: string
  budgetRef: string
  regionRef: string
  budgetLine: string
  value: number
  status: number
  startDay: number
  startMonth: number
  startYear: number
  endDay: number
  endMonth: number
  endYear: number
}

export interface OrganisationCountryBudgetProps {
  organisationsRef: string
  organisationRef: string
  budgetRef: string
  countryRef: string
  budgetLine: string
  value: number
  status: number
  startDay: number
  startMonth: number
  startYear: number
  endDay: number
  endMonth: number
  endYear: number
}

/* IATI Activities Writer */

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

export interface IATIActivityProps {
  humanitarian: boolean
  hierarchy: number
  status: number
  budgetNotProvided: number
  reportingOrg: ReportingOrgProps
  lastUpdated: string
  lang: string
  currency: string
  linkedData: string
  identifier: string
  title: string
  description: string
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
  budgetNotProvided: number
  status: number
  hierarchy: number
  linkedData: string
  description: string
}

export interface IATIActivityAdditionalProps {
  defaultAidType: string
  defaultFinanceType: number
  scope: number
  capitalSpend: number
  collaborationType: number
  defaultFlowType: number
  defaultTiedStatus: number
}

export interface ActivityAdditionalProps {
  activitiesRef: string
  activityRef: string
  additionalRef: string
  defaultAidType: string
  defaultFinanceType: number
  scope: number
  capitalSpend: number
  collaborationType: number
  defaultFlowType: number
  defaultTiedStatus: number
}

export interface ActivityDateProps {
  activitiesRef: string
  activityRef: string
  dateRef: string
  dateType: number
  day: number
  month: number
  year: number
  narrative: string
}

/* IATI Activities Reader */

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
  title: string
  description: string
  identifier: string
  linkedData: string
  lang: string
  currency: string
  lastUpdated: string
  reportingOrgRef: string
  reportingOrgType: number
  reportingOrgIsSecondary: boolean
  status: number
  humanitarian: boolean
  hierarchy: number
  budgetNotProvided: number
}

export interface IATIActivityReport {
  activitiesRef: string
  data: Array<IATIActivityData>
}

export interface IATIActivityReportProps extends PayloadProps {
  data: IATIActivityReport
}

export interface IATIActivityAdditionalData {
  additionalRef: string
  defaultAidType: string
  defaultFinanceType: number
  scope: number
  capitalSpend: number
  collaborationType: number
  defaultFlowType: number
  defaultTiedStatus: number
}

export interface IATIActivityAdditionalReport {
  activitiesRef: string
  activityRef: string
  data: Array<IATIActivityAdditionalData>
}

export interface IATIActivityAdditionalReportProps extends PayloadProps {
  data: IATIActivityAdditionalReport
}

export interface IATIActivityDatesData {
  dateRef: string
  dateType: number
  date: string
  narrative: string
}

export interface IATIActivityDatesReport {
  activitiesRef: string
  activityRef: string
  data: Array<IATIActivityDatesData>
}

export interface IATIActivityDatesReportProps extends PayloadProps {
  data: IATIActivityDatesReport
}



/* Action Types */

export const enum IATIWriterActionTypes{
  TX_INIT =  '@@IATIWriterActionTypes/TX_INIT',
  ORGS_SUCCESS = '@@IATIWriterActionTypes/ORGS_SUCCESS',
  ORGS_FAILURE = '@@IATIWriterActionTypes/ORGS_FAILURE',
  ORGANISATIONS_SUCCESS = '@@IATIWriterActionTypes/ORGANISATIONS_SUCCESS',
  ORGANISATIONS_FAILURE = '@@IATIWriterActionTypes/ORGANISATIONS_FAILURE',
  ORGANISATION_SUCCESS = '@@IATIWriterActionTypes/ORGANISATION_SUCCESS',
  ORGANISATION_FAILURE = '@@IATIWriterActionTypes/ORGANISATION_FAILURE',
  BUDGET_SUCCESS = '@@IATIWriterActionTypes/BUDGET_SUCCESS',
  BUDGET_FAILURE = '@@IATIWriterActionTypes/BUDGET_FAILURE',
  RECIPIENTORGBUDGET_SUCCESS = '@@IATIWriterActionTypes/RECIPIENTORGBUDGET_SUCCESS',
  RECIPIENTORGBUDGET_FAILURE = '@@IATIWriterActionTypes/RECIPIENTORGBUDGET_FAILURE',
  RECIPIENTREGIONBUDGET_SUCCESS = '@@IATIWriterActionTypes/RECIPIENTREGIONBUDGET_SUCCESS',
  RECIPIENTREGIONBUDGET_FAILURE = '@@IATIWriterActionTypes/RECIPIENTREGIONBUDGET_FAILURE',
  RECIPIENTCOUNTRYBUDGET_SUCCESS = '@@IATIWriterActionTypes/RECIPIENTCOUNTRYBUDGET_SUCCESS',
  RECIPIENTCOUNTRYBUDGET_FAILURE = '@@IATIWriterActionTypes/RECIPIENTCOUNTRYBUDGET_FAILURE',
  TOTALEXPENDITURE_SUCCESS = '@@IATIWriterActionTypes/TOTALEXPENDITURE_SUCCESS',
  TOTALEXPENDITURE_FAILURE = '@@IATIWriterActionTypes/TOTALEXPENDITURE_FAILURE',
  DOCUMENT_SUCCESS = '@@IATIWriterActionTypes/DOCUMENT_SUCCESS',
  DOCUMENT_FAILURE = '@@IATIWriterActionTypes/DOCUMENT_FAILURE',
  ACTIVITIES_SUCCESS = '@@IATIWriterActionTypes/ACTIVITIES_SUCCESS',
  ACTIVITIES_FAILURE = '@@IATIWriterActionTypes/ACTIVITIES_FAILURE',
  ACTIVITY_SUCCESS = '@@IATIWriterActionTypes/ACTIVITY_SUCCESS',
  ACTIVITY_FAILURE = '@@IATIWriterActionTypes/ACTIVITY_FAILURE',
  ACTIVITYADDITIONAL_SUCCESS = '@@IATIWriterActionTypes/ACTIVITYADDITIONAL_SUCCESS',
  ACTIVITYADDITIONAL_FAILURE = '@@IATIWriterActionTypes/ACTIVITYADDITIONAL_FAILURE'
}

export const enum IATIReportActionTypes {
  REPORT_INIT =  '@@IATIReportActionTypes/REPORT_INIT',
  ORGS_SUCCESS = '@@IATIReportActionTypes/ORGS_SUCCESS',
  ORGSPICKER_SUCCESS = '@@IATIReportActionTypes/ORGSPICKER_SUCCESS',
  ORGS_FAILURE = '@@IATIReportActionTypes/ORGS_FAILURE',
  ORGANISATIONS_SUCCESS = '@@IATIReportActionTypes/ORGANISATIONS_SUCCESS',
  ORGANISATIONSPICKER_SUCCESS = '@@IATIReportActionTypes/ORGANISATIONSPICKER_SUCCESS',
  ORGANISATIONS_FAILURE = '@@IATIReportActionTypes/ORGANISATIONS_FAILURE',
  ORGANISATION_SUCCESS = '@@IATIReportActionTypes/ORGANISATION_SUCCESS',
  ORGANISATIONPICKER_SUCCESS = '@@IATIReportActionTypes/ORGANISATIONPICKER_SUCCESS',
  ORGANISATION_FAILURE = '@@IATIReportActionTypes/ORGANISATION_FAILURE',
  BUDGET_SUCCESS = '@@IATIReportActionTypes/BUDGET_SUCCESS',
  BUDGET_FAILURE = '@@IATIReportActionTypes/BUDGET_FAILURE',
  RECIPIENTORGBUDGET_SUCCESS = '@@IATIReportActionTypes/RECIPIENTORGBUDGET_SUCCESS',
  RECIPIENTORGBUDGET_FAILURE = '@@IATIReportActionTypes/RECIPIENTORGBUDGET_FAILURE',
  RECIPIENTREGIONBUDGET_SUCCESS = '@@IATIReportActionTypes/RECIPIENTREGIONBUDGET_SUCCESS',
  RECIPIENTREGIONBUDGET_FAILURE = '@@IATIReportActionTypes/RECIPIENTREGIONBUDGET_FAILURE',
  RECIPIENTCOUNTRYBUDGET_SUCCESS = '@@IATIReportActionTypes/RECIPIENTCOUNTRYBUDGET_SUCCESS',
  RECIPIENTCOUNTRYBUDGET_FAILURE = '@@IATIReportActionTypes/RECIPIENTCOUNTRYBUDGET_FAILURE',
  TOTALEXPENDITURE_SUCCESS = '@@IATIReportActionTypes/TOTALEXPENDITURE_SUCCESS',
  TOTALEXPENDITURE_FAILURE = '@@IATIReportActionTypes/TOTALEXPENDITURE_FAILURE',
  DOCUMENT_SUCCESS = '@@IATIReportActionTypes/DOCUMENT_SUCCESS',
  DOCUMENT_FAILURE = '@@IATIReportActionTypes/DOCUMENT_FAILURE',
  ACTIVITIES_SUCCESS = '@@IATIReportActionTypes/ACTIVITIES_SUCCESS',
  ACTIVITIESPICKER_SUCCESS = '@@IATIReportActionTypes/ACTIVITIESPICKER_SUCCESS',
  ACTIVITIES_FAILURE = '@@IATIReportActionTypes/ACTIVITIES_FAILURE',
  ACTIVITY_SUCCESS = '@@IATIReportActionTypes/ACTIVITY_SUCCESS',
  ACTIVITYPICKER_SUCCESS = '@@IATIReportActionTypes/ACTIVITYPICKER_SUCCESS',
  ACTIVITY_FAILURE = '@@IATIReportActionTypes/ACTIVITY_FAILURE',
  ACTIVITYADDITIONAL_SUCCESS = '@@IATIReportActionTypes/ACTIVITYADDITIONAL_SUCCESS',
  ACTIVITYADDITIONAL_FAILURE = '@@IATIReportActionTypes/ACTIVITYADDITIONAL_FAILURE'
}
