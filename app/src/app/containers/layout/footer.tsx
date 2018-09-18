import * as React from 'react'

import { FooterStrings } from '../../utils/strings'

// { textAlign: 'center' }

export type FooterProps = FooterStrings

const Footer: React.SFC<FooterProps> = (props) => {

  return (
    <h5>{props.copyright}</h5>
  )
}

export default Footer
