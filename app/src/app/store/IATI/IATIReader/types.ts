import { PayloadProps, DictData } from '../../types'
import { IATIOrganisationsProps,
         IATIOrganisationBudgetProps,
         IATIOrganisationRecipientBudgetProps,
         IATIOrganisationRegionBudgetProps,
         IATIOrganisationCountryBudgetProps,
         IATIOrganisationExpenditureProps,
         IATIOrganisationDocProps } from '../types'

 export interface ReportProps {
   organisationsRef: string
   organisationRef: string
 }

export interface IATIBudgetReportProps {
  [key: string]: IATIOrganisationBudgetProps
}

export interface IATIRecipientOrgBudgetReportProps {
  [key: string]: IATIOrganisationRecipientBudgetProps
}

export interface IATIRecipientRegionBudgetReportProps {
  [key: string]: IATIOrganisationRegionBudgetProps
}

export interface IATIRecipientCountryBudgetReportProps {
  [key: string]: IATIOrganisationCountryBudgetProps
}

export interface IATITotalExpenditureReportProps {
  [key: string]: IATIOrganisationExpenditureProps
}

export interface IATIDocumentReportProps {
  [key: string]: IATIOrganisationDocProps
}

export interface IATIOrganisationReportProps extends DictData {
  IATIOganisation: IATIOrganisationProps
  totalBudget: IATIBudgetReportProps
  recipientOrgBudget: IATIRecipientOrgBudgetReportProps
  recipientRegionBudget: IATIRecipientRegionBudgetReportProps
  recipientCountryBudget: IATIRecipientCountryBudgetReportProps
  totalExpenditure: IATITotalExpenditureReportProps
  document: IATIDocumentReportProps
}

export interface IATIOrganisationsData extends DictData {
  IATIOganisations: IATIOrganisationsProps
  [key: string]: IATIOrganisationReportProps
}

export interface IATIOrganisationsReport extends DictData {
  [key: string]: IATIOrganisationsData
}

export interface IATIOrganisationsReportProps extends PayloadProps {
  data: IATIOrganisationsReport
}

export const enum OrganisationsReportActionTypes {
  ORGANISATIONS_SUCCESS = '@@OrganisationsReaderActionTypes/GETORGANISATIONS_SUCCESS',
  ORGANISATIONS_FAILURE = '@@OrganisationsReaderActionTypes/GETORGANISATIONS_FAILURE',
  ORGANISATION_SUCCESS = '@@OrganisationsReaderActionTypes/GETORGANISATION_SUCCESS',
  ORGANISATION_FAILURE = '@@OrganisationsReaderActionTypes/GETORGANISATION_FAILURE',
  BUDGET_SUCCESS = '@@OrganisationsReaderActionTypes/GETBUDGET_SUCCESS',
  BUDGET_FAILURE = '@@OrganisationsReaderActionTypes/GETBUDGET_FAILURE',
  RECIPIENTORGBUDGET_SUCCESS = '@@OrganisationsReaderActionTypes/GETRECIPIENTORGBUDGET_SUCCESS',
  RECIPIENTORGBUDGET_FAILURE = '@@OrganisationsReaderActionTypes/GETRECIPIENTORGBUDGET_FAILURE',
  RECIPIENTREGIONBUDGET_SUCCESS = '@@OrganisationsReaderActionTypes/GETRECIPIENTREGIONBUDGET_SUCCESS',
  RECIPIENTREGIONBUDGET_FAILURE = '@@OrganisationsReaderActionTypes/GETRECIPIENTREGIONBUDGET_FAILURE',
  RECIPIENTCOUNTRYBUDGET_SUCCESS = '@@OrganisationsReaderActionTypes/GETRECIPIENTCOUNTRYBUDGET_SUCCESS',
  RECIPIENTCOUNTRYBUDGET_FAILURE = '@@OrganisationsReaderActionTypes/GETRECIPIENTCOUNTRYBUDGET_FAILURE',
  TOTALEXPENDITURE_SUCCESS = '@@OrganisationsReaderActionTypes/GETTOTALEXPENDITURE_SUCCESS',
  TOTALEXPENDITURE_FAILURE = '@@OrganisationsReaderActionTypes/GETTOTALEXPENDITURE_FAILURE',
  DOCUMENT_SUCCESS = '@@OrganisationsReaderActionTypes/GETDOCUMENT_SUCCESS',
  DOCUMENT_FAILURE = '@@OrganisationsReaderActionTypes/GETDOCUMENT_FAILURE'
}
