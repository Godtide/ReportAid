pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationReportBudgets.sol";
import "./Strings.sol";

contract IATIOrganisationReportBudgets is OrganisationReportBudgets {

  bytes32[] reportReferences;

  mapping(bytes32 => bytes32[]) private budgetReferences;
  mapping(bytes32 => mapping(bytes32 => Budget)) private reportBudgets;

  mapping(bytes32 => bytes32[]) private expenditureReferences;
  mapping(bytes32 => mapping(bytes32 => Expenditure)) private reportExpenditure;

  mapping(bytes32 => bytes32[]) private reportOrgRecipientReferences;
  mapping(bytes32 => mapping(bytes32 => RecipientBudget)) private orgRecipientBudget;

  mapping(bytes32 => bytes32[]) private reportOrgRegionReferences;
  mapping(bytes32 => mapping(bytes32 => RegionBudget)) private orgRegionBudget;

  mapping(bytes32 => bytes32[]) private reportOrgCountryReferences;
  mapping(bytes32 => mapping(bytes32 => CountryBudget)) private orgCountryBudget;

  event SetTotalBudget(bytes32 _reportRef, Budget _budget);
  event SetTotalExpenditure(bytes32 _reportRef, Expenditure _expenditure);
  event SetRecipientOrgBudget(bytes32 _reportRef, RecipientBudget _budget);
  event SetRegionBudget(bytes32 _reportRef, RegionBudget _budget);
  event SetCountryBudget(bytes32 _reportRef, CountryBudget _budget);

  function setTotalBudget(Budget memory _budget) public {}
  function setTotalExpenditure(Expenditure memory _expenditure) public {}
  function setRecipientOrgBudget(RecipientBudget memory _budget) public {}
  function setRegionBudget(RegionBudget memory _budget) public {}
  function setCountryBudget(CountryBudget memory _budget) public {}

  function getNumOrganisations() public view returns (uint256) {}
  function getOrganisationReference(uint256 _index) public view returns (bytes32) {}
  function getOrganisationExists(bytes32 _reportRef) public view returns (bool) {}

  function getNumBudgets(bytes32 _reportRef) public view returns (uint256) {}
  function getBudgetReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {}
  function getTotalBudget(bytes32 _reportRef, bytes32 _budgetRef) public view returns (uint256) {}
  function getTotalBudgetStatus(bytes32 _reportRef, bytes32 _budgetRef) public view returns (uint256) {}
  function getTotalBudgetStart(bytes32 _reportRef, bytes32 _budgetRef) public view returns (uint256) {}
  function getTotalBudgetEnd(bytes32 _reportRef, bytes32 _budgetRef) public view returns (uint256) {}

  function getNumRecipientOrgs(bytes32 _reportRef) public view returns (uint256) {}
  function getRecipientOrgReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {}
  function getRecipientOrgBudget(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (uint256) {}
  function getRecipientOrgBudgetStatus(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (uint256) {}
  function getRecipientOrgBudgetStart(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (uint256) {}
  function getRecipientOrgBudgetEnd(bytes32 _reportRef, bytes32 _recipientOrgRef) public view returns (uint256) {}

  function getNumRegions(bytes32 _reportRef) public view returns (uint256) {}
  function getRegionsReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {}
  function getRegionsBudget(bytes32 _reportRef, bytes32 _regionRef) public view returns (uint256) {}
  function getRegionsBudgetStatus(bytes32 _reportRef, bytes32 _regionRef) public view returns (uint256) {}
  function getRegionsBudgetStart(bytes32 _reportRef, bytes32 _regionRef) public view returns (uint256) {}
  function getRegionsBudgetEnd(bytes32 _reportRef, bytes32 _regionRef) public view returns (uint256) {}

  function getNumCountries(bytes32 _reportRef) public view returns (uint256) {}
  function getCountryReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {}
  function getCountryBudget(bytes32 _reportRef, bytes32 _countryRef) public view returns (uint256) {}
  function getCountryBudgetStatus(bytes32 _reportRef, bytes32 _countryRef) public view returns (uint256) {}
  function getCountryBudgetStart(bytes32 _reportRef, bytes32 _countryRef) public view returns (uint256) {}
  function getCountryBudgetEnd(bytes32 _reportRef, bytes32 _countryRef) public view returns (uint256) {}

  function getNumExpenditures(bytes32 _reportRef) public view returns (uint256) {}
  function getExpenditureReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32) {}
  function getTotalExpenditure(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (uint256) {}
  function getTotalExpenditureStatus(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (uint256) {}
  function getTotalExpenditureStart(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (uint256) {}
  function getTotalExpenditureEnd(bytes32 _reportRef, bytes32 _expenditureRef) public view returns (uint256) {}
}
