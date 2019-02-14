import { PayloadProps, DictData } from '../../../../types'
import { IATIOrganisationDocProps } from '../../../types'

export interface IATIDocData {
  [key: string]: IATIOrganisationDocProps
}

export interface IATIDocProps extends PayloadProps {
  data: IATIDocData
}

export interface IATIOrganisationDocData extends DictData {
  [key: string]: IATIDocProps
}

export interface IATIOrganisationDocDataProps extends PayloadProps {
  data: IATIOrganisationDocData
}

export interface IATIOrganisationDocReport extends DictData {
  [key: string]: IATIOrganisationDocDataProps
}

export interface IATIOrganisationDocReportProps extends PayloadProps {
  data: IATIOrganisationDocReport
}
