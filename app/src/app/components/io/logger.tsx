import * as React from 'react'
import Card from '@material-ui/core/Card'

interface LoggerProps {
  log: string[]
}

const Logger: React.SFC<LoggerProps> = (props: LoggerProps) => {

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
