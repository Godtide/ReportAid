import * as React from 'react'

interface LabelledTextInterface {
  label: string
  text: string
}

export type LabelledTextProps = LabelledTextInterface

const LabelledText: React.SFC<LabelledTextProps> = (props) => {

  return (
      <p>{props.label} {props.text}</p>
  )
}

export default LabelledText
