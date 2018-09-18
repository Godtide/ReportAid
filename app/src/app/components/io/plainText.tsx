import * as React from 'react'

interface PlainTextProps {
  text: string
}

const PlainText: React.SFC<PlainTextProps> = (props: PlainTextProps) => {

  return (
    <p>{props.text}</p>
  )
}

export default PlainText
