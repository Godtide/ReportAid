import * as React from 'react'

interface HeadingProps {
  heading: string
}

const Heading: React.SFC<HeadingProps> = (props) => {

  return (
    <h2>{props.heading}</h2>
  )
}

export default Heading
