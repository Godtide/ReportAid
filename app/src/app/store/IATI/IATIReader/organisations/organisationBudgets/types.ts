import { PayloadProps } from '../../../../types'

export interface IATIBudgetData {
  budgetKey: string
  budgetLine: string,
  value: number,
  status: number,
  start: string,
  end: string
}

export interface IATIBudgetReport {
  organisationsRef: string
  organisationRef: string
  data: Array<IATIBudgetData>
}

export interface IATIBudgetReportProps extends PayloadProps {
  data: IATIBudgetReport
}
