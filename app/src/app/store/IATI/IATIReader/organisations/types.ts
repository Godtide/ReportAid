import { PayloadProps, DictData } from '../../../types'
import { IATIOrganisationsProps,
         IATIOrganisationProps,
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

//TS2411: Property 'data' of type 'IATIOrganisationDataProps' is not assignable to string index type 'IATIOrganisationProps'.

export interface IATIOrganisationReportProps extends PayloadProps {
  [key: string]: object
  totalBudget: IATIBudgetReportProps
  recipientOrgBudget: IATIRecipientOrgBudgetReportProps
  recipientRegionBudget: IATIRecipientRegionBudgetReportProps
  recipientCountryBudget: IATIRecipientCountryBudgetReportProps
  totalExpenditure: IATITotalExpenditureReportProps
  document: IATIDocumentReportProps
}

export interface IATIOrganisationsReportData extends PayloadProps {
  [key: string]: IATIOrganisationsProps
}



export interface IATIOrganisationsReport extends PayloadProps {
  IATIOrgtanisations: object
  data: IATIOrganisationReportProps
}

export interface IATIOrganisationsReportProps extends PayloadProps {
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
