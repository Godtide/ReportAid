import * as React from 'react'

interface TaglineInterface {
  tagLine: string
}

export type TaglineProps = TaglineInterface

const Tagline: React.SFC<TaglineProps> = (props) => {

  return (
    <h1>{props.tagLine}</h1>
  )
}

export default Tagline
