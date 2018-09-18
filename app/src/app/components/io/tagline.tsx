import * as React from 'react'

interface TaglineProps {
  tagLine: string
}

const Tagline: React.SFC<TaglineProps> = (props) => {

  return (
    <h1>{props.tagLine}</h1>
  )
}

export default Tagline
