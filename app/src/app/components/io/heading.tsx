import * as React from 'react'

interface HeadingProps {
  heading: string
}

export const Heading: React.SFC<HeadingProps> = (props: HeadingProps) => {

  return (
    <h2>{props.heading}</h2>
  )
}
