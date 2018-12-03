import { DictData } from '../../store/types'

export const getDictEntries = (props: DictData): string => {
  let orgs: string = ``
  Object.keys(props).forEach((key) => {
    let length = 0
    orgs += `**Key**: ${key}, `
    const entries = Object.entries(props[key])
    entries.forEach((entry) => {
      orgs += `**${entry[0]}**: ${entry[1]}`
      length += 1
      length == entries.length ? orgs += `<br />`: orgs += `, `
    })
  })
  return orgs
}
