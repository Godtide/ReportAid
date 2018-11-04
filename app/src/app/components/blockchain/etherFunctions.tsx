import * as React from 'react'
import Web3 from 'web3'
import { ethers } from 'ethers'

interface CallCheckerProps {
  func(): void
  cb(): void
}

const CallChecker = (props: CallCheckerProps) => {

  return (
    // console.log('in func call checker')
    if ((typeof CallCheckerProps.func === 'function') && (typeof CallCheckerProps.cb === 'function')) {
      // console.log('Passed call checker!')
      return true
    } else {
      /* console.log('failed call checker!')
      console.log('func ' + _func)
      console.log('cb ' + _cb) */
      return false
    }
  )
}

interface CallParamsCheckerProps {
  func(): void
  cb(): void
  params: []
}

const CallChecker = (props: CallCheckerProps) => {

  return (
    if ((typeof CallCheckerProps.func === 'function') && (typeof CallCheckerProps.cb === 'function') && (Array.isArray(CallCheckerProps.params))) {
      return true
    } else {
      return false
    }
  )
}

/*

// metamask sets its account to web3.eth.accounts[0]
_setAccount () {
  //console.log("In set account for ", this.account)
}

  _call (_func: Web3HandlerProps, _cb: Web3HandlerProps, _params: []) {
    if (this._callChecker(_func, _cb)) {
      _func(function (err: any, result: any) {
        if (err) {
          console.log(err)
        } else {
          //console.log(result)
          _cb(_caller, result)
        }
      })
    }
  }

  _call1Params (_caller, _func, _params, _cb) {
    if (this._callParamsChecker(_func, _params, _cb)) {
      _func(_params[0], function (err, result) {
        if (err) {
          console.log(err)
        } else {
          //console.log(result)
          _cb(_caller, result)
        }
      })
    }
  }

  _call2Params (_caller, _func, _params, _cb) {
    if (this._callParamsChecker(_func, _params, _cb)) {
      _func(_params[0], _params[1], function (err, result) {
        if (err) {
          console.log(err)
        } else {
          //console.log(result)
          _cb(_caller, result)
        }
      })
    }
  }

  _call3Params (_caller, _func, _params, _cb) {
    if (this._callParamsChecker(_func, _params, _cb)) {
      _func(_params[0], _params[1], _params[2], function (err, result) {
        if (err) {
          console.log(err)
        } else {
          //console.log(result)
          _cb(_caller, result)
        }
      })
    }
  }

  _call4Params (_caller, _func, _params, _cb) {
    if (this._callParamsChecker(_func, _params, _cb)) {
      _func(_params[0], _params[1], _params[2], _params[3], function (err, result) {
        if (err) {
          console.log(err)
        } else {
          //console.log(result)
          _cb(_caller, result)
        }
      })
    }
  }

  _call5Params (_caller, _func, _params, _cb, _isBatch) {
    if (this._callParamsChecker(_func, _params, _cb)) {
      _func(_params[0], _params[1], _params[2], _params[3], _params[4], function (err, result) {
        if (err) {
          console.log(err)
        } else {
          //console.log(result)
          _cb(_caller, result)
        }
      })
    }
  }

  _call6Params (_caller, _func, _params, _cb, _isBatch) {
    if (this._callParamsChecker(_func, _params, _cb)) {
      _func(_params[0], _params[1], _params[2], _params[3], _params[4], _params[5], function (err, result) {
        if (err) {
          console.log(err)
        } else {
          //console.log(result)
          _cb(_caller, result)
        }
      })
    }
  }

  getAccount () {
    return this.account
  }

  getWeb3 () {
    return this.web3
  }

  callHandler (_caller, _func, _cb) {
    this._call(_caller, _func, _cb)
  }

  callParamHandler (_caller, _func, _params, _cb) {
    switch (_params.length) {
      case 1:
        this._call1Params(_caller, _func, _params, _cb)
        break
      case 2:
        this._call2Params(_caller, _func, _params, _cb)
        break
      case 3:
        this._call3Params(_caller, _func, _params, _cb)
        break
      case 4:
        this._call4Params(_caller, _func, _params, _cb)
        break
      case 5:
        this._call5Params(_caller, _func, _params, _cb)
        break
      case 6:
        this._call6Params(_caller, _func, _params, _cb)
        break
      default:
        this._call(_caller, _func, _cb)
    }
  }
}

*/

export {CallChecker, CallParamsCheckerProps}
