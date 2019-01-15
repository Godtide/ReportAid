var StringsLib = artifacts.require("./Strings.sol");
var IATIOrganisations = artifacts.require("./IATIOrganisations.sol");
var IATIOrganisationReports = artifacts.require("./IATIOrganisationReports.sol");
var IATIOrganisationReportDocs = artifacts.require("./IATIOrganisationReportDocs.sol");

module.exports = function(deployer) {
  deployer.deploy(StringsLib);
  deployer.link(StringsLib, IATIOrganisations);
  deployer.link(StringsLib, IATIOrganisationReports);
  deployer.link(StringsLib, IATIOrganisationReportDocs);
  deployer.deploy(IATIOrganisations);
  deployer.deploy(IATIOrganisationReports);
  deployer.deploy(IATIOrganisationReportDocs);
};
