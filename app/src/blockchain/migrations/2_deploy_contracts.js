var StringsLib = artifacts.require("./Strings.sol");
var IATIOrganisations = artifacts.require("./IATIOrganisations.sol");

module.exports = function(deployer) {
  deployer.deploy(StringsLib);
  deployer.link(StringsLib, IATIOrganisations);
  deployer.deploy(IATIOrganisations);
};
