import { PayloadProps, DictData } from '../../../../types'
import { IATIOrganisationRegionBudgetProps } from '../../../types'

export interface IATIRegionBudgetData {
  [key: string]: IATIOrganisationRegionBudgetProps
}

export interface IATIRegionBudgetProps extends PayloadProps {
  data: IATIRegionBudgetData
}

export interface IATIOrganisationRegionBudgetData extends DictData {
  [key: string]: IATIRegionBudgetProps
}

export interface IATIOrganisationRegionBudgetDataProps extends PayloadProps {
  data: IATIOrganisationRegionBudgetData
}

export interface IATIOrganisationRegionBudgetReport extends DictData {
  [key: string]: IATIOrganisationRegionBudgetDataProps
}

export interface IATIOrganisationRegionBudgetReportProps extends PayloadProps {
  data: IATIOrganisationRegionBudgetReport
}
