import { DictData } from '../../store/types'

export const getDictEntries = (props: DictData): string => {
  let orgs: string = ``
  Object.keys(props).forEach((key) => {
    orgs += `<span key=${key}><p><strong>${key}: </strong>`
    let length = 0
    const entries = Object.entries(props[key])
    entries.forEach((entry) => {
      orgs += `${entry[0]} - ${entry[1]}`
      length += 1
      length == entries.length ? orgs: orgs += `, `
    })
    orgs += `</p></span>`
  })
  return orgs
}
