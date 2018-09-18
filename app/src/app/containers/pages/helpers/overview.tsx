import * as React from 'react'
import { connect } from 'react-redux'

import { ApplicationState } from '../../../store'
import { IATIReader } from '../../../store/IATIReader/types'
import { fetchRequest } from '../../..store/IATIReader/actions'

// Separate state props + dispatch props to their own interfaces.
interface OverviewPropsFromState {
  loading: boolean
  data: IATIReader[]
  errors: string
}

// Combine both state + dispatch props - as well as any props we want to pass - in a union type.
type AllProps = OverviewPropsFromState

class Overview extends React.Component<AllProps> {

  public render() {

    const { data } = this.props
    return (
        <p>
          {data}
        </p>
    )
  }
}

const mapStateToProps = (ownProps: HomePropsFromState): HomePropsFromState => {
  return {
    data: ownProps.data,
    errors: ownProps.errors
  }
}

// Now let's connect our component!
// With redux v4's improved typings, we can finally omit generics here.
export default connect(
  mapStateToProps
)(Overview)
