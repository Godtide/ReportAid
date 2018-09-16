import * as React from 'react'

interface HeadingInterface {
  heading: string
}

export type HeadingProps = HeadingInterface

const Heading: React.SFC<HeadingProps> = (props) => {

  return (
    <h2>{props.heading}</h2>
  )
}

export default Heading
