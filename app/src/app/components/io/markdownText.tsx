import * as React from 'react'
import * as Markdown from 'react-markdown'

interface MarkdownTextProps {
  text: string
}

const MarkdownText: React.SFC<MarkdownTextProps> = (props: MarkdownTextProps) => {

  return (
    <div>
      <Markdown escapeHtml={false} source={props.text} />
    </div>
  )
}

export default MarkdownText
