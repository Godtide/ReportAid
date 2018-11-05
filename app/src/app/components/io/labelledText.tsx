import * as React from 'react'

interface LabelledTextProps {
  label: string
  text: string
}

export const LabelledText: React.SFC<LabelledTextProps> = (props: LabelledTextProps) => {

  return (
      <p>{props.label} {props.text}</p>
  )
}
