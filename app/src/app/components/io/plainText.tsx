import * as React from 'react'

interface PlainTextInterface {
  text: string
}

export type PlainTextProps = PlainTextInterface

const PlainText: React.SFC<PlainTextProps> = (props) => {

  return (
    <p>{props.text}</p>
  )
}

export default PlainText
