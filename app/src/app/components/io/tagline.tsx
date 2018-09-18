import * as React from 'react'

interface TaglineProps {
  tagLine: string
}

const Tagline: React.SFC<TaglineProps> = (props: TaglineProps) => {

  return (
    <h1>{props.tagLine}</h1>
  )
}

export default Tagline
