import { PayloadProps, DictData } from '../types'

export interface InfoProps {
  title: string
  data: string
}

export interface InfoData {
  about: InfoProps
  help: InfoProps
  home: InfoProps
  IATIReader: InfoProps
  IATIWriter: InfoProps
  overview: InfoProps
}

export interface InfoPageProps extends PayloadProps {
  data: InfoData
}
