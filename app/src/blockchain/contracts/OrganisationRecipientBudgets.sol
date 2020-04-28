pragma solidity >=0.4.16 <0.7.0;
pragma experimental ABIEncoderV2;

import "./Budgets.sol";

abstract contract OrganisationRecipientBudgets {

  function setRecipientBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef, Budgets.Budget memory _budget) virtual public;

  function getNumRecipientBudgets(bytes32 _organisationsRef, bytes32 _organisationRef) virtual public view returns (uint256);
  function getRecipientBudgetReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) virtual public view returns (bytes32);

  function getRecipientBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (Budgets.Budget memory);
  function getRecipientBudgetLine(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (bytes32);
  function getRecipientBudgetOrg(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (bytes32);
  function getRecipientBudgetValue(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (uint256);
  function getRecipientBudgetStatus(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (uint8);
  function getRecipientBudgetStart(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (bytes32);
  function getRecipientBudgetEnd(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) virtual public view returns (bytes32);
}
