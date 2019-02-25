const StringsLib = artifacts.require("./Strings.sol");
const IATIBudgets = artifacts.require("./IATIBudgets.sol");

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
const IATIActivityDates = artifacts.require("./IATIActivityDates.sol");

module.exports = function(deployer) {

  let orgsAddress;
  let organisationsAddress;
  let organisationAddress;
  let organisationDocsAddress;
  let organisationBudgetsAddress;
  let organisationExpenditureAddress;
  let organisationRecipientBudgetsAddress;
  let organisationRegionBudgetsAddress;
  let organisationCountryBudgetsAddress;
  let activitiesAddress;
  let activityAddress;
  let activityDatesAddress;

  deployer.deploy(StringsLib);
  deployer.link(StringsLib, [IATIBudgets,
                             IATIOrgs,
                             IATIOrganisations,
                             IATIOrganisation,
                             IATIOrganisationDocs,
                             IATIActivities,
                             IATIActivity,
                             IATIActivityDates
                            ]);

  deployer.deploy(IATIOrgs).then(function() {
    orgsAddress = "\"" + IATIOrgs.address + "\"";

  });

  deployer.deploy(IATIOrganisations).then(function() {
    organisationsAddress = "\"" + IATIOrganisations.address + "\"";
  });

  deployer.deploy(IATIOrganisation).then(function() {
    organisationAddress = "\"" + IATIOrganisation.address + "\"";
  });

  deployer.deploy(IATIOrganisationDocs).then(function() {
    organisationDocsAddress = "\"" + IATIOrganisationDocs.address + "\"";
  });

  deployer.deploy(IATIBudgets).then(function() {
    //console.log(IATIBudgets.address)
    deployer.deploy(IATIOrganisationBudgets, IATIBudgets.address).then(function() {
      organisationBudgetsAddress = "\"" + IATIOrganisationBudgets.address + "\"";
    });
    deployer.deploy(IATIOrganisationExpenditure, IATIBudgets.address).then(function() {
      organisationExpenditureAddress = "\"" + IATIOrganisationExpenditure.address + "\"";
    });
    deployer.deploy(IATIOrganisationRecipientBudgets, IATIBudgets.address).then(function() {
      organisationRecipientBudgetsAddress = "\"" + IATIOrganisationRecipientBudgets.address + "\"";
    });
    deployer.deploy(IATIOrganisationRegionBudgets, IATIBudgets.address).then(function() {
      organisationRegionBudgetsAddress = "\"" + IATIOrganisationRegionBudgets.address + "\"";
    });
    deployer.deploy(IATIOrganisationCountryBudgets, IATIBudgets.address).then(function() {
      organisationCountryBudgetsAddress = "\"" + IATIOrganisationCountryBudgets.address + "\"";
    });
  });

  deployer.deploy(IATIActivities).then(function() {
    activitiesAddress = "\"" + IATIActivities.address + "\"";
  });

  deployer.deploy(IATIActivity).then(function() {
    activityAddress = "\"" + IATIActivity.address + "\"";
  });

  deployer.deploy(IATIActivityDates).then(function() {
    activityDatesAddress = "\"" + IATIActivityDates.address + "\"";
  });

  deployer.then( () => {
    console.log("static orgsAddress =", orgsAddress);
    console.log("static organisationsAddress =", organisationsAddress);
    console.log("static organisationAddress =", organisationAddress);
    console.log("static organisationDocsAddress =", organisationDocsAddress);
    console.log("static organisationBudgetsAddress =", organisationBudgetsAddress);
    console.log("static organisationExpenditureAddress =", organisationExpenditureAddress);
    console.log("static organisationRecipientBudgetsAddress =", organisationRecipientBudgetsAddress);
    console.log("static organisationRegionBudgetsAddress =", organisationRegionBudgetsAddress);
    console.log("static organisationCountryBudgetsAddress =", organisationCountryBudgetsAddress);
    console.log("static activitiesAddress =", activitiesAddress);
    console.log("static activityAddress =", activityAddress);
    console.log("static activityDatesAddress =", activityDatesAddress);
  });
};
