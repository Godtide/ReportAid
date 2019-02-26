import { PayloadProps, DictData } from '../../../../types'
import { IATIBudgetProps as BudgetProps } from '../../../types'

export interface IATIBudgetData {
  [key: string]: BudgetProps
}

export interface IATIBudgetProps extends PayloadProps {
  data: IATIBudgetData
}

export interface IATIOrganisationBudgetData extends DictData {
  [key: string]: IATIBudgetProps
}

export interface IATIOrganisationBudgetDataProps extends PayloadProps {
  data: IATIOrganisationBudgetData
}

export interface IATIOrganisationBudgetReport extends DictData {
  [key: string]: IATIOrganisationBudgetDataProps
}

export interface IATIOrganisationBudgetReportProps extends PayloadProps {
  data: IATIOrganisationBudgetReport
}
