import { PayloadProps, DictData } from '../../../types'
import { IATIOrganisationExpenditureProps } from '../../types'

export interface ExpenditureData {
  [key: string]: IATIOrganisationExpenditureProps
}

export interface Expenditure {
  num: number
  data: ExpenditureData
}

export interface OrganisationExpenditureData extends DictData {
  [key: string]: Expenditure
}

export interface OrganisationExpenditureReaderProps extends PayloadProps {
  num: number
  data: OrganisationExpenditureData
}

export const enum OrganisationExpenditureReaderActionTypes {
  NUM_SUCCESS = '@@OrganisationExpenditureReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrganisationExpenditureReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrganisationExpenditureReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrganisationExpenditureReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrganisationExpenditureReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrganisationExpenditureReaderAction/GETREFERENCE_FAILURE',
  NUMEXPENDITURE_SUCCESS = '@@OrganisationExpenditureReaderAction/GETNUMEXPENDITURE_SUCCESS',
  NUMEXPENDITURE_FAILURE = '@@OrganisationExpenditureReaderAction/GETNUMEXPENDITURE_FAILURE',
  EXPENDITUREREF_SUCCESS = '@@OrganisationExpenditureReaderAction/GETEXPENDITUREREF_SUCCESS',
  EXPENDITUREREF_FAILURE = '@@OrganisationExpenditureReaderAction/GETEXPENDITUREREF_FAILURE',
  EXPENDITURE_SUCCESS = '@@OrganisationExpenditureReaderAction/GETEXPENDITURE_SUCCESS',
  EXPENDITURE_FAILURE = '@@OrganisationExpenditureReaderAction/GETEXPENDITURE_FAILURE'
}
