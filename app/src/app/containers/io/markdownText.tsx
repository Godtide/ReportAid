import * as React from 'react'
import * as ReactMarkdown from 'react-markdown'

interface MarkdownTextProps {
  text: string
}

class MarkdownText extends React.Component<MarkdownTextProps> {

  render() {
    return (
      <div>
        <ReactMarkdown escapeHtml={false} source={this.props.text} />
      </div>
    )
  }
}

export default MarkdownText
