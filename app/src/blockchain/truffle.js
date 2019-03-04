module.exports = {
  networks: {
    development: {
      host: "localhost",
      port: 8545,
      network_id: "*",
      gas: "6721975",
    },
    rinkeby: {
      host: "localhost", // Connect to geth on the specified
      port: 8545,
      from: "0x9ef3569b5dc377eb66cd0895af35b36efec94a88", // default address to use for any transaction Truffle makes during migrations
      network_id: 4,
      gas: 4612388 // Gas limit used for deploys
    },
    ropsten: {
      host: "localhost", // Connect to geth on the specified
      port: 8545,
      from: "0x79b0e7de13a17a74ab23fd2e6c69aa3cf93f4e1c",
      network_id: 3,
      gas: 4612388      //make sure this gas allocation isn't over 4M, which is the max
    }
  }
};
