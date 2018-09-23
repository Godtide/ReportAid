// Separate state props + dispatch props to their own interfaces.
export interface AboutProps {
  title: string
  data: string
}

export const enum ActionTypes {
  REQ_DATA = '@@aboutTypes/REQ_DATA'
}
