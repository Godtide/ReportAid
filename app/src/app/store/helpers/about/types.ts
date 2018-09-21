// Separate state props + dispatch props to their own interfaces.
export interface AboutProps {
  readonly title: string
  readonly data: string
}

export const enum AboutActionTypes {
  FETCH_REQUEST = '@@aboutTypes/FETCH_REQUEST',
  REQ_DATA = '@@aboutTypes/REQ_DATA',
  RES_DATA = '@@aboutTypes/RES_DATA',
  FAIL_DATA = '@@aboutTypes/FAIL_DATA'
}
