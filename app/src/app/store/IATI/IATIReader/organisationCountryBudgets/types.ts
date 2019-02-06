import { PayloadProps, DictData } from '../../../types'
import { IATIOrganisationCountryBudgetProps } from '../../types'

export interface CountryBudgetData {
  [key: string]: IATIOrganisationCountryBudgetProps
}

export interface CountryBudgets {
  num: number
  data: CountryBudgetData
}

export interface OrganisationCountryBudgetsData extends DictData {
  [key: string]: CountryBudgets
}

export interface OrganisationCountryBudgetsReaderProps extends PayloadProps {
  num: number
  data: OrganisationCountryBudgetsData
}

export const enum OrganisationCountryBudgetsReaderActionTypes {
  NUM_SUCCESS = '@@OrganisationCountryBudgetsReaderAction/GETNUM_SUCCESS',
  NUM_FAILURE = '@@OrganisationCountryBudgetsReaderAction/GETNUM_FAILURE',
  EXISTS_SUCCESS = '@@OrganisationCountryBudgetsReaderAction/GETEXISTS_SUCCESS',
  EXISTS_FAILURE = '@@OrganisationCountryBudgetsReaderAction/GETEXISTS_FAILURE',
  REF_SUCCESS = '@@OrganisationCountryBudgetsReaderAction/GETREFERENCE_SUCCESS',
  REF_FAILURE = '@@OrganisationCountryBudgetsReaderAction/GETREFERENCE_FAILURE',
  NUMBUDGET_SUCCESS = '@@OrganisationCountryBudgetsReaderAction/GETNUMBUDGET_SUCCESS',
  NUMBUDGET_FAILURE = '@@OrganisationCountryBudgetsReaderAction/GETNUMBUDGET_FAILURE',
  BUDGETREF_SUCCESS = '@@OrganisationCountryBudgetsReaderAction/GETBUDGETREF_SUCCESS',
  BUDGETREF_FAILURE = '@@OrganisationCountryBudgetsReaderAction/GETBUDGETREF_FAILURE',
  BUDGET_SUCCESS = '@@OrganisationCountryBudgetsReaderAction/GETBUDGET_SUCCESS',
  BUDGET_FAILURE = '@@OrganisationCountryBudgetsReaderAction/GETBUDGET_FAILURE'
}
