pragma solidity ^0.5.0;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationBudgets.sol";
import "./Strings.sol";

contract IATIOrganisationBudgets is OrganisationBudgets {

  struct Finance {
    uint256 value;
    bytes32 status;
    bytes32 start;
    bytes32 end;
  }

  bytes32[] orgReferences;
  mapping(bytes32 => Finance) private orgBudgets;
  mapping(bytes32 => Finance) private orgExpenditure;
  mapping(bytes32 => mapping(bytes32 => Finance)) private orgRecipientBudget;
  mapping(bytes32 => mapping(bytes32 => Finance)) private orgRegionBudget;
  mapping(bytes32 => mapping(bytes32 => Finance)) private orgCountryBudget;

  event SetTotalBudget(bytes32 _orgRef, uint256 _value, bytes32 _status, bytes32 _start, bytes32 _end);
  event SetRecipientOrgBudget(bytes32 _orgRef, bytes32 _recipientOrgRef, uint256 _value, bytes32 _status, bytes32 _start, bytes32 _end);
  event SetRegionBudget(bytes32 _orgRef, bytes32 _regionRef, uint256 _value, bytes32 _status, bytes32 _start, bytes32 _end);
  event SetCountryBudget(bytes32 _orgRef, bytes32 _countryRef, uint256 _value, bytes32 _status, bytes32 _start, bytes32 _end);
  event SetTotalExpenditure(bytes32 _orgRef, uint256 _value, bytes32 _status, bytes32 _start, bytes32 _end);

  function setTotalBudget(bytes32 _orgRef, uint256 _value, bytes32 _status, bytes32 _start, bytes32 _end) public {}
  function setRecipientOrgBudget(bytes32 _orgRef, bytes32 _recipientOrgRef, uint256 _value, bytes32 _status, bytes32 _start, bytes32 _end) public {}
  function setRegionBudget(bytes32 _orgRef, bytes32 _regionRef, uint256 _value, bytes32 _status, bytes32 _start, bytes32 _end) public {}
  function setCountryBudget(bytes32 _orgRef, bytes32 _countryRef, uint256 _value, bytes32 _status, bytes32 _start, bytes32 _end) public {}
  function setTotalExpenditure(bytes32 _orgRef, uint256 _value, bytes32 _status, bytes32 _start, bytes32 _end) public {}

  function getNumOrganisations() public view returns (uint256) {}
  function getOrganisationReference(uint256 _index) public view returns (bytes32) {}
  function getOrganisationExists(bytes32 _orgRef) public view returns (bool) {}

  function getNumBudgets(bytes32 _orgRef) public view returns (uint256) {}
  function getBudgetReference(bytes32 _orgRef, uint256 _index) public view returns (bytes32) {}
  function getTotalBudget(bytes32 _orgRef, bytes32 _budgetRef) public view returns (uint256) {}
  function getTotalBudgetStatus(bytes32 _orgRef, bytes32 _budgetRef) public view returns (uint256) {}
  function getTotalBudgetStart(bytes32 _orgRef, bytes32 _budgetRef) public view returns (uint256) {}
  function getTotalBudgetEnd(bytes32 _orgRef, bytes32 _budgetRef) public view returns (uint256) {}

  function getNumRecipientOrgs(bytes32 _orgRef) public view returns (uint256) {}
  function getRecipientOrgReference(bytes32 _orgRef, uint256 _index) public view returns (bytes32) {}
  function getRecipientOrgBudget(bytes32 _orgRef, bytes32 _recipientOrgRef) public view returns (uint256) {}
  function getRecipientOrgBudgetStatus(bytes32 _orgRef, bytes32 _recipientOrgRef) public view returns (uint256) {}
  function getRecipientOrgBudgetStart(bytes32 _orgRef, bytes32 _recipientOrgRef) public view returns (uint256) {}
  function getRecipientOrgBudgetEnd(bytes32 _orgRef, bytes32 _recipientOrgRef) public view returns (uint256) {}

  function getNumRegions(bytes32 _orgRef) public view returns (uint256) {}
  function getRegionsReference(bytes32 _orgRef, uint256 _index) public view returns (bytes32) {}
  function getRegionsBudget(bytes32 _orgRef, bytes32 _regionRef) public view returns (uint256) {}
  function getRegionsBudgetStatus(bytes32 _orgRef, bytes32 _regionRef) public view returns (uint256) {}
  function getRegionsBudgetStart(bytes32 _orgRef, bytes32 _regionRef) public view returns (uint256) {}
  function getRegionsBudgetEnd(bytes32 _orgRef, bytes32 _regionRef) public view returns (uint256) {}

  function getNumCountries(bytes32 _orgRef) public view returns (uint256) {}
  function getCountryReference(bytes32 _orgRef, uint256 _index) public view returns (bytes32) {}
  function getCountryBudget(bytes32 _orgRef, bytes32 _countryRef) public view returns (uint256) {}
  function getCountryBudgetStatus(bytes32 _orgRef, bytes32 _countryRef) public view returns (uint256) {}
  function getCountryBudgetStart(bytes32 _orgRef, bytes32 _countryRef) public view returns (uint256) {}
  function getCountryBudgetEnd(bytes32 _orgRef, bytes32 _countryRef) public view returns (uint256) {}

  function getNumExpenditures(bytes32 _orgRef) public view returns (uint256) {}
  function getExpenditureReference(bytes32 _orgRef, uint256 _index) public view returns (bytes32) {}
  function getTotalExpenditure(bytes32 _orgRef, bytes32 _expenditureRef) public view returns (uint256) {}
  function getTotalExpenditureStatus(bytes32 _orgRef, bytes32 _expenditureRef) public view returns (uint256) {}
  function getTotalExpenditureStart(bytes32 _orgRef, bytes32 _expenditureRef) public view returns (uint256) {}
  function getTotalExpenditureEnd(bytes32 _orgRef, bytes32 _expenditureRef) public view returns (uint256) {}
}
