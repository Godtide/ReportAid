import { PayloadProps, DictData } from '../../../types'
import { IATIOrganisationsProps,
         IATIOrganisationBudgetProps,
         IATIOrganisationRecipientBudgetProps,
         IATIOrganisationRegionBudgetProps,
         IATIOrganisationCountryBudgetProps,
         IATIOrganisationExpenditureProps,
         IATIOrganisationDocProps } from '../../types'

 export interface OrganisationsReportProps {
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

export interface IATIReportProps extends PayloadProps {
  data: IATIOrganisationsReport
}

export const enum IATIReportActionTypes {
  ORGANISATIONS_SUCCESS = '@@OrganisationsReaderActionTypes/READORGANISATIONS_SUCCESS',
  ORGANISATIONS_FAILURE = '@@OrganisationsReaderActionTypes/READORGANISATIONS_FAILURE',
  ORGANISATION_SUCCESS = '@@OrganisationsReaderActionTypes/READORGANISATION_SUCCESS',
  ORGANISATION_FAILURE = '@@OrganisationsReaderActionTypes/READORGANISATION_FAILURE',
  BUDGET_SUCCESS = '@@OrganisationsReaderActionTypes/READBUDGET_SUCCESS',
  BUDGET_FAILURE = '@@OrganisationsReaderActionTypes/READBUDGET_FAILURE',
  RECIPIENTORGBUDGET_SUCCESS = '@@OrganisationsReaderActionTypes/READRECIPIENTORGBUDGET_SUCCESS',
  RECIPIENTORGBUDGET_FAILURE = '@@OrganisationsReaderActionTypes/READRECIPIENTORGBUDGET_FAILURE',
  RECIPIENTREGIONBUDGET_SUCCESS = '@@OrganisationsReaderActionTypes/READRECIPIENTREGIONBUDGET_SUCCESS',
  RECIPIENTREGIONBUDGET_FAILURE = '@@OrganisationsReaderActionTypes/READRECIPIENTREGIONBUDGET_FAILURE',
  RECIPIENTCOUNTRYBUDGET_SUCCESS = '@@OrganisationsReaderActionTypes/READRECIPIENTCOUNTRYBUDGET_SUCCESS',
  RECIPIENTCOUNTRYBUDGET_FAILURE = '@@OrganisationsReaderActionTypes/READRECIPIENTCOUNTRYBUDGET_FAILURE',
  TOTALEXPENDITURE_SUCCESS = '@@OrganisationsReaderActionTypes/READTOTALEXPENDITURE_SUCCESS',
  TOTALEXPENDITURE_FAILURE = '@@OrganisationsReaderActionTypes/READTOTALEXPENDITURE_FAILURE',
  DOCUMENT_SUCCESS = '@@OrganisationsReaderActionTypes/READDOCUMENT_SUCCESS',
  DOCUMENT_FAILURE = '@@OrganisationsReaderActionTypes/READDOCUMENT_FAILURE'
}
