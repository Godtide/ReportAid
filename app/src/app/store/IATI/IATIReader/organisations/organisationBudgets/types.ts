import { PayloadProps, DictData } from '../../../../types'

export interface IATIBudgetData {
  budgetKey: string
  budgetType: number
  budgetLine: string,
  otherRef: string,
  value: number,
  status: number,
  start: string,
  end: string
}

export interface IATIOrganisationBudgetReport {
  organisationsRef: string
  organisationRef: string
  data: Array<IATIBudgetData>
}

export interface IATIOrganisationBudgetReportProps extends PayloadProps {
  data: IATIOrganisationBudgetReport
}
