import * as React from 'react'
import { Card } from 'antd'

interface LoggerInterface {
  log: string[]
}

export type LoggerProps = LoggerInterface

const Logger: React.SFC<LoggerProps> = (props) => {

  let logs = props.log.map((text, index) =>
    <span key={index}>
      {text}
      <br />
    </span>
  )

  return (
      <Card>
        <p>{logs}</p>
      </Card>
  )
}

export default Logger
