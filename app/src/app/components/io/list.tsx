export const getList = (props: String[]): string => {
  let list: string = ``
  props.forEach((value) => {
    list += `${value}<br />`
  })
  return list
}

export const getKeyedList = (props: Object): string => {
  let list: string = ``
  Object.entries(props).forEach((entry) => {
    list += `**${entry[0]}**: ${entry[1]}<br />`
  })
  return list
}
