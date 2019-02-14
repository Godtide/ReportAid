import { PayloadProps, DictData } from '../../../../types'
import { IATIOrganisationCountryBudgetProps } from '../../../types'

export interface IATICountryBudgetData {
  [key: string]: IATIOrganisationCountryBudgetProps
}

export interface IATICountryBudgetProps extends PayloadProps {
  data: IATICountryBudgetData
}

export interface IATIOrganisationCountryBudgetData extends DictData {
  [key: string]: IATICountryBudgetProps
}

export interface IATIOrganisationCountryBudgetDataProps extends PayloadProps {
  data: IATIOrganisationCountryBudgetData
}

export interface IATIOrganisationCountryBudgetReport extends DictData {
  [key: string]: IATIOrganisationCountryBudgetDataProps
}

export interface IATIOrganisationCountryBudgetReportProps extends PayloadProps {
  data: IATIOrganisationCountryBudgetReport
}
