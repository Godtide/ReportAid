import { DictData } from '../../store/types'

export const getDictEntries = (props: DictData): string => {
  let xs: string = ``
  Object.keys(props).forEach((key) => {
    let length = 0
    xs += `**Key**: ${key}, `
    const entries = Object.entries(props[key])
    entries.forEach((entry) => {
      xs += `**${entry[0]}**: ${entry[1]}`
      length += 1
      length == entries.length ? xs += `<br />`: xs += `, `
    })
  })
  return xs
}
