import { PayloadProps, DictData } from '../../../../types'
import { IATIBudgetProps as BudgetProps } from '../../../types'

export interface IATIRecipientBudgetData {
  [key: string]: BudgetProps
}

export interface IATIRecipientBudgetProps extends PayloadProps {
  data: IATIRecipientBudgetData
}

export interface IATIOrganisationRecipientBudgetData extends DictData {
  [key: string]: IATIRecipientBudgetProps
}

export interface IATIOrganisationRecipientBudgetDataProps extends PayloadProps {
  data: IATIOrganisationRecipientBudgetData
}

export interface IATIOrganisationRecipientBudgetReport extends DictData {
  [key: string]: IATIOrganisationRecipientBudgetDataProps
}

export interface IATIOrganisationRecipientBudgetReportProps extends PayloadProps {
  data: IATIOrganisationRecipientBudgetReport
}
