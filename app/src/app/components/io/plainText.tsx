import * as React from 'react'

interface PlainTextWithTitleProps {
  title: string
  text: string
}

interface PlainTextListWithTitleProps {
  title: string
  text: string[]
}

interface PlainTextListProps {
  text: string[]
}

interface PlainTextKeyedListProps {
  list: object
}

interface PlainTextKeyedWithTitleListProps {
  title: string
  list: object
}

export const PlainTextKeyedList: React.SFC<PlainTextKeyedListProps> = (props: PlainTextKeyedListProps) => {

  const keyedList = Object.entries(props.list).map(([key,value]) =>
    <span key={key}>
      <p><b>{key}</b>: {value.toString()}</p>
    </span>
  )

  return (
    <div>
      {keyedList}
    </div>
  )
}

export const PlainTextKeyedWithTitleList: React.SFC<PlainTextKeyedWithTitleListProps> = (props: PlainTextKeyedWithTitleListProps) => {

  const keyedList = Object.entries(props.list).map(([key,value]) =>
    <span key={key}>
      <p><b>{key}</b>: {value.toString()}</p>
    </span>
  )

  return (
    <div>
      <h2>{props.title}</h2>
      {keyedList}
    </div>
  )
}

export const PlainTextList: React.SFC<PlainTextListProps> = (props: PlainTextListProps) => {

  let logs = props.text.map((text, index) =>
    <span key={index}>
      <p>{text}</p>
    </span>
  )

  return (
    <div>
      {logs}
    </div>
  )
}

export const PlainTextListWithTitle: React.SFC<PlainTextListWithTitleProps> = (props: PlainTextListWithTitleProps) => {

  let logs = props.text.map((text, index) =>
    <span key={index}>
      <p>{text}</p>
    </span>
  )

  return (
    <div>
      <h2>{props.title}</h2>
      {logs}
    </div>
  )
}

export const PlainTextWithTitle: React.SFC<PlainTextWithTitleProps> = (props: PlainTextWithTitleProps) => {

  return (
    <div>
      <h2>{props.title}</h2>
      <p>{props.text}</p>
    </div>
  )
}
