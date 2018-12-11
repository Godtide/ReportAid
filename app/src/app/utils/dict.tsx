import { DictData } from '../store/types'

export const getDictEntries = (props: DictData): object[] => {
  let xs: object[] = [{}]
  let index = 0
  Object.values(props).forEach((value) => {
    xs[index] = value
    index += 1
  })
  return xs
}
