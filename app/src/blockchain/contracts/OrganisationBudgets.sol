pragma solidity ^0.4.24;

contract OrganisationBudgets {

  function setTotalBudget(string _orgRef, uint256 _value, string _status, string _start, string _end) public;
  function setRecipientOrgBudget(string _orgRef, string _recipientOrgRef, uint256 _value, string _status, string _start, string _end) public;
  function setRegionBudget(string _orgRef, string _regionRef, uint256 _value, string _status, string _start, string _end) public;
  function setCountryBudget(string _orgRef, string _countryRef, uint256 _value, string _status, string _start, string _end) public;
  function setTotalExpenditure(string _orgRef, uint256 _value, string _status, string _start, string _end) public;

  function getNumOrganisations() public constant returns (uint256);
  function getOrganisationReference(uint256 _index) public constant returns (string);
  function getOrganisationExists(string _orgRef) public constant returns (bool);

  function getNumBudgets(string _orgRef) public constant returns (uint256);
  function getBudgetReference(string _orgRef, uint256 _index) public constant returns (string);
  function getTotalBudget(string _orgRef, string _budgetRef) public constant returns (uint256);
  function getTotalBudgetStatus(string _orgRef, string _budgetRef) public constant returns (uint256);
  function getTotalBudgetStart(string _orgRef, string _budgetRef) public constant returns (uint256);
  function getTotalBudgetEnd(string _orgRef, string _budgetRef) public constant returns (uint256);

  function getNumRecipientOrgs(string _orgRef) public constant returns (uint256);
  function getRecipientOrgReference(string _orgRef, uint256 _index) public constant returns (string);
  function getRecipientOrgBudget(string _orgRef, string _recipientOrgRef) public constant returns (uint256);
  function getRecipientOrgBudgetStatus(string _orgRef, string _recipientOrgRef) public constant returns (uint256);
  function getRecipientOrgBudgetStart(string _orgRef, string _recipientOrgRef) public constant returns (uint256);
  function getRecipientOrgBudgetEnd(string _orgRef, string _recipientOrgRef) public constant returns (uint256);

  function getNumRegions(string _orgRef) public constant returns (uint256);
  function getRegionsReference(string _orgRef, uint256 _index) public constant returns (string);
  function getRegionsBudget(string _orgRef, string _regionRef) public constant returns (uint256);
  function getRegionsBudgetStatus(string _orgRef, string _regionRef) public constant returns (uint256);
  function getRegionsBudgetStart(string _orgRef, string _regionRef) public constant returns (uint256);
  function getRegionsBudgetEnd(string _orgRef, string _regionRef) public constant returns (uint256);

  function getNumCountries(string _orgRef) public constant returns (uint256);
  function getCountryReference(string _orgRef, uint256 _index) public constant returns (string);
  function getCountryBudget(string _orgRef, string _countryRef) public constant returns (uint256);
  function getCountryBudgetStatus(string _orgRef, string _countryRef) public constant returns (uint256);
  function getCountryBudgetStart(string _orgRef, string _countryRef) public constant returns (uint256);
  function getCountryBudgetEnd(string _orgRef, string _countryRef) public constant returns (uint256);

  function getNumExpenditures(string _orgRef) public constant returns (uint256);
  function getExpenditureReference(string _orgRef, uint256 _index) public constant returns (string);
  function getTotalExpenditure(string _orgRef, string expenditureRef) public constant returns (uint256);
  function getTotalExpenditureStatus(string _orgRef, string expenditureRef) public constant returns (uint256);
  function getTotalExpenditureStart(string _orgRef, string expenditureRef) public constant returns (uint256);
  function getTotalExpenditureEnd(string _orgRef, string expenditureRef) public constant returns (uint256);
}
