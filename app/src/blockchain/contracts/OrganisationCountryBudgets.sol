pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Budgets.sol";

abstract contract OrganisationCountryBudgets {

  function setCountryBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef, Budgets.Budget memory _budget) virtual public;

  function getNumCountryBudgets(bytes32 _organisationsRef, bytes32 _organisationRef) virtual public view returns (uint256);
  function getCountryBudgetReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) virtual public view returns (bytes32);

  function getCountryBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (Budgets.Budget  memory);
  function getCountryBudgetLine(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (bytes32);
  function getCountryBudgetCountry(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (bytes32);
  function getCountryBudgetValue(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (uint256);
  function getCountryBudgetStatus(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (uint8);
  function getCountryBudgetStart(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (bytes32);
  function getCountryBudgetEnd(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (bytes32);
}
