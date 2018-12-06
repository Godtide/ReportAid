var StringsLib = artifacts.require("./Strings.sol");
var IATIOrganisations = artifacts.require("./IATIOrganisations.sol");
var IATIOrganisationReports = artifacts.require("./IATIOrganisationReports.sol");

module.exports = function(deployer) {
  deployer.deploy(StringsLib);
  deployer.link(StringsLib, IATIOrganisations);
  deployer.link(StringsLib, IATIOrganisationReports);
  deployer.deploy(IATIOrganisations);
  deployer.deploy(IATIOrganisationReports);
};
