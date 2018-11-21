import * as React from 'react'
import Markdown from 'react-markdown'

interface MarkdownTextProps {
  text: string
}

interface MarkdownTextWithTitleProps {
  title: string
  text: string
}

interface MarkdownListWithTitleProps {
  title: string
  text: string[]
}

interface MarkdownListProps {
  text: string[]
}

export const MarkdownList: React.SFC<MarkdownListProps> = (props: MarkdownListProps) => {

  let logs = props.text.map((text, index) =>
    <span key={index}>
      <Markdown escapeHtml={false} source={text} />
    </span>
  )

  return (
    <div>
      {logs}
    </div>
  )
}

export const MarkdownListWithTitle: React.SFC<MarkdownListWithTitleProps> = (props: MarkdownListWithTitleProps) => {

  let logs = props.text.map((text, index) =>
    <span key={index}>
      <Markdown escapeHtml={false} source={text} />
    </span>
  )

  return (
    <div>
      <h2>{props.title}</h2>
      {logs}
    </div>
  )
}

export const MarkdownText: React.SFC<MarkdownTextProps> = (props: MarkdownTextProps) => {

  return (
    <Markdown escapeHtml={false} source={props.text} />
  )
}

export const MarkdownTextWithTitle: React.SFC<MarkdownTextWithTitleProps> = (props: MarkdownTextWithTitleProps) => {

  return (
    <div>
      <h2>{props.title}</h2>
      <Markdown escapeHtml={false} source={props.text} />
    </div>
  )
}
