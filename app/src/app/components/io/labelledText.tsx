import * as React from 'react'

interface LabelledTextProps {
  label: string
  text: string
}

const LabelledText: React.SFC<LabelledTextProps> = (props) => {

  return (
      <p>{props.label} {props.text}</p>
  )
}

export default LabelledText
