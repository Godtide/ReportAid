class AppStrings {

  static readonly appTitle='ReportAid'
  static readonly appTagline='Using Blockchain Technology to Add Trust to Humanitarian Aid Reporing'
  static readonly copyright='University of Sussex © 2018'
  static readonly author='Created by Steven Huckle'

}

class PathStrings {

  static readonly home='Home'
  static readonly homePath='/'
  static readonly blockchain = 'Blockchain Data'
  static readonly blockchainPath = '/blockchain-data'
  static readonly about='About'
  static readonly aboutPath='/about'
  static readonly overview='Overview'
  static readonly overviewPath='/overview'
  static readonly help='Help'
  static readonly helpPath='/help'
  static readonly writer='Create Record'
  static readonly writerPath='/create-record'
  static readonly reader='Fetch Record'
  static readonly readerPath='/fetch-record'
}

class BlockchainStrings {

  // Using web3 via MetaMask, so these aren't necessary, really
  static readonly host = 'localhost'
  static readonly port = '8545'

  static heading = 'Blockchain Data'
  static APIProvider = 'APIProvider: '
  static networkName = 'Provider Name: '
  static networkChainId = 'Network Chain ID: '
  static ENSAddress = 'Network ENS Address: '
  static networkAccount = 'Account: '
  static setButton = 'Set Blockchain'
  static setButtonTip = 'Sets the blockchain'
}

class HomeStrings {

  static heading = 'Home'

  static info = 'Use this application to record humanitarian aid data and to get humanitarian aid information.<br /><br />Read the [About](#/about) section to learn about the origins of **ReportAid**.<br /><br />The [Overview](/overview) section describes a scenario where an organisation uses **ReportAid** to store information about their funding.<br /><br />The [Help](#/help) section gives brief instructions as to how to use **ReportAid** - in essence, to create a humanitarian aid record, click on the [Create Record](#/create-record) link and fill in all fields. To retrieve a humanitarian aid record, click on the [Fetch Record](#/fetch-record) link.'
}

class AboutStrings {

  static heading = 'About ReportAid'

  static info = '**ReportAid** is the result of an academic paper titled: _Humanitarian Aid - a Blockchain Application That Adds Trust to Aid Provisioning_. The article discusses how the trust mechanisms of blockchain technology might be used to promote transparanecy in humanitarian aid provisioning. The general idea is that blockchains can add trust to the reporting of humanitarian aid funding.<br /><br />For more information about **ReportAid**, please contact s dot huckle at sussex dot ac dot uk.'

}

class OverviewStrings {

  static heading = 'Overview of ReportAid'

  static info = 'blah...'
}

class HelpStrings {

  static heading = 'ReportAid Help'

  static info = '**ReportAid** allows humanitarian aid organisations to record information about funding.<br /><br />Have a read of the [About](#/about) section, which gives some background to the app\'. The [Overview](#/overview) section describes a scenario where an organisation uses **ReportAid** to store information about their funding.<br /><br />To store a humanitarian aid record, click on the [Create Record](#/create-record) link and fill in all fields. That will create some blockchain transactions, which can be signed in [MetaMask](https://metamask.io/). To retrieve information, click on the [Fetch Record](#/fetch-record) link.'
}

class WriterStrings {

  static heading = 'ReportAid Writer'
}

class ReaderStrings {

  static heading = 'ReportAid Reader'
}

export { AppStrings, PathStrings, BlockchainStrings, HomeStrings, AboutStrings, OverviewStrings, HelpStrings, WriterStrings, ReaderStrings }
