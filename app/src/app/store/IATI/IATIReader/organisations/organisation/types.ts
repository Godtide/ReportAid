import { PayloadProps, DictData } from '../../../../types'
import { IATIOrganisationProps } from '../../../types'

export interface IATIOrganisationData extends DictData {
  [key: string]: IATIOrganisationProps
}

export interface IATIOrganisationDataProps extends PayloadProps {
  data: IATIOrganisationData
}

export interface IATIOrganisationReport extends DictData {
  [key: string]: IATIOrganisationDataProps
}

export interface IATIOrganisationReportProps extends PayloadProps {
  data: IATIOrganisationReport
}
