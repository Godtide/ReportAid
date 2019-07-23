# ReportAid

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](/docs/prs.md) [![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](/docs/COPYING.txt)

This is the repository of **ReportAid** - a tool for reporting on humanitarian aid. **ReportAid** is a blockchain implementation of the [IATI Standard](https://iatistandard.org/en/).

A live demo' of [ReportAid](http://4b1bdf7b0f6beeadab5dadaf019cddbc94f618792ea30b8a2f5d957267d5bd92/) is running on the [Rinkeby Ethereum Testnet](https://www.rinkeby.io/) at http://4b1bdf7b0f6beeadab5dadaf019cddbc94f618792ea30b8a2f5d957267d5bd92/. Instructions for running the demo' are available [below](#demoInstructions) - without fulfilling the dependencies listed in those instructions, [ReportAid](http://4b1bdf7b0f6beeadab5dadaf019cddbc94f618792ea30b8a2f5d957267d5bd92/) will not load.

Apologies in advance - help for using [ReportAid](http://4b1bdf7b0f6beeadab5dadaf019cddbc94f618792ea30b8a2f5d957267d5bd92/) is currently 'minimal', at best. Furthermore, the app' is very much a proof of concept, so expect some bugs. It is also missing some functionality as it does not, currently, implement all of the [IATI Standard](https://iatistandard.org/en/). That said, as a proof of concept, the app' serves its purpose nicely. 

The idea for **ReportAid** originally came via a [University of Sussex](https://www.sussex.ac.uk/) Masters Student studying within the university's Science Policy Research Unit (SPRU), who wanted to explore humanitarian blockchains as a means of entering SPRU's [2018 Science, Technology and Innovation Policy Challenge](http://www.sussex.ac.uk/spru/newsandevents/2018/awards/sti-challenge). The original intention of our entry was to use blockchains as a means of delivering trustworthy aid reports, and the [first draft of our policy proposal](docs/SPRUChallenge/SPRUSTIPolicyProposal.md) reflected that intention. However, after reading that first draft, the Masters Student decided to change tack, as she wanted to present blockchains as a novel means of actually providing finance during a humanitarian crisis; in other words, she wanted to leverage the cryptocurrency capabilities of blockchains directly. However, while that is a laudable aim, it was not something I, the original developer of *ReportAid*, thought viable, at the time. Instead, I could see the immediate benefit of a blockchain-based tool that could add trust to aid reports. Therefore, to my mind, a much more practical and immediately realisable aim would be to produce a blockchain implementation of the IATI standard for enhancing humanitarian aid reporting. As a result, **ReportAid**, a blockchain-based proof of concept for reporting on humanitarian aid, came into being.

The intention is to explain the ideas behind **ReportAid** in an academic paper with a proposed title _Humanitarian Aid - a Blockchain Application That Adds Trust to Aid Provisioning_. The current version of the [abstract](docs/abstract.md) gives more detail. Please get in touch with s dot huckle at sussex dot ac dot uk if you would like to contribute to that.

## Running the Demo'

<a name='demoInstructions'>To</a> see the demo' of [ReportAid](http://4b1bdf7b0f6beeadab5dadaf019cddbc94f618792ea30b8a2f5d957267d5bd92/) (and use it), you will need to be running [Firefox](https://www.mozilla.org/) with the [DAT P2P Protocol](https://addons.mozilla.org/en-GB/firefox/addon/dat-p2p-protocol/) and [MetaMask](https://metamask.io/) extensions installed. You will also need a few test Ether in your [MetaMask](https://metamask.io/) wallet, which you can get from the [Rinkeby Faucet](https://faucet.rinkeby.io/). Those test Ether will allow you to sign the transactions necessary to create [IATI Standard](https://iatistandard.org/en/) records that update the blockchain.

## Licensing

GNU General Public License v3.0

Please refer to the file: [COPYING](/docs/COPYING.txt) for the full text.
