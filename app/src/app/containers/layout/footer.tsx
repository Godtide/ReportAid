import * as React from 'react'
import { Layout } from 'antd'

import { FooterStrings } from '../../utils/strings'

// { textAlign: 'center' }

export type FooterProps = FooterStrings

const Footer: React.SFC<FooterProps> = (props) => {

  return (
    <div>
      <Layout style={{ textAlign: 'center' }}>
        <h5>{props.copyright}</h5>
      </Layout>
    </div>
  )
}

export default Footer
