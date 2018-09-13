import * as React from 'react'
import { Layout } from 'antd'

interface FooterProps {
  copyright: string
  style: string
}

// { textAlign: 'center' }

const Footer: React.SFC<FooterProps> = ({ copyright, style }) => (

  const { Footer } = Layout

  return (
    <div>
      <Footer style={{style}}>
        <h5>{copyright}</h5>
      </Footer>
    </div>
)

export default Footer
