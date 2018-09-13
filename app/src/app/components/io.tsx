import * as React from 'react'
import * as ReactDOM from 'react-dom'
import { RouteComponentProps } from 'react-router'

export namespace AppHeading {
  export interface Props extends RouteComponentProps<void> {
    heading: 'blah'
  }
}

export class AppHeading extends React.Component<AppHeading.Props> {

  constructor(props: AppHeading.Props, context?: any) {
    super(props, context)
  }

  render () {

    const { heading } = this.props

    return (
      <div>
        <h1>{heading}</h1>
      </div>
    )
  }
}

class Tagline extends React.Component {

  constructor (props) {
    super(props)
  }

  render () {
    return (
      <div>
        <h1>{this.props.tagLine}</h1>
      </div>
    )
  }
}

class Heading extends React.Component {

  render () {
    return (
      <div>
        <h2>{this.props.heading}</h2>
      </div>
    )
  }
}

class RadOptns extends React.Component {

  constructor(props) {
    super(props)

    this.optns = this.props.optns

    this.state = {
      optn: this.optns[0]
    }
  }

  _handleChange (e) {
    this.setState({
      optn: e.target.value
    })
    this.props.parentFunc(e.target.value)
  }

  render () {

    const RadGroup = Rad.Group
    return (
      <div>
        {this.props.label}
        <Tooltip title={this.props.tip}>
          <RadGroup
            optns={this.optns}
            onChange={this._handleChange.bind(this)}
            value={this.state.optn}
          />
        </Tooltip>
      </div>
    )
  }
}

class PlainTextOutput extends React.Component {

  constructor(props) {
    super(props)
  }

  render () {
    return (
      <div>
        <ReactMarkdown escapeHtml={false} source={this.props.text} />
      </div>
    )
  }
}

class TextOutput extends React.Component {

render () {
    return (
      <div>
        <p>{this.props.label} {this.props.text}</p>
      </div>
    )
  }
}

class TextInput extends React.Component {

  _handleChange (e) {
    this.props.parentFunc(e.target.value)
  }

  render () {
    return (
      <div>
        <Tooltip title={this.props.tip}>
          <Input
            addonBefore={this.props.label}
            defaultValue={this.props.placeHolder}
            onChange={this._handleChange.bind(this)}
          />
        </Tooltip>
      </div>
    )
  }
}

class TextAreaInput extends React.Component {

  _handleChange (e) {
    //console.log(e.target.value)
    this.props.parentFunc(e.target.value)
  }

  render () {

    const { TextArea } = Input

    return (
      <div>
        <Tooltip title={this.props.tip}>
          <p>{this.props.label}
            <TextArea
              rows={this.props.numRows}
              onChange={this._handleChange.bind(this)}
            />
          </p>
        </Tooltip>
      </div>
    )
  }
}

class Date extends React.Component {

  constructor(props) {
    super(props)
    this.format = 'DD/MM/YYYY'
  }

  _handleChange (date) {
    this.props.parentFunc(date._d.toGMTString())
  }

  render () {
    return (
      <div>
        <Tooltip title={this.props.tip}>
          <DatePicker
            icon='calendar'
            placeholder={this.props.label}
            format={this.format}
            onChange={this._handleChange.bind(this)}
          />
        </Tooltip>
      </div>
    )
  }
}

class Select extends React.Component {

  constructor(props) {
    super(props)
  }

  _handleChange (value) {
    this.props.parentFunc(value)
  }

  render () {

    const Optn = Select.Optn
    let children = []
    let size = this.props.selectns.length
    for (let i = 0; i < size; i++) {
      children.push(<Optn key={this.props.selectns[i]}>{this.props.selectns[i]}</Optn>)
    }

    return (
      <Tooltip title={this.props.tip} positn="top">
        <Select placeholder={this.props.label}
                style={{ width: '100%' }}
                onChange={this._handleChange.bind(this)}
                tokenSeparators={[',']}
        >
          {children}
        </Select>
      </Tooltip>
    )
  }
}

class Logger extends React.Component {

  constructor(props) {
    super(props)
  }

  render () {
    let logs = this.props.log.map((text, index) =>
      <span key={index}>
        {text}
        <br />
      </span>
    )
    return (
      <div>
        <Card>
          <p>{logs}</p>
        </Card>
      </div>
    )
  }
}

class ButtonLoad extends React.Component {

  constructor (props) {
    super(props);
  }

  _handleSubmit () {
    this.props.parentFunc()
  }

  render () {
    return (
      <div>
        <Tooltip title={this.props.tip}>
          <Button
            icon={this.props.icon}
            loading={this.props.loading}
            onClick={this._handleSubmit.bind(this)}
          >
              {this.props.label}
          </Button>
        </Tooltip>
      </div>
    );
  }
}

class Button extends React.Component {

  constructor (props) {
    super(props);
  }

  _handleSubmit () {
    this.props.parentFunc()
  }

  render () {
    return (
      <div>
        <Tooltip title={this.props.tip}>
          <Button
            icon={this.props.icon}
            loading={this.props.loading}
            onClick={this._handleSubmit.bind(this)}
          >
            {this.props.label}
          </Button>
        </Tooltip>
      </div>
    );
  }
}

export {AppHeading, Tagline, Heading, RadOptns, PlainTextOutput, TextOutput, TextInput,
        TextAreaInput, Date, Select, Logger, ButtonLoad, Button}

*/
