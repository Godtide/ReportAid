pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

contract OrganisationReportCountryBudgets {

  struct Finance {
    uint256 value;
    bytes32 status;
    bytes32 start;
    bytes32 end;
  }

  struct CountryBudget {
    bytes32 reportRef;
    bytes32 budgetRef;
    bytes32 countryRef;
    bytes32 budgetLine;
    Finance finance;
  }

  function setCountryBudget(CountryBudget memory _budget) public;

  function getReportExists(bytes32 _reportRef) public view returns (bool);
  function getCountryBudgetExists(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bool);

  function getNumReports() public view returns (uint256);
  function getNumCountryBudgets(bytes32 _reportRef) public view returns (uint256);

  function getReportReference(uint256 _index) public view returns (bytes32);
  function getCountryBudgetReference(bytes32 _reportRef, uint256 _index) public view returns (bytes32);

  function getCountryBudget(bytes32 _reportRef, bytes32 _countryRef) public view returns (CountryBudget memory);
  function getCountryBudgetLine(bytes32 _reportRef, bytes32 _budgetRef) public view returns (bytes32);
  function getCountryBudgetValue(bytes32 _reportRef, bytes32 _countryRef) public view returns (uint256);
  function getCountryBudgetStatus(bytes32 _reportRef, bytes32 _countryRef) public view returns (bytes32);
  function getCountryBudgetStart(bytes32 _reportRef, bytes32 _countryRef) public view returns (bytes32);
  function getCountryBudgetEnd(bytes32 _reportRef, bytes32 _countryRef) public view returns (bytes32);
}
