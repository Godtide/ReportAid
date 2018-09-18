import * as React from 'react'

interface TitleProps {
  title: string
}

const Title: React.SFC<TitleProps> = (props: TitleProps) => {

  return (
      <div>
        <h1>{props.title}</h1>
      </div>
    )
}

export default Title
