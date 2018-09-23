// Separate state props + dispatch props to their own interfaces.
export interface InfoProps {
  title: string
  data: string
}

export const enum AboutActionTypes {
  REQ_DATA = '@@aboutTypes/REQ_DATA'
}

export const enum HomeActionTypes {
  REQ_DATA = '@@homeTypes/REQ_DATA'
}

export const enum OverviewActionTypes {
  REQ_DATA = '@@overviewTypes/REQ_DATA'
}

export const enum HelpActionTypes {
  REQ_DATA = '@@overviewTypes/REQ_DATA'
}
