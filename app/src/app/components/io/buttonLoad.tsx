import * as React from 'react'
import { Tooltip, Button } from 'antd'

interface ButtonLoadInterface {
  label: string
  icon: string
  tip: string
  loading: boolean
}

export type ButtonLoadProps = ButtonLoadInterface

const ButtonLoad: React.SFC<ButtonLoadProps> = (props) => {

  return (
    <Tooltip title={props.tip}>
      <Button
        icon={props.icon}
        loading={props.loading}
        onClick={}
      >
          {this.props.label}
      </Button>
    </Tooltip>
  )
}

export default ButtonLoad
