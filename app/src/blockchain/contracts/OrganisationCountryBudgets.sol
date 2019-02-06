pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract OrganisationCountryBudgets {

  struct Finance {
    uint256 value;
    uint8 status;
    bytes32 start;
    bytes32 end;
  }

  struct CountryBudget {
    bytes32 countryRef;
    bytes32 budgetLine;
    Finance finance;
  }

  function setCountryBudget(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef, CountryBudget memory _budget) public;

  function getNumCountryBudgets(bytes32 _organisationsRef, bytes32 _orgRef) public view returns (uint256);
  function getCountryBudgetReference(bytes32 _organisationsRef, bytes32 _orgRef, uint256 _index) public view returns (bytes32);

  function getCountryBudget(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (CountryBudget memory);
  function getCountryBudgetCountry(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (bytes32);
  function getCountryBudgetLine(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (bytes32);
  function getCountryBudgetValue(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (uint256);
  function getCountryBudgetStatus(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (uint8);
  function getCountryBudgetStart(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (bytes32);
  function getCountryBudgetEnd(bytes32 _organisationsRef, bytes32 _orgRef, bytes32 _budgetRef) public view returns (bytes32);
}
