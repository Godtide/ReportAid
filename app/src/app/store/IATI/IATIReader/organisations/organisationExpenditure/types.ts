import { PayloadProps, DictData } from '../../../../types'
import { IATIOrganisationExpenditureProps } from '../../../types'

export interface IATIExpenditureData {
  [key: string]: IATIOrganisationExpenditureProps
}

export interface IATIExpenditureProps extends PayloadProps {
  data: IATIExpenditureData
}

export interface IATIOrganisationExpenditureData extends DictData {
  [key: string]: IATIExpenditureProps
}

export interface IATIOrganisationExpenditureDataProps extends PayloadProps {
  data: IATIOrganisationExpenditureData
}

export interface IATIOrganisationExpenditureReport extends DictData {
  [key: string]: IATIOrganisationExpenditureDataProps
}

export interface IATIOrganisationExpenditureReportProps extends PayloadProps {
  data: IATIOrganisationExpenditureReport
}
