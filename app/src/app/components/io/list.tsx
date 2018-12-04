export const getList = (props: String[]): string => {
  let xs: string = ``
  props.forEach((value) => {
    xs += `${value}<br />`
  })
  return xs
}

export const getKeyedList = (props: Object): string => {
  let xs: string = ``
  Object.entries(props).forEach((entry) => {
    xs += `**${entry[0]}**: ${entry[1]}<br />`
  })
  return xs
}
