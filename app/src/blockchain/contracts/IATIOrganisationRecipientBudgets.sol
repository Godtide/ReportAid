pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

// IATI Organisation Reports
// Steve Huckle

import "./OrganisationRecipientBudgets.sol";
import "./IATIBudgets.sol";

contract IATIOrganisationRecipientBudgets is OrganisationRecipientBudgets {

  IATIBudgets budgets;

  constructor(address _budgets) public {
    require(_budgets != address(0x0));
    budgets = IATIBudgets(_budgets);
  }

  function setRecipientBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef, Budgets.Budget memory _budget) public {
    budgets.setBudget(_organisationsRef, _organisationRef, _budgetRef, _budget);
  }

  function getNumRecipientBudgets(bytes32 _organisationsRef, bytes32 _organisationRef) public view returns (uint256) {
    return budgets.getNumBudgets(_organisationsRef, _organisationRef);
  }

  function getRecipientBudgetReference(bytes32 _organisationsRef, bytes32 _organisationRef, uint256 _index) public view returns (bytes32) {
    return budgets.getBudgetReference(_organisationsRef, _organisationRef, _index);
  }

  function getRecipientBudget(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (Budgets.Budget memory) {
    return budgets.getBudget(_organisationsRef, _organisationRef, _budgetRef);
  }

  function getRecipientBudgetLine(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32) {
    return budgets.getBudgetLine(_organisationsRef, _organisationRef, _budgetRef);
  }

  function getRecipientBudgetOrg(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32) {
    return budgets.getOtherRef(_organisationsRef, _organisationRef, _budgetRef);
  }

  function getRecipientBudgetValue(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (uint256) {
    return budgets.getBudgetValue(_organisationsRef, _organisationRef, _budgetRef);
  }

  function getRecipientBudgetStatus(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (uint8) {
    return budgets.getBudgetStatus(_organisationsRef, _organisationRef, _budgetRef);
  }

  function getRecipientBudgetStart(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32) {
    return budgets.getBudgetStart(_organisationsRef, _organisationRef, _budgetRef);
  }

  function getRecipientBudgetEnd(bytes32 _organisationsRef, bytes32 _organisationRef, bytes32 _budgetRef) public view returns (bytes32) {
    return budgets.getBudgetEnd(_organisationsRef, _organisationRef, _budgetRef);
  }
}
