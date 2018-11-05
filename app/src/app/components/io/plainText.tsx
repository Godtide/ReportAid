import * as React from 'react'

interface PlainTextProps {
  text: string
}

export const PlainText: React.SFC<PlainTextProps> = (props: PlainTextProps) => {

  return (
    <p>{props.text}</p>
  )
}
