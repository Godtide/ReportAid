import * as React from 'react'
import Markdown from 'react-markdown'

interface MarkdownTextProps {
  text: string
}

export const MarkdownText: React.SFC<MarkdownTextProps> = (props: MarkdownTextProps) => {

  return (
    <Markdown escapeHtml={false} source={props.text} />
  )
}
