import * as React from 'react'
import { NavLink } from 'react-router-dom'
import { Layout } from 'antd'
import LayoutContainer from '../../containers/pages/layoutContainer'

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
