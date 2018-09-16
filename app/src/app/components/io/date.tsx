import * as React from 'react'
import { Tooltip, DatePicker } from 'antd'

interface DateInterface {
  tip: string
  icon: string
  format: string
  placeholder: string
}

export type DateProps = DateInterface

const Date: React.SFC<DateProps> = (props) => {

  return (
    <Tooltip title={props.tip}>
      <DatePicker
        icon={props.icon}
        placeholder={props.placeholder}
        format={props.format}
        onChange={}
      />
    </Tooltip>
  )
}

export default Date
