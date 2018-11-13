import * as React from 'react'
import Markdown from 'react-markdown'

interface MarkdownTextProps {
  text: string
}

export class MarkdownText extends React.Component<MarkdownTextProps> {

  render() {
    //console.log(this.props.text)
    return (
      <Markdown escapeHtml={false} source={this.props.text} />
      //<p>{this.props.text}</p>
    )
  }
}
