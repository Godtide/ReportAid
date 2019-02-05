const StringsLib = artifacts.require("./Strings.sol");
const IATIOrgs = artifacts.require("./IATIOrgs.sol");
const IATIOrganisations = artifacts.require("./IATIOrganisations.sol");
const IATIOrganisationReportDocs = artifacts.require("./IATIOrganisationReportDocs.sol");
const IATIOrganisationReportBudgets = artifacts.require("./IATIOrganisationReportBudgets.sol");
const IATIOrganisationReportExpenditure = artifacts.require("./IATIOrganisationReportExpenditure.sol");
const IATIOrganisationReportRecipientBudgets = artifacts.require("./IATIOrganisationReportRecipientBudgets.sol");
const IATIOrganisationReportRegionBudgets = artifacts.require("./IATIOrganisationReportRegionBudgets.sol");
const IATIOrganisationReportCountryBudgets = artifacts.require("./IATIOrganisationReportCountryBudgets.sol");

const IATIActivities = artifacts.require("./IATIActivities.sol");
const IATIActivity = artifacts.require("./IATIActivity.sol");

module.exports = function(deployer) {
  deployer.deploy(StringsLib);
  deployer.link(StringsLib, IATIOrgs);
  deployer.link(StringsLib, IATIOrganisations);
  deployer.link(StringsLib, IATIOrganisationReportDocs);
  deployer.link(StringsLib, IATIOrganisationReportBudgets);
  deployer.link(StringsLib, IATIOrganisationReportExpenditure);
  deployer.link(StringsLib, IATIOrganisationReportRecipientBudgets);
  deployer.link(StringsLib, IATIOrganisationReportRegionBudgets);
  deployer.link(StringsLib, IATIOrganisationReportCountryBudgets);
  deployer.link(StringsLib, IATIActivities);
  deployer.link(StringsLib, IATIActivity);
  deployer.deploy(IATIOrgs);
  deployer.deploy(IATIOrganisations);
  deployer.deploy(IATIOrganisationReportDocs);
  deployer.deploy(IATIOrganisationReportBudgets);
  deployer.deploy(IATIOrganisationReportExpenditure);
  deployer.deploy(IATIOrganisationReportRecipientBudgets);
  deployer.deploy(IATIOrganisationReportRegionBudgets);
  deployer.deploy(IATIOrganisationReportCountryBudgets);
  deployer.deploy(IATIActivities);
  deployer.deploy(IATIActivity);
};
