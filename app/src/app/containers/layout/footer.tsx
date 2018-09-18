import * as React from 'react'

interface FooterProps {
  copyright: string
}

const Footer: React.SFC<FooterProps> = (props: FooterProps) => {

  return (
    <h5>{props.copyright}</h5>
  )
}

export default Footer
