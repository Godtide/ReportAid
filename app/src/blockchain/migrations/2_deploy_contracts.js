var StringsLib = artifacts.require("./Strings.sol");
var IATIOrganisations = artifacts.require("./IATIOrganisations.sol");
var IATIOrganisationReports = artifacts.require("./IATIOrganisationReports.sol");
var IATIOrganisationReportDocs = artifacts.require("./IATIOrganisationReportDocs.sol");
var IATIOrganisationReportBudgets = artifacts.require("./IATIOrganisationReportBudgets.sol");
var IATIOrganisationReportExpenditure = artifacts.require("./IATIOrganisationReportExpenditure.sol");
var IATIOrganisationReportRecipientBudgets = artifacts.require("./IATIOrganisationReportRecipientBudgets.sol");
var IATIOrganisationReportRegionBudgets = artifacts.require("./IATIOrganisationReportRegionBudgets.sol");
var IATIOrganisationReportCountryBudgets = artifacts.require("./IATIOrganisationReportCountryBudgets.sol");

module.exports = function(deployer) {
  deployer.deploy(StringsLib);
  deployer.link(StringsLib, IATIOrganisations);
  deployer.link(StringsLib, IATIOrganisationReports);
  deployer.link(StringsLib, IATIOrganisationReportDocs);
  deployer.link(StringsLib, IATIOrganisationReportBudgets);
  deployer.link(StringsLib, IATIOrganisationReportExpenditure);
  deployer.link(StringsLib, IATIOrganisationReportRecipientBudgets);
  deployer.link(StringsLib, IATIOrganisationReportRegionBudgets);
  deployer.link(StringsLib, IATIOrganisationReportCountryBudgets);
  deployer.deploy(IATIOrganisations);
  deployer.deploy(IATIOrganisationReports);
  deployer.deploy(IATIOrganisationReportDocs);
  deployer.deploy(IATIOrganisationReportBudgets);
  deployer.deploy(IATIOrganisationReportExpenditure);
  deployer.deploy(IATIOrganisationReportRecipientBudgets);
  deployer.deploy(IATIOrganisationReportRegionBudgets);
  deployer.deploy(IATIOrganisationReportCountryBudgets);
};
