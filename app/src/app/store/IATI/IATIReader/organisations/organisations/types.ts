import { PayloadProps, DictData } from '../../../../types'
import { IATIOrganisationsProps } from '../../../types'

export interface IATIOrganisationsData extends DictData {
  [key: string]: IATIOrganisationsProps
}

export interface IATIOrganisationsReportProps extends PayloadProps {
  data: IATIOrganisationsData
}
