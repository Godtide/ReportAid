const StringsLib = artifacts.require("./Strings.sol");
const IATIOrgs = artifacts.require("./IATIOrgs.sol");
const IATIOrganisations = artifacts.require("./IATIOrganisations.sol");
const IATIOrganisation = artifacts.require("./IATIOrganisation.sol");
const IATIOrganisationDocs = artifacts.require("./IATIOrganisationDocs.sol");
const IATIOrganisationBudgets = artifacts.require("./IATIOrganisationBudgets.sol");
const IATIOrganisationExpenditure = artifacts.require("./IATIOrganisationExpenditure.sol");
const IATIOrganisationRecipientBudgets = artifacts.require("./IATIOrganisationRecipientBudgets.sol");
const IATIOrganisationRegionBudgets = artifacts.require("./IATIOrganisationRegionBudgets.sol");
const IATIOrganisationCountryBudgets = artifacts.require("./IATIOrganisationCountryBudgets.sol");

const IATIActivities = artifacts.require("./IATIActivities.sol");
const IATIActivity = artifacts.require("./IATIActivity.sol");

module.exports = function(deployer) {
  deployer.deploy(StringsLib);
  deployer.link(StringsLib, IATIOrgs);
  deployer.link(StringsLib, IATIOrganisations);
  deployer.link(StringsLib, IATIOrganisation);
  deployer.link(StringsLib, IATIOrganisationDocs);
  deployer.link(StringsLib, IATIOrganisationBudgets);
  deployer.link(StringsLib, IATIOrganisationExpenditure);
  deployer.link(StringsLib, IATIOrganisationRecipientBudgets);
  deployer.link(StringsLib, IATIOrganisationRegionBudgets);
  deployer.link(StringsLib, IATIOrganisationCountryBudgets);
  deployer.link(StringsLib, IATIActivities);
  deployer.link(StringsLib, IATIActivity);
  deployer.deploy(IATIOrgs);
  deployer.deploy(IATIOrganisations);
  deployer.deploy(IATIOrganisation);
  deployer.deploy(IATIOrganisationDocs);
  deployer.deploy(IATIOrganisationBudgets);
  deployer.deploy(IATIOrganisationExpenditure);
  deployer.deploy(IATIOrganisationRecipientBudgets);
  deployer.deploy(IATIOrganisationRegionBudgets);
  deployer.deploy(IATIOrganisationCountryBudgets);
  deployer.deploy(IATIActivities);
  deployer.deploy(IATIActivity);
};
