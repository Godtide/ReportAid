import Web3 from 'web3'

class Web3Handler {

  constructor () {
    // Checking if Web3 has been injected by the browser (Mist/MetaMask)
    if (typeof web3 !== 'undefined') {
      console.log('MetaMask!')
      this.web3 = new Web3(web3.currentProvider)
    } else {
      console.log('No web3? You should consider trying MetaMask!')
      let host = '127.0.0.1'
      let port = '8545'
      this.web3 = new Web3(new Web3.providers.HttpProvider('http://' + host + ':' + port))
    }
    this.web3.version.getNetwork((err, netId) => {
      switch (netId) {
        case "1":
          console.log('This is mainnet')
          break
        case "2":
          console.log('This is the deprecated Morden test network.')
          break
        case "3":
          console.log('This is the ropsten test network.')
          break
        case "4":
          console.log('This is the rinkeby test network.')
          break
        default:
          console.log('This is a local, unknown network.')
      }
    })
    this.web3.eth.defaultAccount = this.web3.eth.accounts[0]
    this.account = this.web3.eth.defaultAccount
    console.log('This account ' + this.account)
    setInterval(this._setAccount.bind(this), 3000)
  }

  // metamask sets its account to web3.eth.accounts[0]
  _setAccount () {
    if (this.web3.eth.accounts[0] !== this.account) {
      this.account = this.web3.eth.accounts[0]
    }
    //console.log("In set account for ", this.account)
  }

  _callChecker (_func, _cb) {
    // console.log('in func call checker')
    if ((typeof _func === 'function') && (typeof _cb === 'function')) {
      // console.log('Passed call checker!')
      return true
    } else {
      /* console.log('failed call checker!')
      console.log('func ' + _func)
      console.log('cb ' + _cb) */
      return false
    }
  }

  _callParamsChecker (_func, _cb, _params) {
    if ((typeof _func === 'function') && (typeof _cb === 'function') && (Array.isArray(_params))) {
      return true
    } else {
      return false
    }
  }

  _call (_caller, _func, _cb) {
    if (this._callChecker(_func, _cb)) {
      _func(function (err, result) {
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

export default Web3Handler
