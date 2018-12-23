pragma solidity ^0.5.0;

contract OrganisationBudgets {

  function setTotalBudget(string memory _orgRef, uint256 _value, string memory _status, string memory _start, string memory _end) public;
  function setRecipientOrgBudget(string memory _orgRef, string memory _recipientOrgRef, uint256 _value, string memory _status, string memory _start, string memory _end) public;
  function setRegionBudget(string memory _orgRef, string memory _regionRef, uint256 _value, string memory _status, string memory _start, string memory _end) public;
  function setCountryBudget(string memory _orgRef, string memory _countryRef, uint256 _value, string memory _status, string memory _start, string memory _end) public;
  function setTotalExpenditure(string memory _orgRef, uint256 _value, string memory _status, string memory _start, string memory _end) public;

  function getNumOrganisations() public view returns (uint256);
  function getOrganisationReference(uint256 _index) public view returns (string memory);
  function getOrganisationExists(string memory _orgRef) public view returns (bool);

  function getNumBudgets(string memory _orgRef) public view returns (uint256);
  function getBudgetReference(string memory _orgRef, uint256 _index) public view returns (string memory);
  function getTotalBudget(string memory _orgRef, string memory _budgetRef) public view returns (uint256);
  function getTotalBudgetStatus(string memory _orgRef, string memory _budgetRef) public view returns (uint256);
  function getTotalBudgetStart(string memory _orgRef, string memory _budgetRef) public view returns (uint256);
  function getTotalBudgetEnd(string memory _orgRef, string memory _budgetRef) public view returns (uint256);

  function getNumRecipientOrgs(string memory _orgRef) public view returns (uint256);
  function getRecipientOrgReference(string memory _orgRef, uint256 _index) public view returns (string memory);
  function getRecipientOrgBudget(string memory _orgRef, string memory _recipientOrgRef) public view returns (uint256);
  function getRecipientOrgBudgetStatus(string memory _orgRef, string memory _recipientOrgRef) public view returns (uint256);
  function getRecipientOrgBudgetStart(string memory _orgRef, string memory _recipientOrgRef) public view returns (uint256);
  function getRecipientOrgBudgetEnd(string memory _orgRef, string memory _recipientOrgRef) public view returns (uint256);

  function getNumRegions(string memory _orgRef) public view returns (uint256);
  function getRegionsReference(string memory _orgRef, uint256 _index) public view returns (string memory);
  function getRegionsBudget(string memory _orgRef, string memory _regionRef) public view returns (uint256);
  function getRegionsBudgetStatus(string memory _orgRef, string memory _regionRef) public view returns (uint256);
  function getRegionsBudgetStart(string memory _orgRef, string memory _regionRef) public view returns (uint256);
  function getRegionsBudgetEnd(string memory _orgRef, string memory _regionRef) public view returns (uint256);

  function getNumCountries(string memory _orgRef) public view returns (uint256);
  function getCountryReference(string memory _orgRef, uint256 _index) public view returns (string memory);
  function getCountryBudget(string memory _orgRef, string memory _countryRef) public view returns (uint256);
  function getCountryBudgetStatus(string memory _orgRef, string memory _countryRef) public view returns (uint256);
  function getCountryBudgetStart(string memory _orgRef, string memory _countryRef) public view returns (uint256);
  function getCountryBudgetEnd(string memory _orgRef, string memory _countryRef) public view returns (uint256);

  function getNumExpenditures(string memory _orgRef) public view returns (uint256);
  function getExpenditureReference(string memory _orgRef, uint256 _index) public view returns (string memory);
  function getTotalExpenditure(string memory _orgRef, string memory _expenditureRef) public view returns (uint256);
  function getTotalExpenditureStatus(string memory _orgRef, string memory _expenditureRef) public view returns (uint256);
  function getTotalExpenditureStart(string memory _orgRef, string memory _expenditureRef) public view returns (uint256);
  function getTotalExpenditureEnd(string memory _orgRef, string memory _expenditureRef) public view returns (uint256);
}
