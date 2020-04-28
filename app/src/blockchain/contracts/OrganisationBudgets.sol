pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Budgets.sol";

abstract contract OrganisationBudgets {

  function setBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef, Budgets.Budget memory _budget) virtual public;

  function getNumBudgets(bytes32 _organisationsRef, bytes32 _organisationRef) virtual public view returns (uint256);
  function getBudgetReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) virtual public view returns (bytes32);

  function getBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (Budgets.Budget memory);
  function getBudgetLine(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (bytes32);
  function getBudgetValue(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (uint256);
  function getBudgetStatus(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (uint8);
  function getBudgetStart(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (bytes32);
  function getBudgetEnd(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (bytes32);
}
