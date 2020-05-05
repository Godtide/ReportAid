# ReportAid

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](/docs/prs.md) [![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](/docs/COPYING.txt)

This is the repository of **ReportAid**, a blockchain implementation of the [IATI Standard](https://iatistandard.org/en/), an international standard for reporting on humanitarian aid.

## Table of Contents

- [Usage](#usage)
- [Demo](#demo)
  - [Demo Dependencies](#browser-application-demo-dependencies)  
  - [dat Dependencies](#dat-dependencies)
- [Built Using](#built-using)  
- [Install](#install)
  - [Getting Started](#getting-started)
  - [Dependencies](#dependencies)    
  - [Detailed Instructions](#detailed-instructions)
- [Maintainer](#maintainer)
- [Contributing](#contributing)
- [License](#license)

## Usage

Below describes using and installing the **ReportAid** demo'.

## Demo

**ReportAid** features both a [React](https://reactjs.org/) browser-based application for creating and reading [IATI Standard](https://iatistandard.org/en/) blockchain records, and a [Go](https://golang.org/) RESTful API server for reading those records.

A demo of the **ReportAid** browser application is running both on the web and on the distributed filesystem [dat](https://dat.foundation/).

web: [http://www.reportaid.org](http://www.reportaid.org)
dat: [http://4b1bdf7b0f6beeadab5dadaf019cddbc94f618792ea30b8a2f5d957267d5bd92/](http://4b1bdf7b0f6beeadab5dadaf019cddbc94f618792ea30b8a2f5d957267d5bd92/)

![](/docs/paper/images/reportaidWriteSector.png)

The [Go](https://golang.org/) RESTful API allows you to retrieve [IATI Standard XML](https://iatistandard.org/en/) records from  **ReportAid**.

RESTful API: [http://213.138.113.94:10000/](http://213.138.113.94:10000/)

e.g: [http://213.138.113.94:10000/activities-list](http://213.138.113.94:10000/activities-list)

To get some help using the RESTful API, browse [http://213.138.113.94:10000/help](http://213.138.113.94:10000/help). Please note - the RESTful API is still in the early stages of development - more help is coming soon! indeed, **ReportAid** is an early proof of concept, so you should expect some bugs (please raise an [issue](https://github.com/glowkeeper/ReportAid/issues) if you find any).

### Browser Application Demo Dependencies

**ReportAid** requires the [Firefox](https://www.mozilla.org/) browser with the [MetaMask](https://metamask.io/) extension. The smart contracts of **ReportAid** are running on the [Rinkeby Ethereum Testnet](https://www.rinkeby.io/); hence, for **ReportAid** to interact with those, [MetaMask](https://metamask.io/) should be pointing at [Rinkeby](https://www.rinkeby.io/). You will need a few test Ether in your [MetaMask](https://metamask.io/) wallet, which you can get from the [Rinkeby Faucet](https://faucet.rinkeby.io/). Those test Ether will allow you to sign the transactions necessary to create [IATI Standard](https://iatistandard.org/en/) records that **ReportAid**  supports.

### dat Dependencies

To run **ReportAid** via [dat](https://dat.foundation/), you will need the [Firefox](https://www.mozilla.org/) [dat P2P Protocol](https://addons.mozilla.org/en-GB/firefox/addon/dat-p2p-protocol/) extension installed. Ensure [MetaMask](https://metamask.io/) is pointing at [Rinkeby](https://www.rinkeby.io/), then load the following URL: http://4b1bdf7b0f6beeadab5dadaf019cddbc94f618792ea30b8a2f5d957267d5bd92/

## Built Using

- [node](https://nodejs.org/en/)
- [npm](https://www.npmjs.com/)
- [Ganache](https://github.com/trufflesuite/ganache)
- [Truffle](https://github.com/trufflesuite/truffle)
- [React](https://reactjs.org/) + [Redux](https://redux.js.org/) + [TypeScript](https://www.typescriptlang.org/)
- [Go](https://golang.org/)

## Install

The instruction below enable you to run **ReportAid** app' and RESTful API server on a local, private, Ethereum test network (via [Ganache](https://github.com/trufflesuite/ganache)). Before following the instructions below, please install the [dependencies](#dependencies).

Follow the instructions in the [Ganache](https://github.com/trufflesuite/ganache) repository for downloading and installing Ganache; tl;dr - you need to clone the [Ganache](https://github.com/trufflesuite/ganache) repository, then run `npm install && npm start`.

In the **ReportAid** repository's [/app](/app) directory, type `npm install`. That should install everything listed in [package.json](/app/package.json), which form the components of the REACT-based web frontend to this application.

Now, publish the contracts to your local blockchain (via [Ganache](https://github.com/trufflesuite/ganache)):

1. Change directory to the [Ganache](https://github.com/trufflesuite/ganache) repository.
2. Start [Ganache](https://github.com/trufflesuite/ganache) by typing `npm start`.
3. Ensure [Ganache](https://github.com/trufflesuite/ganache) is running on [http://localhost:8545](http://localhost:8545) (you may need to change its settings to do so).
4. Change to the **ReportAid** repository's [/app](/app) directory.
5. Build the **ReportAid** smart contracts by running `npm run generate`.
6. Deploy the **ReportAid** smart contracts by running `npm run develMigrate`.
7. Edit the **ReportAid** config file [/app/src/app/utils/config.ts](/app/src/app/utils/config.ts) so that the contract static variables contain the addresses generated by `npm run develMigrate`, above.

Now create the web application:

1. Build the REACT frontend by typing `npm run compile` (you can also watch for any changes by running `npm run watch`).
2. Startup an instance of [http-server](https://www.npmjs.com/package/http-server) by typing `npm run start`.

Then fire up a [Firefox](https://www.mozilla.org/) browser and go to the URL [http://localhost:8081](http://localhost:8081). You must have the [Firefox](https://www.mozilla.org/) extension [MetaMask](https://metamask.io/) installed. [Ganache](https://github.com/trufflesuite/ganache) will have generated some test addresses with 100 test Ether that will allow you to sign the transactions necessary to create [IATI Standard](https://iatistandard.org/en/) records that update the blockchain. You need to import one of those addresses' private keys into [MetaMask](https://metamask.io/).

To run the RESTFul API server:

1. Change to the **ReportAid** repository's [/app/src/server](/app/src/server) directory.
2. Edit the server's ethereum config file [./internal/configs/ethereum.go](./internal/configs/ethereum.go) so that the contract constants contain the addresses generated above.
3. Build the server by typing `go build -o ./tmp/goReportAid ./cmd/server/main.go`.
4. Run the server by running `/tmp/goReportAid`.
5. As an alternative to steps 3 and 4, you can use [air](https://github.com/cosmtrek/air), which will build and run the server for you ([air](https://github.com/cosmtrek/air) also features live reload, should you wish to do any development of the server).

The server runs on port 10000, so load a browser, then connect to [http://localhost:10000](http://localhost:10000). To get some help using the server, load [http://localhost:10000/help](http://localhost:10000/help).

### Dependencies

After cloning this repository, download and ensure all dependencies are installed:

- [node](https://nodejs.org/en/)
- [npm](https://www.npmjs.com/)
- [Ganache](https://github.com/trufflesuite/ganache)
- [Truffle](https://github.com/trufflesuite/truffle)
- [http-server](https://www.npmjs.com/package/http-server)
- [Go](https://golang.org/)

## Maintainer

[Steve Huckle](https://glowkeeper.github.io/).

## Contributing

Contributions welcome - if you want to get involved with the app's development, please get in touch via steve at glowkeeper dot uk.

## License

GNU General Public License v3.0

Please refer to the file: [COPYING](/docs/COPYING.txt) for the full text.
