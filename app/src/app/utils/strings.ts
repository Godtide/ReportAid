class AppStrings {

  static readonly appTitle='ReportAid'
  static readonly copyright='University of Sussex © 2018. Created by Steven Huckle'

}

class PathStrings {

  static readonly home='home'
  static readonly homePath='/'
  static readonly about='about'
  static readonly aboutPath='/about'
  static readonly overview='overview'
  static readonly overviewPath='/overview'
  static readonly help='help'
  static readonly helpPath='/help'

}

class HomeStrings {

  static heading = 'Home'

  static info = 'Use this application to record humanitarian aid data and to get humanitarian aid information.<br /><br />Read the [About](/about) section to learn about the origins of **ReportAid**.<br /><br />The [Overview](/overview) section describes a scenario where an organisation uses **ReportAid** to store information about their funding.<br /><br />The [Help](/help) section gives brief instructions as to how to use **ReportAid** - in essence, to store a humanitarian aid record, click on the [Create Aid Record](/create-aid-record) link and fill in all fields. To retrieve information, click on the [Get Aid Record](/get-aid-record) link.'

  static icon='home'

}

class AboutStrings {

  static heading = 'About ReportAid'

  static info = '**ReportAid** is the result of an academic paper titled: _Humanitarian Aid - a Blockchain Application That Adds Trust to Aid Provisioning_. The article discusses how the trust mechanisms of blockchain technology might be used to promote transparanecy in humanitarian aid provisioning. The general idea is that blockchains can add trust to the reporting of humanitarian aid funding.<br /><br />For more information about **Provenator**, please contact s dot huckle at sussex dot ac dot uk.'

  static icon='info'

}

class OverviewStrings {

  static heading = 'Overview of ReportAid'

  static info = 'blah...'

  static icon='book'

}

class HelpStrings {

  static heading = 'ReportAid Help'

  static info = '**ReportAid** allows humanitarian aid organisations to record information about funding.<br /><br />Have a read of the [Overview](/overview), which describes a scenario where an organisation uses **ReportAid** to store information about their funding..<br /><br />to store a humanitarian aid record, click on the [Create Aid Record](/create-aid-record) link and fill in all fields. That will create some blockchain transactions, which can be signed in [MetaMask](https://metamask.io/). To retrieve information, click on the [Get Aid Record](/get-aid-record) link.'

  static icon='question'

}

export { AppStrings, PathStrings, HomeStrings, AboutStrings, OverviewStrings, HelpStrings }
