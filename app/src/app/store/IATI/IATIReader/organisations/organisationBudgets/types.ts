import { PayloadProps, DictData } from '../../../../types'
import { IATIOrganisationBudgetProps } from '../../../types'

export interface IATIBudgetData {
  [key: string]: IATIOrganisationBudgetProps
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
