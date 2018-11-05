import * as React from 'react'

interface TaglineProps {
  tagLine: string
}

export const Tagline: React.SFC<TaglineProps> = (props: TaglineProps) => {

  return (
    <h1>{props.tagLine}</h1>
  )
}
